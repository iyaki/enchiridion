---
title: "Janet for Mortals - Janet Lang"
notion_id: 40a76428-eaa0-4d69-9f59-5f6633ee9c53
notion_url: https://app.notion.com/p/Janet-for-Mortals-Janet-Lang-40a76428eaa04d699f595f6633ee9c53
last_edited: 2023-07-05T19:36:00.000Z
source_url: https://janet.guide/
tags: ["Book", "English", "Janet Lang"]
---
Hello and thank you for coming to my book.

I want to tell you about [Janet](https://janet-lang.org/), because I think that Janet is a very good language and it’s a shame that you haven’t heard of it yet. I like Janet so much that I wrote an entire book about it. Look:

Wow! Look at all those chapters. Just like a real book.

This isn’t one of those chapters yet, though. This is the before part. If this were a real book, which it is, this part might be labeled “Introduction.”

Oh, good.

Janet occupies a really valuable niche: it is a small, simple language that is actually _usable_. It has an elegant simplicity that you might associate with someone’s hobby project, but you can actually run it on Windows. It has built-in concurrency and multithreading, and it is an _excellent_ language for text-wrangling, thanks to the native support for parsing expression grammars (think better regular expressions). It has a simple C FFI, so the package ecosystem is “all of them,” as long as you’re willing to write a few lines of binding code first. And thanks to the lightweight runtime, it’s very easy to use Janet as an embedded scripting language.

My favorite feature of Janet, though, is something that sounds really dumb when I say it out loud: you can actually _distribute_ Janet programs to other people. You can compile Janet programs into statically-linked native binaries and give them to people who have never even heard of Janet before. And they can run them without having to install any gems or set up any virtual environments or download any runtimes or anything else like that.

This means that if you want to use Janet for something, _you can_, and no one even needs to know.

And to be clear, I’m not going to try to convince you to bet your next startup on Janet, or even to use it in any sort of production setting. But I think it’s an excellent language for exploratory programming, scripting, and fun side projects. I’ve personally used Janet in a few capacities:

- [dumb little texty scripts a little too complicated for shell](https://github.com/ianthehenry/privy)
- [dumb little games for weekend tinkering](https://ianthehenry.com/posts/janet-game/day-one/)
- [dumb interactive art playgrounds](https://bauble.studio/)

I have yet to do anything smart with it, though.

## about this book

This book assumes that you already know how to program; I’m not going to waste your time explaining what a `for` loop is. Specifically this book assumes that you already know how to program _in JavaScript_, because it’s better to have something concrete to diff against, and I’m willing to bet that even if it isn’t your first language, you know enough JavaScript to be able to follow along. And if you don’t, then, well, you’ll probably still be fine. All programming languages are basically the same.

With that in mind I’m going to emphasize what makes Janet _different_ early on. I’ll cover the whole language, but I’ll talk about macros and images and PEGs before I talk about, like, `if` statements. Is that a good way for a book to present information? I don’t know. We’re going to find out together.

## about this website

This book is a real book, as previously established. But it is also undeniably a _website_, at this particular moment. As such it has the full power of cyberspace at its disposal, and there are some features you should be aware of before we get started.

The first is that this book contains a repl, and you can summon it whenever you’d like by pressing the escape key on your keyboard. The book will then start downloading like a megabyte of JavaScript and WebAssembly, and once it’s done you will be able to try out Janet right here in the comfort of your browser. No need to install anything; no need to leave the comfort of this book website if you’d like to test something out.

The repl is not _just_ a repl, though. It is also a portal into conversation with me, the author. You can use the repl to report typos or factual errors, ask questions, or express confusion. I won’t be able to _respond_ in the repl, but if you include some kind of contact information in your reports I will make an effort to follow up with you. Here, why don’t you try it now? Open up the repl and type something like this:

```plain text
(report "hey nice book")

```

Fun, right?

## about this author

Oh that’s me.

I’m just a fan of Janet; I am not affiliated with the language in any way. I have no real qualifications to be writing a book about it, and nothing that you read here should be considered authoritative, idiomatic, or educational.
