---
title: "High Variance Management"
notion_id: 995d5ddf-6346-4dc9-99d4-4ea19faeebe6
notion_url: https://app.notion.com/p/High-Variance-Management-995d5ddf63464dc999d44ea19faeebe6
last_edited: 2023-02-01T15:05:00.000Z
source_url: https://blog.sbensu.com/posts/2023-01-18-high-variance-management/
tags: ["Article", "English", "Line/People/Team Management", "Project Management"]
---
_There is a _[_math supplement_](https://blog.sbensu.com/posts/2023-01-18-high-variance-management-math)_ to this post where I build some intuition on the different probability distributions involved. If you are interested, read that later._

# Why variance matters

On Broadway, actors put on a great show eight times a week _consistenly_. If a dancer [actually breaks a leg](https://www.theguardian.com/stage/2012/mar/30/stage-actors-break-a-leg) the play has to stop. So, when Sue, the best dancer in the play, suggests adding a risky-but-awesome jump, the choreographer says no. It doesn’t matter how impressive the jump is if she gets it wrong half of the time. Changes that improve the show but cause mistakes are not worth it. You want high consistency, low variance.

Later, Hollywood decides to make a movie based on the play. Sue is cast and she performs the original choreography perfectly in the first take. But the movie director asks: “What else do you have?”. Sue remembers the risky jump. She tries twice before she lands it, but when she does, it makes the scene that much better.

The movie and the play are working under very different constraints:

- The movie only needs one good take, which will be seen by every viewer.
- The movie can discard any bad takes.

In Hollywood, you should be trying to do the very best in every take (see the appendix), even if that means that you’ll make many mistakes. You want high variance, low consistency.

In [Making Movies](https://www.amazon.com/Making-Movies-Sidney-Lumet/dp/0679756604/), director Sydney Lumet recommends lying to actors so that they loosen up and give you better takes. He says that some actors only dare to try more interesting things after they feel they have an acceptable take. Lumet recommends that you lie to actors and tell them that you have a solid take so that they loosen up and go for more. Lumet is teaching us to elicit higher variance from actors.

But not every part of movie production warrants high variance. Lumet also recommends that you plan shooting days down to 15 minute intervals. He even recommends that you check on the actors' chauffeurs: the best a driver can do is be on time but the worst they can do is to ruin a day of production. Lumet is teaching us to reduce variance from other processes.

Like movies, software projects have parts that require high variance and parts that don't. For most projects, the logging system can be off-the-shelf and predictable. But core parts of the product that require novel design should be as good as they can be. I've found this best illustrated in Ken Kocienda's [Creative Selection](https://www.amazon.com/Creative-Selection-Inside-Apples-Process-ebook/dp/B079DVT6VP/). Kocienda gives us a week-by-week account of how the first iPhone's keyboard was designed and implemented:

- Multiple teams were asked to work on the keyboard in parallel
- Each team made several prototypes
- Prototypes were demoed publicly so that everybody could learn from them
- This process was repeated until there was a winning prototype

The virtual keyboard was not only the first of its kind but also one of the most important UI interactions users had to learn with the new iPhone. Despite its importance, the keyboard's design was not carefully scheduled and "managed". Instead, it was forged via "creative selection", a much more chaotic process. The iPhone managers elicited higher variance from the overall team with that setup.

Managers at software companies rarely think this way. Our intuition and vocabulary for management trace back to [Scientific Management](https://en.wikipedia.org/wiki/Scientific_management). This was developed for factory work and was aimed at dampening variance to make factory workers’ output more predictable. Books like [High Output Management](https://www.amazon.com/High-Output-Management/dp/B08Z8K64S2/) draw from that tradition.

When developing new software products this "reduce variance" approach is the opposite of what you need for key components. Instead, we need High Variance Management.

# Recipes for variance

Feel free to stop reading here. What follows are some recipes to encourage high variance work from a short management career:

## 1. Gather the crew

### Recognize who can do high variance work

Here are a few markers that somebody is looking to do high variance work in a domain:

- They will surprise you with their opinions in that domain (you won’t always agree)
- They are working on this domain for a reason
- It is hard to get them interested in something else

You only need one or two people in this category. The rest of the team needs to contribute with good, dependable work, and not get in the way with unproductive opinions.

### Sell them on it

Not everybody wants to work on something that may likely fail. When recruiting for high variance, make that very clear that you are expecting a wide range of outcomes:

> If this succeeds, it will be huge. But that is low probability. It will most likely fail.

The right people for a high variance project will be thrilled by this proposition. And the rest will wisely choose something else.

### Don’t ask a high variance team to be consistent

A high variance team can only surprise you a couple of times a year. Every time you ask them to work on a project that has an expected outcome, you'll waste half of your yearly surprise budget. If the job requires consistency, give it to a team that is happy with consistency and capable of it.

### Don’t ask a consistent team to go for variance

Here are some signs that you’ve asked for variance from a team that’s not ready for it:

- They spend weeks on brainstorms.
- Instead of rethinking the product, they’ll fix bugs in the existing product.
- They spend their time philosophizing, i.e. “what does ‘good’ mean for this project?”. If the team doesn’t understand the goal from the get-go, they are unlikely to do a good job.

## 2. Frame the problem

Once the team is ready for high variance work, help them frame their projects so that they have enough room for their range to show.

### Ambition

Sometimes teams don’t understand that they could knock it out of the park, and instead they start by copying existing approaches. Snap them out of it:

- If you had to make the best X, how would you do it?
- What are the weirdest examples of X? Why aren’t we doing that?
- What is the state of the art? How does it suck?

### Make a variance budget

In software, like in movies, you don't need variance in _every_ part of the process. Keep an eye out for digressions into interesting but actually misguided investigations. Novel databases and "interesting" architectures are often not where you want the variance but that is often where software engineers want to spend their budget.

### Add constraints

Some open-ended problems are underconstrained. “Imagine the office of the future” is too broad to be useful. “Imagine the CRISPR lab of the future” or “Imagine the ideal software design studio of the future” might generate more ideas.

### Drop constraints

Sometimes the team is mentally blocked by a particular constraint: “we wanted to use GPS but it doesn’t work on desktop.” Drop constraints to unblock them: “Forget desktop, you can figure it out later”. Here are some canned phrases to lift the X constraint from their mind:

- Imagine you could ignore X, what would you do?
- Is there any upside to X? What if X was not a problem but our main feature?
- Are there any users for which X doesn’t matter?

## 3. Dampen the dampeners

Your job as a high variance manager is relatively boring: get rid of the dampeners[1](https://blog.sbensu.com/posts/2023-01-18-high-variance-management/#fn-1).

### Process

Rules like “you are expected to be in the office from 9-5” slow down night owls. Most rules reduce variance somewhere.

### Focus

As a project goes on, many ideas will pop up. Ideas shouldn't be automatically accepted, _or even fully considered_. Choose somebody that is both disagreeable (knows how to disagree) and has good taste to be the decision maker and hold them accountable for the final result.

### Attention and trust

Executive attention often kills creativity. Ideally, the executives are not interested in the project.

If that fails, reward risk-taking whenever you see it to establish trust. Separate the feedback about the result (“I don’t like that color” or "The numbers are weak") from the feedback about their approach (“I like that you are trying more colors, try even more!”). Hopefully, that will let the team take risks without the corresponding fear.

### Normalcy

If one team member constantly points out things as “weird”, something great but weird will get cut out. The more people you have, the more likely one of them will call things weird. Three ways to help with this:

- Don’t add too many people to the team.
- Reassign any team members that continue to doubt the project once it has started.
- Ask people to be precise when they offer constructive feedback: “What parts specifically are weird? What problems will it cause?”

### Vacation

Check people’s energy before starting a project. Make sure they are excited and well rested. If they aren’t, offer them a vacation before the project starts. Exciting work is done in chunks and breaks usually kill a certain momentum (even if necessary).

### Avoid premature deadlines

If you set a deadline before it is clear to the team what to build, they’ll do the obvious thing that fits within the deadline. Instead, give them time to discover what they should be doing.

I really don't know how much about the research phase but I do note that you shouldn't squeeze it into a calendar (see the appendix for Michael Nielsen's notes on that).

Only when the team has a plan, pick a deadline and focus them into executing. To tell if the team has a plan, you can ask individuals what they think the team should do next. Their answers will diverge when ideating and converge when they are ready to execute.

_Thanks to Devon Zuegel, Michael Nielsen, Kanjun Qiu, and Jack Dent for reading drafts and offering valuable feedback._

# Appendix

## Michael Nielsen's notes on managing creativity

> Reflecting on creative work: the manager's or engineer's orientation is toward a list of tasks; the creative researcher's is toward a list of emotional provocations and half-baked hunches.

It seems almost that I don't want a task list. I want a long list of weird-ass provocations which I feel strongly about.

Strong might be with almost any valence. Fear or anger often work just as well as inspiration, oddly!

And you don't check 'em off. You just keep going back for fuel. (Well, sometimes they lose their emotionla valence.)

Michael notes that one should manage a research agenda very differently than an engineering project.

## Real actors on variance vs consistency in Off Camera

In the [Off Camera episode with Hank Azaria](https://offcamera.com/issues/hank-azaria/listen/#archive), Azaria explains how much “perfection and consistency” got in the way of a great performance. Around minute 55m, he talks about an actress that he admires:

> She never did it the same way twice, most of her takes are bad, she didn’t remember more than three lines at a time. She was all over the place.

I’m doing it as if I am in a Broadway stage, every time it is good, consistent, it is in there. Well done, crisp!

And then, in the editing room, she would wipe the floor with me.

In all those takes didn’t matter because there were two or three takes that were breathtaking. She knew what she was doing, it takes a lot of courage.

In the [Off Camera episode with Edward Norton](https://offcamera.com/issues/edward-norton/listen/#archive), Norton talks about Motherless Brooklyn and how he managed to act and direct at the same time. Around 23m:

> [When acting and directing] ... my concern was for the thing that actors do for each other. My hedge on that really was to get actors that were theater experienced, trade-craft pros. And by theater experience I mean people who, as part of their training and professional life, come with the totality of their performance ready. Ready to play, ready to do things. But honestly, text-ready. Like off-book, no fucking around with "I may not know all the lines yet so you are going to cut around that". I had to have people who were ready to rock and who I knew didn’t need a lot from me to get going.

My interpretation from that passage is that Norton knows how much work a director does for actors, and how much actors feed of each other during scenes. Given that it is his first time acting and directing, he can't really afford to work with actors that need a lot of help to produce good takes. In this case, he is looking for consistency, which he knows theater-trained actors can provide.

## Footnotes

1. Paul Graham enumerating dampeners in [The power of the marginal](http://www.paulgraham.com/marginal.html): "the selection of the wrong kind of people, the excessive scope, the inability to take risks, the need to seem serious, the weight of expectations, the power of vested interests, the undiscerning audience, and perhaps most dangerous, the tendency of such work to become a duty rather than a pleasure."[↩](https://blog.sbensu.com/posts/2023-01-18-high-variance-management/#fnref1)
