# specs/ — Especificaciones de enchiridion

Mapa de la documentación. El registro de *por qué* de cada decisión vive en
[`ADR.md`](../ADR.md); estos documentos definen el *qué exacto a construir*.

| Documento | Contenido |
|---|---|
| [vision.md](vision.md) | Problema, producto, usuario, alcance v1 y fuera de alcance |
| [architecture.md](architecture.md) | Componentes, modos de sync, formato del espejo, contrato del renderer, errores, testing |
| [integration.md](integration.md) | API de Notion (endpoints, límites, mapeo de propiedades), distribución, consumo local y CI |
| [implementation-plan.md](implementation-plan.md) | Fases de implementación con criterios de aceptación |

## Estado

- Planificación: **cerrada** (2026-09-20, ADR-01..12).
- Implementación: **no iniciada**. El scaffold JS del commit `378a521` está
  **descartado** (ADR-08); se elimina en la fase 0 del plan.
- Decisiones de stack: ver ADR-08.

> Regla de documentación: las specs describen **qué, para qué y por qué** —
> nunca el cómo. El detalle de implementación vive en el código; las decisiones
> técnicas, en los ADR.
