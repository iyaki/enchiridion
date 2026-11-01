---
title: "Technical Debt is over-used"
notion_id: 1b72b802-9b94-4a02-9b24-687ac7b7cbc8
notion_url: https://app.notion.com/p/Technical-Debt-is-over-used-1b72b8029b944a029b24687ac7b7cbc8
last_edited: 2023-06-20T13:43:00.000Z
source_url: https://peakd.com/hive-168588/@crell/technical-debt-is-over-used
tags: ["English", "Project Management", "Programming", "System Design / Software Architecture", "Article"]
---
The term "technical debt" gets thrown around a lot. Way too much, in fact. Part of that is because it has become a euphemism for "code I don't like" or "code that predates me." While there are reasons to dislike such code (both good and bad), that's not what the term "technical debt" was invented to refer to.

So what does it mean? There's several different kinds of "problematic code," all of which come from different places.

## The original source

The idea of less-than-ideal code making future development harder is as old as I am, at least. The analogy to "debt" comes first from Ward Cunningham in 1992:

> "Shipping first time code is like going into debt. A little debt speeds development so long as it is paid back promptly with a rewrite... The danger occurs when the debt is not repaid. Every minute spent on not-quite-right code counts as interest on that debt. Entire engineering organizations can be brought to a stand-still under the debt load of an unconsolidated implementation, object-oriented or otherwise."

— Ward Cunningham, 1992

