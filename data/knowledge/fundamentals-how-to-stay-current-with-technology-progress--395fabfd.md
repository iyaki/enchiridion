---
title: "Fundamentals - How to stay current with technology progress"
notion_id: 395fabfd-36ee-4208-84cf-0d44b7e51483
notion_url: https://app.notion.com/p/Fundamentals-How-to-stay-current-with-technology-progress-395fabfd36ee420884cf0d44b7e51483
last_edited: 2024-06-05T18:33:00.000Z
source_url: https://blog.ploeh.dk/2024/05/20/fundamentals/
tags: ["Article", "ploeh blog", "English", "Programming", "Career Growth", "Learning"]
---
A long time ago, I landed my dream job. My new employer was a consulting company, and my role was to be the resident [Azure](https://en.wikipedia.org/wiki/Microsoft_Azure) expert. Cloud computing was still in its infancy, and there was a good chance that I might be able to establish myself as a leading regional authority on the topic.

As part of the role, I was supposed to write articles and give presentations showing how to solve various problems with Azure. I dug in with fervour, writing sample code bases and even [an MSDN Magazine article](http://msdn.microsoft.com/en-us/magazine/gg983487.aspx). To my surprise, after half a year I realized that I was bored.

At that time I'd already spent more than a decade learning new technology, and I knew that I was good at it. For instance, I worked five years for Microsoft Consulting Services, and a dirty little secret of that kind of role is that, although you're sold as an expert in some new technology, you're often only a few weeks ahead of your customer. For example, I was once engaged as a [Windows Workflow Foundation](https://en.wikipedia.org/wiki/Windows_Workflow_Foundation) expert at a time when it was still in beta. No-one had years of experience with that technology, but I was still expected to know much more about it than my customer.

I had lots of engagements like that, and they usually went well. I've always been good at cramming, and as a consultant you're also unencumbered by all the daily responsibilities and politics that often occupy the time and energy of regular employees. The point being that while I'm decent at learning new stuff, the role of being a consultant also facilitates that sort of activity.

After more then a decade of learning new frameworks, new software libraries, new programming languages, new tools, new online services, it turned out that I was ready for something else. After spending a few months learning Azure, I realized that I'd lost interest in that kind of learning. When investigating a new Azure SDK, I'd quickly come to the conclusion that, _oh, this is just another object-oriented library_. There are these objects, and you call this method to do that, etc. That's not to say that learning a specific technology is a trivial undertaking. The worse the design, the more difficult it is to learn.

Still, after years of learning new technologies, I'd started recognizing certain patterns. Perhaps, I thought, well-designed technologies are based on some fundamental ideas that may be worth learning instead.

### Staying current [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#ac37913a2b8248e6b51d4506c2da0481)

A common lament among software developers is that the pace of technology is so overwhelming that they can't keep up. This is true. You can't keep up.

There will always be something that you don't know. In fact, most things you don't know. This isn't a condition isolated only to technology. The sum total of all human knowledge is so vast that you can't know it all. What you will learn, even after a lifetime of diligent study, will be a nanoscopic fraction of all human knowledge - even of everything related to software development. You can't stay current. Get used to it.

A more appropriate question is: _How do I keep my skill set relevant?_

Assuming that you wish to stay employable in some capacity, it's natural to be concerned with how your mad [Flash](https://en.wikipedia.org/wiki/Adobe_Flash) skillz will land you the next gig.

Trying to keep abreast of all new technologies in your field is likely to lead to burnout. Rather, put yourself in a position so that you can quickly learn necessary skills, just in time.

### Study fundamentals, rather than specifics [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#c529c0131b284fe1bca42bec0663fc8e)

Those many years ago, I realized that it'd be a better investment of my time to study fundamentals. Often, once you have some foundational knowledge, you can apply it in many circumstances. Your general knowledge will enable you to get quickly up to speed with specific technologies.

Success isn't guaranteed, but knowing fundamentals increases your chances.

This may still seem too abstract. Which fundamentals should you learn?

In the remainder of this article, I'll give you some examples. The following collection of general programmer knowledge spans software engineering, computer science, broad ideas, but also specific tools. I only intend this set of examples to serve as inspiration. The list isn't complete, nor does it constitute a minimum of what you should learn.

If you have other interests, you may put together your own research programme. What follows here are just some examples of fundamentals that I've found useful during my career.

A criterion, however, for constituting foundational knowledge is that you should be able to apply that knowledge in a wide variety of contexts. The fundamental should not be tied to a particular programming language, platform, or operating system.

### Design patterns [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#4f474189809f4d53b447b4005cef1bfd)

Perhaps the first foundational notion that I personally encountered was that of _design patterns_. As the Gang of Four (GoF) wrote in [the book](https://en.wikipedia.org/wiki/Design_Patterns), a design pattern is an abstract description of a solution that has been observed 'in the wild', more than once, independently evolved.

Please pay attention to the causality. A design pattern isn't prescriptive, but descriptive. It's an observation that a particular code organisation tends to solve a particular problem.

There are lots of misconceptions related to design patterns. One of them is that the 'library of patterns' is finite, and more or less constrained to the patterns included in the original book.

There are, however, many more patterns. To illustrate how much wider this area is, here's a list of some patterns books in my personal library:

- [Design Patterns](https://blog.ploeh.dk/ref/dp)
- [Pattern Languages of Program Design 3](https://blog.ploeh.dk/ref/plopd3)
- [Patterns of Enterprise Application Architecture](https://blog.ploeh.dk/ref/peaa)
- [Enterprise Integration Patterns](https://blog.ploeh.dk/ref/eip)
- [xUnit Test Patterns](https://blog.ploeh.dk/ref/xunit-patterns)
- [Service Design Patterns](https://blog.ploeh.dk/ref/service-design-patterns)
- [Implementation Patterns](https://blog.ploeh.dk/ref/implementation-patterns)
- [RESTful Web Services Cookbook](https://blog.ploeh.dk/ref/rest-cookbook)
- [AntiPatterns](https://blog.ploeh.dk/ref/antipatterns)

In addition to these, there are many more books in my library that are patterns-adjacent, including [one of my own](https://blog.ploeh.dk/dippp). The point is that software design patterns is a vast topic, and it pays to know at least the most important ones.

A design pattern fits the criterion that you can apply the knowledge independently of technology. The original GoF book has examples in [C++](https://en.wikipedia.org/wiki/C%2B%2B) and [Smalltalk](https://en.wikipedia.org/wiki/Smalltalk), but I've found that they apply well to C#. Other people employ them in their [Java](https://www.java.com/) code.

Knowing design patterns not only helps you design solutions. That knowledge also enables you to recognize patterns in existing libraries and frameworks. It's this fundamental knowledge that makes it easier to learn new technologies.

Often (although not always) successful software libraries and frameworks tend to follow known patterns, so if you're aware of these patterns, it becomes easier to learn such technologies. Again, be aware of the causality involved. I'm not claiming that successful libraries are explicitly designed according to published design patterns. Rather, some libraries become successful because they offer good solutions to certain problems. It's not surprising if such a good solution falls into a pattern that other people have already observed and recorded. It's like [parallel evolution](https://en.wikipedia.org/wiki/Parallel_evolution).

This was my experience when I started to learn the details of Azure. Many of those SDKs and APIs manifested various design patterns, and once I'd recognized a pattern it became much easier to learn the rest.

The idea of design patterns, particularly object-oriented design patterns, have its detractors, too. Let's visit that as the next set of fundamental ideas.

### Functional programming abstractions [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#c44e7624ea3e4cef9485522146d17a6d)

As I'm writing this, yet another Twitter thread pokes fun at object-oriented design (OOD) patterns as being nothing but a published collection of workarounds for the shortcomings of object orientation. The people who most zealously pursue that agenda tends to be functional programmers.

Well, I certainly like functional programming (FP) better than OOD too, but rather than poking fun at OOD, I'm more interested in [how design patterns relate to universal abstractions](https://blog.ploeh.dk/2018/03/05/some-design-patterns-as-universal-abstractions). I also believe that FP has shortcomings of its own, but I'll have more to say about that in a future article.

Should you learn about [monoids](https://blog.ploeh.dk/2017/10/06/monoids), [functors](https://blog.ploeh.dk/2018/03/22/functors), [monads](https://blog.ploeh.dk/2022/03/28/monads), [catamorphisms](https://blog.ploeh.dk/2019/04/29/catamorphisms), and so on?

Yes you should, because these ideas also fit the criterion that the knowledge is technology-independent. I've used my knowledge of these topics in [Haskell](https://www.haskell.org/) (hardly surprising) and [F#](https://fsharp.org/), but also in C# and [Python](https://www.python.org/). The various [LINQ](https://en.wikipedia.org/wiki/Language_Integrated_Query) methods are really just well-known APIs associated with, you guessed it, functors, monads, monoids, and catamorphisms.

Once you've learned these fundamental ideas, it becomes easier to learn new technologies. This has happened to me multiple times, for example in contexts as diverse as property-based testing and asynchronous message-passing architectures. Once I realize that an API gives rise to a monad, say, I know that certain functions must be available. I also know how I should best compose larger code blocks from smaller ones.

Must you know all of these concepts before learning, say, F#? No, not at all. Rather, a language like F# is a great vehicle for learning such fundamentals. There's a first time for learning anything, and you need to start somewhere. Rather, the point is that once you know these concepts, it becomes easier to learn the next thing.

If, for example, you already know what a monad is when learning F#, picking up the idea behind [computation expressions](https://learn.microsoft.com/dotnet/fsharp/language-reference/computation-expressions) is easy once you realize that it's just a compiler-specific way to enable syntactic sugaring of monadic expressions. You can learn how computation expressions work without that knowledge, too; it's just harder.

This is a recurring theme with many of these examples. You can learn a particular technology without knowing the fundamentals, but you'll have to put in more time to do that.

On to the next example.

### SQL [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#02c8fb3f6fe74fb9bffda719122c60a9)

Which [object-relational mapper](https://en.wikipedia.org/wiki/Object%E2%80%93relational_mapping) (ORM) should you learn? [Hibernate](https://hibernate.org/orm/)? [Entity Framework](https://learn.microsoft.com/ef/)?

How about learning [SQL](https://en.wikipedia.org/wiki/SQL)? I learned SQL in 1999, I believe, and it's served me well ever since. I [consider raw SQL to be more productive than using an ORM](https://blog.ploeh.dk/2023/09/18/do-orms-reduce-the-need-for-mapping). Once more, SQL is largely technology-independent. While each database typically has its own SQL dialect, the fundamentals are the same. I'm most well-versed in the [SQL Server](https://en.wikipedia.org/wiki/Microsoft_SQL_Server) dialect, but I've also used my SQL knowledge to interact with [Oracle](https://www.oracle.com/database/) and [PostgreSQL](https://www.postgresql.org/). Once you know one SQL dialect, you can quickly solve data problems in one of the other dialects.

It doesn't matter much whether you're interacting with a database from .NET, Haskell, Python, [Ruby](https://www.ruby-lang.org/), or another language. SQL is not only universal, the core of the language is stable. What I learned in 1999 is still useful today. Can you say the same about your current ORM?

Most programmers prefer learning the newest, most cutting-edge technology, but that's a risky gamble. Once upon a time [Silverlight](https://en.wikipedia.org/wiki/Microsoft_Silverlight) was a cutting-edge technology, and more than one of my contemporaries went all-in on it.

On the contrary, most programmers find old stuff boring. It turns out, though, that it may be worthwhile learning some old technologies like SQL. Be aware of the [Lindy effect](https://en.wikipedia.org/wiki/Lindy_effect). If it's been around for a long time, it's likely to still be around for a long time. This is true for the next example as well.

### HTTP [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#e4a7c033c0964420a0abbf83a0bbb773)

The [HTTP protocol](https://en.wikipedia.org/wiki/HTTP) has been around since 1991. It's an effectively text-based protocol, and you can easily engage with a web server on a near-protocol level. This is true for other older protocols as well.

In my first IT job in the late 1990s, one of my tasks was to set up and maintain [Exchange Servers](https://en.wikipedia.org/wiki/Microsoft_Exchange_Server). It was also my responsibility to make sure that email could flow not only within the organization, but that we could exchange email with the rest of the internet. In order to test my mail servers, I would often just [telnet](https://en.wikipedia.org/wiki/Telnet) into them on port 25 and type in the correct, [text-based instructions to send a test email](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol).

Granted, it's not that easy to telnet into a modern web server on port 80, but a ubiquitous tool like [curl](https://curl.se/) accomplishes the same goal. I recently wrote how [knowing curl is better](https://blog.ploeh.dk/2024/05/13/gratification) than knowing [Postman](https://www.postman.com/). While this wasn't meant as an attack on Postman specifically, neither was it meant as a facile claim that curl is the only tool useful for ad-hoc interaction with HTTP-based APIs. Sometimes you only realize an underlying truth when you write about a thing and then [other people find fault with your argument](https://blog.ploeh.dk/2024/05/13/gratification#9efea1cadb8c4e388bfba1a2064dd59a). The underlying truth, I think, is that it pays to understand HTTP and being able to engage with an HTTP-based web service at that level of abstraction.

Preferably in an automatable way.

### Shells and scripting [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#e3a250b707b243dabc6609134e864aee)

The reason I favour curl over other tools to interact with HTTP is that I already spend quite a bit of time at the command line. I typically have a little handful of terminal windows open on my laptop. If I need to test an HTTP server, curl is already available.

Many years ago, an employer introduced me to [Git](https://git-scm.com/). Back then, there were no good graphical tools to interact with Git, so I had to learn to use it from the command line. I'm eternally grateful that it turned out that way. I still use Git from the command line.

When you install Git, by default you also install Git Bash. Since I was already using that shell to interact with Git, it began to dawn on me that it's a full-fledged shell, and that I could do all sorts of other things with it. It also struck me that learning [Bash](https://www.gnu.org/software/bash/) would be a better investment of my time than learning [PowerShell](https://learn.microsoft.com/powershell/). At the time, there was no indication that PowerShell would ever be relevant outside of Windows, while Bash was already available on most systems. Even today, knowing Bash strikes me as more useful than knowing PowerShell.

It's not that I do much Bash-scripting, but I could. Since I'm a programmer, if I need to automate something, I naturally reach for something more robust than shell scripting. Still, it gives me confidence to know that, since I already know Bash, Git, curl, etc., I _could_ automate some tasks if I needed to.

Many a reader will probably complain that the Git CLI has horrible [developer experience](https://blog.ploeh.dk/2024/05/13/gratification), but I will, again, postulate that it's not that bad. It helps if you understand some fundamentals.

### Algorithms and data structures [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#a511cfd8d9bf4bbda433dbf70184284a)

Git really isn't that difficult to understand once you realize that a Git repository is just a [directed acyclic graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph) (DAG), and that branches are just labels that point to nodes in the graph. There are basic data structures that it's just useful to know. DAGs, [trees](https://en.wikipedia.org/wiki/Tree_(graph_theory)), [graphs](https://en.wikipedia.org/wiki/Graph_(discrete_mathematics)) in general, [adjacency lists](https://en.wikipedia.org/wiki/Adjacency_list) or [adjacency matrices](https://en.wikipedia.org/wiki/Adjacency_matrix).

Knowing that such data structures exist is, however, not that useful if you don't know what you can _do_ with them. If you have a graph, you can find a [minimum spanning tree](https://en.wikipedia.org/wiki/Minimum_spanning_tree) or a [shortest-path tree](https://en.wikipedia.org/wiki/Shortest-path_tree), which sometimes turn out to be useful. Adjacency lists or matrices give you ways to represent graphs in code, which is why they are useful.

Contrary to certain infamous interview practices, you don't need to know these algorithms by heart. It's usually enough to know that they exist. I can't remember [Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm) off the top of my head, but if I encounter a problem where I need to find the shortest path, I can look it up.

Or, if presented with the problem of constructing current state from an Event Store, you may realize that it's just a left [fold](https://en.wikipedia.org/wiki/Fold_(higher-order_function)) over a [linked list](https://en.wikipedia.org/wiki/Linked_list). (This isn't my own realization; I first heard it from [Greg Young in 2011](https://gotocon.com/cph-2011/presentation/Behavior!).)

Now we're back at one of the first examples, that of FP knowledge. A [list fold is its catamorphism](https://blog.ploeh.dk/2019/05/27/list-catamorphism). Again, these things are much easier to learn if you already know some fundamentals.

### What to learn [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#f109425f27014cd5bd395a74e9575355)

These examples may seems overwhelming. Do you really need to know all of that before things become easier?

No, that's not the point. I didn't start out knowing all these things, and some of them, I'm still not very good at. The point is rather that if you're wondering how to invest your limited time so that you can remain up to date, consider pursuing general-purpose knowledge rather than learning a specific technology.

Of course, if your employer asks you to use a particular library or programming language, you need to study _that_, if you're not already good at it. If, on the other hand, you decide to better yourself, you can choose what to learn next.

Ultimately, if your're learning for your own sake, the most important criterion may be: Choose something that interests you. If no-one forces you to study, it's too easy to give up if you lose interest.

If, however, you have the choice between learning [Noun.js](https://mjvl.github.io/Noun.js/) or design patterns, may I suggest the latter?

### For life [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#94c3f380b556403d82dd9f3cd0c1d1e9)

When are you done, you ask?

Never. There's more stuff than you can learn in a lifetime. I've met a lot of programmers who finally give up on the grind to keep up, and instead become managers.

As if there's nothing to learn when you're a manager. I'm fortunate that, before [I went solo](https://blog.ploeh.dk/2011/11/08/Independency), I mainly had good managers. I'm under no illusion that they automatically became good managers. All I've heard said about management is that there's a lot to learn in that field, too. Really, it'd be surprising if that wasn't the case.

I can understand, however, how just keep learning the next library, the next framework, the next tool becomes tiring. As I've already outlined, I hit that wall more than a decade ago.

On the other hand, there are so many wonderful fundamentals that you can learn. You can do self-study, or you can enrol in a more formal programme if you have the opportunity. I'm currently following a course on compiler design. It's not that I expect to pivot to writing compilers for the rest of my career, but rather,

> 

1. "It is considered a topic that you should know in order to be "well-cultured" in computer science.
2. "A good craftsman should know his tools, and compilers are important tools for programmers and computer scientists.
3. "The techniques used for constructing a compiler are useful for other purposes as well.
4. "There is a good chance that a programmer or computer scientist will need to write a compiler or interpreter for a domain-specific language."

[Introduction to Compiler Design](https://blog.ploeh.dk/ref/introduction-to-compiler-design)

That's good enough for me, and so far, I'm enjoying the course (although it's also hard work).

You may not find this particular topic interesting, but then hopefully you can find something else that you fancy. 3D rendering? Machine learning? Distributed systems architecture?

### Conclusion [#](https://blog.ploeh.dk/2024/05/20/fundamentals/?utm_source=tldrnewsletter#7519e9b6147d49379f545c69871c381a)

Technology moves at a pace with which it's impossible to keep up. It's not just you who's falling behind. Everyone is. Even the best-paid [GAMMA](https://en.wikipedia.org/wiki/Big_Tech) programmer knows next to nothing of all there is to know in the field. They may have superior skills in certain areas, but there will be so much other stuff that they don't know.

You may think of me as a [thought leader](https://x.com/hillelogram/status/1445435617047990273) if you will. If nothing else, I tend to be a prolific writer. Perhaps you even think I'm a good programmer. I should hope so. Who fancies themselves bad at something?

You should, however, have seen me struggle with [C](https://en.wikipedia.org/wiki/C_(programming_language)) programming during a course on computer systems programming. There's a thing I'm happy if I never have to revisit.

You can't know it all. You can't keep up. But you can focus on learning the fundamentals. That tends to make it easier to learn specific technologies that build on those foundations.

### Wish to comment?

You can add a comment to this post by [sending me a pull request](https://github.com/ploeh/ploeh.github.com#readme). Alternatively, you can discuss this post on Twitter or somewhere else with a permalink. Ping me with the link, and I may respond.

![image](https://blog.ploeh.dk/assets/themes/ploeh/images/Portrait15.jpg)

Mark Seemann
