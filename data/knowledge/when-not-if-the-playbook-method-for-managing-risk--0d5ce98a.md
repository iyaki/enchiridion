---
title: "When, not if: The playbook method for managing risk"
notion_id: 0d5ce98a-cc99-49dc-b49a-3d97fffe130b
notion_url: https://app.notion.com/p/When-not-if-The-playbook-method-for-managing-risk-0d5ce98acc9949dcb49a3d97fffe130b
last_edited: 2023-04-16T22:44:00.000Z
source_url: https://leaddev.com/process/when-not-if-playbook-method-managing-risk
tags: ["English", "DevOps", "Site Reliability Engineering", "Productivity", "On Call", "Help Desk", "Article", "LeadDev"]
---
**Traditional risk management isn’t working for engineering organizations. Here, CTO of Identity and Network Access at Microsoft, Yonatan Zunger, shares a better approach.**



Amateurs build systems that perform well; professionals build systems that fail well.

This difference goes to the heart of what it means to be a professional: that you will be responsible not just for rolling the system out the door, but for its continued success and operation down the line. Things will go wrong, and when they do, you will need to be able to deal with them.

This means that one of the core things that professional developers – by which I mean anyone who is responsible for making anything persistent, be it a machine, or a process, or even an idea, that other people will use – need to do when building systems is to think about the ways in which the system might fail, and prepare for those.

Traditional risk management was designed by and for insurers. It’s about estimating likelihoods and impacts of things that can go wrong, and managing a risk budget. This works very well if you’re managing a large risk portfolio and care about its effects on a purely numeric scale – the case it was designed for. But in my experience, it isn’t as useful for the people actually building the systems or making day-to-day decisions about them, because it can result in incentive misalignments and fail at capturing certain common risk patterns.

Instead, I want to discuss an alternative approach based on ideas from site reliability engineering: one which starts, not by asking if something will go wrong, but by assuming it will go wrong, and requiring a plan for what you’ll do when it does. Instead of a risk budget, we’re going to build a playbook for handling disasters, and use the process of building and maintaining that playbook as a tool to prevent them in the first place.

Below, I’ll explain how the traditional and playbook approaches work, discuss the tricky bits of the playbook method, along with real-world examples, and discuss the psychology and organizational management techniques you need to use this method in the field.

## The problem: Why traditional risk management doesn’t help

The standard approach to risk management, which you’ve probably encountered if you’ve ever tried to get something certified or insured, works like this:

- List all the possible things that could go wrong. Assign each of them a ‘likelihood’ and an ‘impact,’ often on a scale of one to five for each. Based on those, give them a risk score that is more or less the product of the two.
- Going in order from the highest risk score down, decide if you’re going to remediate the risk (change things so it goes away), mitigate the risk (figure out a way to reduce its impact if it happens), transfer the risk (make it somebody else’s problem, probably by buying insurance), or simply accept the risk. Then calculate the residual risk by figuring out the likelihood, impact, and risk score after whatever you’ve done.
- Manage your overall risks with things like a maximum residual risk you’ll accept for any single risk, and a total budget of risk score that you need to spread across all your risks.

This system was designed to work from the perspective of someone managing a large portfolio of diverse risks, where all the risks could be boiled down to one (or a few) currencies – that is, from the perspective of an insurer. But when it’s used in an operational team, I’ve seen four major problems emerge time and time again.

- There’s a perverse incentive to ignore risks: if you don’t discuss a particular risk, it doesn’t go on your risk budget, and you don’t have to invest in it. That gives people an incentive to not think about things that could go wrong, which is exactly the opposite of what you want risk management to do.
- People are bad at probability: if you start talking about the likelihood of something going wrong, people start to fall into a mode of saying, “oh, that’ll never happen, we won’t worry about it.” That makes people far too likely to accept rare risks – even if they’re not nearly as rare as people think. This also creates a similar perverse incentive to underestimate likelihoods.
- The math breaks for rare, high-impact events: the product of a very big number and a very small number is not a medium number, it is statistical noise. This is especially pronounced for things like security and privacy risks, where most risks are rare, catastrophic events. A five-point scale won’t cut it. You would need a logarithmic scale to capture most of these things. But normal risk models lose that, and so rare catastrophes get artificially low risk scores.
- The responses break for common events: the above dictionary of responses is based on things that might happen. Things that will happen need to be treated differently. The problem is that the boundary between these two is extremely fuzzy and changes with scale. My friend and colleague, Andreas Schou, put it nicely: “If something happens to one in a million people once a year, at Google we call that ‘six times a day.’”

