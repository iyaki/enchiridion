import { Client } from '@notionhq/client'
import { extractMeta, frontmatter, slugify, blocksToMarkdown, sweepStaleFiles } from './render.js'
import fs from 'node:fs'
import path from 'node:path'
import 'dotenv/config'

const KNOWLEDGE_BASE_DATASOURCE_ID = process.env.KNOWLEDGE_BASE_DATASOURCE_ID
const OUTPUT_DIR = path.resolve(import.meta.dirname, 'knowledge')

if (!process.env.NOTION_TOKEN || !KNOWLEDGE_BASE_DATASOURCE_ID) {
	console.error('NOTION_TOKEN and KNOWLEDGE_BASE_DATASOURCE_ID must be defined in environment variables.')
	process.exit(1)
}

const notion = new Client({
	auth: process.env.NOTION_TOKEN,
	notionVersion: '2025-09-03',
})

// ponytail: full re-sync on every run; add last_edited_after incremental filter if the job gets slow
export async function fetchPages() {
	const pages = []
	let cursor
	do {
		const response = await notion.dataSources.query({
			data_source_id: KNOWLEDGE_BASE_DATASOURCE_ID,
			page_size: 100,
			start_cursor: cursor,
		})
		pages.push(...response.results)
		cursor = response.has_more ? response.next_cursor : undefined
	} while (cursor)
	return pages
}

export async function fetchBlocks(pageId) {
	const blocks = []
	let cursor
	do {
		const response = await notion.blocks.children.list({
			block_id: pageId,
			page_size: 100,
			start_cursor: cursor,
		})
		blocks.push(...response.results)
		cursor = response.has_more ? response.next_cursor : undefined
	} while (cursor)
	return blocks
}

export function writePage(page, blocks) {
	const meta = extractMeta(page)
	const body = blocksToMarkdown(blocks)
	const file = path.join(OUTPUT_DIR, `${slugify(meta.title)}--${meta.notionId.slice(0, 8)}.md`)
	fs.writeFileSync(file, `${frontmatter(meta)}\n${body}\n`)
	return { file, notionId: meta.notionId }
}

async function sync() {
	fs.mkdirSync(OUTPUT_DIR, { recursive: true })
	const started = Date.now()

	const pages = await fetchPages()
	console.log(`Fetched ${pages.length} pages from the knowledge base.`)

	const keptIds = new Set()
	for (const page of pages) {
		try {
			const blocks = await fetchBlocks(page.id)
			const { file } = writePage(page, blocks)
			keptIds.add(page.id)
			console.log(`Synced: ${path.basename(file)}`)
		} catch (error) {
			console.error(`Failed to sync page ${page.id}:`, error.message)
		}
	}

	const removed = sweepStaleFiles(keptIds)
	if (removed.length) console.log(`Removed ${removed.length} stale files.`)

	console.log(`Sync complete: ${keptIds.size} pages kept, ${removed.length} removed, in ${((Date.now() - started) / 1000).toFixed(1)}s.`)
	if (keptIds.size < pages.length) process.exitCode = 1
}

await sync()