As usual, [Wikipedia](https://en.wikipedia.org/wiki/Technical_debt) has a good article on the subject that is worth a read as well.

## The problem of debt

So what's the issue? First, the analogy is imperfect, and as Wikipedia notes, can often lead people to thinking about the code in the wrong way. Financial debt is a very quantifiable, measurable thing. Technical debt is not, and quantifying how much a given design decision hurts future development can range from very difficult to impossible. But just because you can't put a dollar figure on it doesn't mean it won't cost you money.

Second, colloquially a lot of engineers have taken to using the term "technical debt" too broadly. Any kind of problematic code gets called "technical debt" or "code debt," even if it wasn't done deliberately. It could be the code is just buggy, period, or it could be perfectly fine but for the wrong context, etc. These all require different approaches to address.

## Kinds of problematic code

I would break out "problematic code" into four different categories, only some of which would fall under the "technical debt" concept.

### Deliberate debt

Deliberate debt is, essentially, what the term originally referred to. It is code that is not-quite-right or takes shortcuts, and was written that way knowingly and deliberately.

That could mean:

- skipping tests or documentation.
- using known-slow algorithms in certain places.
- hard-coding some value or logic that we know will need to be configurable.
- Violating some separation of concerns to reduce development time.

And so on.

There's nothing wrong with deliberate debt, as long as you are consciously aware of it, and the need to revisit it, and plan ahead for it. It's the "and revisit it" that becomes an issue, as rarely do managers spare the time to clean up corners that have been cut.

As developers, we need to both advocate for addressing deliberate debt _before it is taken on_, and make the effort to allow us to dig out of it later. That could also take many forms.

- Leaving stub tests in place.
- Leaving copious comments about _why_ a certain piece of code was written the way it is, and what should be done to clean it up. (Because whoever gets that task will not remember what you were thinking at the time, including you.)
- Isolating the known-slow algorithm into a component (class or set of classes) that will be easy to swap out with something better later.
- Open an issue in your issue tracker with exactly what steps need to be made, and give it an appropriate priority. ("Appropriate" is always medium to high, not low.)

Etc. These signposts help to identify and track how much "non-visible" work needs to be done, and set you up to be able to do so.

### Accidental debt

I would divide deliberate debt from accidental debt. Accidental debt is what it sounds like: Code that is not-quite-right for some reason that wasn't intentional. That is not the same as a bug, or just sloppy code. (Those are separate categories.) Examples include:

- Not testing certain error pathways you didn't realize were possible.
- Using an algorithm you thought would work, but is actually not great in context.
- Introducing a hard-coded value you didn't realize needed to be configurable.
- Designing the code to have one instance of a thing, when it actually needs N instances of a thing.

Accidental debt code _does work_. It's not buggy, it's just not as flexible or robust as we want or need it to be.

Importantly, there's no blame here. No one did anything wrong, or sub-standard. The code passed QA for a reason. But it's still not-right.

### Contextual debt

Contextual debt is code whose not-rightness is a result of the context the code is running in changing.

For example, at the time it was written, the product requirements said it only had to send an email once a day or so, so the system was built on that (at the time accurate) assumption. Three years later, we're sending 20,000 emails a day because the product definition changed, and the design can't hold up to that much mail traffic.

Or when the app was built, CRUD logic was all that was needed. The company has grown into new markets, though, and is now subject to new regulations, and extensive logging of all actions is now a requirement. Event Sourcing does quite well at that, and would be better in the new situation, but there's already 50,000 LOC that assumes CRUD.

Importantly, as with Accidental Debt, there's no blame here. The original developers did exactly what they were supposed to do, in their context. The code was "right" by every metric. The challenge here is that the definition of "right" has shifted, and the code hasn't.

This can also include the libraries in use in an application. If you're building a PHP application today, for instance, and you use PHP 8.1 or 8.2 for it, you're fine. That's "right." If you're still running that in 5 years, that's not-right. The context will have shifted; PHP 8.2 used to get security support (making running it in production a-OK), but in 5 years it won't. In 5 years, running PHP 8.2 will be a form of contextual debt, just like running PHP 7.2 today is.

Similarly, industry standards may evolve. Ten years ago, building the UI of an application using CSS floats with negative margins and big downloadable CSS frameworks was standard, and what you'd expect a reasonably good CSS engineer to do. Now, though, we have flexbox and grid; they're orders of magnitude better than CSS floats and negative margin tricks. Those still work today; the site won't break for using them. But it still means new design changes need to fight with negative margins that are (and always have been) a pain to deal with, but rewriting the layout to use grid, even if it would be ten times better in the end, has never quite gotten the green light from management.

The people using negative margins in 2013 were not wrong! They were doing cutting edge work, and following [best practices](https://www.garfieldtech.com/blog/best-practices). But the edge has moved and [best practices evolved](https://platform.sh/blog/best-practices-in-deploying-web-apps-updated-for-2020/) and changed, and now we have their code to deal with.

### Sloppy code

This is not a kind of "debt." This is problematic code that isn't just "not quite right" or "not quite right anymore," it's code that was already wrong at the time it was written, and is still wrong today. As hinted at above, though, the definition of "wrong" is a moving target.

20 years ago, if someone wrote a PHP application that used a half-dozen global variables, no one would bat an eye at it. That was standard. If that code survives to today, it's not sloppy, it's contextual debt. It's still a problem, but not in the same way.

If someone wrote a PHP application that used a half-dozen global variables since around 2010, though? That's just sloppy code. The industry standards, the common understanding, and the available tooling and support had evolved to the point that we had better options available, and, in that context, basing your application on global variables is just bad design. And if someone built a global-variable-based app in PHP today, they'd probably get laughed at just as much as someone trying to use tables for web page layout.

## Important distinctions

In all four cases, we have code that is not-right, in the current context. In all four cases, that code inhibits further development in some way. In all four cases, fixing those issues generally has no user-visible improvement, so convincing management and product that they need to be addressed can be hard. (The harder it is, the more dysfunctional your company is.) Yet in all four cases, it does need to be addressed.

The superficial effect on the team is the same in all four cases: Something is not-right in ways that slow us down, and we need to fix it if we want to make faster (or any) progress. However, the psychological context is different. It's really easy for us as developers to assume any code that is not-right is Sloppy Code. While it certainly could be, it could also be contextual debt, deliberate debt, or accidental debt. We should resist the temptation to just classify it all as sloppy code and hate on the previous developers. Someone else will be saying the same about your code in a year. How and why is the code not-right? Was it a deliberate cut corner? Was it an edge case no one thought of? Has the industry just moved on?

Of them, deliberate debt is probably the easiest to address, as it is most likely to have a roadmap for addressing it already defined. In all four cases, however, the sooner it is addressed the cheaper it is to do. The longer you wait to address problematic code, the more it will taint other code built atop it that will also need to be fixed when you get around to addressing it.

Also important to note: "Code I didn't write" appears nowhere on this list. That someone else wrote it, or that it was written 10 or 20 years ago, does not automatically make it problematic or a kind of tech debt. Old code _might_ now be contextual debt, but its authorship has no bearing on that question.

## Fix it fix it fix it fix it

So how do you convince management to give you time to fix it? That's a whole other blog post (or book), which I don't have the time to write right now.

My general answer, quite honestly, is don't. Management doesn't get to decide how you get your job done. If a little (or a lot) [preparatory refactoring](https://martinfowler.com/articles/preparatory-refactoring-example.html) is the best way to get something done, just do it. Keeping the code in a healthy state is your job. If that means the estimate for feature X is a bit larger because you have to first pay down debt Y so that doing X is cheaper... then so be it. That's what the company pays you for: Doing your job well.

In general, I find Kent Beck's glib take to be the most compelling way to approach any software development:

> for each desired change, make the change easy (warning: this may be hard), then make the easy change

When do you clean up problematic code? When doing so will notably improve the thing you're about to do anyway.

For some definition of "notable", which is, of course, contextual.
