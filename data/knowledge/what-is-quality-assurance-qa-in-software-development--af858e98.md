---
title: "What is quality assurance (QA) in software development?"
notion_id: af858e98-fb25-4abb-9f21-1d05b973c459
notion_url: https://app.notion.com/p/What-is-quality-assurance-QA-in-software-development-af858e98fb254abb9f211d05b973c459
last_edited: 2022-12-19T14:07:00.000Z
source_url: https://blog.logrocket.com/product-management/what-is-quality-assurance-qa-software-development/
tags: ["English", "Programming", "Testing", "Project Management", "Product Management", "Line/People/Team Management", "Article", "LogRocket Blog"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Buggy, underperforming solutions rarely delight users or drive business outcomes. High-quality products are what we strive for.

In this guide, we’ll explain what quality assurance (QA) means, describe why it is distinct from testing, and outline the the role QA specialists play in the software development lifecycle.

## Table of contents

- [Common activities done by a QA specialist](https://blog.logrocket.com/product-management/what-is-quality-assurance-qa-software-development/#common-activities-done-by-a-qa-specialist)
- [Summary](https://blog.logrocket.com/product-management/what-is-quality-assurance-qa-software-development/#summary)

## What is quality assurance (QA)?

The process of ensuring product quality is called quality assurance (QA).

Some make the mistake of simplifying quality assurance into just testing, but there’s so much more to it. The actual testing phase is a drop in the ocean within the quality assurance discipline.

## How is quality assurance different from testing?

Quality assurance isn’t the same thing as testing — a true QA specialist is more than just a tester. Testing is a small part of a broader discipline:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Simplifying a bit:

- **Quality assurance (QA)** is a set of processes and practices that aim to ensure high product quality. It includes proactive and reactive activities
- **Quality control (QC)** is a part of quality assurance, and it’s about ensuring the delivered work item meets desired quality standards. It’s done post-factum
- **Testing** is an activity performed as a part of quality control

Ultimately, a QA is responsible for ensuring the quality of the end product, which requires much more than manual testing.

## Role of QA in the software development lifecycle

The most common misconception about quality assurance is that it’s an activity done at the end of the [development lifecycle](https://blog.logrocket.com/product-management/product-managers-role-each-product-lifecycle-stage/). This misunderstanding originates from a waterfall development model:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

However, quality assurance is not left at the end, even in the waterfall model. Although feature testing is one of the last steps, it doesn’t mean the whole quality assurance is non-existent until then. It’s prevalent in every single phase:

### Analysis

Quality assurance includes ensuring the analysis group is built adequately and consists of the right people from the onset. QA specialists do whatever they can to ensure no crucial element is missing.

The more problems that arise during the analysis phase the more unaddressed risks and change requests you’re likely to encounter, which could cause quality to deteriorate.

### Requirement specification

High-quality requirements lead to high-quality solutions. QA managers study requirements religiously, looking for any gaps, contradictions, and unclarities there might be. Finding any logical issues during the requirements phase is much cheaper than finding them in the production code.

### Design

QA specialists play a critical role when the exact details are defined and designed. They ensure everything is covered and the design is free of errors.

They ask themselves meticulous questions such as:

- Does it match the requirements?
- Are all necessary errors stated planned?
- Does it cover edge cases?
- Are the commas in the right place?

The more error-free the design, the more error-free the final product.

### Development

During the development phase, QA specialists mostly prepare for the testing and integration phase. It includes preparing test plans and test cases for further tests. This preparation helps them to perform more robust and efficient testing down the road. Sometimes, QAs also write automation tests alongside the development team.

### Testing

Testing is the execution phase for QAs, where they take the developed increment and execute the test cases they prepared earlier. Depending on the scope of testing and complexity of the increment, they can test functionalities at a high level or triple-check every single input field.

### Deployment

Quality assurance doesn’t end with deployment. After a product or feature is launched, QAs usually double-check to make sure what worked for staging also works for production. They also monitor product quality metrics and ensure the product doesn’t break after new updates.

Regardless of whether your company uses a waterfall model, V-model, or [agile software development framework](https://blog.logrocket.com/product-management/agile-product-management-what-does-it-mean/) such as [scrum](https://blog.logrocket.com/product-management/what-is-scrumban-methodology-how-to-implement/#what-is-scrum) or [kanban](https://blog.logrocket.com/product-management/what-is-scrumban-methodology-how-to-implement/#what-is-kanban), the overall scope of QA activities is rather similar.

No matter what model you adopt, you will always have the analysis, design, development, testing, and deployment stages. The main difference is when these phases take place and how big the pieces of work are.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## QA specialist responsibilities

To provide a snapshot of what a QA specialist does, let’s dive deeper into their day-to-day activities. The exact scope of activities done by a QA depends heavily on the product itself and the company’s approach to ensuring quality:

### Quality assurance

Although quality assurance also includes quality control and testing, let’s focus on activities preventing issues:

### Establishing proper processes

Quality should be built into the entire development process. How problems are analyzed, requirements are documented, and issues are communicated significantly impacts the overall quality.

QA experts ensure that the processes and practices within the team lead to high-quality products.

### Root cause analysis

Whenever there’s an incident, a QA specialist, usually leads or heavily contributes to the post-mortem process. They ensure that the teams learn and improve from each mistake, not just fix and forget.

### Ensuring documentation quality

Proper documentation helps you prevent errors and troubleshoot problems when they appear. In the long run, it also helps you speed up development time.

Although documentation is the whole team’s responsibility, documentation review is often a part of quality assurance practices.

### Performing audits

QAs often ensure the required audits are done or conduct the audits themselves.

For example, an app might need to fulfill specific GDPR requirements, or a card payment processor might require regular external audits. QAs might perform these (pre-) audits and often play a pivotal role in cooperating with external auditors.

### Training

Quality assurance is the responsibility of the whole team. After all, everyone’s actions impact the final outcome.

QA specialists help the team understand best quality practices, explain how to write better documentation, and sometimes teach the team how to write automated tests independently.

### Quality control

Quality control is all about ensuring the product meets the agreed-upon quality standards. The various activities that contribute to keeping quality in check include the following:

### Non-functional requirements monitoring

Non-functional requirements, such as scalability, reliability, performance, and security, should meet some of the quality benchmarks

Whether, in a continuous (e.g., resource usage monitoring) or ad-hoc (e.g., load tests) manner, QAs keep non-functional quality in check with additional support from other specialists like DevOps engineers.

### Automation

As the product scales, it becomes inviable to rely entirely on humans. At some point, it’s essential to automate core test scenarios.

Quality experts prepare and configure test environments and write tests that execute and report specific test cases.

In the long run, automating tasks is more efficient and prevents the team from having to manually test everything every time.

Some organizations rely on QA specialists to write and maintain automated tests, while some decide to have dedicated quality assurance automation (QAA) specialists.

### Maintaining test plans

The product changes and evolves all the time. Thus, the scope of testing should also evolve continuously.

The feature that was critical a year ago might be irrelevant or even depreciated today.

QAs ensure that the scope of all testing (manual, automated, documentation, non-functional, and so on) is adequate to address the current needs of the product.

Ultimately, it’s impossible to test everything, so QAs ensure the team picks the right battles.

### Testing

Testing is a core element of quality control.

There are different types of tests a QA can perform. The most common ones are:

### Smoke tests

A smoke test is the quickest and most rudimentary type of test.

The objective is to test the app’s main flows and functionalities superficially. It might be scripted or done based on a gut feeling. It has two purposes:

- To decide if the product is stable enough for more robust testing; there’s no point in starting scenario tests if the app crashes after launch
- As a sanity check to detect if the app is working. If the app experiences prolonged stability issues, QAs might do a smoke test multiple times in a day to ensure the app is still live

### Scenario tests

In simple words, scenario testing includes designing and executing particular test cases. They have step-by-step instructions on what to test and how to test it:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[Synapse QA](https://synapse-qa.com/2020/03/22/test-case-authoring-tips/)

The reasons to write test cases include:

- Use as input for further test automation
- Ensuring that all core elements are tested regularly
- Combine them into specific test suites for various purposes such as weekly regression tests

The most significant disadvantage is how time-consuming testing can be when writing and maintaining an ever-growing repository of test cases. Complex products can have as many as a few thousand test cases.

### Exploratory tests

These are unscripted, skill-based tests that combine learning, test design, and test execution. QA tests the app without any specific scenario and user flows, often just clicking through the whole app in somewhat random order.

Since the tester has to design and execute tests on the go, past experiences and domain knowledge are extremely valuable.

Exploratory tests serve various purposes. They:

- Allow QAs to learn the product quickly
- Help prepare and prioritize scenario tests
- Can be used as a standalone quality control process

### Regression tests

The fact that something worked once doesn’t mean it will work forever, especially when it comes to developing software.

By changing one application element, a developer could accidentally change component settings used in different parts of the app.

Breaking things that worked in the past is quite common. We call this phenomenon “regression.”

Regression testing means double-checking if things that should work still work. Usually, regression testing is done before releasing a new version of the product, with QAs checking to see if the most critical user flows work as expected.

That means if there’s a release twice a week, a key component, such as the login process, should also be checked twice a week. Hopefully, this testing is done by an automat, not a poor human being.

## Summary

Contrary to popular belief, quality assurance is present at every stage of the [product development lifecycle](https://blog.logrocket.com/product-management/what-is-product-development-lifecycle-stages-examples/), not just the final testing phase. Quality experts ensure the analysis phase is done correctly, review the requirements and designs, and monitor how the app behaves in the product environment.

Additionally, QAs also perform various other activities, such as process improvement, training, documentation review, audits, automation, etc.

Quality is the responsibility of the whole team. However, with the meticulous planning and complex principles that go into ensuring quality, every team could benefit from gaining knowledge from a dedicated QA specialist, or even a few of them.

## [LogRocket](https://lp.logrocket.com/blg/pm-signup) generates product insights that lead to meaningful action

[LogRocket](https://lp.logrocket.com/blg/pm-signup) identifies friction points in the user experience so you can make informed decisions about product and design changes that must happen to hit your goals.

With LogRocket, you can understand the scope of the issues affecting your product and prioritize the changes that need to be made. LogRocket simplifies workflows by allowing Engineering and Design teams to work from the same data as you, eliminating any confusion about what needs to be done.

Get your teams on the same page — try [LogRocket](https://lp.logrocket.com/blg/pm-signup) today.
