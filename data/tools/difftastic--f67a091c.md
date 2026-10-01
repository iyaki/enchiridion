---
title: "Difftastic"
notion_id: f67a091c-7c71-4165-ad83-c92370011d12
notion_url: https://app.notion.com/p/Difftastic-f67a091c7c714165ad83c92370011d12
last_edited: 2023-01-25T19:23:00.000Z
source_url: https://difftastic.wilfred.me.uk/
tags: ["Tool", "English", "Programming"]
---
Difftastic is a structural diff tool that understands syntax. It supports [over 30 programming languages](https://difftastic.wilfred.me.uk/languages_supported.html) and when it works, it's _fantastic_.

Difftastic is open source software (MIT license) and [available on GitHub](https://github.com/wilfred/difftastic).

This copy of the manual describes version 0.43.0. The [changelog](https://github.com/Wilfred/difftastic/blob/master/CHANGELOG.md) records which features and bug fixes are in each version.



Difftastic [detects the language](https://difftastic.wilfred.me.uk/usage.html#language-detection), parses the code, and then compares the syntax trees. Let's look at an example.

```ruby
// old.rs
let ts_lang = guess(path, guess_src).map(tsp::from_language);

```

```ruby
// new.rs
let ts_lang = language_override
    .or_else(|| guess(path, guess_src))
    .map(tsp::from_language);

```

```shell
$ difft old.rs new.rs

1 1 let ts_lang = language_override
. 2     .or_else(|| guess(path, guess_src))
. 3     .map(tsp::from_language);

```

Notice how difftastic recognises that `.map` is unchanged, even though it's now on a new line with whitespace.

A line-oriented diff does a much worse job here.

```shell
$ diff -u old.rs new.rs

@@ -1 +1,3 @@
-let ts_lang = guess(path, guess_src).map(tsp::from_language);
+let ts_lang = language_override
+    .or_else(|| guess(path, guess_src))
+    .map(tsp::from_language);

```

Some textual diff tools also highlight word changes (e.g. GitHub or git's `--word-diff`). They still don't understand the code though. Difftastic will always find matched delimiters: you can see the closing `)` from `or_else` has been highlighted.

If input files are not in a format that difftastic understands, it uses a conventional line-oriented text diff with word highlighting.

Difftastic will also use textual diffing when given extremely large inputs.
