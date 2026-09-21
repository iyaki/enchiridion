---
title: "Platform engineering for engineering managers"
notion_id: e2789cdf-9561-4d77-9f61-a1da4b479f18
notion_url: https://app.notion.com/p/Platform-engineering-for-engineering-managers-e2789cdf95614d779f61a1da4b479f18
last_edited: 2023-04-21T19:30:00.000Z
source_url: https://leaddev.com/process/platform-engineering-engineering-managers
tags: ["English", "Site Reliability Engineering", "DevOps", "SysAdmin", "Infrastructure", "Article"]
---
Making it easier for software developers to do the right thing more easily should be a major priority for engineering leaders. The emerging practice of platform engineering could just be the answer.

Engineering leaders are always pushing their teams and organizations towards faster and safer software delivery practices that enable their organizations to keep up with ever-changing digital business needs. In the new world of microservices, [CI/CD](https://leaddev.com/series/achieving-cicd-excellence), and [DevOps](https://leaddev.com/process/devops-engineering-managers), making best practice the easy thing to do can help accelerate the journey of features from conception to market.

Many of these software development trends have recently converged into the idea of platform engineering, where a dedicated internal platform is built to give developers the tools and pathways to more quickly and safely take an idea from their laptop to a user.

## What is platform engineering?

Platform engineering is a discipline that aims to produce toolchains and workflows that developers use to write, test, and roll out their code. The goal is for these tools to be self-service from the perspective of the developers. In other words, a developer shouldn’t have to pester someone on the platform engineering team to help make use of these tools to deploy or test their application code. Rather, if the platform engineers have done their job, the devs should be able to safely do everything on their own.

There’s [a lot of complexity involved in the deployment of modern applications](https://www.infoworld.com/article/3639050/complexity-is-killing-software-developers.html), all of which adds to a developer’s already-high cognitive load. In other words, a lot of developers are spending too much energy thinking about the plumbing of an application and not enough on its unique business logic.

That plumbing is important, of course, but it won’t necessarily differentiate the application you’re working on from your competitors. Worse, because it can take a lot of experience to understand how all that plumbing works, it’s more likely than not that your most experienced and skilled developers will find themselves burdened with these tasks. From an engineering manager’s perspective, that’s a nightmare of misallocated resources.

The goal of platform engineering is to provide ready-to-use tooling that allows developers to deploy their code without having to worry too much about the underlying infrastructure. This tooling nudges developers into harmonizing their code with the overall infrastructure choices and strategies that the organization has made.

Swedish music streaming service Spotify, an early advocate for platform engineering, calls this the [Golden Path](https://engineering.atspotify.com/2020/08/how-we-use-golden-paths-to-solve-fragmentation-in-our-software-ecosystem/), while the video streaming service Netflix uses the term [paved roads](https://www.infoq.com/news/2017/06/paved-paas-netflix/). The idea is that the platform makes it easy and obvious how to write code that can ultimately be deployed to a standardized environment. While it’s possible (and occasionally necessary) to wander off the paved roads, traveling along them should be the default behavior.

## What is an internal developer platform?

A platform engineering team spends its time building and maintaining an internal developer platform (IDP). Even though an IDP is intended for internal use within your organization, your platform engineering team needs to have the mindset that developers are their customers, and they should work to tailor the IDP to their needs accordingly.

An IDP is an abstraction layer that sits between developers and the infrastructure where they’ll deploy their code. It can take several forms: devs may be able to use its various tools via a command-line interface, through various application programming interfaces (APIs), or even a graphical user interface (GUI) – also sometimes referred to as an internal developer portal.

Developers should be able to use an IDP to change configurations, manage deployment automation, and spin up execution environments and resources. The IDP makes this all happen by integrating with the organization’s automated deployment pipeline. An IDP may also create application configuration templates, assign clusters to specific environment types, and provide facilities for managing permissions.

The difference between an IDP and a platform-as-a-service (PaaS) like Cloud Foundry or Heroku, is that it is tailored specifically to your organization’s ways of doing things, and can be modified over time as part of an ongoing dialogue between your operations and developer team members. A PaaS cannot be customized so extensively: your team must adapt to its way of doing things, rather than the other way around.

## Platform engineering vs DevOps vs SRE

If you’re all-in on bringing together the disciplines of development and operations, many of these principles should sound familiar. After all, the whole point of DevOps is to break down the silos between those who write code and those who deploy it.

While it may sound like platform engineers are doing the operations work for their developers, platform engineering is more an evolution of the philosophy of DevOps, rather than a repudiation of it. By building a self-service toolset that your developers can use as they write and deploy application code, platform engineering can help further the same goals that DevOps was dreamed up to achieve.

In many ways, your platform engineering team is a successor to a traditional operations team. The key difference is that they are working with developers continuously to refine their toolsets and pave the right paths for them to follow.

Similarly, site reliability engineering (SRE) is a modern function tasked with ensuring the whole organization’s underlying infrastructure is as reliable and efficient as possible. Platform engineers, by contrast, are specifically oriented towards the organization’s software delivery processes, and aim to make them as fast and efficient as they can be. A robust platform engineering team can certainly compliment those goals by ensuring that golden paths are safe and reliable ones.

## Platform engineering advantages

Here are some of the key reasons platform engineering may be a worthwhile investment for your organization:

- **Letting team members do what they do best.** By dedicating ops staff to building and maintaining an IDP, you make the best use of everyone’s skills.
- **Reduced complexity and increased velocity.** By establishing “golden paths”, platform engineers can make things simpler for developers, and accelerate product and feature delivery as a result. An IDP can also help new developers get up to speed faster during [onboarding](https://leaddev.com/hiring-onboarding-retention/how-successfully-onboard-remote-engineering-staff-four-weeks).
- **Keep code locked down.** One area we haven’t discussed much so far is security and regulatory compliance. These features are often implemented best when they’re centralized and codified within an IDP, rather than built uniquely into every application and feature.

## Barriers to platform engineering

All that said, there are pitfalls to implementing platform engineering. As an engineering leader, one of your most important tasks is to weigh whether your organization is ready to make such a big shift in how it delivers software.

- **Communication breakdown.** Remember, your team is specifically building an IDP for your developers. That means they need to cater to a spectrum of developer needs and iterate to meet those needs. If the platform ceases to be useful to developers, they will likely find a way to bypass it.
- **Being too rigid (or too flexible).** You don’t want your team constantly scrambling to implement tools that cater to every whim that comes from the developer side – but, at the same time, you don’t want them laying down their own idea of a golden path from high up in an ivory tower and ignoring developer input. The reason you need to keep lines of communication open is to strike a balance that avoids both these negative outcomes.
- **Build a real platform.** As platform engineering gains industry buzz, some organizations may be tempted to simply rebadge their central IT as a “platform team” and call it a day. But you’ll only get real buy-in if you implement the technical and cultural changes we’ve outlined here.

## Platform engineering skills

When it comes time to assemble your platform engineering team, you might look within your company for your best and brightest ops staffers, or look to recruit externally to fill gaps. Either way, there’s a specific set of skills you’ll need to find expertise in:

- Container orchestration, especially [Kubernetes](https://leaddev.com/tech/kubernetes-engineering-managers).
- Microservices and distributed applications.
- Cloud technology.
- API management.
- CI/CD and pipeline management, continuous delivery.
- Infrastructure as code.
- [Observability](https://leaddev.com/tech/observability-engineering-managers), monitoring, and alerting.

Previous experience on a DevOps or SRE team can be helpful, but often isn’t necessary, especially for those who excel in other areas.

While technical skills are important, empathy, collaboration, communication, and documentation are important skills for platform engineers, as a platform engineering team needs the human touch to truly thrive.
