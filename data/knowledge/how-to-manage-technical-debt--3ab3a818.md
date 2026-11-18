---
title: "How to manage technical debt"
notion_id: 3ab3a818-b0f9-425c-ac11-fd73a114eb66
notion_url: https://app.notion.com/p/How-to-manage-technical-debt-3ab3a818b0f9425cac11fd73a114eb66
last_edited: 2023-04-18T19:01:00.000Z
source_url: https://stayrelevant.globant.com/en/technology/process-optimization/how-to-manage-technical-debt/
tags: ["English", "Programming", "System Design / Software Architecture", "Product Management", "Project Management", "Article", "Globant Blog"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_This is a guest article by Javier Bonnemaison. With decades of experience in Lean/Agile consulting, Javier shares his experience with technical debt and how to use it for competitive advantage._

In a digital world, technical decisions are business decisions. Digital technologies already underpin the vast majority of day-to-day business operations, as well as most product and service innovations. As a result, technical strategy has become an essential component of business strategy, and one of its most challenging aspects is . In this article, we will explore the concept of technical debt and its business implications to shed some light on how technical strategy, embodied in investment decisions, directly affects organizational health and performance.

## What do we mean when we talk about “technical debt”?

The term “technical debt” was coined in 1992 by Ward Cunningham, inventor of the Wiki and a pioneer of Extreme Programming. The challenge that Cunningham faced was explaining the need for refactoring, a technical practice that improves the structure of the code without altering its function, to people who may not be conversant with the idea of iterative and incremental software development (now known as Agile) as a process of continuous knowledge discovery.

After reading George Lakoff’s and Mark Johnson’s seminal book “Metaphors We Live By,” Cunningham was inspired to apply the power of metaphor to explain refactoring to non-technical audiences. He reached for the metaphor of financial debt because he was working on financial software at the time and thought the concept would make sense to his customer. In trying to explain the economics of refactoring, he compared interest on financial debt to the costs incurred by neglecting refactoring, which crucially included quickly increasing reductions in productivity and design optionality.

Unfortunately, there is a major flaw in his choice of metaphor when it comes to refactoring. It has caused endless confusion; _unlike financial debt, which is based on a fungible asset such as money, knowledge can be irretrievably lost _with time or personnel changes_. _In fact, the goals that refactoring intends to achieve, ensuring that the code is understandable and reflects the latest understanding of its purpose, are actually a much closer fit to the metaphors of “hygiene” or “fitness” than to financial debt.

The technical debt metaphor has evolved beyond Cunningham’s original refactoring focus to encompass the general category of technical tasks and capabilities that support the organization without directly contributing to creating and delivering new functionality. A common example is the automation of routine tasks to make time available for more productive activities and reduce manual errors. In this case, the financial debt metaphor is more apt since the costs of delaying automation can be fully addressed once sufficient resources are allocated and invested.

Automation is an example of the larger category of_ technical process debt_, but there are many other technical areas that also fit the debt metaphor well, such as_ quality debt_ (outstanding bugs), application lifecycle debt (updating software versions and commercial licenses), and architectural debt (pushing the limits of non-functional requirements such as performance, availability, or scalability.) These are all examples of tradeoffs that, when deliberate, reflect strategic investment decisions that balance business risks with limited resources.

The main practical difference between both categories of technical debt is investment strategy. Technical practices that are better described by hygiene or fitness metaphors, such as refactoring, require continuous investment, albeit often at relatively modest and stable levels. _Technical practices that can be responsibly deferred without inflicting irreversible damage on the code base or the organization are a better fit for the financial debt metaphor and may receive different levels of investment at different times depending on changing business conditions._

## Why does technical debt matter?

Technical debt matters because it conveys vital information from technical to non-technical stakeholders about important technical capabilities that affect business performance. Actively engaging with technical debt concerns is an excellent way for leaders to collaborate in evolving the company’s management model to compete and thrive in a digital world.

Let’s take a look at some examples of how neglecting or ignoring refactoring adversely affects important business needs and capabilities:

- Developer engagement and productivity: Without refactoring, developers have an increasingly hard time making sense of the code to modify and maintain it, which is frustrating, time-consuming, and error-prone. Critical knowledge is increasingly hidden away inside the minds of individual developers, making them a productivity bottleneck and a knowledge-loss risk. These conditions strongly correlate to an increasingly slower development pace, decreased morale, and increased rotation and attrition.
- Quality and customer satisfaction: When developers have a hard time understanding and trusting the code, it takes them longer to fix bugs in production and restore services, and they are also more likely to introduce new bugs in the process.
- Product or service design optionality: When it is not clear what the code does or how it will respond to changes, any new technical options require research and testing, and developers tend to shy away from all but the simplest of technical design approaches, significantly limiting design optionality. Predictability and productivity also suffer since there is more uncertainty about the technical feasibility of any given feature. As a result, there is more variability in the time it takes to assess and complete work. This is the impact that most directly impairs the ability of an organization to respond to market needs (business agility) adequately.

## Managing Technical Debt effectively

Cracking the technical debt conundrum requires us to challenge conventional management thinking in two critical ways: The first is realizing that technical debt is not a problem to be solved but rather one to be managed. The second is seeing technical debt as a way to achieve strategic competitive advantage rather than as a burden.

_Our key challenge is framing; _developing and delivering digital products and services that are fit for purpose in a sustainable way requires a “both/and” approach. Understanding and producing are not mutually exclusive activities but synergistic components of a single value-generation process.

Barry Johson developed the concept of polarity management in 1975 to tackle this very type of challenge. Johnson realized that many seemingly intractable problems are really open-ended challenges. Managing interdependent pairs of concerns involves constant dynamic balancing to maximize each side’s upsides (positive outcomes) while minimizing the downsides.

| Producing | Understanding |
| --- | --- |
| Values  <br>• Delivering tangible outputs <br>• Executing plans and roadmaps <br>• Time to market <br>• Generating revenue through marketing <br>• Focusing on what we know how to do <br>• Optimizing for utilization  | Values  <br>• Delivering positive outcomes <br>• Making the right things/innovating <br>• Achieving product/market fit <br>• Customer, employee & partner experience <br>• Making things right / Reducing waste <br>• Optimizing for learning and optionality  |
| Fears  <br>• Not building the right things <br>• Building things the wrong way <br>• Overcommitment <br>• Burnout  | Fears  <br>• Not getting enough done <br>• Not being equipped to work this way <br>• Losing the sense of urgency <br>• Running out of time or money  |

Figure 1. Sample Polarity map for Technical Debt

Bringing people together in this exercise helps build empathy, understanding, and alignment within the organization and, coupled with leading indicators and balancing countermeasures, provides an actionable blueprint for managing polarities for optimal results.

Importantly, this approach helps build a foundation for establishing a clear evidence-based link between technical practices and operational outcomes. The final piece of the puzzle of integrating technical strategy into business strategy deployment requires connecting operational outcomes with business impacts. Including technical leaders in this process is critical to ensure coherence, traceability, and improved responsiveness to business and market needs.

## Technical strategy is business strategy

As digital technologies take an increasingly prominent role in how organizations operate and differentiate themselves in the marketplace, it is imperative to develop a good working understanding of the links between technical practices, operational outcomes, and business impacts. Only then can executives start to incorporate technical considerations into their strategic palette confidently.

Leaders can accelerate this process by working closely with technical colleagues and partners. In fact, one of the clearest indicators of a successful digital transformation process is the blurring of the lines between technical and non-technical leadership.

Organizations that understand the value of incorporating technical considerations into their overall business strategy and move with a sense of urgency towards a management model that is more responsive to digital needs will achieve a significant competitive advantage.

## Key Takeaways

- The technical debt metaphor originated as a way to convey the view that software development is a process of knowledge discovery.
- Failing to go back and update the code with new understanding may lead to a loss of knowledge or intent and result in an “interest” burden that, among other things, slows down development and limits design optionality.
- The term “technical debt” has evolved to encompass all technical capabilities and practices that support the organization without directly contributing to the creation and delivery of new functionality.
- Technical debt has a framing challenge: it is not a problem to solve but rather a set of operational capabilities to manage in support of the business strategy.
- Polarity thinking (Both/And) can help manage the tension between producing and understanding that is at the core of the technical debt challenge.
- Managing technical debt requires close collaboration between technical and non-technical stakeholders to establish consensus on clear links between technical practices, operational outcomes, and business impacts.
- Technical debt should no longer be seen as a cost or a burden but rather as a set of technical capabilities that support organizational strategy and can be leveraged to achieve a sustainable competitive advantage.

_Michael Feathers also contributed to this article._
