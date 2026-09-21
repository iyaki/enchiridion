---
title: "Noir - An attack surface detector form source code"
notion_id: e0978b4a-3eb6-4a52-bc1a-cc6a75d9a4a7
notion_url: https://app.notion.com/p/Noir-An-attack-surface-detector-form-source-code-e0978b4a3eb64a52bc1acc6a75d9a4a7
last_edited: 2023-08-28T12:56:00.000Z
source_url: https://github.com/hahwul/noir
tags: ["Information Security", "Untried", "Tool", "English"]
---
Noir is an attack surface detector form source code.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Key Features

- Automatically identify language and framework from source code.
- Find API endpoints and web pages through code analysis.
- Load results quickly through interactions with proxy tools such as ZAP, Burpsuite, Caido and More Proxy tools.
- That provides structured data such as JSON and HAR for identified Attack Surfaces to enable seamless interaction with other tools. Also provides command line samples to easily integrate and collaborate with other tools, such as curls or httpie.

## Available Support Scope

### Endpoint's Entities

- Path
- Method
- Param
- Header
- Protocol (e.g ws)

### Languages and Frameworks

<!-- unsupported block: child_database -->

### Specification

| Specification | Format | URL | Method | Param | Header | WS |
| --- | --- | --- | --- | --- | --- | --- |
| Swagger | JSON |  | ✅ |  | X | X |
| Swagger | YAML |  |  |  | X | X |

## Installation

### Homebrew (macOS)

```shell
brew tap hahwul/noir
brew install noir
```

### From Sources

```shell
# Install Crystal-lang
# https://crystal-lang.org/install/

# Clone this repo
git clone https://github.com/hahwul/noir
cd noir

# Install Dependencies
shards install

# Build
shards build --release --no-debug

# Copy binary
cp ./bin/noir /usr/bin/
```

### Docker (GHCR)

```shell
docker pull ghcr.io/hahwul/noir:main
```

## Usage

```plain text
Usage: noir <flags>
  Basic:
    -b PATH, --base-path ./app       (Required) Set base path
    -u URL, --url http://..          Set base url for endpoints
    -s SCOPE, --scope url,param      Set scope for detection

  Output:
    -f FORMAT, --format json         Set output format [plain/json/markdown-table/curl/httpie]
    -o PATH, --output out.txt        Write result to file
    --set-pvalue VALUE               Specifies the value of the identified parameter
    --no-color                       Disable color output
    --no-log                         Displaying only the results

  Deliver:
    --send-req                       Send the results to the web request
    --send-proxy http://proxy..      Send the results to the web request via http proxy

  Technologies:
    -t TECHS, --techs rails,php      Set technologies to use
    --exclude-techs rails,php        Specify the technologies to be excluded
    --list-techs                     Show all technologies

  Others:
    -d, --debug                      Show debug messages
    -v, --version                    Show version
    -h, --help                       Show help

```

Example

```shell
noir -b . -u https://testapp.internal.domains
```

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

JSON Result

```shell
noir -b . -u https://testapp.internal.domains -f json

```

```json
[
  ...
  {
    "headers": [],
    "method": "POST",
    "params": [
      {
        "name": "article_slug",
        "param_type": "json",
        "value": ""
      },
      {
        "name": "body",
        "param_type": "json",
        "value": ""
      },
      {
        "name": "id",
        "param_type": "json",
        "value": ""
      }
    ],
    "protocol": "http",
    "url": "https://testapp.internal.domains/comments"
  }
]
```
