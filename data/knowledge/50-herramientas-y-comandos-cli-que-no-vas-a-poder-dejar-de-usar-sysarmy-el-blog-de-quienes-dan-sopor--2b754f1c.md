---
title: "50 herramientas y comandos CLI que no vas a poder dejar de usar :: Sysarmy — El blog de quienes dan soporte"
notion_id: 2b754f1c-7d23-8160-9353-d3f6337c96c6
notion_url: https://app.notion.com/p/50-herramientas-y-comandos-CLI-que-no-vas-a-poder-dejar-de-usar-Sysarmy-El-blog-de-quienes-dan--2b754f1c7d2381609353d3f6337c96c6
last_edited: 2025-11-26T19:08:00.000Z
source_url: https://sysarmy.com/blog/posts/50-herramientas-cli-no-vas-a-poder-dejar-de-usar/
tags: ["Article", "Sysarmy Blog", "Español", "SysAdmin", "Productivity", "Command Line", "Automation"]
---
![image](https://sysarmy.com/blog/assets/50-cli-command-tools.png)

_La versión original de este post se puede encontrar en _[_dev.to/lissy93_](https://dev.to/lissy93/cli-tools-you-cant-live-without-57f6)_ (inglés)._

Como desarrolladores, pasamos gran parte de nuestro tiempo en el terminal. Hay un montón de herramientas CLI útiles que pueden hacer tu vida en la línea de comandos más fácil, más rápida y más divertida.

Este post describe mi top 50 de herramientas CLI imprescindibles, en las que he llegado a confiar.

Al final del artículo, he incluido algunas secuencias de comandos para ayudarle a automatizar la instalación y actualización de estas herramientas en varios sistemas / distros.

↕️ **Índice**

- [Utilidades](https://sysarmy.com/blog/posts/50-herramientas-cli-no-vas-a-poder-dejar-de-usar/#utilidades)
- [Productividad](https://sysarmy.com/blog/posts/50-herramientas-cli-no-vas-a-poder-dejar-de-usar/#productividad)

## Utilidades

> thefuck es una de esas utilidades sin las que no podrás vivir una vez que la hayas probado. Siempre que escribas mal un comando y obtengas un error, simplemente ejecuta fuck y lo autocorregirá. Usa arriba/abajo para elegir una corrección, o simplemente ejecuta fuck --yeah para ejecutar inmediatamente la más probable.

![image](https://sysarmy.com/blog/assets/50-thefuck.gif)

```plain text
# macOS (vía Homebrew)
brew install thefuck

# Arch Linux
sudo pacman -S thefuck

# FreeBSD
pkg install thefuck

```

> zoxide te permite saltar a cualquier directorio sin necesidad de recordar o especificar su ruta completa. Recuerda los directorios que has visitado, para que puedas saltar rápidamente; ni siquiera necesitas escribir el nombre completo de la carpeta. También dispone de una opción de selección interactiva, mediante fzf, que permite filtrar los resultados de los directorios en tiempo real.

![image](https://sysarmy.com/blog/assets/50-zoxide.gif)

```plain text
# macOS (vía Homebrew)
brew install zoxide

# Arch Linux
sudo pacman -S zoxide

# Debian / Ubuntu
sudo apt install zoxide

# FreeBSD
pkg install zoxide

# Otros (vía Rust Creates)
cargo install zoxide --locked

```

> tldr es una enorme colección de páginas de manual mantenidas por la comunidad. A diferencia de las páginas de manual tradicionales, están resumidas, contienen ejemplos de uso útiles y están coloreadas para facilitar la lectura.

![image](https://sysarmy.com/blog/assets/50-tldr.gif)

```plain text
# macOS (vía Homebrew)
brew install tldr

# Otros (vía NPM)
npm install -g tldr

```

> scc te da un desglose del número de líneas de código Escrito en cada lenguaje para un directorio específico. También muestra algunas estadísticas divertidas, como el coste estimado de desarrollo e información sobre la complejidad. Es increíblemente rápido, muy preciso y tiene soporte para una amplia gama de lenguajes.

![image](https://sysarmy.com/blog/assets/50-scc.webp)

```plain text
# macOS (vía Homebrew)
brew install scc

# Otros (vía go)
go install github.com/boyter/scc/v3@latest

```

> exa es un reemplazo moderno basado en Rust para el comando ls, para listar archivos. Puede mostrar iconos de tipo de archivo, colores, información de archivo/carpeta y tiene varios formatos de salida - árbol, cuadrícula o lista.

![image](https://sysarmy.com/blog/assets/50-exa.webp)

```plain text
# macOS (vía Homebrew)
brew install exa

# Arch Linux
sudo pacman -S exa

# Debian / Ubuntu
sudo apt install exa

```

> duf es genial para mostrar información sobre discos montados y comprobar el espacio libre. Produce una salida clara y colorida, e incluye opciones para ordenar y personalizar los resultados.

![image](https://sysarmy.com/blog/assets/50-duf.png)

```plain text
# macOS (vía Homebrew)
brew install duf

# Arch Linux
sudo pacman -S duf

# Debian / Ubuntu
sudo apt install duf

# FreeBSD
pkg install duf

```

> aria2 es una utilidad ligera, multiprotocolo, de reanudación de descargas para HTTP/HTTPS, FTP, SFTP, BitTorrent y Metalink, con soporte para control a través de una interfaz RPC. Es increíblemente rico en funciones, y tiene toneladas de opciones. También existe ziahamza/webui-aria2 - un buen compañero de interfaz web.

![image](https://sysarmy.com/blog/assets/50-aria2c.png)

```plain text
# macOS (vía Homebrew)
brew install aria2

# Arch Linux
sudo pacman -S aria2

# Debian / Ubuntu
sudo apt install aria2

```

> bat es un clon de cat con resaltado de sintaxis e integración con git. Escrito en Rust, es muy performante, y tiene varias opciones para personalizar la salida y la tematización. Hay soporte para tuberías automáticas y concatenación de archivos.

![image](https://sysarmy.com/blog/assets/50-bat.png)

```plain text
# macOS (vía Homebrew)
brew install bat

# Arch Linux
sudo pacman -S bat

# Debian / Ubuntu
sudo apt install bat

```

> diff-so-fancy le da un mejor aspecto a los diffs para comparar cadenas, archivos, directorios y cambios git. El resaltado de cambios hace que detectar los cambios sea mucho más fácil, y puedes personalizar el diseño y los colores de la salida.

![image](https://sysarmy.com/blog/assets/50-diff.png)

```plain text
# macOS (vía Homebrew)
brew install diff-so-fancy

# Arch Linux
sudo pacman -S diff-so-fancy

# Debian / Ubuntu
sudo apt install diff-so-fancy

```

> entr le permite ejecutar un comando arbitrario cada vez que un archivo cambia. Puedes pasar un archivo, directorio, enlace simbólico o regex para especificar qué archivos debe vigilar. Es realmente útil para reconstruir proyectos automáticamente, reaccionar a logs, pruebas automatizadas, etc. A diferencia de otros proyectos similares, utiliza kqueue(2) o inotify(7) para evitar el sondeo y mejorar el rendimiento.

![image](https://sysarmy.com/blog/assets/50-entr.png)

```plain text
# macOS (vía Homebrew)
brew install entr

# Arch Linux
sudo pacman -S entr

# Debian / Ubuntu
sudo apt install entr

```

> exiftool es una práctica utilidad para leer, escribir, eliminar y crear metainformación para una amplia variedad de tipos de archivos. No vuelva a filtrar accidentalmente su ubicación al compartir una foto.

![image](https://sysarmy.com/blog/assets/50-exiftool.png)

> fdupes se utiliza para identificar y/o eliminar archivos duplicados dentro de directorios especificados. Es útil para liberar espacio en disco cuando tiene dos o más archivos idénticos.

![image](https://sysarmy.com/blog/assets/50-fdupes.png)

> fzf es un buscador de archivos difusos y una herramienta de filtrado extremadamente potente y fácil de usar. También tiene plugins disponibles para la mayoría de los shells e IDEs, para mostrar resultados instantáneos durante la búsqueda. Este post de Alexey Samoshkin destaca algunos de sus casos de uso.

![image](https://sysarmy.com/blog/assets/50-fzf.gif)

```plain text
# macOS (vía Homebrew)
brew install fzf

# Arch Linux
sudo pacman -S fzf

# Debian / Ubuntu
sudo apt install fzf

```

> hyperfine facilita la evaluación comparativa precisa de comandos o scripts arbitrarios. Se encarga de las ejecuciones de calentamiento, de limpiar la caché para obtener resultados precisos y de evitar interferencias de otros programas. También puede exportar los resultados como datos sin procesar y generar gráficos.

![image](https://sysarmy.com/blog/assets/50-hyperfine.png)

```plain text
# macOS (vía Homebrew)
brew install hyperfine

# Arch Linux
sudo pacman -S hyperfine

# Debian / Ubuntu
sudo apt install hyperfine

```

> just es similar a make pero con algunas buenas adiciones. Te permite agrupar los comandos de tus proyectos en recopilaciones, que pueden ser fácilmente listadas y ejecutadas. Soporta alias, argumentos posicionales, diferentes shells, integración dot env, interprulación de cadenas, y casi todo lo que puedas necesitar.

![image](https://sysarmy.com/blog/assets/50-just.png)

```plain text
# macOS (vía Homebrew)
brew install just

# Arch Linux
sudo pacman -S just

# Debian / Ubuntu
sudo apt install just

```

> jq es como sed, pero para JSON - puedes usarlo para cortar y filtrar y mapear y transformar datos estructurados con facilidad. Se puede utilizar para escribir consultas complejas para extraer o manipular datos JSON. También hay un jq playground que puedes usar para probarlo, o formular consultas con información en tiempo real.

> most es un paginador, para leer archivos largos o salidas de comandos. most soporta multi-ventanas y tiene la opción de no ajustar el texto.

> procs es un visor de procesos fácil de navegar, tiene resaltado de color, facilita la clasificación y búsqueda de procesos, tiene árbol Ver y se actualiza en tiempo real.

> rip es una herramienta de borrado segura, ergonómica y eficaz. Le permite eliminar archivos y directorios de forma intuitiva y restaurar fácilmente los archivos eliminados.

![image](https://sysarmy.com/blog/assets/50-rip.webp)

> ripgrep es una herramienta de búsqueda orientada a líneas que busca recursivamente un patrón regex en el directorio actual. Puede ignorar el contenido de .gitignore y omitir archivos binarios. Es capaz de buscar dentro de archivos comprimidos, o sólo buscar en una extensión específica, y entiende los archivos que utilizan varios métodos de codificación

![image](https://sysarmy.com/blog/assets/50-ripgrep.webp)

> rsync te permite copiar archivos grandes localmente o desde o hacia hosts remotos o unidades externas. Puede utilizarse para mantener sincronizados archivos de varias ubicaciones y es perfecto para crear, actualizar y restaurar copias de seguridad.

> sd es una herramienta de búsqueda y reemplazo fácil, rápida e intuitiva, basada en literales de cadena. Puede ejecutarse sobre un fichero, un directorio entero o cualquier texto canalizado.

> tre muestra una lista de archivos en forma de árbol para tu directorio actual o uno especificado, con colores. Al ejecutarlo con la opción -e, numera cada elemento y crea un alias temporal que puedes usar para saltar rápidamente a esa ubicación

![image](https://sysarmy.com/blog/assets/50-tre.webp)

> xsel te permite leer y escribir en el portapapeles de selección X desde la línea de comandos. Es útil para enviar la salida de un comando al portapapeles o para pegar datos copiados en un comando.

> Muestra el uso de ancho de banda, información de conexión, hosts salientes y consultas DNS en tiempo real

> Como top, pero para monitorear el uso de recursos de contenedores en ejecución (Docker y runC). Muestra CPU, memoria y ancho de banda de red en tiempo real, así como el nombre, estado e ID de cada contenedor. También incluye un visor de logs integrado y opciones para gestionar (detener, iniciar, ejecutar, etc.) contenedores.

![image](https://sysarmy.com/blog/assets/50-ctop.webp)

> bpytop es un monitor de recursos rápido, interactivo y visual. Muestra los procesos principales en ejecución, el historial reciente de CPU, memoria, disco y red. Desde la interfaz puedes navegar, ordenar y buscar; también hay soporte para temas de color personalizados.

![image](https://sysarmy.com/blog/assets/50-bpytop.webp)

> glances es otro monitor de recursos, pero con un conjunto de características diferente. Incluye una vista web totalmente responsiva, una API REST y monitoreo histórico. Es fácilmente extensible y puede integrarse con otros servicios.

![image](https://sysarmy.com/blog/assets/50-glances.webp)

> gping puede ejecutar pruebas de ping en varios hosts, mostrando los resultados en un gráfico en tiempo real. También puede usarse para monitorear el tiempo de ejecución, cuando se usa con la opción --cmd.

![image](https://sysarmy.com/blog/assets/50-gping.gif)

> dua-cli te permite ver de forma interactiva el espacio usado y disponible en cada unidad montada, y facilita liberar almacenamiento.

![image](https://sysarmy.com/blog/assets/50-dua-cli.webp)

> speedtest-cli simplemente ejecuta una prueba de velocidad de internet, usando speedtest.net, pero directamente desde la terminal :)

![image](https://sysarmy.com/blog/assets/50-speedtest-cli.webp)

> dog es un cliente de búsqueda DNS fácil de usar, con soporte para DoT y DoH, salidas coloridas y la opción de emitir JSON.

![image](https://sysarmy.com/blog/assets/50-dog.webp)

### Navega por la web, escucha música, revisa correos, gestiona calendarios, lee noticias y más, ¡todo sin salir de la terminal!

> browsh es un navegador de texto moderno, interactivo y en tiempo real, renderizado para TTYs y navegadores. Soporta navegación tanto con ratón como con teclado, y tiene muchas más funciones de lo que esperarías para una aplicación puramente de terminal. Además, ayuda a reducir el consumo de batería que afecta a los navegadores modernos y, con soporte para MoSH, puedes experimentar tiempos de carga más rápidos gracias al menor uso de ancho de banda.

![image](https://sysarmy.com/blog/assets/50-browsh.webp)

> buku es un gestor de marcadores para terminal, con muchas opciones de configuración, almacenamiento y uso. También dispone de una interfaz web opcional y un plugin para navegador para acceder a tus marcadores fuera de la terminal.

![image](https://sysarmy.com/blog/assets/50-buku.webp)

> cmus es un reproductor de música para terminal, controlado con atajos de teclado. Soporta una amplia gama de formatos y códecs de audio, y permite organizar pistas en listas de reproducción y aplicar configuraciones de reproducción.

![image](https://sysarmy.com/blog/assets/50-cmus.webp)

> cointop muestra los precios actuales de criptomonedas y el historial de precios de tu portafolio. Soporta alertas de precio, gráficos históricos, conversión de divisas, búsqueda difusa y mucho más. Puedes probar la demo en la web cointop.sh o ejecutando ssh cointop.sh.

![image](https://sysarmy.com/blog/assets/50-cointop.webp)

> ddgr es como googler, pero para DuckDuckGo. Es rápido, limpio y fácil, con soporte para respuestas instantáneas, autocompletado, bangs de búsqueda y búsqueda avanzada. Respeta tu privacidad por defecto, también tiene soporte para proxy HTTPS y funciona con Tor.

![image](https://sysarmy.com/blog/assets/50-ddgr.webp)

> micro es un editor de código fácil de usar, rápido y extensible, con soporte para ratón. Como viene en un solo binario, la instalación es tan simple como curl https://getmic.ro | bash.

![image](https://sysarmy.com/blog/assets/50-micro.webp)

> khal es una aplicación de calendario para terminales, que muestra los próximos eventos, vistas del mes y de la agenda. Puedes sincronizarlo con cualquier calendario CalDAV, y añadir, editar y eliminar eventos directamente.

![image](https://sysarmy.com/blog/assets/50-khal.webp)

> mutt es un cliente de correo clásico, basado en terminal, para enviar, leer y gestionar correos electrónicos. Soporta todos los principales protocolos de correo electrónico y formatos de buzón, permite archivos adjuntos, CCO/CC, hilos, listas de correo y notificaciones de estado de entrega.

![image](https://sysarmy.com/blog/assets/50-mutt.webp)

> newsboat es un lector y agregador de feeds RSS, para leer las noticias, blogs y seguir las actualizaciones directamente desde el terminal.

![image](https://sysarmy.com/blog/assets/50-newsboat.webp)

> rclone es una práctica utilidad para sincronizar archivos y carpetas con varios proveedores de almacenamiento en la nube. Puede invocarse directamente desde la línea de comandos o integrarse fácilmente en un script como sustituto de las pesadas aplicaciones de sincronización de escritorio.

> taskwarrior es una aplicación CLI de gestión de tareas/todo. Es simple y discreta, pero también increíblemente potente y escalable, con funciones avanzadas de organización y consulta. También hay un montón (¡más de 700!) de [plugins] adicionales (https://taskwarrior.org/tools/) para ampliar su funcionalidad e integrarse con servicios de terceros.

![image](https://sysarmy.com/blog/assets/50-taskwarrior.webp)

> tuir es genial si quieres parecer que estás trabajando, ¡mientras en realidad navegas por Reddit! Tiene combinaciones de teclas intuitivas, temas personalizados, y también puede renderizar imágenes y contenido multimedia. También está haxor para noticias de hackers.

![image](https://sysarmy.com/blog/assets/50-tuir.webp)

> httpie es un cliente HTTP, para pruebas, depuración y uso de APIs. Soporta todo lo que se puede esperar - HTTPS, proxies, autenticación, cabeceras personalizadas, sesiones persistentes, análisis JSON. El uso es simple con una sintaxis expresiva y salida coloreada. Al igual que otros clientes HTTP (Postman, Hopscotch, Insomnia, etc) HTTPie también incluye una interfaz web.

![image](https://sysarmy.com/blog/assets/50-httpie.webp)

> lazydocker es una aplicación de gestión de Docker, que te permite Ver todos los contenedores e imágenes, gestionar su estado, leer registros, comprobar el uso de recursos, reiniciar/reconstruir, analizar capas, podar contenedores, imágenes y volúmenes no utilizados, y mucho más. Te ahorra la necesidad de recordar, escribir y encadenar múltiples comandos Docker.

![image](https://sysarmy.com/blog/assets/50-lazydocker.webp)

> lazygit es un cliente git visual, en la línea de comandos. Fácilmente añadir, confirmar y puch archivos, resolver conflictos, comparar diffs, gestión de registros, y hacer operaciones complejas como aplasta y rebobina. Hay combinaciones de teclas para todo, colores, y es fácilmente configurable y ampliable.

![image](https://sysarmy.com/blog/assets/50-lazygit.webp)

> kdash es una herramienta de gestión de Kubernetes todo en uno. Ver métricas de nodo, ver recursos, transmitir registros de contenedor, analizar contextos y gestionar nodos, pods y espacios de nombres.

> gdb-dashboard añade un elemento visual al depurador de GNU, para depurar programas C y C++. Analiza fácilmente la memoria, pasa por puntos de interrupción y registros Ver.

![image](https://sysarmy.com/blog/assets/50-gdp-dashboard.webp)

## Servicios Externos

> ngrok de forma segura* expone tu localhost a internet detrás de una URL única. Esto te permite compartir lo que estás trabajando con tus colegas remotos, en tiempo real. El uso es muy simple, pero también tiene un montón de características avanzadas para cosas como la autenticación, webhooks, cortafuegos, inspección de tráfico, dominios personalizados / comodín y mucho más.

![image](https://sysarmy.com/blog/assets/50-ngrok.webp)

> tmate te permite compartir instantáneamente una sesión de terminal en vivo con alguien en otra parte del mundo. Funciona a través de diferentes sistemas, soporta control de acceso / auth, puede ser auto-alojado, y tiene todas las características de Tmux

> asciinema es muy útil para grabar, compartir e incrustar fácilmente una sesión de terminal. Ideal para mostrar algo que has construido, o para mostrar los pasos de línea de comandos para un tutorial. A diferencia de los vídeos de grabación de pantalla, el usuario puede copiar y pegar el contenido, cambiar el tema sobre la marcha y controlar la reproducción.

> navi permite navegar por las hojas de trucos y ejecutar comandos. Los valores sugeridos para los argumentos se muestran dinámicamente en una lista. Escribe menos, reduce los errores y ahórrate tener que memorizar miles de comandos. Se integra con tldr y cheat.sh para obtener contenido, pero también puedes importar otras cheatsheets, o incluso escribir las tuyas propias.

> transfer hace que subir y compartir archivos sea realmente fácil, directamente desde la línea de comandos. Es gratis, soporta encriptación, te da una URL única, y también puede ser auto-alojado. He escrito una función de ayuda Bash para hacer el uso un poco más fácil, puedes encontrarla aquí o probarla ejecutando bash <(curl -L -s https://alicia.url.lol/transfer).

![image](https://sysarmy.com/blog/assets/50-transfer-sh.webp)

> surge es un proveedor de alojamiento estático gratuito, que puedes desplegar directamente desde el terminal en un solo comando, ¡sólo ejecuta surge desde dentro de tu directorio dist! Soporta dominios personalizados, certificados SSL automáticos, soporte pushState, soporte de recursos cross-origin - ¡y es gratis!

![image](https://sysarmy.com/blog/assets/50-surge.webp)

> wttr.in es un servicio que muestra el tiempo en un formato digerible en la línea de comandos. Sólo tienes que ejecutar curl wttr.in o curl wttr.in/London para probarlo. Hay parámetros de URL para personalizar los datos devueltos, así como el formato.

![image](https://sysarmy.com/blog/assets/50-wttr-in.webp)

> cowsay es una vaca parlante configurable. Está basado en el original de Tony Monroe

![image](https://sysarmy.com/blog/assets/50-cowsay.webp)

> figlet muestra el texto como arte ASCII

![image](https://sysarmy.com/blog/assets/50-figlet.webp)

> lolcat colorea de arco iris cualquier texto que se le pase

![image](https://sysarmy.com/blog/assets/50-lolcat.webp)

> neofetch imprime información de la distro y del sistema (para que puedas flexionar que usas Arch btw en r/unixporn)

![image](https://sysarmy.com/blog/assets/50-neofetch.webp)

Como ejemplo, estoy usando `cowsay`, `figlet`, `lolcat` y `neofetch` para crear un MOTD personalizado basado en el tiempo que se muestra al usuario cuando se conecta por primera vez. Les saluda por su nombre, muestra información del servidor y la hora, fecha, tiempo e IP. [Aquí está el código fuente](https://github.com/Lissy93/dotfiles/blob/master/utils/welcome-banner.sh).

![image](https://sysarmy.com/blog/assets/50-neofetch-2.webp)

La mayoría de nosotros tenemos un conjunto básico de aplicaciones y utilidades CLI en las que confiamos. Configurar una nueva máquina e instalar cada programa individualmente sería muy tedioso. Por eso, la tarea de instalar y actualizar tus apps de terminal es perfecta para automatizar. [Aquí](https://github.com/Lissy93/dotfiles/tree/master/scripts/installs) tienes algunos scripts de ejemplo que he escrito, que puedes añadir fácilmente a tus dotfiles o ejecutar de forma independiente para asegurarte de que nunca te falte una app.

Para usuarios de macOS, el método más sencillo es usar [Homebrew](https://brew.sh/). Solo tienes que crear un Brewfile (con `touch ~/.Brewfile`), luego listar cada una de tus apps y ejecutar `brew bundle`. Puedes mantener tu lista de paquetes respaldada poniéndola en un repositorio Git. Aquí tienes un ejemplo para empezar: [https://github.com/Lissy93/Brewfile](https://github.com/Lissy93/Brewfile)

En Linux, normalmente querrás usar el gestor de paquetes nativo (por ejemplo, `pacman`, `apt`). Por ejemplo, [aquí tienes un script](https://github.com/Lissy93/dotfiles/blob/master/scripts/installs/arch-pacman.sh) para instalar las apps anteriores en sistemas Arch Linux.

Las aplicaciones de escritorio en Linux pueden gestionarse de forma similar, usando Flatpak. De nuevo, [aquí tienes un script de ejemplo](https://github.com/Lissy93/dotfiles/blob/master/scripts/installs/flatpak.sh) :)

## Conclusión

Y eso es todo: una lista de útiles aplicaciones CLI y un método para instalarlas y mantenerlas actualizadas en todos tus sistemas.

Espero que algunas de estas te sean útiles :)

¡Me encantaría saber cuáles son tus aplicaciones CLI favoritas! Déjalo en los comentarios abajo.

### Información adicional

- Esta lista no incluye lo básico, como Vim, Tmux, Ranger, ZSH, Git, etc., que probablemente ya usas
- Tampoco he incluido nada demasiado específico o solo relevante para pocos usuarios
- No se incluye nada específico de un sistema o que no sea multiplataforma (Linux/Unix, macOS)
- Tampoco apps relacionadas con la terminal pero que no sean CLI (como emuladores de terminal)
- Para la mayoría de los proyectos listados, hay muchas alternativas que logran cosas similares; por brevedad, esas tampoco se incluyeron

Un gran reconocimiento a los autores y comunidades detrás de cada una de estas apps. Sin ellos y su trabajo, nuestra vida en la línea de comandos sería mucho menos genial. Donde ha sido posible, he intentado dar crédito a los autores, pero si me he olvidado de alguno, avísame abajo y haré una actualización.

¿Qué me he perdido? Me encantaría conocer tus aplicaciones CLI favoritas, ¡especialmente si hay alguna increíble que no he incluido!

También me gustaría conocer tus opiniones y sugerencias; siempre busco mejorar :)

### Más información

Si te ha gustado esto, te recomiendo que eches un vistazo también a:

Si eres nuevo en la línea de comandos, [**The Art of Command Line**](https://github.com/jlevy/the-art-of-command-line), de Joshua Levy, es un recurso excelente, al igual que [**Bash Guide**](https://github.com/Idnan/bash-guide), de Adnan Ahmed.

Y si buscas inspiración, te encantará [**r/unixporn**](https://www.reddit.com/r/unixporn/) ⚡

- La versión original de este artículo se puede encontrar en [dev.to/lissy93](https://dev.to/lissy93/cli-tools-you-cant-live-without-57f6) (inglés).
