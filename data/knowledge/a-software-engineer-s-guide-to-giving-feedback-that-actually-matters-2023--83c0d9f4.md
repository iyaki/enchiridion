---
title: "A Software Engineer’s Guide to Giving Feedback that Actually Matters (2023)"
notion_id: 83c0d9f4-1eba-4746-bc35-1e3f7cbba8e8
notion_url: https://app.notion.com/p/A-Software-Engineer-s-Guide-to-Giving-Feedback-that-Actually-Matters-2023-83c0d9f41eba4746bc351e3f7cbba8e8
last_edited: 2023-01-27T17:21:00.000Z
source_url: https://shopify.engineering/software-engineers-guide-to-feedback-that-matters
tags: ["Article", "Shopify Engineering", "English", "Line/People/Team Management", "Communication", "Leadersheep"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

A friend of mine used to get into heated debates all the time. This was before Twitter—well before. Now when you talk to him you would never be able to tell. I asked him to share with me his secret for how he was able to overcome that urge. He told me that he asks himself three questions:

1. Does this need to be said?
2. Does this need to be said by _ me_?
3. Does this need to be said by me _ right now_?

If the answer isn’t a resounding yes for all three, he shuts up.

The idea is that if his feedback will not be heard productively, he simply avoids contributing to the vitriol that fills online spaces today.

Giving poignant feedback is a skill that is not necessarily easy to acquire; yet, it can be a powerful and important tool in a professional organization. Whether you’re providing feedback to a colleague on their work, or managing up and out, it’s important to make sure that feedback will be effective.

This is especially true for engineers working on a team, or a new manager. If you see something wrong in a codebase or architecture, and you don’t say something, you will suffer the consequences and have to deal with an unstable system, or bug fixes for years to come.

You’re also in the trenches with your colleagues. You need to [ trust](https://shopify.engineering/90-day-plan-building-trust-remote-team) them. You’ll spend late nights on deployments and long hours finishing important releases. If someone’s behavior makes it hard for you to work with them, you, the team, and the company will suffer for it.

If your feedback _ is _ productive it leads to growth, both of the individual, and the company.

If your feedback _ isn’t _ productive, you’re simply stroking your own ego.

An environment that runs on ego is not collaborative and people prioritize their personal wins over the company’s goals. This creates an unhealthy competitive tension. No one will grow, no one will learn.

It’s easy to fall into the trap of wanting the other person to know that you’re right and that they’re wrong. Especially when you have imposter syndrome, and are plagued with doubts over whether you rightfully belong in your position. It’s easy to want to show that you’re smart and qualified, and belong at your job.

In my experience as an engineer, and in previous roles as a lead, a manager, and a director, I’ve picked up tips for delivering feedback with intention. Throughout this post I’ll share concrete tips and examples that developers are likely to encounter during the pull request review. Then I will show how to leverage these same tools in more general interactions.

Feedback can be a tool that fosters growth and collaboration, or it can cultivate an environment of unhealthy competition. The question is: what sort of environment do you want to work in?

## The Pull Request Review

Pull request reviews are great. If you don’t do them where you currently work, you should. A pull request is a way to show what new code you are introducing. A pull request review is a process for getting feedback from your peer before deploying your new code; most systems even let you require one before code merges. When done well, they improve the quality of your codebase, prevent engineers from getting siloed, and are a great opportunity for mentoring. In [ today’s remote work environment](https://shopify.engineering/asynchronous-communication-shopify-engineering), pull request reviews create a positive touch point between colleagues.

How productive the review is, though, all depends on how effective the feedback is in your pull requests.

We are all professionals and should treat each other as such. So you might think there isn’t a need to worry about the feelings of the other person. However, even if you give good feedback, you’ll undermine its effectiveness if you ignore the fact that you’re speaking to someone else.

### Don’t Get Bogged Down in Nitpicks

If you’re trying to give productive feedback in your pull request review, a nitpick will ensure that it won’t be heard. A nitpick doesn’t improve the performance or logic of the code, like a style preference—brace usage or file formatting. While a tidy home makes a happy home, no one likes getting nagged for leaving out the oven mitts.

Bogging processes down over trivialities is called the [ Law of Triviality](https://en.wikipedia.org/wiki/Law_of_triviality), Parkinson’s Law, or more colloquially “bike-shedding.” The term came from Parkinson’s example of focusing on the materials for building the staff bikeshed while reviewing nuclear power plant designs, instead of the actual design of the plant. The design of the power plant is a far more complex, and important design to review. This is as true for your code as it is for the design of a nuclear power plant. Instead of focusing on nitpicks, focus on the critical aspects of the code.

Too many nitpicks is an indication that you have a tooling problem.

Keeping uniform style across a codebase does improve the quality of the coding experience. Reviewing—and maintaining—code that is a patchwork of styles will lead to missing important things.

But a pull request is not the time to address these issues. A good linter should be able to enforce these preferences. If pull requests are getting bogged down by nitpicks, revisit your linter, consider implementing it in a pre-commit hook. This tool was designed to enforce these opinionated styles, so you should use the linter instead.

### Raise the Bar for Your Entire Team

A pull request review is not the time to evaluate the skill level of the other engineers on your team. If you foster an environment where people feel judged and impugned they won’t grow, they won’t excel, and they won’t create excellent products. Remember what I said earlier: if your feedback isn’t productive, you’re simply stroking your own ego, without getting the results you want.

I’m not advocating letting poor quality code into your codebase to appease the feelings of your colleagues. I _ am _ advocating that you consider how you’re giving your feedback. Why bother giving feedback that will only fall on deaf ears?

When evaluating code it’s easy to lose sight of the person you’re working with. Writing and reading code requires an analytical frame of mind. Communicating productive feedback takes emotional intelligence. This context switching can be complicated.

So how do you walk that line between giving critical feedback, and making sure it lands?

### **1. Have a conversation as a team.**

Team processes should be reviewed every so often. If you are having trouble with the quality of your team’s code reviews, it’s likely that your team hasn’t discussed important processes in a while.

Context affects everything. If your team regularly discusses growth, and how to raise the bar for everyone, getting feedback is expected.

### **2. Frame the context of your feedback.**

If you’re doing a pull request review in real time consider prefacing the session with your goals, for example:

“We’re here today to help improve the overall quality of our codebase, help you and me grow in our craft, and prevent you from getting siloed by sharing context.”

### **3. Write your feedback twice, once to articulate what you want to say, then to craft it as effective communication.**

This is a great technique for all forms of feedback, whether comments on code, emails, or Slack messages. I typically use Notepad or a doc to write out longer messages so I can think them over. This way I don’t accidentally hit send before my message is ready.

**4. Reinforce the positive too.**

Code evaluation isn’t only about finding what’s wrong, but also finding what’s right. Point out where your colleague did a good job too.

### **5. Ask questions.**

Questions are a great way to level the playing field. If your colleague feels like they have something to contribute to the conversation, the pull request review becomes a give and take. This changes the dynamic from one of finding fault, to a two-way educational process.

The ultimate goal of pull request reviews is to lift up the codebase, and the team along with it through collaboration. Keeping this in mind while providing feedback will enhance the process for everyone.

### **6. Don’t be judge-y.**

You’re not here to evaluate your colleague’s skills. You may find areas where they need improvement. That’s ok. But think about how to relay criticism. I’ve recommended videos and articles when I saw that someone needed help in an area. I’ve also asked for resources when I found myself lacking.

### **7. Make it SMART.**

Just like SMART goals, you can frame your feedback in a way that is Specific, Measurable, Achievable, Realistic, and Timely. If your feedback isn’t specific, it’s hard for the other person to address the issue. Givean example. If it’s a general issue you’d like to address in the codebase overall, mention that it’s a larger conversation, and consider not holding up the entire merge request.

### **8. Separate the person from the behavior.**

When giving critical feedback, it’s easy for the recipient to feel attacked. If you separate the person from the behavior it helps the person accept the feedback more easily. They are not being attacked, it’s simply something that they should do differently. For example, instead of telling someone that they’re a sloppy developer, try saying: “When you don’t [ tophat](https://shopify.engineering/mobile-tophatting-at-shopify-1) thoroughly it leads to more bugs that your other teammates have to handle.”

## Managing Up and Out

You’d be surprised, but you can “manage” your boss and teammates. Every relationship requires a certain level of cultivation, or managing, to remain healthy. With a close friend, that can mean writing a thoughtful note from time to time. With a spouse, you might go on a date night.

In a work setting you should always cultivate relationships with your colleagues, your reports, your peers, and your supervisor.

When someone does something wrong in the workplace, it’s like a dead fish: the longer it stays hidden under the carpet, the more it smells.

Personally, I seek feedback. I like having regular one-on-one meetings with my colleagues. It helps me stay on track. I know what’s going on within my team, what my peers need from me, and what is top of mind for my boss right now.

Management consultant Peter Drucker defined management as “the organ of society specifically charged with making resources productive.” So regular [ communication](https://shopify.engineering/define-boundaries-prevent-burnout) is key for keeping everyone working together productively.

But when something is off—when that dead fish begins to smell—what do you do?

That’s when we ask ourselves my friend’s three questions:

1. **Does this need to be said?**Is this an issue for you, or is it affecting other people? If it’s only bothering you, maybe you don’t have to say something. Maybe you do.Does this need to be said by _ me?_
2. **Will what I say be heard, and be impactful?** Maybe this is something that is better handled by someone closer to the person, or by their manager. If it needs to be said by someone else, but not by you, consider having a conversation with their manager so that the person will get the feedback before it’s too late to improve.
3. ** Does this need to be said by me**_** right now?**_ Some times aren’t good for giving feedback. Is the person, or team, struggling to get an important project done? Are they dealing with something stressful at home? These aren’t excuses to behave badly, but they may affect the impact of your feedback.

Think about how you can provide feedback that will be heard, and be productive. Can you give a feedback sandwich? Compliment, critique, compliment. Finally, are you giving SMART feedback? Feedback should be framed like goals. They should be specific, measurable, achievable, realistic, and timely.

If you ask yourself these questions, and the answer is yes, you’re likely not speaking out just for your own ego. And if you think about who you’re speaking to and how they will receive what you’re saying, you’re likely to be impactful. And your feedback just might make a difference.
