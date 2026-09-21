---
title: "Cómo configurar Nginx para obtener un rendimiento óptimo"
notion_id: af8f207e-12ba-4067-a3df-e5984b3ae14e
notion_url: https://app.notion.com/p/C-mo-configurar-Nginx-para-obtener-un-rendimiento-ptimo-af8f207e12ba4067a3dfe5984b3ae14e
last_edited: 2026-09-21T17:12:00.000Z
source_url: https://www.notion.so/af8f207e12ba4067a3dfe5984b3ae14e
tags: ["Español", "Web Development", "SysAdmin", "Article"]
---
El servidor web Nginx se ha popularizado en los últimos años tanto por su alto desempeño en el balance de carga y el almacenamiento en caché de contenido web dinámico y estático. Esta guía pretende proporcionar algunos consejos para determinar las mejores optimizaciones de rendimiento necesarias en un servidor Nginx para acelerar la entrega de contenido a sus usuarios finales.

## Modificaciones del Worker

Una de las modificaciones más sencillas que puede hacer en su configuración es establecer el número adecuado de conexiones y _workers_.

### Procesos Worker

Si usted tiene un tráfico bajo en su sitio web y además nginx, su base de datos y la aplicación web todos se ejecutan en el mismo servidor entonces abra el archivo `/etc/nginx/nginx.conf` y establezca el valor:

```plain text
worker_processes 1;

```

Si en cambio tiene un tráfico alto o una instancia dedicada para Nginx, fije un _worker_ por cada core del CPU:

```plain text
worker_processes auto;

```

Si desea configurar esto de forma manual, puede usar `grep ^processor /proc/cpuinfo | wc -l` para encontrar el número de procesos que el servidor puede manejar.

### Conexiones Worker

La opción `worker_connections` establece el número máximo de conexiones que cada proceso _worker_ puede procesar a la vez. Por defecto, el límite de este valor es 512 pero muchos sistemas pueden manejar valores más altos.

El tamaño apropiado puede ser descubierto a través de pruebas y es variable con base en el tipo de tráfico que maneje Nginx. Las limitaciones en el core del sistema también pueden encontrarse usando la utilidad `ulimit`:

```plain text
ulimit -n

```

Esto imprimirá un número como resultado:

```plain text
65536

```

También puede definir `use epoll`, un mecanismo escalable de notificación de eventos de Entrada/Salida (I/O) que se activa en la ocurrencia de ciertos eventos, asegurándose de que el I/O es utilizado en su mejor capacidad.

Por último, puede usar `multi_accept` con el fin de que un _worker_ acepte todas las nuevas conexiones al mismo tiempo.

La función `events` debería lucir similar a la siguiente cuando se configure:

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
events {
    worker_connections 66536;
    use epoll;
    multi_accept on;
}

```

## Optimizaciones HTTP y TCP

### KeepAlive

La opción _Keep Alive_ (mantener vivo) permite un menor número de reconexiones desde el navegador.

- Los parámetros `keepalive_timeout` y `keepalive_requests` controlan los ajustes _keep alive_.
- `sendfile` optimiza la entrega de archivos estáticos desde el sistema de archivos, por ejemplo: logotipos.
- `tcp_nodelay` permite que Nginx haga que TCP envíe varios búferes como paquetes individuales.
- `tcp_nopush` optimiza la cantidad de datos enviados en línea a la vez, activando la opción `TCP_CORK` dentro de la infraestructura TCP. `TCP_CORK` bloquea los datos hasta que el paquete haya alcanzado el MSS, lo que es equivalente al MTU menos los 40 o 60 bytes de la cabecera IP.

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
keepalive_timeout 65;
keepalive_requests 100000;
sendfile on;
tcp_nopush on;
tcp_nodelay on;

```

### Tamaño del búfer

Hacer algunos ajustes al tamaño del búfer puede ser muy ventajoso. Si los tamaños del búfer son demasiado bajos, entonces Nginx escribirá en un archivo temporal. Esto es causado por un exceso de I/O en el disco.

