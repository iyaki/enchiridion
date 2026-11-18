---
title: "A Guide to Design Thinking and Continuous Delivery from an Ex-Principal Engineer"
notion_id: cfb2b024-76f7-45d1-9e66-1cb95981372f
notion_url: https://app.notion.com/p/A-Guide-to-Design-Thinking-and-Continuous-Delivery-from-an-Ex-Principal-Engineer-cfb2b02476f745d19e661cb95981372f
last_edited: 2023-01-25T19:29:00.000Z
source_url: https://hackernoon.com/a-guide-to-design-thinking-and-continuous-delivery-from-an-ex-principal-engineer
tags: ["English", "Project Management", "Agile", "Product Management", "Programming", "Article", "Hackernoon"]
---
## How do the big tech companies and unicorn startups succeed in agile product development while maintaining code quality.

### TL;DR

- 🎨 We’ll discuss what is design thinking.
- 🤖 We’ll discuss what is continuous delivery.
- 🏎 We’ll uncover how to utilize design thinking and continuous delivery to build products with agility and quality.
- 📚 Tons of learning materials!

Design thinking and continuous delivery are amazing approaches to product innovation. They are not widely discussed in the agile context, but the synergy between them is crucial to build products that people love.

### What is Design Thinking

Design thinking is about problem-solving. Companies like [Airbnb, Netflix, and Uber](https://online.hbs.edu/blog/post/design-thinking-examples?ref=hackernoon.com) used design thinking to form deep understanding of their users and built features to address the user needs.

> _“Design thinking is a human-centered approach to innovation that draws from the designer’s toolkit to integrate the needs of people, the possibilities of technology, and the requirements for business success.” — _[_Tim Brown, Executive Chair of Ideo_](https://designthinking.ideo.com/?ref=hackernoon.com)

I like it a lot because it’s inclusive in the innovation process. It starts with understanding the users. The goal is to solve a problem with a solution that’s desirable from the human point of view and technically feasible while capturing market opportunities.

### How to Practice Design Thinking

Design thinking follows the flow of **understand**, **explore**, and **materialize**. There are six phases:

- **Empathize**: Conduct research and develop an understanding of your users.
- **Define**: Combine the research and identify where problems exist.
- **Ideate**: Brainstorm wild and creative ideas to address the problems and the business goals.
- **Prototype**: Build lean solutions from a subset of the ideas.
- **Test**: Return to your users and collect feedback.
- **Implement**: Deliver your solution incrementally to your users.

As you can see, design thinking iterates through [product discovery](https://www.producttalk.org/2018/08/effective-product-discovery/?ref=hackernoon.com), [testing](https://hbr.org/1975/05/when-where-and-how-to-test-market?ref=hackernoon.com), and [building](https://www.mojotech.com/blog/building-the-right-product-vs-building-the-product-right/?ref=hackernoon.com). You have the flexibility to run it as lean as in just one week with [the design sprint](https://www.thesprintbook.com/the-design-sprint?ref=hackernoon.com), or you can iterate and repeat through the phases until you find your [product market fit](https://pmarchive.com/guide_to_startups_part4.html?ref=hackernoon.com).

You can read more about the design thinking process:

Now, in engineering, what we want to achieve is to have the infrastructure and workflow to enable such rapid product iterations so that multiple teams can prototype, test, and implement in your product **with speed and at scale**. This is where continuous delivery can help us.

### What is Continuous Delivery

Continuous delivery puts **working software** first. It’s a practice to release software [quickly, safely, and sustainably](https://cloud.google.com/architecture/devops/devops-tech-continuous-delivery?ref=hackernoon.com).

> _“Continuous delivery (CD) is a software engineering approach in which teams produce software in short cycles, ensuring that the software can be reliably released at any time and, when releasing the software, without doing so manually. It aims at building, testing, and releasing software with greater speed and frequency. The approach helps reduce the cost, time, and risk of delivering changes by allowing for more incremental updates to applications in production.” — _[_Wikipedia_](https://en.wikipedia.org/wiki/Continuous_delivery?ref=hackernoon.com)

There are several benefits by integrating continuous delivery. From the [DORA State of DevOps research program](https://www.devops-research.com/research.html?ref=hackernoon.com), which summarizes a 7 years of research on high performance technical delivery and organizational outcomes, we can see a clear benefit in performance and culture. Here’s the screenshot:

The idea is to have an automatic process that builds, tests, analyzes, integrates, and deploys continuously throughout your product cycle. In the case of any error occurring in the process, the process will stop your product changes being integrated in production so that it won’t impact your users. It leads to a greater level of [psychological safety](https://rework.withgoogle.com/print/guides/5721312655835136/?ref=hackernoon.com) that enables us to test our prototypes and collect feedback and data without worrying about breaking the product.

### How to Practice Continuous Delivery

According to the [Accelerate State of DevOps 2021](https://services.google.com/fh/files/misc/state-of-devops-2021.pdf?ref=hackernoon.com), there are [several capability factors](https://cloud.google.com/architecture/devops?ref=hackernoon.com) that serve as indications to software delivery and organizational performance. In this article, I want to highlight a few technical capabilities in the context of agile product development:

- Trunk-based development: developers work in small batches and merge into the trunk frequently.
- Continuous testing: QAs tests early and frequently throughout the delivery cycle.
- Continuous integration: triggers a series of automated tests that provide early feedback.
- Observability practices: monitor your system throughout the delivery cycle.
- Deployment automation: decreases lead time and reduces deployment errors.

You can read more about the concepts, best practices, and KPIs of continuous delivery:

Now that we have discussed both design thinking and continuous delivery, we are ready to see how they can work together to achieve higher speed and code quality.

### Put Them Together in Product Development

Continuous delivery comes into the picture in the **prototype, test, and implement** phase.

Continuous delivery play a part in the prototype, test, and implement phase of design thinking.

Let’s look at the phases separately.

### During Prototype & Test Phase

The two most important focuses during the phases are **learning** and **speed**. You want to be able to put your prototype in front of your users as quickly as you can to test your ideas and hypotheses so you can design the next iterations based on your learning.

To prioritize learning and speed, we can design the continuous delivery process like this:

We are keeping most of the best practices with a few modifications:

**Dirty Pull Requests**

✅ Dos

- Put together the prototype as fast as possible.
- Modularize the code change.
- Mark your prototype module with an obvious prefix. Like **UNSAFE_EXPERIMENT_**.
- Writing tests is not necessary.
- Focus on functionality in code review.

❌ Don’ts

- Don’t code to perfection.
- Don’t block releases with lengthy code reviews.
- Don’t break the build.

**Regression Testing & Continuous Integration**

✅ Dos

- Guard the main user journey. For example: Checkout flow in an E-Commerce site.
- Run only the baseline regression tests in favor of speed.
- Notify there are warnings in the continuous delivery process.

❌ Don’ts

- Don’t let the prototypes that failed the tests go into production.
- Don’t lower the test coverage.

**Analytics and Feature Toggles**

When your prototypes are designed for [quantitative research](https://www.nngroup.com/articles/quant-vs-qual/?ref=hackernoon.com), it requires code changes in the product so that you can analyze the data on a larger scale. This is where analytics and [feature toggles](https://www.martinfowler.com/articles/feature-toggles.html?ref=hackernoon.com) are useful.

Feature toggles gives you the flexibility to target the user segments for your tests and analytics provides you aggregated data to track your KPIs.

The purpose of feature toggles in these phases is to collect data and learn from the intended user segments. Make sure you don’t use toggles to hide prototypes or features from your users. It defeats the purpose of learning.

### During Implement Phase

When you are in the **implement** phase, it means that you have tested the prototypes and you are officially going with one of your ideas. Congratulations!

Unlike in the prototype and test phase, we want to prioritize the best practices of continuous delivery over speed:

**Stacked Pull Requests**

✅ Dos

- Use a [series of pull requests](https://www.michaelagreiler.com/stacked-pull-requests/?ref=hackernoon.com) instead of a large pull request to implement your product or feature.
- Write comprehensive tests.

❌ Don’ts

- Don’t commit large code changes. It makes code review harder and slower.
- Don’t prolong code review as a reviewer. Act (accept, reject, or resign) ASAP.
- Don’t form an elite group that does gatekeeping and blocks the code review.

You can find more about the best practices for pull request authors and reviewers:

**Continuous Testing and Continuous Integration**

✅ Dos

- Set a high test coverage threshold.
- Write end-to-end tests to guard major user journeys.

❌ Don’ts

- Don’t lower code coverage.
- Don’t skip steps in continuous delivery automation.
- Don’t settle for slow automation. Lead time is crucial for product delivery and your business.

### Final Thoughts

Design thinking provides us a human-centric framework to iterate on product innovation, whereas continuous delivery creates a effective and predictable baseline to protect the code quality.

Code quality often lose its priority to lead time in many companies. By following design thinking, we are able to ideate and test the prototypes quickly to learn about the market and users. By integrating continuous delivery, we have a technical capability to collaborate and test on a large scale without breaking our products.

### References

- Wikipedia: [Programming style](https://en.wikipedia.org/wiki/Programming_style?ref=hackernoon.com)
