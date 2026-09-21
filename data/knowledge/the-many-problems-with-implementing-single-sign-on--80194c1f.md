---
title: "The many problems with implementing Single Sign-On"
notion_id: 80194c1f-44b7-4a43-b34b-a8ac393d91da
notion_url: https://app.notion.com/p/The-many-problems-with-implementing-Single-Sign-On-80194c1f44b74a43b34ba8ac393d91da
last_edited: 2023-01-25T19:25:00.000Z
source_url: https://stackoverflow.blog/2022/09/12/the-many-problems-with-implementing-single-sign-on/
tags: ["Information Security", "Programming", "Article", "Stack Overflow Blog", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Without SSO and other enterprise features, a product can only go so far.

Most SaaS startups want to sell to enterprises, but many are unprepared for an enterprise’s most-requested requirement: single sign-on (SSO). SaaS products are often designed for usernames and passwords, not complex integrations with identity providers (IdPs) like Okta, Google, or Active Directory.

When faced with the challenge of building those integrations, many developers will roll up their sleeves, read authentication and authorization specs, and get to work. Unfortunately, there can be so much room for interpretation in these specifications that SSO implementation becomes difficult, slow, and risky. If you are a small startup, or even a mid-sized business with a busy engineering team, this work can be a major drain and slow down your ability to begin acquiring enterprise clients.

It can cost precious weeks (or months) to implement SSO for an enterprise deal and, to many developers’ surprise, what worked for one enterprise customer may not work for the next one. SSO is unexpectedly challenging to do correctly and consistently. Let’s break down why, then offer some alternative approaches.

# What’s so hard about building SSO?

Let me tell you about the first time I tried to implement SSO.

I started Nylas Mail in 2013 after writing the initial lines of code in my dorm room at MIT. [It became a very successful open-source project](https://github.com/nylas/nylas-mail), and we raised more than $10M in an attempt to dethrone Microsoft Outlook. It wasn’t long before we needed to commercialize, which is when we faced a whole new audience of buyers: IT leaders and procurement professionals. We were excited to be speaking to enterprise clients about adoption, but they needed enterprise features to roll out Nylas at scale. As a small startup, we had designed our app for everyday users and not for enterprise adoption, so we weren’t prepared to integrate with an IdP or satisfy other enterprise requirements. The work to add those features proved too much for us, and kept us from commercial market success.

The lesson here is that without SSO and other enterprise features, a product can only go so far.

The key thing to understand about SSO is that it’s an integration problem. Most developers build apps by using OSS packages, such as [Devise](https://github.com/heartcombo/devise), which handles authentication for Ruby on Rails. This works for customers with basic requirements, allowing people to use different models with high modularity. At some point developers end up needing to integrate their product with an IdP, usually beginning with whatever their biggest customer needs. Different enterprises use different IdP solutions: some may use off-the-shelf solutions like Okta or Azure Active Directory, while others have their own custom homegrown solutions.

Any vendor wishing to sell to these companies must integrate with all of these IdPs, which means managing both IdP authentication and native credential-based users (that is, a username and password native to the vendor’s platform). Since none of the IdPs functions identically and there are so many SSO providers, you’ll need to maintain multiple integrations in parallel.

If you want to read about another journey in adding SSO to an enterprise product, [here’s how Stack Overflow did it](https://stackoverflow.blog/2019/07/11/single-sign-on-sso-stack-overflow-okta-integration/).

SSO integration goes beyond building features and functionalities and ensuring they perform. It takes considerable work to seamlessly adapt an app’s login flow. Orchestration and feature enablement are dependent upon IdP functionality and require new business logic with implications for mobile and two-factor authentication (2FA). SSO usually has 2FA built in too, creating additional complexity that must be considered alongside other system-level integrations.

# SAML is half the battle

There are multiple ways to implement SSO. One of the most popular is through an XML-based open standard called Security Assertion Markup Language (SAML). The [SAML specification](http://saml.xml.org/saml-specifications) is flexible and has a number of options to cover a range of possible cases. No two vendors implement the spec the same way. In other words, it’s not implemented consistently and there are several “flavors” of SAML. As a result, it’s rife with opportunities for security issues.

To be clear, it’s not that the spec is designed poorly. Rather, when the protocols were designed, the designers wanted to cover a lot of possibilities. A lot can be done with SAML that’s rarely implemented. But any SSO integration still needs to be ready to support certain edge cases, which is where the vulnerabilities surface.

## SAML payloads

When preparing for an SSO integration, an enterprise vendor will likely conduct a security review of SAML code to look for exploitable code or flows. For example, SAML uses certificates and signatures for payloads, which can be made particularly complicated by nested data structures. This is handy for management of multiple levels of IdP communication, particularly at large enterprises where there may be multiple layers of SAML authentication to pass through. Each layer has signatures that need verification, a process that can be like peeling an onion. A generalized SAML integration can be difficult to implement and check because it’s not always hierarchical and requests between systems can be non-linear.

If every layer isn’t thoroughly checked, malicious actors can misuse the payload. A common SAML exploit is to modify valid responses and inject a different invalid signature from an expired session. The reason this works is simple: one of the SAML SSO developers wrote code that checks for valid signatures, but not through the entire response. Every layer of the SAML payload needs to be investigated through and through. Most organizations don’t have multi-level ID systems, so this isn’t a particularly common exploit, but the protocol supports it because it was designed with this type of flexibility in mind.

## Lack of standardization can easily lead to failures and vulnerabilities

When it comes to SSO, there are thousands of things to get right and a lot of small details to account for. You can implement something and get it working, but robustness will come from testing it against implementations. And while the specification is standardized, those implementations aren’t.

Consider writing a SAML integration for a new organizational platform. The platform may use [canonical XML](https://stackoverflow.com/questions/28145726/proper-xml-canonicalization-for-saml), but may not explicitly declare a namespace, which can cause an authentication failure. Even if both parties on each side of the flow conform to the spec, there’s no guarantee that it will work out of the box. It’s especially painful when one party makes even a small change to their implementation because it can lead to login failures that are especially difficult to troubleshoot. A user might call the IT department and say, “I can’t log in,” launching a wild goose chase in an effort to troubleshoot an issue with a small XML change on a system they don’t control.

SAML’s XML-based nature comes with XML-related challenges. Object ordering and arrangement of nested entities (that is, tags) can cause problems. Attribute mapping is non-standard across platforms. The identifiers for users on platforms are not standardized. Sometimes you get an ID, sometimes an email. Sometimes there’s no identifier and it’s something opaque like a serialized Active Directory string. It might be tempting to use an open-source library—nothing beats the low price of free—but not many open-source packages handle XML well.

## App design can influence SAML functionality and security

SSO integration can also differ based on the structure of an application. For example, I can log into GitHub with my personal Gmail account and then jump into one of my company’s internal systems. GitHub is essentially acting as a second authentication factor for my internal app, allowing me to skip over the primary authentication mechanism for the app. You can have primary and secondary ways of authenticating—and different flows for email and password authentication–but whether those systems work as designed is influenced by the nature of the application itself.

When you’re building an app with SAML SSO, you don’t have to worry about changing usernames or passwords because that can be done outside of your application. However, the flows have to reflect that reality. Some IdP systems, such as Microsoft’s Active Directory, which provides an opaque, unique string, don’t provide an email address. If you don’t get an email address, then you have to figure out another way of identifying and authenticating a user. That method may not work with your application’s data structure or overall architecture.

Many enterprise-ready SaaS apps start SSO logins by asking for the organization’s subdomain. A behind-the-scenes lookup directs you to the login form for that company’s IdP login on that SaaS application. The problem is that a malicious actor might be able to find an exploit by typing a variety of company subdomains and know who’s using SSO and SAML. If that malicious actor is looking to use a newly-discovered security vulnerability, it’s possible that exploit will work on at least one of these IdP login screens.

Building SSO-capable user interfaces is difficult and every company is doing something different. The most common pattern is to have username and password fields with some IdP icons below and in small text, “Sign in with SAML/SSO.” Apple’s iCloud login doesn’t show a password field right away because they check whether the username’s domain is covered by an SSO flow. Slack uses custom CNAMEs in their URLs, like StackOverflow.slack.com and sends you to the right login flow. There’s no standard way of doing it, which is why everyone does it differently.

# Onboarding enterprises and offboarding people

Another complication related to SAML-based SSO is onboarding new customers. Let’s say you’re a SaaS vendor and you’ve just signed your first big contract with a major enterprise. Now it’s time to get those enterprise users into your systems. Your integration engineer will get the SAML integration set up by working with the enterprise’s security architect.

Setting up a SAML integration involves exchanging a set of data parameters, like redirect URLs, and field mapping of SAML attributes. There’s no clear spec for how names should be (such as case sensitivity), so testing will need to be done to ensure both sides work. There’s also the need to upload a certificate, which also has no clear methodology in the spec.

Testing the SAML integration is a huge friction point for companies and it’s almost always done manually, which means the tools for the setup are managed manually as well. At its most basic, companies have used a simple spreadsheet or form for the data parameter exchange. The process is so prone to breakage, it’s often handled with white-glove personalized service. Building a multi-use admin panel for SAML is expensive and difficult, so companies usually proceed with manual work until the project becomes cost-prohibitive.

Coordinating, configuring, implementing, and testing SSO can take weeks if not months of back-and-forth communication. It’s an arduous process, often involving lots of discussion and work around authentication (“here’s the validated identity of a user”) and authorization (“here are the services and features to which the validated user has access”). Enterprise ID systems like SAML only do authentication to prove someone is who they say they are, leaving authorization up to apps and services. SAML should be brokering actual sessions like a one-bit authorization, proving a person is who they say they are, almost like scanning an RFID badge at a security door. What happens after they pass through the door is another thing altogether.

## Session management challenges

When you authenticate with SAML, you’re authenticating a user, but after that validation, you can’t check that it’s still a current active session. You know you had a current session at a point in time when you authenticated, but how long does your session last? If it’s too short—say, 24 hours—it can be a pain, requiring a login every day for every user. There are ways to get around this if you have access to other identity systems, but the point is: session management is challenging.

For B2C SaaS products, logging out users has a negative effect on retention and engagement. Most products have a two-week cookie to keep the session fresh, but enterprise IT admins don’t always want that. SaaS vendors, especially those just selling to their first enterprise customer, may have to rethink session management for their apps. Then there are considerations for how sessions are managed on mobile devices.

## Offboarding users

When an enterprise signs a contract with your product, they bring a whole slew of new users. But not all of those users will stay for the life of the contract. A variety of processes need to be designed for enterprise employee departures. For example:

- What happens to a user’s data when an employee leaves?
- Does the user’s data or account get transitioned to an administrative account?
- Should all sessions be revoked? If so, how soon and how will you find out they left?
- How will shared items be handled?
- What needs to be deactivated and what needs to be preserved?

These are several important questions to consider when selling to enterprise customers. They’re not easy questions to answer, especially if user deactivation is already implemented in an inflexible way.

When an employee departs a client organization, perhaps on bad terms, there’s a chance they have access to sensitive information, important deliverables in progress, or shared documents with impactful financial implications. Without a solid plan for offboarding those users, SaaS vendors can find themselves on the critical path for their enterprise clients’ HR, IT, and information security offboarding processes. If the employee chooses to sabotage or tamper with data to which they should no longer have access, that’s bad news for the vendor.

# Crossing the enterprise chasm

When selling to an enterprise, your product needs to offer enterprise features. Implementing SSO—the most requested enterprise feature—may require changes from application and data architecture to how a dev team operates and interacts with large clients. There’s almost nothing within a SaaS app that doesn’t require reconsideration (or at least a double check) for supporting that first enterprise client.

SSO isn’t just about building an enterprise-level feature, but about maintaining a set of flows that enable different clients with different IdPs to use an app. Building SSO with SAML requires a lot of decisions and business logic, which doesn’t happen organically. Additionally, it’s best practice to ensure there’s more than one person on staff who understands the SAML implementation. Otherwise, if the dev leaves and no one knows the protocols, it can get messy fast. And enterprises don’t want to mess around with SAML.

By far, SSO is the most important enterprise feature to implement, so if there’s only enough budget for one feature, it should be SSO. It’s the gateway to selling to a ton of companies. Going deeper, vendors should consider whether any of the following would be of additional benefit to their SAML SSO implementation:

- Directory integration
- Automatic user lifecycle management
- Deprovisioning users
- Auto-login for audit compliance and e-retention
- SOC 2 compliance
- GDPR
- HIPAA
- Granular role-based access control (RBAC)
- Enterprise key management (bring your own key and encryption)

Before jumping head first into building SSO, take a look at [WorkOS](https://www.workos.com/?utm_source=stackoverflow&utm_medium=blog&utm_campaign=sept2022). SSO is just [one of the features](https://workos.com/single-sign-on?utm_source=stackoverflow&utm_medium=blog&utm_campaign=sept2022) we offer with just a few lines of code in your app. We maintain tight integration with the most popular IdPs, have an easy-to-use administration panel that enterprises love, offer live synchronization with enterprise user directories, and make multi-factor authentication a breeze.

_The Stack Overflow blog is committed to publishing interesting articles by developers, for developers. From time to time that means working with companies that are also clients of Stack Overflow’s through our advertising, talent, or teams business. When we publish work from clients, we’ll identify it as Partner Content with tags and by including this disclaimer at the bottom._

Tags: [partnercontent](https://stackoverflow.blog/tag/partnercontent/), [single sign-on](https://stackoverflow.blog/tag/single-sign-on/)
