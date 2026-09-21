---
title: "Navigating around in your shell"
notion_id: f12716dc-c8ba-4098-a192-1c014102857d
notion_url: https://app.notion.com/p/Navigating-around-in-your-shell-f12716dcc8ba4098a1921c014102857d
last_edited: 2023-11-28T17:46:00.000Z
source_url: https://blog.meain.io/2023/navigating-around-in-shell/
tags: ["Article", "meain/blog", "English", "Shell/Bash"]
---
I have been using terminals for a long time, initially because I thought they looked cool, and later because I genuinely found them to be easier/faster to get stuff done. And since I've been at it for a while, navigating through directories is something I think I've gotten good at. In this blog, I would like to give some tips on ways you can navigate around in your shell quickly.

_I rarely type __`cd`__ to change directories. I get the same feeling you get when you see people google for google on google when I see people typing __`cd ..`__ repeatedly._

## Using `cd` [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#using-cd)

We all know the basics, ie `cd <dir>` to navigate to `dir`. Here a few other things ways you can use `cd`:

- `cd` (no args) is equivalent to `cd ~`
- `cd -` takes you to the previous(not parent) directory
- `cd ..` takes you to the parent directory

## Make use of that `tab` key [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#make-use-of-that-tab-key)

This might sound simple, but you have no idea how many times I've had the painful experience of watching people type out the path by had and even end typing them incorrectly.

## Let's add some useful aliases [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#let%27s-add-some-useful-aliases)

Now that we know some basics, lets add some aliases.

