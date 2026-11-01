---
title: "Bomberos: ¿cómo nos organizamos frente a los incendios en Producción?"
notion_id: aa195c6a-a478-402a-a3d4-a9df654fccd0
notion_url: https://app.notion.com/p/Bomberos-c-mo-nos-organizamos-frente-a-los-incendios-en-Producci-n-aa195c6aa478402aa3d4a9df654fccd0
last_edited: 2023-02-16T14:24:00.000Z
source_url: https://medium.com/redbee/bomberos-ec27c05abd2e
tags: ["Article", "redbee - Medium", "Español", "Help Desk", "On Call", "Project Management", "Productivity"]
---
_Para cumplir con los compromisos del Sprint y resolver las tareas no planificadas creamos este equipo, listo para solucionar cualquier incidente que “entre por la ventana”._

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Junto a uno de nuestros clientes desarrollamos una plataforma por la cual circulan pagos online durante las 24hs de los 365 días del año. Esta característica del sistema hace que un bug en producción tenga potencialmente un gran impacto, no solo en nuestros clientes, que pueden ver reducidas sus ventas, sino en los usuarios finales que pueden no llegar a comprar entradas para un show, pagar un impuesto fuera de término, o perderse una promoción de un artículo electrónico. Por lo tanto, estos incidentes deben ser encarados de una manera rápida y certera.

Idealmente, este tipo de errores no deberían ocurrir y hay que trabajar para minimizarlos, pero somos una industria de personas y, como todo ser humano, cometemos errores más veces de las que nos gustaría.

Cómo atacar este tipo de errores en producción es algo bastante estándar en la industria: existe una primera línea, conocida como **Helpdesk**, que se encarga de resolver inconvenientes triviales (típicamente configuraciones o mal uso de la plataforma por parte de los usuarios) mediante reclamos de los clientes que detectan el error.

Si el problema no puede ser resuelto, la segunda línea, **Soporte**, se encarga de atacarlo, en este caso a través de desarrollos menores que en poco tiempo pueden ser implementados. En el caso de que el error persista o que su corrección requiera desarrollos más complejos, interviene el **Equipo de Desarrollo**, tomando como base el análisis realizado por la línea anterior.

Pero, ¿cómo compite la resolución de estos errores en producción con las historias planificadas para el _sprint_ por el equipo de desarrollo?

# En búsqueda de la solución

Lo primero que consideramos fue qué tan frecuentes son estos casos y qué tanto afecta al desempeño del equipo.

En nuestro equipo observamos que no estábamos cumpliendo con los compromisos del _sprint_. Una de las hipótesis que planteamos fue que esto se debía a que le estábamos dedicando mucho tiempo a _“cosas que entran por la ventana”_, entre ellas, corrección de incidentes productivos. Entonces, decidimos _medir_ para confirmar nuestra suposición.

Empezamos a medir que el _gap_ de tareas que no terminamos en el sprint era muy parecido al tiempo que le dedicamos a la corrección de incidentes productivos (si, también teníamos problemas de calidad de definiciones de las historias y de cómo las estábamos estimando, pero eso será para otro post).

Por ejemplo, al final de uno de nuestros _sprints_ observamos en nuestro _burndown chart_ que la cantidad de tareas no terminadas equivalía a 62 hs (en ese momento aún no estimábamos en esfuerzo relativo con _story_ _points_). Como durante ese _sprint _estuvimos midiendo el tiempo que nos llevaban la corrección de incidentes productivos, en este caso 99, pudimos contrastar esos dos valores:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Este comportamiento, en el que tiempo que le dedicábamos a las tareas no planificadas era mayor al tiempo que nos hubiesen insumido las tareas que nos quedaban sin terminar, lo observamos durante algunos _sprints_. De esta manera corroboramos nuestra hipótesis.

El primer enfoque que utilizamos para resolver este problema fue dedicar una capacidad determinada del _sprint_ a la corrección de incidentes, digamos un 20%, y estimar las tareas del teniendo en cuenta esa reducción en la capacidad.

No nos fue muy bien con esta estrategia. En primer lugar porque es muy difícil saber cuánto es ese 20% y, por otro lado, no teníamos restricciones de cuántas tareas por la ventana tomar al mismo tiempo porque todo era prioritario. Todo el equipo podía estar resolviendo incidentes productivos y hasta no cerrarlos no tomábamos tareas del sprint.

Con estos resultados sobre la mesa decidimos cambiar el enfoque: en vez de un porcentaje de capacidad del equipo decidimos dedicar una cantidad fija de personas a la resolución de estos incidentes. En este caso fueron dos y fueron conocidas dentro del equipo como los **Bomberos**.

_Mauro y Facu, bomberos del Sprint, con su tablero Kanban detrás_

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Estas son algunas de las reglas que seguimos:

- Dado que nuestro equipo cuenta con dos células de trabajo, cada una aporta un bombero por _sprint_.
- El bombero es rotativo, de manera que no sea siempre la misma persona la que toma este trabajo.
- El _sprint_ se planifica teniendo en cuenta la velocidad del equipo sin el bombero incluido. Como todos los _sprints_ hay un bombero, la velocidad sigue siendo la misma.
- Las tareas de bombero se organizan como Kanban, hay un backlog que se va priorizando dinámicamente de manera de ir atacando las tareas que más inconvenientes generan en producción.
- Nunca hay más de dos tareas en progreso, una por cada bombero. Si entra algo nuevo compite con las otras dos.
- Es fundamental la comunicación con el PO, de manera de estar todos alineados en la resolución de lo que genera más valor.
- El resto del equipo encara las tareas de _sprint_ y ese es sólo su foco.
- El bombero puede pedir ayuda a alguno de los demás miembros del equipo en caso de ser necesario, pero es él quién está encargado de la resolución del inconveniente.
- Si el incidente es grande y genera un gran impacto, todos somos bomberos, ya habrá tiempo después para ver qué sucede con el _sprint_ (se cancela, se sacan tareas, etc)

Un punto importante a mencionar es que este esquema no debería ser para siempre. Un equipo que mejora constantemente debe ponerle foco a la calidad del producto que construye para minimizar los incidentes productivos que se puedan originar.

Por otro lado, no debe confundirse con un Esquema de Guardias, también necesario para este tipo de plataformas, en el que haya personas disponibles en todo momento en el caso de que salte alguna alarma de mal funcionamiento.
