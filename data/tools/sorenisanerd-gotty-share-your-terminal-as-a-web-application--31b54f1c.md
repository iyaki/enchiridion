---
title: "sorenisanerd/gotty: Share your terminal as a web application"
notion_id: 31b54f1c-7d23-817b-bbef-f290d1741206
notion_url: https://app.notion.com/p/sorenisanerd-gotty-Share-your-terminal-as-a-web-application-31b54f1c7d23817bbbeff290d1741206
last_edited: 2026-09-21T16:57:00.000Z
source_url: https://github.com/sorenisanerd/gotty
tags: ["Tool", "GitHub", "English", "DevOps", "Web Development", "Backend", "Command Line", "Terminal", "Web Application"]
---
GoTTY is a simple command line tool that turns your CLI tools into web applications.

Original work by Iwasaki Yudai. There would be no GoTTY without him. ❤️

## Installation

## From release page

You can download the latest stable binary file from the Releases page. Note that the release marked Pre-release is built for testing purpose, which can include unstable or breaking changes. Download a release marked Latest release for a stable build.

(Files named with darwin_amd64 are for Mac OS X users)

## Homebrew Installation

You can install GoTTY with Homebrew as well.

```plain text
$ brew install sorenisanerd/gotty/gotty
```

## go get Installation (Development)

If you have a Go language environment, you can install GoTTY with the go get command. However, this command builds a binary file from the latest master branch, which can include unstable or breaking changes. GoTTY requires go1.9 or later.

```plain text
$ go get github.com/sorenisanerd/gotty
```

## Usage

```plain text
Usage: gotty [options] <command> [<arguments...>]
```

Run gotty with your preferred command as its arguments (e.g. gotty top).

By default, GoTTY starts a web server at port 8080. Open the URL on your web browser and you can see the running command as if it were running on your terminal.

## Options

```plain text
--address value, -a value IP address to listen (default: "0.0.0.0") [$GOTTY_ADDRESS] --port value, -p value Port number to liten (default: "8080") [$GOTTY_PORT] --path value, -m value Base path (default: "/") [$GOTTY_PATH] --permit-write, -w Permit clients to write to the TTY (BE CAREFUL) (default: false) [$GOTTY_PERMIT_WRITE] --credential value, -c value Credential for Basic Authentication (ex: user:pass, default disabled) [$GOTTY_CREDENTIAL] --random-url, -r Add a random string to the URL (default: false) [$GOTTY_RANDOM_URL] --random-url-length value Random URL length (default: 8) [$GOTTY_RANDOM_URL_LENGTH] --tls, -t Enable TLS/SSL (default: false) [$GOTTY_TLS] --tls-crt value TLS/SSL certificate file path (default: "~/.gotty.crt") [$GOTTY_TLS_CRT] --tls-key value TLS/SSL key file path (default: "~/.gotty.key") [$GOTTY_TLS_KEY] --tls-ca-crt value TLS/SSL CA certificate file for client certifications (default: "~/.gotty.ca.crt") [$GOTTY_TLS_CA_CRT] --index value Custom index.html file [$GOTTY_INDEX] --title-format value Title format of browser window (default: "{{ .command }}@{{ .hostname }}") [$GOTTY_TITLE_FORMAT] --reconnect Enable reconnection (default: false) [$GOTTY_RECONNECT] --reconnect-time value Time to reconnect (default: 10) [$GOTTY_RECONNECT_TIME] --max-connection value Maximum connection to gotty (default: 0) [$GOTTY_MAX_CONNECTION] --once Accept only one client and exit on disconnection (default: false) [$GOTTY_ONCE] --timeout value Timeout seconds for waiting a client(0 to disable) (default: 0) [$GOTTY_TIMEOUT] --permit-arguments Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB) (default: false) [$GOTTY_PERMIT_ARGUMENTS] --pass-headers Pass HTTP request headers as environment variables (e.g. Cookie becomes HTTP_COOKIE) (default: false) [$GOTTY_PASS_HEADERS] --width value Static width of the screen, 0(default) means dynamically resize (default: 0) [$GOTTY_WIDTH] --height value Static height of the screen, 0(default) means dynamically resize (default: 0) [$GOTTY_HEIGHT] --ws-origin value A regular expression that matches origin URLs to be accepted by WebSocket. No cross origin requests are acceptable by default [$GOTTY_WS_ORIGIN] --ws-query-args value Querystring arguments to append to the websocket instantiation [$GOTTY_WS_QUERY_ARGS] --enable-webgl Enable WebGL renderer (default: true) [$GOTTY_ENABLE_WEBGL] --quiet Don't log (default: false) [$GOTTY_QUIET] --close-signal value Signal sent to the command process when gotty close it (default: SIGHUP) (default: 1) [$GOTTY_CLOSE_SIGNAL] --close-timeout value Time in seconds to force kill process after client is disconnected (default: -1) (default: -1) [$GOTTY_CLOSE_TIMEOUT] --config value Config file path (default: "~/.gotty") [$GOTTY_CONFIG] --help, -h show help (default: false) --version, -v print the version (default: false)
```

