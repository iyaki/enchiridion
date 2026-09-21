---
title: "Create an internal CLI"
notion_id: 66b731cd-5526-4d75-b345-320d7ccdfb08
notion_url: https://app.notion.com/p/Create-an-internal-CLI-66b731cd55264d75b345320d7ccdfb08
last_edited: 2024-09-17T16:29:00.000Z
source_url: https://blog.chay.dev/create-an-internal-cli
tags: ["English", "DevOps", "Site Reliability Engineering", "Project Management", "Article", "blog.chay.dev"]
---
Assumed Audience: programmers who know their way around a terminal.

One useful way to capture and preserve institutional knowledge in a dev team is to grow a collection of useful snippets, scripts, or workflows. This is why almost every repo out there has some way to write and run tasks via `Makefile`s, bash scripts, or language-specific toolings like [scripts in package.json](https://docs.npmjs.com/cli/v10/using-npm/scripts) for Javascript, or [mix tasks](https://hexdocs.pm/mix/1.17.2/Mix.Task.html) for Elixir.

What about org-wide things, like installing useful tools, generating boilerplate code, or running complex AWS commands that no one remembers? Some companies like [Slack](https://slack.engineering/the-joy-of-internal-tools/) and [Shopify](https://blog.nelhage.com/post/stripe-dev-environment/) have their own internal CLI. The modern terminal Warp has a [pretty sick feature](https://docs.warp.dev/features/warp-drive) for documenting and sharing workflows.

It's quite easy to get started with your very own CLI. Let's create one right now for our hypothetical company [Acme Corporation](https://en.wikipedia.org/wiki/Acme_Corporation).

Here's a list of simple requirements:

1. **Common entrypoint** for all commands. All devs should trigger commands by doing `acme <command>` from anywhere, rather than having to navigate to a specific repo first.
2. Devs can easily **contribute** new commands. No need to learn a brand new language or complex syntax.
3. New versions are **distributed easily**. There should be a convenient way to contribute new commands. It should be easy to fetch new updates via `acme update`.
4. Commands should work **across platforms**. For example, `acme download something` command could use `curl` on Linux, and `Invoke-WebRequest` on Windows.
5. Commands should be **discoverable**. Calling `acme list` should enumerate the list of available commands and their short descriptions.

Let's use [`just`](https://github.com/casey/just) for this project. `just` is similar to `make` but designed to be a command runner. It is cross-platform and enables the possibility of running platform-specific commands. [Lots of repos](https://github.com/search?q=justfile) use it.

Some other possible tools we could use:

- Slack's [`magic-cli`](https://github.com/slackhq/magic-cli): it's a fantastic starting point if your devs know Ruby.
- `make`: Makefiles are simple, available out of the box on Linux and Mac machines (but can be installed on Windows machines too), and are very, _very_ widely used.

Install `just`. Follow the instructions [here](https://github.com/casey/just/blob/master/README.md#installation).

Create a folder at `~/acme/cli` and add the following `justfile` at the root:

```plain text
default:
  just --list

# Show arch and os name
os-info:
  echo "Arch: {{arch()}}"
  echo "OS: {{os()}}"

```

The `just` documentation calls commands "recipes", so let's use that word from here on.

When `just` is invoked without a recipe, it runs the first recipe in the justfile. It is a common pattern to name the first recipe "default".

```plain text
$ just
just --list
Available recipes:
    default
    os-info # Show arch and os name

```

Our default recipe runs `just list`, which is the in-built mechanism to enumerate all available recipes, alongside what we documented as comments.

Let's hide the `default` recipe in the list, since there's no point showing it.

Notice that running the recipe prints each command before it is executed. Let's suppress this output using a `@` prefix. This is quite [similar](https://www.gnu.org/software/make/manual/html_node/Echoing.html#Echoing) to how Makefiles work. We can add this prefix on each line we want to suppress, or add it to the recipe name to suppress the output for all lines.

```plain text
[private]
@default:
  just --list

# Show arch and os name
@os-info:
  echo "Arch: {{arch()}}"
  echo "OS: {{os()}}"

```

We want to run these recipes using `acme <command>` instead of `just <command>`. Let's now add the `acme` alias to our `.bashrc` (or whatever rc file you use):

```plain text
alias acme='just --justfile ~/acme/cli/justfile'

```

> 

In the [Forwarding Alias](https://github.com/casey/just?tab=readme-ov-file#forwarding-alias) section of the README, `just`'s creator said "I'm pretty sure that nobody actually uses this feature, but it's there". Hah, in your face, Casey!

Load our new alias using `source ~/.bashrc` or refresh the shell using `exec bash`.

Now, let's try it out:

Alright, now we have our very own `acme` CLI. Roll credits!

> 

Bonus exercise: try creating a `setup` recipe to do this automatically. Remember to cater for different shells and operating systems. Good luck!

### Simple recipes

Here's a simple but useful recipe for retrieving the current AWS IAM identity using `awscli`:

Simplifying commands that no one remembers is probably the top use case of internal CLIs.

Note that we are making the assumption that `awscli` is reasonably cross-platform, so this recipe should work regardless of where we call it from.

> 

Bonus exercise: create a `ensure-aws` recipe that detects the presence of `awscli` and install it automatically if it doesn't exist. Use it as a [dependency](https://github.com/casey/just?tab=readme-ov-file#dependencies) for any recipes that use `aws`.

### Platform-specific recipes

Snippets that involve tools like `systemd` are only relevant for Linux users, should they be exposed only if the dev is on a Linux machine.

By tagging the recipe with a `[linux]` [attribute](https://github.com/casey/just?tab=readme-ov-file#attributes), we can selectively enable a recipe only on Linux. I'm using a macbook now, let's see if it works:

It's not in the list! But what if we try to run to run the recipe anyway?

So it's not just _hidden_, the recipe straight up _does not exist_. This is a great guardrail against devs who might not be aware that the recipe is not available on the respective platform. The error message can be better though..

Let's verify that it works on a Linux machine:

### Cross-platform recipes

No one ever remembers how to get the size of a folder, so let's implement a `acme get-folder-size <path>`.

We add a `[no-cd]` attribute to the command for this recipe to run relatively to where the command is invoked. By [default](https://github.com/casey/just?tab=readme-ov-file#working-directory), `acme` will run with the working directory set to the directory that contains the `justfile`.

Windows doesn't have `du`, so we have to find another way to do it. We should also make it explicit the different shells we use for Linux and Windows:

Now we have a `acme get-folder-size` for both Windows and Linux!

### Scripts

We can embed entire scripts into our recipes! Simply start a recipe with a shebang (the `#!` part) and under the hood, `just` will save the content as a file and execute it.

This is useful if our workflow requires slightly more complex logic, like using control flows (if-else, loops), storing and manipulating variables, etc.

This also means we can utilise programming languages with strong scripting abilities. Some things are easier to do in Python than in Bash.

Not all devs have Python on their machines, and even if they do, they may not have `pillow` installed. We can use `nix` to run scripts with dependencies included:

> Yes, I use nix btw

Rather than rolling our own distribution mechanism, let's just use `git`.

Acme Corp uses GitHub. Create a remote repository first, then turn what we have right now into a git repo and push it up.

Now, anyone with access to this repo can contribute their changes by making a PR. Updating `acme` is also a simple matter of doing a `git pull`. Let's create a recipe to do that:

Note that `acme` runs from the working directory of the justfile by default, so this update recipe can be run anywhere.

> Bonus exercise: create a mechanism for running `acme update` periodically. Use `systemd` or whatever you believe in. Bonus points for creating a `acme setup-auto-update` recipe!

Adoption is critical to the success of an internal tool, and a good `README` that guides a new user to install and explore the tool is a major prerequisite.

`acme` is now ready to be used by all Acme Corp devs! We can now post a Slack message to encourage everyone to try it, and contribute their own snippets.

### Use `nushell` for cross-platform scripts

This introduces yet another dependency, but [`nu`](https://www.nushell.sh/) is truly delightful to use. It offers an intuitive way for manipulating the results of commands, replacing tools like `jq`, `awk`, or `grep`. And it's cross-platform!

### Completions

Completion is the mechanism that allows you to autocomplete subcommands, file paths, options, etc. when you hit the TAB key. Most shells offer this feature, most major CLI tools offer a way to install completions, and most major CLI frameworks - like [click](https://click.palletsprojects.com/en/8.1.x/shell-completion/) for Python, [cobra](https://github.com/spf13/cobra/blob/main/site/content/completions/_index.md) for Golang, and [clap](https://github.com/clap-rs/clap/tree/master/clap_complete) for Rust - can generate them automatically.

Just can generate completions by running `just --completion <shell>`. However, the generated completion works for the `just` command, so we'll need to change the relevant parts from `just` to `acme`. I'll leave this as a bonus exercise for the reader since this is pretty easy, but I encourage checking the modified completion script into the git repo so devs won't have to do it themselves.


