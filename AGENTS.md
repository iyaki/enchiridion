# enchiridion — espejo de la knowledge base de Notion como markdown

Conocimiento del proyecto para agentes que trabajan en este repo.

## Qué es

`enchiridion` baja la knowledge base de Notion (data source `KNOWLEDGE_BASE_DATASOURCE_ID`)
y la escribe como `knowledge/*.md` — un archivo por página, con frontmatter
(`title`, `tags`, `notion_id`, `notion_url`, `source_url`, `last_edited`).

## Convenciones

- JavaScript plano ESM, Node 24, tabs — igual que `content-curator/organizer`.
- Tests con `node:test` + `assert/strict` en archivos `*.test.mjs`, sin red:
  todo lo testeable vive en `render.js` (funciones puras).
- `index.js` es la única pieza con I/O (cliente Notion + fs); mantenerlo así.
- El renderer de bloques es flat: no baja a bloques hijos anidados (toggles,
  sub-páginas). Si hiciera falta, evaluar `notion-to-md`, no crecer el renderer a mano.

## Verificación

- `npm test` — offline, rápido, siempre debe pasar.
- El sync real requiere `NOTION_TOKEN` (no hay en CI más que en el secret del workflow).

## Consumo por agentes (el propósito del proyecto)

En proyectos donde se quiera consultar el conocimiento, agregar al `AGENTS.md`:

```markdown
## enchiridion
Antes de recomendar arquitectura o patrones de diseño, buscá en
`<ruta>/enchiridion/knowledge/` y citá las entradas que usaste.
Si no hay precedentes, decilo explícitamente.
```
