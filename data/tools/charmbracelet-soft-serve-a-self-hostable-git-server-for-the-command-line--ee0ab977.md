---
title: "charmbracelet/soft-serve - A self-hostable Git server for the command line"
notion_id: ee0ab977-e0a2-4d90-b01b-c49917bfd103
notion_url: https://app.notion.com/p/charmbracelet-soft-serve-A-self-hostable-Git-server-for-the-command-line-ee0ab977e0a24d90b01bc49917bfd103
last_edited: 2023-07-28T15:07:00.000Z
source_url: https://github.com/charmbracelet/soft-serve
tags: ["English", "Programming", "Hosting", "Untried", "Tool"]
---
# Soft Serve

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

A tasty, self-hostable Git server for the command line.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

- Easy to navigate TUI available over SSH
- Clone repos over SSH, HTTP, or Git protocol
- Manage repos with SSH
- Create repos on demand with SSH or `git push`
- Browse repos, files and commits with SSH-accessible
- Print files over SSH with or without syntax highlighting and line numbers
- Easy access control with SSH 
- Allow/disallow anonymous access
- Add collaborators with SSH public keys
- Repos can be public or private

## Where can I see it?

Just run `ssh git.charm.sh` for an example. You can also try some of the following commands:

```plain text
# Jump directly to a repo in the TUI
ssh git.charm.sh -t soft-serve

# Print out a directory tree for a repo
ssh git.charm.sh repo tree soft-serve

# Print a specific file
ssh git.charm.sh repo blob soft-serve cmd/soft/root.go

# Print a file with syntax highlighting and line numbers
ssh git.charm.sh repo blob soft-serve cmd/soft/root.go -c -l
```

## Installation

Soft Serve is a single binary called `soft`. You can get it from a package manager:

```plain text
# macOS or Linux
brew tap charmbracelet/tap && brew install charmbracelet/tap/soft-serve

# Arch Linux
pacman -S soft-serve

# Nix
nix-env -iA nixpkgs.soft-serve

# Debian/Ubuntu
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://repo.charm.sh/apt/gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/charm.gpg
echo "deb [signed-by=/etc/apt/keyrings/charm.gpg] https://repo.charm.sh/apt/ * *" | sudo tee /etc/apt/sources.list.d/charm.list
sudo apt update && sudo apt install soft-serve

# Fedora/RHEL
echo '[charm]
name=Charm
baseurl=https://repo.charm.sh/yum/
enabled=1
gpgcheck=1
gpgkey=https://repo.charm.sh/yum/gpg.key' | sudo tee /etc/yum.repos.d/charm.repo
sudo yum install soft-serve
```

