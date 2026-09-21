---
title: "Boring tech is mature, not old"
notion_id: 19954f1c-7d23-8161-8783-df73f596c33e
notion_url: https://app.notion.com/p/Boring-tech-is-mature-not-old-19954f1c7d2381618783df73f596c33e
last_edited: 2025-02-26T21:06:00.000Z
source_url: https://rubenerd.com/boring-tech-is-mature-not-old/
tags: ["English", "System Design / Software Architecture", "Programming", "Article", "Rubenerd (Ruben Schade)"]
---
## [Boring tech is mature, not old](https://rubenerd.com/boring-tech-is-mature-not-old/)

_2025-02-10T18:11:00.000-03:00_

I’ve talked before about how I think NetBSD is “boring”, and that it’s among the highest forms of praise I can give tech as a sysadmin and architect. But I’ve never elaborated why that is.

The opposite of being bored is to be surprised, and that’s not something a sysadmin desires when building, maintaining, scaling, troubleshooting, upgrading, or even replacing a complex system… especially when you’re woken up at 03:30 by a monitoring server. That’s also why the phrase _work doing what you love and you never work a day again in your life_ is an empty platitude of nonsense. But I digress.

Boring tech behaves in predictable ways. It’s a well trodden path others have evaluated, optimised, troubleshooted, and understood. Using tech that has been subjected to all those people hours of use means you’re less likely to run into edge cases, unexpected behaviour, or attributes and features that lack documentation or community knowledge. In other words, when something goes wrong, can you turn to someone or something?

Likewise, tech (generally) doesn’t exist in a vacuum. It interacts with other components and systems, some of which are even conveniently under our control, sometimes. Multiply out the potential for surprises by the number of components and their relative maturity, and your head can start to spin.

This isn’t to say there isn’t room for innovation, or that staying put is a guaranteed recipe for success. What it does teach is that it pays to make informed decisions, and that often times the understood, reliable, boring tech will get you there over something new, shiny or propped up with marketing spin. The number of people I’ve talked with who’ve replaced complicated K8s clusters with a few VMs and seen massive improvements in reliability, cost, and uptime would make some people at the Orange Peanut Gallery more than a little perturbed, for example.

There have been talks about this, perhaps [most famously by Dan McKinley](https://boringtechnology.club/). But there’s been some pushback to this idea, which is intriguing. [Robert Roskam](https://mastodon.social/@raiderrobert/113977311777631282):

> I used to agree with this. Now I don’t think so any more. You should prefer “boring” tech, and boring should be read as has been around for a while and therefore is well understood.Ubiquity is a bad test for well-understood technology. Age as a test for ubiquity is also bad.

There are a few assumptions here:

- Boring should be read as something being around for a while. I wouldn’t necessarily agree with that. I’ve been a DBA, and I never in a million years would call Oracle “boring”. It’s fiendishly complicated and difficult to maintain, and commands above-average salaries in part for that reason. A sign saying “there be Ellison dragons” isn’t boring, it’s frankly terrifying.
- Ubiquity is a test for “boringness”. He’s right here; while age gives something an opportunity to become ubiquitous, it’s not a guarantee. I’d look to the BSDs here; I’d consider them boring, but they’re not exactly widely deployed compared to Penguins.
- Boring should be read as well-understood. That’s true. Time gives something more of a chance of being well-understood, until suddenly it doesn’t and nobody is around with sufficient knowledge and inclination to maintain your COBOL stack. At that stage, I’d say you have a decidedly _un-boring_ issue on your hands.

I’d conclude by suggesting boring tech isn’t old, but _mature_. Maturity not just in the software, but its documentation, community, and track record. _Age_ is often used as an analogue for maturity, but it’s not the same thing. Otherwise I’d be more mature than my Zoomer friends, and I very much doubt that to be the case. `#BIRDISTHEWORD`.

This is why I don’t hesitate to call NetBSD boring, and why I say that’s a compliment.

**Published: **[**Tuesday 11 February 2025**](https://rubenerd.com/year/2025/)**Tagged: **[**software**](https://rubenerd.com/software/)**, **[**bsd**](https://rubenerd.com/tag/bsd/)**, **[**design**](https://rubenerd.com/tag/design/)**, **[**netbsd**](https://rubenerd.com/tag/netbsd/)**Written in: **[**Sydney**](https://rubenerd.com/location/sydney/)

- Published: [Tuesday 11 February 2025](https://rubenerd.com/year/2025/)
- Tagged: [software](https://rubenerd.com/software/), [bsd](https://rubenerd.com/tag/bsd/), [design](https://rubenerd.com/tag/design/), [netbsd](https://rubenerd.com/tag/netbsd/)
- Written in: [Sydney](https://rubenerd.com/location/sydney/)
