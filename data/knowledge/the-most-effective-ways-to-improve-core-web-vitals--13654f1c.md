---
title: "The most effective ways to improve Core Web Vitals"
notion_id: 13654f1c-7d23-81c2-85e2-fd1a281c0f08
notion_url: https://app.notion.com/p/The-most-effective-ways-to-improve-Core-Web-Vitals-13654f1c7d2381c285e2fd1a281c0f08
last_edited: 2024-11-15T20:13:00.000Z
source_url: https://web.dev/articles/top-cwv
tags: ["English", "Web Development", "Article", "web.dev"]
---
# Las formas más eficaces de mejorar las Métricas web esenciales



A lo largo de los años, la comunidad web ha acumulado una gran cantidad de conocimientos sobre la optimización del rendimiento web. Si bien cualquier optimización puede mejorar el rendimiento de muchos sitios, todas a la vez pueden ser abrumadoras y, en realidad, solo algunas de ellas se aplican a un sitio determinado.

A menos que el rendimiento web sea tu trabajo diario, es probable que no sea obvio qué optimizaciones tendrán el mayor impacto en tu sitio. Es probable que no tengas tiempo para todas, por lo que es importante que te preguntes _cuáles son las optimizaciones más impactantes que puedes elegir para mejorar el rendimiento de tus usuarios._

La verdad sobre las optimizaciones de rendimiento es que no puedes juzgarlas solo por sus méritos técnicos. También debes considerar los factores humanos y organizacionales que influyen en la probabilidad de que puedas implementar cualquier optimización determinada. Algunas mejoras en el rendimiento pueden tener un gran impacto en la teoría, pero, en realidad, pocos desarrolladores tendrán el tiempo o los recursos para implementarlas. Por otro lado, es posible que haya prácticas recomendadas de rendimiento de gran impacto que casi todos ya estén siguiendo. En esta guía, se identifican las optimizaciones de rendimiento web que:

- Tener el mayor impacto en el mundo real
- Son relevantes y se aplican a la mayoría de los sitios.
- Son realistas para que la mayoría de los desarrolladores los implementen.

