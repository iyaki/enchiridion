---
title: "Prioritization Techniques Compared"
notion_id: 1ea54f1c-7d23-8107-b646-fb3f9aaf63e3
notion_url: https://app.notion.com/p/Prioritization-Techniques-Compared-1ea54f1c7d238107b646fb3f9aaf63e3
last_edited: 2025-07-26T22:58:00.000Z
source_url: https://itamargilad.com/prioritization-techniques-2/
tags: ["English", "Product Management", "Project Management", "Decision Making", "Productivity", "Article", "Itamar Gilad"]
---
## [Part 1](https://itamargilad.com/prioritization-techniques-1/)

Prioritization is the process of choosing which ideas/bets to invest in and which to postpone or park. It entails two cognitively-hard tasks: 1) evaluating the merit of each idea, and 2) comparing multiple ideas and choosing an order of precedence. If that’s not difficult enough, prioritization decisions are often a flash point of debate and power struggles within the company because they’re deemed crucial for business success. Are we good at it? Probably not. Many product people I interview report lack of good prioritization as a top problem in their companies.

In this 2-part article series I’ll review common prioritization approaches and describe the pros and cons of each. They are:

- Intuition, Consensus and HiPPO
- Rules of Thumb — MoSCoW, Kano, Eisenhower …
- Cost of Delay — CD3, WSJF
- User problems — Double Diamond, Continuous Discovery
- Impact, Confidence, Ease — ICE, RICE

## Prioritization in the Age of Evidence-Guided Development

Before we move forward , it’s important to note that a lot has changed in our thinking about prioritization since the introduction of evidence-guided approaches such as Design Thinking, Lean Startup, and Product Discovery:

- Prioritization in the past was an all-or-nothing decision as to which ideas will be built and shipped and which will not. Today we think of prioritization as **deciding what to test first.**
- Prioritization in the past was based on opinions, consensus, and some data. **Today we also factor in **_**evidence**_, which stems from deliberate attempts to prove/disprove our assumptions about the idea.
- Prioritization in the past was a one-time process designed to create a roadmap or a backlog. Today we **re-evaluate ideas every time we obtain new evidence**, which may change prioritization decisions, roadmaps, and backlogs on the fly.

More broadly, today _idea evaluation_ (which is the term I prefer over prioritization) is just one part of product discovery, the other being _idea validation_.