### Config File

You can customize default options and your terminal by providing a config file to the gotty command. GoTTY loads a profile file at ~/.gotty by default when it exists.

```plain text
// Listen at port 9000 by default port = "9000" // Enable TSL/SSL by default enable_tls = true
```

See the .gotty file in this repository for the list of configuration options.

### Display Customization

GoTTY lets you customize the look of your terminal session — themes, fonts, cursor style, and more — through two channels: a config file for server-side defaults, and an on-screen runtime picker for live tweaking that persists to your browser's localStorage.

### Themes

GoTTY ships with 6 built-in color themes:

Theme Description default Catppuccin Mocha — warm, high-contrast nord Arctic bluish-cold palette dracula Dark purple-based high-contrast solarized-dark Classic solarized dark monokai High-contrast vivid theme light Clean light theme

Set a theme in your config:

```plain text
preferences { theme = "dracula" }
```

You can also override individual colors on top of any theme:

```plain text
preferences { theme = "nord" foreground_color = "#ffffff" background_color = "#1a1a2e" cursor_color = "#e94560" }
```

And override entries in the 16-color ANSI palette (ordered: black, red, green, yellow, blue, magenta, cyan, white, brightBlack, brightRed, brightGreen, brightYellow, brightBlue, brightMagenta, brightCyan, brightWhite):

```plain text
preferences { color_palette_overrides = ["", "", "#00ff00"] }
```

Leave entries empty ("") to keep the theme's default.

### Font Size

Set the terminal font size in pixels (8–48):

```plain text
preferences { font_size = 16 }
```

### Font Family

Choose a monospace font family:

```plain text
preferences { font_family = "'JetBrains Mono', monospace" }
```

The runtime picker makes these fonts available: DejaVu Sans Mono, JetBrains Mono, Fira Code, Source Code Pro, Monaco, Menlo, Cascadia Code, and the system default monospace. Each option in the picker renders its name in its own typeface as a live preview.

### Cursor & Scrollback

```plain text
preferences { cursor_style = "bar" // "block" (default), "underline", or "bar" cursor_blink = true scrollback_lines = 5000 enable_webgl = true // enabled by default }
```

### Alt as Meta Key

For terminal applications that use the Meta modifier (e.g. Emacs), you can have the Alt/Option key send an Escape prefix (M-x → \x1b + x):

```plain text
preferences { alt_is_meta = true }
```

This uses xterm.js's attachCustomKeyEventHandler to intercept Alt+key combinations before they reach the browser's default handling, and sends them as a Meta (Escape) sequence to the terminal. Works on macOS, Linux, and Windows.

Letter keys, digits, space, and common punctuation are supported. Browser-handled shortcuts like Alt+Tab and Alt+F4 are left alone.

### Runtime Picker (🎨 button)

When you open a GoTTY session, a small 🎨 button appears in the bottom-right corner. Click it to open a floating settings panel where you can switch themes, adjust font size, and change font family on the fly. All your choices are saved to localStorage and restored automatically on your next visit. The panel stays open while you tweak multiple settings and closes when you click outside it.

> Note: Runtime picker settings override config defaults for your current browser. Config preferences are re-applied on fresh connections (new browser, cleared storage, incognito mode).

### Security Options

By default, GoTTY doesn't allow clients to send any keystrokes or commands except terminal window resizing. When you want to permit clients to write input to the TTY, add the -w option. However, accepting input from remote clients is dangerous for most commands. When you need interaction with the TTY for some reasons, consider starting GoTTY with tmux or GNU Screen and run your command on it (see "Sharing with Multiple Clients" section for detail).

To restrict client access, you can use the -c option to enable the basic authentication. With this option, clients need to input the specified username and password to connect to the GoTTY server. Note that the credentials will be transmitted between the server and clients in plain text. For more strict authentication, consider the SSL/TLS client certificate authentication described below.

The -r option is a little bit more casual way to restrict access. With this option, GoTTY generates a random URL so that only people who know the URL can get access to the server.

All traffic between the server and clients are NOT encrypted by default. When you send secret information through GoTTY, we strongly recommend you use the -t option which enables TLS/SSL on the session. By default, GoTTY loads the crt and key files placed at ~/.gotty.crt and ~/.gotty.key. You can overwrite these file paths with the --tls-crt and --tls-key options. When you need to generate a self-signed certification file, you can use the openssl command.

```plain text
openssl req -x509 -nodes -days 9999 -newkey rsa:2048 -keyout ~/.gotty.key -out ~/.gotty.crt
```

