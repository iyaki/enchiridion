---
title: "Self-Hosted SaaS Alternatives: Replacing Paid Tools With FOSS Tools"
notion_id: 2ecf59a9-2b9c-43c7-b229-6c2f8fe67bca
notion_url: https://app.notion.com/p/Self-Hosted-SaaS-Alternatives-Replacing-Paid-Tools-With-FOSS-Tools-2ecf59a92b9c43c7b2296c2f8fe67bca
last_edited: 2023-03-09T14:54:00.000Z
source_url: https://tedium.co/2023/03/04/self-hosted-saas-app-alternatives/
tags: ["Article", "Tool", "English", "Hosting", "Privacy", "SysAdmin"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## In a world where every service wants to put you on a subscription forever, is now the time to start self-hosting your apps instead? Let’s look at the landscape.

**Filed under:** [applications](https://tedium.co/tag/applications/), [apps](https://tedium.co/tag/apps/), [google analytics](https://tedium.co/tag/google-analytics/), [saas](https://tedium.co/tag/saas/), [self-hosted](https://tedium.co/tag/self-hosted/)

**Today in Tedium:** Recently, I got a bit of a shocking message in the form of a bill. The bill came from a company named Zapier, which produces a tool for automating tasks of all kinds, such as uploading images to a website or adding lines of text to a spreadsheet. This tool was great for a long time, but then the company decided to raise the price out of nowhere. And not by a little, but by more than three times the cost of what I was paying. I immediately cut back the plan to the minimum needed in response to this unexpected price increase, which honestly kind of upset me a lot. This got me thinking: Why pay all this extra money to another software-as-a-service (SaaS) platform that doesn’t necessarily have my best interests at heart? Is there a better way? So, with that in mind, I’ve been doing some research on tools that I think could be replaced with open-source alternatives without too much trouble, as long as the tools make sense for my needs. But, as always, the rub is always there. Today’s Tedium considers open-source alternatives to a few common categories of software-as-a-service tools. _— Ernie @ Tedium_

[Sponsored By Tech Brew](https://links.swapstack.co/esr)

### Tech news is confusing—so we changed that.

**Join over 450K people reading **[**Tech Brew**](https://links.swapstack.co/esr)—the 3-times-a-week free newsletter covering all updates from the intersection of technology and business. We provide free tech knowledge that’ll help you make more informed decisions.

[Subscribe For Free](https://links.swapstack.co/esr)

_(_[_Internet Archive Book Images/Flickr_](https://flickr.com/photos/internetarchivebookimages/14781263771/)_)_

### Too hot or too cold? Our ranking system for self-hosted apps

Today’s issue will ponder each of these open-source alternative categories by offering three alternatives that fit a “Goldilocks model”—too hot, too cold, or just right. Our goal with this list is not to shame projects, but to offer feedback on what might make a tool a better fit for a broader audience.

### The Methodology

- **Too hot** will refer to tools that are still relatively new and lack a number of features that users might expect from a tool of this nature. Limitations that make the technology a nonstarter will also get a nudge here.
- **Too cold** will discuss tools that are either too complex for the average person’s use case, or have too much cruft to make them easy to parse for regular users. Issues with user interface will also play a factor around where a given tool ranks.
- **Just right** will highlight tools that meet most of the requirements of ease-of-use, feature set, and maturity.
- **Outside the box:** An open-source idea that doesn’t _quite_ fulfill the ask, but could still lead you to a reasonable answer.

For these tests, we either installed the applications locally using docker or [ddev](https://ddev.com/), tested demo versions of each given tool, or spun up cloud instances on [Hetzner](https://www.hetzner.com/) to give each tool a spin in a live environment. I have tried a number of cloud services over the years, and my recommended suggestion for most of these is either to put them on a small machine if it’s for hobby reasons, install them on a cloud service like Hetzner if it’s public-facing, and invest more deeply if it’s more heavyweight.

I’m pretty good at setting up tools at this point, having done so using esoteric approaches such as Node, Rails, and Laravel (along with the ever-present LAMP stack). I’ve set up a Mastodon server on my own, even. If I couldn’t install it, or I had an unnecessary amount of trouble, I’m going to tell you. _(Foreshadowing …)_

_(_[_Stephen Phillps/Unsplash_](https://unsplash.com/photos/shrXn8S8QU)_)_

### Analytics tools: You don’t need Google Analytics

At the moment, the analytics space is probably more disrupted than it’s been in quite some time, in part because Google Analytics is making some major changes that might convince some diehards to jump ship in favor of something more comfortable.

Google Analytics 4 is in the midst of disrupting the lives of site owners and marketers alike, and it feels like if you’re going to head for the exits with your analytics tools, you might as well do it now.

It’s a good time to go self-hosted, too: This space is full of worthy competitors for the first time in many years. And on top of that, with concerns about the General Data Protection Regulation top of mind for many users, the idea of hosting this data in a first-party way suddenly looks a lot more attractive. Anyway, pretty much any of the options listed below could work for most regular people with websites to manage, though at least in my mind, one stands head and shoulders above the rest.

### Self-hosted Google Analytics alternatives

- **Too hot:** [Umami](https://app.umami.is/share/8rmHaheU/umami.is) is only slightly warmer than the winner in this category, and is actually quite similar in many ways. Umami’s hosted cloud service is not quite ready for mainstream use yet, giving its competitor the edge for now. But watch this space.
- **Too cold:** [Matomo](https://matomo.org/) is only a little colder than the alternatives, and not by much. It’s still quite good. Its interface, [which can be viewed in demo form here](https://demo.matomo.cloud/), is more like Google Analytics than the alternatives, but in a way that actually kind of makes sense.
- **Just right:** [Plausible](https://plausible.io/) strikes the perfect balance in terms of cost and usability. The company makes an open-source version that is quite good, but also offers a paid version that costs enough that you know you’re the customer, not the product. Easy to set up and friendly to use.
- **Outside the box:** You could always do with Mastodon does and just eschew analytics entirely.

Beyond the potential-for-disruption reasons, I think Google Analytics should eventually fade out for many companies because it is too close to its source of truth—Google.

_(_[_Plausible_](https://plausible.io/)_)_

That’s to say that I think regulatory concerns will eventually lead Google to weaken the product, and as a result, we should take an approach that shows more respect to our users and doesn’t siphon data where it’s not needed.

I’m still running Google Analytics on Tedium, but I’m also running Plausible, and I expect to drop GA soon.

_(_[_Brett Jordan/Unsplash_](https://unsplash.com/photos/LPZy4da9aRo)_)_

### Email newsletter tools: Old or new, your pick

As we’ve written many times, you don’t need to use a service like Mailchimp, or even a service like Substack. Email is email—it’s the water of the internet, in that it flows everywhere. You can send it however you want.

If you’ve ever looked at different email tools as I have, you’ll know that the problem is not a lack of options but a glut of them, and this is true even of the open-source options. The problem is, email newsletter software is in this weird state where most of the tools are either too old and mature or fairly new and untested.

This space is such that you might want to also consider some of the paid alternatives, such as Sendy, which has many of the bells and whistles of SaaS tools but can be self-hosted and even integrated with WordPress. I’m not recommending Sendy, but only because I want to lean towards FOSS in this as much as I can. I know a lot of people who use it and have good experiences, though.

Anyway, after some looking around, I have to say that it’s kind of a crapshoot—a mixture of tools that are arguably a bit overgrown and tools that are so new that they feel like their chances of long-term survival are low.

_(Listmonk)_

With each of these, you have a number of options, but you will likely be working in the confines of an existing service like Mailgun or Amazon SES to actually send the emails. So sure, you’ll be doing the self-hosting yourself, but Amazon will be managing the message sends for you. You actually want this, because hosting on a new domain can get very costly.

Anyway, a few options to identify:

### Self-hosted Email newsletter platforms

- **Too hot:** [Keila](https://www.keila.io/), a very-polished looking email sender that will fit many users’ needs, but its emphasis on a single flexible template makes it difficult to recommend for power users. (They are working on it, however—and the team seemed responsive to feedback.) [Sendportal](https://sendportal.io/) also has potential, but the fact it hasn’t been updated in more than a year gives us pause.
- **Too cold:** [Mailtrain](https://github.com/Mailtrain-org/mailtrain) has most of the features one might expect out of an email-sending tool, including power-user features like support for MJML, but its user interface can feel imposing to new users. A less-imposing but still-complex alternative is [Mautic](https://www.mautic.org/).
- **Just right:** [Listmonk](https://listmonk.app/) has a clean interface and lightweight feature set. It lacks some bells and whistles that some users might expect, like on-the-fly segmentation, but it’s closer to being a daily driver than any other tool in this category.
- **Outside the box:** Hosting something through your CMS is not a bad idea. [Ghost](https://ghost.org/), while not anywhere near as full-featured as the other tools mentioned for email use cases, strikes a good balance for the average creator.

One challenge I often see with email tools is the ability to manage spam, and with that in mind, I plan to give ListMonk and Keila some deeper tests to determine what’s possible with both of them. I have ended up having to pay for an additional service to manage email lists just because of spam issues, and I would like to stop doing so.

### No-code automation tools: Can you actually ditch Zapier?

Given their significant price increases of late, it might also be worth asking if Zapier is really what you need. If your goal is simply automating tasks, you might be better off taking the hosted service out of it entirely and instead using a tool that can push information around the internet.

The problem is, Zapier tends to be the default option in many cases, and it supports a far wider variety of apps than many other services. The good news is that there are plenty of alternatives, especially in the for-pay space, with a couple that cost significantly less including Pabbly and Make.

_(_[_ActivePieces_](https://www.activepieces.com/)_)_

But honestly, this might just be something you want to self-host. With that in mind, here are a few options:

### Self-hosted Zapier alternatives

- **Too hot:** [Automatisch](https://automatisch.io/) is the closest thing to a direct open-source variant of Zapier at this time, as its design carries the spirit of the paid tool’s interface, but it is brand new, having just launched only in November, and doesn’t have many apps. It feels like one to watch, though. Unlike some of the other options, it falls under a traditional AGPL license, however. I had no trouble installing this locally.
- **Too cold:** [N8n](https://n8n.io/) has the right number of apps, a larger community, and is much more mature than many of its competitors, but its less-intuitive design is potentially a turnoff. (Despite being free to self-host via a community option, it also [does not have an open-source license](https://github.com/n8n-io/n8n/blob/master/LICENSE.md).) It can be run inside a desktop application, though, which is a plus.
- **Just right:** [ActivePieces](https://www.activepieces.com/) seems to have the right momentum to get things going in a good direction; it also has a reasonably priced cloud option, and while it is also still new, it has a lot of potential and a quickly growing app list. (It, too, [doesn’t use](https://github.com/activepieces/activepieces/blob/main/packages/ee/LICENSE) a pure GPL-style license, however. Another responsive-to-feedback team, however.)
- **Outside the box:** If your goal is just automating some basic tasks, the browser extension [Automa](https://www.automa.site/) also has a lot of potential.

I would recommend, if you do decide to go the self-hosted route, you do a little research on webhooks and HTTP headers. Both will be more important in your journey, as you may need to integrate an application that none of these applications currently support. But if you know your way around APIs and webhooks, you can pull in things from lots of different applications, expanding the functionality of these tools far beyond their base use cases.

### Calendar-booking tools: Let’s stick with Calendly for now

Honestly, this was the category I felt had the most potential to be pretty cool, but after a terrible, hours-long process of attempting to install a particular tool, I am just not going to recommend anything here.

Which is too bad. Calendly is a concept that could use FOSS-based disruption, or at least a reasonable alternative. However, there just isn’t anything out there that I feel comfortable saying is worthy of your time.

There’s something that I would have liked to recommend, but I couldn’t get it correctly installed, I imagine that a lot of you won’t be able to do the same. Anyway, here’s what I’ve got:

### Self-hosted Calendly alternatives

- **Too hot:** [Cal.com](https://cal.com/) has a lot of potential to make a lot of people’s lives easier; it’s not only a useful tool, but it also appears to be fairly modern and offers a real alternative to the dance of trying to book a meeting. That means that it is by no means embarrassing to use when you’re trying to set up a meeting. But I’m not recommending it because, at least at this time, it fails at the very basic hurdle of being self-hostable, in part because it uses a very modern application stack that is very temperamental.
- **Outside the box:** If you’re a [NextCloud](https://midrange.tedium.co/issues/cloudy-with-a-chance-of-diy/) user, [there’s a decent appointment app](https://apps.nextcloud.com/apps/appointments) that you can add to the application that covers most of Calendly’s basic use cases. Granted, it’s nowhere near as good as Calendly and doesn’t seamlessly integrate into Zoom, but unlike Cal.com, you can install it without ripping out your hair.

There’s something to be said for open principles and for carrying your application like it’s open in spirit, and something else to be said for the rubber meeting the road.

_(_[_Cal.com_](https://cal.com/)_)_

Cal.com, unfortunately, was stuck hovering above the road. I tried three separate methods to get it to install Cal.com, including my preferred Docker setup and its recommended [Railway](https://railway.app/) method, and it gave me issues at every turn. (Whatever the technical reasons, changing your domain should not be as big a hurdle as it is.)

Their tool may just not be ready for prime time for self-hosters—though if you want _them_ to host it, it might still be a good choice. But as this is a self-hosting guide, I have to call this out, unfortunately.

Suggesting an option like NextCloud, which is essentially a huge suite of tools, might be way too much for most people. But it’s a lot closer than Cal.com at this time, unfortunately.

**So, what am I to make about the world of self-hosting applications?** I think, like FOSS in general, you can tell where a lot of the open-source energy was put into … and where it wasn’t.

I don’t mean that as a slight on project creators. Ultimately developers have different priorities than end users.

Just as an example: I think it would be very useful if there was an open-source competitor to Buffer or Hootsuite that could allow small creators to manage their social media calendars without having to pay a middleman hundreds of dollars a year to do so. I’ve been hunting for a tool like this, and there appears to be one in the works, called [Mixpost](https://mixpost.app/), which works on the PHP framework Laravel. I have yet to test it, but this is promising. Their model makes sense, too—the very basic features are free under an MIT license, while more bells-and-whistles options will be sold under a one-time commercial license.

There is a class of small-scale creator that wants to be able to create and build their own infrastructure without having to spend hundreds of dollars a year on multiple bills for small tools.

But the thing is, this has been an issue for, like, a decade. And while there have been attempts of this nature in the past, only now are we starting to see that we need more FOSS alternatives to SaaS tools.

We have to make room for business models, sure—and I would encourage you, after you find a tool you like, to financially support creators you think are doing things well.

But we need to find a way out of this trap of SaaS. It is yet another manifestation of the platform issue, and an unnecessary one at that at the scale many small creators build.

I shouldn’t have to get a surprise bill for someone to point it out.