An event that goes wrong regularly isn’t a ‘risk’ in the language of traditional risk management, but the border between a risk and a regular event is just a matter of frequency – and causing people to plan differently for one or the other goes poorly.

The combination of mathematical failures and perverse incentives is a witches’ brew. People with an incentive to ignore risks will leap on any mathematical error (either “this is so small, we can ignore it” or “that’s obviously too big, that proves this model is broken”) to ignore more things.

This is true even if these people are well-intentioned. Perverse incentives work on us subconsciously as well as consciously. But those perverse incentives are omnipresent. People are rewarded for building and launching successful systems, and often those rewards disproportionately favor the big news of the initial launch over the long-term health of the system.

If you are either a decision-maker or a builder and maintainer of systems, this method is not going to get you where you need to be. If you are instead a manager of the things that can go wrong, like a member of a ‘horizontal team,’ this method will just create conflict between you and product teams, and not give you the tools to resolve it.

## What is the playbook approach for managing risk?

Consider an alternative approach:

1. List all of the possible things that could go wrong. For each of these things, describe the consequences, and estimate the frequency (not likelihood!) with which it will happen. Explicitly look for combinations of things going wrong, namely, two failures happening at once.
2. For each of the things on that list, write a playbook. How will you detect that this has happened? What, exactly, should the person who detects this (and any other people they then pull in) do to mitigate it? When do you need to escalate, and to whom? If decisions will have to be made, what information will people need in order to make those decisions, how will you get that, and who will have to make the call?

> “Most really catastrophic failures are the result of survivable failures stacking on top of each other in unexpected ways.” – Robert Hansen, software engineer at Google.

And perhaps counterintuitively, that’s it. We’re not going to have an overall risk budget or attempt any numeric calculations at all. We will continually extend the playbooks as we learn more about the system and things actually go wrong; that’s what postmortems are for. Everyone involved should be actively encouraged to add new potential risks to the playbook.

How, and why, does this work? Let’s look at it from the perspective of the people involved.

For the people who run the system, this has an obvious use: it’s literally a core tool for handling emergencies. A pre-written playbook keeps you from having to figure things out from first principles while alarms are going off at three in the morning.

For the people building the system, who are often also the people running it, it helps figure out what you need to build: the playbook brings you face-to-face with what’s going to be happening in the field, and the tools you’re going to need in order to resolve the situation, get the data you need, and so on. That list turns into one of your main product inputs.

For decision-makers, this realigns incentives towards long-term success. There’s no longer an option to say, “oh, this will never happen” and effectively drop something from the list; you just need to say, “if this happens, here’s what we’ll do.” In the short term, this is scary. In the long term, it greatly improves your ability to sleep at night.

## The tricky bit: Addressing prioritization

However, this method runs into two tricky bits. The first is about how to address prioritization (which is much more obvious in the traditional approach), and the second is about how it interacts with legal risk. Let’s start by examining the first of these, and do it by looking at a real-world example.

The **prioritization principle** is that you need a plan for every item on the list, but as frequency drops, your plan can become more manual and laborious to execute, and once you reach a certain frequency, you apply the cutoff point of, “this is no longer your biggest problem.”

To give a practical illustration of how this works, let’s consider BS2, a storage system built by Google which is available commercially as Google Cloud Storage. This service also stores Gmail attachments, photos, videos, and much more. (I was its technical lead, which is how I know this story and can share the thinking behind the decisions).

BS2 is a natively “planet-scale” storage system, which is to say that it treats the existence of individual data centers as an implementation detail generally hidden from the user, much as data center-scale storage systems hide the existence of individual computers and racks.

In BS2, customers specify a “minimum replication rule” like “2+1:” for each file, keep at least two copies on disk and a third on tape. Disk copies affect both availability and data integrity, while tape copies provide additional integrity both at much lower cost and the additional reliability of a separate system which is subject to different risks.

_Side note: Why is that a benefit rather than a cost? Because some single disasters, like software bugs, could affect all the disk replicas simultaneously, but a heterogenous backup would be shielded._

