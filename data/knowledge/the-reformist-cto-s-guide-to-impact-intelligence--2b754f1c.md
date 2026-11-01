---
title: "The Reformist CTO’s Guide to Impact Intelligence"
notion_id: 2b754f1c-7d23-81c2-b72b-cc5967b91336
notion_url: https://app.notion.com/p/The-Reformist-CTO-s-Guide-to-Impact-Intelligence-2b754f1c7d2381c2b72bcc5967b91336
last_edited: 2025-11-26T19:04:00.000Z
source_url: https://martinfowler.com/articles/impact-intel.html
tags: ["English", "Leadership", "Product Management", "Agile", "Change Management", "Article", "Guide", "martinfowler"]
---
The book explains why and how the C-Suite Core should care about this topic. It consists of five parts: (1) Introduction (2) Improving Impact Intelligence (3) Reimagining Initiatives (4) Guidance for low-maturity orgs, and (5) Finale. The article here presents key elements of Part Two although the book gets into greater detail and has more worked examples illustrating simple impact attribution. Besides this, Part Two in the book explores the clash of cultures between businesspeople and data people, and a way to resolve the clash. It also weaves everything together into a framework of eight modules with a suggested adoption sequence. For summary of the other parts, please see the book outline included in this [ free preview](https://www.impactintel.net/Preview_of_Impact_Intelligence_by_Sriram_Narayan.pdf).

Impact Intelligence is the title of my latest [book](https://www.impactintel.net/). It explains how to improve awareness of the business impact of new initiatives. The Classic Enterprise thinks of the expenditure on these initiatives as _discretionary spend_. A software business might account for it as R&D expenditure. Written with a framing of _investment governance_, the book is aimed at the execs who approve investments. They are the ones with the authority to introduce change. They also have the greatest incentive to do so because they are answerable to investors. But they are not the only ones. Tech CXOs have an incentive to push for impact intelligence too.

Consider this. You are a CTO or other tech CXO such as a CIO or CDO (Digital/Data). Your teams take on work prioritized by a Product organization or by a team of business relationship managers (BRM). More than ever, you are being asked to report and improve productivity of your teams. Sometimes, this is part of a budget conversation. A COO or CFO might ask you, “Is increasing the budget the only option? What are we doing to improve developer productivity?” More recently, it has become part of the AI conversation. As in, “Are we using AI to improve developer productivity?”. Or even, “How can we leverage AI to lower the cost per story point?” That’s self-defeating _unit economics_ in overdrive! As in, it aims to optimize a metric that has little to do with business impact. This could, and usually does, backfire.

While it is okay to ensure that everyone pulls their weight, the current developer productivity mania feels a bit much. And it misses the point. This has been stressed [time](https://martinfowler.com/bliki/CannotMeasureProductivity.html) and [again](https://martinfowler.com/bliki/OutcomeOverOutput.html). You might already know this. You know that developer productivity is in the realm of _output_. It matters less than _outcome_ and _impact_. It's of no use if AI improves productivity without making a difference to business outcomes. And that's a real risk for many companies where the correlation between output and outcome is weak.

The question is, how do you convince your COO or CFO to fixate less on productivity and more on overall business impact?

Even if there is no productivity pressure, a tech CXO could still use the guidance here to improve the awareness of business impact of various efforts. Or if you are a product CXO, that's even better. It would be easier to implement the recommendations here if you are on board.

## Impact Trumps Productivity

In factory production, productivity is measured as units produced per hour. In construction, it might be measured as the cost per square foot. In these domains, worker output is tangible, repeatable, and performance is easy to benchmark. Knowledge work, on the other hand, deals in ambiguity, creativity, and non-routine problem-solving. Productivity of knowledge work is harder to quantify and often decoupled from direct business outcomes. More hours or output (e.g., lines of code, sprint velocity, documents written, meetings attended) do not necessarily lead to greater business value. That’s unless you are a service provider and your revenue is purely in terms of billable hours. As a technology leader, you must highlight this. Otherwise, you could get trapped in a vicious cycle. It goes like this.

As part of supporting the business, you continue to deliver new digital products and capabilities. However, the commercial (business) impact of all this delivery is often unclear. This is because impact-feedback loops are absent. Faced with unclear impact, more ideas are executed to move the needle somehow. Spray and pray! A feature factory takes shape. The tech estate balloons.

![image](https://martinfowler.com/articles/impact-intel/image-1.png)

Figure 1: Consequences of Unclear Business Impact

All that new stuff must be kept running. Maintenance (Run, KTLO) costs mount. It limits the share of the budget available for new development (Change, R&D, Innovation). When you ask your COO or CFO for an increase in budget, they ask you to improve developer productivity instead. Or they ask you to justify your demand in terms of business impact. You struggle to provide this justification because of a general deficit of impact intelligence within the organization.

If you’d like to stop getting badgered about developer productivity, you must find a way to steer the conversation in a more constructive direction. Reorient yourself. Pay more attention to the business impact of your teams’ efforts. Help grow impact intelligence. Here’s an introduction.

## Impact Intelligence

_Impact Intelligence_ is the constant awareness of the business impact of initiatives: tech initiatives, R&D initiatives, transformation initiatives, or business initiatives. It entails tracking contribution to _key business metrics_, not just to low-level metrics in proximity to an initiative. [ Figure 2](https://martinfowler.com/articles/impact-intel.html#image-2.jpg) illustrates this with the use of a visual that I call an impact network.

It brings out the inter-linkages between factors that contribute to business impact, directly or indirectly. It is a bit like a _KPI tree_, but it can sometimes be more of a network than a tree. In addition, it follows some conventions to make it more useful. Green, red, blue, and black arrows depict desirable effects, undesirable effects, rollup relationships, and the expected impact of functionality, respectively. Solid and dashed arrows depict direct and inverse relationships. Except for the rollups (in blue), the links don't always represent deterministic relationships. The impact network is a bit like a probabilistic causal model. A few more conventions are laid out in the book.

The bottom row of features, initiatives etc. is a temporary overlay on the impact network which, as noted earlier, is basically a KPI tree1 where every node is a metric or something that can be quantified. I say _temporary_ because the book of work keeps changing while the KPI tree above remains relatively stable.

![image](https://martinfowler.com/articles/impact-intel/image-2.jpg)

Figure 2: An Impact Network with the current Book of Work overlaid.

Typically, the introduction of new features or capabilities moves the needle on product or service metrics directly. Their impact on higher-level metrics is indirect and less certain. Direct or first-order impact, called _proximate impact_, is easier to notice and claim credit for. Indirect (higher order), or _downstream impact_, occurs further down the line and it may be influenced by multiple factors. The examples to follow illustrate this.

The rest of this article features smaller, context-specific subsets of the overall impact network for a business.

### Example #1: A Customer Support Chatbot

What’s the contribution of an AI customer support chatbot to limiting call volume (while maintaining customer satisfaction) in your contact center?

![image](https://martinfowler.com/articles/impact-intel/image-3.jpg)

Figure 3: Downstream Impact of an AI Chatbot

It is not enough anymore to assume success based on mere solution delivery. Or even the number of satisfactory chatbot sessions which [ Figure 3](https://martinfowler.com/articles/impact-intel.html#image-3.jpg) calls _virtual assistant capture_. That’s proximate impact. It’s what the Lean Startup mantra of _build-measure-learn_ aims for typically. However, _downstream impact_ in the form of call savings is what really matters in this case. In general, proximate impact might not be a reliable leading indicator of downstream impact.

A chatbot might be a small initiative in the larger scheme, but small initiatives are a good place to exercise your impact intelligence muscle.

### Example #2: Regulatory Compliance AI assistant

Consider a common workflow in regulatory compliance. A compliance analyst is assigned a case. They study the case, its relevant regulations and any recent changes to them. They then apply their expertise and arrive at a recommendation. A final decision is made after subjecting the recommendation to a number of reviews and approvals depending on the importance or severity of the case. The _Time to Decision_ might be of the order of hours, days or even weeks depending on the case and its industry sector. Slow decisions could adversely affect the business. If it turns out that the analysts are the bottleneck, then perhaps it might help to develop an AI assistant (“Regu Nerd”) to interpret and apply the ever-changing regulations. [ Figure 4](https://martinfowler.com/articles/impact-intel.html#image-4.jpg) shows the impact network for the initiative.

![image](https://martinfowler.com/articles/impact-intel/image-4.jpg)

Figure 4: Impact Network for an AI Interpreter of Regulations

Its proximate impact may be reported in terms of the uptake of the assistant (e.g., prompts per analyst per week), but it is more meaningful to assess the time saved by analysts while processing a case. Any real business impact would arise from an improvement in _Time to Decision_. That’s downstream impact, and it would only come about if the assistant were effective and if the _Time to initial recommendation_ were indeed the bottleneck in the first place.

### Example #3: Email Marketing SaaS

Consider a SaaS business that offers an email marketing solution. Their revenue depends on new subscriptions and renewals. Renewal depends on how useful the solution is to their customers, among other factors like price competitiveness. [ Figure 5](https://martinfowler.com/articles/impact-intel.html#image-5.jpg) shows the relevant section of their impact network.

![image](https://martinfowler.com/articles/impact-intel/image-5.jpg)

Figure 5: Impact Network for an Email Marketing SaaS

The clearest sign of customer success is how much additional revenue a customer could make through the leads generated via the use of this solution. Therefore, the product team keeps adding functionality to improve engagement with emails. For instance, they might decide to personalize the timing of email dispatch as per the recipient’s historical behavior. The implementation uses behavioral heuristics from open/click logs to identify peak engagement windows per contact. This information is fed to their campaign scheduler. What do you think is the measure of success of this feature? If you limit it to _Email Open Rate_ or _Click Through Rate_ you could verify with an A/B test. But that would be proximate impact only.

### Leverage Points

Drawing up an impact network is a common first step. It serves as a commonly understood visual, somewhat like the ubiquitous language of domain driven design. To improve impact intelligence, leaders must address the flaws in their organization’s idea-to-impact cycle ([ Figure 6](https://martinfowler.com/articles/impact-intel.html#image-6.png)). Although it is displayed here as a sequence, iteration makes it a cycle.

Any of the segments of this cycle might be weak but the first (idea selection) and the last (impact measurement & iteration) are particularly relevant for impact intelligence. A lack of rigor here leads to the vicious cycle of spray-and-pray ([ Figure 1](https://martinfowler.com/articles/impact-intel.html#image-1.png)). The segments in the middle are more in the realm of execution or delivery. They contribute more to impact than to impact intelligence.

In systems thinking, leverage points are strategic intervention points within a system where a small shift in one element can produce significant changes in the overall system behavior. [ Figure 6](https://martinfowler.com/articles/impact-intel.html#image-6.png) highlights the two leverage points for impact intelligence: idea selection and impact measurement. However, these two segments typically fall under the remit of business leaders, business relationship managers, or CPOs (Product). On the other hand, you—a tech CXO—are the one under productivity pressure resulting from poor business impact. How might you introduce rigor here?

In theory, you could try talking to the leaders responsible for idea selection and impact measurement. But if they were willing and able, they’d have likely spotted and addressed the problem themselves. The typical Classic Enterprise is not free of politics. Having this conversation in such a place might only result in polite reassurances and nudges not to worry about it as a tech CXO.

This situation is common in places that have grown Product and Engineering as separate functions with their own CXOs or senior vice presidents. Smaller or younger companies have the opportunity to [avoid growing into this dysfunction](https://www.agileorgdesign.com/content/startup-structure-orgdesign-product-engg). But you might be in a company that is well past this orgdesign decision.

We are publishing this article in installments, the next installment will look at five actions to improve impact intelligence, including demand management, impact validation, and using ROP as an alternative to ROI.

To find out when we publish the next installment subscribe to this site's [RSS feed](https://martinfowler.com/feed.atom), or Martin's feeds on [Mastodon](https://toot.thoughtworks.com/@mfowler), [Bluesky](https://bsky.app/profile/martinfowler.com), [LinkedIn](https://www.linkedin.com/in/martin-fowler-com/), or [X](https://twitter.com/martinfowler).

## Additional Resources

To equip yourselves better to appeal to your COO/CFO on this topic, it might be best to read the book which is aimed at them. You could get a sense of its approach by reading the free [ first chapter](https://www.impactintel.net/Preview_of_Impact_Intelligence_by_Sriram_Narayan.pdf) and my article in Armstrong Wolfe's [COO Magazine](https://cdn.armstrongwolfe.com/wp-content/uploads/2025/05/Armstrong-Wolfe-COO-Magazine-Q2-2025-WEB-VERSION.pdf). If you'd like to explore more from a tech lens, you could watch this [podcast](https://www.youtube.com/watch?v=IXauEDLtGuo) by the Tech Lead Journal or this [conference talk](https://www.youtube.com/watch?v=-qlKOAEPggo) I gave earlier this year.

## Acknowledgements

Many thanks to the reviewers: Charles Betz, Dan Leeds, Jeff Foster, Jeff Gardner, Emilia Sherifova, Martin Fowler, Dr.Martin Strunk, Thirukumaran Karunanidhi, Uttam Kini, Andrew Thal, and Chris Chakrit Riddhagni. Additional, special thanks to Martin Fowler for his help and guidance with editing and publishing.

## Footnotes

1: A couple of reviewers said that the impact network reminds them of other formulations. One pointed to cascading OKRs. Another pointed to Gojko Adzic's impact mapping. I guess it could remind you of any formulation that uses a tree-like visual to talk about impact. But a careful reading should convince you that the impact network stands on its own as a specific type of KPI tree. Its seeds were sown in an earlier visual called [Alignment Maps](https://martinfowler.com/bliki/AlignmentMap.html).

**Significant Revisions**
