---
title: "The only developer productivity metrics that matter « genehack.blog"
notion_id: 30d54f1c-7d23-81ee-8169-c1616f0fefc7
notion_url: https://app.notion.com/p/The-only-developer-productivity-metrics-that-matter-genehack-blog-30d54f1c7d2381ee8169c1616f0fefc7
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://genehack.blog/2026/02/the-only-developer-productivity-metrics-that-matter/
tags: ["genehack.blog", "English", "Productivity", "Agile", "Team Management", "Project Management", "Software Development", "Article"]
---
I’m not sure if it’s because I’ve been closely following [Agile is Anarchy](https://agileisanarchy.com/), or some recent “how do we work” conversations at work, or what, but Monday morning this rant popped into my head, almost full blown, and what do I even _have_ a blog for if not for posting rants about software development?

So, here’s the deal, just in case you’ve forgotten or were never told: pretty much every single way management tries to measure software team productivity is bullshit. You’re not measuring what you think you’re measuring, generally; what you _are_ measuring is how good your dev team is at gaming your metrics. (Spoiler: they’re probably going to be extremely good, especially if they’re experienced.) I assume everybody has heard the “we paid bonuses for fixing bugs” story — or if you haven’t, you can probably extrapolate — but the key thing here is: almost every single thing you can measure, the devs can game.

Here are the only two things you should be worried about when it comes to the question of how productive a given team of software developers is:

1. How often does the team _routinely_ ship new versions of the software they build?
2. How often do things break when the team ships a new version?

**THAT’S **_**IT**_. Story points per sprint doesn’t matter; lines of code, doesn’t matter; Jira tickets closed, for the love of everything, do. not. matter.

Are you shipping often? And note, the word “routinely” is emphasized in that question because I’m not talking about hot fixes or “crunch time”, I’m talking about your regular cadence.

Do things keep working when you ship? And I’m not talking about some minor bug here, some typo in an error message or a web page layout misalignment measured in pixels, I’m talking about “production is down and we don’t know for sure when it’ll come back”-scale bugs.

Those two metrics actually matter, because those are the things that have direct impact on your users, your customers — and if your answers to those two questions aren’t “at least once a week, if not more often” and “basically never”, well, then, now you know what you need to be figuring out how to fix.

If you _are_ shipping at least once a week, and things just keep working — if you have to stop and think about the last time prod _really_ broke — then you can start to worry about the third, bonus metric:

1. How often does a new version of the software the team builds spark actual joy in the people who have to use it?

How you go about actually achieving that is an advanced topic, and maybe I’ll write more about that next, but the short version is, if you’re not making at least one user a little bit happier with every new version you ship, you need to reconsider how you’re picking what you decide to work on — and you probably need to spend more time talking directly to those users (_not_ to your product person, although you should be talking to them more too, probably).

Enjoy this post? Or maybe you have a question? Hit me up on [ Mastodon ](https://dementedandsadbut.social/@genehack) and let's talk about it.

Back to [home](https://genehack.blog/)…
