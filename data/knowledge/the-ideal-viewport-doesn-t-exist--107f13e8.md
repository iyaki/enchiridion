---
title: "The ideal viewport doesn’t exist"
notion_id: 107f13e8-a63c-4f7b-b90b-4ba6eed3d50b
notion_url: https://app.notion.com/p/The-ideal-viewport-doesn-t-exist-107f13e8a63c4f7bb90b4ba6eed3d50b
last_edited: 2023-08-29T12:16:00.000Z
source_url: https://viewports.fyi/
tags: ["Website", "Article", "English", "Web Development", "UI/UX"]
---
Before you settle on basing design decisions on a handful of strict breakpoints, make sure you consider the vast fragmentation of screen sizes and browser viewports.

Here at [Set Studio](https://set.studio/), we conducted a little casual experiment to answer “how fragmented are viewport sizes?”. We gathered over **120,000 datapoints** with over **2,300 unique viewport sizes**. The data mainly came from users in the USA and Europe, therefore it is not necessarily representative of a global audience, but still useful for this article.

The experiment only ran for 48 hours, but the data we got was pretty interesting. Let’s dive in and take a look.

## What does 120,000 datapoints represent?

It's important to understand just how many 120,000 is in relative terms. For comparison, let’s presume each datapoint is a person.

Wembley stadium has a capacity of **90,000**, so our datapoints could fill Wembley once and still fill another third of the available capacity.

The population of our home town, Cheltenham, is around **116,000** so our datapoints could almost populate the entire town!

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## What are the top viewport sizes

[buildexcellentwebsit.es](https://buildexcellentwebsit.es/)

This isn't a problem if you build in a [fluid, flexible manner](https://buildexcellentwebsit.es/). That is illustrated in this diagram. If however, you tend to build with very specific breakpoints and hard values for typography, sizing and spacing, you might find that even with the best intentions, **you’re not providing the optimal user experience**.

Let's take an example of a “pixel perfect” UI with a fixed header and/or footer. It might look great when you shrink your development browser down, but how does it look in the conditions outlined earlier? How does it look when you visit from a tiny viewport like a smart watch? How does it look when you visit from a landscape phone?

Based on some of the combinations of aspect ratio and dimensions, we're confident those cases were represented in our data. Also, [people told us too](https://social.vasilis.nl/@vasilis/110586308036626591).

_Before you commit to fixed headers and/or footers, consider how much content your users will actually be able to see in less than ideal conditions_
