---
title: "sshx - Fast, collaborative live terminal sharing over the web"
notion_id: aa0ca645-b666-4011-9c22-d0d2bde0612b
notion_url: https://app.notion.com/p/sshx-Fast-collaborative-live-terminal-sharing-over-the-web-aa0ca645b66640119c22d0d2bde0612b
last_edited: 2023-11-09T19:06:00.000Z
source_url: https://github.com/ekzhang/sshx
tags: ["English", "Shell/Bash", "Office", "Untried", "Tool"]
---
# [sshx](https://github.com/ekzhang/sshx?tab=readme-ov-file#sshx)

A secure web-based, collaborative terminal.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**Features:**

- Run a single command to share your terminal with anyone.
- Resize, move windows, and freely zoom and pan on an infinite canvas.
- See other people's cursors moving in real time.
- Connect to the nearest server in a globally distributed mesh.
- End-to-end encryption with Argon2 and AES.
- Automatic reconnection and real-time latency estimates.
- Predictive echo for faster local editing (à la Mosh).

Visit [sshx.io](https://sshx.io/) to learn more.

## [Installation](https://github.com/ekzhang/sshx?tab=readme-ov-file#installation)

Just run this command to get the `sshx` binary for your platform.

```plain text
curl -sSf https://sshx.io/get | sh
```

Supports Linux and MacOS, on both x86_64 and arm64 architectures. The precompiled Linux binaries are statically linked.

### [CI/CD](https://github.com/ekzhang/sshx?tab=readme-ov-file#cicd)

You can also use sshx in continuous integration workflows to help debug tricky issues, like in GitHub Actions.

```plain text
name: CI
on: push

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      # ... other steps ...

      - run: curl -sSf https://sshx.io/get | sh && sshx
      #      ^
      #      └ This will open a remote terminal session and print the URL. It
      #        should take under a second.
```

We don't have a prepackaged action because it's just a single command. It works anywhere: GitLab CI, CircleCI, CI on your Raspberry Pi, etc.

## [Development](https://github.com/ekzhang/sshx?tab=readme-ov-file#development)

Here's how to work on the project, if you want to contribute.

### [Building from source](https://github.com/ekzhang/sshx?tab=readme-ov-file#building-from-source)

To build the latest version of the client from source, clone this repository and run, with [Rust](https://rust-lang.com/) installed:

```plain text
cargo install --path crates/sshx
```

This will compile the `sshx` binary and place it in your `~/.cargo/bin` folder.

### [Workflow](https://github.com/ekzhang/sshx?tab=readme-ov-file#workflow)

First, start service containers for development.

```plain text
docker compose up -d
```

Install [Rust 1.70+](https://www.rust-lang.org/), [Node v18](https://nodejs.org/), [NPM v9](https://www.npmjs.com/), and [mprocs](https://github.com/pvolok/mprocs). Then, run

```plain text
npm install
mprocs
```

This will compile and start the server, an instance of the client, and the web frontend in parallel on your machine.

## [Deployment](https://github.com/ekzhang/sshx?tab=readme-ov-file#deployment)

The application servers are deployed on [Fly.io](https://fly.io/).

```plain text
fly deploy
```
