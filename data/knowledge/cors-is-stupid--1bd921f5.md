---
title: "CORS is Stupid"
notion_id: 1bd921f5-03de-42af-a43b-a50e10457982
notion_url: https://app.notion.com/p/CORS-is-Stupid-1bd921f503de42afa43ba50e10457982
last_edited: 2024-09-17T16:25:00.000Z
source_url: https://kevincox.ca/2024/08/24/cors
tags: ["English", "Web Development", "Article", "Kevin Cox"]
---
Posted 7 days ago

[CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS), and the browser’s [same-origin policy](https://developer.mozilla.org/en-US/docs/Web/Security/Same-origin_policy#cross-origin_network_access) are often misunderstood. I’m going to explain what they are and what you need to do to stop worrying about them.

First and foremost CORS is a giant hack to mitigate legacy mistakes. It provides both opt-out protections as an attempt to mitigate XSRF attacks against unaware or unmodified sites and opt-in protections for sites to actively protect themselves. But none of these protections are actually sufficient to solve the intended problem. If your site uses [cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies) you **must** take action to be safe. (Ok, not _every_ site, but you better not rely on it. Carefully audit your site or follow these simple steps. Very reasonable seeming patterns can expose you to XSRF vulnerabilities.)

## [The Problem](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#the-problem)

The key problem is how **implicit credentials** are handled in the web. In the past browsers made the disastrous decision that these credentials could be included in cross-origin requests. This opened up the following attack vector.

1. Log in to `https://your-bank.example`.
2. Visit `https://fun-games.example`.
3. `https://fun-games.example` runs `fetch("https://your-bank.example/profile")` to read sensitive information about you like your address and current balance.

This worked because when you logged into your bank it issued you a cookie to access your account details. While `fun-games.example` can’t just steal that cookie, it could make its own requests to your bank’s API and your browser would _helpfully_ attach the cookie to authenticate _you_.

## [The Solution](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#the-solution)

This is where CORS comes in. It describes a policy for how cross-origin requests can be made and used. It is both incredibly flexible and completely insufficient.

The default policy allows making requests, but you can’t read the results. So `fun-games.example` is blocked from reading your address from `https://your-bank.example/profile`. It can also use side channels such as latency and if the request succeeded or failed to learn things.

But despite being incredibly annoying this doesn’t actually solve the problem! While `fun-games.example` can’t read the result, the request is still sent. This means that it can execute `POST https://your-bank.example/transfer?to=fungames&amount=1000000000` to transfer one billion dollars to their account.

This must be one of the biggest security compromises ever made in the name of backwards compatibility. The TL;DR is that the automatically provided cross-origin protections are completely broken. Every single site that uses cookies needs to explicitly handle it.

Yes, **every single site**.

## [Actually Solving the Problem](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#actually-solving-the-problem)

The key defence against these cross-site attacks is ensuring that implicit credentials are not inappropriately used. It is best to start by ignoring all implicit credentials on cross-site requests, then you can add specific exceptions as required.

The best solution is to set up server-wide middleware that ignores implicit credentials on all cross-origin requests. This example strips cookies, if you use HTTP Authentication or TLS client certificates be sure to ignore those too. Thankfully the [`Sec-Fetch-*`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site)[ headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site) are now available on all modern browsers. This makes cross-site requests easy to identify.

```plain text
def no_cross_origin_cookies(req):
	if req.headers["sec-fetch-site"] == "same-origin":
		# Same origin, OK
		return

	if req.headers["sec-fetch-mode"] == "navigate" and req.method == "GET":
		# GET requests shouldn't mutate state so this is safe.
		return

	req.headers.delete("cookie")

```

This provides a safe baseline. If needed you can add specific exceptions for endpoints that are prepared to handle cross-origin implicitly authenticated requests. I would strongly recommend against wide exceptions.

## [Defence in Depth](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#defence-in-depth)

### [Explicit Credentials](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#explicit-credentials)

One of the best ways to avoid this whole problem is by not using implicit credentials. If all authentication is done via explicit credentials then you don’t need to worry about when the browser may add a cookie that you didn’t expect. Explicit credentials can be obtained by signing up for an API token or via an OAuth flow. But either way the most important thing is that logging into one site won’t allow other sites to use these credentials.

The best way to do this is to pass an authentication token in the `Authorization` header.

```plain text
Authorization: Bearer chiik5TieeDoh0af

```

Using the `Authorization` header is a standardized behaviour and will be handled well by many tools. For example this header is likely redacted from logs by default.

But most importantly it must be set explicitly by all clients. This not only solves the XSRF problem but makes multi-account support a breeze.

The major downside is that explicit credentials are unsuitable for server-rendered sites as they aren’t included in top-level navigation. Server-rendering is great for performance so this technique is often unsuitable.

### [`SameSite`](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#samesite-cookies)[ Cookies](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#samesite-cookies)

Even though our server should be ignoring cookies in cross-origin requests it is good to practice avoid including them in the requests in the first place. You should set the `SameSite=Lax` attribute on all of your cookies which will cause the browser to omit them for cross-origin requests.

It is important to remember that **cookies are still included in top-level GET navigations**. You can use `SameSite=Strict` which avoids this, but will make the user appear logged out for the first page load after following a cross-origin link (as that request will lack cookies).

Using `SameSite` cookies also prevents all cross-site form posts with no way to opt-out for a few specific endpoints. Luckily this is a very rare use case and most sites can get by without needing this. I would definitely recommend setting this attribute by default, and only relying on other mechanisms if you actually hit a use case.

## [CORS Policy](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#cors-policy)

A simple copy-pastable policy is this:

That’s it, you’re done.

The effect of this policy is that other sites can only make anonymous requests. This means that you are just as secure as if these requests were made via a CORS Proxy.

### [Shouldn’t I be more Specific?](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#shouldnt-i-be-more-specific)

Probably not. There are a few reasons for this:

1. You may create a false sense of security. Just because a different webpage running in a “correctly functioning” browser can’t make these requests doesn’t mean that they can’t be made. For example CORS proxies are very common.
2. It prevents read-only access to your site. This can be useful for URL previews, [feed fetching](https://kevincox.ca/2022/05/06/rss-feed-best-practices/#cors) or other features. This results in more CORS proxies being run which just harms performance and user privacy.

Remember, CORS is not about blocking access, it is about preventing the accidental reuse of implicit credentials.

## [Rant](https://kevincox.ca/2024/08/24/cors/?utm_source=tldrwebdev#rant)

Why do I need to know all of this stuff, why isn’t the web safe by default? Why do I have to deal with an ineffective policy that makes everything annoying by default without actually solving anything?

IDK, it’s quite annoying. I think most of the reason goes back to backwards compatibility. Sites built features on these security holes, so browsers tried to close them as much as they could without breaking existing sites.

Luckily there may be some sanity on the horizon, with browsers finally willing to break some sites for the good of the user. Major browsers are moving toward top-level domain isolation. They call it different things, Firefox calls it [State Partitioning](https://developer.mozilla.org/en-US/docs/Web/Privacy/State_Partitioning), Safari calls it [Tracking Prevention](https://webkit.org/tracking-prevention/#intelligent-tracking-prevention-itp) and Google likes cross-site tracking cookies, so they have implemented an opt-in [CHIPS](https://developers.google.com/privacy-sandbox/cookies/chips) system.

The main problem is that these approaches are being implemented as privacy features, not security features. This means that they can’t be relied on as they use heuristics to _sometimes_ allow cross-origin implicit credentials. CHIPS is actually better in this respect as it is reliable in supporting browsers, but it only supports cookies.

So it does seem like browsers are moving away from cookies that span top-level contexts, but it is a uncoodinated bumbling. It also isn’t clear if this will be by blocking third-party cookies (Safari), partitioning (Firefox, CHIPS).

- [scvalex@](https://lobste.rs/~scvalex)[lobste.rs - 6 days ago](https://lobste.rs/s/98rp8f)

Shared this post.

- [mary????@](https://bsky.app/profile/mary.my.id)[bsky.app - 6 days ago](https://bsky.app/profile/did:plc:ia76kvnndjutgedggx2ibrem/post/3l2leehullcvq)

kevincox.ca/2024/08/24/cor… "You may create a false sense of security. Just because a different webpage running in a “correctly functioning” browser can’t make these requests doesn’t mean that they can’t be made. For example CORS proxies are very common."

- [Michael Gebetsroither@](https://bsky.app/profile/gebi.bsky.social)[bsky.app - 6 days ago](https://bsky.app/profile/did:plc:ia76kvnndjutgedggx2ibrem/post/3l2leehullcvq#reposted_by_did:plc:zy7bl7rr2ka4nds6pjbdmue5)

Shared a post by bsky.brid.gy/web/did:plc:ia... kevincox.ca/2024/08/24/cor… "You may create a false sense of security. Just because a different webpage running in a “correctly functioning” browser can’t make these requests doesn’t mean that they can’t be made. For example CORS proxies are very common."

- [zerobytes.monster - 6 days ago](https://zerobytes.monster/post/2665963)

Mentioned this post.