- El parámetro `client_body_buffer_size` controla el tamaño del búfer del cliente. La mayoría de los búferes de los clientes vienen del método POST en los envíos de formularios. `128k` es generalmente un buen valor para esta opción.
- `client_max_body_size` establece el tamaño máximo del cuerpo del búfer. Si el tamaño en una solicitud excede el valor configurado en este parámetro, se le devolverá el siguiente error al cliente: 413 (Request Entity Too Large —la entidad solicitada es demasiado grande). Para su referencia, los navegadores no pueden mostrar los errores 413 correctamente. Establecer el tamaño en 0 deshabilita el chequeo de la solicitud del tamaño del cuerpo de búfer del cliente.
- `client_header_buffer_size` maneja el tamaño de la cabecera asignada al cliente. En general, `1k` es una selección bastante sana para este valor.
- `large_client_header_buffers` muestra el número máximo y el tamaño de los búferes para cabeceras más grandes de clientes. 4 cabeceras con búferes de `4k` deberían ser suficientes aquí.
- `output_buffers` establece el número y el tamaño de los búferes utilizados para leer una respuesta en disco. Si es posible, la transmisión de los datos del cliente será pospuesta hasta que Nginx tenga establecido al menos el tamaño de bytes de datos a enviar. El valor cero desactiva el aplazamiento en la transmisión de datos.

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
client_body_buffer_size      128k;
client_max_body_size         10m;
client_header_buffer_size    1k;
large_client_header_buffers  4 4k;
output_buffers               1 32k;
postpone_output              1460;

```

### Cola de conexión

Algunas directivas en el archivo `/etc/sysctl.conf` pueden ser cambiadas para establecer el tamaño de una cola de Linux para las conexiones y los _buckets_. Actualizar los parámetros `net.core.somaxconn` y `net.ipv4.tcp_max_tw_buckets` cambia el tamaño de la cola para las conexiones esperando ser aceptadas por Nginx. Si hay mensajes de error en el log del kernel, aumente el valor hasta que el error desaparezca.

Extracto del archivo: **/etc/sysctl.conf**

```plain text
net.core.somaxconn = 65536
net.ipv4.tcp_max_tw_buckets = 1440000

```

Los paquetes pueden ser almacenados en el búfer de la tarjeta de antes de ser enviados al CPU. Esto se logra estableciendo la reserva máxima con la etiqueta `net.core.netdev_max_backlog`. Consulte la documentación de la tarjeta de red en busca de consejos sobre el cambio de este valor.

### Timeouts

Los tiempos de espera o _timeouts_ puede mejorar el rendimiento drásticamente.

- La opción `client_body_timeout` envía directivas por el tiempo en el que un servidor esperará que un cuerpo sea enviado.
- `client_header_timeout` envía directivas por el tiempo en el que un servidor esperará que una cabecera sea enviada. Estas directivas son responsables por el tiempo en el que un servidor esperará que un cuerpo o cabecera del cliente sea enviado después de una solicitud. Si no se envía ningún cuerpo o cabecera, el servidor emitirá un error 408 "Tiempo de espera de solicitud agotado" (_Request time out_).
- `sent_timeout` especifica el tiempo de espera de la respuesta al cliente. Este _timeout_ no aplica a la transferencia completa, más bien mide únicamente dos operaciones sucesivas de lectura en el cliente.

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
client_header_timeout  3m;
client_body_timeout    3m;
send_timeout           3m;

```

### Entrega de recursos estáticos

Si su sitio entrega recursos estáticos (tales como: CSS, JavaScript e imágenes), Nginx puede almacenar estos archivos en caché por un corto período de tiempo. Añadir esto dentro de su bloque de configuración le dice Nginx que almacene un máximo de 1000 archivos en caché durante 30 segundos, excluyendo los archivos que no hayan sido accedidos en 20 segundos y solo archivos que tengan al menos 5 accesos durante este tiempo. Si no va a implementar con frecuencia puede subir estos valores sin problemas.

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
open_file_cache max=1000 inactive=20s;
open_file_cache_valid 30s;
open_file_cache_min_uses 5;
open_file_cache_errors off;

```

También puede almacenar en caché a través de una ubicación particular. Almacenar archivos en caché por un largo tiempo es beneficioso, especialmente si los archivos tiene un sistema de control de versiones entregado por el proceso de generación o el CMS.

Extracto **/etc/nginx/nginx.conf**

```plain text
location ~* .(woff|eot|ttf|svg|mp4|webm|jpg|jpeg|png|gif|ico|css|js)$ {
    expires 365d;
}

```

### Comprimir contenido con gzip

Para el contenido que es texto plano, Nginx puede usar la compresión gzip para volver a entregar estos recursos al cliente, pero compresos. Los navegadores web modernos aceptarán compresiones gzip y esto recortará varios bytes en cada solicitud que venga en forma de recursos en texto plano. La lista a continuación es una lista "segura" de tipos de contenido comprimible; sin embargo, solo debe habilitar los tipos de contenido que esté utilizando dentro de su aplicación web.

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
gzip on;
gzip_min_length 1000;
gzip_types: text/html application/x-javascript text/css application/javascript text/javascript text/plain text/xml application/json application/vnd.ms-fontobject application/x-font-opentype application/x-font-truetype application/x-font-ttf application/xml font/eot font/opentype font/otf image/svg+xml image/vnd.microsoft.icon;
gzip_disable "MSIE [1-6]\.";

```

