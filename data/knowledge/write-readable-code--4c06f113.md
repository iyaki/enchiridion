---
title: "Write Readable Code"
notion_id: 4c06f113-fe7f-4d74-a201-da369e3e7af6
notion_url: https://app.notion.com/p/Write-Readable-Code-4c06f113fe7f4d74a201da369e3e7af6
last_edited: 2023-04-25T13:19:00.000Z
source_url: https://thoughtbot.com/blog/write-code-to-be-read
tags: ["English", "Programming", "Article", "thoughtbot Blog"]
---
[https://thoughtbot.com/blog/write-code-to-be-read](https://thoughtbot.com/blog/write-code-to-be-read)

Code is read more than it’s written. We write it once, and then read it back. It goes through review if we’re on a team. It is read again when someone else needs to understand, add to, or modify that code. This includes ourselves weeks or months later.

Despite this, we tend to focus on the writing as the main “action”. Writing is very important, but before we write we need to understand the context. We must read before we write. It’s much easier to understand code if it’s written well. Even in write-heavy situations like a new codebase, we eventually have to come back and read our first steps. We should optimize code to be read.

## Names

Names describe what variables, methods, classes are or what they do. They outline the system we’re working with. It’s much easier to write terse names. They keep our lines short and make it easy to type those names again and again. For example cc = CreditCard.find instead of primary_card, or def set_attr instead of set_user_profile_attribute.

The problem is that non-descriptive names like cc or set_attr require further investigation to discover what they are and how they should be used. These examples favor easy writing, not reading.

Consider the concepts you learned to write this code, and try to capture that in names. Consider the why or how something is used instead of what it is. initial_sign_up_profile says a lot more than profile, and lock_stats_table_for_data_export says more than lock_db.

Readability is the goal here, not name length. You can absolutely make unreadable code with long names, especially lots of long names that are too similar. Go for readability, not some arbitrary length metric.

## Abstracting Procedural Logic

The code we write to manipulate a system is different from the way we describe that manipulation. Imagine the process of “showing a modal dialog”. That’s how we’d describe it, even to code-proficient colleagues. We don’t often describe this as “find the appropriate related DOM element and set CSS classes to be visible” but that’s the level that code thinks on. It’s our job to translate between those levels of abstraction.

When you have a long method, the classic fix is extract method. Extract method works by breaking up our unrefined code into named abstractions representing the underlying logic. Again, we’re back to naming, but with a slightly different goal. A good name allows you to describe the functionality in a way that doesn’t require the user to know every internal piece of the system. It allows them to learn (or re-learn) the deeper details as needed.

Here’s an example of showing a modal with JavaScript:

```plain text
function async showModal(event) { const target = event.target; const modal = document.querySelector( event.target.dataset.relatedModalSelector ); if (!modal || !modal.classList.contains("modal")) { return; } for (const element of document.querySelectorAll("modal")) { element.classList.add("hidden"); } const data = modal.dataset; const modalTitle = JSON.parse(data.display)["title"]; const modalContent = await fetchModalData(data.remoteUrl); modal.innerHtml = modalContent; modal.classList.remove("hidden"); }
```

If you already know how the modal system works, this is reasonable to read. But most people don’t keep that information in their heads at all times. Abstracting this procedural logic will help anyone looking at this code with fresh eyes understand where they need to make changes:

```plain text
function async showModal(event) { const modal = this.findPossibleModal(event); if (!this.isValidModal(modal)) { return; } await this.setModalContent(modal); this.hideEveryModal(); this.revealModal(modal) }
```

The refactor makes the necessary steps for displaying modals clear and easily understood. If we need, we can find specific implementation details in extracted methods, and it’s immediately clear what each method is doing. All the pieces exist on a similar level of abstraction; in this case manipulating related DOM elements. The encapsulating method showModal is an abstraction, too, that exists with abstractions on a similar level. It’s easy to imagine other nearby interactions like submitForm, syncUserProgress, or enableFocusMode.

## Testing

When testing, it’s relatively common to isolate the setup phase from the rest of the test using abstractions like let or before. Many tests in the same file require similar (or the same) pieces of context to run, so consolidating that setup feels like a natural way to DRY up a test. Grouping related code can also feel similar to abstraction.

But this makes tests harder to read. That setup code defines the state of the system. More often than not we haven’t seen these tests recently or ever. These pieces of setup are critical to understanding how to fix existing tests or add more. A test separated from its context forces us to memorize that context which distracts from our problem solving skills. A good test tells a story.

Most tests also test a system in multiple states; no single setup can speak for all scenarios. At best, shared setup will have to be redefined for individual tests, scatting that context. At worst, setup is entirely wasted as global setup goes unused. When we put shared setup at the top, we are assuming that all future tests need this particular setup. Write a few more tests and that assumption will likely prove false, causing us to reorganize the whole file or just live with the waste.

Keeping all of that setup inline makes that test much more readable. It’s staggeringly not DRY, but DRY isn’t a useful goal for tests. We do not need tests to be built on reusable abstractions and have a short line count. We need tests to give us predictable confidence in our system and help us refactor.

## Broader Goals

It’s worth remembering that specific metrics like code complexity, test coverage, and “DRY” aren’t goals by themselves. The goal is code that we can easily understand and confidently change to give users the best possible software. Although “readable” is harder to measure, having it as a guiding principle can help us know when to bend or break these quantitative rules and build better software.
