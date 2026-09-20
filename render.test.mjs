import { test } from 'node:test'
import assert from 'node:assert/strict'
import fs from 'node:fs'
import path from 'node:path'
import { slugify, extractMeta, frontmatter, blocksToMarkdown, sweepStaleFiles } from './render.js'

test('slugify strips accents and kebab-cases', () => {
	assert.equal(slugify('El desafío del lenguaje ubicuo'), 'el-desafio-del-lenguaje-ubicuo')
	assert.equal(slugify('  !!! '), 'untitled')
})

test('extractMeta reads title, tags, url and ids from a page object', () => {
	const page = {
		id: 'a1b2c3d4-e5f6-7890-abcd-ef0123456789',
		url: 'https://notion.so/a1b2c3d4',
		last_edited_time: '2026-09-20T10:00:00.000Z',
		properties: {
			Name: { type: 'title', title: [{ plain_text: 'Design Patterns' }] },
			Tema: { type: 'multi_select', multi_select: [{ name: 'arquitectura' }, { name: 'ddd' }] },
			Tipo: { type: 'select', select: { name: 'articulo' } },
			Link: { type: 'url', url: 'https://example.com/post' },
			Otro: { type: 'number', number: 3 },
		},
	}
	const meta = extractMeta(page)
	assert.equal(meta.title, 'Design Patterns')
	assert.deepEqual(meta.tags, ['arquitectura', 'ddd', 'articulo'])
	assert.equal(meta.url, 'https://example.com/post')
	assert.equal(meta.notionId, 'a1b2c3d4-e5f6-7890-abcd-ef0123456789')
})

test('frontmatter escapes quotes and lists tags', () => {
	const md = frontmatter({ title: 'Say "hi"', tags: ['a', 'b"c'], url: null, notionId: 'x', notionUrl: 'u', lastEdited: 't' })
	assert.match(md, /title: "Say \\"hi\\""/)
	assert.match(md, /tags: \["a", "b\\"c"\]/)
	assert.doesNotMatch(md, /source_url/)
})

test('sweepStaleFiles removes only mirrors whose page is gone', () => {
	const dir = fs.mkdtempSync('enchiridion-sweep-')
	const write = (name, id) => fs.writeFileSync(path.join(dir, name), `---\nnotion_id: ${id}\n---\nbody\n`)
	write('kept--aaa.md', 'aaa')
	write('stale--bbb.md', 'bbb')
	fs.writeFileSync(path.join(dir, 'notes.txt'), 'random')
	const removed = sweepStaleFiles(new Set(['aaa']), dir)
	assert.deepEqual(removed, ['stale--bbb.md'])
	assert.equal(fs.readdirSync(dir).sort().join(','), 'kept--aaa.md,notes.txt')
	fs.rmSync(dir, { recursive: true })
})

test('blocksToMarkdown renders the common text blocks', () => {
	const blocks = [
		{ type: 'paragraph', paragraph: { rich_text: [{ plain_text: 'Hello ', annotations: { code: false, bold: false, italic: false, strikethrough: false }, href: null }, { plain_text: 'world', annotations: { code: false, bold: true, italic: false, strikethrough: false }, href: null }] } },
		{ type: 'heading_2', heading_2: { rich_text: [{ plain_text: 'Section', annotations: { code: false, bold: false, italic: false, strikethrough: false }, href: null }] } },
		{ type: 'bulleted_list_item', bulleted_list_item: { rich_text: [{ plain_text: 'one', annotations: { code: false, bold: false, italic: false, strikethrough: false }, href: null }] } },
		{ type: 'numbered_list_item', numbered_list_item: { rich_text: [{ plain_text: 'first', annotations: { code: false, bold: false, italic: false, strikethrough: false }, href: null }] } },
		{ type: 'numbered_list_item', numbered_list_item: { rich_text: [{ plain_text: 'second', annotations: { code: false, bold: false, italic: false, strikethrough: false }, href: null }] } },
		{ type: 'divider', divider: {} },
		{ type: 'code', code: { rich_text: [{ plain_text: 'const x = 1' }], language: 'js' } },
		{ type: 'table', table: {} },
	]
	const md = blocksToMarkdown(blocks)
	assert.match(md, /Hello \*\*world\*\*/)
	assert.match(md, /## Section/)
	assert.match(md, /^- one$/m)
	assert.match(md, /1\. first/)
	assert.match(md, /2\. second/)
	assert.match(md, /```js\nconst x = 1\n```/)
	assert.match(md, /<!-- unsupported block: table -->/)
})
