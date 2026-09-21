---
title: "Bases de Domain Driven Design"
notion_id: eeed1ec1-bb04-4128-bc91-703830a5e2a8
notion_url: https://app.notion.com/p/Bases-de-Domain-Driven-Design-eeed1ec1bb044128bc91703830a5e2a8
last_edited: 2023-04-25T15:01:00.000Z
source_url: https://www.notion.so/Bases-de-Domain-Driven-Design-eeed1ec1bb044128bc91703830a5e2a8
tags: ["The Talking Bit - Fran Iglesias", "Español", "Domain Driven Design", "Article"]
---
Esta es una recopilación de artículos publicados por [Fran Iglesias](https://github.com/franiglesias) en su blog [The Talking Bit](https://franiglesias.github.io/) relacionados con el enfoque de diseño de sistemas, conocido con el nombre de _Diseño Guiado por el Dominio_ (abreviado DDD, por el nombre en inglés con el que Eric Evans lo acuña originalmente: Domain Driven Design).

# Índice

<!-- unsupported block: table_of_contents -->

---

# [Introducción](https://franiglesias.github.io/ddd-intro/)

> Vamos a comenzar un pequeño proyecto para aprender y profundizar en Domain Driven Design.

En este blog suelo tratar todo tipo de temas relacionados con desarrollo de software. Hasta ahora me he interesado especialmente por el testing y el refactoring, tocando algunos aspectos de buenas prácticas en el camino. Dado que he ido abordando temas sueltos según mis intereses en cada momento no puedo decir que exista una unidad, aunque, con el tiempo, hayan ido surgiendo oportunidades de agruparlos, por ejemplo en los libros que hemos publicado.

Ahora quiero profundizar en Domain Driven Design, pero de una forma un poco más sistemática. Es el tema que ha empezado a estar más presente en mi cabeza en los últimos tiempos. A eso hay que unir la demanda, que algunas personas me han comentado, de mostrar una visión aplicada de los distintos temas en proyectos concretos. Así que, esto es lo que me propongo:

Entender el Domain Driven Design
Conocer las herramientas de diseño e implementación que nos proporciona para el desarrollo de software
Explorar las relaciones de DDD con otras metodologías y paradigmas, como la arquitectura hexagonal, BDD, patrones de diseño, etc.
Aplicar este conocimiento en un proyecto práctico real
Mis materiales de referencia serán fundamentalmente:

Domain Driven Design, el libro original de Eric Evans.
Implementing DDD y DDD Distilled, los dos libros de Vaughn Vernon
DDD in PHP, el libro de Carlos Buenosvinos, Christian Soronellas y Keyvan Akbary.
Diversos blogs y libros de otros autores.
El proyecto práctico tendrá que ver con algunos trabajos que realicé en el pasado en mi antiguo sector laboral, el educativo. Así que el objetivo será desarrollar algún ejemplo de aplicación útil para un colegio. Es un dominio en el que tengo experiencia, lo que facilitará toda la parte que tiene que ver, precisamente, con el modelado y su traslación a un sistema de software.

Así que empecemos por lo primero, ¿qué es Domain Driven Design?

## Una introducción a DDD

Domain Driven Design es un paradigma de diseño de software que parte del modelado del espacio problema mediante la identificación de conceptos y procesos clave y su representación en un modelo a través de la colaboración entre los así llamados expertos del dominio y los desarrolladores.

DDD comienza con un proceso de knowledge crunching, es decir de adquisición masiva de conocimiento sobre el dominio. Se trata de aprender todo lo posible sobre el problema que vamos a modelar con el objetivo de comprenderlo y, por tanto, poder generar un modelo que lo represente correctamente. Esto se puede hacer principalmente a través de la conversación con los expertos de dominio, pero también con otras técnicas.

Pero, ¿qué es el dominio?

## El Dominio

El dominio es el espacio del problema, es aquello de lo que trata el negocio y, por tanto, de lo que tratará la aplicación. Forman parte del dominio los conceptos que se manejan, sus relaciones y los procesos que operan en él. El objetivo del DDD es representar el dominio en un modelo que luego se implementará en forma de software.

Normalmente, no debemos considerar el dominio de una forma genérica, sino que depende del contexto. El modelo del dominio para una empresa específica será distinto del de otra empresa. Obviamente habrá muchísimos elementos comunes, pero también existirán suficientes diferencias como para que no sean intercambiables. Esto no impida que se pueda utilizar el paradigma del Domain Driven Design en desarrollar productos genéricos, pero en ese caso hay que pensar no tanto en el dominio genérico, sino en el de las soluciones de software para ese dominio.

Antes hemos mencionado a los expertos de dominio, que son, ni más ni menos, las personas que trabajan día a día en ese negocio y, por tanto, conocen y manejan los conceptos, los entienden y saben explicar cómo se relacionan. Estas personas utilizan un lenguaje para hablar sobre ellos, nombrarlos, describirlos y accionarlos. El Domain Driven Design requiere que sea este lenguaje el que todos los implicados en el desarrollo utilicen y que se utilice como base para nombrar los modelos y expresarlos en el código. A esto le llamamos Lenguaje Ubicuo porque se pretende que se use y esté presente en todas partes: en las conversaciones sobre el dominio, en la documentación, en los diagramas y en el código. Por así decir, un experto del dominio debería poder observar el código e incluso hacerse una idea de lo que se ha expresado ahí.

Para poder comenzar a modelar el dominio, los desarrolladores tienen que adoptar el lenguaje ubicuo y la mejor manera de hacerlo es a través de la conversación con los expertos de dominio. En esta conversación se irán definiendo, pero también descubriendo, conceptos y procesos. De hecho, el diálogo facilita el aprendizaje por las dos partes. Con frecuencia, esta conversación ayuda a los propios expertos de dominio a descubrir, y solucionar, conceptos representados de forma ambigua o que necesitan un término que los represente de forma precisa.

Por ejemplo, en un sistema de información académica para un colegio será necesario dedicar un tiempo a representar correctamente los muchos conceptos distintos que se manejan al usar la palabra “curso”, como año académico, nivel educativo, grupo de clase, etc.

Una de las formas de afrontar esta conversación consiste en hablar acerca de los llamados Eventos de dominio, a través de una técnica conocida Event Storming. En esencia, consiste en identificar las cosas interesantes que pasan en el dominio y, a partir de ellas, extraer y describir los conceptos que intervienen y cómo se relacionan.

Así, siguiendo con el ejemplo del sistema académico, un evento interesante sería que un alumno recibe una calificación, lo que implica al alumno, la calificación en sí, la actividad sobre la que recibe la calificación, la materia o curso, el periodo de evaluación, etc. Ese evento implica que han sucedido otras cosas anteriores, como que el alumno se ha matriculado en el curso correspondiente, que el profesor ha programado una actividad evaluable y un largo etcétera de acontecimientos y conceptos que están implicados.

En el dominio hay un tema central que lo articula y da sentido: el core domain o núcleo del dominio. Es lo que hace único e irremplazable el negocio, lo que no se puede traspasar o subcontratar.

El core domain de un colegio es el aprendizaje, la formación de los estudiantes, su educación y todo lo que la hace posible. Muchas otras partes de su funcionamiento y estructura son necesarios para dar soporte a ese núcleo central y hacer posible que funcione.

Dentro de un dominio es posible identificar subdominios. Estos están formados por aquellos conceptos y procesos que están estrechamente relacionados entre sí. Los subdominios suelen corresponderse con las divisiones propias de una organización y están relativamente aislados entre sí, pero participan igualmente del dominio.

Los subdominios pueden ser clasificados como de soporte y genéricos. Los subdominios de soporte contribuyen al core domain facilitando que se pueda llevar a cabo. Son necesarios y tienen sentido en relación con el core y en ese sentido son específicos.

Los subdominios genéricos, por su parte, son aspectos que no son centrales del negocio, de hecho son comunes a muchos tipos de negocio, y podrían incluso ser subcontratados o externalizados.

Por tanto, el dominio, estructurado en subdominios, configura el espacio del problema.

## Bounded Contexts

El subtítulo del libro fundamental del DDD es “Abordando la complejidad en el corazón del software” y como se desprende del ejemplo anterior, la complejidad de un dominio puede llegar a ser apabullante y es fácil sentirse sobrepasados por ella.

Uno de los primeros pasos para mantener esa complejidad bajo control es identificar los llamados bounded contexts o contextos acotados. Un bounded context es una parte de un dominio en la que viven conceptos y procesos propios, que no participan en otras partes, y en la que los conceptos generales, compartidos entre varios contextos, pueden tener un significado distinto. Los bounded contexts son la representación de los subdominios en el espacio de la solución.

Veámoslo con el ejemplo escolar. El concepto Alumno en un colegio tiene varios significados según el contexto que estemos considerando:

En el contexto académico adquiere conocimientos, realiza actividades, es evaluado por ellas, estudia una serie de materias, etc.
En el contexto de secretaría, el alumno se matricula, se da de baja, tiene un expediente, recibe una titulación.
En el contexto administrativo, los servicios que utiliza tienen un coste, el cual se cobra a través de los correspondientes recibos.
En el contexto de servicios, puede ser usuario del comedor escolar, de las actividades extraescolares, etc.
Estos diferentes contextos son los que llamamos contextos acotados o bounded contexts. En cada uno de ellos, el concepto alumno lo consideramos con distintas propiedades, formando parte de distintos procesos e incluso tiene un ciclo de vida diferente. Se llaman acotados porque existe una frontera entre ellas, una frontera que podríamos considerar de significado.

La identificación de los bounded contexts es posible mediante diversas estrategias. A veces es fácil porque la propia estructura organizativa ya nos permite discriminar estos contextos: las divisiones de una empresa, por ejemplo, nos están dando una pista acerca de posibles contextos. En nuestro ejemplo escolar, es fácil ver que el departamento de administración y el departamento pedagógico son contextos separados. Sin embargo, la secretaría escolar mantiene muchas relaciones conceptuales con el departamento pedagógico, ¿hasta qué punto son contextos separados?.

## Elementos básicos: entidades, value objects y servicios

Para poder responder a esta pregunta tenemos que cambiar un poco de foco.

Hemos dicho que al hablar sobre el dominio identificamos conceptos y procesos. Vamos a ser un poco más precisos:

Algunos de estos conceptos se caracterizan porque conllevan una identidad. El objeto representado (frecuentemente una persona, pero puede ser cualquier cosa) mantiene una individualidad a lo largo de su ciclo de vida y a pesar de los cambios que pueda sufrir en sus propiedades. Por ejemplo, un alumno de un colegio se matricula en distintas asignaturas, pasa por distintos niveles, etc., pero sigue siendo el mismo alumno. Lo mismo podríamos decir de un profesor, de una asignatura, etc.

A estos conceptos que tienen identidad y ciclo de vida los denominamos Entidades.

Otros de estos conceptos se caracterizan por sus propiedades, no nos preocupa su identidad y no cambian a lo largo de su ciclo de vida. Como ejemplos, nos pueden servir una calificación (la nota en sí: notable, aprobado, suspenso), el precio de un servicio, y otros muchos. Son los Value Objects.

Ambos tipos de objetos tienen su propio comportamiento. Es decir: pueden hacer cosas relevantes para el dominio, Sin embargo, es frecuente que ciertas acciones necesiten coordinar de alguna manera varios de estos conceptos. Estamos hablando, entonces, de Servicios.

## Identificación de bounded contexts y módulos

¿Cómo nos ayuda lo anterior a identificar bounded contexts? Tenemos varios tipos de indicios:

Si un conjunto de Entidades, Value Objects y Servicios, tienden a aparecer juntos y cohesionados. Esto también nos podría estar anunciando la existencia de un Módulo, el cual tiene sentido como unidad de agrupación dentro de un bounded context. Sin embargo, para hablar de módulo no debe darse ninguna de las circunstancias siguientes:
Si en ciertos momentos nos interesan unas propiedades de las entidades y en otros momentos otras diferentes, tendríamos una indicación de que existirían al menos dos bounded contexts.
Puede ocurrir que una entidad deje de interesarnos como tal entidad y la podamos considerar como un Value Object. Y lo mismo en el otro caso, un Value Object que en ciertos momentos necesita una identidad. Ese cambio nos señalaría la presencia de un nuevo bounded context.
Intentemos verlo a través de ejemplos:

En nuestro sistema escolar, la entidad alumno aparece asociada con curso, materia, asignatura, profesor, notas, calificaciones, etc, entidades y valores con los que interactúa en diversos procesos. Parece claro que todos ellos formarían un contexto acotado, que bien podría ser el contexto académico.

Por su parte, la misma entidad alumno forma parte del contexto de administración, junto con conceptos como recibos, domiciliaciones, tasas, seguro escolar, etc.

Sin embargo, en el contexto académico son importantes algunas propiedades de la entidad, como puede ser la edad o el historial académico, que en el contexto administrativo resultan ser poco o nada relevantes. En este contexto, interesan propiedades como un número de cuenta, que son irrelevantes en el contexto académico.

En cierto modo, podrían modelarse como entidades distintas en cada uno de los contextos, relacionándose solo a través de su identidad porque representan a la misma persona real.

Estas relaciones entre contextos se conocen como Context Maps y permiten que los bounded contexts puedan, por así decir, hablar entre ellos acerca de las entidades que manejan.

Por otro lado, para el contexto del mantenimiento escolar, los alumnos no interesan por su identidad, sino por su número y alguna propiedad como la edad o si vive cerca o lejos del colegio, que nos sirve para determinar cosas como cuántas sillas y mesas se necesitarán, o cuántas plazas de autobús o de comedor tendríamos que disponer. En este contexto, el alumno no se modelaría como entidad, sino como value object: nos da igual qué persona concreta representa.

Dentro del contexto académico, por otra parte, podemos ver que ciertas entidades y valores aparecen estrechamente relacionados y, por tanto, pueden formar Módulos. Así, el módulo de alumnos puede agrupar a estos, con los grupos de clase y sus tutores. El módulo de asignaturas, con los cursos, niveles y profesores. Esta división en módulos nos aporta manejabilidad del modelo, pero no implica una separación funcional estricta de los conceptos.

## Agregados

A medida que aprendemos sobre el dominio nos daremos cuenta que ciertas entidades y value objects viven tan estrechamente relacionadas que, de hecho, mantienen una relación de inclusión y dependencia. Un ejemplo muy claro es todo lo que tiene que ver con el historial académico de un alumno: calificaciones, actividades, títulos, etc., no tienen sentido ni identidad por sí solos, sino y siempre con relación a un estudiante, de modo que cualquier interacción en ellas debe pasar necesariamente por él. Este tipo de relación forma lo que conocemos como Agregado. Agregado es, por tanto, un conjunto de entidades y value objects que forman un todo cohesionado que se maneja a través de una de las entidades, llamada Raíz del agregado, la cual gestiona todo lo que tiene que ver con las entidades y valores que incluye.

## Recapitulando

En este primer artículo hemos visto algunos de los elementos característicos que nos permiten articular el modelo de un dominio: bounded context, módulos, entidades, value objects, servicios y agregados.

En el próximo, desarrollaremos algunas implicaciones y cómo se concreta en el planteamiento de una aplicación.

---

# [Identificando el dominio y los subdominios](https://franiglesias.github.io/ddd-way-domain/)

> Domain Driven Design no puede tomarse como una guía con instrucciones para ir de un punto a otro siguiendo una ruta determinada. Es más bien, un conjunto de estrategias, métodos y tácticas para orientarse en un territorio desconocido y moverse hacia un objetivo deseado que, para mayor dificultad, es móvil.

Parte de la complejidad que el DDD nos ayuda abordar reside en que tanto el dominio como el conocimiento que tenemos de él es algo cambiante: las necesidades del negocio evolucionan, la precisión y la resolución con las que definimos los conceptos cambian, aprendemos a distinguir y descubrir nuevos matices y ramificaciones. Por simple que pueda parecer un dominio o un subdominio en una primera aproximación, a medida que aumenta nuestra comprensión, aumenta también nuestra capacidad de descubrir nuevas posibilidades, nuevas ramas. Y como no vivimos aislados en el mundo, el entorno hace que el negocio tenga que evolucionar, generando nuevas necesidades y requiriendo nuevas respuestas.

## DDD estratégico

En el artículo anterior introdujimos algunos de los elementos principales y característicos del Domain Driven Design. Buena parte de ellos forman parte de los llamados patrones estratégicos. Con cierta frecuencia se incide más en los patrones tácticos, que serían la concreción de aquellos en el código.

Pero, ¿qué son exactamente los patrones estratégicos y por qué son importantes?

Se podría decir que los patrones estratégicos son herramientas para pensar sobre el dominio y ayudarnos a construir el modelo rico y expresivo que lo representa. Básicamente:

- Conocer y aprender sobre el dominio
- Desarrollar un lenguaje ubicuo para poder hablar sobre él
- Identificar el core domain
- Identificar los diversos subdominios, de soporte y genéricos
- Caracterizar conceptos, procesos y eventos
- Definir lo contextos acotados a partir de los subdominios
- Crear el context map para representar las relaciones entre contextos

Nada de esto es código. Fundamentalmente se trata de definir un problema y construir un modelo que lo represente. Ese modelo será el que luego se exprese en forma de código utilizando los patrones tácticos.

Por cierto, en el artículo anterior mencioné Entidades, Value Objects y Servicios, que en principio forman parte de los patrones tácticos, no de los estratégicos. Sin embargo, diría que es útil tenerlos en cuenta a la hora de lidiar con los estratégicos, en particular para delimitar los bounded contexts.

## Definición del dominio

Una de las cosas que me gustan especialmente del DDD es la capacidad de su proceso para generar beneficios a la organización más allá de dirigir la construcción de su sistema de software. Es el hecho de ayudar a clarificar los propios objetivos y prioridades, los conceptos y los procesos.

Esto surge de la necesidad de explicarlos al equipo de desarrollo, respondiendo a preguntas y discutiendo los matices y los detalles.

Así que diría que un primer valor que puede obtenerse de aplicar Domain Driven Design es responder a la pregunta: ¿entiende la organización cuál es su core domain y es capaz de priorizar sus necesidades articulando sus dominios de soporte y sus dominios genéricos?

### Separando los subdominios genéricos

Sorprendentemente, o no, hay muchas empresas que no tienen bien definida la respuesta a esta pregunta. En particular, en el ámbito de nuestro ejemplo, siempre me ha sorprendido lo difuso que los colegios pueden llegar a tener su core domain y subdominios y, consecuentemente, el modo en que los gestionan.

Cuando, por la razón que sea, las cosas no están claras puede ser buena idea despejar el terreno identificando los subdominios genéricos que, como hemos dicho, son aquellos que podemos encontrar en cualquier organización y que podrían externalizarse o están, de hecho, externalizados, como sería el caso de nuestro hipotético colegio:

- Contabilidad
- Facturación
- Nóminas
- RRHH
- Fiscal
- Legal
- Otros

En este punto conviene volver a recordar que las circunstancias de cada empresa son diferentes y es perfectamente posible que un dominio que normalmente consideraríamos genérico puede ser de soporte en algunas de ellas, y viceversa. La pregunta es si alguno de esos subdominios funciona de una forma particular para nuestra organización o no.

Un tema interesante es que estos subdominios genéricos necesitarán datos procedentes de los subdominio de soporte o del core domain. Esto es algo que se concretará más adelante.

### Identificando los subdominios de soporte

Una vez que hemos identificado los subdominios genéricos sería el momento de enfocarnos en los de soporte: aquellas cosas que son necesarias para desarrollar el core domain y son específicos de nuestra organización, pero que no forman parte del core.

Por tanto, es hora de presentar nuestro colegio.

Llamémosle _Colegio Piruleta_. Se trata de un centro privado con más de 1200 alumnos desde Educación Infantil al Bachillerato. Aparte de la enseñanza reglada correspondiente a esas etapas, el centro ofrece servicios de comedor, actividades extraescolares, transporte escolar, servicio de guardería, biblioteca… Por otra parte, el _Colegio Piruleta_ lleva un tiempo experimentando la introducción de plataformas de enseñanza online investigando el desarrollo de un nuevo modelo de docencia y aprendizaje que le permita distinguirse de la competencia.

El colegio desea disponer de sus propias herramientas para gestionar estos servicios. En unos casos por la dificultad de encontrar soluciones de terceras partes que se ajusten a sus necesidades En otros casos porque consideran algunos de ellos como parte de su core domain y se sienten incómodos utilizando soluciones externas.

Esta breve introducción debería darnos algunas pistas:

Servicios como comedor, transporte, guardería, extraescolares, etc. podrían ser genéricos del sector, pero lo normal sería considerarlos como de soporte, debido a que:

- Ofrecerlos ayuda a captar alumnado, ofreciendo soluciones a las familias para los obstáculos prácticos que podrían desanimarlos a escoger este colegio.
- Contribuyen a la financiación.
- Pueden aprovecharse para reforzar la propuesta de valor que es específica del colegio: su proyecto educativo.

Esto es: no forman parte del core domain, pero ayudan de una manera específica a que este pueda desarrollarse. Requieren herramientas específicas: aplicaciones con las que gestionar los servicios de manera eficaz, permitiendo comunicarse con los subdominios genéricos, de forma que estos puedan adquirir los datos que necesitarán. Por ejemplo, los necesarios para realizar la facturación mensual.

Cada vez nos vamos acercando más al core domain.

### El core domain

Como hemos dicho, el core domain es el que hace única la organización. En el caso de un colegio es su proyecto académico y educativo que se concreta en aspectos metodológicos, de priorización de contenidos, de modo de relación con el alumnado, etc. Aunque podríamos decir que hay elementos que son comunes a cualquier colegio, tenemos que fijarnos en lo que es característico y único.

Por un lado tenemos un aspecto casi administrativo que sería la Secretaría académica, que se ocupa de la matriculación de los alumnos y los diversos procedimientos burocráticos que implica su vida escolar. Tiene un rol que es claramente de soporte y, como veremos más adelante, maneja prácticamente las mismas entidades y valores que el core domain.

Un segundo elemento sería la gestión docente que está estrechamente relacionada con el anterior: todo lo que tiene que ver con la evaluación del aprendizaje, organización de cursos, agrupaciones de alumnos, currículum, asistencia, etc. La pregunta es: ¿Se encuentra dentro del core domain o podemos considerarlo también un subdominio de soporte? La respuesta definitiva nos la tendrían que dar los expertos de dominio.

El tercer elemento es la plataforma de aprendizaje que el colegio quiere construir y utilizar. Este es el que más claramente podríamos identificar como core domain porque afecta directamente al modo en que el colegio quiere desarrollar su actividad principal. Tiene que reflejar el tipo de interacciones de aprendizaje que el centro quiere promover, las metodologías, la forma en que participan e interaccionan los alumnos, la manera de disponer los contenidos, etc.

## Recapitulación

Es importante identificar correctamente los subdominios y el core domain.

Los subdominios genéricos suelen ser comunes a cualquier tipo de organización, al menos dentro del mismo sector, y pueden ser externalizados o resueltos mediante soluciones de terceras partes con una mínima personalización.

Los subdominios de soporte, aún siendo comunes o habituales en otras organizaciones del mismo sector, aparte de ayudar en la consecución de los objetivos del core domain, lo hacen de una manera particular, en sintonía con ese core domain. Tienen que estar bajo nuestro control, aunque puedan externalizarse algunas partes específicas.

Finalmente, el core domain es aquello que hace única a nuestra organización, por lo que requiere de la mayor parte de nuestra atención y esfuerzo tanto de modelado como de implementación.

## Para profundizar

Aparte del libro de Eric Evans, he encontrado estos artículos interesantes acerca de la determinación de los subdominios y de los bounded contexts, los cuales trataremos en próximos artículos.

- [https://vaadin.com/tutorials/ddd/strategic_domain_driven_design](https://vaadin.com/tutorials/ddd/strategic_domain_driven_design)
- [http://gorodinski.com/blog/2013/03/11/the-two-sides-of-domain-driven-design/](http://gorodinski.com/blog/2013/03/11/the-two-sides-of-domain-driven-design/)
- [http://blog.sapiensworks.com/post/2012/04/17/DDD-The-Bounded-Context-Explained.aspx](http://blog.sapiensworks.com/post/2012/04/17/DDD-The-Bounded-Context-Explained.aspx)
- [https://codeburst.io/ddd-strategic-patterns-how-to-define-bounded-contexts-2dc70927976e](https://codeburst.io/ddd-strategic-patterns-how-to-define-bounded-contexts-2dc70927976e)
- [https://medium.com/@naveennegi/thoughts-on-domain-driven-design-in-functional-languages-83c43ec518d](https://medium.com/@naveennegi/thoughts-on-domain-driven-design-in-functional-languages-83c43ec518d)

---

# [El lenguaje ubicuo](https://franiglesias.github.io/ddd-ubiquitous-language/)

> Creo que una de las cosas que provocan que se _haga bola_
 la parte estratégica del DDD tiene que ver con que se trata de una metodología muy orgánica. Es difícil incluso definir un principio y un final del proceso. De hecho, diría que el DDD bien entendido no termina nunca mientras el negocio evoluciona.

Con frecuencia se _despacha_ el DDD con la aplicación de los llamados patrones tácticos, es decir, todo lo que tiene que ver con la implementación del modelo en código. Sin embargo, esto puede llevarnos a lo que se suele llamar _Directory Driven Development_, como si el diseño consistiera en poco más que aplicar arquitectura hexagonal y ya. Y lo mismo se puede decir de otros patrones útiles en torno al DDD, como CQRS, event-driven o incluso Microservicios. DDD no ofrece recetas.

Sin embargo DDD es, ante todo, una conversación.

Comienza con una conversación entre desarrolladores y expertos de dominio acerca de la naturaleza y estructura del dominio, cuyo primer objetivo es delimitar cuál es el problema cuya solución vamos a construir. En esta conversación se van clarificando, separando y organizando los subdominios (de los que hablamos en el artículo anterior), a la vez que construyendo un modelo que los representa. La conversación sigue cuando pasamos del espacio del problema (los subdominios) al espacio de la solución (los contextos acotados), y continúa a medida que avanzamos en el desarrollo del modelo, primero, y la implementación, después. Y en todas estas fases tienen lugar toda suerte de conversaciones, ya sea entre los desarrolladores y los expertos de dominio, o entre desarrolladores de distintos bounded contexts, y todos las personas implicadas en el proyecto.

Y para que esta conversación pueda tener lugar es necesario que se vehicule en un lenguaje que no solo es compartido por todas las implicadas, sino que está presente en todos los niveles de la misma: desde la discusión sobre el dominio hasta su representación en forma de código: en el nombre de los objetos, los servicios, las variables o los tests.

## ¿Documentar el lenguaje ubicuo?

El lenguaje ubicuo se construye a partir del lenguaje que emplean los expertos del dominio. Este lenguaje se compone de términos que designan los conceptos, los procesos y los eventos.

Así que lo primero que nos puede venir a la mente es comenzar recopilando un glosario de estos términos y así documentar el lenguaje para que esté al alcance de todos. Este documento puede crearse a partir de las diversas conversaciones de forma que todos los participantes contribuyan.

Sin embargo, al escribir código nos esforzamos porque sea autodocumentado hasta donde sea posible, porque somos conscientes de que, con el tiempo, la documentación y el código terminarán divergiendo hasta hacerse irreconciliables la una con el otro. Y el caso es que ese mismo riesgo es el que corre la creación de un glosario del lenguaje ubicuo: que con el tiempo el diccionario se estanque mientras el lenguaje crece y evoluciona.

Sabemos que mantener la documentación actualizada a largo plazo es un problema con una solución complicada, porque en un momento dado deja de mantenerse pese a las buenas intenciones iniciales.

[Paul Rayner proporciona algunas pistas en este artículo](http://thepaulrayner.com/blog/2013/05/07/succeeding-with-ddd-documentation/) acerca de cómo tener éxito documentando el proceso de diseño. Además de reconocer el problema, señala algunas cosas interesantes:

Es más importante y productivo documentar como proceso continuo que generar documentación como fin en sí misma. Es decir, documentar mientras se desarrolla, no documentar cuando está el producto terminado.

La documentación debe ser confiable, fácil de cambiar y accesible. En otras palabras: hay que eliminar cualquier dificultad u obstáculo que nos pueda desanimar a la hora de contribuir a esa documentación. El formato concreto es indiferente, ya sea en forma de diagramas hechos a mano, con UML más o menos simplificado, una wiki o documentos markdown en un repositorio.

## Aprendiendo a expresarse en el lenguaje ubicuo

Establecer un lenguaje ubicuo no es un simple proceso de capturar los términos que utilizan los expertos de dominio. En cuanto intentas compilar un lenguaje no formal descubres una increíble cantidad de trampas, sobreentendidos, ambigüedades, polisemias y sinonimias que pueden desafiar la capacidad sistematizadora de cualquiera.

Además, el lenguaje ubicuo tiene _dialectos_. Cada subdominio, cada contexto, puede tener su propia versión un término, que da nombre a un concepto ligeramente distinto porque en cada uno de ellos nos interesa un aspecto particular.

Intentaremos ver esto acudiendo a ejemplos en nuestro proyecto DDD para un colegio.

### Arrancando

Supongamos que nos centramos en el sub dominio académico, en el core domain (o al menos uno de sus dominios de soporte más cercanos). Una forma de empezar puede ser centrarnos en algún evento relevante del dominio. Por ejemplo, un nuevo estudiante se matricula en el colegio. Podemos imaginar una conversación así con la persona responsable de la Secretaría del centro:

– Hablemos entonces del alta de un alumno en el centro.

– De acuerdo. A ese proceso lo llamamos proceso de matrícula o hacer la matrícula.

– Entiendo que se recogen los datos del alumno, sus padres, el curso en el que va a empezar, etc.

– Sí, cada nuevo estudiante aporta sus datos personales. Como es menor, debe ser representado por un adulto. Normalmente serán los padres, pero también hay familias monoparentales o chicos en situación de acogida, con un tutor o tutora legal, etc.

– Vale. O sea que cada nuevo estudiante tiene al menos un representante adulto. ¿Podrían ser más de dos?

– No es habitual, pero realmente tampoco hay nada que lo impida.

Paremos aquí por un momento. Ya han salido varios conceptos y términos:

- Matricula y proceso de matricula.
- Estudiante, también llamado Alumno.
- Representante adulto, típicamente la madre o el padre, pero también un tutor legal.
- Una regla de negocio: todo alumno tiene al menos un representante adulto.

Estamos al principio del proceso por lo que la información va a llegar de forma casi explosiva. Y esa explosión también nos trae los primeros problemas lingüísticos.

- Tenemos dos etiquetas para un mismo concepto: ¿es importante que haya una distinción? ¿Nos está indicando eso algo acerca del dominio?. Es necesario aclararlo.
- El tema del representante adulto: no se le nombra así, pero intuimos que podría haber algún concepto pugnando por salir.

– Vale. Me gustaría profundizar más en esto. Pero antes querría aclarar una cosa: hemos usado la palabra Estudiante y Alumno. ¿Significan lo mismo? ¿Cuál deberíamos usar?

– La verdad es que no hay ninguna diferencia. Habitualmente en los niveles en que trabajamos hablamos de Alumnos, no de Estudiantes.

– De acuerdo, entonces. Volviendo al tema de los padres… Estaba pensando que también se matricularán hermanos, ¿cómo afecta eso a la gestión, las comunicaciones y demás?

– En la medida de lo posible evitamos duplicar comunicaciones generales en el caso de que haya hermanos, sobre todo cuando son en papel, aunque eso ocurre cada vez menos. Pero es verdad que la información de contacto es la misma: domicilio, teléfonos para avisos, etc. Además, las familias prefieren poder acceder a la información de sus hijos de manera unificada.

Otra parada. Ha surgido el concepto Familia que puede ayudar a clarificar el tema de la participación de adultos en el proceso, vamos a tirar un poco de ese hilo:

– ¿Hay alguna consecuencia más del hecho de que haya hermanos en el colegio aparte de lo que se refiere a la comunicación con las familias?

– Aparte de esto, la única cosa relevante ocurre antes de la matrícula, ya que tener hermanos matriculados en el mismo centro da puntos en la solicitud. Pero una vez matriculados la única consecuencia es en este aspecto de la relación con las familias.

– Ya veo. Entonces, ¿podríamos definir Familia como el conjunto de alumnos y sus adultos responsables, por así decir?

– Sí, ciertamente.

– Y que esto tiene importancia en cuanto a la relación con el colegio, la comunicación y el acceso a la información, pero no tiene otros efectos, ¿es correcto?

– Sí, con lo que más tiene que ver con el ámbito de la Secretaría.

Bien. Este último tramo de la conversación nos ha permitido situar las cosas mejor. Hemos hecho salir un concepto que el propio experto de dominio no había expresado inicialmente y que nos permite unificar un asunto que podría salirse fácilmente de control.

Además, nos han señalado unos límites del subdominio de la “Secretaría educativa”: gestión del alumnado y relación con las familias. Pero aún nos falta mucha información, hay un tema que no hemos tratado aún:

– Bien. Hemos dicho que un Alumno se matricula para un curso, como 4º de Primaria o 3º de ESO.

– Sí, correcto.

– ¿Cuándo lo hace?

– En principio la matrícula se hace los primeros días del curso escolar, aunque antes se ha hecho una reserva, si ya era Alumno del centro, o una asignación de plaza cuando es nuevo.

– Vale, pero eso… ¿afecta al proceso en sí?

– Realmente no. Es otro proceso.

– De acuerdo. ¿Cómo se sabe que el alumno se matricula en el curso que le corresponde?

– Bueno. En principio los alumnos se matriculan en un curso en función de su edad, salvo que hayan tenido que repetir alguno o se les haya adelantado en el caso de tener altas capacidades. Esa información consta en su expediente.

– Entiendo que eso está regulado por Ley.

– Sí. Por ejemplo, los alumnos que cumplen 6 años en el año en que comienza el curso escolar se matriculan en Primero de Primaria, los de 7 en Segundo, y así sucesivamente.

Aquí nos surge una situación interesante: un mismo término que estaría refiriéndose a conceptos distintos: curso. Tenemos que aclarar esto, lo que nos traerá un giro inesperado:

– Estamos hablando de curso y curso escolar. Entiendo que son conceptos distintos, ¿no?

– Sí, por un lado hablamos de curso como nivel educativo en que está matriculado un Alumno y, por otro, el curso escolar es el año escolar, que va del 1 de septiembre de un año hasta el 30 de agosto del siguiente.

– ¿Podríamos usar año escolar en lugar de curso escolar para evitar la ambigüedad?

– Sí, aunque es muy habitual usar la otra forma.

– Ya, comprendo. Por otra parte, he observado que el colegio tiene tres grupos por curso, ¿afecta eso al proceso de matrícula?

– En principio, no. Obviamente cada alumno acaba siendo asignado a uno de los tres cursos, pero eso ya es una cuestión de organización pedagógica que se hace una vez matriculado.

– Un momento. Me ha parecido entender que se refería a los tres grupos de un curso como cursos, ¿es eso correcto?

– Hum. Es verdad. Ahora que sale el tema es verdad que muchas veces llamamos curso al grupo. Es habitual que una tutora diga “mi curso” refiriéndose a su grupo de tutoría, o que alguien diga “el curso de 3º B”, por ejemplo.

Tenemos tres significados para una misma palabra y en un mismo contexto, por tanto, tenemos un problema. Para resolverlo, tenemos que ver si es posible utilizar sinónimos que tengan sentido para el dominio:

- Curso escolar: Año escolar
- Curso: Nivel educativo
- Curso: Cada agrupación de alumnos matriculados en el mismo nivel educativo.

Después de aclarar los puntos anteriores, seguimos avanzando en la conversación.

– Una vez que el alumno se ha matriculado, ¿cómo se le identifica dentro del sistema?

– Se le asigna un número de matrícula aunque realmente solo se usa para algunos trámites ya que normalmente identificamos al alumno por su nombre y apellidos o, dentro de un año escolar, por su nivel, grupo y número de clase, que viene dado por su posición en la lista por orden alfabético del apellido.

– De acuerdo. Sin embargo, necesitaremos un identificador único. Entiendo que cabe la posibilidad de que haya alguna coincidencia de nombres y apellidos de vez en cuando.

– Sí, podría pasar, aunque es raro que ocurra dentro del mismo nivel y grupo.

– Ya. Sin embargo, el número de matrícula podría ser suficiente y así el sistema podría ser compatible con los registros que existan ahora.

– De acuerdo.

– Entiendo que es un número que se incrementa con cada matrícula. ¿es necesario ocupar todos los números o pueden quedar huecos de números no usados?

– En principio no hay problema en dejar números sin usar.

Aquí surge otra regla de negocio: cada alumno tiene un identificador único para su vida escolar, pero también se menciona que tiene otros identificadores, como el número de clase que lo identifica dentro de un grupo de alumnos concreto durante un año escolar e, incluso, junto la identificación de su grupo, lo identifica dentro del centro durante ese mismo año escolar. En otras palabras, nos están mencionando distintos contextos.

Se nos aclara que estos identificadores no se usan en el día a día realmente, pero no hay duda de que son necesarios para otorgar identidad a cada alumno individual en el sistema.

– Y con esto, entonces, quedaría terminado el proceso de matrícula.

– Básicamente sí.

De momento, la conversación termina aquí.

## Recopilar la información recogida

El _setup_ ideal para estas conversaciones sería contar con una pizarra blanca o similar en la que poder ir anotando conceptos y dibujando esquemas de forma que todos los participantes puedan aportar. Los acuerdos de significado a los que vamos llegando se registran en forma de documentos que los transcriban, fotos de la pizarra en los momentos en que llegamos a un acuerdo, esquemas que se hayan dibujado, etc.

No es necesario seguir un protocolo formal, sino recoger la información generada de la forma más fidedigna, pero también la más práctica y más a mano para los participantes. Si alguien sabe de UML puede crear diseños simplificados de los conceptos identificados y sus relaciones. Algunos procesos podrán representarse en diagramas de flujo, otros podrían quedar bien resueltos con simples esquemas de cajas y flechas.

También es recomendable incorporar a las notas formularios impresos y plantillas que puedan estar utilizándose para el mismo proceso en la actualidad. Estos documentos nos proporcionan un punto de partida para entender la estructura de los datos que se solicitan o cómo se registran. Podemos hablar sobre cuáles se utilizan realmente, cuáles acaban ignorándose, etc.

Otra forma de iniciar y mantener esta conversación es a través del [Event Storming](https://techbeacon.com/devops/introduction-event-storming-easy-way-achieve-domain-driven-design). Es una técnica en la que se parte de los eventos o sucesos interesantes que suceden en el dominio, buscando todos los conceptos, comandos, servicios, etc, relacionados con lo que se llega a identificar agregados y contextos acotados.

Incluso la metodología Behavior Driven Design nos facilita una forma de hablar sobre el dominio y registrar el lenguaje ubicuo en forma de una documentación dinámica, pero esto quizá nos acerca demasiado a la implementación.

En todo caso, será el tema de otro artículo.

## Más información

Un par de artículos más sobre lenguaje ubicuo:

[https://blog.carbonfive.com/2016/10/04/ubiquitous-language-the-joy-of-naming/](https://blog.carbonfive.com/2016/10/04/ubiquitous-language-the-joy-of-naming/)

[https://arne-mertz.de/2017/07/ubiquitous-language/](https://arne-mertz.de/2017/07/ubiquitous-language/)

---

# [DDD no es lo que te han contado](https://franiglesias.github.io/ddd-is-not-what-they-said/)

> Acabo de leerme, por fin, El Libro sobre _Domain Driven Design_
 de Eric Evans. Después de mucho tiempo leyendo de segunda mano sobre el tema, mi conclusión es: DDD no es exactamente lo que nos han contado.

Por alguna razón las referencias que había escuchado sobre el libro eran bastante negativas. Y es una pena porque he ido posponiendo su lectura más de lo debido, aunque también es verdad que he priorizado otros que tenía más accesibles. Pero, en resumen: el libro me ha gustado mucho y el contenido me ha ayudado a darle sentido a muchas cosas.

En parte puedo entender la visión “negativa” del libro, ya que no es un libro técnico al uso.

Empecemos por ahí, El _Blue Book del DDD_ no es un libro que te proporcione recetas sobre cómo implementar cosas, sino que trata principalmente acerca de cómo pensar y cómo conversar acerca dominio, el espacio del problema, y cómo iniciar y desarrollar el diseño del modelo como espacio de solución. Con todo, es un libro eminentemente práctico que ofrece una sistemática de análisis y de diseño, estructurada en temas para los cuales ofrece una completo abanico de estrategias y patrones de actuación.

Como reza el subtítulo, Domain Driven Design trata sobre _cómo afrontar la complejidad en el corazón del software_.

A lo largo del libro encontrarás muy poco código, apenas unos ejemplos para ilustrar algún punto, pero sí muchos diagramas para ver cómo se plantean los modelos para los distintos casos y cómo pueden evolucionar ante los cambios del dominio. También una buena muestra de conversaciones y mucho contexto para entender los ejemplos que ilustran los distintos temas.

Porque el DDD trata principalmente de eso: de conversaciones, contextos y modelos para entender el problema del dominio y poder desarrollar una solución.

Sin embargo, es frecuente que se hable de DDD asociado a tecnologías o implementaciones concretas: microservicios, CQRS, Event Sourcing o Arquitectura Hexagonal, como si uno implicase las otras. También es frecuente que se insista en los patrones tácticos, los que tienen que ver con la implementación, y no los estratégicos, los que tienen que ver con el diseño, que es de lo que trata el _Blue Book_.

Así, diría que es más popular y conocido el _Red Book_, [Implementing Domain Driven Design](https://www.amazon.es/Implementing-Domain-Driven-Design-Vaughn-Vernon/dp/0321834577), el libro de Vaughn Vernon. Entiendo que es porque es el que explica el cómo hacer, o sea, el que trata de los patrones tácticos. El libro de recetas, por decirlo así. Yo no me he leído el _Red Book_ entero, sino que de momento lo he estado usando como referencia para temas específicos, pero mucha gente se ha leído este y no el de Evans y creo que es un error.

Además, posteriormente el mismo Vaughn Vernon publicó el _Green Book_, [Domain Driven Design Distilled](https://www.amazon.es/Domain-Driven-Design-Distilled-Vaughn-Vernon/dp/0134434420) que hay quien considera como un buen libro de introducción al DDD. En este discrepo por completo. El _Distilled_ es un buen índice para el _Red Book_, pero no aprenderás Domain Driven Design con él.

Posiblemente sea mejor introducción [Domain-Driven Design Reference: Definitions and Pattern Summaries](https://www.amazon.es/Domain-Driven-Design-Reference-Definitions-Summaries/dp/1457501198/) de Evans, de cuya existencia me acabo de enterar porque soy así de espabilado.

La insistencia en los patrones tácticos tiene su aspecto positivo. No dejan de ser buenas prácticas, pero no son exclusivas de DDD, ni mucho menos. De hecho, una de las confusiones más típicas del DDD es con la Arquitectura Hexagonal. Lo que se suele conocer humorísticamente como “directory driven development”.

## DDD y Arquitectura Hexagonal

La Arquitectura Hexagonal y, en general, las arquitecturas limpias son perfectamente compatibles con Domain Driven Design. Esto es así por dos razones:

- La separación en capas, con el dominio en el centro.
- Las reglas de dependencia, con todas las dependencias apuntando hacia adentro, de modo que dominio no tiene dependencias.

Pero no hay nada que diga que la Arquitectura Hexagonal ES la arquitectura propia de una aplicación DDD. De hecho, la Arquitectura Hexagonal es una aplicación del principio de Inversión de Dependencias y sus prácticas asociadas, más que una arquitectura.

En cualquier caso, estructuramos el código en torno a tres carpetas/capas: dominio, aplicación e infraestructura. Pero esto por sí solo no es DDD.

La razón de que esto funcione es que el DDD pide que el modelo de dominio esté representado en código mediante objetos puros del lenguaje, completamente aislados de cualquier detalle técnico como la persistencia. Es decir los llamados _building blocks_ (las entidades, los value objects, los domain events o los servicios) no pueden diseñarse pensando en cómo van a ser persistidos o comunicados en una API, por poner un ejemplo, sino que deben diseñarse como si siempre estuviesen viviendo en memoria.

Pero, como resulta bastante evidente, necesitamos un mecanismo de persistencia de las entidades y agregados, aunque solo sea por pura necesidad técnica.

## DDD y Bases de Datos

El concepto que se encarga de la persistencia en DDD es el de Repositorio, un lugar en el que guardar u obtener entidades (voy a hablar de entidades y agregados indistintamente). Desde el punto de vista del dominio un repositorio es un simple almacenamiento en memoria con el cual puedo:

- Guardar entidades
- Obtener entidades conociendo su identidad
- Obtener subconjuntos de entidades que cumplan una especificación

En un mundo ideal, la interfaz de un Repositorio solo tiene tres métodos:

- store
- retrieve(id)
- findSatisfying(specification)

Para poder hacer el repositorio independiente de implementación aplicamos el principio de inversión de dependencias y en la capa de dominio tenemos una RepositoryInterface, de modo que no nos acoplamos a la tecnología concreta de persistencia.

Esto sigue siendo una buena práctica en términos generales que no es exclusiva del DDD. Lo que sí es propio es el concepto de repositorio como un simple almacén en memoria de entidades y que no puede contener ninguna regla o invariante de negocio. La idea de que un repositorio tenga métodos que son, de hecho, reglas de negocio es una mala práctica.

En este sentido el patrón Specification, definido por [Fowler y Evans](http://www.martinfowler.com/apsupp/spec.pdf), es el camino a seguir.

### DDD y noSQL

Siempre desde el punto de vista del dominio, las bases de datos noSQL serían las que mejor encajan en el concepto de Repositorio. Al fin y al cabo, es relativamente fácil serializar una entidad, por compleja que sea, para persistirla como documento y recuperarla a partir de una clave.

Sí, hay un montón de objeciones técnicas y válidas a esta afirmación: rendimiento, estabilidad, problemas para la gestión de relaciones y un largo etcétera.

### DDD y ORM

En cambio, los ORM, desde el mismo punto de vista, son poco adecuados. Las entidades en el contexto del ORM no son entidades DDD. [Matthias Noback recomienda olvidarse del ORM cuando estemos diseñando entidades](https://matthiasnoback.nl/2018/06/doctrine-orm-and-ddd-aggregates/) porque el funcionamiento del ORM va a interferir con su adecuado diseño. En su lugar, lo que recomienda es simplificar, incluso usando variables privadas en las entidades únicamente con la finalidad de su almacenamiento en la BD e incluso reconstruyendo objetos al vuelo a partir de primitivas.

Todo para mantener Entidades DDD ricas y evitar los modelos anémicos característicos de los enfoques Database-first.

## DDD y CQRS

[Command Query Responsibility Segregation](https://martinfowler.com/bliki/CQRS.html) es otro de esos patrones que se asocia habitualmente con DDD. En parte porque también se relaciona estrechamente con _Event Sourcing_, cosa que tiene sentido.

Simplificando mucho CQRS consiste en la aplicación _extrema_ del CQS (Command Query Segregation) a los modelos de datos.

CQS dice que una función o método ha de ser un comando, que provoca un efecto en el sistema sin obtener información del mismo, o una query, que recupera información de un sistema sin provocar ningún efecto en él.

En la interfaz de repositorio de la que hablamos más arriba, se asume que solo existe un objeto repositorio que gestiona tanto la lectura como la escritura, aunque cada método tiene su propia responsabilidad.

Sin embargo, CQRS separa las operaciones de lectura y escritura en modelos diferentes. De modo que pueden implementarse incluso con tecnologías distintas. Su “reverso tenebroso” es la consistencia: ¿cómo mantenemos consistente la información que se escribe con la que se lee?

Para esto se introduce la noción de _consistencia eventual_ que, explicado muy básicamente, significa que la información acabará por ser consistente en algún momento próximo, cuando el proceso que actualiza los modelos de lectura haya podido realizar todos los cambios recogidos por el modelo de escritura.

Una de las estrategias para lograr esto es justamente mediante eventos que indiquen los cambios, de modo que los modelos de lectura se actualizan como respuesta a esos eventos.

Pero, CQRS es fundamentalmente un patrón de implementación de una solución de persistencia que puede funcionar bien para entornos que requieren un alto rendimiento, y que añade una complejidad excesiva para una gran mayoría de aplicaciones. De nuevo, no hay nada en DDD que implique que CQRS sea un patrón propio, aunque una aplicación diseñada con metodología DDD puede implementar una solución de persistencia CQRS.

## Event Sourcing y DDD

_Event sourcing_ es un enfoque que también se asocia frecuentemente a CQRS y a DDD. Debe ser porque en DDD hablamos de los Eventos de Dominio y un evento es un evento es un evento…. Sí, tiene sentido esta asociación, pero Event Sourcing no es para todo el mundo y tampoco es un patrón que derive del enfoque DDD. Sencillamente, lo que ocurre es que encaja bien.

La idea del _Event Sourcing_ es la siguiente: En un sistema de software tradicional el estado del sistema, o concretamente el estado de una entidad, como podría ser el expediente académico de un alumno, nos habla de su estado en el momento de la consulta, pero no nos dice nada de cómo se ha llegado hasta él, o si ha habido algún tipo de cambio en un período de tiempo. Tradicionalmente, si nos interesa guardar la historia, como en este caso, se guarda explícitamente. Por eso el expediente académico guarda registros de todos los cursos por los que ha pasado el alumno.

Ahora bien, ¿qué pasa si en vez de “actualizar” el estado cada vez que hay un cambio, recogemos todos los eventos que cambian ese mismo estado? Pues pasa que podemos tener _Event Sourcing_.

Si tenemos la historia de eventos, podemos recorrerla, rebobinarla, volver a empezar desde cero y obtener el estado del sistema en un momento dado. Si es necesario, podemos recoger “instantáneas” de ese estado en momentos determinados de la historia. En lugar de tener bases de datos almacenando en tablas la información de cada entidad, podemos tener un almacén de eventos y generar proyecciones específicos de cada vista que deseemos ofrecer en nuestra aplicación, ya sea visual, API o la que sea.

Sería muy prolijo hablar de la cantidad de cosas que se pueden hacer con Event Sourcing, pero imagina poder rebobinar tu sistema a una fecha determinada o que cambies lo que cambies en el software tus datos se pueden adaptar automágicamente porque, en realidad, nunca los guardas en el sentido tradicional.

Pero de nuevo, no hay nada en DDD que diga que _Event Sourcing_ es un patrón preferido o especialmente adecuado. Es una opción tecnológica más a tu disposición para implementar tu aplicación. La salvedad es que el concepto de los Domain Events hace que _Event Sourcing_ encaje fácilmente.

## Microservicios y DDD

Si ha habido un término estrella en los últimos tiempos ha sido el de Microservicios que, ni es nuevo, ni está implícito, ni es particularmente adecuado para DDD.

De nuevo, Microservicios es una forma de implementar una solución tecnológica que, en este caso, consiste en crear aplicaciones (servicios) muy especializadas, comunicándose a través de protocolos estandarizados (típicamente API REST), con los que componer el sistema de software.

Los microservicios han estado de moda en los últimos años. Al principio, como el paradigma al que todo el mundo parecía querer apuntarse. Sin embargo, las dificultades en su implementación han llevado a que últimamente esté de moda hablar de cómo los microservicios no han funcionado en muchísimas ocasiones.

Parte de la asociación de DDD con microservicios puede venir como posible implementación para _bounded contexts_ y el _context map_. En mi opinión es una de esas relaciones bastante traídas por los pelos, lo que posiblemente explique gran parte de las problemáticas encontradas en la implementación de proyectos basados en microservicios.

Más bien pienso que los microservicios podrían surgir de forma natural al identificar módulos muy especializados y extraíbles en un monolito, pero a través de un proceso iterativo y evolutivo.

## Para terminar

No es DDD si el énfasis está en la implementación y los patrones tácticos. Esa es solo una parte de todo lo que DDD supone. La palabra clave en DDD es la última D: Design.

DDD trata de conversaciones, conceptos y modelos. En definitiva, trata de diseño.

---

# [De directory driven a DDD paso a paso](https://franiglesias.github.io/the-way-to-ddd-1/)

> Muchas bases de código que han sido creadas tratando de seguir la metodología DDD se quedan atascadas en ese falso DDD que solemos llamar “directory driven development”, que es básicamente una aplicación de la Arquitectura Hexagonal. Esto es, utilizan la típica distribución de carpetas Domain, Application e Infrastructure, pero el código en ellas está mal organizado y mal distribuido porque en su día no se tenía una comprensión completa de lo que implica Domain Driven Design.

El resultado es que en lugar de la proverbial _Big Ball of Mud_ ahora tenemos tres _Big Balls of Mud_ que, en parte, es algo mejor pero no suficiente.

Así que en este artículo vamos a ver cómo montar un plan que nos permita evolucionar el código a una organización más próxima a DDD que nos facilite llegar a un diseño más sólido, con el que identificar mejor elementos complejos como agregados, módulos o _bounded contexts_.

El punto de partida de esta propuesta es la idea de que el código es una representación ejecutable del conocimiento que tenemos sobre el dominio o, también, una expresión del modelo del dominio que tiene el equipo. Por tanto, se trata de un refactor orientado a tener un mejor conocimiento del dominio, un modelo más articulado del mismo, lo que nos puede llevar a un punto en el que estemos en disposición de profundizar en la flexibilidad del diseño, los _bounded contexts_ como representación de los subdominios y otras muchas mejoras.

Y estas son sus etapas:

En la capa de dominio

- Eliminar menciones técnicas en la capa de dominio y organizarla en torno a conceptos
- Identificar agregados y encapsular en el Aggregate Root
- Empujar comportamiento hacia el dominio
- Eliminar lógica de dominio de los repositorios

En la capa de Aplicación

- Eliminar menciones técnicas en la capa de aplicación
- Organizar la capa de aplicación en torno a use cases
- Mover servicios a la capa de dominio si representan reglas de negocio
- Haz que se lancen eventos anunciando cualquier cosa interesante que haya en el dominio
- Introduce un Command Bus si no lo tienes ya
- Introduce un Event Bus si no lo tienes ya
- Analiza los Use cases y deja en ellos solo el código que ejecuta su intent y mueve la lógica a subscribers

En la capa de infraestructura

- Organizar la capa de infraestructura en torno a conceptos

Siendo realistas, esto va a ser una gran cantidad de trabajo y es una locura intentar hacerlo de una tacada, incluso las fases sueltas. Lo más eficaz es aplicarlo de forma iterativa en cada tarea.

Es decir, en una tarea que tengamos, aplicamos la primera etapa en las clases de dominio que tengamos que tocar. Pero solo la primera. La próxima vez que “pasemos por ahí” aplicamos la siguiente etapa y así sucesivamente.

Intentar refactorizar a DDD como proyecto nunca funciona.

Por otro lado, no existe una receta mágica. Una forma perfecta de organizar un dominio válida para todos los casos. Olvídalo. El código debería reflejar el modelo mental que tenemos, como equipo o empresa, del dominio en el que estamos trabajando o, en todo caso, del subdominio específico que nosotros estamos modelando.

Esencialmente se trata de elaborar el mapa conceptual del dominio o subdominio y eso es un proceso evolutivo y conversacional. La conversación tiene lugar entre los expertos del dominio y el equipo de desarrollo, conversación de la que debe concretarse el lenguaje ubicuo, pero el código va a ser el resultado de expresar el modelo mental que el equipo de desarrollo se forme.

## Eliminar menciones técnicas en la capa de dominio y organizarla en torno a conceptos

Muchas veces la carpeta de dominio tiene una forma como esta:

> Domain
  + Academic
      +- Entity
      |   +- Student.php
      +   +- Teacher.php
      +- VO
      |   +- StudenId.php
      |   +- TeacherId.php
      |   +- Level.php
      |   +- Stage.php
      |   +- Subject.php
      +- Event
      |   +- StudentWasEnrolledInCourse.php
 ...

Te haces a la idea, ¿no? La carpeta de dominio se organiza en base a _technicalities_. Tal vez el primer nivel intente establecer unos conceptos o módulos, pero luego vemos que toman el control aspectos puramente técnicos.

Cuando hablamos del lenguaje ubicuo hablamos de lo que debe incluir, pero no hablamos de lo que no debería incluir, y esta taxonomía de elementos técnicos no debe estar presente en él.

Un problema que puede surgir es intentar organizar prematuramente cosas como los módulos o los agregados. Es mejor ir paso a paso.

Así, el primer paso es intentar estructurar todo en conceptos que, en principio, serán concretos y de este jaez:

> Domain
  + Academic
      +- Student
      |   +- Student.php
      |   +- StudentId.php
      |   +- StudentRepository.php
      |   +- StudentWasEnrolledInCourse.php
      +- Teacher
      |   +- Teacher.php
      |   +- TeacherId.php
      |   +- TeacherRepository.php
      +- Course
      |   +- Level.php
      |   +- Stage.php
      |   +- Subject.php
 ...

Posibles excepciones: puesto que seguramente cada concepto esté involucrado en una buena cantidad de eventos podría ser adecuado tener una carpeta para los eventos. Sin embargo, es posible que también podamos estructurar las carpetas a partir de Servicios, es decir, clases que ejecutan comportamientos que no se pueden atribuir claramente a una entidad o agregado. Algo así:

> Domain
  + Academic
      +- Student
      |   +- Student.php
      |   +- StudentId.php
      |   +- StudentRepository.php
      |   +- EnrollStudentInCourse
      |   |   +- EnrollStudentInCourse.php
      |   |   +- StudentWasEnrolledInCourse.php
      +- Teacher
      |   +- Teacher.php
      |   +- TeacherId.php
      |   +- TeacherRepository.php
      +- Course
      |   +- Level.php
      |   +- Stage.php
      |   +- Subject.php
      |   +- AssessSubject
      |   |   +- AssessSubject.php
      |   |   +- StudentWasAssessedInSubject.php
      |   +- AssessmentPeriod.php
 ...

O así, si lo prefieres:

> Domain
  + Academic
      +- Student
      |   +- Student.php
      |   +- StudentId.php
      |   +- StudentRepository.php
      +- Teacher
      |   +- Teacher.php
      |   +- TeacherId.php
      |   +- TeacherRepository.php
      +- Course
      |   +- Level.php
      |   +- Stage.php
      |   +- Subject.php
      |   +- AssessmentPeriod.php
      +- EnrollStudentInCourse
      |   +- EnrollStudentInCourse.php
      |   +- StudentWasEnrolledInCourse.php
      +- AssessSubject
      |   +- AssessSubject.php
      |   +- StudentWasAssessedInSubject.php
 ...

Como puede apreciarse la impresión de orden aumenta mucho. Cognitivamente hablando, la nueva estructura revela un mayor sentido, ya que las clases más estrechamente relacionadas van juntas. En la práctica, a la hora de hacer cambios, la carga es menor porque posiblemente tendremos que tocar archivos que están en la misma carpeta, sin necesidad de buscar en varias.

Otra ventaja a largo plazo es que nos facilita partir el dominio si en algún momento vemos esa necesidad por el desarrollo de nuestra comprensión del negocio, por ejemplo para detectar módulos o incluso bounded contexts.

Por otra parte, es muy posible que en tu carpeta de dominio no tengas todavía servicios, algo que ya solucionaremos en fases posteriores.

## Identificar agregados y encapsular en el Aggregate Root

Los conceptos de Entidad y Value Object suelen estar bastante claros. Por contra, los agregados pueden ser un poco más complejos de identificar.

Un agregado es una forma de agrupar dentro de un límite de consistencia un conjunto de entidades, y value objects, que mantienen entre ellas una invariante. Dicho así suena muy pedante, así que vamos a intentar explicarlo.

### Límites de consistencia

La idea de consistencia posiblemente ya la tienes. Cuando instancias un objeto procuras que sea consistente, o sea, que contenga toda la información que necesita para crearse y que esta tenga sentido conforme a las reglas de negocio. Por ejemplo, un email ha de ser un email válido, o una persona ha de tener nombre y al menos un apellido.

Cuando son varias las entidades participando debe darse una consistencia entre ellas.

Por ejemplo, en una aplicación para organizar viajes, un viaje tiene que tener diversas etapas, como mínimo origen y destino y puede tener cero o más etapas intermedias. Esto se puede representar así:

```php
class Journey
{
    public function __construct(Place $origin, Place $destination)
    {
    //...
    }
    
    public function addStage(Place $stage)
    {
    //...
    }
}
```

Siendo Journey la raíz del agregado, es la que define los límites en que debe mantenerse una consistencia con las otras entidades participantes.

### Invariantes

Si piensas en el concepto de Validación ya tienes una idea aproximada de lo que sería una invariante. En cierto sentido es una validación que se aplica a la relación entre entidades.

En el ejemplo anterior las invariantes son:

- El viaje ha de tener origen y destino (no se puede crear un viaje sin las dos)
- El viaje puede tener cero o más etapas intermedias (se pueden añadir opcionalmente)

Imagina que para los efectos de esta aplicación una regla de negocio es que solo pueden hacerse un total de 5 etapas en un viaje. Ahora las invariantes son:

- El viaje ha de tener origen y destino.
- El viaje puede tener entre 0 y 3 etapas intermedias.

En código sería algo así:

```php
class Journey
{
    public function __construct(Place $origin, Place $destination)
    {
        $this->stages->append($origin);
        $this->stages->append($destination);
    }
    
    public function addStage(Place $stage)
    {
        if ($this->stages->count() >= 5) {
            throw new AddingStageException('Only 5 stages allowed');
        }
        
        $this->stages->appendAfterOrigin($stage);
    }
}
```

### Ejemplo

Una buena parte de los elementos del DDD pueden ser vistos como límites de consistencia, en los que las reglas de negocio e invariantes deben protegerse a distinta escala, por así decir:

- Value objects y Entidades
- Agregados
- Bounded Context

Así que veamos un ejemplo para explicarlo bien.

Supongamos que nuestro sistema de gestión escolar tiene en un concepto que representa las tareas de aprendizaje que se asignan a cada estudiante y que hemos decidido representar con una entity llamada Task. Para poder llevar un seguimiento, estas tareas tienen un estado, que representaremos mediante un Value Object llamado TaskStatus.

El estado de Task cambia en respuesta a ciertas acciones:

- Task::assign. Cuando se asigna la tarea a un estudiante, se inicia su estado al valor `to do`.
- Task::start.. Cuando el estudiante comienza a trabajar en ella, el estado cambia a `wip`.
- Task::deliver. Una vez que la considera lista, el estado pasa a `review`.
- Task::assess. Cuando es calificada el nuevo estado será `assessed`.
- Task::return. Opcionalmente, el profesor puede devolver la tarea con comentarios para que el estudiante la corrija o la mejore. En ese caso, vuelve al estado `wip`.

Cualquier otro cambio de estado no es válido.

Así que tenemos unas transiciones (start, deliver, assess y return) que cambian los estados de Task (assign, realmente es como un constructor, así que no la vamos a considerar una transición).

Al modelar parece claro que TaskStatus es un ValueObject que “vive” dentro de Task:

```php
class Task
{
    /** @var TaskStatus */
    private $status;
    /** @var StudentId
    private $student;
    
    private function __construct(Student $student)
    {
        $this->status = TaskStatus::toDo();
        $this->student = $student->id();
    }
    
    public static function assign(Student $student): self
    {
        $task = new self($student)
        
        return $task;
    }
}
```

Y también parece claro que Task tiene que ocuparse de mantener su estado. Por ejemplo:

```php
class Task
{
    // ...
    
    public function deliver(): void
    {
        if (!$this->status->isWip()) {
            throw new InvalidTaskTransition('You cannot deliver a task not in wip status');
        }
        $this->status = TaskStatus::review();
    }
}
```

Es decir, una tarea puede entregarse si estaba en estado WIP, pero no en otro estado. El límite de consistencia es la entidad Task, de modo que otras entidades y servicios que no sean Task pueden confiar en que el status de Task será siempre consistente hagan lo que hagan con ella. La invariante que mantienen entre Task y TaskStatus es justamente esa. Y viéndolo desde fuera esta solución parece muy correcta, ya que la entidad se encarga de proteger sus invariantes.

Otra forma de considerar el problema es situarnos dentro de la entidad Task (que bien podría ser un agregado en un cierto contexto) y ver cuáles son los límites de consistencia de TaskStatus. Es cierto que Task se ocupa de mantener su status, pero eso no quiere decir el concepto que representa TaskStatus tenga que ser pasivo. Antes bien, tiene mucho sentido que sea el responsable de mantener sus invariantes en lo que respecta a esos cambios de estado. El ejemplo anterior, planteado de esta manera quedaría:

```php
class TaskStatus
{
    public function deliver(): self
    {
        if (!$this->isWip()) {
            throw new InvalidTaskStatusTransition('You cannot deliver a task not in wip status');
        }
        return TaskStatus::review();
    }
}

class Task
{
    // ...
    
    public function deliver(): void
    {
        $this->status = $this->status->deliver();
    }
}
```

Ahora TaskStatus se encarga de mantener sus invariantes como concepto, encargándose de gestionar las transiciones y liberando a Task de conocer los detalles. De este modo Task puede confiar en TaskStatus, mientras se ocupa de mantener todas las otras invariantes que le conciernan.

TaskStatus define su propio límite de consistencia, colaborando para que Task defina el suyo.

### Volviendo a los agregados

En código, un agregado es una entidad que contiene otras entidades y value objects. Las entidades contenidas en el agregado, a excepción de la entidad raíz, se caracterizan porque no tienen “vida” fuera del mismo, de tal modo que todo acceso ha de pasar por él.

Veámoslo con un ejemplo. Supongamos que en nuestro modelo tenemos la entidad Assessment para modelar las calificaciones que recibe un Student en las diversas actividades educativas. Pues bien, Assessment no tiene una existencia independiente de Student (no se pone una nota a “nadie”, para entendernos, siempre se le pone a una estudiante), Por tanto, Assessment es una entidad agregada en Student, que es la raíz del agregado en ese contexto.

La implicación de esto es que Student tiene que tener métodos para añadir Assessments, por ejemplo, pero encargándose también de instanciarlos. Esquemáticamente podría ser algo así:

```php
class Student
{
    /** AssessmentsCollection */
    private $assessments;
    
    //...
    
    public function assess(DateTime $date, Subject $subject, Calification $calification) {
        $assessment = new Assessmemt($date, $subject, $calification);
        $this->assessment->append($assessment);
        
        $this->raise(new StudentWasAssessed($this->id, $assessment));
    }
    
    //...
}
```

Es decir: el `new` no se hace fuera de Student, sino dentro.

La razón de esto es que el agregado tiene la función de proteger las invariantes y aplicar las reglas de negocio necesarias porque estamos trabajando dentro de los límites de consistencia del agregado.

Pero, por otro lado, Las operaciones sobre un agregado deberán estar dentro de sus límites transaccionales.

Una implicación es que no se tienen repositorios para estas entidades agregadas y que su persistencia debe gestionarla el agregado. Este es lo que se quiere decir cuando se afirma que los repositorios guardan y recuperan agregados. Por tanto, los repositorios tienen que “saber” como montar un agregado con todas sus entidades.

### Límites transaccionales

Del mismo modo que los agregados (pero también las entidades y los value objects) definen límites de consistencia. Los agregados definen límites transaccionales, esto es: los cambios en un agregado se persisten en una transacción, de modo que si falla la persistencia de alguno de los elementos, falla la persistencia del agregado. Y todo eso para garantizar que se mantienen las invariantes.

Veámoslo con algunos ejemplos sencillos a partir del típico agregado Order con sus correspondientes Lines:

- Si fallase la persistencia de una sola Line, pero se mantuviesen las demás, el cliente no recibiría todos los productos que desea.
- ¿Y si fallasen todas las Lines? Nos podrían quedar Orders sin contenido.
- O tal vez, un pedido que se pasa a cobro sin entregar nada…

Y, de momento, basta por hoy, En un próximo artículo intentaremos seguir con agregados y con la lógica de los repositorios.

## Algo para leer

- [Como elegir aggregate roots](http://blog.koalite.com/2015/03/como-elegir-aggregate-roots/)
- [The aggregate](https://lostechies.com/gabrielschenker/2015/05/25/ddd-the-aggregate/)
- [How to Design & Persist Aggregates](https://khalilstemmler.com/articles/typescript-domain-driven-design/aggregate-design-persistence/)

---

# [Del repositorio al dominio](https://franiglesias.github.io/the-way-to-ddd-2/)

> Un caso particular de conocimiento del dominio que suele estar fuera de sitio es el de los repositorios que lo implementan en Infraestructura. En este artículo veremos cómo poner cada cosa en su lugar.

## Mover reglas de negocio de los repositorios a dominio

Cuando un repositorio tiene un método que selecciona una colección de entidades o agregados es muy posible que estemos violando la separación de capas, ejerciendo reglas de negocio en una implementación de infraestructura.

Posiblemente los repositorios sean uno de los componentes de la aplicación peor entendidos y por los que más fácilmente se rompe con DDD.

Vistos desde el dominio, los repositorios son colecciones en memoria de entidades o agregados. Sin embargo, en la práctica y con muchísima frecuencia usamos los repositorios como medio de acceso a la base de datos, cosa que no es correcta en DDD.

¿Y para que usamos estos accesos? Pues, aparte de obtener colecciones de entidades o de agregados, es habitual que los usemos para obtener agregaciones (no en el sentido de agregados de dominio, sino de resumen de datos, estadísticos básicos, etc.) y proyecciones de los datos almacenados, bien para proporcionar información a servicios, bien para mostrarlas en algún tipo de vista, entendiendo esta en un sentido amplio: una parte de la interfaz gráfica, para un _report_, resultado para un _endpoint_ de una API, etc. Este segundo tipo de uso no es propio de los repositorios y lo trataremos en otro lugar.

En DDD los repositorios se usan para:

- Persistir entidades o agregados
- Recuperar entidades o agregados conocidas su identidades
- Recuperar colecciones de entidades que cumplan alguna condición

Para trabajar con un repositorio, que por definición se debe comportar como un almacén en memoria, no podemos presuponer que hay un sistema de base de datos detrás al que podemos interrogar en un lenguaje propio, como SQL, sino que necesitamos alguna abstracción que nos permita definir reglas de dominio, aunque luego las implementemos de la forma que corresponde al mecanismo de persistencia.

Para guardar o recuperar Entidades concretas no necesitamos mucha parafernalia.

Por su parte, los criterios para seleccionar entidades suelen representar reglas de negocio como, por ejemplo, podrían ser recibos impagados, alumnos que promocionan, tareas pendientes de ser evaluadas, y un largo etcétera. Son condiciones de negocio que implementamos examinando las entidades para ver si las cumplen.

Eric Evans propone el patrón [Specification](https://martinfowler.com/apsupp/spec.pdf) para encapsular las condiciones mediante las que se escogen loa agregados. También lo [hemos tratado anteriormente](https://franiglesias.github.io/patron-specification-del-dominio-a-la-infraestructura-1/), aunque seguramente ese artículo necesite una revisión después de tanto tiempo.

La idea de Specification es separar los criterios de la selección del objeto que hace la selección:

> A valuable approach to these problems is to separate the statement of what kind of objects can be selected from the object that does the selection.

De este modo, una Specification encapsula las condiciones que debe cumplir un objeto para ser seleccionado. En la capa de Dominio el resultado de una Specification es un `boolean`: o bien el objeto cumple los criterios o bien no los cumple. La interfaz suele ser esta:

```php
interface StudentSpecification
{
    public function isSatisfiedBy(Object $candidate): bool;
}
```

Una implementación podría ser (muy simplificado todo, para que sea más claro):

```php
class IsEnrolledInPotionsClass implements StudentSpecification
{
    public function isSatisfiedBy(Student $candidate): bool
    {
        return $candidate->isEnrolledInClass('potions-class-id');
    }
}
```

Otra implementación posible es parametrizada:

```php
class IsEnrolledInClass implements StudentSpecification
{
    private ClassId $classId;
    
    public function __construct(ClassId $classId)
    {
        $this->classId = $classId;
    }
    
    public function isSatisfiedBy(Student $candidate): bool
    {
        return $candidate->isEnrolledInClass($this->classId);
    }
}
```

Para no complicar el ejemplo he puesto una condición sencilla, pero podría ser una combinación de condiciones. Incluso es fácil crear _Composite Specifications_ que nos permitan combinar condiciones sencillas mediante operaciones lógicas (AND, OR, NOT), pero ahora no entraré en ese detalle.

El problema del patrón Specification es que tiene dos caras porque en la práctica tiene que resolver dos problemas distintos:

- En la capa de dominio una Specification puede usarse tanto para recorrer una colección de Entidades y filtrar aquellas que la satisfacen, como para validaciones o comprobaciones de todo tipo.

```php
    //...
    $isEnrolledInClass = new IsEnrolledInClass($classId);
    
    if (! $isEnrolledInClass->isSatisfiedBy($student)) {
        throw new StudentNotInClassException($student);
    }
    //...
    $isEnrolledInClass = new IsEnrolledInClass($classId);
    
    $studentsCollection->filterBy($isEnrolledInclass);
```

- En la capa de infraestructura, sin embargo, es habitual que necesitemos que la Specification se traduzca a una cláusula WHERE o equivalente para la solución de persistencia específica. Esto es debido a que es inviable cargar en memoria todas las entidades o agregados para filtrarlos. Es mucho más eficiente hacer una query que devuelta los datos necesarios.

En el _Blue Book_ Evans discute un par de soluciones:

En la más sencilla de ellas, la _Specification_ contendría un método capaz de devolver la SQL que el repositorio debería usar para obtener los datos deseados. Algo más o menos como esto (en una implementación real, podríamos devolver un objeto Query de Doctrine DBAL o cualquier otra implementación que sea adecuada):

```php
class IsEnrolledInClass implements StudentSpecification
{
    private ClassId $classId;
    
    public function __construct(ClassId $classId)
    {
        $this->classId = $classId;
    }
    
    public function isSatisfiedBy(Student $candidate): bool
    {
        return $candidate->isEnrolledInClass($this->classId);
    }
    
    public function asSql(): string
    {
        return 'SELECT * 
        FROM students s 
        LEFT JOIN students_classses sc ON sc.student_id = s.id 
        WHERE sc.class_id = {$this->classId}';
    }
}
```

Esta _Specification_ podría utilizarse así en el Repositorio (el ejemplo es una especie de pseudocódigo, no es una implementación real, pero debería darte la idea):

```php
class SQLStudentRepository implements StudentRepository
{
    // ...
    public function findSatisfying(StudentSpecification $specification)
    {
        $sql = $specification->asSql();
        
        $result = $this->execSql($sql);
        
        return $this->mapResultToAggregates($result);
    }
}
```

El inconveniente principal es que estamos introduciendo detalles de implementación en la capa de dominio. La _query_ depende la implementación concreta del sistema de base de datos, que típicamente puede ser un RDBMS como Postgre o Mysql (lo que supone la posibilidad de que tengamos que lidiar con dialectos concretos de SQL), pero que puede ser cualquier paradigma de persistencia.

Para evitar eso, hay varias posibilidades.

Una de ellas es que la implementación del repositorio se encargue de mantener las distintas _queries_ que sean necesarias, utilizando un Double dispatch para que sea la Specification la que controle cómo se usa. Es algo más o menos como sigue:

El repositorio contiene métodos para obtener los resultados deseados (aquí obvio todo el tema de mapeado y construcción de entidades para facilitar la comprensión):

```php
class SQLStudentRepository implements StudentRepository
{
    // ...
    public function findSatisfying(StudentSpecification $specification)
    {       
        return $specification->findSatisfyingFrom($this);
    }
    
    public function findStudentsInClass(ClassId $classId): string
    {
        $sql = $this->findStudentsInClassSQL();
        
        $result = $this->execSql($sql);
        
        return $this->mapResultToAggregates($result);
    }
    
    private function findStudentsInClassSQL(ClassId $classId): string
    {
        return 'SELECT * 
        FROM students s 
        LEFT JOIN students_classses sc ON sc.student_id = s.id 
        WHERE sc.class_id = {$this->classId}';
    }
}
```

Ahora es la Specification la que dice qué método del repositorio utilizar para obtener los datos:

```php
class IsEnrolledInClass implements StudentSpecification
{
    private ClassId $classId;
    
    public function __construct(ClassId $classId)
    {
        $this->classId = $classId;
    }
    
    public function isSatisfiedBy(Student $candidate): bool
    {
        return $candidate->isEnrolledInClass($this->classId);
    }
    
    public function selectSatisfyingFrom(StudentRepository $repository): StudentCollection
    {
        return $repository->findStudentsInClass($this->classId);
    }
}
```

Esto implica que los métodos como `findStudentsInClass` son públicos y están en la interface de `StudentRepository`, pero normalmente los usaremos siempre a través de la Specification.

El hecho de tener un método en el repositorio por Specification es el principal inconveniente de esta solución. Pero podríamos ir más lejos y aplicar a las Specification el principio de inversión de dependencias y tener implementaciones específicas de la interfaz en dominio.

Usando el mismo ejemplo, sería algo así:

```php
interface StudentSpecification
{
    public function isSatisfiedBy(Object $candidate): bool;
    
    public function selectSatisfying(): string;
}
```

Implementamos la interfaz para tener una clase abstracta que proporcione la funcionalidad de `isSatisfiedBy` al dominio:

```php
abstract IsEnrolledInClass implments StudentSpecification
{
    private ClassId $classId;
    
    public function __construct(ClassId $classId)
    {
        $this->classId = $classId;
    }
    
    public function isSatisfiedBy(Student $candidate): bool
    {
        return $candidate->isEnrolledInClass($this->classId);
    }
    
    abstract public function selectSatisfying(): string;
}
```

Las implementaciones concretas dependen del sistema de persistencia e implementarían solo el método `selectSatisfying`:

```php
class SQLIsEnrolledInClass extends IsEnrolledInClass
{
    public function selectSatisfying(): string
    {
        return 'SELECT * 
        FROM students s 
        LEFT JOIN students_classses sc ON sc.student_id = s.id 
        WHERE sc.class_id = {$this->classId}';
    }
}
```

Para ser usada con el repositorio así:

```php
class SQLStudentRepository implements StudentRepository
{
    // ...
    public function findSatisfying(Specification $specification)
    {   
        $sql = $specification->selectSatisfying();
        
        $result = $this->execSql($sql);
        
        return $this->mapResultToAggregates($result);
    }
}
```

El resultado es que tenemos implementaciones específicas para el sistema de persistencia residiendo en la capa de infraestructura, con una abstracción de la regla en dominio.

## Como gestionar proyecciones de bases de datos

¿Qué quiere decir esto? Intentemos verlo con algún ejemplo.

Nuestra aplicación de gestión escolar ofrece la posibilidad de obtener listas de grupos de clases. La más sencilla de ellas muestra simplemente el nombre de la estudiante y su número de identificación en el grupo. Algo así:

1. Chang, Cho
2. Granger, Hermione
3. Lovegood, Luna
4. Malfoy, Draco
5. Potter, Harry
6. Weasley, Ronald

Usar un repositorio para extraer una colección de agregados `Student` parece un poco excesivo y poco práctico. Sería mucho más eficiente una _query_ directa al sistema de persistencia que nos traiga directamente los datos necesarios.

¿Qué entendemos por proyección si no estamos hablando de _Event Sourcing_? En _Event Sourcing_ hablamos de proyección al referirnos a instantáneas del estado del sistema en un momento dado y generadas para una necesidad determinada.

En nuestro caso, podríamos utilizar una variante de este concepto para referirnos a extracciones del estado del sistema que satisfacen una necesidad específica, como las que podría tener una API, una vista de la UI o un report.

Otra forma de referirse a un concepto similar son los _Read Models_. _Read Model_ es un concepto de CQRS que se refiere a los objetos responsables de la lectura de datos del sistema de persistencia. Igualmente, los _Read Models_ son específicos para las peticiones concretas que se realicen.

Ambos conceptos nos permiten resolver un tema que en DDD nunca queda del todo claro: ¿cómo generamos simples listados o reports sin tener que montar decenas de entidades o agregados? No hay muchas aplicaciones que no requieran en algún momento presentar algún tipo de vista (de nuevo, en sentido genérico, no solo UI, sino también API y similares) en forma de lista con unas pocas propiedades de un agregado.

Aquí entraría este concepto de proyección o read model cuya característica principal es que nos proporciona colecciones de DTOs creados específicamente para la petición que recibe el sistema y que no necesita pasar por la capa de dominio realmente. Esto es así porque se trata más bien de necesidades de la aplicación que no afectan al estado del dominio, ya que solo son lecturas, ni implican realmente reglas del dominio, pues su manipulación se basa normalmente en la aplicación de simples filtros para afinar la selección de información. Con todo, nos ayudan a proporcionar acceso a los consumidores a los agregados individuales en los que estén interesados o la posibilidad de seleccionarlos para operar con ellos.

Tomando nuestro ejemplo de la lista, una posible implementación sería:

El DTO:

```php
class ClassListStudent
{
    public int $number;
    public string $listName;
}
```

Una posibilidad es considerar el servicio que obtiene los datos como un `ReadRepository` (repositorio de lecturas), pero hay que tener en cuenta que:

- No devuelve entidades o agregados por lo que no es un repositorio en el sentido en que lo consideraríamos en dominio
- Por tanto, no tiene que implementar la interfaz de un repositorio

Así que para dejarlo más claro, creo que prefiero optar por considerarlo como un servicio de aplicación y, en todo caso, definir su interfaz para invertir la dependencia con infraestructura:

```php
interface GetClassList
{
    public function forClass(ClassId $classId): array;    
}
```

Y esta sería una implementación muy simple:

```php
class SQLGetClassList implements GetClassList
{
    private Connection $connection;
    
    public function __construct(Connection $connection)
    {
        $this->connection = $connection;
    }
    public function forClass(ClassId $classId): array
    {
        $sql =  'SELECT s.number, s.last_name, s.first_name
        FROM students s 
        LEFT JOIN students_classses sc ON sc.student_id = s.id 
        WHERE sc.class_id = {$classId}';
        
        $result = $this->connection->executeSQL($sql);
        
        $list = array_map(static function ($row) {
            $listRow = new ClassListStudent();
            $listRow->number = $row['number'];
            $listRow->listName = $row['last_name'].' '.$row['first_name']
        }, $result);
        
        return $list;
    }
}
```


