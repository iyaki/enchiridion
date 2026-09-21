---
title: "TalkingBit: La guía definitiva de los dobles de test"
notion_id: 1ab54f1c-7d23-818d-ac25-f2140a9fc111
notion_url: https://app.notion.com/p/TalkingBit-La-gu-a-definitiva-de-los-dobles-de-test-1ab54f1c7d23818dac25f2140a9fc111
last_edited: 2025-04-19T23:25:00.000Z
source_url: https://franiglesias.github.io/test-doubles-guide-2/
tags: ["Article", "The Talking Bit - Fran Iglesias", "Español", "Testing", "Object Oriented Programming"]
---
# Parte 1

En este artículo vamos a explicar todo lo que necesitar saber para utilizar dobles en tus tests.

Para empezar, tenemos que tocar algunas cuestiones teóricas, algo por lo que no me voy a disculpar, pues son imprescindibles para usar correctamente y beneficiarse del uso de dobles. Finalmente, veremos como abordar un caso que nos requiere utilizar distintos tipos de doble para abordar el testeo de un servicio.

## Deja de llamarlos Mocks

Lo primero de todo es un tema de _naming_. Deja de llamar _Mocks_ a todos los dobles de test. Los dobles de test pueden ser dummies, stubs, fakes, spies o mocks. _Mock_ no solamente es un tipo específico de doble de test, sino uno del que no vamos a hablar mucho en este artículo, porque puedes vivir sin él en la mayor parte de los casos.

## Principios de diseño y dobles de test

Para entender y manejar bien los dobles de test necesitas recurrir a varios principios y patrones de diseño de software y testing:

- **Separación Command-Query**: Este principio nos dice que toda función (o método de un objeto) puede ser o bien un comando, que produce un efecto, o cambio, en el sistema, o bien una query (pregunta), que obtiene y nos devuelve una información del mismo sistema. Pero no puede hacer ambas cosas a la vez. Es decir, un comando no puede devolver una respuesta, ni una query puede provocar un cambio en el sistema.
- **Composición**: El principio de composición nos dice que el comportamiento de un objeto es el resultado de la composición de los comportamientos de sus colaboradores. Sabiendo esto, en situaciones de test, debemos decidir si necesitamos aislar al sujeto del test de sus colaboradores.
- **Black-box testing**: Es el tipo de testing que se basa en observar la respuesta devuelta por la unidad bajo test o bien el efecto que ha provocado en el sistema. Se asume que no conocemos la implementación de esa unidad bajo test y solo examinamos sus efectos, por eso decimos que es una caja negra o **black box**.
- **White-box testing**: En este tipo de testing usamos nuestro conocimiento de la implementación de la unidad bajo test para decidir como abordamos las pruebas. Por ejemplo, analizando los flujos de ejecución para decidir qué casos vamos a testear, o bien haciendo aserciones sobre los mensajes que la unidad bajo test y sus colaboradores se pasan.
- **Principio de Inversión de Dependencias**: El principio de inversión de dependencias nos dice que siempre deberíamos depender de abstracciones (interfaces) y nunca de implementaciones. Esto nos permite introducir implementaciones alternativas a las que usamos en producción en los entornos de test.
- **Segregación de Interfaces**: Este principio nos pide diseñar interfaces estrechas (que tengan pocos métodos) a partir de las necesidades de sus consumidores.
- **Inyección de dependencias**: La inyección de dependencias es el patrón por el cual pasamos las dependencias a los objetos que las necesitan en lugar de que sean ellos quienes las instancian. De este modo, podemos explotar, entre otros, el principio de Inversión de Dependencias.
- **Fail fast**: Este principio nos dice que todo módulo que tiene un error debe comunicarlo inmediatamente al módulo que lo haya llamado, el cual tendrá el contexto para tomar la decisión de qué se ha de hacer con el error. Esto nos permite, entre otras cosas, simular fácilmente errores en situaciones de test.

## Dobles de test

Y, por fin, definamos lo que es un doble de test.

Un doble de test es un objeto que reemplaza un colaborador o dependencia de la unidad bajo de test de forma que podamos controlar su influencia en el comportamiento que estamos testeando. Sobre esto último hablaremos dentro de un momento.

En general, utilizaremos dobles de tests para reemplazar aquellas dependencias que suponen un coste, o un obstáculo, para la ejecución del test. Lo normal es que estas dependencias tengan que ver con tecnologías específicas que hemos usado para implementar nuestro sistema.

- Bases de datos de cualquier tipo
- Servicios de terceros a través de una conexión de red
- Servicios intrínsecamente lentos
- Servicios que puedan producir resultados no deterministas (que usualmente dependen de la máquina en la que se ejecuta el código)
- Otros

A continuación, permítime presentarte a los distintos tipos de doble.

### Cuando testeamos queries

Cuando testeamos una _query_ vamos a examinar el resultado que devuelve. En caso de necesitar un doble porque tenemos alguna dependencia (típicamente una base de datos o una API de terceros) usaremos principalmente Stub y Fake. Vamos a conocerlos:

### Stub

Un Stub es un objeto que puede reemplazar a una dependencia y que siempre devuelve una respuesta conocida. Esta respuesta puede:

- Estar pre-programada (hardcoded): el método suplantado devuelve siempre el mismo valor.
- Ser configurable: al construir el stub le pasamos lo que queremos que devuelva.
- Fallar: el método suplantado falla de una forma determinada que nos interesa controlar. Nuestro objetivo es verificar el comportamiento de la unidad bajo test en caso de que esa dependencia falle de esa manera particular.

### Fake

Un Fake es un objeto que implementa un comportamiento definido por una interfaz pero sin el coste de usar una tecnología del mundo real. Si además de eso, puede pasar los mismos tests que le haríamos a esa implementación real, estaríamos hablando de un Fake Verificado que, de hecho, podríamos llegar a usar en producción.

El ejemplo más típico es un repositorio implementado en memoria.

Los Fakes introducen mucho riesgo si no son verificados. Y en ese caso, pueden introducir gran complejidad.

### Cuando testeamos comandos

Cuando testeamos comandos estamos interesadas en verificar que se haya producido un cierto efecto en el sistema. Por desgracia no siempre es posible o conveniente comprobar este efecto.

Por ejemplo, si la unidad bajo test tiene que generar un archivo, siempre podríamos verificarlo examinando el sistema de archivos y cargando el archivo generado para ver si su ubicación y contenido son correctos. Esto es relativamente fácil, pero estos test en el entorno de CI puede ser una fuente de errores. Además, tienen un efecto sobre la velocidad de ejecución de los tests.

Otro ejemplo de comando podría ser el envío de una notificación por email, SMS, Slack o servicio similar. Verificar que un destinatario específico ha recibido esa notificación con el formato y contenido adecuado es, por lo general, impracticable. En su lugar, lo que verificamos es si hemos hecho uso del colaborador adecuado, pasándole la notificación correcta.

El tipo de dobles de test que usamos en esos casos suelen ser Dummies, Spies y Mocks.

### Dummy

Aunque pueda parecer paradójico, el Dummy es un doble de test que usamos cuando no queremos llamar al colaborador. O dicho de otra forma, usamos un Dummy cuando no esperamos que ese colaborador u objeto llegue a usarse, pero lo necesitamos para satisfacer una interfaz.

Los métodos del Dummy devuelven null o directamente fallan, en ese sentido un Dummy es como un **Stub**, pero la respuesta no es la propia de su rol, es un simple chivato. En ambos casos, el test debería fallar en caso de que usemos el objeto (repito, cuando esperamos no usarlo).

### Spy

Un objeto espía es un objeto que reemplaza a la dependencia o colaborador original con el objetivo de recabar información interna de lo que pasa dentro de la unidad bajo test. Su trabajo es tomar nota de las veces que ha sido llamado o de los parámetros de esas llamadas. Y podríamos añadir que debe de hacerlo sin levantar sospechas.

De este modo, una vez que ha terminado la ejecución de la unidad bajo test, no tenemos más que preguntarle al espía sobre la información de nuestro interés y con eso crear nuestras aserciones.

### Mock

Y llegamos por fin al Mock. El trabajo de un Mock es básicamente igual que el del espía, pero en lugar de pasar desapercibido queremos que monte un escándalo en caso de que no se use como esperamos, y eso sin que finalice la ejecución. Es decir, en el momento en el que el Mock detecta que algo no le cuadra, como que le pasan un parámetro que no es el que espera, tiene que lanzar un fallo y provocar que el test no pase.

