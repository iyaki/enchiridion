---
title: "Best practices for inclusive textual websites"
notion_id: dba97589-ac80-439d-82c7-dd1e9a436a2c
notion_url: https://app.notion.com/p/Best-practices-for-inclusive-textual-websites-dba97589ac80439d82c7dd1e9a436a2c
last_edited: 2023-09-13T11:48:00.000Z
source_url: https://seirdy.one/posts/2020/11/23/website-best-practices/
tags: ["English", "Web Development", "Guide", "Article", "Seirdy"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Before you begin

The following applies to minimal websites that focus primarily on text. It does not apply to websites that have a lot of non-textual content. It also does not apply to websites that focus more on generating revenue or pleasing investors than being inclusive.

This is a “living document” that I add to as I receive feedback. See the updated date and changelog after the post title.

If you find the article too long, just read the introduction and conclusion. The table of contents should help you skim.

## Introduction

[ Permalink to section ](https://seirdy.one/posts/2020/11/23/website-best-practices/#introduction)

I realize not everybody’s going to ditch the Web and switch to Gemini or Gopher today (that’ll take, like, at least a month /s). Until that happens, here’s a non-exhaustive, highly-opinionated list of best practices for websites that focus primarily on text. I don’t expect anybody to fully agree with the list; nonetheless, the article should have at least some useful information for any web content author or front-end web developer.

### Inclusive design

My primary focus is [inclusive design](https://100daysofa11y.com/2019/12/03/accommodation-versus-inclusive-design/). Specifically, I focus on supporting _underrepresented ways to read a page_. Not all users load a page in a common web-browser and navigate effortlessly with their eyes and hands. Authors often neglect people who read through accessibility tools, tiny viewports, machine translators, “reading mode” implementations, the Tor network, printouts, hostile networks, and uncommon browsers, to name a few. I list more niches in [the conclusion](https://seirdy.one/posts/2020/11/23/website-best-practices/#conclusion). Compatibility with so many niches sounds far more daunting than it really is: if you only selectively override browser defaults and use plain-old, semantic HTML (POSH), you’ve done half of the work already.

One of the core ideas behind the flavor of inclusive design I present is inclusivity by default. Web pages shouldn’t use accessible overlays, reduced-data modes, or other personalizations if these features can be available all the time. Personalization isn’t always possible: Tor users, students using school computers, and people with restrictive corporate policies can’t “make websites work for them”; that’s a webmaster’s responsibility.

At the same time, many users do apply personalizations; sites should respect those personalizations whenever possible. Balancing these two needs is difficult. Some features conflict; you can’t display a light and dark color scheme simultaneously. Personalization is a fallback strategy to resolve conflicting needs. Disproportionately underrepresented needs deserve disproportionately greater attention, so they come before personal preferences instead of being relegated to a separate lane.

### Restricted enhancement

Another focus is minimalism. [Progressive enhancement](https://en.wikipedia.org/wiki/Progressive_enhancement) is a simple, safe idea that tries to incorporate some responsibility into the design process without rocking the boat too much. I don’t find it radical enough. I call my alternative approach “restricted enhancement”.

Restricted enhancement limits all enhancements to those that solve specific accessibility, security, performance, or significant usability problems faced by people besides the author. These enhancements must be made progressively when possible, with a preference for using older or more widespread features, taking into account unorthodox user agents. Purely-cosmetic changes should be kept to a minimum.

I’d like to re-iterate yet another time that this only applies to websites that primarily focus on text. If graphics, interactivity, etc. are an important part of your website, less of the article applies. My hope is for readers to consider a subset of this page the next time they build a website, and _address the trade-offs they make when they deviate._ I don’t expect—or want—anybody to follow all of my advice, because doing so would make the Web quite a boring place!

Our goal: make a textual website maximally inclusive, using restricted enhancement.

### Prior art

You can regard this article as an elaboration on existing work by the Web Accessibility Initiative (WAI).

I’ll cite the WAI’s [Techniques for WCAG 2.2](https://www.w3.org/WAI/WCAG22/Techniques/) a number of times. Each “Success Criterion” (requirement) of the WCAG has possible techniques. Unlike the Web Content Accessibility Guidelines (WCAG), the Techniques document does not list requirements; rather, it serves to non-exhaustively educate authors about _how_ to use specific technologies to comply with the WCAG. I don’t find much utility in the technology-agnostic goals enumerated by the WCAG without the accompanying technology-specific techniques to meet those goals.

I’ll also cite [Making Content Usable for People with Cognitive and Learning Disabilities](https://www.w3.org/TR/coga-usable/), by the WAI. The document lists eight objectives. Each objective has associated personas, and can be met by several design patterns.

### Why this article exists

Performance and accessibility guidelines are scattered across multiple WAI documents and blog posts. Moreover, guidelines tend to be overly general and avoid giving specific advice. Guidelines from different places tend to contradict each other, especially when they have different goals (e.g., security and accessibility). They also tend to be focused on large corporate sites rather than the simple text-oriented content the Web was made for.

I wanted to create a single reference with non-contradictory guidelines, containing advice more specific and opinionated than existing material. I also wanted to approach the very different aspects of site design from the same perspective and in the same place, allowing readers to draw connections between them.

### xkcd comic: infinite scrolling

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Infinite-scroll means that accidental navigation to a link results in losing your place. From

Toggle comic transcript

### Comic transcript

Megan stands at a desk, touching a book gingerly. Cueball stands behind her.

Cueball Why are you turning the pages like that? Megan If I touch the wrong thing, I’ll lose my place and have to start over. Caption below the panel If books worked like infinite-scrolling webpages

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Searching for the word “good” before and after a “see more” link is clicked. Both situations show a match, but only one of them allows us to view the match. Both screenshots are from the Reddit redesign.

### Example unreadable palette

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This is an unreadable screenshot of [Better Motherfucking Website](http://bettermotherfuckingwebsite.com/). I had set my browser foreground and background colors to white and dark gray, respectively. The website overrode the foreground colors while assuming that everyone browses with a white background.

Toggle screenshot transcript

### Screenshot transcript

A little less contrast.

Black on white? How often do you see that kind of contrast in real life? Tone it down a bit, asshole. I would’ve even made this site’s background a nice `#EEEEEE` if I wasn’t so focused on keeping declarations to a lean 7 fucking lines.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

WebKit’s default dark stylesheet uses dark-colored links that are difficult to read.

This image is an approximation of what halation looks like, cropped from [Essential Accessibility](https://www.essentialaccessibility.com/blog/accessibility-for-people-with-astigmatism).     

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

It’s impossible to discern the number of links in a sequence without some sort of separator. Whitespace alone isn’t sufficient.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Screenshot of the GitHub issues for waifu2x-ncnn-vulkan.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Two screenshots of [the aforementioned ](https://web.archive.org/web/20220808163715/https://www.smashingmagazine.com/2022/06/voice-control-usability-considerations-partially-visually-hidden-link-names/)[_Smashing Magazine_](https://web.archive.org/web/20220808163715/https://www.smashingmagazine.com/2022/06/voice-control-usability-considerations-partially-visually-hidden-link-names/)[ article](https://web.archive.org/web/20220808163715/https://www.smashingmagazine.com/2022/06/voice-control-usability-considerations-partially-visually-hidden-link-names/) after I selected the title text, before and after pressing Tab. The focus moves _backwards_ to the start of the container because the container is focusable. The focus should have moved to an element after the selected text.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

When filtering criteria on [the Quickref Reference page](https://www.w3.org/WAI/WCAG22/quickref/?currentsidebar=%23col_customize&showtechniques=134%2C124&levels=a&technologies=js%2Cserver%2Csmil%2Cpdf%2Cflash%2Csl), a dickbar lists active filters. I increased the zoom level; you may have to add more filters to fill the screen with a smaller font.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

I made sure to leave enough non-interactive space in [my homepage’s webring list](https://seirdy.one/#webrings) to accommodate a 48 px tap target, with extra space in between.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Firefox 99’s default focus indicator, before and after my adjustments.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This page in the [SerenityOS](https://serenityos.org/) web browser. TLS 1.2 support isn’t finished yet; I loaded it from a mirror with a compatible cipher suite.

## Acknowledgements and further reading

[ Permalink to section ](https://seirdy.one/posts/2020/11/23/website-best-practices/#further-reading)

Initial versions of this page were inspired by existing advocates for web minimalism.

Parts of this page can be thought of as an extension to David Copeland’s principles of Brutalist Web Design.

Raw content true to its construction:
1. Content is readable on all reasonable screens and devices.
2. Only hyperlinks and buttons respond to clicks.
3. Hyperlinks are underlined and buttons look like buttons.
4. The back button works as expected.
5. View content by scrolling.
6. Decoration when needed and no unrelated content.
7. Performance is a feature. —[David Copeland](https://naildrivin5.com/), [Brutalist Web Design](https://brutalist-web.design/)

The [250kb club](https://250kb.club/) gathers websites at or under 250kb, and also rewards websites that have a high ratio of content size to total size.

The [10KB Club](https://10kbclub.com/) does the same with a 10kb homepage budget (excluding favicons and webmanifest icons). It also has guidelines for noteworthiness, to avoid low-hanging fruit like mostly-blank pages.

My favorite website club has to be the [XHTML Club](https://xhtml.club/) by [Bradley Taunt](https://bt.ht/), the creator of the original [1mb.club](https://1mb.club/).

Also see [Motherfucking Website](https://motherfuckingwebsite.com/). Motherfucking Website inspired several unofficial sequels that tried to gently improve upon it. My favorite is [Best Motherfucking Website](https://bestmotherfucking.website/).

The [Web Bloat Score calculator](https://www.webbloatscore.com/) is a JavaScript app that compares a page’s size with the size of a PNG screenshot of the full page content, encouraging site owners to minimize the ratio of the two.

One resource I found useful (that eventually featured this article!) was the “Your page content” section of [Your Personal Website](https://www.billdietrich.me/YourPersonalWebSite.html) by [Bill Dietrich](https://www.billdietrich.me/).

If you’ve got some time on your hands, I _highly_ recommend reading the [Web Content Accessibility Guidelines (WCAG) 2.2](https://www.w3.org/TR/WCAG22/). The WCAG 2 standard is technology-neutral, so it doesn’t contain Web-specific advice. For that, check the [How to Meet WCAG (Quick Reference)](https://www.w3.org/WAI/WCAG22/quickref/). It combines the WCAG with its supplementary [list of techniques](https://www.w3.org/WAI/WCAG22/Techniques/).

The WCAG are an excellent starting point for learning about accessibility, but make for a poor stopping point. Much of the content on this page simply isn’t covered by the WCAG. One of my favorite resources for learning about what the WCAG _doesn’t_ cover is [Axess Lab’s articles](https://axesslab.com/articles/).

I’ve learned about a great number of underrepresented ways to browse from the Fediverse, particularly from [this subthread asking people to share](https://pleroma.envs.net/notice/AHqp3TEDFoyz0W4nbc) (requires JavaScript; [plaintext mirror](https://gopher.envs.net/pleroma.envs.net:7070/1/notices/AHqp3TEDFoyz0W4nbc)). Several responses informed updates to this page.

An early version of this article received useful responses when I [posted it to Lobsters](https://lobste.rs/s/akcw1m/opinionated_list_best_practices_for); I incorporated some feedback shortly afterward.

A special thanks goes out to GothAlice for the questions she answered in #webdev on Libera.Chat.
