---
title: "TIL: 8 versions of UUID and when to use them"
notion_id: 13bb043e-6ac3-4f0f-a075-8f7f50791228
notion_url: https://app.notion.com/p/TIL-8-versions-of-UUID-and-when-to-use-them-13bb043e6ac34f0fa0758f7f50791228
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://www.ntietz.com/blog/til-uses-for-the-different-uuid-versions/
tags: ["English", "System Design / Software Architecture", "Programming", "Article", "Guide", "technically a blog (ntietz)"]
---
About a month ago[1](https://www.ntietz.com/blog/til-uses-for-the-different-uuid-versions/?utm_source=tldrnewsletter#not-today), I was onboarding a friend into one of my side project codebases and she asked me why I was using a particular type of UUID. I'd heard about this type while working on that project, and it's really neat. So instead of hogging that knowledge for just us, here it is: some good uses for different versions of UUID.

# What are the different versions?

Usually when we have multiple numbered versions, the higher numbers are newer and presumed to be better. In contrast, there are 8 UUID versions (v1 through v8) which are different and all defined in [the standard](https://datatracker.ietf.org/doc/html/rfc9562).

Here, I'll provide some explanation of what they are at a high level, linking to the specific section of the RFC in case you want more details.

- [UUID Version 1 (v1)](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-1) is generated from timestamp, monotonic counter, and a MAC address.
- [UUID Version 2 (v2)](https://datatracker.ietf.org/doc/html/rfc9562#name-uuid-version-2) is reserved for security IDs with no known details.

[2](https://www.ntietz.com/blog/til-uses-for-the-different-uuid-versions/?utm_source=tldrnewsletter#dce)

- [UUID Version 3 (v3)](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-3) is generated from MD5 hashes of some data you provide. The RFC suggests DNS and URLs among the candidates for data.
- [UUID Version 4 (v4)](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-4) is generated from entirely random data. This is probably what most people think of and run into with UUIDs.
- [UUID Version 5 (v5)](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-5) is generated from SHA1 hahes of some data you proivde. As with v3, the RFC suggests DNS or URLs as candidates.
- [UUID Version 6 (v6)](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-6) is generated from timestamp, monotonic counter, and a MAC address. These are the same data as Version 1, but they change the order so that sorting them will sort by creation time.
- [UUID Version 7 (v7)](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-7) is generated from a timestamp and random data.
- [UUID Version 8 (v8)](https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-8) is entirely custom (besides the required version/variant fields that all versions contain).

# When should you use them?

With eight different versions, which should you use? There are a few common use cases that dictate which you should use, and some have been replaced by others.

You'll usually be picking between two of them: v4 or v7. There are also some occasions to pick v5 or v8.

- Use v4 when you just want a random ID. _This is a good default choice._
- Use v7 if you're using the ID in a context where you want to be able to sort. For example, consider using v7 if you are using UUIDs as database keys.
- v5 or v8 are used if you have your own data you want in the UUID, but generally, you will know if you need it.

What about the other ones?

- [Per the RFC](https://www.rfc-editor.org/rfc/rfc9562.html#section-5.7-4), v7 improves on v1 and v6 and should be used over those if possible. So you usually won't want v1 or v6. If you do want one of those, you _can_ use v6.
- v2 is reserved for unspecified security things. If you _are_ using these, you probably can't tell me or anyone else about it, and you're probably not reading this post to figure out more about them.
- v3 is superceded by v5, which uses a stronger hash. This one is one where you probably _know_ if you need it.

1

Despite the title of "today I learned," I did learn this over a month ago. In between, that month contained a lot of sickness and low energy, and I'm finally getting back into a cadence of having energy for some extra writing or extra coding.

[↩](https://www.ntietz.com/blog/til-uses-for-the-different-uuid-versions/?utm_source=tldrnewsletter#not-today_ref)

2

These were used in [a project](https://en.wikipedia.org/wiki/Distributed_Computing_Environment) that either failed or is extremely secretive. I can't find much information about it and the official page's copyright notice was last updated in 2020.