Así que se podría decir que los Mocks llevan implícitas las aserciones.

El caso es que tanto Spies como Mocks verifican la forma en que la unidad bajo test se comunica con el colaborador doblado y esto añade fragilidad al test.

## ¿Cómo influye un colaborador en el comportamiento de la unidad bajo test?

Puede hacerlo de cuatro maneras:

**No haciendo nada**. En algunas circunstancias el colaborador no hace nada que afecte al comportamiento de la unidad bajo test (esta no se comunica con él). Y en algunos casos el efecto que tiene no nos preocupa. Ejemplos:

- Se le pasa un objeto a la unidad bajo test (o esta lo obtiene de otro colaborador) y simplemente se lo pasa a otro sin utilizarlo.
- En una situación de test, por las razones que sea, ese colaborador no llega a usarse, porque no hay nada que procesar o porque un paso anterior falla y nunca se llega a ese colaborador.
- Un Logger es un caso típico de un colaborador que se usa, pero del cual no nos preocupa normalmente el efecto que pueda tener en el test.

Este es el caso de uso de un Dummy.

**Devolviendo algo**. Con mucha frecuencia, un colaborador influye en el comportamiento produciendo algún resultado que la unidad bajo test necesita para poder completar su trabajo.

Este es el caso de uso de un Stub.

**Recibiendo un mensaje para producir un efecto**. Esto ocurre específicamente en los commands. Un colaborador recibe un mensaje de la unidad bajo test para que produzca un cambio en el sistema, que es lo que esperamos que suceda. Este cambio puede ser lo bastante costoso como para no querer que se produzca en la situación de test.

Este es el caso de uso de un Spy. También de un Mock, pero ¿para qué usar Mocks teniendo espías?

**Fallando**. Todos los colaboradores que puedan tener algún motivo para fallar lo harán alguna vez, por lo que debemos prepararnos para gestionar ese error.

Este es el caso de uso de Stub con un fallo programado.

## Test doble: ¡te elijo a ti!

Podemos seguir el sistema a continuación para escoger el test doble adecuado. Una vez identificado el colaborador que queremos sustituir, iremos haciendo las siguientes preguntas:

- El colaborador, ¿toca una tecnología del mundo real?. 
- **No**: usa directamente el colaborador.
- Sí: pasa a la siguiente pregunta.
- La unidad bajo test, ¿va a usar el colaborador en ese test? 
- **No**: usa un Dummy
- Sí: pasa a la siguiente pregunta.
- El colaborador, ¿debería fallar cuando le llamen? 
- **Sí**: usa un Stub con un fallo programado.
- No: pasa a la siguiente pregunta.
- El colaborador, ¿devuelve una respuesta? 
- **Sí**: usa un Stub.
- **No**: usa un Spy (o un Mock)

## Patrones de uso y buenas prácticas

Es importante reconocer que la función de los dobles de test es reemplazar colaboradores de la unidad bajo test de tal manera que nos permitan controlar su comportamiento de forma que sea predecible y económico en términos de performance y recursos.

En ese sentido, los dobles de tests se sitúan en esa frontera en la que tocamos tecnologías del mundo real, una frontera que los tests no deben traspasar.

Es recomendable escribir tus propios dobles de test y no usar librerías para ello. Las librerías pueden ser cómodas en algunos contextos, pero en la mayor parte de los casos no son necesarias y contribuyen a abusar de dobles excesivamente complejos.

En general, prefiere test sociales. Los tests sociales son aquellos en los que la unidad bajo test no es una única función o clase, sino que puede incluir colaboradores reales. Los dobles solo se usarían cuando tocamos tecnologías concretas o servicios que producen respuestas no deterministas (como el reloj del sistema o el generador aleatorio).

No dobles lo que no poseas. Es recomendable no doblar directamente librerías de tercera parte. Para hacer esto no te quedaría más remedio que usar librerías de dobles o librerías de mocks. En su lugar, introduce una abstracción y aplica el patrón adapter para utilizar la librería. En los tests, haz un doble basado en la abstracción.

Recuerda que los Spies y los Mocks introducen fragilidad en los tests, por lo que debes usarlo todo para testear únicamente el efecto deseado. Y, por supuesto, nunca los utilices cuando hagas tests de queries.

## Anti-patrones o smells en el uso de dobles

En mi experiencia, he identificado tres anti-patrones o smells cuando se usan dobles de test:

**Doble reutilizado**. Aunque hay algunos casos en los que se puede reutilizar el mismo doble, esto introduce el riesgo de añadir complejidad y acoplamiento entre distintos tests. Por ejemplo, si tenemos que añadir alguna lógica en el doble que depende del test en el que se vaya a utilizar.

Como regla general, intenta introducir los dobles necesarios para el test específico en el que estás interesada. Es preferible una cierta duplicación de lógica trivial.

**Doble sabihondo**, Esto ocurre cuando se introduce mucho conocimiento de dominio en el doble. Es decir, cuando ponemos lógica en el doble basada en la misma lógica de dominio que tiene el colaborador original. En el caso de los Stubs, siempre debes simular el mínimo: devolver un valor conocido, sin ninguna lógica que lo calcule.

**Demasiadas expectativas**. He visto muchos tests en las que se establecen expectativas sobre todas y cada una de las llamadas que hace la unidad bajo test a sus colaboradores, independientemente de si es un efecto esperado o de si se trata simplemente de un stub. Esto introduce muchas fuentes potenciales de fallo del test que no tienen que ver con el comportamiento que se está observando.

Esto es debido, sobre todo, a un mal uso de las librerías de dobles y el abuso de Mocks.

Es mucho mejor usar espías y centrarse en el efecto que ese test está verificando.

## Un caso práctico: el servicio de felicitación de cumpleaños

El siguiente es un ejercicio que he diseñado para practicar la introducción de dobles de test. En el ejercicio presentamos un servicio que obtiene los clientes que cumplen años en una fecha dada y les envía un email que incluye un código de descuento generado al vuelo.

Se incluye el esqueleto de un par de tests que hay que terminar, dado que ninguno verifica que se produzca el efecto deseado y que no es otro que se hayan enviado tantos emails como clientes cumpliendo años ese día. Además, habría que añadir al menos un tercer test que compruebe que el mensaje se construye correctamente y se envía al destinatario esperado.

```typescript
/*
 * README
 * Exercise:
 * We have a BirthdayService that runs every day via a cron job
 *
 *  It greets customers with has birthday on that day.
 *  It generates a discount code for them.
 *  It sends an email to them with the discount code.
 *  It logs the email sent.
 *
 * You work is to write the required tests for this functionality.
 * You probably will need to modify the code to make it testable.
 * Use different test doubles for the dependencies.
 *
 * Start by running the test below and fixing the errors.
 * Add assertions to the test that matches the intent of the test.
 *
 * Maybe you need to apply some refactorings to make the code testable in line with the Small Safe Steps workshop.
 *
 * Enrich the exercise by adding more tests:
 *
 * * Make a test to ensure that the service sends the correct email content to the right customer
 * * Make a test to ensure that the service fails gracefully if the email sending fails
 * * Make a test to ensure that the service fails gracefully if the repository fails
 *
 * */

describe('Birthday greetings', () => {
    it('should not send greeting emails if no customer has birthday today', () => {
        const service = new BirthdayService(
            new Customers([]),
            new ProductionEmailSender(),
            new ProductionLogger(),
        )
        service.greetCustomersWithBirthday(new Date())
    })

    it('should send greeting emails to all customers with birthday today', () => {
        const service = new BirthdayService(
            new Customers([
                new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
                new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
            ]),
            new ProductionEmailSender(),
            new ProductionLogger(),
        )

        service.greetCustomersWithBirthday(new Date())
    })
})

```

Cuando intentamos ejecutar el test ocurre lo siguiente:

```plain text
Error: 🤦🏽‍♀️ You are using ProductionCustomerRepository in a test. It will mess up our data.

```

Tal como está escrito el test, se están usando los colaboradores reales. Aun asumiendo que se trata de un entorno de desarrollo local, estaríamos pagando una penalización de rendimiento en el test, por no hablar de la necesidad de preparar la base de datos con una muestra adecuada de registros.

Vamos a centrarnos en el primer test:

