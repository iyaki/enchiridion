---
title: "Seven Shipping Principles"
notion_id: f19dd237-4922-49fb-abaa-5c6e470ba438
notion_url: https://app.notion.com/p/Seven-Shipping-Principles-f19dd237492249fbabaa5c6e470ba438
last_edited: 2026-09-21T17:41:00.000Z
source_url: https://37signals.com/seven-shipping-principles
tags: ["37 Signals", "English", "Programming", "Producer (Individual Contributor)", "Product Management", "Project Management", "Principles", "Article"]
---
[https://37signals.com/seven-shipping-principles](https://37signals.com/seven-shipping-principles)

Core fundamentals that inform how we go about building — and shipping — great software at a sustainable pace.

## We only ship good work

This might sound self-evident, but in reality there are all sorts of pressures that might lead you to ship mediocre or even iffy work. We promised the feature! We spent a lot of time on it! It’s not horrible! From the user’s perspective, it’s actually sorta fine?

None of these excuses are adequate reasons for us to ship anything but good work at 37signals. This means solid implementations in terms of design and programming, CSS and JavaScript, Ruby and operations. Work we’d be pleased to have to revisit 9 months later without a sigh of dread.

Most things we work on at 37signals will be around until the end of the internet. So whatever short-term gain you can capture by shipping short of good, you’ll lose when you have to work on it again later (and eventually you will!).

This isn’t just about productivity, now or in the future, but also about pride in the work. The delight that comes from working with code and design that’s properly built. It’s hard to put a specific price on that, but easy to recognize the smile it puts on your face.

This is not a general license to gold-plate everything. We intentionally constrain ourselves through the cycles in Shape Up, such that we don’t end up spending 2 months on stuff that warranted 2 weeks worth of work. Not every batch of work is going to be a 10/10. But if it’s less than 8/10, it probably shouldn’t go out the door. If it’s less than a 7/10, there’s no way it should go out the door, except as an emergency patch you immediately return to cleanup.

Shipping good work takes discipline and sacrifice. We will afford ourselves both.

## We ship when we’re confident

The reason we do automated and exploratory testing is so we can ship work with confidence that it won’t cause major problems. But exactly how much confidence is required depends both on the team and the criticality of the problem. It’s a measure to gauge in relation to the specifics of the work, not mindlessly delegate to a standardized process.

If you’re working on a small change that’s of low criticality, you don’t need the same due diligence as if you’re working on a big change of high criticality. That might seem obvious, but it’s easy just to run the familiar process, and forget to check whether it makes sense for the particulars.

High criticality work involves anything that mutates or munges data. If you can lose data, it’s high criticality. It also involves anything that makes larger changes to existing and big features. If you’re rewriting how todos in Basecamp work, you better be pretty sure you’ve checked everything twice or thrice. If a bug makes it to production on high criticality work, it’s a big deal.

Lower criticality work is anything that merely changes the presentation of existing data, or works on entirely new sets of data, or deals with screens or features off the critical path. If a bug makes it to production on lower criticality work, it won’t be a big deal. You fix it and move on.

It’s up to the team to gauge whether what they’re working on is high or lower criticality or somewhere in-between. It also depends on who works on it. Someone well-versed in the domain? A lead with 10 years of experience? They’ll deem more things lower criticality than someone who just started recently. This is natural and right.

If you need exploratory testing done by QA to feel confident that your work is solid enough for the criticality at stake, that’s what you should do. Even if it means waiting a bit to be able to ship. But if you’re confident that the work is likely to be good enough for the criticality at stake, you should just ship, stick around to deal with issues, and keep moving forward.

You gauge your confidence, and if in doubt, check with someone more senior. The process, the tools, and the testing are there to increase your confidence. Not to tell you when it’s time to ship or not.

## We ship when the work is finished

It’s rare here to keep work that’s essentially finished behind a feature flag for long if at all. We don’t start a lot of highly speculative projects. If it’s been kicked off, it’s because we’ve decided we need it, and that means it’s going out the kitchen when it’s cooked.

This means we’re not sitting on a large inventory of possibly-maybe projects that linger behind a feature flag until something magical happens. We clear the decks so we can get on with new work.

That runs the risk that we’ll occasionally ship a project that isn’t quite right once we see it in the wild. That’s okay. If it’s bad enough, we’ll fix it then. We don’t have billions of users. They won’t die or leave in droves just because we launch something that perhaps could have been better if it had been done a different way from the start.

The momentum of shipping good work has a quality all of its own. We don’t halt the momentum unless it’s properly bad, we don’t linger around second-guessing ourselves ad nauseam. The show must go on.

## We own the issues after we ship

Clean up your own mess if you make one. Pay attention to the error tracker for a while after launching. Be the first port of call for support when they get feedback from customers. If you did the work, you have the context, and you should put it straight if it’s crooked.

This sometimes means that you will have started something new and then get pulled away to fix something old. If it’s a substantial issue, that’s just what it is. You put the new on hold while you fix the old.

But you also have to keep the line on what “substantial” means. Almost everything we ship is met with a range of feedback. People will want more or less or different or better or whatever no matter how perfect the work. You can spend an eternity refining and reworking even the smallest change if you let yourself. We don’t let ourselves.

So you’re on the hook to make it right, to make it work as intended, but all the feedback that goes beyond that is filed for consideration against all the other ideas and claims for our time.

Ship it, fix it, forget it.

## We don’t ship if it isn’t right

It can be frustrating when work is held in the 11th hour because an issue or an insight that just popped up gives us pause. Why didn’t you mention this three weeks ago? Yes, it’s indeed better to get feedback early, but often that’s just not possible, because the issue or the insight isn’t apparent until it’s go time. That’s when you really check your gut, and the truth spills out.

But however frustrating it is to hit pause for a moment until the gut is settled, it’s far more frustrating to ship something that isn’t right to everyone. Once the work is in the wild, it’s hard to put it back in the box.

It’s more important that what we ship is fundamentally right than letting stuff go out the door because it’s socially easier or doesn’t rub anyone the wrong way. It’s better still if we can get things right early, but when reality denies us that opportunity, we still pull the cord and stop the train if it’s not right.

Everything that goes out has to easily pass the dual questions of Is It Right? and Is It Good?

## We ship our collective best effort

Catching quality issues in implementation, design, or concept is ultimately everyone’s domain. That doesn’t mean all who work here get to determine whether something is important enough to halt shipping or even be addressed. But it does mean that everyone looking at the work is eligible to raise reservations or ideas for improvement.

It’s natural to get a bit defensive when someone points out a perceived deficiency. Don’t fight that feeling, but let it float. Use it to interrogate the issue from both sides, the pro and the con. You don’t have to roll over on every objection, but you do have to consider them.

And if the feedback comes from someone more senior than you, you should default to believing they probably have a point, because they’ve seen more and done more. That’s why they’re senior. It doesn’t mean they’re always right, but the odds are decent that they are, and you’ll learn faster if you take the default stance of the beginner’s mind.

And if the feedback comes from someone more junior than you, make it a teaching moment if you choose to skip the suggestion. Why isn’t this relevant, right, or proportionate to deal with right now? And if the suggestion is adopted, celebrate together.

Regardless of where the insight originates, it’s the higher quality that ultimately sets the bar. If what we’re working on could be better, simpler, faster, and we have the time to make it so, that’s what we’re going to do.

## We ship to our appetite

One of Shape Up’s key concepts is the appetite. Projects are kicked off on a premise that they’re only worth doing if a good version of their pitch can be done within 2, 4, or 6 weeks. Like a great company can be a terrible stock at an exorbitant price, so too can a great pitch become a mistake if it’s pursued far beyond the appetite.

The time constraint imposed by the appetite is meant to force trade-offs and concessions. To curb the ambition that naturally turns every idea into a project that drags on forever by people drawn to perfection.

You’ll nearly always hit the time constraint with more ideas, more minor issues, and more polish to do. That might feel frustrating in the moment. If only I had two more weeks! But if you had two more weeks, chances are you’d just expand your ambitions accordingly, and you’d wish for two more weeks in addition to that at the end.

Constraints force us to make choices and rank what we’d rather have if it’s “or” not “and”. They serve as an objective way to force other stakeholders to accept the art of the possible. Everyone can come up with ideas for more, but it’s much harder to decide on what would you rather if you can’t have both. What we usually need are substitutions, not additions.

Feature creep and blown estimates are the industry standard. Our standard is that we ship the best work within the time we’ve given it, and we hold our heads proud to the compromises that entail.
