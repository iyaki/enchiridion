# Plan de implementación

Estado de referencia: decisiones en `ADR.md` (cerradas), especificación en
`specs/`. Regla transversal: **cero dependencias de terceros** (`go.mod` sin
`require` externos — verificar por revisión del archivo en cada fase).

## Fase 0 — Limpieza

- Borrar el scaffold JS descartado (`index.js`, `render.js`, `render.test.mjs`,
  `knowledge/.gitkeep` se re-crea vacío, `.env.example`, workflow y README
  obsoletos — commits `378a521`, `a9051e2`).
- `go mod init github.com/iyaki/enchiridion` (Node 24 `.tool-versions` fuera;
  `.tool-versions` pasa a `golang 1.2x` — fijar la versión del toolchain local).

**AC**: repo con solo `ADR.md`, `specs/`, `go.mod`, `.gitignore` (`.env`,
`data/` en local dev no... `data/` se commitea: solo `.env`).

## Fase 1 — Cliente Notion (`internal/notion`)

- Structs tipados para: page (properties genéricas), block, query
  request/response, paginación, filtro `last_edited_time`.
- Cliente: POST query, GET blocks children, retry/backoff (429 con
  `Retry-After`, 5xx, máx 3), fail-fast 401/403.
- Tests: `httptest` con fixtures JSON (capturar shapes reales de la API del
  organizer si es posible; si no, del schema documentado).

**AC**: `go test ./internal/notion` verde offline; requests construidos
verificados por assertions en el handler de test.

## Fase 2 — Renderer (`internal/render`)

- `blocksToMarkdown` según la tabla de contrato (architecture.md), incluida
  `table`/`table_row` con `has_column_header`, pipes escapados, `<br>` en
  newlines.
- Inline: code/bold/italic/strikethrough/href en ese orden.
- Golden tests: fixture de bloques → markdown esperado, caso por fila de la
  tabla + caso mixto + marcadores de no-soportados.

**AC**: golden tests verdes; ningún caso de la tabla sin test.

## Fase 3 — Motor de sync (`internal/sync`)

- Selección de modo (sin state / vacío → full; >30d → full; else incremental;
  `--full` fuerza) — tabla de casos testeada.
- Full: paginado completo, escritura `{slug}--{id8}.md`, sweep, reset watermark.
- Incremental: filtro con margen 60s, rename-safe write por `notion_id`,
  avance de watermark.
- Frontmatter con escaping de `"`; slug sin acentos (NFD strip).
- Manejo de errores: por-página continuar + exit 1 final (patrón organizer).

**AC**: tests con `t.TempDir()` cubren: backfill automático, sweep que preserva
vigentes y borra stale, rename que no duplica, watermark que avanza y se
conserva commiteada, fallos parciales → exit 1.

## Fase 4 — CLI (`cmd/enchiridion`)

- `enchiridion sync [--full]`; `ENCHIRIDION_HOME`, `NOTION_TOKEN`,
  `KNOWLEDGE_BASE_DATASOURCE_ID` (faltantes → mensaje claro, exit 1, sin
  llamar API).
- Logging a stdout por página syncada + resumen final (`N pages kept, M
  removed, in Xs`); errores a stderr.
- **Smoke real manual**: primera corrida contra la API con token propio.

**AC**: la corrida real produce el espejo esperado sobre una KB de prueba;
exit codes correctos.

## Fase 5 — CI y release

- `sync-incremental.yml` (nocturno) y `sync-full.yml` (mensual) con
  `ENCHIRIDION_HOME=$GITHUB_WORKSPACE/data`, commit del espejo (patrón
  organizer: exit 1 si hubo fallos parciales).
- `release.yml` con goreleaser (linux amd64/arm64, darwin arm64, checksums).
- Secrets: `NOTION_TOKEN`, `KNOWLEDGE_BASE_DATASOURCE_ID`.

**AC**: corrida manual de cada workflow en GitHub Actions verde; binario de
release instalable y funcional.

## Fase 6 — Devcontainer feature (repo `devcontainer-features`)

- Feature `enchiridion`: options `version`, `enchiridion_home`; install.sh con
  `GITHUB_TOKEN` + download de release + checksum; sin escrever credenciales a
  disco.
- Docs de consumo: snippet `AGENTS.md` + ejemplo de `devcontainer.json`.

**AC**: feature instalado en un devcontainer de prueba desde cero →
`enchiridion sync` corre y puebla el cache.

## Definition of Done (global)

- `go test ./... && go vet ./...` verdes; `go.mod` sin dependencias externas.
- Los 5 criterios de éxito de `specs/vision.md` verificados (los 3 primeros con
  evidencia real, no mocks).
- ADR.md actualizado con cualquier desvío de estas specs.