```typescript
it('should not send greeting emails if no customer has birthday today', () => {
    const service = new BirthdayService(
        new Customers([]),
        new ProductionEmailSender(),
        new ProductionLogger(),
    )
    service.greetCustomersWithBirthday(new Date())
})

```

En este primer test suponemos que no se encuentra ningún cliente que cumpla años en el día de hoy, por lo que no se esperaría enviar ningún mensaje. Lo propio sería comprobar que `ProductionEmailSender` no ha recibido llamadas, por lo que tendríamos que sustituirlo por un espía. Y, antes de eso, tenemos que usar un `Customers` que no sea una implementación de producción. Tenemos bastante trabajo por delante y lo primero sería ver qué pasa dentro de `BirthdayService`:

```typescript
export class BirthdayService {
  private readonly customerRepository: Customers
  private readonly emailSender: ProductionEmailSender
  private readonly logger: ProductionLogger

  constructor(
          customerRepository: Customers,
          emailSender: ProductionEmailSender,
          logger: ProductionLogger,
  ) {
    this.customerRepository = customerRepository
    this.emailSender = emailSender
    this.logger = logger
  }

  greetCustomersWithBirthday(today: Date) {
    const customers = this.customerRepository.findWithBirthday(today)
    customers.forEach((customer) => {
      const discountCode = new DiscountCodeGenerator().generate()
      const template =
              'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                      '{discount}',
                      discountCode.getCode(),
              )
      customer.sendEmail(template, this.emailSender)
      this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
    })
  }
}

```

Lo primero que hace `BirthdayService` es invocar el `findWithBirthday` de su colaborador `customerRepository`, que es de tipo `Customers`. Luego itera la colección de `Customer`, pero como en este primer ejemplo tal colección está vacía podemos despreocuparnos de momento.

Esperaríamos que `Customers` fuese una abstracción… Pero no. Es una implementación concreta y no existe ninguna abstracción.

```typescript
export class Customers {
  private readonly customers: Customer[]

  constructor(customers: Customer[]) {
    this.customers = customers
  }

  findWithBirthday(today: Date): Customer[] {
    throw new Error(
            '🤦🏽‍♀️ You are using ProductionCustomerRepository in a test. It will mess up our data.',
    )
  }
}

```

### Inversión de Dependencias: Customers

Por tanto, nuestro primer objetivo es que se aplique el _Principio de Inversión de Dependencias_. Para ello, hay que introducir una interfaz, hacer que `BirthdayService` depende de ella y así poder reemplazar el servicio `Customers` por un doble de test.

En mi caso, lo primero que hago es cambiar el nombre de `Customers` por `ProductionCustomers`, de este modo dejo claro que se trata de una implementación. Además `Customers`, me parece un nombre mejor para la interfaz.

Usando _IntelliJ_ es posible usar los refactors automáticos, _Rename_ y _Extract Interface_, por lo que no me preocupa que los tests no estén pasando.

El cambio de nombre:

```typescript
export class ProductionCustomers {
  private readonly customers: Customer[]

  constructor(customers: Customer[]) {
    this.customers = customers
  }

  findWithBirthday(today: Date): Customer[] {
    throw new Error(
            '🤦🏽‍♀️ You are using ProductionCustomerRepository in a test. It will mess up our data.',
    )
  }
}

```

La interfaz:

```typescript
export interface Customers {
  findWithBirthday(today: Date): Customer[]
}
```

Por último, tenemos que hacer que BirthdayService dependa de Customers:

```typescript
export class BirthdayService {
  private readonly customerRepository: Customers
  private readonly emailSender: ProductionEmailSender
  private readonly logger: ProductionLogger

  constructor(
          customerRepository: Customers,
          emailSender: ProductionEmailSender,
          logger: ProductionLogger,
  ) {
    this.customerRepository = customerRepository
    this.emailSender = emailSender
    this.logger = logger
  }

  greetCustomersWithBirthday(today: Date) {
    // Removed for clarity
  }
}

```

### Introducción de un Stub de Customers

Ahora que ya tenemos la dependencia invertida es sencillo introducir un doble de test adecuado para este test. En nuestro caso, lo que queremos es una implementación de la interfaz `Customers` que devuelva una colección vacía de objetos `Customer`, por tanto, vamos a introducir un _Stub_.

```typescript
class NoBirthdayTodayCustomers implements Customers {
  findWithBirthday(_: Date): Customer[] {
    return []
  }
}

```

E inyectarlo en el servicio para el caso del test:

```typescript
it('should not send greeting emails if no customer has birthday today', () => {
  const service = new BirthdayService(
          new NoBirthdayTodayCustomers(),
          new ProductionEmailSender(),
          new ProductionLogger(),
  )
  service.greetCustomersWithBirthday(new Date())
})

```

Y con este cambio ya se puede ejecutar el test. Si examinamos el código podemos ver que debido a que en este caso de test no tenemos elementos en la colección, el bucle nunca se ejecuta y no se llaman a los otros colaboradores.

```typescript
export class BirthdayService {
  private readonly customerRepository: Customers
  private readonly emailSender: ProductionEmailSender
  private readonly logger: ProductionLogger

  constructor(
          customerRepository: Customers,
          emailSender: ProductionEmailSender,
          logger: ProductionLogger,
  ) {
    this.customerRepository = customerRepository
    this.emailSender = emailSender
    this.logger = logger
  }

  greetCustomersWithBirthday(today: Date) {
    const customers = this.customerRepository.findWithBirthday(today)
    customers.forEach((customer) => {
      const discountCode = new DiscountCodeGenerator().generate()
      const template =
              'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                      '{discount}',
                      discountCode.getCode(),
              )
      customer.sendEmail(template, this.emailSender)
      this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
    })
  }
}

```

### Espiando el side-effect

De todos modos, nuestro test está incompleto porque no se realiza ninguna comprobación. Es hora de introducirla. En este caso el comportamiento esperado es que no se hagan llamadas al `ProductionEmailSender`, y para ello necesitamos introducir un espía al que podamos preguntarle cuantas veces le hemos pedido que envíe emails.

Por supuesto, ProductionEmailSender es una implementación concreta, así que vamos a ver qué podemos hacer:

```typescript
export class ProductionEmailSender implements EmailSender {
  send(email: string, message: string) {
    throw new Error(
            '🤬 You are using ProductionEmailSender in a test. It will cost lots of money $$.',
    )
  }
}

```

Por suerte, `ProductionEmailSender` implementa una interfaz `EmailSender`, aunque BirthdayService todavía depende de la implementación. Así que tenemos que completar la inversión:

```typescript
export class BirthdayService {
  private readonly customerRepository: Customers
  private readonly emailSender: EmailSender
  private readonly logger: ProductionLogger

  constructor(
          customerRepository: Customers,
          emailSender: EmailSender,
          logger: ProductionLogger,
  ) {
    this.customerRepository = customerRepository
    this.emailSender = emailSender
    this.logger = logger
  }

  greetCustomersWithBirthday(today: Date) {
    // Removed for clarity
  }
}

```

Esto nos permitirá plantear el test de esta forma:

```typescript
it('should not send greeting emails if no customer has birthday today', () => {
  const emailSender = new MessageCountingEmailSender()
  const service = new BirthdayService(
          new NoBirthdayTodayCustomers(),
          emailSender,
          new ProductionLogger(),
  )
  service.greetCustomersWithBirthday(new Date())

  expect(emailSender.countOfSentEmails()).toBe(0)
})

```

Lo que estamos haciendo es introducir una implementación de `EmailSender` que sea capaz de contar la cantidad de veces que se invoca el método `send` y al que le podamos preguntar una vez que se ha ejecutado el servicio.

```typescript
class MessageCountingEmailSender implements EmailSender {
  private msgSent: number = 0
  send(email: string, message: string): void {
    this.msgSent++
  }

  countOfSentEmails(): number {
    return this.msgSent
  }
}

```

Obviamente como el bucle no se ejecuta el contador de mensajes no va a registrar ninguno, así que el test pasa. Podemos verificar que el test es válido porque si cambiamos la línea de `expect` para probar otros valores, el test falla.

Con esto, podemos dar por terminado este test y pasar al siguiente.

### Un nuevo Stub para Customers

Para ejecutar el segundo test necesitaremos un _Stub_ que nos entregue una colección de `Customer`. En este caso, vamos a hacer que tenga una cierta capacidad de configuración:

