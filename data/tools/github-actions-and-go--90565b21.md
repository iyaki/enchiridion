---
title: "Github Actions and Go"
notion_id: 90565b21-e66b-4867-95c8-2438b0c3b1fe
notion_url: https://app.notion.com/p/Github-Actions-and-Go-90565b21e66b486795c82438b0c3b1fe
last_edited: 2023-03-29T23:33:00.000Z
source_url: https://olegk.dev/github-actions-and-go
tags: ["English", "Continuous Integration/Continuous Delivery", "Go", "Tool", "Service"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

> TLDR: See [cristalhq/.github](https://github.com/cristalhq/.github/blob/53335fadafd4a2409b342643789aea60465dbf40/.github/workflows/build.yml) build workflow and how it can be used [cristalhq/jsn](https://github.com/cristalhq/jsn/blob/ab26cf0cee3c27ea8bb30707ee4f923e4f735031/.github/workflows/build.yml)

## Intro

I love open source, and also I love Go. So, a few months ago I decided to build the best CI for Go that I could easily reuse across my projects.

This post shares the results.

Note: Post is based on version `v0.5.0` of [cristalhq/.github](https://github.com/cristalhq/.github) repository.

## Github Actions

You probably know what GitHub Actions are, and because they are well-integrated with GitHub, I am too lazy to look at alternatives.

The only thing that I recommend is to look at [Workflow syntax](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions) which explains what syntax is available for us.

## Dependabot

I've noticed that not all developers are happy with Dependabot notifications (pull requests), and they find them noisy. For me, they do more good than bad. Consider just making the interval parameter bigger (a week or longer).

In almost all my projects I use the following Dependabot configuration:

```plain text
version: 2
updates:
  - package-ecosystem: gomod
    directory: "/"
    schedule:
      interval: daily
  - package-ecosystem: GitHub-actions
    directory: "/"
    schedule:
      interval: daily

```

Save the YAML above in `.github/dependabot.yml` and it's done.

## Basic workflow

The default Github Actions for Go looks like this:

```plain text
name: Go

on:
  push:
    branches: ["main"]
  pull_request:
    branches: ["main"]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: 1.19

    - name: Build
      run: go build -v ./...

    - name: Test
      run: go test -v ./...

```

In many cases it's ok but we can do better. Let's add more good things to it.

## Regular runs

A few times I got stuck with a failing CI because I missed some 3rd party changes and found them too late, running CI regularly prevents that:

```plain text
on:
  schedule:
  - cron: '0 10 * * 1' # run "At 10:00 on Monday"

```

A cron expression is an example, play on [crontab guru](https://crontab.guru/) to find yours.

## Runs on and Go version

Most of my projects do not interact with the operating system, so running on Ubuntu is enough.

It's a good practice to keep your code working on the current and previous Go versions. It's not a strict rule, but it might simplify life for you or your users when a new Go version is released.

The [actions/setup-go@v4](https://github.com/actions/setup-go) action names them as `stable` and `oldstable`:

```plain text
jobs:
  run:
    name: Build
    runs-on: ubuntu-latest
    timeout-minutes: 5 # just in case ¯\_(ツ)_/¯
    strategy:
      matrix:
        go: ['stable', 'oldstable']

```

## First steps:

Just check out the code and install Go. No magic:

```plain text
    steps:
    - name: Check out code
      uses: actions/checkout@v3

    - name: Install Go
      uses: actions/setup-go@v4
      with:
        go-version: ${{ matrix.go }}
        check-latest: true

```

`check-latest` set to `true` will pick the latest Go release for the given Go version. For example, for Go 1.20, it will download Go 1.20.1 (which is the latest at the moment of writing this post).

`actions/setup-go@v4` also supports caching out of the box. Please see [README section](https://github.com/actions/setup-go#caching-dependency-files-and-build-outputs). I don't have this configured because most of my projects are dependency-free or I don't care that much.

## Go steps

Enough to setup the CI job, time to do real work!

### Formatting and go vet

If code isn't formatted there is no sense to build it (or even to review). Same regarding `go vet`. It's the smallest linter but reports good things:

```plain text
      - name: Go Format
        run: gofmt -s -w . && git diff --exit-code

      - name: Go Vet
        run: go vet ./...

```

Thanks to [Daniel Martí](https://github.com/mvdan) `gofmt` is super fast now [Go #43566](https://github.com/golang/go/issues/43566).

### Check dependencies

I love to write dependency-free packages but sometimes I need something to import. So, the dependencies must be verified.

First, we need to check that `go.mod` is correct, download dependencies and verify `go.sum`:

```plain text
      - name: Go Tidy
        run: go mod tidy && git diff --exit-code

      - name: Go Mod
        run: go mod download

      - name: Go Mod Verify
        run: go mod verify

```

You might think that `go mod verify` is an extra step, but this will help you to catch force pushes (or even supply chain attacks) in your dependencies.

I heavily recommend adding this to all your workflows.

### Codegen & build

Maybe it's not the most popular feature in Go and there are different ways to manage this but I prefer and recommend you to commit generated files and always treat them as your code (well, it's yours anyway).

And let's build all packages and ignore the executables (if any, anyway we are not going to use them on CI):

```plain text
      - name: Go Generate
        run: go generate ./... && git diff --exit-code

      - name: Go Build
        run: go build -o /dev/null ./...

```

### Test

Okay, this part might sound overengineered but give me a moment to explain. Most of my projects are quite simple and don't require a database or any separate service to run, it's just a Go package that I use in other projects.

Doing just `go test` is exactly what I'm doing (ignore the flags and `if: ...`, I will explain them in a second):

```plain text
      - name: Go Test
        if: ${{ !inputs.skipTests }}
        run: go test -v -count=1 -race -shuffle=on -coverprofile=coverage.txt ./...

```

But as for Redis client that I'm writing ([cristalhq/redis](https://github.com/cristalhq/redis)), I need to setup Redis to test everything properly which requires additional changes to the workflow file. And I don't want that, I want to reuse as much as possible!

What I did: I added a boolean parameter to the GitHub Actions workflow named `skipTests` which is `false` by default (because most of the projects are simple).

But for the Redis client repository, I change it to `true` and disable tests. I'm running them separately in another workflow file called [e2e.yml](https://github.com/cristalhq/redis/blob/c20a6f2df4e854a5a1ffa0c167e66f0bee675df8/.github/workflows/e2e.yml).

And when tests are disabled I'm only compiling them:

```plain text
      - name: Go Compile Tests
        if: ${{ inputs.skipTests }}
        run: go test -exec /bin/true ./...

```

It's a hacky way to compile all packages and delete executables, see [Go #15513](https://github.com/golang/go/issues/15513).

To add a parameter to the workflow just do:

```plain text
on:
  workflow_call:
    inputs:
      skipTests:
        description: 'Skip tests, useful when there is a dedicated CI job for tests'
        default: false
        required: false
        type: boolean

```

### Test flags

Let's take a closer look at the `go test` flags:

- `v`
- I prefer verbose output where I can see everything
- `count=1`
- prevents caching `go test` results, which might be irrelevant for CI but I prefer to keep it
- `race`
- enable race detector, this increases build and run time but it helps to find concurrency bugs earlier, priceless.
- `shuffle=on`
- run tests in a different order to prevent implicit test order dependencies
- shameless plug: it was my [proposal](https://github.com/golang/go/issues/28592)
- `coverprofile=coverage.txt`
- collect the coverage to upload it later
- `./...`
- just test all the packages

### Benchmark

Benchmarking on the CI in many cases is a bad idea due to noisy neighbours (read other processes on the machine) that will skew your results.

But a rare and bad thing: your benchmark might be broken and you might miss that. Well, this happened to me a few times.

To fix that we can run benchmarks once just to verify that they're still passing:

```plain text
      - name: Go Benchmark
        run: go test -v -shuffle=on -run=- -bench=. -benchtime=1x ./...

```

- `v` & `shuffle=on`
- you know that already
- `run=-`
- means do not run tests
- `bench=.`
- means run every benchmark
- `benchtime=1x`
- tells to run every benchmark once
- `./...`
- do this in every package

I took this suggestion from a good Go issue "Go CI best practices" [Go #42119](https://github.com/golang/go/issues/42119). Again, thanks to Daniel Marti.

### Coverage

In the previous step, we collected code coverage, better to keep it somewhere. I prefer [Codecov](https://about.codecov.io/) but you might use any other:

```plain text
      - name: Upload Coverage
        if: ${{ !inputs.skipTests }}  # upload when we really run our tests
        uses: codecov/codecov-action@v3
        continue-on-error: true  # we don't care if it fails
        with:
          token: ${{secrets.CODECOV_TOKEN}}  # set in repository settings
          file: ./coverage.txt  # file from the previous step
          fail_ci_if_error: false

```

And that's all! This is the build and/or test workflow that I use daily in 100+ Go repositories.

## Vulncheck

Go team released `govulncheck` tool with Go 1.19. It's a simple and very fast tool to check "Does your code contain vulnerable lines?". Check the announcement for more info [go.dev/blog/vuln](https://go.dev/blog/vuln).

And see `vuln.yml` workflow in [cristalhq/.github](https://github.com/cristalhq/.github/blob/53335fadafd4a2409b342643789aea60465dbf40/.github/workflows/vuln.yml).

Thanks to [brandur.org/fragments/govulncheck-ci](https://brandur.org/fragments/govulncheck-ci) where I found this originally.

## How it can be reused?

Quite easily. GitHub Actions allow us to reuse workflows. To use the described above workflow in any of my projects I need a few lines in `.github/workflows/build.yml`:

```plain text
jobs:
  build:
    uses: cristalhq/.github/.github/workflows/build.yml@v0.4.0

  vuln:
    uses: cristalhq/.github/.github/workflows/vuln.yml@v0.4.0

```

(where [cristalhq/.github](https://github.com/cristalhq/.github) is a repository for common stuff across the whole [cristalhq](https://github.com/cristalhq) organisation, like security policies, code of conduct, etc)

One of the projects that use this workflow [cristalhq/jsn](https://github.com/cristalhq/jsn/blob/ab26cf0cee3c27ea8bb30707ee4f923e4f735031/.github/workflows/build.yml). See [Reusing workflows ](https://docs.github.com/en/actions/using-workflows/reusing-workflows) documentation for more.

## Conclusions

Here is the full workflow: [build.yml](https://github.com/cristalhq/.github/blob/53335fadafd4a2409b342643789aea60465dbf40/.github/workflows/build.yml).

That's the basic CI for most of my Go projects. Many things can be added like fuzzing, automatic releases, linting, etc.

But in this post, I wanted to share the core component of my Go CI workflow.

Thanks.

## Links

[DEV](https://dev.to/olegkovalov/github-actions-and-go-4mag) / [Lobsters](https://lobste.rs/s/9fbnlv/github_actions_go) / [Medium](https://medium.com/@olegkovalov/github-actions-and-go-639a9841b715) / [Reddit](https://www.reddit.com/r/golang/comments/11xaw79/github_actions_and_go/) / [Twitter](https://twitter.com/oleg_kovalov/status/1638094023872204800)

### Subscribe to my newsletter

Read articles from  directly inside your inbox. Subscribe to the newsletter, and don't miss out.
