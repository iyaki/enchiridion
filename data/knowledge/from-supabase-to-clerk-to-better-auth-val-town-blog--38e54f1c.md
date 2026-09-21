---
title: "From Supabase to Clerk to Better Auth | Val Town Blog"
notion_id: 38e54f1c-7d23-8119-a138-d302fbb987f9
notion_url: https://app.notion.com/p/From-Supabase-to-Clerk-to-Better-Auth-Val-Town-Blog-38e54f1c7d238119a138d302fbb987f9
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://blog.val.town/better-auth
tags: ["Val Town Blog", "English", "Web Development", "Programming", "Authentication", "Databases", "DevOps", "Article", "Tutorial"]
---
![image](https://imagedelivery.net/iHX6Ovru0O7AjmyT5yZRoA/d58722e1-ec33-4417-59a9-e63119bd6d00/public)

In 2023, I wrote about [how Val Town moved away from Supabase](https://blog.val.town/blog/migrating-from-supabase) and toward a more conventional database setup. We were using a lot of Supabase's functionality, including their authentication. So when it came time to move we found equivalents: [Render](https://render.com/) for the database, and [Clerk](https://clerk.com/) for authentication. But life came at us fast, and by late 2023 we had an issue filed: get off of Clerk. That issue was finally closed a month ago, when we switched to [Better Auth](https://www.better-auth.com/).

Some important context is that Clerk is a major success. They just raised [50 million dollars](https://clerk.com/blog/series-c) and they have lots of satisfied users. Heck, in related news Supabase [raised 100 million dollars at a 5 billion dollar valuation](https://supabase.com/). Congratulations to both of them. I would make a terrible venture capitalist. Whatever opinions and experiences I hold about authentication and row level security are secondary to these numbers and proof. You can't argue with success.

But still, I am happy to have closed that issue and switched to Better Auth. It's been a tough experience, with a lot of workarounds, bugs, and outages. The architecture of Val Town sharply conflicted with Clerk's expectations.

### The core issue

The core issue is that Clerk tried to be your users table _and_ your sessions table. I think they've been shifting away from saying this, but it started from a pretty extreme place: there's a 2021 blog post titled "[Consider dropping your users table](https://clerk.com/blog/offload_user_table)". There's a YouTube video from 2023 called [DELETE your Users table](https://www.youtube.com/watch?v=86sFORO-M3Y). I strongly suggest you don't!

There are two big problems with farming out your users table to a third-party service.

**Clerk was a pretty bad replacement for a users table because it was heavily rate-limited and not very reliable.** When we initially switched, I assumed we could load user data from Clerk's API whenever we needed to. After all, we want to know things like user settings, avatar URLs, and emails. Clerk's SDK made this pretty convenient: the rootAuthLoader, the thing that handles auth for your whole application, had a nice little option called `loadUser` which would do the request for you. Worked great in development. In production, the rate limit for that endpoint was _five requests per second_. For the whole account, across all users. Tough! A pretty bad footgun, that we discovered in production, and was eventually [fixed by removing the option](https://github.com/clerk/javascript/issues/1043).

The rate limiting hit us especially hard for the social aspect of Val Town. For example, let's say you have a social website, so a lot of pages have lists of content from other users, and usernames and avatars to identify them. Clerk's default UIs are based on the assumption that a user only sees _their own avatar_ and needs their own settings and information, and they can get all of that stuff from their nifty [JWT](https://datatracker.ietf.org/doc/html/rfc7519) token. Social websites like Val Town completely break this assumption, and the advice was to sync avatars and other information between Clerk and our users table: so now we have two authorities for that information, and the complexity of _two_ users tables.

So we had to sync Clerk's data to our database by using webhooks, which meant that signing up was convoluted and tricky - for a few moments, a user has a Clerk account but no Val Town database row. Or, because our platform requires usernames, users can be in a state where they have a Clerk account, a database row, but their account is incomplete. Our user settings had to be split between things that Clerk controlled, like auth strategies, and things that we needed, like usernames and editor settings.

**The second is that Clerk became a single point of failure for all our user sessions.** Cookie-based user sessions are usually short-lived with constant refreshing: that way they can be invalidated quickly. But that also means that every few minutes users need to swap their session cookies for new ones. So when someone's login session needed to be refreshed, it was a subdomain of Val Town that passed the request to Clerk that did the refreshing. We didn't have a sessions table or any responsibility over sessions.

That's great, if you're trying to keep avoid any responsibility for authentication, but on the other hand, if Clerk goes down, the whole website goes down. Clerk outages don't just break the login & logout flow, they make the site unusable to people who are already logged in. And Clerk went down pretty often, and went down for long periods of time. [Since May 2025](https://status.clerk.com/) it's been teetering between two and three nines of uptime. There isn't data from before then, but I remember many times that we had a broken site and no way to fix it because of this single point of failure.

A hard lesson you learn building a complex system is that its reliability is the _minimum_ of the combined reliability of its critical parts.

Besides these two major issues, there were [other bugs and problems we encountered](https://github.com/clerk/javascript/issues?q=sort%3Aupdated-desc%20is%3Aissue%20state%3Aclosed%20author%3Atmcw). Most got fixed eventually, but I spent a lot of time battling the "Stale Issue Bot" from auto-closing them.

### Three-ish years

If it was so bad, why didn't we switch away immediately?

First of all, even though this will be the second "switching from X to Y" article I've written, and I'm not trying to make it a habit. Making decisions and sticking with them is good for development velocity and team sanity. We're not trying to rewrite Val Town any more than is absolutely necessary. And writing critique is less fun and positive than building.

And Clerk did some things well. They provided SDKs for all the tech we were using: Remix, Fastify, and Express. They did a decent job of keeping up with the churn of those frameworks, a task that I know is a full-time job. And their administration and anti-abuse measures were decent at helping us solve customer support issues and keep scammers at bay.

Where Clerk definitely shines is relatively simple, heavily frontend apps that don't have a social component, so they don't need a users table. It was incredibly easy to get started, affordable, and the Clerk dashboard is pretty nice.

And there aren't a ton of great options for authentication. The bar for a Clerk replacement was actually pretty high: a lot of open source auth solutions are ancient and semi-abandoned. Authentication-as-a-service platforms had vendor risk and potentially the same problems as those in Clerk. The right level of technical control is hard to nail. We don't want to build authentication from scratch and open Val Town up to new and embarrassing new vulnerabilities, but we also don't want to offload so much responsibility to a provider. Definitely not trusting third-party session management again.

### Better Auth enters the frame

[Better Auth](https://better-auth.com/) checked a lot of boxes right out of the gate: high code quality, good integrations with different frameworks, and truly usable as an independent open source project.

There still is vendor risk: it's a big, complex codebase developed mostly by one company. There's always vendor risk. But we are no longer dependent on a third party staying online in order for sessions & user auth to work.

A close second place was [AuthKit](https://www.authkit.com/) from [WorkOS](https://workos.com/). I trust WorkOS and AuthKit is incredibly slick, but after bouncing between two vendors, it was important to me to find something that could work independently and was open source at the core.

I find Better Auth's dashboard and paid add-ons to be really clever, too. We manage all of our data, and a plugin provides an API on our site that lets their dashboard pull information and do some light user administration. Better Auth's paid service (called 'Infrastructure') is basically stateless in the way that we use it, and uninvolved in session management.

In short, so far it really has been better.

And reluctantly I have to hand it to the LLMs here: with the augmentation of the robots, we were able to take the more complex route of supporting _both_ Better Auth and Clerk for a transitional period of two weeks. Every endpoint that handled authentication would accept either kind of cookie, and users slowly moved over to Better Auth because that was the kind of session that the sign-in page provided. Like anything related to security, close reading, rewriting, and testing of all of the code was necessary to make sure we didn't self-own, and the eventual pure-Better Auth auth was handwritten entirely.

Better Auth also works pretty well with Vals: you can [try out the Better Auth starter template](https://www.val.town/x/templates/better-auth-starter) to add authentication to your code on Val Town.

I've learned a lot along the way. You really do depend on upstream providers for your uptime, and should think hard about how exposed you are to that risk. Products can be good for a lot of use-cases and really successful and still not the right thing for your specific problem. The world of software changes quickly and the right solution might not exist at the moment you need it, but might appear a year later.
