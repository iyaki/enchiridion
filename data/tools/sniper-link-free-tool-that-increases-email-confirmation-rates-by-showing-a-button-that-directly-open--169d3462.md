---
title: "Sniper Link - Free tool that increases email confirmation rates by showing a button that directly opens the user’s email client"
notion_id: 169d3462-b6b6-463b-ab03-233ef05746b7
notion_url: https://app.notion.com/p/Sniper-Link-Free-tool-that-increases-email-confirmation-rates-by-showing-a-button-that-directly-op-169d3462b6b6463bab03233ef05746b7
last_edited: 2024-05-20T19:56:00.000Z
source_url: https://sniperl.ink/
tags: ["English", "Email", "Web Development", "Blogging/Content Creation", "Service"]
---
## Sniper Link is a free tool that increases email confirmation rates by showing a button that directly opens the user’s email client.

## Why?

Using double opt-in for your email newsletter or app can be vital for ensuring that your subscribers are real people.

But because of the extra friction in a second step, it can also be a huge source of missed activation.

## What?

Sniper Link is a widget that you can add to your sign-up forms that links subscribers directly to their email inbox — filtered down to your sender, bypassing spam filters.

You can fully customize the way it looks and behaves, so it blends in to your website.

## How?

Using our pre-built web component:

```plain text
<sniper-link
  recipient="you@example.com"
  from="justin@buttondown.email"
/>
<script src="https://sniperl.ink/v1/sniper-link.js" defer></script>
```

Or by calling our API directly:

```plain text
fetch(`https://sniperl.ink/v1/render?recipient=${
recipient
}&sender=justin@buttondown.email`)
```

You can learn more about both in [the documentation](https://sniperl.ink/docs).

## FAQ?

What email providers do you support?

Gmail, Yahoo, Proton, iCloud, Outlook, HEY, AOL, and Mail.ru, to start. We’re also monitoring for new providers and will add them as they come up.

What platforms do you support?

The redirects are based on the device being used: desktop, iOS, and Android are supported!

Are you doing anything nefarious with my subscribers?

Nope. Don’t want ’em.

I still do not trust you. Can I just supply a domain instead of the full email address?

Yup.
