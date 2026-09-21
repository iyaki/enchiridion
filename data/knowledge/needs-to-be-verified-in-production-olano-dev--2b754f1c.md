---
title: "“Needs to be verified in production” | olano.dev"
notion_id: 2b754f1c-7d23-815f-a712-cddc0111d80a
notion_url: https://app.notion.com/p/Needs-to-be-verified-in-production-olano-dev-2b754f1c7d23815fa712cddc0111d80a
last_edited: 2025-11-26T19:08:00.000Z
source_url: https://olano.dev/blog/verified-in-production/
tags: ["Javier Garzas Blog", "English", "Agile", "Product Management", "Team Management", "Communication", "Testing", "Productivity", "Article"]
---
Some developers think that their job is to close Jira tickets, others that it’s solving business problems. Let’s call these mode 1 (process-oriented) and mode 2 (problem-oriented)[1](https://olano.dev/blog/verified-in-production/#footnote-1). While I’m inclined to think that the second mode is more effective and makes my own work more enjoyable, both are valid and honest. What’s more, most of us switch modes based on our perceived relevance of the task at hand, our interest in it, our current level of energy, etc.

Whether one mode of operation is preferable or even possible at a given organization depends on the prevailing culture and the business processes in place. A healthy organization will ideally be flexible enough to accommodate both, but if it accepts or encourages mode 1 development, it’s necessary that the processes guarantee that a developer playing strictly by the book will arrive at desirable outcomes. For instance, there should be formal instances to disambiguate requirements and flesh out acceptance criteria before programming work begins.

Most companies I’ve worked for have a development process, specified as a ticket lifecycle on Jira or a similar app, that looks more or less like this:

For a mode 1 developer, the job is _Done_ when a pull request is merged and the ticket closed. If changes are necessary or there are bugs to be fixed, they will make their way into the backlog in the form of new tickets.

But, especially when a bug is caught by a user or a stakeholder, someone may ask, perhaps during a retrospective or post-mortem meeting: _Why did we miss this during development? How could we prevent it in the future?_ I imagine there are as many answers to these questions as software teams in the world, but too often I see the conclusion boil down to: _We missed this in our unit testing; we should work on better coverage next time_. I’m [all about automated testing](https://olano.dev/blog/unit-testing-principles), but I find this framing unfortunate.

A particular team I worked with had a slightly refined workflow:

| Backlog | Selected for sprint | In progress | In review | **Awaiting deployment** | **Needs to be verified in production** | Done |
| --- | --- | --- | --- | --- | --- | --- |

While I’m firmly in the “individuals and interactions over processes and tools” camp, every now and then I find myself pining for the fjords those two extra columns, wishing we had them (and the culture they represent) on my current project. Those two columns are loaded; there are many benefits to this little workflow adjustment.

When a pull request gets merged, it’s now the ticket assignee’s responsibility to see the code deployed, which, if not fully automated, may mean taking over release duties. The ‘Awaiting deployment’ column then acts as an indicator of project health: teams with bad release hygiene would struggle to ignore an ever-growing pile of tickets; asymptotic MVPs and kafkaesque legacy modernizations would have no place to hide.

With this updated definition of _done_, the team velocity now suffers when work is not released, so a mode 1 project manager will also have an incentive to push for frequent deploys. This is in their best interest, not least because the taller the ‘Awaiting deployment’ column gets, the more likely your mode 2 devs are interviewing for their next position.

The usefulness of the ‘Awaiting deployment’ column should come as no surprise, as it is common industry knowledge that [deploys are the heartbeat of a company](https://charity.wtf/2019/05/01/friday-deploy-freezes-are-exactly-like-murdering-puppies/). The benefits of ‘Needs to be verified in production’ may be less obvious.

To close a ticket, developers now need to ask themselves: _what is a reasonable way to prove that the code does what I think it does?_ How exactly this verification should be carried out is not strictly defined, and is best left as an exercise to the developer. If a user interface is not available, it may suffice to query a database, check logs and metrics, or probe a production server. This routine procedure adds an incentive to improve the observability of the system.

Hopefully it’s not a big mental leap to go from _how do I prove this is working as expected_ to _does this meet the ticket requirements_ and then to _what problem was I trying to solve?_ That is: the process encourages mode 1 developers to switch—at least retroactively, at least temporarily—to mode 2.

While ad hoc tests and probes lack the good reputation and protection against regressions of automated tests, I can’t overstate how much value can be extracted from a few production checks, how many integration issues they catch, how much they contribute to a general sense of [confidence and ownership](https://olano.dev/blog/software-design-is-knowledge-building) over a software system[2](https://olano.dev/blog/verified-in-production/#footnote-2).

### Notes

[1](https://olano.dev/blog/verified-in-production/#footnote-reference-1)

This framing roughly maps to the [solutions implementer vs problem solver](https://rkoutnik.com/2016/04/21/implementers-solvers-and-finders.html) categorization.

[2](https://olano.dev/blog/verified-in-production/#footnote-reference-2)

This post was proofread by an LLM, I uploaded the full conversation [here](https://olano.dev/blog/verified-in-production-llm-proofread.txt).
