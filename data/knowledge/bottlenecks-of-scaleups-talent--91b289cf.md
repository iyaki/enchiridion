---
title: "Bottlenecks of Scaleups: Talent"
notion_id: 91b289cf-e84d-4b3f-9ae1-0b18a09fef89
notion_url: https://app.notion.com/p/Bottlenecks-of-Scaleups-Talent-91b289cfe84d4b3f9ae10b18a09fef89
last_edited: 2023-04-25T13:16:00.000Z
source_url: https://martinfowler.com/articles/bottlenecks-of-scaleups/02-talent.html
tags: ["English", "Line/People/Team Management", "Project Management", "Article", "Martin Fowler"]
---
[https://martinfowler.com/articles/bottlenecks-of-scaleups/02-talent.html](https://martinfowler.com/articles/bottlenecks-of-scaleups/02-talent.html)

As startups find their product-market fit they grow rapidly, becoming a scaleup. But as they do so, they run into bottlenecks that slow their growth below its potential. We have noticed common bottlenecks and this is part of a series of articles describing them.

## How did you get into the bottleneck?

Forming a startup team starts with hiring from your personal network – your college friends, your cousin’s husband, your former roommates, and old colleagues. This works. Ideally, in the early stages of a company, you need a small, close-knit team that communicates effectively and has personally bought into the company’s goals. The initial experimentation phase will be a tough ride, so you need a totally committed team. The founder’s relationship to the team is what holds it together in the early stages. There will be difficult conversations and decisions that only a close-knit group can have: When to call it on the failing product idea? Which customer segment do we target? How do we find the next 10% of growth? Should we give up equity for funding?

A team like this can carry a company through initial funding and maybe even Series A and B. If the product is a success, it’s gaining traction, and you will quickly have to think about hiring. Its natural, and easiest, to approach the job market in search of candidates that look and behave like you. The new joiners, often senior people, can fit immediately into the initial culture and be productive out of the door. Again, this will work, but only up to a point. In these early days, your extended personal network may be strong enough to source enough people willing to give up a stable job and take on an amount of personal risk. It’s often the founder’s energy and values that convince them to join.

Like most other bottlenecks, the scaling problem happens when the product’s success moves the company into a hyper-growth phase. Invariably, by then the initial team has taken on a lot. The technical founders might still be coding and putting out fires. Product-oriented founders are approving small design changes and at the same time, trying to think about a broader strategy. Without enough resources, there isn’t a choice, but to stretch everyone. Ideally, stresses that come with the growth period are noticed early, and the team can hire before a crisis. More often than not however, indicators lag, and before anyone has had a chance to prioritize hiring, growth is bottlenecked by both capacity and capability.

When companies do expand through aggressive hiring, if it’s not handled carefully, it will cause many problems. It puts a lot of extra work onto the current team. The company is now at a scale where the leadership team can’t spend time interviewing and selling every candidate on the company. Just getting people in the door isn’t enough; you have to be able to incorporate new talent into the company. We frequently hear stories where new hires are not aligned with the founders and initial team, so the hiring investment doesn’t help with the bottleneck.

The story we’re telling is familiar to anyone that has tried to scale a startup. But, is there a better way? How do we recognize the signs that the talent bottleneck is coming? How do we set up our companies better to scale people? That is what we will attempt to answer in this article.

## Signs you are approaching a scaling bottleneck

When you are in the midst of scaling it can be difficult to notice that you are being constrained by talent, until you really feel it, and your business growth is affected. These are some of the signs you can look for.

Figure 1: Hiring process

### Frustrations from employees

Creating a startup is stressful. Working a lot of hours is expected. The strains business growth puts on people can go unnoticed. Concerns are often solved with “let’s just push for this next release, then we can slow down”. Of course – the slowing down never happens.

You need to look for signs of reaching capacity. Listening to people is critical and it requires intentionality. Establishing check-ins from managers and from a people department, as well as anonymous forms of feedback, can give good insights. Concerns about feeling overworked might not be bubbled up by managers reluctant to highlight delays, so it's key to foster a culture of transparency from the beginning – what did it really take to launch a product? Is that sustainable?

You may also notice this frustration surfacing by growing disagreement on deadline dates and ultimately by increased attrition.

### Stretching to hit deadlines, quality is slipping

Rushing to hit deadlines naturally means shortcuts will be taken. This will result in quality problems. They might be visible via user-facing bugs, outages, more customer service calls, delays or problematic releases. It might be covered up by developers firefighting or a customer service team appeasing customers, but this will soon become unsustainable.

There may be other problems with less apparent quality issues that would only surface over time, for example, code standards or testing coverage. We talk about a slippery slope of quality in our technical debt bottleneck. Managers can try to spot the internal problems by giving teams forums to explain the shortcuts they have taken, such as post mortems and retrospectives.

A clear indicator it’s time to add more help is when a team constantly feels the need to take shortcuts but doesn’t have a chance to go back and fix them. But you might only know this is happening if there’s a level of safety established where employees don’t fear repercussions for bringing weaknesses to light.

Pay close attention to the level of testing and automation. Developers might be having to do manual deploys and quality engineers doing repetitive manual testing. Common administrative tasks, if they're being done manually, e.g. with an ad-hoc spreadsheet, will run into problems with scaling.

### Key dependency on people

A growing startup can't afford to have a single point of failure, that if they lose will affect the business critically (the so-called bus factor). An engineer that owns a component that no one else works on. A promoted manager that is still in the weeds running teams. A specialist that builds the core algorithms. As the startup grows, the risk of losing critical members becomes more significant. Finding talent that reduces these dependencies is key.

We can spot this by examining how we’ve designed teams; auditing the owners and contributors to various systems.The in-demand person will likely keep coming up as a dependency on projects.

### Underperforming recruiting

The hiring team itself may be the bottleneck. We often see companies not hiring the amount of recruiters and hiring managers they need to find the right quantity of candidates. We need to look at the current throughput and targets to assess whether to expand the recruiting team. Your time to offer should be < 45 days and your time to start should be < 60 days.

Increasing the hiring team alone may not solve the problem. There may be blockers in the hiring funnel. For example, relying on the judgment of a few. This won't continue to work during a period of rapid growth. The hiring department will need the tools to be able to monitor the funnel data and spot these chokepoints.

### New employee’s expectations aren't being met

The new people who come in the door aren't happy, not producing the result you expected, and you aren’t maximizing their potential. Often a company looks to the new employees as if these failures are their fault - but typically the problem lies within how these team members have been welcomed and set them up for success. We’ve observed typical complaints:

- Leadership hasn’t made space for the new role, they’re either still trying to do it or are micromanaging the new employee.
- New ICs, not given the same level of ownership and access, aren't treated the same way as tenured employees.
- Environment is difficult to get started in, not having tools and the knowledge they need to be productive. E.g. technical documentation about APIs and libraries, or access to infra to be able to run services, or access to design tools.

To identify problems we need to listen closely to the new employees after they join.

### Underestimating growth

Commonly, leaders underestimate how long it can take to hire and build a good team. A startup might see the indicators of growth, but be skittish to commit to hiring in order to reduce their run-rate or extend their runway. At the same time, being able to double down on successes before competitors catch up is critical.

For a quickly growing company, even after hiring, it can take an additional 2-5 months for new hires to be productive, depending on complexity of the domain and the technical platform. All this means it’s necessary to proactively plan for hiring before it becomes a bottleneck.

One way to predict hiring needs is to have a solid platform to monitor the product and business indicators associated with growth. Use trends to hire, rather than simply hire in reaction to obvious problems. Plan well in advance for new product initiatives, and remember to factor in a level of attrition.

## How do you get out of the bottleneck?

Covering a good hiring strategy for startups would require a whole book. These are lessons from our digital scaleup teams that are competing for the best talent.

### Use your technology and innovation as a hiring differentiator

As the company gains traction and becomes more well-known, hiring becomes easier. Currently, the market is highly competitive for technologists. You can attract people based on the impact of your product, the projected success, or the personal interest someone has in your field. This may not be enough, we also recommend making your technology and innovation the things that set you apart as a company.

What interests a technologist is different depending on the candidate, but we find the impact of the work, innovation in the technology and the effectiveness of the technology environment, are big draws for candidates. A scaleup can offer these so it’s important to tell that story to candidates.

### The impact of the work

At a large company a technologist may spend a lot of time on a product or a featureset that never sees the light of day, which is demotivating. The appeal of a startup is that they will work on something that matters, their work will have material differences to the company's success, and their incentives and career will be linked to that. Candidates will have more impact, and the downside – more pressure.

### Innovation in the technology

The ecosystem, stack and tools matter a lot to candidates because this dictates what they will be interacting with every day. The choice you make should weigh both what appeals to candidates and what is dependable to build your product. An older technology will be off putting to candidates. However, a new and shiny technology may be risky because the talent pool is small, and the technology may not be stable.

Often there’s a desire to pick the latest and greatest niche tech, sold under the justification that it will attract top-quality candidates. The company can train candidates in the new technology, but in reality we have seen they may not be willing to learn, for fear it will limit their future job prospects

Innovation is also leveraged in the product itself; through using an emerging technology like Machine learning or Virtual Reality, or because the product design and implementation itself is innovative or unique. This can be compelling for candidates, and should be part of the hiring messaging.

### Effective environments

Technologists want to be effective, and they want to succeed at the job they have been given. This is not measured by lines of code, it’s creating useful software. The reality is a lot of working environments are full of bureaucracy, friction and needless red tape, which leads to less high-quality working software. Chances are candidates have experienced that in previous jobs.

The advantage of a well-run startup is that it will have little baggage and be comparatively effective at software delivery. This message should come through to the candidates. We can do this by talking about the company structure, how products are produced, how people communicate and collaborate.

### Promoting the technology

The job description and the initial interactions with a candidate can go a long way. Instead of just describing the experience and technical skills you want a candidate to have, we find that focusing more on attributes can bring in candidates who are a good fit. Depicting how the company will help them succeed, and including what it's like to actually work there, what a day in the life is like, helps candidates decide if they can see themselves in the role/company. Having candidates talk to an employee will have more weight than a conversation with a recruiter.

Companies that have an open culture on their technology innovation will have an easier time hiring, and we can do this by empowering employees to:

- Talking at conferences about the innovations at their company. Examples: Uber talking about scaling lessons, and Netflix talks about microservices.
- Maintaining a high quality engineering blog. Examples: Etsy's code as craft and Airbnb's tech blog.
- Contributing to open-source projects. Examples: Hashicorp and Thoughtworks.

### Hire more T-shaped technologists than specialists

Another difficult balance to get right is between hiring experts who know a specific tech stack or business domain well and candidates who don’t have the exact experience you need but can learn.

This balance likely changes as you progress. In the beginning, you need a few specialists, who can set patterns for the rest of the team to follow – an infra SME, a seasoned developer who has built a similar scalable architecture, or a data scientist who has worked in the domain of your product. The rest of the group should have relevant experience, but we'd recommend you prioritize flexibility, bias for action, and ability to learn — your archetypal T-Shaped candidate.

Later on in the hyper-growth and optimization phases, there's going to be more room for specialization. There will likely be whole teams that are focused on a single capability, such as observability, front end tech, or data science. However, we often see companies trying to fill too narrow of a gap, which can lead to losing great candidates or taking a long time to find that special person.

A candidate’s deep expertise doesn’t give them a pass on company values. They should go through the same process and hit the same bar e.g. soft skills, like communication and listening skills.

### Utilize Non-Senior Developers

It makes sense to have a small senior (10+ year experience) team in the early stages. However, if the startup continues to hire senior employees as they grow this will quickly become a bottleneck. There is a limited amount of talent in the marketplace, and the demand is very high. They’re also expensive. We recommend that startups alter the balance and include more non-senior talent (2-6 years) in teams.

To hire for non-senior talent, we have to be more flexible on the level of experience and technical skills. Ideally, we want to hire someone that can learn and pick up skills quickly. This requires changes in the interview and sourcing process. We can’t just match against a number of keywords.

To embrace less senior technologists there will have to be a culture shift. A typical anti-pattern we've seen is relying on “hero developers” to do the majority of the work; Senior and tenured, they have written a lot of the core systems and can trouble fix easily. The issue is that they often don’t take the time to bring others up and support the team. Of course, we always need developers that can do heavy lifting, but we find effective teams sacrifice a bit of that individual productivity to increase the productivity of the team.

### Embrace remote working

It’s challenging to make precise recommendations about working habits, as this is evolving. One thing we can say is that the scaleup companies we work with are all embracing remote working. They do this by

- Providing quality remote collaboration tools like video conferencing, long lived group chat rooms, whiteboarding etc.
- Budget to set up a home office environment e.g. ergonomic chair, camera and monitor.
- Reducing the amount of video meetings; a lot of calls is energy-sapping.
- Changing the rituals and practices to better support remote. E.g. making sure in-person and remote groups are on equal footing.

How does this help with hiring? Because creating an efficient remote culture allows startups to leverage wider regional and global talent pools. We’ve seen companies try to embrace remote working while skipping the above steps, but it’s caused a lot of friction for employees, so we’d advise fully committing to remote capabilities once you decide to go this route.

### Example initiatives as you grow

Phase 1

Experimenting

Small founder team hired from personal network

Phase 2

Getting Traction

Referrals from extended network and investors

Create hiring value proposition from product mission

Leverage technology and innovation story to differentiate

Establish mindful and welcoming culture intentionally

Phase 3

(Hyper) Growth

Bring in an experienced hiring leader

Based on projected growth, build hiring team to match capacity

Ensure clear messaging on mission, goals and culture

Sourcing beyond referrals, identify talents pools considering diversity goals

Include hiring in everyone's job responsibilities

Phase 4

Optimizing

Consider expansion to capture talent - global, regional

Invest hiring process for scale; improve consistency, remove friction

Optimize onboarding; time to effectiveness across org

Augment sourcing with AI matching tools

### Invest in the hiring process

### Scaling the hiring team

The phased-approach in the initiative diagram demonstrates how to grow the hiring team incrementally. Expanding as a company moves from experimenting to optimizing phases. It's important to plan early. As a rule of thumb, a recruiter can manage 2-3 hires per month. If you want to grow your team by 36 people in a year, you will need at least one recruiter. Supported by the right tools, administrative support and efficient process.

In addition, we’ve found that for every three recruiters, you should hire a recruiter operations person for interview scheduling and accompanying administrative tasks. This is often missed.

### Streamline the process, practice continuous improvement

To create the best experience for candidates and the most efficient process, our scaleup teams use a lean technique to optimize, similar to the way we optimize other business processes. Using a cross-functional group we map the process, making sure we’re hitting the outcome for all stakeholders ( sourcers, recruiters, managers, interviewers, candidates hired or not). We can then do more detailed research to find the friction and create steps to remove.

This should be an data-driven approach; these are typical data points that are useful:

- Diversity definition and goals – Your baseline metrics for diversity should look like the census data of your office locations and you should strive for incremental improvements every year.
- Success of the hire – adjust the hiring process and job requirements based on feedback from managers.
- Analysis of friction / touchpoints - where can we remove steps to speed up and improve efficacy
- Candidate feedback on recruiting experience - both hired and non-hired
- Market and competitor data - to make evidence-based decisions on comp and benefits. Ensure job titles and descriptions are attractive.
- Interviewer availability and effectiveness – employees are motivated and have enough time to do the required tasks e.g. read resume, preparation, write notes.
- Funnel / Conversion rates - over time to be able to monitor and improve downward trends.

Hiring datasets can be small, often messy, with lots of nuances. We’ve seen small data used to support inaccurate hypotheses. To draw conclusions you should apply statistical techniques, and research techniques to decipher qualitative information. Including a data analyst in your team can help with the research.

Collecting feedback from both candidates and internal participants enables us to continually improve the hiring process. An anonymous survey can be used to capture metrics, like the Net Promoter Score, or responses that are freeform verbatim comments. These inputs can reveal what stages of our process need improvement. An ongoing analysis and calibration will contribute to a high-quality interview experience.

Figure 2: Net promoter score dashboard

Examples of a streamlined hiring process:

Figure 3: Hiring process

### Recruiting and business partnership for planning

We often encounter wildly optimistic hiring plans that have no hope in the realities of today’s market. The recruiting team and business leaders have to work together to make a plan that is reasonable and is able to keep to quality standards. Important guidelines:

Length of time to hire - The length of time needed to hire is difficult to estimate, especially for exec or specialist roles, but hiring teams must try to give the best accuracy they can. It will likely be a range that will also change over time, as the market changes and the company’s profile changes.

Constraints - The hiring team should be transparent about market trends and challenges. While it’s tough to communicate, they’ll be the first ones to see if the company brand is not attractive, why they’re continually losing out to other companies, or if the hiring team doesn't have the capability to find a certain skillset. These constraints are beyond control of the hiring team and will require help from the rest of the company to improve.

### Telling the story

As a company grows and starts to add people, it can no longer rely on the founder to work directly with new employees. There are more people sharing the company mission, goals and ways of working, so conveying a consistent message from the recruiters to HR, to leaders, and even peers - becomes an even bigger challenge. You want to effectively make sure newcomers hear the same messages regardless of who they’re talking to. The message has to be consistent, authentic and clear even if the mission and goals may have changed over time. In the early stages, when goals are especially fluid, it’s prudent to re-examine before a big hiring push.

### Tooling

There’s a growing hire-tech industry making great third party tools that we can utilize so that we don’t have to create systems or complex spreadsheets. Examples of a modern recruiting toolset:

- Applicant Tracking – clear status of the funnel and who is responsible for progressing to the next step. Examples are Greenhouse, ICMS, Avature. An agile-like dashboard is also great to improve transparency.
- Workforce planning – the ability to see the current talent and is what is needed in future, this might start as a spreadsheet, before moving to a more sophisticated tool.
- Sourcing automation – there’s a lot of current innovation using AI to match candidates to capabilities and job descriptions. We recommend exploring them when you have research in later scale phases.
- Referrals management – tracking who is referring, quantity, the thank you recognition and amount.
- Automation - information should flow between systems without the need to open up tickets, a lot of tools automatically integrate to other hiring tools, if not they should be easily scriptable.

A combination of these tools can assist in laying the foundation for clean records for both your candidates and your employees. Select one that is customizable and scalable to grow with your business. Many of these options offer per-seat licenses to suit your budget.

### Everyone has to prioritize recruiting

Unlike our systems, we typically run our product teams at full capacity, if not overcapacity. A hiring push means even more work for everyone: more interviews, sourcing, and hiring decision meetings. Your team might already be frustrated with the pace, so adding more responsibilities is difficult to accept. There’s no way around it – if you are going to maintain the quality, culture and ultimately accept the new joiners in their teams, your employees have to be involved. Hiring shouldn’t be outsourced.

A typical scaling problem is not federating hiring decisions. A small number of people become a bottleneck. It can be uncomfortable for founders to lose some control, but if they have brought new leaders, they should have trust in their hiring decisions.

To make time for hiring, things have to slow down. Build margin for people to be taken away for interviewing– and not just the interview itself; the prep, writing notes, context switching. Managers need to begin planning even earlier than interviewers if there are open roles on a team, so that they can consider how to recruit for those open spots.

Succession planning helps. When a business is growing, it opens up new opportunities for employees. It makes sense to move our top performers into new roles that stretch and challenge them. Their previous teams will need replacements. It’s a good practice for a manager to always know who might replace them and their team leads. Doing this will give the hiring team time to find a candidate before it’s too critical.

### Candidates are interviewing you

A good candidate is interviewing for culture, just as much as the hiring company. Candidates will choose a company where they had a positive interview experience, over one with better salaries and benefits where they had a less favorable experience. Despite offering better salaries and benefits, candidates will often choose companies where they had a positive interview experience. It is quite easy for an interviewer to let their ego or self-centeredness get the better of them and create a very uncomfortable experience for the candidate. A candidate appreciates genuine interest in them and their unique background, not just whether they fit well into a predefined job description.

We strongly recommend cogent and consistent interviewer training. It will provide the framework for knowing what interviewers can and cannot say from a legal point of view, and it will enforce the guardrails of what good looks like. Interview training also is a reminder for interviews to do preparation; read the job description and review the candidate's resume.

Include unconscious bias and awareness training to reinforce that assessments need to be based on capabilities and attributes, and not grounded in a cultural fit for the organization. Interviewer training reminds employees that they are empowered to help recruit their future coworkers. And it may serve as a retention tool to remind people why they are working on the team.

### Finding internal talent

While the company is small, it’s easy for leaders to know the capability of every employee and direct top performers into new opportunities. When you get beyond 50 people, an internal talent program guards against the danger of under-utilizing great people in your team, or creating a culture where only people in the inner circle get promoted. This is started quite easily by using the existing recruiting team. Post job titles internally and interview using a lightweight version of the external hiring process. The difficulty with internal candidates is removing biases from leaders who have only seen someone operating in their current role.

### Diversity won’t just happen

Diversity won’t just happen. It needs intention, planning, and effort. To find people from non-traditional talent pools requires more recruiting steps and time. In our research, many leaders expressed that they’d have started creating a diverse workforce in the early stages. Entering the hyper-growth phase with the need to scaleup capacity, diversity goals can easily be put aside. Before you know it you have a homogeneous workforce, that is difficult to change.

Some of the deliberate things our scaleups do to consider Diversity, Equity and Inclusion (DEI) while hiring:

- Start with the recruiting team; the recruiting workforce itself should reflect the company’s diversity goals.
- Intentional sourcing; e.g. underrepresented minority tech communities, coding bootcamps, geographic focus outside major technology hubs.
- Language in job posting; Go beyond experience and tech skills in job postings; focus on attributes that would make a good match
- Expect evidence from interviewers; vague comments such as “not fitting into the culture” can hide bias.
- Clear diversity definition; your company should have a clear definition and be transparent on targets and initiatives.
- Careful with referrals; if you rely too much on referrals, there is a risk of creating a workforce from the same background, referrals should be at most 30-40% after early growth stages.

### How Thoughtworks grew its talent

While Thoughtworks is a software development consultancy rather than a product company, there are a lot of transferable lessons. Over the last 10 years, Thoughtworks has grown from 1,000 to 10,000 people; not hypergrowth, but it represents significant growth and that put a lot of strain on the business.

### Identity

An important principle was to grow at a sustainable pace, keeping the cultural ideals that were core to Thoughtworks, but also recognizing it would change, that there are differences working at the increased scale. Thoughtworks set out to re-examine their mission by looking at “why does Thoughtworks exist”? This was conducted as a research project involving input from every consultant worldwide. The results were:

- Be an awesome partner for clients and their ambitious missions
- Revolutionize the technology industry.
- Amplify positive social change and advocate for an equitable tech future.
- Foster a vibrant community of diverse and passionate technologists.
- Achieve enduring commercial success and sustained growth.

Using this mission statement, as we scale, we could assess every decision, whether it would help to further the mission.

Thoughtworks also wanted to protect the cultural values that are to us critical to our success. Our values were important to share externally and use within the hiring process as we scaled. They are – Global first, Courageous, Inclusivity, Cultivation, Integrity, Curiosity, Pursuit of Excellence and Autonomous Teams.

### Scaling the hiring team

A case study of the Thoughtworks journey in North America is a good example. Four years ago, the North American Thoughtworks recruiting team was 12 people. That team, structured fairly inconsistently, could barely hire 10 people a month. Fast forward to the present. That team is now 25 people, and can consistently hire 75+ people a quarter. Here are some of the foundational frameworks we implemented.

For every 3 recruiters, we brought on one coordinator; someone to be responsible for the scheduling, travel arrangements, and administrative paperwork for the candidates. This one “pod” of recruiters would be capable of hiring 20-25 hires per quarter or ~100+ hires per year. More senior roles will require more effort and equal 1-2 hires per month for the same work. Be sure you build this deviation into your capacity model.

We also implemented a process we call Joy of Interviewing. Using a set process, we systematically reviewed every role and associated assessment process. We organized our candidate stages, standardized our questions and attributes, and effectively created a repeatable process that helped define our talent bar. With this foundation, we created a system that was easy to scale and replicate from country to country. It also guaranteed that a Senior Developer in the USA would be the same skill level as someone in Munich or London. We review our assessment tools regularly to ensure what we are hiring is still in line with what is required in the marketplace.

Finally, we knew you can’t improve what you don’t measure…so we set out on the journey of capturing meaningful data, and displaying it in a format that made sense. We hired the first ever Talent Data Analyst to help extract and visualize the numbers so we could measure our success by conversion rates and days to offer, to name just two. Using a combination of existing tools, and adding a couple of additional visualization products, we’ve been able to craft dashboards that are easy to read and understand. There is a level of rigor that is required by the associated Talent Teams to ensure that there is a high level of accuracy in order to use this to predictively model and forecast but the effort is worth the end result. With every year, the data we collect will make our hiring estimates closer to reality.

### Thoughtworks University

Like a lot of companies, at a certain scale we created a program to support junior technologists. It has existed since 2005, and has been the key capability to our ability to grow. It is not purely a graduate program, it is designed for anyone that is inexperienced in software development, career changers are very common.

What makes it unique, is that it is run by practitioners, experienced managers and tech leads who will take 3-6 months away to dedicate to teaching the skills and practices they have learnt. The course is 8 weeks, they work in teams to build and deploy a product, working with in a simulated client environment.

It is designed to mimic situations they will experience when working for TW, so they can be fully productive when they hit the ground. Graduates of the program talk about the benefits of the immersion into agile practices and the relationships that they build during that time. Many of our current leaders and managers came from Thoughtworks University.

### Diversity, Equity, and Inclusion as a core mission

As Thoughtworks grew, the goal was to be a company that is equitable, reflective and inclusive of the societies we live in. We aim to include all of society, both in our community and through our tech, by providing talent with a place to belong.

Thoughtworks believes diversity, equity and inclusion have the power to create social change and also to make better software products. By incorporating the perspectives of those from a variety of identities, backgrounds, and lived experiences, we’re better enabled to solve for the needs of the customer/user. The Organisation for Economic Co-operation and Development (OCED) in a study on scaleups entitled “Understanding Firm Growth; Helping SMEs Scale Up” said that “Gender and ethnic diversity are associated with better firm performance in growth-oriented firms...”

### Finding and nurturing diverse talent

One foundational aspect involved redefining what it meant to be a “technologist” at Thoughtworks. Although the definition of “technologist” varies across the tech industry, most definitions center on those in technical or engineering specialists roles, which can often be dominated by those who identify as cis-men. We recognize that not all technologists are engineers, they are anyone who actively participates in the creation of software

We aim to attract talent from non-traditional backgrounds, by not requiring degrees, welcoming career changers. We also partner with community programs to help us increase the representation of women and underrepresented gender minorities, LGBTQIA+, and BIPOC technologists.

### Cultivating and retaining diverse talent

Beyond hiring diverse talent, it’s essential to provide an environment where technologists of differing identities and backgrounds can thrive.

In collaboration with employee resource groups (ERGs), we design initiatives to promote intersectional awareness-building, inclusion training and education, and campaigns that represent the stories and experiences of our diverse talent.

Initiatives include Women in Leadership Development (WiLD), intentional executive sponsorship and employee-led DEI benefits, policy, and reward & recognition working groups.

We further demonstrate this through our metrics, welcoming feedback and engagement from our talent to influence how we improve. At the time of writing, 40.6% of all employees are WUGM (women and underrepresented gender minorities), 38.2% in tech, 62.4% in non-tech and 60% of Executive officers are WUGM. At Thoughtworks University 49% of graduates were women and under-represented gender minorities.

## Summary

We have talked a lot about hiring and attracting talent, just as important is preparing your organization and platform for the employee growth. If a company is going to grow by 2X or 3X, it is impossible to do that effectively without significant change to how the company is structured and operating. Future bottleneck articles including “Organization structure doesn’t support current size” and “Ineffective onboarding process” will describe how to do this.

- Talk about your engineering culture and innovation to attract talent.
- Hire a few specialists; mostly hire adaptable T-Shaped employees.
- Embrace remote practises. Consider Global and regional talent pools as you expand.
- Create a welcoming, mindful and effective environment to integrate and maximise new hire’s abilities.
- Plan hiring initiatives early; building the hiring team, sourcing candidates and interviewing takes time.
- Use data to optimize the hiring process and plans. Incorporate feedback into the hiring process to continually improve it.
- Make interviewing a priority, and give employees training and tools to interview effectively. Design the hiring process to remove bias.
- Take steps to increase the diversity of background, race and thought. Intentionally target non-traditional sources of talent. After the early stages limit the number of referrals.
