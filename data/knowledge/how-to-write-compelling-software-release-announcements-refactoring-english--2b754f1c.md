---
title: "How to Write Compelling Software Release Announcements · Refactoring English"
notion_id: 2b754f1c-7d23-81a9-bd1c-f30c7b4f0f90
notion_url: https://app.notion.com/p/How-to-Write-Compelling-Software-Release-Announcements-Refactoring-English-2b754f1c7d2381a9bd1cf30c7b4f0f90
last_edited: 2025-11-26T17:45:00.000Z
source_url: https://refactoringenglish.com/chapters/release-announcements/
tags: ["Refactoring English", "English", "Product Management", "Communication", "Documentation", "Article"]
---
A release announcement showcases how the user’s experience is better today than it was yesterday. That sounds obvious, but most release announcements seem to forget that there’s a user at all.

So many release announcements just enumerate new features in a way that’s totallly divorced from how real people use the software. The announcement is essentially just a changelog with better writing.

For example, here’s a “changelog” style of announcing a new feature:

Bad: Describe what the dev team did rather than how it benefits the user

Don’t just tell the user that there’s a new button. Tell the user what they can do with that button.

When you create a new event, chances are that the first thing you do is copy/paste information from an existing event. Creating events this way is tedious and slow, so we’re sparing you this toil with the new Duplicate feature.

When you need to create a new event based on an existing event, go to the existing event, and select Options > Duplicate. Duplicating events is 10 times faster than copy/pasting fields.

Note that the example doesn’t boast about what the software can do. It tells the user what they can do with the software. It speaks directly to the user, describing what happens when “you” create events.

## Release notes are not release announcements

Some software companies dump their commit history into a document, add some headings, and call that a release announcement.

Not a release announcement

That is not a release announcement. Those are release _notes_.

Release notes are fine, but they’re bookkeeping. They’re boring.

If you want users to feel excited about your newest release, give them something more interesting to read than a sterile changelog.

## What to feature in a release announcement

- What can the user do in the latest version that they couldn’t before?
- Which workflows became easier?
- Which workflows became faster?

Notice that the questions don’t include, “Which feature took the longest to complete” or “What feature contains the cleverest algorithm?”

I’ve published releases where the flagship feature was a performance improvement that only took a few hours of dev work. Implementation cost doesn’t always correlate with end-user value.

## Call it “faster” not “less slow”

Some bugs are so frustrating and time-consuming to fix that you forget why you’re even fixing them in the first place. You lose sight of how the bug affects the user’s experience.

From the user’s perspective, the new version isn’t just the same software minus a bug — it’s a tangibly better experience.

Focus on the improvement rather than the flaw. Instead of declaring the absence of a bug, celebrate how you improved the user’s experience:

A release announcement should never include the phrase, “various improvements and bugfixes.” You might as well boast that the team proudly breathed air throughout development and used the latest version of the Internet.

If you can’t articulate how a change benefits your users, don’t highlight it in your release announcement. Save the exhaustive list of changes for your release notes, but even there, please leave out “various improvements and bugfixes.”

Screenshots liven up a release announcement and make it easier for users to understand new features.

Make screenshots so obvious that the reader recognizes what you want them to see [without reading the surrounding text](https://refactoringenglish.com/chapters/write-blog-posts-developers-read/#accommodate-skimmers) or zooming in. Use cropping, highlighting, or arrow indicators to draw the reader’s attention to what’s relevant.

The above image makes it difficult for the reader to see which part of the screenshot is relevant. There are so many UI elements in the screenshot, and none of them have focus.

The second screenshot draws the reader’s focus more clearly by using a tighter crop and adding an arrow indicator to identify the relevant part of the UI.

## Keep animated demos short and sweet

Good news: modern browsers support media formats for playing high-quality screen recordings. We’re long past the days of washed-out, pixellated GIFs and clunky YouTube embeds.

- Animated images: WebP and AVIF image formats can show GIF-like animated images, and they enjoy wide [browser](https://caniuse.com/webp) [support](https://caniuse.com/avif).
- Embedded videos: [H.264](https://caniuse.com/?search=mp4) and [WebM](https://caniuse.com/?search=webm)encoded videos play natively in the browser with a simple `<video>` tag. No third-party client library required.

Aim to keep demos under five seconds. 15 seconds should be the maximum. This is both for the sake of respecting the reader’s time and your server’s bandwidth. Nobody wants to sit through a two-minute video about a dialog box.

This 9-second demo is only 852 KB as a [WebP file](https://refactoringenglish.com/chapters/release-announcements/default-expiration.webp) or 100 KB as an [H.264-encoded video](https://refactoringenglish.com/chapters/release-announcements/default-expiration.mp4).

## Plan your release announcement during development

At my last company, I was in charge of both release planning and release announcements.

I once dedicated an entire release to overhauling our product’s self-update feature. The update code was so messy that it took five months to replace it, twice the turnaround of our typical release.

When I finally wrote the release announcement, I realized I’d screwed up: nothing in the release benefitted our users. Sure, future releases would be faster and less error-prone, but the user’s experience today was no better than it was yesterday.

I awkwardly tried to explain to our users why we made them wait five months for a release that essentially did nothing for them. I promised that they’d benefit from these changes soon, as the new update system would accelerate development (and it did), but I was embarrassed that the release primarily made life better for our developers rather than for our users.

From then on, I always considered the release announcement early in release planning. Given the tasks I planned for a sprint, I identified the ones that would be exciting and valuable enough to the user to include in the release announcement. By planning the announcement early, I thankfully avoided ever writing another uncomfortable explanation of a release that offers nothing to the user.

## Real-world examples of compelling release announcements

The release announcement for Gleam 1.11 puts the flagship feature right in the title: a 30% performance improvement. The title makes it crystal clear that the announcement focuses on what the user cares about.

The Gleam release announcement describes every feature in terms of how it makes life easier for their users. They never bore the reader with a dry list of changes. The announcement accompanies every change with code snippets showing how each change in the language and toolset has made life better for Gleam developers.

Figma is a design tool that earned a strong reputation for its relentless focus on user experience. This same focus is evident in their release announcements.

Figma’s UI3 release announcement presents each feature in terms of the user’s experience. Rather than saying, “This button is now here, and we added this dialog,” they tell the user “**You** can now do X to achieve Y.” They don’t just explain what changed; the showcase how the change benefits the user.

My only criticism is that Figma’s release announcement bizarrely uses 8 MB GIFs (yes, multiple) to show 5-second demo animations.

- Highlight new features and changes that excite the user.
- Describe changes in terms of what improved rather than what is no longer bad.
- Leave out vague descriptions like “various improvements and bugfixes.”
- Use screenshots that make the new feature obvious even if the reader doesn’t zoom in or read the surrounding text.
- Include animated demos, but keep them under 15 seconds.

![image](https://refactoringenglish.com/images/refactoring-english-cover-800px.webp)

To improve your writing and further your career, [pre-order now](https://refactoringenglish.com/pre-order) for instant early access and new chapters every month.
