---
title: "EmailEngine - Email REST API"
notion_id: 7bea49ca-74f7-4884-b1db-4bbc3025163f
notion_url: https://app.notion.com/p/EmailEngine-Email-REST-API-7bea49ca74f74884b1db4bbc3025163f
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://emailengine.app/
tags: ["Programming", "Email", "Untried", "Service", "English"]
---
Effortless email integration for your app or service. Simplify IMAP and SMTP complexities, and focus on what truly matters. Build better features, faster.

Self-hosted solution for developers and businesses.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Quickstart

EmailEngine is a self-hosted email automation platform that offers easy access to email accounts through a user-friendly HTTP REST API for receiving and sending emails. The platform actively monitors these accounts and sends webhook notifications for any updates.

Here's how to get started:

1. Check out the installation instructions for your system [here](https://emailengine.app/set-up).
2. Launch EmailEngine from the command line:

```plain text
$ emailengine --dbs.redis="redis://127.0.0.1:6379"
```

1. Open [http://127.0.0.1:3000](http://127.0.0.1:3000/) in your browser and activate a free 14-day trial.
2. Generate an access token and head over to the API Reference page to test out some API calls.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Easy to use REST API

The REST API with JSON payloads you expect it to be. No need to know IMAP internals, weird encodings, or juggle with different IMAP extensions.

Unicode stringsBinary attachmentsPaged listings

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Who can benefit from EmailEngine?

### SaaS CRM services

Integrate your users' email accounts with your platform to track email correspondence activity between users and their contacts.

### Web agencies

Create innovative features for your clients, such as customized newsletter platforms.

### Email hosting providers

Develop a tailored webmail interface on top of your email hosting offerings.

### SMB companies

Monitor and automate your support and info mailboxes.

### Email warmup services

Automate the warmup process between different email accounts to enhance email deliverability.

### Web-hosting providers

Supervise and automate special email accounts like postmaster or abuse mailboxes.

### Enterprise companies

Leverage EmailEngine as an IMAP or SMTP proxy between your legacy applications and MS365 OAuth accounts.

### Cold outreach services

Automate email sending and replying to prospects' messages.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Almost Perfect Deliverability

EmailEngine utilizes the user's own email servers to [send email](https://docs.emailengine.app/using-as-a-transactional-email-service/), giving the appearance to the recipient that the email was sent directly from the user, rather than through a third-party service, resulting in near-perfect deliverability.

Uploads to Sent MailAdds reference headersBounced email detectionStored MJML templates

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

I have been using emailengine for well over a year now and have found it to be simply *tremendous* and I cannot recommend it enough. Why? 1) The speed AND ease of use in interacting with IMAP accounts - from very different providers: standard IMAP, Google, MS Office 365, etc. 2) The simplicity and power of the REST API. The high quality Documentation, Support and Tools to match. 3) The ability to receive notifications and send emails from web apps (where sent emails are also auto created in the users sent items mailbox). 4) The richness and ease of the Webhook implementation, giving real time feedback to users. 5) Many advanced features including Document store support for high speed offline searching, support for Replies & Forwards, Mailmerge, OAuth2, tracking Bounces among others. Lastly, it isn't an exaggeration to say that choosing emailengine was one of - if not - *the best* decisions I have ever made in deciding which third-party software to use. Tom Conlon GO-1 SOFTWARE

Made with  Senja

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Convenient Hosted Authentication

Streamline the authentication process for your users with EmailEngine's [hosted authentication](https://emailengine.app/hosted-authentication) form. Once the email account is authorized, the user will be redirected back to your app.

While you can still register accounts through [API calls](https://api.emailengine.app/#operation/postV1Account), hosted authentication provides a more user-friendly and easier setup option during development.

IMAP/SMTPGmailOutlook

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Effortless Custom Integrations

Easily create [custom webhook](https://docs.emailengine.app/low-code-integrations/) routes using EmailEngine's low-code integration capabilities. Define filter functions and output mappers to transform events into chat messages in Slack or Discord, or send the data to Zapier for additional processing.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Monitoring

EmailEngine exposes a Prometheus metrics collection endpoint for easier [monitoring](https://emailengine.app/monitoring).

GrafanaPrometheus

## FAQ

### Who is EmailEngine intended for?

EmailEngine is ideal for anyone looking to integrate their system with existing email accounts. For instance, a help desk service that operates through the [support@domain.tld](mailto:support@domain.tld) account and converts incoming emails into support tickets.

### How does EmailEngine function?

EmailEngine maintains an open IMAP connection to each registered email account and continuously listens for any changes. When you make a request through the API, EmailEngine converts it into the corresponding IMAP command to execute the desired action.

### Is a license key required?

You can start a free 14-day trial directly from the dashboard without the need for a license key. Once the trial license expires, EmailEngine will stop processing IMAP accounts until a valid commercial license is obtained.

### Do I need to purchase multiple license keys?

Typically, no. You would pay for a [yearly subscription](https://postalsys.com/plans) to Postal Systems, not for multiple license keys for EmailEngine software. If your subscription is active, you can generate as many license keys as needed at no extra cost.

### What are the system requirements for EmailEngine?

EmailEngine requires [Redis](https://redis.io/) as its caching database. As long as you have a reasonably recent version of Redis installed, you should be good to go. For optimal performance, it's recommended to keep the latency between Redis and EmailEngine as low as possible.

### What protocols does EmailEngine support?

EmailEngine utilizes IMAP to access email accounts and SMTP to send emails. Other protocols such as POP3, ActiveSync, or the Exchange Web Services (EWS) SOAP API are not currently supported, but they may be added in the future.

### What are the data compliance considerations with EmailEngine?

As a self-hosted solution, EmailEngine does not send or store any data outside of your network. It only retains [minimal metadata](https://docs.emailengine.app/data-compliance/) necessary for syncing and caching, but not the actual contents of the emails. This ensures that all data remains securely within your network.

### Is it possible to run multiple EmailEngine instances with shared Redis database?

Currently, this is not supported. Horizontal scaling is on the development roadmap, but there is no estimated time of arrival for this feature at this time.

### Why am I receiving a lot of 504 errors for API calls?

IMAP is a single-threaded protocol, meaning that only one command can be executed at a time per account, and you must wait for it to finish before issuing another command. If multiple API requests are made against a single account simultaneously, EmailEngine will queue these requests and process them one at a time. If a queued command cannot be processed within 10 seconds, EmailEngine will abort it with a 504 response.

To resolve this, you have two options: either avoid sending multiple requests at the same time, or increase the 10-second limit. The "Max command duration" option can be found in the configuration options, as documented [here](https://emailengine.app/configuration#general).

### Is EmailEngine similar to Nylas Universal Email API?

For a detailed comparison, you can refer to the comparison article [here](https://docs.emailengine.app/emailengine-vs-nylas/). In brief, while both may seem alike, they are in fact, significantly different.
