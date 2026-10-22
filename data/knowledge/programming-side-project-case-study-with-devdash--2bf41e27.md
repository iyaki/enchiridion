---
title: "Programming Side Project: Case Study With DevDash"
notion_id: 2bf41e27-8fb4-45d5-8ce0-36481fc126ef
notion_url: https://app.notion.com/p/Programming-Side-Project-Case-Study-With-DevDash-2bf41e278fb445d58ce036481fc126ef
last_edited: 2026-09-21T17:22:00.000Z
source_url: https://thevaluable.dev/programming-side-project-example-devdash/
tags: ["Article", "The Valuable Dev", "English", "Programming", "Producer (Individual Contributor)"]
---
[https://thevaluable.dev/programming-side-project-example-devdash/](https://thevaluable.dev/programming-side-project-example-devdash/)

Let sit down comfortably around the campfire, you and me and our common imaginary friends. The night is warm, the crickets are singing, the owls whoop, the mysterious cracks of the fire send sparkles to the night sky, enlightened already by million of stars.

Let me tell you the story of DevDash.

“Is it an animal?” you might ask if you didn’t click on the link.

Behind this singular name is hidden my last side project, a terminal dashboard for developers, written in Golang. Since I released the first early version not long ago, I thought it could be interesting to dive deep into the process of building a side project, with a concrete example.

I’m sorry. This article is not only about building. I try to answer these questions:

- How did I find the idea?
- What are the goals of DevDash? How did I find them, and why?
- How did I manage the project?
- Why did I choose the technology to build it?
- What was the building process?
- How did I promote the project?

If you like creating as much as I do, if you want to show your creations to the world, and if you want to move toward your own goals, stick with me. Let’s learn together from this experience how to create, manage and promote a side project.

## Shaping the General Concept of DevDash

### The Sparkle

How did I come up with the concept behind DevDash?

Each week, I give myself some time for different tasks, all linked to creating or learning. One of this task is called inspiration time. It’s a time when I go on social media (mostly Reddit and Indie Hackers) and other websites in order to give advice, find pain points and ultimately to find inspiration for side projects. Then, I write whatever I found in mindmaps.

In general, I spend 75 minutes a week on this, depending on my priorities.

One day, I stumbled on Reddit into a discussion about having Google Analytics in the terminal, with the possibility to gather data from different accounts. Since I love working in the terminal (my whole development environment is based on it, as I explain here), I thought it was a very good idea.

Reading the conversation and seeing many people liking it, I understood it was a good pain point to solve. Especially since I came across this project on Github not long before, and I thought it would be really fun to create a terminal dashboard.

This was the sparkle, the first version of the idea. I called it TermDash, The Terminal Dashboard.

From there, I let the idea maturing for several weeks, adding and deleting some other minor ideas attached to the main ones. Giving time for the idea to grow and expand is better than acting on it impulsively. You will allow your diffuse mode of thinking to really shape your project idea. That is, you will think about it in the background, adding more details to it without actively thinking about it.

It’s like a fruit on a mindmap branch which get ripes overtime.

I wrote already an article about my idea management, if you want to know more about it.

### The Audience

Time passed. I was thinking from time to time about this TermDash idea. I was wondering if I could find a defined audience which could be interested in the project.

This is essential: before building something, you need to find an audience which could be interested by the concept. Without audience, your project will simply stay in a dark corner of the Internet, abandoned, alone. Overtime, you would lose your motivation to develop it, since nobody would use it.

To keep things simple, I thought building a terminal dashboard primarily for developers would be a good idea, for the following reasons:

- Not many people use a terminal, except developers (in a broad sense).
- I’m a developer. Therefore, it’s easier for me to find what developers need. I’m part of the audience.
- I could still use my side project for my own needs, if nobody was interested by it.

That was it: I wanted to build a terminal dashboard for developers. I renamed my project DevDash, The Terminal Dashboard for Developers.

Hence, I thought it could be nice not only displaying data from Google Analytics, but from other services as well, like Github, Gitlab, Jira, and so on. Every data a developer could need to build, promote and maintain his own side projects, the projects of his clients, or the projects of the company he works with.

DevDash became a more ambitious project, but I knew that my productivity system and my solid habits could allow me to work on it at a good pace, without losing my motivation before releasing the first version.

### Differentiating DevDash

Now that the idea of the project was taking shape, that I roughly found my audience, I needed to look to already existing terminal dashboards and see how DevDash could bring something different.

Nobody wants another tool which does exactly the same things as the ones already available, especially since these tools are developed for a longer time, with more advanced and mature functionalities. It’s like trying to win a marathon while starting hours after everybody else.

The most popular terminal dashboard out there is WTF. It’s a fantastic tool! However, even if I discovered it way before thinking to build a terminal dashboard, I never used it, and I even totally forgot about it.

Why?

- It wasn’t focused enough around my needs, as a developer.
- The configuration was not flexible enough. For the widget git for example, I can only display Branch, Changed file and Recent Commits. What about filtering the commits? Do I really need to display the changed files? Could I see the traffic of my repository’s page?
- It’s not visual enough to me. I want to be able to see everything I need quickly, at a glance. I don’t want to navigate in there with my keyboard to find the information, or simply going through a lot of text.

As I said, I think WTF is a gorgeous project, but it wasn’t fulfilling my needs. That’s when I understood that DevDash needed to have a more flexible configuration, for the user to be able to access quickly the data he specifically needs.

As I was saying above, the project tdash was very appealing to me as well. I really liked the look of it, the clear tables displaying the data. However, the project was very much build around the specific needs of the author, and didn’t include any way to personalize the widgets.

## The Purpose and Goals of DevDash

### A Useful Tool for Developers

First and foremost, DevDash was meant to scratch my own itch. However, I didn’t want to create something only for me. It had to be useful for other developers as well.

I feel more motivated and fulfilled when I help people. It gives me a lot of motivation when people are genuinely interested by what I build. Interest can be positive or negative, and we should seek both. If somebody spends time spotting defects and give negative critics about my project, it’s a mark of interest as well. At the end, ignorance is the worst to deal with.

To take a concrete example, I would like to present to you another one of my side project, testomatic.

When I built testomatic, I didn’t ask myself if somebody would be interested by the project. I just built it, for my needs. Since I had a problem testomatic could solve (automatically running unit tests while saving them), I thought many people would have the same problem.

I was wrong. Nobody was interested by testomatic: it only has three stars on Github. I spoke about it around me, on social media, I tried a few things. This is what can happen if you only scratch your own itch. How do you know if somebody else will be as interested as you are?

The poor testomatic wasn’t a mistake. The main goal of the project was to learn Golang, as I explain here. It was a success. I still use it, too.

However, for DevDash, when I saw the conversation on Reddit where people were genuinely interested by having Google Analytics in the terminal, I knew if I would build it, it would be useful at least for these people.

I knew as well that TUIs are quite popular on Github in general. Another clue that developers could be interested by the project.

You might think that all of this didn’t prove much. It’s not a lot of data to rely on. At the end, the destiny of DevDash could have been the same as testomatic.

Let’s think about it: when I built testomatic, I had 99% of uncertainty if the project would be interesting by other developers.

With DevDash, my uncertainty was lower, since I saw developers speaking about the problems it could solve. I saw how popular TUIs could be. This is what really matters: not being 100% sure that everybody will love your project (it’s impossible), but reducing your uncertainty enough for you to reduce the risks of a total lack of interest.

### Motivation And Joy

One of the goal of DevDash was simple: I needed to enjoy building it!

Finding motivation for a side project can be difficult, especially if you have a daily job. You need all possible sources of motivation to continue building it and improving it. I did a list of these main sources:

- DevDash would be useful for me.
- It’s likely that DevDash would help others.
- DevDash would give me the freedom I don’t have in my daily job to experiment and do whatever I want.
- DevDash would be built with technologies I’m interested in.
- DevDash would be a good way to improve my project management skills.
- DevDash would be a good opportunity to improve my technical skills.

Let’s take another counter example: I developed in 2017 / 2018 another side project, Sharetoall. I spoke about it already on this blog.

This project is still online at the time I write these lines, but I will close it soon.

I did many mistakes with this project, and I learnt a lot from them. The main reason why I’ll close it is simple: I don’t enjoy building it anymore. I still resonate with the goals I had at the time, but each time I try to add some features or fix bug, I need to force myself… too much.

The morale of the story: being only motivated by your goals is not enough. Enjoying building your side project piece by piece is crucial. You need multiple sources of motivation you know will keep you going,

If you want me to come back to my other side project Sharetoall, how I built it and why I stopped maintaining it, don’t hesitate to let a comment or to contact me on social media (you’ll find them on the about page).

### The Core Functionalities of DevDash

The first version of DevDash needed to have the core functionalities I wanted. This is the essence of an MVP (Minimum Viable Product): you need to include in your prototype everything which makes your project unique and valuable. This is the only way to know if your product could help the audience you want to help. It’s not only testing an idea anymore: it’s testing its execution.

Keep in mind that these core features are not perfectly implemented. The goal is, for each new iteration of a side project, to come closer to your goals. Not to reach them 100%. Don’t seek perfection, or you will never launch anything, and, therefore, never learn from it.

### Only the Data You Need

DevDash give the possibility for everybody to display the data they need on a dashboard. No need to have 1234 tabs open on different websites, navigating through filters and pages using weird UIs. I look at you, Google Analytics.

What I want on my dashboards is important data to improve the users experience of my projects, like this blog for example. For that, I want some useful metrics, not a whole array of useless vanity metrics.

### A Flexible Yet Simple Configuration

Since I wanted the user to choose what data he wanted to display, I needed to give them a lot of flexibility to configure their dashboard.

This is the main feature which distinguish DevDash from the other terminal dashboards out there. No need to code anything: creating a dashboard should be as easy as writing some yaml to choose the data displayed, the look of the widgets, their sizes on the screen.

### An Effective Documentation

A strong documentation needs to help the users to use Devdash.

Indeed, giving the option to configure a dashboard in a flexible way introduces a lot of complexity, which needs to be managed.

The documentation is one way to do so. It has three purposes:

1. Quickly explain what DevDash is, in order for the reader to quickly assess if the project is answering his needs or solving a problem he has.
2. Proposing an easy quickstart to allow the reader to try DevDash.
3. Proposing a detailed documentation which describe all the configuration possibilities, for those who wants to have very personalized dashboards.

At the time I write these lines, the whole documentation is in the README of the project. However, this is a lot of information at once, especially for the user who comes on the project’s page on Github and want to simply see what it is. That’s why I will move most of the documentation on a proper website in the future.

## Project Management

### DevDash Mindset

I wrote already an article about the best mindset to have for a side project. Let see how I applied these advice for DevDash.

### Step By Step

I knew I would be able to “work” on DevDash on and off. I had only one rule: working for a minimum of 25 minutes (a pomodoro) at a time, and being focused solely on DevDash during that time. That’s why I wanted to divide my tickets as small as I could, to be able to complete them in this short time frame.

These tickets were parts of bigger milestones. For example:

- Prototyping a first version of the yaml configuration.
- Prototyping a widget with dummy data to test different TUIs.
- Fetching Google Analytics data from an API.
- Displaying the data.

Keep in mind that each of these milestones were constituted of a lot of small tickets, like:

- Creating the basic folder organisation for the project (I tried to follow the advice of this article as an experiment).
- Adding YAML and config dependencies.
- Adding Project / Services / Widget in the configuration.
- Adding the possibility to add widget by row and column, using the configuration.

Creating small tickets is good to keep the motivation flowing: when I move them to “done”, it always feels rewarding. We need these rewards, these little victories to keep our motivation as high as possible.

Another good reason to make small tickets: I want each feature of DevDash to be as simple as possible, while still bringing the project toward the goals I defined. By building these features small step by small step, it’s easier to know when to stop.

I want to develop useful features for the users of DevDash. Therefore, I want as much feedback as possible about these features, before fleshing them out more. If some features are never used, I can easily discard them without losing too much time.

### Consistency and Balance

Now, to keep some consistency and to move forward quickly, in order for my motivation not to go down, I try in general to keep a consistent and balanced life. It’s not perfect of course, but I succeeded to create some good habits over the years.

### Avoiding Useless Stress

First of all, I never imposed deadlines to myself.

You might think that, without deadlines, you would not do anything. You might think that you need to put some pressure on yourself to avoid procrastinating.

I believe that there are other ways to achieve the same result.

If you know how to create and maintain good habits, you won’t need to force yourself for anything, and you won’t need deadlines or other stressful trick. Working on your side project will become automatic. Here’s the thing: more habit you build, more you get better at building them.

With deadlines, you can as well be victim of the student syndrome: waiting for the last minute to do your work. I doubt the quality will follow.

Against procrastination, you need to be aware that the most difficult is to begin. Allow yourself to work on your side project only 10 minutes, and then stop if you want to. It reduces this fear of beginning, drastically. I follow this trick for every of my creative endeavor, and I always work more on them than the imposed 10 minutes. My procrastination has never been that low, and believe me, it was quite high in the past.

Simply because I care about the project. Simply because I like building it. We come back to that, again.

In order to be as productive as possible, only when I know exactly what to do (outcomes are far more important than productivity in a vacuum), I always try to focus entirely on the tasks at end during my pomodoros. By doing so, I avoid having the “I didn’t do enough” stress. There is no “enough”: you need to be aware of your pace and keep at it. You need to be aware as well that results come with patience and consistent commitments.

Since I try to focus as much as I can, I don’t want to work on DevDash for a long period per day. It’s tiring to focus, and at one point you do more harm than good to your codebase, because you’re simply too exhausted to produce quality code.

Quality is always a time savior on a middle / long term, and it allows you to keep the pleasure of building a good side project alive.

### Learning From The Experience

In order to keep some consistency on my productivity, and to see the possibilities to improve it, I tracked the time I spent on the project:

y-axis: pomodoros / x-axis: date

I worked on it 23.75 hours (the blue curve “actual”) in 2018. The time I expected to work on it was 8.33 hours. DevDash was pretty low on my priority list.

I did much more than I expected: as I said, beginning is the most difficult. When I’m in the flow, I can spend hours on it, if I have time. I’m careful as well to stop when something block me for a long time, for my diffuse mode of thinking to kick in. Focusing too long on a problem means as well that you might narrow your focus too much, and miss a whole array of solutions. In short, sometimes, you need to sleep on it.

Doing pomodoros helps in that regard: it reminds you how much time you already spend on your problem and is a good indicator that you should stop if it takes too long.

I enjoyed building DevDash the first half of 2019, as you can see:

y-axis: pomodoros / x-axis: date

At the time I write this article, I worked 67 hours on it, even if I planned to work only 2 hours!

I was very motivated to see the project going further. Consequently, I was working on it during the week-ends, after every other weekly tasks I’ve planned was done.

During these periods of time, I’ve done 50 tickets and discarded 11 of them. It felt good each single time.

Overall, I was able to achieve 1 ticket every 1.5 hour, which is pretty good. Especially when you consider that I had to do a lot of research for basically everything: TUIs, how to connect to the different APIs, how to get the data I wanted and how to format it.

This definitely proved me again that:

- It’s definitely possible to achieve something with a minimum of pressure, working on and off on it.
- Using the Pomodoro technique to stay focus for a given period is really effective.
- Working with small tickets you can easily achieve or discard is motivational and practical. It’s important to be able to throw away some ideas as well, not everything you think about will move you toward your project’s goal.

Of course, other explanations are valid to explain this productivity. For example, I didn’t have to communicate with anybody: I could choose what I wanted to do, how I wanted to do it, without spending time arguing about the why.

It’s true as well that misunderstandings with anybody in the chain of process (management, your fellow colleagues, QA…) can slow down any project.

On the other hand, speaking with colleagues can help you to solve your problems easily, and they can point out some pitfalls you would not be able to see alone. Plus, you can have a good synergy between the different skills of your fellow teammates, which can push you forward. Without even speaking about the mistakes and bug you can sport with pull requests.

At the end, I really believe that deadlines, stress and management-deciding-everything is not necessary to accomplish what you want to accomplish. You can be productive with the techniques I described, and enjoying the process even more.

### The Tools

I already wrote an article about the tools I used to manage my side projects. If you want to know why I use these tools, this is the article you might want to read.

To summarize quickly, I used:

- Productivity Challenge Timer on Android, to measure my time and to follow the Pomodoro technique.
- Trello to manage my tickets.
- I have a mindmap on Mindmeister to organize what I will do on a weekly basis: how many pomodoros I will work on DevDash, or to learn computer science, for example. This is the “expected” red curves on the charts above.
- On the coding side of things, I used my lovely and mouseless development environment.

## Building DevDash

### Technologies and Dependencies: how to choose them?

I chose to develop a CLI for two main reasons:

1. As I stated above, I love the shell. I use it for everything.
2. I needed to only focus on the backend part. No need to build any UI, which was simplifying everything. I’m more of a backend person.

I chose Golang for the programming language:

1. Coding using Golang is a dream. It’s a very simple yet very powerful language. I personally love it!
2. Nowadays, I’m working with Golang full time. I didn’t want the technology to be an obstacle and slowing me down. The goal was not to learn a new language.
3. Golang allows you to build cross-platform binaries very easily. That’s why DevDash is available for Windows, MacOS and Linux! Not a necessity, but a very nice plus.
4. The only thing you need is a binary and a configuration file. Everything is self contained. No need to install thousands of weird dependencies directly on your machine. I look at you, Noje.js and npm.

For the display, after prototyping with different TUIs, I settled down with termui. However, during development, a big rewriting of the library happened and the version I was using was not supported anymore. I tried at one point to use the newest version, but I finally decided that the old one was better for my needs.

Long story short, in order to fix some bugs and add some new functionalities, I had to fork termui. It gave me more control on the display itself, but it means as well that I need to maintain another dependency.

### The Development Process

My development process is pretty straightforward:

1. I’ve got some time! I begin my Pomodoro timer, for 25 minutes (or more) of focused work, or, more accurately, playtime.
2. I pick a ticket in my Trello board. If it takes too much time, I divide it into smaller tickets.
3. I work on the feature.
4. I write the most important unit tests.
5. I come back to 2, or I stop.

This is the perfect scenario. Sometimes, I won’t test anything, just manually, quickly. However, I write the test the next time I work on the project. If I forgot, the next time I need to modify or refactor the same code, I will directly write the tests without doing anything more.

The most “important unit tests” mean: testing the functions which transform the data. Most of the time, it will be data from some API to display them in the TUI, or the data from the config file. Any kind of mapping from one source of data to an output will be tested in priority.

I want to be sure that the data is correct.

I don’t hesitate as well to refactor often when I need it, to be able to make the code more understandable, for my future self and for the others who might want to contribute. The difficult part is to keep a good balance between coding new features and refactoring.

At the end of May 2019, the first early version of DevDash was finished. It was time to spread the word about it, to see if people would be interested by the concept.

As many developers out there, I’m not a marketing guru. I have some basic knowledge, but in general I have difficulties to promote my work. However, I want to get better at it, and DevDash was a good opportunity for that as well.

### The Open Source Platform

For an open source project without website, filling correctly the fields on your open source platform is essential. For DevDash, I’ve chosen Github.

For the viewer to understand as quickly as possible what DevDash is about, I added:

- The description “Highly Configurable Terminal Dashboard for Developers”. It tries to summarize as simply as possible the main ideas: what differentiate DevDash from similar software (“Highly Configurable”), what it is (“Terminal Dashboard”) for whom (“for Developers”).
- A bunch of tags for people to find easily the project on Github.
- A Gopher-logo in the README. The Gopher is Golang’s mascot, and I like it very much. I think it’s always nice to have some visual elements.
- A couple of badges to show reassuring data: the state of my CI (using Travis), the code quality, the type of license. It shows that I care to produce a project of quality and I’m not afraid to be honest.

I don’t spend a lot of time on Social Media, except on Reddit a couple of hours per week. That’s where I decided to begin my “marketing campaign”.

My idea was to post a short message with a screenshot attached to it, for people to see quickly what DevDash is about and how it looks like. Here the posts and the likes they’ve got:

1. First round (2019-05-28):67 likes in r/programming24 likes in r/webdev
2. Second Round (2019-05-29):126 likes in r/golang
3. Third Round (2019-07-16):76 likes in r/commandline

I posted as well to:

- 2019-05-28 - Hacker news. No reaction or feedback.
- 2019-05-29 - Google Nuts. No reaction or feedback.
- 2019-05-31 - 8 likes on Indie Hacker.

Here’s the traffic my page had, courtesy (of the not-yet-released-next-version) of DevDash:

On the top you have the number of Github stars per day the project collected, and at the bottom the number of views the DevDash repository’s page had. Unfortunately, Github doesn’t provide more than one week of page views.

As you can see, the 16th of July, my post on the subreddit commandline brought 332 views, and 27 stars. It can mean many things:

- Not every developer using the terminal is interested by the concept.
- Not everybody understood the concept.
- Not everybody thought about starring the project.

I think the three explanations are possible, especially when you see that there were more likes on Reddit (76) than stars on Github (26) the 16th of July.

At the end, all of that proved me one thing: enough people are interested by DevDash for me to keep my motivation flowing. This is what I wanted to know.

Especially when you think that:

- There are so many projects out there it’s difficult to stand out of the crowd, even a bit. The interest manifested by likes and Github stars is encouraging.
- More than 300 Github stars for a pretty light “promotion” is pretty good.
- Most open source project gain on popularity over time. Like everything else on Internet, we need to keep at it, be patient and stay consistent to have results.

## The Future of DevDash

At the time of writing, I still continue to have fun with DevDash. I add new functionalities which will be interesting for developers.

At least, that’s what I think. And that’s a problem.

In fact, I have a limited amount of qualitative data about the developers interested by the project. I mostly have, for now, quantitative data (likes and stars).

Some developers gave me precious advice (positive and negative) on Reddit (thanks for them!), but I will need more feedback in the future, to know what feature the users themselves would like to use.

That’s why I think more and more about creating a community around the project, using a forum or something similar. Before that, I might create a website and some content to improve the visibility of the project and to improve its documentation.

Right now, my top priority is to make DevDash as easy to use as possible. First, because tackling complexity is fun, and second, because it’s one of the major problem I see right now, especially on the configuration level.

Don’t hesitate to react in the comments, if you have different experiences with different projects of yours. By gathering them, we might learn from each other, build better projects and help more people.
