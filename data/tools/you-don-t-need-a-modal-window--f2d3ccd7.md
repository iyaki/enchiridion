---
title: "You don't need a modal window"
notion_id: f2d3ccd7-a637-48f2-9281-82926b13e8c7
notion_url: https://app.notion.com/p/You-don-t-need-a-modal-window-f2d3ccd7a63748f2928182926b13e8c7
last_edited: 2023-05-26T19:30:00.000Z
source_url: https://youdontneedamodalwindow.dev/
tags: ["Website", "English", "UI/UX", "Frontend"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Just make a separate page.

Maybe you have a master/detail setup on your web app. There's a list of products/documents/whatever, and clicking on one brings up the details.

Don't open the details in a modal window. Have it be a separate page.

**Modal windows can't be bookmarked or shared as links.** Deep linking can be added to modals, but it's complex. What will show up in the background when the link is followed? How much application state needs to be restored? How will the user be confident that it will work?

**Modal windows can't be opened in a new tab.** Even if you implement this, you'll be forcing users to duplicate the page underneath.

**Modal windows make the Back button confusing.** Will it close the modal and return to the page in the background, or return to the page you were before that?

**Modal windows are hard to get right.** The "final boss of accessibility". If you have access to `<dialog>`, it's easier — if you need to support older browsers, good luck. There's a very good chance the modal window you have in production has dealbreaking bugs.

Consider if the modal window is really the best choice. There are many reasons modal dialogs are used when they're not appropriate:

Modal windows can be used for views that don't constitute a "resource" or correspond to a domain entity:

- Alerts
- Confirmation dialogs
- Forms for creating/updating entities
