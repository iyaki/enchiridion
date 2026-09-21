---
title: "Visual Studio Code"
notion_id: 4bb6ba4f-bdb6-4c29-9484-e3dd58f9fdfd
notion_url: https://app.notion.com/p/Visual-Studio-Code-4bb6ba4fbdb64c299484e3dd58f9fdfd
last_edited: 2023-04-25T13:59:00.000Z
source_url: https://code.visualstudio.com/
tags: ["Tool", "English", "Español", "Others", "Programming", "IDE/Extendable Text Editor"]
---
[https://code.visualstudio.com/](https://code.visualstudio.com/)

Your home for multi-agent development

MailList.tsx

```plain text
import { For, Show, createMemo, createSignal, onCleanup, onMount, } from "solid-js"; import { useMail } from "./mail"; import { Item } from "./Item"; const ROW = 56; const PAGE = 24; export function MailList() { const { items, query, load } = useMail(); const [focus, setFocus] = createSignal(0); const [top, setTop] = createSignal(0); const visible = createMemo(() => { const q = query().toLowerCase(); return items() .filter((m) => m.subject .toLowerCase() .includes(q), ) .slice(top(), top() + PAGE); }); const onKey = (e: KeyboardEvent) => { if (e.key === "ArrowDown") { setFocus((i) => Math.min(i + 1, PAGE - 1), ); } else if (e.key === "ArrowUp") { setFocus((i) => Math.max(i - 1, 0), ); } else if (e.key === "Enter") { open(visible()[focus()].id); } }; onMount(() => { load(); window.addEventListener( "keydown", onKey, ); onCleanup(() => window.removeEventListener( "keydown", onKey, ), ); }); return ( <Show when={visible().length} fallback={<Empty />} > <For each={visible()}> {(m, i) => ( <Item mail={m} active={i() === focus()} onOpen={open} /> )} </For> </Show> ); }
```

Refactor MailList component

MailList.tsx

MailListItem

MailList.tsx

useSelection.ts

MailListItem.tsx

+62

MailList.tsx

+18

−74

Mapped the MailList layout

useSelection

react-window

- Create MailListItem.tsx with props for message, selected, and onSelect.
- Wrap it in React.memo with a comparator on message.id + selected.
- Lift useSelection into MailList and pass handlers down.

Make sure arrow‑key focus still works after the split.

focus()

npm test -- MailList

41ms to 12ms

Queued

MailListToolbar

Describe what to build

Sessions

New Session

Refactor MailList component Thinking…

Wire MCP server for telemetry pipeline Editing mcp.config.ts

Add language server for .prisma files Running tests in parser.test.ts

Refactor command palette to fluent API +639 -323 · 2 hrs ago

Scaffold extension for inline chat actions 5 hrs ago

Theme tokens for high-contrast mode +752 -367 · yesterday

Generate typed SDK from OpenAPI 3.1 spec +218 -12 · 2 days ago

Diagnose flaky CI run on macOS arm64 4 days ago

Optimize Live Share session bootstrap +52 -1 · 3 days ago

Add devcontainer for monorepo workspace 6 days ago

Port debug adapter to DAP 1.65 1 wk ago

## Agents that build for you

Hand off tasks to AI agents that autonomously plan, make code changes, run commands, and iterate until the job is done.

For example, assign a CLI-based agent to triage and fix bugs in the background, interact with another agent to implement a feature using live validation in the integrated browser, and delegate a homepage redesign to a cloud agent that opens a pull request for your team to review.

Get started with agents

batch.go

```plain text
package http import ( "context" "encoding/json" "errors" "io" "log/slog" "mime/multipart" "net/http" "time" "golang.org/x/sync/errgroup" "golang.org/x/sync/semaphore" "go.opentelemetry.io/otel" ) type Result[T any] struct { Name string `json:"name"` Value T `json:"value,omitempty"` Error string `json:"error,omitempty"` } type Meta struct { Format string `json:"format"` Width int `json:"width"` Height int `json:"height"` Bytes int64 `json:"bytes"` } const ( perRequestTimeout = 30 * time.Second maxParallel = 8 ) func (s *Server) handleBatch( w http.ResponseWriter, r *http.Request, ) { ctx, span := otel.Tracer("http"). Start(r.Context(), "batch") defer span.End() ctx, cancel := context.WithTimeout( ctx, perRequestTimeout, ) defer cancel() r.Body = http.MaxBytesReader( w, r.Body, s.cfg.MaxBytes, ) if err := r.ParseMultipartForm( s.cfg.MaxBytes, ); err != nil { s.fail(w, http.StatusBadRequest, err) return } files := pickFiles(r.MultipartForm) if len(files) == 0 { s.fail(w, http.StatusBadRequest, errors.New("no files")) return } out := make( []Result[Meta], len(files), ) sem := semaphore.NewWeighted(maxParallel) g, gctx := errgroup.WithContext(ctx) for i, fh := range files { i, fh := i, fh if err := sem.Acquire(gctx, 1); err != nil { break } g.Go(func() error { defer sem.Release(1) out[i] = process(gctx, s, fh) return nil }) } _ = g.Wait() s.log.LogAttrs(ctx, slog.LevelInfo, "batch.done", slog.Int("count", len(out)), slog.Duration("budget", perRequestTimeout), ) w.Header().Set( "Content-Type", "application/json", ) _ = json.NewEncoder(w).Encode(out) } func process( ctx context.Context, s *Server, fh *multipart.FileHeader, ) Result[Meta] { res := Result[Meta]{Name: fh.Filename} f, err := fh.Open() if err != nil { res.Error = err.Error() return res } defer f.Close() data, err := io.ReadAll( io.LimitReader(f, s.cfg.MaxBytes), ) if err != nil { res.Error = err.Error() return res } m, err := s.proc.ExtractCtx(ctx, data) if err != nil { res.Error = err.Error() return res } res.Value = Meta{ Format: m.Format, Width: m.Width, Height: m.Height, Bytes: int64(len(data)), } return res }
```

Batch image processing endpoint

POST /process/batch

server.go

processor.go

config.go

batch.go

+128

server.go

+4

−0

batch_test.go

+96

Mapped the service layout

internal/http

Extract

ExtractCtx

errgroup

http.HandleFunc

internal/

- Add handleBatch in batch.go using a generic Result[Meta] type.
- Cap concurrency with semaphore.NewWeighted(8) and wrap fan‑out in errgroup.
- Wrap the request in an OTel span and a 30s context deadline.
- Wire the route in server.go and add table tests.

go test ./internal/http -run Batch -race

-race

184ms to 31ms

k6

Queued

p99_ms

Describe what to build next

## Any agent, any model

Use the agent harness that fits your workflow. Run agents locally or in the cloud, with Copilot or third-party providers like Claude and OpenAI.

Choose from dozens of models across providers, from fast completion models to advanced reasoning models. Or bring your own key to use any model from any provider.

The same Copilot works across surfaces — VS Code, the Copilot CLI, GitHub.com, and mobile. See how agents work across surfaces.

## All your sessions, one view

Stay productive while multiple agents work on tasks in parallel. Track all your agent sessions from a single view, regardless of where they run.

Quickly filter and monitor sessions, or dive into an individual agent interaction, without switching to a different tool or terminal.

## Your rules, your agents

Ensure agents follow your practices and team workflows. Define custom instructions, add agent skills, or build custom agents tailored to your project.

Connect to external tools and services with MCP servers, or install agent plugins or extensions to expand the agent's capabilities.

## A world-class code editor at its core

VS Code has been the editor of choice for millions of developers for over a decade. AI-powered inline suggestions, intelligent completions, and a rich editing experience make it just as powerful when you're writing code yourself.

Seamlessly switch between working with agents and hands-on coding, all within the same editor.

Get started with VS Code

main.py

```plain text
import numpy as np import pandas as pd iris_data = pd.read_csv("iris_dataset.csv") def describe(species: str) -> pd.Series: 7 subset = data[data["species"] == species] subset = iris_data[iris_data["species"] == species] if subset.empty: raise ValueError(f"{species} missing from sample") return subset[["petal", "sepal"]].agg(["mean", "std"]).loc["mean"] def summary(): 13 for species in np.sort(data["species"].unique()): for species in np.sort(iris_data["species"].unique()): try: stats = describe(species) except ValueError: print(f"{species}: no records") continue print(f"{species}: petal={stats['petal']:.2f} sepal={stats['sepal']:.2f}") if __name__ == "__main__": summary()
```

## Code with extensions

Extend your agents with tools from extensions and Model Context Protocol servers. Or, build your own extension to power your team's unique scenarios.

## Code in any language

VS Code supports almost every major programming language. Several ship in the box, like JavaScript, TypeScript, CSS, and HTML, but extensions for others can be found in the VS Code Marketplace.

JavaScript

TypeScript

Python

C#

C++

HTML

Java

JSON

PHP

Markdown

Powershell

YAML

## Fully customizable

Customize your VS Code UI and layout so that it fits your coding style.

Color themes let you modify the colors in VS Code's user interface to suit your preferences and work environment.

Settings Sync enables you to share your user settings across your VS Code instances with the Settings Sync feature.

Profiles let you create sets of customizations and quickly switch between them or share them with others.

## Code anywhere

Code wherever you're most productive, whether you're connected to the cloud, a remote repository, or in the browser with VS Code for the Web (vscode.dev).

Built-in Source Control empowers you with Git support out-of-the-box. Many other source control providers are available through extensions.

GitHub Codespaces provides cloud-powered development environments for any activity - whether it's a long-term project, or a short-term task like reviewing a pull request.

## Code with rich features

There's a lot more to an editor. Whether it's using built-in features or rich extensions, there's something for everyone.
