---
title: "MicroLighter - A zero-dep syntax highlighter"
notion_id: 3da54f1c-7d23-81c0-95fd-cc8a5c98efec
notion_url: https://app.notion.com/p/MicroLighter-A-zero-dep-syntax-highlighter-3da54f1c7d2381c095fdcc8a5c98efec
last_edited: 2026-09-13T03:05:00.000Z
source_url: https://davatron5000.github.io/microlighter/
tags: ["Tool", "dev.to", "English", "Web Development", "Frontend", "Performance"]
---
## No more spans!

## Blazingly fast!!

## Lazy-loaded grammars!!!

## CSS for humans!!!1

## Install

Add MicroLighter to your project, import a theme, then mark each code block with a `language-*` class.

```plain text
npm install microlighter
```

```plain text
import { highlightAll } from "microlighter";
import "microlighter/themes/github.css";
await highlightAll();
```

```plain text
<pre><code class="language-javascript">const answer = 42;</code></pre>
```

Prefer automatic setup? Import the minified auto-runner instead. It highlights the page immediately and re-highlights when a `syntax-highlight` event is dispatched.

```plain text
import "microlighter/microlighter.min.js";
```

Want copy controls or line numbers? Import the custom element bundle and wrap the same standard `<pre><code>` markup.

```plain text
import "microlighter/micro-lighter-element.min.js";
```

```plain text
<micro-lighter language="javascript" controls="copy" line-numbers>
  <pre><code>const answer = 42;</code></pre>
</micro-lighter>
```

