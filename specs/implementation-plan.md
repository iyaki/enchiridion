# Plan de implementación

Secuencia de entregables. Cada fase define **qué debe ser verdad al terminar**
(criterios de aceptación), no cómo escribirlo. Decisiones de fondo: `ADR.md`.
Comportamiento esperado: `architecture.md` e `integration.md`.

Regla transversal: cero dependencias de terceros (ADR-08) y toda la lógica
testeable sin red (architecture.md, Requisitos no funcionales).

## Fase 0 — Limpieza

Eliminar el scaffold JS descartado (ADR-08) e inicializar el módulo del lenguaje
elegido.

**AC**: el repo contiene únicamente documentación vigente y la base del nuevo
proyecto; sin rastros del scaffold.

## Fase 1 — Acceso a Notion

Cliente del data source: query completo con paginación, lectura del cuerpo de
una página (incluidas tablas anidadas), reintentos según
`architecture.md` (Comportamiento ante errores).

**AC**: verificado offline contra respuestas reales de ejemplo; los errores
definidos producen los comportamientos definidos.

## Fase 2 — Renderer

Transformación de bloques a markdown según el contrato de
`architecture.md` (Contrato de renderizado).

**AC**: cada fila del contrato tiene un caso de prueba con su salida esperada
(exacta, predefinida), incluido el caso mixto y los marcadores de
no-soportados.

## Fase 3 — Motor de sync

Selección automática de modo, full con sweep, incremental con marca de agua,
actualización in-place por identidad de página, frontmatter y nombres de
archivo según el formato del espejo.

**AC**: tabla de casos de `architecture.md` (Modos de sincronización) cubierta
con pruebas sobre directorios temporales: backfill automático, sweep que
preserva vigentes y elimina stale, renames sin duplicados, watermark que
avanza, fallos parciales → salida de error.

## Fase 4 — CLI

`enchiridion sync` con flag para forzar full, variables de configuración según
`integration.md`, logging de progreso y resumen final.

**AC**: configuración faltante produce mensaje claro sin contacto con la API;
códigos de salida según spec; **smoke real manual** contra la API con token
propio produce el espejo esperado.

## Fase 5 — CI y releases del repo

Workflows según `integration.md` (CI de enchiridion): incremental nocturno,
full mensual con sweep, releases con binarios multiplataforma y checksums.
Los workflows de sync corren el scan de secretos sobre el espejo antes de
conmutar `data/` (specs/integration.md — Secretos en el espejo).

**AC**: corrida manual de cada workflow en verde; binario de release instalable
y funcional; `data/` conmutado solo por corridas exitosas.

## Fase 6 — Devcontainer feature (repo `devcontainer-features`)

Feature que instala el binario desde la release y expone la configuración
según `integration.md` (Distribución).

**AC**: un devcontainer desde cero instala el binario, corre el sync y puebla
el cache local.

## Definition of Done (global)

- Los 5 criterios de éxito de `specs/vision.md` verificados — con evidencia
  real, no simulada.
- Suite de pruebas offline completa en verde.
- `make deps-audit` en verde (ADR-08 reforzado mecánicamente, no solo
  documentado).
- `ADR.md` actualizado con cualquier desvío de estas specs.
