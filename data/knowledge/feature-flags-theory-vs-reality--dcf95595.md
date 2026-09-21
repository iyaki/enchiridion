---
title: "Feature Flags: Theory vs Reality"
notion_id: dcf95595-9598-4255-b143-e8e8e6efc8f5
notion_url: https://app.notion.com/p/Feature-Flags-Theory-vs-Reality-dcf9559595984255b143e8e8e6efc8f5
last_edited: 2023-07-12T19:53:00.000Z
source_url: https://bpapillon.com/post/feature-flags-theory-vs-reality/
tags: ["bpapillon", "English", "DevOps", "Continuous Integration/Continuous Delivery", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

A discussion about DevOps in 2023

## **The Rise of Feature Flags in DevOps**

During the second half of the 2010s, the DevOps movement gained massive momentum throughout the SaaS industry. Riding this wave and buoyed by the marketing efforts of LaunchDarkly, feature toggles rapidly became an essential tool for engineering operations used by practically every SaaS company.

In the DevOps context, feature flags are especially great for continuous deployment. By decoupling release from deployment, they allow marketing and engineering to operate independently. This flexibility enables engineers to continuously deploy features in progress.

However, a few downsides were clear even from the beginning. To name a few:

- **They add complexity to the code**, including additional testing overhead.
- **It’s difficult to verify** whether all of your changes are actually “behind” a feature flag as intended.
- **Feature toggle checks aren’t free**; it’s more logic, which means more possible bugs, and the check may take time if a database or external service is involved.

Due to these shortcomings, the [conventional wisdom](https://martinfowler.com/articles/feature-toggles.html) has always been that feature flags should be short-lived and kept to a minimum:

> 

Savvy teams view the Feature Toggles in their codebase as inventory which comes with a carrying cost and seek to keep that inventory as low as possible.

## **Reality Check!**

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

A software developer interacts with their feature management system

As adoption of the feature toggle pattern has spread throughout the industry, a few common problems can be observed.

### **1. Release follow-through and ever-expanding complexity**

If feature flags are meant to be short-lived, then there must be a process for reviewing them and cleaning them up at the appropriate time. This often takes the form of a manual process, such as filing a ticket for a future sprint.

As a general rule of thumb, any process that relies on humans to remember things is not going to be a reliable one, and for most companies, these processes end up being spotty and inconsistent.

A few common issues I’ve seen:

- **Zombie flags**
- In our haste to move on to the next big feature, we often forget to rip out the release flag. This leads to higher “carrying cost”, taking many forms: the readability of the code, the number of automated tests that continually execute each branch of the feature flag logic, and so forth. Over time, this makes our code base harder to work on and leads to annoyance, wasted time, and bugs.
- **Unfinished business**
- Perhaps we forget to finish rolling out our flag, or maybe some customer segments never even get the new feature. This leads to unexpected support tickets months later, at a time when everyone in the organization thought that the rollout was long finished.
- **Ghost flags**
- Sometimes, even when we remember to rip out the feature flag, there may be a communication breakdown with other users, who end up futilely toggling a control that has no effect.

### **2. Unintended or inappropriate usage**

Even with perfect follow-through, a flag meant for release purposes might end up being implemented for long-term feature access. This may seem like a time-saving win, but it results in third-party services becoming load-bearing for cases beyond their intended purpose.

Imagine even a very short outage by the service provider. If we only rely on that service provider for release flags, we probably live with our hardcoded fallback values; but, if we’re using it for something like a permission or entitlement check, there is now essentially an outage of our own.

Many teams understand this danger, and respond by implementing separate systems for short-lived and long-lived flags; the former may use a managed service, while the latter may be a home-grown system that stores its data in the main application database.

Even though we now have a more correct tool for each job, we still end up with a lot of problems in practice:

- **Lack of portability & promotion**
- We’ve mentioned that release flags can sometimes take on a second life as other types of toggles. With this model, if a release flag starts to make sense as a long-lived flag later, there’s no way to promote it to the more appropriate system without making code changes.
- **User confusion & frustration**
- Across these two systems (short and long-lived), there may be a broad base of user types: engineering, product, customer success, marketing, ops. It’s unlikely that these groups share a consistent mental model for the systems, and people tend to just refer to all of it as “flags”. Many users may be unclear as to why there are two separate systems at all, and feel frustrated that they just have to remember which flag lives in which system.
- **Lack of context & readability**
- Exacerbating the user confusion issue, most flags are exposed as a list of keys (e.g. “new_onboarding_flow”, “widgets_v2”) that provide very little context about what these flags do or how they’re meant to be used.
- Without talking to someone, reading code, or separately maintaining documentation, it’s very difficult to know much about the flags, which frustrates non-technical and technical stakholders alike.
- **Testing headaches**
- Developers have to accommodate multiple systems in testing; for example, if you’re using mocks or stubs in your tests to simulate the behavior of these systems, you now need twice as many of these.
- **Feature gaps**
- A long-lived flag service that is homegrown may be able to provide better guarantees with regards to latency and availability; however, the short-lived flag service will almost always be more feature-rich because it’s often provided by a managed service. This reality may incentivize its use over the more reliable system in certain cases.

## Accepting Reality

After observing the feature flag experiment in the wild for some time, I’ve come to the conclusion that the conventional wisdom to limit the number and age of feature flags in your code base is wise, but unrealistic. The fast-paced and cross-functional nature of modern software development create dynamics that are too hard to overcome with best intentions and best practices.

Furthermore, there’s a disconnect between “feature management” as a term of industry and the actual “feature management” that goes on. This term, and the tools that are sold within this market, generally refer only to DevOps use cases like rollout and experimentation. However, we clearly do a lot more managing of features than this.

- Every time a new customer signs up, are their features not being managed?
- When they upgrade to a more expensive plan?
- When a sales negotiation results in a bespoke enterprise plan?
- When a customer success rep enables an add-on?

These are all feature management, but none of them fall into the definition of “feature management” that our tooling and DevOps culture wants to support.

As engineers, it’s time to change our framing of feature management to better align with the businesses we operate in, but to do this, we need new tools.

## Taming complexity

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

A software developer evaluates their adherence to feature management best practices

A software developer evaluates their adherence to feature management best practices

If we accept that we as engineers are powerless to contain the spread of feature management, and perhaps that we are holding back our businesses to the extent we try to do so, then we need to stop relying on manual processes for hygiene and maintenance.

Let’s imagine what capabilities we might need a new feature management tool to have in order to accomplish this. A few possibilities:

- Long-lived and short-lived use cases coexist within the same tool, but are clearly delineated in its interfaces.
- Short-lived flags might come with additional metadata, such as an expiration date by which we expect the flag should no longer be in use.
- Flags should have an owner, either an individual user or perhaps a user role or group.
- Users should be able to add meaning to flags after the fact via metadata and grouping.
- If a flag changes purpose, say from a release flag to an entitlement check, we can simply update this in the tool. Such a change would be tracked in an audit log.
- Policies can be set; for example, short-lived flags must be removed or graduated within a specified amount of time, or perhaps certain metadata (like expiration date) can be made required for certain types of flags.
- Flags can easily be used in relation to one another; for example, one flag might be required in order for another flag to be enabled, or two flags might be incompatible with one another for code reasons. The tool should make it easy to configure such invariants.

With capabilities like this in place, the tool could start to automate some of the maintenance processes that are currently manual.

For example:

- Auditing a codebase for out-of-date flags could be done via a static analysis tool in CI.
- We could fail builds or notify engineering or product managers if certain assumptions are not met.
- Flag owners could receive notifications when flags they are responsible for are out of compliance, or a ticket could automatically be filed in the ticketing system.

If we can automate these processes, then we finally might have a system that holds up to the chaos of the modern software development process and fights back against ever-growing complexity.

## Onward!

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Leaving the old expectations behind as we move to a new framing of feature management

It’s high time that we take another look at what “feature management” means in our industry. If we accept a more expansive view that aligns with how our businesses want to be managing features and build the tools needed to support this, we can free ourselves from the need to adhere to best practices that have proved unrealistic.

If we were to have a tool like this that better suited the natural complexity of feature management, then this would be a great start. However, there are more considerations, such as the architecture of such a tool, that I will explore in upcoming posts.
