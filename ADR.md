# ADR — Registro de decisiones de enchiridion

Cada entrada: contexto → decisión → consecuencias. Las alternativas rechazadas se
anotan para no repetir las mismas discusiones. Última actualización: 2026-09-20.

## Estado del proyecto

Planificación cerrada, implementación iniciada (2026-09-20):
- **Fase 0 ejecutada**: scaffold JS eliminado; módulo Go inicializado
  (`github.com/iyaki/enchiridion`, Go 1.25, cero dependencias).
- **Harness instalado** (patrón reglint/specralph): Makefile de gates
  (`make quality`), lefthook pre-commit (format + coverage + mutation-diff +
  lint + security + arch), golangci-lint, go-arch-lint (sync→notion+render;
  notion y render puros), gremlins, govulncheck/gosec, goreleaser, workflows
  (quality/security/release/update-agent-skills), devcontainer, editorconfig,
  opencode.jsonc con protecciones del harness.

---

## ADR-01 — Propósito y nombre

**Contexto**: se busca que agentes de IA usen la knowledge base de Notion como
fuente primaria de verdad al recomendar arquitectura, patrones o asistir
decisiones. Serie de posts "Yo soy iyaki" como contexto del pipeline actual
(Feedly → triage → KB Notion → curated site).

**Decisión**: nombre **enchiridion** (ἐγχειρίδιον, "lo que tenés en la mano":
manual compacto de consulta permanente; el de Epicteto es el manual clásico de
buenas decisiones). Repositorio propio bajo `iyaki/`.

**Rechazados**: `second-opinion` (colisión semántica: el ecosistema MCP ya lo usa
para "consultar otros LLMs"), `lore` y `grimoire` (ocupados por herramientas de
función idéntica), `heavens-door` (guiño JoJo, descartado por el usuario),
`precedent` / `ground-truth` / `knowledge-judge` (finalistas, no elegidos).

## ADR-02 — Sync vs. query en vivo

**Contexto**: limitaciones de la API de Notion verificadas: ~3 req/s promedio por
integración; `search` solo matchea títulos; `dataSource/query` filtra solo
propiedades; el cuerpo se lee bloque a bloque (paginado, JSON que hay que
renderizar).

**Decisión**: **sync** (espejo programado a archivos locales). La razón
estructural: la API no sabe buscar en cuerpos — "¿qué anoté sobre X?" es
irrespondible en vivo si X no está en un título o tag. Además: el costo de
requests se paga una vez por corrida (no por consulta de agente), auditoría en
git, consultables offline, sin credenciales en los consumidores. `query` en vivo
queda como escape futuro, YAGNI.

## ADR-03 — Formato del espejo

**Decisión**: un `.md` por página en `knowledge/`, frontmatter:

```yaml
title:        # Name.title → plain_text
tags:         # Category (multi_select) + selects/multi_selects dinámicas del classifier
source_url:   # URL.url — ya normalizada (tracking strippado por organizer)
notion_id:    # page.id — identidad estable del espejo
notion_url:   # page.url
last_edited:  # page.last_edited_time
```

Propiedades confirmadas en código de producción (`organizer`, `curator/cms.js`):
`Name` (title), `URL` (url), `Category` (multi_select: Tool, Service, Website,
Note, Framework/Library, Game, ...) y selects dinámicas del clasificador IA.

**Rechazado**: índice único / JSON / SQLite (rompen files-as-API: obligan a
tooling para leer); `INDEX.md` en v1 (grep sobre frontmatter alcanza).

## ADR-04 — Espejo fiel con sweep

**Decisión**: si una página muere en Notion, muere en el espejo. El sweep corre
solo en full sync. Git es el archivo de historia: borrar hoy es recuperable con
`git log`.

**Razón**: un fantasma desactualizado envenena decisiones ("la KB dice X" cuando
X fue borrado). El usuario confirmó que rara vez borra → sweep mensual alcanza.

## ADR-05 — Renderer propio, minimalista

**Contexto**: la KB es texto curado (párrafos, headings, listas, quotes, código,
callouts, tablas — las tablas existen: `getAllPageBlocks` hace fetch recursivo de
children solo para tables).

**Decisión**: renderer propio (~100 líneas) sobre los tipos de bloque usados.
Tablas → markdown. Imágenes: external → link directo; internal → marcador visible
(las URLs internas de Notion expiran en ~1h: commitearlas es inútil, y el propio
`sanitizeBlocks` del organizer ya las descarta al copiar triage → KB).
**Bloques no soportados se marcan visibles** (`<!-- unsupported block: X -->`) —
en un artefacto de verdad, la pérdida silenciosa de contenido es el peor modo de
fallo.

**Rechazado**: `notion-to-md` (riesgo de desactualización frente al datamodel
`2025-09-03`; swap futuro si aparecen bloques exóticos).

## ADR-06 — Scope del espejo: todo el data source

