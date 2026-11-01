---
title: "How To Perform a Successful Incident Postmortem"
notion_id: 3a8b9db3-81dd-4f1a-834c-93f0687e67fd
notion_url: https://app.notion.com/p/How-To-Perform-a-Successful-Incident-Postmortem-3a8b9db381dd4f1a834c93f0687e67fd
last_edited: 2022-12-21T15:22:00.000Z
source_url: https://hackernoon.com/how-to-perform-a-successful-incident-postmortem-that-will-help-your-team-learn-and-grow-ft493355
tags: ["English", "Programming", "Site Reliability Engineering", "Help Desk", "On Call", "Project Management", "Article", "Hackernoon"]
---
## Too Long; Didn't Read

An incident postmortem is a meeting that brings together all of the people that were involved, whether directly or indirectly, to discuss, document and derive value. Postmortems should remain positive to avoid demotivating technical teams further. Learning should be the focus rather than dwelling on what went wrong, and so any past tense discussions should stick to facts instead of opinions. A variety of different stakeholders should get invites to the postmortem, including: Teams that first logged the incident and how long it had been ongoing before being logged.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**Kyle Jones**

When growing a business from a startup to a large enterprise, its software systems also expand in complexity, meaning that encountering incidents is inevitable. These incidents have indirect costs that can include loss of trust in the product or brand. However, incidents occurring are not necessarily a bad thing — it provides the business with new learning opportunities and the chance to improve its operational practices. But how do we learn from such a failure?

## What is a postmortem?

An incident postmortem is a meeting that brings together all of the people that were involved, whether directly or indirectly, to discuss, document and derive value. Postmortems should remain positive to avoid demotivating the technical teams further. Postmortems should also be blameless - this empowers engineers to provide details of their contribution and prevents the establishment of fear culture. Learning should be the focus rather than dwelling on what went wrong, and so any past tense discussions should stick to facts instead of opinions — avoiding phrases that include would have, could have or should have is essential.

## Who should get invites?

A variety of different stakeholders should get invites to the postmortem, including:

- Individuals/Teams that First Logged the Incident
- Individuals/Teams that Responded to the Incident
- Individuals/Teams that Diagnosed the Issue(s)
- Individuals/Teams that Rectified the Issue(s)

It may also be appropriate to invite a user that was directly affected by the incident. Invite a range of different people like this maintains transparency and allows us to glean as much information as possible to document.

## What should get documented?

Thorough documentation of the event is vital. Memories are short, and in time important details fade into obscurity. Some key facts to document are:

- Date and Time that the Incident Started
- Date and Time that the Incident was First Logged
- Which Teams/Individuals Responded to the Incident
- Number of Users/Accounts Affected
- Number of Support Requests Raised
- Date and Time that the Incident was Fixed
- Any Solutions or Mitigations

In addition to the above details getting recorded in a postmortem document, the meeting should also have minutes taken, and a timeline of the incident constructed.

## What should you do after the postmortem?

The postmortem produces discussion points to bring into further sessions and documentation that will prove valuable in the future.

Values recorded should be compared against any service level agreements ([SLAs](https://www.atlassian.com/itsm/service-request-management/slas?ref=hackernoon.com)) that may be in place, to confirm that the incident did not result in any breaches. Any issues identified as a result of the incident should be discussed in-depth, with potential solutions or mitigations planned into the roadmap, alongside rigid delivery dates. These solutions/mitigations should have tickets written to capture the work, each of which should be [SMART](https://www.atlassian.com/blog/productivity/how-to-write-smart-goals?ref=hackernoon.com). Depending on the incident, particularly regarding who first logged the incident and how long it had been ongoing before being logged, improvements to the observability may be required. Observability improvements should be a priority alongside immediate solutions to the faults.

If an external user reported the issue, it might be pertinent to publish the postmortem's findings openly, allowing anyone access. Publishing postmortem outcomes publicly has most [notably been utilized by Monzo](https://monzo.com/blog/2019/09/08/why-monzo-wasnt-working-on-july-29th?ref=hackernoon.com). It enables them to maintain transparency, ensures accountability as a business and has provided their users with greater trust in the brand.

_Previously published at _[_https://kylejones.io/how-to-perform-a-successful-incident-postmortem_](https://kylejones.io/how-to-perform-a-successful-incident-postmortem?ref=hackernoon.com)

L O A D I N G
. . . comments &
