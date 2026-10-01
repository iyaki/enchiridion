---
title: "tinyjs — desktop apps for macOS, Windows, and Linux in ~6 MB"
notion_id: 3e654f1c-7d23-8144-a413-c250d84154c5
notion_url: https://app.notion.com/p/tinyjs-desktop-apps-for-macOS-Windows-and-Linux-in-6-MB-3e654f1c7d238144a413c250d84154c5
last_edited: 2026-09-25T03:23:00.000Z
source_url: https://tinyjs.app/
tags: ["Tool", "Article", "Official Website", "English", "Web Development", "Javascript", "Backend", "Automation"]
---
## Desktop appsin _~6 MB_.

A **JavaScript backend** with full system access, a
    **native webview window** (WebKit on macOS, WebView2 on
    Windows, WebKitGTK on Linux), and nothing else.
    No Electron. No bundled Chromium. No HTTP server. **No ports.**

`curl -fsSL https://tinyjs.app/install | sh`

macOS · apple silicon + intel · MIT

`irm https://tinyjs.app/install.ps1 | iex`

windows 10/11 · prebuilt, no compiler needed · just the WebView2 runtime (preinstalled on win 11) · adds tinyjs to your PATH (open a new terminal) · MIT

`curl -fsSL https://tinyjs.app/install | sh`

linux (beta) · x86_64 + arm64 · same script as macOS, detects Linux · needs libwebkit2gtk-4.1-0 (`sudo apt install libwebkit2gtk-4.1-0`) · MIT

## Shipped app size

Your backend, a ~600 KB window process, and your
    plain HTML/CSS/JS — ~6 MB on all three platforms. The window is the
    webview the OS already ships.

Tauri does the same, and does it well — but its
    backend is Rust, and you compile it. It also ships a window and little
    else: dialogs, clipboard, notifications, global shortcuts, filesystem,
    keychain, autostart, deep links are plugins you add. Here they're just
    the API, and the 6 MB is all of it.

## What you get

**zero ports**

Page ⇄ backend RPC over a Unix socket in a private temp dir. Nothing listens, nothing collides, nothing to scan.

**full system access**

Backend runs on txiki.js — files, sockets, processes, FFI, sqlite, fetch, WebSocket.

**hot reload**

Frontend edits swap into the live window; backend edits restart the process. No build step in dev.

**native chrome**

Real menu bar (About + Quit included), file panels, alert / confirm / prompt — answered by AppKit, not divs.

**signed .app bundles**

`tinyjs build` emits a codesigned, notarization-ready bundle with your icon.

**wrap a hosted app**

Point the window at a URL and it behaves like a browser — dialogs, downloads, popups, find — with a capability gate deciding what that origin may call.

**agent-ready**

Every project ships a skill file, so coding agents already know the whole API.

## See it running

