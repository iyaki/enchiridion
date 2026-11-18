---
title: "Flagsmith - Open Source Feature Flag & Remote Config Service"
notion_id: bf59be64-1be1-4512-b03f-3a0609e554d8
notion_url: https://app.notion.com/p/Flagsmith-Open-Source-Feature-Flag-Remote-Config-Service-bf59be641be14512b03f3a0609e554d8
last_edited: 2026-09-21T17:05:00.000Z
source_url: https://www.flagsmith.com/
tags: ["English", "Programming", "System Design / Software Architecture", "DevOps", "Product Management", "Continuous Integration/Continuous Delivery", "Tool", "Service"]
---
Release features with confidence; manage feature flags across web, mobile, and server side applications. Use our hosted API, deploy to your own private cloud, or run on-premises

[Start For Free](https://app.flagsmith.com/signup)

```plain text
1import flagsmith from 'flagsmith';
2
3flagsmith.init({
4  environmentID: 'QjgYur4LQTwe5HpvbvhpzK',
5});
6
7const App = () => (
8    <Layout
9      darkMode={flagsmith.hasFeature("dark_mode")}
10      designV2={flagsmith.hasFeature("design_v2")}
11    >12      {flagsmith.hasFeature("chat") && <ChatWidget>}
13    </Layout>
14)
15

```

![image](data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNTAwIiBoZWlnaHQ9IjUwMCIgdmlld0JveD0iMCAwIDUwMCA1MDAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik0yNTAuNzkyIDEuNjRjODEuMjQ0LTEwLjgxIDE2Mi43MyAzMy4wNDggMjEwLjcxMyA5NS42NTggNDUuNTgxIDU5LjQ3NSA0Ny42NiAxMzguNzkzIDIyLjAzMyAyMDguMDM2LTIyLjQzIDYwLjYtODQuOTg0IDkxLjA2OC0xNDQuNDAyIDEyMi41OS02Ny42NjEgMzUuODk1LTEzNy40NTkgOTEuMDctMjEwLjI1IDY1LjU1OEM1MS44NTUgNDY2LjQ4NCA1Ljg2MSAzODYuMy4yNTcgMzA5LjE1M2MtNC44NS02Ni43NjcgNjAuMDM0LTExMS4xMTEgMTAzLjMxLTE2NC4yM0MxNDguNzU4IDg5LjQ1NSAxNzcuMjMyIDExLjQyNyAyNTAuNzkyIDEuNjR6IiBmaWxsPSIjRjFFREU0Ii8+PC9zdmc+)

![image](data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNTAwIiBoZWlnaHQ9IjU3MCIgdmlld0JveD0iMCAwIDUwMCA1NzAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+PHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik0yNDYuMDEyIDU2OS42MjdDMTYwLjcyMSA1NzMuODkxIDcwLjU3IDU0MS42NzMgMjIuOTcgNDY3LjYxOWMtNDQuNTMyLTY5LjI3OS0xMy43OTYtMTU2LjY4NCA0LjU1Ni0yMzcuOSAxOC4wODQtODAuMDM0IDI3LjQyLTE3Mi45NzYgOTYuOTgzLTIxMS40NyA3MS40NC0zOS41MzMgMTU3LjE0MS04LjI0IDIyOC41NjggMzEuMzIgNjcuMzY0IDM3LjMxIDEyMi45NTggOTUuNzA2IDE0MC42NjEgMTczLjE3OCAxNy45MTUgNzguMzk5LTMuODAyIDE1OS4yMjktNDkuOTM5IDIyMy44MzMtNDguMjc1IDY3LjU5Ny0xMTcuMTQ3IDExOS4wMTYtMTk3Ljc4NyAxMjMuMDQ3eiIgZmlsbD0iI0YxRURFNCIvPjwvc3ZnPg==)

Trusted by top development teams

Flagsmith provides an **all-in-one platform for developing, implementing, and managing your feature flags**. Whether you are moving off an in-house solution or using toggles for the first time, you will be amazed by the power and efficiency gained by using Flagsmith.

[Learn More](https://docs.flagsmith.com/basic-features/managing-features)

### Manage Flags across multiple platforms

Flagsmith makes it easy to create and manage feature toggles across web, mobile, and server-side applications. Just wrap a section of code with a flag, and then use Flagsmith to manage that feature.

Manage feature flags by development environment, and for individual users, a segment of users, or a percentage. This means quickly implementing practices like canary deployments.

Multivariate flags allow you to use a percentage split across two or more variations for precise A/B/n testing and experimentation.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Remote config

Flagsmith combines the value of feature toggles with the flexibility of remote config, making it easier than ever to test and deploy new features. That’s because **all flags in Flagsmith are capable of being configured for remote config**. What does that mean for you? As well as toggling a flag, with remote config you can customize values for your features, allowing you to deploy functional and visual changes to your users without changing any code or pushing out any updates.

[Learn More](https://docs.flagsmith.com/basic-features/managing-features)

### Remote application management

Remote config lets you alter an application in real-time, without having to wait for a deployment or app store approval.

Configure different elements of your features–like checkout payment options of the color of a button–directly through Flagsmith and release the changes to users in just a few clicks.

Develop user segments based on traits or behaviors, then alter the application experience for subsections of your user base quickly and easily.

Flagsmith makes it easy to store traits associated with your users without modifying your back-end or worrying about transferring data. Do you know who your early adopters are? What about your most active users? Understanding your user base lets you roll out updates strategically.

Want your power users to be the first to test new features? Create detailed user segments based on stored traits, then roll out features based on those segments. As a feature is released, segments can be added or removed without requiring an update or any code changes. Test in production by making internal users or specific teams a segment and only exposing it to these internal segments. Nothing gives a release more confidence than having already been tested in production.

If you’re not sure how a new feature will perform, why roll it out to your entire user base right away? Our platform lets you deploy features to a percentage of your user base. It’s easy to roll out to more users, and eventually to everyone if features perform well; or roll back updates that don’t perform as expected. Flagsmith integrates with popular customer data platforms and analytics products, so you can utilize the tools you already use to analyze the results of tests and fine-tune your application.

The only thing worse than rolling out a feature with bugs is not being able to revert back. Tracking changes to your application lets you pinpoint what edits were made when the error arose so you can undo the offending update. Flagsmith logs all changes for auditing and to make rolling back any features painless. Since the feature was launched with a feature flag, that rollback will be as simple as switching a toggle.

```plain text
1import flagsmith from 'flagsmith/isomorphic';
2import { useFlags, FlagsmithProvider } from 'flagsmith/react';
3
4const App = ({ Component, pageProps, flagsmithState }) => (
5    <FlagsmithProvider serverState={flagsmithState} flagsmith={flagsmith}>6        <Component {...pageProps} />7    </FlagsmithProvider>8)
9
10App.getInitialProps = async () => {
11  await flagsmith.init({ environmentID: "QjgYur4LQTwe5HpvbvhpzK"});
12  return { flagsmithState: flagsmith.getState() }
13}
14
15const HomePage = () => {
16 const flags = useFlags(['chat_widget']);
17 return <>{flags.chat_widget.enabled && <ChatWidget>}</>18}

```


