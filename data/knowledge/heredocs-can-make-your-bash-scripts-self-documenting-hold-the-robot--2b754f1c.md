---
title: "Heredocs Can Make Your Bash Scripts Self-Documenting | Hold The Robot"
notion_id: 2b754f1c-7d23-81cb-ae00-fcd551083199
notion_url: https://app.notion.com/p/Heredocs-Can-Make-Your-Bash-Scripts-Self-Documenting-Hold-The-Robot-2b754f1c7d2381cbae00fcd551083199
last_edited: 2025-11-26T18:48:00.000Z
source_url: https://holdtherobot.com/blog/heredocs-can-make-your-bash-scripts-self-documenting/
tags: ["English", "Shell/Bash", "Automation", "Productivity", "Article", "Hold The Robot"]
---
I have long since come to appreciate the value of writing scripts to avoid someone else (or future me) from having to re-learn and re-solve problems, but something about it has always bugged me.

I am automating a process, but I'm also _documenting_ it, and those two things struggle to coexist.

One option is to write a bash script for the automation and a markdown file for the documentation, but they inevitably end up duplicating information and/or getting out of sync. The other is to just have a single markdown file with a bunch of inline bash that you manually copy into a terminal. But "running" it is clunky, tedious, and easy to mess up.

I tend to prefer the latter despite the annoyances, because "keeping information in sync" is such a big problem. But recently I've been playing with a third option. Rather than maintaining two files or putting bash in markdown; put markdown in bash.

It looks like this (I'm showing you an image since Docusaurus won't syntax highlight this):

![image](https://holdtherobot.com/assets/images/inline-markdown-972100d187dd079b0b3b9ec954770eeb.png)

This is just a bash script that can be executed like normal.

The markdown bit is a "heredoc", which is basically just a multiline string, similar to a triple-quoted string in python. The `<<'delimiter'` starts the string and `delimiter` ends it. Be careful to quote the first delimiter, otherwise you'll get parameter expansion (things like `$HOME` will expand to `/home/myusername`) or even execution in your doc strings (intuitive as always, thanks Bash). I chose `-md-` as a delimiter, but you can choose whatever you like, as long as it's not a string you're going to be using otherwise.

If you precede there heredoc with `cat` it will print to the terminal when you run the script, but you can also leave that out.

I use the vim plugin `preservim/vim-markdown` to get markdown syntax highlighting, concealment, links, and so on. By default, none of that is going to work inside a bash script, but you can fix that by adding the following to `.config/nvim/after/syntax/sh.vim` (create the file and path if needed):

```plain text
syntax region shMarkdown
    \ matchgroup=shMarkdownDelimiter
    \ start=/cat <<'-md-'\s*$/
    \ end=/^-md-\s*$/
    \ contains=@markdownHighlight
    \ containedin=shHereDoc,shHereString
    \ keepend
syntax include @markdownHighlight syntax/markdown.vim
" Link the delimiter to Comment so it's greyed out
highlight link shMarkdownDelimiter Comment

```

And there you go; markdown-ified bash scripting.

There's still plenty of times a markdown file makes more sense, since you're not always writing bash commands that are intended to be run top-to-bottom. I have a file that lists various ffmpeg commands, for example, and I'm only ever going to be copy-pasting things out of that file. But for a runbook style script I really quite like this and I think it's absolutely a better option than maintaining separate scripts and documentation. There's a reason why so many modern codebases use inline documentation, and I think bash scripts should do the same.
