---
title: "How to use Matrix - open network for secure, decentralized communication"
notion_id: dae2f519-a22f-4e11-941d-3814e199da9a
notion_url: https://app.notion.com/p/How-to-use-Matrix-open-network-for-secure-decentralized-communication-dae2f519a22f4e11941d3814e199da9a
last_edited: 2023-02-17T19:35:00.000Z
source_url: https://akselmo.dev/2022/12/29/How-To-Use-Matrix.html
tags: ["Communication", "Information Security", "Privacy", "Article", "Tool", "Service", "English"]
---
As I’ve gotten more into FOSS, I’ve noticed a lot of FOSS projects, like KDE, use [Matrix](https://matrix.org/) protocol.

A protocol that lets you use any client you wish to chat with each other. Basically, it’s IRC but with more features.

And I’m into that! So I decided to move my [Aks_Dev community](https://matrix.to/#/#aksdev-space:matrix.akselmo.dev) from Discord to Matrix. Just to support the project in general but also I do not want my community to be tied to one client. In fact I’ve been using [Matrix-Discord bridge](https://github.com/matrix-org/matrix-appservice-discord) for a while, but it is not the optimal way to chat, IMO. Some things just drop out, like replies from Discord side. It works but.. It would be better if everyone was just on the Matrix side.

On top of the bridge being Very Good but Not Optimal, Discord being proprietary platform can easily follow the path of Twitter. Who knows when Discord decides to add [cryptocurrencies](https://www.nme.com/news/gaming-news/discord-walks-back-cryptocurrency-integration-plans-3093283) or other meaningless crap? Who knows where your data is going to? And if you want to use different client, Discord just can easily ban you because it’s their walled garden and you’re the prisoner of convenience.

_Ahem_..

Enough of ranting, here’s how you get most of Da Matrix, in my holy opinion:

# Create account to one of the instances

Try to avoid Matrix.org if you can. Meanwhile it is the biggest instance, that is also its downside. If Matrix.org runs out of money to host the server, then it’s bye-bye for all the accounts using it.

I think one good alternative to Matrix.org is Mozilla’s instance: [https://wiki.mozilla.org/Matrix](https://wiki.mozilla.org/Matrix)

If you are savvy enough, just self-host your own instance. But if you know how to do that, you probably are not reading this post anyhow.

You can also get your own instance through these hosting services: [https://matrix.org/hosting/](https://matrix.org/hosting/)

Make sure the instance you join is not ACL-banned (Access Control List) by Matrix.org. This will make it harder for you to federate with other servers.

Here is also one list I found from searching: [https://joinmatrix.org/servers/](https://joinmatrix.org/servers/)

Pick one with rules and privacy policy, then make sure you read the rules well. Check that there is nothing _weird_ about the instance.

_But I don’t want to think too much of this!_

Then, go for either Matrix.org or Mozilla’s instance.

# Use Element to set up things

Meanwhile [Element](https://app.element.io/) is not my daily-drive client, I always create the new account using it and set it up wit it. It is made by people who also work with Matrix protocol, so it usually has all the features needed to get you going.

So, use Element to create your account and set it up. I think of Element as the “control panel” of my account, and not the actual chat client I use.

# Pick your client

This is not necessary if you like Element and it has all you need! But I’ve noticed many new Matrix users have problems using Element since it has _a lot of going on!_

So many buttons and widgets. So many things cluttering around. And no custom emoji support (as of writing this)!

This is why I recommend getting another client. My favorite newcomer-friendly client is definitely [Cinny](https://cinny.in/)!

It has custom emoji support, it has less noisy interface, it works with E2EE and basically, it works well as a basic chat client.

There is also desktop version of Cinny, but it’s a bit hidden for some reason: [https://github.com/cinnyapp/cinny-desktop/releases/](https://github.com/cinnyapp/cinny-desktop/releases/)

You can find more clients from here if you’re curious: [Official client list](https://matrix.org/clients/)

However, what I would do is the following:

**Element** as control panel and back-up client. **Cinny** as the main client I use daily, either on desktop or web-browser.

# E2EE verification stuff

_E2EE = end-to-end encryption_

This is where the UX is not very good yet… Basically, when logging in to Cinny, you’ll see a popup that says something about “verify your client/account.”

This means that the encryption key, that is used to encrypt your conversations, will be shared between these clients. So when you verify a device, the client with the key will send it to the other client. In this case it would be from Element to Cinny.

The usual verification dance is following:

1. You have set up your account in Element, but want to now try a new client.
2. You open the new client in another window, for example Cinny, and log in. **Keep Element open!**
3. Cinny or Element warns you that your account is unverified, since there is a new client connected to it.
4. You click the shiny Verify button in either client (I usually always start the verification process from Element)
5. You accept the verification process in the other client.
6. You have to check if some emoji’s match and say Yes in both clients.
7. This dance now should be done.

I think there are better ways to do this, like QR codes or something, but it’s clearly still in WIP to see what kind of verification method will be the one that’ll stick.

If you do not use chats with E2EE enabled, the verification process is not necessary. But I do it anyhow because the popups are annoying, lol.

My community does not use encrypted chats, because bots do not work with those.

# Android client

I don’t know anything about iOS, but on Android, [Fluffychat](https://fluffychat.im/) is what I use. Works well for my usage. Android version of Element is also fine, but it doesn’t have custom emoji support. See the client list linked above for more.

And yes, you’ll have to do the verification dance here as well. Again I recommend starting it from Element on desktop.

# Voice and Video

These are still WIP. I think only Element and [Nheko](https://nheko-reborn.github.io/) clients have voice chat working, and I’m not sure if anything else but Element has video chat working.

For these, you’re probably best served with Discord. If you want FOSS voice-chatting, check out [Mumble](https://www.mumble.info/). It’s the old reliable and I’ve used it for years.

For video chats, I do not know at all what would work. Leave comment if you do, though!

# Quick Summary for the TLDR peeps

1. Use Element to create an account to Matrix.org or Mozilla’s instance.
2. Set up your account inside Element.
3. Download Cinny for your desktop and/or use it from your browser.
4. Use Cinny as your chat client, Element as your “control panel” client.
5. Get some mobile client if you wish.
6. Set up E2EE by verifying the other clients from Element.

# What about XMPP and IRC? Why Matrix?

I actually love both XMPP and IRC. I’ve used both and they still are very reliable ways to chat.

However, most folk I chat with seem to use Matrix, and if you run your own homeserver like I do, you can easily bridge to both XMPP and IRC if needed. That’s why I’ve chosen to stick with Matrix in this case.

Personally I don’t really care about the protocol, I just want something that is not proprietary and the communities I am in use. Matrix seems to bridge with anything so for me it made most sense.

**The best thing we can do, however, is not to fight over protocols but make these open protocols easily to talk to each other, so we can choose anything we wish.**

**To have any chance against proprietary stuff, we must not fight among ourselves, but work together. We are all in this together, after all!**

Aaaand that’s all. I hope this “guide” helped a bit.

Matrix is still very young and figuring things out. It’s complicated and I believe they’re adding features a bit too fast, so things are not always very fleshed out.

But for basic text chatting it works fine for me, and is not proprietary.

So if you want a community to join in with your Matrix account, hop on to Aks_Dev chat!

You can ask me questions about Matrix in there or just comment to this post. :)


