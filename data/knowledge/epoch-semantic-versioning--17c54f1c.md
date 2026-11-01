---
title: "Epoch Semantic Versioning"
notion_id: 17c54f1c-7d23-81d4-99b5-ce4dbb630a23
notion_url: https://app.notion.com/p/Epoch-Semantic-Versioning-17c54f1c7d2381d499b5ce4dbb630a23
last_edited: 2025-02-14T20:39:00.000Z
source_url: https://antfu.me/posts/epoch-semver
tags: ["Programming", "System Design / Software Architecture", "Project Management", "Article", "Website", "Anthony Fu", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If you’ve been following my work in open source, you might have noticed that I have a tendency to stick with zero-major versioning, like `v0.x.x`. For instance, as of writing this post, the latest version of UnoCSS is [`v0.65.3`](https://github.com/unocss/unocss/releases/tag/v0.65.3), Slidev is [`v0.50.0`](https://github.com/slidevjs/slidev/releases/tag/v0.50.0), and `unplugin-vue-components` is [`v0.28.0`](https://github.com/unplugin/unplugin-vue-components/releases/tag/v0.28.0). Other projects, such as React Native is on [`v0.76.5`](https://github.com/facebook/react-native/releases/tag/v0.76.5), and sharp is on [`v0.33.5`](https://github.com/lovell/sharp/releases/tag/v0.33.5), also follow this pattern.

The reason I’ve stuck with `v0.x.x` is my own unconventional approach to versioning. I prefer to introduce necessary and minor breaking changes early on, making upgrades easier, without causing alarm that typically comes with major version jumps like `v2` to `v3`. Some changes might be "technically" breaking but don’t impact 99.9% of users in practice. Breaking changes are relative; even a bug fix can be breaking for those relying on the previous behavior (but that’s another topic for discussion :P). There’s a special rule in SemVer that states **when the leading major version is ****`0`****, every minor version bump is considered breaking**. I’ve been leveraging this rule to navigate the system more flexibly. I kinda abuse that rule to workaround the limitation of SemVer.

Of course, zero-major versioning is not the only solution to be progressive. We can see that tools like [Node.js](https://nodejs.org/en), [Vite](https://vite.dev/), [Vitest](https://vitest.dev/) are rolling out major versions in consistent intervals, with a minimal set of breaking changes in each release that are easy to adopt.

I have to admit that sticking to zero-major versioning isn’t the best practice. While I aimed for more granular versioning to improve communication, using zero-major versioning has actually limited my ability to convey changes effectively. In reality, I’ve been wasting a valuable part of the versioning scheme due to my peculiar insistence.

Thus here, I am proposing to change.

## Epoch Semantic Versioning

[In an ideal world, I would wish SemVer to have four numbers: ](https://x.com/antfu7/status/1679184417930059777)[`EPOCH.MAJOR.MINOR.PATCH`](https://x.com/antfu7/status/1679184417930059777). The `EPOCH` version is for those big announcements, while `MAJOR` is for technical incompatible API changes that might not be significant. This way, we can have a more granular way to communicate changes. Similar we also have [Romantic Versioning that propose ](https://github.com/romversioning/romver)[`HUMAN.MAJOR.MINOR`](https://github.com/romversioning/romver). But of course, it’s too late for the entire ecosystem to adopt a new versioning scheme.

If we can’t change SemVer, maybe we can at least extend it. I am proposing a new versioning scheme called **Epoch Semantic Versioning** (Epoch SemVer for short). Build on top of the structure of `MAJOR.MINOR.PATCH`, extend the first number to be the combination of `EPOCH` and `MAJOR`. To put a difference between them, we use a third digit to represent `EPOCH`, which gives `MAJOR` a range from 0 to 99. This way, it follows the exact same rules as SemVer **without requiring any existing tools to change, but provides more granular information to users**.

The format is simple:

`{EPOCH * 100 + MAJOR}.MINOR.PATCH`

- EPOCH: Increment when you make significant or groundbreaking changes.
- MAJOR: Increment when you make incompatible API changes.
- MINOR: Increment when you add functionality in a backwards-compatible manner.
- PATCH: Increment when you make backwards-compatible bug fixes.

For example, UnoCSS would transition from `v0.65.3` to `v65.3.0`. Following SemVer, a patch release would become `v65.3.1`, and a feature release would be `v65.4.0`. If we introduced some minor incompatible changes affecting an edge case, we could bump it to `v66.0.0` to alert users of potential impacts. In the event of a significant overhaul to the core, we could jump directly to `v100.0.0` to signal a new era and make a big announcement. This approach provides maintainers with more flexibility to communicate the scale of changes to users effectively.

Of course, I’m not suggesting that everyone should adopt this approach. It’s simply an idea to work around the existing system. It will be interesting to see how it performs in practice.

## Moving Forward

I plan to adopt Epoch Semantic Versioning in my projects, including UnoCSS, Slidev, and all the plugins I maintain. I hope this new versioning approach will help communicate changes more effectively and provide users with better context when upgrading.

I’d love to hear your thoughts and feedback on this idea. Feel free to share your comments using the links below!
