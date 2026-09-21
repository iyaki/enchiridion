---
title: "Read-only web apps"
notion_id: 0a0020b2-58bc-4a8e-b68e-022573675c61
notion_url: https://app.notion.com/p/Read-only-web-apps-0a0020b258bc4a8eb68e022573675c61
last_edited: 2024-02-23T22:52:00.000Z
source_url: https://adactio.com/journal/20113
tags: ["Article", "adactio (Jeremy Keith)", "English", "Web Development", "Reflection", "UI/UX"]
---
The most cartoonish misrepresentation of progressive enhancement is that it means making everything work without JavaScript.

No. Progressive enhancement means making sure your _core functionality_ works without JavaScript.

In my book [Resilient Web Design](https://resilientwebdesign.com/), [I quoted Wilto](https://resilientwebdesign.com/chapter5/#Mat%20Marquis):

> Lots of cool features on the Boston Globe don’t work when JS breaks; “reading the news” is not one of them.

That’s an example where the core functionality is readily identifiable. It’s a newspaper. The core functionality is reading the news.

It isn’t always so straightforward though. A lot of services that self-identify as “apps” will claim that even their core functionality requires JavaScript.

Surely I don’t expect Gmail or Google Docs to provide core functionality without JavaScript?

In those particular cases, I actually _do_. I believe that a textarea in a form would do the job nicely. But I get it. That might take a lot of re-engineering.

So how about this compromise…

Your app should work in a read-only mode without JavaScript.

Without JavaScript I should still be able to read my email in Gmail, even if you don’t let me compose, reply, or organise my messages.

Without JavaScript I should still be able to view a document in Google Docs, even if you don’t let me comment or edit the document.

Even with something as interactive as Figma or Photoshop, I think I should still be able to _view_ a design file without JavaScript.

Making this distinction between read-only mode and read/write mode could be very useful, especially at the start of a project.

Begin by creating the read-only mode that doesn’t require JavaScript. That alone will make for a solid foundation to build upon. Now you’ve built a fallback for any unexpected failures.

Now start adding the read/write functionally. You’re enhancing what’s already there. Progressively.

Heck, you might even find some opportunities to provide some read/write functionality that doesn’t require JavaScript. But if JavaScript is needed, that’s absolutely fine.

So if you’re about to build a web app and you’re pretty sure it requires JavaScript, why not pause and consider whether you can provide a read-only version.
