---
title: "try - Simple CLI tool to try Composer (PHP) packages"
notion_id: f5dfc934-df32-4e7c-9d08-e9a18dbe6059
notion_url: https://app.notion.com/p/try-Simple-CLI-tool-to-try-Composer-PHP-packages-f5dfc934df324e7c9d08e9a18dbe6059
last_edited: 2022-12-21T14:37:00.000Z
source_url: https://github.com/marijnvanwezel/try#try---simple-cli-tool-to-try-composer-packages
tags: ["English", "PHP", "Untried", "Tool"]
---
## Folders and files

NameNameLast commit messageLast commit datebinbindocsdocssrcsrc.gitignore.gitignoreLICENSELICENSEREADME.mdREADME.mdcomposer.jsoncomposer.json

- README
- GPL-3.0 license

try makes it super easy to try new Composer packages through the command-line. It was inspired by timofurrer/try.

## Installation

Installation should be done through Composer:

```plain text
composer global require marijnvanwezel/try export PATH=$PATH:~/.config/composer/vendor/bin
```

This installs try as a system-wide binary.

## Usage examples

Try single Composer package:

```plain text
try nikic/php-parser try webmozart/assert
```

Try multiple Composer packages in the same session:

```plain text
try nikic/php-parser webmozart/assert
```

Try a specific version of a package:

```plain text
try webmozart/assert:1.10.0
```
