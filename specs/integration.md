# Integración

## API de Notion

Versión de API: `2025-09-03` (la misma que usa `content-curator/organizer` en
producción con `@notionhq/client` v5.26). Implementación propia sobre `net/http`
con structs tipados — cero dependencias (ADR-08).

### Endpoints usados

| Operación | Request |
|---|---|
| Query del data source | `POST /v1/data_sources/{data_source_id}/query` — body `{page_size, start_cursor, filter?}` |
| Bloques de una página | `GET /v1/blocks/{page_id}/children?page_size=100` |

Headers: `Authorization: Bearer {NOTION_TOKEN}`, `Notion-Version: 2025-09-03`,
`Content-Type: application/json`.

### Paginación

`page_size` máximo 100 en ambos. Iterar mientras `has_more` con `next_cursor`.

### Rate limit y reintentos

~3 req/s promedio por integración (presupuesto compartido por token). Ante 429:
respetar `Retry-After`, backoff exponencial, máximo 3 reintentos por request.
Presupuesto de una sync full para KB de N páginas: ~1 + ceil(N/100) + N + T
requests (T = tablas), ≈ 6-8 min para N=1000 — aceptable en job nocturno
(ADR-02). El incremental solo toca páginas editadas.

### Filtro incremental

```json
{ "timestamp": "last_edited_time",
  "last_edited_time": { "on_or_after": "<watermark - 60s ISO 8601>" } }
```

### Errores

401/403 → fail-fast con mensaje de credenciales (sin fallback silencioso).
`object_not_found` → la integración no tiene acceso al data source: mensaje que
guíe a conectarla en Notion. Resto 4xx → fail por página, log, continuar.

## Mapeo de propiedades (schema real, confirmado en producción)

Fuente de verdad: `content-curator/organizer/index.js` (`getDatabaseSchema`,
`processArticle`) y `content-curator/curator/cms.js`.

| Propiedad Notion | Tipo | Destino frontmatter |
|---|---|---|
| `Name` | `title` | `title` (join de `plain_text`) |
| `URL` | `url` | `source_url` (opcional; ya viene normalizada sin tracking params — `stripTrackingParams` del organizer) |
| `Category` | `multi_select` | `tags` (options conocidas: Tool, Service, Website, Note, Framework/Library, Game, …) |
| otras `select`/`multi_select`/`status` | dinámicas (creadas por el classifier IA) | `tags` |
| — | `page.id` | `notion_id` |
| — | `page.url` | `notion_url` |
| — | `page.last_edited_time` | `last_edited` |

Las selects dinámicas nuevas que cree el classifier entran solas: el schema no se
harcodea, se lee de los `properties` de cada página.

## Distribución (ADR-09)

1. **Releases**: goreleaser — binarios `enchiridion_{os}_{arch}` (linux/arm64 y
   amd64 mínimo; macos/arm64 para desarrollo), checksums. Patrón `web-archiver`.
2. **Devcontainer feature** (repo `devcontainer-features`):
   - Instala el binario de la release con `GITHUB_TOKEN` (repo privado).
   - Opciones: `version` (default `latest`), `enchiridion_home`.
   - El usuario provee `NOTION_TOKEN` y `KNOWLEDGE_BASE_DATASOURCE_ID` como env
     del container (`containerEnv`/`remoteEnv`/secrets del devcontainer).
   - `postCreateCommand` sugerido: `enchiridion sync` (auto-backfill si el cache
     del container está vacío — costo conocido: ADR-10).

## Consumo

### Agente local (caso primario)

Snippet para el `AGENTS.md` de cada proyecto consumidor:

```markdown
## enchiridion
Antes de recomendar arquitectura o patrones de diseño, buscá en
`~/.local/share/enchiridion/knowledge/` y citá las entradas que usaste.
Si no hay precedentes, decilo explícitamente.
```

Sin credenciales de Notion en el proyecto: solo archivos locales.

### CI de un proyecto consumidor

Checkout del repo privado con `GITHUB_TOKEN` y leer `data/knowledge/` — el
espejo conmutado por el CI de enchiridion. Sin token de Notion.

### CI de enchiridion (este repo)

| Workflow | Schedule | Modo | Efecto |
|---|---|---|---|
| `sync-incremental.yml` | nocturno | incremental | commit de `data/` si hay cambios |
| `sync-full.yml` | mensual (día 1, 09:00 UTC) | full + sweep | commit de `data/` |
| `release.yml` | tag `v*` | — | goreleaser |

Todos con `ENCHIRIDION_HOME=$GITHUB_WORKSPACE/data` y secrets del repo.
Permisos: `contents: write`. Commits firmados por bot (`enchiridion-bot`).

## Límites conocidos (registrados, aceptados)

- Rebuild de devcontainer → cache vacío → full sync en postCreate (minutos).
  Mejora futura: named volume (ADR-10).
- Imágenes internas de Notion quedan como marcador (ADR-05); el ingest actual ya
  las descarta al copiar triage → KB, por lo que se espera ~cero presencia.
- Los borrados en Notion tardan hasta el full mensual en reflejarse (ADR-04).
