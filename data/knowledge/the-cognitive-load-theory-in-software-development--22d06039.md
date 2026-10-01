---
title: "The Cognitive Load Theory in Software Development"
notion_id: 22d06039-b447-49df-b2c1-82378eb732d4
notion_url: https://app.notion.com/p/The-Cognitive-Load-Theory-in-Software-Development-22d06039b44749dfb2c182378eb732d4
last_edited: 2022-12-19T14:11:00.000Z
source_url: https://thevaluable.dev/cognitive-load-theory-software-developer/
tags: ["English", "Programming", "Article", "The Valuable Dev"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

“Our codebase is easy to understand. You won’t need more than a couple of hours. You’ll be productive right away! Guaranteed!”

It’s Dave, your colleague developer, explaining to a new employee how one of the codebase of your company shines like a diamond. The three of you work at MegaCorpMoneyMaker, the famous e-commerce. Looking at the codebase, with its millions of lines of code, arbitrary boundaries, and unexpected dependencies, you wonder how Dave can have so much confidence in his “explanations”. Davina, another one of your colleague, seems perplex, too; she listens patiently on a desk not far from yours.

Suddenly, she stands up, and ask Dave:

“What makes you think that this codebase is easy to understand?”

Dave looks at her, surprised, for a couple of second. He then rushes to answer: “Because I wrote most of it, in the pure tradition of The Old Sages: using good OOP design patterns, the DRY principle, and, of course, the all powerful SOLID Principles!”.

Davina takes a deep breath. She looks like she’s again on a mission. She begins to speak to the whole open space: “I’ll be giving a talk about metalearning (learning how to learn) this afternoon, and especially about cognitive load theory. We’ll see how some social experiments can help us understand a bit better our complex codebases”.

Davina will speak about the cognitive load theory and how we can apply its concept to develop better software. More specifically, she’ll cover:

- What the cognitive problems we have regarding software development.
- What is the cognitive load theory.
- What are the different forms of cognitive load
- How to learn more effectively.
- The limits of cognitive load theory.

Are your ready to dive into our brains, to understand a bit better some of its mechanisms?

## Software Development and Cognition: The Software Crisis

The whole tech team is now waiting for Davina’s afternoon talk. She begins in these words:

“It’s only by defining correctly the problems that we can find the best solutions. Dave says that our codebase is easy to understand; I disagree. So before diving into the cognitive load theory by itself, let’s ask ourselves: why is it so difficult to build applications which work, which are reliable, maintainable, _and understandable_?”

### Our Brain is Not Good Enough

“To understand what happens now, it’s useful to know what happened before”, continue Davina. “Let’s look a bit at some history here: why do we think that software development is a difficult endeavor and, most importantly, why?”

Let’s go back, with peace and love, in 1968. From the 7th to the 11th Octover, about 50 experts gathered in Garmisch, Germany, to speak about “software engineering”, a provocative title at the time. The idea was to ground computer science into “theoretical foundations and practical disciplines”, like any other engineering branch.

It’s in this conference that the term “software crisis” was coined. Developers and computer science researchers were already witnessing large software system becoming more and more complex. They were not finished on time, they were costing more than their estimates, and they didn’t always meet the specifications. Many possible causes were discussed during that conference; we can boil them down to a fundamental one, however: our brains are not good enough to process the sheer complexity of a large software system.

“To understand this point, let’s do a thought experiment”, continue Davina. “Let’s imagine that our brain is different in these points:”

1. When we read a codebase, we automatically _understand_ and _remember_ every single detail: all variables, functions, files, packages, namespaces, microservices, and whatnot. In short, every possible module on every possible layers, as well as their relationships.
2. We can then have access to all this knowledge with the sheer power of our mighty will, from the smallest detail to the biggest picture. We can navigate through all this knowledge effortlessly.

“If this was a reality, we would have no problem building software anymore”, you say in a breath. “Bugs wouldn’t exist. Even better: we could get rid of automated tests, continuous integrations, and every single layer of security we try to put in place in case our brain forget, misunderstand, or imagine something dead wrong.”

“We wouldn’t even need computers anymore!” continue Davina. “If we could put the whole reality in our head, with all the details, movement of atoms, and whatver’s happening out there, we would be similar to mentats in Dune, or even closer to the idea some have of God. Not bad, eh?”

But this is not how things are. Instead, we need to use a brain we don’t fully understand, and accept its flaws. As developers, we need to:

- Retain all the necessary information to do what we need to do. In short, having an imperfect, simplified, mental model of our codebase in our head.
- Try to understand how the codebase would run on a computer, that is to try to _simulate the runtime_ in our head.
- Understand the relationships between the different modules we need, at least to solve the problems we try to solve.

That’s a lot of things we need to cram in our heads! Can our brain handle it? Are we smart enough? According to the cognitive load theory, not really.

### Solving the Software Crisis

Let’s admit that our brain is limited How to tackle these limitations?

### Technical solutions

First, we could directly change the codebases themselves, trying to make them [simpler](https://thevaluable.dev/kiss-principle-explained/). Fundamentally, it’s what most of the software literature is about, in one way or another. I won’t dig too much in this subject, there are plenty of articles about complexity everywhere on the internet (including [this blog](https://thevaluable.dev/tags/complexity/)), or in the book shelves.

### Cognitive Solutions

There are downsides to these technical solutions, however. We don’t really fix the root of the problem here (our brain is not good enough), but we try to deal with the consequences. Even if we have powerful tools and workflows, we still make mistakes, we still misunderstand, we still have weird reasoning, and we still create bugs.

So, can we improve our brain?

I don’t know if you tried, but it’s hard as nails. That’s why we don’t really consider that as a solution, and we throw all our energy in the technical ones. We’re also used, as humans, to try to solve our problems with some kind of tools. After all, if we didn’t use sticks, stones, the fire, and the good old caterpillar, we would still go around naked in the savanna, wondering if a lion will ever come in our comfy cave to eat us all.

But I believe that knowing our limitations is important. First, to be a bit more humble: it’s not because we’re convinced about a solution that it’s the good one. It’s not because we’ve 879878 crazy tools that there won’t be any mistake. Let’s not forget that our brain has the tendency to simplify everything, even what should not be simplified.

Understanding how our brain works can help us create better tools, and find better solutions. More specifically, understanding how we learn can give us some clues how to write our codebases, for them to be more understandable by our peers. It can also help us to learn from a codebase we didn’t write ourselves.

## What is Cognitive Load Theory?

Davina pauses. Everybody seems to understand and agree with what she said. Difficult to argue that software systems don’t have bugs. She continues:

“One of the most interesting theory about learning and cognition is the cognitive load theory. Contrary to many good practices we take for granted in our field, the cognitive load theory is supported by many experiments to validate its different theories. Software development is full of gurus telling us what to do without many scientific facts to support their assumptions. The cognitive load theory can, at least, teach us two or three things to make our work even more valuable.”

She looks at the audience. Everybody is listening, curious to know what this cognitive stuff is. “Again, let’s look first at a little bit of history”.

### Where Cognitive Load Theory Comes From

The cognitive load theory was coined by John Sweller in 1988, in the paper [Cognitive Load During Problem Solving: Effects on Learning](https://github.com/Phantas0s/alexandria-library/blob/master/teaching_learning/_PAPERS/1988_cognitive_load_during_problem_solving.pdf). Sweller began to run experiments on our memory to find ways to improve our solving problem skills. We’ll come back to problem solving later (an important skill for developers), to focus for now on the goal of the cognitive load theory.

Basically, it aims to develop teaching and learning techniques taking into consideration how our brains work. It’s highly considered by teachers out there, because it’s backed up by many lab experiments. Again, in software development, we’ve too many assumptions and not enough data to make sure they’re legitimate. Don’t get me wrong: our intuition coming from our experience is valuable; but adding experiments and data can validate our “good practices” even more.

Our memory plays a major role in our work. As I said earlier, it’s difficult to put everything in our head: the details of a codebase, its runtime, all the context around it, its history, the stream of decisions leading to its current state, and so on.

How the cognitive load theory can help us understand a bit better our memory?

### Limited or Infinite Memory?

The cognitive load theory bases many of its conclusions on two critical observations:

1. Our short term memory (also called “working memory”) is limited. We can’t retain many “pieces of information” at once.
2. Our long term memory doesn’t seem to have a limit.

Our memory is both very limited and infinitely large.

### The Working Memory

Our working memory represents the information we are conscious of, at a specific point in time. If you remember part of your shopping list, it’s in your working memory: you might have recalled it from your long term memory, or you might have read it a couple of seconds ago.

George A. Miller, in 1956, wrote a [ground-breaking study](https://github.com/Phantas0s/alexandria-library/blob/master/teaching_learning/_PAPERS/1956_magical_number_seven_plus_minus_two.pdf) about the limitation of our working memory about our working memory. In there, he discusses many studies and experiments, to conclude that our working memory can’t held more than 7 elements at once. This number can vary, but not drastically.

We intuitively know it. That’s why we write things we don’t want to forget. If we could cram all our shopping list in our head for the next 10 days, we wouldn’t need to write them anywhere.

That’s also why it can be so daunting to read the code written by others (or by our past selves): we need to cram all of these new variables, functions, and their relationship in our working memory. I’m not even beginning to speak about performances, another dimension which might be important to consider. So we’re here, on our screen, struggling with too much information at once.

### The Long Term Memory

At that moment, Dave cut Davina: “Since our long term memory seems unlimited, we could store all our codebase in there! Problem solved.”

Davina answers swiftly: “That’s a good idea; but there’s still one problem: we can’t push whatever we want in our long term memory on demand. Also, to be used, the information in our long term memory needs to be brought back in our working memory, for us to effectively “remember it”, that is, consciously using the information. So we’re back to our working memory limitations, again.”

Nonetheless, it’s interesting to look at how the information is store in our working memory, because it can help us to find strategies to push information in our long term memory.

So, how this information is stored? The paper of George Miller speaks about “chunk”, and the cognitive load theory call them “schema”. They’re more or less the same: it’s a group or related information in our long term memory. It’s interesting to note that each schema takes only one space in our precious and limited working memory, even if this schema can contain by itself a lot of different, _related_ “pieces of information.

For example, we can easily retain three words, like “sheep”, “boat”, and “sea”. It’s more difficult to remember each letter individually, like a,e,s,p,e,o,b,e,h,t,s, and e. It’s because we have in your long term memory the schema of a boat; you can visualize some of its details, how it looks like, what it’s used for. In short, we have access to a lot of information under the name “boat”. This schema of boat only use one “place” in our working memory. That’s how you can think about a boat, an umbrella, and a giraffe at the same time.

John Sweller, in his initial paper about the cognitive load theory, speaks about a study which found that the main difference between novice and expert chess players is the number of schemas in their long term memories. These schemas are composed of different arrangements of the pieces on a chessboard, and what to do when they pop up in real life.

Having these chessboard configurations in their long term memories allow experts to recall them easily in their working memory, and use them when needed. Of course, it can take decades to build this sheer amount of schemas in our experts’ brains.

Notice that long term schemas can’t be constructed passively, by listening or reading. Our brain needs to actively _participate_ in the learning process.

Dealing with a codebase is similar: if we know it enough, we know that the file named “Shipment.php” will contains information about the shipment logic in PHP, and that the file “Handler.php” handle the HTTP requests. Again, we use schemas stored in our long term memory to recall what we know, enabling us to understand, reason, and spend our limited brain power on solving more important problems, without overloading our limited working memory.

Davina is now speaking slowly, to be sure that everybody understand this crucial point of the talk: “It begs the question: how to create these schema more effectively? Said differently: how to learn more effectively?”

## Different Forms of Cognitive Load

Davina continues: “to understand how to create these schemas, we need to understand what kind of cognitive load can help us doing so, and try to get rid of the cognitive load preventing us to build them”.

### The Intrinsic load

The _intrinsic load_ is the inherent difficulty for some concepts. Said differently, it can be difficult to create schemas in our long term memory for two reasons:

1. Because the concepts are difficult are complex: they have many parts, and many relationships between these parts.
2. Because the learner lack some prior knowledge to understand the concepts.

For example, if you never programmed before, it will be difficult for you to learn a new programming language. At least, more difficult than an expert programmer, who has already useful schemas in her long term memory. Of course, if the programming language is Haskell (a difficult language to learn by itself), it will be more difficult for the expert than to learn a language very close to what she already knows.

To come back to our codebases, the intrinsic load can be the complexity of the business domain you’re coding for, or the intrinsic difficulty of the task itself. For example, if your manager ask you to develop a disrupting functionality where the users can do everything they want (like a spreadsheet), it will be more complicated than developing a simple login page.

It doesn’t mean that you can’t influence the intrinsic complexity of a task. If you have a better solution which might not bring as many benefits as the one proposed, but 10 times less complex, your stakeholders might agree to follow your advice.

After all, managing complexity is the heart of our job. Mainly to make our code understandable, that is, easier to learn. We can map this idea of “intrinsic load” with what Fred Brook calls “essential complexity” in his paper [No Silver Bullet](http://worrydream.com/refs/Brooks-NoSilverBullet.pdf).

### The Extraneous Load

The _extraneous load_ is the load in our working memories which makes building schemas difficult in our long term memory. To come back to Fred Brook, it’s what he calls the “accidental complexity” in the paper cited above: complexity which could (and should) be avoided. This complexity is often created by the developers themselves (who didn’t find the better solution).

Extraneous load needs to be avoided, for developers to build accurate schemas in their long term memory more easily. We’ll see later in this article some ways to do exactly that.

### The Germane Load

The germane load is the opposite of the extraneous load. It’s the cognitive load helping creating schemas. We want a low amount of extraneous load when we learn, but a high level of germane load.

Now that we’ve defined what we want and what we don’t, how do we improve our learning skills?

## How To Learn More Effectively?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

The audience looks know quite tired. It has been a lot of information at once, and Davina knows it. She tries to summarize:

“We lack natural cognitive tools to perform complex reasoning; more precisely, we can’t juggle with a lot of information at once, unless we’ve access to schemas we’ve created before, in our long-term memory. These schema can be complex, but at the same time they’re not heavy cognitive load for our working memory.”

She continues: “To create these schemas, we need to increase the germane load and decrease the extraneous load. We need to deal with the intrinsic load anyway, because this complexity is part of what we try to learn.”

We could see three areas where learning is important for a developer:

1. The technical dimension: the tech stack and different technical practices used in the codebase.
2. The domain dimension: a codebase _codify_ some business problems and there solutions. We need to be aware of the intent of the code, and the relationship of our different modules with the domain.
3. Problem solving: including analysis of the problem space, exploring the solution space, experimenting, and balancing the solutions’ benefits and drawbacks to take the best possible decision.

How can cognitive load theory can help us in these different learning endeavour?

### Technical Learning

Since our working memory is limited, it’s nearly impossible to correctly understand and learn from a codebase if we don’t have any technical knowledge. Not only we can’t write code if we don’t have this technical baggage; we can’t even understand it.

There is a spectrum between not knowing anything about the tools and the programming language of a codebase, and knowing every single one of them. Our knowledge (our mental model) will never be perfect, but there are many technical schemas we still need to have, to understand the very basics of a codebase.

So, how to acquire these technical schemas? I’ve already wrote about [how to learn a programming language](https://thevaluable.dev/how-to-learn-programming-language/). To summarize, I would:

1. Look at working examples and try to understand them.
2. Isolate concepts and ideas to understand them.
3. Group complementary piece of information together to build schemas.

In parallel, I would practice as soon as possible: building a simple project using the tools I want to learn, to make it more personal and facilitate the creation of these long term schemas. Remember: active learning is what enable the creation of these schemas.

### Working Examples

The learning strategy between a novice and an expert will be different. The first doesn’t know much about software development, and the second has many schemas at hand.

A novice needs to learn everything, so going through already existing codebase, running them, and experimenting with them is a good option. That’s what the cognitive load theory means by “working examples”.

The working examples can take other forms:

- Examples how to use the different interfaces (API) from some documentation (external documentation, comments in the code…).
- Automated tests can describe what the different modules are doing at runtime, helping to build accurate schemas.

It’s not only about learning the codebase, it’s also about understanding its runtime, too. In that sense, it can be valuable to modify it, run it, and try to understand the feedback we get. Starting from an area in the codebase where you feel comfortable and make you way through modules you don’t know can be a good idea. Indeed, in that case, you have already some schemas in place which can help you build new ones.

A more advanced novice (or an intermediate developer) can build their own projects and learn from them, while trying to analyze their own work.

An expert will often need to learn new codebases, or piece of codebases written by others (or their past selves). But they will have access to general schemas helping them in that task; the process becomes then more automatic, less analytics, and therefore less tiring. Their schemas can tell them what are the possible actions to take. This approach has a downside, however: it’s sometimes difficult to find solutions different than the schemas we already have (“thinking out of the box”) and to explore new solutions and ideas. That’s why, [even for an expert](https://thevaluable.dev/expert-blind-spot-software-development/), it’s often beneficial to spend some time to analyze the problem space, to experiment, and to see if other solutions are possible.

After all, all problems are slightly different, and the context can change many of their attributes.

It’s important to note that working examples will be more useful for beginners. As the expertise grow (as more and more schemas are created), working on problems in group can be better to expand the expertise. That’s why working in diverse teams is important.

### Isolating Concepts

We also need to isolate the different concepts we want to learn, to better understand them. It might sound obvious but, again, this time we’ve some scientific facts to back up the claim.

For example, if I tell you that “chat” is the translation of the word “cat” in French, would you say that you “understood” something? Not really.The concept is easy to grasp because the knowledge is isolated. We’ve already have the schemas of word, letters, foreign languages, and so on, so we can use them to instantly understand what a translation is.

We have to actively understand concepts when they get more complex. We then need to store all the necessary “pieces of information” in our limited working memory to create coherent, meaningful schemas out of it. Learning concepts in isolation allow us not to overload our poor working memory. We can then link all these isolated concepts together.

It means that we should isolate what should be isolated in our codebase as much as possible. Said differently, we should limit the dependencies between the different modules we have, precisely for people to understand the system.

### Redundant Information

Pieces of information which are expressed multiple times increase the extraneous load and, as a result, doesn’t help us create our precious schemas. It makes sense: why put the same information in our working memory more than once? When I say “same” information, it could be information formulated differently (but stay inherently the same), like a drawing conveying the same information as a block of text.

On a higher level, that’s where the DRY principle becomes important: not repeating multiple time the same knowledge, even if this knowledge is represented in different ways. In the codebase, or outside of it. If we can, we shouldn’t repeat the same information in comments, in the code, and in some documentation for example.

I’m speaking here about “pieces of information” carrying the same knowledge. I’m not speaking about duplicating the same pieces of code. The difference is important, and the concept is explained further in [this article about the DRY priciple](https://thevaluable.dev/dry-principle-cost-benefit-example/).

### Grouping Complementary Pieces of Information

As we saw, according to the cognitive load theory, schemas are groups of _linked_ information. It means that, even if it’s easier to learn things in isolation, it’s useful to build our schemas by understanding the relationships between the different elements, too.

It’s really close to the duality of cohesion and coupling: it’s easier to learn things in isolation, but it’s also easier when things _related to each other_ are together. This [article about cohesion and coupling](https://thevaluable.dev/cohesion-coupling-guide-examples/) explains further these two concepts.

### Learning the Domain

A codebase is inevitably linked to a business domain, or even more generally, to a context. The complex, real world can change a lot, and these changes will influence the codebase. After all, an application often tries to influence the real world by allowing some sort of automation. In return, the external world can influence a codebase: if the domain has new regulations which makes the software illegal, for example.

Learning the domain is essential. This knowledge can comes from two different sources:

- Any written source: the codebase, some documentation…
- Learning about the domain from your peers: developers, managers, or, even better, from domain experts.

### Learning From The Codebase

To get some good domain knowledge, you might first want to look at the codebase itself. After all, a codebase is the representation of a domain, or at least implement the solutions of the domain’s problems, and the different domain’s opportunities.

But remember: our working memory is limited. To be able to grasp the domain knowledge from the codebase, it’s nice to have:

- Good naming, directly related to the domain knowledge. If we already have schema of this domain knowledge in our long term memory, it would then be easier to match them with the actual modules of the codebase.
- A glossary (what DDD calls “[ubiquitous language](https://martinfowler.com/bliki/UbiquitousLanguage.html)") to have a direct mapping between the naming in the codebase and what it means in the domain.
- High level documentation, in comments directly in the application (easier to update), eventually outside the application too (written somewhere).

If we don’t know anything about the domain (when we begin to work for a different company for example), it will be more difficult to represent it with our long term schemas. As a result, we should learn the foundations step by step.

You can try the instructional advice I gave earlier:

1. Go through working example of the most important domain concepts in the codebase (look at comments, tests, external documentation…)
2. Isolate important concepts in the codebase, and try to understand the elements composing them. It might help creating some schemas in your long term memory.
3. When you really understand a couple of domain concepts in isolation, try to find out in the codebase the different dependencies between them. It can help creating linked schemas.

Learning about a domain is, to me, the most difficult part. It depends on the domain itself, and your prior knowledge; for example, it will be easier to understand an e-commerce than a software to bend aluminium, if such thing exists.

That’s why I could see more and more developers, in the future, specializing in a specific domain, instead of specializing in a specific programming language (or any other technology).

When writing a codebase, it’s important to explicitly state the _intent_ of the domain: what the domain is about, what domain problem the codebase try to solve. For example, for an e-commerce, if everybody considers that a shipment never exists without an order, we could put both our entities “Shipment” and “Order” together; in the same namespace or the same package. It’s a way to “translate” the domain into code.

If you follow this principle, developers (and your future you) will have an easier time to map the knowledge of the domain with the codebase itself.

### Documenting the Domain

Documentation can be useful to understand how a codebase relate to its business domain, but it’s not without its drawbacks. It forces us to hold both source of information (documentation and codebase) and map their knowledge in our limited working memory. In cognitive load theory jargon, it’s called the “split attention effect”. If the documentation is not updated correctly, it’s even worst; we can end up create _wrong_ schemas in our long term memory.

It’s better if the code contains all the important “pieces of information” without the need to document it. High level comments can be useful (explaining some part of the domain, for example), as long as they don’t explain what the code is doing; for the later, the code itself should be self-explanatory as much as possible.

Sometimes, however, documentation is inevitable, at least to tight up the codebase with the business domain in more significant ways.

### Learning From Your Peers

We can’t learn everything from a codebase, because many decisions and use cases won’t appear directly in the code. The next best place to learn about the domain is with proper humans. You know, these bipeds we can’t stop complaining about.

Asking somebody to learn more about key domain concepts is useful, but it’s also a passive way of learning, making difficult the creation of schemas. It’s even better if the explanations have examples and practical applications. Visualizations (like diagrams) can also help create our schemas.

Conversations are useful. That said, we have apparently two working memory: a visual one, and an auditive one. It means that we can potentially double our memory load if we work with both.

It’s also useful to capture conversations in some forms. First, because our schema will be different than the actual conversation. They’re like a map: they’re not the reality in all its details, but our interpretation and understanding of all the concepts and their relationships.

If you explore the domain, it can be useful to write a stream of conversations somewhere, with the date and what you understood. The knowledge might be duplicated in the code, but you don’t necessarily try to learn from both at the same time; both source of information could be considered independent.

## Cognitive Load Theory and Problem Solving

Davina seems very excited to finally speak about problem solving, one of the cornerstone of software development: “The cognitive load theory is also useful to understand how to solve problems more effectively. When our working memory is full, it’s not easy to think about the possible solutions anymore. Our mental model will only struggle to understand the different component of the problem.”

In fact, there are many things to consider when solving a problem:

1. The current problem state.
2. The goal state.
3. The relation between current problem state and goal state.
4. The relation between problem-solving operations.
5. If, when decomposing the problems, there are subgoals, a stack of them need to be maintained.

If we don’t have schemas to help us understanding the problem, it will be difficult to solve it. That’s why understanding the problem is so important: it frees us some cognitive load to be able to think about a solution. It also enable us to create useful schemas we can use later, for the same category of problems.

Experiments also showed that learning some general problem solving skills doesn’t help. As discussed in [another article](https://thevaluable.dev/learning-developer-efficiently-effectively/), the principle of “transfer” (that is, applying what we’ve learned from one context to another) is difficult to do for us, poor humans.

Since the business domain influence the problems we can have in a codebase (except if the problem is 100% technical), my experience tells me that it’s not easy to apply the schemas learned from one domain to another. I remember once a domain expert being surprised that a software developer could switch from one domain to another so easily, when they change companies for example. I don’t think it’s easy, and I’m not sure we’re doing such a good job at it. Again, maybe at one point we’ll see more and more developer specialized in a specific business domain?

I also doubt that solving [kata](https://handwiki.org/wiki/Kata_%28programming%29) (general programming challenges) is useful. I did many of them at one point in my career, and I always had difficulties to reuse what I learned. Simply because I couldn’t transfer the knowledge from the very abstract world of kata to a specific codebase shaped by domain problems. Katas, to me, are too general for being useful.

Additionally, the problems we solve with katas are rarely encountered in a codebase, because these problems are already solved. Searching on The Internet and finding some examples is more useful to me than solving 290128 katas, hopping to see the same exact problem one day in a real life scenario.

I believe that working through codebases used for specific goals will teach us more than these general programming challenges, and enable the creation of our precious schemas we can reuse.

## The Limits of Cognitive Theory

Now that we spoke about many important development practices in the light of the cognitive load theory, let’s ask ourselves: is this theory accurate? What are the limits?

First, let’s look at the studies themselves. Unfortunately, They don’t really take into consideration other factors which can also play big roles in the learning process. For example: the motivation of the learner, how the learner think about her learning capacities, and so on.

Nobody really question the limitations of our working memory. But what’s the limit? Some argue that, in the context of real learning (learning outside of social experiments in laboratory), we can put more information in our working memory.

In fact, these social experiments are quite different from real life. People don’t react the same way during experiments. For example, some participants might try to please the researchers, without even knowing it.

Additionally, these experiments are quite short (with short study time), and the candidates won’t work through subjects they find interesting.

That’s not all: nobody seems to have a good metric to measure the potential cognitive load of a piece of information. When researchers ask candidates to measure this cognitive load on a scale, the results is not consistent from one candidate to another. There’s a lot of subjectivity at play, here. Also, we’re not all equal when we speak about memory: some can remember more, some less. It’s also difficult to understand what we can store and what we can’t.

An objective metric, capable of measuring the cognitive load of some information, could definitely help a lot. Unfortunately, for now, we have to do some trial and errors to see if a codebase is easily understood by the majority of developers.

That said, the cognitive load theory is still useful as a framework, when we have to deal with complex systems. Software codebases fall very often in this category.

## Expanding Our Brain

Davina concludes:

“It has been quite a ride to try to apply the conclusions of the cognitive load theory studies to software development. Let’s summarize what we’ve seen together:

- The information we’re consciously aware of, at any point in time, is stored in our working memory (also called short term memory).
- This working memory is limited: according to lab experiments, we can only keep 7 “pieces of information” in our working memory, more or less.
- We have another type of memory: the long term memory. We store the concepts we understand there. It takes the shape of schemas, many related “pieces of information”.
- It seems that we can store an infinite amount of information in our long term memory.
- A schema from our long term memory don’t take as much space in our working memory than new information.
- To understand our codebases, we first need to make them as simple as possible, to not overload our working memory. We should avoid the “accidental complexity”.
- Codebases need some “essential complexity” to answer the business domain needs, however. That’s why creating schemas of our codebases in our long term memory is important.
- Going through “working examples” is great if you don’t have any idea of the codebase. You can then run part of it, look at the tests, and experiment.
- Isolating concepts which should be isolated, and group together what should be grouped according to the domain can able creating schemas both accurate on the domain and the code sides.
- Redundant information can lead to space in the working memory taken for… nothing. Applying the DRY principle and trying not to repeat the same knowledge in many different places is important.

Davina finish by underlying what makes the job of a developer difficult:

“We write code as much for ourselves than for the others. The computer doesn’t have to understand anything; it only has to follow the different branches we laid out in a logical way. I would argue that our capacity to communicate and _teach_ correctly using our codebases is the most important. It can have heavy consequences on the short and long term, and can bring many misconceptions and bugs overtime. It can be the source of a decline in the quality of the codebase itself.”

She concludes by these words:

“We’re humans, and we have our limits. Knowing them is important: to stay humble, empathetic with our peers, to understand our mistakes, and to improve as human beings.”

### Let's Connect

You'll receive **each month** the last article with additional resources and updates.

[Here's how it looks](https://buttondown.email/thevaluabledev/archive/the-valuable-dev-new-article-about-vim-and-many/)

You can reply to any email if you have questions, problems, or feedback. I'll write back as soon as I can.
