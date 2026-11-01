---
title: "How to create sustainable on-call rotations"
notion_id: 31d34120-f37e-43bc-a36d-ac63777ad315
notion_url: https://app.notion.com/p/How-to-create-sustainable-on-call-rotations-31d34120f37e43bca36dac63777ad315
last_edited: 2026-09-21T17:03:00.000Z
source_url: https://leaddev.com/technical-direction/how-create-sustainable-call-rotations
tags: ["LeadDev", "On Call", "Help Desk", "Health", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

On-call is a powerful tool for engineering teams. Here are some strategies for building more effective, sustainable on-call rotations.

If you manage teams that run software in production, whether you have an end-user-facing web app or an internal service aimed at developers, at some point you’ll probably have to create an on-call rotation for that software.

The way you manage this rotation will have a big impact on your user experience and company reputation (it’s much less embarrassing to discover and fix issues yourself than to hear complaints about your uptime on Twitter). It will also, of course, have a significant impact on the work and lives of the people on call.

So, how can you create a sustainable on-call process that alerts you to user-facing issues that need human intervention, allows you to fix them quickly, and limits the impact on engineers’ real lives outside of work?

If this is your first time setting up an on-call rotation, here are some of the primary considerations around tooling, scheduling, and work management to keep in mind.

### Understanding and achieving your on-call goals

On-call is about alerting people to the presence of issues that require immediate human intervention to fix. There are a couple different important implications of this statement:

First of all, on-call is about _escalating issues that require human intervention_. If an issue is always fixed through the same steps (for example, ‘every time this alert fires, we restart the webserver process’), those steps should be automated. Don’t page someone for something a computer can do itself.

Second, alerts should be things that require a _timely response_. Nearly every ops person I know has a story about getting paged at 3 a.m. because some server’s disk went from 80% full to 81% full and that threshold triggered an alert, despite the fact that the disk wasn’t likely to reach 100% full for another 2 weeks. If it can wait until business hours, it should.

The tools you use as part of your on-call setup should support these goals. This means having separate tools for monitoring (gathering information about your systems), alerting (notifying the appropriate human responder about an issue based on that information), and troubleshooting (helping the human responder gather more information about the issue).

Which tools to use, and which checks to alert on could fill their own articles, so here we’ll just say to start small. If this is the first on-call rotation you’ve set up, start with the basics (for example, ‘is the main URL for our webapp responding with a 200?’) and iterate from there. Start too big, and you run the risk of overwhelming your engineers with meaningless, unactionable alerts, and that way lies alert fatigue, and sadness.

### Reducing the human impact of on-call

The biggest considerations around on-call have to do with the impact on the people involved. On-call is tiring. Even if you don’t get paged, you’re expected to be available. This might mean postponing social plans or medical appointments because you don’t want them to be interrupted, or feeling stuck at home because you don’t want to have to lug a laptop and MiFi device around with you. It removes people’s ability to truly disconnect from work and relax, and that has a significant impact on their personal lives.

That’s why it’s so important to compensate people accordingly. If on-call wasn’t part of their job description when they started, taking on those extra responsibilities should merit a salary adjustment. The company should pay for any associated costs, such as a MiFi or company cell phone. When people get paged outside of business hours, they should be given time off in lieu.

On-call responsibilities should be distributed both fairly and sustainably. The rotation should be set up so that people get sufficient time in between shifts to disconnect and recharge. If you’re setting up a weekly rotation – where one person is on call 24/7 for a week, a different person each week – you need 4 people in that rotation at minimum. If you don’t have that many people available, you need to figure out something else while you scale up the team, because forcing the few people you have into an unsustainable schedule is a great way to burn people out (which might leave you with even fewer people available). Something else might mean having on-call during business hours or weeknights only to start with; the rotation doesn’t need to be perfect to be better than what you had before.

Let the people who are in the rotation figure out a setup that works for them. There won’t necessarily be a schedule that everyone loves, but as the manager, you shouldn’t be dictating the details of a rotation you don’t participate in. You can communicate the requirements to the team – things like ‘we need to have somebody responding within X minutes for critical services Y and Z’ – but let the team work out the implementation details. They might prefer doing daily shifts instead of weekly, or having one person who always handles evenings because they prefer a more nocturnal schedule, allowing someone else with family responsibilities to have their evenings free.

Who should be in the on-call rotation? If you’ve never done on-call before, and you have a team of ops/infra engineers or system administrators, it can be tempting to have those ops people be on call for everything. But remember that one of the goals of on-call is to route information about incidents to the people who are best suited to fix them. If the issue is, ‘the frontend is throwing a bunch of javascript errors’, that’s probably not something the ops team knows how to fix. And if you’re thinking, ‘well, the ops team can just escalate to the frontend team if they need to,’ remember the point about automating the simple responses. If the response to javascript errors is always for the ops team to escalate the incident to the frontend team, you don’t need to bother an ops person for that; just send the alert to where it belongs in the first place. If your devs don’t like being paged for issues, remind them that the ops team doesn’t want to be paged for issues that they didn’t even cause, and then make sure that you’re giving your devs the resources they need to find, resolve, and prevent issues from getting shipped into production in the first place.

### Handling the unplanned work that comes out of on-call

On-call is a dynamic, living thing. Due to the law of entropy, and the fact that computers are terrible, on-call has a tendency to get worse over time if you don’t take proactive steps to make it better. You must build enough slack into the system to allow people to properly deal with the work that will come out of on-call. Remember, bandaids are not real solutions. If something catches on fire at 3 a.m., the on-call engineer might throw a quick hack into place to tide things over until morning (nobody’s likely to be doing their best work in the middle of the night) but they need to have time to actually fix those hacks later. On-call will generate unplanned, interrupt-driven work by its very nature.

When a person is on call, their only responsibilities during the shift should be responding to alerts and doing the extra work generated. In an ideal scenario, where people aren’t getting 40 hours of pages a week, they should be doing things like going through past remediation items, fixing bugs in the backlog, or replacing the 3 a.m. band aid fixes with real, robust solutions. If they get through all the on-call-related tasks and have time to work on their regular planned work, great, but that should be viewed as an added bonus, not something that you depend on with your scheduling and product planning.

Engineers who are on call need to have the authority and autonomy to make changes to the on-call experience. Every alert that pages a human should be actionable. If alerts aren’t actionable, the engineers should be allowed to delete that alert. This might sound obvious, but sometimes managers or executives think that it’s better to err on the side of caution and that they’d rather have a false positive than a false negative, but inactionable alerts lead to alert fatigue leads to burnout. If you as a manager want to send every possible alert ever to your own personal phone, go ahead, but you’ll quickly realize the experience is as unhelpful as it is unpleasant.

Let the on-call engineers – presumably domain experts who were hired for their skills and knowledge, and people at the sharp end doing the day-to-day work – decide what needs to be done. If they say the rotation is getting noisier and noisier and request to stop the line, for example to take some time to fix a backlog of bugs before deploying new features or launching a new product, take that seriously.

### Reflections

An on-call rotation can be a boon to your organization, but it must be managed well in order to be effective. Remember that the purpose of on-call is to alert you to urgent customer-facing problems that require human intervention, not to send alerts for the purpose of sending alerts. Keep in mind that when on-call surfaces these sorts of problems, you need to build enough slack into your planning processes to allow people to fix them properly, because there’s little point alerting on issues if you aren’t going to address them. Finally, go into this with the understanding that on-call impacts people’s real lives outside of work. That additional workload needs to be compensated, both for the work itself, and for the impact that it can have on health, sleep, relationships, and more.



<!-- unsupported block: link_to_page -->
