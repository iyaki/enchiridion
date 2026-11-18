---
title: "Recopilación buenas prácticas de POO"
notion_id: 984fb339-7ac8-4c18-966a-542b28cf0d92
notion_url: https://app.notion.com/p/Recopilaci-n-buenas-pr-cticas-de-POO-984fb3397ac84c18966a542b28cf0d92
last_edited: 2026-09-21T17:06:00.000Z
source_url: https://app.notion.com/p/Recopilaci-n-buenas-pr-cticas-de-POO-984fb3397ac84c18966a542b28cf0d92
tags: ["Article", "The Talking Bit - Fran Iglesias", "Español", "Object Oriented Programming"]
---
Esta es una recopilación de artículos publicados por [Fran Iglesias](https://github.com/franiglesias) en su blog [The Talking Bit](https://franiglesias.github.io/) relacionados con las buenas prácticas que suelen relacionarse al paradigma de programación orientada a objetos.

# Índice

<!-- unsupported block: table_of_contents -->

---

# [Los principios SOLID](https://franiglesias.github.io/principios-solid/)

> Los principios SOLID son cinco principios en el diseño orientado a objetos compilados por Robert C. Martin, aunque quien primero se refirió a ellos con este nombre fue Michael Feathers. Martin los tomó como base para sus trabajos “Clean Code” y “Clean Architecture” y les dio la formulación más conocida.



Se trata de cinco principios que nos proporcionan una guía para crear código que sea reutilizable y se pueda mantener con facilidad. Como tales principios, nos aportan unos criterios con los que evaluar nuestras decisiones de diseño y una guía para avanzar en el trabajo. El código que los respeta es más fácil de reutilizar y mantener que el código que los ignora. Obviamente, no tienen una forma completa de implementarse y, en cada caso, pueden tener una interpretación algo diferente. No son recetas detalladas para aplicar ciegamente, sino que requieren de nosotros una reflexión profunda sobre el código que estamos escribiendo.

Los cinco principios son:

- Principio de única responsabilidad (Single Responsibility Principle)
- Principio abierto/cerrado (Open/Closed Principle)
- Principio de sustitución de Liskov (Liskov Substitution Principle)
- Principio de segregación de interfaces (Interface Segregation Principle)
- Principio de inversión de dependencias (Dependency Inversion Principle)

Si eres un desarrollador único o todo-en-uno en tu empresa o proyecto puedes pensar “para qué necesito que mi código sea mantenible o reutilizable si voy a ser yo quién se ocupe de él en el futuro y quién lleva ocupándose durante estos años y que nadie más va a verlo”. La respuesta es sencilla: dentro de unos pocos meses vas a ser un programador distinto del que eres ahora. Cuando después de varias semanas trabajando en una parte de tu código te dirijas a otra parte a hacer una modificación tal vez acabes preguntándote qué estúpida razón te llevó a escribir esa función o a crear esa clase y no tendrás respuesta. Solo por tu salud mental y laboral deberías escribir un código limpio, legible y mantenible. Seguramente tendrás por ahí secciones completas de código que no te atreves a tocar porque no sabes qué puede pasar si lo haces, o partes que se duplican en varios sitios, sabes que tienes código que apesta y que no te atreverías a enseñar. También por eso, por celo profesional, deberías aplicar los principios SOLID.

Los principios SOLID están muy relacionados entre sí, por lo que suele ocurrir que al intentar cumplir uno de ellos estás contribuyendo a cumplir otros.

Vamos a examinarlos en detalle.

## Principio de única responsabilidad

Este principio dice que las clases deberían tener una única razón para cambiar.

Otra forma de decirlo es que las clases deberían hacer una sola cosa. El problema de esta definición, aparentemente más clara, es que puede estrechar demasiado tu visión: no se trata de que las clases tengan solo un método o algo por el estilo.

En cualquier organización hay secciones o departamentos. Cada uno de ellos puede tener necesidades diferentes con respecto a un asunto. En un colegio, por ejemplo, el profesorado se preocupa de las calificaciones y asistencia de los alumnos, mientras que la administración le interesa saber qué servicios consume el alumno y por cuáles debe facturarle, al servicio de comedor o cantina le interesa saber si el alumno va a comer ese día o no, y al servicio de transporte le interesa saber dónde vive. Al diseñar una aplicación para gestionar un colegio nos encontraremos con estas visiones y peticiones desde los distintos servicios.

Si más de un agente puede solicitar cambios en una de nuestras clases, eso es que la clase tiene muchas razones para cambiar y, por lo tanto, mantiene muchas responsabilidades. En consecuencia, será necesario repartir esas responsabilidades entre clases.

## Principio abierto/cerrado

Este principio dice que una clase debe estar cerrada para modificación y abierta para extensión. Fue enunciado inicialmente por Bertrand Meyer.

La programación orientada a objetos persigue la reutilización del código. No siempre es posible reutilizar el código directamente, porque las necesidades o las tecnologías subyacentes van cambiando, y nos vemos tentados a modificar ese código para adaptarlo a la nueva situación.

El principio abierto/cerrado nos dice que debemos evitar justamente eso y no tocar el código de las clases que ya está terminado. La razón es que si esas clases están siendo usadas en otra parte (del mismo proyecto o de otros) estaremos alterando su comportamiento y provocando efectos indeseados. En lugar de eso, usaríamos mecanismos de extensión, como la herencia o la composición, para utilizar esas clases a la vez que modificamos su comportamiento.

Cuando creamos nuevas clases es importante tener en cuenta este principio para facilitar su extensión en un futuro.

## Principio de sustitución de Liskov

En una jerarquía de clases, las clases base y las subclases deben poder intercambiarse sin tener que alterar el código que las utiliza.

Esto no quiere decir que tengan que hacer exactamente lo mismo, sino que han de poder reemplazarse.

El reverso de este principio es que no debemos extender clases mediante herencia por el hecho de aprovechar código de las clases bases o por conseguir forzar que una clase sea una “hija de” y superar un _type hinting_ si no existe una relación que justifique la herencia (ser clases con el mismo tipo de comportamiento, pero que lo realizan de manera diferente). En ese caso, es preferible basar el polimorfismo en una interfaz (ver el Principio de segregación de interfaces).

Como regla práctica extraída de este principio, se podría decir que la herencia debe usarse solamente para obtener especializaciones a partir de una clase más general.

## Principio de segregación de interfaces

El principio de segregación de interfaces puede definirse diciendo que una clase no debería verse obligada a depender de métodos o propiedades que no necesita.

Supongamos una clase que debe extender otra clase base. En realidad, nuestra clase solo está interesada en dos métodos de la clase base, mientras que el resto de métodos no los necesita para nada. Si extendemos la clase base mediante herencia arrastramos un montón de métodos que nuestra clase no debería tener.

Otra forma de verlo es decir que las interfaces se han de definir a partir de las necesidades de la clase cliente.

## Principio de inversión de dependencia

[La inversión de dependencias no es solo lo que tú piensas](http://blog.koalite.com/2015/04/la-inversion-de-dependencias-no-es-solo-lo-que-tu-piensas/)

El principio de inversión de dependencia dice que:

- Los módulos de alto nivel no deben depender de módulos de bajo nivel. Ambos deben depender de abstracciones.
- Las abstracciones no deben depender de detalles, son los detalles los que deben depender de abstracciones.

Las abstracciones definen conceptos que son estables en el tiempo, mientras que los detalles de implementación pueden cambiar con frecuencia. Una interfaz es una abstracción, pero una clase que la implemente de forma concreta es un detalle. Por tanto, cuando una clase necesita usar otra, debemos establecer la dependencia de una interfaz, o lo que es lo mismo, el type hinting indica una interfaz no una clase concreta. De ese modo, podremos cambiar la implementación (el detalle) cuando sea necesario sin tener que tocar la clase usuaria lo que, por cierto, contribuye a cumplir el principio abierto/cerrado.

---

# [**Más allá de SOLID, principios básicos**](https://franiglesias.github.io/beyond-solid/)

> Los principios de diseño de software no son recetas ni objetivos, son guías para la toma de decisiones y para la evaluación.

## Qué es buen código

Es muy complicado definir qué es buen código. Entran en juego toda una serie de factores, incluyendo preferencias personales o su ajuste a un contexto determinado. El buen código funciona, por supuesto, pero también esperamos que tenga algunas características más: que sea robusto, reutilizable, testeable, bien organizado, legible y sostenible.

¿Qué significa cada una de estas propiedades? ¿Son medibles o auditables? En la mayor parte de los casos podemos dar definiciones aproximadas que, por desgracia, pueden tener mucho que ver con preferencias personales. En cuanto a la medición de las mismas, es posible que solo podamos tomar medidas relativas, del tipo _más que ayer, pero menos que mañana_. Veamos:

Robusto. Diríamos que el código es robusto si está libre de defectos, si la tasa de fallos es reducida o nula, si es capaz de gestionar ciertas condiciones adversas, como la caída de redes o servicios de los que depende.

Reutilizable. El código es reutilizable si podemos usar la misma funcionalidad con distintos datos, incluso en otros contextos.

Testeable. Decimos que el código es testeable cuando podemos construir un test que verifique su comportamiento, con la seguridad de que el resultado del test es función de la ejecución del código probado y no está influenciado por factores externos.

Bien organizado. El código bien organizado nos permite localizar e identificar fácilmente una determinada sección que nos interesa examinar, incluso aunque no lo conozcamos bien.

Legible. El código legible es aquel que al leerlo nos comunica su intención.

Sostenible. Es sostenible aquel código que nos permite modificar o extender la funcionalidad sin provocar defectos y con el mínimo esfuerzo posible.

En conjunto, podemos decir que todas estas características son deseables para cualquier base de código. Sin embargo, es difícil afirmar en qué grado es suficiente haberlas desarrollado para considerarlo satisfactorio. Personalmente, diría que podemos aspirar a un código suficientemente bueno que siempre es mejorable, pero no existe el código perfecto.

En buena medida esto ocurre porque el código es una expresión de un conocimiento compartido sobre el problema que intenta resolver o ayudar a resolver, un conocimiento que siempre es incompleto y refinable. Además, es también el resultado de nuestra capacidad para expresar ese conocimiento en un lenguaje de programación, capacidad que siempre puede crecer como resultado de la experiencia y el aprendizaje.

Puede que sea más fácil definir lo que no es buen código. El mal código carece de las propiedades que hemos mencionando anteriormente:

- No hace lo que se espera que haga, o lo hace de forma manifiestamente incompleta o equivocada.
- Tiene múltiples fallos y problemas, no funciona si alguna condición externa es adversa.
- No es reutilizable.
- No está bien organizado. Encontrar algo en él es un esfuerzo enorme, que lleva a multitud de vías muertas o a descubrir que una misma idea es expresada en diferentes lugares, diseminada en múltiples sitios e incluso ausente.
- No es testeable. Hay que hacer un gran esfuerzo solo para poder empezar a hacer un test, incluyendo tener que transformar el código para hacerlo posible.
- No es legible. No es posible entender qué hace solo con leer el código, hay que buscar otras fuentes de información para hacerse una idea de lo que estaría ocurriendo, incluso aunque exista documentación o comentarios, porque ya han perdido la sincronía con el código.
- No es sostenible. Llevar a cabo una modificación es peligroso porque puede resultar incompleta o tener efectos en lugares insospechados.

## Principios de diseño para crear buen código

A lo largo de la breve historia de la disciplina se han ido desarrollando herramientas, metodologías, lenguajes y paradigmas de programación que, poco a poco, han ido destilando un conjunto de principios que tratan de establecer criterios que nos permitan distinguir entre buen y mal código.

Los principios no son leyes ni dogmas, son ante todo criterios con los que tomar o evaluar ciertas decisiones cuando expresamos un modelo mental en forma de código.

Estos principios interactúan en un sistema. En muchas ocasiones, enfatizar la aplicación de uno de ellos puede dar lugar a tener que relajar la aplicación de otro. No son, por tanto, recetas que se puedan seguir ciegamente. Son más bien estrategias para razonar sobre aquello que queremos conseguir con el código teniendo en cuenta el contexto en que se van a aplicar.

De hecho, unos principios derivan de otros. En algunos casos tienen aplicabilidad en un contexto o paradigma de programación. En otros casos, ni siquiera son principios que hayan nacido en el campo del desarrollo de software.

Intentaré presentar algunos de estos principios, bastante conocidos seguramente, en tres artículos. En el primero, trataremos sobre un pequeño conjunto de principios básicos. En el segundo hablaremos sobre los conocidísimos principio SOLID. Por último, en el tercero, un grupo de principios que he denominado “los consejos pragmáticos”.

## Principios básicos

### Separación de intereses (separation of concerns)

Se atribuye a Edsger Dijkstra y dice simplemente: diferentes partes de un programa se ocupan de diferentes intereses.

Es uno de los principios más fundamentales del desarrollo de software y nos ayuda a estructurar y organizar el software, separando sus componentes en función del tipo de intereses o necesidades de los que se ocupa. Hay partes dedicadas al acceso a bases de datos, a la presentación, al modelo de dominio, a servicios de la aplicación, a comunicarse con otros sistemas, etc.

Se puede decir que todas las arquitecturas basadas en capas, los patrones arquitectónicos e incluso el Domain Driven Design se asientan en este principio, aplicado a distintos niveles de abstracción y en combinación con otros. De hecho, está en el origen algunos otros que veremos más adelante.

En resumen, este principio nos ayuda a separar partes del código de una manera significativa, yendo juntas aquellas que atienen a los mismos intereses, lo cual es un punto de partida para una buena organización.

Un corolario de este principio es que si determinamos que una unidad de software sirve a dos o más intereses, deberíamos partirla y separarla en tantos trozos como necesidades sirve.

### Principio de abstracción

Enunciado por Benjamin Pierce, dice que cada parte significativa de funcionalidad de un programa debe implementarse en un solo lugar del código fuente.

Conocemos una versión más “pragmática” en forma del principio DRY, del que hablaremos en la tercera entrega. Pero en general, este principio lo que nos dice es que toda funcionalidad significativa debería existir en un único lugar accesible al resto del programa, evitando repetir el código que la ejecuta.

Este principio nos ayuda a tener criterios para modularizar código y para favorecer su sostenibilidad y su capacidad de ser reutilizado.

En cuando a la modularidad, una funcionalidad que debe reutilizarse tiene que existir en forma única en un módulo que sea utilizable desde las partes del programa que lo precise. Desde el punto de vista de este consumidor, esa funcionalidad está abstracta en una función o clase, por lo que no necesita preocuparse de los detalles de su implementación.

La sostenibilidad se ve favorecida porque si necesitamos cambiar esa funcionalidad solo existirá un punto en el que debamos intervenir, sin riesgo de que haya distintas partes del código que puedan quedar sin actualizar.

### Separación comando pregunta (command query separation o CQS)

Este principio fue enunciado por Bertrand Meyer y establece que cada método o función debe ser o bien un comando, que provoca un cambio en el sistema ejecutando una acción, o bien una pregunta (o query). que obtiene una información del sistema sin modificarlo.

El meollo de este principio es que una función no debería provocar un cambio en el estado del sistema e informar a la vez de su estado. Un ejemplo en el que se puede ver fácilmente cuál es el tipo de riesgo que supone no cumplir este principio es imaginar un sistema de aceleración y frenado de vehículos. Imagina que la función `frenar` devuelve también la velocidad: ¿de qué velocidad estamos hablando: de la que teníamos antes de aplicar la acción de frenar o de la resultante? Y si quiero saber cuál es la velocidad actual, ¿tengo que acelerar o frenar el vehículo para obtener una respuesta?

Por este tipo de problemas es necesario separar la realización de cambios en el sistema, de la obtención de información acerca de esos cambios.

En consecuencia, normalmente de los comandos no obtenemos respuestas. En su lugar nos preocupa solo saber que se han ejecutado y, en todo caso, si ha habido algo que ha impedido completar esa ejecución. De ahí que sea la práctica habitual que un comando no retorne nada (void) pero puede lanzar excepciones de modo que su emisor pueda saber que ha fallado. Una alternativa más avanzada es que, como resultado de la ejecución, se dispare un evento que notifica a la aplicación el cambio que se acaba de producir, permitiendo que otras partes de la misma actúen en consecuencia.

Por otro lado, aplicando este principio, tendremos la seguridad de que las _queries_ no producen cambios (_side effects_) en el estado del sistema, por lo que devuelven información correcta acerca del mismo. El concepto de _función pura_ tiene que ver con este principio también.

El principio CQS está en la base de los sistemas de mensajería de aplicaciones (commands, queries y events) y también en el patrón _Command Query Responsibility Segregation_ (CQRS) que tiene su aplicación en sistema que requieren un alto rendimiento.

La aplicación del principio CQS favorece la testeablidad del software y contribuye a su robustez.

### Principio del mínimo conocimiento (Ley de Demeter o _No hables con extraños_)

Ian Holland es el promotor de este principio definido como práctica recomendada en un proyecto denominado [Demeter](https://www2.ccs.neu.edu/research/demeter/), del cual tomó su nombre más popular. Su nombre formal tiene que ver con el enunciado del principio, que dice, más o menos, que una unidad de software solo puede hablar con aquellas otras unidades que conozca.

¿Cuales son estas otras unidades que una unidad de software puede conocer? Pues las siguientes:

- Aquellas instanciadas dentro de la propia unidad de software considerada.
- Aquellas que se pasan como parámetros a la unidad considerada.

En el caso de orientación a objetos, asumiendo que la unidad es un método de una clase

- La propia clase en la que se define el método.
- Los métodos y atributos de la clase.

Una consecuencia de este principio es que una unidad de software no debe tener conocimiento de la estructura interna de esas otras unidades con las que puede hablar. No puede usarlas para acceder a propiedades internas de estas colaboradoras, sino que solo puede relacionarse con ellas a través de su interfaz pública.

La aplicación de este principio ayuda mucho a la testeabilidad, sobre todo cuando en el nivel unitario utilizamos dobles de tests. Además, aligera el acoplamiento eliminando dependencias ocultas o transitivas (cuando desde la unidad A, usamos una unidad B porque contiene una tercera unidad C que es la que realmente queremos usar), y favorece la cohesión.

### Cohesión vs Acoplamiento

Una de las primeras definiciones de estas propiedades es de Larry Constantine. Las propiedades de Cohesión y Acoplamiento nos ayudan a definir módulos de software.

La cohesión se refiere al grado en que los elementos de un módulo van juntos. Nos interesa que la cohesión dentro de un módulo sea alta. Una forma práctica de verlo es pensar que “las cosas que cambian juntas deben ir juntas”, lo que se conoce a veces como _principio de covarianza_. Las cosas que no cambian juntas deberían estar separadas.

Para aumentar la cohesión tenemos que reunir en un mismo módulo las unidades de software que tienden a presentarse juntas, a ser usadas juntas y a cambiar juntas. Además, podemos incrementarla aún más sacando de un módulo aquellas unidades que no tienen este tipo de relación con las demás.

Este principio está en la base del empaquetado de software y, por tanto, en la organización del código y su inteligibilidad.

El acoplamiento se refiere al grado de interdependencia entre módulos distintos. Nos interesa que el acoplamiento entre módulos sea lo más bajo posible. El acoplamiento será cero cuando un módulo no utilice a otro, por lo que siempre tendremos un cierto grado de acoplamiento. Ahora bien, este acoplamiento puede ser ligero y controlado por nosotras, algo que se puede conseguir aplicando otros principios, como puede ser el de inversión de dependencias.

El bajo acoplamiento ayuda en la robustez y la testeabilidad.

En conjunto, la relación entre cohesión y acoplamiento son claves en la organización del software y la posibilidad de reutilización de los distintos módulos.

---

# [Más allá de SOLID, los cimientos](https://franiglesias.github.io/beyond-solid-2/)

> A veces, cuando se habla de principios de diseño de software parece que solo existiesen estos cinco. Así que los repasaremos con un poco de contexto.

Los principios SOLID son cinco principios en el diseño orientado a objetos compilados por Robert C. Martin. Puedes encontrarlos [aquí](http://butunclebob.com/ArticleS.UncleBob.PrinciplesOfOod), con enlaces a los artículos originales.

Diría que su popularidad es debida tanto a su inclusión en libros como _Clean Code_, como del afortunado acrónimo con el que se nombra habitualmente el conjunto. Lo cierto es que componen un sistema coherente y mantienen una cierta interdependencia entre ellos. Seguirlos es una gran ayuda para lograr un código sostenible, fácil de entender y de extender, si bien es cierto que poner énfasis en uno de ellos puede perjudicar cumplir otros.

Con todo, yo echo de menos entre ellos el Principio de Mínimo Conocimiento, o Ley de Demeter. Pero de eso podremos hablar después.

## Principio de única responsabilidad (Single Responsibility Principle, SRP)

Este principio dice que las clases deberían tener una única razón para cambiar.

Se puede decir que es una aplicación o reformulación del principio de separación de intereses de Dijsktra. Con frecuencia se intenta explicar o resumir diciendo que las clases o los métodos de las clases deberían hacer una sola cosa, aunque ese es un enfoque erróneo.

El SRP alude a las razones para cambiar el comportamiento de una clase o unidad de software desde el punto de vista de los usuarios de esa clase o unidad. Distintos usuarios podrían requerir en un momento dado distintos tipos de cambios, lo que sería una indicación de que la clase o método está manteniendo muchas responsabilidades. Además, estas responsabilidades pueden verse incluso en un nivel más global desde el punto de visto de los intereses de los _stake holders_ del producto de software.

El SRP tiene más que ver con los propósitos del software que con sus aspectos técnicos. No se trata de las tareas discretas que conlleva cualquier función significativa para el dominio, lo que nos podría llevar a diseños muy atomizados con una gran complejidad al tratar de aislar cada una de esas operaciones en una unidad de software.

## Principio abierto/cerrado (Open Close Principle, OCP)

Este principio dice que una clase debe estar cerrada para modificación y abierta para extensión. Fue enunciado por Bertrand Meyer.

La programación orientada a objetos persigue la reutilización del código. No siempre es posible reutilizar el código directamente, porque las necesidades o las tecnologías subyacentes van cambiando, y nos vemos tentadas a modificar ese código para adaptarlo a la nueva situación.

El principio abierto/cerrado nos dice que debemos evitar eso y evitar tocar el código de las clases que ya está terminado y en producción. La razón es que si esas clases están siendo usadas por otras partes del software generamos un riesgo de que el cambio de comportamiento afecte a los resultados obtenidos en esos otros lugares. En otras palabras, generaremos problemas de compatibilidad hacia atrás o _backwards compatibility_, que darán lugar a defectos en el software.

En lugar de eso, usaríamos mecanismos de extensión, como la herencia o la composición, de modo que las clases existentes pueden seguir sirviendo a sus consumidores actuales, mientras que las nuevas clases aprovechan esa funcionalidad a la vez que añaden nueva.

Cuando creamos nuevas clases es importante tener en cuenta este principio para facilitar su extensión en un futuro.

## Principio de sustitución de Liskov (Liskov Substitution Principle, LSP)

En una jerarquía de clases, las clases base y las subclases deben poder intercambiarse sin tener que alterar el código que las utiliza. El enunciado original es de Barbara Liskov.

El principio no dice que tengan que hacer exactamente lo mismo, obviamente, sino que han de poder reemplazarse sin que el código que las usa necesite ser alterado.

El reverso de este principio es que no debemos extender clases mediante herencia por el hecho de aprovechar código de las clases base o por conseguir forzar que una clase sea una “hija de” para superar un _type hinting_. La herencia no consiste en que varias clases compartan funcionalidad, la herencia trata acerca de clases semánticamente equivalentes pero que tienen comportamientos más o menos especializados.

Otro de sus corolarios es que una clase que extienda de otra no puede modificar la interfaz pública respecto de su clase base añadiendo nuevos métodos. Eso es lo que provocará la violación del principio. Si tenemos la necesidad de hacerlo así, es muy posible que estemos necesitando una nueva familia de clases. De hecho, eso nos lleva al siguiente principio: el de Segregación de Interfaces.

## Principio de segregación de interfaces (Interface Segregation Principle, ISP)

El principio de segregación de interfaces puede definirse diciendo que una clase no debería verse obligada a depender de métodos o propiedades que no necesita. Expresado de otra forma: una interfaz solo debería tener los métodos que sus consumidores necesitan, y ni uno más.

Supongamos que tenemos una clase con ocho métodos públicos y algunos consumidores de la misma que solo requieren dos de esos métodos. Si hacemos que dependan directamente de esa clase (de la totalidad de su interfaz pública) los obligamos a cargar con seis métodos que no necesitan. Ahora bien, si definimos una interfaz para esos dos métodos, y hacemos que sea implementada por la clase, la dependencia de los consumidores será a esa nueva interfaz.

Si la clase de la que hablamos sirve a otros consumidores que utilizan otros métodos, podrían definirse nuevas interfaces. La ganancia aquí es que los consumidores solo dependen de aquellos métodos que usan, lo que nos permitirá cambiar esa dependencia con más facilidad cuando sea necesario.

## Principio de inversión de dependencia (Dependency Inversion Principle, DIP)

El principio de inversión de dependencia dice:

- Los módulos de alto nivel no deben depender de módulos de bajo nivel. Ambos deben depender de abstracciones.
- Las abstracciones no deben depender de detalles, son los detalles los que deben depender de abstracciones.

Las abstracciones definen conceptos que son estables en el tiempo, mientras que los detalles de implementación pueden cambiar con frecuencia. Una interfaz es una abstracción, pero una clase que la implemente de forma concreta es un detalle. Por tanto, cuando una clase necesita usar otra, debemos establecer la dependencia de una interfaz, o lo que es lo mismo, el _type hinting_ indica una interfaz no una clase concreta. De ese modo, podremos cambiar la implementación (el detalle) cuando sea necesario sin tener que tocar la clase usuaria lo que, por cierto, contribuye a cumplir el principio abierto/cerrado.

## ¿Por qué SOLID?

Si utilizamos los principios SOLID para diseñar nuestro software orientado a objetos, podemos tener bastantes garantías de que cumplirá las características que mencionamos en el artículo anterior.

Aún así, pienso que tendríamos que añadir el Principio de Mínimo Conocimiento, o Ley de Démeter, para conseguirlo.

El SRP nos dice que cada unidad de software debe tener una única razón para cambiar, debe ser responsable de un solo interés (no se puede servir a dos amos).

El OCP nos dice que una clase no debería modificarse, sino extenderse para evitar romper el sistema.

El LSP nos dice que solo deberíamos usar la herencia para especializar comportamientos en las jerarquías de clases, lo que nos indica cuando nos conviene usar este mecanismo y cuándo es preferible la composición. Esto nos ayuda con el OCP.

El ISP nos dice que las interfaces deben definirse por las necesidades de sus consumidores, lo que entronca con el SRP y nos impulsa a definir interfaces con pocos métodos.

El DIP nos dice que las dependencias deben hacerse sobre abstracciones, lo que nos facilita usar implementaciones alternativas, manteniendo el control sobre cómo deben interactuar.

En este esquema, el Principio de Mínimo Conocimiento nos ayuda a limitar el modo en que las distintas clases se relacionan, al forzar que lo hagan precisamente a través de sus interfaces, ignorando del todo su estructura interna.

En cierto modo, este principio se sobreentiende en SOLID, pero no está de más hacer explícito lo implícito.

Lástima que estropee el acrónimo.

---

# [**Más allá de SOLID, los consejos prácticos**](https://franiglesias.github.io/beyond-solid-3/)

> Los principios SOLID están sobrevalorados y otras perlas de sabiduría práctica.

Los principios SOLID constituyen un buen sistema de principios, son muy coherentes, pero incompletos en mi opinión. En realidad, creo que faltan dos para tener un juego de criterios realmente potentes.

De uno ya hemos hablado: La ley de Demeter, o Principio de Mínimo Conocimiento. El otro es _Tell, don’t ask_, que completa los 7 principios capitales del desarrollo de software orientado a objetos.

## Tell, don’t ask

No conozco otra formulación de este principio que suene un poco más formal. Lo enunciaron los _pragmáticos_ Andy Hunt y Dave Thomas y dice más o menos esto:

_No deberías tomar decisiones basadas en el estado de un objeto al que llamas que resulten en el cambio de estado de ese objeto._

Este principio tienen que ver con la encapsulación y la ocultación de información de OOP (_information hiding_). La mejor forma de explicarlo es con un ejemplo simple: Supongamos que tenemos un objeto `Score` para llevar la puntuación en un juego. Cuando la jugadora logra puntos, hay que incrementar `Score`. Esto puede hacerse así:

```php
$points = $this->player->score->points();
$this->player->score->setPoints($points + $newPoints);
```

Bien, estamos violando la Ley de Demeter, ¿no?. Resolvamos eso primero:

```php
// Player...

public function points() {
    return $this->score->points();
}

public function setPoints(int $points) {
    $this->score->setPoints($points);
}

//...

$points = $this->player->points();
$this->player->setPoints($points + $newPoints);
```

Hemos hecho que desde fuera de `Player` no tengamos que saber nada acerca de cómo se guardan internamente los puntos o lo que sea. Pero ahora sí tenemos una violación de _Tell, don’t ask_. Básicamente estamos preguntando por un estado para cambiarlo.

Aplicando _Tell, don’t ask_, la cosa debería ser algo más o menos así;

```php
$this->player->winPoints($newPoints);
```

Y ya.

Por dentro podría ser así, aunque desde fuera nos da igual la implementación:

```php
// Player...

public function winPoints(int $points) {
    $currentPoints = $this->score->points();
    $this->score->setPoints($currentPoints + $points);
}
```

Pero podemos hacerlo aún mejor, aplicando el mismo principio, lo que mueve las responsabilidades a donde realmente corresponden:

```php
// Player...

public function winPoints(int $points) {
    $this->score->incrementBy($points);
}
```

_Tell, don’t ask_ junto a la _Ley de Demeter_ son dos poderosas herramientas para ayudar a poner las responsabilidades en los objetos correctos.

De hecho, son una gran ayuda para empezar a modernizar un _legacy_. Típicamente nos vamos a encontrar este tipo de objetos anémicos que podremos alimentar adecuadamente de comportamientos significativos moviéndolos a los objetos adecuados.

Por eso, suelo pensar que los principios SOLID están sobrevalorados: están incompletos sin estos dos.

## Otras perlas de sabiduría

### Keep it simple stupid

Sin coma. Es un consejo de Kelly Jonhson que afirma que la mayoría de los sistemas funcionan mejor si se mantienen simples y sencillos. Es decir, sencillos en el sentido de estúpidos o lo más “tontos” posible.

Una formulación típica es la de “Keep it simple, stupid”, que parece más graciosa, pero que es muy poco útil ya que no define la simplicidad.

La idea de la simpleza aquí es que el sistema necesite la menor cantidad de conocimiento posible, evitando suposiciones sobre lo que entregan o esperan recibir otros sistemas y reduciendo la complejidad de las soluciones. Es decir, nuestro sistema debería ser lo más estúpido, mecánico y predecible que podamos, lo que redundará en que será confiable, fácil de mantener y facil de testear.

### Ley de Gall

En relación con el consejo anterior, John Gall dijo que un sistema complejo que funciona ha evolucionado sin excepción de un sistema más simple que funcionaba. Siempre tienes que empezar con un sistema simple que funcione.

Esta es la base del desarrollo iterativo: Preguntarse: ¿cuál es la forma más simple posible de conseguir esto? Y trabajar a partir de ahí.

### Falla rápido

Este consejo se atribuye a Jim Gray y dice que la responsabilidad de un módulo que falla rápido es detectar errores y dejar el módulo que está el siguiente nivel más alto decida qué hacer.

En la práctica, este consejo nos dice que un módulo que está al final en una cadena de llamadas debería detectar cuanto antes un error y enviarlo al módulo que lo ha llamado. Es éste quien debe responsabilizarse de dar una respuesta adecuada a ese error o pasarlo al siguiente módulo de mayor nivel. El módulo de bajo nivel no tiene que tener el conocimiento necesario para gestionar el error.

Fallar pronto se traduce en el uso de tácticas como cláusulas de guarda, lanzamiento de excepciones, y similares.

### DRY, Don’t Repeat Yourself

Hunt y Thomas reformularon el Principio de Abstracción de Pierce, de una forma bastante interesante: _Cada fragmento de conocimiento debe tener una representación única, no ambigua y autoritativa dentro de un sistema_.

Este principio no se refiere a código. En realidad en un sistema el código se puede repetir muchas veces. El principio se refiere al conocimiento representado en el código. Complementa esto con los “7 principios capitales” y tendrás la clave para desarrollar software realmente valioso, mantenible y duradero.

### YAGNI, You ain’t gonna need it

Ron Jeffries aconseja no añadir funcionalidad hasta que no la vayas a necesitar. Es tan simple, y tan difícil, como desarrollar solo aquello que necesitas para que las cosas funcionen y no pensar en lo que podría necesitarse en un futuro.

Lo que no quita que desarrolles con flexibilidad y capacidad de adaptación a los cambios del futuro. Se refiere a que no insertes funcionalidad en un sistema simplemente porque puedes y a lo mejor un día la podrías querer llegar a usar.

### Peor es mejor

Richard P. Gabriel afirmó que la calidad no necesariamente se incrementa con la funcionalidad. En muchos casos, el problema de intentar incorporar más funcionalidad a un sistema no lo hace mejor y es muy posible que para lograr esa incorporación de forma no justificada por las necesidades de las usuarias se tenga que disminuir su calidad general aumentando los puntos de posible rotura.

Si combinas este consejo con el anterior, lo que tienes es una herramienta para decidir cuando tiene sentido incorporar una funcionalidad a un software, algo que debería estar guiado fundamentalmente por la necesidades de sus consumidores y no por decisiones tomadas fuera de contexto, simplemente porque podemos o porque mola.

---

# [**Más allá de SOLID, los principios olvidados**](https://franiglesias.github.io/beyond-solid-4/)

> Hay mucha vida más allá de los principios SOLID y, sobre todo, mucho antes.

GRASP, General Responsibility Assignment Software Patterns, es un conjunto de patrones o heurísticas para definir el reparto de responsabilidades de un sistema orientado a objetos. Básicamente nos ayudan a responder a la pregunta: ¿a qué clase pertenece esta responsabilidad? Así que se podría decir que cada uno de estos patrones nos proporciona una posible respuesta.

Los principios GRASP son el fundamento de otros muchos principios y nos ayudan a encontrar respuesta a preguntas bastante básicas.

## ¿Quién tiene que crear un objeto? Creator

La responsabilidad de crear un objeto de una cierta clase debería estar en otro objeto tal que cumple una o más de las siguientes condiciones:

- Contiene o agrega instancias de la clase a crear.
- Registra instancias de la clase a crear.
- Usa instancias de de la clase a crear.
- Contiene la información necesaria para instanciar objetos de la clase a crear.

La primera condición: contiene o agrega instancias de otra clase, se aplica, por ejemplo, en los agregados de dominio. En lugar de entregarle objetos ya instanciados, lo correcto sería entregarle los datos necesarios para crearlos de modo que el agregado se responsabiliza de proteger las invariantes de dominio que les afectan.

He aquí un ejemplo. Creamos un objeto `Order` y le añadimos un `Item` pasándole la información necesaria para crearlo:

```ruby
require 'rspec'
require_relative '../src/order.rb'

RSpec.describe 'Order should' do
  it "add items" do
    order = Order.new
    order.add_item 'Item 1', 1, 20
    expect(order.items.length).to eq(1)
  end
end
```

Este es el código. Order.rb

```ruby
require_relative './item.rb'

class Order

  def initialize
    @items = []
  end

  def add_item(name, quantity, price)
    new_item = Item.new name, quantity, price
    @items.push(new_item)
  end

  def items
    @items
  end
end
```

e Item.rb

```ruby
class Item
  def initialize(name, quantity, price)
    @name = name
    @quantity = quantity
    @price = price
  end
end
```

La cuarta condición nos ayuda a entender patrones de construcción como puede ser Factory o Builder, objetos que tienen la información necesaria para instanciar otros objetos.

## ¿Quién debería hacer esto? Information expert

Pon la responsabilidad en la clase que tenga la información necesaria para ejercerla, que es la `information expert` para ese asunto.

Uno de los ejemplos en los que resulta más fácil verlo en acción es en el modelado de una venta en una tienda, así que podemos aprovechar el ejemplo anterior.

Supongamos que necesitamos conocer el importe total del pedido. Normalmente tendremos clases como `Order` e `Item`. Es sencillo ver que `Order`, al contener la colección de `Items` solicitados, es quien tiene toda la información necesaria para poder calcular el importe total, mientras que `Item` sería responsable del importe de cada línea.

Lo podemos ver aquí:

```ruby
require 'rspec'
require_relative '../src/order.rb'

RSpec.describe 'Order should' do

  it "sum total amount" do
    order = Order.new
    order.add_item('Item 1', 2, 20)
    order.add_item('Item 2', 3, 15)

    expect(order.total_amount).to eq(85)
  end
end
```

Item.rb. `Item` sabe calcular el importe de la línea, ya que sabe tanto el precio como la cantidad:

```ruby
class Item
  def initialize(name, quantity, price)
    @name = name
    @quantity = quantity
    @price = price
  end

  def amount
    @price * @quantity
  end

end
```

Order.rb. Mientras, que `Order`, al contener todos los `Items`, puede calcular el total del pedido:

```ruby
require_relative './item.rb'

class Order

  def initialize
    @items = []
  end

  def add_item(name, quantity, price)
    new_item = Item.new name, quantity, price
    @items.push(new_item)
  end

  def items
    @items
  end

  def total_amount
    @items.sum { |item| item.amount }
  end
end
```

## ¿Cómo gestiono variantes basadas en clase de objetos? Polymorphism

El polimorfismo nos ayuda cuando tenemos que hacer distintas cosas en función del tipo de objeto o información que recibimos. La responsabilidad de cada variedad de comportamiento corresponde al tipo, de modo que el objeto consumidor u orquestador no tiene que saber previamente el tipo de objeto que está manejando, simplemente le envía el mensaje para que actúe. Hemos tratado esto más ampliamente en [un artículo sobre Programar sin ifs](https://franiglesias.github.io/programar-sin-ifs/).

El ejemplo clásico de las figuras geométricas nos viene bien aquí. Cada figura tiene unas propiedades ligeramente distintas, pero todas saben dibujarse o calcular su área. El código usuario de las clases no tiene que preguntar primero de qué tipo se trata para dibujarlas. Veámoslo en el siguiente test:

```ruby
require 'rspec'
require_relative '../src/surface_calculator.rb'

RSpec.describe 'Surface Calculator should' do
  it 'calculate surface joining different shapes' do
    surface = SurfaceCalculator.new
    surface.add_triangle 10, 20
    surface.add_square 5
    expect(surface.total).to eq(125)
  end
end
```

SurfaceCalculator es capaz de agregar distintos tipos de objetos Shape, cada uno de los cuales es una especialización:

```ruby
require_relative './triangle'
require_relative './square'

class SurfaceCalculator
  def initialize
    @shapes = []
  end

  def add_triangle(base, height)
    triangle = Triangle.new base, height
    @shapes.append triangle
  end

  def add_square(side)
    square = Square.new side
    @shapes.append square
  end

  def total
    @shapes.sum { |shape| shape.area }
  end
end
```

Como podemos ver en el método `total`, `SurfaceCalculator` solo tiene que preguntarle a cada Shape su área, sin necesidad de preguntarle cuál es su tipo primero. Esto hace que las figuras deban entender el mensaje `area` (o, lo que es lo mismo, tener el método `area`)

He aquí el código de las formas:

Triangle, representa un triángulo:

```ruby
class Triangle
  def initialize(base, height)
    @base = base
    @height = height

  end

  def area
    @height * @base / 2
  end

end
```

Square un cuadrado:

```ruby
class Square
  def initialize (side)
    @side = side
  end

  def area
    @side ** 2
  end
end
```

Podemos ver, por tanto, que cada una de las formas se encarga de realizar el cálculo especializado de su área, liberando a `SurfaceCalculator` de esa responsabilidad, y centrándolo en la suya propia que es agregar todas las áreas de las figuras que contiene.

Esto tiene el beneficio de que será fácil añadir nuevas formas, como Rectangle:

```ruby
class Rectangle
  def initialize(base, height)
    @base = base
    @height = height

  end

  def area
    @height * @base
  end

end
```

El problema es que esta forma de trabajar no es segura, ya que nada nos garantiza la existencia previa del método `area` en los diferentes tipos de forma. ¿Podemos garantizarlo de alguna forma? Vayamos al siguiente principio.

## ¿Cómo diseño los objetos para evitar el impacto de las variaciones? Protected variations

Es el principio tras la definición de interfaces. Lo que buscamos al aplicar este principio es proteger a los componentes del sistema de las variaciones de otros componentes. Se trataría de establecer algún tipo de contrato entre los componentes participantes que obligue a ser capaz de responder a unos mensajes determinados.

En Ruby, el lenguaje usando en estos ejemplos, y en otros lenguajes (Golang, etc.) no existe la idea de definir interfaces explícitamente. En su lugar, tenemos interfaces implícitas. Por eso, el código anterior funciona siempre que usemos objetos que tengan el método `area`. Esto ya sería una buena razón para aplicar el principio `Creator`, pues nos puede facilitar el asegurar que `SurfaceCalculator` solo utiliza formas de las que _sabe_ que le pueden dar una respuesta en el método `area`.

En otros lenguajes, como Java o PHP, tenemos que definir interfaces explícitamente, de modo que el propio intérprete o compilador nos obligan a que las clases las implementen correctamente.

Para mostrar un ejemplo en Ruby, vamos a simular interfaces con una clase base abstracta `Shape` que tiene el método `area`. Si no lo sobreescribimos en una clase hija, lanzará una excepción.

Shape es la clase base que representa una forma genérica.

```ruby
class Shape
  def area
    raise 'Method area not implemented'
  end
end
```

Triangle, representa un triángulo:

```ruby
require_relative 'shape.rb'

class Triangle < Shape
  def initialize(base, height)
    @base = base
    @height = height

  end

  def area
    @height * @base / 2
  end

end
```

Square un cuadrado:

```ruby
require_relative 'shape.rb'

class Square < Shape
  def initialize (side)
    @side = side
  end

  def area
    @side ** 2
  end
end
```

Rectangle:

```ruby
require_relative 'shape.rb'

class Rectangle < Shape
  def initialize(base, height)
    @base = base
    @height = height

  end

  def area
    @height * @base
  end

end
```

Como se puede apreciar es el mismo código, pero ahora hemos definido cada forma como una `Shape` que es capaz de responder al mensaje `area` y decirnos qué superficie tiene.

## ¿Cómo evito el acoplamiento directo? Indirection

El objetivo del patrón Indirection es evitar el acoplamiento directo entre otros dos objetos, normalmente introduciendo un objeto intermediario. El objetivo es permitir que los objetos relacionados puedan evolucionar independientemente.

En los ejemplos anteriores hemos creado una clase `SurfaceCalculator` que puede calcular el área de figuras compuestas por triángulos y cuadrados. Sin embargo, ¿qué tenemos que hacer para que pueda procesar también rectángulos u otras figuras?. Tal como está ahora mismo, `SurfaceCalculator` está acoplada al tipo de figuras que conoce. Por el propio principio Creator, solo puede instanciar figuras conocidas y no puede instanciar figuras que no conoce salvo que la modifiquemos.

Algo así:

```ruby
require 'rspec'
require_relative '../src/surface_calculator.rb'

RSpec.describe 'Surface Calculator should' do
  it 'calculate surface joining different shapes' do
    surface = SurfaceCalculator.new
    surface.add_triangle 10, 20
    surface.add_square 5
    surface.add_rectangle 15, 10
    expect(surface.total).to eq(275)
  end
end
```

La cuestión es que podríamos evitar este acoplamiento interponiendo un mediador entre `SurfaceCalculator` y las figuras que soporta.

Por ejemplo, usando una clase `ShapeFactory` que se encargue de fabricar las formas para `SurfaceCalculator`, de modo que ya no necesita saber qué figuras concretas están disponibles. Algo como esto:

```ruby
require 'rspec'
require_relative '../src/surface_calculator.rb'
require_relative '../src/shape_factory.rb'

RSpec.describe 'Surface Calculator should' do

  it 'calculate surface joining different shapes provided by factory' do
    factory = ShapeFactory.new
    surface = SurfaceCalculator.new
    surface.add factory.make 'triangle', 10, 20
    surface.add factory.make 'square', 5
    surface.add factory.make 'rectangle', 15, 10
    expect(surface.total).to eq(275)
  end
end
```

Este es el código de `ShapeFactory`:

```ruby
require_relative './triangle'
require_relative './square'
require_relative './rectangle'

class ShapeFactory

  def make(shape, *param)
    case shape
    when 'triangle'
      Triangle.new param[0], param[1]
    when 'square'
      Square.new param[0]
    when 'rectangle'
      Rectangle.new param[0], param[1]
    else
      raise "Shape #{shape} not supported"
    end
  end
end
```

Gracias a esto, `SurfaceCalculator` queda más simplificado y puede evolucionar sin preocuparse de ninguna _forma_:

```ruby
class SurfaceCalculator
  def initialize
    @shapes = []
  end

  def add(shape)
    @shapes.append shape
  end

  def total
    @shapes.sum { |shape| shape.area }
  end
end
```

Podríamos discutir dos cosas:

`ShapeFactory` está acoplado a las formas, tiene que tener conocimiento de la que soporta y de las que no. Sin embargo esa es su responsabilidad: saber qué formas se pueden instanciar.

Por otro lado, ¿no hemos dicho que `SurfaceCalculator`, al agregar formas tendría que ser la `Creator`?. Ciertamente, pero en este caso, nada nos impediría hacer que `ShapeFactory` sea colaboradora de `SurfaceCalculator`, permitiéndonos cumplir de nuevo este principio:

```ruby
require 'rspec'
require_relative '../src/surface_calculator.rb'
require_relative '../src/shape_factory.rb'

RSpec.describe 'Surface Calculator should' do
  it 'calculate surface joining different shapes' do
    factory = ShapeFactory.new
    surface = SurfaceCalculator.new factory
    surface.add 'triangle', 10, 20
    surface.add 'square', 5
    surface.add 'rectangle', 15, 10
    expect(surface.total).to eq(275)
  end
end
```

De modo que ahora, `SurfaceCalculator` quedaría así:

```ruby
class SurfaceCalculator
  def initialize(factory)
    @factory = factory
    @shapes = []
  end

  def add(shape, *params)
    @shapes.append @factory.make shape, *params
  end

  def total
    @shapes.sum { |shape| shape.area }
  end
end
```

## ¿Quién gestiona las entradas al sistema? Controller

Controller es una clase queque representa el sistema en general o bien un caso de uso de la aplicación y que gestiona uno o más eventos del sistema enviados por la capa de UI, de la que no forma parte, encargándose de delegar en otros objetos del sistema para obtener la respuesta que debería devolver.

El uso correcto de Controller incluye la posibilidad de procesar eventos estrechamente relacionados, como pueden ser los distintos verbos en REST para un mismo recurso.

## ¿Cómo reduzco el acoplamiento? Low coupling

Ya [hemos hablado de acoplamiento en una entrega anterior](https://franiglesias.github.io/beyond-solid). El acoplamiento es el grado de interdependencia entre objetos. Si dos objetos deben interactuar existe un acoplamiento entre ellos, pero es deseable que sea el mínimo posible. La aplicación de otros principios nos permite mantener el acoplamiento reducido y relajado, como puede ser el uso de la indirección, o [la inversión de dependencias](https://franiglesias.github.io/beyond-solid-2).

## ¿Cómo focalizo las responsabilidades de una clase? High cohesion

En [un artículo anterior](https://franiglesias.github.io/beyond-solid) también mencionamos la cohesión. La cohesión es la fuerza que mantiene focalizadas las responsabilidades de una clase, o de un módulo en su caso. Una clase con alta cohesión es más fácil de entender, ya que no tiene elementos que nos puedan hacer dudar de cuáles son sus responsabilidades.

La alta cohesión se logra, sobre todo, aprendiendo a decir no a las responsabilidades que no corresponden a una clase, especialmente a aquellas que a primera vista sí parecerían adecuadas. También se contribuye a ella identificando aquellas cosas que cambian juntas, ya que, cuando es así, deberían ir juntas.

## ¿Dónde pongo la responsabilidad cuando no puedo asignarla a una clase específica? Pure fabrication

Se trata de una clase que no representa un concepto del dominio, pero que necesitamos cuando no podemos asignar la responsabilidad a una clase que sí lo hace. Este tipo de clase es lo que solemos entender como Servicio. Con frecuencia, los objetos mediadores que usamos al aplicar Indirection son Pure fabrication, también. Es decir, son artificios que introducimos para facilitarnos las cosas. El caso anterior con `ShapeFactory` puede servirnos como ejemplo.

## Referencias

- [GRASP (PDF)](https://www.cs.colorado.edu/~kena/classes/5448/f12/presentation-materials/rao.pdf)
- [GRASP, wikipedia](https://en.wikipedia.org/wiki/GRASP_(object-oriented_design))
- [GRASP – General Responsibility Assignment Software Patterns Explained](http://www.kamilgrzybek.com/design/grasp-explained/)
