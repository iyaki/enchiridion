---
title: "Bash prompt generator"
notion_id: 5026b645-cae4-43ce-86f1-1b58060dafad
notion_url: https://app.notion.com/p/Bash-prompt-generator-5026b645cae443ce86f11b58060dafad
last_edited: 2026-09-21T17:09:00.000Z
source_url: https://robotmoon.com/bash-prompt-generator/
tags: ["English", "Shell/Bash", "Tool", "Service"]
---
user@hostname ~/path/to/directory $

↑

export PS1="\[$(tput setaf 226)\]\u\[$(tput setaf 220)\]@\[$(tput setaf 214)\]\h \[$(tput setaf 33)\]\w \[$(tput sgr0)\]$ "

## Bash prompt PS1

To use the colors you chose, set the PS1 environment variable in your shell:

export PS1="\[$(tput setaf 226)\]\u\[$(tput setaf 220)\]@\[$(tput setaf 214)\]\h \[$(tput setaf 33)\]\w \[$(tput sgr0)\]$ "export PS1="\[\e[38;5;226m\]\u\[\e[38;5;220m\]@\[\e[38;5;214m\]\h \[\e[38;5;33m\]\w \[\033[0m\]$ "

It's up to you to decide between tput and ANSI escape sequences. To persist your customized prompt, export PS1 in ~/.bashrc or ~/.bash_profile

## Bash prompt examples

These are some example color schemes from choosing 4 colors above. Click on the bash prompt previews to view their tput and ANSI PS1 exports.

user@hostname ~/path/to/directory $

Emerald green

user@hostname ~/path/to/directory $

Lemon line

user@hostname ~/path/to/directory $

Fiery orange

user@hostname ~/path/to/directory $

Autumn leaves

user@hostname ~/path/to/directory $

Desert sand

user@hostname ~/path/to/directory $

Ocean blue

user@hostname ~/path/to/directory $

Blue green yellow

user@hostname ~/path/to/directory $

Twilight

user@hostname ~/path/to/directory $

Violet pink

user@hostname ~/path/to/directory $

Monochromatic

## Bash prompt variables

These are the variable substitutions used above.

```plain text
\u user \h hostname \w ~/path/to/directory
```

For reference, this is a list of all the valid bash prompt variables from the PROMPTING section in the bash man pages in `man bash`

```plain text
\a an ASCII bell character (07) \d the date in "Weekday Month Date" format (e.g., "Tue May 26") \D{format} the format is passed to strftime(3) and the result is inserted into the prompt string; an empty format results in a locale-specific time representation. The braces are required \e an ASCII escape character (033) \h the hostname up to the first `.' \H the hostname \j the number of jobs currently managed by the shell \l the basename of the shell's terminal device name \n newline \r carriage return \s the name of the shell, the basename of $0 (the portion following the final slash) \t the current time in 24-hour HH:MM:SS format \T the current time in 12-hour HH:MM:SS format \@ the current time in 12-hour am/pm format \A the current time in 24-hour HH:MM format \u the username of the current user \v the version of bash (e.g., 2.00) \V the release of bash, version + patch level (e.g., 2.00.0) \w the current working directory, with $HOME abbreviated with a tilde (uses the value of the PROMPT_DIRTRIM variable) \W the basename of the current working directory, with $HOME abbreviated with a tilde \! the history number of this command \# the command number of this command \$ if the effective UID is 0, a #, otherwise a $ \nnn the character corresponding to the octal number nnn \\ a backslash \[ begin a sequence of non-printing characters, which could be used to embed a terminal control sequence into the prompt \] end a sequence of non-printing characters
```
