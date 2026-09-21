---
title: "Use Long Options in Scripts"
notion_id: 1c254f1c-7d23-8165-9965-d9e115970704
notion_url: https://app.notion.com/p/Use-Long-Options-in-Scripts-1c254f1c7d2381659965d9e115970704
last_edited: 2025-04-20T18:56:00.000Z
source_url: https://matklad.github.io/2025/03/21/use-long-options-in-scripts.html
tags: ["Article", "matklad (Alex Kladov)", "English", "SysAdmin", "Shell/Bash", "Programming"]
---
Many command line utilities support short form options (`-f`) and long form options (`--force`). Short form is for interactive usage. In scripts, use the long form.

That is, in your terminal, type `$ git switch -c my-new-branch`

In your release infrastructure script, write



```plain text
try shell.exec("git fetch origin --quiet", .{});
try shell.exec(
    "git switch --create release-{today} origin/main",
    .{ .today = stdx.DateUTC.now() },
);
```

Long form options are much more self-explanatory for the reader.
