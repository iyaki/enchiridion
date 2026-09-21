---
title: "Lessons Learned while Working on Large-Scale Server Software"
notion_id: 6eb03bd1-f616-4159-89e5-37eb953bcb5a
notion_url: https://app.notion.com/p/Lessons-Learned-while-Working-on-Large-Scale-Server-Software-6eb03bd1f616415989e537eb953bcb5a
last_edited: 2026-09-21T17:28:00.000Z
source_url: https://ferd.ca/lessons-learned-while-working-on-large-scale-server-software.html
tags: ["English", "Programming", "System Design / Software Architecture", "Project Management", "Article", "My bad opinions (Fred Hebert)"]
---
[https://ferd.ca/lessons-learned-while-working-on-large-scale-server-software.html](https://ferd.ca/lessons-learned-while-working-on-large-scale-server-software.html)

Of all the lessons I've learned, there is one that can summarize them all in 3 simple words: everything is terrible. This text is an attempt to recount some of the hard-earned lessons I have ended up learning, sometimes indirectly, but often personally. Everything is terrible, but our job still is to build something solid and usable on top of that everything. What we build adds to that 'everything', makes it bigger, more terrible.

Existing systems come with all sorts of terrors; each time I touch something, I cause subtle ripples that disrupt the careful balance of all the elements holding it together. Each bug fixed opens the door to a bigger one, hidden deeper, and more violent.

Writing new systems comes with an inescapable feeling of dread. I can imagine everything failing before it has been written, and feel sorry for the person who'll maintain it—possibly me.

That being said, here are fourteen lessons I've learned or live by in order to be more comfortable with the task at hand. I also recommend you read How Complex Systems Fail by Richard I. Cook. It's a much better list than mine.

### 1. Plan for the worst

Always have a plan for the worst case scenario for error conditions. What happens if all the databases are down? What if they're up but have lost all of their data?

Find a general solution for that. And by general solution, it can be as simple as automatically shutting down all operations and returning an error code. Maybe your clients should or will know how to cope with that, wait a bit, retry later. Maybe they'll have a contact number to call to talk to someone letting them know things will be alright, even if the sky is falling right now.

Once this is done, you know what's your bottom of the barrel. You know that if everything goes to hell, it won't be worse than that. This is a recomforting baseline. Any error condition you will not have seen coming, every operator error, every weird circumstances you'll have, it will at worst be as bad as this one.

Handle this worst case scenario and you've just figured out how to handle pretty much all errors in your system. It might be painful, but all other error recovery modes are just optimizations over this one. You do not understand how your system works if you do not understand how it fails.

### 2. The CAP Theorem is Real

Do not forget it. The CAP Theorem is a good model to keep in mind. It's not exaustively listing everything that can go bad, but it forces you into making important decisions. Pick your poison, and stick to it.

Document the option you've taken, and if possible, have every bit of your system that can respect it do so. Idempotence is your best friend. It can be as simple as stamping every message with a recognizable ID that ensures the system will refuse to reprocess the same logical message twice, or it can be fancier (a 'diff' carries enough context to not be reapplied multiple times, for example). But having that property saves a lot of headaches.

### 3. The Fallacies of Distributed Computing are real

Do not forget it either. The eight fallacies of distributed computing will be what you live and die by. The network sucks, latency is unpredictable, bandwidth has a cost and limits, people may invade your network, things will move around, many teams or corporations will be in charge of various parts of your system (and each will have their own agenda), you'll have all kinds of devices, protocols, serialization formats, and some of them will be worse than others—likely the most used ones.

If you can avoid making parts of your system distributed, avoid it. It's leaving the comfort of your home to step into a rain of fire.

If you have to make parts of your system distributed, please, think about it. Hard. Know what that means, how it can fail, die, look healthy (but still be failing), and how it can recover. There's nothing more dangerous than someone going for a stroll over the network without knowing what it entails.

### 4. Back-Pressure or Load-shedding: pick one.

I've written and spoken a lot about this one.

Sooner or later, your system will be overloaded. You have two options: do you stop people from inputting stuff in the system, or do you shed load. Those are inescapable choices, where inaction leads to system failure. System failure means you both stop taking input and lose what was going through the system.

Pick one and prevent (some) losses. Pick one and guide your entire strategy towards optimization and staying away from your accident boundary.

### 5. Debugging is a Science

In Programming Forth (Stephen Pelc, 2011), the author says "Debugging isn't an art, it's a science!" and provides the following (reproduced) diagram:

By far, the easiest bit is 'find a problem'. The difficult bit is to form the right hypothesis that will let you design a proper experiment and prove your fix works. It's especially true of Heisenbugs that ruin your life by disappearing as you observe them.

Embrace systems and platforms that you can dig into. The hard bugs are those you will not have seen coming, and therefore those that have nearly no instrumentation or metrics around them; otherwise, they'd be caught early and not be left to fester until other things mix in with them to create a deadly chimera of failure.

It's no secret I'm a fan of Erlang where I can trace everything and inspect memory live without stopping the system. I swear by these tools more than I'd like to need to, and I frankly can't imagine the pain of investigating without them anymore.

These cut down the debugging loop so much I wish people who never had such tools to assist them could try them on their systems. If you're dealing with the network and know tcpdump and wireshark, imagine the same thing, but for your entire programs. It's the best analogy I can come up with.

### 6. Postel's Law is Hard.

Postel's Law says "Be conservative in what you send, be liberal in what you accept". There is a good interpretation of it that makes a lot of sense, but the colloquial usage of it tends to mean "send good data, try to accept garbage".

That latter form is hogwash. Always start your implementations as strict as possible. It's much harder to respect the lenient version of Postel's Law without accidentally corrupting data than it is to just implement the specification strictly and end up never corrupting anything.

### 7. Don't trust the network.

This is a bit redundant now with mentions of the CAP theorem and the fallacies of distributed computing. But really, just don't trust the network. The network owes you nothing, and it doesn't care about your feelings. It doesn't deserve your trust.

Operating systems will react differently to connections that are made over your localhost (or loopback interface) than those from remote hosts. Your tests may mean nothing. Some kernels will have tricky rare TCP bugs. Some behaviours will only show up when discussing with specific hosts doing specific things with specific configurations and never show up again (or show up all the time). Asymmetric netsplits are real (and they even kill Raft very bad).

The network is a necessary evil, not a place to expand to for fun. Respect End-to-End principles if you want your life to be easier.

### 8. Et tu, System?

The pieces of your infrastructure you trust the most will eventually be your most painful ones. Someone will inevitably trust them so much they'll become 'magic' to your (or another) team. As everyone learns to trust it without ever touching it, it starts to rot and suffer under pressure. You end up having to do difficult (sometimes impossible) scaling operations on components that are now both legacy and critical. Worse, changes will be time-boxed as product development becomes tied to its scaling and performance.

A huge part of a system's success is tied to its operators. Without enough practice, they might just become the riskiest part of the system. And here comes the sense of dread.

### 9. Crash Early, Crash Often

When you're not sure about how to handle an error, let it crash. Violent failures on 1/Nth of your system early on is far better than silent corruption in 100% of it over a long period of time. Errors that declare themselves loudly and early are easy to spot, and easy to thwart. Errors that you notice have been slowly destroying your system from the inside over the course of days, weeks, or months are what is truly painful.

### 10. Deploys Fix Bugs, Cause Failures

Installing and deploying new software is a great way to introduce all kinds of new variables in your system and bring it down. The bigger the deploy, the scarier.

In my experience, a major code deployment that goes terrible causes shame. A major deployment that has minor issues is par for the course. A major deployment that goes well is terrifying because we all make mistakes, and if they weren't violent, they might just be the silent kind that takes a long time to detect.

### 11. Long Running Systems Have Their Own Bugs

Continuous deployment (with rolling upgrades) is standard practice these days, to varying degrees. Restarting nodes frequently means that all the funny behaviours, corruptions, and low-probability events that take a long time to develop and show themselves will remain hidden.

It also means that they'll only show up when everyone's on vacation, during the holidays, or whenever everyone is knee-deep in new project development and new deploys aren't occuring for an extended period of time. This means not deploying changes is also a great way to make your systems go bad.

### 12. Be Ready for a Total Restart

Be ready to restart the entire system from a blank slate, under load. If you can't do it, the day someone or something takes the whole thing down accidentally (or intentionally!), you'll be unable to bring it back up.

This can be the difference between less than one hour of downtime with a quick turnaround, and spending 15 hours trying to bring things back up unsuccessfully.

### 13. There's more Global stuff than variables

Global variables are one kind of thing that can break your programs by spookily changing things at a distance. More insidious are entrenched implicit design decisions. Sometimes, an accidental (or at least non-explicit) technical decision has been made and relied on by other components.

Think of a case where an integer is used as an ID for some events or entries in a system. You change it to a UUID (because that reduces bottlenecks or conflicts and you can split up your system), and suddenly 3 unrelated subsystems break because somehow they used a 'sort' function on the IDs to determine monotonic increments and provide time order to our entries. Woops. All of a sudden, a seemingly small change has to be reverted and cannot be enacted until a parallel mechanism is developed to provide the same feature that was otherwise implicit.

Refactoring of large systems, particularly legacy systems, is fraught with peril because there have been many such decisions, made by both smart and distracted people, that end up acting as invisible glue to large parts of everything you stand on. Refactoring can only truly begin once you've actually learned what a piece of code or some data structure did, the unique properties for which they were written or chosen. Anything else is setting yourself up for failure.

### 14. It's all about people

Humans are the lynchpin holding things together. Systems will live and die by them. Hard-earned lessons about making systems stay alive, operating them, and learning how they should behave to quickly spot what's abnormal take time to develop, and is usually held by humans.

This kind of knowledge usually remains embedded within the teams that develop it, and tends to die when individuals leave or change roles. When new members join the team, it gets transmitted informally, over incident simulations, code reviews, and other similar practices, but never in a really persistent manner. Being aware of that and building channels for more persistent information is crucial.

It also means that when building systems, you should not assume that operators will do things correctly. Expect failure from people. Try to think about tools you can give them to undo their mistakes, because they will happen sooner or later. Have some dread. Be understanding. Know things won't be perfect.

In the end that's how all systems live and die: in a perpetual state of partial failure. Major failures is just what happens when many of them occur at a given time. A system that is 100% healthy is a system that probably needs better monitoring, because that's not normal; it's worrisome.
