---
title: "Your Docker Image Is 1.2GB. Here Is How To Get It Under 80MB. | The Practical Developer"
notion_id: 36d54f1c-7d23-818f-8560-c331bd45675a
notion_url: https://app.notion.com/p/Your-Docker-Image-Is-1-2GB-Here-Is-How-To-Get-It-Under-80MB-The-Practical-Developer-36d54f1c7d23818f8560c331bd45675a
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://the-practical-developer.online/posts/docker-image-from-1gb-to-80mb/
tags: ["English", "Docker", "Node.js", "DevOps", "Backend", "Containerization", "Article", "Tutorial", "The Practical Developer"]
---
A step-by-step optimization of a real Node.js Docker image, from a 1.2GB monster to a 78MB production container. Each technique is benchmarked, copy-paste ready, and explained with the trade-offs.

![image](https://images.unsplash.com/photo-1605745341112-85968b19335b?w=1200&q=80)

Stacked shipping containers at a port

Bloated Docker images are the silent tax on every team that ships containers. Slow CI. Slow deploys. Bigger attack surface. Bigger registry bill. And almost always, it’s fixable in an afternoon.

I took a real Node.js + TypeScript service we ship to production, started from the naive Dockerfile most teams write, and walked it down from 1.2GB to 78MB. Same app, same behavior, six steps, all measured on the same machine. Here is exactly what moved the needle.

## The starting point: 1.2GB

This is the Dockerfile most teams begin with. It works. It is also wasteful in almost every line.

```plain text
FROM node:22

WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

EXPOSE 3000
CMD ["npm", "start"]
```

Build it and check the size:

```plain text
$ docker build -t app:naive .
$ docker images app:naive
REPOSITORY   TAG     SIZE
app          naive   1.21GB
```

1.21GB to ship a service that produces about 4MB of compiled JavaScript. Let’s fix it.

## Step 1: Switch the base image, 1.21GB to 412MB

The `node:22` tag is Debian-based and includes a full toolchain you do not need at runtime. The `slim` variant strips most of it.

| Image | Size |
| --- | --- |
| `node:22` | 1.21GB |
| `node:22-slim` | 412MB |
| `node:22-alpine` | 178MB |

Alpine is even smaller, but it uses musl libc instead of glibc. Most pure-JS apps run fine on it, but anything with native modules (`bcrypt`, `sharp`, `node-gyp` builds) needs extra care, and some packages have subtle musl bugs. I default to `slim` and reach for `alpine` only when I know the dependency tree is clean.

For this post, we will keep going with `slim` to stay realistic about real-world apps.

## Step 2: Use a `.dockerignore`, 412MB to 388MB

`COPY . .` happily copies your `node_modules`, `.git`, build artifacts, local `.env` files, IDE folders, and test fixtures into the image. Even if a later step overwrites `node_modules`, the layer is already in the image history.

Create a `.dockerignore`:

```plain text
node_modules
npm-debug.log
.git
.gitignore
.env*
.vscode
.idea
coverage
dist
build
*.md
test
__tests__
Dockerfile*
.dockerignore
```

Small win on size, big win on rebuild speed and security. Your local `.env.development` is no longer hiding inside a layer that gets pushed to a public registry.

## Step 3: Multi-stage build, 388MB to 198MB

You need TypeScript, eslint, the test framework, and probably a hundred transitive dev dependencies to build the app. You do not need any of them to run it.

A multi-stage build compiles in one image and copies only the artifacts into a second, clean image:

```plain text
# ---- builder ----
FROM node:22-slim AS builder
WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build
RUN npm prune --omit=dev

# ---- runtime ----
FROM node:22-slim
WORKDIR /app

COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./package.json

EXPOSE 3000
CMD ["node", "dist/index.js"]
```

Two things doing real work here:

`npm ci` instead of `npm install`. It uses `package-lock.json` directly, is faster, and is reproducible. Use it in CI and in Docker. Always.

`npm prune --omit=dev` strips dev dependencies after the build. On a typical TypeScript service, that is half of `node_modules` gone.

## Step 4: Layer caching that actually works, same size, 5x faster rebuilds

This is not about size, but it is the single biggest CI win. Most Dockerfiles invalidate `npm ci` on every code change because they `COPY . .` before installing. Order matters.

Already shown above, but worth calling out: copy `package*.json` first, install, then copy the rest. Now `npm ci` is cached as long as your dependencies don’t change.

On a project with ~600 dependencies, this took our cold rebuild from 94 seconds to 18 seconds when only application code changed. Multiplied by every PR, every day, every developer.

## Step 5: Switch to Alpine for the runtime stage, 198MB to 96MB

We can keep the Debian-based builder (compatible with everything) and switch only the runtime to Alpine. The compiled JS does not care about the base OS at runtime as long as no native binaries are calling glibc-specific symbols.

```plain text
# ---- builder ----
FROM node:22-slim AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build
RUN npm prune --omit=dev

# ---- runtime ----
FROM node:22-alpine
WORKDIR /app
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./package.json
EXPOSE 3000
CMD ["node", "dist/index.js"]
```

If you have native modules, build them in a stage that matches the runtime libc. For Alpine that means `node:22-alpine` as the builder too, plus `apk add --no-cache python3 make g++` to compile, then a clean runtime stage.

## Step 6: Drop Node entirely with `distroless`, 96MB to 78MB

Google’s `distroless` images contain just the Node runtime and its TLS roots. No shell, no package manager, no apt, no curl. If something pops a shell inside your container, there is no shell to pop.

```plain text
# ---- runtime ----
FROM gcr.io/distroless/nodejs22-debian12
WORKDIR /app
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./package.json
EXPOSE 3000
CMD ["dist/index.js"]
```

Note the `CMD` syntax change. There is no shell in distroless, so you cannot use the shell form (`CMD node dist/index.js`). The exec form must be used, and the entrypoint is already `node`, so you just pass the script.

Trade-off: you cannot `docker exec -it container sh` for a quick poke around. For debugging, run the same image with the `:debug` tag, which adds a busybox shell. In production, you keep the locked-down version.

## The full picture

| Step | Image | Size | Saved |
| --- | --- | --- | --- |
| 0. Naive | `node:22` | 1.21GB | - |
| 1. `slim` base | `node:22-slim` | 412MB | -67% |
| 2. `.dockerignore` | `node:22-slim` | 388MB | -6% |
| 3. Multi-stage + prune | `node:22-slim` | 198MB | -49% |
| 4. Layer caching | `node:22-slim` | 198MB | (rebuild speed) |
| 5. Alpine runtime | `node:22-alpine` | 96MB | -52% |
| 6. Distroless | `distroless/nodejs22` | 78MB | -19% |

Total: **94% reduction**. Roughly 15× smaller.

## What this actually buys you

Disk and registry costs are the obvious win, but they are usually not the biggest one.

Faster cold starts on Kubernetes and serverless platforms. Pulling 78MB instead of 1.2GB on a node that does not have the layer cached is the difference between a pod ready in 4 seconds and one that takes 40.

Smaller attack surface. The CVE count on `node:22` is in the hundreds. On distroless it is a handful. Your security scanner will stop screaming.

Faster CI. Pushing and pulling smaller images in your pipeline shaves real wall-clock time off every deploy.

Cheaper cross-region replication. If you mirror your registry across regions for DR, you are now moving 6% of the bytes you used to move.

## A few things that did not make the cut

People keep recommending these, and they are usually not worth the complexity:

`docker-slim` (the auto-shrinking tool). It works, but it can silently strip files your app loads at runtime under conditions your test suite does not exercise. Debugging that in production is not a fun afternoon.

Static binaries with `pkg` or `nexe`. They produce small images, but they freeze you to a specific Node version and break dynamic imports. If you are willing to give up flexibility for size, you should probably be writing in Go or Rust, not bundling Node.

`scratch` images. Beautiful in theory. In practice, you spend a week tracking down missing CA certs, missing `tzdata`, and missing nameservers. Distroless gives you 95% of the benefit with none of the pain.

## The takeaway

The naive Dockerfile is fine for a hackathon. For anything you ship more than once a week, the six steps above pay for themselves in a single afternoon. Most of the size comes off in the first three steps. The last three are the difference between “good enough” and “actually production-grade”.

Copy the final Dockerfile, set up your `.dockerignore`, point your CI at it, and stop paying the 1.2GB tax.

_This article was put together by The Practical Developer, with notes from the Docker pipelines we run on real client products. If you are scaling a team that ships software like this, and would rather hire the practice than build it from scratch, Yojji is the international software development studio behind a lot of the work that ends up here. Founded in 2016, with offices across Europe, the US, and the UK, Yojji builds custom web and mobile products with senior engineers, dedicated teams, and a strong bias for actually shipping. Worth a mention if you are looking for the kind of partner who treats Dockerfiles, CI pipelines, and production hygiene as part of the job, not an afterthought._
