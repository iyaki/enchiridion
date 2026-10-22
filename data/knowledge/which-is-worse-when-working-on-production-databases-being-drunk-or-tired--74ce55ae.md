---
title: "Which is worse when working on production databases? Being drunk or tired?"
notion_id: 74ce55ae-e7a7-47ea-b4a9-4267db598079
notion_url: https://app.notion.com/p/Which-is-worse-when-working-on-production-databases-Being-drunk-or-tired-74ce55aee7a747eab4a94267db598079
last_edited: 2023-03-09T15:03:00.000Z
source_url: https://ledgersmbdev.blogspot.com/2023/03/which-is-worse-when-working-on.html
tags: ["English", "Health", "Article"]
---
I have decided to do a series of mini-articles on human factors in database operations. This is the first, and covers fatigue.

In [my talk](https://fosdem.org/2023/schedule/event/postgresql_the_human_factor_why_database_teams_need_crew_resource_management/) at the PostgreSQL devroom of Fosdem, I asked a few questions:

1. How many of you have seen someone work on a production database while drunk? About half the audience.
2. How many times does this cause a major incident? No hands.
3. How many of you have seen someone cause a major incident by working on a production database while tired? Half the audience again raised their hands.

As an industry, we do not take fatigue seriously enough. We appreciate people who come in and work after long disruptive on-call shifts. We don't tell people they are tired and therefore not safe to work on production systems.

We need to do better. Every single major mistake in my career that has caused production problems has been caused either by power distance problems or by fatigue.

I am not saying people should come into work drunk. There are probably a number of contextual aspects to why drunkenness doesn't cause a problem in these cases. However I am saying that people should not touch production systems under fatigue.

Of course this is easier said than done, If we are drunk, we can feel it, but with even light stress, we often don't feel our fatigue. We aren't capable of self-monitoring our conditions in this regard. Fatigue is thus insidious -- it gradually sneaks up on our, invisible, until we make critical mistakes and bad things happen.

While there are reasons to weigh the balance differently in some areas such as operating motor vehicles (to say nothing about flying an aircraft), the fact is that general brain-intensive can be impaired via moderate fatigue [perhaps more than](https://onlinelibrary.wiley.com/doi/full/10.1046/j.1365-2869.1999.00167.x) levels of alcohol we consider unacceptable while driving.

If we value production operations, we should adopt the following rule: friends don't let friends work on the production databases tired.
