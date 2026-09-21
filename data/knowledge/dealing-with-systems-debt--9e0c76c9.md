---
title: "Dealing With Systems Debt"
notion_id: 9e0c76c9-9439-4197-adbe-729722a0af4f
notion_url: https://app.notion.com/p/Dealing-With-Systems-Debt-9e0c76c994394197adbe729722a0af4f
last_edited: 2023-01-25T18:35:00.000Z
source_url: https://hackernoon.com/dealing-with-systems-debt
tags: ["English", "Programming", "Project Management", "Article", "Hackernoon"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_**Pre-amble/tl;dr:**_ _I recently published _[_The Lloyd Braun Principle of Agility_](https://blog.alishahnovin.com/2022/06/the-lloyd-braun-principle-of-agility.html?ref=hackernoon.com)_ mostly as a silly-but-true observation about how that classic Seinfeld line "Serenity now, insanity later" relates back to Agile. I ended that post with a statement about how we need to better prepare for the "Insanity" that comes later. With that in mind, I'm introducing my approach to building productive teams. It's about __**paying down your Systems Debt.**_

_The modern tech team is broken._ _We rely too much on Seniors, and don't leverage Juniors._

_I've been thinking about this problem a lot over the last year, as I've adapted multiple teams to our "new normal." Months ago, I began an article to support the case that hiring managers should hire juniors. To make an effective argument, I knew I had to address a common concern: "First I need seniors to train the juniors..." Quickly the article grew and grew._

_If you're new to management, this is my "treatise" on Management. It's about how to build highly effective and productive teams and is based on more than a decade of successfully managing multiple large teams._

_What follows is a long read - so the __**tl;dr is:**_

_We talk a lot about #TechnicalDebt, but Technical Debt is a symptom of a larger problem I'm dubbing __**Systems Debt.**_

_The recipe is simple:_

_1) Task your Seniors to pay down your Systems Debt, and make work easier._

_2) HIRE. JUNIORS._

_3) Stop fighting attrition. Juniors will stay for an average of 18-24 months. They want varied experiences. They will seek other jobs. As a tech community, we WANT Juniors to gain more experience on their path to Seniority. It makes them better, and it makes us better. So, let's embrace the attrition. Let's start focusing on enabling them to get the most out of those 18 months, and then support them on their next steps._

