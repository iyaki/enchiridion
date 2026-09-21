---
title: "Sprites - Stateful sandboxes"
notion_id: 38c54f1c-7d23-81f9-8486-df4afbe73b21
notion_url: https://app.notion.com/p/Sprites-Stateful-sandboxes-38c54f1c7d2381f98486df4afbe73b21
last_edited: 2026-06-27T02:52:00.000Z
source_url: https://sprites.dev/
tags: ["English", "Web Development", "Frontend", "Javascript", "Framework/Library", "Tool", "sprites.dev"]
---
## Sandboxes aren't enough

Sprites are full Linux computers designed for agents. They're exactly as persistent and disposable as you want them to be, connect securely to external services via Connectors, and have full environment checkpointing and restore. Let's upgrade your agent's accommodations, shall we?

# Install the Sprites CLI

# Log in

# Create a new sprite

# Run a command in it

# Connect to the console

```plain text
# Create a new sprite curl -X PUT https://api.sprites.dev/v1/sprites/my-sprite \ -H "Authorization: Bearer $SPRITES_TOKEN" # Execute a command curl -X POST https://api.sprites.dev/v1/sprites/my-sprite/exec \ -H "Authorization: Bearer $SPRITES_TOKEN" \ -d '{"command": "echo hello"}'
```

```plain text
import { SpritesClient } from '@fly/sprites'; const client = new SpritesClient(process.env.SPRITES_TOKEN!); // Get a Sprite reference const sprite = client.sprite('my-sprite'); // Run a command! const { stdout } = await sprite.exec('echo hello'); console.log(stdout);
```

```plain text
import "github.com/superfly/sprites-go" client := sprites.New("your-auth-token") // Get a sprite handle sprite := client.Sprite("my-sprite") // Run a command - just like exec.Command! cmd := sprite.Command("echo", "hello", "world") output, err := cmd.Output() if err != nil { log.Fatal(err) } fmt.Printf("Output: %s", output)
```

```plain text
# Add to mix.exs {:sprites, github: "superfly/sprites-ex"} # Create a sprite client = Sprites.new(System.get_env("SPRITE_TOKEN")) Sprites.create(client, System.get_env("SPRITE_NAME")) # Run Python sprite = Sprites.sprite(client, System.get_env("SPRITE_NAME")) {output, _} = Sprites.cmd(sprite, "python", ["-c", "print(2+2)"]) IO.write(output)
```

```plain text
# Install: pip install sprites-py import os from sprites import SpritesClient # Create a sprite client = SpritesClient(os.environ["SPRITE_TOKEN"]) client.create_sprite(os.environ["SPRITE_NAME"]) # Run Python sprite = client.sprite(os.environ["SPRITE_NAME"]) output = sprite.command("python", "-c", "print(2+2)").output() print(output.decode(), end="")
```

## How Sprites work

A full Linux computer that keeps its disk, snapshots in a second, answers on its own URL, and reaches the outside world without holding a secret.

### Persistence

- Tiered storage Reads and writes hit a fast local cache, and your data lives durably in object storage behind it. That's what lets a Sprite sleep, move between machines, and come back with its filesystem intact.
- Same disk for every run Your files, installs, and data are on the same paths every run; the environment comes back exactly as you left it.
- Max compatibility It's a normal POSIX filesystem, so anything that reads or writes to disk just works, no special API or SDK.
- Room to work The volume is 100 GB, and you're billed on the storage you actually use rather than a size you have to pick up front. No sizing or resizing a volume.
- S3 Block Device A newer storage backend, in early access. It presents an object storage bucket to the kernel as a real block device and runs ext4 on top, which is what makes checkpoints block-level snapshots instead of file copies. Opt in per organization. Request early access

## Works with your agent. Yes, that one too.

Official plugins for every serious coding agent, plus SDKs and native integrations.

## Pricing

CPU Time Cumulative CPU usage measured by cpu.stat

$0.07 /CPU-hour

