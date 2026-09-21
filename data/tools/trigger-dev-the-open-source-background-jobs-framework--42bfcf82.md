---
title: "Trigger.dev - The open source background jobs framework"
notion_id: 42bfcf82-e0ca-4d32-9d1c-6ccd03047b45
notion_url: https://app.notion.com/p/Trigger-dev-The-open-source-background-jobs-framework-42bfcf82e0ca4d329d1c6ccd03047b45
last_edited: 2023-10-12T19:15:00.000Z
source_url: https://trigger.dev/
tags: ["Tool", "Service", "English", "Infrastructure", "DevOps", "SysAdmin", "Programming"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Supported frameworks

See it in action:

- subscription.ts

```plain text
1client.defineJob({2  id: "subscription-plan-changed",3  name: "Subscription plan changed",4  version: "0.1.1",5  integrations: { slack, resend },6  trigger: stripe.customer.subscriptionUpdated(),7run: async (payload, io, ctx) => {8const user = await db.users.find({ stripeId: payload.customer });10const planId = getNewPlanId(payload);12if (user.planId !== planId) {13await db.users.update(user.id, { planId });15await io.resend.sendEmail("Plan changed email", {16        to: user.email,17        from: "jane@acme.inc",18        subject: "Your plan has changed",19        html: planEmail(payload),22if (isPlanUpgraded(user.planId, planId)) {23await io.slack.postMessage("Notify team", {24          text: `Plan upgraded for ${user.email} to ${planId}`,25          channel: "subscriptions",30});
```

A Webhook example

When a customer's Stripe subscription updates,check if their plan has changed. If it has, update it in your database, send them an email, and notify the team of upgrades.

- tweet-generator.ts

```plain text
1client.defineJob({2  id: "tweet-generator",3  name: "Generate tweets from an idea",4  version: "0.1.4",5  integrations: { openai, twitter },6  trigger: eventTrigger({7    name: "tweet.idea",8    schema: z.object({9      idea: z.string(),10      scheduledTime: z.date(),13run: async (payload, io, ctx) => {14const generatedTweet = await io.openai.createCompletion("Generate ✨", {15      model: "text-davinci-003",16      prompt: tweetPrompt(payload.idea),19await io.wait("Wait 🕘", { date: payload.scheduledTime });21await io.twitter.tweet("Tweet 🐥", {22      text: generatedTweet,26// Send the "tweet.idea" event to trigger the tweet-generator job27await client.sendEvent("tweet.idea", {29  scheduledTime: new Date(),
```

An event example

When one of your users submits an idea, generate a tweet using OpenAI. Wait until the time they selected,then send the tweet using their Twitter authentication.

- `1client.defineJob({2 id: "weekly-user-activity-summary",3 name: "Weekly user activity summary",4 version: "0.1.4",5 integrations: { sendgrid },6 trigger: cronTrigger({9run: async (payload, io, ctx) => {10const users = await db.users.findMany({ summariesEnabled: true });12for (const user of users) {13await io.sendgrid.sendEmail(`Weekly summary for ${user.id}`, {14 to: user.email,15 from: "summary@acme.inc",16 subject: "Your weekly summary",17 html: weeklySummaryEmail(user),21await io.slack.postMessage("Notify team", {22 text: `Weekly summary sent to ${users.length} users`,`

A Scheduled example

Every Friday at 4pm UTC,get all users who have opted to receive a summary, then send them an email. Notify your team that they've all been sent.

## Developer-first features

### Trigger.dev is designed to seamlessly fit into your existing workflow.

## Full visibility of every Run

### View every Task in every Run so you can tell exactly what happened.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Step-by-step

Follow the flow of a Job from the Trigger to the Tasks.

### All the details

See the input and output data of every Task, including retries and detailed errors.

### Re-run Jobs

Quickly re-run a Job to check if you've fixed a bug.

## It all starts with a Trigger

### Trigger Jobs from a webhook, on a schedule or from a custom event.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
client.defineJob({  id: "subscription-plan-changed",  name: "Subscription plan changed",  version: "0.1.1",  integrations: { slack, resend },  trigger:
```

## All the integrations you need

### Easily integrate with hundreds of third-party APIs – including your own.

### Use our built-in integrations

Easily subscribe to webhooks and perform any actions you want.

### Use an existing Node.js SDK

Use `io.runTask`to make it resumable and appear in the dashboard

### OAuth or API key authentication

Use API keys (which never leave your server) or let us handle OAuth for you.

### Bring-your-own authentication

Supply your user’s auth credentials, using Clerk.com, Nango, or rolling your own with our custom auth resolvers.

### We are backed by some of the world’s best investors, founders & operators

Our mission

## The complete open source background jobs framework

### We’re building the most comprehensive and easy-to-use background jobs framework for developers.

| Feature | What it does | Status |
| --- | --- | --- |
| Integration kit | Official Trigger.dev integrations or build your own |  |
| Self-hosting | Host the platform yourself |  |
| Cloud | Just write code, no deployment required |  |
| Dashboard | View every Task in every Run |  |
| Serverless | Long-running Jobs on your serverless backend |  |
| React hooks | Easily update your UI with Job progress |  |
| [Background functions](https://github.com/triggerdotdev/trigger.dev/discussions/400) | Offload long or intense tasks to our infrastructure |  |
| [React frameworks](https://github.com/triggerdotdev/trigger.dev/discussions/411) | Support for Remix, Astro, RedwoodJS & more |  |
| [Long-running servers](https://github.com/triggerdotdev/trigger.dev/discussions/430) | Run Jobs on your long-running backend |  |
| [Polling Triggers](https://github.com/triggerdotdev/trigger.dev/discussions/418) | Subscribe to changes without webhooks |  |
| [Trigger.dev Connect](https://github.com/triggerdotdev/trigger.dev/discussions/441) | Use integrations signed in as your users |  |
| Vercel integration | Easy deploy and preview environment support |  |
| Streaming | Receive data from your Jobs in realtime |  |
| 100+ integrations | Comprehensive support for popular APIs |  |
| File IO | Create Tasks that have file outputs |  |

### We love the open source community. Get involved!

Contribute to Trigger.dev

## Loved by developers

## Frequently asked questions
