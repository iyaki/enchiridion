---
title: "Your incident response wasn’t built for AI - LeadDev"
notion_id: 3df54f1c-7d23-8180-9b3e-cb40bb31ee33
notion_url: https://app.notion.com/p/Your-incident-response-wasn-t-built-for-AI-LeadDev-3df54f1c7d2381809b3ecb40bb31ee33
last_edited: 2026-09-18T01:18:00.000Z
source_url: https://leaddev.com/ai/your-incident-response-wasnt-built-for-ai
tags: ["DevOps", "Information Security", "AI", "Incident Response", "Monitoring", "Software Architecture", "Article", "LeadDev", "English"]
---
You have **1** article left to read this month before you need to [register](https://leaddev.com/register) a free LeadDev.com account.

Estimated reading time: 6 minutes

**Key takeaways:**

- AI systems fail differently: fluent, confident outputs return a 200 OK even when wrong, and identical inputs can produce different outputs due to GPU batching.
- **Nobody owns the quality floor**. Product picks the model, platform carries the pager, and neither controls the failure rate.
- Split SLOs into three tiers:** service**, **behavioral (capped by the model)**, and **containment**, the only tier a team can actually own.

---

[Incident response](https://leaddev.com/technical-direction/incident-response-before-the-incident) assumes three things: that a failure can be reproduced, that it announces itself as an error, and that the person paged has a lever that changes the outcome. Put a [language model in production](https://leaddev.com/leadership/llms-an-operators-view) and all three stop being true at the same time.

This is not a monitoring gap you close by adding dashboards. Your alerting was built to catch a system saying no. These systems fail by saying yes, fluently, with the wrong content, and returning a 200 OK (success, everything worked as expected) while they do it. Availability tooling watching a correctness risk stays green through the entire incident.

## Your incident loop just lost step two

Detect, reproduce, isolate, fix. Step two is where the work usually happens, and it is not available to you.

The [standard objection](https://leaddev.com/software-quality/your-sdlc-is-your-context-engineering) is that it should be. These models have a setting that controls how much randomness goes into choosing each next word, and turning it all the way down makes the model always take its highest scoring option. Teams set it there and assume they have a reproducible system.

They do not. The request never runs alone. Inference servers batch requests arriving at the same moment onto the same GPU, and who lands beside you depends on traffic you do not control. That changes the order the hardware adds numbers in, floating point addition is not perfectly associative, and the scores come back different in the fourth decimal. Almost always invisible. When the top two candidates sit close together, a difference that small flips which one wins, and one different word conditions everything after it.

Same input, different output, nothing changed on your side.

So search your [incident tracker ](https://leaddev.com/technical-direction/incident-response-before-the-incident)for tickets closed as “could not reproduce.” In a deterministic system that is a reasonable disposition. Here it is a category of live defect you have been filing away for months.

## More like this

## The failure returns a 200

There is no error path for being wrong. The model has no way to signal it, and its confidence is not calibrated to its correctness, so you cannot use its certainty as a proxy either.

In a retrieval pipeline the failure is usually upstream and silent. Retrieval returns weakly relevant documents, the generator synthesizes something fluent out of them anyway, and every span in the trace reports success. In [agentic systems](https://leaddev.com/technical-direction/are-you-ready-for-agentic-observability) it compounds, because a wrong step becomes the input to the next one. Nothing in that chain throws.

It also arrives segment shaped rather than uniform. A bad deploy hits everyone and shows up in aggregate within minutes. A model regression lands as 3% of requests concentrated in one customer segment, moving your aggregate quality metric by less than its own noise floor. No dashboard in your organization displays 3%, which is why the escalation reaches you through the account team rather than the pager.

## The fault line nobody drew on the org chart

Product shipped the [AI](https://leaddev.com/ai/best-ai-coding-assistants) feature. Platform inherited the pager. Neither team owns quality.

The people who chose the dependency do not carry the page. The people carrying the page cannot change the model. Ask three leaders in most companies who owns the quality floor for the AI surface and you will get three answers, none of them written down anywhere.

Every control below is really an answer to that question.

![image](https://res.cloudinary.com/leaddev/image/upload/f_auto/q_auto/dpr_auto/c_limit,h_1024,w_1024/next/2026/08/image.png)

_Source: Adora Nwodo_

## What to do instead

Two of those steps, scope and contain, have no equivalent on the left. They are also the two nobody staffs.

The instinct here is to try harder at the model. Better prompts, more [evals](https://leaddev.com/software-quality/from-evals-to-experiments-how-to-ship-successful-ai-initiatives-by-failing-cheaply), a fine tune. I think that is a trap, because you are pushing on a number that belongs to somebody else’s training run. The floor was set in a lab you do not have access to. What you control is what each failure costs.

**Split the SLO into three tiers.**

- Service tier stays exactly as it is. Availability and latency, still necessary, still not sufficient.
- Behavioral tier measures task success against a versioned golden set, sampled continuously in production, reported per segment rather than in aggregate. If you use a model to grade those samples, validate it against human labels first. An unvalidated grader is a metric with unknown bias, which is worse than no metric because people trust it.
- Containment tier is the one you actually own. What fraction of bad outputs are caught before a user sees them, and how many minutes from detection to disablement.

Only the third sits on your timeline. The second is capped by a training run you did not commission, so be straight with your executives about which of these can carry a hard target.

### Make the escape hatch a launch gate

If [on-call](https://leaddev.com/technical-direction/crafting-efficient-call-processes) cannot disable the capability alone, at three in the morning, without a deploy and without waking a second person, it is not production ready. Flag at capability granularity, not service granularity. Define the degraded path before launch, then exercise it in a game day, because an untested kill switch is a belief rather than a control. Put this in the readiness review template. If it lives there it is structural. If it lives in your head it is a negotiation you have to win again every quarter, usually two days before a launch date.

### Replace reproduction with capture

One trace ID carrying the resolved prompt, template version, model and provider version, retrieval document IDs with their relevance scores, the tool call sequence, and the raw output. Those relevance scores are the part teams skip and what makes retrieval incidents diagnosable, because they separate “retrieval missed” from “the generator ignored the context.” Different fixes, different owners.

### Rewrite the postmortem

Root cause is often unrecoverable, and pretending otherwise produces a review that ends with “we improved the prompt.” That is a wish with a commit hash. Ban it unless it ships with an eval case that fails before and passes after. Grade the response instead: time to detection, time to containment, whether the escape hatch was used, whether it worked. It is the only version of that meeting where your engineers do not leave feeling blamed for physics.

One thing to stop while you are at it. Human in the loop is not a control until you measure the reviewer’s catch rate. If your reviewers approve above the mid 90s they are confirming rather than evaluating, and you are counting a rubber stamp as a safety layer. That claim is probably sitting in a compliance document already.

![image](https://res.cloudinary.com/leaddev/image/upload/f_auto/q_auto/dpr_auto/next/2026/03/LDX3-New-York-is-live.png)

**New York • September 15 & 16, 2026**

The pace of change keeps accelerating. See what other leaders are doing about it, at LDX3 New York.

## The actual job

Deterministic systems let you drive the failure rate towards zero and build an entire [operational culture](https://leaddev.com/culture/how-build-intentional-culture) on that trajectory. Probabilistic systems have a floor you did not set and cannot see.

You are being asked to hold availability shaped accountability for a correctness shaped risk, on a dependency you do not version, with a pager carried by the team that did not choose it. The job is not to fix the model. It is to build the containment layer, and to be straight with your executives about which numbers you can move and which ones belong to somebody else.
