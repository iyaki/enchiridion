---
title: "blst-security/cherrybomb - CLI tool that helps you avoid undefined user behaviour by auditing your API specifications, validating them and running API security tests"
notion_id: b3f65260-b5b7-454d-b9ae-ea48e146fdad
notion_url: https://app.notion.com/p/blst-security-cherrybomb-CLI-tool-that-helps-you-avoid-undefined-user-behaviour-by-auditing-your-A-b3f65260b5b7454db9aeea48e146fdad
last_edited: 2023-06-21T15:09:00.000Z
source_url: https://github.com/blst-security/cherrybomb
tags: ["Tool", "English", "REST API", "Backend"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# Stop half-done API specifications

# What is Cherrybomb?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Cherrybomb is an CLI tool written in Rust that helps prevent incorrect code implementation early in development. It works by validating and testing your API using an OpenAPI file. Its main goal is to reduce security errors and ensure your API functions as intended.

# How does it work?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Cherrybomb makes sure your API is working correctly. It checks your API's spec file (OpenAPI Specification) for good practices and makes sure it follows the OAS rules. Then, it tests your API for common issues and vulnerabilities. If any problems are found, Cherrybomb gives you a detailed report with the exact location of the problem so you can fix it easily.

# Get Started

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Installation

### Linux/MacOS:

```plain text

curl https://cherrybomb.blstsecurity.com/install | /bin/bash


```

The script requires sudo permissions to move the cherrybomb bin into **/usr/local/bin/**.

(If you want to view the shell script(or even help to improving it - [/scripts/install.sh](https://github.com/blst-security/cherrybomb/blob/main/scripts/install.sh))

### Containerized version

You can get Cherrybomb through its containerized version which is hosted on AWS ECR, and requires an API key that you can get on that addess(the loading is a bit slow) - [https://cicd.blstsecurity.com/](https://cicd.blstsecurity.com/)

```plain text
docker run --mount type=bind,source=[PATH TO OAS],destination=/home public.ecr.aws/blst-security/cherrybomb:latest cherrybomb -f /home/[OAS NAME] --api-key=[API-KEY]

```

### Get it from crates.io

```plain text
cargo install cherrybomb

```

If you don't have cargo installed, you can install it from [here](https://doc.rust-lang.org/cargo/getting-started/installation.html)

### Building from Sources

You can also build Cherrybomb from sources by cloning this repo, and building it using cargo.

```plain text

git clone https://github.com/blst-security/cherrybomb && cd cherrybomb


```

The main branch's Cargo.toml file uses `cherrybomb-engine` and `cherrybomb-oas` from crates.io.

if you want build those from source too, you can change the following files:

(remove the version number and replace with the path to the local repo)

```plain text
cherrybomb/Cargo.toml:
cherrybomb-engine = version => { path = "cherrybomb-engine" }

```

```plain text
cherrybomb/cherrybomb-engine/Cargo.toml:
cherrybomb-oas = version => { path = "../cherrybomb-oas" }

```

```plain text
cargo build --release
sudo mv ./target/release/cherrybomb /usr/local/bin # or any other directory in your PATH

```

### Profile

Profiles allow you to choose the type of check you want to use.

```plain text
- info: only generates param and endpoint tables
- normal: both active and passive
- intrusive: active and intrusive [in development]
- passive: only passive tests
- full: all the options

```

### Config

With a configuration file, you can easily edit, view, Cherrybomb's options. The config file allows you to set the running profile, location of the oas file, the verbosity and ignore the TLS error.

Config also allows you to override the server's URL with an array of servers, and add security to the request [in development].

Notice that CLI arguments parameter will override config options if both are set.

You can also add or remove checks from a profile using `passive/active-include/exclude`. [in development]

```plain text
cherrybomb --config <CONFIG_FILE>

```

Structure of config file:

```plain text
{
"file" : "open-api.json",
"verbosity" : "normal,
"profile" : " "Normal",
"passive_include" : ["check1, checks2"],
"active_include": ["check3, check4"],
"servers_override" , ["http://server/"],
"security": [{
 "auth_type": "Basic",
 "auth_value" : token_value,
 "auth_scope" : scope_name
 }],
"ignore_tls_errors" : true,
"no_color" : false,
}

```

# Usage

After installing, verify it's working by running

```plain text
cherrybomb --version


```

### OpenAPI specification

`cherrybomb --file <PATH> --profile passive`

Passive Output example:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Generate Info Table

```plain text
cherrybomb --file <PATH> --profile info


```

Parameter table output:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Endpoint table output:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# Integration

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

You can embed it into your CI pipeline, and If you plan on doing that I would recommend that you go to our [website](https://www.blstsecurity.com/?promo=blst&domain=github_integration_link), sign up, go through the [CI pipeline integration wizard](https://www.blstsecurity.com/Loading?redirect=%2FCICD&promo=blst&domain=github_wiz_integration), and copy the groovy/GitHub actions snippet built for you.

Example:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# Support

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Get help

If you have any questions, please send us a message to [support@blstsecurity.com](mailto:support@blstsecurity.com) or ask us on our [discord server](https://discord.gg/WdHhv4DqwU).

You are also welcome to open an Issue here on GitHub.
