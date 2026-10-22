---
title: "[tl;dr] The Staff Engineer's Path"
notion_id: b5ba9919-2d10-477b-971e-d48870d20a09
notion_url: https://app.notion.com/p/tl-dr-The-Staff-Engineer-s-Path-b5ba99192d10477b971ed48870d20a09
last_edited: 2026-09-21T17:03:00.000Z
source_url: https://olano.dev/blog/tldr-the-staff-engineers-path/
tags: ["Español", "Career Growth", "Article", "Apuntes inchequeables (Facundo Olano)"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[_Peopleware_](https://www.goodreads.com/book/show/18895165-peopleware), uno de los libros clásicos de la industria del software, plantea así su premisa:

> The major problems of our work are not so much technological as sociological in nature. (…) The researchers who made fundamental breakthroughs are in a high-tech business. The rest of us are appliers of their work. We use computers and other new technology components to develop our products or to organize our affairs. Because we go about this work in teams and projects and other tightly knit working groups, we are mostly in the human communication business. Our successes stem from good human interactions by all participants in the effort, and our failures stem from poor human interactions.

The main reason we tend to focus on the technical rather than the human side of the work is not because it's more crucial, but because it's easier to do.

Partiendo de una larga educación en "ciencias duras" y pasando los primeros años de nuestras carreras enterrados en absorber los detalles técnicos del trabajo, a algunos programadores nos lleva un tiempo asumir esta realidad.

Entiendo que los autores de _Peopleware_ fueron pioneros en llamar la atención sobre los aspectos humanos de la profesión pero, aunque los principios que proponen siguen siendo válidos, el tratamiento quedó obsoleto. El libro dedica un simpática cantidad de espacio a explicar cómo evitar las interrupciones telefónicas o cómo distribuir los cubículos en la oficina. Es un libro pre-internet. Ese espacio lo ocupa en la bibliografía contemporánea [_The Manager's Path_](https://www.goodreads.com/book/show/33369254-the-manager-s-path), de Camille Fournier, que incorpora la evolución de la industria tecnológica durante las últimas décadas en una propuesta de carrera de management: desde relacionarse con el propio manager, a gestionar equipos, gestionar otros managers y hasta organizaciones enteras. Una de las primeras etapas de ese camino es la de Tech Lead, el desarrollador experimentado que todavía contribuye individualmente, que no tiene la responsabilidad ni la autoridad para manejar gente pero sí ejerce influencia y liderazgo sobre los demás.

Existe el prejuicio de que, llegado a ese punto, el profesional de software tiene que tomar una decisión: o "salta" a la carrera de manager o se "queda" como ingeniero _senior_. No hay vergüenza en ninguna de las dos opciones, pero sí un aire de resignación: o sacrifico el trabajo técnico que me apasiona y que me trajo hasta donde estoy, o me conformo con el mismo puesto por el resto de mi carrera[1](https://facundoolano.github.io/2023-01-16-tldr-the-staff-engineers-path/#fn.1). Pero hay una _tercera posición_. Existe una bifurcación, un camino alternativo para crecer profesionalmente sin gestionar personas, que, en algunas organizaciones, incluye posiciones como _Staff_, _Principal_ y _Distinguished Engineer_ (agrupadas bajo el término _Staff+_).

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

El libro [_The Staff Engineer's Path_](https://www.goodreads.com/book/show/61058107-the-staff-engineer-s-path) de Tanya Reilly, que documenta ese camino menos conocido, es efectivamente el lado B del _Manager's Path_, la otra mitad del plan de carrera[2](https://facundoolano.github.io/2023-01-16-tldr-the-staff-engineers-path/#fn.2), [3](https://facundoolano.github.io/2023-01-16-tldr-the-staff-engineers-path/#fn.3).

## Notas

¿Qué hace exactamente un _Staff Engineer_? En buena medida depende de la organización y de la posición relativa dentro de esa organización. No es lo mismo hacerlo en una empresa que lo contempla en su plan de carrera que hacerlo _de facto_ como Tech Lead de un equipo reducido. Will Larson, que entrevistó a varios Staff Engineers, [identifica cuatro ](https://staffeng.com/guides/staff-archetypes)[_arquetipos_](https://staffeng.com/guides/staff-archetypes):

- El **Tech Lead**, que guía el enfoque y la ejecución de un equipo en particular, generalmente a la par de un manager.
- El **Arquitecto**, que suele dirigir y garantizar la calidad de un área de la organización, posiblemente asistiendo pero sin formar parte de ningún equipo.
- El **Solucionador**, que salta de problema en problema, profundizando para destrabar situaciones complejas y allanarle el camino a los equipos que lo sucedan.
- La **Mano Derecha**, que en organizaciones grandes trabaja directamente para algún ejecutivo, aumenta su alcance y le da perspectiva técnica.

Cada uno es un enfoque distinto, una posible "implementación" de la idea de escalar la influencia y multiplicar el impacto del líder técnico. Cuál o cuáles adopte una persona particular depende de sus inclinaciones y de la estructura y la cultura de la organización en la que se desempeña. Pero en todos los casos sigue valiendo la premisa de _Peopleware_; progresar como líderes técnicos implica mejorar como comunicadores, elegir cuidadosamente cómo usar el tiempo, delegar y empoderar a los demás para abordar problemas de más grandes. Aunque no requiera manejo directo de personas, el Staff Engineer pasa más tiempo en reuniones, más tiempo "surfeando la política", menos programando. De nuevo: no es necesariamente un rol atractivo para cualquiera, y no tiene nada malo optar por dedicarse a la tecnología y maximizar el tiempo de código.

Yonatan Zunger [identifica cuatro disciplinas](https://leaddev.com/sites/default/files/2021-09/Role%20and%20Influence%20The%20IC%20Trajectory%20Beyond%20Staff.pdf) en los equipos profesionales:

- La **habilidad técnica básica** de un puesto, por ejemplo: programación, diseño UI/UX, etc.
- La **gestión de producto**: determinar _qué_ hay que hacer y _por qué_, y elaborar una narrativa al respecto.
- La **gestión de proyecto**: ocuparse de los aspectos prácticos para lograr los objetivos, darle seguimiento a las tareas, eliminar bloqueos.
- La **gestión de personas**: convertir grupos de personas en equipos, ayudarlos a crecer profesionalmente y manejar conflictos.

Cada una es una disciplina distinta y la mayoría de los proyectos requiere de las cuatro, independientemente de quién la ejecute. Me parece interesante distinguirlas porque alguien las termina ejecutando en la práctica, aunque sea implícitamente[4](https://facundoolano.github.io/2023-01-16-tldr-the-staff-engineers-path/#fn.4). Si no hay alguien ocupando un puesto dedicado (por ejemplo un _Project Manager_ o un _Product Owner_), es probable que la responsabilidad recaiga en el Staff Engineer de turno.

El libro se divide en tres partes, por cada uno de los tres "pilares" de los roles Staff+: mirada global (_big-picture thinking_), ejecución de proyectos (_project execution_) y subir el nivel de los ingenieros con los que trabajamos (_leveling up_). De la primera parte, me interesó la propuesta de _evitar los máximos locales_.

El trabajo de desarrollo de software es una larga sucesión de decisiones, constantemente tenemos que optar entre alternativas mediante un análisis de _tradeoffs_, de costos y beneficios. Con la experiencia aprendemos a evitar la arbitrariedad e ignorar nuestras preferencias personales en esas decisiones, las calibramos según los objetivos de equipo. El problema que señala Reilly es que, concentrados en el día a día de nuestro equipo, nos falta el contexto necesario para medir el impacto que nuestras decisiones tienen en el resto de la organización. Optimizamos para el máximo local, una solución ideal para nuestro equipo que no necesariamente es la mejor para el conjunto de la organización. El planteo me interesó porque aplica un razonamiento que ya había visto en otros contextos:

- Hablando del [manejo de la complejidad en el diseño de sistemas](https://facundoolano.github.io/2022-11-28-posdata-sobre-la-complejidad-esencial/), John Ousterhout nos advierte sobre la práctica común de dividir módulos para hacerlos más simples individualmente, al costo de aumentar la complejidad total del sistema. Un caso parecido es el de simplificar la implementación de un componente al costo de complejizar su interfaz.
- En su ensayo [Choose Boring Technology](https://mcfunley.com/choose-boring-technology), citado en el libro, Dan McKinley habla del peligro de "elegir la herramienta correcta para cada trabajo" con este ejemplo: si el lenguaje ideal para un nuevo sistema es Python pero el resto de los sistemas de la organización están implementados en Ruby, la complejidad de tener que manejar un lenguaje nuevo supera el beneficio local de la elección "ideal".

Quizás el capítulo más jugoso es el que trata cómo administrar el tiempo de trabajo. Se da por sentado que a medida que se sube en la "escalera técnica", aumenta la autonomía del ingeniero y es al menos parcialmente responsable de elegir en qué proyectos invertir su tiempo. El libro ataca el problema desde distintos ángulos.

Por empezar, tenemos qué identificar cuán ocupados nos gusta estar. ¿Preferimos estar siempre al máximo de nuestra capacidad y fundirnos cuando surjan emergencias o cambios de prioridades? ¿Preferimos reservar margen de maniobra con el riesgo de aburrirnos un poco mientras tanto?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Hay que asumir que en una organización saludable, siempre va a haber más trabajo disponible del que un Staff Engineer puede atacar. De todo el universo de posibles proyectos que podrían aprovechar nuestra atención, ¿cuál es el que más vale la pena, el que maximiza nuestro aporte? Hunter Walk ([citado por Larson](https://staffeng.com/guides/work-on-what-matters)) propone esta clasificación para orientarnos:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

- Proyectos de **alto impacto/bajo esfuerzo**: son la elección obvia, pero también los primeros que desaparecen a medida que crece la organización.
- Proyectos de **alto impacto/alto esfuerzo**: donde idealmente deberíamos pasar la mayor parte de nuestro tiempo, y entre los que tendremos que priorizar con algún otro criterio.
- Proyectos de **bajo impacto/alto esfuerzo**: los que uno supone que nunca deberían realizarse pero que, si no vigilamos con honestidad, aparecen y absorben energía.
- Proyectos de **bajo impacto/bajo esfuerzo**: lo que Walk llama _snacking_ (porque llenan pero no alimentan). Siempre es tentador atacar proyectos fáciles de completar, y puede ser útil elegirlos de vez en cuando para subir la moral o recuperar la energía, pero si es lo único que hacemos estamos perdiendo el tiempo.

Así como tenemos que considerar el costo/beneficio para la organización, tenemos que ser honestos al gestionar los recursos que tenemos como individuos. Para ilustrar esta idea, Reilly hace una analogía con el juego _The Sims_[_5_](https://facundoolano.github.io/2023-01-16-tldr-the-staff-engineers-path/#fn.5): propone que cada profesional dispone, además de su tiempo, de una serie de recursos (energía, credibilidad, calidad de vida, habilidades, capital social) que se consumen o se recargan según el trabajo que hacemos. Cada proyecto tiene un efecto sobre esos recursos que tenemos que tomar en consideración al elegirlo[6](https://facundoolano.github.io/2023-01-16-tldr-the-staff-engineers-path/#fn.6):

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Si el último proyecto nos dejó exhaustos, lo mejor para nosotros y para la organización es tomarnos vacaciones antes de encarar otro proyecto demandante; si sentimos que nos estamos quedando atrás en habilidades tal vez sea mejor priorizar proyectos con oportunidades de aprendizaje, etc.

## Conclusión

Más allá de que el título de Staff Engineer sea raro fuera de un grupo de organizaciones medianas o grandes, lo interesante del libro de Tanya Reilly es que muestra un camino de crecimiento desde el liderazgo técnico y nos invita a razonar sobre la dinámica del trabajo en equipo y la influencia organizacional a escalas mayores de las que estamos acostumbrados. Nos recuerda, como antes _Peopleware_ y como _The Manager's Path_, que estamos en el negocio de la comunicación y de las interacciones humanas, que los "soft skills" no son territorio exclusivo de los PMs. Por eso, me parecen lecturas valiosas para cualquier profesional del software, independientemente de su posición o de su proyecto de carrera.

## Referencias

Muchas de las ideas y técnicas que discute el libro (y que son difíciles de resumir acá) surgen de otros autores, y uno de los aportes más valiosos de Tanya Reilly es su [curaduría de artículos y charlas](https://noidea.dog/staff-resources). Listo acá los que me gustaron.

- Hillel Wayne, The Crossover Project:
- Tanya Reilly, [Being Glue](https://noidea.dog/glue).

### Notas:

1. En el peor de los casos, esta disyuntiva empuja a gente valiosa a dedicarse al management, una disciplina distinta y para la que quizás no tenga interés o aptitud, dañándose a sí mismo y a las personas que termina manejando.
2. Este libro no es el primer esfuerzo en documentar el rol de Staff Engineer. En su sitio [staffeng.com](https://staffeng.com/), Will Larson reúne guías y entrevistas a profesionales que lo ejercen. _The Staff Engineer's Path_ lo cita en varias oportunidades, así que, entre este libro y una selección de las guías, el material queda bastante bien cubierto.
3. El rol de Staff Engineer se superpone un poco con lo que tradicionalmente, sobre todo en los ambientes "enterprise", se entiende por Arquitecto de Software. Esta visión está cubierta por el libro _Foundations of Software Architecture_, que [comenté en otro post](https://facundoolano.github.io/2020-09-15-tldr-fundamentals-of-software-architecture/), y la secuela de los mismos autores, _Software Architecture: The Hard Parts_.
4. Otra razón por la que me gusta esta clasificación es que separa la gestión de personas de la gestión de proyectos. La gestión de proyectos me parece una aptitud técnica accesible para un ingeniero con experiencia, mientras que la gestión de personas es un mundo aparte. Todos nos cruzamos con esa criatura mitológica, lo que Camille Fournier llama el _Zar de los Procesos_ que, contrario a lo que pedía el [Manifiesto Agile](https://agilemanifesto.org/), desestima a los individuos y sobredimensiona la importancia de seguir un proceso específico a rajatabla. Se trata del revés exacto de la caricatura del programador cuadrado que cree que lo único que importa es el código.
5. La autora expone una versión temprana de esta analogía [en su blog](https://noidea.dog/blog/how-many-vacation-days-does-it-take-to-change-a-lightbulb).
6. La idea no es elegir exclusivamente los proyectos según las necesidades individuales de las personas sino tenerlas siempre presentes, no hacer de cuenta que esas necesidades no existen o no importan.
