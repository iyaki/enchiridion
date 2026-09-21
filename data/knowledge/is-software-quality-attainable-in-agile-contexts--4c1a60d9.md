---
title: "Is software quality attainable in Agile contexts?"
notion_id: 4c1a60d9-ebeb-4dd0-bb9e-7bc119affcd8
notion_url: https://app.notion.com/p/Is-software-quality-attainable-in-Agile-contexts-4c1a60d9ebeb4dd0bb9e7bc119affcd8
last_edited: 2023-08-04T18:22:00.000Z
source_url: https://alediaferia.com/2021/12/27/what-does-software-quality-mean-in-agile/
tags: ["Alessandro Diaferia", "English", "Agile", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Every so often I read a blog post outlining how Agile has got it all wrong and that if you really want high-quality software you have to ditch that methodology in favour of the good old planning and design cycles upfront.

I must admit, this type of article gets me thinking every time. The last one I read prompted me to write this post as a way for me to rethink the concept of Software Quality through Agile eyes, if you will.

I’ve been working in Agile tech contexts through out most of my career and what I’ve heard most often from the detractors of Agile is that you can hardly have _good_ quality and _proper_ architecture with this methodology.

I think there’s a lot to unpack here to give a proper response to this claim so I figured I’d write a blog post focused on the Quality aspect.

# What is Software Quality?

Let’s first define what Quality means. As per [Wikipedia](https://en.m.wikipedia.org/wiki/Software_Quality):

> 

– Software functional quality reflects how well it complies with or conforms to a given design, based on functional requirements or specifications. […]

– Software structural quality refers to how it meets non-functional requirements that support the delivery of the functional requirements, such as robustness or maintainability. It has a lot more to do with the degree to which the software works as needed.

What I get from the definition here is that you can’t have quality without a set of requirements to adhere to. Now, as a detractor, you might argue that you’ve hardly seen extensive specifications in Agile and therefore it’s a natural consequence for _agile_ software not to be of quality.

The reality is that, in Agile, we have _just enough_ requirements and specifications to fulfil the current iteration.

> 

Working software is the primary measure of progress.

[Agile Manifesto](https://agilemanifesto.org/principles.html)

Therefore, strictly speaking, if the software adheres to the definition of that specific iteration then it is of quality.

If you come from non-Agile environments very light requirements in specs might throw you off. The truth **is this is foundational to Agile**:

> 

Simplicity–the art of maximizing the amount of work not done–is essential.

[Agile Manifesto](https://agilemanifesto.org/principles.html)

Agile has no desire to develop any verbose specification upfront: it acknowledges that **in order to build working software, embracing change and iteration is essential**. This makes it quite hard to think about quality in _traditional _terms.

[Agile has no desire to develop any verbose specification upfront: it acknowledges that in order to build working software, embracing change and iteration is essential. ](https://twitter.com/intent/tweet?url=https%3A%2F%2Falediaferia.com%2F2021%2F12%2F27%2Fwhat-does-software-quality-mean-in-agile%2F&text=Agile%20has%20no%20desire%20to%20develop%20any%20verbose%20specification%20upfront%3A%20it%20acknowledges%20that%20in%20order%20to%20build%20working%20software%2C%20embracing%20change%20and%20iteration%20is%20essential.&via=alediaferia&related=alediaferia)[Click To Tweet](https://twitter.com/intent/tweet?url=https%3A%2F%2Falediaferia.com%2F2021%2F12%2F27%2Fwhat-does-software-quality-mean-in-agile%2F&text=Agile%20has%20no%20desire%20to%20develop%20any%20verbose%20specification%20upfront%3A%20it%20acknowledges%20that%20in%20order%20to%20build%20working%20software%2C%20embracing%20change%20and%20iteration%20is%20essential.&via=alediaferia&related=alediaferia)

# Do iterations mean rework?

If you come from a waterfall-like mindset you might classify this iterative approach as _rework_: you might think that light specifications and poor planning have led to this many iterations. Whereas, if you had put enough effort into upfront planning, design and architecture, you’d have had only one development iteration. By now you’d be able to move on to the next project.

Yes, it might happen that if you are extremely lucky and you do a lot of thinking upfront you might come out of a 6 months development phase with a product that is right at that point in time. And that users want. Or, you are probably just optimising a solution for a problem that’s already been solved. The rest of the times lack of feedback during development will make you build something that doesn’t meet your users needs fully.

In Agile, frequent iterations on the same software area are expected, even sought after.

# Just enough quality is what you need

Now, I know that feeling. You have stuck with Agile. You have built and shipped something according to those slim requirements but you already (_think you_) know what’s going to scale and what isn’t. What corner cases you haven’t handled. What is going to happen if your users were to push your software to its limits.

## Resist the urge of anticipating work

**Is an unscalable solution enough to help you validate what your users want?**

If the answer is yes (and it is most of the times) you should not worry about putting an unscalable solution out there.

I’m sure you are opinionated about what is going to scale and what isn’t. Real production usage, though, might inform otherwise or give you a new perspective.

You might have some ideas in mind on what non-functional aspects you want to tackle next: make it more resilient and scalable, for example. If this is the case, I think you should either dismiss them or put them into your backlog. And if you _must_ keep them in your backlog be prepared to revisit them as you gather more usage data.

You’ll probably discover that the way you’re solving your users’ problem is not helping anyone. In that case the whole iteration might go into the bin, together with your ideas on how to make it **scalable** and **resilient**.

On the other hand, what if you have picked the right path and suddenly thousands of users flocked to start using your unscalable solution? First of all, I must tell you, this is an extremely rare possibility. But yes, it does happen from time to time.

If it happens you are going to have an actual production use case to architect your non-functional requirements upon. You’ll be able to base scalability decision upon real usage data. Your software will end up being more solid than you could have anticipated. Plus, you’ll have the confidence that you are building the right solution! **This is a considerable advantage that goes beyond just software engineering. It might be crucial for the company as a whole.**

## Investing in_ quality_ _upfront_ is a gift to your competitors

I’ve read many posts advocating that cycles spent on design and planning upfront will help you embed quality into your software. But how can you embed quality into something that you haven’t validated yet?

Agile methodologies try to move the focus away from endless specification in favour of frequent validation cycles. The purpose here is to make a bet, validate, and iterate to minimise wasted effort.

This means that all those characteristics you feel should be part of your software as you are building it will naturally emerge from iterating and validating if they make sense.

Remember, you want to _**maximise the amount of work not done**_.

Naturally, understanding what the minimum amount of work required before putting your software in front of your (potential?) customers varies according to many factors. From experience I can tell you it’s usually less than what you think it should be.

Nevertheless, I understand that putting garbage in the hands of your customers might not help anyone. At the same time, investing in making your software fully resilient before acquiring your first customer might be the difference between you stuck in development and your competitor winning their Nth customer.

[Making your software fully resilient before acquiring your first customer might be the difference between you stuck in development and your competitor winning their Nth customer ](https://twitter.com/intent/tweet?url=https%3A%2F%2Falediaferia.com%2F2021%2F12%2F27%2Fwhat-does-software-quality-mean-in-agile%2F&text=Making%20your%20software%20fully%20resilient%20before%20acquiring%20your%20first%20customer%20might%20be%20the%20difference%20between%20you%20stuck%20in%20development%20and%20your%20competitor%20winning%20their%20Nth%20customer&via=alediaferia&related=alediaferia)[Click To Tweet](https://twitter.com/intent/tweet?url=https%3A%2F%2Falediaferia.com%2F2021%2F12%2F27%2Fwhat-does-software-quality-mean-in-agile%2F&text=Making%20your%20software%20fully%20resilient%20before%20acquiring%20your%20first%20customer%20might%20be%20the%20difference%20between%20you%20stuck%20in%20development%20and%20your%20competitor%20winning%20their%20Nth%20customer&via=alediaferia&related=alediaferia)

I won’t deny it’s a fine balance to find, though. You have to focus just enough on non-functional requirements so that you can keep iterating on what the users see while solidifying the foundations you build upon.

# Endless upfront specifications? A false feeling of Quality

From my experience I see some of us (software engineers) don’t particularly enjoy the discovery aspect of working in an Agile context. This is something I’ve personally struggled with for a bit at the beginning of my career. **All I cared about was the act of solving a problem** in the most _elegant_ (according to who?) possible way. And it is something I noticed with some colleagues too, being them junior or highly experienced senior engineers.

## I just want to_ do my job_

In these situations a big list of highly detailed requirements might feel like a great thing. There’s a bunch of problems you have to figure out how to solve and you can do that with code. **It’ll be like a playground where you can experiment and hammer at each requirement like a box-ticking exercise until you’re done**.

There’s a great feeling of progress involved. You know exactly where you are and how much is left. Every requirement is like a challenge. There’s very little uncertainty for you in this situation. This feeling is great as human beings. [We don’t particularly enjoy uncertainty](https://www.universityofcalifornia.edu/news/science-what-uncertainty-can-mean-your-mind-and-body) after all.

Maybe you don’t even care about whether your software is going to be used or useful for your company to acquire your next customer. You just care about _doing your job._

It must feel great to complete the development phase having implemented every single requirement, to the exact level of detail in the specification. As a developer you think you have done your job and you can move on.

## Meeting customer demand is also your responsibility

Fortunately, Agile tries to open our eyes to the crude reality out there: you had no idea how your software was going to be used. Turns out your target users can’t get any value out of it. You have just finished implementing a perfectly engineered high-quality piece of garbage that nobody is willing to use because it just makes no sense.

That being said, I don’t think there’s isn’t a place for you in Agile if you don’t want to take part into the discovery process.

You’ll have to accept you are going to work with a slim set of requirements, and you might end up working on the same area of the software across multiple cycles. That’s a good thing! It means your organization is betting on the right things and it’s receiving feedback that they indeed are the right things for your customers.

# Software Quality is emergent

Building software in Agile is a discovery process. Users gradually discover what they want. You work with them to discover how to meet their demand and solve their problems. Additionally, all of this happens while the environment around changes and influences the problems you are trying to solve.

[Building software in Agile is a discovery process. Users gradually discover what they want. You work with them to discover how to meet their demand and solve their problems. All of this happens while the environment around changes and… ](https://twitter.com/intent/tweet?url=https%3A%2F%2Falediaferia.com%2F2021%2F12%2F27%2Fwhat-does-software-quality-mean-in-agile%2F&text=Building%20software%20in%20Agile%20is%20a%20discovery%20process.%20Users%20gradually%20discover%20what%20they%20want.%20You%20work%20with%20them%20to%20discover%20how%20to%20meet%20their%20demand%20and%20solve%20their%20problems.%20All%20of%20this%20happens%20while%20the%20environment%20around%20changes%20and%E2%80%A6&via=alediaferia&related=alediaferia)[Click To Tweet](https://twitter.com/intent/tweet?url=https%3A%2F%2Falediaferia.com%2F2021%2F12%2F27%2Fwhat-does-software-quality-mean-in-agile%2F&text=Building%20software%20in%20Agile%20is%20a%20discovery%20process.%20Users%20gradually%20discover%20what%20they%20want.%20You%20work%20with%20them%20to%20discover%20how%20to%20meet%20their%20demand%20and%20solve%20their%20problems.%20All%20of%20this%20happens%20while%20the%20environment%20around%20changes%20and%E2%80%A6&via=alediaferia&related=alediaferia)

Similarly to Architecture, Quality is a characteristic that gets refined rather than being something planned upfront.

> 

The best architectures, requirements, and designs emerge from self-organizing teams.

If you think light planning and multiple iterations represent **the** **failure of software engineering** think about how wasteful it would be to lock yourself up in a room to architect the highest quality software while your potential users and the environment around them keep changing their needs. By the time you come out of your room you might have produced something that’s no longer relevant.

# Listen for feedback if you want quality

We have explored how embracing Agile means embracing emergence of characteristics in a changing environment. One thing I haven’t stressed enough, though, is that you have to be prepared to receive those signals from the changing environment around you. If you ship an MVP how are you going to evaluate whether it’s solving a problem for your users or not?

## Upfront observability?

Be it customer interviews or instrumentation and observability, this is one of the things that you should invest in as an upfront effort. Agile is all about being able to understand the signals and make the next bet according to the signals from the previous iterations.

I see this aspect is probably one of the most neglected. In my experience it’s hard to stop and think about what kind of signals we’d want to look at to understand if what we’re building is being valuable or not. On the other hand, I think this is the most important thing you have to invest in upfront.

Luckily for us we have plenty of tools at our disposal that help us embed observability into our software. Analytics and metrics platforms usually require a little investment upfront to be able to produce meaningful data. As with everything else in Agile, you’ll discover how to better tune your signals to become more useful.

Do not wait until it’s too late. Embedding observability as an afterthought is a much harder effort: not only you’ll have to make code changes to activate observability, you’ll also have to change how you build software. You will have to change your practices in order for observability to become part of your act of producing software. As with any kind of habit this is going to be hard.

[Embedding observability as an afterthought is hard: it requires practice and mindset change, not just a code change. ](https://twitter.com/intent/tweet?url=https%3A%2F%2Falediaferia.com%2F2021%2F12%2F27%2Fwhat-does-software-quality-mean-in-agile%2F&text=Embedding%20observability%20as%20an%20afterthought%20is%20hard%3A%20it%20requires%20practice%20and%20mindset%20change%2C%20not%20just%20a%20code%20change.&via=alediaferia&related=alediaferia)[Click To Tweet](https://twitter.com/intent/tweet?url=https%3A%2F%2Falediaferia.com%2F2021%2F12%2F27%2Fwhat-does-software-quality-mean-in-agile%2F&text=Embedding%20observability%20as%20an%20afterthought%20is%20hard%3A%20it%20requires%20practice%20and%20mindset%20change%2C%20not%20just%20a%20code%20change.&via=alediaferia&related=alediaferia)

# When doesn’t Agile work?

I based this post on the assumption that we are employing Agile methodologies in a fast-paced and changing environment. I’ve assumed there is a high competitiveness degree between entities chasing market value through software products. In this context I find Agile helps maximising the pace and focus of your efforts around finding the shortest path to product-market-fit.

## Frequent experiments have intrinsic instability

I won’t deny that, in my experience, this approach leads to a higher degree of technological instability in the earlier iterations. Maximising all our efforts on getting early product feedback makes us neglect certain non-functional aspects like resilience, scalability, availability and more.

This happens because we are maximising our efforts on activities that give us a return in terms of confidence that there is appetite for what we are building.

We use techniques like A/B testing, blue-green deployments, feature flagging to make **experiments **that help us test our assumptions. This type of experiments are often bringing into production software that is not _production-quality_.

Hopefully, after having received the feedback we were looking for we will have enough data to make an informed decision about those non-functional requirements that we will invest in in our subsequent iterations.

## Agile might not be right for you

Agile has been formalised around practices that try to set you up for success in highly competitive and changing environments. But not all contexts are like this.

Highly regulated or critical domains where the problem is well known and possibly already solved will probably benefit from a more waterfall-like type of approach. [Avionics software](https://en.wikipedia.org/wiki/Avionics_software) will hardly benefit from [A/B testing](https://en.wikipedia.org/wiki/A/B_testing) or [feature-flagging](https://martinfowler.com/articles/feature-toggles.html) _in production_. Nor it will need to wait for user’s feedback on how to keep the plane in the air.

The Agile methodologies have a strong user-centric focus and are heavily oriented towards minimising effort. Try to understand if Agile really makes sense for you and your organization.

I hope you have enjoyed this post. If you did please consider [**following me on Twitter**](https://twitter.com/alediaferia) where I share my thoughts about Startup Software Engineering.

Finally, if you really, really like what I post you could consider joining my mailing list. I won’t spam you!

By clicking submit, you agree to share your e-mail address with the site owner and MailChimp to receive marketing, updates, and other e-mails from the site owner. Use the unsubscribe link in those e-mails to opt out at any time.
