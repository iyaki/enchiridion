---
title: "monolith - CLI tool for saving complete web pages as a single HTML file"
notion_id: 2c52e50c-706e-45d8-92fb-46c6d742d98e
notion_url: https://app.notion.com/p/monolith-CLI-tool-for-saving-complete-web-pages-as-a-single-HTML-file-2c52e50c706e45d892fb46c6d742d98e
last_edited: 2024-03-25T17:37:00.000Z
source_url: https://github.com/Y2Z/monolith
tags: ["English", "Web Development", "HTML", "?", "Untried", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
 _____     ______________    __________      ___________________    ___
|     \   /              \  |          |    |                   |  |   |
|      \_/       __       \_|    __    |    |    ___     ___    |__|   |
|               |  |            |  |   |    |   |   |   |   |          |
|   |\     /|   |__|    _       |__|   |____|   |   |   |   |    __    |
|   | \___/ |          | \                      |   |   |   |   |  |   |
|___|       |__________|  \_____________________|   |___|   |___|  |___|

```

A data hoarder’s dream come true: bundle any web page into a single HTML file. You can finally replace that gazillion of open tabs with a gazillion of .html files stored somewhere on your precious little drive.

Unlike the conventional “Save page as”, `monolith` not only saves the target document, it embeds CSS, image, and JavaScript assets **all at once**, producing a single HTML5 document that is a joy to store and share.

If compared to saving websites with `wget -mpk`, this tool embeds all assets as data URLs and therefore lets browsers render the saved page exactly the way it was on the Internet, even when no network connection is available.

## Installation

### Using [Cargo](https://crates.io/crates/monolith) (cross-platform)

```plain text
cargo install monolith
```

### Via [Homebrew](https://formulae.brew.sh/formula/monolith) (macOS and GNU/Linux)

```plain text
brew install monolith
```

### Via [Chocolatey](https://community.chocolatey.org/packages/monolith) (Windows)

```plain text
choco install monolith
```

### Via [Scoop](https://scoop.sh/#/apps?q=monolith) (Windows)

```plain text
scoop install main/monolith
```

### Via [MacPorts](https://ports.macports.org/port/monolith/summary) (macOS)

```plain text
sudo port install monolith
```

### Using [Snapcraft](https://snapcraft.io/monolith) (GNU/Linux)

```plain text
snap install monolith
```

### Using [Guix](https://packages.guix.gnu.org/packages/monolith) (GNU/Linux)

```plain text
guix install monolith
```

### Using [AUR](https://aur.archlinux.org/packages/monolith) (Arch Linux)

```plain text
yay monolith
```

### Using [aports](https://pkgs.alpinelinux.org/packages?name=monolith) (Alpine Linux)

```plain text
apk add monolith
```

### Using [FreeBSD packages](https://svnweb.freebsd.org/ports/head/www/monolith/) (FreeBSD)

```plain text
pkg install monolith
```

### Using [FreeBSD ports](https://www.freshports.org/www/monolith/) (FreeBSD)

```plain text
cd /usr/ports/www/monolith/
make install clean
```

### Using [pkgsrc](https://pkgsrc.se/www/monolith) (NetBSD, OpenBSD, Haiku, etc)

```plain text
cd /usr/pkgsrc/www/monolith
make install clean
```

### Using [containers](https://www.docker.com/)

```plain text
docker build -t y2z/monolith .
sudo install -b dist/run-in-container.sh /usr/local/bin/monolith
```

### From [source](https://github.com/Y2Z/monolith)

Dependencies: `libssl` `cargo`

Install cargo (GNU/Linux) Check if cargo is installed

```plain text
cargo -v
```

If cargo is not already installed, install and add it to your existing `$PATH` (paraphrasing the [official installation instructions](https://doc.rust-lang.org/cargo/getting-started/installation.html)):

```plain text
curl https://sh.rustup.rs -sSf | sh
. "$HOME/.cargo/env"
```

Proceed with installing from source:

```plain text
git clone https://github.com/Y2Z/monolith.git
cd monolith
make install
```

### Using [pre-built binaries](https://github.com/Y2Z/monolith/releases) (Windows, ARM-based devices, etc)

Every release contains pre-built binaries for Windows, GNU/Linux, as well as platforms with non-standard CPU architecture.

## Usage

```plain text
monolith https://lyrics.github.io/db/P/Portishead/Dummy/Roads/ -o portishead-roads-lyrics.html
```

```plain text
cat index.html | monolith -aIiFfcMv -b https://original.site/ - > result.html
```

## Options

- `a`: Exclude audio sources
- `b`: Use custom `base URL`
- `B`: Forbid retrieving assets from specified domain(s)
- `c`: Exclude CSS
- `C`: Read cookies from `file`
- `d`: Allow retrieving assets only from specified `domain(s)`
- `e`: Ignore network errors
- `E`: Save document using custom `encoding`
- `f`: Omit frames
- `F`: Exclude web fonts
- `h`: Print help information
- `i`: Remove images
- `I`: Isolate the document
- `j`: Exclude JavaScript
- `k`: Accept invalid X.509 (TLS) certificates
- `M`: Don't add timestamp and URL information
- `n`: Extract contents of NOSCRIPT elements
- `o`: Write output to `file` (use “-” for STDOUT)
- `s`: Be quiet
- `t`: Adjust `network request timeout`
- `u`: Provide custom `User-Agent`
- `v`: Exclude videos

## Whitelisting and blacklisting domains

Options `-d` and `-B` provide control over what domains can be used to retrieve assets from, e.g.:

```plain text
monolith -I -d example.com -d www.example.com https://example.com -o example-only.html
```

```plain text
monolith -I -B -d .googleusercontent.com -d googleanalytics.com -d .google.com https://example.com -o example-no-ads.html
```

## Dynamic content

Monolith doesn't feature a JavaScript engine, hence websites that retrieve and display data after initial load may require usage of additional tools.

For example, Chromium (Chrome) can be used to act as a pre-processor for such pages:

```plain text
chromium --headless --incognito --dump-dom https://github.com | monolith - -I -b https://github.com -o github.html
```

## Proxies

Please set `https_proxy`, `http_proxy`, and `no_proxy` environment variables.

## Contributing

Please open an issue if something is wrong, that helps make this project better.

## Related projects

- Monolith Chrome Extension: [https://github.com/rhysd/monolith-of-web](https://github.com/rhysd/monolith-of-web)
- Pagesaver: [https://github.com/distributed-mind/pagesaver](https://github.com/distributed-mind/pagesaver)
- Personal WayBack Machine: [https://github.com/popey/pwbm](https://github.com/popey/pwbm)
- Hako: [https://github.com/dmpop/hako](https://github.com/dmpop/hako)
- Monk: [https://github.com/monk-dev/monk](https://github.com/monk-dev/monk)

## License

To the extent possible under law, the author(s) have dedicated all copyright related and neighboring rights to this software to the public domain worldwide. This software is distributed without any warranty.

Keep in mind that `monolith` is not aware of your browser’s session
