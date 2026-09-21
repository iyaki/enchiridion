---
title: "How Good Documentation Can Improve Productivity"
notion_id: 652106f5-7f4a-4719-be82-88862c48ea71
notion_url: https://app.notion.com/p/How-Good-Documentation-Can-Improve-Productivity-652106f57f4a4719be8288862c48ea71
last_edited: 2023-01-17T10:59:00.000Z
source_url: https://shopify.engineering/good-documentation-productivity
tags: ["Documentation", "Programming", "Productivity", "Article", "Shopify Engineering"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

When I joined Shopify last year, I knew its engineering culture was top class. What I didn’t know was the company’s documentation culture was even more developed than I previously experienced. Simply put, I was impressed with the amount of documentation about pretty much everything: onboarding, projects, internal tools, internal processes.

It made me think about times in my career when finding a piece of information was an epic journey—documentation was almost non-existent or out-of-date. I had to reach out to people because the knowledge was in their head.

In the tech industry, there’s a problem with internal documentation. We’ve all had this feeling of being eager to get things done when joining a new company. Instead, we’ve found ourselves in a poorly documented codebase, or with unclear instructions on how to set up a project. As engineers move from team to team, project to project, or leave the company, good documentation becomes essential.

Sharing technical context internally is one of the most important and powerful things we can do. It creates space for discussions and problem solving, and lets experts share knowledge to everyone’s benefit.

But internal documentation can be hard: It's a time- and- money-consuming process. Documentation is often spread across different platforms rather than consolidated for easy use. And it doesn’t always get prioritized—shipping code is way more exciting and fulfilling than documenting a process. Navigating new codebases has become one of the biggest challenges in software development.

At Shopify, we’ve tried to overcome that issue by standardizing the way we write, avoiding disparate sources of truth, and creating a writing culture. We treat documentation as a product, and put care into creating comprehensive documentation. The most effective way to encourage others to write is to actually write. It can be as easy as opening a pull request and getting it approved. The benefits of a strong documentation culture are instantaneous.

In a remote-work environment, writing is our default method of sharing knowledge. Everybody has access to everything: All the repos, the dashboards, the wikis, the Google Drive documents. Documentation increases discoverability of content, and by making all information searchable, dev literacy improves. And when dev literacy improves, so does productivity.

## Where We Publish Internal Documentation

Internal documentation is also an operational tool that allows the company to scale. It’s important to be very intentional on where we publish the technical documents. It has to be attractive for engineers and easy to find and navigate. To make a big impact, documentation has to stay fresh and relevant.

Let’s walk through the three main tools we are using to share technical knowledge internally.

### **Markdown Documentation**

We have stand-alone websites written in Markdown and powered by [Docusaurus](https://docusaurus.io/), an open-source templating engine. These are easy to maintain and update since the content lives close to the code and has to go through peer review before being published. The big plus with microsites is that more people will write them because they can do it directly from the integrated development environment and repo they’re familiar with. The advantage of being written in Markdown is that the documentation can easily be transferred to another system. We’re also planning to migrate all the Docusaurus content to the Vault to further increase its discoverability.

### **Github**

Why doesn’t X work? What doesn’t work? What’s next in the pipeline for your team? What’s the status of this bug? Answers to these questions can be found in our Github internal. We use it as a bug/ticket/issue tracker and everybody has access to it. The Github search capability also comes handy when looking for specific bugs or features.

Engineers also use Github and README’s for information on how they should start working on X project. What environment setup is needed? How do we test this code base? Finally, being able to search code in Github allows us to quickly understand how something is implemented or find examples of using API or code.

### **The Vault**

Our custom internal platform. A gold mine. The Vault was created way back in 2012 when Shopify had a little over 150 employees. For seven years, the Vault served us well as a way to share our contextual information with one another. As the company grew to thousands of employees, features needed updating and enhancing to meet growing internal information-sharing needs. The Vault is our main way to “Default to open”, and is the most important source of truth for Shopify context. It’s powered by Github and Markdown editing allows us to create more flexible and customizable page content. Tracking page history is easier thanks to Git versioning.

## Good Documentation = Productivity

Strong documentation culture benefits the organization in several ways, but first and foremost it boosts engineers’ productivity. If you have few places to store common information, there’s no need to spend several hours finding the right person to talk to. Context and answers are a few clicks away.

Having a lot, if not everything, documented prevents teams from having to reinvent the wheel. Engineers go through a lot of success and failures, and can document their processes for other teams’ benefits. Instead of spending enormous amounts of time trying to figure out how to consume an internal API, how to use a specific gem or package, or how to mock dependencies when running unit tests, an engineer can efficiently read up on examples and proven solutions.

Onboarding or to a new team is faster. Not everybody is comfortable asking for help or guidance, especially if they’re a new employee or have joined a new team. With good documentation, everyone can read technical information at their own pace. Newcomers onboard the same way, using the same resources. It reinforces feelings of belonging and unity by creating a common experience from the beginning. Everyone knows what they are building and why. New team members are also encouraged to fill missing gaps in documentation, which is a good way to engage with projects right away, and have a clear understanding of aims.

Written documentation also creates a more inclusive culture. Engineers whose primary language is not English may find it difficult to assimilate spoken information. In general, widely accessible tools and documents create a more inclusive environment because everyone has a say and can feel part of something. Internal information is not just the focus of one dedicated person or particular team, but everyone’s responsibility.

Well-written, clear, and accessible internal documentation encourages collaboration between teams. It’s easier to get started with repositories. Documentation breaks barriers between teams. Engineers will be more inclined to jump in different projects if they feel that documentation is treated with good care. Before using a package or an API, documentation is often the first thing we look for when evaluating options.

Finally, good documentation enables [asynchronous communication](https://shopify.engineering/asynchronous-communication-shopify-engineering), which is especially crucial for [cross time-zone teams](https://shopify.engineering/a-software-engineers-guide-to-working-across-time-zones): everybody has access to the same information and can contribute to it anytime, anywhere, and on any device.

## Three Ingredients of Good Technical Documentation

Documentation quality and scalability is difficult to track. It doesn’t have unit tests asserting that what is published respects a high standard. The tooling to make documentation simple and readable is as important as the technical content itself if you want engineers to write and champion documentation.

Here are three things to consider when starting to write internal documentation, or contributing to established processes:

1. Documentation tool(s) should be integrated into your CI/CD (continuous integration/continuous deployment) pipeline. Devs will be more inclined to write if their content is processed the same way as their code. They’ll be able to see the results right after the CI/CD pipeline is complete and it’ll encourage them to write more and inspire others to do so. Writing docs follows the same principle as writing code.
2. Doc should live somewhere familiar to the engineers. In our case, most documentation is hosted on Github, and chances are that engineers are interacting with GitHub or Gitlab on a daily basis. Searchable capabilities make it easier to discover and access the information, which is a no-brainer when you need to make great decisions quickly.
3. Documentation should be pleasant to look at and user friendly. Readers need to find comfort when going through documentation. The experience should be painless and the information should be presented in an engaging way. Retaining information will increase dev literacy and ultimately better equip your engineering teams to navigate the complex world of software development.

## Write, Write, Write

Lastly, building a support system that encourages and empowers everyone to write is essential to make internal documentation an exceptional product. Writing is a skill that we—engineers—can continually improve, like we do with our code. Here are a few tips that helped me keep up my writing:

- Writing documentation about a big project or feature can be intimidating. Practicing daily on smaller features or even flushing out a bug ticket could help you to get accustomed to the mental and physical concept of technical writing.
- Don’t do it alone: Within my team, we regularly pair to write feature tickets or create diagrams. Having multiple people writing together makes the documentation richer. It also gives different perspectives and contexts, which makes the outcome more powerful.
- Always keep context in mind: You don’t have to give the entire backstory, just the missing pieces so your documentation is clear. Always think about where your documentation will be published and who your audience is. Are you writing for your team only? Is it a highly visible piece of information that will be read across the company? The tone and the amount of information you write may vary depending on your readers.

These are the building blocks from which to build a culture that fosters developer literacy and productivity. Whether you work for a startup or a well-established company, making documents easy to read and find encourages engineers to write and spread knowledge. Documentation lasts forever and is essential to building a 100-year company.

**Michael** is a Senior Developer helping large DTC merchants manage their fleet of first-party POS devices. He lives near Montreal with his partner and two kids, and enjoys soccer, cooking and solving problems with code.

We all get shit done, ship fast, and learn. We operate on low process and high trust, and trade on impact. You have to care deeply about what you’re doing, and commit to continuously developing your craft, to keep pace here. If you’re seeking hypergrowth, can solve complex problems, and can thrive on change (and a bit of chaos), you’ve found the right place. Visit our [Engineering career page](http://www.shopify.com/careers/specialties/engineering?itcat=EngBlog&itterm=Post&shpxid=a23be660-316D-443F-BF0D-F3BE387A359C) to find your role.


