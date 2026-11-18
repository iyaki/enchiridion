---
title: "Guía de estrategias de gestión de deuda técnica"
notion_id: e03fade2-bd19-4054-a97f-51c8d99384b2
notion_url: https://app.notion.com/p/Gu-a-de-estrategias-de-gesti-n-de-deuda-t-cnica-e03fade2bd194054a97f51c8d99384b2
last_edited: 2023-04-18T12:07:00.000Z
source_url: https://medium.com/redbee/gu%C3%ADa-de-estrategias-de-gesti%C3%B3n-de-deuda-t%C3%A9cnica-89cd200898ea
tags: ["Article", "redbee - Medium", "Español", "Programming", "System Design / Software Architecture", "Product Management"]
---
# **¿Qué significa gestionar la deuda técnica?**

El objetivo de una buena **gestión de deuda técnica** es que el equipo la mantenga bajo control de una forma planificada, como parte de su proceso de desarrollo, obteniendo beneficios de ella y evitando sufrir consecuencias indeseadas.

La **gestión de la deuda técnica no es una actividad de una sola vez**; es una parte integral y continua del ciclo de vida del desarrollo de software. ¿Cómo?

Hay **dos formas principales** de mantener la deuda técnica bajo control:

1. **Minimizar** la creación y acumulación de deuda técnica imprudente.
2. **Pagar** la deuda existente de manera eficiente. Para esto, primero hay que identificarla, luego hacerla visible, dimensionarla, priorizarla y, por último, planificar su resolución.

