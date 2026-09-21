---
title: "Build Your Own Redis"
notion_id: 3d615968-933e-414d-bfdc-0b5338e3cc04
notion_url: https://app.notion.com/p/Build-Your-Own-Redis-3d615968933e414dbfdc0b5338e3cc04
last_edited: 2023-02-01T16:50:00.000Z
source_url: https://build-your-own.org/
tags: ["Book", "English", "Programming", "Network", "Databases", "Learning"]
---
Read it [here](https://build-your-own.org/redis/).

## Introduction

Needless to say, the Redis project is quite a success. It’s an important component in backend applications.

Redis could be considered one of the _building blocks_ of modern computing. There are not many projects that fit the such role and stood the test of time. Here are some examples that meet my criteria of the “building block”: NGINX, SQLite, PostgreSQL, Kafka, Linux kernel, etc.

Most of us are not working on projects of such a level, but it is still worthwhile to learn from those projects. It takes higher skill and deeper knowledge to build such projects, thus learning from those projects could be a path to the next level as a software developer. The book is the result of my own learning.

## Why a Book?

My motivations for writing this book:

- Thanks to the open-source culture, anyone can learn from those projects freely by reading the source code. One of my motivations for writing the book is to share what I have learned in the past, and my book is freely available on the web as a giving back. (You can still purchase ebooks or paperbacks.)
- Like most real-world projects, Redis is a complex project built with lots of effort, which can be hard to grasp for beginners. Instead, my book takes an opposite approach: learning by building things _from scratch_; there are not many “from scratch” books.

## Why From Scratch?

A couple of points:

- **To learn faster.** By building things from scratch, concepts can be introduced gradually. Starting from the small, adding things incrementally, and getting the big picture in the end.
- **To learn deeper.** While there are many materials explaining how an existing stuff works, the understanding obtained by reading those materials is often not the same as building the stuff yourself. It is easy to mistake memorization for understanding, and it’s easier to pick up unimportant details than principles and basics.
- **To learn more.** The “from scratch” approach forces you to touch _every_ aspect of the subject — there are no shortcuts to knowledge! And often not every aspect is known to you beforehand, you may discover “things I don’t know I don’t know” in the process.

Summarized in a quote from Feynman: “What I cannot create, I do not understand”.

## Why Redis?

What makes Redis a good target is that it covers two important subjects of software engineering: network programming and data structures.

While there are many guides and books on network programming, I find that there are still some gaps to be filled:

- Some books focus on the usage of socket APIs, and while mastering socket APIs is necessary, it is not sufficient. Network programming (and any other computing subjects) is not about using APIs, knowledge of the event loop, protocols, timers, etc is also needed to do anything useful in production. The socket API is only a tiny aspect of learning network programming, and this book will cover the rest.
- Other books introduce network programming by high-level libraries or frameworks, which have downsides as stated in the “[Why From Scratch](https://build-your-own.org/blog/20230127_byor/#why-from-scratch)” section.

Data structures are interesting topics, and what’s interesting in Redis is that Redis only employs a few basic data structures, yet it is quite versatile and covers quite a lot of real-world use cases.

Although many people learned some basic data structures from textbooks, there is still something more to learn. Data structures implemented in real projects often have some practical considerations which are not touched by textbooks. Learning how data structures are used in a non-toy environment is a unique experience from building Redis.

## The Book

I have split the contents into short chapters, each chapter builds on the previous one, adding a new concept. The full source code is provided for reference purposes, readers are advised to tinker with it or DIY without it.

The code is written as direct and straightforwardly as I could. It’s mostly plain C with minimal C++ features. Don’t worry if you don’t know C, you just have the opportunity to do it in another language by yourself.

The end result is a mini Redis alike with only about 1200 lines of code. 1200 LoC seems low, but it illustrates many important aspects the book attempts to cover.

The techniques and approaches used in the book are not exactly the same as the real Redis. Some are intentionally simplified, and some are chosen to illustrate a general topic. Readers can learn even more by comparing different approaches.

The book is relatively short — the paperback is 120 pages of pure content in a compact size (6x9in), and no fluffs are included.

Read the free web version [here](https://build-your-own.org/redis/). If it’s helpful to you, consider [purchasing](https://build-your-own.org/redis/99_wip.html) the ebook or a hard copy.

Welcome to **build-your-own.org**.

A website for free educational software development materials.

for updates and new books.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

By build stuff from scratch.

With straightforward C/C++ code.
