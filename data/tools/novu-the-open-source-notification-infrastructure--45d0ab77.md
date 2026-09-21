---
title: "Novu - The open-source notification infrastructure"
notion_id: 45d0ab77-a9ca-424c-9cc7-c54d21b5c280
notion_url: https://app.notion.com/p/Novu-The-open-source-notification-infrastructure-45d0ab77a9ca424c9cc7c54d21b5c280
last_edited: 2023-10-12T19:18:00.000Z
source_url: https://novu.co/
tags: ["English", "Programming", "Infrastructure", "Untried", "Tool", "Service"]
---
Simple components and APIs for managing all communication channels in one place: Email, SMS, Direct, and Push

npx novu init

## used by

## How it works?

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Create template

Select channels, add content with {{dynamic}} syntax, and custom rules to control the delivery of notifications.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Connect providers

Use a built in collection of popular providers - Sendgrid, Mailgun, Twilio and many more. Add API key and you're ready to go.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Add trigger

Send an event trigger using one of our community built SDK's, and we will handle it from there.

## We've built it so you don't have to

- Digest

A digest engine that aggregates multiple events in to a single precise notification.

- User preferences

Using Novu API to handle all user preferences and subscriptions across channels. UI components included.

- Priority management

A smart API to centralize all communication channels in a single place: E-mail, SMS, Direct, Push and many more...

- Monitoring

Debug deliverability and analyze sending patterns across multiple channels

- Content management

Manage content for all channels and in multiple languages without the need to redeploy your code

- Timezone awarenessComing soon...

Send transactional notifications based on user's timezone and working hours.

## Fully featured notification center in minutes

Build a real-time notification center using our embeddable components or connect your custom UI with our notification feed API.

