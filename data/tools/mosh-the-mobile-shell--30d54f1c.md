---
title: "Mosh: the mobile shell"
notion_id: 30d54f1c-7d23-81c3-a954-f05ca166b1a8
notion_url: https://app.notion.com/p/Mosh-the-mobile-shell-30d54f1c7d2381c3a954f05ca166b1a8
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://mosh.org/
tags: ["Tool", "Official Website", "English", "DevOps", "Terminal", "Command Line", "Remote Access", "Unix", "Networking"]
---
Remote terminal application that allows **roaming**, supports **intermittent connectivity**, and provides intelligent **local echo** and line editing of user keystrokes.

Mosh is a replacement for interactive SSH terminals. It's more robust and responsive, especially over Wi-Fi, cellular, and long-distance links.

Mosh is free software, available for GNU/Linux, BSD, macOS, Solaris, Android, Chrome, and iOS.

![image](https://mosh.org/mosh.png)

## Change IP. Stay connected.

Mosh automatically roams as you move between Internet connections. Use Wi-Fi on the train, Ethernet in a hotel, and LTE on a beach: you'll stay logged in. Most network programs lose their connections after roaming, including SSH and Web apps like Gmail. Mosh is different.

With Mosh, you can put your laptop to sleep and wake it up later, keeping your connection intact. If your Internet connection drops, Mosh will warn you — but the connection resumes when network service comes back.

SSH waits for the server's reply before showing you your own typing. That can make for a lousy user interface. Mosh is different: it gives an instant response to typing, deleting, and line editing. It does this adaptively and works even in full-screen programs like emacs and vim. On a bad connection, outstanding predictions are underlined so you won't be misled.

You don't need to be the superuser to install or run Mosh. The client and server are executables run by an ordinary user and last only for the life of the connection.

Mosh doesn't listen on network ports or authenticate users. The mosh client logs in to the server via SSH, and users present the same credentials (e.g., password, public key) as before. Then Mosh runs the mosh-server remotely and connects to it over UDP.

Mosh is a command-line program, like ssh. You can use it inside xterm, gnome-terminal, urxvt, Terminal.app, iTerm, emacs, screen, or tmux. But mosh was designed from scratch and supports just one character set: UTF-8. It fixes Unicode bugs in other terminals and in SSH.

Unlike SSH, mosh's UDP-based protocol handles packet loss gracefully, and sets the frame rate based on network conditions. Mosh doesn't fill up network buffers, so Control-C always works to halt a runaway process.