Any good delivery team emphasizes usability. The principles that drive usability appear in many different forms: [the Pareto Principle](https://en.wikipedia.org/wiki/Pareto_principle?ref=hackernoon.com), the [Grandma Test](https://www.forbes.com/sites/johnwinsor/2019/12/23/the-grandma-test-are-you-explaining-what-you-do-in-a-way-that-makes-sense/?sh=5a56348573ca&ref=hackernoon.com), the [Principle of Good Enough](https://en.wikipedia.org/wiki/Principle_of_good_enough?ref=hackernoon.com), and more. And all of them get at something very human: _**Is this harder than it needs to be?**_

We see it all [the time](http://seriouspony.com/blog/2013/7/24/your-app-makes-me-fat?ref=hackernoon.com): Minimalist designs that prioritize ease of use and are hyper-focused on a clear vision. Everything else is a distraction, it's bloat.

There are a whole lot of people talking about leadership, management, work/life balance, and remote/in-person work. Everything is being re-evaluated - and the question lingering in my mind for months now has been simply that: Is work harder than it needs to be? It’s by no means a new question, but it’s a thread that’s woven through most of my career. I love writing code and building products, but a holistic view can reveal that more code is not always the answer. More code means more training and maintenance costs.

If you’re a new manager, keeping this question in mind can be very valuable. Managers are responsible for a lot: you’re helping each individual - accomplish their deliverables and their personal goals. You’re responsible for the team’s cohesion, and direction. You’re often responsible for establishing processes and policies. Then there’s resource allocation and planning, as well as working around PTO schedules. But the guiding light to all this is that same question: _**Is this harder than it should be?**_

When is the last time you evaluated your internal mechanics? Thinking of your delivery team as a product, how easy is it for your consumers (the business/operations team) to achieve their goal? And thinking of your entire business as a product, how easy is it for customers to achieve theirs?

Focusing on the people, the same questions apply: How hard is it for your team to do their work? What structures, frameworks, contrived processes and hierarchies exist that would make your grandma shake her head at how much you've complicated things?

## A Primer on Debt

This whole thought train of critically evaluating your internal processes is, in fact, a part of the Agile philosophy that is often overlooked. Teams claim they are fully Agile, or an "Agile-hybrid" but taking a closer look you realize what they mean is they are simply cutting inconvenient corners to deliver sloppy products faster. When it comes to software, any seasoned developer will know exactly what this becomes a recipe for ever-growing technical debt. The ongoing joke of coders is that the last person did it wrong, and if you want it _right_, you'll need to rebuild it. This isn't due to laziness or lack of ability. Greenfield work is easier because it does away with all the technical debt - but, left to the same bad processes, it's only going to grow once more.

Technical Debt is a symptom, and even when you're good about actively paying it down - you're still only _treating the symptom_. This symptom comes from the mistake of thinking Agile is _just_ about software delivery.

A proper definition of Technical Debt is the 'accumulation of work that needs refactoring.' As mentioned, it arises from taking shortcuts when focused on delivery. Unless you're actively paying it down, it grows rapidly. The result of too much technical debt is brittleness and instability.

At its best Technical Debt is an intentional decision: it's putting something on credit with the intention of paying it off later. In such cases, Technical Debt isn't bad - provided you are diligent about paying the debt down. The more nefarious incarnation of Technical Debt is when you unintentionally add to the debt, or worse - you are unaware that you have added to it.

Technical Debt deals with technology, but there's a similar concept of [Process Debt](https://sdtimes.com/softwaredev/how-much-process-debt-are-you-carrying/?ref=hackernoon.com) - legacy processes that have grown stale, function poorly, create friction, impact cross-team dynamics, or may no longer apply as roles have changed. Process Debt covers our non-technical operational deficiencies.

There are still other forms of debt that impact delivery: Design Debt, Architecture Debt.

Of course, _debt_ isn't inherently evil. Just like you can strategically carry a healthy amount of financial debt, taking on any of these forms of debt is a strategic decision. Instead of paying $20,000 on a car all at once, you take out an auto-loan and spread the payments over 10 years. Similarly - the technology stack you use, the way you've structured your team's skillset, the seniority of those roles, and the processes that drive your business take on a form of debt that gets paid down over a period.

Only - sometimes we don't pay it down. Or, worse: we think we're paying down debt but we're really accruing more.

## Introducing: Systems Debt

I'd like to introduce a concept which I'll call **Systems Debt** _(Systems as in Systems Theory; for more on Systems Theory read Donnella Meadow's fantastic_ [Thinking in Systems](https://www.amazon.com/dp/1603580557/ref=cm_sw_r_apan_glt_i_FGZGX3CZQD735P8T5C9D?ref=hackernoon.com)_)._

**Systems Debt** is the overarching and root cause of the downstream forms of debt that impede or negatively impact delivery - whether through decisions from the business, or technical, architecture, or process decisions. It is the product of taking calculated shortcuts in the business, putting work on _credit. System Debt impedes a functioning system through its structural design._

In Thinking in Systems, Meadows provides a simple example of a system: a bathtub. The input to the system is the faucet, the output being the drain. Meadows explains how different factors can cause the tub to never fill, stay level, or overflow. An optimum system is the water remaining level - with the input (faucet) and output (drain) flowing at the same rate. Systems Debt would be the consequence of taking shortcuts when composing the system. To stretch the bathtub analogy, maybe the faucet is poorly installed and water eventually leaks. Maybe the drain's location leads to water accumulating in certain areas so the tub can never fully drain. Maybe our water needs softening and causes lime to build up.

In these cases, there’s no immediate impact, and it’s still possible to create an optimal system - but what’s hidden is the accumulation of debt that is straining the system: the faucet has to flow faster to compensate for leaks, or the faucet flows slower and the tub takes longer to fill.

If you're familiar with Meadows’s book - many of her examples are rooted in reality with models generally providing a static vision; the system doesn't change over time. It's perfect for her purposes, but when it comes to individuals, teams, operations, and businesses, who we are today is not who we will be tomorrow.

To define Systems Debt in a slightly different way:

### **Systems Debt = Systems Theory + Maturity Model**

With Software teams, you see Systems Debt manifest in a few ways:

- With time, if left unmanaged, teams develop more and more "tribal knowledge". These show up as well-intended habits, shortcuts and workarounds. Because they produce short-term efficiencies, they result in long-term deficiencies being overlooked. If it doesn't get paid down, there's an obvious risk of knowledge loss through regular attrition. This leads to productivity losses as teams have to pivot and ramp up. _("Joe knew this, but now that he's left - we'll have to spend time figuring it out.")_ or an accumulation of technical debt due to unfamiliarity. _("We can't update that component. Sally built it, and any time we touch it, it breaks...")_
- Teams will rely on custom design patterns, unofficial development rules or agreements that maximize local flow, but those patterns are brittle to change. This impedes delivery when something doesn't align with an established pattern, holding the business team hostage to technical decisions. _("We can't spin up a custom instance - we've never designed it that way.")_
- Teams that grow complacent place precedence on maintaining the status quo over disruptive retrospectives. They may still hold retrospectives, but if they've lost faith in the ability to address the real problems, they'll only focus on fixing trivial things. _("We're getting overloaded with incident requests, but whatever - we just have to keep resolving them I guess...")_
- The malaise or indifference that comes with hitting a sustainable speed in spite of addressable friction. The sustainable speed creates a false sense of efficiency. _("Do we really need to change the process? We've hit our stride - why disrupt it?")_

With Product & Operations teams, you see Systems Debt in similar ways:

- The _"every customer call is a fire alarm"_ shortcutting, were engaging the development team to quickly fix minor issues delays releases, lowers velocity, etc. _("I need the team to look at an issue raised by an irate customer, it'll only take 20 minutes...")_
- The _last-person-I-spoke-to-is-my-highest-priority_ chaos forces the team to constantly switch contexts and cannot focus on one problem through to completion.
- When prioritizing is a shortcut through the use of misleading metrics. For example, the high-dollar customer that is no longer aligned to the business objectives. They made sense in your early Start-Up days, but as the business has matured that same customer has become more problematic. _("We don't know if we'll want to retain this customer, but they pay us a lot so we need to do it...")_

To reiterate: All forms of debt are when we elect to take a shortcut now with the intention of fixing it later. Technical Debt is when we do this with code. Process Debt is when we do this with our formal processes. Systems Debt is when we do this at the organizational level. I prefer to see it as 'Systems Debt' instead of 'Organizational Debt' because thinking of the organization as a system, it means that Technical, Process, Design Debt are all directly caused by Systems Debt. The factors that lead us to take on Technical Debt are ultimately related to Systems Debt. _("You can only save so many people from drowning in a river before you start looking upstream to determine why they keep falling in.")_

**As an example:** The development team is releasing a new feature that was properly planned and costed. The team has been on track but in the final stages encounter an issue that forces a dreaded question: Delay the release by addressing the issue properly or do the bare-minimum fix and then resolve it properly in the next iteration? They elect to take on the Technical Debt: _"We'll get it on the next iteration."_

This is now where Systems Debt begins to enter the equation. Will the team really be able to address it? Is the team adequately skilled to refactor? Will the business respond with _"it's good enough, we need to move on"?_ Will future costing reveal it's now become _too expensive_ to do the right way? Will a shift in priorities, or a surge in urgent issues, suddenly delay the fix by another iteration? _Then another iteration, then another..._

Additionally, looking upstream: why did it take so long to encounter the issue? What bad assumptions were made? Were they bad assumptions? There's always the issue you can't determine until late in the game - but then why did it become a question of _delaying_ the release? Were promises made too soon? Should there have been a bigger buffer? Would this have been resolved if Person A (upstream business salesperson) spoke more to Person F (downstream developer) following the [shortest chain](https://www.inc.com/marcel-schwantes/elon-musk-bestbuy-hubert-joly.html?ref=hackernoon.com)?

Another all-too-familiar example comes with the infrastructure, architecture, and hosting models we lock ourselves into early on, based on assumptions on how the business will scale in 3-5 years. A small team may elect to take on infrastructure and architectural debt early on in favour of faster delivery than to adhere to the best DevOps principles.

Of course, it's easy to paint scenarios like this without the specifics - but regardless of the specifics and excuses for those specifics: Systems Debt will accrue. It's inevitable. That's OK - it just requires constant attention and focuses on paying it down to maintainable levels.

## Compensating Shortcuts

We take on debt - System or otherwise - as a shortcut. Before digging deeper into Systems Debt and how to pay it down, let’s first take a step back and ask why are we taking the shortcuts? Shortcuts, just like Debt, aren’t inherently bad - but they should be analyzed closely.

Thinking about physical shortcuts is a great way to start. If you’ve ever been a pedestrian or cyclist, you’ve for sure noticed how things are designed for [vehicles first, pedestrians second.](https://www.strongtowns.org/journal/2018/8/13/a-losing-proposition?ref=hackernoon.com) When you walk, you end up taking many _“shortcuts”_ and don’t follow the roads - but of course, these aren’t shortcuts. These shortcuts are the _as-the-crow-flies_ optimal paths for a pedestrian who can go where cars cannot. In fact, building our routes primarily for cars in pedestrian-heavy areas is [another form](another form) of Systems Debt.

In the business world, we take shortcuts to compensate - for lack of time, lack of budget, lack of resources, lack of accountability, or lack of taking a broader view. Time, Budget and Resourcing all get the spotlight - but accountability and a broad perspective are exactly at the heart of my opening question: _**Is this harder than it should be?**_ When you’re saving people from drowning in a river, it’s taking that look upstream _(broad perspective)_ and pointing to the person _(accountability)_ who keeps pushing people in.

In other words: If you are going to have a serious conversation of Systems Debt, everyone needs to be a part of the conversation. **Localized efforts only get so far.**

## So how do you pay down Systems Debt?

Let's get back to that earlier question: _How easy is it for your delivery team to deliver?_ If you've never considered this question, it's time to get some metrics! These metrics won't necessarily give you the answers but they're an important starting point. When it comes to KPIs, the best advice I'd ever received was that an individual KPI is neither bad nor good. It's objectively just a number, a value. It's your business-as-usual - and for you to decide whether or not you want to adjust that number up or down. If you're a fan of the OKR system or SMART goals, this is great because knowing your KPIs allows you to make better OKRs that are easily quantified.

So let's start with some basics and get down in the weeds. What follows isn't a comprehensive list, and there may be better questions suited for your group. Think of this list as a starting point to asking better questions.

### Delivery:

- What is your Lead and Cycle Time for high-priority/urgent items?
- What is your Lead and Cycle Time for low-priority items?
- How much have you allocated towards paying down Debt?
- How much is your Debt increasing?

These questions may seem familiar to anyone who tracks their team's performance - but remember to ask these at the organization level. The developer's Lead Time may have started from when the ticket was created - but how long did it exist in someone's head?

### Resourcing:

- What is your current team breakdown (senior vs juniors)?
- What is your risk of attrition across seniors and juniors?
- What is the cost of losing someone senior?
- What is the cost of losing a junior?
- What is the cost and duration of hiring?
- What is the length of your interviewing/on-boarding?
- What is the length of your training/ramp-up (and what resources need to be involved, and therefore lose productivity?)

### **Business Backlog:**

- Assuming the Product Manager's Strategic Framework is defined, how many exceptions are made?
- How many Epics/Features are not rooted in objectively defined business cases tied back to both financials and timeframes?
- How does the Sales/Business Dev pipeline ultimately feed into your Software backlog, and what is the Lead & Cycle time for those?
- How frequently do high-level business objectives change, and how much does that chaos impact downstream efforts?
- What are the business's long-term growth objectives, and how do the team's skillset and resourcing plan align?

### Operations Backlog

- How many customer issues are raised daily?
- How many customer issues need to go to Tier 1, 2, 3 support?
- What is the average time to resolution?
- What is the lead time for a customer issue, and what is the time from the initial call to when the development ticket is created?

Getting a simple sense of how long it takes for something to go from inception to availability can be very enlightening - _**especially**_ when it's a customer issue.

There are a number of resources that help improve the metrics above - but the key philosophy behind all this is: 1) Measure, 2) Analyze, 3) Resolve, 4) Iterate. The more issues Tier 3 can offload to Tier 2, the more Tier 2 to Tier 1, and the more Tier 1 can enable the Customer to resolve independently the more productive everyone becomes.

### Developers:

- How long does it take for them to grab the source and build?
- How quickly do you provision credentials to new developers?
- How quickly can they stand up and interact with a working lower environment?
- How long does it take for a developer to release code to production (measured from their hire date)?
- How long does it take to ramp-up on your SDLC and processes?
- If they start mid-Sprint, how long do they sit on the sidelines?
- Overall, what is your current learning curve?

_For comparison: Etsy is a great example of efficiency and makes for a great benchmark. Etsy ensures developers deploy to the production_ [_on their first day_](https://codeascraft.com/2011/02/04/how-does-etsy-manage-development-and-operations/?ref=hackernoon.com)_._

### Holistic View:

- What is the delivery journey from inception to availability? Map out the workflow - can it be optimized?
- What is the customer journey from sales to revenue (and beyond)?
- What are the RPO and RTO?
- How does the business performance compare to high delivery performers as described in [Accelerate: The Science of Lean Software and DevOps](https://www.amazon.com/Accelerate-Software-Performing-Technology-Organizations/dp/1942788339?ref=hackernoon.com)?
- What are the [top 10 nightmare disaster scenarios](https://www.amazon.com/Dealing-Difficult-Customers-Dissatisfied-Disagreeable/dp/1632651173/ref=asc_df_1632651173?tag=bingshoppinga-20&linkCode=df0&hvadid=80882877563758&hvnetw=o&hvqmt=e&hvbmt=be&hvdev=c&hvlocint=&hvlocphy=&hvtargid=pla-4584482455385599&psc=1&ref=hackernoon.com)? _This is something FedEx did to improve their operations and build an internal 'Playbook' to make them the company they are today._

Given all the numbers and metrics behind the above, I'll reiterate that any of these numbers represent your business-as-usual. While they're not intrinsically bad or good, Systems Debt makes it harder to maintain these numbers long-term. Some numbers may be surprising and reveal areas where such debt may have already had an impact.

The next step is to consider how these metrics change overtime - as you've matured and continue to mature. As an example - the engineers who built the core product locked you into an architecture that is nearing the limits of its scalability. In these cases, teams look to how they can pay down the Technical Debt - but what about Systems Debt? Given a limited set of resources, an [increasing risk of attrition](https://gitential.com/why-software-developers-leave/?ref=hackernoon.com), and a maturing business - how do you maintain the delivery KPIs while paying down Technical Debt?

## 'Kill Your Darlings,' Fire Your 'Rock Stars,' Destroy your 'Hidden Factories', Stop Being Helpful

Those are a bunch of bold statements in one heading. The point in it all is that: Being "helpful" to shortcut a process can be dangerous. If we subscribe to the idea that “what gets measured, gets done” the problem with being helpful is it often _doesn't get measured_.

Imagine a customer calling because they accidentally deleted a record in their portal when moving too quickly. They're short on time - and can't go through the process of restoring a record. They call in and your Customer Support employee, wanting to be helpful, immediately escalates to the database engineer who, wanting to be helpful, immediately restores the record. The customer is thrilled, the NPS score goes up. Everything is great, right?

Ignoring the obvious risks of someone updating a production database for a moment, there is a lot of valuable information that gets lost in being helpful:

- Why is doing such a dramatic action so easy that the customer made the mistake in the first place?
- Why is undoing the dramatic action so difficult that they had to call in?
- Why could Customer Support not handle it?
- What productivity impact was the result of being helpful and [context-switching](https://www.citrix.com/fieldwork/employee-experience/cost-of-context-switching.html?ref=hackernoon.com#:~:text=The%20cost%20of%20context%20switching%20on%20employee%20engagement,much%20as%2040%20percent%20of%20our%20productive%20time.)? _(A seemingly simple 20-minute task ends up being more than 35 minutes because of the time it takes to get back into a productive flow.)_

Let's be clear: my headline is bold. But, I'm not advocating against assisting a customer - however, I think such actions should be followed with some root cause analysis. Nothing overly formal - but something to avoid a similar problem in the future.

In one organization, I improved our development team's productivity by 50% simply by implementing a Playbook. A customer calls in and they're courteously greeted by a Customer Support rep who follows a clear workflow towards resolution. If they can't resolve it, then we have a feedback loop so that they only ever escalate an issue once. The result was a more capable and skilled Customer Support team, a development team with fewer interruptions, lower stress overall - and, importantly, happy customers.

The point is that when helpful work happens in the shadows, you can't fix the root cause.

We see the same issue with the Rock Star developer who takes on too much work and responsibility to make up for a lower-skilled team - only for them to grow frustrated, get burnt out, and leave (a cost which can be devastating). Will Larson's great book - an [Elegant Puzzle](https://www.amazon.com/Elegant-Puzzle-Systems-Engineering-Management-ebook/dp/B07QYCHJ7V?ref=hackernoon.com) - does a great job of how to handle your "Rock Stars".

## The Power of Juniors...

Seniors are critical to an organization and product's success - but they are equally one of the greatest risks.

For example, a Senior Developer will know the code base through and through. They know what's documented and what isn't. They'll know where it scales well and where it falls apart. They'll know where the skeletons are buried. We turn to them often - relying on them to build and design features, architect solutions, and help resolve the trickiest bugs. They are the knowledge gurus that can answer any question. They train and mentor the junior staff and are consulted when developing solutions.

Suffice it to say, plenty is asked out of the senior staff. It's an obvious statement, given their experience and perhaps vested interest in the organization's success. However, I'd submit, that this is where we take on the most Systems Debt.

_Any senior staff member that progresses a deliverable is a shortcut that accrues Systems Debt._

I'll reiterate because it's easy to misinterpret. You can have a Senior team member work on a deliverable, but you must plan to pay down the debt that will have accrued in doing so.

Relying on Seniors for deliverables is a broken model - _particularly_ in today's world where there are many Juniors and Skilled Entry-Level candidates, and the risk of Intermediate to Senior attrition is both _high and costly_.

The remote virtual workforce has made it more critical for any organization to elevate the junior/intermediate team while reducing the impact of Senior-staff attrition. This doesn't mean diminishing the Senior staff, but it does mean a structural difference in approach.

Using a generalization, today's Senior teams are largely responsible for the complexities behind systems: their experience and expertise produce mature systems, and the smaller task-based components are implemented by the more junior team members.

This is exactly the model that proves problematic when a Senior team member leaves, and also the structure that accrues Systems Debt. In this situation, the Senior is responsible for complexities and can step in to assist when the junior team cannot deliver (the "Rock Star" developer, as an example).

This problem is further exacerbated by the current staff attrition rate: for example, Junior developers nationally will stay in a role for [18-24 months](https://insights.dice.com/2016/07/08/how-long-do-tech-pros-stay-in-their-jobs/?ref=hackernoon.com) (longer at larger firms). In other words, by the time a Junior reaches a point where they can begin to make more significant contributions, they're on their way out.

Organizations will fight to retain Senior staff, fight (somewhat) to retain Junior staff - and are constantly suffering from a knowledge drain. Ultimately, this is a losing battle - even if staff is retained, or new team members join, they're now in a position of having to pay down a large amount of Systems Debt.

## Make Good Work Easier & Embrace the Inevitable

Imagine a small Michelin-starred restaurant. The head chef is very involved in producing the plates, with dishes too complex to distribute among a team of cooks. The chef _is the restaurant_ in this case.

Contrast this with the broader franchise restaurants. You have a head-chef back at Corporate whose responsibility is not to produce dishes that are consumed by customers. Instead, their goal is to produce reproducible dishes. Dishes that can be easily reproduced (while still being tasty). Dishes that are optimized such that the learning curve is minimal - new cooks can easily be trained to produce the dishes, and the loss of their eventual departure is less impactful. The head chefs also partner with efficiency experts to look at how the franchisee kitchen can be optimized for delivery.

This is the model we should use when we think of the modern team. The Senior team's responsibility should not be product complexity. It should be focused entirely on simplifying delivery: simplifying training and ramp-up, set up times, build times, streamlining lead and cycle times (across the board, from Sales/Product Solutioning, through to Iteration Planning, to Release).

How would you structure things if you know you can only retain people for 18 months? In fact, how would you structure things if you put them on an 18-month contract, with a hard termination at the end? You'd want the ramp-up to be as fast and short as possible. You'd want your team of experts to make sure your new hires can be ramped up in weeks so they can maximize the impact. You'd want to make sure your team of experts can keep a revolving door that maximizes efficiency and never step in to assist (for risk of building Debt).

In building a system that can be more adaptable, that can embrace and leverage short-term employment, you'll subsequently reduce the impact if you were to lose a Senior team member because the knowledge becomes immortalized in process, not in people.

## Tag, You're It

Who taught you how to play tag? No matter where you are in the world, you likely learned to play this game from other kids. Adults don't need to teach the kids to play tag.

We think of _**memes**_ as funny images, but the original definition of a meme _is an element of a culture or system of behavior passed from one individual to another by imitation or other nongenetic means._

Tag is a meme. No one owns the rules. No one is responsible for improving the game. In fact, the rules are simple while still supporting variants like Freeze Tag. Furthermore, it can be adapted to different environments. It's designed for a revolving door of children who eventually turn into teenagers that are too cool.

There's very little Systems Debt to a game like Tag. Compare Tag to other playground games that require more players or more equipment... British Bulldog, Dodgeball, Duck Duck Goose, Cops 'n Robbers, Red Rover. Maybe you've played these games, maybe you didn't. These games carry slightly more Systems Debt. More rules, more equipment, or more players means needing more facilitators.

### So how can we operate like Tag?

1. Leverage Seniors to _**lower the bar**_. Seniors aren't there to play Tag or deliver the spec. Their goal is to lower the bar: how can people train faster, and deliver value faster? How can frequent, incremental, and regular releases with low impact risk? _(*cough cough* DevOps, Agile)_
2. Map out the processes, map out the flows, and identify choke points: Maybe the problem isn't the software team. Maybe it isn't a lack of a dedicated testing team. Maybe the problem is upstream: the business is overloading the delivery pipeline with conflicting priorities?
3. What is the driver for people to take a break? There's absolutely nothing wrong with taking a break - but they are often amazing hints at an undercurrent of frustration. Tracking when and why someone is stepping away can be very informative: Maybe they've stepped away because the code takes a while to compile and deploy. Maybe they've stepped away because they need to think more deeply about a challenge they are facing. _(This isn't a bad thing, but could the problem have been broken down to reduce complexity?)_ Maybe they've stepped away because a customer's need is frustrating them. Maybe they need to come up for air because a new feature is forcing a redesign or revealing a bad assumption.
4. Observe the lack of breaks: Equally telling is when someone is too heads down. They can't step away - their attention is required. This either means there's too large a dependency on someone, or the work and scope has been defined too broadly, or the underlying system is far more complex than it should be.
5. Build a culture of simplification. That is, your Managers, your Seniors, your Intermediates - everyone should be empowered to say: "This is more complicated than it should be - how can we make this easier?" Collect feedback - especially from the juniors. Juniors are not valued enough. Don't make the mistake of thinking their lack of experience means they are naïve. Juniors may not always know there are better ways, but they are very frank about where they spend (and waste) their energy.
6. Whenever a Senior contributes to a delivery, find out why. _Every. Time._ A P0 bug in prod required an urgent hot fix. The low-level code required change. A customer needed a discussion. Executives needed a presentation for an update.

A rising tide lifts all boats. Seniors should be the tide, not the boats.

## The Proof is in the Pudding

A guiding principle throughout my career has been to understand how the problem impacts productivity. It's not been to make more efficient teams but more impactful teams. Product Managers have a mantra of measuring Outcome, not Output. Efficiency matters when you know what you're doing, and you just need to do it faster. The impact is an amorphous, poorly defined, moving target. It requires adapting. It's why Agile principles, OKRs, Lean, and Kanban can be so powerful when used correctly.

Focusing on system-wide outcomes and paying down Systems Debt has given me the opportunity to be impactful in a variety of ways.

- For a SaaS business, mapping out the soup-to-nuts process from the early pre-sales stage through to code deployment into production. Measuring the Lead and Cycle time from the perspective of each stage to identify bottlenecks. In viewing it as one system, we could then identify how to better qualify, better disqualify, and how and when to restructure the sales and implementation process based on the client type. Applying a Kanban approach of pulling work rather than pushing, setting up strict WIP limits (while still accounting for slack - see [Goldratt's Drum-Buffer-Rope](https://www.amazon.com/Goal-Business-Graphic-Novel/dp/0884272079?ref=hackernoon.com) analogy), and formalizing the Definition of Done enabled us to continuously improve the workflow itself. Lastly, one additional principle, in this case, was to _not allow the process to impede productivity_. In creating an accelerated track, things could move faster where possible (i.e. "shortcuts") but it always came with an analysis of where the baseline process failed (i.e. paying down the System Debt.) This streamlined implementations from 6+ weeks down to 2 and became a self-managed system.
- Staffing up with a _**Juniors First Philosophy**_ allowed me to quadruple our team size in 6 months. By reducing ramp-up time, and focusing the Senior team around the Junior's productivity, meant we could _**embrace the inevitable and predict attrition**_. This approach left any Junior hitting their 2-year mark with two attractive options: graduate into a more intermediate role where their focus shifted towards enabling Juniors as they continued to grow their skillset, or work with the team on an amicable exit. This transparency helped combat the 2-year itch that Juniors get when faced with career uncertainty.
- Bringing the Operations and Development teams under one process umbrella. This aligned teams so that rather than parallel workstreams, it was circular: The Operations team’s output was the Development team's input. The Development team's output became the Operations team's input. By implementing a _"living"_ playbook, with strategic feedback loops, the teams became increasingly self-sufficient. This freed up the senior team from being involved in operations and reduced lead times from 7+ months down to 3 weeks.
- Pivoting the SaaS development team's view from software development stages (defining, implementing, validating, release) to client-focused with an emphasis on impact while still prioritizing work that is closest to completion. This enabled the delivery teams to have a better sense of impact and priorities and allowed the business to be more aware and critical of low-impact work/clients.
- Investments in streamlining the iterative development process by cutting down build and deployment times. This was one of those _"Observing the breaks"_ moments where the team would frequently step away due to the helplessness of a long build/deploy process. Importantly, the _fix_ wasn't entirely technical - but a process based. By establishing a new process, the delivery teams became 200% more productive.
- Not building software for process inefficiencies. As much as I love to write code to solve interesting problems, new code should be the last resort. Even if the code has no shortcuts, and no Technical Debt of its own, the code needs to be maintained. You've introduced more Systems Debt - and not addressed the root of the problem. These cases are rampant and are easily overlooked - but they're the equivalent of resolving a roof leak by placing a bucket on the floor. Eliminating these fixes by addressing the root problem creates immeasurable efficiencies and enables the team to work on what matters.

### But I’m just a mid-level manager, how can I implement an org-wide change?

I wrote earlier that localized efforts only get so far and I stand by that. When the entire organization takes a critical look at how to be more efficient, that’s where you see real improvement. Localized efforts only get so far - but they do get far, and when they do, they bring with them the capital of influence.

## The Take-Away

I'll conclude with these final principles:

1. Systems Debt accrues for any business shortcut (software, design, process or otherwise). That, on its own, is OK.
2. Too much debt is a bad thing and requires continuous and intentional paying down.
3. This paying down should be done by the Senior team. The Senior team should also focus on how to de-risk taking on future debt (by looking for local inefficiencies first, then system-wide inefficiencies). Seniors should not work on deliverables, but when they do it should be once. Establish feedback loops that ask _"Why was this necessary and how can we prevent it in the future?"_
4. A good measure of Systems Debt health is to look at the health of your junior team onboarding. Plan for your Junior team to leave after 18 months. That's OK. Enable the Senior team to make them efficient as early as possible.
5. Embolden the voice of the Junior team to identify problems. Even naïve problems are indicative of a more serious problem _("When do I get my email address?"; "How can I get test data?"; "What does our team do?"; "I just grabbed the source code, and I'm getting build errors.")_
6. Plan for any person in any role to leave after 18 months. This is also OK. The Senior Team should be there to work on building the processes and pipeline that enable continuous delivery _in spite of disruptions_. They should be building self-sufficient processes, self-managing & self-organizing teams. The Junior Team should eventually get bored of the work, because it's become repetitive. Team members should constantly and actively make themselves redundant and unnecessary.
7. Plan for those who will stay beyond 18 months. Despite working to make themselves redundant, you will never run out of work that drives efficiencies. Define a [Career Framework](https://dropbox.github.io/dbx-career-framework/?ref=hackernoon.com) that elevates Juniors to Intermediates, Intermediates to Seniors that aligns to outcome-focused efficiencies.
8. Treat the first day as if it were the last day: The first day on the job, employees have the best perspective of what it takes to onboard. They should be thinking: If someone else were to join after I've left, how might this be improved?
9. Hold regular Stupid Questions Only meetings. Allow people to submit (anonymously) their most stupid questions. You'll find the biggest knowledge gaps (and problems) this way. These should be addressed.
10. Schedule regular 'Pay Down' meetings: Bring the Senior Team together and look for inefficiencies. Cost their impact. Cost their resolution. _(This list is why your Senior Team will never run out of work.)_
11. The most important part of Agile is its adaptability. Adjust frequently. Have feedback loops. The thing that was working before isn't going to always work. That's not a problem - that's the norm. Anyone who is resistant to continuous change and experimentation is a bigger part of the problem than they realize.
12. Lastly, and perhaps most importantly: This is not a 'team' exercise. This should be a company-wide exercise. While you can drive local efficiencies within your team, they'll only go so far. That being said, if you're not in a position to influence an org-wide approach get the buy in of your manager, and then create charters around your teams that insulate your team from external factors. Work with your manager to define the scope of your smaller system, its inputs and outputs, and get their sign-off on how you plan to evaluate, grow and pay-down your Systems Debt.

That's all! _(He wrote, acknowledging the full irony of having written his longest article.)_

L O A D I N G
. . . comments &