![image](https://itamargilad.com/wp-content/uploads/2025/04/ProductDiscovery-1024x674.jpg)

Product Discovery. Source: [Evidence Guided](http://evidenceguided.com/)

With product discovery we can explore more ideas while reducing the penalty of choosing wrong. In other words, **if you test your ideas rigorously, accurate prioritization is of far lower importance.** You could pull ideas from a hat and still do OK.

With that in mind, let’s look at some prioritization methods.

## Prioritizing Using Intuition, Consensus and HiPPO

At a basic level we can just choose what our gut, experience, and some data tell us. This may work in early-stage startups where the number of ideas is limited and the founders know intimately the customers, the product, and the data. But as we scale, keeping all prioritization decisions in the hands of the founders becomes ineffective.

Larger and more mature organizations often do prioritization-by-committee, involving managers and stakeholders. The committees may help limit individual biases, but introduce group biases such as groupthink and politics. They also slow down decisions and disempower product teams.

Perhaps a greater problem with prioritization by intuition and opinions is that with all the moving parts in the product, market, and technology, it’s practically impossible to say if idea A will work better than idea B, or will have any positive effect at all. Still, our minds easily fall for heuristics and cognitive biases that convince us that we can make such a call, and that our decisions are well-based and rational.

For all these reasons we’re largely moving away from prioritization-by-opinion (at least on paper). Still, human judgement is always going to be a key component. In fact, I’d argue that the **job of prioritization methods is not to make the decisions for us, but to help us make better judgement calls**.

## Prioritization by Rules of Thumb

These are not necessarily full-fledged prioritization systems, but they do suggest which types of ideas we should prioritize:

- MoSCoW — Must-have, Should-have, Could-have, Won’t-have
- The Kano model — Must-Be, One-Dimensional, Attractive, Indifferent, Reverse
- Jeff Bezos’ one-vs-two-way-door decisions – Reversible ideas vs. irreversible ideas
- The Eisenhower matrix — 2×2 matrix of Important vs. Urgent
- Desirability, Feasibility, Viability — Created by design firm Ideo

Pros:

- Catchy and easy to understand and to communicate
- Frame the discussion and potentially lead to better decisions (but see caveats below)

Cons:

- Very broad and subject to interpretation
- No clear success metrics
- Still largely rely on opinions, consensus, and rank
- Not always true

In summary: I’m not a fan of any of these.

## Prioritization by Cost of Delay

**Cost of Delay (CoD), **introduced by Don Reinertsen in his seminal book “The Principles of Product Development Flow” (2009), is an estimate of how much money the company loses each week by not shipping a specific work item. Reinstern declared CoD the most important metric by which to optimize work.

CoD caught on mostly with Agile practitioners. The Two most popular implementations are _Cost Of Delay Divided By Duration (CD3)_ and _Weighted-Shortest-Job-First (WSJF)_. Here’s a brief overview of both.

### Cost Of Delay Divided By Duration (CD3)

Joshua Arnold created [CD3](https://blackswanfarming.com/cost-of-delay-divided-by-duration/) and has written extensively about it on his website Black Swan Farming. Here’s a quick overview video:

The formula of CD3 is: **Value x Urgency / Duration**.

- [**Value**](https://blackswanfarming.com/understanding-value/) Arnold identifies four types of value: increasing revenue, protecting revenue, reducing costs, and avoiding future costs. Yes, it’s all about the money.
- [**Urgency**](https://blackswanfarming.com/urgency-profiles/) expresses how critical the timing of realizing that value is. Arnold defines 4 urgency profiles, each with its own pattern of value delivery over time. I’m not quite sure what’s the the unit of urgency (presumably 1/week)
- **Duration** is the estimate of how long it will take to complete the project

Here’s an example of a CD3 prioritization table:

![image](https://itamargilad.com/wp-content/uploads/2025/04/CD3-table-1024x268.jpg)

Source: [BlackSwanFarming.com](https://blackswanfarming.com/cost-of-delay-divided-by-duration/)

Note: Arnold also proposed a [non-numerical version of CD3](https://blackswanfarming.com/qualitative-cost-delay/) which was [further adapted by John Cutler](https://cutlefish.substack.com/p/tbm-245-the-magic-prioritization).

### Weighted-Shortest-Job-First (WSJF)

WSJF ( pronounced Wisjif) is the SAFe (Scalable Agile Framework) take on Cost of Delay and therefore quite a popular prioritization method today.

Here’s a quick explainer video by Appfilre:

The formula of WSJF is: **(Business Value + Time Criticality + Risk Reduction) / Estimated size**

- **User-Business Value –** the relative value of an item to the business or customer
- **Time Criticality – **Does the cost of this problem increase over time if we do not act? Is there a deadline? Are there penalties for non-compliance that take effect? Are volumes increasing or remaining steady?
- **Risk Reduction-Opportunity Enablement – **Does this reduce risk, or allow us to take advantage of opportunities not available to us previously?

All values are taken from predefined 5-point scales such as this:

![image](https://itamargilad.com/wp-content/uploads/2025/04/wsjf-value-table.jpg)

Here’s what a  WSJF idea bank might look like:

![image](https://itamargilad.com/wp-content/uploads/2025/04/WSJF-Idea-bank-1024x250.jpg)

Source: [kendis.io](https://kendis.io/scaled-agile-framework/weighted-shortest-job-first-wsjf/)

## Evaluation of Cost-of-Delay Methods

I haven’t practiced CD3 or WSJF myself (although I have consulted companies that use the latter), so these are my opinions only.

### Pros:

- Better than just relying on opinions or rules of thumb
- Focus on consistent business metrics
- Factoring in the urgency of a feature makes intuitive sense — some things are critical by nature

### Cons:

- Optimizing just for money (CD3) — Revenue and costs are lagging indicators and ones that most product teams cannot directly affect (there are too many other factors). Also, sometimes, companies are better off temporarily focusing on lower-level metrics — retention/churn, number of transactions, engagement levels… Lastly, focusing only on money is likely to reduce [customer-focus](https://itamargilad.com/customer-value/), which is very risky in this age of customer choice and fast-moving competitors.
- _Business value_, _Urgency_, and _Risk Reduction_ are vague and hard to estimate** — **In my experience product teams will really struggle to perceive and estimate such things and will have to rely on one of these bad choices: 1) gut-instinct guesses, 2) business stakeholders’ input, or 3) delegating the estimate to business-intelligence or finance. None of these is going to be remotely accurate, and some will be very slow.
- Not factoring Evidence — both implementations of CD3 lack any mention of available evidence which means that scores that are based on one-person’s guesses are just as good as scores based on thorough research and experimentation.
- Is Cost of Delay really the most important optimization? — While some things are time-critical (and maybe we’re losing money not having them), should this be the guiding light for all our activities? What about the mission and the strategy? The goals of the company and its parts? Ethics and legal?

Overall, Cost of Delay strikes me as the sort of revenue/cost optimization that may appeal to certain engineers and execs, but is missing a lot of nuance.  It looks as if it was conceived in the age of [big, expensive projects](https://itamargilad.com/big-projects/) that either were added to a roadmap or were parked. It doesn’t feel like a very scalable, or for that matter, agile process to me.

If you are using one of these methods I would consider these adaptations: 1) Broadening the definition of _Value_ to allow for other metrics, including customer-value (which WSJF seems to have done) 2) Adding evidence-based Confidence (more on this when I talk about ICE) as one of the components.

## Takeaways

In this article I touched on the many limitations of intuition, consensus and HiPPO when it comes to prioritization. Old chestnuts like the Kano model or the Eisenhower Matrix may help frame the discussion somewhat, but won’t take you very far either.

Cost of Delay with its two derivatives — CD3 and and WSJF — warrants more attention, but I I would suggest adapting it to include support various value metrics and some element of evidence-based-confidence.

In part 2 of this article series I’ll discuss two very popular prioritization techniques: Teresa Torres’ Continuous Product Discovery and the popular ICE/RICE method that I’ve written a lot about in the past. This will be an opportunity to contrast a design-oriented approach that derides putting numbers on ideas, and very numerical alternative. I have my praise and my criticisms of both. This article is coming shortly.

## [Part 2](https://itamargilad.com/prioritization-techniques-2/)

In the[ first article](https://itamargilad.com/prioritization-techniques-1/) in this series I reviewed common prioritization approaches, including opinion/consensus/HiPPO, rules-of-thumb, and Cost of Delay.

I also explained why prioritization as part of product discovery — prioritizing what to test — is far more effective than prioritizing what to build. You greatly reduce the cost of prioritization errors, and at the same time allow for more ideas to be accepted, reducing tension and conflict, and improving the odds of finding those rare good ideas.

In this article I’ll focus on the two most common ways to prioritize — prioritizing using user problems and using impact, confidence, and ease (ICE). We’ll see the pros and cons of each, and whether they’re truly mutually-exclusive as some product experts will have you believe.

## Prioritizing by User Problems (aka Opportunities / Needs)

If cost of delay is all about money, then user-problem prioritization is all about the users. The core idea is that business goals should be connected to _user problems_ ( also called _opportunities_ or _underserved needs_). Solving the problems may change user behavior in ways that will drive desired business results. For example if the business goal is to increase the number of signed users, we may look for problems in the sign-up process, prioritize these problems and fix in order of priority, thus an uplift in the number of signups.Derived from design philosophy, this approach often puts a strict order — first map the problem space and identify the most important problems, then think of solutions (aka _ideas_ or _bets_). _Design Thinking_ popularized, by the design firm Ideo, is by far the most influential methodology/philosophy in this space. Teresa Torres’ _Continuous Discovery_ (CD), described in her book [Continuous Discovery Habits](https://www.amazon.es/Continuous-Discovery-Habits-Discover-Products/dp/1736633309) is a modern and popular take, best known by its main artifact: Opportunity Solution Trees (OST).

![image](https://itamargilad.com/wp-content/uploads/2025/05/Opportunity-Solution-Tree-550-x-401.jpeg)

The OST connects business goals to a hierarchy of opportunities found through user research, from which solutions are derived. The assumptions in the solutions are validated using experiments. Source: Teresa Torres / Product Talk

Continuous Discovery, as the name suggests, is a full system for product discovery, but it also includes some implicit multi-level prioritization:

- Prioritizing opportunities — Torres recommends tackling one sub-opportunity at a time. You first pick the most important top-level opportunity, and then choose one of its sub-opportunities. The choice between sibling opportunities is based on: 1) opportunity sizing, 2) market factors, 3) company factors, and 4) customer factors. In her book, Torres explains those in general terms and urges practitioners not to try to use numbers, but rather make a “data-informed, subjective comparison of each of the factors”. The decision is “messy and subjective, and we want to keep it this way”. To ensure practitioners won’t over-dwell on the choice, Torres recommends using Jeff Bezos’ one-way/two-way door criteria to determine which decisions are reversible and which are not.
- Picking solutions/ideas — Through brainstorms the team produces 15-20 ideas per chosen opportunity. Ideas that don’t clearly address the opportunity are then filtered out. Out of the rest, the team picks 3 winners using multiple rounds of dot-voting. Torres is very adamant that you must [Prioritize Opportunities, Not Solutions](https://www.producttalk.org/2019/02/prioritize-opportunities/?srsltid=AfmBOoojP94ZRubsReYiw9I0c1tDh9io322057rI6IFMRRfod-PpO0qf) and specifically calls out the practice of grading ideas in a spreadsheet (see ICE below) as fundamentally flawed.

### My Evaluation of Prioritization by User Problems

I’ll focus mostly on the prioritization method of continuous discovery here as other design-inspired methods vary widely in implementation. While I practiced variants of Design Thinking during my time at Google, I didn’t practice Continuous Discovery (CD) directly, but I’ve consulted companies that did. So these are my opinions, rather than direct hands-on observations:

**Pros**

- CD emphasizes making decision based on evidence (what I call _evidence-guided development_) through a combination of user research and experimentation — exactly as it should be
- CD encourages user research, and especially user interviews, which many companies don’t practice enough. Adopting CD should amplify the customer voice inside the company and elevate customer-centricity.
- Opportunity solution trees offer a practical way to connect business goals to action. I’d argue that’s not the only way, and it is more appropriate for certain types of goals and solutions.
- CD deliberately tries to avoid cognitive load by encouraging fast, imperfect decisions, which makes adoption and practice easier.

**Cons**

I leveled most of my criticism towards the solving-user-problems philosophy in the article [You’re Not Just Solving User Problems](https://itamargilad.com/solvingproblems/). Here are the key points:

- “It’s all about solving customer problems” is too limited a worldview in my opinion. In practice, the company has needs of its own that don’t always overlap with customer needs. For example lowering costs is a legitimate need of the company, but interviewing customers will tell you little about the true opportunities (some of which may be technical or financial). Adopting AI is also a pressing need for many companies, but customer interviews will likely not tell you much.
- The map-the-problem-space-before-the-solution-space model (double diamond) is often overly restrictive and impractical. Many ideas come directly from customers, stakeholders, managers, and the team, not through user research. Rejecting these ideas outright or putting them through the sift of user interviews is impractical and will likely annoy your colleagues. Hence there’s a need for a general idea prioritization system (which can live alongside continuous discovery).
- User research is just one form of research; others include data research, market research, and technology research. Some of the most important innovations in the last decades have emerged from those.
- User research is rich in qualitative meaning, but does not guarantee statistical significance. For example, the choice of interviewees can greatly affect the findings. The opinions and biases of the researcher can inadvertently color the results. Hence I find it’s best to use multiple forms of research and to cross-correlate.
- While I share Torres’ recommendation not to try to make prioritization a pseudo-science, choosing what to work on using broad criteria like “opportunity size” or “market factors” in a deliberately “subjective and messy” way, sounds error-prone to me, especially as there’s no way proposed to test the assumptions behind the opportunities except user interviews. I’m also not a massive fan of team dot-votes as a way to pick ideas.

To be clear, Design Thinking and Continuous Discovery are widely practiced and I met many practicionaris who warmly recommend them. I feel they’re very valuable, but would recommend combining with other methods rather than using exclusively.

**Upcoming Workshops**

Practice hands-on the modern methods of product management, product discovery, and product strategy.

Secure your ticket for the next [**public workshop**](https://ti.to/itamar-gilad-events/)

or book a [**private workshop**](http://itamargilad.com/workshops) or [**keynote**](http://itamargilad.com/keynotes) for your team or company.

![image](https://itamargilad.com/wp-content/uploads/2023/06/brain-fluid.png)

## Prioritizing by Impact, Confidence, and Ease (ICE)

ICE (Impact, Confidence, Ease) was invented by Growth guru Sean Ellis as a way to rank growth _experiments_, but is now widely used to prioritize product and business _ideas_. I’ve written extensively about ICE in my book [Evidence-Guided](https://mybook.to/0hPtd1Y), in my [eBook](https://itamargilad.com/ice-ebook/) on the topic, and in multiple articles.

![image](https://itamargilad.com/wp-content/uploads/2021/12/Idea-Bank.jpg)

To use ICE we need to first collect ideas in an Idea bank, and then estimate three values for each:

- **Impact** — how much does this idea stand to improve the target metric. ICE can flexibly be used to estimate on any metric — from the company [north star or top business metric](https://itamargilad.com/the-three-true-north-metrics-that-your-product-and-business-need/), down to a quarterly key result. Naturally you have to use the same metric across the ideas you wish to compare.
- **Ease** — how easy is it going to be to build and launch this idea in full (without any testing). Usually this is the opposite of person/weeks.
- **Confidence** — how strong is the evidence that we will have this expected impact and ease.

Each value is normalized to a range of 0-10. We multiply the three values, or average them to get the ICE score. The scores are just a hint — knowing what we know now, these ideas look most promising and should be tested first. ICE does not guarantee that these are the best ideas, or that they will even work.

RICE is a derivative of ICE invented by Intercom. It adds a fourth component — Reach — how many users/customers will be impacted. Some ICE practitioners, me included, argue that Reach is simply a component of Impact, and not necessarily a component you always want to factor.

### Evaluation of ICE/RICE

ICE is an area where I feel I have deep experience both as a practitioner and as a coach and I’ve seen many companies benefit from using it. But in no way am I here to endorse the method as the best or only way to prioritize. As you’ll see below ICE has some clear challenges that must be addressed to make it effective.

Pros of ICE/RICE

- Flexible prioritization system that allows teams to focus on any metric. ICE works very well with [Objectives and Key Results](https://itamargilad.com/ebook-okr/) and with metrics hierarchies.
- Confidence, when used correctly, has several important benefits. a) it reflects how reliable and trustworthy the other guesstimates are, b) it encourages testing and validating ideas and making evidence-guided decisions, c) It can act as an important antidote to opinions, biases, and HiPPO.
- ICE is fairly easy to understand by anyone in the company, and it acts as a powerful tool to communicate prioritization reasoning.
- In my experience, switching to looking at ideas through the lenses of impact, confidence and ease, greatly shortens idea discussions and elevates the quality of the decisions. It focuses people to think of the impact on the goals, on costs, and on confidence/risk factor. Often the right decision is that we need to test more to decide and ICE helps drive the point.

Cons of ICE/RICE

- I regularly see product people use opinions and sparse data to estimate Impact and Ease, and yet assign a high Confidence value. This creates very bogus and subjective prioritization which earned ICE many detractors. To address this problem I created the Confidence Meter, which assigns weights to different categories of evidence.

![image](https://itamargilad.com/wp-content/uploads/2022/05/Confidence-Meter-No-CC-1536x1016.jpg)

The Confidence Meter. Download it as a free calculator here

- I regularly see product people use opinions and sparse data to estimate Impact and Ease, and yet assign a high Confidence value. This creates very bogus and subjective prioritization which earned ICE many detractors. To address this problem I created the Confidence Meter, which assigns weights to different categories of evidence.

> 

Here’s an alternative way to prioritize. At any given point you’ll need to pick some low-confidence ideas for early validation (typically handled by the product manager), some medium-confidence ideas for early testing (requiring other team members to help), and a few high-confidence ideas for advanced testing and delivery (requiring heavy engineering and design investment). Keeping all three funnels full ensures you never run out of ideas to work on. So you can split your list of candidate ideas into these three groups by level of confidence, and pick ideas in each.

(source [Evidence Guided](http://evidenceguided.com/))

- ICE requires people to make cognitively-hard assessments of the three values (impact is especially hard). The rule “don’t make me think” applies also to product people and I find that triads/trios can tire of doing too much ICE, and start neglecting the practice. I tried to address this challenge in my book, by breaking ICE prioritization into several stages in the lifecycle of the idea, where most ideas will only require quick guesstimates, while few will require more in-depth analysis. Still this is the biggest challenge of ICE which requires quite a bit of discipline.

![image](https://itamargilad.com/wp-content/uploads/2023/08/Idea-Processing-900px.jpg)

### Final Thoughts

It would have been lovely to have one perfect prioritization system that predicts the future with high probability, but sadly no such thing exists (short of a future-seeing crystal ball, currently out of stock in Amazon). It’s clear that relying too much on intuition and rules of thumbs is very haphazard. The three main approaches: Cost of Delay, User problems, and ICE, all offer marked improvements, but also come with their own caveats and cons.

Counter to what you may hear from purists, you _can_ mix and match methods. For example WSJF (weighted shortest job first) could be made more evidence-guided by incorporating a fifth value — Confidence (perhaps using the [confidence meter](https://itamargilad.com/resources/confidence-meter-calculator/)). Opportunity Solution Trees can be paired with ICE (I know teams that do just that). The old truism holds — there isn’t just one right way to develop products; you need to try the methods and adapt them to your context and needs.

Perhaps the bigger takeaway is that you should see prioritization not in isolation, but as part of a larger system that includes research, idea evaluation, and idea validation. After nearly 30 years in the industry, this evidence-guided approach is the only effective way I’ve found to build high-impact products, and it doesn’t really matter whether you call it GIST, Continuous Discovery or something else.
