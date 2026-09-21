---
title: "LocalStack - A fully functional local cloud stack"
notion_id: a959a29d-89d6-4390-b936-1d6c8823e188
notion_url: https://app.notion.com/p/LocalStack-A-fully-functional-local-cloud-stack-a959a29d89d64390b9361d6c8823e188
last_edited: 2026-09-21T17:08:00.000Z
source_url: https://www.localstack.cloud/
tags: ["English", "AWS", "System Design / Software Architecture", "DevOps", "Programming", "Untried", "Tool", "Service"]
---
## The cloud wasn't built for modern software development.

The cloud wasn't built for modern software development.

It was built to run applications, not AI-native development workflows. LocalStack enables teams and AI agents to validate code without ever touching the cloud.

Get Started

Talk to Sales

Powering the world’s most innovative development teams

See how LocalStack empowers companies of all sizes. See our case studies →

## You don’t need cloud access to test cloud behavior

LocalStack simulates cloud environments on your own infrastructure so software teams and AI agents can validate security, quality, and reliability faster and more safely than the cloud allows.

70%

reduction in cloud development costs for AI-driven teams

“LocalStack is an essential component of our development toolkit and has greatly enhanced our team’s agility & efficiency.”

Janaka AbeywardhanCo-Founder—CPTO at FABR

30%

faster feedback loops for AI agents validating cloud behaviour

“LocalStack has massively decreased the time when the local environment is breaking for our developers. That was a big pain point, and it’s not anymore!”

Yashar HosseinpourLead Developer at Neurolytics

10×

more cloud interactions tested per development cycle

“LocalStack enables faster iterations and quicker feature releases and translates to substantial financial savings.”

Vincent FletcherCo-Founder & CPO at CartonCloud

## Fix bugs that you can’t see in the cloud

LocalStack makes it easier for software teams and AI agents to fix security risks, bugs, and defects that they struggle to diagnose in the cloud.

Using your existing application code and IaC resources, AI agents deploy to a local container that simulates your production app environment’s behavior. No access, cloud costs, or provisioning delays required.

LocalStack surfaces the defects that will cause failures when you push to production so AI agents can validate reliability early in the development pipeline. Spend less time investigating bugs and more time building code that ships to prod.

Agents get instant feedback to validate that code will behave as needed. Fix bugs rapidly with no human intervention or cloud provisioning delays.

Once validated, software teams simply update the endpoint to ship updated code to a staging or production environment. Speed up the pipeline without sacrificing quality or reliability.

## LocalStack for AWS

Eliminate the complexity that holds back early-stage development and testing with support for 120+ AWS services and advanced testing and debugging capabilities.

Find out more

## LocalStack for Snowflake

Speed up workflows and optimize costs for complex data queries using LocalStack’s local development platform for Snowflake.

Find out more

## LocalStack for Azure

Public preview of LocalStack for Azure is coming soon. Join the waitlist to be the first to know.

Join the waitlist

Coming Soon

Use LocalStack to fix bugs in your cloud application environments faster and more effectively. Browse our list of cloud emulators to understand support for your application.

Explore AWS services

Explore Snowflake features

## You are in good company

500M+

500M+

Docker pulls

3M+

3M+

Weekly Docker Pulls

18M+

18M+

Weekly LocalStack Sessions

“LocalStack allows us to move away from costly shared test environments to local environments where developers can experiment freely and test aggressively without stepping on each other.”

John Calhoun,Chief Cloud Architect and Chief Security Officer, PayNetWorx

“With LocalStack we can now empower our devs to iterate quickly without having to perform numerous code commits & waiting for AWS pipelines”

Kevin BretonVP of Engineering at KnowBe4

See how LocalStack empowers companies of all sizes. See our case studies →

## Frequently Asked Questions

What is LocalStack?

LocalStack is a local cloud development platform that makes it easier for software developers and AI agents to validate and ship cloud applications by developing, testing, and debugging in a simulated cloud environment that runs on the customer’s safe, local infrastructure.