(NOTE: For Safari uses, see how to enable self-signed certificates for WebSockets when use self-signed certificates)

For additional security, you can use the SSL/TLS client certificate authentication by providing a CA certificate file to the --tls-ca-crt option (this option requires the -t or --tls to be set). This option requires all clients to send valid client certificates that are signed by the specified certification authority.

## Sharing with Multiple Clients

GoTTY starts a new process with the given command when a new client connects to the server. This means users cannot share a single terminal with others by default. However, you can use terminal multiplexers for sharing a single process with multiple clients.

### Screen

After installing GNU screen, start a new session with screen -S name-for-session and connect to it with gotty in another terminal window/tab through screen -x name-for-session. All commands and activities being done in the first terminal tab/window will now be broadcasted by gotty.

### Tmux

For example, you can start a new tmux session named gotty with top command by the command below.

```plain text
$ gotty tmux new -A -s gotty top
```

This command doesn't allow clients to send keystrokes, however, you can attach the session from your local terminal and run operations like switching the mode of the top command. To connect to the tmux session from your terminal, you can use following command.

```plain text
$ tmux new -A -s gotty
```

By using terminal multiplexers, you can have the control of your terminal and allow clients to just see your screen.

### Quick Sharing on tmux

To share your current session with others by a shortcut key, you can add a line like below to your .tmux.conf.

```plain text
# Start GoTTY in a new window with C-t bind-key C-t new-window "gotty tmux attach -t `tmux display -p '#S'`"
```

## Playing with Docker

When you want to create a jailed environment for each client, you can use Docker containers like following:

```plain text
$ gotty -w docker run -it --rm busybox
```

## Development

You can build a binary by simply running make. go1.16 is required.

To build the frontend part (JS files and other static files), you need npm.

## Architecture

GoTTY uses xterm.js to run a JavaScript based terminal on web browsers. GoTTY itself provides a websocket server that simply relays output from the TTY to clients and receives input from clients and forwards it to the TTY. This xterm + websocket idea is inspired by Wetty.

## Alternatives

### Command line client

- gotty-client: If you want to connect to GoTTY server from your terminal

### Terminal/SSH on Web Browsers

- Secure Shell (Chrome App): If you are a chrome user and need a "real" SSH client on your web browser, perhaps the Secure Shell app is what you want
- Wetty: Node based web terminal (SSH/login)
- ttyd: C port of GoTTY with CJK and IME support

### Terminal Sharing

- tmate: Forked-Tmux based Terminal-Terminal sharing
- termshare: Terminal-Terminal sharing through a HTTP server
- tmux: Tmux itself also supports TTY sharing through SSH)

## License

The MIT License

## Contributors

Thanks goes to these wonderful people (emoji key):

Iwasaki Yudai💻 Soren L. Hansen🐛 💻 Andrea Lusuardi💻 Manfred Touron💻 Stephan💻 Quentin Perez💻 jzl💻 Fazal Majid💻 Immortalin💻 freakhill💻 0xflotus💻 Andy Skelton💻 Artem Medvedev💻 Blake Jennings💻 Christian Jensen💻 Christopher Wilkinson💻 Cyrus💻 David Horsley💻 Jason Cooke💻 Denis Korenevskiy💻 Massimiliano Stucchi💻 Mikhail f. Shiryaev💻 Robert Bittle💻 sebastian haas💻 shoji💻 Shuanglei Tao💻 The Gitter Badger💻 Jacob Zhou💻 zyfdegh💻 fredster33💻 mattn💻 Shinichi Goto💻 ygit💻 Stéphane🐛 Pavol Rusnak🐛 💻 Devan Lai💻 Jeeva Kandasamy💻 Steve Biedermann💻 xgdgsc🐛 💻 flechaig🐛 💻 Fan-SJ🐛 Dustin Martin🐛 Ahmet Alp Balkan🐛 CoconutMacaroon🐛 Danny Ben Shitrit🐛 George-NG🐛 Will Owens🐛 Jaime Pillora🐛 kaisawind🐛 linyinli🐛 LucaMarconato🐛 Kain🐛 Andi Andreas🐛 qigj🐛 shuaiyy🐛 v20z🐛 Yanfeng Qiu🐛 0xcanary🐛 alirezafarnoosh🐛 alphajoza🐛 Cross Nastasi🐛 Guilhem Bonnefille🐛 💻 huiwq1990🐛 imcnanie🐛 Jose Diaz-Gonzalez🐛 💻 Loccy🐛 Mrinal Wahal🐛 stevelaclasse🐛 180909💻 Benoit Tigeot💻 Callum Gare💻 Mike Bentzen💻 zigger💻 老J💻 Łukasz Lach💻 David Nutting💻 Randall McPherson💻 sagar-salvi-unskript💻 uddmorningsun💻 Tim Becker💻

This project follows the all-contributors specification. Contributions of any kind welcome!