## Optimizaciones del sistema de archivos

Estas operaciones en el _filesystem_ mejorar la administración de memoria del sistema y pueden ser añadidas en el archivo `/etc/sysctl.conf`.

### Puertos efímeros

Cuando Nginx actúa como proxy, cada conexión en un servidor de _upstream_ usa un puerto temporal —o efímero.

El rango de puertos locales IPv4 define un valor para el rango de puertos. Una configuración común es: `net.ipv4.ip_local_port_range 1024 65000`.

El _timeout_ de TCP FIN asegura la cantidad de tiempo en el cual un puerto debe estar inactivo antes de ser reutilizado para otra conexión. El valor predeterminado suele ser 60 segundos, pero puede ser reducido a 30 o 15 segundos en la mayoría de los casos:

Extracto del archivo: **/etc/sysctl.conf**

```plain text
net.ipv4.tcp_fin_timeout 15

```

### Escalamiento de la ventana TCP

La opción `The TCP window` es una opción para aumentar el tamaño de la ventana de recepción permitida en el protocolo de control de transmisión o TCP por encima de su antiguo valor de 65.535 bytes. Esta opción TCP, junto con algunas otras, está definida en la hoja de especificaciones IETF RFC 1323, la cual se ocupa de grandes redes. Este valor puede ser definido con la etiqueta: `net.ipv4.tcp_window_scaling = 1`.

### Reservar paquetes antes de eliminarlos

El parámetro `net.ipv4.tcp_max_syn_backlog` determina un número de paquetes a mantener en el _backlog_ antes de que el kernel comience a descartarlos. Un buen valor sería: `net.ipv4.tcp_max_syn_backlog = 3240000`.

### Cerrar conexión tras la falta de respuesta del cliente

`reset_timedout_connection on;` permite que el servidor cierre una conexión después de que un cliente deja de responder. Esto libera la memoria asociada al _socket_.

### Descriptores de archivo

Los descriptores de archivo son recursos del sistema operativo que se usan para manejar cosas como las conexiones y abrir archivos. Nginx puede usar hasta dos descriptores de archivo por conexión. Por ejemplo, si está usado como proxy, hay generalmente un descriptor de archivo por conexión de cliente y otro para la conexión del servidor proxy, aunque esta tasa es mucho más baja si se usan _keep alives_ de HTTP. Para un sistema que sirva a un gran número de conexiones, es probable que sea necesario cambiar esta configuración.

`sys.fs.file_max` define el límite máximo de descriptores de archivo en el sistema. `nofile` define el límite de descriptores de archivo para el usuario. Ambos valores se establecen en el archivo `/etc/security/limits.conf`.

Extracto del archivo: **/etc/security/limits.conf** **

### Registro de errores

El parámetro `error_log logs/error.log warn` define la ubicación y los distintos niveles de severidad escritos en el archivo de registro o _log_ de error. El valor definido para el nivel de log regirá los niveles de registro de errores en este archivo: se registrarán desde el nivel descrito y todos los superiores. Por ejemplo: el nivel `default` de error causará que se registren los mensajes con niveles de severidad: `error`, `crit`, `alert` y `emerg`. Si este parámetro es omitido, entonces se usará `error` por defecto. Los niveles se enumeran a continuación:

- **`emerg`**: las situaciones de emergencia donde el sistema está en un estado no utilizable.
- **`alert`**: situaciones graves donde se requiere acción inmediata.
- **`crit`**: problemas importantes que deben ser abordados.
- **`error`**: un error ha ocurrido. Algo no se pudo completar con éxito.
- **`warn`**: ocurrió algo fuera de lo común, pero no debe causar demasiada preocupación.
- **`notice`**: Ha ocurrido algo normal, pero que vale la pena notar.
- **`info`**: un mensaje de información que podría ser útil.
- **`debug`**: información de depuración que puede ser útil para identificar dónde se está produciendo el problema.

Acceda a los _logs_ con la directiva `log_format` para configurar el formato de los mensajes registrados, así como con la directiva `access_log` para especificar la ubicación del log y su formato.

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
http {
    log_format compression '$remote_addr - $remote_user [$time_local] ' '"$request" $status $body_bytes_sent ' '"$http_referer" "$http_user_agent" "$gzip_ratio"';
    server {
        gzip on;
        access_log /spool/logs/nginx-access.log compression;
    }
}

