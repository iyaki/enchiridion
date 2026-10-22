---
title: "Fir - Reactive html apps in Go & alpine.js"
notion_id: f3c4bedc-f1c7-46e3-ab22-efb754a176e9
notion_url: https://app.notion.com/p/Fir-Reactive-html-apps-in-Go-alpine-js-f3c4bedcf1c746e3ab22efb754a176e9
last_edited: 2026-09-21T16:59:00.000Z
source_url: https://www.notion.so/f3c4bedcf1c746e3ab22efb754a176e9
tags: ["Framework/Library", "English", "Go", "Web Development", "Untried"]
---
The **Fir** toolkit is designed for Go developers with moderate html/css & js skills who want to progressively build reactive web apps without mastering complex web frameworks. It includes a Go library and an Alpine.js plugin.

Scroll below to see a demo & a quickstart guide _↓_ or read about [how it works](https://adnaan.notion.site/Fir-2358531aced84bf1b0b1a687760fff3b).

## 1. Start with html/template

A Fir web page begins as a standard server-side rendered page which reloads to show the new state on user interaction. Click the buttons below to see it in action.

Count: 133

When the html form is submitted, its is handled on the server in `onEvent` functions registered for inc and dec. The action/formaction attribute must be of the format `/?event=inc` where `inc` is the event name. Notice there is **no javascript** in the page.

Go the the directory where you have these files and run:

```plain text
1go run counter.go
```

Open your browser and go to [http://localhost:9867](http://localhost:9867/) to see the counter in action.

- counter.html
-  counter.go

```plain text
 1<!DOCTYPE html>
 2<html lang="en">
 3    <body>
 4        {{ block "count" . }}
 5            <div>Count: {{ .count }}</div>
 6        {{ end }}
 7        <form method="post">
 8            <button
 9                formaction="/?event=inc"
10                type="submit">
11                +
12            </button>
13            <button
14                formaction="/?event=dec"
15                type="submit">
16                -
17            </button>
18        </form>
19    </body>
20</html>

```

## 2. Enhance with [alpinejs](https://alpinejs.dev/)

Later we use fir’s alpinejs plugin to enhance the form submission and receive the re-rendered template as an event. The event is handled by the `$fir.replace()` helper function which updates the inner content of the div on which the event listener is declared. Click the buttons below to see reactivity in action. Open this page in two tabs to see the changes in one tab reflected in the other.

As you can notice, the count value is updated without a page reload. The server side code remains unchanged.

Fir’s magic expression `@fir:event-name:event-state::template-name` piggybacks on [alpinejs event binding syntax](https://alpinejs.dev/directives/on#custom-events) to declare [html/templates](https://pkg.go.dev/html/template) to be re-rendered on the server.

```plain text
1<div
2    @fir:inc:ok::count="$fir.replace()"
3    @fir:dec:ok::count="$fir.replace()">
4    {{ block "count" . }}
5        <div>Count: {{ .count }}</div>
6    {{ end }}
7</div>

```

If the handler response status for event `inc` is `ok` then re-render the template named `count` on the server and return the html output to the event listener as a [CustomEvent](https://developer.mozilla.org/en-US/docs/Web/API/CustomEvent). The [CustomEvent.detail](https://developer.mozilla.org/en-US/docs/Web/API/CustomEvent/detail) property is used by the alpinejs plugin helper `$fir.replace()` to update the div on which the listener is declared.

Alternatively there is a short-hand form for wiring up multiple events with the same action.

```plain text
1<div @fir:[inc:ok,dec:ok]::count="$fir.replace()">
2    {{ block "count" . }}
3    <div>Count: {{ .count }}</div>
4    {{ end }}
5</div>

```

Go the the directory where you have these files and run:

```plain text
1go run counter.go

```

Open your browser and go to [http://localhost:9867](http://localhost:9867/) to see the reactive counter in action.

- counter.html
-  counter.go

```plain text
 1<!DOCTYPE html>
 2<html lang="en">
 3    <head>
 4        <script
 5            defer
 6            src="https://unpkg.com/@livefir/fir@latest/dist/fir.min.js"></script>
 7 8        <script
 9            defer
10            src="https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"></script>

```
