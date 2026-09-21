---
title: "Object Calisthenics"
notion_id: f234441e-e7d5-4ded-9753-eb2a7c46f1e3
notion_url: https://app.notion.com/p/Object-Calisthenics-f234441ee7d54ded9753eb2a7c46f1e3
last_edited: 2023-04-25T15:02:00.000Z
source_url: https://www.notion.so/Object-Calisthenics-f234441ee7d54ded9753eb2a7c46f1e3
tags: ["Object Oriented Programming", "Article", "Guide", "The Talking Bit - Fran Iglesias", "Español"]
---
Esta es una recopilación de artículos publicados por [Fran Iglesias](https://github.com/franiglesias) en su blog [The Talking Bit](https://franiglesias.github.io/) relacionados con la práctica conocida como _Object Calisthenics._



La calistenia es una disciplina de entrenamiento físico que se basa en trabajar con el propio peso buscando desarrollar tanto fuerza como armonía y precisión en el movimiento.

En el campo del software, Jeff Bay introdujo una idea similar en un artículo de la publicación The ThoughtWorks Anthology: Essays on Software Technology and Innovation en 2008. Su propuesta consistía en nueve restricciones al escribir código que tendrán el efecto de mejorar su estructura y diseño.

Estas restricciones están pensadas para forzar ciertos buenos hábitos al escribir código orientado a objetos, contribuyendo a desarrollar un código de mejor calidad y a identificar los elementos que hacen bueno el diseño de un código. Por otro lado, es importante tener en cuenta que la mayoría de ellas tienen sentido en el paradigma de orientación a objetos, por lo que podrían no ser aplicables a otros paradigmas. Aun así, creo que se pueden aprovechar algunas.



La lista de restricciones es la siguiente:

- Un solo nivel de indentación por método
- No usar la palabra clave ELSE
- Encapsular todas las primitivas y strings
- Colecciones de primera clase
- Un punto por línea
- No usar abreviaturas
- Mantener todas las entidades pequeñas
- No más de dos variables de instancia por clase
- No usar getters/setters o propiedades públicas

# Índice

<!-- unsupported block: table_of_contents -->

---

# [Solo un nivel de indentación](https://franiglesias.github.io/calisthenics-1/)

Esta es bastante sencilla de entender, aunque puede que no tanto de aplicar.

La indentación nos ayuda a organizar visualmente el código de modo que cuando un fragmento está, por así decir, contenido en otro se muestra más adentrado en el cuerpo del texto. En Python, la indentación es lo que define los bloques de código mientras que en otros lenguajes estos bloques se definen usando algún tipo de marcador como las llaves, palabras clave como “begin/end”, etc.

El nivel de indentación está fuertemente asociado al nivel de abstracción. Con frecuencia, los bloques de código indentados suponen un cierto nivel de detalle que no se corresponde al nivel de abstracción del método que los contiene. La mezcla de niveles de abstracción hace que sea más difícil comprender el código debido a que tenemos que cambiar nuestro enfoque al entrar y salir de cada bloque.

Los bloques indentados aparecen en estructuras condiciones y en bucles. El problema de estos bloques surge cuando dentro de un bloque indentado aparece la necesidad de introducir una nueva condicional o bucle, resultando en una anidación que genera un nuevo nivel de indentación. Esto incrementa la mezcla de conceptos generales con detalles. Además, hace que tengamos que entrar y salir de distintas ramas del flujo de ejecución. En conjunto, el código así organizado se hace más difícil de leer, de comprender y de mantener en la cabeza.

Veamos un ejemplo no orientado a objetos, tomado de la kata [Theatrical Players](https://github.com/emilybache/Theatrical-Players-Refactoring-Kata) de Emily Bache:

```python
import math


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        if play['type'] == "tragedy":
            this_amount = 40000
            if perf['audience'] > 30:
                this_amount += 1000 * (perf['audience'] - 30)
        elif play['type'] == "comedy":
            this_amount = 30000
            if perf['audience'] > 20:
                this_amount += 10000 + 500 * (perf['audience'] - 20)

            this_amount += 300 * perf['audience']

        else:
            raise ValueError(f'unknown type: {play["type"]}')

        # add volume credits
        volume_credits += max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        if "comedy" == play["type"]:
            volume_credits += math.floor(perf['audience'] / 5)
        # print line for this order
        result += f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'
        total_amount += this_amount

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

En el ejemplo, se puede ver como el bucle `for` introduce un nivel de indentación en el código. Pero dentro de él podemos ver dos estructuras condicionales que añaden hasta dos nuevos niveles.

Para reducir a un solo nivel de indentación el código de un método lo más habitual es extraer la estructura anidada a un método privado, de manera que en su lugar quede una única línea con esa llamada.

Es este ejemplo, la función `statement` calcula el importe de una factura sobre varias actuaciones de una compañía de teatro. Para ello recorre la lista de actuaciones, calculando el importe de cada actuación basándose en características de la obra y de la audiencia y sumándolo todo. La primera estructura condicional contiene los detalles del cálculo del importe de cada actuación.

```python
import math


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            amount = 40000
            if perf['audience'] > 30:
                amount += 1000 * (perf['audience'] - 30)
        elif play['type'] == "comedy":
            amount = 30000
            if perf['audience'] > 20:
                amount += 10000 + 500 * (perf['audience'] - 20)

            amount += 300 * perf['audience']

        else:
            raise ValueError(f'unknown type: {play["type"]}')
        return amount

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)

        # add volume credits
        volume_credits += max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        if "comedy" == play["type"]:
            volume_credits += math.floor(perf['audience'] / 5)
        # print line for this order
        result += f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'
        total_amount += this_amount

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

Si los niveles de anidación son varios podemos empezar por el primero y luego nos vamos moviendo más hacia adentro. Esto choca con la recomendación de empezar a refactorizar por la rama más profunda, pero en este caso parece mejor despejar cada nivel de abstracción desde fuera hacia adentro. Además, el refactor automático de extraer método es lo bastante seguro como para poder hacerlo sin tests.

En nuestro ejemplo, el nuevo método contiene más de un nivel de indentación, pero volveremos a esto más adelante. Vamos a seguir aplanando la función `statement`.

Algunas dificultades que nos podemos encontrar tienen que ver con el uso de variables que se inicializan fuera de la estructura condicional, pero que se modifican en ella. Un paso previo recomendable es agrupar el código relacionado, por ejemplo, las líneas en las que se mencionan las mismas variables deberían ir juntas. De este modo, cuando vayamos a extraer la estructura condicional seremos más conscientes de esas dependencias.

La segunda estructura condicional hace lo que parece ser un cálculo de créditos o puntos para futuros espectáculos (`volume_credits`) y también podemos extraerlo. Como podemos ver, la condicional modifica el valor de una variable que se inicializaba fuera. Por tanto, incluimos todo en la extracción, teniendo en cuenta que `volume_credits` es una variable acumulativa:

```python
        # add volume credits
        volume_credits += max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        if "comedy" == play["type"]:
            volume_credits += math.floor(perf['audience'] / 5)
        # print line for this order
```

Nos quedaría así:

```python
import math


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_amount(perf, play)
        volume_credits += calculate_this_volume_credits(perf, play)
        # print line for this order
        result += f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'
        total_amount += this_amount

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result


def calculate_this_volume_credits(perf, play):
    # add volume credits
    volume_credits = max(perf['audience'] - 30, 0)
    # add extra credit for every ten comedy attendees
    if "comedy" == play["type"]:
        volume_credits += math.floor(perf['audience'] / 5)
    return volume_credits


def calculate_amount(perf, play):
    if play['type'] == "tragedy":
        this_amount = 40000
        if perf['audience'] > 30:
            this_amount += 1000 * (perf['audience'] - 30)
    elif play['type'] == "comedy":
        this_amount = 30000
        if perf['audience'] > 20:
            this_amount += 10000 + 500 * (perf['audience'] - 20)

        this_amount += 300 * perf['audience']

    else:
        raise ValueError(f'unknown type: {play["type"]}')
    return this_amount
```

Un truco simple cuando usas refactor automático es examinar los parámetros que necesita el nuevo método, ya que el análisis que hace la herramienta de refactor identificará todos los necesarios. Esto nos ayuda a descubrir variables temporales que tal vez sean innecesarias o que deberían estar únicamente en el método extraído. Especialmente en el caso de se tenga que devolver su valor.

En este caso, tenemos que separar el cálculo parcial del total. Esta es la secuencia de pasos que he seguido para hacerlo manteniendo los tests en verde en todos los pasos:

En primer lugar, voy a introducir una variable `performance_credits` que contendrá el cálculo parcial:

```python
        # add volume credits
        performance_credits = max(perf['audience'] - 30, 0)
        volume_credits += performance_credits
        # add extra credit for every ten comedy attendees
        if "comedy" == play["type"]:
            volume_credits += math.floor(perf['audience'] / 5)
```

`volumen_credits` solo debería actualizarse cuando se haya completado el cálculo parcial, así que lo muevo al final del fragmento:

```python
        # add volume credits
        performance_credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        if "comedy" == play["type"]:
            volume_credits += math.floor(perf['audience'] / 5)
        volume_credits += performance_credits
```

En la condicional, actualizo `performance_credits` en lugar de `volume_credits`:

```python
        # add volume credits
        performance_credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        if "comedy" == play["type"]:
            performance_credits += math.floor(perf['audience'] / 5)
        volume_credits += performance_credits
```

Ahora ya puedo extraer el cálculo limpiamente:

```python
        performance_credits = calculate_performance_credits(perf, play)
        volume_credits += performance_credits
```

Si te fijas en la parte principal del cuerpo de la función `statement` verás que es mucho más claro y es fácil entender lo que ocurre en un nivel general. Basta con moverse a la función adecuada para poder acceder a los detalles de cada cálculo.

Queda más o menos así, una vez ordenadas las líneas:

```python
    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)

        volume_credits += performance_credits
        # print line for this order
        result += f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'
        total_amount += this_amount
```

Podría argumentarse que para poder extraer las estructuras condicionales y aplanar la indentación tengo que hacer algunos refactors de más. Pero precisamente ese es uno de los beneficios de intentar forzar la regla. Tengo que mejorar la organización del código para tener las condiciones adecuadas que me permitan aplicar la regla de un solo nivel de indentación.

Volvamos ahora a los niveles extra de indentación que aún no hemos tratado. Se han movido todos a la función `calculate_amount`:

```python
def calculate_performance_amount(perf, play):
    if play['type'] == "tragedy":
        this_amount = 40000
        if perf['audience'] > 30:
            this_amount += 1000 * (perf['audience'] - 30)
    elif play['type'] == "comedy":
        this_amount = 30000
        if perf['audience'] > 20:
            this_amount += 10000 + 500 * (perf['audience'] - 20)

        this_amount += 300 * perf['audience']

    else:
        raise ValueError(f'unknown type: {play["type"]}')
    return this_amount
```

Lo más fácil es mover las _patas_ de las condicionales a sus propios métodos, como se puede ver a continuación.

```python
    def calculate_performance_amount(perf, play):
    if play['type'] == "tragedy":
        amount = calculate_amount_for_tragedy(perf)
    elif play['type'] == "comedy":
        amount = calculate_amount_for_comedy(perf)

    else:
        raise ValueError(f'unknown type: {play["type"]}')
    return amount

def calculate_amount_for_comedy(perf):
    amount = 30000
    if perf['audience'] > 20:
        amount += 10000 + 500 * (perf['audience'] - 20)
    amount += 300 * perf['audience']
    return amount

def calculate_amount_for_tragedy(perf):
    amount = 40000
    if perf['audience'] > 30:
        amount += 1000 * (perf['audience'] - 30)
    return amount
```

Al fijarnos en el resultado, podemos observar varias cosas. Una de ellas es que el código de `calculate_amount` sugiere aplicar el patrón _early return_, que clarifica más aún el cuerpo del método, así como suprimir la palabra clave `ELSE`, tema que trataríamos en la siguiente regla. También nos abre la puerta a usar una estructura `switch/case`. Pero si profundizamos, también sugiere fuertemente la posibilidad de introducir orientación a objetos para beneficiarnos del polimorfismo. No lo vamos a hacer en esta ocasión, porque el objetivo del artículo es centrarnos en una regla cada vez.

Este es el resultado hasta el momento. En cada función tenemos un solo nivel de indentación. No es el refactor definitivo, pero ha mejorado sustancialmente la organización y legibilidad del código.

```python
import math


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            amount = calculate_amount_for_tragedy(perf)
        elif play['type'] == "comedy":
            amount = calculate_amount_for_comedy(perf)

        else:
            raise ValueError(f'unknown type: {play["type"]}')
        return amount

    def calculate_amount_for_comedy(perf):
        amount = 30000
        if perf['audience'] > 20:
            amount += 10000 + 500 * (perf['audience'] - 20)
        amount += 300 * perf['audience']
        return amount

    def calculate_amount_for_tragedy(perf):
        amount = 40000
        if perf['audience'] > 30:
            amount += 1000 * (perf['audience'] - 30)
        return amount

    def calculate_performance_credits(perf, play):
        # add volume credits
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        if "comedy" == play["type"]:
            credits += math.floor(perf['audience'] / 5)
        return credits

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)

        volume_credits += performance_credits
        # print line for this order
        result += f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'
        total_amount += this_amount

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

A poco que profundicemos en este ejemplo podemos sentir que está pidiendo a gritos aplicar orientación a objetos, pero precisamente eso es algo que podemos empezar a vislumbrar después de haber aplicado esta simple regla.

## Por qué funciona

Esta regla funciona porque nos ayuda a conseguir que cada método desarrolle su trabajo en un único nivel de abstracción, a la vez que separamos distintas responsabilidades. De este modo, puedes leer cada método y entender qué pasa, sin necesidad de distraerte con detalles que no son relevantes en ese momento. Si necesitas conocer cómo se implementa alguna de las fases no tienes más que revisar el método que se ocupa de ello.

Al separar el comportamiento de ese objeto en pasos implementados por métodos específicos será más fácil también identificar el papel de los colaboradores del objeto, si los hay, así como su aislamiento. De este modo, podremos detectar y solucionar más fácilmente posibles problemas de acoplamiento. A su vez, esta extracción a métodos privados puede ser el primer paso para identificar diferentes responsabilidades en una clase que podrían extraerse a nuevas clases.

Dado que son métodos privados no estamos afectando a la interfaz pública.

## Más allá

### Separar iterador de iteración

Una forma de abordar los bucles es separar el iterador (el bucle) de la iteración (el cuerpo del bucle). Es decir, en lugar de tener un bloque de código, extraemos la totalidad del bloque a un método privado. De este modo, el cuerpo del bucle contendría una sola línea. Una de las ventajas de proceder así es que puede ayudarnos a identificar código que realmente pertenece a la clase del objeto que está siendo procesado en la iteración.

Aplicar esta separación en este ejemplo puede ser un poco complicado, dado que en el bucle for vamos acumulando ni más ni menos que tres variables: `total_amount`, `volume_credits` y `result`.

```python
    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)

        volume_credits += performance_credits
        # print line for this order
        result += f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'
        total_amount += this_amount
```

Vamos a ver si podemos hacer algo al respecto. Lo primero sería extraer una variable para almancenar la línea que estamos calculando en cada iteración:

```python
    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)

        volume_credits += performance_credits
        # print line for this order
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'
        result += line
        total_amount += this_amount
```

Y ahora reunimos las variables acumuladoras:

```python
    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        total_amount += this_amount
        volume_credits += performance_credits
```

Ahora podríamos intentar extraer la parte del cálculo a una nueva función. Sin embargo, en Python las _inner functions_ no son visibles desde dentro de otras _inner functions_, así que tendríamos que pasarlas junto con los parámetros necesarios por lo que este paso no es muy viable. De nuevo, el refactor nos va mostrando que lo más efectivo sería introducir orientación a objetos para este caso.

### Sólo un `if` por método

En el libro [Five lines of code](https://www.manning.com/books/five-lines-of-code), Christian Clausen propone llevar esta restricción un poco más allá. Además de que cada método tenga un único nivel de indentación, sugiere que solo haya una estructura condicional en cada método y que `if` debería ser siempre la primera línea.

Vamos a ver algunos ejemplos en este código y cómo se podrían abordar.

En el primero, podemos ver que hay un extra si la audiencia supera un cierto umbral. En caso contrario, no se incrementa.

```python
    def calculate_amount_for_comedy(perf):
        amount = 30000
        if perf['audience'] > 20:
            amount += 10000 + 500 * (perf['audience'] - 20)
        amount += 300 * perf['audience']
        return amount
```

Podríamos hacer una modificación temporal para verlo más claro:

```python
    def calculate_amount_for_comedy(perf):
        amount = 30000
        if perf['audience'] > 20:
            extra_for_high_audience = 10000 + 500 * (perf['audience'] - 20)
        else:
            extra_for_high_audience = 0
        amount += extra_for_high_audience
        amount += 300 * perf['audience']
        return amount
```

Ahora podemos extraer el bloque condicional:

```python
    def calculate_amount_for_comedy(perf):
        amount = 30000
        extra_for_high_audience = extra_amount_for_high_audience_in_comdey(perf)
        amount += extra_for_high_audience
        amount += 300 * perf['audience']
        return amount

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] > 20:
            extra_for_high_audience = 10000 + 500 * (perf['audience'] - 20)
        else:
            extra_for_high_audience = 0
        return extra_for_high_audience
```

Esto hace que la condición quede como primera línea en `extra_amount_for_high_audience_in_comedy`, que es lo que buscábamos. Ahora limpiamos un poco el código para que quede menos redundante, removiendo variables temporales innecesarias.

```python
    def calculate_amount_for_comedy(perf):
        amount = 30000
        amount += extra_amount_for_high_audience_in_comedy(perf)
        amount += 300 * perf['audience']
        return amount

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] > 20:
            return 10000 + 500 * (perf['audience'] - 20)
        else:
            return 0
```

Podemos aplicar un tratamiento similar para otros tipos de obras. El resultado sería este:

```python
    def calculate_amount_for_tragedy(perf):
        amount = 40000
        amount += extra_amount_for_high_audience_in_tragedy(perf)
        return amount

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] > 30:
            return 1000 * (perf['audience'] - 30)
        else:
            return 0
```

Y lo mismo en el cálculo de créditos:

```python
    def calculate_performance_credits(perf, play):
        # add volume credits
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        if "comedy" == play["type"]:
            credits += math.floor(perf['audience'] / 5)
        return credits
```

Que quedaría así:

```python
    def calculate_performance_credits(perf, play):
        # add volume credits
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        credits += extra_volume_credits_for_comedy(perf, play)
        return credits

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" == play["type"]:
            return math.floor(perf['audience'] / 5)
        else:
            return 0
```

## El resultado

Y este es el resultado final después de aplicar la regla de un solo nivel de indentación y las reglas extra:

```python
import math


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            amount = calculate_amount_for_tragedy(perf)
        elif play['type'] == "comedy":
            amount = calculate_amount_for_comedy(perf)
        else:
            raise ValueError(f'unknown type: {play["type"]}')
        return amount

    def calculate_amount_for_comedy(perf):
        amount = 30000
        amount += extra_amount_for_high_audience_in_comedy(perf)
        amount += 300 * perf['audience']
        return amount

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] > 20:
            return 10000 + 500 * (perf['audience'] - 20)
        else:
            return 0

    def calculate_amount_for_tragedy(perf):
        amount = 40000
        amount += extra_amount_for_high_audience_in_tragedy(perf)
        return amount

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] > 30:
            return 1000 * (perf['audience'] - 30)
        else:
            return 0

    def calculate_performance_credits(perf, play):
        # add volume credits
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        credits += extra_volume_credits_for_comedy(perf, play)
        return credits

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" == play["type"]:
            return math.floor(perf['audience'] / 5)
        else:
            return 0

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        total_amount += this_amount
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

Como puedes comprobar, se han introducido muchos `else`, lo que nos llevará a aplicar una nueva regla. Pero eso será en otra entrega.

---

# [No usar la palabra clave ELSE](https://franiglesias.github.io/calisthenics-2/)

Una estructura condicional puede dar lugar a varias ramas en el flujo de ejecución de modo que si se cumple la condición se sigue un camino, y si no se cumple… pues se sigue otro. O simplemente no se sigue ninguno y se continúa con la siguiente instrucción.

Pero, ¿qué problema hay con `else`? Al fin y al cabo, no indica otra cosa que seguir unas instrucciones específicas para el caso de que no se cumplan las condiciones requeridas en el `if`. Normalmente, el problema no es el hecho de usar `else` per se, sino el contexto en el que lo usamos o la organización de código que se genera usándolo. Se podría decir que utilizar `else` puede ser un _smell_, un síntoma de que algo podría estar mejor diseñado. Si nos obligamos a eliminarlo, podemos mejorar el código.

## Condicionales sencillas, aún más sencillas

Una estructura `if` tiene este aspecto:

```python
// instrucciones previas

if (condicion) then
   // instrucciones si se cumple la condición

// instrucciones posteriores
```

Usamos `else` cuando queremos ejecutar ciertas instrucciones en caso de no cumplirse la condición, de forma alternativa.

```python
// instrucciones previas

if (condicion) then
   // instrucciones si se cumple la condición
else
   // instrucciones alternativas si no se cumple

// instrucciones posteriores
```

Esta estructura ya podría introducir algo de ruido a la hora de leer el programa. Como vimos en el artículo anterior, nos interesa forzar un solo nivel de indentación como máximo para evitar la sobrecarga de seguir el código anidado. La introducción de `else` no añade un nivel de indentación extra, pero implica que tenemos que mantener en la cabeza dos flujos alternativos.

Esto se complica si tenemos que hacer seguimiento de variables que son inicializadas fuera de la estructura condicional, pero manipuladas en ella. También se complica la lectura si el tamaño de uno de los bloques es muy grande, ya que podría ofuscar el otro.

Por esa razón, se recomendaba aislar la estructura condicional en un método o función, de modo que el `if` fuese la primera línea y únicamente hubiese una condicional en ese método:

```python
// instrucciones previas

// instrucciones cuyo resultado depende de una condición

// instrucciones posteriores
```

```python
if (condicion) then
   // instrucciones si se cumple la condición
else
   // instrucciones alternativas si no se cumple
   
return
```

Al aislar de esta manera las condicionales, tanto la rama del `if` como la del `else` retornarán al punto de llamada, ya que no hay más instrucciones que seguir. De hecho, podríamos retornar desde ambas ramas. Es lo que conocemos como patrón `return early`,

```python
if (condicion) then
   // instrucciones si se cumple la condición
   return
else
   // instrucciones alternativas si no se cumple
   return
```

Esto hace redundante la palabra clave `else`, ya que no es necesario asegurar que la condición del if no se cumple.

```python
if (condicion) then
   // instrucciones si se cumple la condición
   return

// instrucciones alternativas si no se cumple
return
```

Retomando el ejemplo del artículo anterior, tenemos varias situaciones en las que se usa else que podríamos examinar. Recordemos el código:

```python
import math


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            amount = calculate_amount_for_tragedy(perf)
        elif play['type'] == "comedy":
            amount = calculate_amount_for_comedy(perf)
        else:
            raise ValueError(f'unknown type: {play["type"]}')
        return amount

    def calculate_amount_for_comedy(perf):
        amount = 30000
        amount += extra_amount_for_high_audience_in_comedy(perf)
        amount += 300 * perf['audience']
        return amount

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] > 20:
            return 10000 + 500 * (perf['audience'] - 20)
        else:
            return 0

    def calculate_amount_for_tragedy(perf):
        amount = 40000
        amount += extra_amount_for_high_audience_in_tragedy(perf)
        return amount

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] > 30:
            return 1000 * (perf['audience'] - 30)
        else:
            return 0

    def calculate_performance_credits(perf, play):
        # add volume credits
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        credits += extra_volume_credits_for_comedy(perf, play)
        return credits

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" == play["type"]:
            return math.floor(perf['audience'] / 5)
        else:
            return 0

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        total_amount += this_amount
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

Aquí tenemos un ejemplo:

```python
    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] > 20:
            return 10000 + 500 * (perf['audience'] - 20)
        else:
            return 0
```

Este caso es bastante sencillo porque el else es redundante:

```python
    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] > 20:
            return 10000 + 500 * (perf['audience'] - 20)
        
        return 0
```

Tenemos formas alternativas. Una de ellas consiste en usar el operador ternario, que funciona especialmente bien cuando queremos expresar un cálculo que se realiza de maneras diferentes.

```python
    def extra_amount_for_high_audience_in_comedy(perf):
        return 10000 + 500 * (perf['audience'] - 20) if perf['audience'] > 20 else 0
