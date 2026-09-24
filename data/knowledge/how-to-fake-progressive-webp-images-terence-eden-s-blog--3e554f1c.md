---
title: "How to fake Progressive WebP Images – Terence Eden’s Blog"
notion_id: 3e554f1c-7d23-81fa-96a5-e2db2a362004
notion_url: https://app.notion.com/p/How-to-fake-Progressive-WebP-Images-Terence-Eden-s-Blog-3e554f1c7d2381fa96a5e2db2a362004
last_edited: 2026-09-24T03:06:00.000Z
source_url: https://shkspr.mobi/blog/2020/04/how-to-fake-progressive-webp-images/
tags: ["Article", "Tutorial", "Terence Eden’s Blog", "English", "Web Development", "Image Editing", "Progressive Enhancement", "HTML", "CSS"]
---
WebP is the [hip new image format](https://developers.google.com/speed/webp) on the scene. It offers unrivalled image compression at superior visual quality. But, in my opinion, it is deficient compared to JPG in one significant aspect. It doesn't have a progressive mode.

Progressive mode is useful because it can quickly load a low resolution preview of an image, and then gradually improve its quality.

WebP, by contrast, just loads up the full image line by line. That's can be annoying on a slow connection.

What if we could change that? Here's a quick-and-dirty hack.
This is a 5472x3080 image which has been compressed down to "zero" quality on WebP and resized to 640x360. It is less than 5KB.

![image](https://shkspr.mobi/blog/wp-content/uploads/2020/04/pi-0-640.webp)

The [original image](https://pixabay.com/photos/raspberry-pi-raspi-electronics-3640738/) is 400 KB as a WebP and 1.6MB as a JPG.

There are [lots](https://jmperezperez.com/svg-placeholders/) of [techniques](https://developers.google.com/web/fundamentals/performance/lazy-loading-guidance/images-and-video) to [load low-res images before the main image loads](https://css-tricks.com/the-blur-up-technique-for-loading-background-images/).

We're going old-skool. Basic HTML + CSS.

Copied HTML to 📋

```plain text
 HTML<img
    width=5472
    height=3080
    alt="Close up of a Raspberry Pi circuit board."
    style="background-image: url('Small.webp');"
    src=Big.webp
>
```

![image](https://shkspr.mobi/blog/wp-content/plugins/tempest-highlight/svg/html.svg)

This sets the _background_ of the image to the low-resolution, low quality version. As the high-resolution image loads, it gradually replaces the image in the background.

There are a few more refinements we could make to this:

Copied HTML to 📋

```plain text
 HTML<img
    width=5472
    height=3080
    alt="Close up of a Raspberry Pi circuit board."
    style="background-color: #0f0;
           background-image: url('Small.webp');"
    loading="lazy"
    src=Big.webp
/>
```

This sets the `background-color` to the dominant colour of the image. That way, even before the background image has loaded, you get something to fill the gap. Because we've set `loading="lazy"` the massive image won't be downloaded until the `img` scrolls into view.

We can even go a step further.  Let's reduce the image down to a thumbnail of 160x90. That takes us down to a ludicrously small 564 Bytes.

![image](https://shkspr.mobi/blog/wp-content/uploads/2020/04/pi-0-160.webp)

Which makes it small enough to stick directly in the HTML once we encode it in Base64:
					Copied CSS to 📋

```plain text
 CSSbackground-image: url("data:image/webp;base64,UklGRiwCAABXRUJQVlA4ICACAACQEwCdASqgAFoAP/3+/3+/vLYyPv+8A/A/iWYIkCfrS2sOTiSUArLU15JWpmwHGTJDQ+i/y0UfRMWET8I422NGGsABPAdIKWnTFvlfPnla/wC0DPVYABJxuECkmqWUZwViSm1vqzvPgAnG+vioXrwz0V3Zx+c/znFFJvKBpH5SceJwmAUi8+8SRjzQvyc/qv3DKe3KfEYnXA+5aEBypmacRUvrvgAA+OUzEDYgEioYgJQRm/l9OIIl60M80PJW00cW6fX1fNpB/L4udWsUF5v6gOR6fB//5/0zuirXi2Z2GoqpVh3aeUDEdVThFtQOZStD/ulvCFO54uEqPVZlD63ukgBgt5Q5KA15Nse9lu4E+4XliSgao9azVod9zUr/XELtUpd2CLSyImxIXq3rRNGthLv0jePmYvlqijrhFxgMsdVer/GmM0C6xEyTC53yBQ+aGJlg/iqxrU5uwNjGNi3Or75QpJVdWW0lEyZlOsIiRsCZ+jTXsDMmgzkB4uaiiQASLl4TIYpK587PmaLb29WRF3hxpAtjR4TPDHne3iY6aBfYZhvMo1N41AMMKe7BLZg3w8LCQ2RHRHnkhgYsdtLL2jSFFpP8JgVtwLZ/KtKa6P7VaHm7aY7aT0UWyEwXPnUbZwo5a4yuK1H+P7YJqLj8yjuKSJZACY6z+6RgRxeoSyytYvQiopOOlc4f1Mdoy2WgfE3HR6Ap5jJiPryYLQAA")
```

![image](https://shkspr.mobi/blog/wp-content/plugins/tempest-highlight/svg/css.svg)

Of course, you don't have to go down to quite this extreme level - choose some settings which make sense for you and your media.

To summarise:

1. Scale the image down
2. Reduce its quality
3. Set it as the background to the	`img` element

## [Endnote](https://shkspr.mobi/blog/2020/04/how-to-fake-progressive-webp-images/#endnote)

I wasn't involved with the development of WebP - but it seems bizarre to me that it doesn't contain a "thumbnail mode". On that ½ MB photo, adding a couple of KB doesn't seem like a huge overhead.

Similarly, the [experimental AVIF](https://netflixtechblog.com/avif-for-next-generation-image-coding-b1d75675fe4) also lacks progressive / thumbnail support.

Anyone know what the reason is?
