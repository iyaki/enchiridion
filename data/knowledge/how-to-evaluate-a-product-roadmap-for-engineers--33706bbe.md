---
title: "How-to Evaluate a Product Roadmap, for Engineers"
notion_id: 33706bbe-eb68-43ca-a626-47d710a84c78
notion_url: https://app.notion.com/p/How-to-Evaluate-a-Product-Roadmap-for-Engineers-33706bbeeb6843caa62647d710a84c78
last_edited: 2026-09-21T16:59:00.000Z
source_url: https://stephen.fm/how-to-evaluate-a-product-roadmap/
tags: ["English", "Product Management", "Producer (Individual Contributor)", "Programming", "Article", "Stephen Puiszis"]
---
Aka, how do I know if a product leader knows what they're doing?

There are thousands of articles, courses, and guides on how to build a roadmap for product managers. Even with all these resources, roadmaps often become a point of contention between product managers and their stakeholders. There are very few guides for engineers, designers, or executives on how they should evaluate or engage with a product manager's roadmap. This post comes from a presentation I put together to answer a simple but difficult question I received from an engineer recently: what makes a great product roadmap, and why? As a former software developer, I thought I was well-equipped to answer this question.

While primarily written for engineers, designers, executives, and other product stakeholders can use this guide to ask questions or discuss it with product leadership or the product manager. For product managers, use this as a checklist to evaluate your roadmaps as you're building them or to engage with your stakeholders.

If you don't know what a product roadmap is, Atlassian has a pretty good [guide on the importance of product roadmaps](https://www.atlassian.com/agile/product-management/product-roadmaps?ref=stephen.fm) and how to build them. Once you get the basics, check out the [Silicon Valley Product Group blog on Product Roadmaps](https://www.svpg.com/product-roadmaps/?ref=stephen.fm).

### How-to Evaluate a Product Roadmap for Engineers

This list attempts to define the ideal for software products from a principles-based approach. Physical products are a different animal and not my expertise.

Product management is often a difficult and demanding job, and the role can be wildly different between organizations, products, scale, etc. It's also very context-dependent. There will always be differences and exceptions. In reality, you might not hit every item on this list all the time, but good roadmaps will hit many, and great ones will hit most.

This criteria is a great starting point for evaluating a roadmap, but at the end of the day, the most important thing is simply to engage with your product leader and the roadmap. Ask questions. Share ideas. Make suggestions. It's through feedback and iterations we build better roadmaps, and ultimately, better products.

### (1) Does the roadmap clearly connect to the higher-level company or product mission, vision, and strategy?

The primary purpose of a product roadmap is to explain how we get to the desired future state or outcome. Roadmaps implement strategy. No matter what level the roadmap is for, whether it's for a business unit, platform, product portfolio, individual product, user journey, etc., a great roadmap should clearly execute upon the higher-level mission, vision, or strategy.

This should be obvious, but the opposite happens all too often in larger companies when you have more teams to coordinate, more corporate politics, and often a portfolio of products to manage (e.g., new, mature, sunsetting).

Consequently, the initiatives, epics, or features on the roadmap need to connect the day-to-day work—pixels, code, copy—to the bigger picture, the future we want to create.

**Can you see how the biggest initiatives execute upon the higher-level mission, vision, or strategy? Can you pick an initiative and see how it affects the company's latest strategy? **If you answer no to either of these questions, ask why we are working on this when the strategy is X?

Aside from the obvious issues with misaligned initiatives and not executing on agreed-upon strategies... large initiatives that don't connect to the higher level are at risk of not receiving the resources they need, get de-prioritized in favor of projects more aligned with strategy, and will face political headwinds.

If the product strategy isn't obvious, ask the product manager. If your product or company doesn't have a well-defined vision and strategy that you can easily articulate yourself, check out Andrew Bozworth's post on [Mission, Strategy, and Tactics](https://boz.com/articles/strategy-tactics?ref=stephen.fm) and Lenny Rachitsky's [article on the topic](“If%20you%20don%27t%20know%20where%20you%20are%20going,%20you%27ll%20end%20up%20someplace%20else.”). Send both to your product manager and work with them to better define your product strategy. This will lead to better outcomes for the product and your teams in the long term.

💡

“If you don't know where you are going, you'll end up someplace else.” — Yogi Berra

### (2) Is the roadmap intuitive, and can it be easily explained without jargon?

A roadmap is a communication tool that defines resource allocations in time, money, effort, or all of the above. The roadmap defines what you and your teams will work on for the next quarter or year. In large companies, this could represent millions of person-hours of work. As a result, it is among one of the highest-leverage activities a product manager can work on.

A great roadmap is all signal and little noise. A product manager's job is to deliver clarity.

If you need a dictionary to understand the roadmap, or it's littered with confusing jargon (synergy!) or acronyms you've never heard of or can't look up quickly, the product manager didn't do a good enough job removing noise. Ask what those things mean.

**When reviewing the roadmap or presentation, did you understand it right away? Does it check out with basic logic?** If not, don't hold back—start asking questions. Invest the time now to better define and understand the roadmap. Changing code in production is expensive, and changing text to clarify the roadmap or strategy document is cheap.

### (3) Is the roadmap outcome-oriented or aligned with customer value?

Product managers exist to drive outcomes for customers and the organization. It should be obvious what the ultimate customer or user benefits are from the roadmap initiatives. If they are unobvious or indirect benefits, they should be spelled out somewhere or easily accessible.

