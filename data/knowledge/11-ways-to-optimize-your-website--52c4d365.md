---
title: "11 Ways to Optimize Your Website"
notion_id: 52c4d365-b88c-4c42-860a-39719a882a5b
notion_url: https://app.notion.com/p/11-Ways-to-Optimize-Your-Website-52c4d365b88c4c42860a39719a882a5b
last_edited: 2026-09-21T16:58:00.000Z
source_url: https://levelup.gitconnected.com/11-ways-to-optimize-your-website-1711b0d42de8?gi=7179e0caed66
tags: ["Article", "Guide", "Level Up Coding by gitconnected", "English", "Web Development", "Frontend", "CSS", "HTML"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You’ve put weeks of hard work into building your website, and now it’s finally ready to go live! However, to ensure your website performs at its best, there are a few things you need to take care of. In this article, we will explore various ways to optimize your website for better user experience and faster loading times, which leads to higher search engine rankings.

> 📧 Subscribe to my newsletter: https://ericsdevblog.ck.page/profile

When the browser requests a webpage, it will first receive an [HTML](https://www.ericsdevblog.com/courses/html-css/) document. The browser will parse the document, and when the parser encounters an external file, the browser will send another request to retrieve it. For example, imagine you have a website with the following structure:

```plain text
.
├── index.html
├── package.json
└── statics
    ├── css
    │   ├── footer.css
    │   ├── header.css
    │   └── main.css
    ├── images
    │   ├── 1.jpg
    │   ├── 2.jpg
    │   ├── 3.png
    │   ├── 4.png
    │   └── 5.jpg
    └── js
        ├── foo.js
        ├── bar.js
        └── index.js
```

To render the `index.html` file, your browser would have to make a total of 12 requests, including 1 HTML file, 3 CSS files, 3 JavaScript files, and 5 images. This process will consume a large amount of time and resources, leading to poor website performance.

Fortunately, there are some ways to enhance the loading speed of this website, such as combining and minimizing the static files, optimizing the images, caching the resources, and many more. We will explore all of these techniques in this article. But first, let’s start with the images.

## Use modern web image formats

To begin with, the images on this website are in either JPG or PNG format, which tend to have larger file sizes and poor performance compared to modern formats such as WebP and AVIF.

There are many cloud-based tools and websites that can convert your images, but the problem with these tools is that you usually have to upload the files for them to be processed, and some of their services are not free. In this article, I’d like to introduce a piece of software called [FFmpeg](https://ffmpeg.org/), which allows you to convert the images locally with one simple command.

If you are using Mac, you can install FFmpeg with [Homebrew](https://brew.sh/):

```plain text
brew install ffmpeg
```

If you are using Windows, use [Winget](https://winget.run/) instead:

```plain text
winget install --id=Gyan.FFmpeg  -e
```

Alternatively, you can simply download the installer from [FFmpeg’s official website](https://ffmpeg.org/download.html).

After the installation process, open the terminal and change into the image directory.

```plain text
cd /path/to/images
```

And then convert the images using the following command:

```plain text
for file in *.{jpg,png}; do ffmpeg -i "$file" -c:v libwebp -q:v 80 "$(basename "$file" .${file##*.}).webp"; done
```

If you are using Windows (CMD), run this command instead:

```plain text
for %i in (*.jpg *.png) do ffmpeg -i "%i" -c:v libwebp -q:v 80 "%~ni.webp"
```

For PowerShell:

```plain text
Get-ChildItem -Path . | Where-Object { $_.Extension -match '\.jpg$|\.png$' } | ForEach-Object { ffmpeg -i $_.FullName -c:v libwebp -q:v 80 ($_.BaseName + ".webp") }
```

Of course, you may have to alter this command to fit your specific condition:

- `{jpg,png}` lists all the image formats in the directory.
- `c:v libwebp` specifies the codec used for WebP. You don't need to change this unless you know what you are doing.
- `q:v 80` sets the compression level for the images. You can adjust the value between `0` (lowest quality, highest compression) and `100` (highest quality, no compression) as needed.

You can play around with the compression level, but in my experience, it is safe to set it to as low as `20` without affecting the image quality too much. Here is a comparison.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

FFmpeg is a very powerful multimedia tool that can handle a wide range of tasks related to audio, image, and video processing. In fact, many of the world’s most famous websites and tools run on top of FFmpeg, such as YouTube, Twitch, VLC Media Player, and so on. Please refer to the [official documentation](https://ffmpeg.org/documentation.html) for details.

## Different images for different viewports

In most cases, your website is created for devices with different screen sizes. For a small screen, it is usually best to use smaller images, and for larger screens, it is best to use large images. Ideally, you could create different versions of the same image using FFmpeg and then embed the images using the `<picture>` element instead of `<img>`.

The `<picture>` element allows you to define multiple sources for the same image, and then the browser can choose different sources based on the viewport size.

```plain text
<picture>
    <source media="(max-width: 600px)" srcset="small.webp">
    <source media="(max-width: 1200px)" srcset="medium.webp">
    <img src="large.webp" alt="example image">
</picture>
```

In this example, the browser will render `small.webp` on a small screen (`<600px`), `medium.webp` on a medium screen (`600px-1200px`), and `large.webp` on a large screen (`>1200px`)

## Lazy load your images

Lastly, you should also lazy load the images if they are not needed immediately.

```plain text
<picture>
    <source media="(max-width: 600px)" srcset="small.webp">
    <source media="(max-width: 1200px)" srcset="medium.webp">
    <img src="large.webp" alt="example image" loading="lazy">
</picture>
```

This ensures that the browser only retrieves the image when the user scrolls down to its location. However, if an image is required for the webpage to render correctly during the initial load, then it is best to set `load` to `eager`, which tells the browser to retrieve it as soon as possible.

```plain text
<picture>
    <source media="(max-width: 600px)" srcset="small.webp">
    <source media="(max-width: 1200px)" srcset="medium.webp">
    <img src="large.webp" alt="example image" loading="eager">
</picture>
```

Secondly, notice that there are 3 CSS files and 3 JavaScript files for this project. It is best to separate the code into different modules during the development stage for more efficient file management, as we’ve discussed in the [HTML & CSS Best Practices](https://www.ericsdevblog.com/courses/html-css/8/) article. However, in the production environment, you’ll want your webpage to download as few external files as possible to improve user experience. And you’ll want the files to be as small as possible.

## Combine and minimize CSS files

There are many frontend tools available for this purpose. For example, [PostCSS](https://postcss.org/) is a popular CSS processor that can combine and minimize your code. With the right plugin, it can even fix your code for compatibility issues, making sure your [CSS](https://www.ericsdevblog.com/courses/html-css/) styles work for all browsers.

PostCSS is built into the web bundlers, which we are going to discuss later. However, if you wish to use PostCSS independently, it can be installed into your project using the following [`npm`](https://www.npmjs.com/) command:

```plain text
npm install postcss postcss-cli postcss-import postcss-preset-env cssnano --save-dev
```

Create a configuration file, `postcss.config.js`, under the project root directory. The configuration should include all necessary plugins.

```plain text
module.exports = {
  plugins: [
    require("postcss-import"),
    require("postcss-preset-env"),
    require("cssnano"),
  ],
};
```

Create an input CSS file. This input file should import all other necessary CSS files.

```plain text
.
├── index.html
├── package-lock.json
├── package.json
├── postcss.config.js
└── statics
    ├── css
    │   ├── footer.css
    │   ├── header.css
    │   ├── main.css
    │   └── styles.css
    ├── images
    └── js
```

`styles.css`

```plain text
@import "./header.css";
@import "./main.css";
@import "./footer.css";
```

Combine and minimize the file using the following command:

```plain text
npx postcss statics/css/styles.css -o dist/styles.css
```

And in your HTML document, you only need to import the output `dist/styles.css` file.

```plain text
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <link rel="stylesheet" href="dist/styles.css">
</head>
<body>
    . . .
</body>
</html>
```

Alternatively, you can also separate the critical and non-critical CSS, load the critical CSS in the head section, and load the non-critical CSS at the end of the body section.

```plain text
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <link rel="stylesheet" href="dist/critical.css">
</head>
<body>
    . . .
    <link rel="stylesheet" href="dist/non-critical.css">
</body>
</html>
```

## Use web bundlers

Nowadays, the complexity of web applications has grown exponentially. You cannot rely on a basic CSS processor to optimize and manage everything. Web bundlers are created to address this challenge. They are designed to process CSS, JavaScript, as well as images, allowing you to manage all of your static assets in one place.

[Webpack](https://webpack.js.org/) is one of the most popular options, with [63.6k stars](https://github.com/webpack/webpack) on GitHub. It packs modules (JavaScript, CSS, images, and so on) into bundled assets that can be executed by the browser. For demonstration purposes, this is how you can install Webpack into your project:

```plain text
npm install webpack webpack-cli --save-dev
```

Assume you have the following project structure:

```plain text
.
├── dist
├── index.html
├── package-lock.json
├── package.json
├── postcss.config.js
├── statics
│   ├── css
│   ├── images
│   └── js
│       ├── bar.js
│       ├── foo.js
│       └── index.js
└── webpack.config.js
```

`foo.js`

```plain text
export default function foo() {
  console.log("foo");
}
```

`bar.js`

```plain text
export default function bar() {
  console.log("bar");
}
```

`index.js`

```plain text
import foo from "./foo.js";
import bar from "./bar.js";

foo();
bar();
console.log("Hello, World!");
```

In this case, `index.js` is the entry point of your project. Create a configuration file (`webpack.config.js`) under the root directory of your project, and make sure the `entry` option points to the `index.js` file.

```plain text
const path = require("path");

module.exports = {
  entry: "./statics/js/index.js",
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "dist"),
  },
};
```

Next, you can run Webpack with the following command:

```plain text
npx webpack --config webpack.config.js
```

A `bundle.js` file will be generated, containing the minimized version of all JavaScript code.

```plain text
(()=>{"use strict";console.log("foo"),console.log("bar"),console.log("Hello, World!")})();
```

By default, Webpack only deals with JavaScript files, but you can extend its capabilities by installing different [loaders](https://webpack.js.org/loaders/). For example, the [css-loader](https://webpack.js.org/loaders/) enables Webpack to process your CSS files, and the [postcss-loader](https://webpack.js.org/loaders/postcss-loader/) makes it compatible with the PostCSS processor we just discussed. Please refer to the linked webpages for details.

Besides Webpack, there are many other popular web bundlers available, such as [Parcel](https://parceljs.org/), [Esbuild](https://esbuild.github.io/), [Rollup](https://rollupjs.org/), and more. They all have their own unique features and strengths, and you should make your decision based on the needs and requirements of your specific project. Please refer to their official websites for details.

Speaking of frontend tools, [Vite](https://vitejs.dev/) is definitely one that we can’t afford to overlook. As your application grows increasingly complex, it is not uncommon for a single application to have hundreds or even thousands of modules. As a result, it often takes an unnecessarily long time for the web bundlers to process all of them before a dev server can be started.

Vite is created to address this issue by providing native support for Hot Module Replacement (HMR), which is a technology that allows developers to apply code updates in real-time without having to refresh the entire page. It also takes a unique approach when it comes to asset bundling. Instead of bundling everything together, it creates smaller bundles for each individual module, and then serves to the browser as needed. This approach allows Vite to have faster build and load times. If you are looking for a frontend build tool that is fast and reliable, definitely give Vite a shot.

## Async vs. defer

Nowadays, JavaScript files are getting more and more complex. They are often heavier than the HTML document, and takes longer to download and process, even when they are combined and minimized.

By default, the browser parses the HTML file line by line, and when it encounters a script, the parser will stop to download the file, read and execute it, and then continue processing the rest of the page.

However, in most cases, it is safe for your JavaScript files to be loaded asynchronously without blocking the parser. To achieve that, you can use the `defer` or `async` attribute.

```plain text
<script src="path/to/script.js" async></script>
```

```plain text
<script src="path/to/script.js" defer></script>
```

Both options instruct the browser to download the script in the background. Their difference is that `async` tells the browser to execute the script right after it is downloaded, while `defer` tells the browser to wait until the parser is completed.

When your webpage contains multiple scripts, `defer` will execute them in their relative order, while `async` will execute the one that is downloaded first, regardless of their order.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
