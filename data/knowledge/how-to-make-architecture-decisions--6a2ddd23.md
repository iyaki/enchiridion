---
title: "How to make architecture decisions"
notion_id: 6a2ddd23-c89b-4125-ba39-7e29c71d5466
notion_url: https://app.notion.com/p/How-to-make-architecture-decisions-6a2ddd23c89b4125ba397e29c71d5466
last_edited: 2024-04-12T16:49:00.000Z
source_url: https://www.ben-morris.com/how-to-make-architecture-decisions/
tags: ["Article", "Ben Morris", "English", "System Design / Software Architecture", "Decision Making"]
---
![image]('https://www.ben-morris.com/images/titles/decisions.jpg')

We want empower engineering teams to make their own decisions, but we should also recognise that autonomous teams cannot solve every problem on their own. Sometimes they may lack the experience, perspective, or technical know-how to drive a decision. Some engineers are reluctant to assert opinions on their colleagues, while others can be a little too eager to do so.

Decision making becomes more difficult as the impact of a decision starts to spread beyond team boundaries. As the scope of a decision increases, more people become involved, different agendas come into play, and it can be difficult to broker agreement. The trade-offs tend to become more nuanced, and the risks are not always immediately obvious.

Larger scale decisions with wide impact need to be made with a sense of expertise and experience, as well as organisational sensitivity. The bigger the impact, the more an experienced architect can add value by helping to foresee any likely outcomes and navigating the potential consequences.

## Knowing what decisions to make

Technology decisions don't have to be made explicitly. You can decide not to make a formal decision, though you need to recognise that this allows lots of smaller decisions to accumulate over time. You need to decide whether this matters, i.e. are you happy to let a decision evolve in a more informal way?

For example, if you choose not provide guidance around which database to use, then your engineering teams will make their own minds up. They may eventually coalesce on a single solution, or they may use different solutions to reflect the varied types of workloads in play. The point is that in the absence of a formal decision, your approach to database standardisation will be determined by lots of individual engineering decisions.

It's up to the architect to foresee these outcomes and decide whether they are comfortable with them. Does it matter where this lands, or does it make sense to constrain the available choices in some way? Will this make life easier in the long term and help you to manage costs, or will it merely undermine agility without delivxering any tangible benefits? Ultimately, it takes wisdom and experience to know _when_ to make a decision.

Decisions can constrain teams and reduce the available options, but this can be a positive thing. Constraints can enable innovation by focusing future efforts on a narrower range of options. This is the sense of creating "guard rails" for engineering teams, where the more generic challenges of engineering are regarded as "solved problems". This allows teams to focus on those decisions that are more directly relevant to their domain.

One important aspect of this is whether decisions are _reversible_. Jeff Bezos famously made a distinction between "Type 1" and "Type 2" decisions, where the former are irreversible decisions that he likened to "one-way doors". This is basic risk management: the more consequential and irreversible the decision, the more care you should take in making it.

The catch is that it's not always that easy to spot a consequential decision. It's also surprisingly easy for apparently reversible decisions to become more consequential than expected. You can manage the risk by making changes smaller and getting feedback more quickly, making it easier to reverse out of those decisions that aren't panning out.

Which decisions should architects be involved in? If we want empowered and autonomous engineering teams, then architects shouldn't be involved in their day-to-day work. They should only get involved in decisions where the impact will be felt beyond the team's boundaries. Again, this is driven by risk management. If the impact of a decision is limited to an individual engineering team, then potential problems tend to be smaller and more manageable.

## When to make decisions

If decisions close off options, then there is an obvious danger in making decisions too early. You can constrain teams unnecessarily by making decisions based on an immature understanding of the problem domain. Sometimes no decision is the wisest choice. If you're not comfortable with a decision, then it might make sense to explore the problem domain to resolve the source of discomfort.

Mary and Tom Poppendieck identified the ["](https://www.ben-morris.com/lean-developments-last-responsible-moment-should-address-uncertainty-not-justify-procrastination/)[_last responsible moment_](https://www.ben-morris.com/lean-developments-last-responsible-moment-should-address-uncertainty-not-justify-procrastination/)["](https://www.ben-morris.com/lean-developments-last-responsible-moment-should-address-uncertainty-not-justify-procrastination/) as a guiding principle for decision making. They defined this as that point at which "_failing to make a decision eliminates an important alternative_". The implication is that there is a "point of no return" at which failing to make an explicit decision can create its own constraints. You should make these decisions explicit, rather than just have them happen to you through inaction.

One problem with the "last responsible moment" is that it's often difficult to recognise when that moment has arrived. It tends to be easier to spot when it has already passed. It can also become an excuse for procrastination. There is never a golden moment where the clouds part to reveal the perfect decision. There are always trade-offs and there is always uncertainty, so decision making requires a degree of courage.

If every decision is delayed until that "last responsible moment" then we risk creating an amorphous decision space that is riven with uncertainty. Too many open decisions create mental clutter that can impede progress. Decisions are built on top of other decisions, and you need a solid foundation and a clear "direction of travel". Engineering teams tend to thrive on certainty, so you shouldn't be afraid to rule out options and opportunities.

## Visibility and collaboration

Although everybody tends to agree on the need for strong decision making and "guard rails" to guide engineering teams, the tricky part comes in recognising where these guard rails should be. You need to take a collaborative and inclusive approach to decision making to encourage buy-in and ensure relevance.

Most organisations have natural leaders that emerge informally in the technology space. These natural leaders are not exclusive to architecture and can be found in engineering, infrastructure, or even product teams. Building a community of natural leaders is vital for establishing consensus and getting buy-in for decisions. A more formal TOGAF-style "architecture board" can help to bake this network of influencers more formally into decision-making.

You also need to make sure that decisions are visible. Architectures usually emerge from the result of a series of imperfect decisions. None of us are reliable witnesses and circumstances change. A log of [lightweight decision records](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions) can help to create a "shared memory" for architecture that documents how the various trade-offs were understood _at the time_. It can also help to keep decisions honest by forcing you to explain how decisions are aligned with the wider business and technical objectives.

Some conflict is inevitable, but that shouldn't make you shy away from making decisions. If somebody really does stand to lose from a decision, then you might not be in the right place. The problem domain might need reframing, or you need to do more work to establish common ground and a shared understanding. That said, so long as decisions are made on an inclusive basis, and the reasons behind them are clearly articulated, then any conflict around them is usually manageable. For the most part, engineering teams usually welcome clear and credible decision making.
