---
title: "steffest/DPaint-js: Webbased image editor, modeled after the legendary Deluxe Paint with a focus on retro Amiga file formats: read and write Amiga icon files and IFF ILBM images"
notion_id: 2f554f1c-7d23-8153-96d5-f02a670d4e94
notion_url: https://app.notion.com/p/steffest-DPaint-js-Webbased-image-editor-modeled-after-the-legendary-Deluxe-Paint-with-a-focus-on--2f554f1c7d23815396d5f02a670d4e94
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://github.com/steffest/DPaint-js
tags: ["English", "Web Development", "Javascript", "Image Editing", "Open Source", "Tool", "Article", "GitHub"]
---
# DPaint.js

Webbased image editor modeled after the legendary [Deluxe Paint](https://en.wikipedia.org/wiki/Deluxe_Paint) with a focus on retro Amiga file formats. Next to modern image formats, DPaint.js can read and write Amiga icon files and IFF ILBM images.

![image](https://github.com/steffest/DPaint-js/raw/master/_img/dpaint-logo.png?raw=true)

Online version available at [https://www.stef.be/dpaint/](https://www.stef.be/dpaint/)

![image](https://github.com/steffest/DPaint-js/raw/master/_img/ui.png?raw=true)

## Main Features

- Fully Featured image editor with a.o. 
- Heavy focus on colour reduction with fine-grained dithering options
- Amiga focus 
- Deluxe Paint Legacy 

## Free and Open

It runs in your browser, works on any system and works fine on touch-screen devices like iPads.

It is written in 100% plain JavaScript and has no dependencies.

It's 100% free, no ads, no tracking, no accounts, no nothing.

All processing is done in your browser, no data is sent to any server.

The only part that is not included in this repository is the Amiga Emulator Files. (The emulator is based on the [Scripted Amiga Emulator](https://github.com/naTmeg/ScriptedAmigaEmulator))

## Building

DPaint.js doesn't need building.

It also has zero dependencies so there's no need to install anything.

DPaint.js is written using ES6 modules and runs out of the box in modern browsers.

Just serve "index.html" from a webserver and you're good to go.

There's an optional build step to create a compact version of DPaint.js if you like.

I'm using [Parcel.js](https://parceljs.org/) for this.

For convenience, I've included a "package.json" file.

open a terminal and run `npm install` to install Parcel.js and its dependencies. Then run `npm run build` to create a compact version of DPaint.js in the "dist" folder.

## Documentation

Documentation can be found at [https://www.stef.be/dpaint/docs/](https://www.stef.be/dpaint/docs/)

## Running offline

Dpaint.js is a web application, not an app that you install on your computer. That being said: DPaint.js has no online dependencies and runs fine offline if you want. One caveat: you have to serve the index.html file from a webserver, not just open it in your browser.

A quick way to do this is - for example - using the [Spark](https://github.com/rif/spark/releases) app.

[Download the binary](https://github.com/rif/spark/releases) for your platform, drop the Spark executable in the folder where you downloaded the Dpaint.js source files and run it. If you then point your browser to [http://localhost:8080/](http://localhost:8080/) it should work.

If you are using Chrome, you can also "install" dpaint.js as app.

![image](https://private-user-images.githubusercontent.com/763047/357537795-fa4a1e8b-4e45-4fe1-9d77-8b1e3364e867.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3Njk0NDQ1NjcsIm5iZiI6MTc2OTQ0NDI2NywicGF0aCI6Ii83NjMwNDcvMzU3NTM3Nzk1LWZhNGExZThiLTRlNDUtNGZlMS05ZDc3LThiMWUzMzY0ZTg2Ny5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjYwMTI2JTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI2MDEyNlQxNjE3NDdaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT0wMWNkOTU1ZmRlYTEyNTJjNzI3OWMwYmU1NzFjNzg5NDRjZDU1MzFmNDgxYjIzZGI3ZDY2YTZlZDBhYjk2NmVmJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCJ9.EscgT4jCLsFX2xa15eHzr8UqJcRXzIYwhk7LJ79MbqE)

It will then show up your Chrome apps and work offline.

## Contributing

Current version is still alpha.

I'm sure there are bugs and missing features.

Bug reports and pull requests are welcome.

### Missing Features

Planned for the next release, already in the works:

- Color Cycling (done)
- Animation support (GIf and Amiga ANIM files) (done)
- Shading/transparency tools that stay within the palette. (done)

Planned for a future release if there's a need for it.

- Support for non-square pixel modes such as HiRes and Interlaced
- PSD import and export
- SpriteSheet support
- Write HAM,SHAM and Dynamic HiRes images
- Commodore 64 graphics modes

## Browser Quirks

Please note that the **Brave** browser is using "[farbling](https://brave.com/privacy-updates/4-fingerprinting-defenses-2.0/#2-fingerprinting-protections-20-farbling-for-great-good)" that introduces random image noise in certain conditions. They claim this is to protect your privacy. Although I totally understand the sentiment, In my opinion a browser should not actively alter the content of a webpage or intentionally break functionality.

But hey, who am I to speak, it's a free world. Just be aware that if you are using Brave, you will run into issues, so please "lower your shields" for this app in Brave or use another browser.

## Color Cycling

Dpaint.js supports Color-Cycling - a long lost art of "animating" a static image by only rotating some colors in the palette. See an example here:

**The_Vision_cycle.mp4**

[Open the layered source file of the above image directly in Dpaint.js](https://www.dpaint.app/?file=gallery%2F2026%2Fthe-vision-layered.json&play=true)
