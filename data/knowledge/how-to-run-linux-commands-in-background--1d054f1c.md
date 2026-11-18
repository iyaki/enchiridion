---
title: "How to Run Linux Commands in Background"
notion_id: 1d054f1c-7d23-81fd-b16a-dbda7a56c83a
notion_url: https://app.notion.com/p/How-to-Run-Linux-Commands-in-Background-1d054f1c7d2381fdb16adbda7a56c83a
last_edited: 2025-07-26T22:24:00.000Z
source_url: https://linuxize.com/post/how-to-run-linux-commands-in-background/
tags: ["Linuxize", "English", "Shell/Bash", "Tutorial"]
---
![image](https://linuxize.com/post/how-to-run-linux-commands-in-background/featured_hu5cd853877ba089617688d74aa6025897_29258_768x0_resize_q75_lanczos.jpg)

Linux Run Command in Background

Typically when you run a command in the terminal, you have to wait until the command finishes before you can enter another one. This is called running the command in the foreground or foreground process. When a process runs in the foreground, it occupies your shell, and you can interact with it using the input devices.

What if the command takes a long time to finish, and you want to run other commands in the meantime? You have several options at your disposal. The most obvious and straightforward option is to start a new shell session and run the command in it. Another option is to run the command in the background.

A background process is a process/command that is started from a terminal and runs in the background, without interaction from the user.

In this article, we will talk about the background processes is Linux. We will show you how to start a command in the background and how to keep the process running after the shell session is closed.

## Run a Linux Command in the Background

To run a command in the background, add the ampersand symbol (`&`) at the end of the command:

```plain text
command &

```

The shell job ID (surrounded with brackets) and process ID will be printed on the terminal:

```plain text
[1] 25177

```

You can have multiple processes running in the background at the same time.

The background process will continue to write messages to the terminal from which you invoked the command. To suppress the `stdout` and `stderr` messages use the following syntax:

```plain text
command > /dev/null 2>&1 &

```

`>/dev/null 2>&1` means redirect `stdout` to `/dev/null` and [`stderr`](https://linuxize.com/post/bash-redirect-stderr-stdout/)[ to ](https://linuxize.com/post/bash-redirect-stderr-stdout/)[`stdout`](https://linuxize.com/post/bash-redirect-stderr-stdout/) .

Use the `jobs` utility to display the status of all stopped and background jobs in the current shell session:

```plain text
jobs -l
```

The output includes the job number, process ID, job state, and the command that started the job:

```plain text
[1]+ 25177 Running                 ping google.com &

```

To bring a background process to the foreground, use the `fg` command:

```plain text
fg
```

If you have multiple background jobs, include `%` and the job ID after the command:

```plain text
fg %1
```

To terminate the background process, use the [`kill`](https://linuxize.com/post/kill-command-in-linux/) command followed by the process ID:

```plain text
kill -9 25177
```

## Move a Foreground Process to Background

To move a running foreground process in the background:

1. Stop the process by typing `Ctrl+Z`.
2. Move the stopped process to the background by typing `bg`.

## Keep Background Processes Running After a Shell Exits

If your connection drops or you log out of the shell session, the background processes are terminated. There are several ways to keep the process running after the interactive shell session ends.

One way is to remove the job from the shell’s job control using the `disown` shell builtin:

```plain text
disown
```

If you have more than one background jobs, include `%` and the job ID after the command:

```plain text
disown %1
```

Confirm that the job is removed from the table of active jobs using the `jobs -l` command. To list all running processes, including the disowned use the [`ps aux`](https://linuxize.com/post/ps-command-in-linux/) command.

Another way to keep a process running after the shell exit is to use `nohup`.

The [`nohup`](https://linuxize.com/post/linux-nohup-command/) command executes another program specified as its argument and ignores all `SIGHUP` (hangup) signals. `SIGHUP` is a signal that is sent to a process when its controlling terminal is closed.

To run a command in the background using the `nohup` command, type:

```plain text
nohup command &
```

The command output is redirected to the `nohup.out` file.

```plain text
nohup: ignoring input and appending output to 'nohup.out'

```

If you log out or close the terminal, the process is not terminated.

## Alternatives

There are a number of programs that allow you to have multiple interactive sessions at the same time.

### Screen

[Screen](https://linuxize.com/post/how-to-use-linux-screen/) or GNU Screen is a terminal multiplexer program that allows you to start a screen session and open any number of windows (virtual terminals) inside that session. Processes running in Screen will continue to run when their window is not visible even if you get disconnected.

### Tmux

[Tmux](https://linuxize.com/post/getting-started-with-tmux/) is a modern alternative to GNU screen. With Tmux, you can also create a session and open multiple windows inside that session. Tmux sessions are persistent, which means that programs running in Tmux continue to run even if you close the terminal.

## Conclusion

To run a command in the background, include `&` at the end of the command.

When you run a command in the background, you don’t have to wait until it finishes before you can execute another one.

If you have any questions or feedback, feel free to leave a comment.
