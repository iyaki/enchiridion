---
title: "Address Tech Debt with SRE Panel Recap"
notion_id: 85099d9c-4ac7-4751-9836-6522ded671c1
notion_url: https://app.notion.com/p/Address-Tech-Debt-with-SRE-Panel-Recap-85099d9c4ac7475198366522ded671c1
last_edited: 2026-09-21T17:35:00.000Z
source_url: https://thoughtbot.com/blog/address-tech-debt-with-sre-panel-recap
tags: ["English", "Site Reliability Engineering", "Article", "thoughtbot Blog"]
---
[https://thoughtbot.com/blog/address-tech-debt-with-sre-panel-recap](https://thoughtbot.com/blog/address-tech-debt-with-sre-panel-recap)

Technical debt can build up quickly for startups as companies prioritize new features over best-practice code. But when it comes to scaling, organizations need to find the right balance of refactoring and feature development. That’s where Site Reliability Engineering (SRE) and service level objectives (SLOs) come to the rescue. We brought together an expert panel of tech leaders to discuss architectural trade-offs related to service level performance and to share their go-to cultural practices for reducing tech debt through SRE. You can view the full discussion on our youtube channel.

Read on for the key takeaways from the discussion.

### What is SRE?

- SRE is a metrics-driven subset of the very large umbrella of DevOps and DevSecOps.
- All SRE work is driven by the metrics that you establish.
- The term SRE was coined by Google, which has a great series of books and workbooks on SRE.
- SRE gives you a prescriptive shared language about how it is going to work.

### How does SRE help you address technical debt?

- Issues related to uptime and deployment are frequently signs that there is technical debt that is actively hurting your app.
- The concept of technical debt is fairly abstract, so being able to share the resulting SRE metrics and real-world effects can make tech debt more concrete for stakeholders.
- SRE is focused on measuring things that affect users, and most stakeholders speak that language.

### As a Site Reliability Engineer, how do you investigate the tech debt?

- The very first phase of SRE implementation is investigation, communication, and just sitting down with the teams, trying to find those things that people don’t want to talk about.
- One favorite ice-breaker is “What’s the worst thing here?” because everyone loves to answer that one.
- The real tech debts are the ones that are not written down: they are in everyone’s heads. You know if you touch this bit of the code it’s going to break everything.

### How can you get developers excited about SRE?

- A lot of tech debt is reflective of a culture that doesn’t listen, so SRE can help improve the developer culture.
- An SRE is like a tech therapist: you just listen and pay attention, giving those people the focus they’ve been lacking, and that will help you find what is broken.
- SRE can help reduce on-call alerts and burnout and make sure issues go to the right person.
- It’s also about improving the developer experience, ensuring that it’s an enjoyable, not a stressful experience.
- Getting the initial pass done is an effort, but when engineers see the charts, graphs and visualizations, they love it - it’s addictive and they start to do it on their own.
- Dashboards and metrics tend to gamify SRE for devs, and there is a natural push to get the numbers as close as they can to 100%

### When people are concerned about failed deployments, how do you help develop their confidence?

- It is very hard for developers to figure out abstractly how close to the edge their system is, but by starting to measure the things that matter to your customers, you can start to see the impact of refactoring or new features.
- When you have a culture that’s built on a shared understanding of Service Level Indicators (SLI) and error budgets, you start to build deployments that inherently try to recover on their own.
- SRE offers a feeling of safety once you start using it, making it safer for developers to push their changes with more confidence.

### How do you make systems more observable?

- Specifically with tools like Lightstep, Jaeger and OpenTelemetry, which is an open-source effort to standardize the big three of observability: logs, metrics and traces.
- With observability, you can trace that exact transaction that broke you in half a second: it’s magic.
- Observability needs standardization, and there’s not a global culture of standardization, so that’s probably going to be our whole industry’s goal for the next 5 or 6 years

### How would you advise someone just getting started on here to begin their SRE journey?

- Get a solid foundation in understanding the cloud service you are working in. e.g. for AWS, use the AWS curriculum to get a certificate (there is a track just for developers).
- The Google SRE book is a real touchstone for thinking about systems.
- Be flexible: as you move in your career, be able to think about systems as generics and thrive in being able to handle unknowns.
- WWC has a cloud and DevOps series.
- If you are new to SRE but your organization already does SRE then go and see the dashboard for your service and look at the SLIs, as SRE is much easier to understand through experience.

### What trade-offs do you have to consider with each platform if you want to implement SRE?

- Tooling is hard and is the most accurate indicator of culture.
- To be successful, an SRE needs to adapt, to flow like water and keep it open tooling-wise, like a swiss army knife.
- Cloud agnostic tooling is great.
- Look at open source, and if big cloud providers have been forced to support a tool, like Kubernetes, that’s a good sign.
- Use the big guardrails of it’s open source, it’s flexible, and people have experience with it.

### How would you talk to senior leaders who haven’t heard of SRE?

- SRE measures the things users care about.
- Observability is a hard problem, but getting a request per second/response time graph that shows how your site has been performing for the past week is quick.
- You can bring that to leadership and show them: this is the real effect. I can show you that the way things are set up now, every time you deploy the site is slow for ten seconds, and that matters to people.
- Buying that little bit of trust means they will listen to what you have to say.
- Bring up the SLOs, and the specific goals for the website and show how SRE can improve them.
- Tie metrics to business outcomes, that’s what usually gets the business happy. E.g. how many of your customers are making it through to checkout, and where are you losing them?

### Which aspects of SRE talk to design or user experience?

- In SRE, you have to understand what your users’ experience is of your site and your business, and then identify metrics to measure that experience.
- Designers tend to understand SRE as they understand the user impact more than backend devs.
- SRE can show a designer what the cost is of a design to page load time using real data.
- At thoughtbot, we’ve found taking a design-driven approach to SRE, as with all things, is more successful.

### When should companies start considering SRE?

- As soon as you have real user traffic, you have all of the problems that SRE solves
- It makes sense to start early, as it could be difficult to implement SRE if you haven’t thought about observability from the beginning
- Implement SRE when you’ve agreed on SLOs and you have an error budget and as soon as you have actual user traffic.

### Which SRE tools are you most excited about.

- Lightstep
- Lens for Kubernetes
- AWS CDK
- CDK for Terraform
- OpenTelemetry

## Panelist Shoutouts

### Camille Clayton

Engineering Manager. Kubernetes Infrastructure at Ticketmaster. Director of Women Who Code DC Cloud DevOps

- Ticketmaster Livenation has a hiring push - feel free to reach out to me on Twitter.
- Side project: remoteworkcalc.com
- Join Women Who Code DC Meetups and Slack

### Emmanuel Apau

CEO @ Mechanicode.io, Former United States Digital Service Expert, Certified Kubernetes Admin. AWS & Azure Certified DevSecOps specialist with 12 years of experience developing innovative automation solutions using DevSecOps & Site reliability best practices for clients. Experience in both the public and private sectors, providing services that engage Agile best practices, scalable cloud architectures, and modern continuous integration & deployment standards.

- mechanicode.io is hiring - Cloud Engineers and DevOps
- My SRE course on Udacity
- Black Code Collective - a community for Black software developers to grow our skills, share knowledge and help each other progress through our careers.

### Joe Ferris

CTO at thoughtbot

- thoughtbot can help with scalability and reliability problems.
- If you have a team interested in adopting SRE and you would like a seasoned hand to get you through that process to that first beautiful graph, we would be very excited to help you get there, reach out here to schedule an introductory call.

Thank you so much to our expert panel for their insight into how SRE and help you tackle tech debt and improve company culture and developer experience.
