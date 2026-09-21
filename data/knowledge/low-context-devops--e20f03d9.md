---
title: "Low-Context DevOps"
notion_id: e20f03d9-20b1-4c30-a1f1-2a24f8f286a2
notion_url: https://app.notion.com/p/Low-Context-DevOps-e20f03d920b14c30a1f12a24f8f286a2
last_edited: 2024-01-30T14:11:00.000Z
source_url: https://www.usenix.org/publications/loginonline/low-context-devops#:~:text=A%20low%2Dcontext%20environment%20is,RC)%20needed%20to%20be%20successful.
tags: ["Article", "USENIX | The Advanced Computing Systems Association", "English", "DevOps", "Productivity"]
---
How often is your work blocked as you wait for an answer from a coworker, spend hours researching how to do something that should be trivial, or waste days on a task only to later find out there was an easier way? These are red-flags that you work in a high-context environment and you lack context. Let’s fix this!

A **high-context environment** is one where a lot of background knowledge (context) is required for a given task. It is full of trips and traps that a new person can fall into. A **low-context environment** is the opposite: Little context is required to be productive; what context is needed is presented to you at the right place and time.

**Low-Context DevOps** is my term for the aspirational goal of creating a work environment that minimizes the **required context (RC)** needed to be successful. In such an environment people are more productive, there is less frustration, higher morale.

Left unchecked, the amount of RC for an environment tends to grow over time and can make any environment insufferable. It must be managed like technical debt: Start off right, permit regressions only reluctantly, periodically create strategic projects to radically reduce it.

In this article I begin with an explanation of the sociology concepts, then focus on five ways to lower context, or leverage low-context techniques.

Keeping these techniques in mind, we can create SRE environments that are less frustrating and more productive.

High and Low Context Environments

High-context culture and low-context culture are extremes in a continuum of how explicit or implicit communication is in a culture. The concept was first introduced in 1959 by the anthropologist Edward T. Hall.

A tight-knit group of friends is a high-context culture. They have in-jokes only understood by those who "had been there." They have traditions and non-obvious ways of doing things that have been refined over the years. They know how to read between the lines: what is said is less important than what is implied.

Contrast this to a low-context environment such as an airport. You'll only be there for a short duration, so there is no expectation that you'll have years of accumulated context and "just know" where things are. What to do, where to do it, and how to do it, is all spelled out clearly and there are information desks and guides ready to assist if you still have questions. The signage is very explicit. Rules are on display for everyone to see. The location of each airline's check-in desk is announced on signage visible from your car, before you even arrive!

### High Context Cultures:

- Communication is implicit.
- Less written/formal information; more collective history
- People have to “read between the line” to understand what's going on
- Relies on long term relationships
- Decisions and activities focus around personal face-to-face relationships, often around a central person who has authority
- Examples: A party with friends, family gatherings, expensive gourmet restaurants with a regular clientele, undergraduate on-campus friendships, hosting a friend in your home overnight

### Low Context Cultures:

