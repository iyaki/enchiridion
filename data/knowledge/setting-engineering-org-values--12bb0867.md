---
title: "Setting engineering org values"
notion_id: 12bb0867-0b1e-4788-9fdd-5754298f42b0
notion_url: https://app.notion.com/p/Setting-engineering-org-values-12bb08670b1e47889fdd5754298f42b0
last_edited: 2023-07-10T17:48:00.000Z
source_url: https://lethain.com/setting-engineering-org-values/
tags: ["English", "Principles", "Entrepreneurship", "Line/People/Team Management", "Article", "Irrational Exuberance (Will Larson)"]
---
_This is an unedited chapter from O’Reilly’s _[_The Engineering Executive’s Primer_](https://www.oreilly.com/library/view/the-engineering-executives/9781098149475/)_._

Uber’s best known corporate value is probably _Super Pumped_, which, in addition to being a one-time company value, is also the title of [Mike Isaac’s account of Uber](https://www.amazon.com/Super-Pumped-Battle-Mike-Isaac/dp/0393652246) and the subsequent television show. However, for me personally, the value I remember most is _Let Builders Build_.

Working in Uber’s infrastructure engineering organization, I once chatted with a product engineering manager who wanted to continue rolling out a new feature that was hammering the production database. I was concerned that ramping up usage would cause wider application instability, and believed they needed to optimize the query before further rollout. The product engineering manager disagreed. They argued that my concern violated the _Let Builders Build_ value; living up to Uber’s values was to approve the accelerated release, even if it risked significant user and business impact. While I personally loved the idea behind the _Let Builders Build_ value, that discussion highlighted just how challenging it can be to use values to operate a company.

Stated values possess no magic for cultural transformation, but with a bit of care they can be useful for consolidating cultural change that you’ve already made. In my experience, company and organizational value statements are both overused and underused, and it’s useful to dig into how to incorporate them:

- What kinds of problems do values solve?
- Should engineering orgs have values at all?
- When does it make sense to establish values out?
- What makes values useful?
- How are engineering values distinct from a technology strategy?
- How should you roll out values?

Finally, I’ll end spending some time discussing some of the values that I’ve personally found effective within the engineering organizations I’ve worked with.

## What problems do values solve?

Stated values make it clear how you want people to make decisions. If senior team members live and role model the values, then those values will become an active, living artifact. If stated values don’t align, then they’ll become a frequently forgotten, running joke.

Building from that, some scenarios where values are useful:

- You’re hiring more and more senior engineers and managers, who come with a significant amount of working experience at other companies. Documented values increase cohesion across the new and existing team. They also avoid the scenario where new hires unknowingly practice their previous companies’ values, causing a cultural rift between new and existing teams
- You’ve made a significant cultural change over the past year to welcome ideas from across the team rather than only accepting top-down ideas, and you want to formalize that change so it persists over time
- The engineering organization generally does a good job of reusing existing coding patterns and services, but a few vocal engineers are advocating for teams setting their own patterns independently of the organization’s existing ones. Formalizing the reuse of existing approaches when possible will prevent a prolonged conflict
- You’ve acquired a very small company of five engineers to join your existing organization of five hundred engineers. You’ve been clear in the acquisition process that you’re looking for the new team to merge into the existing organization, and your documented values help them navigate that change successfully

Values are extremely useful in all of these scenarios, but most importantly recognize what’s missing: they are absolutely not an effective way to start the underlying culture change that they describe. They can celebrate change, and make it permanent, but stated values are not cultural magic.

## Should engineering orgs have values?

Before digging into crafting _useful_ values, it’s worth spending some time discussing whether engineering organizations should have values at all. Org-specific values tend to be controversial. Even if you’ve never heard of them, most companies have a written set of values somewhere, and many executive teams will take it poorly if you introduce engineering values. Why not, they’ll unwind the fishing reel loaded with a leading question, simply use the existing company values instead?

The best case is that you find company values that can be extended to match your goals. You might take [Amazon’s leadership principle](https://www.aboutamazon.com/about-us/leadership-principles) of _Frugality_ and add some engineering-specific nuance: in our pursuit of _Frugality_, we consider the full cost of building and maintaining a new service against the cost of paying a vendor, and we find using a vendor is often more frugal than building it ourselves. Adding an organization-appropriate interpretation to an existing value reinforces that company value, rather than detracting from it, and is an easy sell to the wider executive team.

Sometimes, even if you get very creative at extending the existing values, there simply isn’t an existing value that fits. In that case, there are three paths forward: add a new company value, introduce engineering organization values, or introduce engineering leadership values (for Engineering Managers, Staff-plus Engineers, or ideally both).

Which will work best depends heavily on your company size, your executive team, and the sort of values you’re hoping to add. Some values simply aren’t as relevant outside of engineering: a value around build-versus-buy makes less sense for organizations that don’t have the build option, and might make more sense as an engineering value. Other values might work well for an entire company, such as [creating net-new capacity rather than competing internally to capture existing capacity](https://lethain.com/create-capacity/).

If the value is applicable to the whole company, and there’s excitement within the executive team to add another value, then I’d encourage you to start there. It’s much more work to maintain values than to create them, and adding to the company values will allow you to share the maintenance across the executive team. If it’s not widely applicable, then you have the choice between adding either engineering organizational values or engineering leadership values.

Because people often have a strong internalized belief that a company should have one shared culture, you may find that introducing organization-specific values often run into a surprising amount of friction. If your company already has organizations with org-specific values, then by all means follow that pattern, but otherwise I’d recommend introducing engineering leadership values instead of engineering organizational ones. Engineering leadership values tend to accomplish the goal at hand without running into conflict with the internalized belief that companies should have one, uniform set of values.

Nominally, engineering leadership values that apply to Staff-plus Engineers and Engineering Managers. In practice, almost everyone aspires to be a leader, and will model their behavior on leaders within the team, so leadership values tend to establish themselves as organizational values while side-stepping some of the tripwires that come being explicit.

In addition to personally finding success with leadership values, there are other examples in the wild, notably [Amazon’s Principal Engineering Community Tenets](https://www.amazon.jobs/en/landing_pages/pe-community-tenets). (Note that these “Principal Engineering” tenets apply _really well_ to anyone in engineering, regardless of their title.)

## What makes a useful value?

Companies sure do have a lot of stated values, but I am going to argue here that a significant number of those stated values are simply not useful. Often someone in the value statement creation process gets caught up in the dream of what the company _could_ be, rather than what it _is_, and the values detach from reality. This turns a useful project into something that is at best useless, if not actively misleading. To avoid creating useless values of our own, we’re going to walk through my short definition of what makes a value useful.

A useful value is reversible, applicable, and honest:

1. **Reversible: it can be rewritten to have a different or opposite perspective without being nonsensical.**

Reversibility is a precondition to the next two values, applicability and honest, but it’s also meaningful in its own right. Effective values guide behavior, and it’s only practical to guide behavior when there are multiple, viable approaches. Irreversible values invite lip-service agreement rather than active participation.

**Example:** Uber’s value of _Make magic_ emphasized a willingness to delay product releases that were merely serviceable until they had a spark of delight. Many companies practice the reverse value of, _Ship early and often._ Any given successful company might take either approach.

**Counterexample:** Amazon’s _Are Right, A Lot._ How do you reverse that into a reasonable value? Something like “Are right when it matters” initially sounds like the reverse, but on deeper inspection probably means roughly the same thing. A more extreme reverse like “Are Wrong, A Lot” wouldn’t make much sense unless you extend the value significantly to something like, _Quickly Learn from Mistakes_, which again, is conceptually very similar to _Are Right, A Lot_

1. **Applicable: it can be used to navigate complex, real scenarios, particularly when making tradeoffs.**

Many values are written down and forgotten. To keep a value alive and useful, it needs to be used frequently and visibly by the team. Applicability means a value that contributes to planning sessions, performance reviews, and hiring decisions.

**Example:** Stripe’s value of _Seek feedback_ clarifies the expectation that you should be actively seeking feedback on your work rather than working in isolation. If you’ve completed a [Tech Spec](https://infraeng.dev/tech-spec/) without gathering input, then you know that at Stripe you’re not actually done: you need to seek feedback before finishing the work.

**Counterexample:** Uber’s _Be Yourself._ It’s unclear how you’re supposed to apply this value to a practical scenario. Another counterexample is Stripe’s _Be meticulous in your craft._ It’s a beautiful aspirational value, but it’s sufficiently broad that two reasonable people can’t align on whether or not it applies to any given decision

1. **Honest: it accurately describes real behavior.**

A touch of aspiration is OK, but useful values should explain how effective employees navigate the organization as it exists today. If it’s too aspirational, values become a trap for your best-intention employees. More cynical employees will ignore them anyway, modeling their behavior on the actions of successful internal role-models rather than your stated values. The only way for your entire team to operate on the same values is to describe behavior honestly. (If you want to change company behavior, change the behavior first, and document it second.)

**Example:** Uber’s value of _Meritocracy and Toe-Stepping_ has been criticized externally, but it was a very accurate reflection of the actual internal culture. You went into a meeting with folks at Uber, and that’s how they acted.

**Counterexample:** Amazon’ _Strive to be Earth’s Best Employer._ There are many things uniquely good about working at Amazon, but few would argue that their frugality makes them the best employer. The best employer probably wouldn’t use Amazon’s current vesting cliff (often vesting 20% over the first two years, and 80% over the next two years)

As a test, let’s discuss a few values that meet all three criteria:

- Amazon’s _Dive Deep_ asks leaders to engage with the details at hand. If they’re uncertain about the right path forward, dig in, rather than delegate. This is reversible: many successful companies instead delegate decisions down the hierarchy to whoever has the details rather than expecting leaders to drill in themselves. This is applicable: whenever you’re not sure about a given decision, you should dive into the details before moving forward. This is honest: most folks I’ve worked with from Amazon have embodied this mentality, and are willing to go deep into problems
- Stripe’ _Seek feedback_ asks folks to socialize plans and decisions before finalizing them. This is reversible: other companies would urge moving quickly by avoiding consensus building. As discussed earlier, this is applicable, because you know any major decision needs to be written up and shared. Finally, it’s also honest: this is how people behave at Stripe (when folks complain about Stripe, it’s often that it works this way _too much_)
- Uber’s _Be an Owner, not a Renter_ asks folks to make decisions they’re willing to support long-term rather than assume someone else would absorb the long-term consequences. This is reversible: many early companies want to emphasize doing quick work (e.g. YCombinator’s [Do Things that Don’t Scale](http://paulgraham.com/ds.html)). This is applicable: whenever you’re making a short-term versus long-term tradeoff, you can apply this value. It was also, more or less, honest, although it certainly varied a bit across team (infrastructure teams were incentivized to be very aligned, whereas growth teams were generally incentivized in the other direction)

Each of these values is genuinely useful for deciding if a company is for you, and also for operating successfully within that company. If you pick values like these, you’ll get a lot out of them.

Some folks will argue that good values should be resistant to misuse. For example, my starting story about _Let Builders Build_ might be a bad value because it was misinterpreted by the product engineering manager. I don’t personally worry too much about that. Self-interested individuals will _always_ interpret things in unrealistic ways to their own benefit, and the solution is holding them accountable for their poor behavior, particularly if they are in role modeling roles like Staff-plus Engineer or Engineering Manager.

There are two particularly common categories of non-useful (or colloquially, useless) values that occur frequently enough that they’re worth discussing directly: identity values and prioritization values. Identity values are things like Amazon’s _Are Right, A Lot_, Uber’s _Champion Mindset_ and Stripe’s _Deliver Outstanding Results._ Identity values can be reversible and may be honest, but they are very rarely applicable. How often are you finishing up a project where you think to yourself, “I should focus on being right, rather than wrong, for this one.”

Prioritization values are Uber’s _Celebrate Cities_ and almost every startup’s variant of _Make Big Bold Bets._ These values often are honest, but struggle to be either reversible or applicable. On a literal level, these values are an incomplete investment thesis for planning–Uber should do some work to celebrate cities, and your startup should do some big bets–but are missing too much other context about other priorities to be applicable. More importantly, if you dig into any prioritization value, you will usually find a hidden identity value. For example, making big bets is really about the desire to be innovative and ambitious.

While you should try to avoid useless values, as long as they’re honest, they tend to be inert rather than harmful, so I wouldn’t spend too much time fighting against, particularly them if you’re working on values as a participant (e.g. for the company) rather than as the final decider (e.g. for engineering).

## How are engineering values distinct from a technology strategy?

As described in Rumelt’s [Good Strategy, Bad Strategy](https://www.amazon.com/Good-Strategy-Bad-Difference-Matters/dp/0307886239), a strategy is composed of three parts: circumstances, guiding principles, and concrete action. The best way to think of the relationship between values and a strategy (business, technology, or otherwise) is that useful values generally can serve as a strategy’s guiding principle. Not all guiding principles are values (e.g. how you respond to a current market opportunity is unlikely to be a value), but most values are viable guiding principles.

For example, _Reuse common technologies unless you see 10x improvement_ is a useful engineering value, and could well be a guiding principle you use to address a specific set of circumstances. While values are usually presented in a way that’s divorced from the circumstances that help form them, you can reverse-engineer the implicit circumstances that fed into any given value’s creation.

By default, folks often view it as a failure if their organizational values and strategy have too much overlap, but I see them as tightly connected. Rather than a bad sign, I view overlapping organizational values and strategy as a sign of attentive, detail-oriented leadership.

## When and how to rollout values

Somewhere, there’s a playbook that encourages new executives to quickly publish new values, similar to the nervous desire to show value that culminates in [new executives announcing brand-new technology architectures shortly after joining](https://lethain.com/grand-migration/). Don’t do that. Wait. Wait at _least_ six months. If that’s too long, at minimum wait until you can evaluate whether a given value is honest or aspirational. You can absolutely test some values earlier in small groups, but if you’re confident you _need_ to roll out values quickly after joining a company, I’d push you to consider whether this is the right work or if [you’re retreating into comfortable work](https://lethain.com/reminiscing/).

By focusing on honest values, you’ll have fewer rollout challenges from folks rejecting values outright, but the rollout will still be an involved process. Establishing values should follow [the general patterns of good process rollout](https://lethain.com/good-process-is-evolved/): identify the opportunity, document potential options, involve stakeholders early to build buy-in, test before finalizing to avoid folks feeling trapped, and iterate until it’s useful. It’s easy to announce values, but much harder to introduce values that get used. Reduce that risk by including the wider team, listening, and iterating a few times to make them feel like a shared creation rather than a top-down one.

Once you’ve completed the initial rollout, it’s important to recognize that values are more like a garden than a building. If they’re not part of your daily processes, they’ll probably be forgotten:

- integrate them into your hiring process (including letting candidates opt-out if they don’t like them)
- explicitly talk about them in new-hire onboarding
- update your career ladder to require value-aligned behavior for promotions
- highlight culture-aligned accomplishments in one-on-ones, team meetings, and elsewhere

The one caveat to the standard process rollout is that if you end up writing company values rather than engineering org or engineering leadership values, then you will likely have quite a few peer executives or the CEO involved, and the reality is that executives are hard to hold to a process. This is mostly unavoidable, and you’ll just have to roll with the chaos a bit. (This is one of the reasons I think it’s worthwhile to consider engineering leadership values as opposed to engaging with the company values overall.)

## Some values I’ve found useful

Because useful values need to be reversible, applicable, and honest, there are no universal values. Particularly successful companies export their stated values to the next generation of companies, which tends to make some values appear universal. [Amazon’s ](https://www.amazon.jobs/content/en/our-workplace/leadership-principles)[_Customer Obsession_](https://www.amazon.jobs/content/en/our-workplace/leadership-principles) is a great example: how many companies formed after Amazon include a value around customer-centricity? Almost all of them! How many of those companies replicate Amazon’s genuine focus on customers? Remarkably few.

Although you cannot simply copy values from one company to another, I do think it’s useful to share some of the values that I’ve found particularly valuable in engineering organizations that I’ve worked in. I’m not arguing you should adopt these specifically, but consider the rationale behind them and whether they might apply to your organization as well:

- **Create capacity (rather than capture it).** This value focuses leaders on creating new capacity from outside the company, rather than fighting internally for existing capacity. Creating capacity makes [the first team mindset](https://lethain.com/first-team/) possible, because it aligns incentives across managers who would otherwise be competing with one another for budget. More discussion on this in [Create capacity rather than capture it](https://lethain.com/create-capacity/).

_Reversible?_ Yes, many organizations take a “fungible headcount” view on shifting teams and individuals, which encourages leaders to capture capacity to complete their goals

_Applicable?_ Yes, highly relevant in any prioritization or headcount planning exercise

- **Default to vendors unless it’s our core competency.** You can also write this as: build versus buy decisions should consider the full implementation and maintenance costs. It is often more intellectually interesting to build commodity solutions than to use existing vendor solutions, but the cost of operating, maintaining and extending those solutions is often much higher. Good build versus buy decisions look beyond the initial implementation cost, and don’t prioritize initial fun over long-term maintenance pain.

_Reversible?_ Yes, for example, early Uber had a strong [Not Invented Here](https://en.wikipedia.org/wiki/Not_invented_here) culture, and rarely used external vendors

_Applicable?_ Yes, relevant to any platform or capability build versus buy decision

- **Follow existing patterns unless there’s an order of magnitude improvement.** Engineering organizations often get caught in a slow burning but endless conflict about when to introduce new programming languages and tooling. It’s exceptionally valuable to have a clear answer to align folks on making these decisions. More discussion on this in [Magnitudes of exploration](https://lethain.com/magnitudes-of-exploration/).

_Reversible?_ Yes, many companies encourage introducing new technology _or_ have an infinitely high bar to adopting a new technology stack

_Applicable?_ Yes, whenever selecting the technology stack for a new project

- **Optimize for the {whole, business unit, team}**. Pick _one_ of these options to help the team understand how they should balance decisions between impacting their team and impacting other teams. It’s also a good example of how values and architecture intersect: you can’t _really_ optimize only for your team if you’re in a monolithic service without causing significant problems elsewhere.

_Reversible?_ Yes, there are successful organizations that optimize for the team, ones that optimize for the company, and so on. As long as the approach is aligned with the technical architecture, all of these are very viable approaches

_Applicable?_ Yes, useful whenever making tradeoffs across team and organization, e.g. selecting the technology stack for a new project

- **Approach conflict with curiosity.** One of my foundational beliefs is that most professional conflict between reasonable people is driven by asymmetric information. If you approach conflict with curiosity, you can quickly learn the missing information and generally make the right decision without conflict.

_Reversible?_ Yes, although admittedly on the weaker side. A given successful company might focus on resolving conflict via subject-matter expertise, data, user feedback, and so on

_Applicable?_ Yes, useful in any scenario with conflict

While reversibility and applicability are universal, you’ll have to evaluate for yourself whether these would be honest values for your company or organization.

## Useful but never magic

Documented values are an excellent way to patch a fixed problem, but are not the solution on their own. Good values document how things _actually work_, they don’t transform your company into a totally different one. If you’re experiencing a cultural crisis in your organization, then you need to first fix that issue: organizational values are not magic.

If you do decide that values are the right tool to use for a problem, they’ll serve you well to the extent that they’re reversible, applicable, and honest.
