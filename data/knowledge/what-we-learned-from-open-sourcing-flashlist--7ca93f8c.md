---
title: "What We Learned from Open-Sourcing FlashList"
notion_id: 7ca93f8c-97dc-4e14-856d-1c5346ad965c
notion_url: https://app.notion.com/p/What-We-Learned-from-Open-Sourcing-FlashList-7ca93f8c97dc4e14856d1c5346ad965c
last_edited: 2023-03-30T14:08:00.000Z
source_url: https://shopify.engineering/what-we-learned-from-open-sourcing-flashlist
tags: ["Article", "Shopify Engineering", "English", "Programming", "Project Management", "Product Management"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This blog post is for you—your project, your brilliant idea that is too good to be kept to yourself. Do you want to open-source your project but don't know where to start? In this post I'm going to explain what we learned by open-sourcing FlashList and how you can use it to launch your project.

[FlashList is Shopify’s React Native List library](https://shopify.github.io/flash-list/). This library solved a long-standing problem in the React Native community: bad list performance, by introducing high performance rendering, and removing problems like blank spacing and low fps on some devices. For Shopify, FlashList is really important, as most of our apps are based on lists, and performance is one of the aspects we care about the most.

In this article you will see how we came up from the initial idea, defined a launch strategy, and how we executed this strategy achieving 1,000 Github stars in less than 24 hours. I will explain how important it is to empathize with the end user, how to build authority, and end up with a successful open-source library. Let’s start from the beginning.

## **Defining the Project**

There are two ways of starting a project: You want to do a project but you don’t know what it is going to solve, or because you want to solve a problem you already have.

The first way is difficult. First of all, you’ll need to explore possible problems you want to solve. Once you have identified a problem, you’ll need to go deep and understand it, do market research, find your target audience, etc.

However, if you start because you’re facing the problem, this will be much easier. You’re part of the target audience of your product. You understand the pains that your project can solve, and this will help you find a better solution.

That was our case. I work on the engineering team of the [ Shopify Mobile app](https://www.shopify.com/ca/mobile), which merchants use to manage their Shopify stores on-the-go. When we started migrating the app from native to React Native, we encountered a problem: Lists were not performant enough.

As the app was previously implemented using native technology, it was easy for us to compare. The React Native version of the lists was pretty bad. List items were not rendering while scrolling, or the scroll stopped suddenly. As you may already know, lists are a very common pattern on Mobile. In our app, lists are everywhere: products, orders, inventory, collections to name some ones. We needed to solve this, otherwise the migration to React Native was at risk.

### Finding the Solution

This is when you start experimenting. You code the simplest solution that solves your problem. Of course, it’s not something you can release yet, but it’s what we call a proof of concept. At Shopify, we call this period_ the prototype phase_. We spend some time proving that what we want to build is possible, and trying to answer all the possible problems we will encounter during _ the build phase_, when we’ll be implementing our chosen solution.

I encourage you to do the same with your project. Start trying crazy things at the beginning, you'll have time to clean the code later on.

In the FlashList case, this is how we came up with the idea. A team started prototyping one of the native screens on React Native, using the two most common alternatives for lists at the moment: FlatList and RecyclerListView. FlatList was not performant enough, but was easy to implement; and RecyclerListView was performant, but its API was difficult to understand. So, the idea was simple: can we combine the best of both, using the FlatList API, but gain RecyclerListView performance?

That’s how FlashList⚡️ was born.

### Assembling the Team

You may be thinking: “_Okay, that sounds so cool, but you work at Shopify so for sure you have lots of people to work on a project like this.” _I thought this way before joining the company. The open-source projects look so cool that it gives you the impression you need a big team to accomplish something like this.

In the case of FlashList, we were only four people working on the project: Marek Fořt, Elvira Burchik, Talha Naqvi, and me. So don't feel discouraged if you work on a small team or by yourself.

### Build the Thing

Once you have an idea that solves a problem, the most important aspect is the execution. Two different teams with the same idea could end up in very different places. As we had a small team, we kept the execution fairly simple, and we had a working prototype quite early. At that moment we thought: “_Why keep this only for us and not open-source it?”_

An important note here: if you want to succeed during the execution phase you should be willing to spend time where others won't. Like documentation. We love projects with good documentation, but we hate to write it. Spend a good amount of time on documentation and you'll remove friction from people starting to use your project.

Once you have something meaningful to share, you’ll need to spend time thinking about how you’ll launch it—which leads to the next point.

## **Deciding to Open-Source**

We started planning the launch six months in advance when FlashList was not even close to being finished. We knew we had something worth open-sourcing so we started by building a [ very basic landing page](https://shopify.github.io/flash-list/). The plan—though some of it changed along the way⁠—comprised the following:

### **Landing Page**

You’ll use this to explain all these points: The problem your project is solving, how you are solving it, and how this will help whoever is reading the page. You may also want to think about the goals you want to achieve with your landing page. These were our goals:

- **Primary goal**: Know the features of FlashList compared to its predecessor, FlatList.
- **Secondary goal**: Try out FlashList by installing it in your project.
- **Tertiary goal**: Read the documentation to learn everything about it.

### **Blog Post**

We wanted to explain the story behind what we built and how we built it. You can read [the blog post we created here](https://shopify.engineering/instant-performance-upgrade-flatlist-flashlist). Use a blog post to explain the story to your project. Storytelling is the best tool to sell anything. People will appreciate your work more if they understand what you did, and why you did it.

### **Social Media**

Our plan was to launch on Twitter and LinkedIn whenever the project was ready. At this moment, we weren’t writing anything specific, but we knew social media was necessary for the launch. Sites like Twitter can help you amplify your launch, and if what you are releasing is relevant, it will be retweeted multiple times.

## **Building Buzz Before the Launch**

Now, you have a plan and a working prototype. But, when is a good moment to launch? We weren’t sure when to call it v1.0. We were working on making the library stable, and our apps started using FlashList in production on some of their screens. We started with Shopify Mobile, and then other apps like [Shop](https://shopify.engineering/building-app-clip-react-native) and [ Inbox ](https://shopify.engineering/real-time-buyer-signal-data-pipeline-shopify-inbox) started to use FlashList as well.

But even with all that, it’s hard to see the moment where you can freely say: We are done. Why is that? Because launching a project is SCARY AS HELL. What if we fail? What if nobody cares? What if they care but it's not good enough?

I also have to say that the Shopify brand has a high standard for open-source projects, and we wanted to release something that was at the same level.

Then something unexpected happened.

### **Do Things, Tell People**

Talking about FlashList at conferences was a crucial part of its success, but we weren't originally considering live events as part of the initial launch. [ In June, FlashList was mentioned in an App.js conference talk](https://youtu.be/l2-nR_neb9Y?t=643). There, Monica Restrepo presented our open-source contributions at Shopify. Some team members knew FlashList would be included, and some didn’t. Moreover, some of the FlashList creators attended that conference in-person, watching as their project was explained by somebody outside the team. That was the thing that ripped off the band-aid.

After the talk, developers from App.js started showing interest on Twitter. That was the first good sign that we were onto something cool. As we already knew, FlatList was a big pain for the React Native community and people were desperate for solutions. Here are a few tweets that inspired us to make the next move:

### Build a Community

With buzz building on social media, we started our launch strategy. We decided that having people from Twitter join an invitation-only, closed Discord community would be a good move, and [ tweeted about it](https://twitter.com/ImRestrepo/status/1537061094220615682). We invited people by DM, one by one. We got over 100 people to join. FlashList wasn’t ready, but we knew we had to use this opportunity to create hype around the library. We started showing sneak peeks—videos, images of performance graphs. This gave us a bit more time until we had everything ready to open. But people were impatient, and the memes arrived:

We needed to move quickly or we'd lose momentum.

## **Launch Day and Beyond**

Our initial plan was to publish the [first blog post](https://shopify.engineering/instant-performance-upgrade-flatlist-flashlist) on the day of the launch, but it wasn’t ready yet. If we wanted to keep people engaged we couldn’t wait until this was ready. Even though launching the blog the same day would have helped people better understand the story behind FlashList, we decided to move forward with the launch and publish the blog post later on.

The value proposition was clear with the tweet and the landing, so people immediately engaged even without the full context. The engagement on the [ tweet that Shopify Engineering posted](https://twitter.com/ShopifyEng/status/1542538720276086786) was nuts: More than 1,200 likes and over 300 replies. It rapidly became the tweet with the most engagement since last year’s Black Friday–one of the top moments for Shopify of the year.

On LinkedIn, the announcement post became the most engaged post of all time from the Shopify Engineering account. We were also #2 on the GitHub trending list thanks to all those stars:

I'd say the key for success was having a great library, our preparation, the hype the conference generated around the list, the hunger from the community for a better library, and a bit of good luck⁠.

Whenever you decide to launch, know that it could take off quickly, or that the growth can happen slowly. Don’t be discouraged or think you'll only have one opportunity for the launch. Keep working and understanding your user pains, and success will come eventually.

### Empathize With the Users

Empathizing with users is important from the very beginning, and doesn’t stop after the launch. We created a project to solve a problem we were also experiencing, so we had the added benefit of also being users. If this isn’t true for your project, you’ll need to find more ways to listen and engage with your users.

Remember to emphasize: Think of WHY your product matters and how it helps users with their problem.

Also, if you are passionate about the project, people are going to notice and be passionate about it as well. Talking about FlashList cannot be more of an honor for me, it's one of the best projects I've been on and I hope this gets conveyed to you while reading this.

### **Look for Social Proof**

When you start a project, you have zero credibility. Nobody believes in it. With FlashList, even people inside the company were not optimistic about it. Developers used FlatList because it was the option everyone else was using. Even if it was not the best, it had authority.

I'd say that working at Shopify helps with credibility. Launching something under a brand that is known in the community gives you authority, but also adds pressure to release something that is good enough to be under this name. If you’re an independent developer, you have an advantage as well, as you don't have some high bar to reach and you can launch without pressure.

So think about this question: “_Why should people care about your product?” _Why should people care about FlashList? Our key selling point was performance. But we needed credibility and social proof for people to really believe it. It’s the same as Apple saying their latest Macbook Air performance is top-notch. Until you have Marques Brownlee reviewing the product and giving their personal opinion, you don't know how much of the hype is real.

To solve that you'll need to put your project in the hands of other people. [ This tweet](https://twitter.com/almouro/status/1544669919479996416?ref_src=twsrc%5Etfw%7Ctwcamp%5Etweetembed%7Ctwterm%5E1544669919479996416%7Ctwgr%5E8c0be28dc43e665c6483a6e548e8059020491273%7Ctwcon%5Es1_c10&ref_url=https%3A%2F%2Fshopify.github.io%2Fflash-list%2F) is an example. Alexandre has [ an open-source Flipper plugin to measure performance](https://github.com/bamlab/react-native-flipper-performance-monitor). He tried FlashList and posted the results. That gives you a lot of points.

## **7 Key Things We Learned**

Even if every open-source project is different, there are some general principles that I recommend to anyone thinking of open-sourcing a project:

1. **Coding a library is not everything.** If you want to create a popular open-source library, you need to think about the other aspects I mentioned in the previous section: documentation, promotion, and thinking about the library as a product.
2. **Have a plan.** It doesn't need to be exactly the same as our FlashList plan, but I suggest you follow it as much as possible. The plan will change over time, but you’ll have a direction to follow and a to-do list to complete.
3. **Don't do it all by yourself.** Get your project in the hands of other people as soon as possible. Get external feedback and collaborate with others if possible. You may be biased with it and having more pairs of eyes will help the project grow.
4. **There's never a perfect moment.** There won't be a perfect time to launch. You will always have things that can be improved. But as soon as you consider when and how other people could benefit from your product, that is a good enough time to do it.
5. **Empathize with the user.** Zoom out. Don't think only about yourself or just the tech details of your library. Think about the benefits others could take from your library and document its benefits this way. You can also ask peers outside the project about the value they see in your library.
6. **Create a community.** While creating the closed Discord community, we noticed that giving people sneak peeks in private created a closer relationship with the library. They felt like early adopters, some of them even promoted it without us having to ask.
7. **Build authority with a long-term strategy.** Don’t think only about the first launch. There will be several moments when you can relaunch your product. After the first announcement, you can publish blog posts, technical deep dives, videos about it, and go to conferences. This is an example of a follow-up blog post, and you can see we’ve been in different conferences lately talking about FlashList.

## **Resources**

- [Instant Performance Upgrade: From FlatList to FlashList](https://shopify.engineering/instant-performance-upgrade-flatlist-flashlist)
- [FlashList Landing page](https://shopify.github.io/flash-list/)
- [Shopify React Native Open Source Discord community](https://discord.gg/NqVZ9YA9)
- [How Contributing to OS is Helping Shopify & RN to Build the Future of Mobile](https://www.youtube.com/watch?v=l2-nR_neb9Y)

**David Cortés** is a Development Manager on the React Native Foundations and Mobile Polaris team at Shopify. Connect with David on [Github](https://github.com/davebcn87) and [Twitter](https://twitter.com/davebcn87).

Wherever you are, your next journey starts here! If building systems from the ground up to solve real-world problems interests you, our Engineering blog has stories about other challenges we have encountered. Intrigued? Visit our [Engineering career page](http://www.shopify.com/careers/specialties/engineering?itcat=EngBlog&itterm=Post) to find out about our open positions and learn about [Digital by Design](https://www.shopify.com/careers/work-anywhere?itcat=EngBlog&itterm=CCTA-DD).