Here are the ones I really like. I would suggest adding aliases for going to previous dir, and multiple levels of parent dirs. I initial got this idea from [ohmyzsh/ohmyzsh](https://github.com/ohmyzsh/ohmyzsh). You can find these in my [`dotfiles`](https://github.com/meain/dotfiles/blob/f7ac724386f9dfb24eda53191146795b788b0f06/zsh/.config/zsh/aliases#L4-L9).

```plain text
alias -- -='cd -'
alias ..='cd ..'
alias ...='cd ../..'
alias ....='cd ../../..'
alias .....='cd ../../../..'
alias ......='cd ../../../../..'
```

## Make use of globs [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#make-use-of-globs)

If you know only parts of a path, you can always use globs and let shell figure out the exact path. Use this along with tab completion to make it even more powerful.

For example you can do `cd src/**/color` to find a directory named `color` in the current working directory.

## Auto fix small errors [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#auto-fix-small-errors)

Since my blog is mostly read by humans(I'm sorry bots, you are in a minority), and since humans make mistakes, this one is for you. I would suggest enabling these two options. The first one ignores the case of the argument and the second one auto fixes any small spelling mistakes.

```plain text
# for bash
shopt nocaseglob # ignore case when matching
shopt cdspell # fix common spelling mistakes

# for zsh
setopt nocaseglob # ignore case
setopt correct # correct spelling mistakes
```

_You might be interested in _[_shopt_](https://ss64.com/bash/shopt.html)_ and _[_setopt_](https://zsh.sourceforge.io/Doc/Release/Options.html)_ manuals._

## For the lazy ones [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#for-the-lazy-ones)

Why type `cd` if you don't have to. If you enable to following option, you can just type the name of the folder and if there is no binary by that name in your `$PATH`, your shell will `cd` into it that directory.

```plain text
shopt -s autocd # for bash
setopt auto_cd # for zsh
```

## Remembering a directory [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#remembering-a-directory)

What if you wanna move around a lot, but wanIf you're frequently on the move but need to stay organized and revisit specific directories easily, utilizing pushd and popd can be incredibly beneficial.t to keep track of certain directories specifically to come back to. `pushd` and `popd` are you friends.

```plain text
pushd dir1 # in dir1
cd dir2 # in dir2
pushd dir3 # in dir3
cd dir4
cd dir5
popd # in dir3
popd # in dir1
```

## Bookmarks [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#bookmarks)

[`hash`](https://til.hashrocket.com/posts/xsavbhlrz4-shortcuts-with-hash-d-in-zsh) is a way for you to "bookmark" directories. I use it to "bookmark" directories that I visit often. By default, the bookmarks are only per session, but you can add a bit more code to make it persistent.

The best part is that you can use these as prefixes to construct the full paths. For example if I have bookmark `~/.local/share` as `~share`, I can later do things like `cd ~share/emacs`.

Here are the additions that I have:

### Use `CDPATH` env variable [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#use-cdpath-env-variable)

You can set this env variable with the paths that you would like to always cd into.

For instance, if I set it as `export CDPATH=$HOME/.config:$HOME/.local/share`, I can cd into any dirs in `$HOME/.config` and `$HOME/.local/share` form any location. Moreover, these paths will be incorporated into your shell completions.

## Make a directory and cd into it [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#make-a-directory-and-cd-into-it)

I more of than not find myself having to create a directory and cd into it. I use an function called `take` for this. This is another one of those gems I stole from `oh-my-zsh`.

## Navigating to project root [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#navigating-to-project-root)

Another common usecase is to go the root of the project. Let's also simplify it by using `git` to figure out the root of the project.

## Navigate to a project under this dir [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#navigate-to-a-project-under-this-dir)

Like most people, I too maintain a dir under which I keep all my project. This function lets me pick one of those projects (using `fzf` to filter). It looks for all dirs with a `.git` in them and find their parent dirs.

## Quickly create a tmp dir [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#quickly-create-a-tmp-dir)

Whenever I want to test out something, or look at(clone) a project just temporarily, I used to create a dir under `/tmp` and then do it there. This script automates that. The actual script that I use also does a `git init` so that I can track and diff files.

## One-off dirs [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#one-off-dirs)

One more. I get a lot of data being sent to be as zips, and when I want to look at them, just like before, I create a dir in `tmp`, unzip it into that dir and look into the logs. After a while I automated these steps. This can sound like a me problem, but just wanted to show you the kind of things you can do.

## fzf [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#fzf)

[`fzf`](https://github.com/junegunn/fzf) is a super [versatile](https://github.com/search?q=repo%3Ameain%2Fdotfiles%20fzf&type=code) tool. I use it a lot in scripts, but by default it adds a `ctrl+t` shortcut to select files and folders.

## Intelligent `cd` [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#intelligent-cd)

There's a multitude of tools available that learn the directories you frequently access and allow you to swiftly "jump" into them by typing partial names. I personally rely on rupa/z due to my familiarity with it, but there are several other options worth exploring, such as wting/autojump, ajeetdsouza/zoxide (the Rust one for extra credibility), and gsamokovarov/jump. Each offers its own set of features and advantages, catering to different preferences and workflows.

## Using a TUI file browser [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#using-a-tui-file-browser)

I grew up using [ranger/ranger](https://github.com/ranger/ranger) but switched over to [gokcehan/lf](https://github.com/gokcehan/lf) at some point and I use it as way to select the dir that I wanna jump into. The following code is taken straight from `lf`'s wiki. With this, you can type `lfcd`, then go to the dir you want by navigating a TUI browser and then exit to land on that dir.

## Fuzzy tui navigation [#](https://blog.meain.io/2023/navigating-around-in-shell/?utm_source=tldrnewsletter#fuzzy-tui-navigation)

Another option is to use something like [Canop/broot](https://github.com/Canop/broot) to select a dir. I don't use this much, but have found this useful in the past. You can find how to do that [here](https://github.com/Canop/broot#find-a-directory-then-cd-to-it).

I always wondered if `cd .` had a purpose. I found this out from someone on Mastodon a while ago (unfortunately, I don't remember who).

If you delete a folder from a different terminal session and subsequently recreate a new folder with the same name, using cd . allows you to navigate into the newly created folder without changing your current directory. This command essentially refreshes the current directory, enabling access to the newly created directory with the same name.