```

Otra forma de hacerlo es invertir la condición, dejando el caso residual como una cláusula de guarda. Es especialmente aplicable si se trata de verificar precondiciones de los parámetros que llegan al método o función. De esta forma, centras la atención en la rama más significativa.

```python
    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return 0
        
        return 10000 + 500 * (perf['audience'] - 20)
```

Cualquiera de las tres técnicas te permite suprimir el `else`. La más adecuada dependerá del aspecto que necesites acentuar. Para este ejemplo podrían funcionar las tres bastante bien y resulta difícil decidirse por una de ellas. Quizá en este caso optaría por la condicional invertida.

De este modo, las tres funciones que contienen condicionales simples quedarían así:

```python
    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return 0

        return 10000 + 500 * (perf['audience'] - 20)

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] <= 30:
            return 0

        return 1000 * (perf['audience'] - 30)

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" != play["type"]:
            return 0

        return math.floor(perf['audience'] / 5)
```

## Condicionales complejas

Tenemos otro ejemplo interesante aquí:

```python
    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            amount = calculate_amount_for_tragedy(perf)
        elif play['type'] == "comedy":
            amount = calculate_amount_for_comedy(perf)
        else:
            raise ValueError(f'unknown type: {play["type"]}')
        return amount
```

Se trata de una serie de condicionales encadenadas a través de la clave `else` o `else if`. La estructura condicional maneja un cierto número de condiciones de tal manera, que si no se cumple la inicial, tenemos que verificar si se cumplen otras y actuar en consecuencia.

Esta estructura encadenada se entiende mejor usando `switch`, lo que esconde el `else`, aunque realmente no lo elimina. Sin embargo, Python no tiene `switch` por lo que no incluyo el ejemplo.

De nuevo, podremos usar _return early_ para simplificar la estructura. Primero introducimos el `return`.

```python
    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        elif play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)
        else:
            raise ValueError(f'unknown type: {play["type"]}')
```

Y a continuación, eliminamos los `else`:

```python
    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)
        
        raise ValueError(f'unknown type: {play["type"]}')
```

## Por qué funciona

Eliminar `else` nos obliga a pensar bien nuestras estructuras condicionales. Una estructura condicional siempre hace al menos dos cosas: decidir si se cumple la condición, hacer algo si es así. En el caso de else, hay que añadir una tercera cosa: la acción alternativa.

De hecho, en orientación a objetos, la mera presencia de una estructura condicional puede significar un problema de diseño. Esto ocurre, por ejemplo, cuando la condicional verifica alguna propiedad de un objeto (o de algún concepto del programa que potencialmente pueda ser un objeto). En ese caso, se pone de manifiesto la necesidad de polimorfismo. Nuestro último refactor elo deja muy claro.

Cuando se toma una decisión basada en el tipo de un concepto, debería abordarse mediante polimorfismo.

```python
    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)
        
        raise ValueError(f'unknown type: {play["type"]}')
```

Sin embargo, cuando la decisión se basa en un valor, podríamos recurrir a otros enfoques

```python
    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] <= 30:
            return 0

        return 1000 * (perf['audience'] - 30)
```

De todos modos, la introducción de la orientación a objetos vendrá de la mano de las siguientes reglas, que consisten en empaquetar todas nuestras primitivas y colecciones en objetos. Es decir, representar los conceptos usando objetos.

## El resultado

Después de eliminar la palabra clave `else`, el código queda así:

```python
import math


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)

        raise ValueError(f'unknown type: {play["type"]}')

    def calculate_amount_for_comedy(perf):
        amount = 30000
        amount += extra_amount_for_high_audience_in_comedy(perf)
        amount += 300 * perf['audience']
        return amount

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return 0

        return 10000 + 500 * (perf['audience'] - 20)

    def calculate_amount_for_tragedy(perf):
        amount = 40000
        amount += extra_amount_for_high_audience_in_tragedy(perf)
        return amount

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] <= 30:
            return 0

        return 1000 * (perf['audience'] - 30)

    def calculate_performance_credits(perf, play):
        # add volume credits
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        credits += extra_volume_credits_for_comedy(perf, play)
        return credits

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" != play["type"]:
            return 0

        return math.floor(perf['audience'] / 5)

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        total_amount += this_amount
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

Parece muy claro que conceptos como obra (play) y actuación (performance) están pugnando por salir. Y alguno más. Lo veremos en el artículo siguiente.

---

# [Encapsular todas las primitivas y strings](https://franiglesias.github.io/calisthenics-3/)

Los lenguajes de programación proporcionan tipos de datos básicos, llamados primitivos, con los que podemos representar los diversos conceptos que maneja un programa. Sin embargo, esta representación suele ser imperfecta.

Pensemos por ejemplo, en un precio. El precio se puede representar con un número, pero hay varias características de los números con los que representamos precios que son importantes: son valores positivos, tienen decimales con reglas específicas de redondeo, y suele ser importante conocer la unidad monetaria, entre otros detalles.

Estas características no las proveen los tipos numéricos primitivos habituales. Por esa razón, un tipo específico, que puede estar basado en uno primitivo, pero que encapsule esas reglas es mucho mejor solución. Basta con realizar una encapsulación básica para empezar a obtener beneficios, ya que eso oculta al resto del programa los detalles de implementación del tipo y nos permite que evolucione sin afectar al resto del código. A medida que introducimos comportamiento y validaciones en ese objeto, el programa se beneficia automáticamente.

## Encapsular tipos primitivos simples

Así que volvamos a nuestro ejemplo, en el que tenemos un montón de posibles casos. Para empezar, nos encontramos con los parámetros que se pasan a la función `statement`. Estos nos presentan algunos problemas particulares porque contienen colecciones de cosas, así que vamos a dejarlo para la próxima regla.

Lo primero que nos encontramos es `total_amount`, que representa el importe de la factura y que va acumulando parciales.

```python
import math


def statement(invoice, plays):
    total_amount = 0
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)

        raise ValueError(f'unknown type: {play["type"]}')

    def calculate_amount_for_comedy(perf):
        amount = 30000
        amount += extra_amount_for_high_audience_in_comedy(perf)
        amount += 300 * perf['audience']
        return amount

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return 0

        return 10000 + 500 * (perf['audience'] - 20)

    def calculate_amount_for_tragedy(perf):
        amount = 40000
        amount += extra_amount_for_high_audience_in_tragedy(perf)
        return amount

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] <= 30:
            return 0

        return 1000 * (perf['audience'] - 30)

    def calculate_performance_credits(perf, play):
        # add volume credits
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        credits += extra_volume_credits_for_comedy(perf, play)
        return credits

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" != play["type"]:
            return 0

        return math.floor(perf['audience'] / 5)

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        total_amount += this_amount
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

`total_amount` representa una cantidad de dinero. En este ejemplo, la unidad monetaria resulta ser el centavo como se puede apreciar en la forma en que se usa la función `format_as_dollars`. Básicamente, necesitaremos dos comportamientos: poder acumular y obtener el importe acumulado hasta el momento. Este nuevo tipo se podría llamar `Amount`.

Para no tener los test rotos mucho tiempo voy a introducir el cambio en paralelo, añadiendo la nueva clase, pero sin introducir el cambio hasta el último momento. Por supuesto, puedo hacer esto con TDD.

```python
import unittest


class AmountTestCase(unittest.TestCase):
    def test_contains_amount(self):
        amount_of_300 = Amount(300)
        self.assertEqual(300, amount_of_300.current())


if __name__ == '__main__':
    unittest.main()
```

```python
class Amount:
    def __init__(self, initial_amount):
        self._amount = initial_amount

    def current(self):
        return self._amount
```

Ahora, añadiré un método para acumular importes. Aprovecharé para hacerlo inmutable.

```python
import unittest

from domain.amount import Amount


class AmountTestCase(unittest.TestCase):
    def test_contains_amount(self):
        amount_of_300 = Amount(300)
        self.assertEqual(300, amount_of_300.current())

    def test_can_accumulate_partial_amounts(self):
        amount_of_300 = Amount(300)
        amount_of_500 = amount_of_300.add(Amount(200))
        self.assertEqual(500, amount_of_500.current())


if __name__ == '__main__':
    unittest.main()
```

```python
class Amount:
    def __init__(self, initial_amount):
        self._amount = initial_amount

    def current(self):
        return self._amount

    def add(self, other):
        new_amount = self._amount + other.current()
        return Amount(new_amount)
```

Con esto tengo suficiente para empezar a usarlo. Para ello introduzco una variable `invoice_amount`.

```python
import math

from domain.amount import Amount


def statement(invoice, plays):
    total_amount = 0
    invoice_amount = Amount(0)
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    # ...

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        total_amount += this_amount
        invoice_amount = invoice_amount.add(Amount(this_amount))
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(total_amount/100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

Y, finalmente, solo tendría que reemplazar la variable `total_amount` en la línea que imprime el importe final. La idea es que todos los cambios anteriores ya estén mezclados, de modo que este nuevo cambio ocurra en un único commit y se pueda revertir fácilmente en caso de que falle.

```python
import math

from domain.amount import Amount


def statement(invoice, plays):
    total_amount = 0
    invoice_amount = Amount(0)
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    # ...

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        total_amount += this_amount
        invoice_amount = invoice_amount.add(Amount(this_amount))
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(invoice_amount.current()//100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

De hecho, los tests de statement siguen pasando perfectamente, por lo que podemos quitar `total_amount` ya que ha dejado de usarse.

```python
import math

from domain.amount import Amount


def statement(invoice, plays):
    invoice_amount = Amount(0)
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    # ...

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        invoice_amount = invoice_amount.add(Amount(this_amount))
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(invoice_amount.current()//100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

Una cosa interesante es que `this_amount` también debería ser un `Amount`, así que tendría sentido examinar `calculate_performance_amount`, para que devuelva un tipo `Amount`. Eso nos lleva a una serie de cambios con una mecánica muy similar a la que hemos seguido. Introducimos el código nuevo en paralelo y lo consolidamos en un commit. Finalmente, usamos el nuevo cálculo en un único commit para que deshacerlo sea sencillo. Una vez que confirmamos que no se ha roto nada, eliminamos el código viejo.

En este caso, lo que voy a hacer es introducir un objeto Amount en las funciones que realizan el cálculo y, provisionalmente, dejaré que todavía no devuelvan el tipo `Amount`, sino el primitivo calculado con Amount. En una segunda fase adaptaré el código llamante para que espere el tipo `Amount`. He aquí un ejemplo:

```python
    def calculate_amount_for_comedy(perf):
        amount = 30000
        amount += extra_amount_for_high_audience_in_comedy(perf)
        amount += 300 * perf['audience']
        return amount
```

Primer paso es calcularlo en paralelo:

```python
    def calculate_amount_for_comedy(perf):
        base_amount = Amount(30000)
        amount_with_extra = base_amount.add(Amount(extra_amount_for_high_audience_in_comedy(perf)))
        comedy_amount = amount_with_extra.add(Amount(300 * perf['audience']))
        
        amount = 30000
        amount += extra_amount_for_high_audience_in_comedy(perf)
        amount += 300 * perf['audience']
        return amount
```

Una vez hecho un commit con esos cambios, utilizaré el nuevo cálculo, pero sin devolver todavía el objeto:

```python
    def calculate_amount_for_comedy(perf):
        base_amount = Amount(30000)
        amount_with_extra = base_amount.add(Amount(extra_amount_for_high_audience_in_comedy(perf)))
        comedy_amount = amount_with_extra.add(Amount(300 * perf['audience']))

        amount = 30000
        amount += extra_amount_for_high_audience_in_comedy(perf)
        amount += 300 * perf['audience']
        return comedy_amount.current()
```

Como los tests siguen pasando puedo consolidar el cambio y eliminar el código que ya no uso.

```python
    def calculate_amount_for_comedy(perf):
        base_amount = Amount(30000)
        amount_with_extra = base_amount.add(Amount(extra_amount_for_high_audience_in_comedy(perf)))
        comedy_amount = amount_with_extra.add(Amount(300 * perf['audience']))

        return comedy_amount.current()
```

Por supuesto, puedo evitar el uso de variables temporales:

```python
    def calculate_amount_for_comedy(perf):
        return Amount(30000)\
            .add(Amount(extra_amount_for_high_audience_in_comedy(perf)))\
            .add(Amount(300 * perf['audience']))\
            .current()
```

El mismo cambio se puede aplicar en muchos lugares. Usaremos el mismo procedimiento, aunque no lo voy a mostrar para no alargar el artículo innecesariamente. Así es como quedará, teniendo en cuenta que todavía estoy dejando que las funciones retornen el primitivo.

```python
import math

from domain.amount import Amount


def statement(invoice, plays):
    invoice_amount = Amount(0)
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)

        raise ValueError(f'unknown type: {play["type"]}')

    def calculate_amount_for_comedy(perf):
        return Amount(30000)\
            .add(Amount(extra_amount_for_high_audience_in_comedy(perf)))\
            .add(Amount(300 * perf['audience']))\
            .current()

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return Amount(0).current()

        return Amount(10000 + 500 * (perf['audience'] - 20)).current()

    def calculate_amount_for_tragedy(perf):
        return Amount(40000)\
            .add(Amount(extra_amount_for_high_audience_in_tragedy(perf)))\
            .current()


    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] <= 30:
            return Amount(0).current()

        return Amount(1000 * (perf['audience'] - 30)).current()

    def calculate_performance_credits(perf, play):
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        credits += extra_volume_credits_for_comedy(perf, play)
        return credits

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" != play["type"]:
            return 0

        return math.floor(perf['audience'] / 5)

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount/100)} ({perf["audience"]} seats)\n'

        result += line
        invoice_amount = invoice_amount.add(Amount(this_amount))
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(invoice_amount.current()//100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

Ahora iré desde dentro hacia afuera cambiando el tipo retornado para dejar de usar el primitivo. Iré paso a paso, para asegurarme de que lo hago bien pasando los tests cada vez. Este es el primero:

```python
    def calculate_amount_for_comedy(perf):
        return Amount(30000)\
            .add(extra_amount_for_high_audience_in_comedy(perf))\
            .add(Amount(300 * perf['audience']))\
            .current()

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (perf['audience'] - 20))
```

Y sigo paso a paso hasta que los cambio todos. La idea es que solamente use `Amount::current` cuando sea necesario para imprimir la factura.

```python
import math

from domain.amount import Amount


def statement(invoice, plays):
    invoice_amount = Amount(0)
    volume_credits = 0
    result = f'Statement for {invoice["customer"]}\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)

        raise ValueError(f'unknown type: {play["type"]}')

    def calculate_amount_for_comedy(perf):
        return Amount(30000) \
            .add(extra_amount_for_high_audience_in_comedy(perf)) \
            .add(Amount(300 * perf['audience']))

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (perf['audience'] - 20))

    def calculate_amount_for_tragedy(perf):
        return Amount(40000) \
            .add(extra_amount_for_high_audience_in_tragedy(perf))

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] <= 30:
            return Amount(0)

        return Amount(1000 * (perf['audience'] - 30))

    def calculate_performance_credits(perf, play):
        credits = max(perf['audience'] - 30, 0)
        # add extra credit for every ten comedy attendees
        credits += extra_volume_credits_for_comedy(perf, play)
        return credits

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" != play["type"]:
            return 0

        return math.floor(perf['audience'] / 5)

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({perf["audience"]} seats)\n'

        result += line
        invoice_amount = invoice_amount.add(this_amount)
        volume_credits += performance_credits

    result += f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n'
    result += f'You earned {volume_credits} credits\n'
    return result
```

Como se puede ver, el código no ha cambiado demasiado y seguramente hay espacio para muchas mejoras, pero tenemos que proceder de manera sistemática. Así que vamos a buscar otro primitivo que podamos reemplazar.

`volume_credits` tiene un funcionamiento similar a `Amount`, pero significa una cosa distinta, así que vamos a introducir una clase `Credits`, que representará ese contexto. Y usaremos la misma aproximación: introducir la nueva clase, usarla en paralelo y, finalmente, sustituirla.

```python
import unittest

from domain.credits import Credits


class CreditsTestCase(unittest.TestCase):
    def test_contains_credits(self):
        self.assertEqual(100, Credits(100).current())  # add assertion here

    def test_accumulates_credits(self):
        initial_credits = Credits(100)
        extra = Credits(100)
        self.assertEqual(200, initial_credits.add(extra).current())


if __name__ == '__main__':
    unittest.main()
```

```python
class Credits:

    def __init__(self, initial_credits):
        self._credits = initial_credits

    def current(self):
        return self._credits

    def add(self, more_credits):
        return Credits(self._credits + more_credits.current())
```

Los cambios en el código los hacemos de la misma manera que antes. El resultado será más o menos este:

```python
import math

from domain.amount import Amount
from domain.credits import Credits


def statement(invoice, plays):
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    result = f'Statement for {invoice["customer"]}\n'

    # ...

    def calculate_performance_credits(perf, play):
        return Credits(max(perf['audience'] - 30, 0)).\
            add(extra_volume_credits_for_comedy(perf, play))


    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" != play["type"]:
            return Credits(0)

        return Credits(math.floor(perf['audience'] / 5))

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)
        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({perf["audience"]} seats)\n'

        result += line
        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    result += f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n'
    result += f'You earned {volume_credits.current()} credits\n'
    return result
```

Nuestro siguiente candidato es `result`, que es un string que va acumulando las líneas que se imprimirán en la factura. De hecho, podríamos incorporar el concepto de `Printer` como objeto encargado de _imprimir_ las líneas que se le pasan, en lugar de un simple almacén de líneas para devolver al final. Queremos que funcione más o menos como indica este test:

```python
import unittest


class PrinterTestCase(unittest.TestCase):
    def test_can_print_lines(self):
        printer = Printer()

        printer.print("Line 1")
        printer.print("Line 2")

        expected = "Line 1Line 2"

        self.assertEqual(expected, printer.output())


if __name__ == '__main__':
    unittest.main()
```

De momento lo implementamos así, que es más o menos como está en el código original y será suficiente para lo que necesitamos:

```python
class Printer:
    def __init__(self):
        self._lines = ""

    def print(self, line):
        self._lines += line

    def output(self):
        return self._lines
```

Par integrarlo, procedemos del mismo modo que antes. Primero lo introducimos en paralelo y dejamos que el último cambio sea muy simple. El resultado, una vez eliminado el código anterior es este:

```python
import math

from domain.amount import Amount
from domain.credits import Credits
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    printer.print(f'Statement for {invoice["customer"]}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)

        raise ValueError(f'unknown type: {play["type"]}')

    def calculate_amount_for_comedy(perf):
        return Amount(30000) \
            .add(extra_amount_for_high_audience_in_comedy(perf)) \
            .add(Amount(300 * perf['audience']))

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (perf['audience'] - 20))

    def calculate_amount_for_tragedy(perf):
        return Amount(40000) \
            .add(extra_amount_for_high_audience_in_tragedy(perf))

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] <= 30:
            return Amount(0)

        return Amount(1000 * (perf['audience'] - 30))

    def calculate_performance_credits(perf, play):
        return Credits(max(perf['audience'] - 30, 0)). \
            add(extra_volume_credits_for_comedy(perf, play))

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" != play["type"]:
            return Credits(0)

        return Credits(math.floor(perf['audience'] / 5))

    for perf in invoice['performances']:
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)

        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({perf["audience"]} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

## Encapsular estructuras de datos nativas

Nos quedan varios objetos interesantes. En particular Play y Performance, que son centrales en este dominio. Veamos cómo los podemos tratar.

En principio, estos objetos están tratados como diccionarios (o hash, o array asociativo, según el lenguaje). La tentación es intentar crear desde cero un objeto que reproduzca esa estructura. Sin embargo, vamos a seguir un enfoque más simplista. Por el momento solo vamos a encapsular esos diccionarios y añadir métodos que nos permitan acceder recuperar los valores de sus claves.

Una vez hecho esto, que será el primer paso, podremos hacer evolucionar la estructura interna sin que el resto del código tenga que preocuparse de ello. La razón para hacerlo así es evitar mezclar distintos objetivos en una única acción de refactor.

Después de examinar el código pienso que voy a empezar por Performance. En el bucle de la función statement se recorren las distintas Performances que se van a facturar y se opera con sus datos. En principio, una performance tiene las siguientes propiedades:

- playID, que hace referencia a la obra representada
- audience, que representa la cantidad de pública asistente

Así que introduciré la clase `Performance`, que tendrá por el momento dos métodos públicos: `play_id` y `audience`. En esta ocasión no voy a hacer tests, ya que son métodos triviales y su comportamiento quedará cubierto por los tests que ya tenemos.

```python
class Performance:
    def __init__(self, perf):
        self._data = perf

    def audience(self):
        return self._data['audience']

    def play_id(self):
        return self._data['playID']
```

Sospecho lo que estás pensando, pero de momento lo único que quiero es estar seguro de que el cambio funcionará. Ten en cuenta que este ejemplo es muy sencillo. En situaciones en que las estructuras de datos sean más complejas, este paso previo sirve para explorar las responsabilidades del objeto sin preocuparnos de su estructura interna.

Ahora toca introducirlo. Será aquí:

```python
    for perf in invoice['performances']:
        performance = Performance(perf)
        play = plays[perf['playID']]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)

        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({perf["audience"]} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)
```

Es ahora cuando podemos empezar a usarlo. Primero, en este nivel de abstracción:

```python
    for perf in invoice['performances']:
        performance = Performance(perf)
        play = plays[performance.play_id()]
        this_amount = calculate_performance_amount(perf, play)
        performance_credits = calculate_performance_credits(perf, play)

        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')
```

Y ahora viene algo interesante. Tenemos un par de funciones a las que les pasamos las variables `perf` (que representa una `Performance`) y `play` para hacer cálculos con ellas:

```python
    def calculate_performance_amount(perf, play):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)

        raise ValueError(f'unknown type: {play["type"]}')
```

De hecho, podemos ver que esta función llama a otras que utilizan `perf` como único parámetro. Esto nos está indicando que este comportamiento es propio de `Performance`. Es decir, sería responsabilidad de `Performance` calcular el importe facturable. Básicamente me estoy refiriendo a estas funciones:

```python
    def calculate_amount_for_comedy(perf):
        return Amount(30000) \
            .add(extra_amount_for_high_audience_in_comedy(perf)) \
            .add(Amount(300 * perf['audience']))

    def extra_amount_for_high_audience_in_comedy(perf):
        if perf['audience'] <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (perf['audience'] - 20))

    def calculate_amount_for_tragedy(perf):
        return Amount(40000) \
            .add(extra_amount_for_high_audience_in_tragedy(perf))

    def extra_amount_for_high_audience_in_tragedy(perf):
        if perf['audience'] <= 30:
            return Amount(0)

        return Amount(1000 * (perf['audience'] - 30))
```

¿Cómo voy a hacer este cambio? La verdad es que se me ocurren un par de maneras, aunque muy similares. El objetivo es copiar y adaptar el código que ahora está en funciones internas en `statement` para que sean métodos en `Performance`. La dificultad está en cómo hacer esto sin romper el test que tenemos.

Vamos con la primera forma. El primer paso es copiar el código de las funciones en `Performance` y adaptarlo de manera que no haya errores. Debería quedar más o menos así:

```python
from domain.amount import Amount


class Performance:
    def __init__(self, perf):
        self.data = perf

    def audience(self):
        return self.data['audience']

    def play_id(self):
        return self.data['playID']

    def calculate_amount_for_comedy(self):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy()) \
            .add(Amount(300 * self.audience()))

    def extra_amount_for_high_audience_in_comedy(self):
        if self.audience() <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (self.audience() - 20))

    def calculate_amount_for_tragedy(self):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy())

    def extra_amount_for_high_audience_in_tragedy(self):
        if self.audience() <= 30:
            return Amount(0)

        return Amount(1000 * (self.audience() - 30))
```

Ahora tenemos que pasar `performance` en vez de `perf` a la función en la línea:

```python
        this_amount = calculate_performance_amount(perf, play)
```

En estos casos, lo que suelo hacer es añadir un nuevo parámetro y luego reemplazar su uso, hasta que el viejo parámetro queda sin usar. Cuando verifico que todo funciona correctamente, elimino el viejo.

```python
    def calculate_performance_amount(perf, play, performance):
        if play['type'] == "tragedy":
            return calculate_amount_for_tragedy(perf)
        if play['type'] == "comedy":
            return calculate_amount_for_comedy(perf)

        raise ValueError(f'unknown type: {play["type"]}')
```

En este punto puedo hacer commit antes de realizar el cambio importante, que sería hacer que `performance` ejecute el cálculo:

```python
    def calculate_performance_amount(perf, play, performance):
    if play['type'] == "tragedy":
        return performance.calculate_amount_for_tragedy()
    if play['type'] == "comedy":
        return performance.calculate_amount_for_comedy()

    raise ValueError(f'unknown type: {play["type"]}')
```

He hecho el cambio y los tests siguen pasando, así que puedo eliminar el parámetro `perf` y también las funciones internas que ya no necesito.

```python
import math

from domain.amount import Amount
from domain.credits import Credits
from domain.performance import Performance
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    printer.print(f'Statement for {invoice["customer"]}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(play, performance):
        if play['type'] == "tragedy":
            return performance.calculate_amount_for_tragedy()
        if play['type'] == "comedy":
            return performance.calculate_amount_for_comedy()

        raise ValueError(f'unknown type: {play["type"]}')

    def calculate_performance_credits(perf, play):
        return Credits(max(perf['audience'] - 30, 0)). \
            add(extra_volume_credits_for_comedy(perf, play))

    def extra_volume_credits_for_comedy(perf, play):
        if "comedy" != play["type"]:
            return Credits(0)

        return Credits(math.floor(perf['audience'] / 5))

    for perf in invoice['performances']:
        performance = Performance(perf)
        play = plays[performance.play_id()]
        this_amount = calculate_performance_amount(play, performance)
        performance_credits = calculate_performance_credits(perf, play)

        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

¿Empieza a tener mejor pinta? Parece que sí. Hacemos lo mismo con `calculate_performance_credits`. No voy a poner todo el detalle del proceso, pero es la misma idea: mover el código a `Performance`, adaptándolo y cambiando los usos de las funciones internas por llamadas al objeto. Finalmente, eliminar el código que no necesitamos.

Así es como queda `Performance`:

```python
import math

from domain.amount import Amount
from domain.credits import Credits


class Performance:
    def __init__(self, perf):
        self.data = perf

    def audience(self):
        return self.data['audience']

    def play_id(self):
        return self.data['playID']

    def calculate_amount_for_comedy(self):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy()) \
            .add(Amount(300 * self.audience()))

    def extra_amount_for_high_audience_in_comedy(self):
        if self.audience() <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (self.audience() - 20))

    def calculate_amount_for_tragedy(self):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy())

    def extra_amount_for_high_audience_in_tragedy(self):
        if self.audience() <= 30:
            return Amount(0)

        return Amount(1000 * (self.audience() - 30))

    def calculate_performance_credits(self, play):
        return Credits(max(self.audience() - 30, 0)). \
            add(self.extra_volume_credits_for_comedy(play))

    def extra_volume_credits_for_comedy(self, play):
        if "comedy" != play["type"]:
            return Credits(0)

        return Credits(math.floor(self.audience() / 5))