```typescript
class CustomersWithBirthdayToday implements Customers {
  private customers: Customer[]
  constructor(customers: Customer[]) {
    this.customers = customers
  }

  findWithBirthday(_: Date): Customer[] {
    return this.customers
  }
}

```

Y lo usamos en el test:

```typescript
  it('should send greeting emails to all customers with birthday today', () => {
  const service = new BirthdayService(
          new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
          ]),
          new ProductionEmailSender(),
          new ProductionLogger(),
  )

  service.greetCustomersWithBirthday(new Date())
})

```

Al ejecutar el test nos encontramos con este mensaje. Si bien hemos cambiado `ProductionCustomers` por su doble, no lo hemos hecho todavía con `ProductionEmailSender`, por lo que el test fallará:

```plain text
Error: 🤬 You are using ProductionEmailSender in a test. It will cost lots of money $$.

```

En nuestro caso, podemos empezar usando el mismo espía que usamos en el otro test, ya que este segundo test se basa igualmente en la cuenta de mensajes:

```typescript
it('should send greeting emails to all customers with birthday today', () => {
  const emailSender = new MessageCountingEmailSender()
  const service = new BirthdayService(
          new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
          ]),
          emailSender,
          new ProductionLogger(),
  )

  service.greetCustomersWithBirthday(new Date())

  expect(emailSender.countOfSentEmails()).toBe(2)
})

```

### El colaborador que no hacía nada

Ahora, al ejecutar el test, cambia el error:

```plain text
Error: ️😱 You are using ProductionLogger in a test. It will increase our bills by zillions $$.

```

Necesitamos suplantar `ProductionLogger` con un doble. Nosotras usaremos un _Dummy_ porque, si bien necesitamos una instancia de `ProductionLogger` para poder instanciar el propio `BirthdayService`, no tiene influencia en su comportamiento.

Si estudiamos la implementación actual de ProductionLogger podemos ver que es una clase concreta, por lo que lo mejor sería introducir una abstracción, invertir la dependencia y así tener vía libre para crear un doble de test.

```typescript
export class ProductionLogger {
  log(level: string, message: string) {
    throw new Error(
      '️😱 You are using ProductionLogger in a test. It will increase our bills by zillions $$.',
    )
  }
}

```

Se podría argumentar que una alternativa es extender la clase `ProductionLogger` y sobreescribir sus métodos para introducir un doble de test. Esto funcionaría, pero es una de esas soluciones de “pan para hoy, hambre para mañana”. No es descabellado pensar que podríamos tener que cambiar nuestra librería de logs en algún momento y disponer de una abstracción hace que ese cambio sea trivial.

Por tanto, tal como hemos hecho con otros colaboradores, introducimos una interfaz `Logger` a partir de `ProductionLogger` y definimos un `DummyLogger`, que simplemente no hará nada.

```typescript
class DummyLogger implements Logger {
  log(level: string, message: string): void {
  }
}

```

El test queda así y con estos cambios ya pasa:

```typescript
it('should send greeting emails to all customers with birthday today', () => {
  const emailSender = new MessageCountingEmailSender()
  const service = new BirthdayService(
          new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
          ]),
          emailSender,
          new DummyLogger(),
  )

  service.greetCustomersWithBirthday(new Date())

  expect(emailSender.countOfSentEmails()).toBe(2)
})

```

Podemos ver que el test falla si esperamos valores de `countOfSentEmails` distintos de 2, lo que nos indica que el test está verificando lo que queríamos.

### Verificando que se envía el mensaje correcto

Con esto, tenemos resuelta la primera parte del ejercicio. Sin embargo, todavía nos queda un buen trabajo por delante. Por ejemplo, no tenemos tests que verifiquen que se construye el mensaje correcto y que se envía a la persona correcta. Si estudiamos el código de `BirthdayService` podemos observar un par de cosas:

```typescript
greetCustomersWithBirthday(today: Date) {
  const customers = this.customerRepository.findWithBirthday(today)
  customers.forEach((customer) => {
    const discountCode = new DiscountCodeGenerator().generate()
    const template =
            'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
            )
    customer.sendEmail(template, this.emailSender)
    this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
  })
}

```

La primera es que para hacer este test necesitamos seguir usando `BirthdayService` como unidad bajo test. Además, vemos que el generador de códigos de descuento también nos va a generar alguna dificultad. Esencialmente, necesitamos poner un espía en lugar del `EmailSender` que se encargue de recopilar los detalles de los correos enviados para ver si están correctamente formados.

En esencia, lo que queremos es tener este test:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
  const emailSender = new MessageContentSpyEmailSender()
  const service = new BirthdayService(
          new CustomersWithBirthdayToday([
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
          ]),
          emailSender,
          new DummyLogger(),
  )

  service.greetCustomersWithBirthday(new Date())

  expect(emailSender.lastMessageContent()).toBe('Happy birthday, Jane Doe! Here is your discount code: {discount}')
  expect(emailSender.lastMessageRecipient()).toBe('jane@example.com')
})

```

Como se puede ver, vamos a introducir un nuevo espía. Podrías estar preguntándote si no sería más sencillo modificar el espía existente y añadirle más métodos y capacidades. Sin embargo, eso no haría más que introducir un acoplamiento alto entre tests. En los tests anteriores nos preocupaba contar el número de mensajes enviados, que es algo distinto a examinar su contenido. Por tanto, escribamos un nuevo doble, que no cuesta tanto y nos dejará todo limpio y ordenado:

```typescript
class MessageContentSpyEmailSender implements EmailSender{
  private lastMessageRecipientValue: string
  private lastMessageContentValue: string

  send(email: string, message: string): void {
    this.lastMessageRecipientValue = email
    this.lastMessageContentValue = message
  }

  lastMessageContent(): string {
    return this.lastMessageContentValue
  }

  lastMessageRecipient(): string {
    return this.lastMessageRecipientValue
  }
}

```

Si ejecutamos el test ahora, nos devolverá el siguiente mensaje:

```plain text
AssertionError: expected 'Happy birthday, Jane Doe! Here is you…' to be 'Happy birthday, Jane Doe! Here is you…' // Object.is equality
Expected :Happy birthday, Jane Doe! Here is your discount code: {discount}
Actual   :Happy birthday, Jane Doe! Here is your discount code: FP1BRI

```

Hasta ahora no sabíamos el código de descuento, por lo que vamos a sustituirlo en el test por el que se acaba de generar. Seguramente ya te has dado cuenta de lo que va a pasar:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
  const emailSender = new MessageContentSpyEmailSender()
  const service = new BirthdayService(
          new CustomersWithBirthdayToday([
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
          ]),
          emailSender,
          new DummyLogger(),
  )

  service.greetCustomersWithBirthday(new Date())

  expect(emailSender.lastMessageContent()).toBe(
          'Happy birthday, Jane Doe! Here is your discount code: FP1BRI',
  )
  expect(emailSender.lastMessageRecipient()).toBe('jane@example.com')
})

```

Efectivamente, el test va a seguir fallando porque el código de descuento se calcula de forma aleatoria cada vez. En estas condiciones, el test no va a pasar nunca porque el resultado es no determinista.

```plain text
AssertionError: expected 'Happy birthday, Jane Doe! Here is you…' to be 'Happy birthday, Jane Doe! Here is you…' // Object.is equality
Expected :Happy birthday, Jane Doe! Here is your discount code: FP1BRI
Actual   :Happy birthday, Jane Doe! Here is your discount code: DDRJ2L

```

El problema es que `DiscountCodeGenerator` es el colaborador “no determinista” y está instanciado dentro de `BirthdayService`, no se inyecta y no hay forma, aparentemente, de librarse de esa limitación y poder escribir un test en condiciones.

```typescript
greetCustomersWithBirthday(today: Date) {
  const customers = this.customerRepository.findWithBirthday(today)
  customers.forEach((customer) => {
    const discountCode = new DiscountCodeGenerator().generate()
    const template =
            'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
            )
    customer.sendEmail(template, this.emailSender)
    this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
  })
}

```

En realidad, tenemos varias alternativas:

**Escribir un test basado en propiedades**: en lugar de testear por el mensaje exacto usar una expresión regular para describir que el mensaje contiene una parte fija de texto y otra que es una combinación aleatoria de los caracteres A-Z y 0-9. Por ejemplo:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
  const emailSender = new MessageContentSpyEmailSender()
  const service = new BirthdayService(
          new CustomersWithBirthdayToday([
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
          ]),
          emailSender,
          new DummyLogger(),
  )

  service.greetCustomersWithBirthday(new Date())

  expect(emailSender.lastMessageContent()).toMatch(
          /Happy birthday, Jane Doe! Here is your discount code: [A-z0-9]{6}/,
  )
  expect(emailSender.lastMessageRecipient()).toBe('jane@example.com')
})

