---
title: "Toasts are Bad UX"
notion_id: 7fcf27cb-f62e-485f-a5d3-6679095b060a
notion_url: https://app.notion.com/p/Toasts-are-Bad-UX-7fcf27cbf62e485fa5d36679095b060a
last_edited: 2024-09-10T20:10:00.000Z
source_url: https://maxschmitt.me/posts/toasts-bad-ux
tags: ["English", "UI/UX", "Article", "Max Schmitt"]
---
The core problem is that toasts always show up far away from the user's attention.

Take a look at this example from YouTube:

[https://s3.eu-central-1.amazonaws.com/maximilianschmitt.me/v2/post-media/toasts-bad-ux/youtube-toast.webm](https://s3.eu-central-1.amazonaws.com/maximilianschmitt.me/v2/post-media/toasts-bad-ux/youtube-toast.webm)

## The Problems with the YouTube Toast

In this particular example, the entire interaction is quite jarring:

- I click the "Save" button on the right-hand side of the screen
- A modal appears in the middle of the screen
- The toast appears in the bottom left corner

And there are a few more problems in this particular example:

- The toast is delayed without a loading indicator
- If I check or uncheck a checkbox in the modal, I need to wait multiple seconds for the previous toast to disappear before I get the confirmation toast for the latest action
- The "Undo" button in the toast is unnecessary because the user can just click the checkbox again

## The Solution: No Toast

Here is a simple redesign of the "Save" interaction that solves all the problems above:

![image](https://maxschmitt.me/post-media/toasts-bad-ux/youtube-toast-simplified@2x.jpg)

A simplified version of the YouTube toast interaction

- The playlists are shown right beneath the button instead of in a modal
- After checking/unchecking a checkbox, a loading indicator is shown
- When the loading indicator disappears, it implies the action has completed
- No toast necessary!

## 2 More Examples

### 1. Confirming that Something was Added/Removed

When archiving an email in Gmail, a toast appears showing confirmation. But by archiving the email, the email disappears from the list, which already implies the action was successful.

![image](https://maxschmitt.me/post-media/toasts-bad-ux/gmail-toast@2x.jpg)

A toast in Gmail confirming that an email was archived

Note

We do have to consider the undo-functionality and that the toast feedback can be useful when using keyboard shortcuts.

### 2. Confirming that Something was Copied

A toast is shown after something was copied to the clipboard. In this example, the button already includes a confirmation so the toast is entirely unnecessary.

![image](https://maxschmitt.me/post-media/toasts-bad-ux/iconfinder-toast@2x.jpg)

A toast confirming that something was copied to the clipboard

## It Could be Worse

What's worse than a toast? No feedback at all.

So if you don't have time to design or build a better feedback mechanism, a toast is better than nothing.
