---
title: "Vim and Composability"
notion_id: d5770ca9-fbc2-4836-a47e-c40ec937950a
notion_url: https://app.notion.com/p/Vim-and-Composability-d5770ca9fbc24836a47ec40ec937950a
last_edited: 2023-04-25T13:37:00.000Z
source_url: https://ferd.ca/vim-and-composability.html
tags: ["English", "Programming", "SysAdmin", "Linux", "IDE/Extendable Text Editor", "Article", "My bad opinions (Fred Hebert)"]
---
[https://ferd.ca/vim-and-composability.html](https://ferd.ca/vim-and-composability.html)

I've been a Vim user for a few years now. Before then I used emacs and emacs-like editors (without ever going too deep into them), but recurrent wrist pain forced me to change my usage patterns. Vim, as it turned out, was great for my wrist, and where I would not be able to work after 4-6 hours, I could now work for way longer than 8 every day, which I still hope to avoid.

The first thing about vim that appealed to me was the use of modes. When you think of it, using ctrl+x as a prefix for file handling shortcuts in emacs, or just holding ctrl or meta is modal editing. The difference is that Vim has different ones, and the keys leading to them only need to be pressed once.

That's where, I think, the biggest part of the learning curve is spent. Going from that series of implicit modes based on a given shortcut or prefix (ctrl, meta, etc.) to an explicit set that must be switched in and out of manually.

Where the fun begins is when you figure out Vim is its own language for the common shortcuts.

First of all, it has verbs, some of which are:

```plain text
d: delete D: delete line c: change (overwrite) C: change line y: yank (copy) Y: yank line >: indent <: deindent .: repeat previous action
```

It has actions to move around, some being:

```plain text
h,j,k,l: left, down, up, right (arrow keys also work for this) w: next word b: previous word 0: start of line $: end of line (: beginning of sentence ): end of sentence
```

Commands that select words or directions (hjkl and w for example) can also take a length prefix, so that 10w means '10 next words'.

There are also ways to search for the next instance of something, including:

```plain text
/regex/: search next instance n: next search result N: previous search result t: 'till (select) T: 'till backwards f: find (select) F: find backwards
```

The last 6 ones have to be followed by a character: tp will put the cursor before the next p in the text, and fp will put the cursor on the next p

It has a few modifiers that can be chained with verbs, notably:

```plain text
i: inside, excluding the character following it a: around, including the character following it
```

You build your commands by chaining them together:

```plain text
dl : delete 1 character to the right y$ : copy from the current cursor to the end of line ct; : change the content of the line from the current cursor position until the ';' cf; : change the content of the line from the current cursor position including the ';' da) : delete everything inside the parentheses, including the parens ci] : change the text within the square brackets while leaving the square brackets in place >3j : indent the next 3 lines >/hello : indent the text from the current position until the next line that contains 'hello' <a}: remove one level of indentation for the current block within { }, including the brackets. cit: change content within XML tags d5l : delete 5 characters to the right x5.: delete a character, then do it 5 times more (a length may prefix a verb!)
```

The commands become complex, but they're built from simple components expressing movements and specific actions.

You also have specific contexts, such as the visual mode (v), which will let you select text, the same as if you were dragging a mouse around. For example v3jtf will mean 'select (v) 3 lines down (3j) and then 'till the letter f (tf)'. That selection can then be passed to any of the verbs (d, c, >, <, etc.) to operate on it. In this case, doing v3jtfd would delete the complex selection.

Vim commands look like arcane incantations, but they're just composing a smallish set of functions together. Once that is learned, remembering specific commands and committing them to muscle memory is the hardest part. After that, each new command discovered has large benefits and usability because if it works in one context, it likely works in all other ones.
