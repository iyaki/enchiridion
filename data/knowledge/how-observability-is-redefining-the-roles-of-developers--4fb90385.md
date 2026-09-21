---
title: "How observability is redefining the roles of developers"
notion_id: 4fb90385-52e2-41db-a162-44e165c6090a
notion_url: https://app.notion.com/p/How-observability-is-redefining-the-roles-of-developers-4fb9038552e241dba16244e165c6090a
last_edited: 2023-01-25T18:32:00.000Z
source_url: https://stackoverflow.blog/2022/07/18/how-observability-is-redefining-the-roles-of-developers/
tags: ["Stack Overflow Blog", "English", "DevOps", "Site Reliability Engineering", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Your DevOps team might already be singing the praises of their observability platforms. But developers can benefit from them, too.

You’re tracking a bug through production. You look through the logs. The one thing you need isn’t there… Dead end. A few years ago, I was tracking a production issue with a server that triggered a request to a database read due to [cache misses](https://en.wikipedia.org/wiki/CPU_cache#Cache_miss). This skyrocketed our cost due to high read volume. Unfortunately, there was no way to know what triggered this since there was no logging on cache misses.

The flip side of this is that adding logging here would have helped track the cache miss, but would have skyrocketed our logging ingestion costs. Logging storage is amazingly expensive. This is something I run into a lot; a developer who needs to add a log to production has to go through PR, approval, merge, [CI/CD](https://stackoverflow.blog/2021/12/20/fulfilling-the-promise-of-ci-cd/) etc., only to find out that another log is needed. In our team, we nicknamed this the CI/CD cycle of death. We love the basic process, but it wasn’t designed for use as a “poor man’s debugger.”

It’s a pain we all feel. Developer observability is the name for a family of tools designed to solve this pain. Existing [observability](https://stackoverflow.blog/2021/09/08/observability-is-key-to-the-future-of-software-and-your-devops-career/) tools were designed with DevOps in mind, developer observability shifts observability left into the software development lifecycle. In this article, I’ll explain what they are, how they are built, what they offer, and how to pick the one that fits your needs (without endorsing specific tools).

## A complex new world

Today’s software world is pretty different from the one we had when I started programming. When I was a young programmer, monitoring your production server meant walking over and kicking the hardware to hear the hard drive spin up. “Yep, it’s working.”

This is obviously no longer tenable. Our modern scale just doesn’t allow for it. We now have the DevOps team guarding production—that’s a good thing. Since we adopted DevOps, hired SREs, implemented CI/CD, etc., production has become far more stable and startups have scaled far more effectively than ever before.

But production bugs are still here. The progress we’ve made as an industry has a significant downside as those production bugs are MUCH harder to track than they were in the past. The scale makes it hard. Cloud, container orchestration, [serverless](https://stackoverflow.blog/2020/05/18/you-want-efficient-application-scaling-go-serverless/) etc., let us deploy hundreds or thousands of instances immediately. This enables responsiveness, reliability, and flexibility like never before but also presents concurrency problems like never before. Data corruption at scale due to a bug or misconfiguration are rampant. Only a portion of our production is accessible to us—this makes debugging it even harder.

Very few developers use observability and monitoring tools. Most observability tools vendors build their products with DevOps in mind and don’t target engineers. It makes sense: DevOps handles production. But the new generation of tools presents an option: what if you could debug issues right in production?

What if you could do that with no risk?

## What’s developer observability?

Developer observability is a new pillar of observability adapted for the needs of developers. Unlike typical observability solutions, it’s aimed directly at developers and not at DevOps. As such, it provides a direct connection between the source code and the observable production.

Developer observability includes the two following distinctive properties:

- Based on user requests
- Works with source code

Typical observability tools place instrumentation throughout the application—e.g. on every web service entry point and often deeper. These tools use the instrumentation to sample data and send information. As such they push observability data to their management server.

Developer observability tools do nothing by default. A developer needs to explicitly add observability to specific source file line (or lines). It works in a “pull” mode.

A good analogy would be that developer observability is like a debugger whereas current observability tools are like a profiler. When you run with a debugger, it doesn’t do much until you add breakpoints to extract information. A profiler constantly gets information while running. Both are very useful and both serve different use cases.

## Logging on demand

[Logging in production](https://lightrun.com/best-practices/logging-best-practices-mdc-ingestion-and-scale/) can be invaluable in tracking thread related problems. Because of the scale of production, some concurrency issues only show up there.

Unfortunately extensive production logging isn’t something we can realistically do for most use cases. If we add a log in every method entry/exit, our logs will blow up. They will become unreadable, skyrocket our storage costs, and slow down the performance of the server. Adding a few logs to a specific server can make a big difference in the debugging process without a noticeable impact on the logs overall.

This is where developer observability tools can step in. A typical feature in these tools is the ability to add a new log into production without changing the code. A developer could add a log on on the method entry and exit points directly in their IDE. Since loggers normally include the details of the thread, we could inspect the log to see potential race conditions.

Under the hood, the developer observability tools rely on an agent service installed on your production server. It adds the log for you as if you wrote it in the code yourself. To keep production segregated and safe, these tools communicate externally to a management server. Your IDE connects directly to that server and has no direct access to production. Since production is involved, these tools include safety features such as sandboxing to prevent a method invoked from a log from changing state. E.g. I can add a log such as: `“User {user.getUserId()} reached myMethod”`.

Some tools verify that the method invocation is indeed read-only.

We can then review the log to check if different threads access the state. This works reasonably well for simple cases, but there are still several challenges we need to deal with:

- **Performance impact of new logs** – Some tools provide the capability of sandboxing requests, which will pause logs if they take up too much CPU.
- **Problems that might not be reproducible on a single container/server **– I glossed over the fact that when you add a new log, you can target a specific agent (application process). Instead of that, we can often target tags and the log will instantly be applied to all applicable tags.
- **Noise in our logs** – Some tools can log to the application logger. That means logs appear as if you called them in code. With piping, we can redirect the added logs to the IDE UI and remove all the noise (and cost) from the actual log.

## Deep insight into production

Logs are great when we have a general sense of the problem we’re facing. But there are many cases, such as transaction failures, that can be more amorphous. We need to see more details such as call stack and variable values in order to get our bearings.The problem is we don’t necessarily know what we’re looking for but we might know it when we see it.

When working locally, we would add a breakpoint and look at the local variables and stack frames. If this is an occasional failure, we can use a conditional breakpoint to grab the information in case of a failure. You can do a similar thing with a developer observability tool. The one difference is that you can’t break since stepping-over in production isn’t practical. You can’t “hold” the production server thread.

Some tools refer to this capability as snapshots, others call it capture or non-breaking breakpoints.

A production Spring Boot application could occasionally get transaction rollbacks. Using our developer observability tool, we placed a conditional snapshot on a Spring internal class (`TransactionAspectSupport`). We then received the full stack trace and all the variable values for the failed transaction. Upon reviewing the state, we could understand the root cause of the failed transactions.

Conditional snapshots (like we used in this example) are very much like conditional breakpoints. We can use a boolean condition referencing the source code to narrow the scope so we’ll only receive the applicable snapshot. Conditions can be anything; for example, “user.getId() == 5999965”. Notice that in this case I used Java to define the condition but usually it would be in the language of the current environment, you also have access to variables, methods/functions in the scope of the Snapshot.

One of the hardest things to debug is the nasty bugs that happen once in a blue moon. We can’t reproduce them locally and we get a “weird” stack from the server. We know where the problem is, but can’t imagine what would cause it!

In these cases we can place a conditional snapshot on the applicable line but increase its expiry time. Most tools in this field implicitly expire actions to reduce overhead, though this is sometimes configurable. Then we can come back the next day or a week later when the problem has been reproduced for us.

At this point, we will have the stack and the values of all the applicable variables in the stack. This is a godsent for this nasty class of hidden bugs.

## Is anyone using this code?

Even in a moderately sized codebase, it may not be obvious if code deployed to production actually gets called in production. Unfortunately, the answer is usually a shrug. We can use the “find usage” capability of the IDE but it only provides some of the story. The code might be “reachable” on a technical level, but no user will ever actually reach that line.

A few years ago we had a feature in an app and later on removed it from the UI. We had no way of knowing if people turned on that setting in the UI and just left it. So the backend code to support that “long gone feature” was still around.

That also meant we had unit tests covering it to increase coverage and with every refactor we had to make sure it works. A simple log showed that it was still used. But we wanted to get a better sense of the numbers.

A counter increments every time the counter line is reached and can be added like a snapshot or log. It can be conditional just like the other actions, so we can count the number of times people from a specific country reached a specific line of code. This is remarkably helpful when we need to make architectural decisions about the code.

We can leverage the Pareto principle (the 80-20 rule) to focus our optimizations on the code that’s actually used for future growth and improvement. By using counters, we can discover the area of the code that’s actually used.

## Pinpoint performance issues

A very common mistake is the N+1 queries in ORM (object relational mapping) tools. This happens when a single operation that should have fetched the entire result set ends up triggering a new query for every row.

You might overlook these errors as the database gets many queries and it’s often hard to associate the code with the resulting SQL. Locally this can go unnoticed without causing issues, but in larger production datasets, the performance impact can be significant. Unfortunately, in production the volume of queries is so big it’s even harder to notice the specific set of small queries that cause this. Since each individual query will appear to be performant, even a seasoned DBA might miss this.

A typical observability tool will probably point us at the general problem. For example, suppose web service X is performing badly. A single web service might trigger many operations in this container and possibly through microservices. How can we narrow this down?

When debugging code locally I frequently save the current clock time then a few lines below that print out the difference between the current time and the original time. This provides accurate low level measurements on the performance of a block of code. This is a common pattern that’s sometimes [built into the language APIs](https://www.rdocumentation.org/packages/tictoc/versions/1.0.1/topics/tictoc). The name tictoc refers to the sound of the wall clock and represents the two calls to it: the tic and the toc. We can add such a log to our production but then our production logs will be filled with printouts that are hard to read and quantify.

Metrics let us measure the performance of a block of code over time. We can mark a region in the IDE and add a measurement that works like the other actions we discussed. Conveniently, we can use metrics to narrow the scope. For example, if a specific user is experiencing a performance problem, we can configure a conditional metric on their user ID to see the specific lines of code that are at fault.

Thanks to this, we noticed a misconfiguration in our Spring Boot transaction behavior that triggered redundant queries. Since the code looked efficient on the surface and was indeed efficient when reached from a different path, we never suspected a problem!

## Tracking a zero day vulnerability

The recent [Log4j](https://stackoverflow.blog/2022/01/19/heres-how-stack-overflow-users-responded-to-log4shell-the-log4j-vulnerability-affecting-almost-everyone/) vulnerability was tough. It was easy to exploit before a patch was available and it was very hard to test against. Many developers had no idea that they used Log4j because it was a dependency of third-party code that might have been vulnerable itself!

I’m not a security expert, but the scope of the problem was immense. Log4J is in far more places than people even imagined, companies didn’t even know they used Java and were vulnerable. The severity was immediately clear to me as the bug allowed easy remote code execution. On the same day the issue went public, we added a snapshot into the vulnerable Log4j file in our project. This didn’t solve the problem or stop a malicious hacker. But if someone would have exploited this vulnerability we would have gotten information about the attack.

I later used this approach in the [Spring4Shell](https://www.datadoghq.com/blog/spring4shell-vulnerability-overview-and-remediation/) exploit as well. In that case, I could use the exploit to verify that none of our servers were vulnerable against that specific attack.

## Security implications of developer observability

Tracking zero days isn’t as impactful if the tool itself is vulnerable or exposes our production servers to risk. All the tools in this field (that I am aware of) don’t expose production in any way. An agent is added to the production applications and it communicates to the vendor server.

Developers only have access to the vendor server and not into production. This way the DevOps still maintains 100% control and isolates production without risk. There are many other security-oriented features that tools might incorporate such as PII reduction, certificate pinning, sandboxing, etc., but block lists are the most important one!

Some of you might have read this article thinking about remote debugging. This has many drawbacks/problems but one shines above all. Imagine a developer in your company placing a breakpoint on the user authentication code and siphoning off user credentials.

60% of company security breaches come from inside the organization. Disgruntled engineers could use tools like these for their own ends. This might also violate privacy regulations such as GDPR by effectively exposing private information.

Blockists are the solution to that problem. In them, you can specify the files/classes that should be excluded from actions. An engineer can’t add an action to such files and is effectively blocked from there.

When setting up the server environment these areas must be mapped to prevent malicious intent.

## Final word

Developer observability tools are a debugger designed for code running in production. They’re a seismic shift in deconstructing the DevOps/Developer silos, as they let us peer into production without the associated risks.

There are many use cases for which we can apply the power of these tools. In this article, I barely scratched the surface of what’s possible. The creativity of developers using these tools never ceases to amaze me. I encourage all of you to go through the list of features and review some of the top vendors in the field. Some of these tools are completely free on smaller scales, so you can conduct an investigation/proof of concept without the procurement hassle.

Tags: [observability](https://stackoverflow.blog/tag/observability/)
