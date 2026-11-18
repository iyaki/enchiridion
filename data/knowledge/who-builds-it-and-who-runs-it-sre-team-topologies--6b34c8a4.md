---
title: "Who builds it and who runs it? SRE team topologies"
notion_id: 6b34c8a4-732b-4f55-a887-b5fc618397e2
notion_url: https://app.notion.com/p/Who-builds-it-and-who-runs-it-SRE-team-topologies-6b34c8a4732b4f55a887b5fc618397e2
last_edited: 2023-03-29T23:35:00.000Z
source_url: https://stackoverflow.blog/2023/03/20/who-builds-it-and-who-runs-it-sre-team-topologies/
tags: ["Stack Overflow Blog", "English", "DevOps", "Site Reliability Engineering", "Line/People/Team Management", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Ad-hoc SRE principles can get you on the right track, but if you want to sustain it long term, you'll need organizational structure.

Site reliability engineering (SRE) can emerge as a bottom-up initiative to run services in an organization and grow into a successful practice fulfilling SRE principles. While ad-hoc SRE can help developers maintain code in production, to sustain the practice long-term, an appropriate organizational structure for SRE is needed. In this article, we explore SRE team topologies—ways to organize for SRE that stood the test of time.

# SRE principles vs. SRE organizational structure

To begin with, we need to distinguish between fulfilling the [SRE principles](https://sre.google/workbook/how-sre-relates/#background-on-sre) and an organizational structure for SRE. The SRE principles are:

- Operations is a software problem
- Work to minimize toil
- Automate this year’s job away
- Move fast by reducing the cost of failure
- Share ownership with developers
- Use the same tooling, regardless of function or job title

It is vitally important to understand that the SRE principles do not dictate any organizational structure. Rather, the SRE principles can be followed by teams embedded in several different organizational structures.

An SRE practice where the SRE principles are followed can succeed either with a central SRE team, without a central SRE team, or with several central SRE teams comprising an SRE organization. With this, what are the options to organize well for SRE?

# Who builds it? Who runs it?

Organizing for SRE must start with a fundamental decision: “Who builds and who runs the services?” This gives rise to [several options](https://www.stevesmith.tech/blog/who-runs-it) ranging from the traditional “you build it, ops run it” to the modern “you build it, you run it.” The main options in-between are “you build it, you and SRE run it” and “you build it, SRE run it.” In [_Establishing SRE Foundations_](https://www.amazon.com/Establishing-Foundations-Step-Step-Organizations/dp/0137424604), these options are aligned on the so-called “who builds it, who runs it” spectrum. The spectrum is shown in the figure below.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

(Image attribution: “[Establishing SRE Foundations](https://www.amazon.com/Establishing-Foundations-Step-Step-Organizations/dp/0137424604)”)

What is important to understand about the options on the spectrum are the incentives they provide for the development teams to implement reliability. With “you build it, you run it,” the incentives are maximized because developers are on-call and do not want to be woken up in the middle of the night due to reliability issues. This will prompt the developers to do everything possible to implement reliable services, though it does add yet another responsibility to developers. These incentives diminish with every other option.

With “you build it, ops run it,” the incentives are minimal and can lead to the notorious chasm between development and operations teams. The chasm results in developers throwing their code over the wall to operations engineers. In this case, neither the code is written with operability in mind nor the operations engineers possess the knowledge to operate it. We therefore exclude this option in the considerations below.

Other differences between the options on the “who builds it, who runs it” spectrum include knowledge synchronization between teams, incident resolution times, service handover for operations, establishment of an SRE organization, etc.

# Setting up an organizational structure for SRE

Once an organization selects an option from the “who builds it, who runs it” spectrum, they can set up an organizational structure for SRE. To do so, the following questions need to be answered:

- Which teams are in the development organization?
- Which teams are in the operations organization?
- Which teams are in the SRE organization, if it is to be created?

The cross product of

- you build it, you run it
- you build it, you and SRE run it
- you build it, SRE run it

and

- development organization
- operations organization
- SRE organization

yields nine sensible SRE Team Topologies. These are described in detail in [_Establishing SRE Foundations_](https://www.amazon.com/Establishing-Foundations-Step-Step-Organizations/dp/0137424604). In the next section, we provide an overview of the topologies.

# SRE Team Topologies

The SRE team topologies are embedded in the development, operations, and SRE organizations of an enterprise. To avoid ambiguity, here are the primary responsibilities of the three organizations:

| Organization | Primary responsibilities |
| --- | --- |
| Development organization | Build products Depending on the SRE team topology: Run products to the extent agreed |
| Operations organization | Provide tools as a serviceDepending on the SRE team topology:Build and run the SRE infrastructure Run products to the extent agreed |
| SRE organization | Depending on the SRE team topology: Build and run the SRE infrastructure Run products to the extent agreed |

That is, a selected SRE team topology determines to a great extent the primary responsibilities of the development, operations, and, if it exists, SRE organization. Below is the list of nine SRE team topologies from [_Establishing SRE Foundations_](https://www.amazon.com/Establishing-Foundations-Step-Step-Organizations/dp/0137424604).

**SRE Team Topology 1:**

| Development organization | You build it, you run it with no dedicated SRE role. Every developer is an SRE on rotation |
| --- | --- |
| Operations organization | SRE infrastructure team |
| SRE organization | Does not exist |

This is a classic “you build it, you run it” SRE team topology as followed by [Amazon](https://queue.acm.org/detail.cfm?id=1142065), for example.

**SRE Team Topology 2:**

| Development organization | You build it, you run it with a dedicated SRE role in the team |
| --- | --- |
| Operations organization | SRE infrastructure team |
| SRE organization | Does not exist |

This SRE team topology introduces a dedicated SRE role in the development team. That is, unlike the SRE team topology 1, not every developer is an SRE on rotation here.

**SRE Team Topology 3:**

| Development organization | You build it, you run it with a dedicated SRE role in the team and a dedicated developer on rotation |
| --- | --- |
| Operations organization | SRE infrastructure team |
| SRE organization | Does not exist |

This SRE team topology is a combination of the SRE team topologies 1 and 2. There is a dedicated SRE role in the team that runs the product together with another developer on rotation.

**SRE Team Topology 4**

| Development organization | You build it, you and SRE run it with a dedicated SRE team |
| --- | --- |
| Operations organization | SRE infrastructure team |
| SRE organization | Does not exist |

This SRE team topology introduces a dedicated SRE team placed in the development organization. The members of the SRE team run the product in a shared on-call together with the developers from development teams.

**SRE Team Topology 5**

| Development organization | You build it, you & SRE run it |
| --- | --- |
| Operations organization | Dedicated SRE team and SRE infrastructure team |
| SRE organization | Does not exist |

This SRE team topology places a dedicated SRE team into the operations organization. Like in the SRE team topology 5, the members of the SRE team run the product in a shared on-call together with the developers from the development teams.

**SRE Team Topology 6**

| Development organization | You build it, you and SRE run it |
| --- | --- |
| Operations organization | SRE tool chain procurement and administration |
| SRE organization | Dedicated SRE team and SRE infrastructure team |

This SRE team topology introduces a dedicated SRE organization. The SRE team running the product together with the development teams is in the SRE organization. The SRE infrastructure team building and running the SRE infrastructure is in the SRE organization too. The shared on-call is the same as in the SRE team topologies 4 and 5. This is roughly the SRE team topology employed by Facebook with their production engineering organization. At Facebook, it is called the “[centralized reporting, embedded locality](https://thenewstack.io/facebooks-integrated-approach-to-building-engineering-teams/)” model.

**SRE Team Topology 7**

| Development organization | You build it, SRE run it with a dedicated SRE team |
| --- | --- |
| Operations organization | Dedicated SRE infrastructure team |
| SRE organization | Does not exist |

This SRE team topology places the responsibility of running the product onto a dedicated SRE team placed in the development organization. However, if the services fall below an agreed service level, the SRE team “returns the pager” to the development team until the agreed service level is reached again.

**SRE Team Topology 8**

| Development organization | You build it, SRE run it |
| --- | --- |
| Operations organization | Dedicated SRE team and SRE infrastructure team |
| SRE organization | Does not exist |

This SRE team topology places the responsibility of running the product onto a dedicated SRE team placed in the operations organization. As in SRE team topology 7, if the services fall below an agreed service level, the SRE team “returns the pager” to the development team until the agreed service level is reached again.

**SRE Team Topology 9**

| Development organization | You build it, SRE run it |
| --- | --- |
| Operations organization | SRE tool chain procurement and administration |
| SRE organization | Dedicated SRE team and a dedicated SRE infrastructure team |

This SRE team topology places the responsibility of running the product onto a dedicated SRE team placed in the SRE organization. As in SRE team topology 7, if the services fall below an agreed service level, the SRE team “returns the pager” to the development team until the agreed service level is reached again. This is the SRE team topology employed by [Google](https://itrevolution.com/articles/how-google-sre-and-developers-collaborate/).

In addition to the differences in organizational structure, different SRE team topologies vary in other areas such as knowledge synchronization between teams and organizations, effort for service handover for operations, incident resolution times, and more. An often overlooked difference is the SRE cultural identity created by an SRE team topology.

# SRE cultural identity

An SRE cultural identity is based on three identity dimensions: a product-centric identity, an incident-centric identity, and a reliability user experience-centric identity. A product-centric SRE identity is when SREs strongly identify themselves with the product they run. They are not just SREs, they are (for example) Microsoft Office 365 SREs taking pride in the product. This is typical when SREs are placed in the development organization.

An incident-centric identity is when SREs are focused on having as few incidents as possible in products they run. These SREs pride themselves in metrics like only having just a few incidents a year. This is typical when SREs are placed in the operations organization.

A reliability user experience-centric identity is when SREs are focused on achieving the user experience of reliable products for the products they run. These SREs pride themselves in having SLOs tracking the user experience well, having the SLOs fulfilled by the products they run, etc. This is typical when SREs are placed in a dedicated SRE organization.

An SRE team topology spawns an SRE cultural identity triangle with the vertices: product-centric identity, incident-centric identity, and reliability user experience-centric identity. A particular SRE team topology will lean more towards one of the vertices on the SRE identity triangle.

# Transition to the selected SRE team topology

Once an SRE team topology has been selected, the question of transitioning from the current setup to the selected one becomes important. If a new SRE organization gets established during the transition, it needs to be positioned within the overall product delivery organization.

The SRE organization can be viewed as a cost center, an asset, a business partner, or a business enabler. The goal of the newly minted head of the SRE organization is to position the organization as much as possible to be the business enabler.

Within the SRE organization, an SRE career path needs to be established to provide a proper career ladder for SRE professionals as they grow their skill and practice. A defined SRE career path also helps attract SRE talent to the company.

# Summary

SRE principles can be fulfilled by many organizational structures. In this article, nine SRE team topologies were presented, which can be widely found in the industry. A decision to choose a particular SRE team topology needs to be made taking into account the current organizational setup and culture, the envisioned target organization and SRE cultural identity, knowledge synchronization requirements between teams, and other factors.

More details on how the decision can be made are available in the talk “[Establishing SRE Foundations: Aligning The Organization On Ops Concerns Using SRE Team Topologies](https://videos.itrevolution.com/watch/778246010/)” from the DevOps Enterprise Summit US 2022 and the corresponding book [_Establishing SRE Foundations: A Step-by-Step Guide to Introducing Site Reliability Engineering in Software Delivery Organizations_](https://www.amazon.com/Establishing-Foundations-Step-Step-Organizations/dp/0137424604) by the author.
