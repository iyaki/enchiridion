---
title: "Pesky little scripts"
notion_id: 5ea02d4c-a402-4aa6-b7dc-e0409fce4dcd
notion_url: https://app.notion.com/p/Pesky-little-scripts-5ea02d4ca4024aa6b7dce0409fce4dcd
last_edited: 2023-12-05T17:32:00.000Z
source_url: https://rednafi.com/misc/pesky_little_scripts
tags: ["English", "Shell/Bash", "Article", "Redowan's Reflections"]
---
I like writing custom scripts to automate stuff or fix repetitive headaches. Most of them are shell scripts, and a few of them are written in Python. Over the years, I’ve accumulated quite a few of them. I use Git and GNU stow[1](https://rednafi.com/misc/pesky_little_scripts/?utm_source=tldrnewsletter#fn:1) to manage them across different machines, and the workflow[2](https://rednafi.com/misc/pesky_little_scripts/?utm_source=tldrnewsletter#fn:2) is quite effective. However, as the list of scripts grows larger, invoking them becomes a pain because the tab completion results get cluttered with other system commands. Plus, often I even forget the initials of a script’s name and stare at my terminal while the blinking cursor facepalms at my stupidity.

I was watching this amazing talk[3](https://rednafi.com/misc/pesky_little_scripts/?utm_source=tldrnewsletter#fn:3) by Brandon Rhodes that proposes quite an elegant solution to this problem. It goes like this:

> All your scripts should start with a character as a prefix that doesn’t have any special meaning in the shell environment. Another requirement is that no other system command should start with your chosen character.

That way, when you type the prefix character and hit tab, only your custom scripts should appear and nothing else. This works with your aliases too!

The dilemma here is picking the right character that meets both of the requirements. Luckily, Brandon did the research for us. Turns out, the shell environment uses pretty much all the characters on the keyboard as special characters other than these 6:

```plain text
@ _ + - : ,

```

Among them, the first 5 requires pressing the Shift key, which is inconvenient. But the plain old comma `,` is right there. You can start your script or alias names with a comma `,` and it’ll be golden.

My tab completion looks like this:

```plain text
rednafi@air:~/canvas/rednafi.com
$ ,
,brclr                    ,clear-cache              ,docker-prune-containers  ,redis
,brpre                    ,docker-nuke              ,docker-prune-images      ,www

```

All my aliases start with `,` too so that they also appear in the list with the custom scripts. Fin!

1. [GNU stow](https://www.gnu.org/software/stow/) [↩︎](https://rednafi.com/misc/pesky_little_scripts/?utm_source=tldrnewsletter#fnref:1)
2. [Dotfile stewardship for the indolent](https://rednafi.com/misc/dotfile_stewardship_for_the_indolent) [↩︎](https://rednafi.com/misc/pesky_little_scripts/?utm_source=tldrnewsletter#fnref:2)
3. [Activation energy — Brandon Rhodes](https://www.youtube.com/watch?v=pybtvFFRYFs) [↩︎](https://rednafi.com/misc/pesky_little_scripts/?utm_source=tldrnewsletter#fnref:3)

## Recent posts

- [Escaping the template pattern hellscape in Python](https://rednafi.com/python/escape_template_pattern/)
- [Finding flow amid chaos](https://rednafi.com/misc/finding_flow_amid_chaos/)
- [The diminishing half-life of knowledge](https://rednafi.com/misc/diminishing_half_life_of_knowledge/)
- [Oh my poor business logic](https://rednafi.com/misc/oh_my_poor_business_logic/)
