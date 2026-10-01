---
title: "Dotfiles matter!"
notion_id: 60ec991d-4630-4f0f-a937-edbe53c3b0c8
notion_url: https://app.notion.com/p/Dotfiles-matter-60ec991d46304f0fa937edbe53c3b0c8
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://dotfiles-matter.click/
tags: ["English", "Linux", "Shell/Bash", "Article", "Website"]
---
```plain text
| Go         | https://github.com/adrg/xdg                  |
| Rust       | https://github.com/whitequark/rust-xdg       |
| Python     | https://github.com/srstevenson/xdg-base-dirs |
| Haskell    | https://github.com/willdonnelly/xdg-basedir  |
| JavaScript | https://github.com/folder/xdg                |
| Java       | https://github.com/omajid/xdg-java           |
| Ruby       | https://github.com/rubyworks/xdg             |
| C++        | https://github.com/azubieta/xdg-utils-cxx    |

```

(Note for Go users: Go’s `os.UserConfigDir` is just a hardcoded $HOME/.config (literally so unhelpful))

### TL;DR: Where should you store application data?

### Linux/Unix

As of 2023, most major Linux distributions and desktop environments adhere to the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html), which defines standard locations for different types of files: cache, configuration data, and state (… more data). The specification defines environment variables which if defined provide an absolute path for that kind of data. Following are the expected locations for application files in order of importance (each path begins from $HOME):

```plain text
| Cache            | $XDG_CACHE_HOME  | $HOME/.cache                      | Non-essential non-persistent data files, logs, thumbnails etc.   |
| Configuration    | $XDG_CONFIG_HOME | $HOME/.config                     | Persistent configuration data files                              |
| Application Data | $XDG_DATA_HOME   | $HOME/.local/share                | Data files like history, databases, logs, "stuff"                |
| Executable Files | n/a (sigh)       | $HOME/.local/bin                  | User-facing executable files (i.e. the output of `cc -o <file>`) |
| Runtime Data     | $XDG_RUNTIME_DIR | Should be provided by your system | This is for runtime data (sockets, pipes, whatever.)             |

```

There are more directories, but don’t overcomplicate things. Even just using `$XDG_CACHE_HOME` and `$XDG_CONFIG_HOME` is a lifesaver.

### MacOS

Apple provides the [File System Programming Guide](https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileSystemOverview/FileSystemOverview.html#//apple_ref/doc/uid/TP40010672-CH2-SW1), which defines expected storage locations, again for: cache, configuration and state (data). Following are the expected locations for application files (where ~ is $HOME, the current user’s home directory):

```plain text
| Cache                           | ~/Library/Caches              | Non-essential, non-persistent application data                     |
| Application Data, Configuration | ~/Library/Application Support | Persistent configuration data files and "private" application data |
| Applications                    | ~/Applications                | Applications                                                       |

```

When writing a cross-platform application, some developers like to enable XDG support on Darwin, too (see Linux/Unix section).

### Please share the word.

Users matter. Please help the movement to clean up users’ $HOME directories and standardise configuration by spreading the word, opening issues (preferably with PRs, not to annoy maintainers!), and contributing to a world where application data can roam in nature preserve of specification.

### Contact

If you have opinions or corrections please share them with me at `offsetcyan <at> proton.me`. Thanks!
