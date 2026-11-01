---
title: "How to streamline your processes as an engineering manager"
notion_id: 132c958e-1e83-4c7a-9d57-4c502ea0d9db
notion_url: https://app.notion.com/p/How-to-streamline-your-processes-as-an-engineering-manager-132c958e1e834c7a9d574c502ea0d9db
last_edited: 2026-09-21T17:30:00.000Z
source_url: https://leaddev.com/leadership/how-streamline-your-processes-engineering-manager
tags: ["Article", "LeadDev", "English", "Line/People/Team Management", "Productivity"]
---
[https://leaddev.com/managing-managers/how-streamline-your-processes-engineering-manager](https://leaddev.com/managing-managers/how-streamline-your-processes-engineering-manager)



<!-- unsupported block: link_to_page -->

You have 1 article left to read this month before you need to register a free LeadDev.com account.

As a manager, many of your tasks are quite repetitive. With a few simple steps around these processes, you can start saving precious time.

As programmers, we’re accustomed to noticing patterns and abstracting them away for reuse throughout our applications. Thankfully, the same approach can also be used when you transition into management.

I manage teams of teams, so being very careful about how I use my time and setting up systems for myself are the difference between being organized at the beginning of every new day and being in a reactive state. Since I want to use as much of my daily mental capacity where it matters, I set up some reproducible patterns for myself and my teams, so I can spend my time where it counts.

Here I’m sharing a breakdown of the manager ops practices that help me to streamline work for myself and my team. Many of these approaches can work for you too, and can be adjusted to your own organization and needs.

### Templates

1:1 templates

It’s completely fine to create one-off docs for meetings and 1:1s, but if you begin scaling your team, you may notice that there are some repetitive tasks in these endeavors that can be abstracted so that you’re not starting from scratch every time. Here are a few examples:

I typically have three templates running at any given time:

- 1:1 templates for direct reports
- 1:1 templates for skip levels
- 1:1 templates for peers (this one can be used for folks above, below, and around in other depts as well)

For direct reports, I typically kick off the meetings with topics like their career, their values, whether they have context on any projects they want to share with me, and if there is anything I could be doing to support them. I also give them space to talk about what isn’t going well, and if they want to brag about anything. I have noticed the last two topics are a very American approach – if you have an international team, you may want to adjust your framing here. The Culture Map by Erin Meyer has some great guidance on this.

Over time I also take a 30/60/90 approach to talking about work and career ladders with my reports (i.e. breaking down the work they work to be doing over the next 30, 60, and 90 days – more about this process here). I then split up the documentation into small in-doc templates (Notion has some interesting functionality for setting these mini-templates up).

For skip levels, I ask many of the things in the direct report template, as well as if they are experiencing any bottlenecks and what motivates them. I also ask questions about any specific team matters that I don’t ordinarily have full visibility into. I start all meetings asking how they would like to spend the time, so that we can cover items in the agenda or completely deviate as they see fit. Having an agenda set doesn’t mean you have to stick to it, it can be useful to kick the conversation off either way.

For peers, I typically ask if there is any context on shared projects I should know about, if there’s anything I can do to support them, and if there’s anything I can do to improve cross-department communication.

Meeting templates

Google Calendar now has a nice feature that allows you to add meeting notes fairly easily. This can come in handy in a pinch. but you can take this a step further and set up a template that is specific to the type of meeting and topic, because agendas and takeaways vary.

Types of meetings:

- Weekly team meetings
- Topicals
- Brainstorms

The format of the templates change depending on the broader purpose. For instance, for topicals, we make sure to capture the purpose, action items to take, and docs that are dependencies. Brainstorming sessions, on the other hand, are often much more open-ended and may not necessarily pressure people to have such clear outcomes.

Automating weekly standups

At the start of every week, I used to send all of my direct reports a form. In that form, I would ask what top three to five things they accomplished last week and what three to five things they want to accomplish this week. The form also included some status fields to update, and a field for anything on their mind. I’d automated the form so everyone on the team received each other’s answers as they wrote them. The answers were also automatically imported into a Google spreadsheet, kicking off a serverless function to build a dashboard out of the status fields.

Then, in our weekly standups, we could review the dashboard and talk about the most vital items and any anomalies.

This process kept us organized, provided visibility into our tasks, allowed people to reflect, and kept me from having to manually solicit items from the team or have to build dashboards etc. from scratch on my own. As the team grew, this came in very handy as my individual focus time dwindled.

Full disclosure: some members of my team liked this approach because it provides consistency and is a forcing function to planning out our tasks for the week and providing information between team members. Some didn’t like the approach because they were against the additional layer of process. Your mileage may vary depending on the culture of the teammates involved. That’s why it’s important to ask for feedback and adjust over time, as well.

Weekly to-do list template

So far we’ve talked about how to set up processes for your team, but not about how to set up processes for yourself. Personally, I like to use Notion’s Weekly Agenda Template to create a daily sheet for myself.

Rather than setting it up every week, I use a template that has some consistent items on it:

- Weekly 1:1s
- Weekly team syncs and meetings
- Any individual work I have to do that repeats
- Things I want to accomplish daily (for me thats: drinking eight cups of water, going on a walk, reading, writing in a gratitude journal)
- Household tasks like driving the kids, etc

From there, once I put the template to use for the week, I add any other meetings or tasks to that individual week. I set this all up with scripts for myself so that I remove as much of the manual work as possible.

There are also some great tools that help do this for you like Sunsama, which has a great approach to assisting you with tasks if you’re not a nerd like me who loves to set up her own python scripts.

The main takeaway is to notice what you do again and again or what patterns you have week over week and try to automate as much as possible.

### Knowledge sharing

The way information flows in an organization is directly correlated to the health and productivity of the team, as well as how well important stakeholders understand your work. In order to facilitate high transparency without constantly sharing information in one-off interactions, or making the people on your team individually responsible, we want to create as many low-friction systems for our team as we can. Ideally the time they take to set up is recouped over time, and your team knows how to find the information they need easily.

Handbooks

If your team or org doesn’t already have a handbook, it can be a good practice to create one. A handbook is a central place to find all the docs necessary to participate in your organization.

Some examples of what we keep in a handbook (even if it’s just links to something else):

- Design doc locations
- Incident management processes
- Onboarding processes and documentation
- Cross-functional and partner team contacts
- Asynchronous hygiene
- Any HR type policies that are team-related, like education budget
- An anonymous feedback form

Team calendars

It can be nice to have a team-wide calendar of events and important dates so that people can plan. An example of such a calendar can include:

- Annual planning cycles
- Company-wide performance review timings, including promotion cycles
- Team tech talks
- Office hours
- Important holidays
- No-meetings weeks

The only major issues I’ve had with creating such a calendar is that sometimes your engineering org calendar is dependent on other groups in the company being transparent and ahead of schedule on timing for the year, and sometimes that doesn’t happen. This can make it tough to stub out a calendar consistently. My suggestion is to stub out the dates you think are correct and be clear with a prefix that it’s not settled. Example: [EVENTUAL DATE TO BE DETERMINED]: Annual Planning.

### Soliciting wider feedback

It’s important to create an opportunity for folks to give you feedback, not just in a 1:1 setting, but in a wider, more anonymous format. It can be helpful to run surveys now and then to gather such feedback, and a lot of those processes can be automated.

Developer pulse surveys

I run developer pulse surveys with my group every quarter or so, with a consistent base set of questions (for example, ‘Does our team see failures/outages as opportunities to improve the system?’). This helps our team look at impact over time, and see if things are improving or declining, and decide what actions to take. I built a lot of the questions from information learned in the great book Accelerate by Nicole Forsgren, Jez Humble, and Gene Kim.

Once the form is set up, it’s fairly easy to automate both cadence and collecting the responses.

Post-event surveys

After team events such as offsites and hack weeks, it’s important to solicit feedback about how the events went and what could be better. Just like with the developer pulse surveys, a lot of the questions can be standardized and a base template can be used for the events with some deviations.

### Trackers

I could write an entire article on team trackers alone. Certainly, if you have good program management staff, this should be part of their work and how you collaborate effectively with them as cross-functional partners. If you’re at a company that doesn’t have program management, it can be a big job to break your team’s work up. I wrote a post a while ago on how to do exactly that.

### Wrapping up

There are a ton of small tasks involved in our work as managers, and some can be quite repetitive. We have to interrupt our own focus to help create wider systems for the teams, but if we set up these processes with care, we can make sure we’re optimizing our own time, so we can spend time where it really matters: with our people.