```

Y ahora reemplazamos las llamadas a las funciones por `Performance`:

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.performance import Performance
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    printer.print(f'Statement for {invoice["customer"]}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def calculate_performance_amount(play, performance):
        if play['type'] == "tragedy":
            return performance.calculate_amount_for_tragedy()
        if play['type'] == "comedy":
            return performance.calculate_amount_for_comedy()

        raise ValueError(f'unknown type: {play["type"]}')

    for perf in invoice['performances']:
        performance = Performance(perf)
        play = plays[performance.play_id()]
        this_amount = calculate_performance_amount(play, performance)
        performance_credits = performance.calculate_performance_credits(play)

        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

El panorama se ha ido despejando al introducir objetos que han _atraído_ comportamiento y eso que aún nos queda por traer a colación el objeto `Play`.

Pero ahora nos fijamos en estas líneas. Hay una falta de _simetría_ que ralla un poco:

```python
        this_amount = calculate_performance_amount(play, performance)
        performance_credits = performance.calculate_performance_credits(play)
```

Está claro que `calculate_performance_amount` es un comportamiento de `Performance`, es hora de llevarlo a su lugar. Hacemos exactamente lo mismo. Copiar y adaptar. Luego reemplazar.

```python
import math

from domain.amount import Amount
from domain.credits import Credits


class Performance:
    def __init__(self, perf):
        self.data = perf

    def audience(self):
        return self.data['audience']

    def play_id(self):
        return self.data['playID']

    def calculate_amount_for_comedy(self):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy()) \
            .add(Amount(300 * self.audience()))

    def extra_amount_for_high_audience_in_comedy(self):
        if self.audience() <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (self.audience() - 20))

    def calculate_amount_for_tragedy(self):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy())

    def extra_amount_for_high_audience_in_tragedy(self):
        if self.audience() <= 30:
            return Amount(0)

        return Amount(1000 * (self.audience() - 30))

    def calculate_performance_credits(self, play):
        return Credits(max(self.audience() - 30, 0)). \
            add(self.extra_volume_credits_for_comedy(play))

    def extra_volume_credits_for_comedy(self, play):
        if "comedy" != play["type"]:
            return Credits(0)

        return Credits(math.floor(self.audience() / 5))

    def calculate_performance_amount(self, play):
        if play['type'] == "tragedy":
            return self.calculate_amount_for_tragedy()
        if play['type'] == "comedy":
            return self.calculate_amount_for_comedy()

        raise ValueError(f'unknown type: {play["type"]}')
```

Un detalle que quiero destacar de Performance es el uso de la auto-encapsulación. Esto consiste en no acceder directamente a las propiedades de una clase, sino a través de métodos que podrían ser privados. De este modo, el resto del código de la clase no tiene que saber nada acerca de su estructura y me da libertad para cambiarla en cualquier momento, como veremos más adelante.

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.performance import Performance
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    printer.print(f'Statement for {invoice["customer"]}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for perf in invoice['performances']:
        performance = Performance(perf)
        play = plays[performance.play_id()]
        this_amount = performance.calculate_performance_amount(play)
        performance_credits = performance.calculate_performance_credits(play)

        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

Mejoremos un poco el nombre de las cosas:

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.performance import Performance
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    printer.print(f'Statement for {invoice["customer"]}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for perf in invoice['performances']:
        performance = Performance(perf)
        play = plays[performance.play_id()]
        this_amount = performance.amount(play)
        performance_credits = performance.credits(play)

        line = f' {play["name"]}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

Nos queda introducir un objeto para representar una obra, que será `Play`. Por supuesto, hay una relación estrecha entre `Performance` y `Play` pero, de momento, no nos vamos a ocupar de eso. Simplemente queremos introducir el concepto y luego, ya veremos a dónde nos lleva.

Lo primero que hago es revisar qué cosas necesitamos de Play:

- `name`, para crear líneas de concepto en la factura.
- `type`, para saber qué tipo de obra es, ya que implica precios diferentes.

Esencialmente, hacemos lo mismo que con `Performance`. Empezamos simplemente encapsulando la estructura de datos de la manera más simple posible:

```python
class Play:
    def __init__(self, data):
        self._data = data

    def name(self):
        return self._data['name']

    def type(self):
        return self._data['type']
```

Como primer paso, reemplazamos la representación actual por el objeto. `Play` se usa sobre todo en `Performance`, pero hay un uso en `statement` que, de momento, necesitamos tener en cuenta:

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.performance import Performance
from domain.play import Play
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    printer.print(f'Statement for {invoice["customer"]}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for perf in invoice['performances']:
        performance = Performance(perf)
        play = Play(plays[performance.play_id()])
        this_amount = performance.amount(play)
        performance_credits = performance.credits(play)

        line = f' {play.name()}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

```python
import math

from domain.amount import Amount
from domain.credits import Credits


class Performance:
    def __init__(self, perf):
        self.data = perf

    def audience(self):
        return self.data['audience']

    def play_id(self):
        return self.data['playID']

    def calculate_amount_for_comedy(self):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy()) \
            .add(Amount(300 * self.audience()))

    def extra_amount_for_high_audience_in_comedy(self):
        if self.audience() <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (self.audience() - 20))

    def calculate_amount_for_tragedy(self):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy())

    def extra_amount_for_high_audience_in_tragedy(self):
        if self.audience() <= 30:
            return Amount(0)

        return Amount(1000 * (self.audience() - 30))

    def credits(self, play):
        return Credits(max(self.audience() - 30, 0)). \
            add(self.extra_volume_credits_for_comedy(play))

    def extra_volume_credits_for_comedy(self, play):
        if "comedy" != play.type():
            return Credits(0)

        return Credits(math.floor(self.audience() / 5))

    def amount(self, play):
        if play.type() == "tragedy":
            return self.calculate_amount_for_tragedy()
        if play.type() == "comedy":
            return self.calculate_amount_for_comedy()

        raise ValueError(f'unknown type: {play.type()}')
```

Esta es únicamente un primer paso. Dentro de un momento, veremos algunas ideas para proseguir con el refactor basándonos en las oportunidades que nos proporciona haber introducido objetos.

## Por qué funciona

La regla de encapsular todos los primitivos en objetos funciona porque, de entrada, nos ayuda a separar responsabilidades entre los diversos conceptos que participan en el programa. Además, contribuye a ocultar algunos detalles de implementación, haciendo más fácil entender qué está pasando.

Los objetos nos permiten encapsular reglas de negocio y aislar los detalles de implementación entre las distintas partes del código. Esto ayuda, además, en que esas mismas partes puedan evolucionar de forma independiente, sin afectar al funcionamiento del conjunto del programa. Ninguna parte del programa necesita saber, por ejemplo, los detalles estructurales de Performance o Play. Simplemente, les pasan mensajes para que proporcionen la información solicitada. La forma en que se calcula no es importante para el objeto que envía el mensaje, pero igualmente la obtiene.

A medida que hemos ido introduciendo objetos, hemos podido reducir el tamaño de la función `statement` y que su código sea mucho más expresivo. Por supuesto, es mejorable, pero ahora no están mezclados la mayor parte de detalles. En conjunto, hay mucha más cantidad de código, pero es mucho más legible y fácil de mantener.

Esto ocurre porque los objetos funcionan como _atractores_ de comportamiento. Una vez que descubrimos un objeto que participa en el programa, resulta fácil asignarle responsabilidades y extraerlas del código inicial. Por otro lado, los objetos nos ayudan a garantizar que los datos que encapsulan cumplen las reglas de dominio requeridas. No necesitamos verificarlo constantemente.

## Más allá

### Agregación de objetos

Al introducir objetos se va clarificando el escenario del programa y las relaciones entre los distintos conceptos. En nuestro ejercicio, por ejemplo, se aprecia muy bien que `Play` es un elemento de `Performance` y, salvo por conocer el nombre de la obra para poder imprimir la factura, la función `statement` no necesita saber ni que existe.

Así que podemos transformar `Performance` para usar `Play`. Sin embargo, antes nos vendría bien cambiar el modo en que Performance guarda su información. Es ahora cuando se pueden apreciar los beneficios de la auto-encapsulación. Sólo tengo que cambiar unas pocas líneas:

```python
class Performance:
    def __init__(self, perf):
        self._audience = perf['audience']
        self._play_id = perf['playID']

    def audience(self):
        return self._audience

    def play_id(self):
        return self._play_id

    # ...
```

De esta forma, es más fácil añadir una nueva propiedad:

```python
class Performance:
    def __init__(self, perf, plays):
        self._audience = perf['audience']
        self._play_id = perf['playID']
        self._play = Play(plays[perf['playID']])
        
    def audience(self):
        return self._audience

    def play_id(self):
        return self._play_id

    def play(self):
        return self._play
```

Y dar soporte al cambio en la instanciación, así como en el único uso directo que hace `statement` de `Play`.

```python
    for perf in invoice['performances']:
        performance = Performance(perf, plays)
        play = Play(plays[performance.play_id()])
        this_amount = performance.amount(play)
        performance_credits = performance.credits(play)

        line = f' {performance.play().name()}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)
```

Nos queda eliminar el paso de `Play` a los métodos `amount` y `credits`. Pero será bastante fácil:

```python
import math

from domain.amount import Amount
from domain.credits import Credits
from domain.play import Play


class Performance:
    def __init__(self, perf, plays):
        self._audience = perf['audience']
        self._play_id = perf['playID']
        self._play = Play(plays[perf['playID']])

    def audience(self):
        return self._audience

    def play_id(self):
        return self._play_id

    def play(self):
        return self._play

    def calculate_amount_for_comedy(self):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy()) \
            .add(Amount(300 * self.audience()))

    def extra_amount_for_high_audience_in_comedy(self):
        if self.audience() <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (self.audience() - 20))

    def calculate_amount_for_tragedy(self):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy())

    def extra_amount_for_high_audience_in_tragedy(self):
        if self.audience() <= 30:
            return Amount(0)

        return Amount(1000 * (self.audience() - 30))

    def credits(self, play):
        return Credits(max(self.audience() - 30, 0)). \
            add(self.extra_volume_credits_for_comedy(self.play()))

    def extra_volume_credits_for_comedy(self, play):
        if "comedy" != self.play().type():
            return Credits(0)

        return Credits(math.floor(self.audience() / 5))

    def amount(self):
        if self.play().type() == "tragedy":
            return self.calculate_amount_for_tragedy()
        if self.play().type() == "comedy":
            return self.calculate_amount_for_comedy()

        raise ValueError(f'unknown type: {self.play().type()}')
```

Y tras eso, eliminar el parámetro innecesario:

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.performance import Performance
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    printer.print(f'Statement for {invoice["customer"]}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for perf in invoice['performances']:
        performance = Performance(perf, plays)
        this_amount = performance.amount()
        performance_credits = performance.credits()

        line = f' {performance.play().name()}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

Sí, lo sé. Se pueden ver algunas cosillas cuestionables todavía. Vamos a seguir permitiendo que sean las reglas de Calisthenics las que nos guíen en el proceso y veremos si se arreglan o no.

## El resultado

El código ha evolucionado muchísimo tras aplicar la regla de encapsular primitivos en objetos. Sin embargo, todavía nos quedan algunos por atacar. Particularmente `invoice` y `plays`, pero los dejaremos para la próxima regla que nos pide hacer Colecciones de primera clase, lo que significa encapsular cada colección es su propia clase.

Si observamos el código desde el punto de vista de refactoring está claro que aún nos queda mucho trabajo por hacer y algunos _smells_ son evidentes y no están siendo tratados. Esto tienen un motivo en el contexto de estos artículos y no es otro que queremos ver si aplicar las reglas de forma sistemática nos conduce eventualmente a un mejor diseño. Hasta ahora creo que puede decirse que sí, con algunas salvedades, pero también es cierto que estamos aplicando cada regla una por una. En otras circunstancias estaríamos usando las reglas allí donde se viesen aplicables sin importar el orden.

En cualquier caso, en este momento podemos observar algunos efectos positivos, ya que las responsabilidades se han ido distribuyendo en objetos y funciones.

---

# [Colecciones de primera clase](https://franiglesias.github.io/calisthenics-4/)

La traducción literal a español no refleja muy bien lo que implica esta regla, pero es bastante sencilla. Se trata de encapsular en un objeto toda estructura de datos que represente una colección de tal manera que la única propiedad de este objeto sea esa misma estructura, con los métodos que necesitemos para tener acceso a los datos. Y en el fondo no es más que una extensión de la regla anterior. Unificadas ambas, podríamos decir que cualquier estructura de datos nativa del lenguaje debería ser encapsulada, da igual lo simple (primitivos) o compleja que sea (colecciones).

El motivo es aislarte de la estructura de datos de tal forma que el resto del programa no esté acoplado a la misma. Esto nos permite cambiar la estructura sin tener que tocar el resto del código cuando tengamos alguna razón para ello.

En el ejemplo que estamos usando en esta serie tenemos un par de buenos casos: `plays` y `performances`, dentro de `invoice`.

## Colección con acceso por clave

Este es el caso de `plays`. Accedemos a un elemento de esta colección dada una clave, que en este caso es el ID de la obra. La responsabilidad de `plays` en este sistema es actuar como una especie de catálogo en el que consultar las obras que la compañía puede representar. Simplemente, necesitamos un método `get_by_id`, que nos devuelva la obra solicitada.

```python
class Plays:
    def __init__(self, data):
        self._data = data

    def get_by_id(self, play_id):
        return Play(self._data[play_id])
```

Únicamente tenemos un uso y es fácil reemplazarlo:

```python
# ...
    for perf in invoice['performances']:
        performance = Performance(perf, Plays(plays))
        this_amount = performance.amount()
        performance_credits = performance.credits()

# ...
```

```python
import math

from domain.amount import Amount
from domain.credits import Credits


class Performance:
    def __init__(self, perf, plays):
        self._audience = perf['audience']
        self._play_id = perf['playID']
        self._play = plays.get_by_id(self._play_id)

# ...
```

Fíjate que no se trata de refactorizar la estructura en sí y cambiarla por otra que pueda ser más eficiente o apropiada. Se trata simplemente de no usar directamente ninguna estructura nativa, como si fuese una dependencia de terceros a la que no queremos acoplarnos.

Recuerda también aplicar YAGNI (no lo vas a necesitar), e introduce solo los métodos que tu código necesite para funcionar.

## Colección iterable

La única diferencia significativa entre el caso anterior y este, en el que vamos a encapsular la colección de performances, es que queremos poder iterar los elementos de esta colección, ya sea mediante un bucle `for` como el que tenemos en el ejemplo, ya sea mediante otro enfoque.

En python podemos hacer iterable una clase definiendo el método `__iter__` para que devuelva una clase iteradora, la cual debe contener el método `__next__`:

```python
class Performances:
    def __init__(self, data, plays):
        self._data = data
        self._plays = plays

    def __iter__(self):
        return PerformancesIterator(self)

    def by_index(self, index):
        return Performance(self._data[index], self._plays)

    def size(self):
        return len(self._data)


class PerformancesIterator:
    def __init__(self, performances):
        self._performances = performances
        self._current = 0

    def __next__(self):
        if self._current >= self._performances.size():
            raise StopIteration

        result = self._performances.by_index(self._current)
        self._current += 1
        return result
```

En el cuerpo de `statement` hacemos de esta manera:

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.performance import Performances
from domain.play import Plays
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    performances = Performances(invoice['performances'], Plays(plays))

    printer.print(f'Statement for {invoice["customer"]}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for performance in performances:
        this_amount = performance.amount()
        performance_credits = performance.credits()

        line = f' {performance.play().name()}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

Este cambio ha sido un poco más elaborado y ha conllevado algunas modificaciones interesantes. Por ejemplo, la instanciación de `Performance` ocurre dentro de `Performances`, así que `statement` ya no necesita conocer cómo se construye un objeto `Performance`.Luego profundizaré en algunas consecuencias de esto.

Como he mencionado antes, lo único que hemos hecho ha sido mover la estructura de datos original (un diccionario) dentro de una nueva clase. De esta forma, el código de statement, no conoce los detalles de implementación de `Performances` (o de `Plays`) pero sigue pudiendo acceder a la información que necesita. En el futuro podríamos cambiar esto sin necesidad de afectar a `statement`, lo que es una ventaja importante.

Por eso, aunque ahora mismo el código dentro de `Performances` nos parezca menos que bueno, podremos cambiarlo en cualquier momento sin miedo de romper cosas en múltiples lugares. Los cambios ocurrirán únicamente en un sitio (dentro de `Performance`), maximizando la mantenibilidad y manteniendo localizados los errores potenciales.

## Por qué funciona

Al igual que ocurre con la regla anterior, encapsular colecciones nos permite desacoplarnos de la estructura nativa de datos. Esto es una gran ventaja porque nos aporta libertad a la hora de cambiar esta estructura y la gestión de los datos en ella.

Además, este tipo de cambios suele generar algunas ventajas más. Las estructuras nativas están diseñadas para cubrir numerosos casos de uso, por lo que son genéricas y pueden incluir numerosos métodos que no vamos a necesitar o que introducen confusión a la hora de utilizarlos. Al encapsular en una clase, podemos definir cómo queremos que el resto del programa interactúe con ella de forma inequívoca, usando incluso un lenguaje apropiado a nuestro dominio.

Por otro lado, estos procesos de encapsulación ayudan a descubrir y modelar mejor relaciones entre conceptos, sugiriendo dónde deben ir las distintas responsabilidades.

## Más allá

A medida que aplicamos las reglas de Object Calisthenics el código no solo va tomando mejor forma, sino que también desvela áreas que pueden mejorar.

Esto ocurre porque, en general, las reglas nos fuerzan a organizar mejor el código. No arreglan los problemas, pero contribuyen a despejar el paisaje de una forma parecida a lo que ocurre cuando, por ejemplo, organizamos las piezas de un puzzle por colores o texturas antes de empezar. Se podría decir, que gracias a esta manera de trabajar conseguimos dividir un problema grande en partes manejables.

Así, por ejemplo, tras el último cambio podemos ver que `invoice` es la última estructura de datos nativa que nos queda por arreglar. Pero también vemos que podríamos mejorar cosas en la forma en que instanciamos `Performance`.

### Encapsular estructuras de datos

Como he mencionado más arriba, tanto la regla de hoy “Encapsular colecciones” como la del pasado artículo “Encapsular primitivos” son dos caras de una misma moneda” encapsular cualquier estructura de datos nativa. Esto es, cualquier concepto que aparece en nuestro dominio debería ser representado por un objeto que se puede implementar usando la estructura de datos que más nos convenga, pero sin que el resto del código tenga que saber qué estructura en concreto estamos usando.

En este ejercicio he dejado `invoice` para el final para analizarlo con calma. En principio, un objeto `Invoice` nos debería proporcionar el nombre del cliente (para imprimir la factura) y la lista de actuaciones.

```python
class Invoice:
    def __init__(self, data):
        self._data = data

    def customer(self):
        return self._data['customer']

    def performances(self):
        return self._data['performances']
```

Y reemplazarlo sus usos en el código resultaría trivial:

```python
def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    inv = Invoice(invoice)

    performances = Performances(inv.performances(), Plays(plays))

    printer.print(f'Statement for {inv.customer()}\n')

    # ...
```

De entrada, es fácil ver que `Invoice` nos pide más responsabilidades. Por ejemplo, la instanciación de `Performances` debería ocurrir en `Invoice`. Podríamos hacerlo así:

```python
from domain.performance import Performances
from domain.play import Plays


class Invoice:
    def __init__(self, data, plays):
        self._data = data
        self._customer = data['customer']
        self._performances = Performances(data['performances'], Plays(plays))

    def customer(self):
        return self._customer

    def performances(self):
        return self._performances
```

Y usarlo de esta manera:

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.invoice import Invoice
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    inv = Invoice(invoice, plays)

    printer.print(f'Statement for {inv.customer()}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for performance in inv.performances():
        this_amount = performance.amount()
        performance_credits = performance.credits()

        line = f' {performance.play().name()}: {format_as_dollars(this_amount.current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(this_amount)
        volume_credits = volume_credits.add(performance_credits)

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

Ahora está claro que la lógica para calcular `invoice_amount` y `volume_credits` está reclamando fuertemente formar parte de `Invoice`, cosa que tiene su complicación dada la forma en que se _imprime_ la factura. Ya llegaremos a esto, pero ahora se ve claramente que hay dos responsabilidades diferentes: el cálculo de las líneas y totales de la factura y la impresión de las mismas. Nuestro problema es que ahora aparecen entrelazadas.

¿Hay algo que podamos hacer aquí? Una posibilidad es eliminar variables temporales, lo que reduce bastante el ruido, aclarando algunas cosas, pero _ensuciando_ otras.

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.invoice import Invoice
from domain.printer import Printer


def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    inv = Invoice(invoice, plays)

    printer.print(f'Statement for {inv.customer()}\n')

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    for performance in inv.performances():
        line = f' {performance.play().name()}: {format_as_dollars(performance.amount().current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(performance.amount())
        volume_credits = volume_credits.add(performance.credits())

    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')

    return printer.output()
```

### Evolución interna de los objetos

Hemos dicho que al encapsular estructuras de datos en objetos, la evolución interna de estos se hace de forma transparente para el resto del código. Esto nos permite hacer cambios sin romper funcionalidades, especialmente si estamos protegidas por tests.

Vamos a ver unos ejemplos.

Tras la transformación anterior, alguien podría argumentar que llamamos dos veces a `performance.amount()`, lo que podría tener consecuencias en, ejem, performance.

```python
    # ...
    for performance in inv.performances():
        line = f' {performance.play().name()}: {format_as_dollars(performance.amount().current() / 100)} ({performance.audience()} seats)\n'
        printer.print(line)

        invoice_amount = invoice_amount.add(performance.amount())
        volume_credits = volume_credits.add(performance.credits())

    # ...
```

Si esto te supone mucho problema, un patrón _memoization_ podría ayudar. Básicamente, se trata de mantener una _cache_ del cálculo, de la cual el código que llaman no tiene que saber ni que existe. Por ejemplo, esta implementación bastante ingenua:

