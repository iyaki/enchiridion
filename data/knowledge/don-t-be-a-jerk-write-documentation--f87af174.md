---
title: "Don't be a Jerk: Write Documentation"
notion_id: f87af174-6f3f-4dc2-a5e4-be1c7471e979
notion_url: https://app.notion.com/p/Don-t-be-a-Jerk-Write-Documentation-f87af1746f3f4dc2a5e4be1c7471e979
last_edited: 2023-04-25T13:37:00.000Z
source_url: https://ferd.ca/don-t-be-a-jerk-write-documentation.html
tags: ["English", "Programming", "Producer (Individual Contributor)", "Documentation", "Article", "My bad opinions (Fred Hebert)"]
---
[https://ferd.ca/don-t-be-a-jerk-write-documentation.html](https://ferd.ca/don-t-be-a-jerk-write-documentation.html)

Developers hate writing documentation. And generally, users hate developers who don't write documentation. There's a subset of users who dislike going through documentation and will prefer to go directly in the source for whatever they're doing, but they're certainly not a majority. My position is that if you do not have any kind of documentation available, you're a jerk.

### I don't have the time

The classic cop out to producing documentation is not having the time. Of course, it's entirely valid and possible that more urgent things take priority over documentation. However, developers who use time as an excuse should be aware that time they do not pour into documentation is time their users instead have to invest in it. To quote Scott Berkun:

> When 100 people are listening to you for an hour, that’s 100 hours of people’s time devoted to what you have to say. If you can’t spend 5 or 10 hours preparing for them, thinking about them, and refining your points to best suit their needs, what does that say about your respect for your audience’s time? It says your 5 hours are more important than 100 of theirs, which requires an ego larger than the entire solar system.

The original quote is about public speaking, but it's easy to translate in terms that are relevant to programming and software. Users give their time away to learn to use your libraries and your applications in the hopes it saves them time in the future. When they need to figure stuff out, their appreciation of your work is instantly diminished, and you'll have to take even more time for indirect support (which could frankly be easy to recycle into documentation).

Frankly, not having documentation is a net loss in terms of time for everyone involved. People will ask you questions that will be repeated dozens of times over e-mail, IRC, github issues, or in person. And every person who asks you something is likely someone who has went through all the hoops and asked other people before bothering you. You'll likely spend way more time explaining your application to many people in a one-by-one manner than you would had you just written some documentation in the first place.

### Just read the source

Telling a user to "just read the source" is a very jerk-ish thing to do. It's the "Why are you bothering me? Go away!" of programming. You're basically telling the user to reverse engineer whatever pile of crap you put out there (because most code out there, including mine, is likely a pile of crap) using their own time, while you're sitting there with the knowledge they need, just not willing to distribute it.

### Contributions are welcome

This one is probably the worst one. It says “My time is so much more important than yours that you should figure everything out yourself, and then reward me with the fruits of your labour while you’re at it.”

It's somewhat of a valid excuse when documentation currently exists and users are unhappy with it, but it is a jerk thing to say when nothing is available.

### Doc is hard to maintain

This is a fair point. My response to this one is usually similar to the one given to "I don't have the time". Good documentation, much like good code, requires a lot of effort and takes time. The more time it takes to write decent documentation, the more time it would likely take your users to figure things out on their own, and the more frustrated they should be.

There's also an interesting question that has to be asked related to this: if code changes too much to be documented without monumental effort, does it mean that the same code changes too much to be worth using by other people?

If the answer is "no", then nothing keeps you from documenting these unchanging parts, the one users would likely care about.

### Documentation is useless

Well I don't really have much to write here regarding people who think documentation is useless. The obvious thing to say is that misleading documentation can be worse than no documentation, but it's an argument that likely does not carry over with most other types of documentation.

There's a definite amount of people who care about documentation and find it generally useful (over 90% of Erlang users polled for high-level documentation -- comments in code score much lower, a bit above 75% See related poll results), and should not be entirely ignored.

### A few pointers to writing minimal but helpful docs

One thing I want to mention regarding documentation is what documentation is in the first place, or what its point is. At its simplest level, documentation is being able to transfer knowledge from one person to another one (and not having one person say "just reinvent all that information yourself"). What changes is how persistent the media by which it is done are, and how accessible they are.

Talking to someone, a chat session, an e-mail conversation or mailing lists are all different ways to transfer knowledge, but they're ultimately not accessible to third parties compared to other media. They're volatile and while the information transferred there is usually very clear and well-adapted to its target audience, it's lost to the rest of the world.

README files, diagrams, examples and whatnot are more durable, can be kept within the repository, and anyone with access to the source can get to the information. Heck, tests fit this category if they're clear enough in their intent. Websites, FAQs, and Wikis can be maintained and accessed independently of the source, and managed by different people than those who care deeply about the source. Diagrams and examples can come in here too. Then you get other media such as books, talks and slides, screencasts, or even a support line by phone. They're not searchable, but they're alternative means of documentation that are usually worth their weight in gold.

I'm not advocating one type of documentation as better than any other.Such a judgment would be circumstantial at best. The thing is, most people who have shitty or no docs at all for their project, at one point or another, end up explaining it to someone directly, either over chat, face to face, or over another private channel. This explanation is documentation, and the question that person asked is an example of a need for better documentation.

Now, if you only ever explain over volatile media or formats, such questions will need to be repeated over and over by different people, and the answers will need to be repeated over and over by you (or a few dedicated users). Everybody's losing time.

Take whatever volatile format you used, and persist it over to a more stable one. It will take a bit of work, but whatever you had, if it answered your user's question, should be decent documentation. Copy/paste chat logs, shell input/output, e-mail conversations. Extract the vital aspects, explain context, why you did a given things and you have a decent base, with very little effort. Understanding your user's problem and offering a solution was the hard part.

### Know your Audience

Not all documentation is equal. Not everyone seeks the same information, and catering to specific types of readers will go a long way to help make your documentation more focused and helpful. I generally categorize people into 3 groups:

- Newcomers
- Regular users
- Contributors

They're not strict and clearly defined categories. There's a continuum between all poles and a newbie in one area is a contributor or expert in another one of your system, but I've personally found them to be useful guidelines.

### Newcomers

For newcomers, the documentation needs to be similar to Journalism's five (or six) Ws and explain a few common points that everybody is going to wonder about at some point:

- What does it do?
- Why did you write it (or didn't use an existing solution)?
- Who should be using it, and for what reasons? (this may include a license)
- How should it be used (how to build it, how to configure it, how to run it and get started)?
- Where can additional information be found?
- NOT THE SOURCE

These are all things that will pretty much always need to be answered, and you might as well get it out of the way at first. Additional information is always nice, but if the first 4 points cannot be answered, the newcomer is missing vital information regarding whether they should use your software/library or not. Growing documentation from that point for newcomers might be simpler after this.

### Regular Users

Regular users are people who probably have an idea of what they're looking for in your application or library. They've used it for a while, know or suspect some functionality exists but need more information, or are encountering problems related to using your material and want to figure out what they're doing wrong.

For them, the following can be useful:

- Reference Manual
- Examples (also useful to newcomers)
- EDoc / JavaDoc / $LanguageDoc (also useful for contributors)
- Wiki, Websites, etc.
- API descriptions
- NOT THE SOURCE

They're the usual places where people can go to look things up, maybe search for an extension to the knowledge they already have. This is the type of documentation that is the hardest to maintain, by far. It's usually fairly technical, and when it's complete, documents nearly every behaviour your code base can have.

### Contributors

If you want people to give back some effort and contribute code (or documentation) back, you have to take their hand for a while. Usually, the first pull request you'll have may be worthy of comments, and will include things such as a discussion including programming style and other demands ("has it been tested?"). Documentation for contributors should predict a few of these questions and preemptively answer them. This includes:

- The project structure and architecture: where should I head into the source to find about some functionality? Where should new features go?
- The project's principles: there's a reason why things were written the way they were, and a contributor shouldn't break these principles inadvertently.
- Tests: where to find them, how the contributor should use them to make sure nothing broke, and what do you expect for code additions.
- Issues and roadmap. Some users just want to give back or are looking for a knowledge base to expand, a place to communicate with you with problems they encounter. Contributors should also be able to know whether what they're working on is going to become useless in the next iteration of your project.
- The repository (THE SOURCE!)

Note that the kind of users you get (newcomers, regulars, and/or contributors) may dictate the kind of documentation you should write, but that the opposite is also true: the kind of documentation you make available may bring or guide specific types of users to your project. If you want contributors and people to try your application in production and report back, you should pave the way for them to do it. This is larger than documentation, but documentation is a definitive part of it.

### Get Started

This is the part of the blog post that becomes way more anecdotal and difficult to apply to broad audiences. I'll be basing myself off personal experience here, and how I do things. I implicitly assume I'm usually writing decent documentation, but I frankly have no empirical evidence to prove I do so. If you're into cold hard facts, you may want to skip this section.

### Talking to yourself

There's no perfect way to get started, and to some extent, a blank documentation page is as intimidating as a blank page for an essay, or an empty repository for a software project you don't know how to tackle.

One approach I like for documentation is to pretend I'm having a discussion with a new user. This might be similar to rubber ducking or personas in marketing and design. The user has no idea what the code does, what the project does. He or she has a given level of experience in the language(s), platforms, and problem domains the code base deals with. He or she may have a given budget, organization size, or language level that may direct where things go.

I imagine the discussion I would have with that user (and it may be different if I initiate it or they initiate it). How does the conversation go? What questions do they ask?

This procedure alone usually helps me get started. By adding different users, different scenarios, and repeating the exercise, documentation gets to expand and be more comprehensive. It may also end up being split up into different categories for different people, or distributed over different platforms and media.

### Problems everywhere

Another very quick way for me to get started is to find a problem a user would have, then showing how to solve it. If the reader ends up having a similar problem, they'll possibly find that the documentation suits them perfectly at least once.

To find such problems, I ask myself what are going to be common pitfalls. It usually help to take a break from a project for a few days, come back to it (uninstall it and start from a fresh setup), and try to use it for a while. If you make a mistake, users who haven't written the code (and thus know it far less than you do) will likely have similar questions.

You can answer these questions by fixes to the program itself or documentation. When you have users, every question you get about your project becomes a potential piece of documentation.

### “Just RTFM”

Informal communications you have with users, errors you encounter, or procedures you have with your application on a day-to-day basis all have the potential to become very useful documentation.

When you have a bit of documentation, you won't have any guarantee it is good documentation (you likely never will), but hey, maybe you'll have earned the smug rights to tell your users to RTFM. Then when they don't, you may have to consider writing a better FM.
