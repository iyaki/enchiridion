---
title: "The Fallacy of Strongly Typed Languages"
notion_id: 2107cb6c-23be-4e5f-8ac0-61c06978be0c
notion_url: https://app.notion.com/p/The-Fallacy-of-Strongly-Typed-Languages-2107cb6c23be4e5f8ac061c06978be0c
last_edited: 2023-01-13T18:00:00.000Z
source_url: https://hackernoon.com/the-fallacy-of-strongly-typed-languages
tags: ["Programming", "Article", "Hackernoon", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Today I want to share a piece of enlightenment that struck me this year. I have to mention that during most of my career I worked as a Java developer and therefore this may be specific to the domains where strongly-typed languages are in use rather than to a whole industry in general. Java is a statically typed language like C#, Golang, or Typescript and to work with the data, we have to declare the shapes and structure of the objects we will use. As I realized this is only sometimes the best approach in programming.

A few years ago I got to be part of a system with most of its guts written in NodeJS, which is a dynamically typed language. That system was written very rapidly and therefore was an absolute mess. Working with it was painful and I wouldn't say I liked it. Not only did it have an enormous codebase but having to debug something and understand how features work was a headache and a challenge.

One of the main reasons the system was difficult to comprehend was the lack of information about the objects being passed around in the code. When I jumped to a file with the needed logic, I couldn't get a clue about what structures were processed, where they came from, or even what those fields were - numbers, strings, or objects. From that random part of a code I couldn't get this knowledge even if you traced back the origin of a code flow - some controller or external invocation. The code usually looked like this.

```plain text
let answers = {}

(request) => {
	if (request.questions) {
		answers = callExternalSystem(request.questions.sections)
	}
	if (!answers.detailed) {
		answers = callAnotherExternalSystem(request.questionsSetId)
	}
	return answers
}

```

Can you understand what the result of this function will be? What will be sent to external systems? What if you need to add some additional information to the response of this method? Every time I had to deal with almost all of that NodeJS code, this frustration caught me. But it would be the same if it was python without types or any other language. I hated that I always found myself in a situation where I didn't know what was going on and to have a bit of understanding I had to contact the author or go and try luck with finding those external services APIs.

I couldn't understand how authors worked with such vague tooling and information in the data flow. Wouldn't it be better even for authors if in 6 months, when they don't remember anything about the system, types helped them to see the data in use right at the moment of reading the code?

In the same company where I worked they asked me to write a service that would proxy requests to several downstream services. This was going to be a temporary solution for multi-region support of the calling service. Based on a few fields, I needed to decide which downstream service to call. It took about two weeks to collect models of each service, implement nice DTOs and clean data flow, and resolve all the inquiries regarding the API with the authors of downstream services. Everything was sound and clear until, while implementing the service, I realized that I used only 3 fields to route the request flow, and all that work could have been omitted. After a few days of agonizing about this fact, I accepted the sour reality that it would have been better both code and business-wise if I didn't implement all of this. It would have been even better for the business since they would have saved the developer time. From the code perspective, there was no use in all DTOs and data model investigations.

From that moment I became suspicious about the bare minimum of work I need to do and what business actually want - a quick shortcut for their problem, a long-term solution, or some compromise between these two.

In one of the following projects, I heeded to implement a feature, and while browsing through the codebase, I noticed a comment that this code implemented a Tolerant Reader design pattern. After reading a bit about what it is, it became clear that it was the solution I needed to implement back then in that multi-region proxy. It turned out that the design where the application doesn't care about models of data that it accepts, except for the actual fields it uses, is quite common. Even more, this pattern is the default programming approach to handling and passing data in NodeJS applications. They knew that! For some reason, in the NodeJS world they don't spend time to model everything if the data is only to be sent to some other location or if the business logic requires only a few fields from the request. The perfectionist engineer ego was again offended and in a fight with the realization that all the tricks learned about Generics, Type Erasure at Runtime, Schemas, and so on are indeed needless in many cases. But I had to move on.

In several months I joined a new company and was asked to fix the issue with one downstream service we were using wrong. One of the crucial things that I knew about was that there was a high sales season coming up to us in a week, and this fix was critical for the business. I was stressed since for investigation, reading all the code, and understanding what was going I only had 3 days. To my surprise, when I saw NodeJS code similar to what I showed before, I wasn't terrified and didn't feel any anger. Now I was gloating over the fact that this was typeless code. The first time I said to myself - aha! I know I don't know anything, but I don't need to! I will only need to add a few checks and an extra field to the request! This was precisely the case, and I spent about 2 hours implementing and sending the fix to production. Everyone was happy, even me.

With time it got to me that using schemaless databases or key-value storages like MongoDB or DynamoDB is the same approach and is valuable for the same reasons. Not having to think and work around the schema of the data and the freedom to solve problems as they appear allows us to get feedback sooner and implement solutions faster. Combining these 2 design approaches - tolerant reader pattern and schemaless databases significantly bootstraps and simplifies development.

When I create a new service today, I ask myself the following questions. How soon does the business need this service? What maintenance period is going to be for it? Does it need to be performant or handle lots of data? How tolerant will the required logic be of the failures and errors? Answers to these questions help me limit the upgoing service's complexity. Modeling data inside the service is a time-consuming task. Using a database with the schema requires much more planning and investigation, even before bootstrapping the code. Even adding tests may be redundant for some scenarios. In the end, any business needs a working solution as soon as possible. Since then, saving time and reducing complexity has become my main criteria when thinking about implementation devices.

What do you think? Let me know in the comments.

Happy coding and more straightforward solutions for you all!

L O A D I N G
. . . comments &
