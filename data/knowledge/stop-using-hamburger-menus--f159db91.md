---
title: "Stop Using Hamburger Menus"
notion_id: f159db91-17b6-4780-a425-a3e5ddc29a06
notion_url: https://app.notion.com/p/Stop-Using-Hamburger-Menus-f159db9117b64780a425a3e5ddc29a06
last_edited: 2023-05-12T10:59:00.000Z
source_url: https://bt.ht/hamburgers/
tags: ["bt.ht - Brad Taunt Blog", "English", "UI/UX", "Article"]
---
I recently [tooted about my hatred of website hamburger menus](https://fosstodon.org/@tdarb/110264983268249599) which was met with a surprising amount of support from other users. It seems like most people _don't actually like hamburger menus_. So why do we, as developers, keep using them in our products and designs? Is it our only option? Or is it because we have become conditioned to expect it?

## The Core Problem with Hamburger Menus

The biggest headache when coming across these menus on the web is the complete disregard for **accessibility**. Performance and solid user experience is almost always thrown out the window in favour of a "prettier" design layout. You might have made the overall designer "cleaner" for your users, but you sacrificed all usability to do so.

I challenge you to visit a webpage or web app with a hamburger menu and try to navigate solely with your keyboard and screen-readers (or better yet - try these on mobile!). Within seconds you will find a whole mess of issues. Now try the same test with JavaScript disabled... Yikes.

## "But I Have No Choice!"

I see this argument pop-up frequently when taking to design leaders or developers. I call bullshit on this excuse. You _absolutely_ have the choice to avoid implementing bad designs - that's your job! Either you're not fighting hard enough against those pushing for it, or you're just trying to build a "pretty" portfolio.

## Best Alternative: Sitemap Footer

So instead of just whining about hamburger menus, I will actually offer up a solid replacement: **sitemap footers**. Simply place all your website/application links into the bottom footer and link directly to them from your header. Be sure to also include some form of "Top of the page" link for quick access back to the initial scroll view.

That's it. There is nothing else you need to do for this to _just work_. It might sound oversimplified and that's because it is. Looking for an example? This very website utilizes this technique, so give it a spin! Try using just your keyboard or even better - use a screen reader. Disable JS and CSS and watch it work flawlessly still.

**Pros:**

- Keyboard navigation accessible
- Excellent screen-reader support
- Works on all devices/screens by default (no media queries!)
- Stays out of the way until called upon (UX goodness)
- Requires ZERO CSS or JavaScript

**Cons:**

- Footer can become large with many links (although I _really_ don't see this as a big deal)

## No Excuse

There really is no excuse to still be using hamburger menus. Users expect them to be present only because we as designers have conditioned them think that way. They deserve a better experience on the web. The _least_ we can do is improve something as simple as website navigation...
