---
title: "DefaultTrialRetire"
notion_id: 5c93137a-52ec-4ae3-bc40-0b137c6e0e41
notion_url: https://app.notion.com/p/DefaultTrialRetire-5c93137a52ec4ae3bc400b137c6e0e41
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://martinfowler.com/bliki/DefaultTrialRetire.html
tags: ["Product Management", "System Design / Software Architecture", "Article", "Martin Fowler", "English"]
---
[https://martinfowler.com/bliki/DefaultTrialRetire.html](https://martinfowler.com/bliki/DefaultTrialRetire.html)

Within each normal-sized team, limit the choice of alternatives for any class of technology to three. These are: the current sensible default, the one we're experimenting with as a trial, and the one that we hate and want to retire.

The conversation goes like this: We want to introduce a new messaging technology. How many do we have already in place? Oh we have three in active use, including one that's considered legacy and we're partway through migrating off and one that we experimented with previously but didn't gain traction. Ok, so we're at our limit now. If we want to add another messaging tech then we have two choices. Either migrate all of our apps off the legacy tech, or properly rid ourselves of the failed experiment. This is quite closely related to the idea of capping the number of Innovation Tokens in use within your teams.

At a team level these kinds of limits are relatively easy to maintain and discuss and act upon, because we have common priorities and ways of working and high trust, high bandwidth communication. At the scope of the whole organisation the challenge is similar, but getting alignment takes a lot longer and doing actual migration and consolidation work can take a long time - so we sometimes have to allow for more variation in technology. We also use different techniques to discuss and communicate the status of our preferred technologies.

An approach we use at MYOB to engage our whole organisation in broader decisions about technology is by publishing our own MYOB Technology Radar, following the format of the Thoughtworks Technology Radar. This approach of building our own radar involves taking input from all of our verticals and teams, and making a clear statement on what technologies we encourage teams to adopt, trial, or more importantly which ones to keep clear of.
