---
title: "vercel-labs/scriptc: TypeScript-to-Native Compiler"
notion_id: 3ea54f1c-7d23-81d6-a297-f80e55180991
notion_url: https://app.notion.com/p/vercel-labs-scriptc-TypeScript-to-Native-Compiler-3ea54f1c7d2381d6a297f80e55180991
last_edited: 2026-09-29T04:10:00.000Z
source_url: https://github.com/vercel-labs/scriptc
tags: ["Tool", "GitHub", "English", "Programming", "TypeScript", "WebAssembly", "Automation", "Developer Tools"]
---
![image](https://camo.githubusercontent.com/1292ed5fc232d2004f06ae411b281e2cbc9dabff7a6a222082eee753fb2528b4/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4c4142532d4558504552494d454e542d3061306130612e7376673f7374796c653d666f722d7468652d6261646765266c6f676f3d56657263656c266c6162656c436f6c6f723d303030303030)

![image](https://camo.githubusercontent.com/7d3c4d47fe3535ef1c6fc498c34401f65fefaa4235f86288aae09cad71a9f833/68747470733a2f2f696d672e736869656c64732e696f2f6e706d2f762f736372697074632e7376673f7374796c653d666f722d7468652d6261646765266c6162656c436f6c6f723d303030303030)

![image](https://camo.githubusercontent.com/dd558b49cebaf6a58b3057027ed7e667e5d461d615978302b44ab39d859277dd/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6c6963656e73652f76657263656c2d6c6162732f736372697074632e7376673f7374796c653d666f722d7468652d6261646765266c6162656c436f6c6f723d303030303030)

![image](https://camo.githubusercontent.com/3b3ce038b638df35a9fbcad3f9d86c121498d3128240e5f99d590cfbff0f6659/68747470733a2f2f696d672e736869656c64732e696f2f6e706d2f646d2f736372697074632e7376673f7374796c653d666f722d7468652d6261646765266c6162656c436f6c6f723d303030303030266c6162656c3d6e706d253230646f776e6c6f616473)

scriptc compiles TypeScript and JavaScript to typed IR, readable C, textual LLVM IR, native assembly and objects, native executables, and WebAssembly modules. It uses the TypeScript compiler for parsing and type checking. Source outputs require only Node. On macOS 15+ arm64, ordinary LLVM-tier executables use scriptc's bundled helper and precompiled runtime pack; clang is only the platform linker driver and does not compile program or runtime C.

Static builds include a small native runtime, but no Node or JavaScript engine. Code that cannot compile statically is reported as a diagnostic. For npm packages and `any`-typed code, `--dynamic` embeds [quickjs-ng](https://github.com/quickjs-ng/quickjs) explicitly.

scriptc is experimental and targets macOS, Linux, Windows, and WebAssembly via WASI Preview 1.

## Installation

The compiler requires Node.js 24 or newer. `--emit=ir|c|llvm` needs only Node. `--emit=asm|obj` uses the matching optional platform helper installed with scriptc on supported macOS, Linux, and Windows hosts (and for WASI), but needs no compiler, archiver, linker, or SDK. Ordinary LLVM-tier executable builds need a platform linker driver and SDK/sysroot, but use the bundled helper plus precompiled runtime pack rather than compiling generated or runtime C. Set `SCRIPTC_LINKER` to choose that driver. Explicit C builds, LLVM fallbacks, `--sanitize`, and the deprecated `SCRIPTC_CC=clang|zigcc` compatibility route additionally need a C compiler. The executables it produces do not require Node.

## Build a program

Create `hello.ts`:

```plain text
const who = process.argv.length > 2 ? process.argv[2] : "world";
console.log(`hello, ${who}`);
```

Compile and run it in one step:

```plain text
$ scriptc run hello.ts
hello, world
```

Or write a standalone executable:

```plain text
$ scriptc build hello.ts -o hello
$ ./hello ctate
hello, ctate
```

Or stop at a source-level compiler artifact without invoking clang, an archiver, or a linker:

```plain text
$ scriptc build hello.ts --emit=ir >/dev/null
$ ls .scriptc/
hello.ir.json
$ scriptc build hello.ts --emit=c >/dev/null
$ ls .scriptc/
hello.c
hello.ir.json
$ scriptc build hello.ts --emit=llvm >/dev/null
$ ls .scriptc/
hello.c
hello.ir.json
hello.ll
$ scriptc build hello.ts --emit=asm >/dev/null
$ ls .scriptc/
hello.c
hello.ir.json
hello.ll
hello.s
$ scriptc build hello.ts --emit=obj >/dev/null
$ ls .scriptc/
hello.c
hello.ir.json
hello.ll
hello.o
hello.s
```

Different output kinds accumulate in `.scriptc/`; rebuilding a kind updates its file.

- `-emit=obj` writes a relocatable program object, not a standalone library. It has undefined `scr_*` runtime references and a required `scr_runtime_abi_v3` marker; `scriptc build --lib --profile ...` remains the self-contained archive interface. The helper runs on macOS 15+ arm64 and emits artifacts with an `arm64-apple-macosx14.0.0` deployment target. Sanitized assembly/object emission is rejected until the helper's AddressSanitizer pipeline matches the executable path.

External object consumption is experimental. Use
`--print=native-link-info` to emit the object and print a versioned JSON recipe
containing its target, `main` entry, exact `@scriptc/runtime` source pack,
required system libraries, FFI inputs, and ABI marker. The recipe never uses
hidden scriptc cache paths. See [`examples/native-object`](https://github.com/vercel-labs/scriptc/blob/main/examples/native-object)
for C-driver and direct Apple-linker builds.

## Use Node APIs

Supported Node APIs compile to the native runtime. For example, `server.ts`:

```plain text
import { createServer } from "node:http";
const server = createServer((req, res) => {
  res.setHeader("content-type", "application/json");
  res.end(JSON.stringify({ path: req.url }));
});
server.listen(8080, () => {
  console.log("listening on http://localhost:8080");
});
```

```plain text
$ scriptc build server.ts -o server
$ ./server
listening on http://localhost:8080
```

## Check static coverage

`scriptc coverage` shows how much of a program can compile statically and gives a coded diagnostic for every dynamic or unsupported site.

```plain text
$ scriptc coverage hello.ts
  statements analyzed   2
  compile statically    2  (100%)
  fully static — this program has no dynamic remainder.
```

## Build WebAssembly

WASI and other cross-target builds require Zig. Its bundled WASI libc produces a portable WASI Preview 1 module through the production LLVM backend:

Install Zig and make sure the `zig` executable is available on your `PATH`. `SCRIPTC_CC=zigcc` is scriptc's selector for invoking Zig's `cc` subcommand; `zigcc` is not a standalone executable.

```plain text
$ SCRIPTC_CC=zigcc SCRIPTC_TARGET=wasm32-wasi scriptc build hello.ts --no-keep-c -o hello.wasm >/dev/null
$ file hello.wasm
hello.wasm: WebAssembly (wasm) binary module version 0x1 (MVP)
$ SCRIPTC_CC=zigcc SCRIPTC_TARGET=wasm32-wasi scriptc run hello.ts
hello, world
```

The WASI target supports the same executable language tiers as the native targets, including async/await, promises, generators, timers, stdin/readline events, callback and promise filesystem APIs, and `--dynamic`. APIs that require capabilities absent from portable WASI Preview 1—network sockets/fetch, child processes, OS signals, and filesystem watching—fail before linking with `SC3002`; sanitizer builds, native FFI, and library-mode archive builds are target diagnostics too. See [platform support](https://scriptc.dev/platforms) for the precise boundary.

## Use npm packages

Pass `--dynamic` to embed an npm package's JavaScript in the executable. The result does not read `node_modules` at runtime.

```plain text
import pc from "picocolors";
console.log(pc.green("hello from scriptc"));
```

```plain text
$ npm install picocolors
$ scriptc build cli.ts --dynamic -o cli
$ ./cli
hello from scriptc
```

## Documentation

See the [quickstart](https://scriptc.dev/quickstart) and [CLI reference](https://scriptc.dev/cli) for the complete workflow. The docs also describe [npm dependencies](https://scriptc.dev/dependencies), [native FFI](https://scriptc.dev/ffi), [platform support](https://scriptc.dev/platforms), and the current [limitations](https://scriptc.dev/limitations).

## Development

```plain text
$ pnpm install && pnpm -r build
$ vercel link && vercel env pull  # writes a project-scoped VERCEL_OIDC_TOKEN
$ pnpm test:sandbox
```

The normal workspace build needs no local LLVM installation. To rebuild a
native helper/runtime pack, install CMake, Ninja, and the pinned LLVM 22
development package on that target host, then run the matching
`@scriptc/llvm-<platform>` and `@scriptc/runtime-<platform>` `build:native`
scripts. The macOS full test suite also uses those generated artifacts.

`pnpm test:sandbox` loads `.env.local`, preflights Vercel authentication and
project access, and uses the managed `vercel/sandbox/universal` image by
default. It installs the repository-pinned Node, pnpm, and LLVM toolchain plus
ScriptC dependencies in each disposable Sandbox before building the uploaded
worktree. Set
`SCRIPTC_SANDBOX_IMAGE` to a fully qualified VCR reference only to use the
optional prebuilt image from `pnpm test:sandbox:image`.
The prebuilt image keeps the roughly four-minute fast path; cold managed-image
runs take longer because they install the pinned toolchain in each Sandbox.

`VERCEL_OIDC_TOKEN` is preferred. For access-token authentication, set
`VERCEL_TOKEN`, `VERCEL_TEAM_ID`, and `VERCEL_PROJECT_ID`; team and project are
never inferred from `SCRIPTC_SANDBOX_IMAGE`. The legacy VCR command used by
`pnpm test:sandbox:image` cannot authenticate with an OIDC JWT, so image builds
use `VERCEL_TOKEN` when available or the existing Vercel CLI login; OIDC claims
still select the VCR team and project. The test corpus runs each program
under Node and as a compiled native binary, then compares stdout, stderr, and
exit codes byte for byte. The full gate also runs the corpus with
AddressSanitizer and the runtime reference-count audit.
