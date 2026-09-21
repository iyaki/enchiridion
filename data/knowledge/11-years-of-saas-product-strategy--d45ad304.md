---
title: "11 years of SaaS product strategy"
notion_id: d45ad304-7da0-4256-be32-0a831faad55d
notion_url: https://app.notion.com/p/11-years-of-SaaS-product-strategy-d45ad3047da04256be320a831faad55d
last_edited: 2023-10-12T21:39:00.000Z
source_url: https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy
tags: ["Article", "Alex Ghiculescu's Newsletter", "English", "Product Management", "Entrepreneurship"]
---
In the last year at [Tanda](https://www.tanda.co/), we’ve gone from being a single-product company, to launching a bunch of new products that all work together really well.

This has got me reflecting on the different phases of product strategy over the last decade, as we went from no product to a market leader.

If you find this interesting you might also enjoy [_11 years of hosting a SaaS_](https://ghiculescu.substack.com/p/11-years-of-hosting-a-saas).

## Phase 0: The roadmap is just an idea

This is the phase that all products, projects, and companies start at. There is nothing built, there is only an idea. In this phase all you should do is build enough of the idea that you can present it to people who’d use or buy it and see what they think. Startup people call that an MVP.

Planning at this stage is actually quite useful, but it looks very different to planning later on. Anyone experienced with product planning at a mature company will need to unlearn a lot to succeed here. That’s because “planning” here is writing 2 sentence summaries of what needs to get built.

> We are going to make an employee time clock that runs on an Android device. It will connect to a website in the cloud so you don’t need any dedicated IT support.

Some people say you should start with a market, and then tailor your product to that. Go out and interview 100 similar potential buyers, figure out something they want, and build exactly that. In my experience, this advice is often given and (correctly) never followed. If you don’t already know about something that people want that you can explain in 2 sentences, then you shouldn’t be starting a business.

If you insist on the customer research approach, at least work on your script. When wannabe founders ask me if they can do “a customer development interview to identify potential pain points and riff on product ideas”, all I hear is that the call will be a waste of my time.

## Phase 1: It’s obvious what to build next

You’ve built the first two sentences of your product, and managed to get some people to use it (hopefully to pay for it too). At this stage it can seem pretty obvious what to build next, because you have way less features than what customers expect. You’re on the way to product market fit.

For Tanda, this was the case for award interpretation, rosters, the early, simple, versions of timesheets and leave, and payroll exports. ([2](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-2-136210108))

Why were they obvious? Obvious to who? Here’s how I think about it. When talking to a customer or prospect, if instead of asking “does it do X?”, they ask “how do you do X?”, and you can’t confidently answer “oh, you _don’t_ do X, and here’s why”, then that is an obvious gap.

1) The customer assumed the functionality is there, so it’s an obvious requirement to them.

2) You don’t have a justification for it not being there ([3](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-3-136210108)), so implicitly it’s obvious to you.

When it’s obvious what to build next, product planning is easy. Pick one obvious thing, build the simplest version of it that you can, and repeat. Steer clear of “prioritising”; if you need to prioritise things then implicitly admitting that it’s not obvious what needs to get built, and during this phase you should only build truly obvious things. Pick whichever obvious thing you got asked about most recently, build it, and repeat.

Seasoned product managers will warn you that if you do this too much you will end up building all sorts of random disparate features that don’t work well together and are not consistent. They’re not wrong, but they miss the part where if you don’t do that you’ll go broke. During this phase, you should build lots of things, accept some of them will suck and attract bad customers, and use this intel to identify who your ideal customers are. Then gradually remove the features, and customers, you don’t want.