En conjunto, estas son las formas más realistas y eficaces de mejorar tus métricas de [Métricas web esenciales](https://web.dev/articles/vitals?hl=es-419). Si es la primera vez que utilizas el rendimiento en la Web o si todavía no estás decidiendo qué te dará el mayor retorno de la inversión, este es el mejor punto de partida.

## Interaction to Next Paint (INP)

Como métrica más reciente de las Métricas web esenciales, [Interaction to Next Paint (INP)](https://web.dev/articles/inp?hl=es-419) tiene algunas de las mayores oportunidades de mejora. Sin embargo, como muchos menos sitios superan el [umbral](https://web.dev/articles/inp?hl=es-419#good-score) de experiencias "buenas" en comparación con su [predecesor obsoleto](https://web.dev/blog/fid?hl=es-419), es posible que seas uno de los muchos desarrolladores que aprenden a optimizar la capacidad de respuesta de la interacción por primera vez. Comienza con estas técnicas fundamentales para conocer las formas más eficaces de mejorar el INP.

### 1. Realiza una pausa con frecuencia para dividir tareas largas.

Las tareas son cualquier trabajo discreto que realiza el navegador, como la renderización, el diseño, el análisis, la compilación o la ejecución de secuencias de comandos. Cuando una tarea supera los 50 milisegundos de duración, se convierte en una [tarea larga](https://web.dev/articles/optimize-long-tasks?hl=es-419). Las tareas largas son problemáticas porque pueden impedir que el subproceso principal responda rápidamente a las interacciones del usuario.

Si bien siempre debes esforzarte por hacer el menor trabajo posible en JavaScript, puedes ayudar al subproceso principal [dividiendo las tareas largas](https://web.dev/articles/optimize-long-tasks?hl=es-419). Para ello, [cede el subproceso principal con frecuencia](https://web.dev/articles/optimize-long-tasks?hl=es-419#use_asyncawait_to_create_yield_points), de modo que las [actualizaciones de renderización](https://web.dev/articles/top-cwv?hl=es-419#inp-rendering) y otras interacciones del usuario puedan ocurrir antes.



La [API de Scheduler](https://web.dev/articles/optimize-long-tasks?hl=es-419#scheduler-api) te permite poner en cola el trabajo con un sistema de prioridades. Específicamente, la API de [scheduler.yield()](https://developer.mozilla.org/docs/Web/API/Scheduler/yield) divide las tareas largas y, al mismo tiempo, se asegura de que las interacciones se puedan controlar sin renunciar a su lugar en la cola de tareas.

**Si divides las tareas largas, le das al navegador más oportunidades para realizar tareas críticas que bloquean a los usuarios.**

**Lee para obtener más información:** [Optimiza las tareas largas](https://web.dev/articles/optimize-long-tasks?hl=es-419).

### 2. Evita el código JavaScript innecesario

[Los sitios web envían más JavaScript que nunca](https://httparchive.org/reports/state-of-javascript#bytesJs), y la tendencia no parece cambiar. Cuando envías demasiado JavaScript, creas un entorno en el que las tareas compiten por la atención del subproceso principal. Esto puede afectar la capacidad de respuesta de tu sitio web, especialmente durante ese período crucial de inicio.

Sin embargo, este no es un problema irresoluble y tienes opciones:

- Utiliza las funciones de [Baseline](https://web.dev/baseline?hl=es-419) disponibles en la plataforma web, en lugar de implementaciones redundantes basadas en JavaScript.
- Usa la [herramienta de cobertura](https://developer.chrome.com/docs/devtools/coverage/?hl=es-419) en Chrome DevTools para encontrar código sin usar en tus secuencias de comandos. Si reduces el tamaño de los recursos necesarios durante el inicio, puedes asegurarte de que las páginas dediquen menos tiempo a analizar y compilar código, lo que proporciona una experiencia del usuario inicial más fluida.
- Usa la [división de código](https://web.dev/articles/reduce-javascript-payloads-with-code-splitting?hl=es-419) para crear un paquete independiente para el código que no es necesario para la renderización inicial, pero que se usará más adelante.
- Si utilizas un administrador de etiquetas, [optimiza tus etiquetas](https://web.dev/articles/tag-best-practices?hl=es-419) periódicamente. Puedes quitar las etiquetas anteriores con código sin usar para reducir el espacio en JavaScript de tu Administrador de etiquetas.

**Lee para obtener más información:** [Cómo quitar el código que no se usa](https://web.dev/articles/remove-unused-code?hl=es-419).

### 3. Evita las actualizaciones de renderización grandes

La ejecución de JavaScript es solo un factor que afecta la capacidad de respuesta de tu sitio web. La renderización es un tipo de trabajo costoso en sí mismo y, durante las actualizaciones de renderización grandes, es posible que tu sitio web responda aún más lento a las interacciones de los usuarios.

La optimización del trabajo de renderización no es un proceso sencillo y depende de lo que intentes lograr. Aun así, estas son algunas medidas que puedes tomar para asegurarte de que las tareas de renderización no se conviertan en tareas largas:

- Reorganiza las lecturas y escrituras del DOM en tu código JavaScript para evitar el [diseño forzado](https://web.dev/articles/avoid-large-complex-layouts-and-layout-thrashing?hl=es-419#avoid_forced_synchronous_layouts) y la [paginación excesiva de diseños](https://web.dev/articles/avoid-large-complex-layouts-and-layout-thrashing?hl=es-419#avoid_layout_thrashing).
- [Mantén pequeños los tamaños del DOM](https://web.dev/articles/dom-size-and-interactivity?hl=es-419). El tamaño del DOM y la intensidad del trabajo de diseño están correlacionados. Cuando el renderizador tiene que actualizar el diseño de un DOM muy grande, el trabajo necesario para volver a calcular su diseño puede aumentar significativamente.
- [Usa la contención de CSS](https://web.dev/articles/content-visibility?hl=es-419) para renderizar de forma diferida el contenido del DOM fuera de la pantalla. No siempre es sencillo, pero si aislas las áreas que contienen diseños complejos, puedes evitar el trabajo de diseño y renderización innecesario.

**Lee para obtener más información:** [Evita los diseños grandes y complejos, y la paginación excesiva de diseños](https://web.dev/articles/avoid-large-complex-layouts-and-layout-thrashing?hl=es-419).

## Largest Contentful Paint (LCP)

El [Largest Contentful Paint (LCP)](https://web.dev/articles/lcp?hl=es-419) es la métrica web esencial con la que los desarrolladores suelen tener más problemas: el [40%](https://httparchive.org/reports/chrome-ux-report#cruxFastLcp) de los sitios del Informe sobre la experiencia del usuario en Chrome no cumple con el [umbral de LCP recomendado](https://web.dev/articles/lcp?hl=es-419#what_is_a_good_lcp_score) para brindar una buena experiencia del usuario. El equipo de Chrome recomienda las siguientes técnicas como las formas más eficaces de mejorar el LCP.

### 1. Asegúrate de que el recurso de LCP sea detectable desde la fuente HTML y tenga prioridad

El equipo de Chrome observó lo siguiente en relación con el LCP en la Web:

- Según el [2022 Web Almanac](https://almanac.httparchive.org/en/2022/) de HTTP Archive, el [72%](https://almanac.httparchive.org/en/2022/performance#fig-8) de las páginas para dispositivos móviles tienen una imagen como elemento de LCP.
- Un análisis de los datos de usuarios reales de Chrome muestra que la mayoría de los orígenes con un LCP deficiente dedican [menos del 10%](https://web.dev/blog/common-misconceptions-lcp?hl=es-419) de su tiempo de LCP del p75 a _descargar_ la imagen del LCP.
- Entre las páginas con un LCP bajo, la carga de sus imágenes de LCP se retrasa en el cliente en [1,290 milisegundos](https://web.dev/blog/common-misconceptions-lcp?hl=es-419#real_navigation_performance_data) en el percentil 75, lo que representa más de la mitad del presupuesto para una experiencia rápida.
- De las páginas en las que el elemento LCP era una imagen, el [39% de esas imágenes](https://almanac.httparchive.org/en/2022/performance#lcp-static-discoverability) tenía URLs de origen que no se podían detectar en la respuesta HTML inicial (como `<img src="...">` o `<link rel="preload" href="...">`), lo que permitiría que [el escáner de carga previa del navegador](https://web.dev/articles/preload-scanner?hl=es-419) las detectara lo antes posible.
- Según el Web Almanac, [solo el 0.03%](https://almanac.httparchive.org/en/2022/performance#lcp-prioritization) de las páginas aptas aprovechaba el atributo HTML `fetchpriority` para dar prioridad a los recursos, incluidos aquellos que podrían mejorar el LCP de una página con relativamente poco esfuerzo.

Estas estadísticas son reveladoras, ya que los desarrolladores tienen una gran oportunidad para reducir el [retraso en la carga de recursos](https://web.dev/articles/optimize-lcp?hl=es-419#lcp-breakdown) y la [duración de la carga de recursos](https://web.dev/articles/optimize-lcp?hl=es-419#reduce-resource-load-duration) para las imágenes de LCP.



![image](data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' width='24' height='24' viewBox='195 190 135 135'%3E%3Cdefs%3E%3ClinearGradient id='s-a' x1='132.6' x2='134.4' y1='111.7' y2='-105.3' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-color='%23d2d2d2' /%3E%3Cstop offset='.5' stop-color='%23f2f2f2' /%3E%3Cstop offset='1' stop-color='%23fff' /%3E%3C/linearGradient%3E%3ClinearGradient id='s-b' gradientUnits='userSpaceOnUse' /%3E%3ClinearGradient id='s-c' x1='65.4' x2='67.4' y1='115.7' y2='17.1' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-color='%23005ad5' /%3E%3Cstop offset='.2' stop-color='%230875f0' /%3E%3Cstop offset='.3' stop-color='%23218cee' /%3E%3Cstop offset='.6' stop-color='%2327a5f3' /%3E%3Cstop offset='.8' stop-color='%2325aaf2' /%3E%3Cstop offset='1' stop-color='%2321aaef' /%3E%3C/linearGradient%3E%3ClinearGradient id='s-d' x1='158.7' x2='176.3' y1='96.7' y2='79.5' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-color='%23c72e24' /%3E%3Cstop offset='1' stop-color='%23fd3b2f' /%3E%3C/linearGradient%3E%3CradialGradient id='s-i' cx='-69.9' cy='69.3' r='54' gradientTransform='matrix(.9 -.01 .04 2.72 -9 -120)' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-color='%2324a5f3' stop-opacity='0' /%3E%3Cstop offset='1' stop-color='%231e8ceb' /%3E%3C/radialGradient%3E%3CradialGradient id='s-j' cx='109.3' cy='13.8' r='93.1' gradientTransform='matrix(-.02 1.1 -1.04 -.02 137 -115)' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-opacity='0' /%3E%3Cstop offset='1' stop-color='%235488d6' stop-opacity='0' /%3E%3Cstop offset='1' stop-color='%235d96eb' /%3E%3C/radialGradient%3E%3C/defs%3E%3Crect width='220' height='220' x='22' y='-107' fill='url(%23s-a)' ry='49' transform='matrix(.57 0 0 .57 187 256)' /%3E%3Cg transform='translate(194 190)'%3E%3Ccircle cx='67.8' cy='67.7' fill='url(%23s-c)' paint-order='stroke fill markers' r='54' /%3E%3Ccircle cx='-69.9' cy='69.3' fill='url(%23s-i)' transform='translate(138 -2)' r='54' /%3E%3C/g%3E%3Cellipse cx='120' cy='14.2' fill='url(%23s-j)' rx='93.1' ry='93.7' transform='matrix(.58 0 0 .58 192 250)' /%3E%3Cg transform='matrix(.58 0 0 .57 197 182)'%3E%3Cpath fill='%23cac7c8' d='M46 192h1l72-48-7-9-66 57Z' /%3E%3Cpath fill='%23fbfffc' d='M46 191v1l66-57-7-9-59 65Z' /%3E%3Cpath fill='url(%23s-d)' d='m119 144-7-9 66-57-59 66Z' /%3E%3Cpath fill='%23fb645c' d='m105 126 7 9 66-57-1-1-72 49Z' /%3E%3C/g%3E%3Cpath stroke='%23fff' stroke-linecap='round' stroke-miterlimit='1' stroke-width='1.3' d='m287 278 3-2m-12-17 8-2m-8-3h4m-4-13 8 2m-8 3h4m-1-13 7 3m-4-11 7 4m-2-11 6 6m0-12 6 7m1-11 4 6m4-10 3 7m5-9 2 7m15-7-1 7m10-5-3 7m11-4-4 7m11-2-5 6m16 7-7 4m10 4-7 3m10 6-8 1m8 16-8-2m5 10-7-3m4 11-7-4m2 11-6-5m0 11-5-6m-2 11-4-7m-4 11-3-8m-6 10-1-8m-16 8 2-8m-10 5 3-7m-11 4 4-7m-11 2 5-6m-8 3 3-3m4 8 2-3m5 8 2-4m6 7 1-4m8 5v-4m8 4v-4m9 3-1-4m9 1-2-4m9 0-2-4m9-2-3-3m8-4-3-2m8-5-4-2m7-6-4-1m5-8h-4m4-8h-4m3-9-4 1m1-9-4 2m-1-9-3 2m-2-9-3 3m-4-8-2 3m-5-8-2 4m-6-6-1 3m-8-5v4m-8-4v4m-9-2 1 3m-9 0 2 3m-9 1 2 3m-9 2 3 3m-8 4 3 2m-8 5 4 2m-7 6 4 1m-4 25 4-1m-2 5 7-3m-6 7 4-2m-2 6 7-4m-13-21h8m41-41v-8m0 99v-8m49-42h-8' transform='translate(-65 8)' /%3E%3C/svg%3E)

Cuando el problema es la demora en la carga de recursos, es fundamental recordar **que puede ser demasiado tarde para lograr un buen LCP si una página debe esperar a que CSS o JavaScript se carguen por completo antes de que las imágenes puedan comenzar a cargarse**. Además, se puede reducir la duración de carga de recursos de una imagen de LCP si se vuelve a priorizar para que reciba más ancho de banda y, por lo tanto, se cargue más rápido con el atributo HTML `fetchpriority`.

Si tu elemento de LCP es una imagen, la URL de la imagen debe ser detectable en la respuesta HTML para reducir el retraso en la carga de recursos. Estas son algunas sugerencias para lograrlo:

- **Carga la imagen con un elemento ****`<img>`**** con el atributo ****`src`**** o ****`srcset`****.** No uses atributos no estándar, como `data-src`, que requieran JavaScript para renderizarse, ya que siempre serán más lentos. El [9%](https://almanac.httparchive.org/en/2022/performance#lcp-lazy-loading) de las páginas oculta su imagen de LCP detrás de `data-src`.
- **Prefiere la renderización del servidor (SSR) en lugar de la renderización del cliente (CSR),** ya que la SSR implica que el marcado completo de la página (incluida la imagen) está presente en la fuente HTML. Las soluciones de CSR requieren que se ejecute JavaScript para que se pueda descubrir la imagen.
- **Si se debe hacer referencia a tu imagen desde un archivo CSS o JS externo, puedes incluirla en la fuente HTML con una etiqueta ****`<link rel="preload">`****.** Ten en cuenta que el [análisis de precarga](https://web.dev/articles/preload-scanner?hl=es-419) del navegador no puede detectar las imágenes a las que hacen referencia los estilos intercalados. Por lo tanto, aunque se encuentren en la fuente HTML, su descubrimiento podría seguir bloqueado durante la carga de otros recursos, por lo que la precarga puede ser útil en estos casos.

Además, puedes acortar la duración de carga de un recurso asegurándote de que el recurso de LCP se cargue con anticipación y con prioridad alta:

- **Agrega el atributo ****`fetchpriority="high"`**** a la etiqueta ****`<img>`**** o ****`<link rel="preload">`**** de tu imagen de LCP.** Esto aumenta la prioridad del recurso de imagen para que pueda comenzar a cargarse antes.
- **Quita el atributo ****`loading="lazy"`**** de la etiqueta ****`<img>`**** de tu imagen de LCP.** Esto evita la demora de carga causada por confirmar que la imagen aparece en el viewport o cerca de él.
- **Aplaza los recursos no esenciales siempre que sea posible.** Mover estos recursos al final de tu documento, cargar [imágenes de forma diferida](https://web.dev/articles/browser-level-image-lazy-loading?hl=es-419) o [iframes](https://web.dev/articles/iframe-lazy-loading?hl=es-419), o bien cargarlos de forma asíncrona con JavaScript ayudará a despejar el camino para que los recursos más importantes, como la imagen LCP, se carguen más rápido.

**Lee los siguientes artículos para obtener más información:** [Cómo eliminar la demora en la carga de recursos](https://web.dev/articles/optimize-lcp?hl=es-419#1_eliminate_resource_load_delay) y [Mitos comunes sobre cómo optimizar el LCP](https://web.dev/blog/common-misconceptions-lcp?hl=es-419).

### 2. Intenta lograr navegaciones instantáneas

La experiencia del usuario ideal es no tener que esperar a que se cargue una página. Las optimizaciones de LCP, como la visibilidad y la priorización de recursos, son eficaces para reducir el tiempo que un usuario espera a que se cargue y se renderice el elemento LCP, pero existe un límite físico respecto de qué tan rápido se cargan esos bytes en la red y se renderizan en una página. Mucho antes de alcanzar ese límite, se requiere un esfuerzo prohibitivo para reducir solo unos pocos milisegundos más. Por lo tanto, para lograr navegaciones instantáneas, debemos adoptar un enfoque radicalmente diferente.

Las navegaciones instantáneas intentan cargar y renderizar la página _antes_ de que el usuario comience a navegar allí. De esta manera, la página renderizada previamente se puede mostrar de inmediato con un LCP casi nulo. Las restauraciones y las especulaciones son dos maneras de hacerlo. Cuando un usuario navega hacia atrás o hacia adelante a una página que visitó anteriormente, se puede restablecer rápidamente desde una memoria caché integrada y aparece exactamente como la dejó el usuario. Como alternativa, las aplicaciones web pueden intentar predecir adónde irá un usuario a continuación y, si es correcto, la siguiente página ya se habrá cargado y renderizado cuando el usuario navegue allí.

La [memoria caché atrás/adelante (bfcache)](https://web.dev/articles/bfcache?hl=es-419) permite restablecer las páginas visitadas anteriormente. Para usarla, debes asegurarte de que tus páginas cumplan con los [criterios de elegibilidad de bfcache](https://web.dev/articles/bfcache?hl=es-419#optimize). Los motivos habituales por los que las páginas no son aptas para la bfcache son que se publican con directivas de almacenamiento en caché [`no-store`](https://web.dev/articles/bfcache?hl=es-419#minimize-no-store) o tienen objetos de escucha de eventos [`unload`](https://web.dev/articles/bfcache?hl=es-419#never-use-the-unload-event).

El restablecimiento de páginas totalmente renderizadas mejora no solo el rendimiento de carga, sino también la estabilidad del diseño. Puedes obtener más información sobre bfcache y su eficacia para mejorar el CLS en la sección [Cómo garantizar que las páginas sean aptas para bfcache](https://web.dev/articles/top-cwv?hl=es-419#cls-bfcache).



![image](data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' width='24' height='24' viewBox='195 190 135 135'%3E%3Cdefs%3E%3ClinearGradient id='s-a' x1='132.6' x2='134.4' y1='111.7' y2='-105.3' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-color='%23d2d2d2' /%3E%3Cstop offset='.5' stop-color='%23f2f2f2' /%3E%3Cstop offset='1' stop-color='%23fff' /%3E%3C/linearGradient%3E%3ClinearGradient id='s-b' gradientUnits='userSpaceOnUse' /%3E%3ClinearGradient id='s-c' x1='65.4' x2='67.4' y1='115.7' y2='17.1' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-color='%23005ad5' /%3E%3Cstop offset='.2' stop-color='%230875f0' /%3E%3Cstop offset='.3' stop-color='%23218cee' /%3E%3Cstop offset='.6' stop-color='%2327a5f3' /%3E%3Cstop offset='.8' stop-color='%2325aaf2' /%3E%3Cstop offset='1' stop-color='%2321aaef' /%3E%3C/linearGradient%3E%3ClinearGradient id='s-d' x1='158.7' x2='176.3' y1='96.7' y2='79.5' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-color='%23c72e24' /%3E%3Cstop offset='1' stop-color='%23fd3b2f' /%3E%3C/linearGradient%3E%3CradialGradient id='s-i' cx='-69.9' cy='69.3' r='54' gradientTransform='matrix(.9 -.01 .04 2.72 -9 -120)' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-color='%2324a5f3' stop-opacity='0' /%3E%3Cstop offset='1' stop-color='%231e8ceb' /%3E%3C/radialGradient%3E%3CradialGradient id='s-j' cx='109.3' cy='13.8' r='93.1' gradientTransform='matrix(-.02 1.1 -1.04 -.02 137 -115)' xlink:href='%23s-b'%3E%3Cstop offset='0' stop-opacity='0' /%3E%3Cstop offset='1' stop-color='%235488d6' stop-opacity='0' /%3E%3Cstop offset='1' stop-color='%235d96eb' /%3E%3C/radialGradient%3E%3C/defs%3E%3Crect width='220' height='220' x='22' y='-107' fill='url(%23s-a)' ry='49' transform='matrix(.57 0 0 .57 187 256)' /%3E%3Cg transform='translate(194 190)'%3E%3Ccircle cx='67.8' cy='67.7' fill='url(%23s-c)' paint-order='stroke fill markers' r='54' /%3E%3Ccircle cx='-69.9' cy='69.3' fill='url(%23s-i)' transform='translate(138 -2)' r='54' /%3E%3C/g%3E%3Cellipse cx='120' cy='14.2' fill='url(%23s-j)' rx='93.1' ry='93.7' transform='matrix(.58 0 0 .58 192 250)' /%3E%3Cg transform='matrix(.58 0 0 .57 197 182)'%3E%3Cpath fill='%23cac7c8' d='M46 192h1l72-48-7-9-66 57Z' /%3E%3Cpath fill='%23fbfffc' d='M46 191v1l66-57-7-9-59 65Z' /%3E%3Cpath fill='url(%23s-d)' d='m119 144-7-9 66-57-59 66Z' /%3E%3Cpath fill='%23fb645c' d='m105 126 7 9 66-57-1-1-72 49Z' /%3E%3C/g%3E%3Cpath stroke='%23fff' stroke-linecap='round' stroke-miterlimit='1' stroke-width='1.3' d='m287 278 3-2m-12-17 8-2m-8-3h4m-4-13 8 2m-8 3h4m-1-13 7 3m-4-11 7 4m-2-11 6 6m0-12 6 7m1-11 4 6m4-10 3 7m5-9 2 7m15-7-1 7m10-5-3 7m11-4-4 7m11-2-5 6m16 7-7 4m10 4-7 3m10 6-8 1m8 16-8-2m5 10-7-3m4 11-7-4m2 11-6-5m0 11-5-6m-2 11-4-7m-4 11-3-8m-6 10-1-8m-16 8 2-8m-10 5 3-7m-11 4 4-7m-11 2 5-6m-8 3 3-3m4 8 2-3m5 8 2-4m6 7 1-4m8 5v-4m8 4v-4m9 3-1-4m9 1-2-4m9 0-2-4m9-2-3-3m8-4-3-2m8-5-4-2m7-6-4-1m5-8h-4m4-8h-4m3-9-4 1m1-9-4 2m-1-9-3 2m-2-9-3 3m-4-8-2 3m-5-8-2 4m-6-6-1 3m-8-5v4m-8-4v4m-9-2 1 3m-9 0 2 3m-9 1 2 3m-9 2 3 3m-8 4 3 2m-8 5 4 2m-7 6 4 1m-4 25 4-1m-2 5 7-3m-6 7 4-2m-2 6 7-4m-13-21h8m41-41v-8m0 99v-8m49-42h-8' transform='translate(-65 8)' /%3E%3C/svg%3E)

La renderización previa de la siguiente página que visita un usuario es otra forma eficaz de mejorar de forma significativa el rendimiento del LCP, y es posible gracias a la [API de Speculation Rules](https://developer.chrome.com/docs/web-platform/prerender-pages?hl=es-419). Sin embargo, para obtener estos beneficios, asegúrate de que se rendericen previamente las páginas correctas. Las especulaciones incorrectas desperdician recursos en el servidor y en el cliente, lo que podría perjudicar el rendimiento. Por lo tanto, cuanto menos seguro estés de cuál será la siguiente página, más conservador debes ser con la renderización previa. Cuando tengas dudas, tus datos de estadísticas pueden darte la confianza para renderizar previamente las páginas con mayor probabilidad de que se visiten a continuación.

**Lee para obtener más información:** [Memoria caché atrás/adelante](https://web.dev/articles/bfcache?hl=es-419) y [Procesamiento previo de páginas en Chrome para navegaciones instantáneas de páginas](https://developer.chrome.com/docs/web-platform/prerender-pages?hl=es-419).

### 3. Usa una CDN para optimizar el TTFB

La recomendación anterior se enfocó en las navegaciones instantáneas, que proporcionan la mejor experiencia posible a los usuarios, pero podría haber situaciones en las que no se apliquen las técnicas de bfcache y carga especulativa. Supongamos que un usuario sigue un vínculo de origen cruzado a tu sitio, donde la respuesta inicial del documento HTML bloquea de forma eficaz el LCP. El navegador no puede comenzar a cargar ningún subrecurso hasta que recibe el primer byte de la respuesta. Cuanto antes suceda, antes podrá comenzar a suceder todo lo demás.

Este tiempo se conoce como [tiempo hasta el primer byte (TTFB)](https://web.dev/articles/ttfb?hl=es-419). Las mejores formas de reducir el TTFB son las siguientes:

- Publica tu contenido lo más cerca posible de tus usuarios geográficamente.
- Almacena ese contenido en caché para que se pueda entregar rápidamente si se vuelve a solicitar en un futuro cercano.

La mejor manera de hacer ambas cosas es [usar una CDN](https://web.dev/articles/content-delivery-networks?hl=es-419). Las CDN distribuyen tus recursos a servidores perimetrales en todo el mundo, lo que limita la distancia que esos recursos deben recorrer por cable hasta los usuarios. Las CDN también suelen tener controles de almacenamiento en caché detallados que pueden modificarse según las necesidades de tu sitio.

Las CDN también pueden entregar y almacenar en caché documentos HTML, pero, según el Almanac web, solo el [29% de las solicitudes de documentos HTML se entregaron desde una CDN](https://almanac.httparchive.org/en/2022/cdn#cdn-adoption). Esto significa que los sitios tienen una oportunidad significativa para obtener ahorros adicionales.

Estas son algunas sugerencias para configurar las CDN:

- Almacena en caché documentos HTML estáticos, aunque sea por poco tiempo. Por ejemplo, ¿es importante que el contenido siempre esté actualizado? ¿O puede estar desactualizada unos minutos?
- Explora si puedes mover la lógica dinámica que se ejecuta en tu servidor de origen al [periférico](https://en.wikipedia.org/wiki/Edge_computing), que es una función de la mayoría de las CDN modernas.

Cada vez que puedes entregar contenido directamente desde el perímetro y evitar un viaje a tu servidor de origen, se obtiene un aumento de rendimiento. Incluso en los casos en los que _debes_ realizar todo el recorrido hasta el origen, las CDN suelen estar optimizadas para hacerlo más rápido, por lo que es una ventaja de cualquier manera.

**Lee para obtener más información:** [Optimiza el tiempo hasta el primer byte](https://web.dev/articles/optimize-ttfb?hl=es-419).

## Cumulative Layout Shift (CLS)

El [Cambio de diseño acumulado (CLS)](https://web.dev/articles/cls?hl=es-419) es una medida de la estabilidad visual de una página web. Si bien CLS es la métrica con la que la mayoría de los sitios tienden a lograr un buen rendimiento, alrededor de [una cuarta parte de ellos](https://httparchive.org/reports/chrome-ux-report#cruxSmallCls) aún no cumplen con el [umbral recomendado](https://web.dev/articles/cls?hl=es-419#what_is_a_good_cls_score), por lo que sigue existiendo una gran oportunidad para que muchos sitios mejoren la experiencia del usuario.

### 1. Configura tamaños explícitos en cualquier contenido que se cargue desde la página

Los [cambios de diseño](https://web.dev/articles/cls?hl=es-419#layout_shifts_in_detail) suelen ocurrir cuando el contenido existente se mueve después de que otro contenido termina de cargarse. La forma principal de mejorar CLS es reservar el espacio requerido con la mayor anticipación posible.

La mejor manera de corregir los cambios de diseño causados por imágenes sin tamaño es **configurar explícitamente los atributos ****`width`**** y ****`height`** o sus propiedades de CSS equivalentes. El [72%](https://almanac.httparchive.org/en/2022/performance#explicit-dimensions) de las páginas tiene al menos una imagen sin tamaño. Sin un tamaño explícito, estas imágenes tienen una altura inicial de `0px`, lo que puede provocar cambios de diseño cuando se cargan estas imágenes y el navegador descubre sus dimensiones. Esto representa una gran oportunidad para la Web colectiva, y esa oportunidad requiere menos esfuerzo que algunas de las otras recomendaciones sugeridas en esta guía.

Las imágenes no son los únicos factores que afectan el CLS. Los cambios de diseño pueden deberse a otro contenido que, por lo general, se carga después de que se renderiza la página inicialmente, incluidos los anuncios de terceros o los videos incorporados. La propiedad [`aspect-ratio`](https://web.dev/articles/aspect-ratio?hl=es-419) puede ayudarte en este caso. Es una función de CSS de [Baseline ampliamente disponible](https://web.dev/baseline?hl=es-419) que permite a los desarrolladores establecer explícitamente una relación de aspecto tanto en imágenes como en elementos que no sean imágenes. Esto te permite establecer un `width` dinámico (por ejemplo, en función del tamaño de la pantalla) y hacer que el navegador calcule automáticamente la altura adecuada, de la misma manera que lo hace para las imágenes con dimensiones.

Sin embargo, no siempre es posible conocer el tamaño exacto del contenido dinámico. Incluso si no sabes el tamaño exacto, puedes reducir la gravedad de los cambios de diseño. **Configurar un ****`min-height`**** razonable** casi siempre es mejor que permitir que el navegador use la altura predeterminada de `0px` para un elemento vacío. El uso de un `min-height` también suele ser una solución directa, ya que permite que el contenedor crezca hasta la altura del contenido final si es necesario, solo que reduce esa cantidad de crecimiento a un nivel más tolerable.

**Lee para obtener más información:** [Optimiza el cambio de diseño acumulado](https://web.dev/articles/optimize-cls?hl=es-419#images-without-dimensions).

### 2. Asegúrate de que las páginas sean aptas para la bfcache

Como se indicó anteriormente en esta guía, la [memoria caché atrás/adelante](https://web.dev/articles/bfcache?hl=es-419) (bfcache) carga instantáneamente una página anterior o posterior del historial del navegador a partir de una instantánea de memoria. Si bien la bfcache es una optimización de rendimiento significativa a nivel del navegador que mejora el LCP, también elimina por completo los cambios de diseño. De hecho, [la introducción de bfcache en 2022](https://chromium.googlesource.com/chromium/src/+/refs/heads/main/docs/speed/metrics_changelog/2022_01_bfcache.md) fue responsable de la mayor mejora en el CLS que vimos ese año.

A pesar de esto, [una cantidad significativa de sitios web](https://almanac.httparchive.org/en/2022/performance#bfcache-eligibility) no son aptos para la bfcache y, por lo tanto, se están perdiendo este beneficio gratuito de rendimiento web. A menos que tu página cargue información sensible que no quieras que se restablezca de la memoria, asegúrate de que tus páginas sean aptas para usar la bfcache.

Los propietarios de sitios deben verificar si las páginas son [aptas para la bfcache](https://web.dev/articles/bfcache?hl=es-419#optimize_your_pages_for_bfcache) y corregir los motivos por los que no lo son. Chrome [tiene un verificador de bfcache en Herramientas para desarrolladores](https://web.dev/articles/bfcache?hl=es-419#test_to_ensure_your_pages_are_cacheable) y también puedes usar la [API de Not Restored Reasons](https://developer.chrome.com/docs/web-platform/bfcache-notrestoredreasons?hl=es-419) para detectar motivos de inelegibilidad en el campo.

**Obtén más información:** [Memoria caché atrás/adelante](https://web.dev/articles/bfcache?hl=es-419).

### 3. Evita las animaciones y transiciones que usan propiedades CSS que inducen el diseño

Otra fuente común de cambios de diseño es cuando se animan los elementos. Por ejemplo, los banners de cookies y otros banners de notificaciones que se deslizan desde la parte superior o inferior suelen contribuir a la CLS. Esto es particularmente problemático cuando estos banners despliegan otro contenido, pero incluso cuando no lo hacen, animarlos puede afectar CLS.

Si bien los datos de HTTP Archive no pueden conectar de forma concluyente las animaciones con los cambios de diseño, los datos muestran que las páginas que animan cualquier propiedad de CSS que _podría_ afectar el diseño tienen un 15% menos de probabilidades de tener un CLS “bueno” que las páginas en general. Algunas propiedades tienen un CLS peor que otras. Por ejemplo, las páginas con animaciones de ancho de `margin` o `border` tienen un CLS "deficiente" casi el doble que el porcentaje de páginas que, en general, se consideran deficientes.

Quizás no sea sorprendente, ya que cada vez que realices la transición o la animación de _cualquier_ propiedad de CSS que induzca el diseño, se producirán [cambios de diseño](https://web.dev/articles/cls?hl=es-419#layout_shifts_in_detail). Si esos cambios de diseño no se producen dentro de los 500 milisegundos posteriores a una interacción del usuario, afectarán a la CLS.

Lo que puede sorprender a algunos desarrolladores es que esto es cierto incluso en los casos en que el elemento se toma fuera del flujo normal del documento. Por ejemplo, los elementos de posicionamiento absoluto que animan `top` o `left` provocan cambios de diseño, incluso si no están empujando otro contenido. Sin embargo, si, en lugar de animar `top` o `left`, animas `transform:translateX()` o `transform:translateY()`, el navegador no actualizará el diseño de la página, lo que evitará los cambios de diseño.

Por mucho tiempo, preferir la animación de las propiedades de CSS que se pueden actualizar en el subproceso compositor del navegador ha sido [una práctica recomendada para el rendimiento](https://web.dev/articles/animations-guide?hl=es-419) porque mueve ese trabajo del subproceso principal a la GPU. Además de ser una práctica recomendada de rendimiento general, también puede ayudar a mejorar el CLS.

Como regla general, nunca animes ni realices transiciones de propiedades CSS que requieran que el navegador actualice el diseño de la página, a menos que lo hagas en respuesta a un toque del usuario o a una pulsación de tecla (aunque [no ](https://web.dev/articles/cls?hl=es-419#user-initiated_layout_shifts)[`hover`](https://web.dev/articles/cls?hl=es-419#user-initiated_layout_shifts)). Siempre que sea posible, prefiere las transiciones y animaciones con la propiedad [`transform`](https://developer.mozilla.org/docs/Web/CSS/transform) de CSS.

La auditoría de Lighthouse [Evitar animaciones no compuestas](https://developer.chrome.com/docs/lighthouse/performance/non-composited-animations/?hl=es-419) advierte cuando una página anima propiedades de CSS potencialmente lentas.

**Para obtener más información, consulta:** [Cómo optimizar los cambios de diseño inducidos por animaciones](https://web.dev/articles/optimize-cls?hl=es-419#animation).

## Conclusión

Mejorar el rendimiento de la página puede parecer abrumador, en especial si se tienen en cuenta muchos consejos que se deben tener en cuenta en toda la Web. Sin embargo, si te enfocas en esta breve lista de las prácticas recomendadas más eficaces, puedes abordar el problema con un enfoque renovado y, con suerte, mejorar las Métricas web esenciales de tu sitio web.

Si deseas ir más allá de las optimizaciones que se mencionan aquí, lee estas guías para obtener más información:

- [Optimiza la INP](https://web.dev/articles/optimize-inp?hl=es-419)
- [Optimiza el LCP](https://web.dev/articles/optimize-lcp?hl=es-419)
- [Optimiza la métrica CLS](https://web.dev/articles/optimize-cls?hl=es-419)
