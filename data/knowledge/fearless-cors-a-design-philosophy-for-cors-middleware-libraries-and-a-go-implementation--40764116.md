---
title: "Fearless CORS: a design philosophy for CORS middleware libraries (and a Go implementation)"
notion_id: 40764116-38f6-4951-bd9b-a0baaaa067a7
notion_url: https://app.notion.com/p/Fearless-CORS-a-design-philosophy-for-CORS-middleware-libraries-and-a-Go-implementation-4076411638f64951bd9ba0baaaa067a7
last_edited: 2023-02-17T19:56:00.000Z
source_url: https://jub0bs.com/posts/2023-02-08-fearless-cors/
tags: ["Web Development", "Information Security", "REST API", "Network", "Article", "English"]
---
In this post, I investigate why developers struggle with CORS and I derive _Fearless CORS_, a design philosophy for better CORS middleware libraries, which comprises the following twelve principles:

This post also introduces [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors), an open-source, production-ready CORS middleware library for [Go](https://go.dev/) that adheres to _Fearless CORS_.

Familiarity with the CORS protocol is a prerequisite for reading this post; if you need to brush up on CORS, [MDN Web Docs](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a good resource. Some familiarity with Go, Java, and JavaScript is beneficial but not required. This post is quite long, but its structure should allow you to dip in and out of it as you please.

[Cross-Origin Resource Sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a mechanism that lets servers instruct browsers to relax, for select clients, some restrictions (in terms of both sending and reading) enforced by the [Same-Origin Policy (SOP)](https://developer.mozilla.org/en-US/docs/Web/Security/Same-origin_policy) on cross-origin network access. The protocol relies on special HTTP headers such as `Origin`, `Access-Control-Request-Method`, `Access-Control-Allow-Origin`, etc. Because a picture is worth a thousand words, let me show you a typical example of a successful CORS handshake:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

1. Bob visits Carol’s website (`https://carol.com`).
2. Carol’s client uses [the Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch) to issue a `GET` request to a resource on Sarah’s server (`https://sarah.com`). Carol includes a header named `Authorization` in his request.
3. The request is such that Bob’s browser first issues a [_preflight_](https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request)[ request](https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request) (which uses the [`OPTIONS`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/OPTIONS)[ method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/OPTIONS)) to the server.
4. The server sends a preflight response, which instructs Bob’s browser that Carol’s client is allowed to send her _actual_ (`GET`) request.
5. The browser sends Carol’s actual request to the server.
6. The server sends back a response, which also instructs Bob’s browser that Carol’s client is allowed to read the response in question.
7. The browser complies and gives Carol’s client access to Sarah’s response.

This is CORS in a nutshell. In practice, some daring server developers choose to implement CORS by “manually” setting the relevant HTTP response headers, either at the application level or [at the reverse-proxy level](https://enable-cors.org/server_nginx.html). But, as expressed by [Jake Archibald](https://twitter.com/jaffathecake)’s [pithy statement](https://jakearchibald.com/2021/cors/),

> CORS (Cross-Origin Resource Sharing) is hard.

CORS is, at the very least, more intricate than meets the eye. In practice, developers tend to rely instead on some middleware library, which can leverage the full power of their programming language to abstract some of CORS’s complexity.

Although CORS has been, since its inception in the late 2000s, an instrumental mechanism for the development of the modern Web, it remains a frequent source of confusion and exasperation for developers. At the time of writing this post, the number of [Stack Overflow questions tagged with “cors”](https://stackoverflow.com/questions/tagged/cors) hovers around 12,500, and a disconcertingly large fraction of those questions comes from a place of anguish. Some view CORS issues as a rite of passage into Web-development mastery. Others take to social media to give way to gloom or vent their frustration:

> I've been building web sites since 1998 and nothing has ever been as frustrating as dealing with CORS. Total clown town.

[August 28, 2014](https://twitter.com/adrianholovaty/status/505041023492300800?ref_src=twsrc%5Etfw)

> I don't think I mentioned it today...But in case you didn't know...I HATE CORS!!It's always CORS. The bane of my webdev existence. pic.twitter.com/Xv0Tm54O4a

[November 4, 2021](https://twitter.com/LaylaCodesIt/status/1456223619797880832?ref_src=twsrc%5Etfw)

> CORS has to be the single must frustrating part of any development I do. I've solved it so many times and yet every time I build an API I still end up down a Google rabbit hole 🤷‍♂️

> CORS issues keep me up at night.No matter how many times I build apps, I'm debugging CORS for hours upon hours... pic.twitter.com/DFAvTIzXQC

[January 13, 2023](https://twitter.com/andrewbrown/status/1613901675051507712?ref_src=twsrc%5Etfw)

[Productivity in software development is notoriously hard to measure](https://martinfowler.com/bliki/CannotMeasureProductivity.html), but obstacles to productivity are hard to ignore; and, if such public outcry as sampled above is anything to go by, one can only shudder at the thought of the cumulative time and money that has been wasted on troubleshooting CORS issues over the years.

There is a flip side to this: because CORS deactivates some browser defences, misconfiguring it can compromise, not just functionality, but also security. Insecure CORS configurations are perhaps less decried than dysfunctional ones but, [when they do occur](https://hackerone.com/hacktivity?querystring=cors%20misconfiguration), they have the potential to expose stakeholders to [devastating cross-origin attacks](https://portswigger.net/research/exploiting-cors-misconfigurations-for-bitcoins-and-bounties). In fact, server developers fighting a losing battle against a dysfunctional CORS configuration may be driven, as a last-ditch attempt to just “make things work”, to adopt an overly permissive CORS policy, such as [one that throws most of the SOP out the window](https://stackoverflow.com/a/9866124/2541573)!

Why this continual weeping and gnashing of teeth? Is the protocol itself to blame for this misery? No; I argue that CORS’s reputation as a productivity killer and security footgun is mostly undeserved. Are developers simply too dumb to bend CORS to their will, then? Developers in general would certainly benefit from familiarising themselves better with the protocol before using it or [opting for farfetched and dangerous alternatives](https://fosterelli.co/developers-dont-understand-cors), but assigning the blame entirely to them would be unfair. What about the tools? Contrary to popular opinion, I think browsers do a rather fine job in assisting developers with their CORS issues. This only leaves CORS libraries out of account…

Over the past few months, I’ve been driven to study the CORS protocol with more attention than I had formerly given to it. I’ve spent hours perusing current and old specifications, code spelunking in both prominent and obscure CORS libraries, rummaging through pull requests and GitHub issues, reading reports of CORS insecure configurations, sifting through countless Stack-Overflow questions about CORS and [answering them when I could](https://stackoverflow.com/search?q=user%3A2541573%20%5Bcors%5D%20is%3Aanswer)… At last, I came to the conclusion that developers' difficulties with CORS largely result from certain infelicities in CORS middleware libraries.

But identifying culprits isn’t enough: how shall we escape this predicament? In reaction to the shortcomings I perceive in existing CORS middleware libraries, I started gathering my thoughts about how I would design such a library from scratch, with the benefit of hindsight, and unemcombered by unfortunate past design decisions or promises of backwards compatibility.

This post enunciate my vision for a better CORS middleware library, one that designs CORS misconfigurations—both dysfunctional and insecure configurations—out of existence. Nicknamed _Fearless CORS_, my design philosophy ramifies into twelve language-agnostic principles, in a manner reminiscent of [Heroku’s Twelve-Factor-App methodology](https://12factor.net/). This post also presents [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors), a production-ready CORS middleware library for Go that adheres to _Fearless CORS_.

What follows is one third murder mystery, one third design manifesto, and one third billboard for [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors). I hope you enjoy it! My views may strike you as stubborn and contentious, and my tone as acerbic and pompous, but I’ve endeavoured to remain lucid, if only occasionally harsh, in my criticism of existing CORS libraries. Of course, I don’t have a monopoly on truth; if you disagree with _Fearless CORS_, I’d like to hear from you; and if you find a bug in [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) or identify a missing feature, please do [file an issue on GitHub](https://github.com/jub0bs/fcors/issues/new).

## The twelve principles of Fearless CORS

### 1. Optimise for readability

Most CORS middleware libraries—to the notable exception of [Spring’s](https://spring.io/guides/gs/rest-service-cors/) and [.NET Core’s](https://learn.microsoft.com/en-us/aspnet/core/security/cors?view=aspnetcore-7.0)—adopt an imperative style. Here is an example featuring [rs/cors](https://pkg.go.dev/github.com/rs/cors):

```plain text
c := cors.New(cors.Options{
  AllowedOrigins: []string{"https://example.com"},
  AllowedMethods: []string{
    http.MethodGet,
    http.MethodPost,
    http.MethodPut,
    http.MethodDelete,
  },
  AllowedHeaders: []string{"Authorization"},
})

```

For configuration purposes, though, a declarative style is often preferable to an imperative one in terms of readability. Here is how you would express an equivalent CORS policy with [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors):

```plain text
cors, err := fcors.AllowAccess(
  fcors.FromOrigins("https://example.com"),
  fcors.WithMethods(
      http.MethodGet,
      http.MethodPost,
      http.MethodPut,
      http.MethodDelete,
  ),
  fcors.WithRequestHeaders("Authorization"),
)

```

The differences are more than cosmetic. By contrast with [rs/cors](https://pkg.go.dev/github.com/rs/cors), [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) requires less ceremony. In particular, through its use of [variadic function parameters](https://go.dev/ref/spec#Function_types), [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) does away with those distracting [slice literals](https://go.dev/ref/spec#Composite_literals) (`[]string{}`). Moreover, through its use of a pattern known in the Go community as [_functional options_](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis), [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) obviates the need for exposing an intermediate configuration struct and lets you express your desired CORS policy in a more declarative style via a small embedded [domain-specific language](https://www.martinfowler.com/bliki/DomainSpecificLanguage.html). If you ignore the repeated ocurrences of the package name (“fcors”) in [qualified identifiers](https://go.dev/ref/spec#Qualified_identifiers), you’ll likely find the configuration code largely free of visual noise.

By saving developers a few keystrokes in this manner, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) can afford to err on the side of verbosity and self-documentation for the names of its options, e.g.

Not exposing any configuration struct presents another advantage: the impossibility to build upon an existing CORS policy. Although option values can be shared across calls to [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors)’s middleware factory functions (named `AllowAccess` and `AllowAccessWithCredentials`), this constraint tends to promote _colocation_ of CORS-configuration code. It also discourages the use of many different CORS policies, which [OWASP frowns upon](https://owasp.org/Top10/A01_2021-Broken_Access_Control/#how-to-prevent) anyway:

> 

Implement access control mechanisms once and re-use them throughout the application, including minimizing Cross-Origin Resource Sharing (CORS) usage.

As a result of these design decisions, configuration code written with [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) is comparatively easier to read and, hence, to review.

### 2. Strive for a simple and cohesive API

In [_The Computer Scientist as Toolsmith II_](https://www.cs.unc.edu/~brooks/Toolsmith-CACM.pdf), the late [Fred Brooks](https://en.wikipedia.org/wiki/Fred_Brooks) articulated an idea (slightly paraphrased here) that still resonates with me:

> Two criteria for success in a tool are:

- It must be so _easy_ to use that a seasoned developer _can_ use it, and
- It must be so _productive_ that seasoned developers _will_ use it.

However, the size of a library’s API can hurt both ease of use and productivity. [Joshua Bloch](https://twitter.com/joshbloch), author of [_Effective Java_](https://www.pearson.com/en-us/subject-catalog/p/effective-java/P200000000138/9780134685991) and of [Java’s Collections Framework](https://docs.oracle.com/en/java/javase/19/docs/api/java.base/java/util/package-summary.html#JavaCollectionsFramework), [insists](https://www.youtube.com/watch?v=aAb7hSCtvGw&t=24m19s) that the best libraries distinguish themselves as having a small _conceptual surface area_, and enjoins library authors to maximise _power-to-weight ratio_. Likewise, [John Ousterhout](https://twitter.com/johnousterhout), author of [_A Philosophy of Software Design_](https://web.stanford.edu/~ouster/cgi-bin/aposd.php), argues that software _modules_—libraries, in this case—should be _deep_ rather than _shallow_; in particular, between two libraries that have feature parity, the one whose interface is smaller is preferable:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Many CORS middleware libraries, no doubt due to their long development history, are not as _deep_ as they ideally could be. For example, [rs/cors](https://pkg.go.dev/github.com/rs/cors)’s [`Options`](https://pkg.go.dev/github.com/rs/cors#Options)[ struct type](https://pkg.go.dev/github.com/rs/cors#Options) provides no fewer than three different ways of expressing a CORS policy’s allowed origins:

- `AllowedOrigins`, of type `[]string`;
- `AllowedOriginFunc`, of type `func (string) bool`; and
- `AllowedOriginRequestFunc`, of type `func (*http.Request, string) bool`.

This corner of [rs/cors](https://pkg.go.dev/github.com/rs/cors)’s API strikes me as _shallow_. In particular, `AllowedOriginRequestFunc` subsumes—is strictly more powerful than—`AllowOriginFunc`. This overlap in functionality between the two functions is unfortunate, as it unnecessarily bloats the conceptual surface area of the library’s API. I shall revisit those two function fields [further down in this post](https://jub0bs.com/posts/2023-02-08-fearless-cors/#10-render-insecure-configurations-impossible), as I believe custom callbacks are misfeatures in their own right.

Besides, all three options lack self-documentation. What happens if, say, `AllowedOrigins` is used in conjunction with `AllowOriginFunc`? Are those options somehow additive? If not, which one has precedence over the other? Definite answers to these questions are not immediately apparent; you will only find them in the library’s documentation.

With [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors), developers can specify their allowed origins only via two orthogonal and (mutually incompatible) options named [`FromOrigins`](https://pkg.go.dev/github.com/jub0bs/fcors#FromOrigins) and [`FromAnyOrigin`](https://pkg.go.dev/github.com/jub0bs/fcors#FromAnyOrigin). Its comparatively smaller API also lends itself better to auto-completion by IDEs. Moreover, in a bid to minimise clutter and dispel any ambiguity, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) opts for non-additive options and raises a run-time error in the face of repeated and/or conflicting option calls:

```plain text
cors, err := fcors.AllowAccess(
  fcors.FromAnyOrigin(),
  fcors.WithMethods("GET", "POST", "PUT"),
  fcors.WithMethods("DELETE"), // ❌ (conflicts with earlier call)
)

```

### 3. Provide support for Private Network Access

[Private Network Access (PNA)](https://wicg.github.io/private-network-access/) is a W3C initiative that strengthens the Same-Origin Policy by denying clients in more public networks (e.g. the public Internet) access to more private networks (e.g. `localhost`); the goal is to mitigate a class of [cross-site-request-forgery attacks](https://owasp.org/www-community/attacks/csrf). PNA also extends CORS by providing a server-side opt-in mechanism for such access.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

At this stage, Private Network Access remains a moving target: its specification still has draft status and the W3C is [still debating whether to rename it to “Local Network Access”](https://github.com/WICG/local-network-access/issues/91). Nevertheless, PNA is gradually [getting support in Chromium](https://developer.chrome.com/blog/private-network-access-preflight/). Unfortunately, many CORS middleware libraries—[Gin’s](https://github.com/gin-contrib/cors/issues/81) is but one example—still lack support for PNA, which forces their users to look elsewhere for a solution.

Besides, PNA happens to be a source of new difficulties for library authors. One assumption that used to be true is that browsers would issue a preflight request only in reaction to a client’s attempt to send a _cross-origin_ request. Some CORS libraries, such as [Gin’s](https://github.com/gin-contrib/cors), still [actively rely on this assumption](https://github.com/gin-contrib/cors/blob/1d5e083a75f60fe6f09ef9463d896ef30eb2be5f/config.go#L70). In this regard, PNA throws a spanner in the works: as a mitigation against DNS rebinding, PNA-compliant browsers may now [issue a preflight request even in reaction to a client’s attempt to send a ](https://bugs.chromium.org/p/chromium/issues/detail?id=1299676)[_same-origin_](https://bugs.chromium.org/p/chromium/issues/detail?id=1299676)[ request](https://bugs.chromium.org/p/chromium/issues/detail?id=1299676)! Accordingly, the implementation of Gin’s CORS middleware library will require amendments, which may break clients who rely on the library’s current behaviour.

Because PNA is a natural extension of the CORS protocol, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) fully supports it through advanced options hidden away in [its ](https://pkg.go.dev/github.com/jub0bs/fcors/risky)[`risky`](https://pkg.go.dev/github.com/jub0bs/fcors/risky)[ companion package](https://pkg.go.dev/github.com/jub0bs/fcors/risky). Moreover, among [other similar constraints](https://jub0bs.com/posts/2023-02-08-fearless-cors/#10-render-insecure-configurations-impossible), [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) prohibits developers from enabling private network access from all origins, a practice that the team leading the PNA specification effort [actively discourages](https://developer.chrome.com/blog/private-network-access-preflight/#examples):

> The server can set `Access-Control-Allow-Origin: *`, though this is dangerous and discouraged. Private network resources should rarely be accessible to all origins, so think carefully about the risks involved in setting such a header.

Warning: In case the W3C opts for “Local Access Network” instead, I reserve the right to rename, for consistency reasons, the relevant options in a future v0.x version of [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors). This would be a breaking change, but [semantic versioning does allow breaking changes for versions prior to v1](https://semver.org/#spec-item-4).

### 4. Categorise requests correctly

The Fetch standard [tells us](https://fetch.spec.whatwg.org/#cors-preflight-request) that a preflight request is one that, at the very least,

In one of his posts, Jake Archibald [laments](https://jakearchibald.com/2021/cors/#cors-requests) that some servers incorrectly take the presence of an `Origin` header as an unmistakable cue that a request is cross-origin. I am similarly saddened by reverse proxies and CORS middleware libraries like [Express.js’s](https://github.com/expressjs/cors/blob/f038e7722838fd83935674aa8c5bf452766741fb/lib/index.js#L163) that fail to recognise that preflight requests form a _strict_ subset of all `OPTIONS` requests:

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This blunder unfortunately precludes proper handling of non-preflight `OPTIONS` requests, which never make it past the CORS middleware to the next HTTP handler in the chain. It also tends to reverberate throughout the codebase.

Rather than identifying the problem and tackling it at the root, the maintainers of [Express.js’s CORS middleware](https://expressjs.com/en/resources/middleware/cors.html) unfortunately opted instead to [retrofit the library](https://github.com/expressjs/cors/pull/40) to provide their users a kludgey way of letting non-preflight `OPTIONS` requests through, in the form of a new [configuration option named ](https://expressjs.com/en/resources/middleware/cors.html#configuration-options)[`preflightContinue`](https://expressjs.com/en/resources/middleware/cors.html#configuration-options). [The addition of an ](https://github.com/rs/cors/commit/a4c23809d15590a94f37d870f991f0c49998591c)[`OptionsPassthrough`](https://github.com/rs/cors/commit/a4c23809d15590a94f37d870f991f0c49998591c)[ option](https://github.com/rs/cors/commit/a4c23809d15590a94f37d870f991f0c49998591c) to [rs/cors](https://pkg.go.dev/github.com/rs/cors) and [the addition of an ](https://github.com/gorilla/handlers/commit/4d725479f7154fff401333083d412376c1b35570)[`IgnoreOptions`](https://github.com/gorilla/handlers/commit/4d725479f7154fff401333083d412376c1b35570)[ option](https://github.com/gorilla/handlers/commit/4d725479f7154fff401333083d412376c1b35570) to [gorilla/handlers](https://github.com/gorilla/handlers) were similarly motivated. Such options are pure manifestations of accidental complexity and make their library’s API harder to apprehend.

By contrast, because [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) correctly categorises `OPTIONS` requests, it requires no such cruft.

### 5. Validate configuration and fail fast

The sheer number of CORS middleware libraries that perform little to no configuration validation and unconditionally produce a middleware, even a dysfunctional one, baffles me. In this regard, those libraries truly do their users a disservice.

For example, developers not very familiar with the concept of [Web origin](https://developer.mozilla.org/en-US/docs/Glossary/Origin) may attempt to list, among the origins allowed by their CORS policy, a value that is actually not a valid origin, such as [a URI that lacks a scheme](https://stackoverflow.com/questions/56959505/quarkus-blocked-by-cors-policy) or [one that contains a path](https://stackoverflow.com/q/70353729/2541573). The following code snippet draws inspiration from [a real-life example](https://github.com/gofiber/fiber/issues/1441) involving [Fiber](https://gofiber.io/), a Web framework for Go:

For context, here is the signature of Fiber’s [factory function ](https://github.com/gofiber/fiber/blob/07ab88278bff1ee8c82ac1256a2239708294feff/middleware/cors/cors.go#L75)[`cors.New`](https://github.com/gofiber/fiber/blob/07ab88278bff1ee8c82ac1256a2239708294feff/middleware/cors/cors.go#L75):

```plain text
func New(...cors.Config) fiber.Handler

```

The use of a variadic parameter is itself [questionable](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis), but what I want to point out is that this function never “fails”: it does not return an error, nor does it [panic](https://go.dev/blog/defer-panic-and-recover). Instead, in my example, it completely overlooks the configuration issue caused by the invalid origin value and returns a middleware that happens to be dysfunctional because that middleware rejects _all_ CORS requests, even those whose origin is `https://example.com`. Incidentally, I find this practice of ignoring failure cases—or sweeping them under the rug—especially puzzling when followed by Go libraries, because diligent error reporting is so central to Go’s culture.

Users of CORS libraries that lack configuration validation may find out that nothing works as they expected only much later, e.g. after having deployed their server and testing their client against it; and, even then, those developers may be at a loss to pinpoint the root cause of the dysfunction they observe. True, browsers (Chromium, in this case) do, in such cases, provide an illuminating error message:

> Access to fetch at [REDACTED] from origin ‘[https://example.com](https://example.com/)’ has been blocked by CORS policy: Response to preflight request doesn’t pass access control check: The ‘Access-Control-Allow-Origin’ header has a value ‘[https://example.com/'](https://example.com/%27) that is not equal to the supplied origin. Have the server send the header with a valid value, or, if an opaque response serves your needs, set the request’s mode to ‘no-cors’ to fetch the resource with CORS disabled.

But this rather loose feedback loop is detrimental to developer experience. Immediately presenting developers with the brutal truth, in the form of a helpful _server-side_ error message, would be preferable to delaying their disappointment. Accordingly, CORS middleware libraries should categorically refuse to produce a middleware bound to be dysfunctional and should error out as early as middleware instantiation.

Here is another example of a a general lack of configuration validation. Some CORS libraries, like [rs/cors](https://pkg.go.dev/github.com/rs/cors), allow their users to configure the status code that preflight responses should use, but few of them check that the status code in question is an [_ok status_](https://fetch.spec.whatwg.org/#ok-status) (i.e. an integer in the 200-299 range), which is [a necessary condition for preflight to succeed](https://fetch.spec.whatwg.org/#cors-preflight-fetch-0):

```plain text
c := cors.New(cors.Options{
  AllowedOrigins: []string{"https://example.com"},
  OptionsSuccessStatus: 300, // incorrect
})

```

Still unconvinced that lack of configuration validation is problematic? Here is another example featuring the [Spring framework](https://spring.io/):

```plain text
@Override
public void addCorsMappings(CorsRegistry corsRegistry) {
  corsRegistry.addMapping("/**")
    .allowedOrigins("*")
    .allowedHeaders("Content-Type,Authorization"); // incorrect
}

```

This Java code snippet raises no exception, but the resulting CORS middleware does _not_ allow request headers `Content-Type` and `Authorization`. Why not? Because [Spring’s ](https://github.com/spring-projects/spring-framework/blob/64de6de725dcbee0ea9b14ee4ab6c264f32421ee/spring-webflux/src/main/java/org/springframework/web/reactive/config/CorsRegistration.java#L97)[`allowedHeaders`](https://github.com/spring-projects/spring-framework/blob/64de6de725dcbee0ea9b14ee4ab6c264f32421ee/spring-webflux/src/main/java/org/springframework/web/reactive/config/CorsRegistration.java#L97)[ method](https://github.com/spring-projects/spring-framework/blob/64de6de725dcbee0ea9b14ee4ab6c264f32421ee/spring-webflux/src/main/java/org/springframework/web/reactive/config/CorsRegistration.java#L97) is variadic and expects their users to specify their allowed request-header names, not as a single string value, but one by one as separate arguments:

```plain text
@Override
public void addCorsMappings(CorsRegistry corsRegistry) {
  corsRegistry.addMapping("/**")
    .allowedOrigins("*")
    .allowedHeaders("Content-Type", "Authorization"); // correct
}

```

If Spring instead failed fast in the event of an invalid user-specified request-header name—and `Content-Type,Authorization` most definitely is not a [valid request-header name](https://datatracker.ietf.org/doc/html/rfc7230#section-3.2)—the framework would spare its users [much frustration](https://stackoverflow.com/questions/75022849/webflux-cors-no-access-control-allow-origin/75025792).

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In stark contrast with those libraries, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) adopts a decidedly [defensive](https://en.wikipedia.org/wiki/Defensive_programming) approach: it steadfastly validates CORS configuration (origins, headers, methods, etc.) and errors out at middleware instantiation rather than produce a dysfunctional middleware.

Moreover, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) guides developers unfamiliar with CORS towards a cruft-free configuration, by aggressively rejecting [request headers](https://fetch.spec.whatwg.org/#forbidden-request-header), [response headers](https://fetch.spec.whatwg.org/#forbidden-response-header-name), and [method names](https://fetch.spec.whatwg.org/#forbidden-method) that the Fetch standard _forbids_:

```plain text
cors, err := fcors.AllowAccess(
  fcors.FromAnyOrigin(),
  fcors.WithMethods("CONNECT"),              // ❌
  fcors.WithRequestHeaders("Cookie"),        // ❌
  fcors.ExposeResponseHeaders("Set-Cookie"), // ❌
)

```

As a result of its strict stance on configuration validation, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) enables a tighter development feedback loop than its competition does. The next principle expands on this idea, but in the specific context of the CORS protocol’s so-called _wildcard exception_.

### 6. Treat CORS as a compilation target

One peculiarity of the CORS protocol that trips many Web developers up is a deliberate constraint colloquially known as the _wildcard exception_. The Fetch standard remains the authoritative (if somewhat dry) source of truth about CORS, but [MDN Web Docs does a good job at demystifying the wildcard exception](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS#credentialed_requests_and_wildcards). In a nutshell, credentialed requests (e.g. requests that include cookies) are incompatible with uses of the wildcard (`*`) in any of the following CORS response headers:

- `Access-Control-Allow-Origin`
- `Access-Control-Allow-Methods`
- `Access-Control-Allow-Headers`
- `Access-Control-Expose-Headers`

Browsers simply deny clients access to such responses and raise an error, either during preflight (when applicable) or after receiving the response to the actual request from the server.

The wildcard exception indisputably is a beneficial constraint: it plays an important role in preventing incautious server developers from allowing credentialed access from more Web origins than they ought to, and, therefore, in protecting stakeholders from [cross-origin attacks meant to exfiltrate their sensitive data](https://portswigger.net/research/exploiting-cors-misconfigurations-for-bitcoins-and-bounties).

You would expect a well-designed CORS middleware library to prohibit dysfunctional CORS configurations that disregard the wildcard exception. Unfortunately, many libraries don’t go through that trouble. Here is an example featuring [the CORS middleware library of Express.js](https://expressjs.com/en/resources/middleware/cors.html):

```plain text
const express = require('express')
const cors = require('cors')
const app = express()
const port = 8081

const corsOptions = {
  origin: '*',
  credentials: true,
}

app.get('/hello', cors(corsOptions), (req, res) => {
  res.send('Hello, World!')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})

```

Sending a `GET` CORS request to `/hello` yields a response that contains the following headers:

```plain text
Access-Control-Allow-Origin: *
Access-Control-Allow-Credentials: true

```

As a result, browsers will allow anonymous access from all origins, but won’t allow credentialed access precisely because of the wildcard exception. Clients who send credentialed `GET` requests to `/hello` will indeed be confronted with a CORS error of this kind:

> Access to fetch at [REDACTED] from origin [REDACTED] has been blocked by CORS policy: The value of the ‘Access-Control-Allow-Origin’ header in the response must not be the wildcard ‘*’ when the request’s credentials mode is ‘include’.

You can imagine how this error message can elicit puzzlement and frustration among developers unfamiliar with the wildcard exception. In fact, seldom a day goes by without some distraught developers [asking for help about the wildcard exception on Stack Overflow](https://stackoverflow.com/q/19743396/2541573). Not a great developer experience…

A well-designed CORS middleware library would [fail fast](https://ieeexplore.ieee.org/stamp/stamp.jsp?arnumber=1331296) and would not allow server developers to proceed with such a dysfunctional CORS configuration as one that ignores the wildcard exception.

But developers’ misfortunes do not end there, because some CORS middleware libraries make even more regrettable design decisions. Well aware of the wildcard exception, those libraries bend over backwards to spare their users a dysfunctional CORS middleware and give them an insecure one instead! Here is an example featuring the [Echo framework](https://echo.labstack.com/)’s [CORS middleware](https://echo.labstack.com/middleware/cors/):

```plain text
package main

import (
  "net/http"

  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {
  e := echo.New()
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins:     []string{"*"},
    AllowCredentials: true,
  }))
  e.GET("/hello", hello)
  e.Logger.Fatal(e.Start(":8081"))
}

func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

```

Sending a `GET` CORS request to `/hello` from Web origin `https://attacker.com` yields a response that contains the following headers:

```plain text
Access-Control-Allow-Origin: https://attacker.com
Access-Control-Allow-Credentials: true

```

😱 Yikes! Echo’s CORS middleware actively bypasses the wildcard exception by unconditionally reflecting the request’s origin in the `Access-Control-Allow-Origin` header, thereby leaving the door wide open to [cross-origin attacks](https://portswigger.net/research/exploiting-cors-misconfigurations-for-bitcoins-and-bounties)! This regrettable design decision has plagued and still plagues many CORS middleware libraries, to the point that such conscientious security researchers as [Evan Johnson](https://github.com/go-chi/cors/pull/6) and [Jianjun Chen](https://github.com/rs/cors/issues/55) felt compelled to publicly call some of them out, with varying degrees of success.

By contrast with the great majority of CORS middleware libraries, [jubobs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) sidesteps the wildcard-exception difficulty altogether. In the spirit of [John Ousterhout’s philosophy of software design](https://web.stanford.edu/~ouster/cgi-bin/aposd.php), [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) _pulls complexity downwards_: it abstracts the complexity associated with the wildcard exception by treating CORS as a lower-level compilation target—an idea reminiscent of one that Mike West once floated about [Content Security Policy](https://www.w3.org/TR/CSP3/):

> : … perhaps we can start looking at CSP as a (complicated, low-level) compilation target. @hillbrad @randomdross

[October 5, 2016](https://twitter.com/mikewest/status/783755021645217792?ref_src=twsrc%5Etfw)

[jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) indeed prevents dysfunctional configurations by prohibiting explicit use of the wildcard (`"*"`) and forcing users to describe their desired CORS policy in a higher-level, more declarative fashion instead. For instance, the following code snippet raises a run-time error:

Developers who wish to allow anonymous access from any origin should instead write

```plain text
fcors.AllowAccess(
  fcors.FromAnyOrigin(), // ✅
)

```

The library does use the wildcard under the hood when possible, but affords developers blissful oblivion to this implementation detail.

Besides, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) embraces [Yaron Minsky](https://twitter.com/yminsky)’s [principle of ](https://blog.janestreet.com/effective-ml-revisited/#make-illegal-states-unrepresentable)[_making illegal states unrepresentable_](https://blog.janestreet.com/effective-ml-revisited/#make-illegal-states-unrepresentable): far from insecurely bypassing the wildcard exception, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) leverages Go’s type system in order to prevent (at compile time!) server developers from allowing credentialed access from all origins.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 7. Provide no default configuration

Most CORS middleware libraries provide some default configuration. For instance, [rs/cors](https://pkg.go.dev/github.com/rs/cors) [exports the following function](https://github.com/rs/cors/blob/fcebdb403f4d4585c705318c0e4d6d05a761a4ab/cors.go#L195), which returns a “default” CORS middleware:

```plain text
// Default creates a new Cors handler with default options.
func Default() *Cors

```

As a prospective user of this function, you may find its documentation cryptic and ask yourself the following questions:

- What exactly are those “default options”?
- What behaviour does the resulting CORS middleware exhibit?
- Is the resulting CORS middleware suited for my needs?

To find definite answers, you’d have no other choice but to dig into the library’s source code.

Lacking documentation is one thing, but I want to draw your attention to a deeper issue. On the one hand, and contrary to popular belief, activating CORS never _strengthens_ security but always _weakens_ security (to varying degrees).

> That so many people routinely conflate the Same-Origin Policy and CORS annoys me to no end... No metaphor is perfect but, if the former is akin to a seatbelt, the latter is the button that unbuckles that seatbelt. #infosec

[July 5, 2022](https://twitter.com/jub0bs/status/1544366638446907394?ref_src=twsrc%5Etfw)

Some widespread CORS configurations may strike you as innocuous. For instance, you may perceive a blanket `Access-Control-Allow-Origin: *` policy as a harmless default, but even such an innocent-looking CORS policy [can prove insecure in some contexts](https://security.stackexchange.com/a/254684/127436). On the other hand, configuration in general should be [secure by default](https://en.wikipedia.org/wiki/Secure_by_default).

From this apparent conflict, I see but one escape: the only sane default consists in **not enabling CORS at all**. Thefore, my stance is that a well-factored CORS middleware library should _not_ provide any default configuration. instead, it should, at the cost of some verbosity, force users to be deliberate about their CORS configuration.

Accordingly, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) provides no default configuration. For instance, developers who wish to allow anonymous access from any origin must be deliberate and explicit about it in their CORS configuration:

```plain text
cors, err := fcors.AllowAccess(
  fcors.FromAnyOrigin(),
)

```

### 8. Do not preclude legitimate configurations

Some CORS middleware libraries are rife with [false affordances](https://en.wikipedia.org/wiki/Affordance#False_affordances), i.e. things that look like they should be possible to do but are actually not. Here are two examples.

To reduce traffic, browsers feature an internal cache dedicated to preflight responses. Developers can tune the expiration time of an individual preflight response from the preflight cache by including an [`Access-Control-Max-Age`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age)[ header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age) in that response and specifying the desired _max age_ (in seconds) as a nonnegative integer:

```plain text
Access-Control-Max-Age: 30

```

Omitting the `Access-Control-Max-Age` header altogether causes browsers to cache the prelight response in question for a default duration of five seconds, whereas specifying a max-age value of `0` instruct browsers not to cache the preflight response at all:

```plain text
Access-Control-Max-Age: 0

```

Unfortunately, [rs/cors](https://pkg.go.dev/github.com/rs/cors), among other libraries, [ignores this distinction](https://github.com/rs/cors/blob/fcebdb403f4d4585c705318c0e4d6d05a761a4ab/cors.go#L333) and, therefore, altogether prevents its users from disabling preflight caching, should they wish to do so.

> The "functional options" pattern is superior to a config struct in many ways, but one of its benefits is that it allows package authors to draw a distinction between zero value and default behaviour. #golang

[October 18, 2022](https://twitter.com/jub0bs/status/1582296106246885378?ref_src=twsrc%5Etfw)

> A telling example is CORS's Acess-Control-Max-Age response header:- Setting it with a value of 0 (int's zero value) leads to no caching of the preflight response.- Omitting it altogether leads browsers to cache the preflight with a default value of 5 seconds.

[October 18, 2022](https://twitter.com/jub0bs/status/1582296109502042112?ref_src=twsrc%5Etfw)

> Yet most #golang CORS middleware that use a config struct rather than functional options interpret a 0 max-age value as a cue to omit that header in the preflight response. 🤷

[October 18, 2022](https://twitter.com/jub0bs/status/1582296111640764417?ref_src=twsrc%5Etfw)

[jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) doesn’t suffer from this limitation:

```plain text
cors, err := fcors.AllowAccess(
  fcors.FromAnyOrigin(),
  fcors.WithRequestHeaders("Authorization"),
  fcors.WithMaxAgeInSeconds(0), // disables preflight caching
)

```

Several CORS middleware libraries normalise HTTP-method names to uppercase. Those libraries, either inadvertently or deliberately, neglect the fact that, [according to the Fetch standard](https://fetch.spec.whatwg.org/#methods), method names are case-sensitive; this illustrative [instance of Jake Archibald’s CORS playground](https://jakearchibald.com/2021/cors/playground/?prefillForm=1&requestMethod=patch&requestUseCORS=1&preflightStatus=204&preflightAllowOrigin=https%3A%2F%2Fjakearchibald.com&preflightAllowMethods=PATCH&responseStatus=200&responseAllowOrigin=https%3A%2F%2Fjakearchibald.com) should be enough to convince you.

Such unwarranted case normalisation causes problems for clients that send requests whose method is not uppercase—and not some case-insensitive match for one of `DELETE`, `GET`, `HEAD`, `OPTIONS`, `POST`, or `PUT`, [names for which the Fetch standard carves out an exception](https://fetch.spec.whatwg.org/#concept-method-normalize). Here is an example featuring [rs/cors](https://pkg.go.dev/github.com/rs/cors):

```plain text
c := cors.New(cors.Options{
  AllowedOrigins: []string{"https://example.com"},
  AllowedMethods: []string{"patch"},
})

```

Assume then that a client running in the context of Web origin `https://example.com` in a Fetch-compliant browser sends a `patch` request (note the lower case) to the server. Because [rs/cors](https://pkg.go.dev/github.com/rs/cors) [internally normalises method names to uppercase](https://github.com/rs/cors/blob/fcebdb403f4d4585c705318c0e4d6d05a761a4ab/cors.go#L320), the prelight response would contain

```plain text
Access-Control-Allow-Methods: PATCH

```

as opposed to

```plain text
Access-Control-Allow-Methods: patch

```

Because method names are case-sensitive, browsers rule this as a mismatch and fail CORS preflight with an error message of this kind:

> 

Access to fetch at [REDACTED] from origin ‘[https://example.com](https://example.com/)’ has been blocked by CORS policy: Method patch is not allowed by Access-Control-Allow-Methods in preflight response.

Very confusing! By contrast, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) respects the case sensitivity of HTTP methods and therefore never gives rise to this kind of problem.

### 9. Ease troubleshooting by eschewing shortcuts during preflight

CORS is configured on the _server_ side, but applied by the _browser_ on the _client_ side. Because the CORS procotol involves two actors (three if you count the browser), figuring out how to fix a specific CORS issue remains challenging. Despite recent and commendable efforts by the Chromium team in that area (see [Jecelyn Yeen](https://twitter.com/jecfish)’s [DevTools State of the Union 2022](https://www.youtube.com/watch?v=YqWEqYa-evk&t=564s)), troubleshooting CORS issues can still slow down Web development to a snail’s pace.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

In practice, a dysfunctional CORS configuration manifests itself as an error in the browser’s Console tab, much to the chagrin of developers who often perceive those error messages as puzzling and unhelpful:

> CORS is probably one of the most frustrating things I've worked with. It would be nice if it gave you an error message on the console.

[July 22, 2012](https://twitter.com/mcolyer/status/227163952608575488?ref_src=twsrc%5Etfw)

In my opinion, however, this bad reputation is undeserved: most browsers, such as [Firefox](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS/Errors), indeed do boast a rich variety of informative CORS error messages. Unfortunately, most CORS middleware libraries are implemented in such a way that they give rise to only a small subset of those error messages in the browser, thereby masking the root cause of many CORS issues. For instance, a CORS error message that developers typically have to face is the following:

> 

Access to fetch at [REDACTED] from origin [REDACTED] has been blocked by CORS policy: Response to preflight request doesn’t pass access control check: No ‘Access-Control-Allow-Origin’ header is present on the requested resource.

This error message is so common that the last sentence is contained verbatim in [roughly 20% of the questions tagged with “cors” on Stack Overflow](https://stackoverflow.com/search?q=%22No%20'Access-Control-Allow-Origin'%20header%20is%20present%20on%20the%20requested%20resource.%22%20%5Bcors%5D). In many cases, though, the issue stems, _not_ from the origin of the request, but [from a ](https://stackoverflow.com/questions/71181738/no-access-control-allow-origin-header-despite-specifying-allowedorigins)[_request header_](https://stackoverflow.com/questions/71181738/no-access-control-allow-origin-header-despite-specifying-allowedorigins)[ that happens to be disallowed by the server’s CORS policy](https://stackoverflow.com/questions/71181738/no-access-control-allow-origin-header-despite-specifying-allowedorigins). The difficulty in troubleshooting this kind of issues is compounded by clients libraries like [Axios](https://axios-http.com/) which, in some cases, [silently include headers, such as ](https://github.com/axios/axios/blob/d83316db4a242252db3a2ed7728cb43f8f8f4189/lib/defaults/index.js#L99)[`Content-Type`](https://github.com/axios/axios/blob/d83316db4a242252db3a2ed7728cb43f8f8f4189/lib/defaults/index.js#L99)[, to requests](https://github.com/axios/axios/blob/d83316db4a242252db3a2ed7728cb43f8f8f4189/lib/defaults/index.js#L99).

To be absolutely fair, I must concede that the blame for this sad state of affairs likely rests, not with CORS middleware libraries themselves, but with the original CORS specification (since obsoleted by the [Fetch standard](https://fetch.spec.whatwg.org/)), namely the [W3C Cross-Origin Resource Sharing working draft](https://www.w3.org/TR/2014/REC-cors-20140116/). The [version dated 17 March 2009](https://www.w3.org/TR/2009/WD-cors-20090317/) (☘️) of that working draft saw the addition of a section entitled _Server Processing Model_, which, among other things, prescribes how servers should handle preflight requests. Among those normative requirements, here is the passage (only slightly redacted) that is most relevant to the present discussion:

> If [the value of the `Access-Control-Request-Method` header] is not a case-sensitive match for any of [the allowed methods], **do not set any additional headers and terminate this set of steps**. […]

1. If any of the [header names listed in the `Access-Control-Request-Headers` header] is not a ASCII case-insensitive match for any of the [allowed headers], **do not set any additional headers and terminate this set of steps**. […]
2. If the URL supports credentials, add a single `Access-Control-Allow-Origin` header, with the value of the `Origin` header as value, and add a single `Access-Control-Allow-Credentials` header with the case-sensitive string “`true`” as value. Otherwise, add a single `Access-Control-Allow-Origin` header, with either the value of the `Origin` header or the string “” as value.

In essence, the W3C working draft prescribes servers take shorcuts during preflight: servers ought to omit all CORS headers (`Access-Control-Allow-Origin` too!) from the response if preflight fails. This leaves very little contextual information to the browser, which can only complain about the absence of the `Access-Control-Allow-Origin` header—because that header happens to be the first CORS header to get processed during a [CORS check](https://fetch.spec.whatwg.org/#concept-cors-check).

Here is an example. Assume that Sarah configures CORS on her server (`https://sarah.com`) to allow Carol’s client (`https://carol.com`) with a request header named `Foo`. In her request to Sarah’s server, Carol happens to include a header named `Bar`, which Sarah’s CORS policy disallows:

```plain text
const headers = {
  'Foo': 'whatever',
  'Bar': 'yolo',
};
fetch('//sarah.com', {headers: headers})
  .then(res => res.text())
  .then(console.log);

```

In the resulting preflight request, Sarah’s CORS middleware sees a header it disallows:

```plain text
OPTIONS / HTTP/1.1
Host: sarah.com
Origin: https://carol.com
Access-Control-Request-Method: GET
Access-Control-Request-Headers: bar,foo
-snip-

```

And because Sarah’s CORS middleware follows the W3C’s normative requirements for servers, it replies with a 403 status and omits _all_ CORS headers from the preflight response!

```plain text
HTTP/1.1 403 Forbidden
-snip-

```

Even though Sarah does allow the Web origin of Carol’s client in her CORS policy, when preflight fails, the browser presents Carol with an infuriatingly confusing CORS error message:

> Access to fetch at ‘[https://sarah.com](https://sarah.com/)’ from origin ‘[https://carol.com](https://carol.com/)’ has been blocked by CORS policy: Response to preflight request doesn’t pass access control check: No ‘Access-Control-Allow-Origin’ header is present on the requested resource.

This misleading message hides the root cause of the problem: Carol’s use of a request-header name (`Bar`) that Sarah’s CORS policy happens to disallow.

The W3C’s normative requirements for servers endured subsequent versions of the working draft more or less unchanged; and, because initial development of many CORS middleware libraries was contemporaneous with the W3C working draft’s heyday, most of those libraries—to the notable exception of [Express.js’s](https://expressjs.com/en/resources/middleware/cors.html)—diligently complied and never looked back.

In my view, **those problematic normative requirements for servers are the main reason why so many server developers have been and still are routinely led astray by CORS error messages.**

In this context, taking such shortcuts and _failing fast_ is actually detrimental, because it often robs developers of a good explanation as to what caused their CORS issues.

In my earlier example, if Sarah’s middleware instead responded to Carol’s preflight request with

```plain text
HTTP/1.1 200 OK
Access-Control-Allow-Origin: https://carol.com
Access-Control-Allow-Headers: foo

```

then the browser would still fail preflight but would present Carol with a much more actionable CORS error message:

> Access to fetch at ‘[https://sarah.com](https://sarah.com/)’ from origin ‘[https://carol.com](https://carol.com/)’ has been blocked by CORS policy: Request header field bar is not allowed by Access-Control-Allow-Headers in preflight response.

Carol would then immediately realise that she needs to either drop the `Bar` header from his request, or kindly ask Sarah to also allow that request header in her CORS policy.

Rather than freeing themselves from the W3C’s injunctions and getting rid of shortcuts during preflight, some libraries, such as [rs/cors](https://pkg.go.dev/github.com/rs/cors), [elected to add support for logging in a bid to ease troubleshooting](https://github.com/rs/cors/pull/3). However, I find such a decision less defensible in light of the changes brought by the Fetch standard. By comparison with the W3C working draft, the Fetch standard is indeed [far less prescriptive](https://fetch.spec.whatwg.org/#http-responses) about how servers ought to handle preflight requests:

> Ultimately server developers have a lot of freedom in how they handle HTTP responses and these tactics can differ between the response to the CORS-preflight request and the CORS request that follows it […]

Sometimes you feel like throwing your hands up in the air, but [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) [has got the love you need to see you through](https://www.youtube.com/watch?v=PQZhN65vq9E)! Because my library altogether eschews shortcuts during preflight, you should find CORS issues much easier to troubleshooting with it than with other libraries.

Interestingy, getting rid of such shortcuts [was once suggested for](https://github.com/rs/cors/issues/54#issue-317389430) but never implemented in [rs/cors](https://pkg.go.dev/github.com/rs/cors). A missed opportunity, if you ask me.

### 10. Render insecure configurations impossible

[Earlier in this post](https://jub0bs.com/posts/2023-02-08-fearless-cors/#validate-configuration-and-fail-fast), I stressed that CORS libraries should not produce dysfunctional middleware in the event of a misconfiguration. There’s a flip side to this. One of [Joshua Bloch](https://twitter.com/joshbloch)’s design principles that has stayed with me [ever since I encountered it](https://www.youtube.com/watch?v=aAb7hSCtvGw&t=6m02s) is the following:

> A good library is one that is not only easy to use but _hard to misuse_.

The second part of this maxim is especially relevant when the library in question is concerned with such security-critical mechanisms as CORS.

I hope you will forgive this somewhat paternalising metaphor: a seat belt should arguably be designed in such a way that no unattended toddler can unbuckle it; similarly, a good CORS library should prevent developers from producing dangerously permissive CORS middleware.

Unfortunately, too many CORS middleware libraries fail to effectively protect developers from themselves because they insufficiently rein in their power.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

I’ve already touched upon this topic in earlier in this post (see [Provide support for Private Network Access](https://jub0bs.com/posts/2023-02-08-fearless-cors/#3-provide-support-for-private-network-access) and [Treat CORS as a compilation target](https://jub0bs.com/posts/2023-02-08-fearless-cors/#6-treat-cors-as-a-compilation-target)), but here are more ways in which CORS middleware libraries should prevent their users from shooting themselves in the foot.

### Do not support the origin

Security researcher [James Kettle](https://twitter.com/albinowax) warns us (and [rightly so](https://www.cynet.com/wp-content/uploads/2016/12/Blog-Post-BugSec-Cynet-Facebook-Originull.pdf)!) that allowing the `null` origin in one’s CORS configuration [is rarely (if ever) a good idea](https://portswigger.net/research/exploiting-cors-misconfigurations-for-bitcoins-and-bounties); the only justifiable use case I can think of is a deliberately vulnerable resource in the context of [a Web-security lab about CORS misconfiguration](https://portswigger.net/web-security/cors/lab-null-origin-whitelisted-attack). Yet CORS libraries, by and large, are quite content to tolerate the `null` origin.

By contrast, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) completely forbids the `null` origin; attempts to allow it result in a run-time error:

Another, perhaps more subtle, CORS misconfiguration consists in allowing _insecure origins_, e.g. origins whose scheme is `http` as opposed to `https`, such as `http://example.com`. As [demonstrated by James Kettle](https://portswigger.net/research/exploiting-cors-misconfigurations-for-bitcoins-and-bounties), doing so enables an [active network attacker](https://en.wikipedia.org/wiki/Man-in-the-middle_attack) to steal potentially sensitive data from its victim, even if the latter only ever willingly interacts with the vulnerable resource over a secure connection (i.e. using HTTPS). Unfortunately, none of the CORS libraries I’ve reviewed guard their users against this pitfall. Even if you’re aware of it, you are a mere one-character typo away from misconfiguring your CORS middleware, and no such tiny typo should be this consequential.

Again, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) has your back: and disallows _most_ `http` origins by default; it does carve out an exception for origins whose host is `localhost` or a [loopback IP address](https://en.wikipedia.org/wiki/Loopback#Virtual_loopback_interface), though, because such origins are typically harmless and so commonly used in pre-prod environments. By default, attempts to allow one or more insecure origins result in a run-time error:

```plain text
cors, err := fcors.AllowAccess(
  fcors.FromOrigins("http://example.com"), // ❌
)

```

If you need to _deliberately_ allow one or more insecure origins, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) lets you do so under the condition that you also activate an advanced option provided by [its companion package, aptly named “risky”](https://pkg.go.dev/github.com/jub0bs/fcors/risky):

```plain text
cors, err := fcors.AllowAccess(
  fcors.FromOrigins("http://example.com"), // ✅
  risky.TolerateInsecureOrigins(),
)

```

> Beyond secure defaults, a piece of software should ideally be designed in such a way that an insecure configuration always require more keystrokes than a secure configuration. Looking at you, CORS middleware libraries... 😇 #infosec

[February 5, 2022](https://twitter.com/jub0bs/status/1489995943722332160?ref_src=twsrc%5Etfw)

It is common among CORS libraries, such as [Express.js’s](https://expressjs.com/en/resources/middleware/cors.html), to support regular expressions as a means to allow a set of origins, e.g. [all subdomains of some base domain](https://github.com/expressjs/cors/issues/42):

```plain text
/^https:\/\/.*\.example\.com$/

```

This feature affords a great deal of flexibility… perhaps too much. Security misconfigurations due to flawed regexps are indeed so common that it puts the judiciousness of such a feature into question. Even seasoned regexp wranglers can sometimes be faulted, e.g. when they forget to [escape periods standing for label separators](https://jub0bs.com/posts/2022-08-04-scraping-the-bottom-of-the-cors-barrel-part1/#unescaped-periods-in-the-etld1-part-of-a-regexp) or to anchor their regexp at both ends:

```plain text
/https:\/\/.*.example.com/
/^https:\/\/.*\.example\.com/

```

Besides, some regular-expression engines—though thankfully [not Go’s](https://pkg.go.dev/regexp)—are prone to catastrophic backtracking, which, if exploited by an attacker on a vulnerable server, can cause a [denial of service](https://owasp.org/www-community/attacks/Regular_expression_Denial_of_Service_-_ReDoS). Therefore, I don’t think CORS libraries should support regexps. [Rob Pike likely would agree](https://commandcenter.blogspot.com/2011/08/regular-expressions-in-lexing-and.html):

> 

Regular expressions are, in my experience, widely misunderstood and abused.

Some CORS middleware libraries, like [rs/cors](https://pkg.go.dev/github.com/rs/cors), deserve praise for [refusing to support regexps](https://github.com/rs/cors/issues/131#issuecomment-1203882405) and supporting only a limited variety of pattern types, but most of them still stop short of guarding their users against overly permissive patterns. For instance, [rs/cors](https://pkg.go.dev/github.com/rs/cors) tolerates `https://*` as an origin pattern, which is insecure because it matches _any_ `https` origin, including, say, `https://attacker.com`!

[jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) is adamant in restricting the types of origin patterns it supports. Some examples follow:

- `https://*.example.com`, in which the asterisk marks exactly one DNS label;
- `https://**.example.com`, in which the two asterisks mark one or more DNS labels;
- `https://example.com:*`, in which the asterisk marks an arbitrary (possibly implicit) port number.

[jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) supports no other types of origin patterns. In particular, all of the following patterns are illegal and cause middleware instantiation to fail: `https://*`, `https://example.*`, `https://*.example.com:*`.

[jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) even goes the extra mile for your safety: by default, it prevents developers from inadvertently allowing arbitrary subdomains of a [public suffix](https://publicsuffix.org/) (e.g. `com`), because such domains (e.g. `random-attacker-666.com`) are by definition registrable by _anyone_, including malicious actors. By default, attempts to allow arbitrary subdomains of a public suffix result in a run-time error:

Again, if you need to _deliberately_ allow arbitrary subdomains of one or more public suffixes, you’ll need to explicitly open another escape hatch provided by the [risky package](https://pkg.go.dev/github.com/jub0bs/fcors/risky):

```plain text
cors, err := fcors.AllowAccess(
  fcors.FromOrigins("https://*.com"), // ✅
  risky.SkipPublicSuffixCheck(),
)

```

I’ve kept the most dangerous misfeature of them all for last: custom callbacks. They come in numerous CORS libraries under different names and shapes:

- Express.js’s CORS configuration object [accepts a callback (or even a boolean!) for its ](https://expressjs.com/en/resources/middleware/cors.html#configuring-cors-w-dynamic-origin)[`origin`](https://expressjs.com/en/resources/middleware/cors.html#configuring-cors-w-dynamic-origin)[ property](https://expressjs.com/en/resources/middleware/cors.html#configuring-cors-w-dynamic-origin);
- [rs/cors](https://pkg.go.dev/github.com/rs/cors)’s configuration struct exports a function field named `AllowOriginFunc` of signature `func(string) bool`;
- Echo’s CORS middleware’s configuration struct exports a function field also named `AllowOriginFunc` but of signature `func(string) (bool, error)`;

By letting their users specify custom callbacks, those CORS middleware libraries give their users near total flexibility for specifying how CORS requests should be accepted or rejected, e.g. by querying some database for the list of allowed origins, as [suggested by the documentation of Express.js’s CORS middleware](https://expressjs.com/en/resources/middleware/cors.html#configuring-cors-w-dynamic-origin):

```plain text
var corsOptions = {
  origin: function (origin, callback) {
    // db.loadOrigins is an example call to load
    // a list of origins from a backing database
    db.loadOrigins(function (error, origins) {
      callback(error, origins)
    })
  }
}

```

This “DIY” approach to CORS often leads to frightful results, precisely because custom callbacks can so easily be abused by developers. For example, I’ve seen many developers specify a [predicate that unconditionally returns ](https://github.com/rs/cors/issues/80)[`true`](https://github.com/rs/cors/issues/80) while also allowing credentialed requests, thereby unwittingly opening the door to authenticated cross-origin attacks from _any_ Web origin; here is an example using [rs/cors](https://pkg.go.dev/github.com/rs/cors):

```plain text
c := cors.New(cors.Options{
  AllowOriginFunc: func (origin string) bool { return true }, // 😱
  AllowCredentials: true,
})

```

There’s an even more pernicious form of custom callbacks, present in both [rs/cors](https://pkg.go.dev/github.com/rs/cors) and [Chi](https://github.com/go-chi/chi) (yet another Web framework for Go), which grants the predicate access to the request object:

```plain text
func(*http.Request, string) bool

```

Motivations for this feature include the desire to vary responses on the basis of the presence and/or value of some request headers, such as `Cookie`, `Authorization`, or some non-standard HTTP header. Do you see anything amiss? The answer is quite subtle…

Perhaps the most glaring issue with that predicate’s signature is that, since the request’s origin is accessible via the [`*http.Request`](https://pkg.go.dev/net/http#Request) parameter anyway, the string parameter is redundant; another manifestation of accidental complexity right there.

However, one much more severe issue with such a predicate is that it’s almost guaranteed to lead to [cache poisoning](https://owasp.org/www-community/attacks/Cache_Poisoning)! If you vary responses on the basis of the presence and/or value of some request header, HTTP indeed mandates that you ist the name of that header in your responses' [`Vary`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary)[ header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary); doing this signals to caching intermediaries that the header in question should be part of the cache key. But because the predicate lacks a [`http.ResponseWriter`](https://pkg.go.dev/net/http#ResponseWriter) parameter, it gives you no way of populating the `Vary` header as required! You have to remember to manually do so _outside_ of the predicate. If you forget, as I believe most developers do, to adequately populate the `Vary` header, caching intermediaries are then liable to serve inappropriate (and possibly malicious) cached responses to your clients.

The examples above may be sufficient to convince you that CORS middleware libraries are better off staying away from those custom-callback features. In my experience, most use cases for CORS don’t require this level of customisation anyway.

By contrast, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) curtails developers' power (for their own good) and eschew all forms of custom callbacks. Moreover, as discussed [earlier in this post](https://jub0bs.com/posts/2023-02-08-fearless-cors/#7-treat-cors-as-a-compilation-target), its [`FromAnyOrigin`](https://pkg.go.dev/github.com/jub0bs/fcors#FromAnyOrigin)[ option](https://pkg.go.dev/github.com/jub0bs/fcors#FromAnyOrigin) is rendered incompatible (at compile time!) with credentialed access.

### 11. Guarantee configuration immutability

Some CORS middleware libraries support dynamic updates to CORS configuration, thereby obviating the need for a server deployment. Here is a proof of concept featuring Express.js:

```plain text
const express = require('express')
const cors = require('cors')
const app = express()
const port = 3000

const corsOptions = {
  origin: 'http://example.com',
}
app.use(cors(corsOptions));

app.get('/', (req, res) => {
  res.send('Hello World!')
})

app.post('/change-allowed-origin', function (req, res, next) {
  corsOptions.origin = 'http://example.org';
  res.send('Done!')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})

```

Once that server has been started, sending a single `POST` request to `/change-allowed-origin` will indeed cause an update of the server’s CORS configuration: the allowed origin, which beforehand was `https://example.com` will become `https://example.org`.

```plain text
$ curl -sD - -o /dev/null \
  -H "Origin: https://example.org" \
  localhost:3000

```

```plain text
HTTP/1.1 200 OK
X-Powered-By: Express
Access-Control-Allow-Origin: http://example.com
Vary: Origin
-snip-

```

```plain text
$ curl -XPOST localhost:3000/change-allowed-origin
Done!

```

```plain text
$ curl -sD - -o /dev/null \
  -H "Origin: https://example.org" \
  localhost:3000

```

```plain text
HTTP/1.1 200 OK
X-Powered-By: Express
Access-Control-Allow-Origin: http://example.org
Vary: Origin
-snip-

```

Whether this [affordance](https://en.wikipedia.org/wiki/Affordance) is intentional or incidental (and due to a lack of defensive copying) is unclear to me. Though expedient and [sometimes desirable](https://stackoverflow.com/q/75103900/2541573), it is arguably more harmful than good.

Insofar as CORS relaxes some of the restrictions enforced by the SOP, it is a _security-critical_ mechanism. Accordingly, a server’s CORS configuration should be subject to [_change control_](https://en.wikipedia.org/wiki/Change_control): more specifically, I argue that any update to CORS configuration warrants careful vetting and should be followed by a server restart. [Custom callbacks](https://jub0bs.com/posts/2023-02-08-fearless-cors/#10-render-insecure-configurations-impossible) and insufficient defensive copying stand in the way of this principle. Besides, if developers perceive server startup as prohibitively slow, they should, in my opinion, focus their efforts on reducing startup time rather than on avoiding the need to restart the server.

A middleware built with [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) is effectively immutable, and any amendment to the corresponding CORS policy requires a server restart. Although this constraint does not guarantee that CORS-policy changes will get reviewed, I believe that it at least nudges developers to adopt such a practice.

### 12. Focus performance optimisation on middleware execution

As discussed earlier in this post, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) performs more validation at startup than most other CORS libraries do; moreover, it relies heavily on the functional-options pattern at startup, which is [less performant](https://www.evanjones.ca/go-functional-options-slow.html) than simply exposing a configuration struct type to developers. Therefore, building a CORS middleware with [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) is comparatively slower and more memory-hungry, though not prohibitively so (only by a factor of about 20), than building an equivalent middleware with [rs/cors](https://pkg.go.dev/github.com/rs/cors):

```plain text
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkStartup_rs_cors-8  2016933    589.9 ns/op   371 B/op   8 allocs/op
BenchmarkStartup_fcors-8      85252  12130.0 ns/op  6664 B/op  76 allocs/op

```

In my opinion, that’s tolerable. For HTTP middleware, performance at _initialisation_ indeed matters less than performance at _execution_: incurring a modest performance penalty when initialising a middleware is acceptable if invocations of that middleware require relatively few CPU cycles and cause inconsequential amounts of heap allocations.

[jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) makes such a tradeoff: it frontloads as much of the necessary work as possible at middleware initialisation in order to spare the hot path, i.e. middleware execution. That initial investment, both in time and memory, is typically recouped after only a modest number of middleware invocations. And if your use case is such that server startup shall not suffer even a few microseconds' delay, you can always riff on [Mat Ryer](https://twitter.com/matryer)’s [tricks for lazy middleware initialisation](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html#sync-once-to-setup-dependencies-10).

You may be tempted to relegate execution time of a middleware to secondary importance, because network I/O is likely to be the dominating factor in performance. However, this is only true if middleware invocations remain inexpensive; and many CORS middleware libraries, in part because [they give developers ](https://jub0bs.com/posts/2023-02-08-fearless-cors/#do-not-support-custom-callbacks)[_too much_](https://jub0bs.com/posts/2023-02-08-fearless-cors/#do-not-support-custom-callbacks)[ freedom](https://jub0bs.com/posts/2023-02-08-fearless-cors/#do-not-support-custom-callbacks), simply cannot provide such guarantee.

As often in software design, [constraints liberate](https://wiki.c2.com/?LiberatingConstraint=). By [restricting the kind of origin patterns that developers can specify in their CORS configuration](https://jub0bs.com/posts/2023-02-08-fearless-cors/#disallow-dangerous-origin-patterns), [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) is free to rely on an efficient tree-like data structure for fast origin lookup. And by denying developers a means to [specify their CORS policy as a custom callback](https://jub0bs.com/posts/2023-02-08-fearless-cors/#do-not-support-custom-callbacks), [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) prevents them from inadvertently carrying out expensive computations (such as compiling a regexp) or I/O-intensive tasks (such as querying a database) during execution of their CORS middleware.

The consequence of those deliberate constraints is that a middleware built by [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) is fast and incurs little to no heap allocation. I plan to release more representative benchmarks in the coming weeks, but here are a few data points showing how a middleware built with [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) typically fares against a functionally equivalent middleware built with [rs/cors](https://pkg.go.dev/github.com/rs/cors):

```plain text
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkNoCORSMiddleware-8   10899271  111.4 ns/op   94 B/op  0 allocs/op
BenchmarkActual_rs_cors-8      3810650  309.9 ns/op  209 B/op  2 allocs/op
BenchmarkActual_fcors-8        8711205  118.2 ns/op   94 B/op  0 allocs/op
BenchmarkPreflight_rs_cors-8   4191823  275.9 ns/op  246 B/op  0 allocs/op
BenchmarkPreflight_fcors-8    18226157  127.0 ns/op    0 B/op  0 allocs/op

```

As you can see, [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) is nimble and purrs quietly.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Thank you for your attention and your patience through this long and winding post! At this stage, I hope that at least some of my arguments for _Fearless CORS_ have swayed you.

If you’re a Gopher in need of a dependable CORS middleware library, you are welcome to take [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors) for a spin. I’ll wait (an indefinite amount of time) for feedback before releasing version 1.0.0 but, as far as I’m concerned, the library is feature-complete and production-ready. Please let me know on [Twitter](https://twitter.com/jub0bs) or [Mastodon](https://infosec.exchange/@jub0bs) if it makes your life easier!

Although I’ve endeavoured to avoid past design mistakes of others, I’ve likely made brand new ones myself. If you disagree with _Fearless CORS_, or perceive shortcomings (bugs, missing features, misfeatures, etc.) in [jub0bs/fcors](https://pkg.go.dev/github.com/jub0bs/fcors), please [open an issue on GitHub](https://github.com/jub0bs/fcors/issues/new)

Finally, if your language of choice isn’t Go and lacks a good CORS library, feel free to draw inspiration from _Fearless CORS_ to implement a better CORS library for that language; and do let me know how it goes!

I’m indebted to both Michael Smith (a.k.a. [sideshowbarker on Stack Overflow](https://stackoverflow.com/users/441757/sideshowbarker) and elsewhere) and [Mat Ryer](https://twitter.com/matryer) for generously taking the time to read an early draft of this particularly lengthy post and giving me valuable tips on how to improve it.

I’m also grateful to [Anne Van Kesteren](https://twitter.com/annevk) and [Jake Archibald](https://twitter.com/jaffathecake) for clarifying aspects of the CORS protocol that initially tripped me up, and to [Titouan Rigoudy](https://developer.chrome.com/authors/titouan/) for elucidating some of the finer points of Private Network Access.

Finally, although I’ve never interacted with them directly, [Joshua Bloch](https://twitter.com/joshbloch) and [John Ousterhout](https://twitter.com/johnousterhout), through their books and talks, had a profound impact on _Fearless CORS_.
