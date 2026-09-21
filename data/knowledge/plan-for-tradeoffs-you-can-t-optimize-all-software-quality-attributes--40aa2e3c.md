---
title: "Plan for tradeoffs: You can’t optimize all software quality attributes"
notion_id: 40aa2e3c-bec1-4f81-bac4-01a04ee25196
notion_url: https://app.notion.com/p/Plan-for-tradeoffs-You-can-t-optimize-all-software-quality-attributes-40aa2e3cbec14f81bac401a04ee25196
last_edited: 2026-09-21T17:37:00.000Z
source_url: https://stackoverflow.blog/2022/01/17/plan-for-tradeoffs-you-cant-optimize-all-software-quality-attributes/
tags: ["English", "Product Management", "Project Management", "System Design / Software Architecture", "Article", "Stack Overflow Blog"]
---
[https://stackoverflow.blog/2022/01/17/plan-for-tradeoffs-you-cant-optimize-all-software-quality-attributes/](https://stackoverflow.blog/2022/01/17/plan-for-tradeoffs-you-cant-optimize-all-software-quality-attributes/)



The next software app I use had better not have any bugs in it: no 404 “page not found” errors or help screens that don’t match the form with which I’m working. It shouldn’t use much memory or slow down my computer, and it ought to free up all the memory it used when it’s done. The app should be completely secure: no one can steal my data or impersonate me. It should respond instantaneously to my every command and be completely reliable. I don’t want to see any “internal server error” or “application is not responding” messages. The user interface should never let me make a mistake. I should be able to use the app on any device I want with instantaneous downloads and no timeouts. It should let me import and export any data I need from other sources. Oh, yes, I almost forgot—this app should be free too.

Doesn’t that sound like a fabulous app? It sure does! Are my expectations reasonable? Of course not!

It’s impossible to get the best of all possible worlds for every aspect of a software system’s capabilities and characteristics. There are inevitable trade-offs among various aspects of quality—increasing one often causes an unavoidable decrease in another. Consequently, an essential part of requirements analysis is understanding which quality characteristics are the most important so that designers can address them appropriately.

Software project teams must consider a broad set of quality attributes as they explore requirements. The terms design for excellence and DfX refer to quality attributes as well, where X is a property of interest that designers strive to optimize. When people talk about nonfunctional requirements, they’re usually referring to quality attributes.

Nonfunctional requirements aren’t directly implemented in software or hardware. Instead, they serve as the origin for derived functionality, architectural decisions, or design and implementation approaches. Some nonfunctional requirements impose constraints that limit the choices available to the designer or developer. For instance, an interoperability requirement could constrain a product design to use certain standard interfaces.

I’ve seen lists of more than 50 software quality attributes, organized into various hierarchies and groupings. Few projects will need to worry about accounting for so many attributes. The first list shows some quality attributes that every software team should consider as they learn what quality means for their product. Physical products that contain embedded software have some additional quality attributes, such as those shown in the second list (Koopman 2010, Sas and Avgeriou 2020).

Some important quality attributes for software systems

- Availability - Can I use the system when and where I need to?
- Conformance to standards - Does the system comply with all applicable standards for functionality, safety, communication, certification, and interfaces?
- Efficiency - Does the system use computer resources economically?
- Installability - Can I easily install, uninstall, and reinstall the system and upgrades?
- Integrity - Does the system protect against data inaccuracy, corruption, and loss?
- Interoperability- Does the system connect well with others to exchange data and services?
- Maintainability - Can developers easily modify, correct, and enhance the system?
- Performance - Does the system respond sufficiently quickly to user actions and external events?
- Portability - Can the system be migrated to different platforms easily?
- Reliability - Does the system run when it’s supposed to without failing?
- Reusability - Can developers reuse portions of the system in other products?
- Robustness - Does the system respond sensibly to erroneous inputs and unexpected operating conditions?
- Safety - Does the system protect users from harm and property from damage?
- Scalability - Can the system easily expand to accommodate more users, data, or transactions?
- Security - Does the system protect against malware attacks, intruders, unauthorized users, and data theft?
- Usability - Can users easily learn how to use the system to accomplish their tasks?
- Verifiability - Can testers determine whether the software was implemented correctly?

Some additional quality attributes for physical products containing embedded software

- Durability - Will the product hold up well under normal usage conditions?
- Extensibility - Can new functionality, sensors, or other hardware be added easily to the product without disrupting its functioning?
- Fault handling - Does the product detect, recover from, and log faults that occur?
- Manufacturability - Is the product easy and cost-effective to manufacture?
- Resource usage - Does the product retain enough slack capacity in the resources consumed for memory, network bandwidth, power, processor capacity, and so forth?
- Serviceability - Can people efficiently perform preventive and corrective maintenance on the product?
- Sustainability - Does the product have minimum adverse environmental impacts over its lifecycle, from extraction of raw materials to manufacturing, usage, and disposal?
- Upgradability - Can the product easily be enhanced by adding or replacing components?

One property that doesn’t appear in either of these tables is cost. As with product functionality, designers must balance the value of achieving some desired quality goal against the cost of achieving it. For instance, everyone would like the software they use to be available for use all the time, but achieving that can be expensive.

One of my consulting clients had a manufacturing control computer system with an availability requirement of 24 hours a day, 365 days a year (366 in leap years), with zero downtime acceptable. They met that requirement by having redundant computer systems. They could install software updates on the offline system, test the software, cut over to put that system online, and then update the second system. Having two independent computer systems was expensive, but it was cheaper than not manufacturing their product when the control system was down.

Software, hardware, and user experience designers all need to know which quality attributes are most important, which aspects of those often multidimensional attributes are paramount, and the target goals. It’s not enough to simply say, “The system shall be reliable” or “The system shall be user friendly.” The business analyst (BA), requirements engineer, product owner, or product manager needs to ask questions during requirements elicitation to understand just what stakeholders mean by “reliable” or “user-friendly.” How would we be able to tell whether the system was reliable or user-friendly enough? What are some examples of not being reliable or user-friendly?

The more precisely the BA can state the stakeholders’ quality expectations, the easier it is for designers to make good choices and assess whether they’ve reached the goal. Roxanne Miller (2009) provides many examples of clearly written quality attribute requirements in numerous categories. When possible, state quality goals in measurable and verifiable ways to guide design decisions. Consider using Planguage, a keyword language invented by software pioneer Tom Gilb that permits precise, quantitative, and measurable specification for such vague attributes as “availability” and “performance” (Simmons 2001, Gilb 2005). Specifying requirements this carefully takes some time, but that’s time well-spent compared to restructuring a product after it fails to meet customer expectations.

Designers can optimize their solution approach for just about any quality parameter, depending on what they’ve been told—or what they think—is most important. Without guidance, one designer might optimize for performance, another for usability, and a third for portability across delivery platforms. The project’s requirements explorations need to identify which attributes are more important than others to guide the designers in the most important direction for business success. That is, you need to prioritize nonfunctional requirements just as you do functionality.

Prioritization is essential because of the trade-offs between certain pairs of quality attributes. Increasing one quality attribute often requires the designer to compromise in some other areas (Wiegers and Beatty 2013). Here are some examples of quality attribute conflicts that demand trade-off decisions:

- Multifactor authentication (MFA) is more secure than a simple password login, but it reduces usability because of the additional steps (and possibly devices) involved.
- A product or component that’s designed to be reusable might be less efficient than if the code for that functionality were optimized for a single application. Provided the performance penalty is acceptable, however, it might still be sensible to create a reusable component.
- Optimizing a system for performance could reduce its portability if the developers exploited specific operating system or language properties to squeeze out every bit of performance.
- Optimizing certain aspects of a complex quality attribute could degrade others. For instance, within the broad usability domain, designing a system to be easy for new or occasional users might make the system less efficient for an expert.

On the other hand, some pairs of quality attributes exhibit synergies. Designing a system for high reliability will enhance several other attributes:

- Availability: If the system doesn’t crash, people can use it.
- Integrity: The risk of data loss or corruption from a system failure is reduced.
- Robustness: The product is less likely to fail because of an unexpected user action or environmental condition.
- Safety: If a product’s safety mechanisms work reliably, nobody gets hurt.

Quality attribute interactions demonstrate why the project team must understand early on what quality means to the key stakeholders and align everyone’s work toward those objectives. Stakeholders who don’t work with BAs to shape this understanding are at the mercy of designers who will make their best guesses. If you don’t explore nonfunctional requirements during elicitation and specify them precisely, you can’t ensure that designers will build in the properties that customers value.

Development teams need to understand which attributes require close attention early on so that they can make appropriate architectural design choices. A system’s architecture affects multiple attributes, including availability, efficiency, interoperability, performance, portability, reliability, safety, scalability, and security. Because compromises are often needed, if architects don’t know which attributes to prioritize, they might make design choices that don’t lead to the desired outcomes.

It’s costly to go back late in development or after release and reengineer the system’s architecture to remedy shortcomings in quality. Building systems incrementally without early knowledge of the most significant quality goals can lead to problems that might be hard to rectify, particularly if both hardware and software are involved. As is so common with software projects, spending a little more time up front to better understand the quality goals can lead to less expensive and more durable solutions.

Requirements explorations naturally focus on functionality. That’s what users see when they use an app, and that’s what designers focus on when they conceive solutions. However, exploring a product’s many, often conflicting quality characteristics from the beginning provides a solid foundation for customer satisfaction and robust design.

Gilb, Tom. 2005. Competitive Engineering: A Handbook for Systems Engineering, Requirements Engineering, and Software Engineering Using Planguage. Oxford, England: Elsevier Butterworth-Heinemann.

Koopman, Philip. 2010. Better Embedded System Software. Pittsburgh: Drumnadrochit Press.

Miller, Roxanne E. 2009. The Quest for Software Requirements. Milwaukee: MavenMark Books.

Sas, Darius, and Paris Avgeriou. 2020. “Quality attribute trade-offs in the embedded systems industry: an exploratory case study.” Software Quality Journal 28(2):505–534. https://doi.org/10.1007/s11219-019-09478-x.

Simmons, Erik. 2001. “Quantifying Quality Requirements Using Planguage.” http://www.geocities.ws/g/i/gillani/SE%272%20Full%20Lectures/ASE%20-%20%20Planguage%20Quantifying%20Quality%20Requirements.pdf.

Wiegers, Karl, and Joy Beatty. 2013. Software Requirements, 3rd Ed. Redmond, WA: Microsoft Press.

Tags: quality, design, quality attributes, nonfunctional requirements

Karl Wiegers is the author of Software Development Pearls, from which this article is adapted. He’s the Principal Consultant at Process Impact and the author of numerous books on software development, project management, design, and other topics.