```

Esta solución nos permite avanzar sin tocar más el código. El problema es que _tapa_ un defecto de diseño como es el de tener una dependencia _dura_, como es la de un colaborador que devuelve resultados no deterministas embebido en el código, con lo que estamos completamente acopladas.

En todo caso, nada nos impide proceder a un refactor orientado a deshacer este acoplamiento, extraer, invertir e inyectar la dependencia, de modo que en el futuro no tengamos muchas dificultades para cambiarla o para testearla. Pero esto vamos a dejarlo para una segunda parte, ya que el artículo tiene bastante contenido.

## Conclusiones

Los dobles de test son herramientas que nos ayudan a poner nuestro código a prueba en un entorno controlado y aislándonos de detalles tecnológicos que pueden influir de forma significativa en los aspectos funcionales y no funcionales de nuestro código.

Por desgracia, tenemos un problema de nomenclatura, ya que es común referirnos a todos los dobles de test con el término Mock, que denomina un tipo específico que, además, no es siquiera el más recomendable para usar.

En el artículo hemos mencionado una serie de principios y patrones de diseño que deben estar presentes en nuestro código para beneficiarnos del uso de dobles de test. También hemos descrito los diversos tipos y en qué casos de uso nos conviene utilizarlos.

El artículo se completa con un ejercicio práctico en el que es explora como introducir diferentes tipos de dobles de test, incluso cuando el código no aplica correctamente los principios y patrones de diseño mencionados.

En la siguiente entrega veremos como extraer una dependencia acoplada para poder doblarla en situación de test. También exploraremos la introducción de Stubs que fallan, lo que nos permitirá practicar la creación de test enfocados a verificar que nuestros objetos saben lidiar con los errores de sus colaboradores.

Publicado: March 2, 2025

[https://franiglesias.github.io/test-doubles-guide-1/](https://franiglesias.github.io/test-doubles-guide-1/)



# Parte 2

Vamos por la segunda parte de este artículo en la que seguimos explicando como usar dobles de test.

En esta ocasión vamos a analizar un caso un poco más complejo de dependencia totalmente acoplada y luego nos centraremos en los dobles que simulan errores.

Aquí tienes el índice del artículo, por si prefieres saltar directamente a alguno de los puntos:

- [El caso de la dependencia acoplada](https://franiglesias.github.io/test-doubles-guide-2/#el-caso-de-la-dependencia-acoplada) 
- [Seams: una solución temporal](https://franiglesias.github.io/test-doubles-guide-2/#seams-una-soluci%C3%B3n-temporal)
- [Inversión de dependencias al rescate](https://franiglesias.github.io/test-doubles-guide-2/#inversi%C3%B3n-de-dependencias-al-rescate)
- [Dobles de test que simulan fallos](https://franiglesias.github.io/test-doubles-guide-2/#dobles-de-test-que-simulan-fallos) 
- [Base de datos innacesible](https://franiglesias.github.io/test-doubles-guide-2/#base-de-datos-innacesible)
- [Lógica justificada en un doble de test](https://franiglesias.github.io/test-doubles-guide-2/#l%C3%B3gica-justificada-en-un-doble-de-test) 
- [Valores distintos en cada llamada](https://franiglesias.github.io/test-doubles-guide-2/#valores-distintos-en-cada-llamada)
- [Stubs que fallan unas veces y otras no](https://franiglesias.github.io/test-doubles-guide-2/#stubs-que-fallan-unas-veces-y-otras-no)
- [Más sobre reintentos](https://franiglesias.github.io/test-doubles-guide-2/#m%C3%A1s-sobre-reintentos)
- [Conclusiones](https://franiglesias.github.io/test-doubles-guide-2/#conclusiones)

Puedes encontrar el [código en este repositorio: birthday-service-kata](https://github.com/franiglesias/birthday-service-kata).

## El caso de la dependencia acoplada

El grado máximo de acoplamiento de un objeto con otro ocurre cuando uno de los objetos tiene que saber instanciar otro, conocimiento que debe añadirse al de saber como usarlo. Es el caso que tenemos aquí, `BirthdayService` está completamente acoplado a `DiscountCodeGenerator`.

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
    }

    greetCustomersWithBirthday(today: Date) {
        const customers = this.customerRepository.findWithBirthday(today)
        customers.forEach((customer) => {
            const discountCode = new DiscountCodeGenerator().generate()
            const template =
                'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
                )
            customer.sendEmail(template, this.emailSender)
            this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
        })
    }
}

```

`DiscountCodeGenerator` produce un resultado no determinista. Es decir, no podemos saber con antelación qué nos va a devolver. Eso incrementa la dificultad para hacer un test. Por el momento, la hemos superado usando un test basado en propiedades en el que, en lugar de esperar un resultado concreto, esperamos un resultado que cumple unas características que hemos podido describir con una expresión regular:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
    const emailSender = new MessageContentSpyEmailSender()
    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
        ]),
        emailSender,
        new DummyLogger(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.lastMessageContent()).toMatch(
        /Happy birthday, Jane Doe! Here is your discount code: [A-z0-9]{6}/,
    )
    expect(emailSender.lastMessageRecipient()).toBe('jane@example.com')
})

```

Con esto, hacemos pasar el test y no tenemos necesidad de nada más. Podríamos decir que el test cubre a `DiscountCodeGenerator` y ya sería suficiente.

Pero este enfoque me parece equivocado por una razón que voy a tratar de explicar, y que se traduce en que tenemos que hacer el test así porque `BirthdayService` está acoplado a este `DiscountCodeGenerator` específico y, por tanto, el test también lo está.

Es más, si `DiscountCodeGenerator` cambiase por el motivo que sea y el código generado tuviese una forma distinta el test fallaría por una razón totalmente ajena a la naturaleza del test. Es cierto que `BirthdayService` es responsable de componer el mensaje, pero claramente no es responsable de generar el código. De hecho, su responsabilidad con respecto al código es únicamente incluirlo en la notificación, tenga la forma que tenga.

En consecuencia, este test es frágil. Va a romper por una razón que no tiene que ver con su intención y está poniendo bajo test algo que no es de su incumbencia.

Pero, además, en el ejercicio no se aprecian otros problemas de una dependencia dura: consumo de tiempo y recursos, complejidad de instanciación, side-effects que pueda acarrear, etc. Imagina el efecto en el test si la ejecución de este componente consumiese varios segundos.

### Seams: una solución temporal

Un abordaje habitual de este tipo de test es introducir un _Seam_ para controlar el comportamiento de la dependencia. Nuestro objetivo es eliminar del test el efecto de la dependencia real (`DiscountCodeGenerator`). Puede ser una buena aproximación inicial para entender mejor qué es lo que pasa en el código y asegurarnos que tenemos controlada la dependencia.

Un _Seam_ es un punto en el que podemos hacer cambios en el código sin afectar al resto de la unidad. En la práctica de nuestro ejercicio se traduce en aislar el uso de la dependencia en un método que luego podamos sobreescribir.

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
    }

    greetCustomersWithBirthday(today: Date) {
        const customers = this.customerRepository.findWithBirthday(today)
        customers.forEach((customer) => {
            const discountCode = this.getDiscountCode()
            const template =
                'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
                )
            customer.sendEmail(template, this.emailSender)
            this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
        })
    }

    protected getDiscountCode(): DiscountCode {
        return new DiscountCodeGenerator().generate()
    }
}

```

En este caso no hay más usos de `DiscountCodeGenerator` por lo que podemos pasar al siguiente paso. Extendemos `BirthdayService` creando una clase derivada `TestableBirthdayService` en la cual sobreescribimos el método `getDiscountCode` de forma que devuelva un valor conocido. Es como hacer un _Stub_.

```typescript
class TestableBirthdayService extends BirthdayService {

    protected getDiscountCode(): DiscountCode {
        return new DiscountCode('FP1BRI');
    }
}

```

