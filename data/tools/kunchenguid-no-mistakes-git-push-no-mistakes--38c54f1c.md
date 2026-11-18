---
title: "kunchenguid/no-mistakes: git push no-mistakes"
notion_id: 38c54f1c-7d23-813b-8cf0-cb074fb5f591
notion_url: https://app.notion.com/p/kunchenguid-no-mistakes-git-push-no-mistakes-38c54f1c7d23813b8cf0cb074fb5f591
last_edited: 2026-06-27T02:52:00.000Z
source_url: https://github.com/kunchenguid/no-mistakes
tags: ["Tool", "Guide", "GitHub", "English", "Git", "DevOps", "Continuous Integration/Continuous Delivery", "Automation", "Programming"]
---
# `git push no-mistakes`

![image](https://camo.githubusercontent.com/a6e7364de9ef0dfef6f6b34f51217229c04d6dd7a269f3572dd707b20dd75604/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f616374696f6e732f776f726b666c6f772f7374617475732f6b756e6368656e677569642f6e6f2d6d697374616b65732f72656c656173652e796d6c3f7374796c653d666c61742d737175617265266c6162656c3d72656c65617365)

![image](https://camo.githubusercontent.com/5932bcc57e63b240092f97c5f94ba4a837456b4d666b9606f0d18f0bfe047a5a/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f706c6174666f726d2d6d61634f532532302537432532304c696e757825323025374325323057696e646f77732d626c75653f7374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/758198a034b7ac6257103933f6320aab2366ec20e24f6c1dfea5a105eeefd1ce/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f582d406b756e6368656e677569642d626c61636b3f7374796c653d666c61742d737175617265)

![image](https://camo.githubusercontent.com/3eb64fccfb56bff0f7bcb6f4f7c563d48298802e5ebaab7604de88ba219a3338/68747470733a2f2f696d672e736869656c64732e696f2f646973636f72642f313433393930313833313033383736333039323f7374796c653d666c61742d737175617265266c6162656c3d646973636f7264)

### Kill all the slop. Raise clean PR.

![image](https://raw.githubusercontent.com/kunchenguid/no-mistakes/main/demo.gif)

`no-mistakes` puts a local git proxy in front of your real remote. Push to `no-mistakes` instead of `origin`, and it spins up a disposable worktree, runs an AI-driven validation pipeline, forwards upstream only after every check passes, and opens a clean PR automatically.

- **Non-blocking** - the pipeline runs in an isolated worktree without disrupting your work.
- **Agent-agnostic** - `claude`, `codex`, `rovodev`, `opencode`, `pi`, or `acp:<target>` via `acpx`.
- **Agent-native** - `/no-mistakes` lets the coding agent that wrote the change also gate it: it runs the pipeline, applies the safe fixes, and escalates the rest to you.
- **Human stays in charge** - auto-fix or review findings, your call.
- **Clean PRs by default** - push, open PR, watch CI, and auto-fix failures in one shot.

Full documentation: [https://kunchenguid.github.io/no-mistakes/](https://kunchenguid.github.io/no-mistakes/)

## Install

```plain text
curl -fsSL https://raw.githubusercontent.com/kunchenguid/no-mistakes/main/docs/install.sh | sh
```

Windows, Go install, and build-from-source instructions are in the [installation guide](https://kunchenguid.github.io/no-mistakes/start-here/installation/).

## Quick Start

```plain text
$ no-mistakes init
  ✓ Gate initialized

    repo  /Users/you/src/my-repo
    gate  no-mistakes → /Users/you/.no-mistakes/repos/abc123def456.git
  remote  git@github.com:you/my-repo.git
   skill  /no-mistakes installed for agents

  Push through the gate with:
  git push no-mistakes <branch>

$ git checkout my-branch

# do some work in the branch...

$ git push no-mistakes
  * Pipeline started

  Run no-mistakes to review.

$ no-mistakes
# opens the TUI for the active run
```

## Three ways to trigger the gate

Every change runs through the same pipeline. Pick the entry point that fits how you're working when the change is ready:

- **`git push no-mistakes`** - the explicit Git path. Push a committed branch to the gate remote instead of `origin`.
- **`no-mistakes`** - the TUI. Run it after making changes (no commit needed) and a wizard walks you through creating a branch, committing, and pushing through the gate, then attaches to the run. `no-mistakes -y` does all of that automatically.
- **`/no-mistakes`** - the agent skill. Tell the coding agent that wrote the change to gate it. It runs the pipeline, applies the safe fixes itself, and stops to ask you about anything that needs a human call.

`no-mistakes init` installs the `/no-mistakes` skill for Claude Code and other agents. Under the hood the skill drives `no-mistakes axi`, a non-interactive TOON interface to the same approval flow.

See the [quick start](https://kunchenguid.github.io/no-mistakes/start-here/quick-start/) for the full first-run walkthrough.

## Development

```plain text
make build   # Build bin/no-mistakes with version info
make test    # Run go test -race ./... (excludes the e2e suite)
make e2e     # Run the tagged end-to-end agent journey suite
make e2e-record # Re-record e2e fixtures when agent wire formats change
make lint    # Check generated skill drift and run go vet ./...
make skill   # Regenerate skills/no-mistakes/SKILL.md
make fmt     # Run gofmt -w .
make demo    # Regenerate demo.gif and demo.mp4 (needs vhs and ffmpeg)
make docs    # Build the Astro docs site in docs/dist
```

See `Makefile` for the full target list.

`make e2e-record` overwrites `internal/e2e/fixtures/` from the real `claude`, `codex`, and `opencode` CLIs, spends real API quota, and should be reviewed before committing.
