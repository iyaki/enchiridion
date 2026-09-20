// Pure helpers: no Notion client, no process env — all offline-testable
import fs from 'node:fs'
import path from 'node:path'

export function slugify(title) {
	return title
		.normalize('NFD')
		.replace(/[\u0300-\u036f]/g, '')
		.toLowerCase()
		.replace(/[^a-z0-9]+/g, '-')
		.replace(/^-+|-+$/g, '') || 'untitled'
}

export function extractMeta(page) {
	let title = 'Untitled'
	const tags = []
	let url = null

	for (const prop of Object.values(page.properties)) {
		switch (prop.type) {
			case 'title':
				if (prop.title.length) title = prop.title.map(r => r.plain_text).join('')
				break
			case 'select':
				if (prop.select) tags.push(prop.select.name)
				break
			case 'multi_select':
				tags.push(...prop.multi_select.map(o => o.name))
				break
			case 'status':
				if (prop.status) tags.push(prop.status.name)
				break
			case 'url':
				if (prop.url) url = prop.url
				break
		}
	}

	return {
		title,
		tags: [...new Set(tags)],
		url,
		notionId: page.id,
		notionUrl: page.url,
		lastEdited: page.last_edited_time,
	}
}

export function frontmatter(meta) {
	const lines = [
		'---',
		`title: "${meta.title.replaceAll('"', '\\"')}"`,
		`notion_id: ${meta.notionId}`,
		`notion_url: ${meta.notionUrl}`,
		`last_edited: ${meta.lastEdited}`,
	]
	if (meta.url) lines.push(`source_url: ${meta.url}`)
	if (meta.tags.length) lines.push(`tags: [${meta.tags.map(t => `"${t.replaceAll('"', '\\"')}"`).join(', ')}]`)
	lines.push('---')
	return lines.join('\n')
}

function richText(items) {
	return items
		.map(r => {
			let text = r.plain_text
			if (!text) return text
			if (r.annotations.code) text = `\`${text}\``
			if (r.annotations.bold) text = `**${text}**`
			if (r.annotations.italic) text = `_${text}_`
			if (r.annotations.strikethrough) text = `~~${text}~~`
			if (r.href) text = `[${text}](${r.href})`
			return text
		})
		.join('')
}

// ponytail: flat renderer, no recursive child fetching (toggles/sub-pages render as markers);
// swap in notion-to-md if exotic blocks start mattering
export function blocksToMarkdown(blocks) {
	const out = []
	let numbered = 0
	for (const block of blocks) {
		const type = block.type
		const data = block[type]
		switch (type) {
			case 'paragraph':
				numbered = 0
				out.push(richText(data.rich_text))
				break
			case 'heading_1':
			case 'heading_2':
			case 'heading_3':
				numbered = 0
				out.push(`${'#'.repeat(Number(type.slice(-1)))} ${richText(data.rich_text)}`)
				break
			case 'bulleted_list_item':
				numbered = 0
				out.push(`- ${richText(data.rich_text)}`)
				break
			case 'numbered_list_item':
				out.push(`${++numbered}. ${richText(data.rich_text)}`)
				break
			case 'quote':
				numbered = 0
				out.push(`> ${richText(data.rich_text)}`)
				break
			case 'code':
				numbered = 0
				out.push('```' + (data.language || '') + '\n' + data.rich_text.map(r => r.plain_text).join('') + '\n```')
				break
			case 'divider':
				numbered = 0
				out.push('---')
				break
			case 'callout':
				numbered = 0
				out.push(`> ${richText(data.rich_text)}`)
				break
			case 'bookmark':
			case 'embed':
			case 'link_preview':
				numbered = 0
				out.push(`[${data.url}](${data.url})`)
				break
			case 'image':
				numbered = 0
				out.push(`![image](${data.file?.url || data.external?.url || ''})`)
				break
			case 'toggle':
				numbered = 0
				out.push(`**${richText(data.rich_text)}**`)
				break
			case 'child_page':
				out.push(`<!-- child page: ${data.title} -->`)
				break
			default:
				out.push(`<!-- unsupported block: ${type} -->`)
		}
	}
	return out.join('\n\n')
}

// Removes mirrors of pages deleted (or filtered out) in Notion, so the mirror never lies
export function sweepStaleFiles(keptIds, dir) {
	const removed = []
	for (const file of fs.readdirSync(dir)) {
		if (!file.endsWith('.md')) continue
		const content = fs.readFileSync(path.join(dir, file), 'utf8')
		const match = content.match(/^notion_id: (\S+)$/m)
		if (!match || !keptIds.has(match[1])) {
			fs.rmSync(path.join(dir, file))
			removed.push(file)
		}
	}
	return removed
}