Y ahora podemos reemplazarlo en el test y cambiar la aserción. De esta forma, el test me está diciendo que `BirthdayService` inserta el código generado por `getDiscountCode` en la notificación, sin importarle como lo genera.

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
    const emailSender = new MessageContentSpyEmailSender()
    const service = new TestableBirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
        ]),
        emailSender,
        new DummyLogger(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.lastMessageContent()).toBe(
        'Happy birthday, Jane Doe! Here is your discount code: FP1BRI',
    )
    expect(emailSender.lastMessageRecipient()).toBe('jane@example.com')
})

```

### Inversión de dependencias al rescate

Por supuesto, el objetivo debería llegar a ser invertir, extraer e inyectar, la dependencia. En nuestro ejemplo esto supone un cierto coste y riesgo porque hay que alterar la signatura del constructor y podríamos tener varios usos del mismo. Afortunadamente, disponemos de varias posibilidades, algunas dependiendo del lenguaje: sobrecarga de constructores, parámetros opcionales, etc.

En el lado positivo, hay que decir que haber optado por empezar con un Seam puede ser una buena idea. Es una forma de unificar los usos de la dependencia de forma inocua, lo que facilita el trabajo a continuación. En realidad, los pasos son bastante sencillos: se trata de promover la dependencia para convertirla en un colaborador e inicializarla en el constructor:

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: DiscountCodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = new DiscountCodeGenerator()
    }

    greetCustomersWithBirthday(today: Date): void {
        const customers = this.customerRepository.findWithBirthday(today)
        customers.forEach((customer) => {
            const discountCode = this.getDiscountCode()
            const template =
                'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
                )
            customer.sendEmail(template, this.emailSender)
            this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
        })
    }

    protected getDiscountCode(): DiscountCode {
        return this.discountCodeGenerator.generate()
    }
}

```

Esto no cambia nada en el comportamiento, pero nos pone en una mejor disposición. Es el momento de introducir una interfaz, como abstracción del rol de _generador de códigos_.

```typescript
export interface CodeGenerator {
    generate(): DiscountCode
}

```

Y también de usarla:

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: CodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = new DiscountCodeGenerator()
    }

    // Removed for clarity
}

```

Ya estamos cerca de inyectarla. En Typescript nos bastaría con hacerlo así:

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: CodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
        discountCodeGenerator: CodeGenerator = undefined
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = discountCodeGenerator ?? new DiscountCodeGenerator()
    }

    // Removed for clarity
}

```

O también así:

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: CodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
        discountCodeGenerator: CodeGenerator = new DiscountCodeGenerator()
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = discountCodeGenerator
    }

    // Removed for clarity
}

```

En ambos casos, mantenemos la compatibilidad con el código existente (los tests siguen pasando) y añadimos la posibilidad de inyectar la dependencia. En el caso del test, podemos introducir un doble:

```typescript
class CodeGeneratorStub implements CodeGenerator {
    generate(): DiscountCode {
        return new DiscountCode('FP1BRI')
    }
}

```

Y utilizarlo en lugar del servicio real en el test:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
    const emailSender = new MessageContentSpyEmailSender()
    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
        ]),
        emailSender,
        new DummyLogger(),
        new CodeGeneratorStub(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.lastMessageContent()).toBe(
        'Happy birthday, Jane Doe! Here is your discount code: FP1BRI',
    )
    expect(emailSender.lastMessageRecipient()).toBe('jane@example.com')
})

```

Y con esto, hemos inver†ido la dependencia y desacoplado `BirthdayService` de sus colaboradores, abriendo la puerta a unos tests mucho más sólidos.

## Dobles de test que simulan fallos

En los tests no tenemos cubierta la posibilidad de que nuestro servicio falle por problemas en los colaboradores. Por ejemplo, BirthdayService será sensible a que falle la conexión de `Customers` a la base de datos física, pero también que falle `EmailSender` o, en su caso, `CodeGenerator`. De hecho, esa posibilidad de fallo es uno de los motivos por los que tenemos que introducir dobles de test: por la posibilidad no reproducible de que el colaborador falle y por la dificultad de provocar un fallo del colaborador.

Pero con un doble de test es fácil conseguir ambas cosas: un colaborador que no falle o un colaborador que falle de una forma determinada, que es de lo que nos vamos a ocupar a continuación.

Como no tenemos ninguna gestión de errores en nuestro código de ejemplo podemos suponer que vamos a trabajar en TDD para añadirla.

### Base de datos innacesible

Supongamos que hemos averiguado que necesitamos gestionar un error de tipo `DatabaseUnavailable` cuando, por el motivo que sea, falla nuestra implementación en producción de `Customers`. Para simplificar, vamos a suponer que el comportamiento de `BirthdayService` en esa situación debe ser hacer un log informando del fallo y terminar con normalidad.

Como es de esperar, necesitaremos un _Stub_ de `Customers` que tire un error `DatabaseUnavailable`, y un espía que nos diga qué mensaje se ha puesto en el log.

```typescript
it('should log and stop when database is unavailable', () => {
    const emailSender = new MessageCountingEmailSender()
    const logger = new LoggerMessageSpy()
    const service = new BirthdayService(
        new UnavailableCustomers(),
        emailSender,
        logger,
        new CodeGeneratorStub(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(logger.lastMessage()).toBe('Customer database is not available')
    expect(logger.lastLevel()).toBe('ERROR')
    expect(emailSender.countOfSentEmails()).toBe(0)
})

```

He aquí el Stub que falla de Customers. Es una implementación que lo único que hace es fallar de una manera determinada:

```typescript
class UnavailableCustomers implements Customers {
    findWithBirthday(today: Date): Customer[] {
        throw new DatabaseUnavailable('Customer database is not available')
    }
}

```

Específicamente, con este error:

```typescript
class DatabaseUnavailable implements Error {
  name: string
  message: string

  constructor(message: string) {
    this.message = message
  }
}

```

Por su parte, para espiar el mensaje que recibirá el logger tenemos:

```typescript
class LoggerMessageSpy implements Logger{
    private message: string
    private level: string

    log(level: string, message: string): void {
        this.message = message
        this.level = level
    }

    lastMessage() {
        return this.message
    }

    lastLevel() {
        return this.level
    }
}

```

Como es de esperar el test fallará, puesto que todavía no hemos implementado nada en el servicio.

```plain text
Customer database is not available

```

El código que necesitamos sirve para capturar el error y reaccionar de la forma que queremos:

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: CodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
        discountCodeGenerator: CodeGenerator = new DiscountCodeGenerator()
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = discountCodeGenerator
    }

    greetCustomersWithBirthday(today: Date): void {
        let customers: Customer[]
        try {
            customers = this.customerRepository.findWithBirthday(today)
        } catch (e) {
            this.logger.log('ERROR', e.message)
            return
        }

        customers.forEach((customer) => {
            const discountCode = this.getDiscountCode()
            const template =
                'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
                )
            customer.sendEmail(template, this.emailSender)
            this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
        })
    }

    protected getDiscountCode(): DiscountCode {
        return this.discountCodeGenerator.generate()
    }
}

```

Por supuesto, podemos hacerlo de una forma un poco más limpia:

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: CodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
        discountCodeGenerator: CodeGenerator = new DiscountCodeGenerator(),
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = discountCodeGenerator
    }

    greetCustomersWithBirthday(today: Date): void {
        const customers: Customer[] = this.getCustomersWithBirthday(today)
        customers.forEach((customer) => {
            const discountCode = this.getDiscountCode()
            const template =
                'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
                )
            customer.sendEmail(template, this.emailSender)
            this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
        })
    }

    private getCustomersWithBirthday(today: Date): Customer[] {
        let customers: Customer[] = []
        try {
            customers = this.customerRepository.findWithBirthday(today)
        } catch (e) {
            this.logger.log('ERROR', e.message)
        }
        return customers
    }

    protected getDiscountCode(): DiscountCode {
        return this.discountCodeGenerator.generate()
    }
}

```

## Lógica justificada en un doble de test

Hasta ahora hemos escrito dobles de test con una lógica muy sencilla. Lo más complicado ha sido ir tomando nota de los parámetros recibidos o ir contando las invocaciones a un método. Sin embargo, hay ocasiones en las que necesitamos simular un comportamiento un poco más elaborado. Algunos ejemplos:

- _Stubs_ que pueden ser llamados repetidas veces y queremos, o necesitamos, que cada vez nos devuelvan algo distinto.
- _Stubs_ que fallan unas veces y otras funcionan bien.
- Cuando nos interesa testear que la unidad bajo test reintenta una operación.

