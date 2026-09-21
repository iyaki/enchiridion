---
title: "Your Code Doesn't Have to Be a Mess"
notion_id: 900593e7-decc-4ab9-b558-c60439b44c9d
notion_url: https://app.notion.com/p/Your-Code-Doesn-t-Have-to-Be-a-Mess-900593e7decc4ab9b558c60439b44c9d
last_edited: 2026-09-21T17:40:00.000Z
source_url: https://danielsieger.com/blog/2022/07/25/your-code-doesnt-have-to-be-a-mess.html
tags: ["Daniel Sieger Blog", "English", "Programming", "Article"]
---
[https://www.danielsieger.com/blog/2022/07/25/your-code-doesnt-have-to-be-a-mess.html](https://www.danielsieger.com/blog/2022/07/25/your-code-doesnt-have-to-be-a-mess.html)

Jul 25, 2022

> Fools ignore complexity. Pragmatists suffer it. Some can avoid it. Geniuses remove it. - Alan Perlis

If you’ve been developing software for a while, you know that code has this natural tendency to turn into a mess. Keeping software simple over time is a challenge that keeps me thinking. My last post left you hanging without much concrete advice. This time I will outline a few high-level strategies to keep software simple.

## Define Clear Goals

Define clear goals what problem your software is trying to solve. Write them down, keep them visible. Always keep them in mind when making decisions about adding a new feature or accepting a proposed change. Do one thing and do it well. Don’t try to solve a myriad of problems at the same time. Follow basic Unix philosophy.

## Setup Constraints

Define your non-goals as well: What problems does your software explicitly not solve? Constrain the scope of your software. Setup technical constraints on programming languages, libraries, or models of computation you are using. Embrace those constraints and stick to them. They will force you to think harder about the problem at hand. Get creative within the bounds of your constraints. Working with constraints can foster creativity and innovation. However, be aware when you need to bend your code too much.

## Say No

There is almost always an infinite demand for new features and functionality. Learning to say no is crucial to keep your software from growing wildly into too many directions. This can be super tough or downright impossible, notably in a commercial context. But even then: Do what is needed, but not more. Be aware of feature creep. But also be aware of when saying no would be unreasonable.

## Eliminate Waste

Regularly pare down your code, and be relentless while doing it. Remove any unused features, dead code, debugging helpers, or prototype code used during development. Version control has your back. Eliminating waste is a core principle of lean software development. Borderline but fun: Only accept PRs that reduce net code size for a while. Be aware of code obfuscation though.

## Minimize Dependencies

Be picky about adding any dependencies. Don’t treat code re-use as a holy grail. Carefully consider the pros and cons. Dependencies might break, disappear, turn into garbage, or become a security risk. Consider to vendor by default. If the functionality in question is crucial for your core business: Consider doing it yourself.

Obviously, the above strategies are not equally applicable in all settings. Your mileage may vary a lot between commercial development, an open source project, or creative coding.

Similarly, striving for simplicity can be much harder if you are working on software that tackles complex problems. Some problem domains have an inherent complexity that can not be negotiated away. However, even in those cases you can strive for simplicity in the building blocks of your complex system.

If you are developing software on a team, there is also an important social dimension. Some folks just revel in complexity, and it can be tough to develop a sense for simplicity with them. I’ve got the suspicion that the C++ crowd is particularly vulnerable to this disposition, but that’s another story.

Finally, keep in mind that simplicity is not a value by itself. It is a means to an end: produce better software, and be able to maintain and evolve it over time.

## References and Further Reading

- Simple Made Easy by Rich Hickey
- Unix philosophy
- Lean software development
- Epigrams in Programming by Alan Perlis
- Clean Code by Robert C. Martin
- A Philosophy of Software Design by John Ousterhout

The discussion on HackerNews makes an interesting read, thanks to all participants!
