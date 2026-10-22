---
title: "A Flexible Framework for Effective Pair Programming (2023)"
notion_id: 1d481607-93fa-405e-9a01-e4f14988adfd
notion_url: https://app.notion.com/p/A-Flexible-Framework-for-Effective-Pair-Programming-2023-1d48160793fa405e9a01e4f14988adfd
last_edited: 2023-01-27T02:09:00.000Z
source_url: https://shopify.engineering/a-flexible-framework-for-effective-pair-programming
tags: ["Programming", "Productivity", "Article", "Guide", "Shopify Engineering", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[Pair programming](https://shopify.engineering/pair-programming-explained) is one of the most important tools we use while mentoring early talent in the [Dev Degree](https://devdegree.ca/) program. It’s an agile software development technique where two people work together, either to share context, solve a problem, or learn from one another. Pairing builds technical and communication skills, encourages curiosity and creative problem-solving, and brings people closer together as teammates.

In my role as a Technical Educator, I’m focused on setting new interns joining the Dev Degree program up for success in their first 8 months at Shopify. Because pair programming is a method we use so frequently in onboarding, I saw an opportunity to streamline the process to make it more approachable for people who might not have experienced it before. I developed this framework during a live workshop I hosted at RenderATL. I hope it helps you structure your next pair programming session!

# Pair Programming in Dev Degree

_“Pair programming was my favorite weekly activity while on the Training Path. When I first joined my team, I was constantly pairing with my teammates. My experiences on the Training Path made these new situations manageable and incredibly fun! It was also a great way to socialize and work technically with peers outside of my school cohort that I wouldn't talk to often. I've made some good friends just working on a little project for the weekly module.” - Mikail Karimi, Dev Degree Intern_

In the [Dev Degree program](https://shopify.engineering/dev-degree-a-big-bet-on-software-education), we use mentorship to build up the interns’ technical and soft skills. As interns are usually early in their career journey, having recently graduated from high school or switched careers, their needs differ from someone with more career experience. Mentorship from experienced developers is crucial to prepare interns for their first development team placement. It shortens the time it takes to start making a positive impact with their work by developing their technical and soft skills like problem solving and communication. This is especially important now that Shopify is [digital by design](https://www.shopify.ca/careers/work-anywhere), as learning and working remotely is a completely new experience to many interns.

All first year Dev Degree interns at Shopify go through an eight-month period known as the _Training Path_. During this period, we deliver a set of courses designed to teach them all about Shopify-specific development. They’re mentored by Student Success Specialists, who coach them on building their soft skills like communication, and Technical Instructors , who focus on the technical aspects of the training. Pair programming with a peer or mentor is a great way to support both of these areas of development.

Each week, we allocate two to three hours for interns to pair program with each other on a problem. We don’t expect them to solve the problem completely, but they should use the concepts they learned from the week to hone their technical craft.

We also set up a bi-weekly 30 minute pair programming sessions with each intern. The purpose of these sessions is to provide dedicated one-on-one time to learn and work directly with an instructor. They can share what they are having trouble with, and we help them work through it.

_“When I’m switching teams and disciplines, pair programming with my new team is extremely helpful to see what resources people use to debug, the internal tools they use to find information and how they approach a problem. On my current placement, I got better at resolving problems independently when I saw how my mentor handled a new problem neither of us had seen.” Sanaa Syed, Dev Degree Intern_

As we scale up the program, there are some important questions I keep returning to:

- How do we track their progress most effectively?
- How do we know what they want to pair on each day?
- How can we provide a safe space?
- What are some best practices for communicating?

I started working on a framework to help solve these issues. I know I’m not the only one on my team who may be asking themselves this. Along the way, an opportunity arose to do a workshop at RenderATL. At Shopify, we are encouraged to learn as part of our professional development. Wanting to level up my public speaking skills, I decided to talk about mentorship through a pair programming lens. As the framework was nearly completed, I decided to crowdsource and finish the framework together with the RenderATL attendees.

# Crowdsourcing a Framework for All

On June 1st, 2022, Shopify hosted free all-day workshops at RenderATL called _Heavy Hitting React at Shopify Engineering_. It contained five different workshops, covering a range of topics from specific technical skills like React Native to broader skills like communication. We received a lot of positive feedback, met many amazing folks, and made sure those who attended gained new knowledge or skills they could walk away with.

For my workshop, _Let’s Pair Program a Framework Together_, we pair programmed a pair programming framework. The goal was to crowdsource and finish the pair programming framework I was working on based on the questions I mentioned above. We had over 30 attendees, and the session was structured to be interactive. I walked the audience through the framework and got their suggestions on the unfinished parts of the framework. At the end, the attendees paired up and used the framework to work together and draw a picture they both wanted to draw.

Before the workshop, I sent a survey internally asking developers a few questions about pair programming. Here are the results:

- 62.5% had over 10 years of programming experience
- 78.1% had pair programmed before joining Shopify
- 50% of the surveyor pair once or twice a week at Shopify

When asked _“What is one important trait to have when pair programming?_”, this is what Shopify developers had to say:

**Communication**

- Expressing thought processes (what you’re doing, why you’re making this change, etc.)
- Sharing context to help others get a thorough understanding
- Use of visual aids to assist with explanation

**Empathy**

- Being aware of energy levels
- Not being judgemental to others

**Open-mindedness**

- Curious to learn
- Willingness to take feedback and improve
- Don’t adhere to just one’s opinion

**Patience**

- Providing time to allow your partner to think and formulate opinions
- Encourage repetition of steps, instructions to encourage question asking and learn by doing

Now, let’s walk through the crowdsourced framework we finished at RenderATL.

_Note: For those who attended the workshop, the framework below is the same framework that you walked away with, but with more details and resources._

# The Framework

This framework covers everything you need to run a successful pair programming session, including: roles, structure, agenda, environment, and communication. You can pick and choose within each section to design your session based on your needs.

## 1a. Pair Programming Styles

There are many different ways to run a pair programming session. Here are the ones we found to be the most useful, and when you may want to use each depending on your preferences and goals.

### Driver and Navigator

Think about this style like a long road trip. One person is focused on driving to get from point A to B. While the other person is providing directions, looking for future pit stops for breaks, and just observing the surroundings. As driving can be taxing, it’s a good idea to switch roles frequently.

The driver is the person leading the session and typing on the keyboard. As they are typing, they’re explaining their thought process. The navigator, also known as the observer, is the person observing, reviewing code that’s being written, and making suggestions along the way. For example, suggesting refactoring code and thinking about potential edge cases.

If you’re an experienced person pairing with an intern or junior developer, I recommend using this style after you paired together a few sessions. They’re likely still gaining context and getting comfortable with the code base in the first few sessions.

### Tour Guide

This style is like giving someone a personal tour of the city. The experienced person drives most of the session, hence the title tour guide. While the partner is just observing and asking questions along the way.

I suggest using this style when working with someone new on your team. It’s a great way to give them a personal tour to how your team’s application works and share context along the way. You can also flip it, where the least experienced person is the tour guide. I like to do this with the Dev Degree interns who are a bit further into their training when I pair with them. I find it helps bring out their communication skills once they’ve started to gain some confidence in their technical abilities.

### Unstructured

The unstructured style is more of a freestyle way to work on something together, like learning a new language or concept. The benefit of the unstructured style is the team building and creative solutions that can come from two people hacking away and figuring things out. This is useful when a junior developer or intern pairs with someone at their level. The only downside is that without a mentor overseeing them, there’s a risk of missteps or bad habits going unchecked. This can be solved after the session by having them share their findings with a mentor for discussion.

We allocated time for the interns to pair together. This is the style the interns typically go with, figuring things out using the concepts they learned.

## 1b. Activities

When people think of pair programming, they often strictly think of coding. But pair programming can be effective for a range of activities. Here are some we suggest.

### Code Reviews

I remember when I first started reviewing code, I wasn’t sure what exactly I was meant to be looking for. Having an experienced mentor support with code reviews helps early talent pick up context and catch things that they might not otherwise know to look for. Interns also bring a fresh perspective, which can benefit mentors as well by prompting them to unpack why they might make certain decisions.

### Technical Design and Documentation

Working together to design a new feature or go through team documents. If you put yourself in a junior developer’s shoes, what would that look like? It could look like a whiteboarding session mapping out logic for the new feature or improving the team’s onboarding documentation. This could be an incredibly impactful session for them. You’ll be helping broaden their technical depth, helping future teammates onboard faster, and sharing your expertise along the way.

### Writing Test Cases Only

Imagine you’re a junior developer working on your very first task. You have finished writing the functionality, but haven’t written any tests for it. You tested it manually and know it works. One thing you’re trying to figure out is how to write a test for it now. This is where a pair programming session with an experienced developer is beneficial. You work together to extend testing coverage and learn team-specific styles when writing tests.

### Onboarding

Pairing is a great way to help onboard someone new onto your team. It helps the new person joining your team ramp up quicker with your wealth of knowledge. Together you explore the codebase, documentation, and team-specific rituals.

### Hunting Bugs

Put yourselves in your users’ shoes and go on a bug hunt together. As you test functionalities on production, you'll gain context on the product and reduce the number of bugs on your application. A win-win!

## 2. Set an Agenda

Pair programming can be used for more than just writing code. Try pairing on other problems and using tools like a whiteboard to work through ideas together.

Setting an agenda beforehand is key to making sure your session is successful.

Before the session, work with your partner to align on the style, activity, and goals for your pair programming session. This way you can hit the ground running while pairing and work together to achieve a common goal. Here are some questions you can use to set your agenda:

- What do you want to pair on today?
- How do you want the session to go? You drive? I drive? Or both?
- Where should we be by the end of the session?
- Is there a specific skill you want to work on?
- What’s blocking you?

## 3. Set the Rules of Engagement

Think of your pair programming session like a classroom. How would you make sure it’s a great environment to learn?

_"After years of frequent pair programming, my teammates and I have established patterns that give the impression that we are always right next to each other, which makes context sharing and learning from peers much simpler." -Olakitan Bello, Developer_

Now that your agenda is set, it’s time to think about the environment you want to have during the session. Imagine yourself as a teacher. If this was a classroom, how would you provide the best learning environment?

### Be Inclusive

Everyone should feel welcomed and invited to collaborate with others. One way we can set the tone is to establish that “There are no wrong answers here” or “There are no dumb questions.” If you’re a senior colleague, saying “I don’t know” to your partner is a very powerful thing. It shows that you’re a human too! Keep accessibility in mind as well. There are tools and styles available to tailor pair programming sessions to the needs of you and your partner. For example, there are alternatives to verbal communication, like using a digital whiteboard or even sending messages over a communication platform_. _Invite people to be open about how they work best and support each other to create the right environment.

### Remember That Silence Isn’t Always Golden

If it gets very quiet as you’re pairing together, it’s usually not a good sign. When you pair program with someone, communication is very important to both parties. Without it, it’s hard for one person to perceive the other person’s thoughts and feelings. Make a habit of explaining your thought process out loud as you work. If you need a moment to gather your thoughts, simply let your partner know instead of going silent on them without explanation.

### Respect Each Other

Treat them like how you want to be treated, and value all opinions. Someone's opinions can help lead to an even greater solution. Everyone should be able to contribute and express their opinions.

### Be Empathic Not Apathetic

If you’re pair programming remotely, displaying empathy goes a long way. As you’re pair programming with them, read the room. Do they feel flustered with you driving too fast? Are you aware of their emotional needs at the moment?

As you’re working together, listen attentively and provide them space to contribute and formulate opinions.

### Mistakes Are Learning Opportunities

If you made a mistake while pair programming, don’t be embarrassed about it. Mistakes happen, and are actually opportunities to learn. If you notice your partner make a mistake, point it out politely—no need to make a big deal of it.

## 4. Communicate Well

Communication is one of the most important aspects of pair programming.

Pair programming is all about communication. For two people to work together to build something, both need to be on the same page. Remote work can introduce unique communication challenges, since you don’t have the benefit of things like body language or gestures that come with being in the same physical room. Fortunately, there’s great tooling available, like [Tuple](https://tuple.app/), to solve these challenges and even enhance the pair programming experience. It’s a MacOS only application which allows people to pair program with each other remotely. Users can share their screen and either can take control to drive. The best part is that it’s a seamless experience without any additional UI taking up space on your screen.

During your session, use these tips to make sure you’re communicating with intention.

### Use Open-ended Questions

Open-ended questions lead to longer dialogues and provide a moment for someone to think critically. Even if they don’t know, it lets them learn something new that they will take away from the session. With closed questions, it’s usually a “yes” or a “no” only. Let’s say we’re working together on building a grouped React component of buttons, which one sounds more inviting for a discussion:

- Is ButtonGroup a good name for the component? (Close-ended question)
- What do you think of the name ButtonGroup? (Open-ended question)

Other examples of open-ended questions:

- What are some approaches you took to solving this issue?
- Before we try this approach, what do you think will happen?
- What do you think this block of code is doing?

### Give Positive Affirmations

Encouragement goes a long way, especially when folks are finding their footing early in their career. After all, knowing what’s gone right can be just as important as knowing what’s gone wrong. Throughout the session, pause and celebrate progress by noting when you see your partner do something well.

For example, you and your partner are building a new endpoint that’s part of a new feature your team is implementing. Instead of waiting until the feature is released, celebrate the small win.

Here are a few example messages you can give:

- It looks like we marked everything off the agenda. Awesome session today.
- Great work on catching the error. This is why I love pair programming because we work together as a team.
- Huge win today! The PR we worked together is going to help not only our team, but others as well.

### Communication Pitfalls to Avoid

No matter how it was intended, a rude or condescending comment made in passing can throw off the vibe of a pair programming session and quickly erode trust between partners. Remember that programming in front of someone can be a vulnerable experience, especially for someone just starting out. While it might seem obvious, we all need a reminder sometimes to be mindful of what we say. Here are some things to watch out for.

### Passive-Aggressive (Or Just-Plain-Aggressive) Comments

Passive-aggressive behavior is when someone expresses anger or negative feelings in an indirect way. If you’re feeling frustrated during a session, avoid slipping into this behavior. Instead, communicate your feelings directly in a constructive way. When negative feelings are expressed in a hostile or outwardly rude way, this is aggressive behavior and should be avoided completely.

Passive-aggressive behavior examples:

- Giving your partner the silent treatment
- Eye-rolling or sighing in frustration when your partner makes a mistake
- Sarcastic comments like “Could you possibly code any slower?”
- Subtle digs like “Well…that’s an interesting idea.”

Aggressive behavior examples

- “Typical intern mistake, rookie.”
- “This is NOT how someone at your level should be working.”
- “I thought you knew how to do this.”

### Absolute Words Like Always and Never

Absolute words means you are 100 percent certain. In oral communication, depending on the use of the absolute word, it can sound condescending. Programming is also a world full of nuance, so overgeneralizing solutions as _right_ or _wrong_ is often misleading. Instead, use these scenarios as a teaching opportunity. If something is usually true, explain the scenarios where there might be exceptions. If a solution rarely works, explain when it might. Edge cases can open some really interesting and valuable conversations.

For example:

- “You never write perfect code”
- “I always review code”
- “That would never work”

Use alternative terms instead:

- For always, you can use usually
- For never, you can use rarely

_“When you join Shopify, I think it's overwhelming given the amount of resources to explore even for experienced people. Pair programming is the best way to gain context and learn. In this digital by design world, pair programming really helped me to connect with team members and gain context and learn how things work here in Shopify which helped me with faster onboarding.” -Nikita Acharya, Developer_

I want to thank everyone at RenderATL who helped me finish this pair programming framework. If you’re early on in your career as a developer, pair programming is a great way to get to know your teammates and build your skills. And if you’re an experienced developer, I hope you’ll consider mentoring newer developers using pair programming. Either way, this framework should give you a starting point to give it a try.

In the pilot we’ve run so far with this framework, we’ve received positive feedback from our interns about how it allowed them to achieve their learning goals in a flexible format. We’re still experimenting and iterating with it, so we’d love to hear your feedback if you give it a try! Happy pairing!

Raymond Chung is helping to build a new generation of software developers through Shopify’s Dev Degree program. As a Technical Educator, his passion for computer science and education allows him to create bite-sized content that engages interns throughout their day-to-day. When he is not teaching, you’ll find Raymond exploring for the best bubble tea shop. You can follow Raymond on [Twitter](https://twitter.com/rchung95), [GitHub](https://github.com/rchung95), or [LinkedIn](https://www.linkedin.com/in/raymondchung95/).

Wherever you are, your next journey starts here! If building systems from the ground up to solve real-world problems interests you, our Engineering blog has stories about other challenges we have encountered. Intrigued? Visit our [Engineering career page](http://www.shopify.com/careers/specialties/engineering?itcat=EngBlog&itterm=Post) to find out about our open positions and learn about [Digital by Design](https://www.shopify.com/careers/work-anywhere?itcat=EngBlog&itterm=CCTA-DD).

### Get stories like this in your inbox!

Stories from the teams who build and scale Shopify. The commerce platform powering millions of businesses worldwide.

Share your email with us and receive monthly updates.