While prioritisation is a bad thing during this phase, scoping is not. It’s tempting to get carried away by making a really sophisticated version of your obvious need. Either one that does too many niche things, or one that is too polished. This is where the saying “[if you’re not embarrassed by your v1, you shipped too late](https://twitter.com/reidhoffman/status/847142924240379904?lang=en)” comes from; the first version of every obvious feature should look a bit ugly, and should only have functionality that 100% of customers will use exactly the same way. All the other stuff, you can come back to later. It’s tempting to plan ahead because you’re “sure” you’ll need some functionality soon. Resist this temptation; even if “soon” actually comes soon (it often doesn’t!) it’s still better to aggressively restrict your scope and add the extra bits later.

During this phase you will get asked for lots of things that are not actually obvious needs. For us, a mobile app for employees was the most classic example. We didn’t have one on the app store for the first 4 years of Tanda. We got asked “how do employees install the app?” all the time. Our answer: there’s a mobile website ([4](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-4-136210108)), and everything important also happens over email and SMS ([5](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-5-136210108)). So you see - we’d explain to customers - you don’t _need_ an app, but what you do need is a really accurate award interpreter, and only we have that.

We didn’t say this because we were fundamentally opposed to native apps. Eventually we built one. But for a long time, it just wasn’t as important as other things we could have built with our really small team.

## Phase 2: It’s really hard to know what to build next

All of a sudden you’ll find that all the really obvious stuff has been built. Suddenly prioritisation goes from being really easy, to really hard. You’re drowning in feature requests, bug reports, suggestions (that are phrased as requests), complaints about technical debt, and a million other inputs. And while most of them don’t seem like bad ideas, none of them seem critical either.

Now you have to actually prioritise.

I found this transition really difficult. I’d spent years building whatever felt like a good idea and being right more often than not. Now everyone was saying we couldn’t keep doing that and we needed to create a lot of new process to protect ourselves from us building the wrong thing. In hindsight, I should have been more skeptical of the sudden need for drastic change. But I wasn’t, and so we lent into transforming how we did product strategy.

We got a lot of things wrong over many years while trying to find the right of balancing our intrinsic scrappiness with the new and constantly growing number of customers that relied on us. Here’s some things that didn’t work:

**Succumbing to pressure from big customers**

Big customers will always ask you to build something that everyone knows only they will use. A million people have written the advice that if 1 customer dictates your roadmap then you need to say no. It’s still hard to resist, the first time you have that opportunity.

I thought the reason you aren’t meant to do this is because then you’ll never get off the treadmill of building whatever that customer wants. We were on that treadmill for years with our first enterprise customer, but it actually wasn’t that hard to step off once we really committed to do it. We got very firm about explaining how we’d like to work (citing [a book](https://basecamp.com/shapeup) helped make it more official), and they respected that and let us work our way.

The real problem is that when you are building something a single customer is pressuring you for, it’s incredibly hard to build it in a way that different customers will be able to use. ([6](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-6-136210108))

Before we learned this, we built a few features for big important customers, and those features work really well if you’re the same size, and in the same industry, as those customers. For anyone else, they’re too confusing to bother. It took a long time to internalise this and properly try to make them work for all customers. We still haven’t perfected that but it’s a lot better than when we started.

Scarred by getting this wrong a few times, we made a big decision: we won’t build features if only a single customer wants them. This sounded good on paper. In practice it resulted in…

**“Prioritisation” by sorting by MRR**

This is a phenomenon when any internal proposal for a feature (sometimes bug reports too!) includes a list of prospects or customers that have suggested the feature, and their corresponding revenue. Prioritisation becomes picking whichever item has the biggest revenue associated with it.

The implication is “if we build this feature we’ll win all these new customers” or “if we don’t build this feature we’ll lose all these customers”. It’s never true. It’s very rare to see someone buy or churn off a business software product exclusively because of one feature’s existence. If that’s happening often, and you want those customers, you don’t need prioritisation!

This methodology is one of those things where there’s enough merit to it that it sounds reasonable in theory, but actually is a total mess in practice. Salespeople and account managers can be very convincing, and it’s really hard to resist the allure of just building what they ask. At the same time, it’s bad to not listen to them at all. To make good product decisions when it’s not obvious what to build next, you need to walk the tightrope between those extremes.

**Forced innovation**

Worn down by MRR-based-development, a few teams drew a line in the sand. “From now on, we’re only going to bet on big, chunky, innovative things.” This was very popular internally. Developers love working on “big” projects. ([7](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-7-136210108))

Everyone loves innovative new things. What’s not to like?

The problem with forcing innovation is that it assumes there is an endless supply of game changing ideas floating around your company, and the only reason you haven’t built them yet is because the product team didn’t want to. If this is true for your company, you’re on the verge of being obscenely rich. But for most companies this isn’t the case. [You don’t make good products because you really want to, you make good products by fostering the conditions in which great products can be made](https://stratechery.com/2017/amazons-operating-system/). An environment where every 6 weeks teams pick which big new innovation they’ll be launching is not one that invents anything new.

That’s what we learned, from trying to force this. We saw lots of complex features with fancy names come from this, but nothing that substantially changed the lives or fortunes of our customers.

**Data-driven development**

Like most product teams, we spent a lot of time buying and setting up product analytics tools like [Amplitude](https://amplitude.com/). Like most product teams, we spent virtually no time using the data inside Amplitude to make good decisions.

Amplitude is a cool tool and it’s well built. I’m sure there are some kinds of product, and some kinds of company, where it’s really useful for finding diamonds of insight. But in our B2B SaaS world, it only uncovered really obvious things we already knew (“lots of people log in every day”, “everyone uses user profiles regularly”).

Same goes for recording tools like [FullStory](https://www.fullstory.com/). We were able to uncover a few bugs that were hard to replicate otherwise - and I’m glad we did - but we did not discover brilliant untapped product insights from watching screen recordings of our settings pages.

All these tools are a bad proxy for talking to customers. ([8](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-8-136210108))

## Phase 3: Getting aligned

This is where we’re at today. As we’ve grown as a company, we’ve found that the biggest impact we get when building features comes when the rest of the company can also align its work with that feature.

This sounds obvious, but it’s a pretty big change from the early days of having the feature idea over breakfast, shipping it after dinner, and hoping your customers discover it.

For example, a few years ago, we added a bunch of improvements to how we handle different kinds of lunch breaks in the shifts people work. To market this, we sent all our prospects a [Kit Kat](https://www.creativemoment.co/creative-classic-how-and-why-have-a-break-have-a-kit-kat-has-lasted-60-years) in the mail, as well as a flyer about what we’d built. Mailing people stuff is a surprisingly big logistical challenge, but it worked really well to draw attention to what we were doing and get people talking about it. It also got the whole company talking about the same thing. Between designing and building the features and putting Kit Kats in boxes, everyone was aware of what was going on and could make sure they understood the new features well so they could explain them to customers.

I’ve found there’s two cadences that really matter, and if you can get product happening on those cadences then things will work well.

- If you do your product thinking on the same cadence as that every other team prioritises work on, you can work together better.
- If you do your product thinking on a cadence that matches how customers will deploy new features, they’ll be more comfortable with your pace of deployment.

In our case it turned out the answer to both of these was every 3 months. We use [OKR](https://en.wikipedia.org/wiki/Objectives_and_key_results)s across the entire company. Every 3 months the OKRs get refreshed. About a month before that, we pick what the product theme for the next quarter will be. OKRs flow down from that, and the individual pieces of work get shaped from that.

This means every team knows broadly what areas we are thinking about, where to expect improvement, and is able to align it with their work.

This is where we are at today, I’m sure in the future I’ll write an update about all the stuff that we got wrong here.

## Phase _n_: adding another product

We’re building and launching new products at the moment, which has got me reflecting on this, because those new products are still in Phase 1 (or 0!). It’s so obvious what to build next because there’s still lots of stuff missing! But as a result, we’re getting it done really quickly, knowing we’ll get some things wrong and have to rebuild some stuff along the way.

I wish I’d thought more about _when_ to add a second product a long time ago. It’s a bit like having kids: most people assume they’ll do it one day, but want to wait until they’re ready. If you wait for an email telling you you’re ready, you’ll never do it. So it is with many things, and I found that once we started building a second product it was not as hard as expected, and much more rewarding. I was most surprised by how much happier customers who buy more products seem to be, which I think has really energised everyone.

Should we have done it earlier? I think so. It would have been harder, we would have been less ready, we’d have gotten more things wrong, and we’d have been more stretched between products - but there’s a lot of sunk cost fallacy to that line of thinking.

More usefully: should someone who’s just moving from Phase 1 to Phase 2 consider adding a second product as an alternative on going deeper on the existing one? Absolutely you should at least consider it, even if you decide not to after weighing it up.

## If I had a time machine

If I had a time machine to go back to 2012 and give myself a few pointers, what would I say? 3 big tips:

**Be less afraid of adding new products/markets**. I was always really reluctant to consider new products, or new countries (which is basically a new product in our world). It seemed like a lot of hard work that would distract from focusing on the core product. I underestimated how much it would energise the team and increase focus.

**Change/remove features if they aren’t working**. There’s a lot of product management advice that boils down to “say no to feature ideas”. This is easy to prescribe, and hard to do. But I wish I’d gotten better at changing features, or scrapping them, if they weren’t working well for the majority of customers or were clearly too confusing.

**The best product strategy is to build what your customers say they’ll buy**. As cliche as it is, there is always temptation to come up with other sources of inspiration for the right thing to build. These should be resisted. If you build what your customers are asking for you’ll build too much stuff, it won’t work together properly, and it’ll be confusing in your UI. But if you build stuff your customers aren’t asking for, you won’t be troubled by having customers for much longer.

Thanks for reading Alex Ghiculescu's Newsletter! Subscribe for free to receive new posts and learn from more of my mistakes.

[1](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-anchor-1-136210108) We’ve found that customers that buy multiple products tend to be much happier, but that’s another story.

[2](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-anchor-2-136210108) If you don’t live in the workforce management world you might not know what all these features are. In terms of how fundamental they are, they’re on par with building Github without an online code browser.

[3](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-anchor-3-136210108) “We haven’t built it yet” is a great (and the most likely!) explanation, but not a great justification.

[4](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-anchor-4-136210108) try explaining that to boomers

[5](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-anchor-5-136210108) Luckily, boomers loved this. It turns out lots of people [outside the bubble](https://ghiculescu.substack.com/p/software-companies-that-sell-software) struggle to install apps. Even as customers were requesting an app they were also dreading having to help all their staff install the app and log in!

[6](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-anchor-6-136210108) There’s a really similar challenge if your product deals with any regulation and you’re across more than one country. It’s easy to tell yourself that you can build features that will work in every country, and hard to actually do that.

[7](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-anchor-7-136210108) “Big” often means [greenfield](https://ghiculescu.substack.com/p/the-greenfield-project-trap), even if you didn’t intend it to.

[8](https://ghiculescu.substack.com/p/11-years-of-saas-product-strategy?utm_source=tldrfounders#footnote-anchor-8-136210108) I’m less bearish on [Canny](https://canny.io/), but we’ve also struggled to get it right. At least with Canny you’re kind of talking directly to customers.