With support for AWS, Snowflake, and a soon-to-be-released platform for Microsoft Azure, LocalStack uses the cloud service providers’ actual APIs to simulate cloud behavior in a lightweight container that customers can run locally, including on developer laptops, in CI/CD runners, and in centralized infrastructure that is distributed for access by developers and AI agents.

This allows AI agents and software developers to test how the application will behave once deployed to the production environment without needing to provision cloud infrastructure.

AI-native software development workflows increase the volume of new application code, creating bottlenecks in testing and validating the security, quality, and reliability of those applications before they ship to the production cloud environment. In many cases, a human intervenes to perform basic testing and debugging for AI-generated code in traditional cloud-native environments. This slows down software delivery, increases costs on the cloud bill, and makes it difficult to identify security risks and other bugs that hide in the logs when deploying to cloud infrastructure. Using the containerized nature of the local environment, LocalStack surfaces bugs, security risks, and other failures that go unseen in cloud-native staging or QA environments. And since code updates are re-deployed to the container, AI agents and software developers validate their fixes instantly–without the natural delays that occur when provisioning cloud infrastructure.

By moving development, testing, and validation processes to the local environment, LocalStack customers have reported 10x improvements in productivity, resulting in faster deployment frequency, lead time for changes, and hours saved in the development pipeline. Other benefits include reductions in charges on the cloud bill and improved application quality and reliability through more effective testing.

LocalStack is trusted by more than 1,500 organizations globally, including software teams at some of the world’s largest organizations building the world’s most popular applications. Adoption has scaled above 400 million Docker pulls all-time. To learn more about bringing LocalStack to your workflows, talk to our sales team.‍

How accurate is the emulation? Will my code work the same in production?

LocalStack is built for high fidelity. Your application uses the same API calls, SDKs, and Infrastructure as Code tools (CDK, Terraform, Pulumi) that it would against a live cloud environment, so what works locally works in production. We continuously validate our emulation against real AWS APIs to keep parity tight and we partner with cloud service providers to ensure our emulated services align with the latest updates to theirs.

Which cloud platforms does LocalStack support?

LocalStack currently emulates AWS (120+ services) and Snowflake. Public preview for LocalStack for Azure is coming soon. You can join the waitlist to be the first to hear when LocalStack for Azure is available.

Can LocalStack support production applications?

LocalStack is built for development and testing, not production. Since LocalStack simulates how the cloud services will behave when deployed, it’s a safe, fast environment to build and validate cloud applications before they ship. However, the cloud is still the right fit for supporting the production demands of a customer-facing application.‍Once your code is ready, it deploys to real AWS or Snowflake as normal. Think of LocalStack as the place where mistakes are cheap and caught more easily. Iterate freely, break things, fix them, and arrive at production with confidence rather than crossed fingers.

Is LocalStack just for AI agents?

No. While many of our users integrate LocalStack into their AI-native software development workflows, many also rely on LocalStack for on-demand staging environments and CI/CD integration.

Many of our users rely on LocalStack to enable concurrent testing by eliminating shared cloud staging environments for on-demand development and iteration. Since cloud-hosted staging environments often crash when multiple developers push code simultaneously, this forces developers to take turns in the staging enviroment–resulting in a queue that delays work and disrupts developer productivity.

Since LocalStack can run independently without cloud access or costs, software teams rely on LocalStack to build, iterate, and test on-demand.

Meanwhile, CI/CD integration allows those teams to speed up integration and other tests by eliminating the provisioning delays that can slow those stages down.

Think of LocalStack as the platform for building and scaling cloud applications, regardless of whether it’s deployed by an AI agent, software developer, or CI pipeline.

How does LocalStack fit into CI/CD pipelines?

You can drop LocalStack into any CI environment — GitHub Actions, GitLab CI, CircleCI, Jenkins — with a single Docker image. Every pull request gets a fully isolated, ephemeral cloud environment with no external dependencies and no shared state between runs. Faster pipelines. Zero cloud costs for test infrastructure.

How does LocalStack help teams building agentic AI workflows?

AI agents that interact with cloud services need a safe place to run, fail, and iterate without risk.

