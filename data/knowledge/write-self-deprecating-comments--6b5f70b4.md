---
title: "Write self-deprecating comments"
notion_id: 6b5f70b4-a954-46ec-b9ff-a56e6ef50a71
notion_url: https://app.notion.com/p/Write-self-deprecating-comments-6b5f70b4a95446ecb9ffa56e6ef50a71
last_edited: 2026-09-21T17:34:00.000Z
source_url: https://thepugautomatic.com/2021/02/write-self-deprecating-comments/
tags: ["Article", "The Pug Automatic", "English", "Programming"]
---
[https://thepugautomatic.com/2021/02/write-self-deprecating-comments/](https://thepugautomatic.com/2021/02/write-self-deprecating-comments/)

Written February 27, 2021. Tagged Code style.

Comments and code easily get out of sync, but there are tricks to lessen the impact.

Instead of

```plain text
PaymentAPI.call( mode: "X", # Disable 3D Secure verification. timeout: 12, # The smallest value that avoids errors.)
```

, write

```plain text
PaymentAPI.call( mode: "X", # "X": Disable 3D Secure verification. timeout: 12, # 12 secs is the smallest value that avoids errors.)
```

This double-entry bookkeeping means you can easily tell when the code and comment drift apart:

```plain text
PaymentAPI.call( mode: "Y", # "X": Disable 3D Secure verification. timeout: 15, # 12 secs is the smallest value that avoids errors.)
```

Whether the discrepancy is caught immediately by the author, or in review, or by another developer far down the line, it will be explicitly clear that the comment was not intended for the current value.

This technique is a great fit for short-and-cryptic values like these. Longer values would be annoying to repeat, but also tend to be more self-documenting.

Content and design © Henrik Nyh (@henrik@ruby.social). Code is under a MIT License unless otherwise stated.

Pug art by Johanna Öst; other graphics are under a CC BY License.

Powered by Eleventy.