[Read Docs](https://docs.novu.co/notification-center/getting-started)

## Built by the community

Open-source is in the heart of Novu. We keep all the source code and work publicly available. Join our community driven project with over 3,000+ developers from around the world who contribute code and help building the modern notification infrastructure.

### Join the community:

[27k](https://github.com/novuhq/novu)

## Component based

Novu API-first approach, means that you can use just what you need, when you need it.

[Get Started](https://web.novu.co/)

## Simple to use outgoing communication layer

[Show all services](https://github.com/novuhq/novu/tree/main/providers)

### [Twilio](https://github.com/novuhq/novu/tree/main/providers/twilio)

[SMS](https://github.com/novuhq/novu/tree/main/providers/twilio)

[**Slack**](https://github.com/novuhq/novu/tree/main/providers/slack)[
](https://github.com/novuhq/novu/tree/main/providers/slack)[Chat](https://github.com/novuhq/novu/tree/main/providers/slack)

[**Mailgun**](https://github.com/novuhq/novu/tree/main/providers/mailgun)[
](https://github.com/novuhq/novu/tree/main/providers/mailgun)[Email](https://github.com/novuhq/novu/tree/main/providers/mailgun)

[**Postmark**](https://github.com/novuhq/novu/tree/main/providers/postmark)[
](https://github.com/novuhq/novu/tree/main/providers/postmark)[Email](https://github.com/novuhq/novu/tree/main/providers/postmark)

[**Sendgrid**](https://github.com/novuhq/novu/tree/main/providers/sendgrid)[
](https://github.com/novuhq/novu/tree/main/providers/sendgrid)[Email](https://github.com/novuhq/novu/tree/main/providers/sendgrid)

[**MS Teams**](https://github.com/novuhq/novu/tree/main/providers/ms-teams)[
](https://github.com/novuhq/novu/tree/main/providers/ms-teams)[Chat](https://github.com/novuhq/novu/tree/main/providers/ms-teams)

[**SES**](https://github.com/novuhq/novu/tree/main/providers/ses)[
](https://github.com/novuhq/novu/tree/main/providers/ses)[Email](https://github.com/novuhq/novu/tree/main/providers/ses)

[**Plivo**](https://github.com/novuhq/novu/tree/main/providers/plivo)[
](https://github.com/novuhq/novu/tree/main/providers/plivo)[SMS](https://github.com/novuhq/novu/tree/main/providers/plivo)

[**SendinBlue**](https://github.com/novuhq/novu/tree/main/providers/sendinblue)[
](https://github.com/novuhq/novu/tree/main/providers/sendinblue)[Email](https://github.com/novuhq/novu/tree/main/providers/sendinblue)

[**Discord**](https://github.com/novuhq/novu/tree/main/providers/discord)[
](https://github.com/novuhq/novu/tree/main/providers/discord)[Chat](https://github.com/novuhq/novu/tree/main/providers/discord)

[**Mailjet**](https://github.com/novuhq/novu/tree/main/providers/mailjet)[
](https://github.com/novuhq/novu/tree/main/providers/mailjet)[Email](https://github.com/novuhq/novu/tree/main/providers/mailjet)

[**Mandrill**](https://github.com/novuhq/novu/tree/main/providers/mandrill)[
](https://github.com/novuhq/novu/tree/main/providers/mandrill)[Email](https://github.com/novuhq/novu/tree/main/providers/mandrill)

```plain text
1import { Novu } from '@novu/node';2
3const novu = new Novu(process.env.NOVU_API_KEY);4
5await novu.trigger('<WORKFLOW_TRIGGER_ID>',6  {7    to: {8      subscriberId: '<UNIQUE_SUBSCRIBER_IDENTIFIER>',9      email: 'john@doemail.com',10      firstName: 'John',11      lastName: 'Doe',12    },13    payload: {14      name: "Hello World",15      organization: {16        logo: 'https://happycorp.com/logo.png',17      },18    },19  }20);21
```

## An infrastructure that speaks your language

Community built server-side SDK's for your preferred programming language

[View SDKs](https://docs.novu.co/api/client-libraries)[Read Docs](https://docs.novu.co/overview/introduction)

### Loved by engineers from around the world

Explore tweets from engineers worldwide and see why they're fans of our company's innovations.

- [Todos for today: #ship a @GleapSDK update that utilizes @novuhq for amazing notifications 🎉](https://twitter.com/lukasboehler/status/1696793039841144916)[Lukas@lukasboehler](https://twitter.com/lukasboehler/status/1696793039841144916)
- [migrating to @novuhq be like...](https://twitter.com/psteinroe/status/1602958750847062017?s=20&t=GazBEYVRhI2ch6xP7Wqn5A)[Philipp Steinrötter@psteinroe](https://twitter.com/psteinroe/status/1602958750847062017?s=20&t=GazBEYVRhI2ch6xP7Wqn5A)
- [Thanks to a great tool called @novuhq, we can easily implement notifications into our upcoming v0.5.0 release.](https://twitter.com/doinfinehq/status/1671123804049874947)[Doinfine@doinfinehq](https://twitter.com/doinfinehq/status/1671123804049874947)
- [Novu make notification management much easier. They're doing a great job with the service they offer.](https://twitter.com/NikkiSiapno/status/1696509202993426884)[Nikki Siapno@NikkiSiapno](https://twitter.com/NikkiSiapno/status/1696509202993426884)
- [So excited about the rise of the notifications infrastructure space (+ open source 🔥)](https://twitter.com/rauchg/status/1557048605042565120)[Guillermo Rauch@rauchg](https://twitter.com/rauchg/status/1557048605042565120)
- [The best solution for notifications.](https://twitter.com/csaba_kissi/status/1696056864373416109)[Csaba Kissi@csaba_kissi](https://twitter.com/csaba_kissi/status/1696056864373416109)
- [Amazon Simple Notification Service: @novuhq](https://twitter.com/nathan_tarbert/status/1692654472952959300)[Nathan🔸Tarbert@nathan_tarbert](https://twitter.com/nathan_tarbert/status/1692654472952959300)
- [Finally, someone (Novu) made an open-source notification center](https://twitter.com/FGRibreau/status/1686631921797767168)[Francois-Guillaume Ribreau@FGRibreau](https://twitter.com/FGRibreau/status/1686631921797767168)

## Ready to send your first notification?

### Self-Hosted

[Read Docs](https://docs.novu.co/overview/docker-deploy)

[Get Started](https://web.novu.co/)
