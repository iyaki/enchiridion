---
title: "Deuda Técnica: lo que necesitás saber para poder gestionarla"
notion_id: 1712f138-6533-43f3-8068-5af1e216e98b
notion_url: https://app.notion.com/p/Deuda-T-cnica-lo-que-necesit-s-saber-para-poder-gestionarla-1712f138653343f380685af1e216e98b
last_edited: 2023-02-16T13:56:00.000Z
source_url: https://medium.com/redbee/deuda-t%C3%A9cnica-lo-que-necesit%C3%A1s-saber-para-poder-gestionarla-8ee255b73501
tags: ["Article", "redbee - Medium", "Español", "Programming", "Project Management", "System Design / Software Architecture"]
---
Para poder gestionar adecuadamente la deuda técnica, es necesario contar con algunas definiciones que le permitan al equipo ponerse de acuerdo en la mejor estrategia para abordarla. En este artículo, revisamos todo sobre este concepto, tan popular entre los equipos.

Una definición bien genérica de Deuda Técnica puede ser:

_**“El costo y los intereses a pagar a futuro por no hacer las cosas del todo bien hoy”.**_

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

El [concepto de Deuda Técnica](http://c2.com/doc/oopsla92.html) no es nuevo. Fue acuñado en 1992 por [Ward Cunningham](https://es.wikipedia.org/wiki/Ward_Cunningham), uno de los firmantes del Manifiesto Ágil, haciendo una analogía entre el código mal escrito y una deuda financiera.

Hoy en día esta comparativa puede aplicarse no sólo al diseño y escritura del código, sino también a otros aspectos del proceso de construcción de productos digitales, como por ejemplo el diseño de interfaces y la usabilidad.

Pero, ¿cómo es esta analogía? Veámoslo con un ejemplo.

Cuando una persona toma un préstamo, por ejemplo para comprar un auto, incurre en una deuda. El beneficio que obtiene al hacerlo es que cuenta con el auto antes de lo que lo hubiera hecho hasta tener todo el dinero necesario. Si paga regularmente las cuotas, la deuda contraída se cancela y, más allá de haber pagado de más, no se genera mayor problema.

Sin embargo, si la persona no paga la cuota, se le aplica una penalidad en forma de interés (generalmente creciente) y se acumula cada vez que no realiza el pago, es decir, empieza a deber mucho más dinero del que pidió prestado y se hace más difícil pagarla.

En caso de que la persona no pueda pagar las cuotas durante mucho tiempo, el interés acumulado puede hacer que la deuda total sea tan grande que la persona tenga que declararse en quiebra.

De forma similar, cuando un equipo opta por una solución rápida en lugar de una solución adecuada, introduce una deuda técnica. El beneficio que obtiene es que puede lanzar el producto antes de lo que lo hubiera hecho construyéndose de la forma correcta.

El equipo no va a tener grandes problemas si luego de lanzar el producto va corrigiendo aquello que no hizo del todo bien en la primera ocasión. Pero si el equipo no hace esas correcciones, éstas se irán acumulando haciendo que sea cada vez más difícil hacerlo.

El riesgo reside en que, con el tiempo, y en un caso extremo, podría ser imposible corregir esos problemas, y el producto podría tener que descartarse o hacerse de nuevo (sería nuestra quiebra).

Siguiendo la analogía, la deuda técnica está compuesta por el **capital** y los **intereses, **y podemos decir que:

- El **capital** corresponde a lo que no se hizo en la búsqueda del beneficio. Por ejemplo, el costo del auto. El capital asociado a la deuda técnica es proporcional al esfuerzo que el equipo emplearía en eliminarlo.
- Los **intereses** corresponden al esfuerzo extra que se tiene que pagar al querer hacer lo que no se hizo en su momento, o al querer hacer algo nuevo. Por ejemplo, es el costo extra que se paga sobre el valor del auto si se lo hubiera comprado sin tomar la deuda. Si no se paga el interés, se acumula con el tiempo y dificulta la implementación de cambios posteriores.
- **La tasa de interés** no es fija, sino que es en función del tiempo: cuanto más se ignore o se posponga, mayor será la deuda con el tiempo.
- El interés es de **naturaleza compuesta**, ya que el código de nuevas funcionalidades introducidas se entrelaza con el código existente con deuda técnica aumentando aún más la deuda.

Algunos **ejemplos típicos de deuda técnica** y conocidos por los equipos son:

- Hardcodeos, como textos de frontend, mensajes de error, números mágicos, contraseñas, etc.
- Código que no cumple convenciones ni buenas prácticas.
- Diseño sin seguir patrones.
- Falta de tests unitarios y/o otros tipos de tests automatizados.
- Utilizar un framework que tenga limitaciones para extenderlo posteriormente.
- Utilizar una arquitectura que no sea escalable para las necesidades del negocio.
- Desarrollar un front end sin considerar la usabilidad.
- Utilizar versiones viejas de librerías.
- Falta de logging y trazabilidad.
- No considerar seguridad.

Ahora que ya comprendemos lo que es la deuda técnica, es hora de aprender a gestionarla. Para eso, es necesario diferenciarla de algunos otros conceptos.

# No todo es deuda técnica

Es común para varios equipos etiquetar como deuda técnica algunas cosas que no lo son, como por ejemplo bugs, funcionalidades incompletas y mejoras técnicas del sistema.

Para visualizar qué es deuda técnica y que no, podemos ayudarnos del **Cuadrante de Kruchten**:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

En este cuadrante, se visualiza la deuda técnica dentro del conjunto de características que componen un producto.

Lo que nos indica este gráfico es que **no todo es deuda técnica**, pero también, que la deuda técnica es parte de las características del producto, por los tanto, _**los ítems de deuda técnica tienen que estar dentro del (mismo) backlog**_ que las funcionalidades nuevas, los bugs y los ítem de arquitectura. Esta definición es clave para poder gestionar la deuda técnica.

# La deuda técnica no necesariamente es mala

La palabra “deuda” puede sonar como algo malo, pero, así como ocurre con la deuda financiera, no necesariamente lo es.

El [Cuadrante de Martin Fowler](https://martinfowler.com/bliki/TechnicalDebtQuadrant.html) puede ayudarnos a distinguir mejor entre la deuda técnica “buena” y la “mala”.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Los ejemplos dentro del cuadrante anterior se tratan de Deuda Técnica, porque (más temprano que tarde) el equipo deberá hacer cambios al código que ya se tiene funcionando, con el costo que eso implica.

- Los ejemplos 1 y 2 son **Deuda Técnica Deliberada**, porque ambos equipos decidieron explícitamente tomar estas decisiones. Las causas de este tipo de deuda son claras y conocidas. Al ser intencional, es más probable que se registre esta deuda y que el equipo pueda gestionarla de manera eficiente en el futuro.
- En cambio, en los ejemplos 3 y 4 los equipos no sabían que estaban incurriendo en deuda técnica mientras estaban desarrollando su producto. Por eso se trata de **Deuda Técnica Inadvertida**. Las causas de la misma no son claras para el equipo y no la notarán en el momento, sino que será visible más adelante.
- Los ejemplos 1 y 4 se tratan de **Deuda Técnica Prudente**. En el ejemplo 1 porque el beneficio justifica el costo. En el ejemplo 4 porque, al desarrollar la primera versión, el equipo hizo lo mejor que pudo con el conocimiento que tenía en ese momento. Que el equipo encuentre una mejor forma de hacerlo luego de haber aprendido cómo funciona el producto, no solo es entendible, sino también esperable.
- **Deuda Técnica Imprudente** son los ejemplos 2 y 3. El equipo del ejemplo 2 incurre en este tipo de deuda porque es negligente al decidir no testear todos los casos. En el caso del equipo del ejemplo 3 la deuda es por incompetencia, ya que nadie tenía los conocimientos necesarios como para haber realizado la API según las buenas prácticas definidas en REST.

Entender esta clasificación permite al equipo **reconocer qué tipo de deuda evitar y qué tipo de deuda gestionar**.

# La deuda técnica es una función del tiempo

La deuda técnica es una función del tiempo porque evoluciona (o no) a medida que el sistema (o parte de él) evoluciona.

Ese progreso puede tener dos formas:

1. El sistema directamente no evoluciona, no crece, entonces la deuda técnica no tiene ningún costo o interés. Por ejemplo: un script de migración de datos de una base a otra que sólo se ejecutará una vez.
2. El sistema evoluciona con más funcionalidades y/o más requerimientos no funcionales como mayor cantidad de usuarios. Este es el caso más habitual.

Ahora, algunas consideraciones a tener en cuenta sobre la evolución del código:

- Si un ítem de deuda técnica se encuentra en una porción del código que no se ve afectada por una evolución futura, entonces se trata de **deuda técnica potencial**. Y, cuando afecta a la evolución, es **deuda técnica real**.
- Un sistema puede tener una deuda potencial muy grande en un momento dado. En función del siguiente incremento de la evolución, una parte (o toda) esa deuda técnica, puede transformarse en deuda real.

Esta distinción es útil para priorizar qué ítems de deuda técnica se resolverán, ya que conviene **enfocarse principalmente en la deuda real primero** y luego en la deuda potencial, dependiendo de su probabilidad de ocurrencia.

El siguiente gráfico, adaptado de [_Managing Technical Debt: Reducing Friction in Software Development_](https://www.amazon.com/-/es/Philippe-Kruchten/dp/013564593X), muestra el punto de inflexión (T3) en el cual la deuda técnica deja de ser un activo (el beneficio supera el costo) para transformarse en un pasivo.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Monitorear este momento en el tiempo es un factor clave a la hora de tomar deuda técnica deliberada y de gestionar la existente.

# Gestionando la deuda técnica adecuadamente

Los conceptos que vimos anteriormente son la base para entender cómo gestionar de forma adecuada la deuda técnica. [En un próximo artículo comentaré sobre algunas estrategias](https://medium.com/p/89cd200898ea) para hacerlo, como evitar la deuda técnica imprudente, identificar la deuda técnica inadvertida, registrar los ítems de deuda técnica existentes en el product backlog, dimensionarlos, priorizarlos y planificar su resolución.

¿Qué tipos de deuda técnica han experimentado ustedes? ¿Cómo la gestionaron?
