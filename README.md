# enchiridion

*The handbook agents consult before deciding.* — Del griego ἐγχειρίδιον, "lo que
tenés en la mano": manual compacto de consulta permanente, como el de Epicteto,
pero para decisiones de arquitectura.

Espeja una knowledge base de Notion como archivos markdown greppables, para que
agentes de IA la usen como fuente primaria de verdad al recomendar arquitectura,
patrones de diseño o asistir decisiones.

## Estado

Planificación cerrada (`specs/`, `ADR.md`); implementación en curso — fase 0/1 de
[`specs/implementation-plan.md`](specs/implementation-plan.md). El comando `sync`
aún no está implementado.

## Desarrollo

Requisitos: Go 1.25. Tooling (lint, seguridad, mutación, arquitectura) se
instala con `.devcontainer/install-go-tools.sh`.

```sh
make quality          # gates completos: test, lint, race, flaky, coverage, mutation, security, arch
make test             # suite offline
make build            # binario en bin/enchiridion
make help             # todos los targets
```

Convenciones y reglas para agentes: [`AGENTS.md`](AGENTS.md). Decisiones:
[`ADR.md`](ADR.md). Especificaciones: [`specs/`](specs/README.md).

## Consumo (una vez implementado)

Ver `specs/integration.md` — distribución vía releases + devcontainer feature,
cache central en `~/.local/share/enchiridion/`, y el snippet de política para
los `AGENTS.md` de los proyectos consumidores.
