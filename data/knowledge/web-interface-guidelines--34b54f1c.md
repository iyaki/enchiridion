---
title: "Web Interface Guidelines"
notion_id: 34b54f1c-7d23-81b3-9f81-fb8cbed1faf6
notion_url: https://app.notion.com/p/Web-Interface-Guidelines-34b54f1c7d2381b39f81fb8cbed1faf6
last_edited: 2026-04-23T02:28:00.000Z
source_url: https://vercel.com/design/guidelines
tags: ["Article", "Guide", "Vercel", "English", "UI/UX", "Web Development", "Frontend", "Productivity", "Design"]
---
Interfaces succeed because of hundreds of choices. This is a living, non-exhaustive list of those decisions. Most guidelines are framework-agnostic, some specific to React/Next.js. [Feedback is welcome](https://github.com/vercel-labs/web-interface-guidelines/tree/main).

- [**Clear focus.**](https://vercel.com/design/guidelines/#clear-focus) Every focusable element shows a visible focus ring. Prefer `:focus-visible` over `:focus` to avoid distracting pointer users. Set `:focus-within` for grouped controls.
- [**Match visual & hit targets.**](https://vercel.com/design/guidelines/#match-visual-hit-targets) Exception: if the visual target is < 24px, expand its hit target to ≥ 24px. On mobile, the minimum size is 44px.
- [**Mobile input size.**](https://vercel.com/design/guidelines/#mobile-input-size) `<input>` font size is ≥ 16px on mobile to prevent iOS Safari auto-zoom/pan on focus. Or set `<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />`.
- [**Minimum loading-state duration.**](https://vercel.com/design/guidelines/#minimum-loading-state-duration) If you show a spinner/skeleton, add a short show-delay (~150–300 ms) & a minimum visible time (~300–500 ms) to avoid flicker on fast responses. The `<Suspense>` component in React does this automatically.
- [**URL as state.**](https://vercel.com/design/guidelines/#url-as-state) Persist state in the URL so share, refresh, Back/Forward navigation work e.g., [nuqs](https://nuqs.dev/).
- [**Optimistic updates.**](https://vercel.com/design/guidelines/#optimistic-updates) Update the UI immediately when success is likely; reconcile on server response. On failure, show an error & roll back or provide Undo.
- [**Ellipsis for further input & loading states.**](https://vercel.com/design/guidelines/#ellipsis-for-further-input-loading-states) Menu options that open a follow-up e.g., "Rename…" & loading/processing states e.g., "Loading…", "Saving…", "Generating…" end with an ellipsis.
- [**Design forgiving interactions.**](https://vercel.com/design/guidelines/#design-forgiving-interactions) Controls minimize finickiness with generous hit targets, clear affordances, & predictable interactions, e.g., [prediction cones](https://x.com/JohnPhamous/status/1657083267299028992).
- [**Overscroll behavior.**](https://vercel.com/design/guidelines/#overscroll-behavior) Set `overscroll-behavior: contain` intentionally e.g., in modals/drawers.
- [**Autofocus for speed.**](https://vercel.com/design/guidelines/#autofocus-for-speed) On desktop screens with a single primary input, autofocus. Rarely autofocus on mobile because the keyboard opening can cause layout shift.
- [**No dead zones.**](https://vercel.com/design/guidelines/#no-dead-zones) If part of a control looks interactive, it should be interactive. Don’t leave users guessing where to interact.
- [**Deep-link everything.**](https://vercel.com/design/guidelines/#deep-link-everything) Filters, tabs, pagination, expanded panels, anytime `useState` is used.
- [**Clean drag interactions.**](https://vercel.com/design/guidelines/#clean-drag-interactions) Disable text selection & apply [`inert`](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/inert) (which prevents interaction) while an element is dragged so selection/hover don't occur simultaneously.
- [**Links are links.**](https://vercel.com/design/guidelines/#links-are-links) Use `<a>` or `<Link>` for navigation so standard browser behaviors work (Cmd/Ctrl+Click, middle-click, right-click to open in a new tab). Never substitute with `<button>` or `<div>` for navigational links.
- [**Implementation preference.**](https://vercel.com/design/guidelines/#implementation-preference) Prefer CSS, avoid main-thread JS-driven animations when possible.
- [**Compositor-friendly.**](https://vercel.com/design/guidelines/#compositor-friendly) Prioritize GPU-accelerated properties (`transform`, `opacity`) & avoid properties that trigger reflows/repaints (`width`, `height`, `top`, `left`).
- [**Necessity check.**](https://vercel.com/design/guidelines/#necessity-check) Only animate when it clarifies cause & effect or when it adds deliberate delight e.g., [the northern lights](https://x.com/JohnPhamous/status/1831380516509278561).
- [**Never **](https://vercel.com/design/guidelines/#never-transition-all)[**`transition: all`**](https://vercel.com/design/guidelines/#never-transition-all)[**.**](https://vercel.com/design/guidelines/#never-transition-all) Explicitly list only the properties you intend to animate (typically `opacity`, `transform`). `all` can unintentionally animate layout-affecting properties causing jank.
- [**Cross-browser SVG transforms.**](https://vercel.com/design/guidelines/#cross-browser-svg-transforms) Apply CSS transforms/animations to `<g>` wrappers & set `transform-box: fill-box; transform-origin: center;`. Safari historically had bugs with transform-origin on SVG & grouping avoids origin miscalculation.
- [**Deliberate alignment.**](https://vercel.com/design/guidelines/#deliberate-alignment) Every element aligns with something intentionally whether to a grid, baseline, edge, or optical center. No accidental positioning.
- [**Balance contrast in lockups.**](https://vercel.com/design/guidelines/#balance-contrast-in-lockups) When text & icons sit side by side, adjust weight, size, spacing, or color so they don’t clash. For example, a thin-stroke icon may need a bolder stroke next to medium-weight text.
- [**Responsive coverage.**](https://vercel.com/design/guidelines/#responsive-coverage) Verify on mobile, laptop, & ultra-wide. For ultra-wide, zoom out to 50% to simulate.
- [**No excessive scrollbars.**](https://vercel.com/design/guidelines/#no-excessive-scrollbars) Only render useful scrollbars; fix overflow issues to prevent unwanted scrollbars. On macOS set ["Show scroll bars" to "Always"](https://support.apple.com/guide/mac-help/change-appearance-settings-mchlp1225/mac#:~:text=or%20status%20bars.-,Show%20scroll%20bars,-Scroll%20bars%20appear) to test what Windows users would see.
- [**Let the browser size things.**](https://vercel.com/design/guidelines/#let-the-browser-size-things) Prefer flex/grid/intrinsic layout over measuring in JS. Avoid layout thrash by letting CSS handle flow, wrapping, & alignment.
- [**Icons have labels.**](https://vercel.com/design/guidelines/#icons-have-labels) Convey the same meaning with text for non-sighted users.
- [**Don’t ship the schema.**](https://vercel.com/design/guidelines/#dont-ship-the-schema) Visual layouts may omit visible labels, but accessible names/labels still exist for assistive tech.
- [**Locale-aware formats.**](https://vercel.com/design/guidelines/#locale-aware-formats) Format dates, times, numbers, delimiters, & currencies for the user’s locale.
- [**Prefer language settings over location.**](https://vercel.com/design/guidelines/#prefer-language-settings-over-location) Detect language via `Accept-Language` header & `navigator.languages`. Never rely on IP/GPS for language.
- [**Non-breaking spaces for glued terms.**](https://vercel.com/design/guidelines/#non-breaking-spaces-for-glued-terms) Use a non-breaking space &nbsp; to keep units, shortcuts & names together: `10 MB` → `10&nbsp;MB`, `⌘ + K` → `⌘&nbsp;+&nbsp;K`, `Vercel SDK` → `Vercel&nbsp;SDK`. Use `&#x2060;` for no space.
- [**Enter submits.**](https://vercel.com/design/guidelines/#enter-submits) When a text input is focused, Enter submits if it's the only control. If there are many controls, apply to the last control.
- [**Labels everywhere.**](https://vercel.com/design/guidelines/#labels-everywhere) Every control has a `<label>` or is associated with a label for assistive tech.
- [**Submission rule.**](https://vercel.com/design/guidelines/#submission-rule) Keep submit enabled until submission starts; then disable during the in-flight request, show a spinner, & include an idempotency key.
- [**Don’t block typing.**](https://vercel.com/design/guidelines/#dont-block-typing) Even if a field only accepts numbers, allow any input & show validation feedback. Blocking keystrokes entirely is confusing because the user gets no explanation.
- [**No dead zones on controls.**](https://vercel.com/design/guidelines/#no-dead-zones-on-controls) Checkboxes & radios avoid dead zones; the label & control share a single generous hit target.
- [**Error placement.**](https://vercel.com/design/guidelines/#error-placement) Show errors next to their fields; on submit, focus the first error.
- [**Don’t trigger password managers for non-auth fields.**](https://vercel.com/design/guidelines/#dont-trigger-password-managers-for-non-auth-fields) For inputs like “Search” avoid reserved names (e.g., password), use `autocomplete="off"` or a specific token like `autocomplete="one-time-code"` for OTP fields.
- [**Text replacements & expansions.**](https://vercel.com/design/guidelines/#text-replacements-expansions) Some input methods add trailing whitespace. The input should trim the value to avoid showing a confusing error message.
- [**Windows **](https://vercel.com/design/guidelines/#windows-select-background)[**`<select>`**](https://vercel.com/design/guidelines/#windows-select-background)[** background.**](https://vercel.com/design/guidelines/#windows-select-background) Explicitly set `background-color` & `color` on native `<select>` to avoid dark-mode contrast bugs on Windows.
- [**Preconnect to origins.**](https://vercel.com/design/guidelines/#preconnect-to-origins) Use `<link rel="preconnect">` for asset/CDN domains (with crossorigin when needed) to reduce DNS/TLS latency.
- [**Subset fonts.**](https://vercel.com/design/guidelines/#subset-fonts) Ship only the code points/scripts you use via unicode-range (limit variable axes to what you need) to shrink size.
- [**Don’t use the main thread for expensive work.**](https://vercel.com/design/guidelines/#dont-use-the-main-thread-for-expensive-work) Move especially long tasks to [Web Workers](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API) to avoid blocking interaction with the page.
- [**Hue consistency.**](https://vercel.com/design/guidelines/#hue-consistency) On non-neutral backgrounds, tint borders/shadows/text toward the same hue.
- [**Set the appropriate color-scheme.**](https://vercel.com/design/guidelines/#set-the-appropriate-color-scheme) Style the `<html>` tag with `color-scheme: dark` in dark themes so that scrollbars and other device UI have proper contrast.
- [**Text anti-aliasing & transforms.**](https://vercel.com/design/guidelines/#text-anti-aliasing-transforms) Scaling text can change smoothing. Prefer animating a wrapper instead of the text node. If artifacts persist set `translateZ(0)` or `will-change: transform` to promote to its own layer.
- [**Avoid gradient banding.**](https://vercel.com/design/guidelines/#avoid-gradient-banding) Fading content to dark colors using css masks can cause banding. [Background images can be used instead](https://x.com/JohnPhamous/status/1724491202148675590).

These preferences reflect Vercel’s brand & product choices. They aren’t universal guidelines.

- [**Active voice.**](https://vercel.com/design/guidelines/#active-voice) 
- [**Action-oriented language.**](https://vercel.com/design/guidelines/#action-oriented-language) 
- [**Use consistent placeholders.**](https://vercel.com/design/guidelines/#use-consistent-placeholders)
- [**Consistent currency formatting.**](https://vercel.com/design/guidelines/#consistent-currency-formatting) In any given context, display currency with either 0 or 2 decimal places, never mix both.
- [**Separate numbers & units with a space.**](https://vercel.com/design/guidelines/#separate-numbers-units-with-a-space) 
- [**Default to positive language.**](https://vercel.com/design/guidelines/#default-to-positive-language) Frame messages in an encouraging, problem-solving way, even for errors. 
- [**Error messages guide the exit.**](https://vercel.com/design/guidelines/#error-messages-guide-the-exit) Don’t just state what went wrong—tell the user how to fix it. 
- [**Avoid ambiguity.**](https://vercel.com/design/guidelines/#avoid-ambiguity) Labels are clear & specific.

Use these guidelines with AI coding agents. Audit all generated interfaces.

Install `/web-interface-guidelines` to review UI code:

Supports Antigravity, Claude Code, Cursor, Gemini CLI, OpenCode, & Windsurf.

For other agents, use the [command prompt](https://raw.githubusercontent.com/vercel-labs/web-interface-guidelines/main/command.md) directly.

Add [AGENTS.md](https://agents.md/) to your project so agents apply these guidelines during generation.

We’re hiring people who live for these details. [Check out the job postings](https://vercel.com/careers?function=Design).
