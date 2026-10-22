---
title: "How Do You Like What You’ve Built?"
notion_id: 17454f1c-7d23-81a8-a97d-ef938a4055fd
notion_url: https://app.notion.com/p/How-Do-You-Like-What-You-ve-Built-17454f1c7d2381a8a97def938a4055fd
last_edited: 2025-02-12T22:08:00.000Z
source_url: https://morrisbrodersen.de/how-do-you-like-what-you-built/
tags: ["English", "Product Management", "Line/People/Team Management", "Article", "Morris Brodersen"]
---
_January 3rd, 2025_

Recently, a colleague asked me for a code review on some complex UI changes. The idea was to make an error-prone form easier to use by displaying warnings in the right moment, and automating various state changes in the background (think “smart configurator”). The requirements were given upfront, so my colleague was in execution mode, turning these requirements into code.

Glancing over the pull request, nothing stood out as problematic, but I didn’t feel ready to approve the changes before playing with the new UI.

However, before testing it myself, I asked (via chat message):

> How do you like the new behavior?

I went to grab a coffee; when I came back, they had answered with roughly the following:

> Not sure, to be honest.

_A couple minutes pass..._

It feels a bit off, like sometimes I can’t really make the changes I want as a user.

_Another pause..._

Actually there’s still another bug I need to fix, let me get back to you. This needs to be rock-solid for our users!

I can only assume what happened in the meantime, but it seems asking them for their opinion (and _not_ responding for a while) was a cue for [sanding the UI](https://blog.jim-nielsen.com/2024/sanding-ui/), with great results: In the end, a problematic requirement was dropped and multiple commits were added before the actual review.

Later, knowing how the author had engaged with their solution, reviewing and approving the changes felt very comfortable.

I find it remarkable how a simple question can have such an outsized impact. Of course, my colleague’s particular reaction to it was fantastic, and not to be taken for granted. They took ownership and got to work, and were able to critically look at the requirements and their solution. Thank you!

In any case, I believe asking engineers how _they_ like what they’ve built to be a useful trick to try from time to time. I suspect only upside. ∎