En principio, las librerías de dobles podrían facilitarnos esto, pero en realidad no es tan complicado programarlo para las necesidades de nuestro test.

Ahora bien, a veces es simplemente cuestión de pensarlo detenidamente, por lo que en cada ejemplo a continuación voy a intentar explicar por qué es muy probable que no necesites ese esfuerzo.

### Valores distintos en cada llamada

Primero, analicemos este test:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
    const emailSender = new MessageContentSpyEmailSender()
    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
        ]),
        emailSender,
        new DummyLogger(),
        new CodeGeneratorStub(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.lastMessageContent()).toBe(
        'Happy birthday, Jane Doe! Here is your discount code: FP1BRI',
    )
    expect(emailSender.lastMessageRecipient()).toBe('jane@example.com')
})

```

Una objeción que me han puesto a veces es que deberíamos hacer un test con varios ejemplos de `Customer`, para verificar que efectivamente se itera la lista y que se envía el mensaje correcto a los distintos clientes. Si te digo la verdad, no tengo clara la necesidad de hacer este test, pero tiene sentido. ¿De cuántos clientes debería ser la muestra? Puesto a ejercer de abogado del diablo, nada me garantiza que al llegar a un cierto número de clientes, el bucle deje de iterar o algo.

En fin, bromas aparte, supongamos que con dos clientes ya tenemos suficiente. Un posible abordaje sería introducir un nuevo espía que en lugar del último mensaje, sea capaz de decirme el mensaje enviado en una iteración determinada.

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
    const emailSender = new MessageContentByIterationEmailSender()
    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
        ]),
        emailSender,
        new DummyLogger(),
        new CodeGeneratorStub(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.messageContent(0)).toBe(
        'Happy birthday, John Doe! Here is your discount code: FP1BRI',
    )
    expect(emailSender.messageContent(1)).toBe(
        'Happy birthday, Jane Doe! Here is your discount code: FP1BRI',
    )
})

```

Es fácil ver que hay un problema con este espía, ya que está acoplado al orden en que introducimos los clientes. Se me ocurre una forma mejor: pedirle el mensaje que hemos enviado a cada dirección de email. De este modo, da igual el orden en que hayamos metido los datos e incluso da igual la cantidad de clientes que usemos en el test. Es más, con esto verificamos no solo el contenido sino que lo enviamos a la dirección correcta:

```typescript
class MessageContentByAddressEmailSender implements EmailSender {
    private messages: Record<string, string> = {}
    send(email: string, message: string): void {
        this.messages[email] = message
    }
    messageContent(email: string): string {
        return this.messages[email]
    }
}

```

El test es el siguiente:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
    const emailSender = new MessageContentByAddressEmailSender()
    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
        ]),
        emailSender,
        new DummyLogger(),
        new CodeGeneratorStub(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.messageContent('john@example.com')).toBe(
        'Happy birthday, John Doe! Here is your discount code: FP1BRI',
    )
    expect(emailSender.messageContent('jane@example.com')).toBe(
        'Happy birthday, Jane Doe! Here is your discount code: FP1BRI',
    )
})

```

Puedes añadir más clientes, o cambiar su orden, y el test seguirá funcionando perfectamente. Sin embargo, casi puedo escuchar la pregunta: Pero, ¿qué pasa con el código de descuento? ¿No debería ser distinto cada vez?

Bien, pues no tenemos más que introducir un nuevo _Stub_ que pueda darnos diferentes códigos de descuento que tengamos controlados. Este que mostramos nos devolverá cada código que le hayamos pasado en el constructor en el mismo orden.

```typescript
class MultipleCodeGeneratorStub implements CodeGenerator {
    private codes: string[]

    constructor(...codes: string[]) {
        this.codes = codes
    }

    generate(): DiscountCode {
        return new DiscountCode(this.codes.shift())
    }
}

```

Lo que nos hace posible escribir este test:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
    const emailSender = new MessageContentByAddressEmailSender()
    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
        ]),
        emailSender,
        new DummyLogger(),
        new MultipleCodeGeneratorStub('FP1BRI', 'FP20ZI'),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.messageContent('john@example.com')).toBe(
        'Happy birthday, John Doe! Here is your discount code: FP1BRI',
    )
    expect(emailSender.messageContent('jane@example.com')).toBe(
        'Happy birthday, Jane Doe! Here is your discount code: FP20ZI',
    )
})

```

Ahora bien, tenemos algunos problemas. Como no hay relación entre el código de descuento y el cliente, ahora el test estará acoplado al orden en que pasamos los códigos de ejemplo, así que perdemos la capacidad de ordenar los clientes de otra forma. Y, por otro lado, el doble de test funcionará mientras haya códigos, por lo que el número de clientes de ejemplo debe ser igual al número de códigos.

Puedes pensar que estos problemas no son especialmente importantes o que puedes convivir con ellos.

Sin embargo, te planteo de nuevo la reflexión que hicimos más arriba. En este test no estamos verificando que se genera un código correcto, solo verificamos que se incluye el código obtenido en el mensaje que enviamos al cliente. Nos da igual el código mientras se pueda expresar como un `string` e incluir en la plantilla del mensaje. En consecuencia, para este test realmente no nos importa que el doble de `CodeGenerator` devuelva siempre el mismo código. Este test, por tanto, sería suficiente, y además es mucho más flexible:

```typescript
it('should send a well formed message with the discount code to the right customer', () => {
    const emailSender = new MessageContentByAddressEmailSender()
    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
            new Customer('Jane Doe', 'jane@example.com', new Date('2005-02-14')),
        ]),
        emailSender,
        new DummyLogger(),
        new CodeGeneratorStub(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.messageContent('john@example.com')).toBe(
        'Happy birthday, John Doe! Here is your discount code: FP1BRI',
    )
    expect(emailSender.messageContent('jane@example.com')).toBe(
        'Happy birthday, Jane Doe! Here is your discount code: FP1BRI',
    )
})

```

En todo caso, tendrías que testear específicamente `DiscountCodeGenerator` para verificar que siempre saca un código distinto.

### Stubs que fallan unas veces y otras no

Este caso de uso es un poco más _tricky_, pero aceptemos que nos interesa comprobar que nuestra unidad bajo test es capaz de hacer reintentos y lograr sus objetivos si tiene la oportunidad. O dicho de otra forma, que si su colaborador falla en una primera ocasión, es capaz de probar de nuevo y conseguir su objetivo.

Nuestro ejemplo de enviar un email cuadra muy bien con nuestro objetivo, ya que no es descabellado pensar que un sistema de envío de emails al enviar muchos mensajes seguidos pueda dejar de aceptar nuevos intentos durante unos segundos. En cualquier caso, es un mecanismo habitual.

Así que vamos a necesitar un doble de `EmailSender` que, por ejemplo, falle la primera vez que se llame, pero que funcione al segundo intento. De este modo vamos a poder probar que `BirthdayService` sabe reaccionar a esa situación. Ahora bien, no quiero mezclar una cosa con otra, así que voy a hacerlo mediante un decorador que será el que falle y delegará en otro doble la cuenta de mensajes.

En construcción se pueden especificar las veces que queremos que falle el doble antes de simular que permite el envío:

```typescript
class FailingEmailSender implements EmailSender {
    private emailSender: EmailSender
    private timesFailing: number
    constructor(emailSender: EmailSender, timesFailing: number) {
        this.emailSender = emailSender
        this.timesFailing = timesFailing
    }

    send(email: string, message: string): void {
        if (this.timesFailing > 0) {
            this.timesFailing--
            throw new Error('Email sending failed')
        }
        this.emailSender.send(email, message)
    }
}

```

Y este será el test, en el que hacemos que el `EmailSender` falle una vez antes de permitir enviar emails.

```typescript
it('should retry to send email if EmailSender fails', () => {
    const emailSender = new MessageCountingEmailSender()
    const timesToFailBeforeAcceptingMessage = 1
    const failingEmailSender = new FailingEmailSender(emailSender, timesToFailBeforeAcceptingMessage)
    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
        ]),
        failingEmailSender,
        new DummyLogger(),
    )

    service.greetCustomersWithBirthday(new Date())

    expect(emailSender.countOfSentEmails()).toBe(1)
})

```

