---
title: "Security needs to shift left into the software development lifecycle"
notion_id: f26616a6-aa5d-4b17-9402-dd8e90b3451b
notion_url: https://app.notion.com/p/Security-needs-to-shift-left-into-the-software-development-lifecycle-f26616a6aa5d4b179402dd8e90b3451b
last_edited: 2026-09-21T17:24:00.000Z
source_url: https://stackoverflow.blog/2022/05/05/security-needs-to-shift-left-into-the-software-development-lifecycle/
tags: ["Stack Overflow Blog", "English", "DevOps", "Information Security", "Article"]
---
[https://stackoverflow.blog/2022/05/05/security-needs-to-shift-left-into-the-software-development-lifecycle/](https://stackoverflow.blog/2022/05/05/security-needs-to-shift-left-into-the-software-development-lifecycle/)

Before the cloud-native era, software products were often developed through sequential work by several teams across multiple disciplines. One team did their work, then threw it over the wall to the next team. the software development life cycle (SDLC) was a linear process:

- The software development team iteratively built the software to run on that infrastructure
- QA and product management coordinated validation efforts on functional correctness and user acceptance
- The information security team tested for security and compliance concerns relative to the requirements of the particular customer, vertical, or industry
- The IT team defined, purchased, installed, and configured the necessary infrastructure
- Finally, the operations team deployed the software to production

This provisioning and release cycle took weeks, if not months. Additional concerns like security and privacy were handled outside the SDLC. It wasn’t a huge deal—all your code and dependencies ran on on-prem servers, so you could just scan your libraries and be done.

Today’s systems are anything but easy. Deployment happens much faster, but the attack surface of an application is much greater. Many monoliths have been replaced by service architectures: smaller, limited-function programs working together in virtual containers in scalable cloud computing. Each container is capable of communicating with every other container, so the operating system is distributed across an orchestrated network that grows in complexity as additional services enter the ecosystem. Services often communicate over open networks through API or RPC calls, which can then be exploited by malicious actors.

Today, developers can import new libraries with a few lines of code and deploy that library to a production container in minutes. But as the Log4j issue has shown us, your dependencies can change or become vulnerable long after you’ve been using them. The industry is highly dependent upon open-source software they did not write and cannot fix on their own.

This evolution has made a significant improvement in the software industry’s ability to deliver quality software quickly. The industry calls this integration into the SDLC “shift left” because it has shifted these processes earlier (toward the left) in the SDLC. Software engineering practitioners are thinking beyond the code itself as design, infrastructure, and testing are becoming primary concerns for developers.

But there are other concerns—namely security—that need to move from an afterthought to a primary concern. In this article, we’ll talk about the movement to shift security concerns left, new security challenges cloud-native teams face, how observability is the key to solving these challenges, and what shifting left looks like for real-world developers.

When a concern is integrated into the SDLC—shifted left—it becomes part of every step. The requirements and design consider it, development implements, tests cover it, and it’s deployed to infrastructure provisioned with it in mind.

Look how tightly testing has been integrated into the SDLC. While it is a phase in the SDLC, philosophies such as test driven development make it part of every phase of the cycle. Modern software teams increasingly make defect corrections either as the work is being done, or as close to that time as possible. Found a bug? Fix and redeploy to your production environment in a few minutes. Or better yet, spin up a test environment for your specific pull request right now and test it before it makes it into the main branch.

The more concerns—testing, infrastructure, security—the full team understands, the better decisions they can make, and the less work has to be corrected after code has been deployed (when the cost to fix is far greater and unhappy users or clients notice the issues as well). The speed and flexibility of today’s cloud-enabled technologies further empower software teams.

Shifting left is about moving these traditionally holistic concerns earlier in the process, where teams can act on them during development. For decades, software security and testing were predominantly on the right end of the SDLC; that is, outside of the standard cycle of design, develop, and test. Shifting left brings those processes into the development cycle where they become part of design, development, and testing.

Shifting left requires additional work for developers, as it forces you to expand beyond your historical duties. An increasing number of items have become first-class concerns in the SDLC, as developers now manage design, architecture, infrastructure, and testing using code, DevOps practices, and advanced CI/CD pipelines. Yes, you are busier than ever, but the trade-off is that you are more involved in holistic aspects of your code’s design and execution.

The complexity of the cloud and cloud application architectures means that you can no longer keep security separate from what developers work on every day.

Cloud-native application stacks and containerization architectures give software developers great power. They can introduce new microservices and APIs in response to business demands with very low technical overhead. While this greatly accelerates the pace at which you can get your software into customers’ hands, it also substantially increases the complexity of back-end platforms and expands their attack surface.

Security considerations on legacy systems were simpler. They were often installed on user desktops or deployed as monolithic architectures on a single server. They had the benefit of limited exposure, as all libraries and other dependencies were on-prem and easy to scan for vulnerabilities. But as we know, these simpler, monolithic applications are becoming a thing of the past.

Today every piece of an application and the SDLC often takes place in a cloud environment. Not only does the production application run in the cloud, testing environments do too. You may be able to test a single service locally, but testing the full set of services requires a bit more room. Many code repositories have become SaaS tools as organizations have grown. Build processes and tests often happen seamlessly within these tools.

It’s very difficult—if not impossible—to secure a system that can be so easily changed by an open-source library and a bit of integration code. These free software packages aren’t always transparent about their security vulnerabilities or communication requirements, either. So, you are left to determine what APIs are communicating and how. There’s no easy way to create a map of that communication and check it against the software specifications.

Most of us are familiar with the old school method of debugging through print statements. Monitoring and logging take this a step further, providing records of system behaviors. But if you want to really know what’s going on in a system, reading logs after noticing problems won’t cut it. You need to know what’s happening in real time.

Many organizations have a significant number of microservices and APIs in production. There’s only so much you can keep in your head and in written documentation, especially for services that evolve regularly.

Observability tools are very useful for understanding the security posture of such an ecosystem. You don’t understand your security posture unless you can visualize it. A lot of tools are command-line-based but don’t show you the big picture at a glance. Instead of poring over logs to trace an error through a chain of services, a good observability tool will make problem areas more visible.

At Cisco, where I work, we created an open-source tool called APIClarity for Kubernetes clusters that helps developers map out what APIs are communicating with each other. Few tools offer full-stack observability, but tools like API Clarity can help. It utilizes a service mesh framework to capture and analyze API traffic to determine whether it’s something anomalous that requires further inspection.

When an application uses an API, it’s not always one that’s part of your architecture. Plenty of times, those are external dependencies called as API endpoints. For those, you’re relying on a spec for information on how it works, and that spec may be inaccurate or outdated. The traffic, though, doesn’t lie.

Observing a cloud-based application in action clarifies the work a development team needs when shifting left to make their software more robust and secure.

Shift left security means looking at security threats from a development perspective. Security and privacy are no longer “problems for the security team,” but a code-level concern. Code is the cause of—and solution to—most security issues, so making those issues a first-class concern in the development process means you’ll end up with a safer product with less development time. In cloud-native architectures, developers are empowered to fix a wide range of problems because so much is defined within the code.

For a developer, security means adding extra checks into their work:

- How do I make sure code is secure and written to spec?
- What does communication look like between containers?
- What do APIs look like for various systems?
- What external APIs am I consuming?
- How do I minimize my infrastructure’s software attack surface?
- How do we know what systems need to communicate?
- Can we visualize these communication patterns in a map or a sequence diagram?

Adding these checks in the SDLC is a chore. As a programmer myself, I understand. Developers want the benefits of security tooling without any downsides. The more friction something adds to your workflow, you’re less likely to adopt it.

Shifting left must be more than an occasional reminder to scan some code. Security reviews should be part of the development cycle, like a checkpoint in a CI/CD workflow. The DevOps team can automate some of this, though everyone should get involved to determine where it makes the most sense. More engineering departments are picking up DevOps skills, and those skillsets translate easily into skills that DevSecOps roles require.

I like to joke that security is great until it changes how I write code. Security should not get in the way of a developer closing tickets in a sprint. On the other hand, application features should not come at the expense of application security.

If security shifts left correctly, little changes for a developer. Security checkpoints get added, perhaps with a gate that doesn’t affect most workflows. A lot of security checks can be automated: some can integrate into existing CI/CD pipelines; some can integrate directly into IDEs and other everyday tools. It should be as convenient as an automated test suite.

With a tool like Cisco Secure Application, it’s as simple as adding an application policy. If the new code meets the security policy, then build and deploy processes continue without a hitch. If not, there’s some debugging to do. Simple tools that keep developers happily working are the keys to building a more secure SDLC.

Instead of designing solutions that ignore security, developing code that needs to be patched, and running tests that miss common security concerns, your SDLC should include checks on security at every step of the way. This means thinking beyond the work of writing business logic code and considering threats to application and infrastructure security.

“Shift left” means we get more out of the SDLC. We’re no longer relegating privacy, security, and quality to a secondary priority—instead, they become first-class citizens in our development organizations.

We’re no longer saying things like, “I do dev and you do DevOps.” I was discussing this with Steven Augustus, head of open source at Cisco and part of the Kubernetes tooling SIG. He told me that everyone wants to be called something different: SREs, DevOps, etc. People want to specialize, and that’s great, but the organization needs to cohere and communicate. If it all converges, it’s not so different and isolated.

In a sense, shifting left allows people to get more work done themselves. Most orgs don’t have these roles well-defined anyway, so we think shifting left is an inevitable reality for software teams.

At Cisco, we want to work with the community to learn and grow. We are working with developers all over the world, building open source tools to help develop better, more secure applications.

We have a lot of Kubernetes tooling for shifting left, including observability, optimization, and a service mesh. A lot of people understand DevOps, but you may not be able to implement it due to the organizational structures and working processes in your company. If you’re going to KubeCon and want to learn more, Cisco and I will be there with my podcast, Cloud Unfiltered. Come visit us!

The Stack Overflow blog is committed to publishing interesting articles by developers, for developers. From time to time that means working with companies that are also clients of Stack Overflow’s through our advertising, talent, or teams business. When we publish work from clients, we’ll identify it as Partner Content with tags and by including this disclaimer at the bottom.
