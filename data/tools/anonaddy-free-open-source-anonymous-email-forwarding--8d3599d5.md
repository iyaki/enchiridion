---
title: "AnonAddy - Free, Open-source Anonymous Email Forwarding"
notion_id: 8d3599d5-e35c-4b59-8333-ba3b771f6506
notion_url: https://app.notion.com/p/AnonAddy-Free-Open-source-Anonymous-Email-Forwarding-8d3599d5e35c4b598333ba3b771f6506
last_edited: 2026-09-18T00:54:00.000Z
source_url: https://anonaddy.com/
tags: ["Email", "Information Security", "Untried", "Service", "English"]
---
## How Does It Work?

### 1. Register Your Username

Let's say your username is **johndoe**. You can now use *@johndoe.anonaddy.com (or .me) as your email. Where * denotes any valid local part for an email address.

If you would like to remain anonymous choose a username that is not linked to your real name or identity and that you haven't used anywhere else.

You can also create aliases at **shared domains** if you are concerned about others linking alias ownership to you.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 2. Create Aliases

Either **on the fly** or **generated** beforehand. Next time you are signing up to a website or newsletter, simply make up a new alias and enter that instead of your real email address.

For example, if you are on vuejs.org and you want to sign up to their newsletter you could simply enter **vuejs@johndoe.anonaddy.com** (or .me).

We'll automatically create the alias in your dashboard as soon as it receives its first email. You can reply to emails and send from aliases anonymously too!

### 3. Manage Aliases

Let's say a spammer gets hold of one of your aliases and starts sending unsolicited email to it. You can simply toggle a switch in your dashboard and deactivate that alias.

Our system will then silently **discard** any further emails and you won't be forwarded anything else for that alias.

You can also delete the alias. Then our system will **reject** any emails and respond with an error.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Features You'll Love

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Generate new aliases in a couple of clicks straight from your browser using the [open-source](https://github.com/anonaddy/browser-extension) browser extension.

The extension is available for download on [Firefox](https://addons.mozilla.org/en-GB/firefox/addon/anonaddy/) and [Chrome](https://chrome.google.com/webstore/detail/anonaddy/iadbdpnoknmbdeolbapdackdcogdmjpe)!

Also available on other chromium based browsers such as Brave and Vivaldi.

### GPG/OpenPGP Encryption

Bring your own GPG/OpenPGP public keys and add them per recipient.

You can then easily toggle encryption on and off. With encryption on, all forwarded messages will be encrypted with your public key. Only you will be able to decrypt them with the corresponding private key. You can even hide and encrypt the email subject!

This is great if you are using Gmail or Outlook and wish to prevent any inbox snooping.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If you have your own domain you can add it and use it exactly like your username subdomain aliases e.g. **alias@example.com**.

You can also enable or disable the **catch-all** functionality for each of your domains.

You can then manage your aliases and deactivate/delete any that start receiving spam!

### Blend Into The Crowd

If you don't like the fact that all your aliases contain your username then you can generate random unique aliases from your dashboard.

The generated aliases will look something like this **x481n904@anonaddy.me** or **circus.waltz449@anonaddy.me**

This prevents anyone linking ownership of the alias to you.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Reply **anonymously** to forwarded emails, the sender will receive the email as if it has come from the alias.

You can even initiate an email conversation by sending an email from one of your aliases.

Your real email address is **not revealed** when replying or sending from an alias.

### Add Additional Usernames

You can add additional usernames to your account and use them exactly like the one you signed up with.

So if you signed up as johndoe you can add johnsmith as an additional username and then use **anyalias@johnsmith.anonaddy.com** too.

This can be used to compartmentalise your aliases. You could have a username for work emails a different one for personal emails etc.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

If you'd like an alias to go to more than one recipient you can easily add multiple recipients from your dashboard.

You can even add recipients to an alias as it is created by doing **alias+2.3.4@user.anonaddy.com**

Where 2,3 and 4 are the keys for existing recipients in your account.

### API Access

Manage your aliases, recipients, domains and additional usernames using the AnonAddy API.

In order to use the API you first need to generate an API access token in your account settings.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Pricing