Y aquí tenemos un código que hace pasar el test. Por cierto, que si quieres verificar que BirthdayService lo va a intentar tres veces antes de darse por vencido puedes cambiar el valor de `timesToFailBeforeAcceptingMessage`. Spoiler: el test no pasará si pones que `EmailSender` falle tres veces, que es el número de reintentos que hemos decidido permitir.

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: CodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
        discountCodeGenerator: CodeGenerator = new DiscountCodeGenerator(),
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = discountCodeGenerator
    }

    greetCustomersWithBirthday(today: Date): void {
        const customers: Customer[] = this.getCustomersWithBirthday(today)
        customers.forEach((customer) => {
            const discountCode = this.getDiscountCode()
            const template =
                'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
                )
            let retries = 3
            while (retries > 0) {
                try {
                    customer.sendEmail(template, this.emailSender)
                    break
                } catch (e) {
                    this.logger.log('ERROR', e.message)
                    retries--
                }
            }
            this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
        })
    }

    private getCustomersWithBirthday(today: Date): Customer[] {
        let customers: Customer[] = []
        try {
            customers = this.customerRepository.findWithBirthday(today)
        } catch (e) {
            this.logger.log('ERROR', e.message)
        }
        return customers
    }

    protected getDiscountCode(): DiscountCode {
        return this.discountCodeGenerator.generate()
    }
}

```

De nuevo, nos conviene hacer un poco de refactor para dejar el método público más legible.

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: CodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
        discountCodeGenerator: CodeGenerator = new DiscountCodeGenerator(),
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = discountCodeGenerator
    }

    greetCustomersWithBirthday(today: Date): void {
        const customers: Customer[] = this.getCustomersWithBirthday(today)
        customers.forEach((customer) => {
            const discountCode = this.getDiscountCode()
            const template =
                'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
                )
            this.sendMessage(customer, template)
            this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
        })
    }

    private getCustomersWithBirthday(today: Date): Customer[] {
        let customers: Customer[] = []
        try {
            customers = this.customerRepository.findWithBirthday(today)
        } catch (e) {
            this.logger.log('ERROR', e.message)
        }
        return customers
    }

    protected getDiscountCode(): DiscountCode {
        return this.discountCodeGenerator.generate()
    }

    private sendMessage(customer: Customer, template: string): void {
        let retries = 3
        while (retries > 0) {
            try {
                customer.sendEmail(template, this.emailSender)
                break
            } catch (e) {
                this.logger.log('ERROR', e.message)
                retries--
            }
        }
    }
}

```

Seguramente podríamos haber hecho un _Spy_ con toda la funcionalidad, pero me parece que el decorador nos ha permitido no solo aprovechar el espía existente, sino ser un poco más rigurosas con la separación de responsabilidades.

### Más sobre reintentos

¿Cómo probamos que `BirthdayService` es capaz de dejar de enviar un mensaje si el servicio falla en todos los intentos? A primera vista podría tener sentido usar un test similar al anterior, junto con un espía que nos diga el mensaje que hemos puesto en el `Logger`. No lo he mencionado explícitamente, pero nuestro servicio procura no romperse si no puede mandar un email y registra el fallo para que podamos hacer algo posteriormente.

Si nos fijamos en el código veremos que se hace un log con cada intento de envío fallido de un mensaje. Eso supone que si un mensaje no se puede enviar aparecerá logado tres veces. Necesitamos un espía de logs que pueda contarlo, tal vez incluso teniendo en cuenta el email al que se ha tratado de enviar el mensaje.

```typescript
class CountingLogger implements Logger {
    private entries: string[] = []

    log(level: string, message: string): void {
        this.entries.push(message)
    }

    countOfEntriesFor(address: string): number {
        const entriesForAddress = this.entries.filter((entry) =>
            entry.includes(address),
        )
        return entriesForAddress.length
    }
}

```

Por otro lado, si el peso de verificar los reintentos recae en el log, resulta que `EmailSender` puede ser mucho más simple. Sencillamente: puede fallar siempre.

```typescript
class AlwaysFailingEmailSender implements EmailSender {
    send(email: string, message: string): void {
        throw new Error('Email sending failed')
    }
}

```

El test queda así:

```typescript
it('should retry to 3 times before failing to send', () => {
    const emailSender = new AlwaysFailingEmailSender()
    const logger = new CountingLogger()

    const service = new BirthdayService(
        new CustomersWithBirthdayToday([
            new Customer('John Doe', 'john@example.com', new Date('1990-02-14')),
        ]),
        emailSender,
        logger,
    )

    service.greetCustomersWithBirthday(new Date())

    expect(logger.countOfEntriesFor('john@example.com')).toBe(3)
})

```

Pero si lo ejecutamos falla. Nos salen cuatro entradas, ¿qué está pasando? Pues que si nos fijamos en este fragmento, estamos _logando_ como enviados todos los mensajes, incluso los fallidos:

```typescript
greetCustomersWithBirthday(today: Date): void {
    const customers: Customer[] = this.getCustomersWithBirthday(today)
    customers.forEach((customer) => {
        const discountCode = this.getDiscountCode()
        const template =
            'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                '{discount}',
                discountCode.getCode(),
            )
        this.sendMessage(customer, template)
        this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
    })
}

```

Es cuestión de cambiar la línea del logger a un lugar más apropiado. Así queda finalmente `BirthdayService`.

```typescript
export class BirthdayService {
    private readonly customerRepository: Customers
    private readonly emailSender: ProductionEmailSender
    private readonly logger: Logger
    private readonly discountCodeGenerator: CodeGenerator

    constructor(
        customerRepository: Customers,
        emailSender: ProductionEmailSender,
        logger: Logger,
        discountCodeGenerator: CodeGenerator = new DiscountCodeGenerator(),
    ) {
        this.customerRepository = customerRepository
        this.emailSender = emailSender
        this.logger = logger
        this.discountCodeGenerator = discountCodeGenerator
    }

    greetCustomersWithBirthday(today: Date): void {
        const customers: Customer[] = this.getCustomersWithBirthday(today)
        customers.forEach((customer) => {
            const discountCode = this.getDiscountCode()
            const template =
                'Happy birthday, {name}! Here is your discount code: {discount}'.replace(
                    '{discount}',
                    discountCode.getCode(),
                )
            this.sendMessage(customer, template)
        })
    }

    private getCustomersWithBirthday(today: Date): Customer[] {
        let customers: Customer[] = []
        try {
            customers = this.customerRepository.findWithBirthday(today)
        } catch (e) {
            this.logger.log('ERROR', e.message)
        }
        return customers
    }

    protected getDiscountCode(): DiscountCode {
        return this.discountCodeGenerator.generate()
    }

    private sendMessage(customer: Customer, template: string): void {
        let retries = 3
        while (retries > 0) {
            try {
                customer.sendEmail(template, this.emailSender)
                this.logger.log('INFO', customer.fillWithEmail('Email sent to {email}'))
                break
            } catch (e) {
                this.logger.log(
                    'ERROR',
                    customer.fillWithEmail(`${e.message} when sending to {email}`),
                )
                retries--
            }
        }
    }
}

```

## Conclusiones

Seguramente se podrían analizar más casos de tests en los que podríamos usar distintos tipos de dobles, pero creo que tenemos una muestra bastante representativa en este artículo. El ejercicio en sí buscaba forzar algunas situaciones típicas de testing, que pueden darse tanto con el test a posteriori como en situaciones de TDD.

Si hago recuento de todos los dobles de test que he introducido en el ejercicio me salen 12. Puede que esta cifra te parezca exagerada, pero si lo examinas con calma, verás que cada uno de ellos es extremadamente sencillo y enfocado en su tarea. Además, hemos minimizado el acoplamiento entre tests.

Por otro lado, para introducir dobles hemos tenido que aplicar varios principios de diseño. En especial, el de Inversión de Dependencias. También nos hemos beneficiado, aunque no lo he mencionado explícitamente, de la Segregación de Interfaces. En este ejemplo, todas las interfaces son muy sencillas, con un único método, lo que abarata enormemente la creación de dobles de test.

En ningún momento hemos usado librerías de mocks. De hecho, no hemos usado ningún _Mock_. En todos aquellos casos en los que necesitábamos saber cómo se habían comunicado la unidad bajo test y su colaborador, hemos ido muy bien servidas por un espía. Además, y gracias a eso, las aserciones o expectativas han quedado bien reflejadas en el propio test.

Publicado: March 3, 2025

[https://franiglesias.github.io/test-doubles-guide-2/](https://franiglesias.github.io/test-doubles-guide-2/)