```python
import math

from domain.amount import Amount
from domain.credits import Credits


class Performance:
    def __init__(self, perf, plays):
        self._audience = perf['audience']
        self._play_id = perf['playID']
        self._play = plays.get_by_id(self._play_id)
        self._amount = None

    # ...

    def amount(self):
        if self._amount is not None:
            return self._amount

        if self.play().type() == "tragedy":
            tragedy = self.calculate_amount_for_tragedy()
            self._amount = tragedy
            return tragedy
        if self.play().type() == "comedy":
            comedy = self.calculate_amount_for_comedy()
            self._amount = comedy
            return comedy

        raise ValueError(f'unknown type: {self.play().type()}')
```

Alternativamente, podrías [utilizar esta clase memoize de Graham Jenson](https://maori.geek.nz/python-decorator-to-memoize-instance-methods-ad4f6a05f1dc), con lo que te bastaría decorar el método `amount` con un `@memoize`.

Como puedes ver, al tener objetos con responsabilidades bien definidas y un contrato claro con sus usuarios, introducir mejoras es muchísimo más fácil y seguro.

Otro asunto interesante es que cuando instanciamos `Performance`, seguimos pasando la colección completas de obras. Pero no tenemos por qué hacerlo así, ya que ahora es más fácil montar `Performance` con la obra (`Play`) que le corresponde. Este es el código que tenemos ahora:

```python
class Performances:
    def __init__(self, data, plays):
        self._data = data
        self._plays = plays

    def __iter__(self):
        return PerformancesIterator(self)

    def by_index(self, index):
        return Performance(self._data[index], self._plays)

    def size(self):
        return len(self._data)
```

Y este el cambio que proponemos:

```python
class Performances:
    def __init__(self, data, plays):
        self._data = data
        self._plays = plays

    def __iter__(self):
        return PerformancesIterator(self)

    def by_index(self, index):
        return Performance(self._data[index], self._plays.get_by_id(self._data[index]['playID']))

    def size(self):
        return len(self._data)
```

Mientras que `Performance` podría quedar así:

```python
class Performance:
    def __init__(self, perf, play):
        self._audience = perf['audience']
        self._play = play
        self._amount = None

    # ...
```

Pero entonces resulta que podemos tener un constructor mucho más natural:

```python
class Performance:
    def __init__(self, audience, play):
        self._audience = audience
        self._play = play
        self._amount = None
```

Y usarlo de esta otra forma:

```python
    def by_index(self, index):
        play = self._data[index]
        return Performance(play['audience'], self._plays.get_by_id(play['playID']))
```

## Resultado

Object Calisthenics nos está ayudando a despejar el diseño del código, identificando objetos y repartiendo responsabilidades. Gracias a ello tenemos un código que, aunque es más grande, está organizado en objetos cada vez más especializados en sus tareas, de modo que la comprensión del sistema es mejor, a la vez que se hace más mantenible y, como acabamos de ver, incluso más optimizable.

---

# [Un punto por línea](https://franiglesias.github.io/calisthenics-5/)

Esta regla nos pide no encadenar llamadas a objetos proporcionados por otros objetos, de tal forma que solo tengamos un punto (una flecha en PHP) en cada línea de código. Puede parecer fácil de aplicar, pero vamos a poder identificar varias situaciones en las que la regla no es relevante, así como diferentes soluciones cuando sí lo es.

## Las interfaces fluidas son correctas

Las interfaces fluidas no se ven afectadas por esta regla. Las interfaces fluidas devuelven el mismo objeto al que se pasa el mensaje, por lo que podemos seguir enviándole mensajes sin límite, lo que parece una oportunidad de aplicar la regla. Pero no lo es. En todo caso, es cierto que poner un punto por línea mejora mucha la legibilidad. El objetivo, y ventaja, de la interfaz fluida es poder enviar varios mensajes a un mismo objeto en un orden dado y que, además, se pueda entender como una operación unitaria.

En nuestro código, hacemos algo así con Amount, aunque cada vez se devuelva una instancia distinta es semánticamente el mismo objeto:

```python
    def calculate_amount_for_comedy(self):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy()) \
            .add(Amount(300 * self.audience()))
```

## Filtración de propiedades

Fijémonos ahora en esta línea:

```python
line = f' {performance.play().name()}: {format_as_dollars(performance.amount().current() / 100)} ({performance.audience()} seats)\n'
```

Para saber el título de la actuación, tenemos que pedirle a `Performance` la obra y obtener su título. De este modo, se revela un detalle de implementación de `Performance` que el resto del código no tiene por qué conocer. El comportamiento que se quiere de `Performance` es que sea capaz de decirnos el título de la obra que se representa, da igual si lo tiene guardado o le pregunta a `Play` o tiene alguna otra forma de obtenerlo o construirlo.

Por eso, una forma más adecuada sería:

```python
line = f' {performance.title()}: {format_as_dollars(performance.amount().current() / 100)} ({performance.audience()} seats)\n'
```

De tal forma que ahora el mundo exterior no tiene ningún detalle sobre cómo hace Performance para proporcionar el título de la obra representada:

```python
class Performance:
    def __init__(self, audience, play):
        self._audience = audience
        self._play = play
        self._amount = None

    def audience(self):
        return self._audience

    def play(self):
        return self._play

    def title(self):
        return self._play.name()
```

De este modo, el resto del código reduce su acoplamiento de `Performance` y esta puede modificar la forma en que obtiene el título sin afectar a sus consumidores.

Con todo, en este caso concreto podría haber otras soluciones, pero no voy a tratarlas en este momento, ya que me estoy limitando a aplicar la reglas de Calisthenics. Pero, en cualquier caso, creo que se ve muy bien cómo aplicar una regla va desvelando mejores soluciones, pero también problemas de diseño más profundos que requieren soluciones más elaboradas. Es decir: intentar aplicar la regla nos lleva a pensar más a fondo en ciertas decisiones de diseño.

En este caso, la solución es aceptable porque tiene sentido que `Performance` tenga como una de sus responsabilidades saber el nombre de la obra representada.

## Más filtración de conocimiento

Veamos este fragmento:

`performance.amount().current()`

Aquí tenemos un problema aparentemente similar. El método `Performance.amount()` nos devuelve un objeto y `statement` invoca un método en ese objeto devuelto. ¿Podemos aplicar la misma solución que antes?

Aparentemente sí, añadiendo a `Performance` un método que nos proporcione ese valor, algo así como:

```python
class Performance:
    def __init__(self, audience, play):
        self._audience = audience
        self._play = play
        self._amount = None

# ...

    def amount_value(self):
        return self.amount().current()

# ...
```

Si lo pensamos un poco a fondo, veremos que no es nada correcto. Y eso es porque, de hecho, el método `Amount.current()` no debería existir, ya que en realidad expone una propiedad del objeto `Amount`. El método existe porque necesitamos obtener el primitivo contenido en el objeto. En otras palabras: intentar aplicar esta regla va más allá de simplemente encapsular el código en un nuevo método. Debería hacernos reflexionar sobre el diseño.

Una mejor solución es delegar y pasar el objeto a alguien que sepa comunicarse con él, Con todo, todavía presenta problemas, pero los tendremos que examinar en otro momento:

```python
    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    def format_line(title, audience, amount):
        return f' {title}: {format_as_dollars(amount.current() / 100)} ({audience} seats)\n'

    for performance in inv.performances():
        line = format_line(performance.title(), performance.audience(), performance.amount())
        printer.print(line)
        invoice_amount = invoice_amount.add(performance.amount())
        volume_credits = volume_credits.add(performance.credits())
```

## Un caso muy sutil

¿Notas algo problemático aquí?

```python
    for performance in inv.performances():
        printer.print(formatted_line(performance.title(), performance.audience(), performance.amount()))
        invoice_amount = invoice_amount.add(performance.amount())
        volume_credits = volume_credits.add(performance.credits())
```

Pues es un caso muy sutil de violación de esta regla. `statement` recibe objetos `Performance` que no tendría que conocer. Es una situación similar a la que acabamos de describir en el apartado anterior.

Podríamos abordarla así, pero los problemas son evidentes.

```python
def process_performance(performance, invoice_amount, volume_credits, printer):
    printer.print(formatted_line(performance.title(), performance.audience(), performance.amount()))
    invoice_amount = invoice_amount.add(performance.amount())
    volume_credits = volume_credits.add(performance.credits())
    return invoice_amount, volume_credits

for performance in invoice.performances():
    invoice_amount, volume_credits = process_performance(performance, invoice_amount, volume_credits, printer)
```

Tenemos que pasar variables que serán retornadas, aparte del objeto `Performance`. Y para completarlo, el nuevo método devuelve dos valores.

Hay varias razones por las que está pasando esto. Por un lado, el hecho de `Invoice` sea, por el momento, un objeto muy anémico, ya que debería ser responsable de calcular tanto el importe total como los créditos. Por otra parte, en el bucle están pasando varias cosas: se calculan los importes parciales, se van acumulando los dos totales y además se envían las líneas para imprimir.

Nos conviene separar las responsabilidades. Primer paso:

```python
    for performance in invoice.performances():
        invoice_amount = invoice_amount.add(performance.amount())

    for performance in invoice.performances():
        volume_credits = volume_credits.add(performance.credits())

    for performance in invoice.performances():
        printer.print(formatted_line(performance.title(), performance.audience(), performance.amount()))
```

Segundo paso. Pongamos juntas las cosas relacionadas:

```python
    invoice_amount = Amount(0)
    for performance in invoice.performances():
        invoice_amount = invoice_amount.add(performance.amount())

    volume_credits = Credits(0)
    for performance in invoice.performances():
        volume_credits = volume_credits.add(performance.credits())

    printer.print(f'Statement for {invoice.customer()}\n')
    for performance in invoice.performances():
        printer.print(formatted_line(performance.title(), performance.audience(), performance.amount()))
    printer.print(f'Amount owed is {format_as_dollars(invoice_amount.current() // 100)}\n')
    printer.print(f'You earned {volume_credits.current()} credits\n')
```

Se debería ver claro que esta lógica pertenece a `Invoice` y la podríamos pasar sin mucha dificultad.

```python
from domain.amount import Amount
from domain.credits import Credits
from domain.performance import Performances
from domain.play import Plays


class Invoice:
    def __init__(self, data, plays):
        self._data = data
        self._customer = data['customer']
        self._performances = Performances(data['performances'], Plays(plays))

    def customer(self):
        return self._customer

    def performances(self):
        return self._performances

    def amount(self):
        invoice_amount = Amount(0)
        for performance in self.performances():
            invoice_amount = invoice_amount.add(performance.amount())

        return invoice_amount

    def credits(self):
        volume_credits = Credits(0)
        for performance in self.performances():
            volume_credits = volume_credits.add(performance.credits())

        return volume_credits
```

Y así quedaría `statement`, después de limpiar un poco el código.

```python
from domain.invoice import Invoice
from domain.printer import Printer


def statement(invoice_data, plays):
    def formatted_line(title, audience, amount):
        return f' {title}: {format_as_dollars(amount.current() / 100)} ({audience} seats)\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    printer = Printer()

    invoice = Invoice(invoice_data, plays)

    printer.print(f'Statement for {invoice.customer()}\n')
    for performance in invoice.performances():
        printer.print(formatted_line(performance.title(), performance.audience(), performance.amount()))
    printer.print(f'Amount owed is {format_as_dollars(invoice.amount().current() // 100)}\n')
    printer.print(f'You earned {invoice.credits().current()} credits\n')

    return printer.output()
```

## Por qué funciona

Esta es una regla que nos remite al Principio de Mínimo Conocimiento o Ley de Demeter y su objetivo es evitar acoplarnos a detalles internos de otros objetos. Nos fuerza a considerar los objetos como cajas negras con las que nos podemos comunicar, pero no saber cómo funcionan por dentro.

Cuando un objeto usa otro lo hace a través de su interfaz pública. La interfaz pública define los mensajes que un objeto puede recibir y las respuestas que puede devolver. Este es el máximo de conocimiento que un objeto debería tener sobre otro para minimizar el acoplamiento. Todo conocimiento a mayores incrementa el acoplamiento. Ese conocimiento incluye saber cómo comunicarse con objetos que son devueltos. La acción del consumidor debería limitarse a pasar ese objeto para que sea empleado en otro sitio.

En general, que haya puntos del código en que aplicar esta regla nos revela errores de diseño. Le estamos pidiendo a objetos comportamientos que no les corresponden, usando un conocimiento íntimo de su estructura.

## El resultado

Por un lado, esta regla nos ayuda a mover responsabilidades a su lugar adecuado. Pero también suele destapar problemas que requieren reconsiderar nuestro diseño. No basta con introducir un método para ocultar una llamada encadenada.

Por eso, el resultado en este momento resulta un poco insatisfactorio. Tendremos que esperar a las reglas restantes para alcanzar mejores soluciones.

```python
from domain.invoice import Invoice
from domain.printer import Printer


def statement(invoice_data, plays):
    def formatted_line(title, audience, amount):
        return f' {title}: {format_as_dollars(amount.current() / 100)} ({audience} seats)\n'

    def format_as_dollars(amount):
        return f"${amount:0,.2f}"

    printer = Printer()

    invoice = Invoice(invoice_data, plays)

    printer.print(f'Statement for {invoice.customer()}\n')
    for performance in invoice.performances():
        printer.print(formatted_line(performance.title(), performance.audience(), performance.amount()))
    printer.print(f'Amount owed is {format_as_dollars(invoice.amount().current() // 100)}\n')
    printer.print(f'You earned {invoice.credits().current()} credits\n')

    return printer.output()
```

---

# [No usar abreviaturas](https://franiglesias.github.io/calisthenics-6/)

Este artículo debería ser bastante breve, ya que no hay muchos casos en nuestro ejercicio de ejemplo. No obstante, forzaremos algunos ejemplos para ver los problemas del uso de abreviaturas.

La regla nos pide no usar abreviaturas para nombrar variables, objetos o funciones. El objetivo, por supuesto, es que el código sea lo más autoexplicativo posible. Si nos encontramos una abreviatura puede ocurrir que no conozcamos la referencia, puede que sea ambigua y no pueda entender bien el significado ni siquiera por el contexto, o puede ser simplemente confusa.

## Abreviatura por conflicto de nombres

Aquí tenemos un ejemplo de uso de una abreviatura. En este caso para no generar un conflicto de nombres entre el parámetro que pasa los datos y la variable que contiene el objeto `Invoice`. Este tipo de atajos vienen de no tomar suficiente tiempo para pensar un nombre adecuado.

```python
def statement(invoice, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    inv = Invoice(invoice, plays)

    printer.print(f'Statement for {inv.customer()}\n')
```

La pregunta es, ¿quién de los dos tiene el derecho a llamarse propiamente `invoice`? A medida que hemos ido aplicando reglas e introduciendo objetos, también necesitamos cambiar nombres. Al principio, `invoice` designaba una estructura de datos que representaba una factura. Sin embargo, al introducir el objeto `Invoice`, el parámetro pasa a ser un simple transporte de datos.

Por esa razón, realmente tiene más sentido hacer algunos cambios en los nombres. Esta es una posible solución:

```python
def statement(invoice_data, plays):
    printer = Printer()
    invoice_amount = Amount(0)
    volume_credits = Credits(0)
    invoice = Invoice(invoice_data, plays)

    printer.print(f'Statement for {invoice.customer()}\n')
```

## Lo que nos dice un nombre

La abreviatura `inv` puede ser confusa si no tenemos contexto. Por ejemplo, es habitual que signifique `inverso`, así que siempre es preferible poner nombres completos, significativos e inequívocos. Es preferible pasarse por nombre largo que por nombre corto.

En `Performance` tenemos el método `extra_amount_for_high_audience_in_comedy`, cuyo nombre es extremadamente largo. Sin embargo, es inequívoco y dice exactamente lo que hace. A veces, el contexto nos puede proporcionar suficientes pistas. Este método es llamado desde `calculate_amount_for_comedy`, por lo que podríamos considerar acortarlo a `extra_for_high_audience`. Pero existe otro método de nombre similar en la misma clase: `extra_amount_for_high_audience_in_tragedy`. Así que para diferenciarlos deberíamos mantener la referencia al tipo de obra.

Por supuesto, en realidad estos nombres nos están insistiendo en la necesidad de abordar el polimorfismo de `Play`, pero es algo que vamos a dejar para otro artículo más adelante. La lección aquí es que reflexionar sobre los nombres nos ayudará a alcanzar un mejor diseño.

En cualquier caso, si un nombre resulta incómodo por ser demasiado largo, siempre tienes la oportunidad de refactorizar.

## Abreviaturas aceptables

Algunas abreviaturas son de uso común. Por ejemplo, `vat` por `value added tax`.

## Convenciones problemáticas

Existen algunas convenciones que usan nombres abreviados o especialmente cortos. Un ejemplo son los bucles, en los que se suelen usar nombres de variables como `i`, `j` o `k`. En su lugar es recomendable usar alternativas: `index`, `position`, `counter`, son mucho más explícitas y más difíciles de confundir.

```python
    for i in range(0, 3):
        print(i)
```

Frente a:

```python
    for counter in range(0, 3):
        print(counter)
```

En general, usar variables de una sola letra es confuso. ¿Qué es `p`? Incluso teniendo el contexto, una variable de una única letra nos obliga a pensar dos veces.

```python
    for p in invoice.performances():
        printer.print(formatted_line(p.title(), p.audience(), p.amount()))
        invoice_amount = invoice_amount.add(p.amount())
        volume_credits = volume_credits.add(p.credits())
```

Además, es poco práctico. Si tienes que hacer una búsqueda de texto para encontrar la variable puede ser una odisea. Dentro del archivo nos salva que para tareas de refactor los IDE suelen usar el árbol sintáctico, pero si la búsqueda es de texto normal… ¡Buena suerte!

Esto ocurre también con nombres cortos, pero demasiado genéricos, como `get`, `add`, etc., que son comunes a infinidad de librerías.

## Por qué funciona

No usar abreviaturas nos fuerza a pensar nombres significativos, lo que ayuda a que el código se explique mejor por sí mismo. Esto permite que sea más fácil incorporar más personas a los proyectos y hacerlo más mantenible en el largo plazo. Puede que con el tiempo nos olvidemos de lo que significaban las abreviaturas, por lo que usar nombres completos será una ventaja.

---

# [Mantener todas las entidades pequeñas](https://franiglesias.github.io/calisthenics-7/)

Esta regla suele generar discusión porque vamos a poner un límite totalmente arbitrario al tamaño de las entidades de código. Esto se refiere a clases, al número de métodos, al cuerpo de funciones, al número de archivos en un paquete, etc. Por ejemplo, esta es una propuesta más o menos típica:

- 10 archivos por paquete o carpeta
- 50 líneas por clase
- 5 líneas por método o función
- 2 argumentos por método o función

Así que se trata de recorrer el código buscando áreas que superen estos límites.

El objetivo, como ocurre en todas las reglas de Calisthenics, es que tratar de forzar la aplicación de las reglas nos traiga como resultado un código mejor diseñado, más fácil de entender y de mantener. En el caso de esta, lo que buscamos obtener es un sistema de objetos pequeños muy simples.

Lo cierto es que después de todos los cambios resultado de aplicar las reglas anteriores, nos encontramos con relativamente pocos casos problemáticos. Pero alguno hay.

## Paquetes y sub-paquetes

Por ejemplo, el paquete domain, que contiene casi todo el código que hemos generado, no llega a 10 archivos. En parte es porque tenemos algunos archivos que contienen dos clases, algo que no está recomendado en todos los lenguajes. Puedes verlo como una forma de contribuir a esta regla, haciendo que el módulo de Python se pueda considerar como un sub-paquete y forzando que no contenga más de 10 clases o funciones.

En general, en el caso de encontrarnos con paquetes de más de 10 archivos, deberíamos plantearnos agruparlos por algún criterio en sub-paquetes cohesivos.

## Clases grandes

Tenemos una clase que tiene más de 50 líneas. `Performance` contiene gran parte de la lógica del programa pero, ¿podemos reducir su tamaño? O bien, ¿necesita realmente ser tan grande? Además, el método `amount` tiene unas 10 líneas, con lo cual también supera el límite de cinco que habíamos definido.

```python
class Performance:
    def __init__(self, audience, play):
        self._audience = audience
        self._play = play
        self._amount = None

    def audience(self):
        return self._audience

    def play(self):
        return self._play

    def title(self):
        return self._play.name()

    def calculate_amount_for_comedy(self):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy()) \
            .add(Amount(300 * self.audience()))

    def extra_amount_for_high_audience_in_comedy(self):
        if self.audience() <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (self.audience() - 20))

    def calculate_amount_for_tragedy(self):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy())

    def extra_amount_for_high_audience_in_tragedy(self):
        if self.audience() <= 30:
            return Amount(0)

        return Amount(1000 * (self.audience() - 30))

    def credits(self):
        return Credits(max(self.audience() - 30, 0)). \
            add(self.extra_volume_credits_for_comedy())

    def extra_volume_credits_for_comedy(self):
        if "comedy" != self.play().type():
            return Credits(0)

        return Credits(math.floor(self.audience() / 5))

    def amount(self):
        if self._amount is not None:
            return self._amount

        if self.play().type() == "tragedy":
            tragedy = self.calculate_amount_for_tragedy()
            self._amount = tragedy
            return tragedy
        if self.play().type() == "comedy":
            comedy = self.calculate_amount_for_comedy()
            self._amount = comedy
            return comedy

        raise ValueError(f'unknown type: {self.play().type()}')
```

Parte del problema de `Performance` es que se ocupa de varias cosas. Gran parte de su lógica depende del tipo de obra representada, así que tiene que preguntarle a `Play` por su tipo y hacer cálculos basados en eso. Esto nos remite a la última regla que nos pide no exponer _getters_, _setters_ o propiedades públicas de los objetos que, a su vez, se basa en la aplicación del principio “Tell, don’t ask”. En pocas palabras: si tienes que preguntar un objeto por una información, para actuar con base en esa información, entonces haz que el objeto se encargue de hacerlo.

De hecho, si la lógica estuviese en `Play` podríamos reducir el tamaño de la clase `Performance`. Vamos a empezar por ahí.

Fundamentalmente, podemos mover algunos métodos de `Performance` a `Play`, así que simplemente los copio y los adapto. Cuando los tenga listos, podré reemplazarlos. Voy con los relacionados con el tipo _Comedy_. Un detalle importante es que ahora tenemos que pasar el argumento de audiencia para permitir el cálculo.

```python
class Play:
    def __init__(self, data):
        self._data = data

    def name(self):
        return self._data['name']

    def type(self):
        return self._data['type']

    def calculate_amount_for_comedy(self, audience):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy(audience)) \
            .add(Amount(300 * audience))

    def extra_amount_for_high_audience_in_comedy(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))
```

Ahora puedo introducirlos en lugar de los existentes, que puedo eliminar a continuación, una vez que he comprobado que los tests siguen pasando igualmente.

```python
    def amount(self):
        if self._amount is not None:
            return self._amount

        if self.play().type() == "tragedy":
            tragedy = self.calculate_amount_for_tragedy()
            self._amount = tragedy
            return tragedy
        if self.play().type() == "comedy":
            comedy = self.play().calculate_amount_for_comedy(self.audience())
            self._amount = comedy
            return comedy

        raise ValueError(f'unknown type: {self.play().type()}')
```

Y pasará lo mismo con las obra de tipo _Tragedy_, moviendo los métodos relacionados y reemplazando las llamadas. Quedará así:

```python
class Play:
    def __init__(self, data):
        self._data = data

    def name(self):
        return self._data['name']

    def type(self):
        return self._data['type']

    def calculate_amount_for_comedy(self, audience):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy(audience)) \
            .add(Amount(300 * audience))

    def extra_amount_for_high_audience_in_comedy(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))

    def calculate_amount_for_tragedy(self, audience):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy(audience))

    def extra_amount_for_high_audience_in_tragedy(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(1000 * (audience - 30))
```

Y reduciremos el tamaño de `Performance` porque nos libramos de bastantes métodos.

```python
    def amount(self):
        if self._amount is not None:
            return self._amount

        if self.play().type() == "tragedy":
            tragedy = self.play().calculate_amount_for_tragedy(self.audience())
            self._amount = tragedy
            return tragedy
        if self.play().type() == "comedy":
            comedy = self.play().calculate_amount_for_comedy(self.audience())
            self._amount = comedy
            return comedy

        raise ValueError(f'unknown type: {self.play().type()}')
```

De hecho, todavía podemos quitar un poco más de código a `Performance` puesto que tenemos que hay un cálculo de créditos que depende de la obra:

```python
class Play:
    def __init__(self, data):
        self._data = data

    def name(self):
        return self._data['name']

    def type(self):
        return self._data['type']

    def calculate_amount_for_comedy(self, audience):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy(audience)) \
            .add(Amount(300 * audience))

    def extra_amount_for_high_audience_in_comedy(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))

    def calculate_amount_for_tragedy(self, audience):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy(audience))

    def extra_amount_for_high_audience_in_tragedy(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(1000 * (audience - 30))

    def credits(self, audience):
        if "comedy" != self.type():
            return Credits(0)

        return Credits(math.floor(audience / 5))
```

Con lo que `Performance` se reduce hasta la mitad de líneas:

```python
class Performance:
    def __init__(self, audience, play):
        self._audience = audience
        self._play = play
        self._amount = None

    def audience(self):
        return self._audience

    def play(self):
        return self._play

    def title(self):
        return self._play.name()

    def credits(self):
        return Credits(max(self.audience() - 30, 0)). \
            add(self.play().extra_volume_credits_for_comedy(self.audience()))

    def amount(self):
        if self._amount is not None:
            return self._amount

        if self.play().type() == "tragedy":
            tragedy = self.play().calculate_amount_for_tragedy(self.audience())
            self._amount = tragedy
            return tragedy
        if self.play().type() == "comedy":
            comedy = self.play().calculate_amount_for_comedy(self.audience())
            self._amount = comedy
            return comedy

        raise ValueError(f'unknown type: {self.play().type()}')
```

Por supuesto puedes argumentar: pero si has movido el código de una clase a otra. Ahora `Play` es mucho más grande. Y es cierto, pero ahora contiene casi toda la lógica que le pertenece.

## Un método largo

Con todo, el método `amount` sigue teniendo más de cinco líneas. Hemos adelgazado la clase, pero no el método más grande. Podemos mover parte de este código a `Play`. Aquí tenemos un pequeño obstáculo pues implementamos la memoización de una forma que nos complica un poco. Pero podemos arreglarlo. El primer paso es separar la memoización del cálculo:

```python
    def amount(self):
        if self._amount is not None:
            return self._amount

        self._amount = self.calculate_amount()
        
        return self._amount

    def calculate_amount(self):
        if self.play().type() == "tragedy":
            return self.play().calculate_amount_for_tragedy(self.audience())
        if self.play().type() == "comedy":
            return self.play().calculate_amount_for_comedy(self.audience())

        raise ValueError(f'unknown type: {self.play().type()}')
```

Gracias a este cambio, además resulta que reducimos el tamaño del método `amount`, y el nuevo método también cumple la limitación a un máximo de cinco líneas. De hecho, ahora `amount` se encarga básicamente de la memoización y `Play` del cálculo. Más interesante aún es que se ha reducido el acoplamiento. `Play` no sabe nada de `Performance`, pero lo mejor es que esta no sabe nada de `Play`. Es decir: únicamente sabe que le puede pedir `amount` y `credits`, pero no tiene que saber cómo se hace el cálculo.

Este nuevo método es que queremos trasladar a `Play`, que sigue estando dentro del límite de tamaño.

```python
class Play:
    def __init__(self, data):
        self._data = data

    def name(self):
        return self._data['name']

    def type(self):
        return self._data['type']

    def calculate_amount_for_comedy(self, audience):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy(audience)) \
            .add(Amount(300 * audience))

    def extra_amount_for_high_audience_in_comedy(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))

    def calculate_amount_for_tragedy(self, audience):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy(audience))

    def extra_amount_for_high_audience_in_tragedy(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(1000 * (audience - 30))

    def credits(self, audience):
        if "comedy" != self.type():
            return Credits(0)

        return Credits(math.floor(audience / 5))

    def amount(self, audience):
        if self.type() == "tragedy":
            return self.calculate_amount_for_tragedy(audience)
        if self.type() == "comedy":
            return self.calculate_amount_for_comedy(audience)

        raise ValueError(f'unknown type: {self.type()}')
```

Por supuesto, ahora queda más claro que nunca que `Play` necesita especializarse en dos clases `Tragedy` y `Comedy`. Pero no vamos a abordar ese cambio ahora, sino cuando la última regla nos lo pida.

## Solo dos argumentos

La función statement tiene varios problemas relacionados con esta regla. Claramente, tiene más de cinco líneas en el cuerpo, incluso sin contar las _inner functions_. Además, una de estas funciones recibe más de dos parámetros. Vamos a ver algunas soluciones. Aquí está:

```python
    def formatted_line(title, audience, amount):
        return f' {title}: {format_as_dollars(amount.current() / 100)} ({audience} seats)\n'
```

Recuerda que extrajimos esta función porque necesitábamos que alguien pudiese manejar `Amount` debido a la regla de no más de n punto por línea. Esto nos impediría usar la solución más inmediata que sería pasar el objeto `Performance`. Pero de hacerlo así volveríamos a romper la regla anterior. Por supuesto, hay más problemas ahí, pero de momento consideremos otras opciones.

Cuando una función recibe muchos parámetros una posibilidad es introducir un `Objeto parámetro`. Los contructores de los objetos no están limitados por esta regla, así que podríamos introducir algo como esto:

```python
class Line:
    def __init__(self, title, audience, amount):
        self.title = title
        self.audience = audience
        self.amount = amount.current()
```

Y cambiar la función `formatted_line` para usarlo:

```python
    def formatted(line):
        return f' {line.title}: {format_as_dollars(line.amount / 100)} ({line.audience} seats)\n'
```

Y se podría usar así:

```python
    printer.print(f'Statement for {invoice.customer()}\n')
    for performance in invoice.performances():
        printer.print(formatted(Line(performance.title(), performance.audience(), performance.amount())))
    printer.print(f'Amount owed is {format_as_dollars(invoice.amount().current() // 100)}\n')
    printer.print(f'You earned {invoice.credits().current()} credits\n')
```

Pero es que, además, ahora tendría todo el sentido mover esa función a `Line`.

```python
class Line:
    def __init__(self, title, audience, amount):
        self.title = title
        self.audience = audience
        self.amount = amount.current()

    def amount_as_dollars(self):
        return f"${self.amount/100:0,.2f}"

    def formatted(self):
        return f' {self.title}: {self.amount_as_dollars()} ({self.audience} seats)\n'
```

Este cambio genera algún problema porque duplicamos el código que da formato a `Amount` introduciendo el riesgo de que ocurran divergencias. Una forma de resolverlo podría ser introducir un patrón decorador:

```python
class FormattedAmount:
    def __init__(self, amount):
        self.amount = amount

    def dollars(self):
        return f"${self.amount.current() / 100:0,.2f}"
```

De modo que se pueda usar cuando sea necesario, haciendo un par de pequeños cambios:

```python
class Line:
    def __init__(self, title, audience, amount):
        self.title = title
        self.audience = audience
        self.amount = amount

    def formatted(self):
        return f' {self.title}: {FormattedAmount(self.amount).dollars()} ({self.audience} seats)\n'
```

```python
def statement(invoice_data, plays):
    invoice = Invoice(invoice_data, plays)

    printer = Printer()
    printer.print(f'Statement for {invoice.customer()}\n')

    for performance in invoice.performances():
        line = Line(performance.title(), performance.audience(), performance.amount())
        printer.print(line.formatted())

    printer.print(f'Amount owed is {FormattedAmount(invoice.amount()).dollars()}\n')
    printer.print(f'You earned {invoice.credits().current()} credits\n')

    return printer.output()
```

## Más oportunidades de acortar métodos

La función `statement` sigue siendo demasiado larga. Por supuesto, en ocasiones nos encontraremos con que es muy difícil o imposible hacer un método más pequeño por lo que se trata de no obsesionarse. Recordemos que estamos haciendo un ejercicio para entrenar nuestra capacidad de descubrir oportunidades para aplicar las reglas. ¿Tenemos algún punto más que podamos reducir?

Parte del problema con `statement` es que es una función y tiene un par de líneas de inicialización de objetos. Además, al ser una función nos complica la extracción de bloques de código. Por ejemplo, el bucle que procesa las `Performance` podría extraerse para mantener un único nivel de abstracción. Quizá podríamos introducir el concepto de `StatementPrinter` para llevarnos toda esa lógica de ahí y tener más libertad para manipularla.

```python
class StatementPrinter:
    def __init__(self, printer):
        self.printer = printer

    def print(self, invoice):
        self.printer.print(f'Statement for {invoice.customer()}\n')

        for performance in invoice.performances():
            line = Line(performance.title(), performance.audience(), performance.amount())
            self.printer.print(line.formatted())

        self.printer.print(f'Amount owed is {FormattedAmount(invoice.amount()).dollars()}\n')
        self.printer.print(f'You earned {invoice.credits().current()} credits\n')

        return self.printer.output()
```

De este modo, `statement` simplemente actúa como una especie de _caso de uso_:

```python
def statement(invoice_data, plays):
    invoice = Invoice(invoice_data, plays)

    statement_printer = StatementPrinter(Printer())

    return statement_printer.print(invoice)
```

Esto nos da algunas opciones. Por ejemplo:

```python
class StatementPrinter:
    def __init__(self, printer):
        self.printer = printer

    def print(self, invoice):
        self.printer.print(f'Statement for {invoice.customer()}\n')

        self.print_lines(invoice)

        self.printer.print(f'Amount owed is {FormattedAmount(invoice.amount()).dollars()}\n')
        self.printer.print(f'You earned {invoice.credits().current()} credits\n')

        return self.printer.output()

    def print_lines(self, invoice):
        for performance in invoice.performances():
            self.print_details(performance)

    def print_details(self, performance):
        line = Line(performance.title(), performance.audience(), performance.amount())
        self.printer.print(line.formatted())
```

Una cuestión es que `Printer` ahora se refiere a un mecanismo concreto de impresión, así que es mejor cambiarlo de nombre y ubicación. Por otro lado, `StatementPrinter`, `Line` o `FormattedAmount` son objetos que hemos introducido aunque aún no hemos ubicado correctamente.

```python
from domain.invoice import Invoice
from domain.statement_printer import StatementPrinter
from infrastructure.console_printer import ConsolePrinter


def statement(invoice_data, plays):
    invoice = Invoice(invoice_data, plays)

    statement_printer = StatementPrinter(ConsolePrinter())

    return statement_printer.print(invoice)
```

## Por qué funciona

La razón de que esta regla funcione es que al querer reducir el número de líneas que contiene una clase o un método nos obliga a buscar líneas de código muy relacionadas entre sí, o sea que mantengan alta cohesión, y que puedan moverse juntas a un nuevo método o incluso a otra clase. A un nuevo método si contribuyen a la misma responsabilidad de la clase, y a otra clase si representan una responsabilidad ajena.

Cuando separamos un gran bloque de código de una clase en métodos más pequeños altamente cohesivos es fácil identificar responsabilidades, de modo que podemos analizar si realmente corresponden a la clase o deberían irse a otro lugar. Estos métodos y clases más pequeños son más fáciles de testear porque tienden a hacer una sola cosa. También son más fáciles de mantener por su pequeño tamaño, ya que podemos entender de un vistazo su propósito y si algo va mal con ellos.

Por supuesto, no siempre es posible forzar un método a tener un determinado tamaño, incluso cuando tiene una responsabilidad bien definida y sus líneas tienen mucha cohesión. En cualquier caso, siempre es buena idea intentar analizar los métodos largos en busca de oportunidades de hacerlos más pequeños.

## El resultado

[Puedes consultar el proyecto en Github](https://github.com/franiglesias/theatrical-plays-kata)

---

# [No más de dos variables de instancia por clase](https://franiglesias.github.io/calisthenics-8/)

Otra regla que se presta a mucha discusión es esta y puede ser considerada un auténtico _tour de force_, porque ¿qué entidad de negocio no necesita una buena cantidad de propiedades? ¿Y pretendes que únicamente sean dos?

De nuevo, una regla de calisthenics nos propone una restricción especialmente artificial que nos obliga a reflexionar sobre nuestro diseño y cómo podríamos mejorarlo. [Un artículo anterior de este blog](https://franiglesias.github.io/calistenics-and-value-objects/) planteaba un ejercicio en el que se mostraba un ejemplo de cómo hacerlo en un tipo de datos bastante comunes en muchos negocios.

De hecho, en el ejemplo de las obras teatrales no tenemos más que un par de casos discutibles. Esto es debido en parte a lo reducido del problema, pero también porque hemos ido extrayendo todo el conocimiento a objetos pequeños.

## El caso de Performance

La clase Performance contiene tres propiedades o variables de instancia:

```python
class Performance:
    def __init__(self, audience, play):
        self._audience = audience
        self._play = play
        self._amount = None
```

Lo que nos encontramos en `Performance` es que las variables de instancia son, por decirlo así, irreconciliables. Representan cosas completamente diferentes. De hecho `_amount` tiene un significado puramente técnico, siendo una variable que usamos para poder realizar una optimización por lo que podríamos decir que Performance solo tiene dos propiedades: `_audience` y `_play`.

Precisamente, `Play` también tiene dos propiedades, aunque en este momento únicamente muestra una:

```python
class Play:
    def __init__(self, data):
        self._data = data

    def name(self):
        return self._data['name']

    def type(self):
        return self._data['type']
```

Esto es consecuencia de que simplemente hemos encapsulado una estructura de datos nativa, pero no significa que `Play` tenga una única propiedad. Sus dos propiedades se manifiestan en dos métodos _getter_, de los que tendremos que hablar en el siguiente artículo.

Vamos a refactorizar eso:

```python
class Play:
    def __init__(self, data):
        self._name = data['name']
        self._type = data['type']

    def name(self):
        return self._name

    def type(self):
        return self._type
```

Volvamos por un momento a `Performance`. Una consecuencia interesante de nuestro diseño es que el resto del programa no necesita saber de la existencia de `Play`, ya que todo el comportamiento de `statement` ocurre a través de `Performance`. Desde este punto de vista, `Play` sería irrelevante y podríamos haberla fusionado con `Performance`. De este modo, `Performance` podría tener este aspecto:

```python
class Performance:
    def __init__(self, audience, play_name, play_type):
        self._audience = audience
        self._title = play_name
        self._type = play_type
        self._amount = None
```

¿Recuerdas cuando `Performance` era demasiado grande porque se ocupaba de responsabilidades de `Play`? En aquel momento hubiésemos podido prescindir del objeto `Play` que entonces no era más que una simple _Data Class_ (un objeto que solo tiene datos pero no comportamiento) y podríamos haber fusionado sus propiedades con las de `Performance`.

Esencialmente, lo que quiero decir es que cuando una clase tiene muchas propiedades, es muy probable que esté tratando de ocuparse de demasiadas responsabilidades. Si agrupamos propiedades cohesivas y extraemos nuevas clases a partir de ellas, lo más seguro es que se llevarán consigo comportamientos de la clase contenedora.

Más pequeño y más simple.

## El caso de `Line`

Otra clase con más de dos propiedades es `Line`:

```python
class Line:
    def __init__(self, title, audience, amount):
        self.title = title
        self.audience = audience
        self.amount = amount
```

`Line` tiene tres propiedades por una buena razón, su tarea es algo así como representar un registro que tiene tres campos. Se trata de un ejemplo bastante claro de no poder reducir el número de variables por debajo del límite marcado.

Pero, ¿acaso `Line` no es la versión impresa de `Performance`? A lo mejor no necesitamos pasar las tres propiedades separadas, sino que `Performance` ya las agrupa. `Line` es como un decorador.

```python
class FormattedPerformance:
    def __init__(self, performance):
        self._performance = performance

    def formatted(self):
        return f' {self._performance.title()}: {FormattedAmount(self._performance.amount()).dollars()} ({self._performance.audience()} seats)\n'
```

Y la usaríamos así:

```python
def print_lines(self, invoice):
    for performance in invoice.performances():
        self.print_details(performance)

def print_details(self, performance):
    line = FormattedPerformance(performance)
    self.printer.print(line.formatted())
```

Este enfoque es interesante. Nos permite cumplir la regla de las dos variables de instancia reemplazando `Line` que tiene tres por `FormattedPerformance` que solo tiene una.

Pero todavía nos queda una regla que aplicar y va a poner en cuestión muchas de estas decisiones.

## Por qué funciona

Tanto esta como la regla anterior ponen énfasis en que las clases se ocupen de pocas cosas a la vez. Cuantas menos mejor. Para lograr eso nos fuerza a intentar cumplir con unos límites totalmente arbitrarios, que nos obligan a pensar en la cohesión de nuestro código.

La cohesión es el grado en que cada línea de código se relaciona con las demás dentro de su misma unidad (blqque, método, clase…). Cuando la cohesión es máxima, todas las líneas de código tienen que estar ahí, ninguna sobra. Para que esto ocurra, los bloques de código tienen que ser pequeños, minimizando la posibilidad de una parte del código realmente no esté contribuyendo a las responsabilidades de esa unidad.

Con las propiedades (o variables de instancia) ocurre lo mismo. Cuantas más haya en una clase, más probable es que exista una falta de cohesión. En algunos casos, el problema vendrá dado porque esas propiedades no corresponden realmente a esa clase. En otros casos, lo que ocurre es que algunas de esas propiedades son altamente cohesivas entre ellas, indicando que pueden agruparse en un objeto que represente un concepto al que contribuyen y que podemos extraer.

## El resultado

[Puedes consultar el proyecto en Github](https://github.com/franiglesias/theatrical-plays-kata)

---

# [No usar getters, setters o propiedades públicas](https://franiglesias.github.io/calisthenics-9/)

El objetivo de la regla es evitar que te bases en tu conocimiento del estado de los objetos de forma que acoples el resto del código a ese estado. En su lugar, los objetos solo deberían exponer comportamiento, minimizando la posibilidad de acoplarse a detalles de implementación. Por lo general, intentamos aplicar un principio llamado _Tell, don’t ask_, de modo que en lugar de preguntar a un objeto sobre su estado (ask), le pedimos que haga cosas.

En nuestro ejemplo hay varios casos de estos. Vamos a verlos y plantear posibles soluciones.

## El caso de Invoice y StatementPrinter

En este código podemos ver que `StatementPrinter` le pregunta un montón de cosas a `Invoice`. Podríamos decir que el comportamiento de `Invoice` parece ser darle información sobre su estado a `StatementPrinter`.

```python
class StatementPrinter:
    def __init__(self, printer):
        self.printer = printer

    def print(self, invoice):
        self.printer.print(f'Statement for {invoice.customer()}\n')

        self.print_lines(invoice)

        self.printer.print(f'Amount owed is {FormattedAmount(invoice.amount()).dollars()}\n')
        self.printer.print(f'You earned {invoice.credits().current()} credits\n')

        return self.printer.output()

    def print_lines(self, invoice):
        for performance in invoice.performances():
            self.print_details(performance)

    def print_details(self, performance):
        self.printer.print(FormattedPerformance(performance).formatted())
```

De hecho, `StatementPrinter` sabe muchas cosas de `Invoice`. Por ejemplo, sabe que `Invoice` tiene `Customer`, `Amount`, `Credits` e incluso `Performances`. Literalmente, conoce su estructura interna.

Para intentar aligerar ese conocimiento voy a empezar a separar cosas. Haré un ejemplo paso a paso con `Customer`. Lo primero es extraer una variable `customer` para no usar directamente la invocación a `Invoice.customer()`. Lo que quiero es que haya un método en `StatementPrinter` que pueda imprimir la línea del cliente sin saber nada directamente de `Invoice`.

```python
    def print(self, invoice):
        customer = invoice.customer()
        self.printer.print(f'Statement for {customer}\n')

        self.print_lines(invoice)

        self.printer.print(f'Amount owed is {FormattedAmount(invoice.amount()).dollars()}\n')
        self.printer.print(f'You earned {invoice.credits().current()} credits\n')

        return self.printer.output()
```

Ahora extraigo el método:

```python
    def print(self, invoice):
        customer = invoice.customer()
        self.fill_customer(customer)

        self.print_lines(invoice)

        self.printer.print(f'Amount owed is {FormattedAmount(invoice.amount()).dollars()}\n')
        self.printer.print(f'You earned {invoice.credits().current()} credits\n')

        return self.printer.output()

    def fill_customer(self, customer):
        self.printer.print(f'Statement for {customer}\n')
```

Y me deshago de la variable temporal:

```python
    def print(self, invoice):
        self.fill_customer(invoice.customer())

        self.print_lines(invoice)

        self.printer.print(f'Amount owed is {FormattedAmount(invoice.amount()).dollars()}\n')
        self.printer.print(f'You earned {invoice.credits().current()} credits\n')

        return self.printer.output()

    def fill_customer(self, customer):
        self.printer.print(f'Statement for {customer}\n')
```

Hago lo mismo con las demás datos:

```python
    def print(self, invoice):
        self.fill_customer(invoice.customer())

        self.print_lines(invoice)

        self.fill_amount(invoice.amount())
        self.fill_credits(invoice.credits())

        return self.printer.output()

    def fill_credits(self, credits):
        self.printer.print(f'You earned {credits.current()} credits\n')

    def fill_amount(self, amount):
        self.printer.print(f'Amount owed is {FormattedAmount(amount).dollars()}\n')

    def fill_customer(self, customer):
        self.printer.print(f'Statement for {customer}\n')
```

También modifico el método `print_lines` para mantener el paralelismo:

```python
class StatementPrinter:
    def __init__(self, printer):
        self.printer = printer

    def print(self, invoice):
        self.fill_customer(invoice.customer())
        self.fill_lines(invoice.performances())
        self.fill_amount(invoice.amount())
        self.fill_credits(invoice.credits())

        return self.printer.output()

    def fill_credits(self, credits):
        self.printer.print(f'You earned {credits.current()} credits\n')

    def fill_amount(self, amount):
        self.printer.print(f'Amount owed is {FormattedAmount(amount).dollars()}\n')

    def fill_customer(self, customer):
        self.printer.print(f'Statement for {customer}\n')

    def fill_lines(self, performances):
        for performance in performances:
            self.print_details(performance)

    def print_details(self, performance):
        self.printer.print(FormattedPerformance(performance).formatted())
```

Es cierto que seguimos haciendo llamadas de tipo _getter_ a `Invoice`, pero esto nos prepara para los siguientes pasos. Queremos no preguntarle cosas a `Invoice`. En su lugar, `Invoice` podría darle a `StatementPrinter` la información, sin desvelar sus detalles. Para ello usaremos un patrón _Visitor_.

Así que en lugar de preguntarle a `Invoice` por su información, esta rellena los datos que `StatementPrinter` necesita.

```python
    def print(self, invoice):
        invoice.fill(self)
        
        return self.printer.output()
```

De esta manera:

```python
    def fill(self, statement_printer):
        statement_printer.fill_customer(self.customer())
        statement_printer.fill_lines(self.performances())
        statement_printer.fill_amount(self.amount())
        statement_printer.fill_credits(self.credits())
```

Este cambio aún está incompleto porque todavía `StatementPrinter` sigue preguntando a `Performance`.

```python
    def fill_lines(self, performances):
        for performance in performances:
            self.print_details(performance)

    def print_details(self, performance):
        self.printer.print(FormattedPerformance(performance).formatted())
```

En parte tendríamos que deshacer lo que hicimos al aplicar reglas anteriores porque no queremos que `StatementPrinter` sepa ningún detalle. Así que vamos a reintroducir un método que imprima una línea de detalles de la performance:

```python
    def fill_line(self, title, amount, audience):
        self.printer.print(f' {title}: {FormattedAmount(amount).dollars()} ({audience} seats)\n')
```

De este modo, `Invoice` puede controlar el modo que se rellena `StatementPrinter`, que ya no necesita saber ni siquiera cuantas líneas necesitará imprimir, pues de eso se encargará `Invoice`.

```python
    def fill(self, statement_printer):
        statement_printer.fill_customer(self.customer())
        for performance in self.performances():
            statement_printer.fill_line(performance.title(), performance.amount(), performance.audience())
        statement_printer.fill_amount(self.amount())
        statement_printer.fill_credits(self.credits())
```

Así es como queda `StatementPrinter`:

```python
class StatementPrinter:
    def __init__(self, printer):
        self.printer = printer

    def print(self, invoice):
        invoice.fill(self)

        return self.printer.output()

    def fill_credits(self, credits):
        self.printer.print(f'You earned {credits.current()} credits\n')

    def fill_amount(self, amount):
        self.printer.print(f'Amount owed is {FormattedAmount(amount).dollars()}\n')

    def fill_customer(self, customer):
        self.printer.print(f'Statement for {customer}\n')

    def fill_line(self, title, amount, audience):
        self.printer.print(f' {title}: {FormattedAmount(amount).dollars()} ({audience} seats)\n')


class FormattedAmount:
    def __init__(self, amount):
        self.amount = amount

    def dollars(self):
        return f"${self.amount.current() / 100:0,.2f}"
```

Y así queda `Invoice`:

```python
class Invoice:
    def __init__(self, data, plays):
        self._data = data
        self._customer = data['customer']
        self._performances = Performances(data['performances'], Plays(plays))

    def customer(self):
        return self._customer

    def performances(self):
        return self._performances

    def amount(self):
        amount = Amount(0)
        for performance in self.performances():
            amount = amount.add(performance.amount())

        return amount

    def credits(self):
        volume_credits = Credits(0)
        for performance in self.performances():
            volume_credits = volume_credits.add(performance.credits())

        return volume_credits

    def fill(self, statement_printer):
        statement_printer.fill_customer(self.customer())
        for performance in self.performances():
            statement_printer.fill_line(performance.title(), performance.amount(), performance.audience())
        statement_printer.fill_amount(self.amount())
        statement_printer.fill_credits(self.credits())
```

Algunos comentarios sobre lo que acabamos de hacer:

- Ahora tenemos que los métodos de `Invoice` solo son llamados por `Invoice`, así que los podríamos marcar como privados. En Python podemos hacer esto prefijando sus nombres.
- Una pregunta legítima que podemos hacer es si `Invoice` ahora sabe demasiado de `StatementPrinter` dato que hay cuatro métodos que tiene que conocer para poder usarlo.
- 
- Para este caso específico podemos plantear esta solución. Al fin y al cabo, lo que hacemos con `StatementPrinter` es rellenar una plantilla. Podríamos tener entonces un método `fill` más genérico en el que indicamos que plantilla queremos rellenar. Algo similar a lo que se muestra a continuación. `Invoice` solo tiene que conocer un método:

```python
    def fill(self, statement_printer):
        statement_printer.fill('customer', self.customer())
        for performance in self.performances():
            statement_printer.fill('line', performance.title(), performance.amount(), performance.audience())
        statement_printer.fill('amount', self.amount())
        statement_printer.fill('credits', self.credits())
```

Y `StatementPrinter` ya no tiene que exponer detalles tampoco:

```python
    def fill(self, template, *args):
        getattr(self, 'fill_' + template)(*args)
```

¿Y qué pasa con `Performance`? Sigue exponiendo _getters_. Así que podríamos hacer algo similar:

```python
    def fill(self, statement_printer):
        statement_printer.fill('line', self.title(), self.amount(), self.audience())
```

Y ahora `Invoice` no tiene más que decirle a `Performance` que rellene su parte:

```python
    def fill(self, statement_printer):
        statement_printer.fill('customer', self.customer())
        for performance in self.performances():
            performance.fill(statement_printer)
        statement_printer.fill('amount', self.amount())
        statement_printer.fill('credits', self.credits())
```

A continuación, lo suyo sería hacer privados todos estos _getters_ o incluso eliminarlos.

El patrón de relación que nos ha quedado entre `Invoice` y `StatementPrinter` se llama _Double Dispatch_, pero podemos simplificar un poco las cosas de esta manera. `StatementPrinter` ya no sabe nada de `Invoice`:

```python
class StatementPrinter:
    def __init__(self, printer):
        self.printer = printer

    def print(self):
        return self.printer.output()

    def fill(self, template, *args):
        getattr(self, '_fill_' + template)(*args)

# ...
```

Y la función `statement` queda así, _and I think it’s beautiful_:

```python
from domain.invoice import Invoice
from domain.statement_printer import StatementPrinter
from infrastructure.console_printer import ConsolePrinter


def statement(invoice_data, plays):
    statement_printer = StatementPrinter(ConsolePrinter())

    invoice = Invoice(invoice_data, plays)
    invoice.fill(statement_printer)

    return statement_printer.print()
```

### El caso especial de `Play`

El problema con Play está aquí:

```python
    def credits(self, audience):
        if "comedy" != self.type():
            return Credits(0)

        return Credits(math.floor(audience / 5))

    def amount(self, audience):
        if self.type() == "tragedy":
            return self.calculate_amount_for_tragedy(audience)
        if self.type() == "comedy":
            return self.calculate_amount_for_comedy(audience)

        raise ValueError(f'unknown type: {self.type()}')
```

`Play` tiene que preguntarse “¿qué tipo de obra soy?”, para decidir como realizar el cálculo que le piden. Esto es muy similar a una violación del principio _Tell, don’t ask_, ya que tiene que consultar una propiedad para poder escoger el algoritmo adecuado.

Los objetos tienen propiedades por algo. Normalmente, la razón de ser de esas propiedades es ser capaces de regular el comportamiento del objeto. Las propiedades tienen un papel similar al de los coeficientes de una ecuación y operan junto con los parámetros que se pasan a los métodos para calcular un resultado.

Sin embargo, propiedades que modelan el _tipo_ de un objeto son harina de otro costal. Aportan el criterio para decidir qué algoritmo utilizar al realizar el cálculo. Pero si un objeto es de un _tipo_, esto debería reflejarse en el código por su _clase_. Cuando un objeto de una clase tiene _tipo_, y ese tipo determina la forma en que efectúa su comportamiento, lo que ocurre es que la clase debería tener variantes especializadas basadas en su tipo, ejecutando su comportamiento en su forma particular.

En nuestro ejemplo, está muy claro que hay dos tipos de obras: comedias y tragedias. Ambos tipos son obras teatrales, pero para los efectos de nuestro ejemplo, calculan sus importes y sus créditos de forma diferente.

¿Cómo podemos refactorizar `Play` para extraer las dos subclases? Vamos a ver un procedimiento bastante mecánico. En primer lugar duplicamos `Play` para crear la clase `Tragedy`, que extenderá de la misma `Play`:

```python
class Tragedy(Play):
    def __init__(self, data):
        self._name = data['name']
        self._type = data['type']

    def name(self):
        return self._name

    def type(self):
        return self._type

    def calculate_amount_for_comedy(self, audience):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy(audience)) \
            .add(Amount(300 * audience))

    def extra_amount_for_high_audience_in_comedy(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))

    def calculate_amount_for_tragedy(self, audience):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy(audience))

    def extra_amount_for_high_audience_in_tragedy(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(1000 * (audience - 30))

    def credits(self, audience):
        if "comedy" != self.type():
            return Credits(0)

        return Credits(math.floor(audience / 5))

    def amount(self, audience):
        if self.type() == "tragedy":
            return self.calculate_amount_for_tragedy(audience)
        if self.type() == "comedy":
            return self.calculate_amount_for_comedy(audience)

        raise ValueError(f'unknown type: {self.type()}')
```

El siguiente paso es reemplazar todas las condicionales sobre `self.type()` por `True`. En nuestro ejemplo, solo tenemos un caso en el método `amount`:

```python
    def amount(self, audience):
        if True:
            return self.calculate_amount_for_tragedy(audience)
        if self.type() == "comedy":
            return self.calculate_amount_for_comedy(audience)

        raise ValueError(f'unknown type: {self.type()}')
```

Probablemente, el IDE habrá empezado a señalar que la condicional es redundante porque ahora siempre se cumple. En mi caso está señalando que el resto del código del método no se ejecutará nunca. Así que podemos borrarlo:

```python
    def amount(self, audience):
        if True:
            return self.calculate_amount_for_tragedy(audience)
```

De hecho, nos sobra la condición:

```python
    def amount(self, audience):
        return self.calculate_amount_for_tragedy(audience)
```

Al hacer esto, dejamos de llamar a varios métodos, los que ejecutaríamos si el tipo fuese _comedy_. También los borramos porque no se llaman en más sitios. Nos va quedando esto:

```python
class Tragedy(Play):
    def __init__(self, data):
        self._name = data['name']
        self._type = data['type']

    def name(self):
        return self._name

    def type(self):
        return self._type

    def calculate_amount_for_tragedy(self, audience):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy(audience))

    def extra_amount_for_high_audience_in_tragedy(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(1000 * (audience - 30))

    def credits(self, audience):
        if "comedy" != self.type():
            return Credits(0)

        return Credits(math.floor(audience / 5))

    def amount(self, audience):
        return self.calculate_amount_for_tragedy(audience)
```

El siguiente paso es cambiar todas las condicionales que buscan tipos que no sean “tragedy” para reemplazarlas por `False`. En `Tragedy` ya no se da ese caso. Sin embargo, en `credits` tenemos una condición inversa que en el contexto de `Tragedy` equivale a comprobar si el tipo es _tragedy_. Así que en realidad, la condición siempre se cumplirá:

```python
    def credits(self, audience):
        if True:
            return Credits(0)

        return Credits(math.floor(audience / 5))
```

Todo el código fuera de la condición no se ejecuta y lo borramos, por lo que el resultante será:

```python
    def credits(self, audience):
        return Credits(0)
```

El método `amount` llama sin más a otro método, así que podríamos integrar este último, así como eliminar referencias superfluas en el nombre del método que nos dice el importe extra. `Tragedy` quedará así y podremos eliminar también la propiedad `type` y todo lo relacionado con ella:

```python
class Tragedy(Play):
    def __init__(self, data):
        self._name = data['name']

    def name(self):
        return self._name

    def extra_amount_for_high_audience(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(1000 * (audience - 30))

    def credits(self, audience):
        return Credits(0)

    def amount(self, audience):
        return Amount(40000).add(self.extra_amount_for_high_audience(audience))
```

Aplicamos el mismo tratamiento a `Comedy`. Empezamos duplicando `Play` y reemplazando todas las condicionales que verifican el tipo de tal modo que aquellas que chequean que el tipo es _comedy_ sean siempre `True` y las que no siempre `False`:

```python
class Comedy(Play):
    def __init__(self, data):
        self._name = data['name']
        self._type = data['type']

    def name(self):
        return self._name

    def type(self):
        return self._type

    def calculate_amount_for_comedy(self, audience):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy(audience)) \
            .add(Amount(300 * audience))

    def extra_amount_for_high_audience_in_comedy(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))

    def calculate_amount_for_tragedy(self, audience):
        return Amount(40000) \
            .add(self.extra_amount_for_high_audience_in_tragedy(audience))

    def extra_amount_for_high_audience_in_tragedy(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(1000 * (audience - 30))

    def credits(self, audience):
        if False:
            return Credits(0)

        return Credits(math.floor(audience / 5))

    def amount(self, audience):
        if False:
            return self.calculate_amount_for_tragedy(audience)
        if True:
            return self.calculate_amount_for_comedy(audience)

        raise ValueError(f'unknown type: {self.type()}')
```

A continuación, eliminaríamos todo el código muerto y que no se ejecuta porque ya no será llamado nunca.

```python
class Comedy(Play):
    def __init__(self, data):
        self._name = data['name']

    def name(self):
        return self._name

    def calculate_amount_for_comedy(self, audience):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience_in_comedy(audience)) \
            .add(Amount(300 * audience))

    def extra_amount_for_high_audience_in_comedy(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))

    def credits(self, audience):
        return Credits(math.floor(audience / 5))

    def amount(self, audience):
            return self.calculate_amount_for_comedy(audience)
```

Y rematamos integrando y cambiando nombres:

```python
class Comedy(Play):
    def __init__(self, data):
        self._name = data['name']

    def name(self):
        return self._name

    def extra_amount_for_high_audience(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))

    def credits(self, audience):
        return Credits(math.floor(audience / 5))

    def amount(self, audience):
        return Amount(30000) \
            .add(self.extra_amount_for_high_audience(audience)) \
            .add(Amount(300 * audience))
```

Ahora vamos a ver como utilizar las nuevas clases especializadas. El lugar en el que se instancian objetos Play es aquí:

```python
class Plays:
    def __init__(self, data):
        self._data = data

    def get_by_id(self, play_id):
        return Play(self._data[play_id])
```

Una forma sencilla sería introducir un método factoría en Play que nos entregue la subclase adecuada:

```python
class Play:
    # ...

    @staticmethod
    def create(data):
        if data['type'] == "tragedy":
            return Tragedy(data)
        if data['type'] == "comedy":
            return Comedy(data)

        raise ValueError(f'unknown type: {data["type"]}')

    # ...
```

Y usarla:

```python
class Plays:
    def __init__(self, data):
        self._data = data

    def get_by_id(self, play_id):
        return Play.create(self._data[play_id])
```

Ahora, no queda más que eliminar todos los métodos y propiedades innecesarias en `Play`:

```python
class Play:
    @staticmethod
    def create(data):
        if data['type'] == "tragedy":
            return Tragedy(data)
        if data['type'] == "comedy":
            return Comedy(data)

        raise ValueError(f'unknown type: {data["type"]}')

    def credits(self, audience):
        pass

    def amount(self, audience):
        pass
```

El método factoría `create` decide qué subtipo concreto de `Play` se usará. Si en el futuro necesitamos dar soporte a más tipos no tenemos más que añadir una nueva clase y una nueva condición.

## Por qué funciona

Esta regla suele ser más difícil de aceptar o entender si vienes de un estilo de programación procedural en el que conocer y controlar el estado lo es todo. Pero en programación orientada a objetos, cada objeto es responsable de su propio estado y de como implementa sus comportamientos. Por tanto, lo más importante es saber quién debe encargarse de qué, en lugar de tratar de obtener su estado y operar con él.

Cada objeto debe operar con su estado y comunicarse con otros objetos cuando necesite algo, o cuando quiera enviarles algo.

Al preguntar por una propiedad de otro objeto estamos acoplándonos a ese objeto, porque sabemos qué propiedad nos interesa y cómo obtenerla. Si usamos ese dato para un cálculo, es muy posible que ese objeto al que le preguntamos deba ejecutar ese cálculo. Por supuesto, puede ocurrir que el cálculo requiera alguna información del objeto que llama. Pero en ese caso la puede pasar como parámetro.

La regla de no usar _getters_, _setters_ o propiedades públicas, nos fuerza a pensar en los objetos como cajas negras a las que podemos pedirles que hagan cosas. En pocas palabras, la regla nos dice que no debemos acceder al estado interno los objetos del sistema. Si necesitamos algo de ellos, tenemos que poder pedirles que lo hagan, aportando información si es necesario. Algunos lenguajes como Ruby fuerzan que todas las propiedades de un objeto sean privadas por defecto, aunque es fácil introducir _getters_ o _setters_.

El hecho de no poder acceder al estado de los objetos es beneficioso para evitar el acoplamiento. Nos permite cambiar las implementaciones de los objetos de forma transparente al resto del sistema.

## El resultado

[Puedes consultar el proyecto en Github](https://github.com/franiglesias/theatrical-plays-kata)

---

# Aplicaciones prácticas

## [M](https://franiglesias.github.io/calistenics-and-value-objects/)[**ejorar el diseño de las clases**](https://franiglesias.github.io/calistenics-and-value-objects/)

> En este artículo presentamos un ejercicio que puede servir para adquirir soltura a la hora de reconocer patrones que nos permitan organizar objetos complejos mediante value objects.

Un problema típico para una aplicación de _e-commerce_ es el de requerir una dirección de envío y una dirección de facturación para cada pedido. Ambas son muy similares estructuralmente, pero diferentes en cuanto a significado. La dirección de envío es relevante para el sistema logístico y puede requerir un extra de instrucciones especiales, mientras que las de facturación es relevante en términos de legalidad fiscal.

Representar una dirección es un problema aparentemente simple pero en cuanto empiezas a escarbar un poco empiezan a salir todo tipo de inconveniencias. ¿Cuán complicada puede ser una dirección? Pues aparentemente mucho y eso ciñéndonos solo a direcciones postales españolas y sin considerar problemas de formato, simplemente intentando representarlas correctamente en nuestro dominio.

Supongo que este formulario te sonará, tanto en versión digital como en papel:

```plain text
Tipo de vía
Nombre de la vía
Número
Portal
Bloque
Escalera
Piso
Puerta
Código postal
Localidad
Provincia
```

_En algún caso he visto el campo extra: tipo de numeración, para indicar numeraciones de calles no estándar, como el punto kilométrico cuando la vivienda se encuentra en una carretera, pero con lo anterior creo que ya nos llega._

Aparte, la dirección necesita contar con el nombre del destinatario y, en el caso de la de facturación, con el identificador fiscal. Es decir, el formulario tendría esta pinta para la dirección de entrega:

```plain text
Nombre
Primer Apellido
Segundo Apellido
Teléfono de contacto
Tipo de vía
Nombre de la vía
Número
Portal
Bloque
Escalera
Piso
Puerta
Código postal
Localidad
Provincia
```

Y así para la dirección de facturación:

```plain text
Nombre
Primer Apellido
Segundo Apellido
NIF
Tipo de vía
Nombre de la vía
Número
Portal
Bloque
Escalera
Piso
Puerta
Código postal
Localidad
Provincia
```

Quince campos para cada dirección. No está mal.

¿Hay formas de simplificar eso? Sí. Una sencilla es agrupar algunos campos de modo que su número se reduzca. Veamos por ejemplo, una versión reducida de la dirección de envío:

```plain text
Nombre
Apellidos
Teléfono de contacto
Dirección
Código postal
Localidad
Provincia
```

Mucho mejor, ¿no? Al fin y al cabo, la dirección se puede expresar en un solo campo fácilmente y eso no supone ningún problema. ¿O sí?

¿Qué ocurre si el cliente solo pone el nombre de la calle y el número de portal, pero no indica el piso? ¿Y si en su edificio hay dos escaleras y no indica cuál? ¿Se entregará su paquete según lo previsto?

Estos problemas, y otros muchos similares, hacen que necesitemos que los clientes nos proporciones sus direcciones de una manera precisa y que no se olviden de ningún dato. Por eso, presentamos formularios con todos esos espacios para que el cliente sea consciente de todo lo que necesita decirnos para poder entregarle su pedido con el mínimo posible de demoras e inconvenientes. Así que en vez de lidiar con siete manejables campos nos toca hacerlo con quince por dirección, lo que hace un total de treinta campos sin contar el resto de los que necesita el pedido.

La siguiente restricción que quiero introducir viene de las [Object Calisthenics](https://williamdurand.fr/2013/06/03/object-calisthenics/). La Calistenia es un sistema de ejercicios destinados a [conseguir gracia y belleza en el movimiento](https://es.wikipedia.org/wiki/Calistenia). En nuestro caso son ejercicios que nos pueden servir para automatizar ciertas buenas prácticas. Su aplicación [contribuye a que el código se ajuste a mejores principios](https://keyvanakbary.com/object-calisthenics-mejora-tu-diseno-orientado-a-objetos/), como SOLID, Demeter o DRY, y, como mínimo, nos obliga a reflexionar.

En este caso vamos a aplicar dos de las reglas:

- Envolver primitivas en objetos
- Sólo dos variables de instancia

– ¿Cómo?

- Envolver primitivas en objetos
- Sólo dos variables de instancia

– A ver, es que la primera me suena hasta razonable, pero… ¿Dos variables de instancia? Estamos hablando de quince campos para especificar una dirección… ¿Y quieres que use solo dos variables de instancia?

– Exactamente. Pero deja que me explique…

La primera regla, a poco que la pienses, tiene que ver con los _Value Objects_. Podemos representar valores usando tipos escalares (`string`, `int`, `float`, etc.) y aunque es una práctica habitual tiene sus problemas. Los tipos escalares solo nos imponen algunas restricciones muy generales sobre los datos aceptables.

Pongamos por caso, el código postal. En España, el código postal es un número de cinco dígitos, lo que nos sugiere representarlo con un `int`. Pero, en realidad, normalmente es mejor representarlo con un tipo `string` porque aunque tiene forma numérica, no es un número.

Una razón es que puede empezar con un 0 y esto nos daría problemas para mantenerlo si usamos un valor de tipo `int`. Sin embargo, no tiene una semántica numérica, por así decir. El código postal usa los dos primeros dígitos para representar la provincia (así que va de 00 a 52) y los otros tres para identificar el distrito postal, de modo que una localidad grande puede tener varios códigos postales y un mismo código postal puede aplicarse a varias poblaciones muy pequeñas. En último término, un código postal:

- Es un tipo string
- Tiene 5 caracteres
- Todos son numéricos
- Los dos primeros (empezando por la izquierda) representan un número entre 00 y 52.
- Los tres restantes representan un número que va de 000 a un límite distinto según la provincia.

Ningún tipo escalar nos permite cumplir todas estas restricciones, por lo que es buena idea crear un _Value Object_ `PostalCode` que nos permita validar el código conforme a estas reglas y así asegurarnos que solo podemos crear códigos postales válidos o, al menos, reducir al mínimo la posibilidad de introducirlos no válidos.

En otros casos ocurre que un concepto se representa con varios campos de información que siempre van juntos. Así, por ejemplo, un nombre de persona es el nombre de pila y dos apellidos, siendo así que constituyen una unidad. Los tres campos se representan mediante tipos `string` y la forma de indicar que van juntos es reuniéndolos en un _Value Object_ `PersonName` que nos permita manejarlos como un todo.

– Pero habías hablado de dos variables de instancia, ¿cómo se come eso con un Value Object que necesitará tres?

La segunda regla dice que no usemos más de dos variables de instancia. Dos es un número totalmente arbitrario y, de hecho, es posible que no siempre podamos llegar a ese objetivo de una manera razonable o sostenible. La regla sería: minimizar la cantidad de variables de instancia, pero forzar un límite arbitrario nos obliga a no conformarnos y tratar de llegar al objetivo si es posible y tiene sentido. Recuerda: Calistenia es un ejercicio con el que automatizar una forma de escribir nuestro código.

Esto nos lleva de nuevo al recurso de los _Value Objects_. Como hemos visto hace un momento, algunos conceptos están representados con varios campos que siempre van juntos. Constituyen una unidad que se representa encapsulándolos en un único objeto.

En nuestro caso, tenemos 30 campos que modelan la dirección de envío y la dirección de facturación, así que empecemos a agruparlos:

```plain text
Dirección de envío
    Nombre
    Primer Apellido
    Segundo Apellido
    Teléfono de contacto
    Tipo de vía
    Nombre de la vía
    Número
    Portal
    Bloque
    Escalera
    Piso
    Puerta
    Código postal
    Localidad
    Provincia
Dirección de facturación
    Nombre
    Primer Apellido
    Segundo Apellido
    NIF
    Tipo de vía
    Nombre de la vía
    Número
    Portal
    Bloque
    Escalera
    Piso
    Puerta
    Código postal
    Localidad
    Provincia
```

En el primer nivel ya tendríamos solo dos variables de instancia (dirección de envío y dirección de facturación), así que vamos bien. Ahora vamos a agrupar las otras variables según si cambian juntas como un todo o no:

```plain text
Dirección de envío
    Destinatario
        Nombre
        Primer Apellido
        Segundo Apellido
        Teléfono de contacto
    Dirección Postal
        Tipo de vía
        Nombre de la vía
        Número
        Portal
        Bloque
        Escalera
        Piso
        Puerta
        Código postal
        Localidad
        Provincia
Dirección de facturación
    Titular
        Nombre
        Primer Apellido
        Segundo Apellido
        NIF
    Dirección Postal
        Tipo de vía
        Nombre de la vía
        Número
        Portal
        Bloque
        Escalera
        Piso
        Puerta
        Código postal
        Localidad
        Provincia
```

De nuevo, hemos conseguido reducir a dos las variables de instancia dentro de este nivel. Además, hemos identificado un tipo de dato que nos vale igualmente para ambos tipos de direcciones y que es la Dirección Postal.

Vamos a los siguientes niveles a ver qué descubrimos.

El Destinatario del envío puede representarse también con dos campos que, dentro de ese contexto, pueden cambiar de forma separada. Lo mismo ocurre en el caso del titular de la factura:

```plain text
Dirección de envío
    Destinatario
        Nombre de Persona
            Nombre
            Primer Apellido
            Segundo Apellido
        Teléfono de contacto
    Dirección Postal
        Tipo de vía
        Nombre de la vía
        Número
        Portal
        Bloque
        Escalera
        Piso
        Puerta
        Código postal
        Localidad
        Provincia
Dirección de facturación
    Titular
        Nombre de Persona
            Nombre
            Primer Apellido
            Segundo Apellido
        NIF
    Dirección Postal
        Tipo de vía
        Nombre de la vía
        Número
        Portal
        Bloque
        Escalera
        Piso
        Puerta
        Código postal
        Localidad
        Provincia
```

Y, otra vez, hemos conseguido reducir a dos variables de instancia. Pero hay algo más, hemos conseguido una nueva regularidad, ya que Nombre de Persona es un tipo de Value Object que podemos usar exactamente igual en ambos lados.

Vamos a forzar un poco aquí. El nombre de una persona puede estar formado por dos campos: nombre de pila y apellidos, de modo que podemos encapsular estos en un nuevo tipo de objeto Apellidos. Normalmente nos puede interesar tener campos separados para los dos apellidos debido a la existencia de apellidos compuestos y no siempre quedaría claro cuál es el punto de corte si viniesen en un único campo.

```plain text
Dirección de envío
    Destinatario
        Nombre de Persona
            Nombre
            Apellidos
                Primer Apellido
                Segundo Apellido
        Teléfono de contacto
    Dirección Postal
        Tipo de vía
        Nombre de la vía
        Número
        Portal
        Bloque
        Escalera
        Piso
        Puerta
        Código postal
        Localidad
        Provincia
Dirección de facturación
    Titular
        Nombre de Persona
            Nombre
            Apellidos
                Primer Apellido
                Segundo Apellido
        NIF
    Dirección Postal
        Tipo de vía
        Nombre de la vía
        Número
        Portal
        Bloque
        Escalera
        Piso
        Puerta
        Código postal
        Localidad
        Provincia
```

Lo hemos logrado de nuevo, dos variables en un objeto.

Vayamos a la Dirección Postal. Una forma de agrupar los campos tiene que ver con el aspecto de la dirección al que se refieren. Una parte indica la ubicación dentro de la localidad, mientras que la otra nos indica la localidad. Podríamos organizar esos campos así, lo que nos dejaría dos variables en el nivel de la Dirección Postal:

```plain text
Dirección de envío
    Destinatario
        Nombre de Persona
            Nombre
            Apellidos
                Primer Apellido
                Segundo Apellido
        Teléfono de contacto
    Dirección Postal
        Dirección
            Tipo de vía
            Nombre de la vía
            Número
            Portal
            Bloque
            Escalera
            Piso
            Puerta
        Localidad
            Código postal
            Localidad
            Provincia
Dirección de facturación
    Titular
        Nombre de Persona
            Nombre
            Apellidos
                Primer Apellido
                Segundo Apellido
        NIF
    Dirección Postal
        Dirección
            Tipo de vía
            Nombre de la vía
            Número
            Portal
            Bloque
            Escalera
            Piso
            Puerta
        Localidad
            Código postal
            Localidad
            Provincia
```

Ahora, tenemos unos cuantos campos en Dirección y podríamos intentar hacer grupos con ellos. Por ejemplo, Tipo de vía y Nombre de la vía, se refieren a un único concepto. Número. Portal y bloque, indican el acceso dentro de la finca (no se me ocurre como denominar este concepto de otra forma). Y escalera, piso y puerta, indican la vivienda del cliente dentro de un portal concreto de la finca.

```plain text
Dirección de envío
    Destinatario
        Nombre de Persona
            Nombre
            Apellidos
                Primer Apellido
                Segundo Apellido
        Teléfono de contacto
    Dirección Postal
        Dirección
            Via
                Tipo de vía
                Nombre de la vía
            Número
            Acceso
                Portal
                Bloque
            Vivienda
                Escalera
                Piso
                Puerta
        Localidad
            Código postal
            Localidad
            Provincia
Dirección de facturación
    Titular
        Nombre de Persona
            Nombre
            Apellidos
                Primer Apellido
                Segundo Apellido
        NIF
    Dirección Postal
        Dirección
            Via
                Tipo de vía
                Nombre de la vía
            Número
            Acceso
                Portal
                Bloque
            Vivienda
                Escalera
                Piso
                Puerta
        Localidad
            Código postal
            Localidad
            Provincia
```

Todavía podríamos hacer algo más. Vía y número forman una unidad que nos permite identificar la finca en la que se encuentra la dirección. Mientras que los otros dos, acceso y vivienda nos indican la ubicación dentro de la finca. Empieza a ser complicado encontrar nombres para estos conceptos:

```plain text
Dirección de envío
    Destinatario
        Nombre de Persona
            Nombre
            Apellidos
                Primer Apellido
                Segundo Apellido
        Teléfono de contacto
    Dirección Postal
        Dirección
            Finca
                Via
                    Tipo de vía
                    Nombre de la vía
                Número
            Ubicación
                Acceso
                    Portal
                    Bloque
                Vivienda
                    Escalera
                    Piso
                    Puerta
        Localidad
            Código postal
            Localidad
            Provincia
Dirección de facturación
    Titular
        Nombre de Persona
            Nombre
            Apellidos
                Primer Apellido
                Segundo Apellido
        NIF
    Dirección Postal
        Dirección
            Finca
                Via
                    Tipo de vía
                    Nombre de la vía
                Número
            Ubicación
                Acceso
                    Portal
                    Bloque
                Vivienda
                    Escalera
                    Piso
                    Puerta
        Localidad
            Código postal
            Localidad
            Provincia
```

Hemos forzado mucho las cosas, pero casi hemos conseguido hacer que cada objeto contenga tan solo dos atributos. Podríamos seguir, pero lo voy a dejar aquí en parte para que hagas el ejercicio de intentar dar este último paso.

---

## [Adelgazar las clases](https://franiglesias.github.io/calistenics-and-small-classes/)

> En este artículo presentamos un ejercicio que puede servir para adquirir soltura a la hora de escribir clases más compactas con métodos más expresivos y sencillos de entender.

La idea de estos ejercicios es imponer unas restricciones artificiales para forzar respuestas que nos obliguen a pensar más allá de las soluciones convencionales.

En este artículo vamos a aplicar estas restricciones:

- Un solo nivel de indentación
- No usar else
- Mantener las entidades pequeñas

### Ejercicios

Hace algún tiempo estuve practicando escribir algoritmos y estructuras de datos clásicos mediante TDD, y aunque se trata de clases relativamente pequeñas, tienen algunas estructuras que podrían mejorarse aplicando las restricciones propuestas, así que vamos a ver algunos ejemplos y cómo los podemos hacer evolucionar.

Como nota de interés decir que voy a hacer la mayor parte de los refactors con las herramientas que proporciona el IDE (PHPStorm).

Empecemos con `BubbleSort`, el algoritmo de ordenación más sencillo e intuitivo, aunque poco eficiente para muchos elementos:

```php
<?php

namespace Dsa\Algorithms\Sorting;

class BubbleSort
{
    public function sort(array $source): array
    {
        $length = count($source);
        for ($i = 0; $i < $length; $i++) {
            for ($j = 0; $j < $length; $j++) {
                if ($source[$i] < $source[$j]) {
                    $tmp = $source[$i];
                    $source[$i] = $source[$j];
                    $source[$j] = $tmp;
                }
            }
        }

        return $source;
    }
}
```

El test es este, por cierto, y está realizado con `PHPSpec`:

```php
<?php

namespace spec\Dsa\Algorithms\Sorting;

use Dsa\Algorithms\Sorting\BubbleSort;
use PHPSpec\ObjectBehavior;


class BubbleSortSpec extends ObjectBehavior
{
    function it_is_initializable()
    {
        $this->shouldHaveType(BubbleSort::class);
    }

    public function it_sorts_an_array_of_integers()
    {
        $source = [1123, 45, 76, 23, 87, 234, 34, 12, 36];
        $this->sort($source)->shouldBe([12, 23, 34, 36, 45, 76, 87, 234, 1123]);
    }
}
```

En este caso tenemos tres niveles de indentación y la restricción es que cada método solo puede tener uno como máximo. Para ello, extraeremos el `for` anidado a su propio método con ayuda del IDE, asegurándonos de que los tests siguen pasando:

```php
<?php

namespace Dsa\Algorithms\Sorting;

class BubbleSort
{
    public function sort(array $source): array
    {
        $length = count($source);
        for ($i = 0; $i < $length; $i++) {
            $source = $this->compareEveryElementWithCurrent($source, $length, $i);
        }

        return $source;
    }

    private function compareEveryElementWithCurrent(array $source, int $length, int $i): array
    {
        for ($j = 0; $j < $length; $j++) {
            if ($source[$i] < $source[$j]) {
                $tmp = $source[$i];
                $source[$i] = $source[$j];
                $source[$j] = $tmp;
            }
        }

        return $source;
    }
}
```

El test pasa, así que no hemos roto nada. Podemos ver algunos detalles mejorables en esta extracción:

- El parámetro `$length` es innecesario, ya que lo podemos obtener fácilmente de `$source`, así que lo omitiremos.
- Podríamos haber utilizado una estructura `foreach` en lugar de `for`.
- Cabe la posibilidad de pasar el elemento del array en vez de su índice.

Este es un punto interesante: el hecho de realizar la extracción del método nos provoca unas cuantas reflexiones sobre nuestras decisiones anteriores y posibles mejoras en el código. Así que antes de proseguir vamos a aplicar algunas.

```php
<?php

namespace Dsa\Algorithms\Sorting;

class BubbleSort
{
    public function sort(array $source): array
    {
        foreach ($source as $i => $element) {
            $source = $this->compareEveryElementWithCurrent($source, $i);
        }

        return $source;
    }

    private function compareEveryElementWithCurrent(array $source, int $i): array
    {
        $length = count($source);
        for ($j = 0; $j < $length; $j++) {
            if ($source[$i] < $source[$j]) {
                $tmp = $source[$i];
                $source[$i] = $source[$j];
                $source[$j] = $tmp;
            }
        }

        return $source;
    }
}
```

De momento, pasar el elemento y no el índice no parece viable, pero hemos conseguido algunas mejoras, ya que ahora no necesitamos calcular explícitamente la longitud de `$source`, lo que significa una línea menos y somos más explícitos en representar la idea de que recorremos el array. Seguimos manteniendo un único nivel de indentación que, además, solo tiene una línea.

Ahora tenemos dos niveles de indentación dentro del método `compareEveryElementWithCurrent`, así que vamos a tratarlos de la misma manera: extrayendo a un método.

```php
<?php

namespace Dsa\Algorithms\Sorting;

class BubbleSort
{
    public function sort(array $source): array
    {
        foreach ($source as $i => $element) {
            $source = $this->compareEveryElementWithCurrent($source, $i);
        }

        return $source;
    }

    private function compareEveryElementWithCurrent(array $source, int $i): array
    {
        $length = count($source);
        for ($j = 0; $j < $length; $j++) {
            $source = $this->swapElementsIfCurrentIsLower($source, $i, $j);
        }

        return $source;
    }

    private function swapElementsIfCurrentIsLower(array $source, int $i, int $j): array
    {
        if ($source[$i] < $source[$j]) {
            $tmp = $source[$i];
            $source[$i] = $source[$j];
            $source[$j] = $tmp;
        }

        return $source;
    }
}
```

Bien, hay un par de cosas interesantes por aquí:

- Solo tenemos un nivel de indentación en cada método o nivel de abstracción.
- En el nombre del método hacemos una referencia a un _current element_, así que estaría bien expresarla en código cambiando el nombre la variable `$i` a `$currentIndex` o similar, de forma que sea explícita la referencia.
- Podríamos aplicar el mismo tratamiento de convertir el `for` en `foreach` y ahorrarnos la variable `$length`.

Hagamos algunos de estos cambios:

```php
<?php

namespace Dsa\Algorithms\Sorting;

class BubbleSort
{
    public function sort(array $source): array
    {
        foreach ($source as $i => $element) {
            $source = $this->compareEveryElementWithCurrent($source, $i);
        }

        return $source;
    }

    private function compareEveryElementWithCurrent(array $source, int $currentElementIndex): array
    {
        foreach ($source as $j => $element) {
            $source = $this->swapElementsIfCurrentIsLower($source, $currentElementIndex, $j);
        }

        return $source;
    }

    private function swapElementsIfCurrentIsLower(array $source, int $i, int $j): array
    {
        if ($source[$i] < $source[$j]) {
            $tmp = $source[$i];
            $source[$i] = $source[$j];
            $source[$j] = $tmp;
        }

        return $source;
    }
}
```

Con esto hemos _aplanado_ los niveles de indentación en todos los métodos. La parte negativa es que la clase ha crecido en número de líneas, pero se ve compensado por el hecho de que cada método explica mejor qué hace y podemos profundizar en la explicación en la medida que necesitemos.

Hay un par de cosas que también merece la pena señalar:

- Vamos pasando el array `$source` de método en método para procesarlo y devolverlo. Nos surge la pregunta de si podría pasarse por referencia, para evitar los retornos y mantener los cambios o incluso guardarlo en la clase como propiedad y operar sobre ella. Respecto a esta última opción diría que no, ya que el algoritmo encapsulado en la clase no tiene por qué tener estado, y guardar el array sería otorgarle uno. Sobre lo de pasar el array por referencia, es una opción que dejaría el código un poco más conciso, pero tal vez tengamos otras maneras de hacerlo así que no lo vamos a aplicar por ahora.
- Por otro lado, en el método `swapElementsIfCurrentIsLower` nos queda un bloque de tres líneas que bien podría merecer ser extraído a su propio método para explicitar su intención.

De paso, extendemos el cambio de nombre de `$i` a todos sus usos para que quede más claro en todo momento:

```php
<?php

namespace Dsa\Algorithms\Sorting;

class BubbleSort
{
    public function sort(array $source): array
    {
        foreach ($source as $currentElementIndex => $element) {
            $source = $this->compareEveryElementWithCurrent($source, $currentElementIndex);
        }

        return $source;
    }

    private function compareEveryElementWithCurrent(array $source, int $currentElementIndex): array
    {
        foreach ($source as $j => $element) {
            $source = $this->swapElementsIfCurrentIsLower($source, $currentElementIndex, $j);
        }

        return $source;
    }

    private function swapElementsIfCurrentIsLower(array $source, int $currentElementIndex, int $j): array
    {
        if ($source[$currentElementIndex] < $source[$j]) {
            $source = $this->swapElements($source, $currentElementIndex, $j);
        }

        return $source;
    }

    private function swapElements(array $source, int $currentElementIndex, int $j): array
    {
        $tmp = $source[$currentElementIndex];
        $source[$currentElementIndex] = $source[$j];
        $source[$j] = $tmp;

        return $source;
    }
}
```

Gracias a este refactor, no tenemos que bajar a las tripas del mecanismo de intercambio para entender qué está pasando con los elementos.

A continuación vamos a intentar hacer algunas mejoras que nos ayuden a limpiar un poco el código y reducir el número de líneas y lo vamos a hacer aprovechando que cada paso del algoritmo está representado en su propio método.

Nos centraremos en `swapElements` y aplicaremos un tratamiento un poco radical, prescindiendo de pasar `$source` y de la variable temporal. Pasaremos los elementos a intercambiar por referencia al método y los reasignaremos mediante una pizca de la poca [syntactic sugar](https://en.wikipedia.org/wiki/Syntactic_sugar) que ofrece PHP:

```php
<?php

namespace Dsa\Algorithms\Sorting;

class BubbleSort
{
    public function sort(array $source): array
    {
        foreach ($source as $currentElementIndex => $element) {
            $source = $this->compareEveryElementWithCurrent($source, $currentElementIndex);
        }

        return $source;
    }

    private function compareEveryElementWithCurrent(array $source, int $currentElementIndex): array
    {
        foreach ($source as $j => $element) {
            $source = $this->swapElementsIfCurrentIsLower($source, $currentElementIndex, $j);
        }

        return $source;
    }

    private function swapElementsIfCurrentIsLower(array $source, int $currentElementIndex, int $j): array
    {
        if ($source[$currentElementIndex] < $source[$j]) {
            $this->swapElements($source[$currentElementIndex], $source[$j]);
        }

        return $source;
    }

    private function swapElements(int &$currentElement, int &$swapWith): void
    {
        [$swapWith, $currentElement] = [$currentElement, $swapWith];
    }
}
```

Se puede decir ahora que cada método tiene el mínimo de niveles de indentación posibles, así como el mínimo de líneas posible.

Es ahora cuando pasaremos `$source` por referencia, ahorrándonos varios `return`, excepto en el método público principal.

```php
<?php

namespace Dsa\Algorithms\Sorting;

class BubbleSort
{
    public function sort(array $source): array
    {
        foreach ($source as $currentElementIndex => $element) {
            $this->compareEveryElementWithCurrent($source, $currentElementIndex);
        }

        return $source;
    }

    private function compareEveryElementWithCurrent(array &$source, int $currentElementIndex): void
    {
        foreach ($source as $j => $element) {
            $this->swapElementsIfCurrentIsLower($source, $currentElementIndex, $j);
        }
    }

    private function swapElementsIfCurrentIsLower(array &$source, int $currentElementIndex, int $j): void
    {
        if ($source[$currentElementIndex] < $source[$j]) {
            $this->swapElements($source[$currentElementIndex], $source[$j]);
        }
    }

    private function swapElements(int &$currentElement, int &$swapWith): void
    {
        [$swapWith, $currentElement] = [$currentElement, $swapWith];
    }
}
```

### Eliminando else

En esta ocasión vamos a revisar un Binary Search Tree, que está a medio arreglar. Es decir: hay algunos primeros intentos de poner el código en mejor estado, pero aún había quedado mucho margen de mejora.

Como podemos ver, hay algunas zonas con dos niveles de indentación y bastantes usos de `else`.

```php
<?php

namespace Dsa\Structures;

class BinarySearchTree
{
    /**
     * @var BinarySearchNode
     */
    private $root;

    public function insert($value)
    {
        $new = new BinarySearchNode($value);
        if (! $this->root) {
            $this->root = $new;
        } else {
            $this->insertNew($this->root, $new);
        }
    }

    public function insertNew(BinarySearchNode $current, BinarySearchNode $new)
    {
        if ($new->getValue() < $current->getValue()) {
            if (! $current->getLeft()) {
                $current->setLeft($new);
            } else {
                $this->insertNew($current->getLeft(), $new);
            }
        } else {
            if (! $current->getRight()) {
                $current->setRight($new);
            } else {
                $this->insertNew($current->getRight(), $new);
            }
        }
    }

    public function isInTree($value)
    {
        if (! $this->root) {
            return false;
        }

        return $this->contains($this->root, $value);
    }

    public function contains(BinarySearchNode $current = null, $value)
    {
        if (! $current) {
            return false;
        }

        return $this->findNode($current, $value) ? true : false;
    }

    public function getParentValueOf($value)
    {
        if ($value == $this->root->getValue()) {
            return null;
        }

        return $this->findParent($this->root, $value)->getValue();
    }

    private function findParent(BinarySearchNode $current, $value)
    {
        if ($value < $current->getValue()) {
            if (! $current->getLeft()) {
                return null;
            } elseif ($current->getLeft()->getValue() == $value) {
                return $current;
            } else {
                return $this->findParent($current->getLeft(), $value);
            }
        } else {
            if (! $current->getRight()) {
                return null;
            } elseif ($current->getRight()->getValue() == $value) {
                return $current;
            } else {
                return $this->findParent($current->getRight(), $value);
            }
        }
    }

    private function findNode(BinarySearchNode $current = null, $value)
    {
        if (! $current) {
            return false;
        }

        if ($current->getValue() == $value) {
            return $current;
        } elseif ($value < $current->getValue()) {
            return $this->findNode($current->getLeft(), $value);
        } else {
            return $this->findNode($current->getRight(), $value);
        }
    }
}
```

El test que nos va a proteger en este proceso es el siguiente:

```php
<?php

namespace spec\Dsa\Structures;

use Dsa\Structures\BinarySearchTree;
use PHPSpec\ObjectBehavior;

class BinarySearchTreeSpec extends ObjectBehavior
{
    function it_is_initializable()
    {
        $this->shouldHaveType(BinarySearchTree::class);
    }

    public function it_knows_that_value_is_not_in_empty_tree()
    {
        $this->shouldNotBeInTree(34);
    }

    public function it_knows_that_value_is_in_tree_if_it_is_the_unique()
    {
        $this->insert(15);
        $this->shouldBeInTree(15);
    }

    public function it_knows_that_other_values_are_not_in_tree()
    {
        $this->insert(15);
        $this->shouldNotBeInTree(76);
    }

    public function it_knows_what_values_are_in_the_tree()
    {
        $this->prepareAnExampleTree();

        $this->shouldBeInTree(34);
        $this->shouldBeInTree(12);
        $this->shouldBeInTree(76);
    }

    public function it_knows_what_values_are_not_in_the_tree()
    {
        $this->prepareAnExampleTree();

        $this->shouldNotBeInTree(65);
        $this->shouldNotBeInTree(2);
        $this->shouldNotBeInTree(123);
    }

    public function it_knows_that_the_parent_of_the_root_is_null()
    {
        $this->prepareAnExampleTree();

        $this->getParentValueOf(15)->shouldBeNull();
    }

    public function it_can_find_the_parent_of_a_node()
    {
        $this->prepareAnExampleTree();

        $this->getParentValueOf(12)->shouldBe(7);
        $this->getParentValueOf(76)->shouldBe(68);
        $this->getParentValueOf(7)->shouldBe(15);
    }

    protected function prepareAnExampleTree(): void
    {
        $this->insert(15);
        $this->insert(34);
        $this->insert(7);
        $this->insert(68);
        $this->insert(12);
        $this->insert(4);
        $this->insert(76);
    }
}
```

La estructura Binary Search Tree se caracteriza porque cada nodo tiene dos hijos. Cuando se inserta un nodo se compara con el nodo raíz. Si no existe, porque todavía no se habían añadido elementos al árbol, se hace que el nuevo nodo sea la raíz. Si existe el nodo raíz, se inserta el nuevo nodo bajo él usando el método `insertNew`. Esto es lo que hace el método `insert`, que es con el que añadimos elementos al árbol.

Vamos allá. Empezamos con el método `insert`, que contiene un `else`:

```php
public function insert($value)
{
    $new = new BinarySearchNode($value);
    if (! $this->root) {
        $this->root = $new;
    } else {
        $this->insertNew($this->root, $new);
    }
}
```

En este caso, nos basta con regresar en la rama del `if`:

```php
public function insert($value): void
{
    $new = new BinarySearchNode($value);
    
    if (! $this->root) {
        $this->root = $new;
        return;
    }

    $this->insertNew($this->root, $new);
}
```

Alternativamente, podríamos haber invertido la condicional para que fuese positiva, que es más fácil de leer:

```php
public function insert($value): void
{
    $new = new BinarySearchNode($value);
    if ($this->root) {
        $this->insertNew($this->root, $new);
        return;
    }

    $this->root = $new;
}
```

El método `insertNew` añade nodos bajo un nodo determinado. En un árbol binario como este, cada nodo puede tener dos hijos (y así recursivamente), de tal modo que el nodo hijo de la izquierda contiene valores menores que el nodo padre, y el otro nodo contiene valores mayores. Si alguno de los nodos hijos ya existe, se le intenta añadir el nuevo nodo de forma recursiva, hasta que se encuentra una “rama” libre en la que colocarlo.

Hablando de `insertNew`, alcanza los dos niveles de indentación y tiene dos `else`, ¿qué podemos hacer al respecto?

```php
public function insertNew(BinarySearchNode $current, BinarySearchNode $new)
{
    if ($new->getValue() < $current->getValue()) {
        if (! $current->getLeft()) {
            $current->setLeft($new);
        } else {
            $this->insertNew($current->getLeft(), $new);
        }
    } else {
        if (! $current->getRight()) {
            $current->setRight($new);
        } else {
            $this->insertNew($current->getRight(), $new);
        }
    }
}
```

Ante de nada, voy a invertir las condicionales negativas:

```php
public function insertNew(BinarySearchNode $current, BinarySearchNode $new): void
{
    if ($new->getValue() < $current->getValue()) {
        if ($current->getLeft()) {
            $this->insertNew($current->getLeft(), $new);
        } else {
            $current->setLeft($new);
        }
    } else {
        if ($current->getRight()) {
            $this->insertNew($current->getRight(), $new);
        } else {
            $current->setRight($new);
        }
    }
}
```

Lo primero que podemos observar es que todas las patas de los condicionales no llevan a la salida y no hay procesamiento antes ni después. Dicho de otro modo, en cada pata podemos poner un `return`. Esto nos facilitará eliminar los `else` porque los hace innecesarios.

```php
public function insertNew(BinarySearchNode $current, BinarySearchNode $new): void
{
    if ($new->getValue() < $current->getValue()) {
        if ($current->getLeft()) {
            $this->insertNew($current->getLeft(), $new);
            return;
        }

        $current->setLeft($new);

        return;
    }

    if ($current->getRight()) {
        $this->insertNew($current->getRight(), $new);
        return;
    }

    $current->setRight($new);
}
```

El test demuestra que este cambio no afecta a la funcionalidad, hemos eliminado los `else` y uno de los casos de dos niveles de indentación.

Para poder aplanar el método necesitamos extraer los dos caminos de ejecución principales a sus propios métodos, haciendo explícita su intención:

```php
public function insertNew(BinarySearchNode $current, BinarySearchNode $new): void
{
    if ($new->getValue() < $current->getValue()) {
        $this->insertNewNodeAsLeftChild($current, $new);
    }

    $this->insertNodeAsRightChild($current, $new);
}
    
private function insertNewNodeAsLeftChild(BinarySearchNode $current, BinarySearchNode $new): void
{
    if ($current->getLeft()) {
        $this->insertNew($current->getLeft(), $new);

        return;
    }

    $current->setLeft($new);
}

private function insertNodeAsRightChild(BinarySearchNode $current, BinarySearchNode $new): void
{
    if ($current->getRight()) {
        $this->insertNew($current->getRight(), $new);

        return;
    }

    $current->setRight($new);
}
```

Las dos vías de ejecución de `insertNew` son ahora explícitas y su código es prácticamente el mismo. Esto es interesante porque pone de relieve una de las interpretaciones erróneas del principio `Don't Repeat Yourself`. El principio DRY se refiere a conocimiento, no a código, aunque a veces sean coincidentes. En este caso tenemos la misma estructura, pero significa cosas diferentes: cómo tratar un valor mayor y cómo tratar un valor menor.

El método `findParent` tiene varios problemas, dos niveles de indentación, condicionales anidadas y cinco `else`, aparte de algunos defectos que podemos arreglar de paso.

```php
private function findParent(BinarySearchNode $current, $value)
{
    if ($value < $current->getValue()) {
        if (! $current->getLeft()) {
            return null;
        } elseif ($current->getLeft()->getValue() == $value) {
            return $current;
        } else {
            return $this->findParent($current->getLeft(), $value);
        }
    } else {
        if (! $current->getRight()) {
            return null;
        } elseif ($current->getRight()->getValue() == $value) {
            return $current;
        } else {
            return $this->findParent($current->getRight(), $value);
        }
    }
}
```

Primero, una limpieza general. Cuando tenemos varios return deberíamos indicar el _return type_, a fin de garantizar que todos ellos son consistentes. En este caso el método puede devolver un `BinarySearchNode` o `null` si no se ha encontrado.

```php
private function findParent(BinarySearchNode $current, $value): ?BinarySearchNode
{
    if ($value < $current->getValue()) {
        if (! $current->getLeft()) {
            return null;
        } elseif ($current->getLeft()->getValue() === $value) {
            return $current;
        } else {
            return $this->findParent($current->getLeft(), $value);
        }
    } else {
        if (! $current->getRight()) {
            return null;
        } elseif ($current->getRight()->getValue() === $value) {
            return $current;
        } else {
            return $this->findParent($current->getRight(), $value);
        }
    }
}
```

Tenemos condicionales negativas. En este caso es un poco más delicado invertirlas al tener tres ramas y podría cambiar el comportamiento. Pero como tenemos tests, vamos a ver qué pasa si lo hacemos:

```php
private function findParent(BinarySearchNode $current, $value): ?BinarySearchNode
{
    if ($value < $current->getValue()) {
        if ($current->getLeft()) {
            if ($current->getLeft()->getValue() === $value) {
                return $current;
            } else {
                return $this->findParent($current->getLeft(), $value);
            }
        } else {
            return null;
        }
    } else {
        if ($current->getRight()) {
            if ($current->getRight()->getValue() === $value) {
                return $current;
            } else {
                return $this->findParent($current->getRight(), $value);
            }
        } else {
            return null;
        }
    }
}
```

La cosa empeora por un lado porque han aumentado los niveles de indentación, pero por otro lado ha mejorado porque algunos de los `else` se han vuelto completamente prescindibles, así que los quitamos sin más.

```php
private function findParent(BinarySearchNode $current, $value): ?BinarySearchNode
{
    if ($value < $current->getValue()) {
        if ($current->getLeft()) {
            if ($current->getLeft()->getValue() === $value) {
                return $current;
            } else {
                return $this->findParent($current->getLeft(), $value);
            }
        }
    } else {
        if ($current->getRight()) {
            if ($current->getRight()->getValue() === $value) {
                return $current;
            } else {
                return $this->findParent($current->getRight(), $value);
            }
        }
    }
}
```

Como todas las ramas tienen su `return` creo que será una buena idea quitar los 3 `else` que quedan:

```php
private function findParent(BinarySearchNode $current, $value): ?BinarySearchNode
{
    if ($value < $current->getValue()) {
        if ($current->getLeft()) {
            if ($current->getLeft()->getValue() === $value) {
                return $current;
            }

            return $this->findParent($current->getLeft(), $value);
        }
    }
    if ($current->getRight()) {
        if ($current->getRight()->getValue() === $value) {
            return $current;
        }

        return $this->findParent($current->getRight(), $value);
    }
}
```

Los tests siguen pasando y el método está un poquito más aplanado. En este momento, parece buena idea deshacer parte de lo avanzado antes, ya que si invierto algunas condicionales puedo reducir niveles de indentación sin extraer métodos:

```php
private function findParent(BinarySearchNode $current, $value): ?BinarySearchNode
{
    if ($value < $current->getValue()) {
        if (! $current->getLeft()) {
            return null;
        }

        if ($current->getLeft()->getValue() === $value) {
            return $current;
        }

        return $this->findParent($current->getLeft(), $value);
    }
    
    if (! $current->getRight()) {
        return null;
    }

    if ($current->getRight()->getValue() === $value) {
        return $current;
    }

    return $this->findParent($current->getRight(), $value);
}
```

Es posible que hubiésemos llegado al mismo punto de no haber invertido las condicionales al enfrentar por primera vez el método. Es lo de menos, lo importante es movernos de manera segura por el código.

Al hacer este cambio queda de manifiesto que el método tiene dos posibles flujos, así que podemos hacerlo explícito moviendo cada bloque a su propio método:

```php
private function findParent(BinarySearchNode $current, $value): ?BinarySearchNode
{
    if ($value < $current->getValue()) {
        return $this->findParentThroughLeftBranch($current, $value);
    }

    return $this->findParentThroughRightBranch($current, $value);
}
    
private function findParentThroughLeftBranch(BinarySearchNode $current, $value): ?BinarySearchNode
{
    if (! $current->getLeft()) {
        return null;
    }

    if ($current->getLeft()->getValue() === $value) {
        return $current;
    }

    return $this->findParent($current->getLeft(), $value);
}

private function findParentThroughRightBranch(BinarySearchNode $current, $value): ?BinarySearchNode
{
    if (! $current->getRight()) {
        return null;
    }

    if ($current->getRight()->getValue() === $value) {
        return $current;
    }

    return $this->findParent($current->getRight(), $value);
}
```

El nuevo `findParent` es ahora mucho más fácil de entender y también lo son sus ramas. Incluso las condicionales negativas actúan ahora como cláusulas de guarda lo que hace que resulte mucho más evidente su función (no hacer nada si no hay nada con lo que trabajar). Ambas ramas extraídas básicamente nos dicen que si el valor buscado coincide con el del hijo, izquierdo o derecho respectivamente, del nodo es que es su nodo padre. Y que si no coincide, siga buscando.

El último método que podemos planchar es `findNode`, que aquí muestro ya con los arreglos necesarios para ponerlo al día. Eliminar los `else` debería ser fácil:

```php
private function findNode(?BinarySearchNode $current, $value): ?BinarySearchNode
{
    if (! $current) {
        return null;
    }

    if ($current->getValue() === $value) {
        return $current;
    } elseif ($value < $current->getValue()) {
        return $this->findNode($current->getLeft(), $value);
    } else {
        return $this->findNode($current->getRight(), $value);
    }
}
```

El método busca el nodo que corresponde a un valor. Si el actual coincide lo devuelve y, si no, busca por el lado izquierdo si es menor y por el derecho si es mayor. El método _aplanado_ queda así:

```php
private function findNode(?BinarySearchNode $current, $value): ?BinarySearchNode
{
    if (! $current) {
        return null;
    }

    if ($current->getValue() === $value) {
        return $current;
    }

    if ($value < $current->getValue()) {
        return $this->findNode($current->getLeft(), $value);
    }

    return $this->findNode($current->getRight(), $value);
}
```

Nuestro `BinarySearchTree` es bastante menos intimidante ahora:

```php
<?php

namespace Dsa\Structures;

class BinarySearchTree
{
    /**
     * @var BinarySearchNode
     */
    private $root;

    public function insert($value): void
    {
        $new = new BinarySearchNode($value);
        if ($this->root) {
            $this->insertNew($this->root, $new);

            return;
        }

        $this->root = $new;
    }

    public function insertNew(BinarySearchNode $current, BinarySearchNode $new): void
    {
        if ($new->getValue() < $current->getValue()) {
            $this->insertNewNodeAsLeftChild($current, $new);
        }

        $this->insertNodeAsRightChild($current, $new);
    }

    public function isInTree($value): bool
    {
        if (! $this->root) {
            return false;
        }

        return $this->contains($this->root, $value);
    }

    public function contains(?BinarySearchNode $current, $value): bool
    {
        if (! $current) {
            return false;
        }

        return $this->findNode($current, $value) ? true : false;
    }

    public function getParentValueOf($value)
    {
        if ($value === $this->root->getValue()) {
            return null;
        }

        return $this->findParent($this->root, $value)->getValue();
    }

    private function findParent(BinarySearchNode $current, $value): ?BinarySearchNode
    {
        if ($value < $current->getValue()) {
            return $this->findParentThroughLeftBranch($current, $value);
        }

        return $this->findParentThroughRightBranch($current, $value);
    }

    private function findNode(?BinarySearchNode $current, $value): ?BinarySearchNode
    {
        if (! $current) {
            return null;
        }

        if ($current->getValue() === $value) {
            return $current;
        }

        if ($value < $current->getValue()) {
            return $this->findNode($current->getLeft(), $value);
        }

        return $this->findNode($current->getRight(), $value);
    }

    private function insertNewNodeAsLeftChild(BinarySearchNode $current, BinarySearchNode $new): void
    {
        if ($current->getLeft()) {
            $this->insertNew($current->getLeft(), $new);

            return;
        }

        $current->setLeft($new);
    }

    private function insertNodeAsRightChild(BinarySearchNode $current, BinarySearchNode $new): void
    {
        if ($current->getRight()) {
            $this->insertNew($current->getRight(), $new);

            return;
        }

        $current->setRight($new);
    }

    private function findParentThroughLeftBranch(BinarySearchNode $current, $value): ?BinarySearchNode
    {
        if (! $current->getLeft()) {
            return null;
        }

        if ($current->getLeft()->getValue() === $value) {
            return $current;
        }

        return $this->findParent($current->getLeft(), $value);
    }

    private function findParentThroughRightBranch(BinarySearchNode $current, $value): ?BinarySearchNode
    {
        if (! $current->getRight()) {
            return null;
        }

        if ($current->getRight()->getValue() === $value) {
            return $current;
        }

        return $this->findParent($current->getRight(), $value);
    }
}
```

---

# [**Comentarios finales**](https://franiglesias.github.io/calisthenics-10/)

> En este artículo reviso algunas cuestiones que se han planteado y exploro algunas líneas de desarrollo que quedaban pendientes.

## Aplicar object calisthenics, ¿mejora el diseño?

Sí, aplicando las reglas de _object calisthenics_ el diseño del código mejora incluso aunque no comencemos a introducir patrones de diseño. En otras palabras, _calisthenics_ te ayuda incluso si no tienes mucha experiencia en diseño de software.

En líneas generales, reducir el tamaño de los bloques de código y aplanar las estructuras indentadas ayuda a tener bloques y métodos más cohesivos centrados en torno a una responsabilidad.

Encapsular primitivos y estructuras de datos nativas abre la puerta a asignar mejor las responsabilidades y mover comportamientos a los objetos a los que corresponden.

## ¿Hay un orden adecuado para aplicar las reglas?

No. Las reglas se aplican según lo necesitamos o nos parece más evidente que se pueden aplicar. Muchas veces, aplicar una regla genera situaciones que se abordan aplicando otra. Así que en realidad, lo que hacemos es observar fragmentos de código que violan una u otra regla y los arreglamos lo mejor posible.

El proceso es, por tanto, iterativo. Empiezas aplicando una regla cuya utilidad ves clara y vas haciendo pequeños commits con los cambios que ves que mejoran tu código. En algún momento, descubrirás oportunidades para aplicar otras y así sucesivamente.

## ¿Por dónde empezar?

Empieza aplicando la regla que te resulte más fácil o cuyos casos sean más evidentes. Por ejemplo, no usar abreviaturas es fácil de aplicar en casi cualquier código. Aplanar estructuras condicionales suele ser muy evidente y el refactor _extraer método_ es sencillo de aplicar en un IDE moderno.

Encapsular tipos primitivos y estructuras de datos no es difícil, pero ya supone un trabajo extra porque tenemos que asegurar que en todos sus usos podemos hacer la sustitución. Sin embargo, una vez introducido un concepto como objeto, mover comportamiento viene de forma casi natural.

Eliminar la palabra clave ELSE puede ser complicado si no aislamos las estructuras condicionales previamente, para lo cual es bueno haber aplicado antes la regla de unh solo nivel de indentación.

No usar getter o setters puede ser muy sencillo en algunos casos, pero no es evidente como hacerlo en otros. En uno de los ejemplos de estos artículos, introdujimos el patrón Visitor para hacerlo, pero no es uno de los más sencillos de aplicar precisamente.

## ¿Debo aplicar las reglas exhaustivamente en todo el código?

No. Céntrate sobre todo en la lógica de dominio, que es la que más te interesa que sea fácil de entender y de mantener en el futuro. Las mejoras del código en esta área son más prioritarias, porque los objetos tienen mayor significación. En las partes de implementación de infraestructura, los beneficios pueden no ser tan importantes, lo que no debería justificar un diseño chapucero.

Usa tu buen juicio. Céntrate en el código que sea importante.

## Más consideraciones y ejemplos sobre algunas reglas

### Más sobre encapsular primitivas

Me he dejado algunos valores primitivos sin encapsular. El criterio de prioridad para encapsular primitivas sería algo así como: Encapsula primitivos en objetos cuando:

- El primitivo representa un concepto relevante del dominio o negocio de la aplicación
- El primitivo tiene reglas validación o comportamiento asociado que no es soportado por el propio tipo, lo que básicamente indica que el concepto es importante para el dominio

Por ejemplo, tras aplicar la última regla a `Play` y extender en dos subclases, quedó de manifiesto que el cálculo de importe extra en relación con la audiencia era un comportamiento asociado al concepto de Audiencia. De hecho, el IDE señala esos métodos como candidatos a ser métodos estáticos. Por ejemplo, en `Tragedy` es así:

```python
class Tragedy(Play):
    # ...

    def extra_amount_for_high_audience(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(1000 * (audience - 30))

    #...
```

Y en `Comedy`, así:

```python
class Comedy(Play):
    # ...

    def extra_amount_for_high_audience(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))

    # ...
```

Como se puede ver ninguno de los dos métodos depende de la clase que los contiene. Es cierto que podríamos extraer sus valores como propiedades de su clase. Sin embargo, fíjate que todo el cálculo se refiere solo al concepto de `Audience`. Hay un límite por encima del cual se genera un `Amount` extra. Si no se supera el límite `Amount` es cero.

Si igualamos la estructura de los métodos para que se parezcan lo más posible, quedaría una cosa así. Para `Tragedy`:

```python
    def extra_amount_for_high_audience(self, audience):
        if audience <= 30:
            return Amount(0)

        return Amount(0 + 1000 * (audience - 30))
```

Y para `Comedy`:

```python
    def extra_amount_for_high_audience(self, audience):
        if audience <= 20:
            return Amount(0)

        return Amount(10000 + 500 * (audience - 20))
```

Podríamos introducir una clase `Audience` que nos calcule el `Amount` extra, pasándole los parámetros necesarios:

```python
class Audience:
    def __init__(self, audience):
        self.audience = audience

    def amount(self, threshold, minimum, coeficient):
        if self.audience <= threshold:
            return Amount(0)

        return Amount(minimum + coeficient * (self.audience - threshold))
```

Y podemos usarlo:

```python
class Tragedy(Play):
    def __init__(self, data):
        self._name = data['name']

    def name(self):
        return self._name

    def extra_amount_for_high_audience(self, audience):
        return Audience(audience).amount(30, 0, 1000)
    
    def credits(self, audience):
        return Credits(0)

    def amount(self, audience):
        return Amount(40000).add(self.extra_amount_for_high_audience(audience))
```

O más simplificado:

```python
class Tragedy(Play):
    def __init__(self, data):
        self._name = data['name']

    def name(self):
        return self._name

    def credits(self, audience):
        return Credits(0)

    def amount(self, audience):
        return Amount(40000).add(Audience(audience).amount(30, 0, 1000))
```

Ahora tendría sentido introducir las propiedades de Tragedy que representan los parámetros:

```python
class Tragedy(Play):
    def __init__(self, data):
        self._name = data['name']
        self.threshold = 30
        self.minimum_amount = 0
        self.coefficient = 1000

    def name(self):
        return self._name

    def credits(self, audience):
        return Credits(0)

    def amount(self, audience):
        return Amount(40000).add(Audience(audience).amount(self.threshold, self.minimum_amount, self.coefficient))
```

Nos quedaría código relacionado con `Audience` y la posibilidad de instanciar el objeto desde el principio. Personalmente, cuando se trata de refactors suele empezar a introducirlo lo más adentro y voy “sacando” el objeto un paso cada vez. Así, por ejemplo, en el caso de `credits`, la lógica tiene que ver con Audience, pero no está tan claro como aplicar la relación.

### Más sobre límites de tamaño: parámetros y propiedades

Introducir `Audience` ha generado un problema, ya que la función para calcular el extra requiere tres parámetros y además hemos introducido tres propiedades más en las clases `Play`, ni más ni menos. En este caso, puede ser de aplicación el patrón `Parameter Object` para agruparlos. Sería algo así como `ExtraAmountData`:

```python
class ExtraAmountData:
    def __init__(self, threshold, minimum_amount, coefficient):
        self._threshold = threshold
        self._minimum_amount = minimum_amount
        self._coefficient = coefficient

    def threshold(self):
        return self._threshold

    def minimum_amount(self):
        return self._minimum_amount

    def coefficient(self):
        return self._coefficient
```

Esto se tendría que aplicar más o menos así. En Audience:

```python
    def extra_amount(self, extra_amount_data):
        if self.audience <= extra_amount_data.threshold():
            return Amount(0)

        return Amount(extra_amount_data.minimum_amount() + extra_amount_data.coeficient() * (
                    self.audience - extra_amount_data.threshold()))
```

Pero esto, sin embargo, no pinta bien. `Audience` no debería ser la responsable de calcular el extra, sino que es un dato necesario para hacerlo. Tendría más sentido que otro objeto dirija el cálculo sin exponer todos sus datos. Se podría considerar una especie de calculadora del importe extra basada en la audiencia, con coeficientes definidos por cada tipo de obra. Así que vamos a cambiar el concepto por completo.

```python
class ExtraAmountByAudience:
    def __init__(self, threshold, minimum_amount, coefficient):
        self._threshold = threshold
        self._minimum_amount = minimum_amount
        self._coefficient = coefficient

    def amount(self, audience):
        if audience <= self._threshold:
            return Amount(0)
        return Amount(self._minimum_amount + self._coefficient * (audience - self._threshold))
```

Y esto se usaría así:

```python
class Tragedy(Play):
    def __init__(self, data):
        self._name = data['name']

    def name(self):
        return self._name

    def credits(self, audience):
        return Credits(0)

    def amount(self, audience):
        return Amount(40000).add(ExtraAmountByAudience(30, 0, 1000).amount(audience))
```

Los tres parámetros de ExtraAmountByAudience son bastante crípticos. Una posible solución es usar un patrón builder:

```python
class ExtraAmountByAudience:
    def __init__(self):
        self._threshold = 0
        self._minimum_amount = 0
        self._coefficient = 1

    def when_audience_greater_than(self, threshold):
        self._threshold = threshold
        return self

    def minimum_amount_of(self, minimum_amount):
        self._minimum_amount = minimum_amount
        return self

    def and_coefficient(self, coefficient):
        self._coefficient = coefficient
        return self
```

Con lo cual, podemos hacer una construcción más expresiva:

```python
class Comedy(Play):
    def __init__(self, data):
        self._name = data['name']

    def name(self):
        return self._name

    def credits(self, audience):
        return Credits(math.floor(audience / 5))

    def amount(self, audience):
        calculator = ExtraAmountByAudience().
            when_audience_greater_than(20).
            minimum_amount_of(10000).
            and_coefficient(500)

        return Amount(30000)
            .add(calculator.amount(audience))
            .add(Amount(300 * audience))
```

¿Sobre-ingeniería?

## Comentarios de lectores

El objetivo de los artículos no era tanto llegar a un diseño de código `final`, como a mostrar que aplicando las reglas de Calisthenics es posible mejorar el diseño del software a través de dos caminos. El más simple consiste en aplicar las reglas tal cual. El segundo consiste en avanzar a partir de ese punto, descubriendo oportunidades para aplicar patrones de refactoring más avanzados.

Algunos lectores habéis comentado áreas en las que se podría mejorar el código.

### Acoplamiento temporal al imprimir la factura

[josemi](https://franiglesias.github.io/calisthenics-9/#comment-6009753237) hace un par de sugerencias interesantes. Por ejemplo, señala un caso de acoplamiento temporal dado que `StatementPrinter` no controla el orden en que se imprimen los elementos del `Statement`. Esto es debido a que no hay una separación entre la obtención de los datos y su impresión. El método `fill` obtiene el dato e imprime la línea. De ese modo, el control lo tiene `Invoice`, así que bastaría cambiar el orden de las llamadas en `Invoice` para _romper_ la impresión del Statement.

Esta es una primera aproximación muy basta, pero suficiente para hacernos a la idea y que elimina el acoplamiento temporal:

```python
class StatementPrinter:
    def __init__(self, printer):
        self._printer = printer
        self._customer = None
        self._amount = None
        self._credits = None
        self._lines = []

    def print(self):
        self._printer.print(f'Statement for {self._customer}\n')
        for line in self._lines:
            self._printer.print(f' {line["title"]}: {FormattedAmount(line["amount"]).dollars()} ({line["audience"]} seats)\n')

        self._printer.print(f'Amount owed is {FormattedAmount(self._amount).dollars()}\n')
        self._printer.print(f'You earned {self._credits.current()} credits\n')

        return self._printer.output()

    def fill(self, template, *args):
        getattr(self, '_fill_' + template)(*args)

    def _fill_credits(self, credits):
        self._credits = credits

    def _fill_amount(self, amount):
        self._amount = amount

    def _fill_customer(self, customer):
        self._customer = customer

    def _fill_line(self, title, amount, audience):
        self._lines.append({"title": title, "amount": amount, "audience": audience})
```

Ahora podría cambiar el orden de las líneas en Invoice, sin afectar al resultado:

```python
    def fill(self, statement_printer):
        for performance in self._performances:
            performance.fill(statement_printer)
        statement_printer.fill('credits', self._credits())
        statement_printer.fill('amount', self._amount())
        statement_printer.fill('customer', self._customer)
```

### Más sobre colecciones de primera clase

Otra sugerencia de [josemi](https://franiglesias.github.io/calisthenics-9/#comment-6009753237) es que la clase `Performances`, que contiene la colección de actuaciones se encargue también de controlar el orden en que se envían las líneas a `StatementPrinter`, en lugar de `Invoice`. Me parece una propuesta interesante. Sería una aplicación del principio _Tell, don’t ask_. `Invoice` le pide a `Performances` que realice la coordinación y cálculos que ahora mismo se hacen en `Invoice`, que quedaría así:

```python
from domain.performance import Performances
from domain.play import Plays


class Invoice:
    def __init__(self, data, plays):
        self._data = data
        self._customer = data['customer']
        self._performances = Performances(data['performances'], Plays(plays))

    def _amount(self):
        return self._performances.amount()

    def _credits(self):
        return self._performances.credits()

    def fill(self, statement_printer):
        statement_printer.fill('credits', self._credits())
        statement_printer.fill('amount', self._amount())
        statement_printer.fill('customer', self._customer)
        self._performances.fill(statement_printer)
```

Mientras que Performances podría quedar así, una vez eliminado el código para hacerla iterable que ya no es necesario:

```python
class Performances:
    def __init__(self, data, plays):
        self._data = data
        self._plays = plays

    def amount(self):
        amount = Amount(0)
        for data in self._data:
            performance = self._performance(data)
            amount = amount.add(performance.amount())

        return amount

    def _performance(self, data):
        return Performance(data['audience'], self._plays.get_by_id(data['playID']))

    def credits(self):
        volume_credits = Credits(0)
        for data in self._data:
            performance = self._performance(data)
            volume_credits = volume_credits.add(performance.credits())

        return volume_credits

    def fill(self, statement_printer):
        for data in self._data:
            performance = self._performance(data)
            performance.fill(statement_printer)
```

### El resultado

[Puedes consultar el proyecto en Github](https://github.com/franiglesias/theatrical-plays-kata)


