---
title: "Me he comprado un NAS para dejar de depender de Google Photos y Dropbox: configuración paso a paso"
notion_id: 00c9ebc2-6d43-4b80-b297-a1404ab9ec39
notion_url: https://app.notion.com/p/Me-he-comprado-un-NAS-para-dejar-de-depender-de-Google-Photos-y-Dropbox-configuraci-n-paso-a-paso-00c9ebc26d434b80b297a1404ab9ec39
last_edited: 2023-02-09T01:26:00.000Z
source_url: https://www.xataka.com/otros-dispositivos/me-he-comprado-nas-para-dejar-depender-google-photos-dropbox-configuracion-paso-a-paso-1
tags: ["Español", "Hosting", "Article", "Xataka | Genbeta"]
---
2020 fue un año muy movido en lo tecnológico. La pandemia hizo que **empresas tecnológicas como Zoom triunfaran** de un día para otro con fortísimas subidas en bolsa mientras que el resto del mercado se hundía. Sin embargo, si hubo una noticia relevante a nivel tecnológico que define el 2020, curiosamente no tiene mucho que ver con la pandemia: [Google Fotos anunción que dejaría de ser gratis en junio de 2021](https://www.xataka.com/aplicaciones/google-fotos-golpe-para-quienes-llevaban-anos-subiendo-su-vida-fotos-a-nube).

Este movimiento de Google fue, para mí, el aleteo de una mariposa en mi _statu quo_ digital. A lo largo de los años había ido creándome un ecosistema de servicios en la nube y offline para tener mis datos a buen recaudo, pero **este movimiento de Google hizo que me replanteara todo**. Tanto que decidí sustituir estos servicios por un NAS, un dispositivo en red que almacene todos mis datos: una nube privada.

## Mi escenario antes del NAS

Antes del anuncio de Google Fotos tenía un verdadero follón de mis datos digitales. Y el motivo era, como no, la historia. Igual que la historia hace que sigan existiendo las motos a pesar de lo peligrosas que son para sus pasajeros o que en EEUU siga siendo legal comprar armas, la historia justificaba mi caos digital (o al menos tranquilizaba mi conciencia).

Hace como diez años que me regalaron una Raspberry Pi. Después de trastear un poco con ella decidí usarla como **media center**. Para ello le instalé una distribución de GNU/Linux con Kodi y le conecté un disco duro de 1 TB, que acabó evolucionando con los años a 3 TB.

La idea inicial era usarlo no solo como media center sino también como **repositorio para copias de seguridad**. Sin embargo era algo incómodo, lo intenté con rsync, samba, etc. pero al final había que hacer tareas de forma manual, tenía también datos importantes en el móvil (¡las fotos!) y al final quedó exclusivamente para media center.

**Por el camino apareció Dropbox**, que permitía tener un repositorio para datos en la nube y además permitía hacer un backup automático para fotos desde el móvil. Estaba muy bien e hice todos los trucos necesarios para tener más de los 2 GB que venían de forma gratuita. Sin embargo el almacenamiento se acabó llenando y empecé a pagar (primero 100 euros al año, pero luego lo subieron a 120 euros al año).

Al final convivía con tres ecosistemas distintos: la Raspberry Pi para películas y series, Dropbox para datos y Google Fotos para fotos y vídeos

Ya que pagaba y tenía 2 TB de espacio lo empecé a usar de forma bastante frecuente para backup de datos importantes, no solo para las fotos. El PC se podía sincronizar con una aplicación sencilla y proporcionaba la ventaja de poder acceder a esos datos desde el móvil, el tablet u otro PC. **Era realmente cómodo**.

Entonces [llegó Google Fotos](https://www.xataka.com/aplicaciones/fotos-de-google-el-gran-reto-de-organizar-una-vida-de-fotos), con **su increíble aplicación para el móvil**. Esta proporcionaba también backup de fotos pero sobre todo una aplicación móvil para usarlas. En Dropbox la subida automática era limitada, permitía hacer backups, sí, pero no visualizar las fotos y compartirlas como con Google. Y lo que es mejor, Google Fotos llegó con almacenamiento ilimitado gratuito (con una pérdida despreciable de calidad).

Al final convivía con tres ecosistemas distintos: la Raspberry Pi para películas y series, Dropbox para datos y Google Fotos para fotos y vídeos. Y convivían en un equilibro estable, sin incidentes. **Pero entonces Google decidió acabar con la gratuidad**.

## Pagar cada vez más por mi _status quo_

Al empezar a a cobrar por almacenamiento Google estaba poniéndome otra vez en una situación por la que ya había pasado: cuando Dropbox decidió subir el precio de 100 a 120 euros ya me planteé cambiar cosas, pero o **no investigué suficiente o no tenía ánimos para hacer cambios en aquel momento**.

Y, de hecho, cuando Google anunció el fin de la gratuidad, me planteé pagar. Al final los 100 GB de Google cuestan 20 euros al año, no es mucho, que además podría compartir con mi mujer con la cuenta familiar. **Pero estoy seguro de que esos 100 GB se acabarían quedando cortos; y quizá Dropbox podría subir precios otra vez**. Entonces lo lógico sería migrar todo a Google. Es decir, los cambios que no quise acometer en el pasado iban a llegar quisera o no. Es entonces cuando escuché [un podcast sobre los NAS](https://mixx.io/2020/12/30/como-tener-tu-propio-google-fotos-dropbox-y-spotify-en-tu-casa/) y me abrió los ojos.

**Un NAS es un pequeño ordenador en red** con capacidad para alojar internamente varios discos duros. Originalmente estaban muy pensados para tener datos, pero en la actualidad ofrecen muchos servicios: media center para películas y series, servicio de música en la nube (privada), realizar descargas por torrent, servicio de fotos, instalación de aplicaciones de terceros... Era precisamente lo que necesitaba.

## ¿Qué NAS elegir?

Como hemos comentado, un NAS no es solo almacenamiento así que a la hora de elegir uno debía adaptarse a mis necesidades. Por un lado tenía que sustituir a mi Raspberry Pi, es decir, actuar de media center pero también gestionar las descargas. Por otro tenía que servir de backup del ordenador. Y por último, gestionar las fotos. **Y por supuesto que tuviera aplicaciones para el móvil**.

Aunque hay muchas marcas de NAS, muchas se centran en la parte de backups. Hay dos destacan por su calidad y los múltiples frentes que pretenden cubrir: **Synology y QNAP**.

Algunos dicen que QNAP es el Microsoft de los NAS, por su versatilidad, y que Synology es el Apple, por sus acabados. Mirando las comparativas efectivamente QNAP ofrece un poco más por el mismo precio y su ecosistema permite configurar más allá que en Synology, pero realmente yo no quería nada que se saliera de lo anterior.

Examinando bien los dos **me gustó más Synology** por sistema operativo, más sencillo de usar y con aplicaciones móviles que parecen más depuradas. Rápidamente me centré en esta marca.

Lo primero destacable es que el NAS tenía que cubrir el principal uso por el que no me quedaba solo con Dropbox: la aplicación de fotos. Y en este caso Synology acababa de lanzar una nueva aplicación, **Synology Photos**, que aunaba dos aplicaciones anteriores (Moments y Photo Station). Synology Photos permite, al igual que con Google Fotos, tener un aplicación en el móvil y en el escritorio (vía web) con todas las fotos, subida automática, crear álbumes, reconocimiento de caras, etc.

Una vez tomada la decisión de la marca a considerar había que elegir modelo, ya que hay muchos. **Para un entorno doméstico lo normal es centrarse en los NAS de dos bahías**, ya que son más económicos y no estamos en un entorno profesional donde hay que ampliar capacidad de forma habitual.

Dentro de los modelos de dos bahías hay unas cuantas opciones, pero yo quería **una versión suficientemente potente para hacer transcoding de vídeo** (por ejemplo, si lo uso de media center para ver películas en el móvil, en lugar de transmitirlas en 4K las puede mandar en 720 para ahorrar datos). También quería tener soporte para docker, ya que así permite instalarse múltiples aplicaciones para GNU/Linux. Eso sí, la versión de docker es algo antigua, esperemos que la actualicen pronto porque algunos paquetes empiezan a dar problemas. El NAS de Synology que cumplía todos los requisitos es el [DS200+](https://www.xataka.com/redirect?url=https%3A%2F%2Fwww.amazon.es%2FSynology-DS220-Caja-Almacenamiento-Red%2Fdp%2FB08BS4PW1B%2F&category=otros-dispositivos).

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Cuando se compra un NAS suele venir sin discos duros. Estos se compran aparte. Y es otra decisión a considerar. ¿Uno o dos discos duros? **¿Especiales para NAS?** Si se eligen dos discos, ¿qué tipo de RAID elegimos?

En mi caso decidí tener **una capacidad total de 6 TB**, porque ya tenía unos 3 TB de datos, quería tener espacio para crecer y también porque iba a compartir este almacenamiento con mi familia.

Sobre los discos duros, Synology recomienda usar solo los que están aprobados y lo lógico es usar discos especiales para NAS ya que están usándose de forma continua. Aunque tenía mis dudas (pues no le voy a dar un uso empresarial) finalmente decidí usar discos duros especiales para NAS, concretamente **Seagate Ironwolf NAS**.

**Pero quedaba la duda del RAID**. RAID es la forma de combinar los discos duros que [hemos explicado en detalle en alguna ocasión](https://www.xataka.com/basics/raid-discos-duros-que-sus-principales-tipos). En mi caso la duda era entre RAID 0 y RAID 1. RAID 0 básicamente combina los discos duros (para tener 6 TB necesitaría dos discos de 3 TB), pero si un disco duro falla los datos que estén ahí se pierden. RAID 1 replica los datos (para tener 6 TB necesitas tener dos discos de 6 TB), lo que da seguridad.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Una de las cosas que me hacía dormir bien cuando tenía todos los datos importantes en Dropbox y Google Fotos es que sabía que no iba a pasar nada. Ellos se encargan de replicar los datos y tener sus propios RAID para evitar pérdidas por fallos de hardware. Como quiero seguir durmiendo bien, **decidí usar RAID 1**, aunque fuera algo más caro. Luego veremos, en la configuración, que elegí SHR, [un RAID específico que Synology](https://nascompares.com/2016/07/06/what-is-shr-and-what-is-the-difference-between-synology-hybrid-raid-and-ordinary-raid/) con alguna ligera ventaja.

En definitiva, mi compra fue de un [Synology DS220+](https://www.xataka.com/redirect?url=https%3A%2F%2Fwww.amazon.es%2FSynology-DS220-Servidor-bah%25C3%25ADas-Discos%2Fdp%2FB08DKZ4GQ9%2Fref%3Dsr_1_1%3F__mk_es_ES%3D%25C3%2585M%25C3%2585%25C5%25BD%25C3%2595%25C3%2591%26dchild%3D1%26keywords%3DSynology%2BDS220%252B%26qid%3D1617192086%26sr%3D8-1&category=otros-dispositivos) y dos discos duros [Seagate Ironwolf NAS 6 TB](https://www.xataka.com/redirect?url=https%3A%2F%2Fwww.amazon.es%2FSeagate-ST6000VN001-IronWolf-NAS-Disco%2Fdp%2FB083JBZC5K%2Fref%3Dsr_1_3%3F__mk_es_ES%3D%25C3%2585M%25C3%2585%25C5%25BD%25C3%2595%25C3%2591%26dchild%3D1%26keywords%3Dseagate%2Bironwolf%2Bnas%2B6tb%26qid%3D1617192235%26sr%3D8-3&category=otros-dispositivos). Total 685 euros, que espero amortizar algo menos de cinco años dejando de pagar Dropbox y no pagando Google Fotos.

## La configuración: primeros pasos

Nada más llegar el DS220+ monté los discos duros y lo enchufé. Es bastante sencillo, **los discos duros se colocan sobre una carcasa de plástico y se introducen en las ranunras**. Luego el NAS se conecta al router de casa por Ethernet y se enciende.

El control del NAS se hace desde un PC a través de un navegador web. Hay dos formas de hacerlo: o bien ponemos `http://synologynas:5000` en el navegador o si sabemos la IP que ha cogido del router, entonces `http://direccion_ip:5000`.

Lo primero que veremos a continuación es una pantalla de inicio que nos invita a configurar el NAS. Le damos a configurar y **comenzará a instalar el sistema operativo en los discos duros**.

El siguiente es ponerle nombre al servidor y poner nombre y contraseña al usuario principal, que además tendrá permisos de administrador. **Es importante que la contraseña sea segura porque tendrá acceso desde Internet**.

El siguiente paso es **configurar el QuickConnect**. Este es un servicio gratuito que ofrece Synology para evitar que haya que estar configurando el reenvío de puertos en el router. Gracias a QuickConnecto tendremos una dirección web que podemos usar cuando estemos fuera de casa para acceder al NAS. Aparte de cómodo es muy conveniente cuando usamos las aplicaciones del móvil, ya que podemos usar el alias de QuickConnect para acceder al NAS estemos dentro o fuera de casa, de lo contrario habría que saber la IP pública que tenemos en cada momento si el operador nos la cambia y estar cambiando por la IP privada si estamos dentro casa.

Una vez terminada la configuración llegamos a la pantalla de inicio, que tiene pinta de ser **un sistema operativo completo**, pero en realidad es una sesión web. Lo llaman DSM (DiskStation Manager) y va muy fluido.

Lo siguiente que hace que hacer es actualizar el sistema operativo a la última versión, por motivos de seguridad. Una vez hecho esto y estando en la versión 7 de DSM, **el aspecto es bastante más moderno**, lo cual es de agradecer.

Por último, y antes de empezar a usar de verdad el NAS, hay que **configurar el RAID de los discos duros**. Para ello hay que ir al menú principal -> Administrador de almacenamiento y elegir lo que queremos hacer con las unidades. En mi caso fue crear un volumen de almacenamiento con ambos discos duros en SHR con tolerancia a fallos de 1 unidad. El sistema de archivos recomendado es Btrfs y es el que usé.

Una vez formateados los discos duros, es conveniente probar si hay fallos. Eso se hace realizando una prueba S.M.A.R.T., desde el mismo Administrador de almacenamiento. Y menos mal que lo hice pues en uno de los discos **había sectores defectuosos que además iban aumentando**, tuve que cambiarlo por uno nuevo.

Lo bueno de tener los discos en RAID es que se puede perfectamente quitar el disco duro defectuoso, cambiarlo y poner el nuevo en caliente sin dejar de trabajar con el NAS. Cuando faltaba un disco duro salía un aviso de que no había protección disponible pero cuando le puse el nuevo, lo añadí al grupo de almacenamiento y se regeneraron los datos en este segundo disco sin tener que hacer nada más.

Si alguna vez decido ampliar la capacidad es muy sencillo de hacer: basta con comprar dos discos de mayor tamaño. Primero hay que sustituir uno y esperar a que los datos se regeneren. Una vez terminado hay que cambiar el segundo y de nuevo se regenarán los datos en él. Así se puede **ampliar el almacenamiento sin interrumpir el servicio y sin perder configuraciones**.

## Copiando los datos al NAS

Una vez completamente configurado lo más básico del NAS era el momento de empezar a usarlo. Antes de nada se puede ver que la principal aplicación para moverse por los archivos es File station. Es un navegador de archivos muy intuitivo que permite arrastrar para copiar o mover ficheros. **Los datos normalmente se colocan en la carpeta home**, pero hay algunas como video o music que están fuera de home (normalmente para que puedan acceder todos los usuarios). Los administradores, además, pueden ver una carpeta llamada homes, en las que cuelgan todas las carpetas home de los usuarios.

**Los archivos y carpetas se pueden compartir**, creando un enlace que cualquiera con acceso a Internet pueda descargar. Este enlace puede tener una validez (por ejemplo, 10 días), puede ser completamente público, protegido por contraseña o incluso puede ser necesario tener un usuario en el NAS para poder descargarse. Y, además, se puede limitar el número de veces que se puede descargar.

Una vez ya había trasteado un poco con mi nuevo NAS lo primero que iba a hacer era traerme todos los datos de Dropbox. Para ello basta con ir al Centro de paquetes dentro de DSM, e instalar **Cloud sync**, que permite sincronizar el NAS con múltiples servicios en la nube (en uno o dos sentidos).

Ya con todos los datos de Dropbox en el NAS llegó el momento de poner a funcionar **Synology Photos**. Para ello hay que instalar la aplicación desde el Centro de paquetes y copiar las fotos en la carpeta Photos del directorio home.

Es entonces cuando Synology Photos empieza a **indexar los datos** y puede tardar bastante. También se puede instalar la aplicación en el móvil (tanto para Android como para iOS), donde se puede habilitar la subida automática.

El siguiente paso fue instalar el Video station, también desde el Centro de paquetes. Esta aplicación permite convertir el NAS en un media center. Solo hay que indicarle las carpetas donde están las películas y las series y él se encarga de buscarles carátulas y la información (basada en el nombre del archivo) y es como si tuvieras un Netflix directamente en el navegador, móvil o tablet (hay aplicaciones para Android e iOS). **Para ver los contenidos en una televisión se usa DNLA**, cosa que soportan las SmartTV, Chromcast y otros dispositivos conectados a la televisión.

Dentro del DS220+ también está la aplicación de Plex, por si se prefiere esta a Video station, pero por defecto no soporta la transcodificación, [hay que toquetear la configuración](https://nascompares.com/2020/07/30/synology-nas-plex-driver-fix-for-h-256-hevc-ds920-ds220-ds720-and-ds420/). No es muy complicado pero de momento prefiero ceñirme a Video Station.

Download station es una aplicación que permite **descargar desde torrents, páginas web o incluso emule**. Es muy sencilla de usar y por supuesto tiene aplicación móvil para controlarla de forma remota.

Por último configuré un par de aplicaciones extra que me parecieron interesantes. Por un lado Audio Station, que permite escuchar la música en streaming. Le indicas las carpetas donde está la música y se puede escuchar desde el PC o el móvil desde cualquier sitio, estilo Spotify.

También me instalé Note station, un clon de Evernote. Normalmente usaba bastante Evernote en su tramo gratuito, pero las limitaciones a veces se me hacían cuesta arriba aunque no tanto como para pagar. **Ahora con Note station tengo todas mis notas en varios dispositivos sin problema** (y tiene una opción para importar las libretas de Evernote así que el cambio ha sido muy limpio).

## Backup de los datos

Por último quedaba un detalle. Aunque tener dos discos duros en RAID 1 me da bastante tranquilidad siempre queda la posibilidad de perder los datos de otras formas: robos, incendio, etc. Estamos hablando ya de cosas bastante remotas pero sinceramente, al menos las fotos merecen un backup extra.

Como tenía cuenta de Dropbox por unos meses seguí usando Cloudsync para sincronizar todo allí. Pero después de eso [me decanté por Backblaze](https://www.xataka.com/otros-dispositivos/asi-puedes-hacer-backup-nas-nube-que-precio-sea-problema), con un coste de 0,005 dólares por GB al mes. Los backups los hago también con Cloudsync (usando el servicio S3) pero por la noche únicamente.

## Cuentas para familiares

Otro de los motivos por los que me decidí por el NAS, aparte de que en costes se amortiza relativamente rápido, es que se pueden crear varias cuentas y podría lograr que mi familia, que usaba Google Photos, **no tuviera que pagar cuando se les llenara el almacenamiento**.

Dentro de Panel de control -> usuarios, se pueden crear nuevos usuarios, con su correspondiente login y password. Estos usuarios **pueden tener cuota de espacio y capacidad para acceder a todas, alguna o ninguna de las aplicaciones instaladas**.

Además, para simplificar su uso, les podemos denegar permiso para acceder al DSM, que quizá pueda complicar su uso, y que entonces **solo puedan acceder a determinadas aplicaciones** (el File Station, el Video station, Synology Photos, etc.).

Lo primero que hice fue crear una cuenta a mi mujer y **pasar todas las fotos que tenía en Google Photos a Synology Photos** (ya que ella no tenía copia de seguridad en Dropbox). Para ello, una vez creada su cuenta, descargué las fotos desde Google Takeout (hace no mucho sacamos [una guía detallada](https://www.xataka.com/basics/como-descargar-recuperar-todas-fotos-google-fotos)) y las puse en su carpeta Photos. Una vez indexadas, ya puede acceder a las fotos desde un PC o desde la aplicación móvil.

También creé cuentas para mis padres y suegros y a ellos les simplifiqué bastante el uso, dejándoles acceder solo a las aplicaciones más relevantes. Además creé alias para que fuera más sencillo. Esto permite que accedan directamente a las aplicaciones desde un navegador web. Por ejemplo, para acceder a File station basta con que pongan en el navegador `http://QuickConnect.to/miquickconnect/file`, y para acceder a Synology Photos `http://QuickConnect.to/miquickconnect/foto`. Si acceden desde aplicaciones móviles no hace falta, ya que cada aplicación de Synology se conecta exclusivamente a dicho servicio. Solo tienen que poner el quickconnect que tiene mi NAS, su usuario y su contraseña.

## Las aplicaciones móviles y de escritorio

Una de las cosas que más me gusta de Synology es **su gran ecosistema de aplicaciones**. Prácticamente cada aplicación que tiene el NAS tiene su correspondiente acceso desde el móvil o tablet, y eso es muy útil.

Por ejemplo, para acceder a los archivos tenemos DS file, que se parece bastante a lo que ofrece Dropbox. Con DS video podemos acceder a Video station como si tuviéramos nuestra aplicación de Netflix. DS finder permite gestionar el NAS, e incluso acceder al escritorio DSM desde el móvil. **Y Synology Photos es muy parecida a Google Fotos**.

**También existen aplicaciones para Windows, Mac y GNU/Linux**. Por ejemplo hemos comentado que Note station [tiene aplicaciones](https://www.synology.com/es-es/dsm/software_spec/note_station) no solo para móvil. Basta con ir a la página de Synology y [explorar los paquetes](https://www.synology.com/es-es/dsm/packages) para ver lo que tiene disponibles en escritorio.

## Otras aplicaciones que todavía no uso

El centro de paquetes de Synology es apasionante, hay un montón de aplicaciones que pueden **simplificar mucho la vida de los usuarios**. Hay unas cuantas que las tengo localizadas pero todavía no me he animado a dar el salto.

Existe una aplicación llamada [Synology Drive Server](https://www.synology.com/es-es/dsm/software_spec/synology_drive) que es muy parecida a lo que ofrece Dropbox o Google Drive en escritorio, **sincronización en tiempo real de archivos**. Pero está muy enfocado a equipos, archivos compartidos, versionado y no acaba de convencerme, de momento prefiero usar File station para tener backup de mis archivos y hacerlos manualmente o por rsync.

También está [Synology Calendar](https://www.synology.com/es-es/dsm/packages/Calendar), por si alguna vez quiero **librarme de la tiranía de Google**. Esto, junto a varias aplicaciones de [correo electrónico](https://www.synology.com/es-es/dsm/packages?group=All&search=mail) podrían hacer que diera el paso definitivo, pero tener tu propio servidor de correo siempre da problemas (por temas de reconocimiento de spam).

Otra aplicación interesante es el [Synology Chat Server](https://www.synology.com/es-es/dsm/packages/Chat), una especie de slack o Microsoft Teams completamente privado, con aplicaciones para escritorio y móvil.

Y por supuesto se puede poner [un servidor web](https://www.synology.com/es-es/dsm/packages/WebStation) o directamente [un Wordpress](https://www.synology.com/es-es/dsm/packages/WordPress) en el NAS. **Al final el DS220+ es un ordenador con GNU/Linux y una capa muy potente por encima**.

En cada paquete se puede ver qué modelos de NAS están soportados. Como dije al principio el DS220+ es bastante potente y soporta casi todo, pero **algunos modelos más baratos no tienen esa capacidad**.

En definitiva, estoy bastante contento con el NAS, con el que llevo ya un año. Tengo un setup parecido a lo que tenía antes, sin riesgos de tener que aumentar los costes, con mis datos más controlados y **con muchas posibles extensiones que hacer en el futuro**.
