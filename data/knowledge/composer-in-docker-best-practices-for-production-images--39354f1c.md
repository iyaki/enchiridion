---
title: "Composer in Docker: best practices for production images"
notion_id: 39354f1c-7d23-8198-a53b-e2a8c6751a2f
notion_url: https://app.notion.com/p/Composer-in-Docker-best-practices-for-production-images-39354f1c7d238198a53be2a8c6751a2f
last_edited: 2026-07-04T02:39:00.000Z
source_url: https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images
tags: ["English", "Docker", "PHP", "Continuous Integration/Continuous Delivery", "DevOps", "Guide", "nth-root.nl"]
---
Published on June 29th, 2026

- PHP
- Composer
- Docker

![image](https://nth-root.nl/resized/nic-wortel-jGB5X-d/80x80.jpg)

By

**Nic Wortel**

Learn how to use Composer to build production-ready Docker images for PHP applications, including best practices for caching dependencies and optimizing image size.

In this guide

- [Getting the Composer binary into the build](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#getting-the-composer-binary-into-the-build)
- [Running the Composer install command](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#running-the-composer-install-command)
- [Excluding the local vendor directory with .dockerignore](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#excluding-the-local-vendor-directory-with-dockerignore)
- [Cache optimizations for faster consecutive builds](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#cache-optimizations-for-faster-consecutive-builds)
- [Optimizing the autoloader](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#optimizing-the-autoloader)
- [Installing private packages](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#installing-private-packages)
- [Putting it together](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#putting-it-together)
- [Wrapping up](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#wrapping-up)

**Updated on 30-06-2026: added the **[**Pinning the Composer image using a digest**](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#pinning-the-composer-image-using-a-digest)** section.**

Docker containers are a reliable way to deploy your application.
As the image contains not only your PHP code but also all the dependencies your application needs to run,
you can be sure that it will run the same way in production as it does on your local machine.

Docker should be able to build your application image from your source code, without relying on any manual steps.
This means that instead of copying a pre-installed `vendor` directory into the image,
your Dockerfile should run `composer install` as part of the build process.
This ensures an image which can be built in any environment (both locally and in CI/CD pipelines).

Running Composer during your Docker build comes with some challenges, and there are different approaches which come with
pros and cons. This guide demonstrates some best practices when it comes to building production-ready Docker images.

## Getting the Composer binary into the build

Composer is a _build-time dependency_ for your application image.
This means that you need to have the Composer binary available in the build environment, but it does not need to be
present in the final image.
In fact, it is a best practice to ensure that the Composer binary is not present in the final image.

Your first guess might be to install the Composer binary using the installation script from the Composer website, similar
to how you would install it on your local machine.
But while the installation script is a convenient way for manual installation, it is not suitable for automated
installation as it verifies the downloaded file against a hash which changes with every release.

The recommended way of using Composer in Docker is by using
the [official Composer image](https://hub.docker.com/_/composer).
However, you should not base your production image on the Composer image: it is not meant to be used as a production image
and your application might require a different PHP version or extensions than the ones provided by the Composer image.
Instead, the official Composer image's documentation [recommends copying the Composer binary from the Composer image into
your own image](https://hub.docker.com/_/composer#php-version--extensions):

```plain text
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer
```

While this provides a more reliable way to get the Composer binary into your build environment than the installer script,
it still has the drawback that the Composer binary becomes part of your final image.
Even if you remove it in a later step, it will still be present in the image history and increase the size of your image.

### Using a multi-stage build

A [multi-stage build](https://docs.docker.com/build/building/multi-stage/) is a Docker feature
that lets you split your Dockerfile into multiple stages and copy artifacts from intermediate stages into the final image.
This allows you to install Composer in an intermediate stage, run `composer install`, and then copy only the `vendor`
directory into the final image:

```plain text
FROM composer:2 AS build-stage
WORKDIR /app
COPY composer.json composer.lock ./
RUN composer install --no-dev
FROM php:8.5-fpm
WORKDIR /app
COPY --from=build-stage /app/vendor/ vendor/
COPY . .
```

The above example has one drawback: if your application depends on a specific PHP version or extensions (listed in
`composer.json`), the `composer install` command will fail because the Composer image's runtime does not match the
platform requirements.
While this issue can be suppressed by passing the `--ignore-platform-reqs` flag to Composer, doing so would silently
ignore any missing platform requirements, which can lead to runtime errors in production.

A more reliable approach is to base the build stage on the same PHP image as your final image, and (again)
copy the Composer binary into it from the official Composer image:

```plain text
FROM php:8.5-fpm AS base
RUN docker-php-ext-install pdo_mysql
WORKDIR /app
FROM base AS build-stage
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer
COPY composer.json composer.lock ./
RUN composer install --no-dev
FROM base
COPY --from=build-stage /app/vendor/ vendor/
COPY . .
```

### Bind-mounting the binary

An alternative approach to the multi-stage build is a single-stage build where you bind-mount the Composer binary.
Modern versions of Docker can bind-mount files from another image (or the build context) into a `RUN` command.
This lets you mount the Composer binary straight from the official image for the duration of the install,
without copying it into your own image.

```plain text
RUN --mount=type=bind,from=composer:2,source=/usr/bin/composer,target=/usr/bin/composer \
    composer install --no-dev --no-interaction
```

The mount only exists for the duration of the `RUN` command, so the Composer binary will not be present in the final image.
This is the best approach if you want to avoid complicated multi-stage builds,
but it requires a recent version of Docker with BuildKit enabled.

### Installing Composer's own runtime requirements

Composer has a few runtime requirements of its own, such as `openssl` to download packages over HTTPS and `zip` to
extract them.
The official Composer image includes these dependencies, but if you are copying or mounting the Composer binary into
your own image you need to ensure that your image (or build stage) has them as well.

When you base your image on the official `php` image, most of these dependencies are already present, but the `zip`
extension and the `unzip` binary are not, which means that you will need to install them yourself:

```plain text
FROM base AS build-stage
RUN apt-get update && apt-get install -y --no-install-recommends libzip-dev unzip \
    && docker-php-ext-install zip \
    && rm -rf /var/lib/apt/lists/*
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer
```

This shows the main benefit of using a multi-stage build over bind-mounting the executable: you can install Composer's
dependencies in the build stage without adding them to your final image.

In some cases Composer has a dependency on `git` to download packages from Git repositories.
Composer 2.10
[disabled the automatic fallback to source checkouts when the dist/zip install fails](https://getcomposer.org/changelog/2.10.0),
but some applications load (private) dependencies straight from a Git repository.
If that is the case, add `git` to the `apt-get install` line in your build stage as well.

### Pinning the Composer image using a digest

Image tags (such as `composer:2` or `composer:2.10.1`) are mutable, meaning that they can be overwritten to point to a
newer image.
While that sounds useful, because you will automatically get the latest version of Composer when building your image, it
also means that any breaking changes or bugs in a newer release of Composer can suddenly break your build.
On top of that, a malicious actor who gains push access to the image registry could overwrite an existing tag to point
to a malicious version of the image as part of a supply-chain attack, allowing them to compromise your application.

Similar to how Composer's lock file locks your dependencies to fixed versions, you
can [pin images in your Dockerfile to an immutable digest](https://docs.docker.com/reference/cli/docker/image/pull/#pull-an-image-by-digest-immutable-identifier)
too:

```plain text
COPY --from=composer:2.10.1@sha256:7725eb4545c438629ae8bde3ef0bb9a5038ef566126ad878442a69007242d267 /usr/bin/composer /usr/bin/composer
```

This pins Composer to the exact image matching that digest, even when the tag is updated to point to a newer
image.
The tag could even be omitted (`composer@sha256:<hash>` resolves to the same image), but keeping it documents the
intended version in a human-readable way.

The digest for any image can be found in the image registry UI, or by pulling it:

```plain text
$ docker pull composer:2.10.1
2.10.1: Pulling from library/composer
Digest: sha256:7725eb4545c438629ae8bde3ef0bb9a5038ef566126ad878442a69007242d267
Status: Image is up to date for composer:2.10.1
docker.io/library/composer:2.10.1
```

While pinning the image does not guarantee its trustworthiness (pinning a compromised image will just force Docker to
always pull that compromised image), it does mean that the image cannot suddenly change between builds.
Any update of the image becomes a deliberate change, tracked in version control and validated by your CI pipeline, just
like running `composer update`.
The same principle can be applied to the `php` base image as well.

The trade-off is that a pinned image never gets updated by itself, requiring you to manually update the digest to stay
up-to-date.
Tools like [Dependabot](https://docs.github.com/en/code-security/dependabot) and [Renovate](https://docs.renovatebot.com/)
can take care of this by tracking the image registry and opening a pull request to bump the tag and the digest when an
image is updated.
That allows you to run your CI pipeline against the newer image version, and merge the PR once the build passes.

## Running the Composer install command

Now that the Composer binary is available in the build environment, we can add a `RUN composer install` instruction to
run Composer.

The default behavior of Composer is optimized for development environments, and to build a production-ready image we
should provide some extra options:

- `-no-dev`, to skip the installation of development dependencies (such as unit testing and static analysis libraries)
- `-no-interaction`, to prevent Composer from asking interactive questions which would hang the build
- `-no-progress`, to prevent Composer from rendering progress bars
- `-no-scripts`, to prevent Composer from executing scripts, see [Preventing unnecessary Composer reinstalls](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#preventing-unnecessary-composer-reinstalls)

Combining those, a typical `composer install` instruction looks like this:

```plain text
RUN composer install --no-dev --no-interaction --no-progress --no-scripts
```

## Excluding the local `vendor` directory with `.dockerignore`

Every `COPY . .` instruction copies your entire project into the image, including the local `vendor` directory.
That overwrites the dependencies Composer installed during the build, which may target different platform requirements
or include dev dependencies.

A `.dockerignore` file tells Docker which paths to leave out of the build context, much like `.gitignore` does for Git.
Excluding `vendor` keeps your built dependencies intact, and keeps the build context small as a bonus:

```plain text
# Include anything else that shouldn't end up in the image, such as .git/, .env.local, or node_modules/
vendor/
```

## Cache optimizations for faster consecutive builds

Downloading and extracting dependencies can take a long time, especially for large applications with many dependencies.
This can add up quickly when you are rebuilding your image frequently during development or in a CI/CD pipeline.
While the first build will always take the longest, you can optimize subsequent builds by caching the result of
`composer install` and reusing it until something changes.

### Preventing unnecessary Composer reinstalls

Docker caches the result of each instruction and reuses the cached layer until something it depends on changes.
This makes subsequent builds more efficient, but only when you are careful about the order of instructions in your
`Dockerfile`.

A `COPY` instruction's cache is invalidated when any copied file changes, which makes the following `Dockerfile` quite
inefficient:

```plain text
COPY . .
RUN composer install --no-dev --no-interaction --no-progress --no-scripts
```

Editing a single PHP file invalidates the `COPY` layer, which invalidates the `composer install` layer below it, forcing
a full reinstall of all dependencies.

To avoid this, copy only the files Composer needs first, run `composer install`, then copy the rest of the files:

```plain text
COPY composer.json composer.lock ./
RUN composer install --no-dev --no-interaction --no-progress --no-scripts
COPY . .
```

Now the `composer install` layer is only invalidated when `composer.json` or `composer.lock` changes, which is exactly
what you want.

The `--no-scripts` flag is important here: many post-install scripts (Symfony's `auto-scripts`, for example) expect the
full application to be present, but at this point only `composer.json` and `composer.lock` have been copied and running
them would result in an error.
Instead, run `composer run-script post-install-cmd` after the source files have been copied, or run the relevant scripts
as individual `RUN` instructions.

### Sharing the Composer cache between builds

The optimized layer ordering prevents unnecessary `composer install` runs, but doesn't help when we're actually updating
one of the dependencies.
A change in `composer.json` or `composer.lock` still requires a full download of all dependencies,
because the Composer cache is part of the (now invalid) layer and is not persisted between builds.

[Cache mounts](https://docs.docker.com/build/cache/optimize/#use-cache-mounts) are a BuildKit feature that allows you to
persist the Composer cache between builds without including it in the image:

```plain text
RUN --mount=type=cache,target=/tmp/composer-cache \
    COMPOSER_CACHE_DIR=/tmp/composer-cache \
    composer install --no-dev --no-interaction --no-progress --no-scripts
```

[`COMPOSER_CACHE_DIR`](https://getcomposer.org/doc/03-cli.md#composer-cache-dir) is an environment variable which points
Composer to the mounted cache directory, and the cache mount keeps that directory between builds.
After the first build, changing a dependency forces Composer to perform a fresh install, but it will only have to
download the new dependency.

## Optimizing the autoloader

Composer generates an autoloader which can load classes and interfaces when you use them in your code.
This enables you to use those classes and interfaces without littering your codebase with `require_once` expressions.
By default, the Composer autoloader checks the filesystem before resolving a classname.
This is convenient for development environments because it will discover new classes without having to regenerate the
autoloader, but adds unnecessary overhead in a production environment where autoloaded files shouldn't change anyway.

As part of the Docker image build, we can optimize the Composer autoloader by generating a class map which maps all
known class names to file paths, reducing the number of filesystem checks.
This can be achieved by passing the `--optimize-autoloader` flag to the `composer install` command or the `--optimize`
flag to the separate `dump-autoload` command:

```plain text
RUN composer dump-autoload --no-dev --optimize
```

The `dump-autoload` command is especially useful when
the [source files are not yet copied into the image](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#preventing-unnecessary-composer-reinstalls)
when `composer install` is executed.

Authoritative class maps are an enhancement of the standard class maps, by disabling all filesystem checks and only
relying on the class map.
This makes autoloading faster, but can break autoloading for classes which are not included in the class map
(because they were generated at runtime, for example).

```plain text
RUN composer dump-autoload --no-dev --classmap-authoritative
```

Test this carefully before enabling it in production, and revert to `--optimize` if something breaks.

## Installing private packages

Many applications depend on (internal or third-party) private packages, hosted in a private Git or Packagist repository.
These packages are not publicly accessible and require some form of authentication to download.
While it is tempting to simply copy the credentials into the image or set them as an environment variable, this is a
security risk: anyone who can pull the image can see the credentials in the image history, even if you delete them in a
later layer.

Even when using a multi-stage build and copying only the `vendor` directory into the final image, the credentials are
still present in the intermediate build stage and exported build cache.

Instead, a [build secret mount](https://docs.docker.com/build/building/secrets/) should be used to securely pass the
credentials during the build:

```plain text
RUN --mount=type=secret,id=composer_auth,env=COMPOSER_AUTH \
    composer install --no-dev --no-interaction --no-progress --no-scripts
```

The `env=COMPOSER_AUTH` option exposes the secret as an environment variable for the duration of the `composer install`
command, which is the environment variable Composer reads credentials from.

To make this work, the secret needs to be passed as an argument of the `docker build` command:

```plain text
COMPOSER_AUTH='{"http-basic":{"repo.packagist.com":{"username":"token","password":"..."}}}' \
    docker build --secret id=composer_auth,env=COMPOSER_AUTH .
```

This reads the value of the `COMPOSER_AUTH` environment variable and passes it as a secret with ID `composer_auth` to
the build.
Alternatively, it's possible to source the secret value from a file:

```plain text
docker build --secret id=composer_auth,src=auth.json .
```

Because the mount declares `env=COMPOSER_AUTH`, the file's contents are exposed through that environment variable,
so `auth.json` must contain a valid `COMPOSER_AUTH` value (the same JSON structure Composer's own `auth.json` uses).
Just make sure to add the file to `.gitignore` so you don't accidentally commit it.

## Putting it together

A complete single-stage Dockerfile combining the binary mount, manifests-first ordering,
the cache mount, and the recommended flags:

```plain text
FROM php:8.5-fpm
RUN docker-php-ext-install pdo_mysql
RUN apt-get update && apt-get install -y --no-install-recommends libzip-dev unzip \
    && docker-php-ext-install zip \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY composer.json composer.lock ./
RUN --mount=type=bind,from=composer:2.10.1@sha256:7725eb4545c438629ae8bde3ef0bb9a5038ef566126ad878442a69007242d267,source=/usr/bin/composer,target=/usr/bin/composer \
    --mount=type=cache,target=/tmp/composer-cache \
    COMPOSER_CACHE_DIR=/tmp/composer-cache \
    composer install --no-dev --no-interaction --no-progress --no-scripts
COPY . .
RUN --mount=type=bind,from=composer:2.10.1@sha256:7725eb4545c438629ae8bde3ef0bb9a5038ef566126ad878442a69007242d267,source=/usr/bin/composer,target=/usr/bin/composer \
    composer dump-autoload --no-dev --classmap-authoritative
```

This example installs `zip` and `unzip` directly in the final image (see
[Installing Composer's own runtime requirements](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#installing-composers-own-runtime-requirements)).
Bind-mounting the binary can't isolate those runtime dependencies the way a build stage can,
so they remain in the final image.

The multi-stage equivalent moves the install into a build stage and copies the result forward,
so the runtime image never contains Composer or the download cache:

```plain text
FROM php:8.5-fpm AS base
RUN docker-php-ext-install pdo_mysql
WORKDIR /app
FROM base AS build-stage
RUN apt-get update && apt-get install -y --no-install-recommends libzip-dev unzip \
    && docker-php-ext-install zip \
    && rm -rf /var/lib/apt/lists/*
COPY --from=composer:2.10.1@sha256:7725eb4545c438629ae8bde3ef0bb9a5038ef566126ad878442a69007242d267 /usr/bin/composer /usr/bin/composer
COPY composer.json composer.lock ./
RUN --mount=type=cache,target=/tmp/composer-cache \
    COMPOSER_CACHE_DIR=/tmp/composer-cache \
    composer install --no-dev --no-interaction --no-progress --no-scripts
COPY . .
RUN composer dump-autoload --no-dev --classmap-authoritative
FROM base
COPY --from=build-stage /app/vendor/ vendor/
COPY . .
```

The final `COPY . .` in each example relies on a `.dockerignore` that excludes the local `vendor/`
(see [Excluding the local ](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#excluding-the-local-vendor-directory-with-dockerignore)[`vendor`](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#excluding-the-local-vendor-directory-with-dockerignore)[ directory with ](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#excluding-the-local-vendor-directory-with-dockerignore)[`.dockerignore`](https://nth-root.nl/en/guides/composer-in-docker-best-practices-for-production-images#excluding-the-local-vendor-directory-with-dockerignore)):
without it, that copy would overwrite the `vendor/` built in the image with your local one.

Both examples skip Composer scripts with `--no-scripts`.
If your application defines post-install scripts (such as Symfony's `auto-scripts`), run them with a separate
`RUN composer run-script post-install-cmd` after `COPY . .`, once the full source is present.

## Wrapping up

A short checklist for installing dependencies in a Docker build:

- Keep Composer out of the runtime image, by bind-mounting the binary or using a Composer stage, and pin the image by digest.
- Copy `composer.json` and `composer.lock` before the source
so dependency installs are cached separately from code changes.
- Add a `.dockerignore` that excludes `vendor/` so `COPY . .` can't overwrite the dependencies built in the image.
- Add a cache mount for Composer's download cache to speed up rebuilds.
- Pass auth tokens as build secrets, never as `ENV` or a committed `auth.json`.
- Use `-no-dev`, `-no-interaction`, `-no-progress`, and `-no-scripts`,
and optimize the autoloader after the source is in place.

![image](https://nth-root.nl/resized/nic-wortel-jGB5X-d/200x200.jpg)

Are you looking to implement PHP, Composer, or Docker in your project or team?
                    As a PHP consultant I can save you time and money by guiding you through the
                    process and helping you avoid costly mistakes.
