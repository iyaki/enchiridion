---
title: "Notes on software quality"
notion_id: 36d54f1c-7d23-8181-8c0f-daaa5ab28143
notion_url: https://app.notion.com/p/Notes-on-software-quality-36d54f1c7d2381818c0fdaaa5ab28143
last_edited: 2026-05-27T03:34:00.000Z
source_url: https://anthonyhobday.com/blog/20260410
tags: ["English", "Quality", "Software Development", "Testing", "Principles", "Management", "Article", "anthonyhobday"]
---
Jump to:

- [How I think about quality](https://anthonyhobday.com/blog/20260410/#how-i-think-about-quality)
- [Universal signals of quality](https://anthonyhobday.com/blog/20260410/#universal-signals-of-quality)
- [The six signals of quality in software](https://anthonyhobday.com/blog/20260410/#the-six-signals-of-quality-in-software)
- [The benefits of software quality](https://anthonyhobday.com/blog/20260410/#the-benefits-of-software-quality)
- [Beliefs about quality I want to disprove](https://anthonyhobday.com/blog/20260410/#beliefs-about-quality-i-want-to-disprove)
- [Quality is impossible at scale](https://anthonyhobday.com/blog/20260410/#quality-is-impossible-at-scale)
- [Anecdotal evidence that quality is impossible at scale](https://anthonyhobday.com/blog/20260410/#anecdotal-evidence-that-quality-is-impossible-at-scale)
- [Companies with dedicated quality efforts](https://anthonyhobday.com/blog/20260410/#companies-with-dedicated-quality-efforts)

## How I think about quality

### Quality is the absence of problems

“The absence of problems” is the best definition I can come up with for quality.

The most practical way to measure quality is to test with many different people, and have many experts take a look. If thorough testing and 100 experts can’t find a problem, the thing is probably perfect.

### Quality is a spectrum up to “perfection”

Perfection is impossible. The agreed-upon best-designed software in the industry has noticeable problems. But the closer you can get, the better.

The closer you try to get to perfection, the harder it is. This is the same “diminishing returns” effect you’ll see anywhere quality is involved.

### Quality relies on the organisation

Leadership and culture both define and allow quality. Even if they don’t believe they do. Even if they haven’t thought about it. Their actions and beliefs will either enable or limit quality.

Quality is a result of ability and appetite. Does the organisation have people who can produce high quality (ability)? And does the organisation allow quality to be achieved (appetite)?

~100% of the time I’ve seen an organisation make high quality software, it is because the leaders wanted it. If the leaders don’t want quality, it will be much harder.

But don’t forget that whatever you work on, there is an amount of quality you personally can achieve. Even in the most rigid design-system-controlled interface, there are still higher quality and lower quality choices you can make.

An organisation can set objective measures of quality (e.g. bugs reported). Hopefully those objective measures don’t blind the organisation to the subjective measures (e.g. “does an expert think this is high quality?”).

### Quality gets harder with scale

The larger the organisation or software, the harder quality is. High quality is impossible past a certain point of scale. Some organisations are incapable of making high quality software.

Importantly, this is a natural result of scale. It’s not a complaint or a problem to be solved. It cannot be solved. It’s a trade-off, like almost everything else in software.

### Quality is not necessary, but ideal

I wish every organisation cared about quality. It’s clear that many of them don’t, and don’t need to. Many commercially successful organisations don’t prioritise quality.

But quality is ideal. Everyone wants to live in a higher-quality world. It feels good to create quality. Things last longer if they’re higher quality. And over the long term, commercial success is harder if people realise your brand doesn’t represent quality.

## Universal signals of quality

These are the signals of quality that I’ve found apply to everything. People and artefacts.

1. **Appearance:** In (over)simple terms, the more beautiful something is, the higher quality we think it is. I assume this is because beauty requires effort, and so does quality in general. Any evidence that we have put care into something is evidence of further care.
2. **Association:** We judge whether something is high quality by what it is associated with. A common approach to this is social proof, where someone or something is associated with other someones or somethings. We judge a thing by the company it keeps.
3. **Cost:** If someone or something costs more, we might assume that is because it is worth more. But “cost” does not have to mean “money”. If something takes more time, for example, we might assume it is of higher quality.
4. **Performance:** Performance could be quantitative like “top speed”, or qualitative like “how does it make me feel”. It could be results-focused like “did it help me complete my task”. In simple terms, if appearance is “how it looks”, performance is “how it does”.

## The six signals of quality in software

1. **Reliability:** Does the software technically perform as expected, all of the time? This means no bugs, no downtime, no errors, and so on.
2. **Speed:** Does the software react to input instantly, or as near as practically possible? For those tasks that must take longer, are they as fast as possible?
3. **Clarity:** Does the person using the software understand everything?
4. **Efficacy:** Can the person do what they need to do with the software?
5. **Efficiency:** Can the person do what they need to do with the software as easily as possible?
6. **Beauty:** Is the software as aesthetically pleasing as possible?

## The benefits of software quality

Some might argue that quality is an obvious good. That the reasons quality is important are naturally understood by everyone. But we’ve still got a lot of crap software, so 🤷.

1. **Attract employees:** People want to work for companies that care about quality.
2. **Satisfaction as creators:** Everyone involved will feel better about themselves.
3. **Fight fewer fires:** Higher quality means fewer problems to fix afterwards. Fewer bugs, fewer reliability issues, etc.
4. **Fight software entropy:** All software suffers as it grows. It becomes bloated, harder to use and look at. A commitment to quality pushes back against entropy.
5. **Less employee churn:** Care for software generally means care for the people who make it. It encourages skill growth.
6. **Easier to sell:** If quality is part of your company identity, the product is easier to sell. Word of mouth will do some of the marketing for you.
7. **Attract users:** Quality attracts people. They’re more likely to use your software.
8. **Users become fans:** Usable, effective, satisfying software turns users into fans. Fans want to support the company, forgive it more easily, and are happy to help it improve.
9. **A competitive moat:** Quality is hard. Your competitors probably won’t put the effort in.
10. **Make more money:** People universally pay more for quality.
11. **Beauty improves society:** Beautiful artefacts universally improve how we live.
12. **Inspire quality in others:** Other people and companies will be inspired by your work.
13. **Leave a legacy:** No product lasts forever. But people talk about good software for decades.

## Beliefs about quality I want to disprove

I’ve had software interface quality on my mind for most of my career. Here are some things I believe are true. I would prefer that they are not.

1. Quality is easiest when one person can keep the whole interface in their head.
2. The more people who work on some software, the harder high quality is.
3. The more high quality things you have in your life, the better your life is.
4. The lower the quality of your product, the more people will abandon it.
5. More people will talk about your product if quality is high.
6. Quality depends on permission from whoever is most powerful at an organisation.
7. Quality cannot be a bottom-up effort, driven only by the people who do the work.
8. High quality always requires a sacrifice of something, e.g. scope, money, growth, or time.
9. High quality always requires someone to have a strong personal interest/obsession with it.
10. A focus on quality means a sacrifice of growth.
11. Quality is both harder and more important for software that is used often.
12. Quality is a competitive moat because so few organisations focus on it.
13. You should ideally hire good people if you want high quality.
14. If you cannot hire good people you should be willing to iterate a lot.
15. You need to iterate a lot more than expected to get high quality.
16. Some organisational cultures are incapable of high quality.
17. The larger the organisation, the harder quality becomes. At a certain size of organisation high quality is impossible.
18. The larger the software, the harder quality becomes. At a certain size of software high quality is impossible.
19. If something is low quality in some software, it’s usually a reliable sign that other things are low quality.
20. Some types of quality are not always immediately obvious.
21. Perfection is impossible.
22. Products and companies can fail because they focused on quality at the expense of other things.
23. Quality requires that everything be given a lot of attention.
24. Even though everyone agrees quality is positive, it is hard to convince people to focus on it.
25. Some people want high quality in their personal lives, but don’t care enough about it in their professional lives.
26. Many commercial goals hurt quality, e.g. growth.
27. Quality is a constant effort. Problems can slip in quietly, especially as changes are made.
28. The only reliable way to measure quality is to have lots of experts look at the software.
29. Focus on quality has diminishing returns. It takes much more effort to achieve “90%” quality than it did to achieve “80%” quality.
30. The systems in place to help create software (e.g. design systems) can also hurt efforts to focus on quality.
31. People associate expressive visual styles with visual quality. They may expect expressive visual styles in high quality software even when it’s not needed.
32. People associate high cost with quality.
33. People associate hype/community interest with quality.
34. There are many ways to compete that cost less than quality, e.g. features, price, sales experience.
35. People may leave a company because they cannot focus on quality as much as they’d like to.
36. Novelty is more immediately attractive to people than high quality.
37. High quality is often not noticed until later.
38. Organisations and people will almost never tell the truth about the low quality of their products.

## Quality is impossible at scale

Good interface design becomes harder as software grows. Eventually, world-class interface design is impossible. Here’s are some thoughts.

### There are too many relationships to manage

Good interface design is coherent. “Coherent” means the relationships between all of the things in the interface work well. But as the number of things goes up, the number of relationships goes up even faster. Eventually it’s impossible for people to properly consider all of those relationships.

This principle also applies to people. The more people who work on software, the harder it is to manage the relationships between all of them. Things are lost in translation. People disagree on the goal, and pull in different directions.

You can manage this with stricter processes. But eventually you’ll spend more time on process than on design.

### Some people don’t care enough

The more people you hire, the more likely you are to hire people who don’t care enough about good interface design. Good interface design needs to be valued by everyone who can affect the work. That includes developers, designers, product managers, and often the CEO.

### Some commercial goals hurt good design

If a company wants some numbers to go up for commercial reasons, this often hurts interface design quality. For example, adding more advertising to a website can help a company to become profitable. But it can also make the interface worse.

One way to grow is to add features that people request. But often each added feature makes the interface worse. This is hard to avoid, because of the first point above.

## Anecdotal evidence that quality is impossible at scale

Patrick McKenzie [tweeted about the culture of quality at Stripe](https://twitter.com/patio11/status/1176866754682707969?s=20&t=7ov-AYTrfpj1_yvA8U33LQ). Patrick ended with “none of it is sufficient. We are actively dissatisfied with where quality is at the moment”.

Benji Taylor of Family and Honk [said that](https://twitter.com/benjitaylor/status/1571994317014863872?s=20&t=O9yvjV9aHuuddP6PU55xpw) “big corporate software seems to be getting buggier each year, while indie software [seems] to be getting smoother and more reliable”.

Some notable people replied to Benji’s tweet:

> 

In their article [Optimizing For Feelings](https://browsercompany.substack.com/p/optimizing-for-feelings), The Browser company says that “as our everyday software tools and media became global for the first time, the hand of the artist gave way to the whims of the algorithm. And our software became one-size-fits-all in a world full of so many different people. All our opinions, beliefs, and ideas got averaged out”.

David Khourshid, founder of Stately, [says that](https://twitter.com/DavidKPiano/status/1571838435656192000?s=20&t=oGG_OH9QyU3LGB2hCo6eew) “adding even a single ‘feature’ (new states + transitions) to an app can introduce an exponentially greater number of ways the app can be used”.

Benji Taylor [tweeted that](https://twitter.com/benjitaylor/status/1509242039262593025?s=20&t=3T3ZkLvoiLYyMCmOhGEb5A) “something all these products have in common is that they are … built by small teams… Too often products lose their soul once they reach a certain scale”.

Paul Stamatiou [said that](https://twitter.com/Stammy/status/1558122812988088322?s=20&t=c3YScLRvsv7fZqZf3vR0Mg) “responses here so far confirm my theory that… high-quality software products are only made by small teams”.

Nathan A. Curtis [said](https://twitter.com/nathanacurtis/status/1517194393836560388?s=20&t=fk5JnUuHPGvuvgSYhHlWsQ) “I love Figma and value the experimentation they’re doing to add more and more. But I can sense it… their UI is ‘Adobe Creative Suite’ing. And that carries a usability cost.” Note that he said this in April of 2022, months before Adobe acquired Figma.

Repeating an argument Steve Jobs made, [Marty Cagan said](https://twitter.com/lennysan/status/1561876506548113408?s=20&t=mR4WJOPNyNR0oRQrVSBxHg), “as a company gets bigger, product historically became less important… Good product people don’t want to work there anymore… They go to a company that values product. I think that’s a better explanation than any other I’ve heard.”

Steve Ruiz [said](https://twitter.com/steveruizok/status/1571559529350942721?s=20&t=2Qm-i57bcMFxiOTZTOIi9w), “unfortunately taste doesn’t scale”.

Packy McCormick [wrote about this](https://www.notboring.co/p/indistinguishable-from-magic). Here are some good bits:

> 

Paul Graham [said](https://twitter.com/paulg/status/1560349112922279936?s=20&t=_2skSuvZBzNGHx2c3ehEGg), “One theme I’ve noticed talking to people in Silicon Valley this trip is that Google has become terminally bureaucratic”.

[On the Tim Feriss show](https://tim.blog/2018/12/24/the-tim-ferriss-show-patrick-collison/), Stripe co-founder Patrick Collison said “Just think of any big, major company … for whatever reason, they can’t turn capital into good software. And it would be immensely valuable for them if they could. But they can’t. Or at least, they don’t. And I don’t think it’s for lack of trying or lack of realizing this. And so I think actually, small companies don’t realize how much of an upper hand they have here where if they can create a product that is so much better … it’s going to be really freaking hard for a big company to effectively compete because, again, this organizational transformation into being good at software is just so profoundly hard.”

Dan Luu [has said](https://danluu.com/culture/) that “while it’s not exactly common, there are plenty of small firms out there with a culture of excellence that generally persists without heavyweight processes or big incentives”, but that “this doesn’t work at big tech companies since they’ve all gone through a hypergrowth period where it’s impossible to maintain such extreme (by mainstream standards) cultural values.”

Kevin Yien [says that](https://twitter.com/kevinyien/status/1577102698062962689?s=20&t=ZfcKRVc-DyHKIVCIIvG1MQ) “you know you are no longer a startup when your [organisationl] structure prevents well-intentioned people from fixing problems”.

Cemre Güngör, a product manager at Instagram, [says that](https://twitter.com/gem_ray/status/1412915293274329094?s=20&t=iK3Zb15dzlafbFB5eGbAVg) “Big tech companies hire expensive senior designers but still end up with poor design because it’s much harder to make structural changes to incentives and slavish metric chasing. There’s great talent in most companies already. They’re just stifled by the organization around them.”

Craig Mod [said](https://craigmod.com/essays/software_slump/) “… the three primary pieces of software on macOS are probably Finder, Safari, and Mail. To have two of these show signs of instability is … just weird. It shouldn’t happen, especially when these are new, critical bugs in decades-old programs. It makes you wonder … what’s broken with the development cycle to allow for these bugs to ship… we’re seeing once-reliable applications suffer from feature creep and bloat. Perhaps this is endemic to the very nature of public companies and their conflation of features with user growth? For example: Dropbox has gone from a svelte, hyper-reliable file syncing service to a bloated curiosity that pegs the cpu at 200% for unclear reasons.”

The Financial Times’ Tim Bradshaw [said](https://www.ft.com/content/9ab775cc-d596-4e0b-aa3e-34e9111ec431) “… developers that respond by bundling ever more features into a super app risk frustrating users. Single-purpose apps are faster and easier to navigate. By contrast, the super-app approach involves cramming as many features as possible into one bloated piece of software… Spotify’s interface, for instance, has become cluttered and confusing since it added podcasts… This is my real issue with super apps: most solve a problem for the company, not the customer.”

Adam Michela, on the topic of Airbnb’s misleading pricing design, [says](https://twitter.com/soopa/status/1579596501424406528?s=12&t=n8XUHFkK4bjsMYR47rZq2w) “the problem is that Airbnb’s chief metric is Nights Booked (units sold) and manipulating users with basic pricing psychology moves units. It takes a Jobs-like chief to prioritize user experience over business performance.”

Adam Michela also [said that](https://twitter.com/soopa/status/1583476515634716673?s=20&t=ELkal_-ek_OLKknqVggVwg) “We’ve helped hundreds of 2-200 person startups build products … I can confidently agree that the ~30 person moment (5-10 engineers and 2 designers) is the quality sweet spot.”

Dennis Brotzky, coufounder of Fey, [says that](https://twitter.com/brotzky_/status/1580943383841083392?s=20&t=NZ36_x9Bq2CaFaDshl8qqw) “the smallest and most focused teams deliver the best results. Keep your team small and give them time and freedom.”

Michael Karnjanaprakorn, who founded Skillshare, [says that](https://twitter.com/mikekarnj/status/1581287479436488705?s=20&t=-0LxZk9WVOMnKNoVU-uFgQ) “quality and speed goes down as team size gets bigger”.

Rasmus Andersson [says that](https://twitter.com/rsms/status/1583335925056086016?s=20&t=ELkal_-ek_OLKknqVggVwg) “I’m convinced that—generally speaking—there’s an inverse correlation between number of people working on a software product and its quality (from end-user perspective)”.

Kyle Rubenok, a PM at Microsoft, [says that](https://twitter.com/krubenok/status/1583210021068103680?s=20&t=ELkal_-ek_OLKknqVggVwg) “the best iOS design is exclusively in indie/small team apps today.”

Cameron Moll [says](https://twitter.com/cameronmoll/status/1584908753274634242?s=20&t=g63b7MWNAJ0C8wBA4k7SZg) “quality and craft don’t seem to scale as well as headcount.”

Box Baxley [says](https://twitter.com/ThisIsBobBaxley/status/1584923766244589570?s=20&t=g63b7MWNAJ0C8wBA4k7SZg) “Nothing about creativity scales—and it never has. Production scales. Creativity doesn’t.”

Hardik Pandya [says](https://twitter.com/hvpandya/status/1587676031921504257?s=20&t=ByQQTnB-jFJd2OXKjMj-3Q), “Don’t be afraid of large teams. Be afraid of small & determined teams.”

Packy McCormick [says](https://www.notboring.co/p/capital-and-taste) “taste is a fickle thing on which to hang the fate of a company. At some point, growth pressures can get in the way of making sure that every detail is perfect, that every hire is great. Taste, like trust, takes a lifetime to build and minutes to destroy.”

Jasmine Friedl [said](https://uxdesign.cc/measuring-design-quality-with-heuristics-44857efa514) “while these [measures of quality] worked for us, we acknowledged that it wouldn’t scale much past our small team; for example, a team with hundreds of designers shipping daily wouldn’t likely have the bandwidth for manual reviews.”

Paul Stamatiou wrote an [article about craft](https://paulstamatiou.com/craft/) which has some gems in it:

> 

In a 1995 interview for TV show “Triumph of the Nerds”, Steve jobs said “It turns out the same thing can happen in technology companies that get monopolies, like IBM or Xerox. If you were a product person at IBM or Xerox, so you make a better copier or computer. So what? When you have monopoly market share, the company’s not any more successful. So the people that can make the company more successful are sales and marketing people, and they end up running the companies. And the product people get driven out of the decision making forums, and the companies forget what it means to make great products. The product sensibility and the product genius that brought them to that monopolistic position gets rotted out by people running these companies that have no conception of a good product versus a bad product. They have no conception of the craftsmanship that’s required to take a good idea and turn it into a good product. And they really have no feeling in their hearts, usually, about wanting to really help the customers.”

Gergely Orosz [says](https://twitter.com/GergelyOrosz/status/1504366157805101056?s=20&t=YZt87iLR2ApRlrm_WdCEZw) “Very few if any people and teams [at a large tech company] get rewarded for fixing low impact issues impacting 0.002% of users. Also few get punished by not doing so. At any Big Tech, incentives that drive promotions and performance reviews end up driving a surprisingly large amount of decisions. Most of these incentives are impact, revenue generation at scale & savings at scale. Fixing small issues fits none of these. In fact, fixing one-off issues is one of the most expensive things to do. It would - and does! - pull profits down. This is bad for business.”

Dan Luu [says](https://danluu.com/nothing-works/) “The pervasiveness of … technical decisions made without serious technical consideration, is a major reason that the selection pressure on companies to make good products is so weak. There is some pressure, but it’s noisy enough that successful companies often route around making a product that works, like in the Mongo example from above, where Mongo’s decision to loudly repeat demonstrably bogus performance claims and making demonstrably false correctness claims was, from a business standpoint, superior to focusing on actual correctness and performance; by focusing their resources where it mattered for the business, they managed to outcompete companies that made the mistake of devoting serious resources to performance and correctness.”

Arvind Narayanan [says](https://twitter.com/random_walker/status/1182636136843354112?s=20&t=60rPs-IIdDOu3lip-R5O6A) “some products are sold directly to the end user, and are forced to prioritize usability. Other products are sold to an intermediary whose concerns are typically different from the user’s needs. Such products don’t HAVE to end up as unusable garbage, but usually do.”

Nick Lockwood [says](https://twitter.com/nicklockwood/status/1534827111877312513?s=20&t=LfjcGQG9Ze9IYd5E1hWLUQ) “development is like Zeno's paradox, where effort rises exponentially. You never actually reach completion, you just get ever nearer to it, with each unit of effort getting you fractionally closer… this fact applies *fractally*, to every feature. At any point in a product's lifecycle, it will be the same (or more) effort to fix a small bug or papercut than to implement a major new feature (or rather, to implement 80% of it, minus the last few bugs and papercuts).”

Alexandr Wang [says](https://twitter.com/alexandr_wang/status/1594062781673721856?s=20&t=00CKJLccHetOTmKZ5Ql2uA) “Unintuitive fact of engineering: Big teams (>8 people) ship less than small teams. Not only less per capita—less overall!”

On Twitter, Dan Luu [said](https://twitter.com/danluu/status/1603952302897131520?s=20&t=n52BQUC33hjl3FkryTUpGg) “… a sentiment I’ve heard from literally all of the highest impact/most effective people I’ve talked to at large companies: You can make a big difference, but you’re constantly fighting a self-sabotaging organization”.

In his article [Tiktok’s enshittification](https://pluralistic.net/2023/01/21/potemkin-ai/#hey-guys), Cory Doctorow said “here is how platforms die: first, they are good to their users; then they abuse their users to make things better for their business customers; finally, they abuse those business customers to claw back all the value for themselves. Then, they die”.

George Kedenburg III wrote an article, “[The Cost of Craft](https://gk3fyi.substack.com/p/the-cost-of-craft)”, on this topic:

> 

Patrick McKenzie [wrote that](https://twitter.com/patio11/status/1648716242885435392?s=20) “… as software companies get larger … they increasingly make it difficult to get relatively simple software done, which leads to surprising results … You just need a landing page, but the process of getting a landing page done requires three quarters of lead time and sixteen reviews, and the Landing Page as a Service fits on a purchasing card and takes two hours…”

Adam Stoddard [wrote about craft at scale](https://aaadaaam.com/notes/craft-at-scale/), and said that “Funded startups and public tech companies need to keep the high-growth engine running, which inevitably leads progressively larger and more complex product offerings that attempt to attract more and more users. No single team can deliver on that, which leads directly to disjointed and/or competing teams, and a steady erosion of vision that’s backfilled by process.”

Joel Beukelman [said](https://twitter.com/_bklmn/status/1664279030861017088?s=20), “After ~8 years at Google I have a deep respect for product at scale, but my $.02 is there’s some sparkle lost when a team gets too big.”

Chief Product and Engineering Officer Claire Vo [wrote a post](https://x.com/clairevo/status/1705637278226506135?s=20) on this topic. Some excerpts:

> 

Karri Saarinen, co-founder of Linear, [wrote a post about quality](https://x.com/karrisaarinen/status/1710331542369914903?s=20) and said, “As companies grow, they often have a hard time with quality, and usually just give up on it. [The] main reason is that quality is something which cannot be easily measured or defined. As the companies scale, the way they operate or make decisions [is] based more on measurements.”

In [a conversation with Linear about quality](https://linear.app/quality/01), Dick Costolo said, “The best product teams will stay as small as possible, as long as possible. It just becomes too difficult for the product to look and feel like the outcome of a singular simple vision for what the future could be as you layer more and more and more people into the thing.”

## Companies with dedicated quality efforts

Many companies say they care about quality, but don’t have specific efforts dedicated to it. Here are some that do.

### Automattic

Automattic employs a Chief Quality Officer. Here’s what Lance Willett told me about the role:

> 

You can read more about Lance’s role [here](https://fullstackleader.blog/2025/07/14/head-of-quality/).

### GitLab

GitLab has a “UX Paper Cuts” team, which they describe this way:

> 

### HubSpot

HubSpot created a “Cohesion Studio”, which they describe this way:

> 

### Linear

In 2022, [Linear had a month dedicated](https://linear.app/changelog/2022-12-01-polishing-season-2022) to interface design polish: “For the rest of the year, polishing will be our sole focus. No new feature work, just quality of life improvements.”

They also have a permanent policy related to bug fixes:

> 

Finally, for changes that aren’t bugs, they have “Quality Wednesdays”:

> 

Ron Goldin and some other people at Meta introduced quality metrics:

> 

### Microsoft

After public outcry about the quality of Windows 11, Microsoft announced a dedicated effort to improve:

> 

### Operate

> 

### Shopify

Shopify started a team to focus on quality:

> 

### Sonos

Sonos has appointed a “quality ombudsman”.

> 

### Stripe

Stripe hired Malthe Sigurdsson as “Head of Craft”:

> 

### Zed

[Twice a year](https://x.com/zeddotdev/status/2001008164301627903?s=20), Zed has a “quality week”. Here’s how they describe it on their website:

> 

## Interface quality of life improvements

Here are some quality of life improvements you can make to an interface.

1. Generous mouse paths in nested menus, where the menu doesn’t close if the mouse isn’t always within its boundaries.
2. Coyote time for keyboard shortcuts, where a multi-key shortcut isn’t cancelled as soon as you release the wrong key first.
3. Stretch/bounce motion effects in scrolling list views, to help show that someone has reached the end of the list.
4. Smooth and fast expand and collapse animations for all elements that expand and collapse (as opposed to no animation).
5. Starting expansion and collapse animations from the direction of the element that triggered the animation, so that it feels as if it’s expanding away from that element.
6. Smooth and fast animated movement for elements that move from one place to another (as opposed to no animation).
7. Animation to let you know that elements in the edge of your vision have appeared or changed.
8. Menus and other pop-up elements open ~instantly but take slightly longer to fade away.
9. Optical adjustment of visually unbalanced elements so they look properly aligned.
10. If an element that contains text is changed to an “edit” mode, keep the text in the same place as the mode changes.
11. Don’t move elements as the result of an interaction near or with that element, unless that’s the interaction’s purpose.
12. If elements are added or removed to a list, maintain the subjective scroll position so that elements visually stay in the same place.
13. Open context menus so that the menu actions are always in the same place relative to the cursor.
14. On mobiles open the correct type of keyboard for the field selected.
15. If clicking something opens a new interface where the mouse is, don’t put something risky under the current mouse cursor position, in case the person clicks again by accident.
16. Object permanence: elements that feature in more than one view persist throughout transitions from one view to another, even if they move or change shape/size.
17. Unless explicitly requested, information entered by a person is never lost if that person leaves the screen.
18. If a moveable element (e.g. a scrollable horizontal stack of containers) has natural positions, snap it into those positions.
19. Show changes to something in real time, or if the thing is not visible, show a preview of the changes as options are changed.
20. Delay opening e.g. a tooltip, but if the mouse then moves to a neighbouring element which also has a tooltip, show the new tooltip instantly.
