---
title: "Scaling Translations at Spotify"
notion_id: 0d281d66-c737-4ee3-a7dc-d092df3bbf13
notion_url: https://app.notion.com/p/Scaling-Translations-at-Spotify-0d281d66c7374ee3a7dcd092df3bbf13
last_edited: 2023-01-27T02:17:00.000Z
source_url: https://engineering.atspotify.com/2022/09/scaling-translations-at-spotify/
tags: ["English", "Programming", "System Design / Software Architecture", "Article", "Spotify Engineering"]
---
Last year, we added support for 36 new languages to our products in one go, for a total of 62 languages. This article describes how we delivered on such an immense localization effort at Spotify. We called the project Scaling Translations.

## The business case

We believe that localization is key for engaging with our users across the world. Industry peers have indicated that translations have been one of their strongest drivers of growth over the years. When we first started thinking about Scaling Translations, Spotify was available in 26 languages. By increasing the number of supported languages from 26 to more than 60, we believed we could expand our reach and impact those users whose native languages were not supported at that time. We could also expand our total addressable market and make our product more accessible to an even more diverse user base.

## The software engineering efforts needed

To support the business case, it was essential to understand what was required on the tech side to pull off the effort.

First, we needed to assess our technical landscape. We asked ourselves a number of questions — what capabilities exist today to support our ambitions? What are the pain points within our current architecture? How will these pain points be amplified when the number of supported languages has more than doubled? What new problems do we foresee both technically and organizationally? What capabilities do we need to sustainably support all of these languages, both short and long term?

While adding support for new languages is nothing new at Spotify, historically it’s been done in support of launching our services in a new country or region, without adapting and developing the overall architecture. Adding just one new language has not been as straightforward as one might think. Spotify’s tech stack is large and complex, and it has grown significantly over the years as a result of hypergrowth. [Conway’s law](https://en.wikipedia.org/wiki/Conway%27s_law) manifests itself in our architecture.

The effort behind Scaling Translations was unprecedented — we needed to make sure we had a solid strategy, so our users could have the best possible experience. We needed to create a completely new approach to the problem.

We spent two months on a quest to find answers to the questions we had at the start of our journey. It involved talking to software engineers, product managers, engineering managers, colleagues on the business side — all across the company. We also had to roll up our sleeves and dig through numerous different source code repositories, consult system diagrams, and analyze production metrics and telemetry. After concluding our deep dive, we summarized our findings in a comprehensive report — the 25-page write-up covered extensive details of existing systems, their interactions, use cases, business metrics, what capabilities we needed to develop, and the effort it would take. We proposed a direction for how to realize the business case and product vision using technology, and we also included several considered, but discarded, alternatives for the project for added context.

Once the technical effort had been defined, we secured the required funding and staffing. Six months ahead of the scheduled release date, we hired a small, dedicated engineering team to focus on this project. This was the first time we’d had dedicated engineers work alongside our localization team on a joint project to launch new languages. The engineering team worked relentlessly to address the many pain points and required improvements outlined in the technical scoping document, collaborating with other engineering teams and business stakeholders across the company.

Efforts included defragmenting the existing tech stack, automating mundane linguistic quality-assurance tasks for our mobile apps, and using sophisticated automation to update 30+ backend services en masse to successfully integrate support for the new languages. At our 2021 [Spotify Stream On](https://newsroom.spotify.com/stream-on/) event, we announced the support of 36 new languages. Our team was ready to start rolling out the new languages the following day, and we successfully launched the project on time, within budget, and with great quality.

## The localization effort

In the past, launching in a new market with a new language took significantly more time. To simultaneously launch 36 languages in six months (while maintaining and delivering a good user experience), we honed in on six key areas:

1. **Creating a minimum viable product:** We stripped down the product to what matters most to users in local markets. Here at Spotify, we love the idea of delivering quickly, learning from results, iterating, and expanding, if needed.
2. **Staying nimble:** This is a philosophy that governed the whole process. We had to think about how to hit the ground running with a minimum number of resources and as little ramp-up time as possible. We had to make a number of hard decisions to allow for progress and to avoid paralysis by analysis.
3. **Establishing strong partnerships with vendors:** We included our vendors in the process from day one. They provided market and language research and assistance in building the business case. They participated in strategic decisions well before production started and built linguistic teams and quality processes quickly and efficiently. Most importantly, they shared our approach and philosophy for this project. We provided them with context and visibility, enabling them to understand what we had set out to accomplish and how. Their support and investment in the process was critical to the success of the project.
4. **Seeking cross-functional support:** To be able to deliver on the work, we needed the support of dozens of teams at Spotify. Some of the biggest challenges included identifying the teams on which we had dependencies, negotiating buy-in and aligning with their roadmaps, and extending the nimble philosophy to these teams.
5. **Balancing flexibility with prioritization:** As awareness about the project increased, more people brought forth questions and expressed interest in incorporating their areas of the product in the initiative. While we welcomed this, we had to make tough calls on scope based on prioritization. The reality was that scope changes were part of the process throughout, and we subsequently ended up with lots of learnings about what an MVP really meant while still providing a quality experience to our users. The key to making this project happen was getting things off the ground and providing our stakeholders with as much visibility as possible to get them on board.
6. **Communication: **For a project of this size and complexity, clear and timely communication targeted accordingly to a given audience was paramount; it was the common thread in the successful execution of the above elements. We set up dedicated Slack channels to streamline communication, established a regular meeting cadence for the critical workstreams, and held daily localization team stand-ups dedicated to this project alone, which allowed us to maintain a healthy level of visibility and nimbleness in collective decision making. We created a dedicated Jira dashboard for functional bug reporting from localization testing, which helped streamline the process of triaging and fixing bugs.

## Conclusion

Scaling Translations proved to have a positive business impact on our user base and drastically reduced the percentage of users not served their preferred language — while also making Spotify a more diverse and inclusive product from a language perspective. The initiative proved that challenging traditional ways of thinking about international expansion and product launches can be a worthwhile effort in order to move faster as a business. In fact, challenging the status quo _is_ possible to do while keeping a high level of quality in the final product and providing users with a culturally relevant experience. It also proved to us the value of being nimble, courageous, and laser-focused on fixing the most important problems quickly while iterating with fast feedback loops and tight scope and stakeholder management.

This project was a huge challenge for all involved. It required a change of mindset and in ways of working, and through a very solid business case with alignment across the organization and using technology to our advantage, it turned out to be a huge success. The project has now opened the doors for new globalization opportunities at Spotify that were al
