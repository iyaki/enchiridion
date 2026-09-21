---
title: "Three layers to secure a software development organization"
notion_id: cb9a0de0-1a34-4f22-90f5-22ab9022ea73
notion_url: https://app.notion.com/p/Three-layers-to-secure-a-software-development-organization-cb9a0de01a344f2290f522ab9022ea73
last_edited: 2023-02-17T19:52:00.000Z
source_url: https://stackoverflow.blog/2023/02/09/three-layers-to-secure-a-software-development-organization/
tags: ["Article", "Stack Overflow Blog", "English", "Information Security"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This affects the individual developer writing insecure code, the engineering team blindly trusting their dependencies, and the organization thinking that their best bet is to roll their own security controls.

Every year, universities, colleges, and bootcamps around the world graduate brand new software developers who have never taken a course on secure coding or application security. In fact, they may not have been taught anything about security at all.

This lack of security understanding has implications for software development on three layers: the individual developer writing insecure code, the engineering team blindly trusting their dependencies, and the organization thinking that their best bet is to roll their own security controls. In this article, I’ll talk about each of these layers, how a lack of security knowledge affects it, and how individuals and organizations can create software that follows security best practices.

## Hello (insecure) world

From the very first lesson, we are taught to code insecurely and incorrectly. The first thing we do in the famous [‘hello world’](https://stackoverflow.blog/2020/03/05/a-modern-hello-world-program-needs-more-than-just-code/) lesson is to put data on the screen: “Hello world!” The second thing is we ask the user for their name: “What is your name?” The user enters their name, we take that data and reflect it back to the screen to say “Hello <insert user’s text here>”. There is no mention of validating the user input to see if it’s potentially malicious, nor is the learner taught to output encode the information, which would disable any potential code the user has entered.

If we added more to the lesson, validating the input (and rejecting it if it was bad) and then output encoding the data before mirroring it to the screen, then we would be teaching every new coder how to avoid cross-site scripting (XSS). But we don’t teach that as part of the lesson. We show them how to do it _insecurely_. I suspect this is part of the reason that XSS is so prevalent across the internet: we’ve been drilling it into new coders how to create this vulnerability from the very first day.

### Secure coding best practices

If everyone is learning the wrong way from the start, how can we fix this? There are several things, but let’s start with how we build software: the system development life cycle (SDLC).

Some common security activities that can be added to the SDLC include:

- **Security requirements**: The measures that need to be put in place for the application to be considered secure. You’re building a [serverless](https://stackoverflow.blog/2020/05/18/you-want-efficient-application-scaling-go-serverless/) app? Cool! Here’s a list of security requirements we need you to add to your project’s requirements document to ensure it’s safe enough to put it on the internet.
- **Threat modeling**: A brainstorming session that includes at least one security professional, a product owner and/or business representative, and a member of the technical team. The goal is to identify as many potential threats to the product as possible, then work to mitigate the ones that seem the most damaging or dangerous. This is often done in the design phase, but it could be done at any point afterwards if you didn’t have time earlier in the SDLC.
- **Code scans**: During the coding phase, you could scan your own code with a static application security testing tool (SAST) or a software composition analysis tool (SCA) or throw your app on a dev server and scan it with a [DAST](https://stackoverflow.blog/2022/11/30/continuous-delivery-meet-continuous-security/) tool.
- **Break it early**: During the testing phase, you could hire a penetration tester to perform thorough automated and manual testing of your entire system.
- **Protect code in production**: During the release phase, you could set up [monitoring, logging, and alerting](https://stackoverflow.blog/2022/10/12/how-observability-driven-development-creates-elite-performers/) for your application. You could install it behind a runtime application security protection (RASP) or web application firewall (WAF).

All the items listed above are just the beginning: you can organize your secure SDLC in any way that works for your organization and teams. You can use some of these or none of these; the key is finding activities and tooling that work for _your_ organization.

We can create a _secure_ SDLC by adding security activities throughout out processes. That said, we have to follow the security steps every time, not just sometimes or when it’s convenient if we want to reliably produce secure software.

If you add one single security step or verification to your SDLC, you will create more secure applications. If you add more than one step, you will create _even better_ applications. [Software is always a tradeoff](https://stackoverflow.blog/2022/01/17/plan-for-tradeoffs-you-cant-optimize-all-software-quality-attributes/): we don’t want to spend a million dollars on security for an application that will only earn h a couple thousand dollars worth of value. We _do _want to ensure we are protecting our systems such that they meet the risk tolerance of our organization. In more plain language, we should work with the security team to decide exactly how much security is required, but generally it’s safe to assume that ‘more is more’ when it comes to security.

## Zero trust, but not the marketing version

Quite often when designing systems, we create ‘implied trust’; that is, one or more parts of the system don’t verify something they should or could. When someone logs into a system, a security control verifies the user’s identity (via the username and password combination) and then grants them access to the parts they are allowed to access. In programming, we generally call the first part authentication (who are you?) and the second part authorization (should you be here?). If we skip this step, we are creating _implied trust_. We don’t care who the person is and are assuming that the person is allowed to be there.

Zero trust means never, ever, having any implied trust. Every single possible part of a system is locked down and only opened if it must be. This means closing all ports except the ones you need. It means blocking all connections except the ones you know you need. It means always verifying everything before using it. No trust in anything at all, even the other systems you have as dependencies.

Although zero trust is quite a lot of work to implement, it _works_. And it works _well_.

Examples of zero trust that you could implement in programming:

- Validating, sanitizing, and escaping all inputs (in that order)
- Each API, serverless, container, and app is an island; treat them like that. They should not trust each other!
- Authentication and authorization for everyone! Every single integration, even with systems built and maintained by your team, requires authentication and authorization, for every single connection.
- Output encoding, because you never know if someone changed something. Do not trust your user’s input or input that may have been tampered with!

I have never seen an organization implement every single possible variation of zero trust, but that’s okay. Applying zero trust principles to as many systems as possible is good enough.

## Buy, borrow, _then_ build

Building security controls is hard. They are complex in nature to build, they are tested far more often and more aggressively than any other feature (thanks to malicious actors), and there are few public examples to work from due to companies wanting to protect their intellectual property. This puts us in a situation where building our own custom security control ends up being expensive, time consuming, and potentially even dangerous.

With this in mind, security professionals (and software architects, for that matter) often recommend we go in this order when we decide if we will use a pre-existing component or write our own:

- **Buy it **if we can afford it with the budget we have. The reason for this is that it’s available immediately, we don’t need to assign resources to create and thus they are free for other work, we don’t need to maintain it ourselves (which is almost always the reason custom software costs more than buying it), and it is very likely to have had significantly more security testing performed against it (big customers demand intensive testing and high security assurance). Although the price tag may seem steep, when you calculate the risk, the wait, and the long-term cost of maintaining the system, it is almost always cheaper to buy than build.
- **Borrow it.** This can mean open source software, online examples, code from other teams, using the features in your framework, or even having an entire system supplied (this happens in government and nonprofit organizations more often than for-profit/private). Borrowing it is more of an expression than a reality, we mean ‘use code that is open source or other code you can use for free, but that you did not write’. The advantages of this are similar to the ‘buy’ section: the borrowed code is available immediately, it’s free, and we don’t have to maintain it ourselves. That said, if it’s given for free, it might not have had any security testing on it, meaning you should verify that it’s secure before using it if at all possible.
- If neither of those is an option, we **build it**.

This saves our organizations money, time and risk, despite the fact that building our own from scratch is usually way more fun.

With this in mind, whenever possible, use the security features found within your programming framework, such as: encryption, authentication, and session management. They are tried, tested, and true!

Next, use third-party components (libraries, packages, gems, etc.), as they have (usually) been tested quite thoroughly. Verify third-party code and components via a software composition analysis (SCA) tool and reverify often.

To summarize the buy/borrow/build philosophy: don’t write your own security features unless you absolutely have to. It’s hard to get right!

In closing, although we may not have been taught it in school, we can start with the steps above to begin creating more secure applications and code. By creating a secure system development lifecycle, avoiding implied trust within the systems we build, and only creating our own security controls when it is absolutely necessary, we will reduce the risk to our organizations in a consistently positive way. The organizations we work for are counting on us to be their first line of defense, and with these new strategies, you should be up to the task
