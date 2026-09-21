---
title: "Keep Stuff Linkable"
notion_id: 8a4b9c3e-026a-45b6-b81f-c32ba4aea1c1
notion_url: https://app.notion.com/p/Keep-Stuff-Linkable-8a4b9c3e026a45b6b81fc32ba4aea1c1
last_edited: 2023-04-22T19:55:00.000Z
source_url: https://animaomnium.github.io/keep-stuff-linkable/
tags: ["Blogging/Content Creation", "Reflection", "Article", "Crash Lime", "English"]
---
_We interrupt this doomscroll __
to bring you..._

# [Crash Lime](https://animaomnium.github.io/)

Delicious performance 
64 bites at a time!

You’ve spent hours absorbing [incoming](https://news.ycombinator.com/) [bitstreams](https://lobste.rs/), and a seed of an idea has germinated in your mind. You fire up your [favorite text editor](https://neovim.io/), plant it down, and spend a couple hours letting the idea grow out. You’ve finished! You are about to publish your freshly-grown bitstream on the interwebs when you get _that sinking feeling_ in your gut: something’s missing…

You scan over the post. Is anything wrong? Nope: argument is solid, formatting is A-Ok: [oll korrect](https://en.wikipedia.org/wiki/OK). But wait… what’s that?

Where are all the links?

No [links](https://en.wikipedia.org/wiki/Hyperlink), no game. (It is the _Web_ you are publishing to, after all.) Sighing, you stumble around with [Google](https://duckduckgo.com/) for a bit before giving up. Maybe some other time. If only there were a better way…

Finding links sucks.

It may be my lack of discipline, but keeping track of the references principally responsible for each little bit of text I write is _hard_. I have so many linkless posts waiting to be published: I could go ahead and publish them as-is, but by doing so I feel as though I’d be treating you, dear reader, unjustly.

I wish there was a faster way to link the posts I write. I want to [write about things that are happening ](https://simonwillison.net/2023/Apr/16/web-llm/)[_now_](https://simonwillison.net/2023/Apr/16/web-llm/). If I wait a week to hunt down references, things will have already moved on. [Writing consistently](https://blog.codinghorror.com/how-to-achieve-ultimate-blog-success-in-one-easy-step/) requires rhythm, and nothing interrupts a consistent rhythm more effectively than haphazardly tumbling down [internet rabbit-holes](https://en.wikipedia.org/wiki/Rabbit_Hole) while in search of the perfect link.

In a perfect world, I imagine a little companion reading everything alongside me. He records the references and key ideas of each piece. After I write a post, he’d comb through my post sentence by sentence, linking every important phrase to its source. Now, I haven’t quite done this, but I’ve developed a quick-and-dirty [first-order approximation](https://en.wikipedia.org/wiki/Order_of_approximation), which should hopefully let me link stuff with greater ease. Let’s dive in.

_Linkoln_, no pun intended, is the name of my companion. It’s a short hacky [Python](https://www.python.org/) script I wrote this morning, so I could link and publish this post. All links in this post are Linkoln’s fault, not mine.

Linkoln parses [wikilinks](https://en.wikipedia.org/wiki/Help:Link) out of a [markdown](https://commonmark.org/help/) document, and searches the [world wide web](https://en.wikipedia.org/wiki/World_Wide_Web) to find a [hyperlink](https://en.wikipedia.org/wiki/Hyperlink) for each one.

Here’s what Linkoln does, on a more concrete level. Given a Markdown post with wikilinks:

```plain text
# Thoughts on Rust
[[programming language:Rust]] is a [[systems programming language]] bootstrapped from [[rust prehistory|OCaml]].

```

Linkoln normalizes the post, replacing each wikilink with the best corresponding hyperlink it could find on the web, using [Google](https://duckduckgo.com/):

```plain text
# Thoughts on Rust
[Rust][1] is a [systems programming language][2] bootstrapped from [OCaml][3].
[1]: https://www.rust-lang.org
[2]: https://en.wikipedia.org/wiki/System_programming_language
[3]: https://github.com/graydon/rust-prehistory

```

Linkoln supports three types of wikilinks:

1. Literal Links: `[[text]]`
2. Query Links: `[[query|text]]`
3. Context Links: `[[context:text]]`

Here’s a quick breakdown of each link type:

_Literal Links_ search the given query and include the query verbatim. For example:

```plain text
[[GitHub]]

```

Searches “GitHub” and becomes:

```plain text
[GitHub][0]
[0]: https://github.com

```

_Query Links_ let you use a different query than the text of the link. For example:

```plain text
[[notes on a smaller rust|Rust *could* be easier]]

```

Searches “notes on a smaller rust” and becomes:

```plain text
[Rust *could* be easier][0]
[0]: https://boats.gitlab.io/blog/posts/notes-on-a-smaller-rust

```

_Context Links_ are useful when qualifying a search for an otherwise generic term. The two halves are [concatenated](https://en.wikipedia.org/wiki/Concatenation) to form the entire query:

```plain text
[[wikipedia language:Python]]

```

Searches “wikipedia language Python” and becomes:

```plain text
[Python][0]
[0]: https://en.wikipedia.org/wiki/Python_(programming_language)

```

Despite its simplicity, I’ve been getting a lot of mileage out of Linkoln. I can keep all my link-searching activity in one place: Once I’ve finished a post, I can bracket off terms to link, qualifying searches as necessary with advanced [google-foo](https://www.urbandictionary.com/define.php?term=googlefoo), and Linkoln takes care of the rest.

Linkoln is by no means a silver bullet: it’s not intended to be one. The point of this post is not to highlight some gimmicky throwaway python script, but to express, perhaps, a little worry over why I think linking stuff is so important.

Despite Linkoln’s reliance on them, the need for web-wide search engines could be said to be a failure in the organizational structure of the web. As the proliferation of [GPT-4](https://openai.com/research/gpt-4) leads to the crystallization of the [Dead Internet](https://www.theatlantic.com/technology/archive/2021/08/dead-internet-theory-wrong-but-feels-true/619937/), how will we find a single live page in a soup of procedurally generated web-gloop?

Perhaps links _are_ dead. Why link when ChatGPT can explain? Why post and upvote when [attention-maximizing algorithms](https://gantry.io/blog/papers-to-know-20230110) can recommend? Perhaps we’re at the end of the old-web, now a corner relegated to hobbyists, as all text ever written is absorbed in a single differentiable scream.

For us hobbyists, however, perhaps links _aren’t_ dead: they’re vitally important. Links lend authority. Trace a hop away from your homepage, maybe two: can you still trust what you read? Does it matter?

So in this deluge, link _more_, not less. Don’t link to stuff you don’t trust, SEO or otherwise. If using a tool like Linkoln—script, chatbot, or otherwise—make sure you vet what you’re linking to. You are not _just_ linking, but building a Web of Trust: _this_ is _key_.

Links aren’t dead. Neither is the old web. Keep the dream of an open web alive: Keep Stuff Linkable!

Happy linking!

[_Lobsters_](https://lobste.rs/stories/new?url=https%3A%2F%2Fanimaomnium.github.io%2Fkeep-stuff-linkable%2F)_ • _[_HN_](https://news.ycombinator.com/submitlink?u=https%3A%2F%2Fanimaomnium.github.io%2Fkeep-stuff-linkable%2F&t=Keep%20Stuff%20Linkable)_ • _[_Reddit_](https://www.reddit.com/submit?url=https%3A%2F%2Fanimaomnium.github.io%2Fkeep-stuff-linkable%2F)_ • _[_Twitter_](http://twitter.com/share?url=https%3A%2F%2Fanimaomnium.github.io%2Fkeep-stuff-linkable%2F&text=Keep%20Stuff%20Linkable)

_Thank you for tuning in!_

Check out more posts by heading [Home](https://animaomnium.github.io/).

Subscribe to this blog via [Atom/RSS](https://animaomnium.github.io/atom.xml).

Copyright © 20xx Insert Employer Here, All Limes Preserved.
