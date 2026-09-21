---
title: "Creating an integrated business and technology strategy"
notion_id: a53d8844-2ae9-43c8-bcce-2e9c4e4f96fc
notion_url: https://app.notion.com/p/Creating-an-integrated-business-and-technology-strategy-a53d88442ae943c8bcce2e9c4e4f96fc
last_edited: 2023-08-29T12:19:00.000Z
source_url: https://martinfowler.com/articles/creating-integrated-tech-strategy.html
tags: ["Article", "Martin Fowler", "English", "Entrepreneurship", "Decision Making", "Product Management"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Creating an integrated business and technology strategy

How do you create a technology strategy? The conventional approach suggests you start with your current state, determine your future state and build the roadmap to get there. But there is a nuance in that approach that isn't quite right. What often results from following this is a big wish list of all the things that could be done. A powerful technology strategy is as much about what is left out as it is about what is included. Furthermore, technology strategies are often created in isolation, separate from business or product strategies. They are commonly created after the business strategy has been agreed upon. The result being infeasible business strategies which can not be achieved without considerable cost or time.

The challenges with this conventional approach shouldn't be too surprising. After all, if you were to build a health strategy your doctor wouldn't start with a full body scan and tell you how to fix all the symptoms. They would start with your health objectives and the outcomes that you are after, then investigate if your body was capable of achieving your goals, and if not put a remediation plan in place.

I would like to challenge this conventional approach to creating technology strategies, and offer up a different way to create yours. Start with the objectives and outcomes of your organisation. As the organisation considers the different strategic directions that they could traverse to achieve the goals, follow specific lines of inquiry to investigate if your current environment is capable of achieving the proposed strategic direction. The recommendations that result from the different investigations inform the feasibility of that direction, and can be used to formulate a remediation plan. Additionally, because technology is considered as the business strategy is being formed, technology itself can be the driving force behind ideas for new revenue streams. In doing so, your technology strategy will be integrated with the business strategy because it is born together with the business strategy.

### How to use this article

In this article, we look at eleven prevalent strategic directions that organisations traverse, grouped into four broad categories.

| Section | Strategic Direction |
| --- | --- |
| Growing the business | [Expand to complementary products](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#expand-complementary) |
|  | [Expand to new markets or regions](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#expand-market) |
|  | [Expand customer segments](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#expand-segments) |
|  | [Inorganic growth](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#inorganic-growth) |
| Building a strong foundation | [Accelerate time to value with improved efficiency and productivity.](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#time-to-value) |
|  | [Increase customer satisfaction with improved product quality](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#quality) |
|  | [Reduce Cost and Minimize Operational Risk](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#risk) |
|  | [Enhanced competitive advantage by enabling data driven decision making](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#data-driven) |
| Supporting the people | [Culture](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#culture) |
|  | [Internal and back office systems](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#internal-systems) |
| Responding to the ever changing future | Emerging technologies and market trends |

For each strategic direction identified, we provide examples of the lines of inquiry that you can use to investigate how feasible the strategic direction is. We also provide activities that you can do to help answer the lines of inquiry. Many activities span across many lines of inquiry, which allow you to concentrate your efforts into the synthesis of the activity rather than the activity itself.

Here's the format we use for these directions:

### Direction

This is the implication on technology for this strategic direction

### Key Business Questions

- Any questions that you should ask the business to inform the decisions you make

Lines of Inquiry

### Name of the line of inquiry

Description of the investigation

Questions about the topic - click to expand

Some questions to ask that will guide your investigation

What to look for:

A description of things to look out for

Activities: Activities that may help the investigation

Questions about other aspects - click to expand

More questions to ask that will guide your investigation

What to look for:

A description of things to look out for

Activities: Activities that may help the investigation

By selecting the relevant line of inquiry, and using the example questions as a springboard to guide your investigation, you will be one the right tack to creating your own integrated business and technology strategy.

## Growing the business

There are just a few key avenues for business growth:

- Offering complementary products: This involves selling additional products or services to your existing customers. This can be a great way to increase revenue and customer loyalty.
- Expanding into new markets: This involves selling your products or services in new geographic regions or to new customer segments. This can be a great way to grow your business, but it can also be a challenge. It is important to carefully research new markets before expanding into them.
- Attracting new customer segments: This involves identifying new groups of people who could be interested in your products or services. You can do this by conducting market research, analyzing customer data, and identifying trends in your industry.

Organizations typically focus on one or two of these growth strategies at a time, while maintaining and growing their existing business. This is often achieved through organic growth, but at times, the growth is accelerated through mergers and acquisitions, joint ventures or other strategic alliances (ie inorganic growth), which accelerates value realization along one or more of these growth strategies.

### Expand to complementary products

Offering new products to your customers could result in significant changes throughout your systems, from the customer buying experience to shared-service backend systems such as payments and invoicing, distribution and warehouse management, and business reporting. How different the product is from your existing product suite will impact how large the change needs to be.

### Key Business Questions

- How different are the two product types? Will business processes need to change?
- What business capabilities can be shared across products? Eg Payments, Distribution, Inventory.
- What are undifferentiated capabilities that require significant changes? What products exist on the market today that meet the needs of these undifferentiated capabilities?
- Do you need a single customer view over the products?
- What is the opportunity cost to the existing product with expansion a complementary area? Will investment dilute across the products, degrading the experience of the original product?

Lines of Inquiry

### Product representation

The representation of product types in the code base can have a big effect on how easy it is to add new ones or adjust the categories as new products appear.

Questions about product representation

How are product categories represented in systems?

What to look for:

The representation of product types in the code base can have a big effect on how easy it is to add new ones or adjust the categories as new products appear. Look for assumptions in other systems for the product line information. Product codes may be hard-coded elsewhere, some systems may only assume certain kinds of product. Look for changes throughout the stack, from the UI as you offer different customer experiences and navigation through your site, to new entries in your domain model and potentially new table structures in your database.

Activities: Code inspection • Database inspection • Domain modeling

### Changes to business processes

Business processes may not be suited for future products under consideration. As a result, expanding into complementary products might necessitate larger system changes. If the blast radius is wide, how easy is it to make large changes through the system?

Questions about spread of change

How many other systems need to change if a product change occurs in one system?

What to look for:

Business processes may not be suited for future products under consideration. Changes may be need to be made to shared-service back-end systems such as payment and invoicing systems, distribution and warehouse management systems and business reporting

Activities: Business process mapping

### Shared Capability

As you add complementary products to your offering, you will need to identify which capabilities should be shared across the products and what special treatment is required to shared capabilities. Moving toward a digital platform architecture will allow you to reuse shared capabilities if you expose them via APIs.

Questions about shared capabiltiies

What capabilities are shared across product types? How do shared capabilities need to change to support new products?

What to look for:

Sharing capabilities across products allows you to focus on the value added differentiators of the new product. Consolidating the shared capabilities improves speed to market. However, be watchful over mandating the sharing of capabilities where the processes differ between product offerings as it may result in needless complication and debt within the capability.

Activities: Business capability mapping • Business process map

Questions about build or buy capabities

Should you build or buy the capability? Are there newer products on the market that can replace the undifferentiated capabilities that require significant change?

What to look for:

Capabilities that don't add to product differentiation can safely be assigned to packaged software, either installed or SaaS. Developing a list of such capabilities can drive considering current vendor offerings. Changes in product line can help clarify what capabilities are important to product differentiation. Re-examine the vendors offering suitable products, both due to changes over time and from reevaluating differentiating factors.

Activities: Business capability mapping • Vendor product scan • Operating Model Quadrant from Enterprise Architecture as Strategy • Build vs buy analysis

### Expand to new markets or regions

As you expand into new geographies, you will be faced with the challenges of running a global platform that needs to cater for regional differences brought about by different local integration's, differing buyer personas, and different processes. Some capabilities will need to be provided in a global platform, others need to have flexibility to allow for these regional differences.

You will also need to contend with government regulation requirements such as GDPR, data sovereignty and regulatory compliance like SOX and APRA. This impacts where your data is processed and where it is stored. It also might introduce new features to handle compliance requirements.

### Key Business Questions

- What are the differences between the customers in each market?
- What are the regulatory compliance requirements for the new market?
- Do you need to change language, unit formats or factor in time zone conversions? Are there any cultural differences that require UI changes (eg the colour black has different interpretations in Korea than North America)?
- Do you need to introduce new tax calculations or additional reporting requirements?

Lines of Inquiry

### Diversification or Rationalisation of capabilities

As you expand into new geographies, you will be faced with the challenges of running a global platform that needs to cater for regional differences brought about by different local integrations, differing buyer personas, different processes. Some capabilities will need to be provided in a global platform, others need to have flexibility to allow for these regional differences.

Questions about different capabilities

How does your platform support different capabilities? How do they change across the markets?

What to look for:

Identify the capabilities that can be diversified, replicated, unified or coordinated across the markets. Diversified and replicated capabilities can exist for each market, where you want to look to offering unified capabilities through a global platform.

Activities: Business capability mapping • Operating Model Quadrant from Enterprise Architecture as Strategy

### Regulation and compliance

You may need to contend with government regulation requirements such as GDPR, data sovereignty and regulatory compliance like SOX and APRA. This impacts where your data is processed and where it is stored. It also might introduce new features to handle compliance requirements.

Questions about infrastructure and deployment

How easy is it to replicate your infrastructure?Do you have one-click deployment of your infrastructure-as-code?

What to look for:

If you need to host your infrastructure in the new market due to regulatory, compliance or even performance reasons, you need to make deployment repeatable across your systems. This reduces inconsistencies across regions, which leads to high triage times when a fault occurs. It's too common for organizations to set up separate teams to customize and operate each region, resulting in configuration drift / snowflakes as code, and manual maintenance. Equal emphasis needs to go to maintenance - cost, staff effort, and results (things are patched, consistent, old versions retired, fixes of all sizes rolled out quickly and comprehensively). The running costs to scale geographically becomes near-linear, and improvements and fixes are not distributed quickly or consistently.

Activities: Deployment strategy • Infrastructure as code

Questions about data centres

Does your infrastructure run in data centers or the cloud? How long will it take to commission a new data center in a new country?

What to look for:

Commissioning new data centers in new markets in order to be compliant may take a long time. If there is no repeatability in the process, and it takes a while to commission, you could consider a move to the cloud. Cloud-based infrastructure will be easier to adhere to data sovereignty requirements in new markets but you may need to look at ways to partition your database to be compliant.

Activities: Infrastructure architecture diagrams • Data storage

Questions about third party system compliance

How do your third-party systems handle data? Will they make you non-compliant?

What to look for:

Look for any third-party systems which make you non-compliant by passing regulated information outside the country. This might include an examination of OLAs and SLAs.

Activities: 3rd party inspection, review of OLAs and SLAs

### Internationalisation

Expanding into new markets may introduce internationalisation changes such as languages, time, money and units. There may also be taxes to consider, and new timezone or daylight saving changes to be factored into.

Questions about language and unit conversion

How easy is it to translate language files and unit formats?

What to look for:

Look for UI frameworks that enable this by default. Retrofitting internationalisation into UIs can be a tiresome process. Content length can change across languages. Look for dynamic UI elements that can accommodate longer or shorter text elements. If the UI framework has already been configured for translation files, a nice experiment to run is to change all translations to something small like “XXX”, and to something very long to see how the page responds.

Activities: Code inspection

Questions about time, money and units

How easy is it to introduce new tax calculations, or date/time conversions? What assumptions does the code make about the units or the currency it uses?

What to look for:

Look for well factored code bases that isolate changes.

Activities: Code inspection

Questions about language translations

What is the process to translate the content? How does this process impact continuous delivery?

What to look for:

Will you need to add time into your development process to allow for an agency to translate your files? This extra wait time can affect your small feedback loops. Alternatively, will you need to introduce the translation process as part of the design process?

Activities: Path to production • Value stream mapping

### Expand customer segments

Ideally, offering your existing product to new customer segments should see little changes through your system. However, at times new customer segments could introduce new operational processes, new customer journeys or new channels experiences. For example, a bank that expands existing credit cards into the sub-prime market introduces completely new operational processes to manage increased risk around debt, regulatory issues due to responsible lending, new way of doing collections (due to higher numbers and earlier interventions) and new marketing strategies. Where as a move from B2C to B2B might mean introducing a new API customer channel. Expanding to more mobile customers might need move to a native mobile experience.

You might want to gather different customer insight as you move, including customer sentiment, adoption and usage so you might add new reporting requirements or alter existing reports to also report on the new segments.

### Key Business Questions

- What are the differences between the customers segments?
- What operational process changes are required to support the new customer segment?
- Are you moving between B2C and B2B?
- What are the different expectations of interaction for the new customer segment?

Lines of Inquiry

### Customer journey changes

A new customer base might need new customer experiences. If their customer journey differs from your existing customer journey, you will need to make changes to your system

Questions about front end code

How easily can you offer a new experience, new pages, new navigation or new portal?

What to look for:

Front end code which is difficult to change will make it challenging to offer new customer experiences. Look for hard-coded navigation elements, difficult to change front end code, and lack of front end tests.

Activities: User experience debt • Front end code inspection • Test coverage

Questions about front end testing

How is the front end code tested?

What to look for:

Untested code, especially untested JavaScript, will make it risky to change the experiences. Front ends which are only tested through slow (and flaky) end-to-end tests will increase the time it takes to develop new experiences.

Activities: Test coverage

### Channel strategy

A new customer experience might also open the need for different customer channels. Moving from a B2C to a B2B offering may provide you with an opportunity to streamline your experience behind customer-centric APIs. Expanding into partner networks will require integration into the networks.

Questions about different UI experiences

Will you need to offer a different mobile or website experience?

What to look for:

When incorporating a new UI experience, such as extending your website to a native mobile app, consider how your digital platform accommodates and integrates core business functionalities while providing flexibility for various front-end interactions. If you have a microservice or micro front-end website in place, what architectural adjustments such as BFF (Backend For Frontend) mobile adapters are required.

Activities: Architecture design

### Inorganic growth

Inorganic growth through mergers and acquisitions (M&A), joint ventures, or other strategic alliances is often a faster way to accelerate the growth of a business along the three aforementioned axes. It in itself drives a different line of investigation.

### Key Business Questions

- What were the value drivers for the acquisition? How are you protecting these?
- What is the long term view of the acquisition - are you merging it into the business over time, or keeping it separate?
- What are the long term plans for the acquisition? Will you divest this asset once your company grows?

Lines of Inquiry

### Independently run businesses

This is the case where the acquired businesses continues to run independently to the organisation. This might be the first phase of the acquisition, or it might be to enable an easy sale of the asset in the future. While the technology organisations and the systems remain separate, there is value in loose-coupling (e.g. via restful APIs) to integrate the two running systems.

Questions about shared business capabilities

What business capabilities need to be integrated? Are these capabilities exposed through APIs? Do the APIs expose the inner workings of the systems, or are they nice facades which describe the behaviour? How stable are the APIs? Public-facing APIs should be as stable as possible, because of the large amount of coordination (and work) across the partner ecosystem. How secure are the APIs?

What to look for:

Even though the acquired businesses continues to run independently, there is an assumption that some business capabilities need to be integrated (e.g. Finance) but possibly not all. A business capability map is a quick first step to illustrate the cross-over at a business level before dropping to the API. API Strategy is a way to unlock existing business capabilities by building an API ecosystem to allow innovation at scale. It helps reduce time to market, by creating an ecosystem of APIs that are easy to operate, integrate and consume. APIs differ from traditional approaches that focus on system integration. By focusing the integration of the businesses through well defined APIs you will increase your ability to innovate into the future.

Activities: Business capability mapping • API strategy review and inspection of existing APIs

### Tight integration independently run

This is the case where you have two businesses running independently but you want tight integration between the two so that you can amplify the customer value created.

Questions about API strategy

What is the API strategy? Are different integration options available than publicly-facing APIs?

What to look for:

Investigation into the API strategy applies as equally in the independently run case as it does in this case, however there are a few other levers that you can take advantage of. For instance, you could share event streams or take advantage of shared storage.

Questions about common domain models

How different are the domain models? Is there alignment across the two, or is there significant work required to represent similar concepts across the systems? Are identifying entities such as customers the same across organisations? How can you connect the two entities? How will data be replicated across the systems?

What to look for:

As you look to work closely with the acquired business, you will need to look to consolidate across domain models. If you focus on getting this alignment in the domain model, the data matching and integration of systems will become easier to configure.

Activities: Domain modeling and consolidation

Questions about user experience across SSO or unified dashboards

Can you offer a seamless user experience to your customers by offering single sign-on, or dashboard views of the acquired company within your own systems? Can you get better or more transparent intelligence for support teams across the two systems?

What to look for:

A unified customer experience across the two products will increase customer satisfaction and improve retention. It will also enable the amplification of the value that both businesses together provide your customers, as the customer doesn't face the burden of working across two disjointed systems.

Activities: Customer journey mapping

### Complete merge

This is the case where the acquired systems will be first class systems within the ecosystem. This results in the rationalisation and consolidation of systems, migration of systems into the one system, integration into the runtime system, and merge of operational systems for observability and monitoring. It might also introduce different data archiving mechanisms

Questions about business capabilties

What capabilities now exist within the organisation? Which systems need to be rationalised? Common examples are CRM, CMS and payment systems

What to look for:

The Operating Model Quadrant from Enterprise Architecture as Strategy, is relevant for organizations wanting to unify capability across their group of companies, or when addressing consolidation of inorganic growth.

Activities: Business capability mapping

Questions about data migration and security

How do you migrate the acquired systems into your technology stack? What is the security posture of the acquired systems? Will there need to be work done to update the runtimes and libraries? What data needs to be migrated into the new system? What data can be archived? What data can be deleted?

What to look for:

Successful acquisitions dedicate resources and people to integrating the businesses. As a technology leader your input maybe required to help the integration efforts understand how the systems will be integrated. Runtime environments may need to be consolidated, data might need to be migrated into the new system, or archived for compliance reasons. A security review of the systems may highlight security remediation that needs to happen within the system before the migration. The systems themselves may need consolidation as libraries and frameworks are unified across the technology stack, to improve the developer experience moving across products.

Questions about operational aspects

What operational aspects need to be consolidated? How do you migrate log data into the new systems? Do you need to update log formats? What runtime information should be surfaced for improved observability?

What to look for:

Streamlining and consolidate the operational aspects of the systems reduces the operational cost and time it takes to manage across multiple disparate systems. You may need to change how running systems are observed, which could include changing the logging information, frequency or changing severity level of logs.

Activities: Cross functional requirements • Review instrumentation used to operate the systems

## Building a strong foundation

Any business that wants to grow needs to be built on strong and stable foundations. Oftentimes, technology strategies focus singularly on the ways that technology leaders can improve and continue to build a strong foundation, and so this section may feel familiar to technical readers. However for an integrated business and technology strategy to be successful, business leaders also need to understand how this focus will enable and support the business to grow. It is important that the improvements to the engineering organisation and the platforms they look after align with the themes that resonate with the rest of the organisation. These themes are:

- Accelerate time to value with improved efficiency and productivity. A good technology foundation can help businesses automate tasks, streamline processes, and improve communication. This can lead to significant improvements in efficiency and productivity, which can free up time and resources for other areas of the business.
- Increased customer satisfaction with improved product quality. A strong technology foundation can help businesses provide better customer service, offer more personalized experiences, and make it easier for customers to do business with them. This can lead to increased customer satisfaction, which can lead to increased sales and revenue.
- Reduce cost and minimize operational risk. A good technology foundation looks to reduce costs of the technology systems if the business need is to contain costs. Effective operational risk management helps companies prevent or minimize losses and safeguard their operations and reputation.
- Enhanced competitive advantage by enabling data driven decision making. A strong technology foundation can help businesses stay ahead of the competition by giving them access to new technologies, data, and insights. This can help them develop new products and services, improve their marketing and sales efforts, and gain a competitive edge in the market.

### Accelerate time to value with improved efficiency and productivity.

Accelerating time to value reduces the time it takes to deliver measurable benefits or achieve desired outcomes from a particular initiative, product, or project. It focuses on maximizing the efficiency and effectiveness of the value creation process. Three areas that negatively impact time to value include the inability to easily and confidently change code, poor developer experience and waste within the delivery process.

### Key Business Questions

- Is the current pace of delivery keeping up with the pace of customer demand?

Lines of Inquiry

### Code quality

Well written and structured code, supported by appropriate test coverage is easy to modify to new feature requests. Teams can go at a rapid pace when they are confident that changes they make will not inadvertently introduce hidden defects.

Questions about code structure

Is the code well structured? Does the code follow the common patterns and practices of the language? Is it at least internally consistent?

What to look for:

Does it mirror the domain? Consider this at multiple levels: within classes or files, all the way to component boundaries. How do they interact with each other? (consider domain modeling). Dimensions to assess the code against are: size, complexity, coupling, cohesion.

Activities: Code toxicity analysis

Questions about test strategy

Are there tests? Do the tests follow the test pyramid? If not, where are the gaps? If you change the code, does this also break the tests, is it possible to run unit and integration tests independently? Are more advanced techniques like parameterized or mutation testing being used?

What to look for:

Test coverage is easy to measure but won't always give a good indication of the test quality - other things to look at include number of tests; flakiness; execution time; consistency/ structure of tests; naming; and coupling to code

Activities: Test code coverage • Build times, failures, historic data

Questions about defects

Where are defects found?

What to look for:

Defects caught in pre-prod or prod have exponentially higher cost to remediate and disrupt the flow of value added work

Activities: Build times, failures, historic data

### Developer experience

[Developer Experience](https://martinfowler.com/articles/developer-effectiveness.html) and experience is key to engineering excellence leading to desired business outcomes and organisational performance. Developer experience (DX) encompasses all aspects of a developer’s interaction with an organisation, its tools and systems. Engineering platforms, that provide self-service capabilities help automate and streamline every stage of software development journey from ideation to go-live and collecting customer feedback leading to excellent developer experience

Questions about feedback loops

How quickly can teams get feedback on changes they make?

What to look for:

Look for long feedback loops such as long build times, validating cross functional requirements are upheld

Activities: Value stream map

Questions about access to knowledge

How easy is it to find the right information at the right time? How do new team members learn about the code? Is the code well documented?

What to look for:

What documentation exists, is it up to date, is it relevant / well organized / easy to find, does each repo have clear documentation about purpose of code and how to test and run it.

Activities: Documentation inspection

Questions about developer productivity accelerators

What accelerators, starter kits, paved roads are available? How long does it take a new team member to become productive?

What to look for:

Sensible defaults, convention over configuration, safe guard rails and good onboarding documentation improve the speed to productivity for new team members.

Activities: Measure time to productivity for the last batch of new hires

Questions about cognitive overload

How much cognitive overload do people face? How often do people need to context switch? What is the cost of misunderstood integrations, abstractions, and data?

What to look for:

Cognitive taxes create quality issues, slowing delivery significantly.

Activities: Value stream map

### Delivery Process

Remove the waste that exists within your delivery process. The waste is negatively affecting your speed to market. Waste may be in handoffs between groups, approval boards that slow down continuous releases, or rework in the system.

Questions about waste

Where is the waste within the delivery process? How many digital or human interactions are needed to in the delivery process? Do systems add to or detract from the optimum flow of the role being played?

What to look for:

Teams are often caught up on building a lot of things that are non-essential to business value. Examples of waste within the delivery process include cognitive load/context switching, ability to find the right information at the right time, friction in the development experience and slow quality feedback loops. Look for developer ecosystems that are fragmented with a need to navigate into multiple places and systems to get the job done. Measuring the size of work queues between stages of workflow is a good way to quantify bottlenecks. Look at flow metrics and cycle and lead times.

Activities: Value stream map • Measure lead and cycle times • Measure flow metrics

Questions about developer friction

Where do our teams face friction in delivering on our big bets? How easy is it to make a code change?

What to look for:

Measure time to put a single line code change into production. Long lag times for small changes discourage frequent updates, leading to larger, riskier updates and reduce responsiveness to business opportunities.

Activities: Value stream map

### Increase customer satisfaction with improved product quality

Improving your product quality increases customer's satisfaction and can have a big impact on customer retention. Product quality is impacted by the build quality itself, but is also impacted by performance issues, technical debt and complexity which results in increased reliance on call centres. There are often dangerous parts of systems that teams are reluctant to change due to the technical debt and lack of a safety net around it. If any feature development needs to be done in those areas, it will be slower and deployment will be riskier unless you tackle that debt. Maybe now is the time.

### Key Business Questions

- Why do customers stop using your product? What are customer retention rates?

Lines of Inquiry

### Address build quality issues

Unfortunately, IT systems have notoriously been buggy which impacts the overall product quality. Modern agile software delivery practices have significantly increased the build quality of modern codebases. Nevertheless, systems continue to have areas which have defects, or are difficult and complicated to understand which makes changes in these parts risky.

Questions about defect rates

What areas have the highest defect or incident rates?

What to look for:

A heatmap of defect rates for the components can show the most significant areas to improve first.

Activities: Incident and defect reports

Questions about risky parts of the codebase

Where are the riskiest parts of the codebase? Which parts do teams fear changing? What upcoming features need to change these areas? What is the cost of change in relation to cost of permanent improvement fix?

What to look for:

If future feature development is planned for the risky parts, it might be more effective to properly fix this area. Typically architectural changes are required to address the root cause of these areas..

Activities: Interview with teams • Analysis of historical velocity of changes through different components • Feature improvements, business initiative review

### System performance under increased load

It will increase the number of site visits to your website and increase your customer load. How capable is your system today of performing with the anticipated new load? Look for signs that your system is struggling under regular days as well as irregular events such as Black Friday sales. Identifying the breakpoints on these days gives you clues to what needs to be addressed for any additional load.

Questions about handling expected loads

How far can you scale up? Can you handle your peak expected volumes today? What is your equivalent of Black Friday sales? How does the site currently handle the increased load?

What to look for:

Look for periods where you are reaching capacity and notice patterns over time. If you are nearing capacity at the anticipated loads after growth you will need to start addressing system performance now.

Activities: Incident and performance log inspection • Performance, load and soak tests • Chaos engineering testing • Experimentation to determine bottlenecks and headroom in volume • Tests around dynamic scaling of infrastructure

### Call center complaints

To uncover areas for improving product quality that will make the biggest impact on your customers, consult your customer support team. They often hold valuable insights into where your systems are falling short of customer expectations. However, be cautious of survivorship bias, similar to the WW2 planes returning with bullet holes - it's important to consider not only customer complaints, but also where and why customers drop out of your funnel to get a comprehensive understanding of the issues.

Questions about call centre patterns

Which product features receive the most call centre interactions? What do the call centre staff spend the majority of their time doing?

What to look for:

Look for patterns in the call centre data. Match any patterns in this data with a heatmap across the product defect list.

Activities: Interview with call centre teams • Review of call centre data analysis

### Address technical debt

Technical debt is a metaphor for choosing an easy solution now that will make it harder to make changes or add new features later. It is often incurred when developers choose to use quick hacks or workarounds instead of taking the time to write clean, well-organized code. The technical debt that is building up in your system can have significant impact on the quality of the product.

Questions about user experience debt

Where is your User Experience debt? What improvements could be made to improve your product quality?

What to look for:

Activities:

Questions about how code changes

How frequently is code changed? How many people are touching the code and how often (ownership / knowledge)? Has the code been recently changed?

What to look for:

Look for the components which frequently change together which imply they are tightly coupled. Hotspots of change indicate a high churn rate.

Activities: Git commit log analysis

### Reduce Cost and Minimize Operational Risk

Operational risk refers to the potential loss resulting from internal process failures, system issues, or external events. This includes errors, fraud, system failures, and other disruptions that can impact business operations and financial performance. Examples include cyber-attacks, employee errors, and supply chain disruptions, which can affect a company's reputation and regulatory compliance. Effective operational risk management helps companies prevent or minimize losses and safeguard their operations and reputation.

### Key Business Questions

- What are the biggest risks relating to technology on the Risk Register?

Lines of Inquiry

### Cloud cost optimization

Cloud is increasingly a major component of technology budgets. In a lot of cases, actual cloud costs are exceeding budgets and the savings that helped rationalize a move to the cloud are not materializing. Understanding a business's cloud spend and sizing it appropriately to reduce waste and optimize the rates paid requires cooperation across the whole organization, with finance, product and engineering teams working together to ensure that the cloud spend is in proportion to the value the assets bring in.

Questions about cloud cost spend

Do you have visibility into your cloud spend? Does the cloud bill exceed budgets? Can you reconcile cloud spend to business consumption?

What to look for:

With the move to the cloud, some financial controllers and technology leaders alike have lost the financial visibility that they once had over infrastructure bought and run in data centres. Infrastructure buying centres have moved from the procurement office to the engineers keyboard, with credit cards stored on file with cloud vendor accounts. All too often, financial offices are only seeing the cloud spend once the bill arrives at the end of the month. There is a disconnect between the engineers who are spinning up new infrastructure and the controllers of the budget.

Activities: Cloud bill analysis

Questions about cloud cost controls

What cloud cost controls or alerting are in place?

What to look for:

Democratising the financial decision making around cloud spend to the engineers that are responsible for instantiating the cloud components needs to be done with the right financial controls and alerts in place. Influencing the culture of the organisation to operate on the cloud is important to help optimise cloud costs with governance and collaboration as a discipline. Tagging infrastructure and surfacing the spend on operational dashboards, alongside uptime and usage metrics and alerting, will provide the right balance between financial control and governance.

Activities: Tagging infrastructure • Operational dashboard configuration

Questions about orphaned systems

What orphaned systems are you still being billed for that you could turn off?

What to look for:

Retiring or sunsetting systems that you are no longer using is the fastest way to reduce your cloud bill. Take a critical look at everything that appears on the bill to make sure that you are still getting value from it.

### Data governance

Digital operation gives us the opportunity to capture data from every aspect of the business, together with the opportunity to analyze this data to better understand how everything works. But this data also presents a risk, as privacy violations can undermine the benefits.

Questions about data governance

What degree of data governance is applied across the organization? Can you trust the data? What is the quality of the information that is stored?

What to look for:

Data governance refers to the overall management, control, and protection of an organization's data assets. It encompasses the processes, policies, and standards that ensure the quality, integrity, availability, and security of data throughout its lifecycle. Data governance aims to establish a framework that governs how data is collected, stored, used, and shared across an organization.

Activities: Review of data management, policy, control and organisation philosophy towards its treatment of data

### End of life software

The use of products or systems which have reached or are approaching the end of life (EOL) agreements with the supplier poses significant risks to an organization. The risks include security vulnerabilities for products that no longer receive security patches from the vendor, compliance and regulatory non-compliance which may result from using unsupported versions, business disruption from failure of the systems, data loss and recovery challenges, and increased concentration risk if the vendor goes out of business. To mitigate these risks, organizations must adopt proactive end-of-life management strategies. This includes regularly evaluating the product lifecycle of their technology stack, planning for timely upgrades and replacements, and ensuring compatibility with future systems.

Questions about software versions that are in use

Are any technologies or framework/library versions used approaching End of Life (EOL) support?

What to look for:

End-of-life (EOL) software versions are no longer supported by the software vendor. This means that the vendor will no longer provide security updates, bug fixes, or new features for the software. As a result, using EOL software can pose a number of risks to businesses, including security vulnerabilities, performance problems, compliance issues and data loss. EOL versions provides an opportunity for a current market scan to migrate onto newer technologies that can enable the next growth phase.

Activities: Review of vendor contracts, inspect vendors EOL version list

### Supply chain disruption

The SolarWinds hack, also known as the Sunburst hack, was a major cyberattack that occurred in 2020. The hackers targeted SolarWinds, a Texas-based company that provides IT management software to businesses and governments around the world. The hackers were able to insert malicious code into SolarWinds' Orion software, which is used by thousands of customers. This allowed the hackers to gain access to the networks of these customers, including government agencies, Fortune 500 companies, and think tanks. The SolarWinds hack was a major security breach, and it has raised concerns about the security of supply chains. SolarWinds is a trusted supplier of software to businesses and governments, and the fact that it was hacked shows that no organization is immune to cyberattacks. How would your organisation respond to similar disruptions in the supply chain?

Questions about supply chain distruption

What could disrupt your digital supply chain? What vulnerabilities exist within our system? How can we address these?

What to look for:

By having backup plans in place in case of disruptions, companies can reduce the impact of disruptions on their operations. Use technology to improve visibility into the supply chain. By having real-time data on the location of goods and materials, companies can identify and respond to disruptions more quickly.

Activities: Review of supply chain

### Enhanced competitive advantage by enabling data driven decision making

Providing better access to data insights is a way to support internal staff in making data-based decisions. While many people claim to make data-driven decisions, in reality, they often make a decision first and then look for data to support their choice. This is what many data platforms and "business intelligence" systems focus on. However, true data-driven decision-making involves sensing what is happening in your environment and using that information to understand what decisions need to be made. A sound technology strategy should prioritize creating meaningful decision-making capabilities within your organization through a robust data platform. This can be achieved by embedding intelligence and machine learning into every decision, customer touchpoint, service, and product you provide. This approach can lead to improved decision-making, streamlined operations, and better customer experiences based on new data-driven insights that go beyond intuition and gut feelings.

Lines of Inquiry

### Enabling easy data access

Data Platforms enable data consumers to easily discover and access data they need in the right format at the right time, for effective decision-making and creating data solutions using the right set of tools to create maximum business value.

Questions about ease of data access

How easy is it for anyone in the organization to access relevant data? Is it self-service? What is the turn around time? Who owns and protects the data?

What to look for:

Define data architecture to cater to the analytics use cases. Design and execute data platforms based on best practices from data engineering, software engineering and continuous delivery. You want to be able to scale with ease with different data sources and users, enabling analytics use cases leading to richer insights.

Activities: Value stream map • Interview with internal teams to understand their usage and needs of business data • Data platform

Questions about data agregation

To what extent is data aggregated across business lines, products, sales, revenues, complaints etc.

What to look for:

Easy access to the data may not be as useful to the organisation if the data is a disconnected bunch of facts. Build curated views and governed datasets with data from various sources for analytics, machine learning, services, applications, etc.

## Supporting the people

Technology leaders play a vital role in supporting the organisation and its employees by leveraging technology to improve business processes, enabling data-driven decision making that is required to drive innovation, efficiency, and strategic growth.

> 

Contrary to popular belief, digital transformation is less about technology, and more about people. You can pretty much buy any technology, but your ability to adapt to an even more digital future depends on developing the next generation of skills[[1]](https://martinfowler.com/articles/creating-integrated-tech-strategy.html#footnote-talent)

Having a robust digital talent strategy is a competitive advantage in today’s fiercely competitive market. This enables businesses to have the right talent and have the right competencies to meet current and future demand to meet business goals or to stay on track for digital transformation aspirations.

### Culture

You might be thinking that culture is too touchy feely to go into a technology strategy. However, according to research by DevOps Research and Assessment (DORA), “organizational culture that is high-trust and emphasizes information flow is predictive of software delivery performance and organizational performance in technology”. The DORA research has also been backed up with further research from [Google's Project Aristotle](https://rework.withgoogle.com/print/guides/5721312655835136/)

The DORA report cited research by sociologist Dr. Ron Westrum. Westrum's research noted that such a culture influences the way information flows through an organization. Westrum provides three characteristics of good information:

- It provides answers to the questions that the receiver needs answered.
- It is timely.
- It is presented in such a way that the receiver can use it effectively.

DORA research shows that changing the way people work changes culture; this is echoed in the work of John Shook, who spoke of his experiences in transforming culture: "The way to change culture is not to first change how people think, but instead to start by changing how people behave—what they do." DORA group [provides useful information](https://cloud.google.com/architecture/devops/devops-culture-westrum-organizational-culture) on how to implement better organizational culture.

Lines of Inquiry

### Leadership

Leadership plays a pivotal role in establishing and shaping organizational culture. They set the tone and values, serve as role models, and communicate the organization's vision and values. By making hiring and promotion decisions aligned with cultural fit, empowering and holding employees accountable, recognizing and rewarding cultural alignment, and adapting the culture to changes, leaders foster a strong and sustainable culture. Their long-term vision and consistent efforts to embed cultural elements influence employee behavior and ultimately impact the organization's performance.

Questions about culture and leadership

Do we have the right culture and capabilities within our organization to thrive? What is the leadership style we promote? Is it conducive to a positive culture?

What to look for:

A strong leader should have a clear vision and purpose, communicating it effectively to inspire others. Emotional intelligence is vital, as is acting with integrity and ethics. Decisiveness, empowerment, adaptability, collaboration, and team-building are essential traits for fostering a positive work environment. Accountability, resilience, and a focus on results are also indicative of effective leadership. Additionally, good leaders prioritize mentorship and development to ensure the growth of their team members, contributing to the organization's long-term success. Look for leaders that promote diverse and equitable culture that promotes inclusion.

Activities: Employee surveys e.g. CultureAmp or Workday Peakon Employee Voice surveys

### Knowledge sharing and learning

Knowledge sharing can take many forms, from wiki or confluence pages to lunch-and-learn weekly meetups. Human based knowledge systems often have key go-to people that know where all the information lives. A good technology strategy should outline a plan to replicate these people, break down any information silos that exist and provide the tooling and mechanism that is necessary to enable the right information to be shared to the right people, and at the right time.

Questions about knowledge sharing

Do you foster knowledge sharing? How easy is it for people to learn from the people around them? Do your teams have the tools to easily add to the body of knowledge? Are they internally motivated to share with others? On the flip side, how accepting are your people to learn from others? Are they humble and vulnerable enough to be open to listening to others? Can they access the body of knowledge easily?

What to look for:

Look at incentives, recognition or promotions that encourage knowledge sharing behaviours.

Questions about access to the right tools

Do we foster knowledge sharing, and do people have the tools to both add to the body of knowledge and access and learn from it?

What to look for:

Look for systems that are in place that are easy to access, navigate and contribute back into.

### Employer brand - attract and retain talent

Employer brand refers to the reputation and perception of an organization as an employer. It represents the collective attributes, values, and culture that the organization exhibits to attract and retain talent. Employer brand allows you to attract and retain top talent, build trust and credibility and support business goals. It creates a positive image of the organization as an employer and helps in attracting the right people who align with the company's values and goals.

Questions about employee brand

How attractive is your employee brand? How easy is it to attract and hire talent suited to work in a complex / chaotic environment that is typical of today's digital businesses?

What to look for:

Look at retention data and exit interviews to see if there are recurring patterns

Activities: Retention data analysis • Exit interviews • Employer review website analysis (eg Glassdoor) • Employee Voice Survey

### Internal and back office systems

Technology leaders usually develop and support internal and back office systems. These systems directly contribute to the pleasure and satisfaction that workers derive from their job. Internal and back office systems or processes should offer employees a frictionless experience. But are your internal facing systems serving their purpose? We've all used back office systems that made us feel like we are running in treacle. Without focus on these systems for their delightful user experience, it's little wonder that CRM, ERM and timesheet systems feel like they are stuck in the 90s.

### Key Business Questions

- What systems are key to the running of your business?
- What is lacking from the systems that your teams use?

Lines of Inquiry

### Streamlining Business Processes

Technology can play an important role streamlining business processes, identifying and eliminating unnecessary steps or activities in a process in order to make it more efficient and effective. The goal of streamlining processes is to reduce waste, improve efficiency, and increase productivity.

Questions about software requests

How long does it take to grant requests for software, tools and access?

What to look for:

Shorten the turn around time of systems access requests so that employees can get back to value-added tasks

Activities: Value stream map • Service blueprints

Questions about waste in business processes

How many in-progress work items are waiting between steps in a business process? Do systems add to or detract from the optimum flow of the role being played? How many digital or human interactions are needed to complete each business process? How many different communication channels do employees use?

What to look for:

Measuring the size of work queues (inventory) between stages of workflow is a good way to quantify bottlenecks. Supporting many communication channels increases costs, and many channels adds confusion as people aren't sure which ones to use. Handoff points typically create friction

Activities: Value stream map • Service blueprints

Questions about proccess across different countries

How many business tasks are supported by different systems in different geographies?

What to look for:

Some systems need differences due to country-specific regulation and cultural differences. But many systems are developed locally while having substantially the same behavior as the rest of the world. Local business units often can't wait for features to be prioritized globally, but this results in many similar systems which are inefficient to sustain.

Activities: Value stream map • Service blueprints

### User Sentiment

For any product, it's important to understand how your customers view the offering. The same kind of customer analysis needs to be made on internal systems, so that we can accurately judge how to better improve our employee's ability to help the enterprise and to identify impediments so that we can remove them.

Questions about negagtive sentiment

How many questions, complaints or negative sentiment about systems in chat channels?

What to look for:

Chat comments can be an indicator for employee satisfaction with systems and processes. Sentiment analysis tools are often used for customer comments, they can also be used to assess internal systems.

Questions about support tickets

How many support tickets received related to systems?

What to look for:

Systems that get an inordinately high number of support tickets should be reviewed for enhancement or replacement

Questions about system by-passers

Percentage of people who use system/process vs something else?

What to look for:

If staff vote with their feet for external systems, then we should review whether the internal systems are worthwhile.

Questions about aging systems

How many systems are modern compared to legacy?

What to look for:

Old, tired systems not only wear down on employee sentiment towards your organisation but can also can harm productivity. In the same way that you want to focus on compelling customer experience, treat your employees as customers of your internal products, and provide them with the necessary modern tools to achieve their goals. It is important to treat these systems with the same care as you do customer facing products. That means treating your internal users as the customers of these systems and applying the same product thinking approach, customer research and customer service that you do to your products.

We're releasing this article in installments. The next, and final, installment will look at the last direction: Emerging Technologies and Market Trends.

To find out when we publish the next installment subscribe to the site's [RSS feed](https://martinfowler.com/feed.atom), Martin's [X (twitter) stream](https://twitter.com/martinfowler), or [Mastodon feed](https://toot.thoughtworks.com/@mfowler).

## Footnotes

1: see [Digital Transformation Is About Talent, Not Technology](https://hbr.org/2020/05/digital-transformation-is-about-talent-not-technology)

Significant Revisions

_22 August 2023: _Published culture and internal systems

_16 August 2023: _Published data-driven organization and reducing costs and risks

_15 August 2023: _Published improved efficiency and improved product quality

_10 August 2023: _Published Expand customer segments and inorganic growth

_08 August 2023: _Published first two strategic directions: expanding to complementary products and new markets.
