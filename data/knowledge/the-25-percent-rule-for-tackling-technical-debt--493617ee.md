---
title: "The 25 Percent Rule for Tackling Technical Debt"
notion_id: 493617ee-ad2b-4efd-b7e1-723131481242
notion_url: https://app.notion.com/p/The-25-Percent-Rule-for-Tackling-Technical-Debt-493617eead2b4efdb7e1723131481242
last_edited: 2023-02-22T18:42:00.000Z
source_url: https://shopify.engineering/technical-debt-25-percent-rule
tags: ["Article", "Shopify Engineering", "English", "Programming", "Project Management"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Let’s talk about technical debt. Let’s talk about practical usable approaches for actually paying it down on a daily, weekly, monthly, and yearly basis. Let’s talk about what debt needs to be fixed now versus what can wait for better planning.

But first, let’s talk about talking about technical debt.

# Catching Smoke

Talking about technical debt is like trying to catch smoke. When you think you have the problem described and addressed, someone can say “but what about” and now the whole solution seems to fall apart. For example, when someone is talking about larger debt projects, like big refactors or rewrites, and someone says “What about the small improvements people can make daily?” So, then, that’s talked about for a bit, and someone says, “But we also need time to make bigger refactors.” And the cycle repeats. Even if we catch a little bit of the smoke, enough squeezes out and we have to catch that smoke, but then we lose our grip on the first batch, and in the end people are just frustrated.

Recently, my team was talking about technical debt and were discussing remediations. I noticed a few patterns in how we were talking that I think are worth sharing.

Of all the ways we spoke, I heard four types of debt primarily described:

1. Yearly Debt
2. Monthly Debt
3. Weekly Debt
4. Daily Debt

# Yearly Debt (Rewrites)

_Yearly Debt_ is the kind where after lots of conversations, someone concludes a rewrite is the only solution. Sometimes a rewrite may be the only solution. Sometimes you may have a Ship of Theseus problem on your hands where you need to slowly and methodically replace parts until the system is the same but different.

Sometimes, though, this isn’t really _debt_. It’s possible that your dilemma is the result of growth or changing markets. In that respect, calling it debt does a disservice to our success, and distracts from solving the problem of growth. Other times you may need to realize your system doesn’t need to be Google scale and will work fine for quite some time without intervention.

Nevertheless, it’s easy to point at this sort of dilemma as _debt_, which mistakenly implies that some previous decision or mistake is responsible, when it’s possible your position is the result of your own success, or, you know, the world just changing.

# Monthly Debt (New Projects)

_Monthly Debt_ is debt that takes months to pay off. While working in a system we can clearly identify the problem and may have some vague sense of a solution. It’s possible we’ve even done some prototyping as a proof of concept, but any serious implementation requires a small team and several months. Or maybe we shipped a prototype that had surprising success and now we need to clean up the code and stabilize it. Realistically this type of _debt_ can take over a year to address, but it’s different in scope from The Yearly Debt.

Monthly Debt likely requires making a case to senior product or engineering leaders that the work needs to be done. This may be a frustrating point for some, as we should feel empowered to make our codebases better as we go along, but at this scope the equation changes. Instead, we need to take time to consider what the cost and benefits are of addressing the _debt_, and plan accordingly by assigning people and teams with clear goals. Or, by doing nothing and accepting the tradeoffs in that.

# Weekly Debt (New Cards)

_Weekly Debt_ is the debt that can be solved by adding a _card_ or _issue_ to a sprint. They’re a discovery or side effect of the work we’re doing. One great example would be if we go to implement some new code and find that we have a more efficient or cleaner way of doing it. We could even refactor some adjacent code to use our new code, again simplifying things and improving our internal APIs. However, including those changes in our current PR would make review very distracting and isn’t necessary to ship our changes.

This kind of debt ought to be worked on by someone, but it doesn’t have to be worked on by the discoverer. Instead, we can add it to our sprint, or at worst, our backlog. Nevertheless we don’t want to leave it behind as it creates two code paths that need to be maintained, thus adding to complexity.

Weekly Debt and Daily Debt face similar problems, so let’s look at that one before addressing them further.

# Daily Debt (Tidying)

_Daily Debt_ shouldn’t be called _debt_ by many standards. Daily Debts are the opportunities we pass by to make the code we’re working in better. Maybe while reading some code, we struggled for a long time to understand what it was doing. After considerable time, and maybe a walkthrough from the original author, we have an idea of how to refactor it.

We may attempt to do this for an hour or two. Maybe we succeed and include it in our PR as a change we want to include alongside ours. Or maybe we fail to implement any significant changes, but realize extracting a few methods with clear names would go a long way to prevent the same confusion we experienced. Daily Debt never feels like debt. Instead, it’s what makes code feel not ergonomic. It’s what makes code _annoying_ to work with.

# Why Don’t We Fix Debt?

I think there are two primary reasons we don’t address these types of _debt_.

First, we are afraid of wasting time. Even if we are in a great development culture that values being good stewards of our code bases, we still face pressure to ship. Each team and individual feels these pressures in different ways, but knowing how to reconcile the two is difficult. Should we use five percent of our time? Ten percent? The problem with those numbers is they lump together the different kinds of debt, all of which require different amounts of time to address. Additionally, typically addressing Monthly and Yearly debt is done in a very different manner than Daily and Weekly Debt.

Secondly, it’s also possible people don’t address Daily Debt because it’s boring. Refactoring doesn't often get the praise it should. We celebrate the launch of new products. I also often see code deletion celebrated. We offer the platitude, “The best code is no code.”

It can be very hard to encourage people to address Daily and Weekly Debts if they fear their time wouldn’t be used well or they won’t be recognized for the positive effects. Particularly, if you don’t feel yourself to be a strong engineer, you may feel it’s not worth the time to improve some code, as you’re not really sure what to do.

So how much time should we spend?

# The 25 Percent Rule

It may sound like a lot, but anything less does not allow for the various types of debt to all be addressed. Let’s break it down.

## Ten Percent for Daily Debt

We should spend ten percent of our time (four hours a week) making the code we’re working in better. That doesn’t mean we’ll always spend those four hours every week, but engineers should be told they’re _allowed_ and _flat out encouraged_ to use up to four hours a week if they want to try and improve code in an area they encounter if they feel it would benefit from it.

Again, this shouldn’t count as _extra_ time being spent on technical debt. It’s an enshrining of the practice of refactoring into the team ethos and culture. We must emphasize this is on code that we’re already working on in our regular daily work, it’s not going searching for random _bad code_ to fix. It’s being empowered to improve things while we are working on it.

## Ten Percent for Weekly Debt

Additionally, we should spend another ten percent of our time on Weekly Debt. Now, this may not happen on the individual level. But as a team, we should feel empowered to put items on our project boards that can be worked on. It’s possible four hours each on a team of four (16 hours) may result in one team member working on a card or two over two days.

This allows for bigger changes to happen. It’s possible that always working on these issues immediately would be very distracting for a team and would hurt a person’s momentum. Instead, we carve out time for someone else to gain momentum on making these changes.

## Five Percent for Monthly and Yearly Debt

Finally, the remaining five percent is spent between Monthly and Yearly Debt. This amounts to two hour long meetings a week somewhere to talk about planning. Obviously, this wouldn’t be enough time to address those things. Instead, it’s on the intellectual work of figuring out what we know about the problems and prioritizing if something even needs to be worked on at all, and if so, how and when it can be done. These may already take place in the form of backlog meetings or long term strategic planning, which should take into account the current state of the system.

# Fixing Debt is Done Through Culture

Addressing technical debt is rarely about making time for large fixes. It’s about setting strong examples for improving code in our daily work. It’s about celebrating the ability to refactor code to make it easier to work with. It’s about the ability to see when code complexity and atrophy change our codebases from a nuisance to a hindrance. It’s about understanding when to choose the good over the perfect and knowing when it’s time to go back and make it better.

**John DeWyze** is a Staff Developer at Shopify. He has a passion for healthy teams, psychological safety, and the human side of tech. He is a proud dad, enjoys playing bridge and softball, and is an aspiring pizzaiolo.

We all get shit done, ship fast, and learn. We operate on low process and high trust, and trade on impact. You have to care deeply about what you’re doing, and commit to continuously developing your craft, to keep pace here. If you’re seeking hypergrowth, can solve complex problems, and can thrive on change (and a bit of chaos), you’ve found the right place. Visit our [Engineering career page](http://www.shopify.com/careers/specialties/engineering?itcat=EngBlog&itterm=Post) to find your role.


