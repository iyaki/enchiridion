---
title: "Empowering engineering teams to perform code reviews"
notion_id: 7c8c063f-bbf3-488c-839a-34f55fc15ec3
notion_url: https://app.notion.com/p/Empowering-engineering-teams-to-perform-code-reviews-7c8c063fbbf3488c839a34f55fc15ec3
last_edited: 2023-03-30T13:37:00.000Z
source_url: https://leaddev.com/code-reviews-docs/empowering-engineering-teams-perform-code-reviews
tags: ["English", "Programming", "Article", "LeadDev"]
---
You don't have time to review every pull request. Here's how to empower developers to perform code reviews instead.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

So, what should you do instead? The answer is to empower your team to perform code reviews effectively. Here are seven quick tips for doing just that.

### 1. Document code review expectations

Documentation is essential for establishing standards and sharing knowledge around code reviews. Team members should be aware of the expected size, scope and structure for each pull request.

Detail pull request _author_ expectations, for example:

- The expected context and details
- When to refactor – in a separate PR, or within the same one?
- How to drive disagreements to resolution

Detail pull request_ reviewer _expectations, for example:

- Get to them fast
- Be kind – give a _reason_ for every comment
- Favor pragmatism, not perfectionism

### 2. Teach your team to review effectively

Beyond documentation, make time in person or virtually to teach the team some code review best practices. Emphasize the importance of kindness and clarity in code review comments. Make sure they understand when blocking is or isn’t justified.

Examples where blocking is_ justified:_

- Flaws
- Risks
- Over engineering

Examples where blocking is _not justified _(in these situations, they should comment and approve):

- Style
- Nitpicks
- Urgent fixes

### 3. Integrate automated review tools

Use linters and formatters to ensure consistent style. Ideally, your teammates wouldn’t discuss style at all during code reviews.

Set up automated static code analysis tools. These will help catch obvious flaws, and identify best practice opportunities.

### 4. Establish paradigms in your codebase

Introduce design patterns and structure to the codebase, that others can leverage and build on top of.

Document interfaces and modules, and how they’ll interact. Designs are always evolving, so keep this documentation up to date.

### 5. Enhance your delivery systems outside of code review

Code reviews are critical in preventing and detecting defects. But since you can’t perform every review, it’s a good idea to strengthen systems that contribute towards the same goal.

Do the work to strengthen your release pipelines. Prevent defects with robust, reliable, fast tests. The sooner you can test, the better. Add automated checks that run tests before developers can merge pull requests.

Similarly, you can detect defects faster with monitoring and alarms. You can mitigate defects faster with automatic rollback of changes with negative impact.

### 6. Be aware of what’s being shipped to production

Know when to scan a pull request, and when to do a thorough dive. For low-risk changes, let your teammates handle it. For high-risk changes, step in when necessary to prevent negative impact.

Integrating Slack and GitHub works well for this, specifically for pull request events and comments. You can read through a Slack channel intermittently, for awareness.

### 7. Give your peers some space to learn from mistakes

You can’t prevent everything. Software will never be perfect, and occasional defects will slip through. Treat these events as learning experiences and teachable moments. Help your peers grow.

Remember, you can’t write and review all the code for your team. If you could, hiring others would be pointless. Instead, put your team in position to ship better software, faster.