LocalStack gives every agent run its own local cloud development sandbox that surfaces the bugs and defects that will disrupt application behavior in production, all without granting the agent access to a cloud account, driving up costs on the cloud bill, or forcing them to wait for code to provision to the cloud.

AI agents can simply deploy the application to the LocalStack endpoint, review the recommended issues in the environment, and validate their code fixes instantly.

Once validated, software teams can deploy that code to the production environment seamlessly by simply swapping out the endpoint.

Can LocalStack run on Kubernetes?

Yes. LocalStack can be deployed directly into a Kubernetes cluster using the LocalStack Operator, a Helm chart, or standard Kubernetes manifests. When running in Kubernetes mode, services that would normally spin up Docker containers (Lambda, ECS, RDS) create Kubernetes pods instead, giving you dynamic scaling and native cluster orchestration. This makes LocalStack a fit for shared hosted dev environments, platform-level dev sandboxes, and Kubernetes-native CI/CD pipelines.

‍Talk to our sales team to discuss Kubernetes integration.

Does LocalStack integrate with internal developer platforms?

Yes. LocalStack is designed to fit into platform engineering workflows, not sit outside them. On Kubernetes, platform teams can use the LocalStack Operator to provision isolated cloud environments on demand—surface these through any IDP that manages workloads (Backstage, Port, and similar). LocalStack's ephemeral instance API lets you spin environments up and tear them down programmatically to enable so self-service dev sandboxes.

Enterprise plans also include SSO (Azure AD, Okta) and SCIM provisioning, so identity and access management flows through your existing stack.

Want to explore how LocalStack fits your platform? Talk to our team.

What does LocalStack actually cost me (or save me)?

LocalStack’s products offer pricing that varies based on use case.

The LocalStack for AWS Hobby plan provides free access to basic service emulation for 30 AWS services for non-commercial users. Hobby plan can support basic AWS service emulation to streamline testing and debugging, but lacks many of the advanced features, integrations, and security capabilities that large software teams need.

LocalStack for AWS Base plan costs $39 per user per month when billed annually, and $45 per user per month when purchased on a monthly commitment. The Base plan adds support for more AWS services and advanced features for collaboration, testing, and debugging that are better aligned to small teams of developers.

LocalStack for AWS Ultimate plan costs $89 per user per month and expands support for AWS services as well as more advanced capabilities. This includes automating the generation and testing of IAM policies and the ability to replicate live AWS resources in the local environment for more realistic testing.

LocalStack for AWS Enterprise is priced on a custom basis and offers advanced capabilities, centralized delivery and management, SSO, enterprise-grade support, and the ability to run LocalStack fully offline or in air-gapped environments.

LocalStack for AWS no longer imposes limits on CI/CD builds on a monthly basis. While legacy pricing provided a monthly set of credits for CI/CD builds in LocalStack–which varied by plan–LocalStack removed those limitations in March 2026. Today, all LocalStack for AWS plans offer unlimited operation in CI/CD pipelines.

LocalStack for Snowflake Base plan is available at $29 per user per month when billed annually and $35 per user per month when billed monthly. The Base plan includes CI/CD integration and the ability to snapshot and restore environment state on-demand.

LocalStack for Snowflake Enterprise plan is priced on a custom basis and offers customized usage, enterprise-grade security, and dedicated support.

Visit our Pricing page to learn more.https://www.localstack.cloud/pricing

Is LocalStack just for individual developers, or does it scale to teams?

Both. Individual developers use LocalStack to build faster and test locally without cloud access. Teams use it to improve AI-native testing and debugging, eliminate shared staging environments, and speed up CI/CD pipelines. Enterprise plans allow centralized management and distribution and advanced cloud emulation and development capabilities, and include SSO, team management, and dedicated support.

How do I get started?

To get your AI agent started with LocalStack, visit this page.‍You can also browse our Pricing page to understand which plan aligns with your needs.

You'll need Docker — then it's a single pip install localstack or docker pull localstack/localstack to have your first local AWS environment running in under a minute. From there, point your existing AWS SDK or CLI at LocalStack and keep building exactly as you do today.

## Launch yourself in the world of local cloud development

Try for free

Talk to Sales
