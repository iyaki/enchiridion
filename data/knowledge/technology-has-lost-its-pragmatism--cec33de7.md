---
title: "Technology Has Lost Its Pragmatism"
notion_id: cec33de7-54d4-4218-a1bf-67a54d07c633
notion_url: https://app.notion.com/p/Technology-Has-Lost-Its-Pragmatism-cec33de754d44218a1bf67a54d07c633
last_edited: 2023-02-18T01:56:00.000Z
source_url: https://danielbmarkham.com/technology-has-lost-its-pragmatism/
tags: ["Reflection", "Programming", "Article", "English"]
---
So we have two questions here: One, how do we draw boundaries around our computer programs such that we minimize the risk of any one concept failing? Two, how do we define any one term such that it is impossible for there to be a misunderstanding or for the program to crash?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

_This is part of a series of essays about the _[_Four Philosophies That Drive Technology_](https://wiki.info-ops.org/index.php?title=Main_Page#Philosophy)

It's difficult to lose something you never realized you had.

When I was a kid, my family had a mundane friendly argument over who threw away dad's back-scratcher. Dad had been using it for years and suddenly it was gone. Nobody was able to answer his question. In fact, nobody was really sure what he was talking about, although the term "back-scratcher" is not that ambiguous.

After a bit of childhood detective work, it was determined that mom had thrown away the spaghetti fork. When we stopped eating pasta, it was put away. For a while, the family used it to reach far-away objects in the pantry. Then we started using something else to reach pantry items. Dad had absconded with the spaghetti fork (which he only knew as an object-reacher that was unused), only to repurpose it as a back-scratcher. When mom came across it years later, she recognized that it was no longer being used as a spaghetti fork or an object-reacher and rightly thrown it away.

Everybody was speaking the same language in the same culture, yet miscommunication resulted in Bad Things Happening. Meanings of things shift over time and from context-to-context. As pointed out in the [Tripartite Semiotics essay](https://danielbmarkham.com/the-biggest-problem-in-real-world-computer-programming/), simply because something has an unambiguous name you can look up in a dictionary doesn't mean that this name is useful for all communications regarding the object the name refers to.

This is not a rule about cooking implements. In fact, it's quite a profound rule I keep coming across in situations from Bounded Contexts in Message Queuing systems to random run-of-the-mill social media flame wars. The rule is a critical part of determining compilation units, features, product roadmaps, startup canvases, corporate divisions ... the list goes on and on.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[From a recent tweet with a friend writing a book on Event Store](https://twitter.com/danielbmarkham/status/1623803017001680896)

When I was writing a book on backlogs and a second book on writing better programming, I came across a question I had rarely heard addressed before in generic terms: how do you know when/how to **stop** developing? I'm not talking about making your unit tests pass. My question is at the generic level. Creating an artificial example, let's say I'm adding email to my phone app I'm developing. When adding features, I can move to a more detailed level of coding. Perhaps I start doing some SMTP programming over a webhook with a custom back-end. I can move to a more generic level of coding. Perhaps I get a library that handles email. I implement interfaces to an email gateway in that library, and just to be sure I have all of the features I might need, I "move out horizontally" and implement interfaces to a message class, a message handler class, an email response class, and so on.

I can move to using more broad abstractions. I can move to using more detailed custom-made code. I can move to more simple code that uses a much more complex library. Each of these choices has benefits and drawbacks. Which choice do I make?

The usual answer is that you add enough code that you might need in order to not need to come back and code later. Technically that's an answer, but there's no definitions of what you might need, or why or why not you might need to come back. Even if there were, there's no information on how you might decide whether to "come back later" or recode, or extend, or genericize .. and so on. It sounds like we've provided an answer, but not really. We can code something a thousand different ways and rightfully claim that I'm following this rule. And we do.

Programmers are not alone in this problem by any means. Although it can be described as a programming problem, the generic problem is actually a problem in philosophy. Which philosophy should I use? Just like I need to code that email feature in my app and there's a thousand ways of doing it, there are thousands of philosophies in life and we are forced to choose among them whether we realize that or not.

Just like there are three million ways to do email, since before people started writing there have been really smart people coming up with systems for determining truths, making tough choices, understanding the world around us, and living a good life. And just like that email feature problem, there's no clear way to choose one over the other. The generic advice "Pick one that works for you" seems to be an answer, but just like the coding answer, once we dive in we're left with more unknowns than before. It just appears to be an answer.

In the last couple of centuries, as these philosophies have become translated, written, published, and taught in school, it's become more and more evident that none of them seem to work with the others. Personally, I would like to practice Stoic Existentialism in an Anti-Solipsistic way, but I struggle with how these three philosophies can exist together, if they can exist together at all.

If you thought email was tough to code, once you branch out into real-world decisions with vague criteria, like how to spend social welfare money or whether religious beliefs are philosophies or ideologies, it gets worse. By the mid 1800s, really smart people were arguing in good faith about a multitude of complex issues using all sorts of similar names to mean things slightly different. We had a thousand back-scratchers. Some were missing, some were being used the wrong way, some folks didn't know what a spaghetti fork was, and so on. It seemed impossible to have any reasonable public discussion about anything complex without ending up talking semantics or getting the dictionary out. (This continues today and it drives some folks crazy)

I was exposed to the answer to this problem 20 or more years ago reading essays by programmers who became startup founders. I can't remember the exact one, but the general answer was "Don't feel like you need to attach to one philosophy or not and stick to it. Instead, pick and choose what you need as you need it for the life problem in front of you. Tomorrow, with a different problem, pick another one."

I liked this, but I was unsure of when to pick different philosophies, when to change, how to know if the one I was using was working, and so on. Then I came across this quote from Charles Sanders Pierce.

> Consider what effects, which might conceivably have practical bearings, we conceive the object of our conception to have. Then the whole of our conception of those effects is the whole of our conception of the object. (Peirce 1878/1992, p. 132)

(From the [Internet Encyclopedia of Philosophy](https://iep.utm.edu/peircepr/#H2))

Peirce was not a programmer. In the late 1800s there were no programmers. Instead, he was a master logician, doing the best he could to emulate how future computers might reason. I keyed on the word "object". He was saying that while we understand an object/concept by the name we give it, the only real way to deal with it is to use the things that it does, not the name.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

I can't keep using that table for my homework. What does the word "table" mean?

People who are not programmers might think this kind of nit-picking is silly. Can't ChatGPT already describe the types of tables? Isn't "table" a unique word with a lot of history?

These rebuttals are true but not relevant. People can only deal with ambiguities. That's our life. Our brains fill in the gaps and we move along. AI Chat works because the person chatting with the computer is the one doing all the work. There's no red light when you misunderstand something. Conversation continues along. It's only when there's some problem that we end up diving into a word or phrase, only to find that each of us means something subtly different.

Programmers deal with this with every class and variable they create in a computer program. We name them and we use them according to whatever name we've given them. Programs are compiled webs-of-names that exist in some sort of logical synchronicity. Good programmers know that every name we use is wrong in some way or another. This is why talking to a programmer can be so frustrating sometimes. They keep asking you to define simple things that everybody knows. They know that everybody doesn't really know much, but the program has to work anyway.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Random Domain-Driven Design diagram [sourced here](https://www.mirkosertic.de/blog/2013/04/domain-driven-design-example/). The larger boxes are presumably [Bounded Contexts](https://dzone.com/articles/bounded-context-in-my-view).

## Borders, Boundaries, and Contexts are a PITA

Here's what Peirce and many philosophers since him have realized that we all need to know: once you become meticulous and demand that all the terms have firm meanings and work together, the whole thing falls apart. One word is likely to fail when talking about writing a program. Seventeen certainly _will_ fail.

So we have two questions here: One, how do we draw boundaries around our computer programs such that we minimize the risk of any one concept failing? Two, how do we define any one term such that it is impossible for there to be a misunderstanding or for the system to break?

Functional Domain-Driven Design (DDD) seeks to make programs impossible to be in a state of error by tightly-defining the types used, as shown above. It is certainly possible to program a type that will do only the things you want it to and will never crash the larger program. But all you've done here is pretend that humans don't exist. Remember that this is a human language problem, not a programming problem. Programmers are stuck with it because we're the ones stuck trying to assemble various pieces of human language into a running program. You can make the program solid, but all you've really done then is change what might be a discussion about what one term might mean into an incredibly-complex discussion about how some web of terms logically interoperate. Don't be surprised when that fails.

Pragmatism gives us the answer. The thing is only what we use it for, nothing else. Humans will drag a bunch of other crap into a problem if you let them. You're a human. You will too. But as far as the code goes, it's what it does, nothing else.

Translation: code exists to pass a test. Types, names, and boundaries exist to pass a test. Nothing else matters, and in fact anything else can and will be a problem at some point if the code continues running. **The test can only be an external behavior valued by some human**. The back-scratcter was a back-scratcher to my dad. That's all it was. To everyone else it was junk.

Code should not exist unless it's justified by passing some test, and code that passes that test in addition to others is suspect due to semantic leakage. The other tests the code is passing is dependent on other people with other definitions of things, and it's a guarantee that their things aren't going to be the same as your things, no matter what the names are.

_**A boundary is defined by one external behavior test. This will demand what goes inside the code such that you can make it so that it cannot fail. You can try to create a boundary based on multiple external behaviors, but it quickly (if not immediately) becomes impossible to make it such that it cannot fail. In fact, the risks rise exponentially that it will fail, and in unpredictable ways.**_

There are some logical consequences that fall out of this pragmatic programming maxim, some of which might make you uncomfortable.

- If you show a variable/class/module name to another person, you've changed what the thing is assumed to do without knowing it
- If everything just exists to pass one test, then feature changes, rewrites, and so forth are all replacements, not modifications. Once the tests pass, you should never open that code again
- Therefore variable names shouldn't go in the code, at least by the time you're done. Business meaning should exist outside executables
- Databases and other coupling exists, of course, but domain interaction is a business feature itself. Therefore it should associate with a group of behaviors and whether the database is up or down shouldn't matter. Otherwise you've got too much coupling
- Stated differently, complex class hierarchies and libraries become complex because creators try to stuff as much domain knowledge and intent as possible into the code, figuring that it'll be needed by others. They rightfully play it safe. Pragmatism, on the other hand, means that the goal of programming just the opposite: the removal of all intent from source code
- Because of this, imperative and object-oriented coding cannot be pragmatic. They depend on grouping together human language concepts that we know to be loosely-defined and broken. The more they group together, the worse understanding and maintaining becomes
- Functional programming can be pragmatic, but only when decoupled and pure. The names of the functions matter, the names and types of the data does not. It's certainly possible to code that way and it's been done for decades, but it's not the way we teach computer programming. We suffer because of that
- If this looseness is a problem in programming, it's certainly a problem everywhere else in life. We just don't expect our twitter feed to compile. The answer here is very similar: agree on tests and remove definitions and terms as much as possible from contentious conversations. Where there's no contention, double-check anyway. There's always meaning shift and semantic leakage.

At some point, all human language is going to be bullshit, but minimizing the terms used, creating well-delineated and firm boundaries, and being clear about acceptance tests can help a programmer or anybody else create a "web of meaning" such that enough terms will hang together for a useful conversation to occur. That social analogue is outside the scope of this essay.

P.S. There are some wonderful quotes about how pragmatism is sterile, useless, and so on. I think adding to the confusion is that pragmatism was conceived as a formal-system-to-human-system tool. If you only use it in human endeavors it can certainly cause you to self-justify being a moron, but it's not used in that way here (or in that way by many of its founders, either) c.f.

> "Pragmatism is an intellectually safe but ultimately sterile philosophy"
