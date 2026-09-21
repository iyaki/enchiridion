---
title: "Supporting, influencing, and leading as a security practitioner"
notion_id: cba0148d-a70b-4e8a-aa33-54d8c9ae8fce
notion_url: https://app.notion.com/p/Supporting-influencing-and-leading-as-a-security-practitioner-cba0148da70b4e8aaa3354d8c9ae8fce
last_edited: 2026-09-21T17:43:00.000Z
source_url: https://leaddev.com/leadership/supporting-influencing-and-leading-security-practitioner
tags: ["English", "DevOps", "Article", "LeadDev"]
---
[https://leaddev.com/security/supporting-influencing-and-leading-security-practitioner](https://leaddev.com/security/supporting-influencing-and-leading-security-practitioner)



<!-- unsupported block: link_to_page -->

You have 1 article left to read this month before you need to register a free LeadDev.com account.

Pushing security practices is a bit like selling insurance; people know they need it but nobody enjoys the associated costs.

I learned this the hard way when I switched roles from developer to security advisor. I was given security responsibility over a set of products without being made part of their development team. I was to help them improve their security posture from the outside.

My plan was quite simple – and naive. I’d talk to the decision-makers; they’d see the value inherent in employing security best practices; we’d identify the gaps in the current status; make a plan to fill those up; they’d incorporate those in their objectives; and everyone would be happy. I’ll pause here and wait until you stop laughing and compose yourself. I’m sorry if you spilled your drink.

In the years since I’ve advised on many products as a lead security architect and have learned a lot about what works and what doesn’t when it comes to building relationships with engineering teams and getting their buy-in. Here I’m going to walk you through a few different approaches I’ve tried over the years – starting with the ones where I learned by failure – before sharing how I finally began to find success.

### My initial approach

### Caring more about the project than the people responsible for it

When I first started out as a security advisor, there was a particular team that became my own pet project. I cared about what they did, and how and why they did it. I thought my background as a developer fully equipped me to understand their challenges and how to solve them. I believed I could think like them, understand them, and influence them purely because I had been them.

I was wrong. Eventually, I learned that an email was being passed around the team: ‘Do not respond to our security advisor anymore. Every time someone answers a question, he finds further problems’.

I quickly came to learn that you shouldn’t care more about something than the people responsible for it. If you’re not the one doing it, don’t expect too many things to be done the way you think is best just because you’re saying so. There are many competent and talented devs on the team; your role is to work with them, not to dictate. They won’t take kindly to it, and you won’t be able to help them.

### Bugging people

I thought that by regularly telling people, ‘secure code is quality code’, I could get their buy-in. But this didn’t work. When I said to folks, ‘you have to do X, Y, and Z because not doing them hurts your security posture’, their answer was the same: ‘We’ve not had a problem until now and we’ll deal with it when we do’.

I kept poking at them: ‘Your authorization is broken’ ‘you’re not storing your secrets securely’; ‘that authentication could be done better’. I was basically a human static code analysis tool, with a rather smaller false-positive ratio but with a much louder alert system. All this cajoling, coaxing, and bugging resulted in a few small advances but didn’t help to create the significant organizational change around security that was needed.

### Taking a hands-off approach

Over time, I started moving towards what I’ve recently heard Brook Schoenfield, master security architect, describe as radical support: ‘I wish for the best outcome for you, no matter if I can join with your effort or agree (or not) with it. If I can’t actively support your effort, at least I can get out of your way’. Brook goes further: ‘radical support means allowing people the opportunity to make mistakes and then to learn from them. It implies I reserve power and influence for preventing catastrophe rather than getting my own way’.

Taking a step back and allowing people to learn from their mistakes is something that should always be part of a support approach, but in my experience going too far with it created some problems. As I was trying to influence existing teams and hierarchies from the outside, being too hands-off meant I ended up alienating myself from the group. My perceived absence pushed people away rather than bringing them in. Radical support is a great leadership tool in general, but it wasn’t enough when trying to lead from the sidelines.

(I must make it clear, I’m referring only to my own experience with radical support. Brook’s has been different and more successful than mine, based on his own style of mentorship, support, and team building.)

### Practicing empathy

Adam Shostack, thought leader in threat modeling, posits that security folks ‘need soft skills such as active listening, respect, and assumption of good intentions’. He also later adds ‘patience’ and ‘collaboration’ to the list. It’s by using these skills that we can help and influence a development team. You can be a great technologist, but if you can’t deliver on these soft skills, you won’t be an effective one. The problem is that these skills don’t translate into a clear strategy or plan of action, but they are definitely necessary when delivering a strategy or a plan.

This brings us to my fourth approach of practicing empathy. Empathy is commonly described as the ability to relate to the emotions of others. Exercising empathy was my attempt to use my soft skills and relate to the developers I was working with. We commiserated together around how we really want to do the right thing if only management would allow the time and resources. But alas, they wouldn’t, so we didn’t. I recognized the body language of the people I interacted with when we hit a particular piece of the design that was open to many avenues of attack: ‘We know,’ it showed, ‘but we can’t do anything about it.’

Practicing empathy helped me to share in the problem, but it didn’t seem to help me solve the problem.

### Where had I been going wrong?

My attempts at empathy didn’t seem to work. Radical support alone and treating the team like my own personal project didn’t work. Bugging the team relentlessly led to a few small changes but not the large-scale organizational cultural shift around security we needed.

My only option was to interrogate where I’d failed and why. When I reflected, I realized that my attempts at empathy perhaps weren’t as effective as they could’ve been. I noticed my empathy had pretty quickly turned into commiseration. I’d been relating to what the team was feeling because I’d been in their shoes. But because I was sharing in their feelings, I too very quickly became cynical and started to believe these were the cards we’d been dealt and there was little we could do to change things.

Once I’d identified the emotional trap I’d been falling into, I read up on empathy to try to understand where I’d been going wrong. This was a big turning point, and what I learned guided me towards a sliver of success.

### The missing piece: Cognitive empathy and perspective-taking

I have the luck to be married to a professor with a PhD in organizational studies. She directed me to research empathy, where I learned it is classified into affective empathy, somatic empathy, and the one that is more appropriate for our discussion, cognitive empathy: ‘the capacity to understand another’s perspective or mental state’.

I realized I’d been practicing affective empathy rather than cognitive empathy. I was simply trying to empathize with and feel what the developers and managers were feeling, rather than working to adopt their perspectives. This meant I was identifying their situation but not understanding it. Naturally, we all felt frustrated: them because I was pointing at things they didn’t have the will, power, time, or resources to fix, and me because I wasn’t able to impact on how things were being done. I had to move to cognitive empathy.

By itself, cognitive empathy subdivides into three levels:

- Perspective-taking: adopting or relating to another’s perspective
- Fantasy: identifying with fictional characters
- Tactical/strategic empathy: the use of perspective-taking to achieve a goal

Perspective-taking is the one that looked useful to me: by adopting the perspectives of the decision-makers, we can grasp how they’re likely to perceive our plans and attempts at persuasion. In this way, we can position ourselves for the best ‘sell’ of our security needs (there are plenty of resources out there explaining the link between perspective-taking and issue selling). Notice that in this context, a decision-maker is also the most junior developer, who has to make active choices when considering the time and resources required by security best practices, has to invest time in training and considering available options, up to the C-suite who has to invest in security tools.

With this in mind, there are a few things we can do to successfully get buy-in from developers and product owners:

- Tailoring the pitch Instead of doom and gloom, appeal to the team’s own motivations. For example, rather than saying, ‘if you don’t use static code analysis the product will fail horribly’, position it as, ‘static code analysis when used correctly can reduce the time it takes between identifying the issue and correcting it, without the developer needing to wait until the issue is found in post-deployment testing’. This translates into time and costs savings, which is much more persuasive. Don’t focus on the importance of your goals (i.e. improving overall security posture) but the impact on their goals (i.e. saving time and money and building better products).
- Managing both sides of the conversation Consider the different points of view folks might have about your plans. For example, ask yourself if the security you’re suggesting is too much for a particular system, or if the security is hindering the user experience. Product managers will value these issues above the abstract concept of security.
- Suggesting solutions, rather than giving instructions Make your recommendations in the spirit of, ‘let me help you secure the thing you built’, rather than, ‘you need to do this to fix the thing you built’. You’ll see that folks are much more responsive.
- Finding the right timing to push Take time to understand the constraints the team is operating under and respect them. Understand that sometimes this will mean a compromise on your side, and that’s okay. The important thing is working with the team to improve security posture within their time constraints.

Supporting, influencing, and leading as a security practitioner certainly comes with its challenges. But it is possible! In my experience, it takes a combination of these approaches to build valuable allies who will support your work down the road. You may not get what you want, but you just might get what you need.
