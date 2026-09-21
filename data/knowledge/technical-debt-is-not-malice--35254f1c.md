---
title: "Technical debt is not malice"
notion_id: 35254f1c-7d23-81b2-9a12-ebd2f51a342b
notion_url: https://app.notion.com/p/Technical-debt-is-not-malice-35254f1c7d2381b29a12ebd2f51a342b
last_edited: 2026-04-30T02:43:00.000Z
source_url: https://phpunit.expert/articles/technical-debt-is-not-malice.html
tags: ["English", "Technical Debt", "Software Development", "Empathy", "Programming", "Team Management", "Article", "Note", "phpunit.expert"]
---
![image](https://phpunit.expert/img/articles/technische-schulden-sind-keine-bosheit.jpg)

The image shows five wooden blocks with symbols representing interpersonal skills. A hand deliberately highlights the cube with the red heart and the inscription "EMPATHY". The warm, calm atmosphere emphasises empathy as a technical skill in software development.

> 

The above is a comment by Henning Schwentner on a statement by Michael Plöd:

> 

When we talk about “legacy code”, we usually mean code that has accumulated technical debt over the years: shortcuts that someone took, workarounds that were never cleaned up, design decisions that look questionable today. A common reaction is mockery: “Who writes code like this?”

That debt rarely comes from malice, laziness, or incompetence. It comes from circumstances we did not witness: tight deadlines, missing information, tools that did not exist yet. The question “Who writes code like this?” usually answers itself with: someone doing their best under exactly those circumstances.

## Stories hidden in code

The developer who picked a particular framework 20 years ago probably did not pick the worst one, they picked one that was considered a good choice back then. The team that chose the quick solution over the clean one usually did so because shipping that week mattered more than shipping in three weeks' time. The architect whose decisions look strange today made them under assumptions that no longer hold.

Dismissing such code also dismisses the people who wrote it, without knowing the constraints they were under. On top of the tone, this is a practical problem: from a position of judgement, you miss the information you would need to actually improve the system.

> 

Empathy here is not a vague virtue but a technical skill that we can learn and improve through deliberate practice.

Instead of “why was this done this way?” with an unspoken “... because this is nonsense”, a more useful question is “what problem was this solving? what constraints were in play? what did the person know that I do not?” That shifts your stance from judging to learning and most of the time, you do actually learn something.

## Write for the next person

The uncomfortable flip side is this: the code we write today is tomorrow's legacy. If we ask for empathy towards past decisions, we owe the same to whoever reads our code in five or ten years. That is mostly a writing task:

- Commit messages that explain the “why”, not just the “what”
- Pull request descriptions that surface the reasoning and the constraints
- Code comments where the decision is not obvious from the code itself
- Documentation that captures the current behaviour and the assumptions behind it
- Tests that show not only what happens, but which edge cases matter and why

I provided more detail on some of these artefacts, such as architecture decision records and technical debt records, in an earlier [article](https://phpunit.expert/articles/modern-php-development.html).

These are not overhead for overhead's sake. They are the only form in which context survives staff turnover and the passage of years. If we do it well, the person who runs into our code ten years from now will not have to decipher it first.

I have over 35 years' experience developing software, including almost 30 years working with PHP. I have also been developing PHPUnit for over 25 years. The knowledge I have gained during this time is reflected in my articles, but this is just the tip of the iceberg.

If you and your team want to achieve measurable progress, I would be happy to support you with [targeted advice and individual coaching](https://phpunit.expert/consulting-and-coaching.html). [Let's get talking!](https://phpunit.expert/en.html#contact)