```

### Registro condicional de eventos

Los _logs_ condicionales pueden ser completados si el administrador del sistema solo quiere registrar ciertas solicitudes. El ejemplo a continuación excluye el registro de eventos para los códios estatus 2XX y 3XX en HTTP:

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
map $status $loggable {
    ~^[23]  0;
    default 1;
}

```

### Desactivar el registro de logs completamente

El registro de logs puede ser desactivado completamente si tiene una metodología de _logs_ alternativa o si no le importa tener algún registro de eventos en su servidor. Desactivar el registro de eventos puede ser llevado a cabo con las siguientes directivas en el servidor:

Extracto del archivo: **/etc/nginx/nginx.conf**

```plain text
server {
    listen       80;
    server_name  example.com;
    access_log  off;
    error_log off;
}

```

### Supervisión de actividades

También se puede establecer la supervisión de actividades para ver respuestas JSON en tiempo real. Con la siguiente configuración, la página web `status.html` ubicada en `/usr/share/nginx/html` puede ser solicitada a través de la URL: `http://127.0.0.1/status.html`.

## Archivos de ejemplo

Después de hacer los ajustes sugeridos tendremos tres archivos para mejorar el desempeño de Nginx en su sistema. A continuación se muestran fragmentos completos a continuación.

### sysctl.conf

Extracto del archivo: **/etc/sysctl.conf**

```plain text
net.core.somaxconn = 65536
net.ipv4.tcp_max_tw_buckets = 1440000
net.ipv4.ip_local_port_range = 1024 65000
net.ipv4.tcp_fin_timeout = 15
net.ipv4.tcp_window_scaling = 1
net.ipv4.tcp_max_syn_backlog = 3240000

```

### limits.conf

Extracto del archivo: **/etc/security/limits.conf**

### nginx.conf

Extracto del archivo: **nginx.conf**

```plain text
pid /var/run/nginx.pid;
worker_processes  2;

events {
    worker_connections   65536;
    use epoll;
    multi_accept on;
}

http {
    keepalive_timeout 65;
    keepalive_requests 100000;
    sendfile         on;
    tcp_nopush       on;
    tcp_nodelay      on;

    client_body_buffer_size    128k;
    client_max_body_size       10m;
    client_header_buffer_size    1k;
    large_client_header_buffers  4 4k;
    output_buffers   1 32k;
    postpone_output  1460;

    client_header_timeout  3m;
    client_body_timeout    3m;
    send_timeout           3m;

    open_file_cache max=1000 inactive=20s;
    open_file_cache_valid 30s;
    open_file_cache_min_uses 5;
    open_file_cache_errors off;

    gzip on;
    gzip_min_length  1000;
    gzip_buffers     4 4k;
    gzip_types       text/html application/x-javascript text/css application/javascript text/javascript text/plain text/xml application/json application/vnd.ms-fontobject application/x-font-opentype application/x-font-truetype application/x-font-ttf application/xml font/eot font/opentype font/otf image/svg+xml image/vnd.microsoft.icon;
    gzip_disable "MSIE [1-6]\.";

    # [ debug | info | notice | warn | error | crit | alert | emerg ]
    error_log  /var/log/nginx.error_log  warn;

    log_format main      '$remote_addr - $remote_user [$time_local]  '
      '"$request" $status $bytes_sent '
      '"$http_referer" "$http_user_agent" '
        '"$gzip_ratio"';

    log_format download  '$remote_addr - $remote_user [$time_local]  '
      '"$request" $status $bytes_sent '
      '"$http_referer" "$http_user_agent" '
        '"$http_range" "$sent_http_content_range"';

    map $status $loggable {
        ~^[23]  0;
        default 1;
    }

    server {
        listen        127.0.0.1;
        server_name   127.0.0.1;
        root         /var/www/html;
        access_log   /var/log/nginx.access_log  main;

        location / {
            proxy_pass         http://127.0.0.1/;
            proxy_redirect     off;
            proxy_set_header   Host             $host;
            proxy_set_header   X-Real-IP        $remote_addr;
            proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
            proxy_connect_timeout      90;
            proxy_send_timeout         90;
            proxy_read_timeout         90;
            proxy_buffer_size          4k;
            proxy_buffers              4 32k;
            proxy_busy_buffers_size    64k;
            proxy_temp_file_write_size 64k;
            proxy_temp_path            /etc/nginx/proxy_temp;
        }

        location ~* .(woff|eot|ttf|svg|mp4|webm|jpg|jpeg|png|gif|ico|css|js)$ {
            expires 365d;
        }
    }
}

```

## Recursos adicionales

Puede consultar los siguientes recursos en busca de información adicional con respecto a este tema. Aunque este material es provisto esperando que sea útil, tenga en cuenta que no podemos dar fe de la actualidad o precisión de los contenidos externos.
