---
title: "Asynchronous Communication is the Great Leveler in Engineering"
notion_id: 4a3580bc-6ded-4b51-bcad-5048a16c25d4
notion_url: https://app.notion.com/p/Asynchronous-Communication-is-the-Great-Leveler-in-Engineering-4a3580bc6ded4b51bcad5048a16c25d4
last_edited: 2023-04-25T13:17:00.000Z
source_url: https://shopify.engineering/asynchronous-communication-shopify-engineering
tags: ["English", "Help Desk", "On Call", "Productivity", "Communication", "Article", "Shopify Engineering"]
---
[https://shopify.engineering/asynchronous-communication-shopify-engineering](https://shopify.engineering/asynchronous-communication-shopify-engineering)

In March 2020—the early days of the pandemic—Shopify transitioned to become a remote first- company. We call it being Digital by Design. We are now proud to employ Shopifolk around the world.

Not only has being Digital by Design allowed our staff the flexibility to work from wherever they work best, it has also increased the amount of time that they are able to spend with their families, friends, or social circles. In the pre-remote world, many of us had to move far away from our hometowns in order to get ahead in our careers. However, recently, my own family has been negotiating with the reality of aging parents, and remote working has allowed us to move back closer to home so we are always there for them whilst still doing the jobs we love.

However, being remote isn’t without its challenges. Much of the technology industry has spent decades working in colocated office space. This has formed habits in all of us that aren’t compatible with effective remote work. We’re now on a journey of mindfully unraveling these default behaviors and replacing them with remote-focused ways of working.

If you’ve worked in an office before, you’ll be familiar with synchronous communication and how it forms strong bonds between colleagues: a brief chat in the kitchen whilst getting a coffee, a discussion at your desk that was prompted by a recent code change, or a conversation over lunch.

With the layout of physical office spaces encouraging spontaneous interactions, context could be gathered and shared through osmosis with little specific effort—it just happened.

There are many challenges that you face in engineering when working on a globally distributed team. Not everyone is online at the same time of the day, meaning that it can be harder to get immediate answers to questions. You might not even know who to ask when you can’t stick your head up and look around to see who is at their desks. You may worry about ensuring that the architectural direction that you’re about to take in the codebase is the right one when building for the long term—how can you decide when you’re sitting at home on your own?

Teams have had to shift to using a different toolbox of skills now that everyone is remote. One such skill is the shift to more asynchronous communication: an essential glue that holds a distributed workforce together. It’s inclusive of teams in different time zones, it leaves an audit trail of communication and decision making, encourages us to communicate concisely, and enables everybody the same window into the company, regardless of where they are in the world.

However, an unstructured approach can be challenging, especially when teams are working on establishing their communication norms. It helps to have a model with which to reason about how best to communicate for a given purpose and to understand what side-effects of that communication might be.

## The Spectrum of Synchronousness

When working remotely, we have to adapt to a different landscape. The increased flexibility of working hours, the ability to find flow and do deep work, and the fact that our colleagues are in different places means we can’t rely on the synchronous, impromptu interactions of the office anymore. We have to navigate a continuum of communication choices between synchronous and asynchronous, choosing the right way to communicate for the right group at the right time.

It’s possible to represent different types of communication on a spectrum, as seen in the diagram below.

The spectrum of communication types

Let’s walk the spectrum from left to right—from synchronous to asynchronous—to understand the kinds of choices that we need to make when communicating in a remote environment.

- Video calls and pair programming are completely synchronous: all participants need to be online at the same time.
- Chats are written and can be read later, but due to their temporal nature have a short half-life. Usually there’s an expectation that they’ll be read or replied to fairly quickly, else they’re gone.
- Recorded video is more asynchronous, however they’re typically used as a way of broadcasting some information or complimenting a longer document, and their relevance can fade rapidly.
- Email is archival and permanent and is typically used for important communication. People may take many days to reply or not reply at all.
- Written documents are used for technical designs, in-depth analysis, or cornerstones of projects. They may be read many years after they were written but need to be maintained and often represent a snapshot in time.
- Wikis and READMEs are completely asynchronous, and if well-maintained, can last and be useful forever.

## Shifting to Asynchronous

When being Digital by Design, we have to be intentionally more asynchronous. It’s a big relearning of how to work collaboratively. In offices, we could get by synchronously, but there was a catch: colleagues at home, on vacation, or in different offices would have no idea what was going on. Now we’re all in that position, we have to adapt in order to harness all of the benefits of working with a global workforce.

By treating everyone as remote, we typically write as a primary form of communication so that all employees can have access to the same information wherever they are. We replace meetings with asynchronous interactions where possible so that staff have more flexibility over their time. We record and rebroadcast town halls so that staff in other timezones can experience the feeling of watching them together. We document our decisions so that others can understand the archeology of codebases and projects. We put effort into editing and maintaining our company-wide documentation in all departments, so that all employees have the same source of truth about teams, the organization chart, and projects.

This shift is challenging, but it’s worthwhile: effective asynchronous collaboration is how engineers solve hard problems for our merchants at scale, collaborating as part of a global team. Whiteboarding sessions have been replaced with the creation of collaborative documents in tools such as Miro. In-person Town Halls have been replaced with live streamed events that are rebroadcast on different time zones with commentary and interactions taking place in Slack. The information that we all have in our heads has needed to be written, recorded, and documented. Even with all of the tools provided, it requires a total mindset shift to use them effectively.

We’re continually investing in our developer tools and build systems to enable our engineers to contribute to our codebases and ship to production any time, no matter where they are. We’re also investing in internal learning resources and courses so that new hires can autonomously level up their skills and understand how we ship software. We have regular broadcasts of show and tell and demo sessions so that we can all gather context on what our colleagues are building around the world. And most importantly, we take time to write regular project and mission updates so that everyone in the company can feel the pulse of the organization.

Asynchronous communication is the great leveler: it connects everyone together and treats everyone equally.

## Permanence in Engineering

In addition to giving each employee the same window into our culture, asynchronous communication also has the benefit of producing permanent artifacts. These could be written documents, pull requests, emails, or videos. As per our diagram, the more asynchronous the communication, the more permanent the artifact. Therefore, shifting to asynchronous communication means that not only are teams able to be effective remotely, but they also produce archives and audit trails for their work.

The whole of Shopify uses a single source of truth—an internal archive of information called the Vault. Here, Shopifolk can find all of the information that they need to get their work done: information on teams, projects, the latest news and video streams, internal podcasts and blog posts. Engineers can quickly find architecture diagrams and design documents for active projects.

When writing design documents for major changes to the codebase, a team produces an archive of their decisions and actions through time. By producing written updates to projects every week, anyone in the company can capture the current context and where it has been derived from. By recording team meetings and making detailed minutes, those that were unable to attend can catch up later on-demand. A shift to asynchronous communication means a shift to implied shift to permanence of communication, which is beneficial for discovery[g][h][i], reflection and understanding.

For example, when designing new features and architecture, we collaborate asynchronously on design documents via GitHub. New designs are raised as issues in our technical designs repository, which means that all significant changes to our codebase are reviewed, ratified and archived publicly. This mirrors how global collaboration works on the open source projects we know and love. Working so visibly can be intimidating for those that haven’t done it before, so we ensure that we mentor and pair with those that are doing it for the first time.

## Establishing Norms and Boundaries

Yet, multiple mediums of communication incur many choices in how to use them effectively. When you have the option to communicate via chat, email, collaborative document or GitHub issue, picking the right one can become overwhelming and frustrating. Therefore we encourage our teams to establish their preferred norms and to write them down. For example:

- What are the response time expectations within a team for chat versus email?
- How are online working hours clearly discoverable for each team member?
- How is consensus reached on important decisions?
- Is a synchronous meeting ever necessary?
- What is the correct etiquette for “raising your hand” in video calls?
- Where are design documents stored so they’re easily accessible in the future?

By agreeing upon the right medium to use for given situations, teams can work out what’s right for them in a way that supports flexibility, autonomy, and clarity. If you’ve never done this in your team, give it a go. You’ll be surprised how much easier it makes your day to day work.

The norms that our teams define bridge both synchronous and asynchronous expectations. At Shopify, my team members ensure that they make the most of the windows of overlap that they have each day, setting aside time to be interruptible for pair programming, impromptu chats and meetings, and collaborative design sessions. Conversely, the times of the day when teams have less overlap are equally important. Individuals are encouraged to block out time in the calendar, turn on their “do not disturb” status, and find the space and time to get into a flow state and be productive

A natural extension of the communication norms cover when writing and shipping code. Given that our global distribution of staff can potentially incur delays when it comes to reviewing, merging, and deploying, teams are encouraged to explore and define how they reach alignment and get things shipped. This can happen by raising the priority of reviewing pull requests created in other timezones first thing in the morning before getting on with your own work, through to finding additional engineers that may be external to the team but on the same timezone to lean in and offer support, review, and pairing.

## Maintaining Connectivity

Once you get more comfortable with working ever more asynchronously, it can be tempting to want to make everything asynchronous: stand-ups on Slack, planning in Miro, all without needing anyone to be on a video call at all. However, if we look back at the diagram one more time, we’ll see that there’s an important third category: connectivity. Humans are social beings, and feeling connected to other, real humans—not just avatars—is critical to our wellbeing. This means that when shifting to asynchronous work we also need to ensure that we maintain that connection. Sometimes having a synchronous meeting can be a great thing, even if it’s less efficient—the ability to see other faces, and to chat, can’t be taken for granted.

We actively work to ensure that we remain connected to each other at Shopify. Pair programming is a core part of our engineering culture, and we love using Tuple to solve problems collaboratively, share context about our codebases, and provide an environment to help hundreds of new engineers onboard and gain confidence working together with us.

We also strongly advocate for plenty of time to get together and have fun. And no, I’m not talking about generic, awkward corporate fun. I’m talking about hanging out with colleagues and throwing things at them in our very own video game: Shopify Party (our internal virtual world for employees to play games or meet up). I’m talking about your team spending Friday afternoon playing board games together remotely. And most importantly, I’m talking about at least twice a year, we encourage teams to come together in spaces we’ve invested in around the world for meaningful and intentional moments for brainstorming, team building, planning, and establishing connections offline.

Asynchronous brings efficiency, and synchronous brings connectivity. We’ve got both covered at Shopify.

James Stanier is Director of Engineering at Shopify. He is also the author of Become an Effective Software Manager and Effective Remote Work. He holds a Ph.D. in computer science and runs theengineeringmanager.com.

Wherever you are, your next journey starts here! If building systems from the ground up to solve real-world problems interests you, our Engineering blog has stories about other challenges we have encountered. Intrigued? Visit our Engineering career page to find out about our open positions and learn about Digital by Design.
