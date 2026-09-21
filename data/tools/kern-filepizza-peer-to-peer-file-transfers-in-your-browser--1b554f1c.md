---
title: "kern/filepizza - Peer-to-peer file transfers in your browser"
notion_id: 1b554f1c-7d23-8128-bec8-cb4cd4a5f234
notion_url: https://app.notion.com/p/kern-filepizza-Peer-to-peer-file-transfers-in-your-browser-1b554f1c7d238128bec8cb4cd4a5f234
last_edited: 2025-04-19T23:51:00.000Z
source_url: https://github.com/kern/filepizza
tags: ["Tool", "Service", "English", "Office"]
---
A hosted instance of FilePizza is available at [file.pizza](https://file.pizza/).

![image](https://camo.githubusercontent.com/dbaa632801e564800edb58e1fd300604cae6ba5784adf8b6bab89935e1453147/687474703a2f2f696d67732e786b63642e636f6d2f636f6d6963732f66696c655f7472616e736665722e706e67)

![image](https://github.com/kern/filepizza/raw/main/public/images/wordmark.png)

### Peer-to-peer file transfers in your browser

_Cooked up by _[_Alex Kern_](https://kern.io/)_ & _[_Neeraj Baid_](https://github.com/neerajbaid)_ while eating Sliver @ UC Berkeley._

Using [WebRTC](http://www.webrtc.org/), FilePizza eliminates the initial upload step required by other web-based file sharing services. Because data is never stored in an intermediary server, the transfer is fast, private, and secure.

A hosted instance of FilePizza is available at [file.pizza](https://file.pizza/).

## What's new with FilePizza v2

- A new UI with dark mode support, now built on modern browser technologies.
- Works on most mobile browsers, including Mobile Safari.
- Transfers are now directly from the uploader to the downloader's browser (WebRTC without WebTorrent) with faster handshakes.
- Uploaders can monitor the progress of the transfer and stop it if they want.
- Better security and safety measures with password protection and reporting.
- Support for uploading multiple files at once, which downloaders receive as a zip file.
- Streaming downloads with a Service Worker.
- Out-of-process storage of server state using Redis.

## Development

```shell
$ git clone https://github.com/kern/filepizza.git
$ pnpm install
$ pnpm dev
$ pnpm build
$ pnpm start
Running with Docker
```

```shell
$ pnpm docker:build
$ pnpm docker:up
$ pnpm docker:down
```

## Stack

- Next.js
- Tailwind
- TypeScript
- React
- PeerJS for WebRTC
- View Transitions
- Redis (optional)

## FAQ

**How are my files sent?** Your files are sent directly from your browser to the downloader's browser. They never pass through our servers. FilePizza uses WebRTC to send files. This requires that the uploader leave their browser window open until the transfer is complete.

**Can multiple people download my file at once?** Yes! Just send them your short or long URL.

**How big can my files be?** As big as your browser can handle.

**What happens when I close my browser?** The URLs for your files will no longer work. If a downloader has completed the transfer, that downloader will continue to seed to incomplete downloaders, but no new downloads may be initiated.

**Are my files encrypted?** Yes, all WebRTC communications are automatically encrypted using public-key cryptography because of DTLS. You can add an optional password to your upload for an extra layer of security.

## License & Acknowledgements

FilePizza is released under the [BSD 3-Clause license](https://github.com/kern/filepizza/blob/main/LICENSE). A huge thanks to [iblowyourdesign](https://dribbble.com/iblowyourdesign) for the pizza illustration.
