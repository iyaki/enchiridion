---
title: "Resend - Email for developers"
notion_id: 14d4753f-8c12-4faa-80b2-c72333b0a4d9
notion_url: https://app.notion.com/p/Resend-Email-for-developers-14d4753f8c124faa80b2c72333b0a4d9
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://resend.com/
tags: ["Email", "Untried", "Service", "English"]
---
Companies of all sizes trust Resend to deliver their most important emails.

## Integrate this morning

A simple, elegant interface so you can start sending emails in minutes. It fits right into your code with SDKs for your favorite programming languages.

```plain text
1import { Resend } from 'resend';23const resend = new Resend('re_123456789');45(async function() {6  try {7    const data = await resend.emails.send({8      from: 'onboarding@resend.dev',9      to: 'delivered@resend.dev',10      subject: 'Hello World',11      html: '<strong>it works!</strong>'12    });1314    console.log(data);15  } catch (error) {16    console.error(error);17  }18})();
```

## First-class developer experience

We are a team of engineers who love building tools for other engineers.

Our goal is to create the email platform we've always wished we had — one that _just works_.

## Test Mode

Simulate events and experiment with our API without the risk of accidentally sending real emails to real people.

[Learn more](https://resend.com/docs/dashboard/emails/send-test-emails)

## Modular Webhooks

Receive real-time notifications directly to your server. Every time an email is delivered, opened, bounces, or a link is clicked.

[Learn more](https://resend.com/docs/dashboard/webhooks/introduction)

## Develop emails using React

Create beautiful templates without having to deal with <table> layouts and HTML.

Powered by react-email, our open source component library.

Get StartedCheck the Docs

```plain text
1import { Body, Button, Column, Container, Head, Heading, Hr, Html, Img, Link, Preview, Row, Section, Text, Tailwind } from '@react-email/components';2import * as React from 'react';34const WelcomeEmail = ({5  username = 'newuser',6  company = 'ACME',7}: WelcomeEmailProps) => {8  const previewText = `Welcome to ${company}, ${username}!`;910  return (11    <Html>12      <Head />13      <Preview>{previewText}</Preview>14      <Tailwind>15      <Body className="bg-white my-auto mx-auto font-sans">16        <Container className="my-10 mx-auto p-5 w-[465px]">17          <Section className="mt-8">18            <Img19              src={`${baseUrl}/static/example-logo.png`}20              width="80"21              height="80"22              alt="Logo Example"23              className="my-0 mx-auto"24            />25          </Section>26          <Heading className="text-2xl font-normal text-center p-0 my-8 mx-0">27            Welcome to <strong>{company}</strong>, {username}!28          </Heading>29          <Text className="text-sm">30            Hello {username},31          </Text>32          <Text className="text-sm">33            We're excited to have you onboard at <strong>{company}</strong>. We hope you enjoy your journey with us. If you have any questions or need assistance, feel free to reach out.34          </Text>35          <Section className="text-center mt-[32px] mb-[32px]">36              <Button37                pX={20}38                pY={12}39                className="bg-[#00A3FF] rounded text-white text-xs font-semibold no-underline text-center"40                href={`${baseUrl}/get-started`}41              >42                Get Started43              </Button>44          </Section>45          <Text className="text-sm">46            Cheers,47            <br/>48            The {company} Team49          </Text>50        </Container>51      </Body>52      </Tailwind>53    </Html>54  );55};5657interface WelcomeEmailProps {58  username?: string;59  company?: string;60}6162const baseUrl = process.env.URL63  ? `https://${process.env.URL}`64  : '';6566export default WelcomeEmail;
```

| **Welcome to ****ACME****, user!**<br>Hello newuser,<br>We're excited to have you onboard at **ACME**. We hope you enjoy your journey with us. If you have any questions or need assistance, feel free to reach out.[Get Started](https://resend.com/get-started)<br>Cheers,<br>The ACME Team |
| --- |

## Reach humans, not spam folders

> Resend is transforming email for developers. Simple interface, easy integrations, handy templates. What else could we ask for.

Guillermo Rauch

CEO at Vercel

## Everything in your control

All the features you need to manage your email sending, troubleshoot with

detailed logs, and protect your domain reputation – without the friction.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Beyond expectations

Resend is driving remarkable developer experiences that enable success

stories, empower businesses, and fuel growth across industries and individuals

- "Our team loves Resend. It makes email sending so easy and reliable. After we switched to Dedicated IPs, our deliverability improved tremendously and we don't hear complaints about emails landing on spam anymore."Vlad MatsiiakoVlad MatsiiakoCo-founder of Infisical
- "I've used Mailgun, Sendgrid, and Mandrill and they don't come close to providing the quality of developer experience you get with Resend."Brandon StrittmatterBrandon StrittmatterCo-founder of Outerbase
- "Resend is an amazing product. It was so easy to switch over. I feel confident knowing that our important emails are in good hands with Resend. Everyone should be using this."Shariar KabirShariar KabirFounder at Ruby Card
- "All of our customers are located in South America, so having a solution that could send emails from the region closest to our users is very important. Resend's multi-region feature is a game-changer for us."Giovanni KeppelenGiovanni KeppelenCTO & Partner at VOA Hoteis
- "The speed and ease of integrating with the product was incredible, but what really stood out was their intricate knowledge of email and relentless support day or night. Oh and we also ended up winning Product of the week."Sam DuckerSam DuckerCo-founder of Anyone
- "As a developer I love the approach that the Resend team is taking. Its so refreshing. They are also extremely user-centric and helpful in terms of getting you up and running, sending beautiful emails that deliver."Hahnbee LeeHahnbee LeeCo-Founder at Mintlify
- "The Resend team have built a great product in a space that hasn't seen 10x innovation for years. Engineering peers are raving about Resend - it's such a smoother dev experience."Roberto RiccioRoberto RiccioHead of Product at Alliance
- "If you're a developer or working on a startup, you're going to love Resend's approach to emailing."Joe DeMariaJoe DeMariaCo-founder & CEO of SpecCheck
- "We were up and running with Resend in no time. It was seamless to integrate into our existing workflow and gave us a tremendous amount of visibility into our email capabilities. Simple to say, it was a no-brainer."Ty SharpTy SharpCo-founder & CEO of InBuild
- "Resend not only streamlines our emails to accommodate our expanding customer base, but their team also offered valuable hands-on support during the transition from our old API. Their product is visually stunning and seamlessly integrates with React Email."Thiago CostaThiago CostaCo-founder of Fey and Narative
- "As of our last deployment all of our emails are using Resend. We are loving the development experience of React Email - not having to leave my dev environment to develop new emails is a game-changer."Adam RankinAdam RankinFounding Engineer at Warp
- "Working with Resend has been amazing. By using Webhooks, I'm able to track email opened/clicked events via Segment and log those events in LogSnag for visibility. I highly believe in the people behind Resend."Taylor FacenTaylor FacenFounder of Finta
- "Resend is super easy to set up. Loving the modern approach the team is taking with supercharging email. Never been a fan of other clunky tools."Brek GoinBrek GoinFounder of Hammr
- "Our team loves Resend. It makes email sending so easy and reliable. After we switched to Dedicated IPs, our deliverability improved tremendously and we don't hear complaints about emails landing on spam anymore."Vlad MatsiiakoVlad MatsiiakoCo-founder of Infisical
- "I've used Mailgun, Sendgrid, and Mandrill and they don't come close to providing the quality of developer experience you get with Resend."Brandon StrittmatterBrandon StrittmatterCo-founder of Outerbase
- "Resend is an amazing product. It was so easy to switch over. I feel confident knowing that our important emails are in good hands with Resend. Everyone should be using this."Shariar KabirShariar KabirFounder at Ruby Card
- "All of our customers are located in South America, so having a solution that could send emails from the region closest to our users is very important. Resend's multi-region feature is a game-changer for us."Giovanni KeppelenGiovanni KeppelenCTO & Partner at VOA Hoteis
- "The speed and ease of integrating with the product was incredible, but what really stood out was their intricate knowledge of email and relentless support day or night. Oh and we also ended up winning Product of the week."Sam DuckerSam DuckerCo-founder of Anyone
- "As a developer I love the approach that the Resend team is taking. Its so refreshing. They are also extremely user-centric and helpful in terms of getting you up and running, sending beautiful emails that deliver."Hahnbee LeeHahnbee LeeCo-Founder at Mintlify
- "The Resend team have built a great product in a space that hasn't seen 10x innovation for years. Engineering peers are raving about Resend - it's such a smoother dev experience."Roberto RiccioRoberto RiccioHead of Product at Alliance
- "If you're a developer or working on a startup, you're going to love Resend's approach to emailing."Joe DeMariaJoe DeMariaCo-founder & CEO of SpecCheck
- "We were up and running with Resend in no time. It was seamless to integrate into our existing workflow and gave us a tremendous amount of visibility into our email capabilities. Simple to say, it was a no-brainer."Ty SharpTy SharpCo-founder & CEO of InBuild
- "Resend not only streamlines our emails to accommodate our expanding customer base, but their team also offered valuable hands-on support during the transition from our old API. Their product is visually stunning and seamlessly integrates with React Email."Thiago CostaThiago CostaCo-founder of Fey and Narative
- "As of our last deployment all of our emails are using Resend. We are loving the development experience of React Email - not having to leave my dev environment to develop new emails is a game-changer."Adam RankinAdam RankinFounding Engineer at Warp
- "Working with Resend has been amazing. By using Webhooks, I'm able to track email opened/clicked events via Segment and log those events in LogSnag for visibility. I highly believe in the people behind Resend."Taylor FacenTaylor FacenFounder of Finta
- "Resend is super easy to set up. Loving the modern approach the team is taking with supercharging email. Never been a fan of other clunky tools."Brek GoinBrek GoinFounder of Hammr

## Email reimagined.Available today.