When designing this system, we had to consider events that could take out components ranging from a single hard disk, up to the entire system. Because the system uses existing data center-scale storage systems (like Colossus and BigTable) as its primary storage components, those systems’ plans for handling failures of units smaller than a data center (hard disks, computers, racks, rows) already manage those failure modes, and those plans are exposed via the SLA’s provided by those teams.

However, those systems deliberately consider data center-scale failures out of scope. Their plan for if an entire data center loses power, for example, is, “this system loses power too, and stops working; we have a plan for how to bring the system back up if and when power is restored.” That is a very appropriate termination of the risk list for a data center-scale storage system, but it means that the planet-scale system needs its own plans for data center-scale failures.

Data center failures can be broken into a few categories based on how bad the failure is. There is scheduled maintenance, which happens every few months for each data center. There is an emergency power-off (EPO), where you have an hour or so of warning, which often happens if something like a regional brownout is happening. There is a zero-notice EPO, where you have five minutes or less of warning, which happens if, say, the building is on fire, or the cooling system fails. Lastly, there is an irreversible EPO, where the data center switches off and isn’t coming back, perhaps because the fire actually destroyed enough things that whatever comes next is effectively a new data center.

When building our playbook, outages with an hour’s notice or more were easy to deal with. We stop reading from or writing to that data center, routing requests elsewhere, and instead have the data center focus on making sure that any files it has which have recently been created, and aren’t fully replicated yet, get replicated. That way, when the data center switches off completely, things only become unavailable if they’ve intentionally picked a replication policy of keeping only one live data center copy – something which essentially nobody does, for exactly this reason. This process could be fully automated, requiring zero human action in the ordinary case.

_Side note: Response plans live on an automation spectrum. At one end are totally manual responses, where humans need to be alerted, using processes and human judgement to know what to do and tools to identify and correct the issue. As a problem becomes more frequent, it makes sense to invest more in tools that can identify and remediate common cases more automatically, escalating to humans only in exceptional cases or when judgment is needed. The endpoint of this spectrum is a fully automated solution where, except in very unusual cases, humans don’t need to be involved at all. Beyond this is rearchitecture, where the system is changed so that this failure mode is no longer possible in the first place. As you move down this spectrum, you are paying more of an up-front investment to get lower maintenance costs down the road. The hardest case is where common failures require extensive human judgment to resolve. These generally are not amenable to scaling, but nonetheless require resolution. Privacy and abuse issues commonly fall into that category._

Zero-notice outages, reversible or otherwise, are trickier. If you’re relying on a background process to ensure that all files have the right number of replicas, new files could easily become unavailable or even lost outright (if the outage is irreversible). We dealt with this by changing the way new files are created to very urgently create a minimum number of replicas of each. This meant that zero-notice outages and outages with notice end up working the same way, and require no human intervention.

But what about outages that span more than a data center? For example, often many data centers will be located at a single facility, sharing services like power and network, and multiple facilities might be in a single metro area, which is subject to overall risks like natural disasters. Having two copies of the file won’t help you if both are destroyed by the same tornado!

To handle this, we came up with the concept of “meteor impacts”. These are events that cause the zero-notice, irreversible deactivation of an entire metro area at once. To manage meteor impact risks, we changed the algorithm that distributes replicas to be aware of geography, so that those minimum replicas that we were counting on for availability and integrity wouldn’t be at risk of the same natural disaster.

Now we consider stacked disasters. What happens if there are two meteor impacts? If they are spaced widely enough in time the existing plan works fine: the automatic restoration of minimal replication policies by background processes of the system would get the system back up to full robustness without human intervention. However, two meteor impacts in a short time window would not be helped. To resolve this, we recommended that people use the “2+1” replication strategy, and made it the default: losing any two metros would, at worst, impact availability for some files for a day or so, but not lose integrity.

This, in fact, is our answer for any number of multiple meteor impacts. As it happens, the 2+1 strategy is enough for GCS to offer an 11 9’s annual data integrity SLA, which (perhaps unsurprisingly) turned out to be more than enough for nearly all customers – and those who needed more could easily select 3+1, or even higher.

But we knew we didn’t need to worry about triple meteor impacts even before we did this calculation, for a simple reason: if three distant metro areas have been simultaneously destroyed by natural disasters, then whatever is going on here, the integrity of your files in the cloud is almost certainly no longer your biggest problem in life.

