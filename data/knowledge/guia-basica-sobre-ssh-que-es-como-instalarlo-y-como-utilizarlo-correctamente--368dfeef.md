---
title: "Guía básica sobre SSH: Qué es, cómo instalarlo y cómo utilizarlo correctamente"
notion_id: 368dfeef-7230-4d41-9467-840dbc95d610
notion_url: https://app.notion.com/p/Gu-a-b-sica-sobre-SSH-Qu-es-c-mo-instalarlo-y-c-mo-utilizarlo-correctamente-368dfeef72304d419467840dbc95d610
last_edited: 2026-09-21T17:24:00.000Z
source_url: https://raiolanetworks.com/blog/ssh/
tags: ["Español", "SysAdmin", "Network", "Article", "Guide", "Raiola Networks Blog"]
---
[https://raiolanetworks.es/blog/ssh/](https://raiolanetworks.es/blog/ssh/)

SSH sirve para conectarte y gestionar un servidor remoto mediante una conexión cifrada, ejecutando comandos sin desplazarte al centro de datos. Emplea cifrado simétrico (AES, blowfish) para las comunicaciones y cifrado asimétrico opcional, con claves generadas mediante ssh-keygen, para autenticar al usuario.

SSH es el protocolo con el que se administra un servidor sin estar físicamente delante de él, y existe por necesidad: desde que hay equipos informáticos que funcionan a modo de servidor, es decir, ordenadores que ofrecen servicios a otros usuarios como un servidor de hosting compartido de Raiola Networks, hizo falta alguna forma de administrarlos en remoto.

En este articulo vamos a ver qué es SSH, cómo funciona, qué tipos de cifrado utiliza y cuáles son los comandos más utilizados, para que sepas lo que estás haciendo de verdad cuando abres una conexión.

Índice del artículo

- Qué es SSH y para qué sirve
- Cómo funciona SSH
- Tipos de cifrado
- Cómo instalar y configurar SSH
- Cómo instalar un servidor SSH
- Cómo instalar un cliente SSH
- Cómo acceder por SSH
- Cómo acceder por SSH en MAC
- Cómo acceder por SSH en Windows
- Cómo acceder por SSH en Linux

¡Recibe nuestros contenidos en tu correo!

- Recibirás nuestra newsletter de forma totalmente gratuita. ¡Prometemos no enviarte spam!
- Responsable: RAIOLA NETWORKS, S.L. CIF: B27453489 Avda de Magoi, 66, Semisótano, Dcha., 27002 Lugo (Lugo) Teléfono: +34 982776081 e-mail: info@raiolanetworks.es Finalidad: Atender solicitudes de información, ejecución de la contratación de servicios y remisión de comunicaciones comerciales. Legitimación: Consentimiento del interesado y contratación de productos y/o servicios del Responsable. Destinatarios: No se ceden datos a terceros, salvo obligación legal. Personas físicas o jurídicas directamente relacionadas con el Responsable Encargados de Tratamiento adheridos al Privacy Shield. Derechos: Acceder, rectificar y suprimir los datos, portabilidad de los datos, limitación u oposición a su tratamiento, derecho a no ser objeto de decisiones automatizadas, así como a obtener información clara y transparente sobre el tratamiento de sus datos. Información adiccional: Para obtener información más detallada, puede consultar nuestra política de privacidad

## Qué es SSH y para qué sirve

Como ya te comentaba en la introducción, SSH (acrónimo de Secure Shell) es un protocolo de red para administrar un servidor de forma remota, desde tu oficina o desde tu despacho en casa, sin desplazarte al centro de datos donde está alojado.

Para entender SSH primero tienes que saber qué es telnet, su predecesor.

Con telnet puedes conectarte a un equipo remoto y mover archivos de una carpeta a otra o eliminarlos, o sea, usar los comandos típicos de SSH (como cp o rm), y también cosas más complejas como crear una base de datos e insertar contenido en ella.

La gran desventaja de telnet es que las conexiones van en texto plano. Cualquiera que espíe tu red entiende toda la conversación entre el servidor y tu ordenador y puede capturar tu contraseña sin ningún esfuerzo.

En la siguiente imagen puedes ver un ejemplo de una conexión telnet a un servidor VPS:

Te explico lo que significa todo lo que ves en pantalla:

- telnet 91.134.16.2: Se realiza una conexión de tipo telnet al servidor 91.134.16.2
- nuevo login: root El servidor me pide un usuario y utilizo root, el superadministrador en sistemas operativos de tipo unix.
- Password: Aunque no se ven puntos ni asteriscos, aquí he tecleado mi contraseña.
- [ root@raiolanetworks.servidordepruebas.com ] # ls -a Aquí ejecuto el comando ls con el modificador -a para listar todos los contenidos de un directorio, incluidos archivos ocultos.

Como te comentabamos, SSH sucedió a telnet y lo dejó relegado a herramienta de diagnóstico, porque las conexiones por SSH sí son seguras.

Una conexión SSH utiliza algoritmos de cifrado actuales, así que aunque alguien espíe tu red no obtiene más que una cadena ininteligible de caracteres que no puede descifrar.

Resumiendo: SSH es un protocolo de red cifrado que te permite manejar un servidor remoto como si lo tuvieras delante.

## Cómo funciona SSH

En su forma más básica, SSH te conecta a un equipo remoto con un usuario y una contraseña, y con un par de comandos ya estás manejando el servidor.

Antes de abrir el terminal y conectarte, vamos con los distintos métodos de cifrado, que son los que explican por qué este protocolo es tan seguro. Te va a servir además para entender cómo funcionan las comunicaciones seguras de cualquier otro protocolo de red, como FTPS, SFTP o HTTPS.

### Tipos de cifrado

La forma en que el protocolo SSH gestiona las conexiones es tremendamente segura.

Cuando aprendas a iniciar una conexión SSH desde un terminal podrás comprobarlo tú mismo, revisando los tipos de certificado que ofrecen tu cliente SSH y el servidor al que te estás conectando, con el comando ssh -vv usuario@servidor.

Vamos con los tipos de cifrado que existen y cuáles utilizar para que tus comunicaciones sean perfectamente seguras.

### Cifrado simétrico o criptografía de clave secreta

La criptografía simétrica es la más antigua de todas y se basa en que dos ordenadores que se comunican entre sí comparten un secreto. De ahí el segundo nombre que recibe este método, cifrado de clave secreta.

La clave compartida puede ser una palabra, un número o una cadena aleatoria de caracteres, y sirve para modificar de una manera preestablecida los mensajes que se envían los dos equipos. Sabiendo la clave y la forma en que los mensajes son cifrados con ella, se codifican y se descodifican.

Aquí la seguridad depende casi en exclusiva de usar un secreto fuerte, y no tanto del método de cifrado, que un atacante puede conocer sin problema: mientras no tenga la clave, no puede descifrar el mensaje.

SSH utiliza este método de cifrado para las comunicaciones.

No vamos a entrar en los detalles técnicos de cómo se hace en este articulo, pero cuando te conectas por SSH el programa que usas a modo de cliente y el servidor negocian un tipo de cifrado: AES, blowfish o arcfour en sus distintas variantes, por ejemplo.

Cuando se ponen de acuerdo, el cliente y el servidor crean un secreto en paralelo mediante el algoritmo de intercambio de claves, basado en el protocolo diffie-hellman. Los dos llegan a la misma clave por su cuenta, compartiendo cierta información pública y manipulándola con información secreta.

### Cifrado asimétrico

En contra de lo que se puede pensar, el protocolo SSH solamente utiliza cifrado asimétrico para la autenticación de usuarios, y ni siquiera es obligatorio, porque lo normal es conectarse con nombre de usuario y contraseña. Lo que hace es añadir una capa de seguridad extra a las conexiones SSH.

Se basa en que cada equipo de la comunicación tiene un par de claves: una clave pública, que como su nombre indica puede mostrarse a cualquiera, y una clave privada.

Para que veas cómo funciona, primero te enseño a generar esas claves. Con el comando ssh-keygen generas una pareja de clave pública y privada como las que se usan para las conexiones seguras https o SSH.

```plain text
ssh-keygen -t rsa -b 4096 -f clave_secreta -N ''
```

Como ves, el método usado para generar la clave es el cifrado RSA. Existen otros muchos, como EDCSA o ed25519. Esta sería la explicación para el comando anterior:

- -b 4096: El número de bits de la clave privada es 4096.
- -f clave_secreta: La clave se guarda en un archivo llamado clave_secreta.
- -N‘’: Evita que solicite contraseña a la hora de usarla (sí, se puede añadir seguridad extra a las claves para que nadie pueda usarlas sin contraseña).

Estas claves funcionan de modo que la clave privada, que es la parte realmente importante, no sea accesible para otros. Es la llave con la que se descifran los mensajes.

La clave pública sí puede compartirse, porque a partir de una misma clave privada pueden generarse infinidad de claves públicas distintas. De hecho, una clave pública viaja en la comunicación inicial entre el cliente y el servidor SSH.

Una vez dispones de una clave pública y otra privada, puedes iniciar la conexión SSH con cifrado asimétrico: si la clave privada figura como 'autorizada' en el servidor, la autenticación es más rápida y más segura.

Lo que ocurre por dentro es que el servidor, cuando comprueba que tu clave está autorizada, crea un mensaje aleatorio cifrado con la clave pública y te lo devuelve para ver si eres capaz de descifrarlo.

Tu clave privada lo descifra y lo devuelve. Si concuerda con el que generó al principio, la autenticación es correcta y entras por SSH sin usar contraseña.

### Hashing

Ahora vamos a ponernos todavía más técnicos y hablar de matemáticas, que es donde se llama función hash a la que no tiene inversa. Traducido al lenguaje humano, esto quiere decir que un mensaje cifrado con una función hash no puede descifrarse.

Evidentemente, este es el método más seguro para almacenar cualquier tipo de información. Así es como los sistemas operativos (o WordPress, nuestro CMS preferido) guardan las contraseñas en sus bases de datos: una vez has pasado un mensaje por una función hash, no hay forma de recuperar el original.

SSH implementa este cifrado en cada mensaje que se envían el cliente y el servidor. Primero codifica la información con la clave que salió del cifrado simétrico, y con el mensaje ya preparado genera un hash único mediante un método preestablecido (md5, sha o el que se haya acordado durante la fase de negociación).

Esto añade una capa extra de seguridad (otra más): aunque un atacante haya obtenido el secreto compartido de la conexión cifrada asimétrica, no puede modificar la información sin que los equipos involucrados se den cuenta, porque al modificarla, el hash varía inevitablemente.

Cuando el receptor comprueba que lo recibido tiene el mismo hash que le envió su interlocutor, sabe que la información es fiable y que nadie la ha alterado.

Vamos a verlo con un ejemplo:

```plain text
echo -n "Raiola Networks" | shasum c7ccfb61f727966a2f7a7ed9e5824f457bbcb10a
```

```plain text
echo -n "Raiola manda y no el panda" | shasum b59996e8011e11a286d685ef12f1cb2fe3e1f2f4
```

Como ves, la cadena cifrada de 'Raiola Networks' y la de 'Raiola manda y no el panda' es distinta y única: del resultado no podemos sacar el texto inicial de ninguna forma.

En cambio, si vuelves a pasar el texto original por la misma función hash, sale exactamente el mismo resultado. De ahí que te sirva para certificar que el mensaje original no se ha alterado.

## Cómo instalar y configurar SSH

Cuando hablamos de instalar SSH puedo referirme a dos cosas distintas:

- Servidor SSH: el programa que configuras en el servidor y que acepta las conexiones SSH. Si estás alojado con nosotros, no te preocupes, porque normalmente no tendrás que instalarlo: en los productos de Raiola Networks se te proporciona instalado y preconfigurado por defecto.
- Cliente SSH: la aplicación con la que abres la conexión y que, en términos informáticos, actúa a modo de cliente. Según el sistema operativo que tengas, puede que ya venga uno instalado por defecto o que necesites instalarlo.

### Cómo instalar un servidor SSH

Te explico la instalación en dos familias de Linux distintas, las basadas en Debian y las basadas en Red Hat, que son dos de las más importantes y usan sistemas de instalación diferentes.

El comando que necesitas para instalar un servidor SSH en Ubuntu (y sistemas basados en Debian) es:

```plain text
apt install -y openssh-server
```

El comando que necesitas para instalar un servidor SSH en sistemas operativos como AlmaLinux, basados en RedHat es:

```plain text
dnf install -y openssh-server
```

Como ves, la diferencia no es demasiado grande: cambia el comando de instalación, pero el paquete es el mismo.

Ahora vamos con la instalación de verdad, porque quiero mejorar la seguridad del servidor al que me conecté antes por telnet y aplicarle los métodos de cifrado que acabo de explicarte.

Como mi servidor tiene un Linux basado en Red Hat (CentOS), voy a instalar un servidor SSH con el comando 'dnf install -y openssh-server'

Cuando acaba la instalación, configuro el servidor para mejorar su seguridad: cambio el puerto, impido el acceso del usuario root mediante contraseña y establezco un tiempo máximo de vida para las conexiones SSH, de forma que mueran si hay demasiado tiempo entre un mensaje y otro.

- Para cambiar el puerto, modificaré la directiva Port de 22 a 11022, el que usan los planes de hosting compartido en Raiola Networks para conexiones SFTP.
- Para impedir que el usuario root pueda conectarse con contraseña, tengo que modificar el valor de PermitRootLogin de Yes a prohibit-password.
- Voy a impedir también que nadie pueda conectarse con contraseñas vacías que, aunque parezca raro, puede suceder. Modifico el valor de PermitEmptyPasswords a No.
- Por último, para establecer un timeout de 5 minutos en la conexión, doy a ClientAliveInterval el valor de 300 segundos.
- En cuanto hayas hecho estos cambios, los guardas y reinicias el servicio sshd (para que tome las nuevas configuraciones) con el comando systemctl restart sshd

Antes de desconectarte, tienes que autorizar la clave que creaste, porque has configurado que root solo pueda acceder con clave privada. El archivo donde la guardas es /root/.ssh/authorized_keys.

Ahora, desde un nuevo terminal, pruebas la conexión.

Y si quieres ponerte creativo y mostrar un mensaje propio para que nadie adivine qué servidor SSH usas y en qué versión, edita el banner de bienvenida cambiando el valor de Banner por un fichero de texto personalizado, como hago yo aquí:

```plain text
echo ‘Bienvenido a tu nuevo servidor en Raiola Networks!’ >> /etc/motd
```

### Cómo instalar un cliente SSH

A diferencia del servidor, existen clientes SSH para todos los sistemas operativos, así que te explico cómo instalar uno en Windows, Linux y MAC.

### Cómo instalar SSH en MACOS

MacOS es un sistema operativo de tipo Unix, así que te lo pone muy fácil: entre sus herramientas preinstaladas cuenta con un cliente SSH.

Para utilizarlo, solo tienes que buscar la aplicación Terminal y abrirla. Con eso ya puedes establecer tu primera conexión SSH.

### Cómo instalar SSH en Windows

Windows, a diferencia de macOS, no trae cliente SSH por defecto, así que necesitas uno. Hay multitud de aplicaciones; yo voy a usar Putty:

- Pulsa en el enlace marcado en la captura:
- Descarga Putty pulsando en el enlace correspondiente a tu sistema operativo. El mío es de 64 bits:
- Busca el archivo descargado y haz doble clic para ejecutarlo. Acepta que se ejecute:
- Pulsa Siguiente para comenzar la instalación:
- Si la ruta en que se va a instalar el programa es válida, pulsa Siguiente:
- De nuevo, deja las opciones por defecto si no quieres modificar alguna y pulsa Siguiente:
- Finaliza la instalación:

### Cómo instalar SSH en Linux

Al igual que en MacOS, los equipos Linux suelen traer un cliente SSH instalado, así que solo tienes que ir a "Inicio" > "Herramientas" > "Terminal", abrir la aplicación y usar SSH exactamente igual que en MacOS.

## Cómo acceder por SSH

Con el cliente SSH instalado, toca conectarse al servidor; vamos a ver cómo se utiliza en cada sistema operativo.

### Cómo acceder por SSH en MAC

Con el terminal abierto del paso anterior, para tu primera conexión SSH solo tienes que escribir:

```plain text
ssh usuario@servidor
```

Yo voy a conectarme a un servidor optimizado de Raiola Networks con los datos del email de bienvenida:

```plain text
ssh root@raiolanetworks.servidordepruebas.com
```

El programa te pedirá la contraseña y, esto es muy importante, mientras la escribes no aparecerán puntos, asteriscos ni nada similar: parece que no pasa nada. Es una medida de seguridad para que nadie sepa la longitud de tu clave contando los puntos o los asteriscos. Escríbela, pulsa ENTER y ya estás conectado.

### Cómo acceder por SSH en Windows

Ya tienes instalado Putty, así que ahora vas a hacer tu primera conexión SSH desde Windows.

Para conectarte, tienes que usar el nombre del servidor o su dirección IP y pulsar en Open:

Ahora se te muestra la clave pública del servidor; acéptala para confiar en ella.

Esto te pedirá el nombre de usuario y, una vez establecido, la contraseña.

Al igual que en MacOS, no aparecerán caracteres al escribir la contraseña. Escríbela, pulsa ENTER y ya has hecho tu primera conexión SSH.

### Cómo acceder por SSH en Linux

En Linux estableces la conexión exactamente igual que hiciste antes en MacOS. Abre un terminal y teclea:

```plain text
ssh root@raiolanetworks.servidordepruebas.com
```