[Try the custom element demo](https://davatron5000.github.io/microlighter/custom-element.html)[Read more in the docs](https://github.com/davatron5000/microlighter#readme)

## Code samples

Supports the TextMate grammar format used by VS Code.

## Markup

```plain text
<!-- This div centers the div that centers the button. -->
<div class="centerer">
  <div class="centered">
    <button id="victory" type="button">Center me</button>
  </div>
</div>
<style>
  .centerer {
    display: grid;
    min-height: 100vh;
    place-items: center;
  }
  .centered {
    display: flex;
    align-items: center;
    justify-content: center;
  }
</style>
<script>
  document.querySelector("#victory").addEventListener("click", event => {
    event.target.textContent = "Still centered.";
  });
</script>
```

## Markdown

```plain text
# Yet Another Todo App
![build: passing](https://img.shields.io/badge/build-passing-brightgreen)
![coverage: probably](https://img.shields.io/badge/coverage-probably-yellow)
![works: here](https://img.shields.io/badge/works-on_my_machine-blue)
A **blazingly fast** todo app for developers who found the other
47,000 todo apps insufficiently opinionated.
## Install
    npm install yet-another-todo
## Usage
Run `npm start`, add "write documentation" to the list, then ignore it.
## Documentation
Coming soon™
> Production-ready, once we decide what production means.
```

## Python

```plain text
def readable_version(groups: list[list[str]]) -> list[str]:
    words = []
    for group in groups:
        for sentence in group:
            for word in sentence.split():
                if len(word) > 3:
                    words.append(word.lower())
    return words
def pythonic_version(groups: list[list[str]]) -> list[str]:
    return [
        word.lower()
        for group in groups
        for sentence in group
        for word in sentence.split()
        if len(word) > 3
    ]  # Pythonic.
assert readable_version([["Simple is better than nested"]]) == \
    pythonic_version([["Simple is better than nested"]])
```

## SQL

```plain text
-- Find meetings that could have been emails
SELECT
  m.title,
  COUNT(a.user_id) AS witnesses,
  SUM(m.duration_minutes) / 60.0 AS hours_wasted
FROM meetings AS m
JOIN attendees AS a ON a.meeting_id = m.id
WHERE m.could_have_been_email = TRUE
  AND m.outcome = 'another_meeting'
GROUP BY m.id, m.title
HAVING COUNT(a.user_id) > 2
ORDER BY hours_wasted DESC
LIMIT 10;
```

## C++

```plain text
#include <iostream>
#include <memory>
#include <string>
class Meeting {
public:
    explicit Meeting(std::string agenda)
        : agenda_(std::move(agenda)) {}
    ~Meeting() {
        std::cout << "This meeting has finally ended.\n";
    }
private:
    std::string agenda_;
};
std::unique_ptr<Meeting> make_meeting(std::string agenda) {
    return std::make_unique<Meeting>(std::move(agenda));
}
int main() {
    auto meeting = make_meeting("Discuss the next meeting");
    auto sequel = std::move(meeting); // Ownership is now perfectly clear.
    return 0;
}
```

## TSX

```plain text
import { useCallback, useEffect, useMemo, useState } from "react";
export function YetAnotherTodoApp(): JSX.Element {
  const [done, setDone] = useState(false);
  const completedCount = useMemo(() => Number(done), [done]);
  const toggle = useCallback(() => setDone(value => !value), []);
  useEffect(() => {
    console.log(`${completedCount} productivity units achieved`);
  }, [completedCount]);
  return (
    <label>
      <input type="checkbox" checked={done} onChange={toggle} />
      Write a simpler todo app
      {/* TODO: replace this with a state management library. */}
    </label>
  );
}
```

## Supported languages

Every grammar below ships as its own module and loads only when a page actually uses it. That includes formats like Git diff and config languages like TOML that aren't strictly "programming languages" but still get real TextMate scoping.

### Web

- HTML
- CSS
- SCSS
- JavaScript
- TypeScript
- TSX
- Svelte
- Vue
- HEEx

### Systems

- C
- C++
- Rust
- Go
- Assembly

### Application

- Java
- C#
- Kotlin
- Swift
- Objective-C
- Dart
- PHP
- Elixir

### Scripting

- Python
- Ruby
- Bash
- Perl
- PowerShell
- Lua
- R

### Data, config & docs

- JSON
- YAML
- TOML
- SQL
- Markdown
- GraphQL
- Dockerfile
- Git diff

## Playground

Preview a language

```plain text
/* Theme a single token */
::highlight(keyword) {
  color: var(--syntax-keyword);
}
```

```plain text
section .data
    msg db "tokens", 0
section .text
    global _start
_start:
    mov eax, 1
    int 0x80
```

```plain text
#!/usr/bin/env bash
set -euo pipefail
# Rebuild dist/ and run the test suite
npm run build && npm test
```

```plain text
#include <stdio.h>
int main(void) {
    printf("tokens: %d\n", 42);
    return 0;
}
```

```plain text
public class Token {
    public string Scope { get; set; } = "keyword";
    public int Start { get; set; } = 0;
}
```

```plain text
#include <string>
struct Token {
    std::string scope;
    int start = 0;
};
```

```plain text
/* Theme a single token */
::highlight(keyword) {
  color: var(--syntax-keyword);
}
```

```plain text
class Token {
  final String scope;
  final int start;
  const Token(this.scope, {this.start = 0});
}
```

```plain text
FROM node:20-slim
WORKDIR /app
COPY . .
RUN npm ci && npm run build
```

```plain text
defmodule Tokenizer do
  @moduledoc "Turns source into tokens."
  def tokenize(source), do: String.split(source, " ")
end
```

```plain text
diff --git a/src/index.js b/src/index.js
--- a/src/index.js
+++ b/src/index.js
@@ -1,2 +1,2 @@
-const answer = 0;
+const answer = 42;
```

```plain text
package main
import "fmt"
func main() {
	fmt.Println("tokens:", 42)
}
```

```plain text
# Fetch a token by id
query GetToken($id: ID!) {
  token(id: $id) {
    scope
  }
}
```

```plain text
<%!-- Render the current count --%>
<.button phx-click="increment" :if={@enabled}>
  <%= gettext "Count" %>: {@count}
</.button>
```

```plain text
<!-- A comment stays a comment -->
<button type="button" aria-pressed="false">42</button>
```

```plain text
public class Token {
    String scope = "keyword";
    int start = 0;
}
```

```plain text
// Highlight everything on the page
import { highlightAll } from "microlighter";
await highlightAll();
```

```plain text
{
  "name": "microlighter",
  "version": "1.0.0"
}
```

```plain text
data class Token(val scope: String, val start: Int = 0)
```

```plain text
-- Tokenize a tiny slice of source
local function tokenize(source)
    return source
end
```

```plain text
# MicroLighter
A **tiny** syntax highlighter for the [CSS Custom Highlight API](https://developer.mozilla.org/).
```

```plain text
@interface Token : NSObject
@property (nonatomic, strong) NSString *scope;
@end
```

```plain text
#!/usr/bin/env perl
use strict;
use warnings;
my @tokens = split / /, "const x = 42";
```

```plain text
<?php
class Token {
    public function __construct(private string $scope) {}
}
```

```plain text
# Rebuild dist/ and run the tests
function Invoke-Build {
    npm run build
}
```