Memory Time Actual memory usage

$0.04375 /GB-hour

Storage Time Storage usage in GB-hours

HOT

$0.000683 /GB-hour

COLD

$0.000027 /GB-hour

### Examples

### Claude Code Session

4-hour coding session with bursts to 100% of 8 CPUs and 8 GB RAM, averaging 30% of 2 CPUs and 1.5 GB

CPU (2.4 CPU-hrs) $0.17

Memory (6 GB-hrs) $0.26

Hot storage (5 GB × 4 hrs) $0.01

Cold storage (10 GB × 4 hrs) $0.00

Total $0.44

### Web App

30 hours of wake time per month (~5 concurrent users avg), averaging 10% of 2 CPUs and 1 GB RAM

CPU (6 CPU-hrs) $0.42

Memory (30 GB-hrs) $1.31

Hot storage (3 GB × 30 hrs) $0.06

Cold storage (5 GB × 732 hrs) $0.10

Total $1.89 / month

## Need a truly wild number of Sprites? No problem.

Talk to us

## FAQ

- CPU, RAM, and hot storage, metered per hour of active use: CPU at $0.07/CPU-hour, RAM at $0.04375/GB-hour, and hot storage at $0.000683/GB-hour (≈ $0.50/GB-month). Nothing is charged per sprite. A sprite that exists but does nothing costs nothing beyond its storage.
- Sprites have three states: running (billed), warm (not billed), and cold (not billed). Running → warm happens within seconds of the idle monitor seeing no activity. Warm → cold can take much longer, but that doesn't matter for cost, since warm isn't billed either. So if you're seeing compute charges, your sprites are genuinely active.
- Four things reset the idle timer: an in-flight HTTP/API request, output to a session or exec'd process's stdout (redirecting to a file or detaching tmux doesn't count), an open TCP connection, or an active task (sprite-env tasks create, max 1 hour, renewable). More in the idle-detection docs.
- Plans bundle an allowance of CPU-hours, RAM GB-hours, and storage GB, not unlimited usage. Anything above the allowance bills at standard rates on top of the plan fee. Example: Hero ($100/mo) includes 1,200 CPU-hours, 4,800 RAM GB-hours, and 150 GB of storage; a real Hero account that used 1,554 CPU-hours and 28,386 RAM GB-hours in a month had a genuine overage.
- Each user can grant one $30 trial credit, and each org can receive at most one. So if you activate a second org yourself, you've already spent your grant and that org gets nothing. An org activated by a user who's never granted one still gets it.
- Compare your billable units against each plan's included hours; a bigger plan isn't automatically cheaper. One customer's usage was close to Mythic's 28,000 included RAM GB-hours, but at $2,000/mo the plan cost more than the overage it saved. RAM is almost always the line that dominates, so reducing per-sprite memory usually beats upgrading.
- Limits are per-org and scale with the plan. Hero allows 100 concurrently running sprites and 100 warm sprites; cold sprites are unlimited, and higher tiers raise those ceilings. If you hit the ceiling you'll see max sprites per org exceeded on provision. That's a limit, not an extra charge; nothing bills for being blocked. Sprite creation rate is also tiered: 10 sprites/minute on pay-as-you-go, rising with plan tier from 60/minute on Adventurer up to 240/minute on Mythic.
- The invoice is authoritative. Bills are only finalised at month end, and we'll reconcile against the underlying billing data and credit any discrepancy before then.
- No. Sprites bandwidth isn't metered today. (Any egress lines on your invoice come from Machines, not Sprites.)
- No. Credits apply to Sprites usage only. The monthly subscription fee is billed separately and isn't covered.
- From Hero upward, yes. Hero, Champion, and Legend include Standard email support; Epic and Mythic include Premium. Pay-as-you-go and the plans below Hero (Adventurer, Veteran) have community support. You don't need to buy a separate support subscription on top. If you're paying for both, tell us and we'll sort it out.

$30 in trial credits

Create, like, 500 Sprites free.
