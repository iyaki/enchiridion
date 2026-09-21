---
title: "semgrep - Lightweight static analysis for many languages. Find bug variants with patterns that look like source code"
notion_id: 10d2ac6b-e240-484c-88c9-92905d5fccfc
notion_url: https://app.notion.com/p/semgrep-Lightweight-static-analysis-for-many-languages-Find-bug-variants-with-patterns-that-look--10d2ac6be240484c88c992905d5fccfc
last_edited: 2023-10-05T13:17:00.000Z
source_url: https://github.com/returntocorp/semgrep
tags: ["Tool", "English", "Information Security", "Programming"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Code scanning at ludicrous speed.

Semgrep is a fast, open-source, static analysis engine for finding bugs, detecting vulnerabilities in third-party dependencies, and enforcing code standards. Semgrep analyzes code locally on your computer or in your build environment: **code is never uploaded**. [Get started →.](https://github.com/returntocorp/semgrep?utm_source=tldrwebdev#getting-started-)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Language support

Semgrep supports 30+ languages.

| Category | Languages |
| --- | --- |
| GA | C# · Go · Java · JavaScript · JSX · JSON · PHP · Python · Ruby · Scala · Terraform · TypeScript · TSX |
| Beta | Kotlin · Rust |
| Experimental | Bash · C · C++ · Clojure · Dart · Dockerfile · Elixir · HTML · Julia · Jsonnet · Lisp · Lua · OCaml · R · Scheme · Solidity · Swift · YAML · XML · Generic (ERB, Jinja, etc.) |

### Getting started 🚀

1. [From the CLI](https://github.com/returntocorp/semgrep?utm_source=tldrwebdev#option-1-getting-started-from-the-cli)
2. [From the Semgrep Cloud Platform](https://github.com/returntocorp/semgrep?utm_source=tldrwebdev#option-2-getting-started-from-the-semgrep-cloud-platform-recommended)

For beginners, we recommend starting with the [Semgrep Cloud Platform](https://github.com/returntocorp/semgrep?utm_source=tldrwebdev#option-2-getting-started-from-the-semgrep-cloud-platform-recommended) because it provides a visual interface, a demo project, result triaging and exploration workflows, and makes setup in CI/CD fast. Scans are still local and code isn't uploaded. Alternatively, you can also start with the CLI without logging in and navigate the terminal output to run one-off searches.

### Option 1: Getting started from the CLI

1. Install Semgrep CLI

```plain text
# For macOS
$ brew install semgrep

# For Ubuntu/WSL/Linux/macOS
$ python3 -m pip install semgrep

# To try Semgrep without installation run via Docker
$ docker run --rm -v "${PWD}:/src" returntocorp/semgrep semgrep

```

1. 

Go to your app's root directory and run `semgrep scan --config auto`. This will scan your project with the default settings.

1. 

[Optional, but recommended] Run `semgrep login` to get the login URL for the Semgrep Cloud Platform. Open the login URL in the browser and login.

### Option 2: Getting started from the Semgrep Cloud Platform (Recommended)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

1. 

Register to [semgrep.dev](https://go.semgrep.dev/login-ghrmgo)

1. 

Explore the demo app

1. 

Scan your project by navigating to `Projects > Scan New Project > Run scan in CI`

1. 

Select your version control system and follow the wizard to add your project. After this setup, Semgrep will scan your project after every pull request.

1. 

[Optional but recommended] If you want to run Semgrep locally, follow the steps in the CLI section.

### Notes:

1. 

Visit [Docs > Running rules](https://semgrep.dev/docs/running-rules/) to learn more about `auto` config and other rules.

1. 

If there are any issues, [please ask for help in the Semgrep Slack](https://go.semgrep.dev/slack)

1. 

To run Semgrep Supply Chain, [contact the Semgrep team](https://semgrep.dev/contact-us). Visit the [full documentation](https://semgrep.dev/docs/getting-started/) to learn more.

### Semgrep Ecosystem

The Semgrep ecosystem includes the following products:

- Semgrep OSS Engine - The open-source engine at the heart of everything (this project).
- [Semgrep Cloud Platform (SCP)](https://semgrep.dev/login) - Deploy, manage, and monitor SAST and SCA at scale using Semgrep, with [free and paid tiers](https://semgrep.dev/pricing). Integrates with continuous integration (CI) providers such as GitHub, GitLab, CircleCI, and more.
- [Semgrep Code](https://semgrep.dev/products/semgrep-code) - Scan your code with Semgrep's Pro rules and Semgrep Pro Engine to find OWASP Top 10 vulnerabilities and protect against critical security risks specific to your organization. Semgrep Code is available on both [free and paid tiers](https://semgrep.dev/pricing).
- [Semgrep Supply Chain (SSC)](https://semgrep.dev/products/semgrep-supply-chain) - A high-signal dependency scanner that detects reachable vulnerabilities in open source third-party libraries and functions across the software development life cycle (SDLC). Semgrep Supply Chain is available on both [free and paid tiers](https://semgrep.dev/pricing).

and:

- [Semgrep Playground](https://semgrep.dev/editor) - An online interactive tool for writing and sharing rules.
- [Semgrep Registry](https://semgrep.dev/explore) - 2,000+ community-driven rules covering security, correctness, and dependency vulnerabilities.

Join hundreds of thousands of other developers and security engineers already using Semgrep at companies like GitLab, Dropbox, Slack, Figma, Shopify, HashiCorp, Snowflake, and Trail of Bits.

Semgrep is developed and commercially supported by [Semgrep, Inc., a software security company](https://semgrep.dev/).

### Semgrep Rules

Semgrep rules look like the code you already write; no abstract syntax trees, regex wrestling, or painful DSLs. Here's a quick rule for finding Python `print()` statements.

Run it online in Semgrep’s Playground by [clicking here](https://semgrep.dev/s/ievans:print-to-logger).

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Examples

Visit [Docs > Rule examples](https://semgrep.dev/docs/writing-rules/rule-ideas/) for use cases and ideas.

| Use case | Semgrep rule |
| --- | --- |
| Ban dangerous APIs | [Prevent use of exec](https://semgrep.dev/playground/s/lglB) |
| Search routes and authentication | [Extract Spring routes](https://semgrep.dev/playground/s/Y6wD) |
| Enforce the use secure defaults | [Securely set Flask cookies](https://semgrep.dev/playground/s/6KwW) |
| Tainted data flowing into sinks | [ExpressJS dataflow into sandbox.run](https://semgrep.dev/playground/s/qEpR) |
| Enforce project best-practices | [Use assertEqual for == checks](https://semgrep.dev/playground/s/oEox), [Always check subprocess calls](https://semgrep.dev/playground/s/zENk) |
| Codify project-specific knowledge | [Verify transactions before making them](https://semgrep.dev/playground/s/p8zk) |
| Audit security hotspots | [Finding XSS in Apache Airflow](https://semgrep.dev/playground/s/KPwj), [Hardcoded credentials](https://semgrep.dev/playground/s/2Br8) |
| Audit configuration files | [Find S3 ARN uses](https://semgrep.dev/playground/s/jEKD) |
| Migrate from deprecated APIs | [DES is deprecated](https://semgrep.dev/playground/r/java.lang.security.audit.crypto.des-is-deprecated.des-is-deprecated), [Deprecated Flask APIs](https://semgrep.dev/playground/r/python.flask.maintainability.deprecated.deprecated-apis.flask-deprecated-apis), [Deprecated Bokeh APIs](https://semgrep.dev/playground/r/python.bokeh.maintainability.deprecated.deprecated_apis.bokeh-deprecated-apis) |
| Apply automatic fixes | [Use listenAndServeTLS](https://semgrep.dev/playground/s/1Ayk) |

### Extensions

Visit [Docs > Extensions](https://semgrep.dev/docs/extensions/) to learn about using Semgrep in your editor or pre-commit. When integrated into CI and configured to scan pull requests, Semgrep will only report issues introduced by that pull request; this lets you start using Semgrep without fixing or ignoring pre-existing issues!

### Documentation

Browse the full Semgrep [documentation on the website](https://semgrep.dev/docs). If you’re new to Semgrep, check out [Docs > Getting started](https://semgrep.dev/docs/getting-started/) or the [interactive tutorial](https://semgrep.dev/learn).

### Metrics

Using remote configuration from the [Registry](https://semgrep.dev/r) (like `--config=p/ci`) reports pseudonymous rule metrics to semgrep.dev.

Using configs from local files (like `--config=xyz.yml`) does **not** enable metrics.

To disable Registry rule metrics, use `--metrics=off`.

The Semgrep [privacy policy](https://semgrep.dev/docs/metrics) describes the principles that guide data-collection decisions and the breakdown of the data that are and are not collected when the metrics are enabled.

### More

- [Frequently asked questions (FAQs)](https://semgrep.dev/docs/faq/)
- [Contributing](https://semgrep.dev/docs/contributing/contributing/)
- [Build instructions for developers](https://github.com/returntocorp/semgrep/blob/develop/INSTALL.md)
- [Ask questions in the Semgrep community Slack](https://go.semgrep.dev/slack)
- [CLI reference and exit codes](https://semgrep.dev/docs/cli-usage)
- [Semgrep YouTube channel](https://www.youtube.com/c/semgrep)
- [License (LGPL-2.1)](https://github.com/returntocorp/semgrep/blob/develop/LICENSE)

### Upgrading

To upgrade, run the command below associated with how you installed Semgrep:

```plain text
# Using Homebrew
$ brew upgrade semgrep

# Using pip
$ python3 -m pip install --upgrade semgrep

# Using Docker
$ docker pull returntocorp/semgrep:latest
```