You can also download a binary from the [releases](https://github.com/charmbracelet/soft-serve/releases) page. Packages are available in Alpine, Debian, and RPM formats. Binaries are available for Linux, macOS, and Windows.

Or just install it with `go`:

```plain text
go install github.com/charmbracelet/soft-serve/cmd/soft@latest
```

A [Docker image](https://github.com/charmbracelet/soft-serve/blob/main/docker.md) is also available.

## Setting up a server

Make sure `git` is installed, then run `soft serve`. That’s it.

This will create a `data` directory that will store all the repos, ssh keys, and database.

To change the default data path use `SOFT_SERVE_DATA_PATH` environment variable.

```plain text
SOFT_SERVE_DATA_PATH=/var/lib/soft-serve soft serve
```

When you run Soft Serve for the first time, make sure you have the `SOFT_SERVE_INITIAL_ADMIN_KEYS` environment variable is set to your ssh authorized key. Any added key to this variable will be treated as admin with full privileges.

Using this environment variable, Soft Serve will create a new `admin` user that has full privileges. You can rename and change the user settings later.

Check out [Systemd](https://github.com/charmbracelet/soft-serve/blob/main/systemd.md) on how to run Soft Serve as a service using Systemd. Soft Serve packages in our Apt/Yum repositories come with Systemd service units.

### Server Settings

Once you start the server for the first time, the settings will be in `config.yaml` under your data directory. The default `config.yaml` is self-explanatory and will look like this:

```plain text
# Soft Serve Server configurations

# The name of the server.
# This is the name that will be displayed in the UI.
name: "Soft Serve"

# Log format to use. Valid values are "json", "logfmt", and "text".
log_format: "text"

# The SSH server configuration.
ssh:
 # The address on which the SSH server will listen.
 listen_addr: ":23231"

 # The public URL of the SSH server.
 # This is the address that will be used to clone repositories.
 public_url: "ssh://localhost:23231"

 # The path to the SSH server's private key.
 key_path: "ssh/soft_serve_host"

 # The path to the SSH server's client private key.
 # This key will be used to authenticate the server to make git requests to
 # ssh remotes.
 client_key_path: "ssh/soft_serve_client"

 # The maximum number of seconds a connection can take.
 # A value of 0 means no timeout.
 max_timeout: 0

 # The number of seconds a connection can be idle before it is closed.
 idle_timeout: 120

# The Git daemon configuration.
git:
 # The address on which the Git daemon will listen.
 listen_addr: ":9418"

 # The maximum number of seconds a connection can take.
 # A value of 0 means no timeout.
 max_timeout: 0

 # The number of seconds a connection can be idle before it is closed.
 idle_timeout: 3

 # The maximum number of concurrent connections.
 max_connections: 32

# The HTTP server configuration.
http:
 # The address on which the HTTP server will listen.
 listen_addr: ":23232"

 # The path to the TLS private key.
 tls_key_path: ""

 # The path to the TLS certificate.
 tls_cert_path: ""

 # The public URL of the HTTP server.
 # This is the address that will be used to clone repositories.
 # Make sure to use https:// if you are using TLS.
 public_url: "http://localhost:23232"

# The stats server configuration.
stats:
 # The address on which the stats server will listen.
 listen_addr: ":23233"
# Additional admin keys.
#initial_admin_keys:
# - "ssh-rsa AAAAB3NzaC1yc2..."
```

You can also use environment variables, to override these settings. All server settings environment variables start with `SOFT_SERVE_` followed by the setting name all in uppercase. Here are some examples:

- `SOFT_SERVE_NAME`: The name of the server that will appear in the TUI
- `SOFT_SERVE_SSH_LISTEN_ADDR`: SSH listen address
- `SOFT_SERVE_SSH_KEY_PATH`: SSH host key-pair path
- `SOFT_SERVE_HTTP_LISTEN_ADDR`: HTTP listen address
- `SOFT_SERVE_HTTP_PUBLIC_URL`: HTTP public URL used for cloning
- `SOFT_SERVE_GIT_MAX_CONNECTIONS`: The number of simultaneous connections to git daemon

## Configuration

Configuring Soft Serve is simple and straightforward. Use the SSH command-line interface to manage access settings, users, and repos.

Try `ssh localhost -i ~/.ssh/id_ed25519 -p 23231 help` for more info. Make sure you use your key here.

For ease of use, instead of specifying the key, port, and hostname every time you SSH into Soft Serve, add your own Soft Serve instance entry to your SSH config. For instance, to use `ssh soft` instead of typing `ssh localhost -i ~/.ssh/id_ed25519 -p 23231`, we can define a `soft` entry in our SSH config file `~/.ssh/config`.

```plain text
Host soft
 HostName localhost
 Port 23231
 IdentityFile ~/.ssh/id_ed25519

```

Now, we can do `ssh soft` to SSH into Soft Serve. Since `git` is also aware of this config, you can use `soft` as the hostname for your clone commands.

```plain text
git clone ssh://soft/dotfiles
# make changes
# add & commit
git push origin main
```

> 

```plain text
-i
```

### Access Levels

Soft Serve offers a simple access control. There are four access levels, no-access, read-only, read-write, and admin-access.

`admin-access` has full control of the server and can make changes to users and repos.

`read-write` access gets full control of repos.

`read-only` can read public repos.

`no-access` denies access to all repos.

### Authentication

Everything that needs authentication is done using SSH. Make sure you have added an entry for your Soft Serve instance in your `~/.ssh/config` file.

By default, Soft Serve gives ready-only permission to anonymous connections to any of the above protocols. This is controlled by two settings `anon-access` and `allow-keyless`.

- `anon-access`: Defines the access level for anonymous users. Available options are `no-access`, `read-only`, `read-write`, and `admin-access`. Default is `read-only`.
- `allow-keyless`: Whether to allow connections that doesn't use keys to pass. Setting this to `false` would disable access to SSH keyboard-interactive, HTTP, and Git protocol connections. Default is `true`.

```plain text
$ ssh -p 23231 localhost settings
Manage server settings

Usage:
 ssh -p 23231 localhost settings [command]

Available Commands:
 allow-keyless Set or get allow keyless access to repositories
 anon-access Set or get the default access level for anonymous users

Flags:
 -h, --help help for settings

Use "ssh -p 23231 localhost settings [command] --help" for more information about a command.
```

> 

When `allow-keyless` is disabled, connections that don't use SSH Public Key authentication will get denied. This means cloning repos over HTTP(s) or git:// will get denied.

Meanwhile, `anon-access` controls the access level granted to connections that use SSH Public Key authentication but are not registered users. The default setting for this is `read-only`. This will grant anonymous connections that use SSH Public Key authentication `read-only` access to public repos.

`anon-access` is also used in combination with `allow-keyless` to determine the access level for HTTP(s) and git:// clone requests.

## Authorization

Admins can manage users and their keys using the `user` command. Once a user is created and has access to the server, they can manage their own keys and settings.

To create a new user simply use `user create`:

```plain text
# Create a new user
ssh -p 23231 localhost user create beatrice

# Add user keys
ssh -p 23231 localhost user add-pubkey beatrice ssh-rsa AAAAB3Nz...
ssh -p 23231 localhost user add-pubkey beatrice ssh-ed25519 AAAA...

# Create another user with public key
ssh -p 23231 localhost user create frankie '-k "ssh-ed25519 AAAATzN..."'

# Need help?
ssh -p 23231 localhost user help
```

Once a user is created, they get `read-only` access to public repositories. They can also create new repositories on the server.

Users can manage their keys using the `pubkey` command:

```plain text
# List user keys
ssh -p 23231 localhost pubkey list

# Add key
ssh -p 23231 localhost pubkey add ssh-ed25519 AAAA...

# Wanna change your username?
ssh -p 23231 localhost set-username yolo

# To display user info
ssh -p 23231 localhost info
```

## Repositories

You can manage repositories using the `repo` command.

```plain text
# Run repo help
$ ssh -p 23231 localhost repo help
Manage repositories

Usage:
 ssh -p 23231 localhost repo [command]

Aliases:
 repo, repos, repository, repositories

Available Commands:
 blob Print out the contents of file at path
 branch Manage repository branches
 collab Manage collaborators
 create Create a new repository
 delete Delete a repository
 description Set or get the description for a repository
 hide Hide or unhide a repository
 import Import a new repository from remote
 info Get information about a repository
 is-mirror Whether a repository is a mirror
 list List repositories
 private Set or get a repository private property
 project-name Set or get the project name for a repository
 rename Rename an existing repository
 tag Manage repository tags
 tree Print repository tree at path

Flags:
 -h, --help help for repo

Use "ssh -p 23231 localhost repo [command] --help" for more information about a command.
```

To use any of the above `repo` commands, a user must be a collaborator in the repository. More on this below.

### Creating Repositories

To create a repository, first make sure you are a registered user. Use the `repo create <repo>` command to create a new repository:

```plain text
# Create a new repository
ssh -p 23231 localhost repo create icecream

# Create a repo with description
ssh -p 23231 localhost repo create icecream '-d "This is an Ice Cream description"'

# ... and project name
ssh -p 23231 localhost repo create icecream '-d "This is an Ice Cream description"' '-n "Ice Cream"'

# I need my repository private!
ssh -p 23231 localhost repo create icecream -p '-d "This is an Ice Cream description"' '-n "Ice Cream"'

# Help?
ssh -p 23231 localhost repo create -h
```

Or you can add your Soft Serve server as a remote to any existing repo, given you have write access, and push to remote:

```plain text
git remote add origin ssh://localhost:23231/icecream

```

After you’ve added the remote just go ahead and push. If the repo doesn’t exist on the server it’ll be created.

```plain text
git push origin main

```

Repositories can be nested too:

```plain text
# Create a new nested repository
ssh -p 23231 localhost repo create charmbracelet/icecream

# Or ...
git remote add charm ssh://localhost:23231/charmbracelet/icecream
git push charm main
```

### Deleting Repositories

You can delete repositories using the `repo delete <repo>` command.

```plain text
ssh -p 23231 localhost repo delete icecream
```

### Renaming Repositories

Use the `repo rename <old> <new>` command to rename existing repositories.

```plain text
ssh -p 23231 localhost repo rename icecream vanilla
```

### Repository Collaborators

Sometimes you want to restrict write access to certain repositories. This can be achieved by adding a collaborator to your repository.

Use the `repo collab <command> <repo>` command to manage repo collaborators.

```plain text
# Add collaborator to soft-serve
ssh -p 23231 localhost repo collab add soft-serve frankie

# Remove collaborator
ssh -p 23231 localhost repo collab remove soft-serve beatrice

# List collaborators
ssh -p 23231 localhost repo collab list soft-serve
```

### Repository metadata

You can also change the repo's description, project name, whether it's private, etc using the `repo <command>` command.

```plain text
# Set description for repo
ssh -p 23231 localhost repo description icecream "This is a new description"

# Hide repo from listing
ssh -p 23231 localhost repo hidden icecream true

# List repository info (branches, tags, description, etc)
ssh -p 23231 localhost repo icecream info
```

To make a repository private, use `repo private <repo> [true|false]`. Private repos can only be accessed by admins and collaborators.

```plain text
ssh -p 23231 localhost repo icecream private true
```

### Repository Branches & Tags

Use `repo branch` and `repo tag` to list, and delete branches or tags. You can also use `repo branch default` to set or get the repository default branch.

### Repository Tree

To print a file tree for the project, just use the `repo tree` command along with the repo name as the SSH command to your Soft Serve server:

```plain text
ssh -p 23231 localhost repo tree soft-serve
```

You can also specify the sub-path and a specific reference or branch.

```plain text
ssh -p 23231 localhost repo tree soft-serve server/config
ssh -p 23231 localhost repo tree soft-serve main server/config
```

From there, you can print individual files using the `repo blob` command:

```plain text
ssh -p 23231 localhost repo blob soft-serve cmd/soft/root.go
```

You can add the `-c` flag to enable syntax coloring and `-l` to print line numbers:

```plain text
ssh -p 23231 localhost repo blob soft-serve cmd/soft/root.go -c -l

```

Use `--raw` to print raw file contents. This is useful for dumping binary data.

## The Soft Serve TUI

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Soft Serve serves a TUI over SSH for browsing repos, viewing files and commits, and grabbing clone commands:

```plain text
ssh localhost -p 23231
```

It's also possible to “link” to a specific repo:

```plain text
ssh -p 23231 localhost -t soft-serve
```

You can copy text to your clipboard over SSH. For instance, you can press c on the highlighted repo in the menu to copy the clone command [1](https://github.com/charmbracelet/soft-serve#user-content-fn-osc52-496cb481288e5546042b7d44b05d2138).

## Hooks

Soft Serve supports git server-side hooks `pre-receive`, `update`, `post-update`, and `post-receive`. This means you can define your own hooks to run on repository push events. Hooks can be defined as a per-repository hook, and/or global hooks that run for all repositories.

You can find per-repository hooks under the repository `hooks` directory.

Globs hooks can be found in your `SOFT_SERVE_DATA_PATH` directory under `hooks`. Defining global hooks is useful if you want to run CI/CD for example.

Here's an example of sending a message after receiving a push event. Create an executable file `<data path>/hooks/update`:

```plain text
#!/bin/sh
#
# An example hook script to echo information about the push
# and send it to the client.

refname="$1"
oldrev="$2"
newrev="$3"

# Safety check
if [ -z "$GIT_DIR" ]; then
 echo "Don't run this script from the command line." >&2
 echo " (if you want, you could supply GIT_DIR then run" >&2
 echo " $0 <ref> <oldrev> <newrev>)" >&2
 exit 1
fi

if [ -z "$refname" -o -z "$oldrev" -o -z "$newrev" ]; then
 echo "usage: $0 <ref> <oldrev> <newrev>" >&2
 exit 1
fi

# Check types
# if $newrev is 0000...0000, it's a commit to delete a ref.
zero=$(git hash-object --stdin </dev/null | tr '[0-9a-f]' '0')
if [ "$newrev" = "$zero" ]; then
 newrev_type=delete
else
 newrev_type=$(git cat-file -t $newrev)
fi

echo "Hi from Soft Serve update hook!"
echo
echo "RefName: $refname"
echo "Change Type: $newrev_type"
echo "Old SHA1: $oldrev"
echo "New SHA1: $newrev"

exit 0
```

Now, you should get a message after pushing changes to any repository.

## A note about RSA keys

Unfortunately, due to a shortcoming in Go’s `x/crypto/ssh` package, Soft Serve does not currently support access via new SSH RSA keys: only the old SHA-1 ones will work.

Until we sort this out you’ll either need an SHA-1 RSA key or a key with another algorithm, e.g. Ed25519. Not sure what type of keys you have? You can check with the following:

```plain text
$ find ~/.ssh/id_*.pub -exec ssh-keygen -l -f {} \;
```

If you’re curious about the inner workings of this problem have a look at:

- [golang/go#37278](https://github.com/golang/go/issues/37278)
- [https://go-review.googlesource.com/c/crypto/+/220037](https://go-review.googlesource.com/c/crypto/+/220037)
- [golang/crypto#197](https://github.com/golang/crypto/pull/197)

## Feedback

We’d love to hear your thoughts on this project. Feel free to drop us a note!

- [Twitter](https://twitter.com/charmcli)
- [The Fediverse](https://mastodon.social/@charmcli)
- [Discord](https://charm.sh/chat)

## License

[MIT](https://github.com/charmbracelet/soft-serve/raw/main/LICENSE)

Part of [Charm](https://charm.sh/).

Charm热爱开源 • Charm loves open source

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Footnotes

1. 

Copying over SSH depends on your terminal support of OSC52. Refer to [go-osc52](https://github.com/aymanbagabas/go-osc52) for more information.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->
