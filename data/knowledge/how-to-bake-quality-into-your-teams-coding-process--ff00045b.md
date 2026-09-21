---
title: "How to bake quality into your teams’ coding process"
notion_id: ff00045b-4746-4f87-b45c-5178962a9ae3
notion_url: https://app.notion.com/p/How-to-bake-quality-into-your-teams-coding-process-ff00045b47464f87b45c5178962a9ae3
last_edited: 2023-02-13T19:28:00.000Z
source_url: https://leaddev.com/process/how-bake-quality-your-teams-coding-process
tags: ["Article", "LeadDev", "Programming", "System Design / Software Architecture", "DevOps", "Product Management", "Testing", "Project Management", "Site Reliability Engineering", "Communication"]
---
As engineering leaders, we’re responsible for establishing clear patterns for writing [high-quality code](https://leaddev.com/building-better-software/guide-measuring-and-improving-code-quality). Without these standards, our teams will accumulate technical debt very quickly, and before long, features will become too costly to build, release cycles will stretch out for far too long, and our users will ultimately suffer.

In this article, I’ll share some straightforward strategies for ensuring a [solid baseline of code quality](https://leaddev.com/tech/four-pillars-code-health) for your teams, beyond writing documentation.

## The challenges of documenting code conventions

Maintaining technical documentation for the product you’re building is essential and challenging work. It isn’t only vital for cataloging years of company lore and the reasons for making many individually impactful choices over the lifespan of a system, but for detailing all of the code conventions and standard practices adopted by your engineering teams.

Creating a clear, up-to-date description of what your company is building, and how the teams are building it, is essential for onboarding new hires and providing direction on maintaining code quality throughout your systems. The goal is to have all of your engineers writing code in a singular, cohesive style, following a common set of principles, avoid wasting time with analysis paralysis, and working together efficiently.

However, in practice, there are three common obstacles to achieving this:

1. The first is that it can be **difficult to update your documentation alongside your codebase**. As your teams commit changes and ship releases, all of these updates need to be reflected within the documentation. This is time-consuming work, and in the daily grind of sprints and release cycles, it’s often neglected.
2. Another common issue is **consistency**. When it comes to style guides, it’s not as simple as writing down how you should write the code: you must also apply it in practice. No matter how often a particular style choice is corrected during code reviews, keeping your code consistent across all of your teams and projects is nearly impossible. Once that consistency is lost, the style guide becomes little more than an idea without authority, and the code continues to drift.
3. The last obstacle is **discoverability**. As your organization grows in size, the challenge of eliminating silos becomes steeper and more critical, as the number of teams and projects grows and product use increases. Similarly, your documentation becomes more fragmented and spread apart, increasing the likelihood that different teams will invest time in solving the same recurring problems within your codebase.

## How to bake quality into your coding process

Fortunately, there’s a simple approach that can address all of these problems. Rather than constantly having your documentation trying to catch up with the code, let the code itself maintain its quality. By making a few significant changes to your development process, you can avoid the pain of having an increasingly divergent and fragmented codebase. There are three core steps to this approach:

### 1. Outsource code reviews to robots

The first step is to ensure you have a set of compilers, linters, and formatters working on every piece of code that you ship. There are many official and community-supported tools, regardless of what language or environment you’re working in, and whether it’s application or infrastructure code (I won’t bother to list them here since they can be discovered easily through a simple web search).

No matter what tools you choose, make sure they’re integrated into your CI process, so they can enforce a strict set of patterns and stop any commits that deviate from the standards. These tools will serve as the most effective code reviewers in your organization, and will provide an enormous productivity boost.

In addition to the essential tools for compiling, linting, and formatting your code, there are many paid products available that take a more comprehensive approach to static code analysis, and offer the ability to identify [smells within your code](https://martinfowler.com/bliki/CodeSmell.html). Depending on how mature your company is or the problems your teams are struggling with, these tools may justify the cost rather than dealing with a buggy, legacy project that’s core to your business and needs more solid documentation. Otherwise, you may be better off sticking with community-supported tools until a specific need arises.

Once you have all these tools in place, you’ll want to settle on a common set of rules for all of them to follow. When it comes to formatters, they are [notoriously rigid](https://black.readthedocs.io/en/stable/index.html) [by design](https://prettier.io/docs/en/why-prettier.html) to prevent the many nit-picking arguments that developers tend to have over style, so everyone can get back to working on problems that actually matter. The same principle should be applied to all configurations for your static code analyzers: make straightforward decisions with a small group of people without much fanfare. If anyone has any issues with the style choices made later down the road, you can create a simple process for making adjustments over time. Still, typically, people can settle into the style of the projects they're working on and quickly forget about whatever gripes they may have once had.

### 2. Integrate code quality with templates

Now that you’ve developed a centralized set of rules for static code analysis and configurations for your CI tooling, the next challenge is propagating these configurations throughout your projects. Where possible, you should publish these configurations in their own modules with proper versioning. Once these modules are integrated into your existing projects, an automated dependency updater, such as [dependabot](https://github.com/dependabot), can manage any updates shipped to these modules. It’s generally a good idea to use an automatic dependency tool for overall code quality and security, as it relates to any other third-party packages your projects may depend on. Regarding your core configurations, it becomes a powerful way to ship changes to all projects.

Now that existing projects are sorted, you need a simple way to bootstrap new projects that applies all of these internally published configurations. One way to do this is with the use of template projects. [GitHub](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-template-repository) and [GitLab](https://docs.gitlab.com/ee/user/project/working_with_projects.html#create-a-project-from-a-custom-template) both offer methods for creating new projects from a template, but the same effect can be achieved with a simple starter project that is then cloned or forked when kicking off a new project.

Template projects offer several benefits. You can skip a lot of the busy work that comes with initializing a new project and have something ready to run and deploy in just a couple of steps. You can also write the starter with other essential style choices in mind, particularly the kind that can not easily be imposed with a linter. This might include file naming, folder structure, test scripts, etc.

A compelling approach to template projects is to build them with minimally functional modules. Of course, a downside to this is that it’s a decent chunk of code that gets immediately stripped out of every project after initialization. However, the benefit is demonstrating to newcomers the patterns for developing applications with high code quality, building towards a paved road for development.

### 3. Iterate on and improve code quality as a team

At this point, you have centralized configurations for static code analyzers and CI tools. You have them all written inside a project template, and a simple set of components written into the project to demonstrate proper development patterns for the given language, framework, and stack. The last piece here is to make sure your teams iterate and build on the lessons learned in battle.

As engineers build with the set of tools chosen for your projects, they will continue to find ways to increment and improve on them. It’s important to have a process established for developers to contribute back into the core tooling and templates, preferably in the style of [innersourcing](https://en.wikipedia.org/wiki/Inner_source). Building a development culture that promotes and encourages innersourcing and cross-team collaboration enables teams to develop higher quality code, without being siloed to solve similar problems.

In turn, you’re teaching them to grow and boost one another. It’s one thing for a junior developer to pull a ticket over into the done column for their team; getting a PR merged that solves an issue shared by every team in their organization is much more significant. These types of wins can really boost your teams’ spirit, and inspire team members to search for new ways to contribute to the overall code quality of your product.



<!-- unsupported block: link_to_page -->
