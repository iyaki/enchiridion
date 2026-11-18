---
title: "GoatCounter – Open source web analyticshttps://www.goatcounter.com/"
notion_id: ce3855b2-c811-46d5-ba47-e50ec8c33d6d
notion_url: https://app.notion.com/p/GoatCounter-Open-source-web-analyticshttps-www-goatcounter-com-ce3855b2c81146d5ba47e50ec8c33d6d
last_edited: 2023-08-31T18:46:00.000Z
source_url: https://www.goatcounter.com/
tags: ["English", "Blogging/Content Creation", "Untried", "Decision Making", "Tool", "Service"]
---
_**Easy**_** web analytics. **_**No tracking**_** of personal data.**

GoatCounter is an [open source](https://github.com/arp242/goatcounter) web analytics platform available as a free donation-supported hosted service or _self-hosted_ app. It aims to offer easy to use and meaningful privacy-friendly web analytics as an alternative to Google Analytics or Matomo.

[Why I made GoatCounter](https://www.goatcounter.com/why)

[Sign up](https://www.goatcounter.com/signup)

Already have an account? Sign in at _yourcode_.goatcounter.com. [Forgot?](https://goatcounter.com/user/forgot)

[Live demo](https://stats.arp242.net/)

The main dashboard

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Additional information

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The “text view”, and filtering of paths

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Features

**Privacy-aware**; doesn’t track users with unique identifiers and doesn't need a GDPR notice. Fine-grained **control over which data is collected**. Also see the [privacy policy](https://www.goatcounter.com/privacy) and [GDPR consent notices](https://www.goatcounter.com/gdpr).

**Lightweight** and **fast**; adds just ~3.5KB of extra data to your site. Also has JavaScript-free "tracking pixel" option, or you can use it from your application's middleware or **import from logfiles**.

Identify **unique visits** without cookies or persistently storing any personal data ([technical details](https://github.com/arp242/goatcounter/blob/master/docs/sessions.markdown#goatcounters-solution)).

Keeps useful statistics such as **browser** information, **location**, and **screen size**. Keep track of **referring sites** and **campaigns**.

**Easy**; if you've been confused by the myriad of options and flexibility of Google Analytics and Matomo that you don't need then GoatCounter will be a breath of fresh air.

**Accessibility** is a high-priority feature, and the interface works well with assistive technology such as screen readers.

100% committed to **open source**; you can see exactly what the code does and make improvements, or **self-host** it for any purpose. See [the GitHub page](https://github.com/arp242/goatcounter).

**Own your data**; you can always export all data and **cancel at any time**.

Integrate on your site with just **a single script tag**:

`<script data-goatcounter="https://yoursite.goatcounter.com/count"
        async src="//gc.zgo.at/count.js"></script>`

The JavaScript integration is a good option for most, but you can also use a **no-JavaScript image-based tracker**, integrate in your **backend middleware**, or **parse log files**.

## Documents

Some documents about GoatCounter that don’t fit the [documentation page](https://www.goatcounter.com/help):

- [Why I made GoatCounter](https://www.goatcounter.com/why)
- [Notes about GoatCounter's design](https://www.goatcounter.com/design)
- [Analytics on personal websites](https://www.arp242.net/personal-analytics.html)
- [Why GoatCounter ignores Do Not Track](https://www.arp242.net/dnt.html)

## Pricing

GoatCounter.com is currently offered for free for reasonable public usage. Running your personal website or small-to-medium business on it is fine, but sending millions of pageviews/day isn’t.

You can [self-host GoatCounter](https://github.com/arp242/goatcounter) easily if you want to use it for more serious purposes.

Donations are accepted via [Github Sponsors](https://github.com/sponsors/arp242/) to cover server costs.