- Communication is explicit
- There are rules, you are told the rules
- Knowledge tends to be codified, public, external, and accessible
- More interpersonal connections of shorter duration
- Knowledge is more often transferable
- Examples: Large airports, a chain supermarket, a cafeteria, a convenience store, sports where rules are clearly laid out, a motel [[1]](https://www.usenix.org/publications/loginonline/low-context-devops#ref1)

**Required content (RC)** is the amount of context required to be successful at a given task in a particular environment. Performing a task requires certain skills and knowledge. However the amount of RC depends on the environment, not the task itself.

For example the task of making toast requires very little skill and knowledge: Put bread in the toaster, press a button, and wait. If you are visiting my house, making toast has a low RC: The bread is in a breadbox next to the toaster; the cabinets have glass doors making it easy to find a plate. However the same task at a friend’s house has a higher RC: Finding a clean plate is a challenge and you need to know that the toaster and microwave oven can’t be used at the same time without triggering the circuit breaker. How could anyone be expected to know all that?

Low Context DevOps

**Low-Context DevOps** is my term for the aspirational goal of creating a work environment that minimizes the required context needed to be successful. In such an environment people are more productive, there is less frustration, morale is higher.

I would rather work in a low-context environment. I hope you agree!

I want to spend more time working, less time blocked and frustrated by roadblocks and information gaps. If I need to know something, I want that information to be accessible and easy to find. In fact, the hyperlink to that information should magically appear before my eyes at the right time and place. I shouldn't have to seek it out.

Being low-context is not something we can "bolt on at the end". We must employ low-context thinking from the start. The context required grows as the system evolves and becomes more complex. We need to be in a constant journey to reduce RC. Occasionally we must institute strategic projects to reign in RC.

In these next sections I’ll explore:

- Fixing the most common high-context environment: New employee onboarding.
- Using low-context thinking to encourage compliance with policies or recommended practices
- Designing systems so they are low-context from the start
- Using ubiquitous documentation to reign in or prevent increases in RC

### New Employee Context

How can we reduce the RC that stands between a new employee and being productive?

A new employee, by definition, lacks context. Being a new employee can feel confusing, disorienting, and frustrating. Assimilating into a high-context environment is frustrating for the person and inefficient for the company.

I conducted an informal survey of friends, asking them how long after being hired did they have the tools (PC/laptop and software) and access (ability to connect, and the proper permissions) to do their job. It varied by industry but multiple months was typical, with the banking industry being the worst at 6 months minimum.

There's an obvious inefficiency when new employees sit idle blocked from being able to work. What percentage of your budget goes down the drain as a result? How much more profitable would companies be without such waste?

This is not just about new employees. Once hired, employees typically change projects frequently; each time they become "the new person". This trend is increasing over time.

Inefficiency aside, it is also demotivating. Enthusiasm is at a peak when you first join a company or a project. Shouldn't that enthusiasm be leveraged to do productive work? Instead we squander it on fighting a bureaucracy for 6 months begging to be added to the right LDAP group so you can access the test database. No, that group didn't work. What about this group? No? Ok, I'll ask around and see if anyone can figure out what's wrong.

I reframe the problem as follows: We take people with the least context and expect them to guess their way through creating their workplace. We expect people that don’t yet fully understand the job to pick their tools. We take people that are, by definition, the least skilled at navigating our corporate systems and expect them to figure out what they need and how to get it.

The solution is easy to explain, difficult to implement: New employee's should be given the environment they need, not expected to create it. This includes their workstation (a PC or laptop, fully configured and ready to use), access to the systems they need to do their job, permissions on said systems specific to their responsibilities.

Doing this properly lowers the bar for new employees. It virtually eliminates the RC needed to become a productive part of their team.

Reducing the context this way is difficult. New employees can't fix this problem. They're new and chiefly powerless to make improvements. Experienced employees don't feel the pain, thus they've lost the motivation to fix the problem. Those that do try to fix this problem quickly realize how difficult it is to solve. Even the smallest obvious change requires working across silos that aren't used to having to communicate: Human Resources, IT, InfoSec, Engineering, and possibly others! Finding the right people is difficult enough. Getting them all in the same room at the same time to agree that this is a problem is harder. Actually working though all the politics, inertia, and red tape isn't the kind of thing taught at University. I majored in computer science, not hostage negotiation.

I won’t dig too deeply into this issue, as Gene Kim’s new book, The Unicorn Project, covers these issues so well. The book illuminates the source of the problem and offers strategic and tacitcal solutions. Like his previous book The Phoenix Project, it is told as a story.

### Encouraging Specific Practices

People comply with policies better when we reduce the context required to follow such policies.

There are plenty of behaviors we wish to encourage: Recommended practices in coding, design, testing, operations, and so on. As SREs we often must enforce rules, such as security or regulatory requirements.

It is better to get people to adopt a new behavior by making compliance easier than non-compliance. More carrot, less stick. Don't muck up the old ways to make them more difficult, provide systems that make the new way easier.

This is known as "make right easy" or "the lazy path should be the desired path." However my coworker Jamey Turcic said it best when he said he prefers to "trick his coworkers so that they fall into the pit of success."

Often non-compliance is due to a lack of context. The person didn’t know a rule existed, or they tried to comply but were unsuccessful, so they reverted to old behaviors.

This reminds me of an old story. A sysadmin was complaining that his users were wasting paper by not printing on both sides of paper. “Don’t they know how to select duplex when they print?” He was taking for granted that his context wasn’t shared universally. He was spending his time trying to teach people to remember how to select duplex printing instead of reducing the context people need to be successful. Eventually he set the default to be duplex globally and required extra effort (and context) to print 1-sided. In other words, he made the lazy way the right way.

Maybe your engineering department would like to require moving to continuous integration (CI) for all software projects. You could chastise teams that don't use CI, but that wouldn't be very effective. More effective would be to provide a build farm that compiles code faster than their existing solution. Developers hate to wait for compiles to complete. Oh, and if they end up with a CI system that embodies all the recommended practices, sure, that's nice too.

I was impressed when Google moved their entire (at the time) 10,000-person engineering organization to CI this way. The fastest, easiest, way to build even the smallest "Hello, World!" program is to use their new CI system (known as "Blaze"), and you get all that CI goodness, including fancy dashboards, metrics, and so on, for free. These features, on by default, made it easier to comply with engineering goals and standards.

This also relates to corporate policies. We can't expect people to know every policy. Penalizing them for not knowing every new policy is unfair. However if our systems make it easy to comply by default, people will comply without even realizing it.

### The Systems and Tools We Build

The systems we build should be designed with low-context in mind. People without context should be guided to the recommended practices.

Companies that maintain superior corporate culture do so by providing the tools that sustain and encourage that culture. Rules and norms fade into the background because it's "just how things are done." Culture is all the things that we do without realizing it. I didn't realize that a past employer's culture was so deeply data-driven until I left. We had powerful tools at our fingertips so being data-driven was just a part of everything. At my next employer trying to explain it was as difficult as a fish trying to explain what water feels like.

The systems we build and the tools we use should have low-context design built in so that the recommended practices we want to encourage fade into the background. How could you not do CI the way we recommend, with the code-cover testing, security fuzzing, artifact CVE analysis, push-button deployment, and so on? The lazy path should enable all that and more.

We can use low-context techniques to encourage or enforce practices in all of our foundational tools and infrastructure such as:

- Ticket systems, bug tracking systems
- Monitoring / observability systems
- Configuration management systems
- OS installation and patching
- CI/CD pipeline systems
- Container / artifact repository systems
- Source code control systems
- Chat and collaboration infrastructure

Low-context design can not be "bolted on at the end". Instead we must keep this in mind from the start. The key is to think deeply about the user's experience or path. A new user follows a certain path: They decide if they should be using the system, they get set up to use it their first time, and they continue to use it. Product Managers may refer to this as "personas". We might name these personas the Decider, The New User, The Experienced User. Identify what context those personas need and design that need out of the system.

Design is iterative. Early in my career I thought a design was done when it covered all the functional requirements. Eventually I learned that being feature-complete is just the start. We must then keep iterating on the design, refining it to improve the non-functional requirements such as achieving our goal of being low-context.

There are a myriad ways to achieve a low-context design but I find that a few simple techniques go a long way:

### Carefully constructed defaults:

Be thoughtful when deciding defaults. Favor the recommended path. Consider the most common paths taken by personas. Favor safety. For example, default to the dev database and require users to opt-in, or change a default to use the production database.

### Templates:

Templates hide complexity and focus the user on what's important. It is great when powerful and flexible CI/CD systems allow you to do anything anyway you want. What's better than that? A template that lets you not have to care about all that flexibility. People want to be guided by a trusted advisor. Users prefer to only have to specify what's unique about their situation. Sure, the template should be created by someone that will put time and thought into the best way to do something. Everyone else is willing to just follow along.

An easy way to make a template is to find a good example and strip it down. A template for a wiki page or design doc is often easiest to create by taking an exemplary document and stripping out everything except the headings. A template that drives a CI/CD system might start by comparing two working examples and parameterizing the differences.

### Convention over configuration:

Convention over configuration means a user should only have to specify the unconventional aspects of a configuration. The concept was introduced by David Heinemeier Hansson to describe attempts to decrease the number of decisions that a developer using a framework is required to make. When we lack context, each decision we need to make is a painful burden.

I recently had an experience that reminded me of the power of Convention Over Configuration.

We were replacing an older system that was super flexible. For each project we could specify the database name, which users had access, which users were admins, the names of cloud resources it accessed, where the log files were stored, and so on and so on. There must have been 100 or so very flexible settings.

All that flexibly was... terrible. The system was a tangled mess of unique snowflakes. Inconsistent naming leads to frequent mistakes and typos. Testing was difficult because the number of combinations to be tested was out of control. The required context was very high... how could you know that one project is named differently because of that thing that happened before you were hired?

The new system (we'll pretend it is called "thing") was designed with Convention Over Configuration in mind. For a project FOO, the database was called fooDB, the users had access if they were in LDAP group thing-foo-users, the admins were listed in LDAP group thing-foo-users-admins, the cloud resources had names prefixed with foo, the log files were in /var/log/thing/foo. A person with no context would easily guess which logs were in /var/log/thing/margaret.

### The Right Information at the Right Time and Place

We can reduce context but we can't eliminate it. Therefore we need ubiquitous documentation. Documentation is ubiquitous when it provides the right information, the right amount of information, where we need it, when we need it.

I saw an example of well-designed documentation on a recent visit to Paddington Station in London. Rather than overhead signs directing you to the taxi stand, there is a black line on the floor with the word "taxi" repeated every few meters along with arrows all pointing the same direction. You simply walk along the line, following the arrow, and soon you are at the taxi stand. It is the right amount (one word plus an arrow), at the right place and time (at your feet when you get off a train).

Another example is when Apple released macOS Catalina and changed the default Terminal shell to Zsh. If you are still using Bash you receive a warning message with the command needed to change your default, and a URL for more information. Apple could have just left people to Google-search for help. Providing a specific URL was a small effort with a huge effect.

A deep-link URL that brings you to the exact information the user needs is preferred over a link to the main page of a large corpus of text that the user then must navigate to to find the information they want.

Such links should be placed strategically in error messages, dashboards, portals, and anywhere people may desire additional context.

Documentation doesn't just magically appear. Management must set expectations around documentation quantity, quality, and timeliness.

For example, managers should set expectations that documentation is written as part of the project, not as an afterthought. Discourage people from providing work estimates like, "5 days plus 1 day to update the documentation." That's a 6-day project. Don't treat documentation as an "extra" that can be cut.

When we save documentation for the end of the project, we create an opportunity to skip it. Instead, encourage people to write documentation along the way.

### Continuous Improvement

Left unchecked, RC increases. We must be in a constant battle to reduce RC otherwise it gets out of control. Left unchecked, it becomes increasingly difficult and more laborious to be successful in that environment.

One way to win this battle is to update documentation in small increments. Writing in small increments is very powerful. From the author’s perspective, the writing is easier because the information is fresher in their mind. From the consumer’s perspective, the documentation is more up to date, as there is less lag before the new knowledge makes its way into the document.

I would rather see SREs ship a few paragraphs each day than wait a year to receive a large tome. This means people need to get comfortable with other people seeing their drafts. One update to the wiki might be a draft, the next update might refine it; this is ok if you label information was being “draft quality” so others are forewarned.

Lastly, incremental documentation is documentation that is more likely to be written. I’ve heard people say that they’ll document something “when things slow down” or “when they can find a solid chunk of time to write”. Well, that mythical “slow time” is never going to come. This is not a writing strategy, it is a procrastination technique.

Instead I encourage people to write in small batches: Update a sentence here, add an example there, revise an example into a parameterized bash one-liner, etc. Every time I close a ticket or merge a PR, I pause to update the related documentation even if it is just a single sentence. I might not know “the way” to do something, but I can list how it was done last time and include a note that this is “for reference only”.

Every time I close a ticket I ask myself what can I add to the wiki so that the next person can do that kind of ticket faster, or without having to poke around and guess what needs to be done?

That person might be a co-worker but it might be "me, six months from now." I don't know if you've met "me, six months from now" but he's a great guy. Just about as smart as "me, now" but can't remember as many details. Most of the documentation I write is really a note for "me, six months from now" or my other friend "me, when I get paged at 4am". That guy, gosh, he needs help!

Another motivation for writing documentation is that it makes it easier to take vacations. Engineers constantly put off taking vacations out of fear that the company can't survive without them. Lack of relaxation time leads to low morale and eventually burn-out. Ironically when someone burns out they often leave the company, and their 2-weeks notice is spent writing documentation. That mythical week of time to write finally appears!

Well-maintained documentation also makes it easier to change projects. A person recently asked me why a particular engineer is allowed to hop to the most exciting projects while other people are stuck doing the same thing all the time. Opportunities are easier to grab when your documentation is always complete.

Another barrier to writing is called the blank screen syndrome. Staring at a blank screen trying to decide what to write is intimidating: Where should I start? What's the scope? How much or little do I need to write? Who's the audience? Will the reader need detailed hand-holding or can I assume a certain level of experience?

Faced with all those questions it is easier to find some other project to work on. Problem solved!

My two favorite ways to solve the blank screen syndrome is templates and repurposing text.

I like to have a template for every situation. For example, at Stack Overflow every internal service has a "service doc" which lists basic facts, common operational tasks, and so on. The service docs are relatively consistent since they all start with the same template.

The blank screen syndrome is avoided because you start with a template, not a blank screen. You fill in the sections that apply and leave everything else blank. Often the first iteration is just filling in the title. That's ok! Better to have a good start that other people can fill in than never get started at all.

I also like to repurpose text. By that I mean I find where people are already writing and encourage them to take their excellent words and memorialize them on the wiki.

For example, when someone replies to an email with a long explanation, I pause and give them a compliment: What a great write-up! Thank you! Then I go further: This is so good, could you take a moment to paste that into our wiki?

Many people say they dislike writing but are very comfortable writing long descriptions in emails, in chat systems, on Stack Overflow.

Psychologically these systems don't feel like "writing documentation". It's just answering a question. The blank screen syndrome is avoided because "it's just answering someone's question". Inherent to the question is where to start, the scope, and the audience.

I'm involved in Stack Overflow for Teams, which is like stackoverflow.com but it gives a private, secure, place for an organization to maintain their own corpus of questions and answers. I find that many engineers who strongly dislike writing documentation will gladly answer questions from a coworker. I have a feed of questions tagged with keywords relevant to my job and I race to answer questions faster than my coworkers. Many company's inner source initiatives are enhanced by registering tags for various internal project names. Supporting those internal projects becomes gamified as people race to respond to questions tagged with their project's name.

Who Will Make This Happen?

I hope I have convinced you to think about the high-context and low-context cultures and how you can lower the required context of your engineering environment.

I hope you don't wait for some manager to come along and make these changes. Change isn't made by managers. It is made by leaders. While not everyone is (or wants to be) a manager, everyone can be a leader.

You can lead by taking the time to think out the personas for your next project, and basing the design around low-context ideals. When other people notice, show them what they did. People want to copy people that are successful.

We can lead by going first and making it easier for others to follow. You might be the only person on your team with the ability to build that new shiny new CI/CD system, but the right "getting started" guide makes it easy for others to follow in your footsteps.

You can lead by not waiting for permission. Create templates for your wiki and share them with everyone. If people see you using it, they'll follow. You don't need permission.

One of my favorite high-context environments is New York Penn Station. There's a sign that I walk by frequently directing people to the 7th Ave Subway. To understand that sign requires context: In particular, the knowledge that the 7th Ave Line hasn't been called that for 30 years [[2]](https://www.usenix.org/publications/loginonline/low-context-devops#ref2). What fascinates me about this sign is that it is back-lit by a lightbulb. Lightbulbs don't last 30 years. Someone has been changing that lightbulb for decades and is not empowered to update the sign.

Be a leader. Stop changing the lightbulb and start fixing the sign!

References:

[1] [http://www.culture-at-work.com/highlow.html](http://www.culture-at-work.com/highlow.html)

[2] [https://en.wikipedia.org/wiki/IRT_Broadway%E2%80%93Seventh_Avenue_Line](https://en.wikipedia.org/wiki/IRT_Broadway%E2%80%93Seventh_Avenue_Line)

Article Categories:

SRE

Programming

Last updated February 8, 2023

Authors:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Tom an internationally recognized author, speaker, system administrator, and DevOps advocate. He manages the SRE teams at Stack Overflow, Inc, and previously worked at Google, Bell Labs/Lucent, AT&T, and others. His books include Time Management for System Administrators (O'Reilly), The Practice of System and Network Administration (3rd edition), and The Practice of Cloud System Administration. In 2005, he received the USENIX SAGE Outstanding Achievement Award.

- [Log in](https://www.usenix.org/user/login?destination=node%2F273653%23comment-form) or [Register](https://www.usenix.org/user/register?destination=node%2F273653%23comment-form) to post comments
