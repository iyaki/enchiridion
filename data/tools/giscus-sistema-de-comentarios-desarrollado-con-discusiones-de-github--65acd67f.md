---
title: "giscus - Sistema de comentarios desarrollado con Discusiones de GitHub"
notion_id: 65acd67f-e4d7-41e3-b9e5-2a5c628707cd
notion_url: https://app.notion.com/p/giscus-Sistema-de-comentarios-desarrollado-con-Discusiones-de-GitHub-65acd67fe4d741e3b9e52a5c628707cd
last_edited: 2026-09-18T00:54:00.000Z
source_url: https://giscus.app/
tags: ["Español", "English", "Others", "Communication", "Blogging/Content Creation", "Programming", "Tool", "Framework/Library"]
---
Un sistema de comentarios desarrollado con [Discusiones de GitHub](https://docs.github.com/en/discussions). ¡Permite a sus visitantes dejar comentarios y reacciones en su sitio web a través de GitHub! Inspirado en [utterances](https://github.com/utterance/utterances).

- Sin seguimiento, sin anuncios, siempre gratis. 📡 🚫
- No se necesita base de datos. Todos los datos se almacenan en Discusiones de GitHub.
- Obtiene automáticamente nuevos comentarios y ediciones de GitHub. 🔃

> 

**Note:** giscus todavía está en desarrollo activo. GitHub también sigue desarrollando activamente Discusiones y su API. Por lo tanto, algunas características de giscus pueden romperse o cambiar con el tiempo.

## Cómo funciona

Cuando se carga giscus, la [API de búsqueda de discusiones de GitHub](https://docs.github.com/en/graphql/guides/using-the-graphql-api-for-discussions#search) se usa para encontrar la discusión asociada con la página en función de la forma de mapeo elegida (URL, `pathname`, `<title>`, etc.). Si no se puede encontrar una discusión que coincida, el bot giscus creará automáticamente una discusión la primera vez que alguien deje un comentario o una reacción.

Para comentar, los visitantes deben autorizar la [aplicación de giscus](https://github.com/apps/giscus) para [publicar en su nombre](https://docs.github.com/en/developers/apps/identifying-and-authorizing-users-for-github-apps) utilizando el flujo de GitHub OAuth. Alternativamente, los visitantes pueden comentar sobre la Discusión de GitHub directamente. Puede moderar los comentarios en GitHub.

## Configuración

### Idioma

Selecciona el idioma en el que se mostrará giscus. ¿No encuentras tu idioma? Puedes [Constribuir](https://github.com/giscus/giscus/blob/main/CONTRIBUTING.md#adding-localizations) con la traducción.

Español

### Repositorio

Selecciona el repositorio al que se conectará giscus. Asegúrate de lo siguiente:

1. El **repositorio es **[**público**](https://docs.github.com/en/github/administering-a-repository/managing-repository-settings/setting-repository-visibility#making-a-repository-public), de lo contrario, tus visitantes no podrán ver los comentarios.
2. La aplicación de [**giscus**](https://github.com/apps/giscus)** está instalada**, de lo contrario, tus visitantes no podrán comentar ni reaccionar.

Un repositorio **público** de GitHub. Es el repositorio donde se vincularán las discusiones.

### Página ↔️ Mapeo de discusiones

Elige la forma de mapeo, entre la página y las discusiones

**El título de la discusión contiene el ****`pathname`**** de la página**

giscus buscará una discusión cuyo título contenga el `pathname` de la página.

**El título de la discusión contiene la ****`URL`**** de la página**

**El título de la discusión contiene el ****`<title>`**** de la página.**

giscus buscará una discusión cuyo título contenga la etiqueta HTML `<title>` de la página.

**El título de la discusión contiene el valor establecido en el atributo ****`og:title`**

giscus buscará una discusión cuyo título contenga el valor de la etiqueta HTML [`<meta property="og:title">`](https://ogp.me/).

**El título de la discusión contiene un término específico**

**Número de discusión específico**

giscus cargará una discusión específica por número. Esta opción **no** admite la creación automática de discusiones.

**Usar coincidencia estricta de títulos**

Evite discrepancias debido al método de búsqueda difuso de GitHub cuando hay varias discusiones con títulos similares. Consulte [ la documentación ](https://github.com/giscus/giscus/blob/main/ADVANCED-USAGE.md#data-strict) para obtener más detalles.

### Categoría de discusión

Selecciona la categoría donde se crearán nuevas discusiones. Se recomienda usar una categoría con el tipo **Announcement** para que solo los mantenedores y giscus puedan crear nuevas discusiones.

Categoría de discusión

**Buscar solo discusiones en esta categoría**

### Características

Selecciona las funcionalidades específicas que quieres habilitar.

**Habilitar reacciones para la publicación principal**

**Emitir metadatos de discusión**

Los metadatos de discusión se enviarán periódicamente a la ventana principal (la página de incrustación). Para una demostración, habilite esta opción y abra la consola de su navegador en esta página. Consulte [ la documentación ](https://github.com/giscus/giscus/blob/main/ADVANCED-USAGE.md#imetadatamessage) para obtener más detalles.

**Coloque el cuadro de comentarios encima de los comentarios**

El cuadro de entrada de comentarios se colocará encima de los comentarios, para que los usuarios puedan dejar un comentario sin desplazarse hasta el final de la discusión.

**Cargar los comentarios perezosamente**

La carga de los comentarios se aplazará hasta que el usuario se desplace cerca del contenedor de comentarios. Esto se hace agregando [`loading="lazy"`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe#attr-loading) al elemento `<iframe>`.

### Tema

Selecciona un tema que coincida con su sitio web. ¿No puede encontrar uno que lo haga? [Contribuya](https://github.com/giscus/giscus/blob/main/CONTRIBUTING.md#creating-new-themes) con un tema nuevo.

Preferred color scheme

### Habilitar giscus

Agregue la siguiente etiqueta ` <script> ` a la plantilla de su sitio web donde desea que aparezcan los comentarios. Si existe un elemento con la clase ` giscus `, los comentarios se mostrarán allí.

No has configurado tu [repositorio](https://giscus.app/es#repository) y/o [categoría](https://giscus.app/es#category). Los valores de esos campos no se mostrarán hasta que los complete.

```plain text
<script src="https://giscus.app/client.js"
        data-repo="[URL REPOSITORIO]"
        data-repo-id="[REPOSITORIO ID]"
        data-category="[NOMBRE CATEGORÍA]"
        data-category-id="[ID CATEGORÍA]"
        data-mapping="pathname"
        data-strict="0"
        data-reactions-enabled="1"
        data-emit-metadata="0"
        data-input-position="bottom"
        data-theme="preferred_color_scheme"
        data-lang="es"
        crossorigin="anonymous"
        async>
</script>
```

Puede personalizar el diseño del contenedor utilizando los selectores `.giscus` y `.giscus-frame` de la página de incrustación.

Si está usando giscus, considere [recomendar 🌟 giscus en GitHub](https://github.com/giscus/giscus) y agrega [`giscus`](https://github.com/topics/giscus) topic [en tu repositorio](https://docs.github.com/en/github/administering-a-repository/classifying-your-repository-with-topics)! 🎉

## Uso avanzado

Puede agregar configuraciones adicionales (por ejemplo, permitir orígenes específicos) siguiendo la [guía de uso avanzado](https://github.com/giscus/giscus/blob/main/ADVANCED-USAGE.md).

Para usar giscus con React, Vue o Svelte, consulte la [biblioteca de componentes de giscus](https://github.com/giscus/giscus-component).

## Migrando

Si ha utilizado anteriormente otros sistemas que usan GitHub Issues (p.ej. [utterances](https://github.com/utterance/utterances), [gitalk](https://github.com/gitalk/gitalk)), puedes [convertir los issues existentes en discusiones](https://docs.github.com/en/discussions/managing-discussions-for-your-community/moderating-discussions#converting-an-issue-to-a-discussion). Después de la conversión, asegúrese de que el mapeo entre los títulos de la discusión y las páginas sea correcto, entonces giscus utilizará automáticamente las discusiones.

## Sitios que usan giscus

## Contribución

## Pruébalo 👇👇👇

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
