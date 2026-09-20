# enchiridion

> **⚠ DESCARTADO**: este scaffold JS implementa decisiones revocadas (ADR-08).
> El stack definitivo es Go stdlib-only. Ver `ADR.md` — el estado de verdad de las
> decisiones es ese registro, no este README. Se elimina al iniciar la implementación.

*The handbook agents consult before deciding.* — Del griego ἐγχειρίδιον, "lo que tenés en la mano": manual compacto de consulta permanente, como el de Epicteto, pero para decisiones de arquitectura.

Espeja la [knowledge base de Notion](https://iyaki.notion.site/066daa9a7abb4c029724323209c85ca6) como archivos markdown grepeables. Existe porque la API de Notion **no busca en el cuerpo de las páginas** (solo títulos y propiedades) y cada consulta en vivo gasta el presupuesto de ~3 req/s de la integración; el sync lo paga una vez por día y deja el conocimiento versionado, auditable y consultable offline por cualquier agente con shell.

## Uso

```sh
npm install
cp .env.example .env   # completar NOTION_TOKEN
npm run sync           # Notion → knowledge/*.md
npm test
```

El workflow `.github/workflows/sync.yml` corre el sync diariamente y commitea los cambios.

## Cómo lo consume un agente

En el `AGENTS.md` de cualquier proyecto:

```markdown
## enchiridion
Antes de recomendar arquitectura o patrones de diseño, buscá en
`<ruta>/enchiridion/knowledge/` y citá las entradas que usaste.
Si no hay precedentes, decilo explícitamente.
```

El agente greppea `knowledge/**/*.md`; el frontmatter (`tags`, `notion_url`, `source_url`) permite refinar por tema y volver a la fuente.

## Convenciones

- JavaScript plano ESM, tabs, `node:test` — igual que `content-curator/organizer`.
- `index.js` orquesta (red + fs); `render.js` es puro y se testea offline.
- El renderer soporta los bloques de texto comunes; los exóticos quedan marcados como `<!-- unsupported block: X -->`.
- ponytail: sync completo en cada corrida; incremental si el job llegara a doler.

## Roadmap (solo si algo medible lo pide)

- CLI de búsqueda con ranking (`enchiridion query`) cuando grep alcance mal.
- Adaptador MCP como wrapper del mismo núcleo.
