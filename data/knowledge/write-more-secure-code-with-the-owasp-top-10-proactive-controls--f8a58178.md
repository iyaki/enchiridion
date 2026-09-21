---
title: "Write more secure code with the OWASP Top 10 Proactive Controls"
notion_id: f8a58178-2e92-42fa-9450-4e9117e5a690
notion_url: https://app.notion.com/p/Write-more-secure-code-with-the-OWASP-Top-10-Proactive-Controls-f8a581782e9242fa94504e9117e5a690
last_edited: 2023-04-25T14:09:00.000Z
source_url: https://github.blog/2021-12-06-write-more-secure-code-owasp-top-10-proactive-controls/
tags: ["Article", "Github Blog", "English", "Programming", "Information Security", "DevOps"]
---
[https://github.blog/2021-12-06-write-more-secure-code-owasp-top-10-proactive-controls/](https://github.blog/2021-12-06-write-more-secure-code-owasp-top-10-proactive-controls/)

As software becomes the foundation of our digital—and sometimes even physical—lives, software security is increasingly important. But developers have a lot on their plates and asking them to become familiar with every single vulnerability category under the sun isn’t always feasible. Even for security practitioners, it’s overwhelming to keep up with every new vulnerability, attack vector, technique, and mitigation bypass. Developers are already wielding new languages and libraries at the speed of DevOps, agility, and CI/CD. Yet the reality remains that security is a shared responsibility.

“From my experience all software developers are now security engineers whether they know it, admit to it, or do it. Your code is now the security of the org you work for.”

– Jim Manico, OWASP Top 10 Proactive Controls co-leader

## The limits of “top 10” risk list

Projects such as the OWASP Top 10 Security Risks have always been a reference to drive developer security training, but these kinds of “top 10 risks” lists are not without some concerns:

First, security vulnerabilities continue to evolve and a top 10 list simply can’t offer a comprehensive understanding of all the problems that can affect your software. Entirely new vulnerability categories such as XS Leaks will probably never make it to these lists, but that doesn’t mean you shouldn’t care about them.

Secondly, these lists might not even apply to your technology stack. Do you need to learn about SQL injections when your application doesn’t use databases? It’s certainly worthwhile to try to understand the most common types of vulnerabilities and even to perform the mental exercise of trying to apply those concepts to the concrete technology stack you use in your daily jobs. (If untrusted input shouldn’t be used to craft a database query, it’s also probably not a good idea to use it to craft an LDAP one or a JNDI lookup, for example.)

Expecting developers to know every single vulnerability category and to be up to date with the latest attack vectors simply does not scale. So, what can you do to help prevent the introduction of these specific types of vulnerabilities in your code, even without deep knowledge or understanding of the vulnerabilities classes themselves? The good news: by consistently applying defensive programming concepts as developers, you reduce your odds of introducing vulnerabilities. At the very least, you reduce the odds of them being exploited in the event that a vulnerability does make its way into your code.

## The OWASP Top 10 Proactive Controls: a more practical list

The OWASP Top 10 Proactive Controls is a lesser-known OWASP project that is aimed at helping developers prevent vulnerabilities from being introduced in the first place by focusing on defensive techniques and controls, as opposed to any specific known risks or vulnerabilities. This approach is suitable for adoption by all developers, even those who are new to software security. It provides practical awareness about how to develop secure software.

“Software developers are the foundation of any application. But building secure software requires a security mindset. Unfortunately, obtaining such a mindset requires a lot of learning from a developer. The OWASP Top 10 Proactive Controls aim to lower this learning curve.”

– Jim Manico, OWASP Top 10 Proactive Controls co-leader

The Top 10 Proactive Controls, in order of importance, as stated in the 2018 edition are:

- C1: Define Security Requirements
- C2: Leverage Security Frameworks and Libraries
- C3: Secure Database Access
- C4: Encode and Escape Data
- C5: Validate All Inputs
- C6: Implement Digital Identity
- C7: Enforce Access Controls
- C8: Protect Data Everywhere
- C9: Implement Security Logging and Monitoring
- C10: Handle All Errors and Exceptions

Some of them may resonate with you and some of them may not. In this series, I’m going to introduce the OWASP Top 10 Proactive Controls one at a time to present concepts that will make your code more resilient and enable your code to defend itself against would-be attackers. When possible, I’ll also show you how to create CodeQL queries to help you ensure that you’re correctly applying these concepts and enforcing the application of these proactive controls throughout your code.

Let’s get started with a brief introduction of the concepts we’ll be exploring:

### C1: Define Security Requirements

Just as functional requirements are the basis of any project and something we need to do before writing the first line of code, security requirements are the foundation of any secure software. In the first blog post of this series, I’ll show you how to set the stage by clearly defining the security requirements and standards of your application. You’ll learn about the OWASP ASVS project, which contains hundreds of already classified security requirements that will help you identify and set the security requirements for your own project.

### C2: Leverage Security Frameworks and Libraries

If you devote your free time to developing and maintaining OSS projects, you might not have the time, resources, or security knowledge to implement security features in a robust, complete way. In this blog post, I’ll discuss the importance of establishing the different components and modules you’ll need in your project and how to choose frameworks and libraries with secure defaults. Two great examples of secure defaults in most web frameworks are web views that encode output by default (providing XSS attack defenses) as well as built-in protection against Cross-Site Request Forgeries. Sometimes though, secure defaults can be bypassed by developers on purpose. So, I’ll also show you how to use invariant enforcement to make sure that there are no unjustified deviations from such defaults across the full scope of your projects.

### C3: Secure Database Access

Database injections are probably one of the best-known security vulnerabilities, and many injection vulnerabilities are reported every year. In this blog post, I’ll cover the basics of query parameterization and how to avoid using string concatenation when creating your database queries.

### C4: Encode and Escape Data

No matter how many layers of validation data goes through, it should always be escaped/encoded for the right context. This concept is not only relevant for Cross-Site Scripting (XSS) vulnerabilities and the different HTML contexts, it also applies to any context where data and control planes are mixed. In this post, Senior Application Security Engineer Jason White will show you how to identify the characters with special meaning for any given context and how to properly encode them so they cannot be used to break out of the context they’re being written to.

### C5: Validate All Inputs

An easy way to secure applications would be to not accept inputs from users or other external sources. Obviously, that’s not a practical option. So, we need to validate the inputs to our applications. The phrase that possibly applies best here is “trust, but verify.” You can’t control or know what the inputs are that will come to your application, but you do know the general expectations of what those inputs should look like (a phone number, a zip code, or some other format). This is where validation comes into play. Checking and constraining those inputs against the expectations for those inputs will greatly reduce the potential for vulnerabilities in your application.

### C6: Implement Digital Identity

Authentication is used to verify that a user is who they claim to be. It’s important to carefully design how your users are going to prove their identity and how you’re going to handle user passwords and tokens. This should include processes and assumptions around resetting or restoring access for lost passwords, tokens, etc. In this post, you’ll learn how using standard and trusted libraries with secure defaults will greatly help you implement secure authentication.

### C7: Enforce Access Controls

Once authentication is taken care of, authorization should be applied to make sure that authenticated users have the permissions to perform any actions they need but nothing beyond those actions is allowed. In this post, you’ll learn more about the different types of access control and the main pitfalls to avoid.

### C8: Protect Data Everywhere

You need to protect data whether it is in transit (over the network) or at rest (in storage). Some of this has become easier over the years (namely using HTTPS and protecting data in transit). Performing cryptographic operations still often has sharp edges. You may even be tempted to come up with your own solution instead of handling those sharp edges. In this post, I’ll help you approach some of those sharp edges and libraries with a little more confidence.

### C9: Implement Security Logging and Monitoring

Incident logs are essential to forensic analysis and incident response investigations, but they’re also a useful way to identify bugs and potential abuse patterns. In this blog post, you’ll learn what should be logged and how.

### C10: Handle All Errors and Exceptions

Details of errors and exceptions are useful to us for debugging, analysis, and forensic investigations. They are generally not useful to a user unless that user is attacking your application. In this blog post, you’ll learn more about handling errors in a way that is useful to you and not to attackers. This includes making sure no sensitive data, such as passwords, access tokens, or any Personally Identifiable Information (PII) is leaked into error messages or logs.

Stay tuned for the next blog posts in this series to learn more about these proactive controls in depth. I’ll keep this post updated with links to each part of the series as they come out.

Follow GitHub Security Lab on Twitter for the latest in security research

## Written by

## Related posts

## Explore more from GitHub

### Docs

Everything you need to master GitHub, all in one place.

Go to Docs

### The ReadME Project

Stories and voices from the developer community.

Learn more

### GitHub Actions

Native CI/CD alongside code hosted in GitHub.

Learn more

### Enterprise content

Executive insights, curated just for you

Get started
