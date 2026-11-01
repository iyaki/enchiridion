---
title: "fx - Terminal JSON viewer"
notion_id: 252cb384-e331-494e-8f50-9211abbec11e
notion_url: https://app.notion.com/p/fx-Terminal-JSON-viewer-252cb384e331494e8f509211abbec11e
last_edited: 2023-04-22T01:09:00.000Z
source_url: https://fx.wtf/
tags: ["English", "Shell/Bash", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

- _Function eXecution_

## Features

- Mouse support
- Streaming support
- Preserves key order
- Preserves big numbers

## Install

```plain text
brew install fx
```

```plain text
snap install fx
```

```plain text
scoop install fx
```

```plain text
pacman -S fx
```

```plain text
pkg install fx
```

```plain text
go install github.com/antonmedv/fx@latest
```

Or download [pre-built binary](https://github.com/antonmedv/fx/releases) via:

```plain text
curl https://fx.wtf | sh
```

## Usage

Start the interactive viewer via:

```plain text
fx data.json
```

Or

```plain text
curl ... | fx
```

Type `?` to see full list of key shortcuts.

Pretty print:

```plain text
curl ... | fx .
```

### Reducers

Write reducers in your favorite language: [JavaScript](https://github.com/antonmedv/fx/blob/master/doc/js.md) (default), [Python](https://github.com/antonmedv/fx/blob/master/doc/python.md), or [Ruby](https://github.com/antonmedv/fx/blob/master/doc/ruby.md).

```plain text
fx data.json '.filter(x => x.startsWith("a"))'
```

```plain text
fx data.json '[x["age"] + i for i in range(10)]'
```

```plain text
fx data.json 'x.to_a.map {|x| x[1]}'
```

## Documentation

See full [documentation](https://github.com/antonmedv/fx/blob/master/doc/doc.md).

## Themes

Theme can be configured by setting environment variable `FX_THEME` from `1` to `9`:

```plain text
export FX_THEME=9
```

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Add your own themes in [theme.go](https://github.com/antonmedv/fx/blob/master/pkg/theme/theme.go) file.

## License

[MIT](https://github.com/antonmedv/fx/blob/master/LICENSE)
