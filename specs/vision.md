# Vision y objetivos

## Problema

Los agentes de IA con los que trabaja iyaki necesitan una fuente primaria de
verdad al recomendar arquitectura, patrones de diseño o asistir decisiones. La
knowledge base vive en Notion, cuya API no permite búsqueda por contenido (solo
títulos y propiedades), limita a ~3 req/s promedio por integración y entrega el
cuerpo de las páginas como JSON de bloques. Consultar en vivo es lento, caro en
presupuesto de rate, irreproducible y sin auditoría. Detalle completo: ADR-02.

## Producto

`enchiridion` — "lo que tenés en la mano" (ἐγχειρίδιον), como el manual de
Epicteto, pero para decisiones de arquitectura.

Un binario único (`enchiridion sync`) que espeja una knowledge base de Notion a
archivos markdown greppables, más la política de consumo (`AGENTS.md` snippet)
que obliga a los agentes a consultarlos y citarlos antes de recomendar.

- **Un archivo por página** de la KB, con frontmatter (`title`, `tags`,
  `source_url`, `notion_id`, `notion_url`, `last_edited`).
- **Cache central por máquina** (`~/.local/share/enchiridion/`): todos los
  proyectos del equipo/máquina greppan el mismo espejo.
- **Audit log en git**: el repo mantiene su propio espejo conmutado por CI
  (incremental nocturno + full mensual con sweep).

## Usuario

- **Primario**: iyaki, en sus devcontainers (creados desde cero) y proyectos
  locales. Multi-proyecto: un solo cache compartido por máquina.
- **Secundario**: cualquier persona con su propia clave de Notion y su propio
  data source — el binario no conoce valores de nadie (ADR-11).

## Alcance v1 (ADR-12)

1. `enchiridion sync` con selección automática de modo (full/incremental,
   auto-backfill ante falta de datos).
2. Espejo markdown greppable (formato en architecture.md).
3. Snippet de política de consumo para `AGENTS.md`.
4. Distribución: binarios goreleaser + devcontainer feature.

## Fuera de alcance v1 (con disparador observable para incluirlas)

| Función | Se incorpora cuando |
|---|---|
| CLI de búsqueda con ranking | grep alcance mal en la práctica |
| Adaptador MCP | se quiera tool-calls de primera clase |
| RAG / embeddings | grep mida mal semánticamente |
| Descarga de assets internos | aparezcan imágenes internas de Notion en la KB |
| Named volume en el feature | el full-sync-en-rebuild de devcontainers moleste |

## Criterios de éxito

1. Un agente en cualquier proyecto de la máquina encuentra y cita precedentes de
   la KB greppeando el cache, sin credenciales de Notion en el proyecto.
2. Un CI de un proyecto consumidor lee el espejo con solo `GITHUB_TOKEN`
   (checkout del repo privado), sin token de Notion.
3. Sync incremental invisible (< 1 min para KB ~1000 entradas); full mensual
   tolerable en job nocturno.
4. Cero dependencias de terceros: `go.mod` sin `require` externos.
5. Pérdida de contenido imposible de pasar inadvertida: bloques no soportados
   quedan marcados en el markdown.