**Decisión**: baja Tool, Service, Website, Game, etc. incluidos. Para decisiones
de arquitectura, las entradas Tool/Framework son de las más consultadas ("¿qué
usamos para X?"). El filtro por tipo es decisión de presentación de cada
consumidor (trivial vía `tags`). El filtro "solo artículos" del curated site NO
se replica: es de la capa de presentación.

## ADR-07 — Cadencia y modos de sync

**Decisión** (definida por el usuario):
- **Incremental** en cada startup/corrida local (filtro `last_edited_time
  on_or_after` + margen de solape por clock skew). Sin sweep.
- **Full** mensual en GH Action, con sweep.
- **Selección automática de modo**: sin state o `knowledge/` vacío → full
  (auto-backfill: la herramienta resuelve la falta de datos sola, sin paso
  manual); watermark >30 días → full; resto → incremental. Flag `--full` fuerza.
- **Watermark** commiteada junto al espejo (`.sync-state.json`) — un clone fresco
  incrementa correcto.
- **Renames**: reescritura por lookup de `notion_id` en frontmatter (el path
  conserva el slug viejo, el contenido el nuevo) — cero duplicados entre fulls.

## ADR-08 — Stack: Go stdlib-only

**Contexto**: criterio del usuario (era AI-first): tecnologías lo más verificables
y seguras posibles. Los proyectos JS previos (`content-curator`,
`knowledge-base-clasificator`) preceden a esa era y no son la referencia.

**Decisión**: **Go, cero dependencias de terceros**. Los 2 endpoints usados
(`dataSources/query`, `blocks/children`) van sobre `net/http` +
`encoding/json` con structs tipados: toda la cadena de red y parsing vive en el
repo, supply chain de tamaño 1, `go vet` / `govulncheck`, binario estático único.

**Rechazados**: Rust (máxima garantía, costo de desarrollo injustificado para un
sync), TypeScript estricto (tipos pero conserva runtime + árbol de dependencias),
JS puro (el criterio previo, superado). **Nota**: el scaffold JS del commit
`378a521` implementa decisiones ya revocadas; descartado.

## ADR-09 — Distribución: repo privado + binarios + devcontainer feature

**Decisión**:
- Repositorio **privado** (el usuario no busca publicarlo).
- Binarios por release vía goreleaser (patrón `web-archiver`).
- Feature propio en el repo `devcontainer-features` que instala el binario con
  `GITHUB_TOKEN` y expone `NOTION_TOKEN` / `KNOWLEDGE_BASE_DATASOURCE_ID` como env
  del container. Los proyectos consumidores no clonan este repo.

## ADR-10 — Cache central de conocimiento

**Contexto**: si los consumidores solo instalan el binario, ¿dónde greppan los
agentes?

**Decisión**: espejo central por máquina en `~/.local/share/enchiridion/knowledge/`
(XDG), override con `ENCHIRIDION_HOME`. Todos los proyectos de la máquina
comparten un solo espejo y watermark. El snippet de política en el `AGENTS.md` de
cada proyecto apunta a esa ruta absoluta. El repo mantiene además su propio
espejo commiteado (CI: incremental nocturno + full mensual) — audit log en git, y
los CI consumidores pueden leer el espejo con solo `GITHUB_TOKEN`, sin token de
Notion.

**Costo registrado**: rebuild de devcontainer = cache vacío → full sync en
postCreate (minutos). Correcto por el auto-backfill; mejora futura: named volume
para el data dir en el feature.

## ADR-11 — Config multi-usuario, cero hardcode

**Decisión**: cada usuario provee su propio `NOTION_TOKEN` y su
`KNOWLEDGE_BASE_DATASOURCE_ID` — el binario no conoce ni el token ni la KB de
nadie. Ningún valor específico de iyaki vive en el código (el ID del datasource
de su KB va solo en su config).

## ADR-12 — Alcance v1

**Decisión**: sync (auto full/incremental) + archivos greppables + política de
consumo (`AGENTS.md` snippet). El producto v1 es el espejo y la política que
obliga a consultarlo.

**Postergados con disparador observable**: CLI de búsqueda con ranking (cuando
grep alcance mal), adaptador MCP (wrapper del mismo núcleo, cuando se quiera
tool-calls de primera clase), RAG/embeddings (cuando grep mida mal), descarga de
assets internos (cuando aparezcan imágenes internas en la KB), named volume en el
feature (cuando el full-en-rebuild moleste).

## ADR-13 — Trigger de consumo: skill global + snippet por-proyecto

**Contexto**: sin un mecanismo de trigger, ningún agente usaría la herramienta —
la única referencia existente era la conversación misma. El consumo es
policy-based (archivos + instrucciones), no tool-calls.

**Decisión**:
1. **Skill global** (primaria): la skill vive como artefacto versionado del repo
   en `.agents/skills/enchiridion/SKILL.md`, con una description que enumera
   disparadores concretos (elegir librerías, diseñar módulos, recomendar
   patrones, resolver disputas técnicas, citar precedentes). Instalación única
   vía `npx skills add` (o symlink `.omp/skills` para omp) → dispara en todas
   las sesiones, en cualquier directorio.
2. **Snippet por-proyecto** (complemento): regla explícita en el `AGENTS.md` de
   proyectos específicos, con los mismos disparadores concretos — la abstracción
   "arquitectura" sola perdía casos como "¿qué ORM uso?".

La instalación de ambos es **paso post-implementación** (nada se instala hasta
que la fase 4 produzca el primer sync real).

**Rechazados**: presencia del tool en el toolset vía MCP como recordatorio
per-turn (over-engineering, ADR-12); hooks que inyecten resultados
automáticamente ante ciertos edits (matching difuso + costo por turno, sin
evidencia de necesidad).

---

## Verificación del sync real (pendiente de implementación)

La única pieza no verificable offline: las llamadas vivas a la API. Smoke manual
con token propio la primera vez + el workflow en CI.
