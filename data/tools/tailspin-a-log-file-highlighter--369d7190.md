---
title: "tailspin - A log file highlighter"
notion_id: 369d7190-7174-46b4-89a2-68d821e2a59a
notion_url: https://app.notion.com/p/tailspin-A-log-file-highlighter-369d7190717446b489a268d821e2a59a
last_edited: 2023-11-06T11:34:00.000Z
source_url: https://github.com/bensadeh/tailspin
tags: ["English", "Site Reliability Engineering", "DevOps", "SysAdmin", "Untried", "Infrastructure", "Network", "Tool"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

A log file highlighter

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### [Features](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#features)

- 🪵 View (or `tail`) any log file of any format
- 🍰 No setup or config required
- 🌈 Highlights numbers, dates, IP-addresses, UUIDs, URLs and more
- ⚙️ All highlight groups are customizable
- 🧬 Easy to integrate with other commands
- 🔍 Uses `less` under the hood for scrollback, search and filtering

### [Table of Contents](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#table-of-contents)

- [Overview](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#overview)
- [Installing](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#installing)
- [Highlight Groups](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#highlight-groups)
- [Watching folders](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#watching-folders)
- [Customizing Highlight Groups](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#customizing-highlight-groups)
- [Working with ](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#working-with-stdin-and-stdout)[`stdin`](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#working-with-stdin-and-stdout)[ and ](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#working-with-stdin-and-stdout)[`stdout`](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#working-with-stdin-and-stdout)
- [Using the pager ](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#using-the-pager-less)[`less`](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#using-the-pager-less)
- [Settings](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#settings)

## [Overview](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#overview)

`tailspin` works by reading through a log file line by line, running a series of regexes against each line. The regexes recognize patterns like dates, numbers, severity keywords and more.

`tailspin` does not make any assumptions on the format or position of the items it wants to highlight. For this reason, it requires no configuration or setup and will work predictably regardless of the format the log file is in.

## [Installing](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#installing)

### [Package Managers](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#package-managers)

The binary name for `tailspin` is `spin`.

```shell
# Homebrew
brew install tailspin

# Cargo
cargo install tailspin

# AUR
paru -S tailspin

# Nix
nix-shell -p tailspin

# NetBSD
pkgin install tailspin
```

### [From Source](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#from-source)

```shell
cargo install --path .
```

Binary will be placed in `~/.cargo/bin`, make sure you add the folder to your `PATH` environment variable.

## [Highlight Groups](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#highlight-groups)

### [Dates](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#dates)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[date]
style = { fg = "magenta" }
# To shorten the date, uncomment the line below
# shorten = { to = "␣", style = { fg = "magenta" } }

[time]
time = { fg = "blue" }
zone = { fg = "red" }
# To shorten the time, uncomment the line below
# shorten = { to = "␣", style = { fg = "blue" } }
```

### [Keywords](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#keywords)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[[keywords]]
words = ['null', 'true', 'false']
style = { fg = "red", italic = true }

[[keywords]]
words = ['GET']
style = { fg = "black", bg = "green" }
border = true

# You can add as many keywords as you'd like
```

### [URLs](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#urls)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[url]
http = { faint = true }
https = { bold = true }
host = { fg = "blue", faint = true }
path = { fg = "blue" }
query_params_key = { fg = "magenta" }
query_params_value = { fg = "cyan" }
symbols = { fg = "red" }
```

### [Numbers](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#numbers)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[number]
style = { fg = "cyan" }
```

### [IP Addresses](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#ip-addresses)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[ip]
segment = { fg = "blue", italic = true }
separator = { fg = "red" }
```

### [Quotes](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#quotes)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[quotes]
style = { fg = "yellow" }
token = '"'
```

### [Unix file paths](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#unix-file-paths)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[path]
segment = { fg = "green", italic = true }
separator = { fg = "yellow" }
```

### [HTTP methods](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#http-methods)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config See Keywords

### [UUIDs](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#uuids)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[uuid]
segment = { fg = "blue", italic = true }
separator = { fg = "red" }
```

### [Key-value pairs](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#key-value-pairs)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[key_value]
key = { faint = true }
separator = { fg = "white" }
```

### [Unix processes](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#unix-processes)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Details

Config

```plain text
[process]
name = { fg = "green" }
separator = { fg = "red" }
id = { fg = "yellow" }
```

## [Watching folders](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#watching-folders)

`tailspin` can listen for newline entries in a given folder. Watching folders is useful for monitoring log files that are rotated.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

When watching folders, `tailspin` will start in follow mode (abort with Ctrl + C) and will only print newline entries which arrive after the initial start.

## [Customizing Highlight Groups](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#customizing-highlight-groups)

### [Overview](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#overview-1)

Create `config.toml` in `~/.config/tailspin` to customize highlight groups.

Styles have the following shape:

```plain text
style = { fg = "color", bg = "color", italic = false, bold = false, underline = false }
```

### [Disabling Highlight Groups](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#disabling-highlight-groups)

To disable a highlight group, set the `disabled` field to true:

```plain text
[date]
disabled = true
```

### [Adding Keywords](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#adding-keywords)

To add custom keywords, either include them in the list of keywords or add new entries:

```plain text
[[keywords]]
words = ['MyCustomKeyword']
style = { fg = "green" }

[[keywords]]
words = ['null', 'true', 'false']
style = { fg = "red", italic = true }
```

## [Working with ](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#working-with-stdin-and-stdout)[`stdin`](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#working-with-stdin-and-stdout)[ and ](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#working-with-stdin-and-stdout)[`stdout`](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#working-with-stdin-and-stdout)

By default, `tailspin` will open a file in the pager `less`. However, if you pipe something into `tailspin`, it will print the highlighted output directly to `stdout`. This is similar to running `spin [file] --print`.

To let `tailspin` highlight the logs of different commands, you can pipe the output of those commands into `tailspin` like so:

```plain text
journalctl -f | spin
cat /var/log/syslog | spin
kubectl logs -f pod_name | spin
```

## [Using the pager ](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#using-the-pager-less)[`less`](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#using-the-pager-less)

### [Overview](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#overview-2)

`tailspin` uses `less` as its pager to view the highlighted log files. You can get more info on `less` via the **man** command (`man less`) or by hitting the h button to access the help screen.

### [Navigating](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#navigating)

Navigating within `less` uses a set of keybindings that may be familiar to users of `vim` or other `vi`-like editors. Here's a brief overview of the most useful navigation commands:

- j/k: Scroll one line up / down
- d/u: Scroll one half-page up / down
- g/G: Go to the top / bottom of the file

### [Follow mode](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#follow-mode)

When you run `tailspin` with the `-f` or `--follow` flag, it will scroll to the bottom and print new lines to the screen as they're added to the file.

To stop following the file, interrupt with Ctrl + C. This will stop the tailing, but keep the file open, allowing you to review the existing content.

To resume following the file from within `less`, press Shift + F.

### [Search](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#search)

Use / followed by your search query. For example, `/ERROR` finds the first occurrence of **ERROR**.

After the search, n finds the next instance, and N finds the previous instance.

### [Filtering](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#filtering)

`less` allows filtering lines by a keyword, using & followed by the pattern. For instance, `&ERROR` shows only lines with **ERROR**.

To only show lines containing either `ERROR` or `WARN`, use a regular expression: `&\(ERROR\|WARN\)`.

To clear the filter, use & with no pattern.

## [Settings](https://github.com/bensadeh/tailspin?utm_source=tldrnewsletter#settings)

```plain text
-f, --follow                 Follow the contents of the file
-t, --tail                   Start at the end of the file
-p, --print                  Print the output to stdout
-c, --config-path PATH       Path to a custom configuration file
-t, --follow-command 'CMD'   Follows the output of the provided command
```
