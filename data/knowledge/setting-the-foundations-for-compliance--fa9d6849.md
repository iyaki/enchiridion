---
title: "Setting the foundations for compliance"
notion_id: fa9d6849-5930-469f-9c18-7498a6e5f0cb
notion_url: https://app.notion.com/p/Setting-the-foundations-for-compliance-fa9d68495930469f9c187498a6e5f0cb
last_edited: 2023-01-27T18:02:00.000Z
source_url: https://github.blog/2023-01-26-setting-the-foundations-for-compliance/
tags: ["Article", "Github Blog", "English", "DevOps", "Site Reliability Engineering", "Information Security", "Change Management", "Project Management"]
---
While compliance is foundational to delivering software around the world, there may be instances where developers get frustrated with policy enforcement slowing down their workflow.

Since compliance is what enables the world’s software to be run across regions and enterprises with different security, data, and privacy requirements and regulations, we at GitHub are well-practiced in balancing business needs with developer happiness. After all, [100 million developers](https://github.blog/2023-01-25-100-million-developers-and-counting/) and 90% of the Fortune 100 build their software on GitHub.

Let’s take a deep dive into how you can give your developers the tools to do their best work and also meet your business’s compliance needs.

## Compliance, defined

A strict dictionary [definition](https://www.merriam-webster.com/dictionary/compliance) of compliance focuses on following a desire, demand, proposal, or regimen or to coercion; as well as, conformity in fulfilling official requirements. From this definition, you can see why compliance isn’t often discussed in the development community that thrives on flexibility, openness, and collaboration.

A definition that is more developer and cloud-native friendly can be found in the description of an [Office of Compliance](https://martinfowler.com/articles/devops-compliance.html) (OoC). An OoC views compliance as “the desired set of outcomes to achieve and the process by which systems must be validated before they can be deployed to production environments.” Defining an ideal state and the process to reach that state sounds a lot better than being coerced to follow the official requirements.

## Setting the groundwork for compliance

While compliance may seem like an overwhelming concept to tackle, there are some very practical concepts that should be understood first to help set you up for success.

### Understanding your population

This may seem obvious, but it needs to be stated that you can’t address your software development compliance needs if you don’t know where your code is and who your developers are.

If your code is centrally managed in a platform like GitHub, you’ll have this first concept covered. If you don’t have your code in a central platform, you’ll need to first try and find all of the different tools being used and somehow get a list of all the developers that have access to these tools. As you can imagine, this can be a time consuming process and there is no guarantee that your population is complete.

### Understanding access

Once you understand your population, you’ll need to understand who has access to what. Your code is probably some of the most important information in your organization and you need to ensure it’s secure but available to your developers. The [CIA Triad](https://www.microsoft.com/en-us/security/business/security-101/what-is-information-security-infosec) offers a structure to help ensure your code is secure and only people who need to access it can do so.

- **Confidentiality**: organizations should enact measures that allow only authorized users access to information.
- **Integrity**: importance of accurate, reliable data, and permit no unauthorized user to access, alter, or otherwise interfere with it
- **Availability**: guarantee that authorized users have dependable, consistent access to data as they need it.

Some ways that GitHub ensures that individuals and enterprises can control access to their code are by enabling [access permissions](https://docs.github.com/en/get-started/learning-about-github/access-permissions-on-github) for all of your [repositories](https://docs.github.com/en/repositories) and other resources so that each user only has access to what they need to do their job (also known as the principle of least privilege). The concept of an enterprise account is built into the core of GitHub. This enables you to manage access holistically, integrate with your existing provider and synchronize teams for ease of management, and set guard rails to ensure a secure foundation for your engineering teams.

With [100 million](https://github.blog/2023-01-25-100-million-developers-and-counting/) of the world’s developers on GitHub, we ensure that it’s available to all individual and enterprise developers whenever they need it. We publish [GitHub availability](https://www.githubstatus.com/) reports on a monthly basis as a commitment to our communities, making sure their code is available when they need it.

Using a central tool like GitHub Enterprise can help reduce the management overhead required. You can check out our recent blog post on [3 tips on consolidating your toolkit](https://github.blog/2022-10-26-3-strategies-for-consolidating-your-toolkit-and-boosting-productivity/) if your current tooling environment is overly complex and you need some help.

### User attestation

At this stage you understand where your code is, your population of users, and the controls to ensure the CIA Triad are addressed. Next, you need to have processes in place to periodically attest to the access that your users have. If an employee moves to a different role or leaves the company, a process needs to be in place to ensure their access is revoked.

In future blogs, we’ll explore how automation with [GitHub Actions](https://github.com/features/actions), and your existing developer workflows with [pull requests](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests), can be leveraged to optimize compliance related to user attestation.

### Continuous Compliance

Everyone is probably familiar with the term [Continuous Integration (CI) / Continuous Delivery (CD)](https://resources.github.com/ci-cd/), but the concept of doing something continuously can also be applied to compliance as well. One of the major benefits of CI/CD is [speed](https://resources.github.com/ci-cd/#ci-cd-benefits). Imagine a world where your compliance testing and audits are fast and as seamless as your DevOps CI/CD pipeline.

Continuous Compliance won’t magically appear overnight, but if the foundational pieces of compliance are understood and you have started putting the groundwork in place, you’ll have a few major benefits that will make compliance testing and audits faster.

### 1. No surprises.

One of the worst scenarios you can have when compliance testing or audits start is to be surprised and ill-prepared. Nobody wants to have to create new artifacts because testers or auditors are asking for them.

### 2. Common understanding.

The major goal of the [DevOps Audit Defense Toolkit](https://cdn2.hubspot.net/hubfs/228391/Corporate/DevOps_Audit_Defense_Toolkit_v1.0.pdf?t=1453925000299) was to improve communication, create a common understanding and “educate IT management and practitioners on the audit process so they can demonstrate to auditors they understand the business risks and are properly mitigating those risks.” Improved communication and sharing of information between IT and Auditors means more efficient audits, and, according to Forrester, [substantial cost savings](https://resources.github.com/forrester/).

### 3. Built into developer workflows.

If compliance has been designed from the ground up in your software delivery lifecycle, your developers will be meeting required controls and gathering artifacts as part of their daily flow.

### AI-Enabled compliance

At GitHub, we are already leveraging machine learning to help [find security vulnerabilities](https://github.blog/2022-02-17-leveraging-machine-learning-find-security-vulnerabilities/) that are important for security and compliance. AI is starting to be adopted in the manufacturing industry to aid[ auditing and process inspection](https://www.businesswire.com/news/home/20221215005002/en/DarwinAI-brings-AI-powered-Visual-Quality-Inspection-to-market-and-announces-a-funding-round-led-by-BDC-Capital-to-accelerate-growth). Gartner has also identified generative AI as a technology that [banks will be leveraging ](https://www.gartner.com/en/newsroom/press-releases/2022-05-24-gartner-identifies-three-technology-trends-gaining-tr)soon for fraud detection and risk modeling.

The advantages and potential of using AI and machine learning for continuous compliance are very compelling. In a future blog, we’ll look more at what’s around the corner.

## Next steps

Once these foundational aspects of compliance are well understood, you can start thinking about the next steps to implement these concepts. That’s the hard part. In our next post in this series, we’ll be looking at several practical ways to meet compliance needs while keeping your developers happy and in the flow.

Ready to increase developer velocity and collaboration while remaining secure and compliant? See how [GitHub Enterprise can help](https://github.com/enterprise).
