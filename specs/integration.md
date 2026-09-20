# Integración

Contratos con sistemas externos y condiciones de borde. El registro de decisiones
vive en `ADR.md`.

## Integración con Notion

**Qué consume**: el contenido completo (propiedades + cuerpo) de todas las
páginas de un data source de knowledge base, más — solo cuando existen — los
bloques de tabla anidados.

**Restricciones externas que condicionan el diseño** (detalle en ADR-02):

- ~3 requests/segundo promedio por integración, presupuesto compartido por token.
- La búsqueda de la API no cubre el cuerpo de las páginas (solo títulos y
  propiedades): por eso existe el espejo.
- El cuerpo llega paginado y en un formato propio de bloques que debe
  transformarse a markdown.
- Misma superficie de API que `content-curator/organizer` usa en producción
  (versión `2025-09-03`), para no abrir un segundo frente de compatibilidad.

**Presupuesto**: una sync full de N páginas cuesta ~N+1+T requests (T = tablas).
Para N ≈ 1000 son minutos — invisible en job nocturno, intolerable dentro de la
sesión de un agente: otra razón del modelo sync (ADR-02). El incremental solo
toca páginas editadas.

## Mapeo de propiedades (schema real, confirmado en producción)

Fuente de verdad: `content-curator/organizer/index.js` y
`content-curator/curator/cms.js`.

| Propiedad Notion | Tipo | Destino |
|---|---|---|
| `Name` | título | `title` |
| `URL` | url | `source_url` (opcional; ya viene normalizada sin parámetros de tracking por el organizer) |
| `Category` | multi_select | `tags` (options conocidas: Tool, Service, Website, Note, Framework/Library, Game, …) |
| otras categorías/temas | select / multi_select dinámicas (creadas por el curador automático) | `tags` |
| — (metadata de página) | — | `notion_id`, `notion_url`, `last_edited` |

Las clasificaciones dinámicas nuevas que cree el curador automático entran sin
cambios: el schema no se hardcodea.

## Configuración (contrato con el usuario)

Cada usuario provee la suya — el binario no conoce valores de nadie (ADR-11).

| Variable | Obligatoria | Significado |
|---|---|---|
| `NOTION_TOKEN` | sí | integración de Notion del usuario, con acceso al data source |
| `KNOWLEDGE_BASE_DATASOURCE_ID` | sí | data source a espejar |
| `ENCHIRIDION_HOME` | no (default `~/.local/share/enchiridion`) | raíz del cache local |

**Least privilege**: el sync solo lee — la integración de Notion debe crearse con
capacidades de solo lectura (Read content, Read user info sin email). Espejo el
contenido de la KB tal cual es; si el token puede escribir, cualquier bug del
sync tendría alcance destructivo sin necesidad.

**Secretos en el espejo**: la KB es texto libre — puede llegar a contener un
token. Antes de conmutar `data/` en los workflows de sync, correr el scan de
secretos sobre el espejo (gitleaks); un espejo con secreto no se conmuta.

## Distribución (ADR-09)

1. **Releases**: binarios autocontenidos por plataforma (linux amd64/arm64,
   macos arm64) con checksums, adjuntos a cada release del repo.
2. **Devcontainer feature** (repo `devcontainer-features`): instala el binario
   desde la release usando el token de GitHub del usuario; expone las variables
   de configuración al container; permite fijar versión y raíz del cache.
   Opciones: `version` (default: última), `enchiridion_home`.

## Consumo

El trigger de consumo tiene dos mecanismos complementarios (ADR-13):

1. **Skill global** (caso primario): la skill vive en este repo
   (`.agents/skills/enchiridion/SKILL.md`) y se instala una sola vez con
   `npx skills add` (o symlink para omp). Dispara en todas las sesiones, en
   cualquier directorio.
2. **Snippet por-proyecto**: para proyectos donde se quiera la regla explícita
   en el `AGENTS.md`:

```markdown
## enchiridion

Fuente primaria de verdad: espejo de la knowledge base en
`~/.local/share/enchiridion/knowledge/` (override: `$ENCHIRIDION_HOME`).

Consultá el espejo (rg/grep) ANTES de responder cuando la tarea implique:
- elegir o recomendar una librería, framework o herramienta
- definir la estructura o el diseño de un módulo o servicio
- recomendar patrones de diseño o arquitectura
- resolver una disputa técnica entre alternativas
- citar cómo se resolvió algo antes

Cité las entradas usadas (archivo + `notion_url`). Sin precedentes, decilo
explícitamente. Cache inexistente: informalo — nunca inventes precedentes.
Actualizar: `enchiridion sync`.
```

Sin credenciales de Notion en el proyecto: solo archivos locales.

### CI de un proyecto consumidor

Checkout del repo privado con `GITHUB_TOKEN` y lectura del espejo conmutado en
`data/`. Sin token de Notion.

### CI de enchiridion (este repo)

| Workflow | Schedule | Comportamiento |
|---|---|---|
| Incremental | nocturno | sincroniza y conmuta `data/` si hubo cambios |
| Full | mensual (día 1) | sincroniza todo, aplica sweep, conmuta `data/` |
| Release | tag `v*` | publica binarios por plataforma |

Corridas con fallos parciales terminan en error visible — nunca se conmuta un
espejo parcial sin señal (patrón organizer).

## Límites conocidos (registrados, aceptados)

- Rebuild de devcontainer → cache vacío → sync full en el arranque (minutos).
  Mejora futura: volumen persistente para el cache (ADR-10).
- Imágenes internas de Notion quedan como marcador (ADR-05); el ingest actual ya
  las descarta al copiar triage → KB, por lo que se espera ~cero presencia.
- Los borrados en Notion tardan hasta el full mensual en reflejarse (ADR-04).