![image](https://miro.medium.com/v2/resize:fit:700/0*LZuZKyKj_5gxZvSK)

A continuación se describen algunas estrategias para hacerlo, dentro de 6 categorías, que son las mismas que vimos al principio: minimizar, identificar, visibilizar, dimensionar, priorizar y planificar la deuda técnica.

# **1 — Minimizar la creación y acumulación de deuda técnica**

Las estrategias a aplicar para minimizar la creación de deuda técnica nueva y la acumulación de la existente, dependen del tipo de deuda técnica del que se trate.

![image](https://miro.medium.com/v2/resize:fit:700/0*wztU71rGtAz_KQiq)

A — Para evitar la [**deuda técnica inadvertida e imprudente**](https://medium.com/p/8ee255b73501), la mejor estrategia es que el **equipo cuente con los conocimientos y habilidades necesarias** para construir el producto que está desarrollando. Se deben diseñar los equipos para que tengan una distribución de roles, seniorities y experiencia que aseguren que no pasarán de largo este tipo de deuda. No solo se trata de realizar cambios de personas para que esto suceda, sino que se puede recurrir a capacitaciones y/o mentoring de los miembros actuales.

B — La mejor defensa para cortar por lo sano la [**deuda técnica deliberada e imprudente**](https://medium.com/p/8ee255b73501)** **viene de la mano de la **definición y la aplicación de buenas prácticas de ingeniería** en el equipo, ya que al hacer explícito lo que se espera que suceda con el código, hay menos probabilidades de desconocimiento.

Algunas de estas prácticas que pueden atajar la deuda técnica son _Pair Programming, Code Reviews, Pull Requests, Sesiones de Diseño, Tests Automatizados, Convenciones de Código, Builds y Deploys automatizados, etc_.

C — Si, además, la aplicación de algunas de estas prácticas está incluida dentro de una **sólida Definition of Done**, el equipo tendrá una barrera de contención explícita que le permitirá evitar este tipo de deuda técnica.

Algunos puntos que se pueden incluir dentro de esta _Definition of Done_ son:

- _Se diseñó la solución respetando los lineamientos de arquitectura._
- _El código fue revisado por un par._
- _La funcionalidad contiene tests unitarios._
- _La cobertura de tests unitarios no disminuye._
- _Se respetan las convenciones de código establecidas por el equipo._
- _Contiene logs que permiten la trazabilidad._
- _El código fue analizado por una herramienta de análisis estático y no arrojó code smells ni vulnerabilidades de seguridad._

D — Por último, antes de llegar a la decisión de tomar [**deuda técnica deliberada y prudente**](https://medium.com/p/8ee255b73501), conviene hacerse algunas **preguntas para tratar de evitarla**:

- _¿Obtenemos algún beneficio al tomar este tipo de deuda?_
- _¿El beneficio obtenido es mayor al costo e intereses a pagar en el futuro?_
- _¿Los plazos que queremos acelerar tienen sentido? ¿Son realistas? ¿Podemos esperar un poco más y hacer las cosas mejor?_
- _¿Es posible balancear el costo y el beneficio? Por ejemplo, ¿es posible reducir un poco los beneficios y achicar los costos?_

La [**deuda técnica inadvertida y prudente**](https://medium.com/p/8ee255b73501) es inevitable (y esperable). Por lo tanto, lo único que se puede hacer es gestionarla con las estrategias que se mencionan a continuación.

# **2 — Identificar la deuda técnica existente**

Aunque todos los sistemas tienen algún tipo de deuda técnica, esta no siempre es conocida por el equipo. Por lo tanto, el primer paso en la gestión de deuda técnica es poder identificarla. Estas son algunas estrategias:

![image](https://miro.medium.com/v2/resize:fit:700/0*borOuwaIeWCcpn9f)

A — Todo integrante del equipo tiene alguna idea sobre dónde existe deuda técnica. Mediante una **sesión de equipo** cada persona puede aportar su conocimiento y, entre todas, construir un **listado** de la deuda técnica conocida.

B — Existen **herramientas de análisis de código** que el equipo puede utilizar para obtener información que sirva de punto de partida sobre dónde ir a buscar deuda técnica existente.

Las herramientas, como [SonarQube](https://www.sonarqube.org/), analizan el código y otorgan métricas sobre duplicidad, complejidad ciclomática, porcentaje de cobertura de tests, code smells, spelling, cantidad de líneas de código, etc.

El repositorio de código también es fuente de información. Por ejemplo, puede indicar qué archivos se tocan con mayor frecuencia, cuántos archivos se modifican por funcionalidad, qué paquetes son los más utilizados.

Otro tipo de herramientas que se pueden utilizar son las que analizan el código en busca de vulnerabilidades de seguridad.

Es probable que este tipo de herramientas proporcionen una larga lista de problemas, y esos problemas pueden o no ser deuda técnica. Es por eso que, al utilizarlas, es necesario comprender las métricas que arrojan y analizar en mayor detalle de qué se trata, ya que las herramientas no tienen toda la información contextual sobre cómo fue construido el código.

C — Analizar **métricas sobre bugs** y **reclamos de usuarios** puede también darnos una idea de dónde ir a buscar posible deuda técnica. Si identificamos componentes o funcionalidades que destaquen por sobre los demás, podemos ir a revisarlos con mayor detalle.

D — Durante los** refinamientos y sesiones de diseño**, cuando el equipo esté trabajando en una nueva funcionalidad y detecte que para resolverla de forma óptima tiene que refactorizar un determinado módulo, puede marcarlo como Deuda Técnica si decide no resolverlo en esa ocasión.

E — Si, como resultado de una **code review** o durante un **pull request**, algún integrante del equipo detecta deuda técnica, puede marcarla, de manera que no se pierda.

F — En los **comentarios en el código** también se puede encontrar información sobre la deuda técnica existente. Muchos desarrolladores suelen dejar _TODOs_ en el código, que pueden revisarse en búsqueda de deuda técnica. Además, métodos con comentarios explicando qué hace el código suelen ser candidatos a ser refactorizados, y por lo tanto, pueden ser deuda técnica.

G — Se puede comparar la arquitectura actual con el diseño original en búsqueda de deuda técnica relacionada a **desvíos de la arquitectura original**. También es posible utilizar tests de arquitectura como [ArchUnit](https://www.archunit.org/) que indiquen cuando no se están respetando las dependencias o las capas, que limiten niveles de herencia o restrinjan el uso de determinadas excepciones.

H — Existe la posibilidad de revisar el **proceso de build y puesta en producción** de un sistema para identificar deuda técnica. Si este proceso es manual, lleva mucho tiempo, es tedioso y/o es propenso a errores, probablemente exista deuda técnica allí.

# **3 — Visibilizar (registrar) los ítems de deuda técnica**

No alcanza con identificar la deuda técnica existente. Para poder gestionarla, además hay que visibilizarla.

![image](https://miro.medium.com/v2/resize:fit:700/0*sZ7B2-0W2yLm2lyI)

A — Algunos equipos deciden armar un listado con los **ítems de deuda técnica** conocida, por ejemplo, en una planilla. Si bien esta estrategia da visibilidad, es recomendable **registrar los Ítems de Deuda Técnica dentro del Product Backlog**. De esta manera, el equipo tiene una sola fuente de información con respecto al trabajo a realizar, ya sean funcionalidades nuevas, corrección de bugs o pago de deuda técnica.

B — Al reflejar los ítems de deuda técnica dentro del product backlog, es conveniente **identificarlos como otro tipo particular de PBI **para distinguirlos de los demás, ya que requieren un tratamiento diferente**. **Esto se puede hacer creando un tipo de _issue custom_ en la herramienta de _issue tracking_ que utilice el equipo, agregando una etiqueta particular a los tipos existentes o bien utilizando una convención de nombres que permita identificar que se trata de un ítem de deuda técnica.

C — Para evitar que el product backlog se transforme en un basurero de deuda técnica y que los que pierdan visibilidad sean los demás ítems, es conveniente manejar una **granularidad de ítems de deuda técnica adecuada**. Se pueden agrupar dichos ítems por componente en que se encuentran, por funcionalidad que afectan, por naturaleza de la deuda técnica, etc.

D — Como cualquier otro PBI, los **ítems de deuda técnica deben tener atributos** que permitan entender de qué se tratan y que no sean simplemente un título.

Algunos de estos atributos pueden ser:

- Un **nombre** que lo identifique.
- Una **descripción** que explique de qué se trata. Por ejemplo, el lugar dentro del código en qué se encuentra, desde cuándo está, por qué se generó, las métricas obtenidas de las herramientas de análisis estático de código, alguna idea de solución, etc.
- Algún tipo de **clasificación** que sea útil para el equipo: por tipo de deuda técnica, por causa, por componente, etc.
- **Dependencias** y **relaciones** con otros PBIs. Si un ítem de deuda técnica tiene dependencia o relación con algún otro ítem del backlog (ya sea otro ítem de deuda técnica o no), puede ser conveniente marcar esa relación.
- El **costo de resolverla, **es decir, una estimación del esfuerzo requerido para pagarla.
- El **beneficio esperado de resolverla** y/o el **costo de no resolverla, **es decir, el impacto que genera que siga existiendo.

Si el equipo utiliza **comentarios del tipo TODO** en el código, dentro de ese comentario se puede generar una referencia al ítem del backlog en el que está reflejado, para buscar más información y tener una mejor trazabilidad.

Identificar y visibilizar la deuda técnica no es algo que se hace una sola vez, sino que es un trabajo continuo. Cada vez que el equipo incurra en deuda técnica deliberada, o cuando identifique deuda técnica inadvertida, debe registrarla de la misma manera para mantener el backlog actualizado.

# **4 — Dimensionar (estimar) los ítems de deuda técnica**

Medir la deuda técnica ciertamente es una tarea difícil. Pero poder asignar un valor numérico a los ítems de deuda técnica le permite al equipo, y en particular al Product Owner, decidir cuándo y cómo atacarlos.

En particular, los atributos a dimensionar son:

- El **costo de resolverla.**
- El **beneficio esperado de resolverla** y/o el **costo de no resolverla.**

![image](https://miro.medium.com/v2/resize:fit:700/0*nVhoso7y7-1vcE9e)

A — Para asignar un valor numérico al **costo de resolver un ítem de deuda técnica**, el equipo puede utilizar las mismas herramientas y técnicas de estimación que utiliza para estimar los otros ítems de backlog. Por ejemplo, [**estimación en esfuerzo relativo con story points mediante planning poker**](https://medium.com/p/6fdac9150730).

Es importante que se estime el costo de estos ítems en la misma unidad en la que están reflejados las estimaciones de los demás ítems del backlog, para que puedan compararse.

B — El equipo también puede utilizar las mismas **técnicas de slicing** que utiliza con los demás PBIs en el caso de que un ítem de deuda técnica arroje una estimación muy grande que no entre en un sprint.

C — A la hora de estimar un ítem de deuda técnica hay que **tener en cuenta el capital y los intereses**. Esto implica estimar no solo hacer lo que originalmente no se hizo, sino también deshacer parte o todo de lo que sí se hizo.

D — Para la estimación de ítems de deuda técnica cobra mayor importancia **considerar el esfuerzo de las pruebas, **ya que la resolución de estos ítems generalmente involucra algún tipo de refactorización, que, sin tener tests automatizados, puede ser muy riesgosa.

E — Como la deuda técnica es una función del tiempo, la estimación del costo de resolverla también lo es. Es por eso que es importante,** revisar periódicamente la estimación** del esfuerzo que requiere resolver un ítem de deuda técnica, ya que puede cambiar más adelante.

F — Es recomendable dejar **registro de cuándo se hizo la estimación**, para, cuando se vuelva a revisar, poder determinar cuánto cambiaron las condiciones iniciales, y, eventualmente, volver a estimar.

G — Una de las razones por las cuales los ítems de deuda técnica suelen quedar en el backlog sin resolverse es que no suele estar claro para el equipo (en particular para el PO) el valor que se obtiene al atacarlos. Por lo tanto, expresar (y cuantificar) los **beneficios que se obtendrían al resolver los ítems de deuda técnica** es una buena estrategia. Esto no suele ser fácil, pero se puede recurrir a una **conjetura informada** (educated guess) para tratar de dimensionar el beneficio.

Algunos ejemplos de cómo expresar estos beneficios son: reducir a la mitad el time to market, aumentar el uptime al 98%, reducir al 70% la cantidad de bugs, etc.

H — Otra estrategia (no excluyente) que se puede utilizar, es ir por la negativa: En lugar de tener una dimensión del beneficio que aportaría resolver un ítem de deuda técnica, se puede tener una dimensión del **costo que implica no resolverla**, es decir, qué pasaría si se deja la deuda tal y cómo está.

[Martin Fowler propone una idea de cómo ](https://martinfowler.com/bliki/EstimatedInterest.html)[**medir los intereses**](https://martinfowler.com/bliki/EstimatedInterest.html)[ que se pagan](https://martinfowler.com/bliki/EstimatedInterest.html): cuando el equipo complete un PBI, se estima el esfuerzo real que les llevó realizarlo y el esfuerzo que le hubiera llevado si no hubiera tenido deuda técnica. La diferencia entre ambas estimaciones es el interés que se pagó.

Por ejemplo: “me llevó un día más porque tuve que agregar tests unitarios que faltaban” o “tuve que refactorizar 5 clases porque no entendí qué hacían”

Claramente este enfoque no es el mejor porque se está suponiendo un estado ideal del sistema. Y además supone que el interés que se paga es sólo en ese PBI que se tocó.

I — Para mejorar este último problema, si se crea un ítem de deuda técnica, por cada PBI que se desarrolla y es afectado por ese ítem de deuda técnica, se puede ir registrando el interés, de la forma antes mencionada, que fue pagando.

Con este **registro de intereses pagados**, el equipo puede tener una dimensión del costo de no resolver ese ítem de deuda técnica. Como alternativa a este enfoque, también se puede hacer una **conjetura informada** sobre el costo y trabajar con eso.

# **5 — Priorizar los ítems de deuda técnica**

Tener una idea del costo de resolver un ítem de deuda técnica y el beneficio que se obtiene al hacerlo (o el costo de no resolverla) ayuda a priorizar esos ítems en relación con los demás del backlog.

Si el equipo expresa el beneficio de resolver la deuda técnica o el costo de no resolverla en términos de negocio, la conversación con el PO sobre si un ítem de deuda técnica es más prioritario que otro ítem de negocio va a ser más fácil.

![image](https://miro.medium.com/v2/resize:fit:700/0*frpwIgyQI4u7wvFm)

A — Lo primero a considerar al priorizar un ítem de deuda técnica es [**determinar si se trata de deuda técnica real o potencia**](https://medium.com/redbee/deuda-t%C3%A9cnica-lo-que-necesit%C3%A1s-saber-para-poder-gestionarla-8ee255b73501)[l](https://medium.com/redbee/deuda-t%C3%A9cnica-lo-que-necesit%C3%A1s-saber-para-poder-gestionarla-8ee255b73501). La deuda técnica real es siempre más prioritaria que la potencial. Para hacer esa distinción, es necesario conocer la dirección futura del producto, algo que, generalmente, sabe el PO.

B — Para priorizar un ítem de deuda técnica se puede hacer un simple **análisis de Retorno de Inversión (ROI)**. Se compara el costo de resolver la deuda técnica y el beneficio que se obtendrá (o el costo de no resolverla). Este análisis no necesariamente debe ser muy matemático o exhaustivo.

C — Una vez obtenido el ROI de ese ítem de deuda técnica, se hace una comparativa con los demás ítems del backlog y se **ordenan según la prioridad**.

# **6 — Planificar la resolución de un ítem de deuda técnica**

Una vez priorizados los ítems de deuda técnica, junto a los demás ítems del backlog, es momento de agregarlos a un sprint en la planning y resolverlos.

![image](https://miro.medium.com/v2/resize:fit:700/0*16wCIQInTs3eOXl1)

A — La clave es **considerar a la resolución de ítems de deuda técnica como un trabajo continuo**. Es decir, dedicar en todos los sprints una proporción para resolverla.

B — Se puede seguir la **regla del 80/20**. Aquí el equipo planifica los sprints conteniendo un 20% de esfuerzo dedicado a ítems de deuda técnica y el 80% restante a los demás. Estos porcentajes pueden variar según las necesidad del equipo, pero esa es una proporción razonable. Seguir esta regla hace que la resolución de ítems de deuda técnica se convierta en algo habitual, para el equipo y para el PO que prioriza.

C — Algo a tener en cuenta es que suele ser desmotivante para los desarrolladores trabajar todo el tiempo en ítems de deuda técnica. Por eso es recomendable que no sean siempre las mismas personas quienes resuelvan los ítems de deuda técnica, sino que se vayan **rotando**.

D — Existen situaciones en las cuales la deuda técnica es tan abrumadora que no alcanza asignar porcentajes de sprint para su resolución. En estos casos pueden ser necesarios **sprints enteros** para hacerlo. Llegado a ese extremo, es recomendable **no prolongar por mucho tiempo** esta situación, ya que el equipo puede aburrirse de refactorizar código viejo, y eso impacta en la motivación.

E — En el caso de que existan ítems de deuda técnica que tengan un costo bajo de resolución, se los puede ir **resolviendo cuando se los vea**. Por ejemplo, si al desarrollar una funcionalidad, una persona se topa con una porción de código que necesita refactoring como parte de un ítem de deuda técnica, y su costo es bajo (en resolución y en impacto en pruebas de resto del código), se lo puede hacer en el momento.

F — Además, si llegando al final del sprint, ya se resolvieron todos los ítems del sprint backlog, se puede **utilizar ese tiempo remanente del sprint para solucionar ítems de deuda técnica**.

La clave para estas dos últimas estrategias, es tener ítems de deuda técnica de bajo costo, empezar por los más prioritarios y siempre actualizar el backlog con lo que se resuelve.

# **La gestión de deuda técnica como parte del proceso de desarrollo**

Las estrategias descritas anteriormente son útiles para que los equipos puedan gestionar de forma eficiente su deuda técnica, es decir, incorporar su resolución como parte del proceso de desarrollo. Cuáles se utilizan y en qué medida, dependerá de qué tan problemática sea la deuda técnica en determinado momento.
