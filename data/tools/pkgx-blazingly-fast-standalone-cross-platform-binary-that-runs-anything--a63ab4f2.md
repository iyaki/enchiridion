---
title: "pkgx - Blazingly fast, standalone, cross‐platform binary that runs anything"
notion_id: a63ab4f2-109a-45e8-8f4a-79394d4047ca
notion_url: https://app.notion.com/p/pkgx-Blazingly-fast-standalone-cross-platform-binary-that-runs-anything-a63ab4f2109a45e88f4a79394d4047ca
last_edited: 2023-10-12T21:16:00.000Z
source_url: https://pkgx.sh/
tags: ["Tool", "English", "Producer (Individual Contributor)", "Productivity", "Shell/Bash", "Programming", "Untried"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### `pkgx` is a blazingly fast, standalone, cross‐platform binary that _runs anything_

other ways to install

### What’s better than a package manager? No package manager

$ bun

command not found: bun

^^ type `pkgx` to run that

$ pkgx

running `bun`…

Bun: a fast JavaScript runtime, package manager, bundler and test runner.

# …

### Run Any Version

node 16? python 2? postgres 12? **nps**.

$ pkgx node@16

Node.js v16.20.1

>

### Zero System Impact

We don’t install packages. _We cache them_. Just like `npx` caches & executes node packages, `pkgx` caches everything else (including `npx`).

$ pkgx rustc --version

rustc 1.72.1

$ which rustc

command not found: rustc

# ^^ _still_ not installed!

### Shell Integration

`pkgx` can _optionally_ integrate with your shell giving it **pkging powers**.

When our investors ask why this is cool we just shrug and say “do you even dev?”.

$ deno

command not found: deno

$ env +deno^1.35

added ~/.pkgx/deno.land/v1.36.1

$ deno --version

deno 1.36.1

`deno`’s not installed, `pkgx` just added it to your shell session. It’ll be gone when you `exit` 🥹

Explore the docs

### “Max Howell, the mind behind Homebrew, is shaking up the foundation of development once again with his new creation, `pkgx`”

### Run Anywhere

### Wherever you work, `pkgx` works too.

macOS

- >= 11
- Intel and Apple Silicon

Linux

- glibc >=2.28
- `x86_64` & `arm64`

Windows

- WSL2
- _Native coming soon!_

Docker

Sure you could memorize the weird naming conventions of `apt`.

But wouldn’t you rather _just run?_

FROM ubuntu

RUN curl https://pkgx.sh | sh

RUN pkgx python@3.10 -m http.server 8000

CI/CD

- uses: pkgxdev/setup@v1

- run: pkgx npm@10 start

Other CI/CD Providers

Editors

Just Works™ in VSCode? nps.

The Deets on That

Scripts

Isn’t it time you had more than just Bash and POSIX in your scripts?

We got you.

#!/usr/bin/env -S pkgx +gh +gum fish

if gum confirm "Are you sure you want to deploy?"

gh release create v1.0.0

end

Explore the docs

### The Foundation of your Toolset

### The tools you need for work, where and when you need them.

$ cd myproj

myproj $ dev

found cargo.toml, package.json; env +cargo +npm

$ cargo build

Compiling myproj v0.1.0

# …

All the tools you need for work? Check.

Developer Environments

Developer environments provide the tools you need when working in those directories. When you step away—so do they.

`dev` docs

Reading your Keyfiles

`dev` works by examining the _keyfiles_ in your project root. If we see `cargo.toml` we know that means you want Rust. If we see `pyproject.toml` we know you want Python, but also we read the TOML to **determine what package manager you want**.

Constraining Your Dependencies

You can constrain your dependencies to a specific version, or a range of versions by adding YAML front matter to your project keyfiles.

Keyfile YAML Front MatterExplore the docs

### Trusted by 10k Engineers

### And built by them too.`pkgx` is open source through and through.

Packagers Who Care

We go the extra mile so you don’t have to.

- Our `git` is configured to ignore `.DS_Store` files 😍
- We configure package managers to install to `~/.local/bin` (and ensure that’s in the `PATH`) unless they can write to `/usr/local/bin`
- We automatically install great git extensions like `git absorb` _when you type them_
- We configure other version managers like `pyenv` to automatically install the Pythons you ask for

“The UNIX Philosophy” is in our DNA

We study and preach it, worship and practice it. It’s 100% who we are and everything we believe in and `pkgx` is UNIX through and through.

Open Source is in our DNA too

Our founder, Max Howell, created Homebrew the package manager used by tens of millions of developers around the world. He understands how important community is to open source.

He built it before.

He’s building it again.
