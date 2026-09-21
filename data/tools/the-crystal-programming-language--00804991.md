---
title: "The Crystal Programming Language"
notion_id: 00804991-db2a-419c-bc96-e6e0d2bb825e
notion_url: https://app.notion.com/p/The-Crystal-Programming-Language-00804991db2a419cbc96e6e0d2bb825e
last_edited: 2023-06-02T01:04:00.000Z
source_url: https://crystal-lang.org/
tags: ["English", "Programming", "Untried", "Tool"]
---
## Latest release [1.8.2](https://crystal-lang.org/2023/05/09/crystal-1.8.2-released/)

## Syntax

Crystal’s syntax is heavily inspired by Ruby’s, so it feels natural to read and easy to write, and has the added benefit of a lower learning curve for experienced Ruby devs.

[Start learning Crystal with the Language Reference](https://crystal-lang.org/reference/getting_started/)

## Type system

Crystal is statically type checked, so any type errors will be caught early by the compiler rather than fail on runtime. Moreover, and to keep the language clean, Crystal has built-in type inference, so most type annotations are unneeded.

[Read more about Crystal's type system](https://crystal-lang.org/reference/syntax_and_semantics/types_and_methods.html)

## Null reference checks

All types are non-nilable in Crystal, and nilable variables are represented as a union between the type and nil. As a consequence, the compiler will automatically check for null references in compile time, helping prevent the dreadful [billion-dollar mistake](https://www.infoq.com/presentations/Null-References-The-Billion-Dollar-Mistake-Tony-Hoare).

Running the previous file:

## Macros

Crystal’s answer to metaprogramming is a powerful macro system, which ranges from basic templating and AST inspection, to types inspection and running arbitrary external programs.

[Read more about macros](https://crystal-lang.org/reference/syntax_and_semantics/macros.html)

## Concurrency Model

Crystal uses green threads, called fibers, to achieve concurrency. Fibers communicate with each other using channels, as in Go or Clojure, without having to turn to shared memory or locks.

[Read more about Crystal's concurrency model](https://crystal-lang.org/reference/guides/concurrency.html)

## C-bindings

Crystal has a dedicated syntax to easily call native libraries, eliminating the need to reimplement low-level tasks.

[Learn how to bind to C libraries](https://crystal-lang.org/reference/syntax_and_semantics/c_bindings/)

## Dependencies

Crystal libraries are packed as Shards, and distributed via Git without needing a centralised repository. Built in commands allow dependencies to be easily specified through a YAML file and fetched from their respective repositories.

[Read more about Shards in the repo](https://github.com/crystal-lang/shards)

### Crystal top sponsors

[Meet all](https://crystal-lang.org/sponsors)

[Manas Technology Solutions](https://manas.tech/)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[Nikola Motor Company ](https://nikolamotor.com/) [  PlaceOS](https://place.technology/)

### Some of our CI runs here

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Articles from our blog

Nicolás di Tada 15 May 2023

### [Changes in the Steering Committee](https://crystal-lang.org/2023/05/15/update-in-steering-committee/)

[Juan Wajnerman is stepping down from his position, with Beta Ziliani taking his place.](https://crystal-lang.org/2023/05/15/update-in-steering-committee/)

_chat_bubble_outline_ 0 Comments

23 Mar 2023

### [LLVM opaque pointer support has landed](https://crystal-lang.org/2023/03/23/llvm-opaque-pointers/)

[Updates to the LLVM bindings bring support for LLVM 15+ and significant improvements in codegen performance in the next release.](https://crystal-lang.org/2023/03/23/llvm-opaque-pointers/)

_chat_bubble_outline_ 2 Comments

06 Mar 2023

### [Reveal type in Crystal](https://crystal-lang.org/2023/03/06/reveal-type-in-crystal/)

[Porting reveal_type from Sorbet to Crystal.](https://crystal-lang.org/2023/03/06/reveal-type-in-crystal/)

_chat_bubble_outline_ 1 Comment

[More articles](https://crystal-lang.org/blog)

## Release Notes

|  | [Crystal 1.8.2 is released!](https://crystal-lang.org/2023/05/09/crystal-1.8.2-released/) | 09 May 2023 |
| --- | --- | --- |
|  | [Crystal 1.8.1 is released!](https://crystal-lang.org/2023/04/20/1.8.1-released/) | 20 Apr 2023 |
|  | [Crystal 1.8.0 is released!](https://crystal-lang.org/2023/04/14/1.8.0-released/) | 14 Apr 2023 |
|  | [Crystal 1.7.3 is released!](https://crystal-lang.org/2023/03/07/1.7.3-released/) | 07 Mar 2023 |
|  | [Crystal 1.7.2 is released!](https://crystal-lang.org/2023/01/23/1.7.2-released/) | 23 Jan 2023 |
|  | [Crystal 1.7.1 is released!](https://crystal-lang.org/2023/01/17/1.7.1-released/) | 17 Jan 2023 |
|  | [Crystal 1.7.0 is released!](https://crystal-lang.org/2023/01/09/1.7.0-released/) | 09 Jan 2023 |
|  | [Crystal 1.6.2 is released!](https://crystal-lang.org/2022/11/03/1.6.2-released/) | 03 Nov 2022 |
|  | [Crystal 1.6.1 is released!](https://crystal-lang.org/2022/10/21/1.6.1-released/) | 21 Oct 2022 |
|  | [Crystal 1.6.0 is released!](https://crystal-lang.org/2022/10/06/1.6.0-released/) | 06 Oct 2022 |