This is the real termination rule. You stop considering risks when, if that event happens, you no longer care about what happens to the system. The point at which this happens depends on what you’re doing. It’s different for a dating app than for an implanted pacemaker – but there will always be some set of risks you simply can’t do anything about and shouldn’t bother planning for.

## The other tricky bit: Legal risks

Although much of the point of this approach is to eliminate perverse incentives, it’s worth recognizing that whenever a system – even a perverse one – is used for a long time, there’s generally a reason why it’s being used, and other things may be depending on its behaviors. In particular, one cost of the playbook approach is that you lose the flexibility to quietly ignore risks that you don’t want to talk about. While this is often something you want to lose, it shouldn’t be done blindly.

In particular, this can create new kinds of legal risk. Under many laws around the world, “we couldn’t have predicted this” is an important legal defense. If you did know something was a risk and failed to act, you can be significantly more liable. This means that having something in your playbook, or even there being a record of you having discussed whether it should be in your playbook, can create a huge legal liability if after you handle it someone is still hurt or suffers a loss. (Note that I am not a lawyer, and if you’re delving into this, you should get to know your counsel well and discuss this with them!)

In my experience, this is actually the biggest obstacle to adopting a playbook approach, but it can also be one of the approach’s great strengths. There are a few ways to manage this risk.

One approach – the one I strongly favor – is using blunt honesty. In this one, you put the thing in the playbook anyway, and work with lawyers and other experts to make sure that your playbook also contains a clear argument for why you stop at a certain point and why this is justifiable. This actually makes for better playbooks. You define a stopping point for attempted remediations very intentionally, with a clear understanding (that you can even explain to an affected customer) for why you can’t do more.

Side note: While I was at Google, Larry Page taught me something very useful about defining stopping points. Whenever you came to him with a plan to grow a system by 10x, he would ask “why not 100x?” If you came with 100x, he would ask “why not 1,000x?” The key to a successful meeting was to have thought this through ahead of time and have a very clear argument for why you were going to stop at some size, such as the need for a much more fundamental design shift to add that next factor of 10, or because you needed the learnings of 100x to figure out how to do 1,000x. The same principle applies here.

From a legal perspective, if this ever comes to court, you not only have an argument pre-baked, but it’s evidently baked into your procedures, which makes that argument considerably stronger. (Of course, this approach also requires that you follow your own procedures, but if you’re writing procedures down and not following them, you are already setting yourself up for gigantic legal and practical nightmares and, once again, this playbook will not be the worst of your problems.)

This doesn’t work if there are things that you don’t want to invest in which you couldn’t justify in the press or in court. This is where searching for plausible deniability, and playbooks – especially ones where people are encouraged to contribute to them – become a foolish strategy. If this is your goal, then I no longer have advice for you; whatever you are doing is not honest work.

## Getting people on board

Organizations are rarely openly unethical. It’s much more common to teeter on the verge of dubious conduct, with people incentivized to take unwise shortcuts that they don’t actually want to take, or admit to themselves that they are taking. In this scenario, plausible deniability becomes an internal way of functioning. It’s even more common to have a cultural environment where risks simply aren’t discussed, because it violates a norm of optimism or of not raising conflictual issues.

In these cases, the playbook approach can be particularly powerful in avoiding or repairing such problems. A single team can easily adopt it independently of wider company practices, as an additional part of their own launch process. A playbook is a hard thing to argue with as it is so directly tied to operational work. The efficacy of having a playbook in improving incident response makes it an easily copied and evangelized idea within the larger organization, making it easy to encourage adoption without having to first get top-level buy-in.

Adding hard questions into playbooks, by simply including them alongside other (preferably similar but less-controversial) problems, is a low-stress way to raise difficult issues. Framing a problem as, “how do we deal with this kind of event?” turns a major strategic question into a very practical, tactical conversation about day-to-day work. It also positively frames an issue as something that will be dealt with, and limits showstopper responses, such as blocking launches or un-launching products. It opens up a coherent discussion about when such a thing should be done, rather than if it should ever be done.

## Reflections

The playbook method is both a practical and a psychological tool. Practically, it lets you plan ahead for how to handle problems, and make complex planning decisions ahead of time. Psychologically, it undoes perverse incentives and makes hard conversations easier.

It doesn’t replace traditional risk management – that is still very appropriate in cases that fall into the insurer model – but in my experience, it is an excellent tool to use in the cases where that method falls short.
