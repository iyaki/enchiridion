---
title: "Deployment reliability at GitHub"
notion_id: fb98ecb5-3f7a-4181-93ec-a7c3e7e43a81
notion_url: https://app.notion.com/p/Deployment-reliability-at-GitHub-fb98ecb53f7a418193eca7c3e7e43a81
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://github.blog/developer-skills/github/deployment-reliability-at-github/
tags: ["English", "DevOps", "SysAdmin", "Site Reliability Engineering", "Article", "Github Blog"]
---
[https://github.blog/2021-02-03-deployment-reliability-at-github/](https://github.blog/2021-02-03-deployment-reliability-at-github/)

Welcome to another deep dive of the Building GitHub blog series, providing a look at how teams across the GitHub engineering organization identify and address opportunities to improve our internal development tooling and infrastructure.

In a previous post of this series, we described how we improved the deployment experience for github.com. When we describe deployments at GitHub, the deployment experience is an important part of what it takes to ship applications to production, especially at GitHub’s scale, but there is more to it: the actual deployment mechanics need to be fast and reliable.

## Deploying GitHub

GitHub is deployed to two types of “targets”: multiple Kubernetes clusters and directly to bare metal hosts. Those two targets have different needs and characteristics, such as different number of replicas, different runtimes, etc.

The deployment process of GitHub is designed to be an invisible event for users—we deploy GitHub tens of times a day (yes, even on a Friday) without impact on our users.

When implementing deployments for a monolithic application, we must keep into account the impact that the deployment process has on the internal users of the tool as well. Hundreds of GitHub engineers work at the same time on new features and bug fixes on the same codebase and it’s critical that they can reliably deploy to production. If deployments take too long or if they are prone to fail (even if there is no impact on users), it will mean that developers at GitHub will spend more time getting those features out to users.

For these reasons, we asked ourselves the following questions:

- How long does it take to get code successfully running in production?
- How often do we roll back changes?
- How often do deployments require any kind of manual intervention?

One thing was sure from the beginning, we needed data to answer those questions.

## Measure all the things

We instrumented our tooling to send metrics on several important key aspects, including, but not limited to:

- Duration of CI builds.
- Duration of individual steps of the deployment pipeline.
- Total duration of a deployment pipeline.
- Final state of a deployment pipeline.
- Number of deployments that are rolled back.
- Occurrences of deployment retries in one of the steps of the pipeline.

As well as more general metrics related to the overall delivery:

- How many pull requests we deploy/merge every week.
- How long it takes to get a pull request from “ready to be deployed” to “merged”.

We used these metrics to implement several improvements to our deployment tooling: we generally made our deployments more reliable by analyzing those metrics, but we also introduced changes that allow us to tolerate some classes of intermittent deployment failures, introducing automatic retries in case of problems.

Additionally, instrumenting our deployment tooling allowed us to identify problems sooner when they happen so that we can react in a timely fashion.

## Better visibility in the deployment process

As we mentioned, GitHub is a monolithic rails app that is deployed to Kubernetes and bare metal servers, with the customer facing part of GitHub being 100% deployed to Kubernetes. When we deploy a new version of GitHub we need to start hundreds of pods in multiple Kubernetes clusters.

A few months ago, our deploy tooling did not print much information on what was going on behind the scenes with our Kubernetes deployments. This meant that whenever a deployment failed, for example due to an issue that we didn’t previously detect in stages before canary, we would have to dig into what happened by directly asking Kubernetes.

At GitHub, we don’t require engineers deploying to understand the internals of Kubernetes. We abstract Kubernetes in a way that is easier to deal with and we have tooling in place to debug GitHub without directly accessing specific Kubernetes clusters.

We analyzed internal support requests to our infrastructure teams and found the possibility to reduce toil by making it easier to figure out what went wrong when deploying to Kubernetes.

For those reasons, we introduced changes to our tooling to provide better information on a deployment while it is being rolled out and proactively providing specific lower level information in case of failures, which includes a view of the Kubernetes events without the need to directly access Kubernetes itself.

This change allowed us to have better detailed information on the progress of a deployment and to increase the visibility on errors in case of failures, which reduces the time to detect a problem and ultimately reduces the overall time needed to deploy to production.

## An SLO based approach to deployment reliability

When deployments fail, there is no impact for GitHub customers: deploys are automatically halted before there can be customer-facing issues and to do so we heavily rely on Kubernetes, for example by using readiness probes.

However, deploying GitHub tens of times a day means that the longer a deployment takes, the less things we can ship!

To make sure that we can successfully keep shipping new features to our customers, we defined a few service level objectives (SLOs) to keep track of how fast and reliable deploying GitHub is.

SLOs are usually defined for things like the success rate or latency of a web application, but they can be used for pretty much anything else. In our case, we started using SLOs to set reliability objectives and keep track of how much time it takes to deploy PRs to production which allows us to understand when we need to shift our focus from new features to improvements to the overall shipping flow of GitHub.

At GitHub we have a dedicated team that is responsible for the continuous deployment of applications, which means that we develop tools and best practices but ultimately also help Hubbers ship their applications. These SLOs are now an integral part of the team dynamics and influence the priorities of the team, to ensure that we can keep shipping hundreds of pull requests every week.

## Conclusion

In this post we discussed how we make sure that we keep Hubbers shipping new features and improvements over time. Since we started taking a look at the problem, we significantly improved our deployment process, but more importantly we introduced SLOs that can guide investments to further improve our tools and processes so that GitHub users can keep getting fresh new features all year round.

## Written by

## Related posts

### The cost of saying yes has changed

The cost of writing code dropped; the cost of owning it didn’t. A framework for deciding which changes are actually cheap in the AI era.

## Explore more from GitHub

### Docs

Everything you need to master GitHub, all in one place.

Go to Docs

### GitHub

Build what’s next on GitHub, the place for anyone from anywhere to build anything.

Start building

### Customer stories

Meet the companies and engineering teams that build with GitHub.

Learn more

### GitHub Universe 2026

Join us October 28-29 in San Francisco or online for GitHub Universe, our flagship developer event uniting people, agents, and the world’s code.

Register now
