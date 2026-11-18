---
title: "Why Self-host? | Roman Zipp"
notion_id: 2b754f1c-7d23-8141-89c2-e6bcf8935259
notion_url: https://app.notion.com/p/Why-Self-host-Roman-Zipp-2b754f1c7d23814189c2e6bcf8935259
last_edited: 2025-11-26T19:08:00.000Z
source_url: https://romanzipp.com/blog/why-a-homelab-why-self-host
tags: ["Article", "Roman Zipp", "English", "Privacy", "Self-hosted", "Information Security", "DevOps"]
---
![image](https://romanzipp.com/cdn-cgi/image/width=600px,quality=90,format=auto/https://cdn-a.romanzipp.com/blog/why-a-homelab-why-self-host/cover.jpg)

I recently shared my current Homelab setup with a colleague and was asked a pretty simple question I just took for granted... _**why**_?

Why go through the hassle of configuring servers, installing applications, setting up containers and spending quite a substantial amount of money on hardware that will not even run under optimal data center conditions (consumer-grade internet connection, no failover, no auto migrations)?

I will also give some specific recommendations on what you could and maybe _**should**_ self-host.

## Privacy

You saw that coming.

Privacy is not a god-given right but has to be fought for. Big Tech and governments (like with [chat control](https://fightchatcontrol.eu/) in the EU) want to shine light in every part of your personal life. Self-hosting services can reduce or even completely mitigate the risk of being surveilled. But it also requires a lot of technical knowledge so you can make a difference and educate your family or friends and even host some services for those who don't have the capabilities.

### Calendar & Contacts

Your calendar says more about you than you probably think. Apart from your full identity it can also give away information about regular contacts, family, coworkers, confidential business meetings, your health information such as medical appointments, sleep and workout routine, legal obligations, financial information like scheduled loans, subscriptions, political beliefs through scheduling to visit a protest and even let's other profile your behavior for identity theft to find out when you're available and when not.

Same goes for contacts, your social graph can say so much about you, combined with metadata such as queries for certain contacts and creation dates. Did you recently add an unusual amount of new contacts with the same sex, first name only and phone number? You must be dating. Just created a contact for a doctor? Looks like you're visiting a therapist.

Most people don't even think about where their social graph data is stored and probably assume it comes with their phone when in reality that data is being processed by Google, Apple, Samsung or whoever knows.

I don't want a single company holding all that sensitive information and possibly deriving data points from that. Even with Apple's [Advanced Data Protection](https://support.apple.com/guide/security/advanced-data-protection-for-icloud-sec973254c5f/web) your calendars and contacts are not end-to-end encrypted.

### Location

Many many years ago I was running an Android phone with Google services like Google Maps. One day I was looking for a feature in my Google account and saw that GMaps recorded my location history for years with detailed geocoordinates about every trip and every visit.

I was fascinated but also **scared about that** since I've never actually enabled it myself. I do like the fact that I could look up my location for every point in time but I want to **be in control** about that and know that **only I have access** to that data.

### So much more

It's beyond the scope of this post to list every possible way, your data can be traced back to you and argument why you should be conscious about that. I want to motivate you to start a new journey!

## Sovereignty

Digital sovereignty for me means to be in control of, choosing what I do with and controlling who I share my data with.

You constantly [hear about cases](https://www.reddit.com/r/googleworkspace/comments/1kes3ab/locked_out_of_google_workspace_for_3_days_support/) of tech companies locking down accounts with no apparent reason and it even happened to me in the past with Google. I do not want to be at the mercy of a giant tech firm which you can not even contact or if - get annoyed by a [garbage AI chat bot](https://pivot-to-ai.com/2025/10/08/salesforce-replaces-help-with-agentforce-ai-customers-outraged/) (see my [Microsoft rant](https://romanzipp.com/blog/rant-microsoft-edge-api) for more fun). _Besides - why are there no regulatory requirements for tech companies to provide a way to get in contact with an actual human?_

I like protocols and file standards, no "Gmail" API - we call that thing SMTP and IMAP (yes, they are dated but the best we currently have. Thus I still welcome the new [JMAP initiative](https://jmap.io/)). Another paragraph without bashing Microsoft? Hell no, Big Tech like Microsoft really wants you to use their AI-Copoilit-enabled-365-Office-Live-Outlook _spyware_ software - that's why they have recently [disabled SMTP access](https://www.reddit.com/r/Outlook/comments/1g56ejp/did_microsoft_recently_discontinued_smtp_support/) for Office 365 accounts.

## What to self-host

Let's get to the bread and butter of this article and give some straightforward examples on what to self-host.

Some of those applications need to be available outside of your local network if you don't want to be constantly connected to a VPN. I will write more about **how to do that securely** and all available options in an upcoming post. If you want to read that, [subscribe to the RSS feed](https://romanzipp.com/rss).

### Hardware

I'm fortunate enough to work at a company ([enum.co](https://enum.co/)) where digital sovereignty is not just a phrase. That's why I got provided with three mini servers where I'm running a highly available **Kubernetes cluster** (which my boss also helped me set up, thanks [Max](https://www.linkedin.com/in/maxheyer/)!). Also... more on that in a later blog post.

### Calendar & Contacts

As stated above, calendar and contact data is more sensitive than one might think. This is why I am hosting my own CalDAV / CardDAV server.

There are some options on servers for you which all have their ups and downs. Here are just a few:

- [Radicale](https://radicale.org/v3.html#getting-started) (Python, basic web ui, only single user, does not work with apple devices from my experience)
- ⭐ [Baïkal](https://github.com/sabre-io/Baikal) (PHP, active development, advanced web ui, multi-user)
- [DAViCal](https://gitlab.com/davical-project/davical) (PHP, haven't tried)
- [Xandikos](https://www.xandikos.org/) (Python, No built-in authentication, no web ui)
- [Nextcloud](https://nextcloud.com/groupware/) (PHP, If you're already using it go for it - too bloated for me)

Baïkal Web UI

![image](https://romanzipp.com/cdn-cgi/image/width=900px,format=auto/https://cdn-a.romanzipp.com/blog/why-a-homelab-why-self-host/baikal.jpeg)

Being conscious about what other can do with your calendar and contact data also mean to review, which apps have access to your contact book and calendar.

Oh no, I said the forbidden phrase: Self-hosted mail server. I was always told to never under any circumstances do that. But it's really not that deep.

"Recent" developments like [Stalwart](https://stalw.art/) or [Mailcow](https://mailcow.email/) made it really easy and straighforward to self-host email. Beware I'm not talking about marketing mails but rather personal inboxes.

Of course, you don't want to self-host your mail server at home since it requires a static IP and needs to be reachable from the whole internet. Going into that, you want to start with a clean IP address. Choose a hoster you _trust_, get a server, look up the IP address in mail blacklists and repeat until you get a clean one. After setting up the server, you want to make sure you can receive mail and every required protocol has been correctly configured. I found the [internet.nl online test tool](https://internet.nl/) to be super usefull to ensure everything works. Start by sending mails to Google, Microsoft and Yahoo addresses to check if your mails are getting redirected to SPAM. Iterate on that, check DNS, DMARC, SPF, TLS etc.

I will probably write a detailed blog post on that in the future.

### Smart Home

When I started hosting my own [Home Assistant](https://www.home-assistant.io/) instance a couple of years ago it was just an experiment to see what I can do since I wasn't really missing anything with Apple Homekit. Since then more and more smart home companies went bankrupt, sunsetted their cloud services, jacked up prices or put free services behind a paywall.

For me, Home Assistant paid off a couple of weeks ago when I heard that [Philips Hue will force users to create an account](https://consumerrights.wiki/w/Philips_Hue_starts_requiring_an_account_for_the_hue_app) just to use **any feature** for their lights, they already paid real money for. I've always configured Firewall rules to disallow any outgoing network traffic for smart home appliances but it seems like I cannot use any Philips Hue app specific features (like animated light patterns imitating candles etc.) even on my local network. I haven't looked into this but I hope there's some community plugin which emulates this functionality.

I will **never**, under any circumstances, create an online account for an appliance I will only use locally.

_Also I am now obsessed with tracking energy usage and plan on building and developing a Raspberry Pi + camera device which tracks energy usage of gas meter via machine vision._

Home Assistant

![image](https://romanzipp.com/cdn-cgi/image/width=900px,format=auto/https://cdn-a.romanzipp.com/blog/why-a-homelab-why-self-host/hass.jpeg)

### RSS Aggregator

I am subscribed to many news sites and blogs over RSS which is by itself already decentralized and sovereign. This is why self-hosting an RSS aggregator is kind of optional and only the last mile to go.

On my iPhone and Mac I'm running [NetNewsWire](https://netnewswire.com/), in my opinion the best RSS reader, even open source and backed by [incredible people](https://inessential.com/). NetNewsWire comes with a native integration for [FreshRSS](https://freshrss.org/index.html) - a feed aggregator that also provides many more features like filtering and lets you subscribe to sources which don't natively provide an RSS feed.

### Location Tracker

I've deployed an instance of [dawarich](https://dawarich.app/) (German for "_I was there_") which is a server for ingesting and viewing geolocation data. It also allows you to choose from many available mobile apps which can track send your current location to the server. At the time of writing this includes:

- [official dawarich app](https://apps.apple.com/de/app/dawarich/id6739544999?itscg=30200&itsct=apps_box_badge&mttnsubad=6739544999) (always shows a navigation icon in the iOS notch)
- [Overland](https://overland.p3k.app/) (high battery drain for me)
- ⭐ [Owntracks](https://owntracks.org/) (works best for me on iOS, only app settings are crazy confusing)
- [PhoneTrack](https://f-droid.org/packages/net.eneiluj.nextcloud.phonetrack/)

dawarich Web UI

![image](https://romanzipp.com/cdn-cgi/image/width=900px,format=auto/https://cdn-a.romanzipp.com/blog/why-a-homelab-why-self-host/dawarich.jpeg)

### Ideas & Outlook

I recently re-worked my homelab and went from a single big server to a 3 node Kubernetes cluster. This gives me much more flexibility in the kind of applications I can host.

This is a list of apps and tools I want to have a look at:

- [EteSync](https://www.etesync.com/): End-to-end encrypted CalDAV & CardDAV
- [AnyType](https://doc.anytype.io/anytype-docs/advanced/data-and-security/self-hosting/self-hosted): Self-hosting my own AnyType server instance
- [Iimmich](https://immich.app/) or [ente](https://ente.io/): Moving from iCloud photos to self-hosting
- [Passbolt](https://www.passbolt.com/): Password manager (no I really don't like Bitwarden)
- [BirdNET](https://github.com/tphakala/birdnet-go): Monitoring bird species outside with a microphone
- [penpot](https://penpot.app/): Like Figma but free & open source
- [Habitica](https://habitica.com/static/home): Habit manager
- [vert](https://vert.sh/): File converter
- [InvoiceShelf](https://github.com/InvoiceShelf/InvoiceShelf): Invoice manager

There's also [selfh.st](https://selfh.st/) - a great resource where can spend hours finding self-hostable applications.

## Read more...