![image](https://tinyjs.app/images/amp.webp)

**amp** _— a Winamp for the desktop_

Player, playlist, 10-band EQ and six visualizer engines, each pane
          a real native window that snaps, docks and windowshades. Plus a
          podcast deck and a fullscreen 80s hi-fi mode with VU needles and a
          world-radio globe. Plays **mp3 · m4a/aac · flac · wav · aiff ·
          caf · ogg/opus** — and, synthesized on the fly in-app, **MIDI**
          (real SoundFont banks) and **tracker modules** (mod · s3m · xm ·
          it, via the OpenMPT engine). A webview can't play any of those
          last two; amp renders them itself.

[⬇ .dmg · 7.4 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/amp-v0.12.0/amp-0.12.0.dmg)[⬇ .zip · 7.4 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/amp-v0.13.0/amp-0.13.0-win.zip)⬇ .tar.gz
              [x86_64](https://github.com/tarwin/tinyjsapp-examples/releases/download/amp-v0.12.0/amp-0.12.0-linux-x86_64.tar.gz)[arm64](https://github.com/tarwin/tinyjsapp-examples/releases/download/amp-v0.12.0/amp-0.12.0-linux-arm64.tar.gz)
            
            [source ↗](https://github.com/tarwin/tinyjsapp-examples/tree/main/amp)

![image](https://tinyjs.app/images/nib.webp)

**Nib** _— the best Markdown editor for docs_

One window per document, editor/preview split, clickable task
          boxes, themed PDF and HTML export, lossless closing. Double-click
          any `.md` in Finder and it opens
          here.

[⬇ .dmg · 5.6 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/nib-v0.2.0/nib-0.2.0.dmg)[⬇ .zip · 4.8 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/nib-v0.3.0/nib-0.3.0-win.zip)⬇ .tar.gz
              [x86_64](https://github.com/tarwin/tinyjsapp-examples/releases/download/nib-v0.2.0/nib-0.2.0-linux-x86_64.tar.gz)[arm64](https://github.com/tarwin/tinyjsapp-examples/releases/download/nib-v0.2.0/nib-0.2.0-linux-arm64.tar.gz)
            
            [source ↗](https://github.com/tarwin/tinyjsapp-examples/tree/main/nib)

![image](https://tinyjs.app/images/platter.webp)

**Platter** _— a record player, not a music player_

Albums, one side at a time, and deliberately slow. Point it at a
          music folder and the scanner reads it the way a person would:
          `Artist/Album/` trees, year
          prefixes stripped, **CD1**/**CD2** merged back into one LP,
          an artist's strays filed as _Singles_. No library to commit to
          — it ships with a sample record to spin, and connects Spotify if
          you'd rather bring your own.

[⬇ .dmg · 6.0 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/platter-v0.4.0/platter-0.4.0.dmg)[⬇ .zip · 4.9 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/platter-v0.4.0/platter-0.4.0-win.zip)⬇ .tar.gz
              [x86_64](https://github.com/tarwin/tinyjsapp-examples/releases/download/platter-v0.4.0/platter-0.4.0-linux-x86_64.tar.gz)[arm64](https://github.com/tarwin/tinyjsapp-examples/releases/download/platter-v0.4.0/platter-0.4.0-linux-arm64.tar.gz)
            
            [source ↗](https://github.com/tarwin/tinyjsapp-examples/tree/main/platter)

![image](https://tinyjs.app/images/worldclock.webp)

**World Clock** _— every city you care about, in the menu bar_

The tray is the app: your home city as an emoji, ticking every
          second — **🌉 4:45p** — and a click drops a vibrancy popover
          centred under the icon with each city's time, day offset and a
          day/night dot. It dismisses itself on focus loss like a real
          popover, and stops redrawing entirely while hidden — the backend
          alone keeps the bar ticking. Turn on cycling and it rotates through
          the cities.

**Best on macOS and Linux**, whose menu bar and app indicator
          can show live text. It runs on Windows, but a tray icon there
          carries no label — the clock becomes something you hover to read,
          which rather misses the point.

[⬇ .dmg · 4.4 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/worldclock-v0.3.4/worldclock-0.3.4.dmg)[⬇ .zip · 3.8 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/worldclock-v0.3.5/worldclock-0.3.5-win.zip)⬇ .tar.gz
              [x86_64](https://github.com/tarwin/tinyjsapp-examples/releases/download/worldclock-v0.3.5/worldclock-0.3.5-linux-x86_64.tar.gz)[arm64](https://github.com/tarwin/tinyjsapp-examples/releases/download/worldclock-v0.3.5/worldclock-0.3.5-linux-arm64.tar.gz)
            
            [source ↗](https://github.com/tarwin/tinyjsapp-examples/tree/main/worldclock)

![image](https://tinyjs.app/images/kitchen-sink.webp)

**Tiny Deck** _— the whole API on one deck_

Thirteen tabs of live demos: shell, files, HTTP, GPU, WASM, FFI,
          windows, tray, hotkeys, share sheets, screenshots, battery,
          clipboard, Spotlight. If you want to know what a call does, press it.

[⬇ .dmg · 5.6 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/kitchen-sink-v0.18.0/kitchen-sink-0.18.0.dmg)[⬇ .zip · 4.9 MB](https://github.com/tarwin/tinyjsapp-examples/releases/download/kitchen-sink-v0.17.0/kitchen-sink-0.17.0-win.zip)⬇ .tar.gz
              [x86_64](https://github.com/tarwin/tinyjsapp-examples/releases/download/kitchen-sink-v0.18.0/kitchen-sink-0.18.0-linux-x86_64.tar.gz)[arm64](https://github.com/tarwin/tinyjsapp-examples/releases/download/kitchen-sink-v0.18.0/kitchen-sink-0.18.0-linux-arm64.tar.gz)
            
            [source ↗](https://github.com/tarwin/tinyjsapp-examples/tree/main/kitchen-sink)

## Quick start

### terminal

```plain text
tinyjs new myapp
cd myapp
tinyjs dev
# window opens, hot reload
tinyjs build
# dist/myapp.app, signed
# (win/linux: dist/myapp.exe or dist/myapp)
```

### src/main.js — the whole backend

```plain text
export const api = {
  hello: async ({ name }) =>
    `hi ${name}`,
};
```

### in the page

```plain text
await tiny.api.call('hello', { name });
tiny.menu.set([…]);
await tiny.dialog.openFile();
```

### how it fits

**your page**your HTML, CSS and JS

_tiny.api.call · tiny.api.on_

**launcher**WebKit · WebView2 · WebKitGTK

_unix socket · named pipe on windows_

**backend**txiki.js

## Questions

|  | backend | window | ships | ports |
| --- | --- | --- | --- | --- |
| electron | Node.js | bundled Chromium | ≥ 150 MB | none |
| tauri | Rust, compiled | system webview | ~10 MB | none |
| neutralino | none — page-side only | system webview | ~3 MB | localhost ws |
| tinyjs | JavaScript (txiki.js) | system webview | ~6 MB | none |

Everything below is the honest version, caveats
    included — including the parts where another tool is the better answer.

**How is this different from Electron?**

**How is this different from Tauri?**

**What about Neutralino?**

**Will my app look the same on all three platforms?**

**Can I use React, Vue, Svelte, TypeScript, npm?**

**Do I need Node installed?**

**Can I build a Windows app from my Mac?**

**Is it production ready?**

**Can I wrap a website I already have?**

**How do I sign, notarize and ship updates?**

**What about iOS and Android?**

**Why ~6 MB and not 1 MB?**

![image](https://tinyjs.app/images/rakali-badge.svg)
