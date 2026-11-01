---
title: "You’re Not Taking On Enough Tech Debt | SingulariTea ☕"
notion_id: 30354f1c-7d23-81d4-b0bb-d433e71bd56f
notion_url: https://app.notion.com/p/You-re-Not-Taking-On-Enough-Tech-Debt-SingulariTea-30354f1c7d2381d4b0bbd433e71bd56f
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://singularitea.bearblog.dev/tech-debt/#fnref-1
tags: ["English", "Technical Debt", "Software Development", "Productivity", "Artificial Intelligence (AI)", "Reflection", "Article", "SingulariTea"]
---
**TL;DR - The cost of servicing technical debt is plummeting because of LLMs; assuming coding models keep improving. You're better off taking more technical debt in your projects and bet on the fact that LLMs will clean up the debt in the future.** [1](https://singularitea.bearblog.dev/tech-debt/?utm_source=tldrnewsletter%2F#fn-1)

Through most of my software engineering journey, I understood the concept of technical debt as taking shortcuts while development that bite you at some later point. [Here](https://blog.codinghorror.com/paying-down-your-technical-debt/) [are](https://enterprisersproject.com/article/2020/6/technical-debt-explained-plain-english) [some](https://www.ibm.com/think/topics/technical-debt) other explanations and takes if you're unfamiliar with the term.

A step change in my understanding of the term came when I read [Avery's](https://x.com/apenwarr) [post](https://apenwarr.ca/log/20230605) on extending the tech debt metaphor.

A few ideas from the post that help me make my point:

### Not all (Tech) Debt is bad

> 

> 

The idea being that debt in service of more [insert a metric you care about], is probably prudent to take. This is well understood in software development as Knuth's famous maxim - _premature optimization is the root of all evil_.

This comes with the usual caveat of when and how much debt you should take. It's easier to take shortcuts at the start of the project/company/codebase (smaller principal). The shortcuts shouldn't be so egregious that the time spent on fixing the shortcut (higher rate of interest) is not worth the upside.

### Debt to income ratios

> 

If you can grow your revenue/DAUs/MAUs fast enough, you could probably throw more resources down the line to make up for the shortcuts you took.

## Extending the Metaphor Further

As far as Avery went in extending the analogy, I would like to go further.

### The Risk Free Interest Rate

The risk free rate of return is a theoretical concept on how much return could you get on an asset assuming no underlying risk. It is most commonly proxied by the US Fed Rate (although that could be changing slowly). Most structured debt in the world is directly downstream from this number. Your mortgage is some percentage points plus this rate. You should be buying more risky assets when Fed rates are low and fleeing to low risk assets when the Fed rates are high.

This risk free rate decides how much liquidity/money there is in the system and is the **single most important** number in finance. It is revisited up to 8 times a year and there's a central committee who sense the [vibes](https://www.npr.org/sections/money/2022/05/31/1101774189/fear-the-vibe-shift-are-we-entering-a-recession) of the economy and decide this number. If it seems a bit arbitrary, it is (by design).

When extending this analogy to technology, the risk free rate of return is akin to how much we can expect underlying technologies to improve without any intervention from us. Because the long arc of technology is up and to the right - tech's risk free rate has monotonically fallen over the years.

An example of this was Moore's Law in the '90s and early oughts, CPUs were reliably getting faster every generation and developers were incentivized to work on features over speed optimizations because the hardware would catch up. Another example is software tooling getting better over time meant software development as a vocation was accessible to more people.

Over the years, we have internalized a rate of improvement in underlying technologies. The advent of coding with LLMs, however, challenge our intuitive rate.

### LLMs Lower the Risk Free Rate of Interest (dramatically)

The cost of writing a useful[2](https://singularitea.bearblog.dev/tech-debt/?utm_source=tldrnewsletter%2F#fn-2) line of code is falling every month. The mess we (or our LLMs) create by taking shortcuts in code today is a problem to be tackled by a future LLM.

It seems intuitive but horrid to say - the amount of care we put into writing code should be trending down. I say this for people writing code in service of _something_. Not for people who take joy in writing code for its own sake (me!). _Artisanal Coding_ has its place like handmade goods do[3](https://singularitea.bearblog.dev/tech-debt/?utm_source=tldrnewsletter%2F#fn-3), but that will be a shrinking size of the code we write.

### You Need to Be a (Slightly) Worse Coder

We have all grown up under a regime of what counts as good coding practices. For instance Harold Abelson from Structure and Interpretation of Computer Programs (SICP):

> 

Or, Robert C. Martin:

> 

> 

Or Brian Kernighan:

> 

Or Jeff Attwood:

> 

They are all still true but ...less[4](https://singularitea.bearblog.dev/tech-debt/?utm_source=tldrnewsletter%2F#fn-4) important than they were before. Handcrafting routines, writing helpful comments, having the right level of abstraction are still good principles but less consequential than they were before. Most code going forward _is_ going to be read by machines.

As software engineers trained in the BCE (Before Claude Era), we have ourselves calibrated to taking a certain amount of tech debt allowed by our circumstances. If we're writing mission critical software going on a rocket, we would (understandably, and hopefully) take fewer shortcuts but when writing some front-end help page that 3 other people in the world are going to see, it is okay to take some shortcuts.

We might be due for a recalibration. Not taking shortcuts does come at some cost, and the bar to **not** take a shortcut just went down.

### Redeeming the Vibe Coder (partially)

[The](https://www.databricks.com/blog/passing-security-vibe-check-dangers-vibe-coding) [Vibe](https://www.darkreading.com/application-security/vibe-coding-innovation-demands-vigilance) [Coder](https://www.kaspersky.com/blog/vibe-coding-2025-risks/54584/) [understandably](https://www.csoonline.com/article/4116923/output-from-vibe-coding-tools-prone-to-critical-security-flaws-study-finds.html) [gets](https://devclass.com/2026/01/15/vibe-coded-applications-full-of-security-blunders/) [a](https://www.aikido.dev/blog/vibe-coding-security) [bad](https://www.securityjourney.com/post/10-professional-developers-on-the-true-promise-and-peril-of-vibe-coding) [rap.](https://www.contrastsecurity.com/glossary/vibe-coding) However they might be on to something. As seasoned software engineers have underestimated the fall in the risk free rate, most vibe coders have the opposite problem. They overestimate what the underlying technology can deliver. However, over time (if model improvements, continue), they'll be more right than wrong.

### The Consequences of Long Term Lowering Interest Rates (software's ZIRP era)

When interest rate fall in the economy, capital flies to a lot of stupid things. [Remember](https://www.coindesk.com/opinion/2023/07/05/the-broke-ape-yacht-crash-lessons-for-justin-bieber-and-other-nft-collectors) [those](https://www.cnn.com/2023/02/10/business/crypto-nft-bored-ape-moonpay-lawsuits/index.html) [million](https://www.cnbc.com/2024/05/02/bored-ape-yacht-club-nfts-floor-price-sinks-ceo-announces-layoffs.html) [dollar](https://techstory.in/justin-biebers-1-3-million-ape-is-now-worth-less-than-a-used-honda/) [monkey](https://hypebeast.com/2021/10/bored-ape-yacht-club-nft-3-4-million-record-sothebys-metaverse) [JPEG](https://decrypt.co/84474/bored-ape-yacht-club-ethereum-nft-breaks-sales-record) [crypto](https://www.hollywoodreporter.com/business/business-news/celebrity-promoters-sued-over-bored-ape-nft-endorsements-1235279115/) [scams?](https://fortune.com/2022/12/12/jimmy-fallon-justin-bieber-celebs-sued-bored-ape-nft-promotions/) They were all downstream from low interest rate - the term coined for these ventures was ZIRP - Zero Interest Rate Phenomenon. When interest rates go up, these projects clear out, hopefully leaving behind durable, value accretive businesses.

However, in this case, I don't think the interest rate is going to rise, the amount of crummy software just keeps rising, the enshittification/ensloppification of software has begun, we will need new infrastructure/heuristics to figure out good software from bad.

Notes:
