---
title: "From Development to Real Users: How to Create a Web Performance Story"
notion_id: beb55471-8a42-4c70-8346-b77a4a06f344
notion_url: https://app.notion.com/p/From-Development-to-Real-Users-How-to-Create-a-Web-Performance-Story-beb554718a424c708346b77a4a06f344
last_edited: 2023-01-27T17:11:00.000Z
source_url: https://engineering.atspotify.com/2022/09/from-development-to-real-users-how-to-create-a-web-performance-story/
tags: ["Product Management", "Article", "Guide", "Spotify Engineering", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Some of the most common questions asked when it comes to work with performance are, How do you convince stakeholders that improving the performance of your project is actually worth the investment? How can you prove that the work is necessary to begin with? Or prove that you have shipped improvements? And what is the impact of certain changes on users in different scenarios?

The answer to those questions? Data!

## Step 1: Get to know your lab data

The start of your performance story begins with your development environment, or lab data. At Spotify, web applications have access to [lighthouse-ci](https://github.com/GoogleChrome/lighthouse-ci) as a [Backstage](https://engineering.atspotify.com/2020/03/what-the-heck-is-backstage-anyway/) integration, which makes measuring performance for our web applications seamless.

With this integration, we gain an understanding of how our code will potentially affect real users. We leverage the powerful cli and [Node-API](https://github.com/GoogleChrome/lighthouse-ci/tree/main/packages/utils) to create a feedback report posted as a comment on each of the project’s pull requests, comparing against the latest results from master on key areas such as Web Vitals and performance budgets. This integration also makes use of the great visualization tools from [Lighthouse CI server](https://github.com/GoogleChrome/lighthouse-ci/blob/main/docs/getting-started.md#the-lighthouse-ci-server), where you can see audit data for each run in detail and compare between two runs to learn how they diverge and why. This also gives you a very important perspective on the history of your metrics and trends.

Additionally, you can leverage the Lighthouse configuration through the [lighthouserc.json](https://github.com/GoogleChrome/lighthouse-ci/blob/main/docs/configuration.md) file and change how the profiling will be run against parameters from your client base according to your analytics data. This way, your profiling results will better represent your users.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Lighthouse CI server audit view: This view gives you an idea of progress over time and helps you identify reasons for regressions and improvements made on the codebase within each commit.

## Step 2: Get to know your users

Collecting real users’ data to build better lab profiling tools is only an initial step. In addition to a good lab data setup, it’s vital to know how each change affects real users. Your lab data can’t be the only source of truth, because it doesn’t have enough representation to reveal how your project runs given the different conditions out there.

Lab data is used as a sort of smoke test that can highlight issues that may appear in field data. So to understand how real users experience your application, field data is necessary to determine what real impact particular changes can have.

## Step 3: Identify what to measure

For live applications, it’s important to understand which metrics are the most impactful and meaningful for your user’s experience. And it is for that purpose that we have the [Web Vitals project](https://web.dev/vitals/), providing both APIs and tooling to identify core, user-centric metrics for applications. Understanding those metrics is also important, to give background on how they affect the user experience.

So take some time to explore the official documentation on the core metrics and other metrics supported by Web Vitals library:

### Core metrics

[LCP](https://web.dev/lcp/)

“Largest Contentful Paint (LCP) is an important, user-centric metric for measuring perceived load speed because it marks the point in the page load timeline when the page’s main content has likely loaded — a fast LCP helps reassure the user that the page is useful.”

“First Input Delay (FID) is an important, user-centric metric for measuring load responsiveness because it quantifies the experience users feel when trying to interact with unresponsive pages — a low FID helps ensure that the page is usable.”

“Cumulative Layout Shift (CLS) is an important, user-centric metric for measuring visual stability because it helps quantify how often users experience unexpected layout shifts — a low CLS helps ensure that the page is delightful.”

### Other supported metrics

“First Contentful Paint (FCP) is an important, user-centric metric for measuring perceived load speed because it marks the first point in the page load timeline where the user can see anything on the screen — a fast FCP helps reassure the user that something is happening.”

“Time to First Byte (TTFB) is a foundational metric for measuring connection setup time and web server responsiveness in both the lab and the field. It helps identify when a web server is too slow to respond to requests. In the case of navigation requests — that is, requests for an HTML document — it precedes every other meaningful loading performance metric.”

### Experimental metrics

“Interaction to Next Paint (INP) is an experimental metric that assesses responsiveness. INP observes the latency of all interactions a user has made with the page, and reports a single value which all (or nearly all) interactions were below. A low INP means the page was consistently able to respond quickly to all — or the vast majority — of user interactions.”

## Metrics on live environment

Utilizing Web Vitals to capture real users’ metrics, we can leverage the Google Analytics export to BigQuery and assemble a dashboard for real-time metrics to monitor our performance in the wild.

These tools have been put together in a very helpful [video](https://www.youtube.com/watch?v=xg47r3Y6K8I) and [article](https://web.dev/vitals-ga4/) from Google I/O 2021. We at Spotify have successfully integrated those tools and currently have a dashboard that we use to monitor our performance with real users.

## Seeing from your user’s perspective

Let’s start by checking one of our performance experiments as an example. I’ll utilize our Server Side Rendering (SSR) case study for the mobile web player on the Show page back in January 2022.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Graph from our real user data monitoring tool, showing our state before SSR

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Graph from our real user data monitoring tool, showing our state after SSR.

My team had a theory that SSR could improve LCP for the mobile web player, one of the main areas that needed improvement. But without historical data for context, it would be extremely difficult to know either how bad our LCP was or how our work would impact real users. With our monitoring in place, we could assess the improvements our work would bring to real users.

That’s the value real user data can give — you can kickstart investigations and dive deep into why certain metrics need improvement. The data will lead you to the parts of your project that need attention and where your users are struggling the most. And it will also help you validate that the work you’ve done has had the intended impact.

After finding the areas in need of improvement, we perform in-depth profiling of those areas, utilizing tools like the built-in Lighthouse integration on Google Chrome and the [performance tab](https://developer.chrome.com/docs/devtools/evaluate-performance/). After gathering the data from the profiling by your tool of choice, you will understand which parts need to be fixed and work on the improvement areas it highlighted.

Each performance problem needs a different fix, and there’s always an “it depends” clause. But investigating performance problems based on data is much easier than exploring what might be wrong and trying to come up with blanket solutions that might not bring the biggest impact to real users.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In the figure above, one of our pages, the Playlist page, had a high LCP value. With that in hand and after a thorough analysis, we were able to ship a series of improvements to lower that LCP value and improve our user experience (e.g., prioritizing the LCP element for loading and rendering).

## Closing thoughts

Creating a performance story can only start with the data. Having a better understanding of your app’s state and continuously measuring and monitoring it is key to making sure you can ship the most impactful results to your users. And keeping track of performance journeys can help you better understand your product and the relationship between what you ship and what the users actually get.

With good tooling in place, you can identify key user’s journeys in your app and identify ways to better instrument those and ship impactful results in return. Trying to apply corrections on a general basis may not always work, because not all applications and user journeys are the same.

Tags: [web](https://engineering.atspotify.com/tag/web/)