Look for features that enhance user experience, solve problems, or offer unique value. These are more likely to drive customer satisfaction and loyalty. If the vast majority of work doesn't benefit or potentially benefit our customers in the future (e.g., via an experiment/test), ask why. Sometimes, roadmaps will have initiatives in them simply because the CEO says so, but a good product leader can strike a balance. Tech debt might need to be addressed. Other times, a roadmap needs to address regulatory or compliance issues, but a good product leader should be able to strike a balance. Most businesses, especially in tech, need to be constantly implementing new features that delight their customers to stay competitive or relevant.

### (4) Is the roadmap flexible or iterative?

If you've spent any time around early-stage ventures, you know the earlier the product or business, the more likely the strategy and goals will inevitably change.

On a long enough time horizon, change is the only constant.

A flexible or iterative roadmap mitigates risk and allows us to learn faster. It allows the team to adapt to changing market conditions.

If you see large or extra-large releases, work with your product manager to see if you can split initiatives into multiple releases or sequences to de-risk them. Can you launch certain features quickly to gather feedback and de-risk others? Can you progressively roll out features via feature flagging? The product manager would love any ideas here.

Finally, a word on dates. They are every team's worst enemy. Executives love them because they provide predictability and allow them to plan and forecast. Dates are often just a necessary evil.

Dates associated with important objectives, like a major public launch at a big conference, are practical and valuable. These _can_ be motivating and help focus teams. However, if you see many dates associated with initiatives, inquire about their logic. Are these dates other teams are depending on us for? Have we publicly announced this date? Why? Do these take into account risk? Anyone who's delivered software before knows you're unlikely to hit your deadlines unless your teams have a well-refined discovery and delivery process and cadence (see questions #5 and #6).

If you're looking to convince your product manager for a roadmap format that's less Gantt chart or date-driven, have them out the [Now, Next, Later](https://www.prodpad.com/blog/invented-now-next-later-roadmap/?ref=stephen.fm) format used at Netflix and thousands of other organizations.

### (5) Are the roadmap initiatives scoped and prioritized based on evidence?

Scoped: In an ideal world, discovery will have already happened for most, if not all, of the important initiatives. Discovery is essential because it de-risks projects and helps us learn what is the right thing to build. Evidence and discovery help us determine the right thing to build faster. If it hasn't happened yet, is discovery accounted for in the roadmap? Why or why not? You might need to budget in time for this.

Some projects will inevitably fail, but you don't want to waste your time consistently building the wrong product or feature. Text on a product requirements document is cheap to change; building the wrong product could be catastrophically expensive. Spend time with your product manager to clarify what you're building. Is there a way you could prototype the feature, product, or some aspect of it to learn faster?

Prioritized: Not all projects are created equal, and product leadership should have prioritized everything on the roadmap. Scoping should have happened already because that impacts prioritization, sequencing, risk, etc. You want to be working on the highest priority or highest impact items. If that's not the case, ask for context. You should voice your opinion if you think something should be prioritized higher or lower. But most importantly, explain why you think that.

### (6) Does the roadmap identify major dependencies or risks?

There are any number of things that could derail projects. Some of these are unknown or unknowable at the time, while others are known. A great roadmap explicitly calls out or considers knowable risks or dependencies.

What are some examples of dependencies? In a large organization, the biggest ones often rely on third parties, other teams, or business units to deliver some aspect of the solution. They often create sequencing and coordination problems.

- Budget for required hardware, software, etc.
- Access to specific data
- Legal or regulatory filings, certifications, or licenses being renewed or granted.
- An internal application team needs a platform product team to make changes
- New vendor implementation

As an engineer, you'll likely be able to spot technical dependencies easily. Did the product manager take that into account? If not, call them out!

Similarly, from a risk standpoint, initiatives often have endless risks. Rather than enumerating every risk, focus on the most important ones for the most important projects from your domain expertise. Consider running a Pre-Mortem exercise for the most important initiatives with your product leader to discover how things could go off the rails and how to mitigate them.

### (7) Does the roadmap feel aggressive but achievable?

Every company has a different philosophy for goal setting or OKR. I believe a great roadmap sets high but achievable goals for teams to hit. Slight or moderate discomfort is often good and forces teams to develop creative solutions. Necessity is the mother of invention. Good goals promote team and personal growth and can be motivating. Conversely, too high of a bar can be demotivating and exert too much pressure on the team. You be the judge of the overall roadmap and provide feedback.

### (8) Does the roadmap take on appropriate risk?

A great roadmap takes on an appropriate or healthy amount of risk. For most products or businesses, you must consistently find new ways to delight customers or differentiate from competitors. Not every project will work out or can be de-risked in advance.

What's an appropriate amount of risk? That's very context-dependent. If you have an unreliable product or have performance issues, then improving what you already have rather than adding new features is a good idea. On the flip side, if you have customers who are already very happy, the roadmap should have some projects that will _delight_ them (e.g., an easter egg) or have experiments looking for ways to significantly improve the product experience.

Finally, does it contain some experiments for us to learn from? Until GPT-1000 comes out, we aren't going to know everything in advance. A great product roadmap builds in some time for experimentation to discover what should come next. If there are no experiments on the roadmap, ask the product manager why.

### (9) Is the roadmap easily referenceable later?

Is the roadmap in a format that is accessible, findable, or searchable? Ideally, is it a living roadmap that your work can directly link to? This probably won't be as important for a new business with a small team as your first few roadmaps might be very simple, and there's less to coordinate. However, this is especially important for larger companies. The reality is that people will forget essential things; no one has a perfect memory. Key people probably weren't in the meeting; they were on vacation, parental leave, or dealing with a fire drill. For any number of reasons, you or your team will want to reference the roadmap later or share it with others. Ask the product manager to make it easy to do so. If they're a good product manager, they'll be happy to hear this.

_Thanks to Justin Ziccardi, Duke Best, and Hacker News for reading early drafts and providing feedback._
