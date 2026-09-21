---
title: "Rethinking the Modern Web"
notion_id: da26b196-6506-4d75-ab0a-86cd3c7f9c58
notion_url: https://app.notion.com/p/Rethinking-the-Modern-Web-da26b19665064d75ab0a86cd3c7f9c58
last_edited: 2023-04-22T01:19:00.000Z
source_url: https://dev.to/oxharris/rethinking-the-modern-web-5cn1
tags: ["Article", "dev.to", "English", "Web Development", "Reflection"]
---
Frontend has an engineering problem!

It is one field that has forged about the most enviable tooling ecosystem by roughly every measure - from growth rate to technical wizardry - and along with that an incredible amount of thought leadership on its every move,

and yet at its bottom line:

- the average website performance is only regressing with us!
- the average developer productivity is only grinding to a halt!

So it turns out, all of that "vibrant ecosystem" isn't really translating to more "accessible", "functional" apps on the _user_ front, and neither to more "speed" and "productivity" on the _developer_ front!

The closer you get to Frontend's bottlenecks and to how much of that is tooling-induced the more you are wondering if we just have created more problems with our tools than we've solved!

---

---

---

---

## The Counterintuitive Equations!

Consider how self-defeating it gets...

### ...on the _user_ front

Given [real world network and device-capability factors](https://infrequently.org/2017/10/can-you-afford-it-real-world-web-performance-budgets/#js-is-your-most-expensive-asset), performance starts with _shipping less bytes_, especially when it coms to JavaScript - [your most expensive asset](https://infrequently.org/2017/10/can-you-afford-it-real-world-web-performance-budgets/#js-is-your-most-expensive-asset). (Truth is, all of the extra things like instant navigations and transitions that will delight your users need your app to _first be accessible_!)

Unfortunately, SPA frameworks, who have this as their primary call, have their engineering model going the _opposite_ way: **towards more JavaScript**! These may follow different programming paradigms... but all have got the same _JavaScript_ bottom-line!

- `[HTML, CSS, JS]` > `Build_Step` > `[JS, JS, JS]` e.g. Svelte, Vue
- `[JS, JS, JS]` > `Build_Step` > `[JS, JS, JS]` e.g. React

> 

Your final bundle here is the sum of your _framework_ bundle, plus your _essential_ JS, plus the _induced_ (usually larger) JS from page structure, content, and styling - the obvious math to Frontend's [terrifying payloads](https://httparchive.org/reports/page-weight#bytesJs)!

With your application's entire weight now culminating in JavaScript... and its non-JS aspects - page structure, content, and styling - now being accounted for in JS, you're now wrongly set up for the real world! It soon becomes clear how much of a knife edge you're on as your application becomes a real thing and starts to regress with every weight gained!

Not surprising... the idea of a "highly-optimized", "accessible", "functional" application quickly falls apart!

Reporting on this in the 2017 Real-world Web Performance Budgets ([here](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#:~:text=real%20world%20network%20and%20device-capability%20factors)), Alex Russell [relates](https://infrequently.org/2017/10/can-you-afford-it-real-world-web-performance-budgets/#:~:text=i%27ve%20seen%20teams%20that%20have,%20faster%22%20experiences%20under%20real-world%20conditions.):

> 

_"I've seen teams that have just finished re-building on a modern tech stack cringe for an hour as we walk them through the experience of using their 'better', 'faster' experiences under real-world conditions."_

Tim Kadlec also [notes from field experience](https://timkadlec.com/remembers/2020-04-21-the-cost-of-javascript-frameworks/#:~:text=what%20is%20clear%3A%20right%20now%2C%20if%20you%E2%80%99re%20using%20a%20framework,in%20the%20best%20of%20scenarios.&text=if%20you%20are%20going%20to%20use%20one%20of%20these%20frameworks,meantime.):

> 

_"What is clear: right now, if you’re using a framework to build your site, you’re making a trade-off in terms of initial performance—even in the best of scenarios. [...] If you are going to use one of these frameworks, then you have to take extra steps to make sure you don’t negatively impact performance in the meantime."_

It is often excluded in the conversation, but many are winding up with more tooling-induced performance problems than they ever would otherwise! How entirely counterproductive!

### ...on the _developer_ front

Performance is all about being able to tame complexity with the least amount of overhead. This calls for the abstractions/tools for managing complexity to hold to "a 'first do no harm' principle... make sure you are at least no less productive with overhead than you were without it." ([Swyx](https://dev.to/swyx/what-drives-optimal-overhead-2p3a#:~:text=a%20%22first%20do%20no%20harm%22,without%20it.))

Question is: are we really improving at managing complexity? Not when it feels like we're drowning in **more complexity than ever**, with an entire curriculum of [deep programming concepts](https://auth0.com/blog/glossary-of-modern-javascript-concepts/) taking the place of what used to be the "HTML" kind of problems, and along with that a doze of [compilers](https://dev.to/this-is-learning/a-look-at-compilation-in-javascript-frameworks-3caj) - that themselves have to cater to a number of new paradigms, magic syntaxes and dialects!

We must now rigor through the surprisingly high amount of _code-level complexity_, and pay the price of a _compile-step_ (with [Webpack](https://webpack.js.org/) and [babel](https://babeljs.io/) (or similar beast) and a barrage of configurations, plugins and extensions) to have a working web page! But that's not all, we must also bend along with mind-bending _runtime behaviours_ (of hooks and friends), and _debug_ through difficult-to-grok transforms of our code!

At this point... you're now entirely cracking a different grade of nut! Everything easy is hard again[1](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fn1)!

Jelan [shares her frustration](https://news.ycombinator.com/item?id=33134021#:~:text=I%20have%20a,this%20space%20now.):

> 

_"I'm getting more and more frustrated with all the added layers of complexity needed to work with most common frontend frameworks. I’ve hit a point where it just doesn’t seem like the end justifies the means in the vast majority of cases anymore."_

And Chris Coyer [puts a fitting analogy to that](https://css-tricks.com/the-great-divide/#:~:text=ironically,cougar%20problem.):

> 

_"Ironically, while heaps of tooling add complexity, the reason they are used is for battling complexity. Sometimes it feels like releasing cougars into the forest to handle your snake problem. Now you have a cougar problem."_

Look what was going to land you in the pit of success[2](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fn2) winding you up with **more complexity than you had at first**! (On the [essential/accidental](https://en.wikipedia.org/wiki/No_Silver_Bullet)[3](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fn3) scale, that's you having your "accidental" (tooling-induced) complexity trump your "essential" (real problem-space) complexity!) And again, that's utterly counterproductive!

If there's anything that's obvious, it is the one thing that ruins the experience each time: **that "JavaScript"**; this time, not the "JS" in the conventional "HTML, CSS, JS" approach, but the "JS" in our new "all-JS" equations:

- the "_all-JS_ bottom-line" engineering model - which brings all of page structure, content, and styling to JS (`Build_Step => [JS, JS, JS]`)!
- the "_all-JS_ front-line" engineering model - which defaults to JS for authoring page structure, content, and styling (`[JS, JS, JS] => Build_Step`)!

## Looking Back: The Paradigm Shift!

Signaled by React's [Rethinking Best Practices](https://www.youtube.com/watch?v=x7cQ3mrcKaY) pitch back in March 2013, frontend's new era has for the most part sat on what seems to be a universal axiom: **the traditional web is awful; abstract it**! Everything "traditional" about the web application story - from authoring to runtime - has since been doomed for _an abstraction_ or _a re-engineering_ - with the conventional "HTML, CSS, JS" and the DOM being the first casualty of war, and JavaScript being the new default instinct!

Whole new ecosystems of breakaway technologies have spun, each staying at a false dichotomy with web fundamentals and staying defensive about that! Here's for an insight into the thought process...

from HTML to JavaScript:

- Mike Turley (on JSX): [Why JavaScript is Eating HTML](https://css-tricks.com/why-javascript-is-eating-html/)
- Mark Dalgleish (on CSS-in-JS): [A Unified Styling Language](https://medium.com/seek-blog/a-unified-styling-language-d0c208de2660)

from the web platform to abstractions:

- Rich Harris: [Why I don't use web components](https://dev.to/richharris/why-i-don-t-use-web-components-2cia)
- Ryan Carniato: [Maybe Web Components are not the Future?](https://dev.to/ryansolid/maybe-web-components-are-not-the-future-hfh)

### ...from HTML to JavaScript

A surprising exodus ensues!

"HTML has been the running punchline in the web development community lately. [...] We are in the midst of a very large problem in the web development field, where HTML is being left in the rearview mirror in place of JavaScript", [writes Mark Steadman](https://www.deque.com/blog/javascript-frameworks-the-lost-art-of-html/#:~:text=HTML%20has%20been%20the%20running%20punchline%20in%20the%20web%20development%20community%20lately.&text=We%20are%20in%20the%20midst%20of%20a%20very%20large%20problem,JavaScript.)

React's [infamous movement](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#:~:text=Rethinking%20Best%20Practices) may have had an initial context of just [JSX](https://en.wikipedia.org/wiki/JSX_(JavaScript)) and, later, [CSS-in-JS](https://en.wikipedia.org/wiki/CSS-in-JS)! But unfortunately, this has swept and overturned how a broader share of the frontend tooling space thinks! There's now multiple JS-first agendas everywhere you look - even on web platform proposal boards!

Consider as honorable mentions:

- [HTML Modules](https://github.com/WICG/webcomponents/blob/gh-pages/proposals/html-modules-explainer.md) - a proposed "JS-first" _import_ mechanism for "HTML".
- [CSS Modules](https://github.com/css-modules/css-modules) - a "JS-first" _import_ mechanism for "CSS".

Slowly, some of what belongs in the HTML problem space are now beginning to land as JavaScript feature proposals - sometimes downright namesquatting HTML!

> 

But lest you asked how "HTML Modules" and friends don't solve a problem: maybe they do! But those would certainly not be the "HTML" kind of problems, but the "accidental limitation" type of problem in the JS-first world! Proposals like this are only symptomatic of having gone the "unconventional" way; you must endlessly build JS-first bridges, seek a JavaScript metaphor for every other HTML/CSS thing, and wish the rest of the web were colored [yellow](https://en.wikipedia.org/wiki/File:Unofficial_JavaScript_logo_2.svg)! (You can tell clearly how these aren't a feature, but a means to an end!) And notice how this comes even at the risk of _tight-coupling_ a supposed _general-purpose_, _DOM-agnostic_ programming language with the DOM! (See how HTML Modules in principle tortures the JavaScript language to produce DOM primitives!)

It has turned out to be not just more JavaScript being smuggled in through the back door, but also more HTML and CSS being thrown out the window! (Probably the greater harm!) Think the case regarding the removal - rather than improvement - of [HTML Imports](https://web.dev/imports/) from the spec in favour of [ES6 Modules](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Modules) - or possibly now the name-squatting [HTML Modules](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#:~:text=HTML%20Modules) above! It just happened that "the JS-centric zeitgeist won" this over - to put it in [Brad Frost's words](https://bradfrost.com/blog/link/why-were-breaking-up-with-css-in-js/#:~:text=the%20js-centric%20zeitgeist%20version%20won%2C%20which%20is%20why%20we%20don%E2%80%99t%20have%20html%20imports)!

> 

And is someone pointing to the removal of Scoped CSS in favour of Shadow DOM based CSS? We'll come to that!

JS-first has this way led to setting fire on many good thinking around HTML just to have them reborn in JavaScript! An irreversible, all-in proposition!

Needles to say is how therefore many highly-missing features in HTML aren't even in the radar on proposal boards! When you look around, really, how much "developer mind-share" does HTML have anymore to push those?

This has been so deep that if you dared to contend with the status quo and build a specialised career on real UI development skills that means less engineering, you'd find yourself in the marginalised half of the [Great Divide](https://css-tricks.com/the-great-divide/) and [risk not having a career](https://news.ycombinator.com/item?id=33134021#:~:text=the%20one%20redeeming%20quality%20of%20doing%20this%20kind%20of%20work%20is%20that%20it%20is%20in%20very%20high%20demand%2C%20and%20i%20worry%20that%20the%20price%20of%20becoming%20more%20specialized%20or%20doing%20something%20more%20enjoyable%20with%20less%20bloat%20is%20that%20it%20becomes%20much%20harder%20to%20find%20jobs.)! (So then, [many devs have had to acquiesce against their own will](https://bradfrost.com/blog/link/why-were-breaking-up-with-css-in-js/#:~:text=Again%20we%E2%80%99re%20all,and%20learned%20React.) considering that there are bills to pay!)

### ...from the web platform to abstractions

Suddenly, the platform is abandoned!

Born out of contempt and mischaracterization of the web platform, the focus of modern abstractions has come to seem somewhat like _re-engineering_ web languages, _replicating_ platform features and APIs, and _duplicating_ the browser's efforts! The whole idea of using the web platform has remained an unattractive option for breakaway technologies. To put it in [Alex Russell's words](https://infrequently.org/2017/10/web-components-the-long-game/#:~:text=the%20incentives%20of%20framework%20authors%20are%20not%20aligned%20with%20compatibility), "The incentives of framework authors are not aligned with compatibility". I often see "web standards" being touted only where it makes a big news or where the performance gains are the incentive!

Looking back at how React and its ecosystem has stayed detached over the years, Mikeal Rogers relates:

> 

_"I remember when React was launched, the whole thing was about DOM diffing. The value of it is this virtual DOM thing. Then we made the DOM fast, and who gives a shit now. But we’re still using React because of – I don’t know. [...] And then now we have Web Components and they can’t adopt it, because they’re on their own pattern, so we can’t take this feature upgrade from the platform. [And] I think there's a ton of other examples of this where the platform starts to catch up, and then the frameworks can’t."_ - [JS Party – Episode #89 | Changelog](https://changelog.com/jsparty/89#:~:text=i%20remember%20when%20react%20was%20launched%2C%20the%20whole%20thing%20was%20about%20dom%20diffing,we%E2%80%99re%20still%20using%20react%20because%20of%20%E2%80%93%20i%20don%E2%80%99t%20know&text=and%20then%20now%20we%20have%20web%20components%20and%20they%20can%E2%80%99t%20adopt%20it,platform%20starts%20to%20catch%20up%2C%20and%20then%20the%20frameworks%20can%E2%80%99t.)

We just seem to be disposed to trading the "norm" and yet searching for the "kind" in the _abstract_ world!

And what has been a particularly sad implication of the framework-first culture? The more we've invested in breakaway technologies and side ecosystems, _the less we've learned about our real problem space and all the opportunities to actually move the web forward_! It turns out, HTML and its ecosystem remains _undertooled_ to date, whereas the framework web continues to flourish! You're now almost guaranteed to get stranded as being in a desert land building anything in vanilla HTML and the DOM:

> 

_"It’s really depressing that most useful tools these day are made for React projects and/or require React knowledge to set up and use. This locks out many of us who are not using React for everything & who still prefer the vanilla route for their projects."_ - [Sara Soueidan](https://twitter.com/sarasoueidan/status/1098960689455075330?s=21)

You just realise how much everyone seems to have disconnected from the real "problem-space" type of problems and gotten buried in "abstract" problems, leaving us with framework-specific solutions to very common problems!

It turns out, undertooling remains a real deterrent for everyone who has craved the simplicity of the vanilla web. (And this should also be one valid challenge to the _challenge_ thrown here by Remy Sharp when he asked: [what's stopping you from using ](https://remysharp.com/2021/02/11/the-web-didnt-change-you-did#:~:text=dear%20reader%20-%20let%20me%20ask%20you%20this%2C%20and%20i%20hope%20you%20ask%20your%20colleagues%20the%20same%3A%20what%27s%20stopping%20you%20from%20using%20exactly%20method%20today%3F)[_exactly_](https://remysharp.com/2021/02/11/the-web-didnt-change-you-did#:~:text=dear%20reader%20-%20let%20me%20ask%20you%20this%2C%20and%20i%20hope%20you%20ask%20your%20colleagues%20the%20same%3A%20what%27s%20stopping%20you%20from%20using%20exactly%20method%20today%3F)[ [that] method today?](https://remysharp.com/2021/02/11/the-web-didnt-change-you-did#:~:text=dear%20reader%20-%20let%20me%20ask%20you%20this%2C%20and%20i%20hope%20you%20ask%20your%20colleagues%20the%20same%3A%20what%27s%20stopping%20you%20from%20using%20exactly%20method%20today%3F))

With so much having gone wrong in the tooling space, what follows is only expected: a community at large suffering a terrible blind spot for web fundamentals! Ground truths are now debated everywhere, "best practices" are literally going numb ([hi Pete](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#:~:text=%20Rethinking%20Best%20Practices)), and standards are increasingly losing their place among developers! (Consider a survey: [What do you know about Web Standards?](https://www.smashingmagazine.com/2019/01/web-standards-guide/#why-am-i-telling-you-this).)

Go see a typical framework-speaking dev - even in their seniors of roles; go check the typical learning path of the coming generation; and go see what the modern Frontend job descriptions on job boards are saying! Surprisingly, the modern developer career is now little about web technologies themselves and all "about [the] intricacies of the most popular frameworks", to put it in [Frédéric Bonnet's words](https://dev.to/fredericbonnet/the-third-age-of-web-development-kgj#:~:text=about%20mastering%20the%20intricacies%20of%20the%20most%20popular%20frameworks)! See that?

There's a destructive _information gap_ and a _culture erosion_[4](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fn4) now also overtaking us!

Suffice to say, the past decade has been **a life of alienation** from conventional wisdom, almost entirely spent _trading_ HTML - and burning our ships on the go - in an irreversible bet on _JavaScript_; turning our backs and throwing dirt on the web platform to bank futures on abstractions! Yet, only a few are really talking about the almost irreversible proposition of this decade-long shift...

> 

_"In elevating frontend to the land of Serious Code we have not just made things incredibly over-engineered but we have also set fire to all the ladders that we used to get up here in the first place."_ - [Laura Buns](https://dev.to/walaura/the-web-without-the-web-aeo#:~:text=In%20elevating%20frontend,here%20in%20the%20first%20place.)

## Towards a Faster Web: A New Quest!

As the ills of the Framework era become more and more palpable, the [bells and whistles of the decade](https://blog.daftcode.pl/hype-driven-development-3469fc2e9b22) are losing their charms on people! Folks are now coming around on our lived reality, and many are actively going back on their life's bets in search of sanity, at the risk of bringing people with torches and pitchforks to their door!

- Hajime Yamasaki Vukelic: [What Got Me Writing Vanilla JavaScript again](https://javascript.plainenglish.io/what-got-me-writing-vanilla-js-again-2c53756c8a4c)
- Sam Magura: [Why We're Breaking Up with CSS-in-JS](https://dev.to/srmagura/why-were-breaking-up-wiht-css-in-js-4g9b)

But what about the philosophy at the tooling layer: is it now time for a rethink? It depends on how much of a rethink you are asking of!

On the one hand, folks have been hard at work with identifying and addressing inherent overheads in the current system. (But of course, not with a view to rethinking the overall JS-first philosophy!)

For example, realising that the idea of shipping applications as just JavaScript wasn't working anything good on the user front, React and friends have since slid back to the idea of being able to send HTML over the network. This has called for major architectural rework (think [React 18's Streaming SSR architecture](https://reactjs.org/blog/2022/03/29/react-v18.html#new-suspense-features)); this being in addition to the existing idea of static site generation with build tools like [Gatsby](https://www.gatsbyjs.com/). Server-Side Rendering ([SSR](https://web.dev/rendering-on-the-web/)) and Static Site Generation ([SSG](https://web.dev/rendering-on-the-web/)) have been such a feat for the status quo!

So, right now, HTML and progressive enhancement have again become an attractive option to what has been the "all-JS or nothing" camp. (In fact, [Remix](https://remix.run/) has this as a new bragging right! In its own words: who knew?) Put together, all of the new takes here just represent a new game plan: a long walk back to the basics; a bit of a _compromise_ on its [initial takes](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#:~:text=Rethinking%20Best%20Practices), but a bit of some greater good: towards a faster, functional Frontend!

On the other hand, a new wave of innovation is taking center stage with a more ambitious take than the ongoing long, painful walk back to the basics: this time, the idea of an outright HTML-first, "progressive enhancement" architecture from start! (SitePen Engineering gives the perfect [Intro to HTML-first Frontend Frameworks](https://www.sitepen.com/blog/intro-to-html-first-frontend-frameworks) featuring [Qwik](https://qwik.builder.io/), [Marko](https://markojs.com/), [Astro](https://astro.build/), [11ty](https://www.11ty.dev/), [Fresh](https://fresh.deno.dev/), and [Enhance](https://enhance.dev/); and this is absolutely worth your time!)

I find this especially interesting because, finally, everyone can see that the idea of a faster web didn't have to be the hard, compute-intensive, and yet elusive thing we've had for years! It is now clearer than ever how much of an absolute illusion we've been on... in staying _fixated to JS_ and yet _subscribed to HTML_, albeit cleverly! You'd also realise: HTML over the network didn't have to be any "modern" wisdom or such a big deal of an innovation, or any feat! Not when this is something the web has had from genesis! (Streaming SSR from day 1 of PHP, anyone?)

Sad that we seem to only live after the fact when it comes to tooling, but thanks to the new wave of frameworks for shedding the light that openly challenges the status quo! Finally, we've come to a point that we can be hopeful about: the prospect of having "accessible", "functional" applications at half the price! (Hopefully related is how we've now for the first time - since the JS-making epoch - [recorded a less steep increase](https://almanac.httparchive.org/en/2022/javascript#:~:text=from%202021%20to%202022%2C%20an%20increase%20of%208%25%20for%20mobile%20devices%20was%20observed%2C%20whereas%20desktop%20devices%20saw%20an%20increase%20of%2010%25.&text=is%20less%20steep%20than%20in%20previous%20years) in the amount of JavaScript shipped to users!)

So are we good now? Well, not just yet!

Not when it seems that we've come only halfway with the idea: sending pages as HTML but authoring them as JS; addressing the tradeoff problem and its overheads on the _user_ front but leaving it unchanged on the _developer_ front! Clear yet? _**Much**_** of the "HTML" in our new HTML-first equations is still the **_**tooling-yielded HTML**_**, not the **_**hand-authored HTML**_; same as before where HTML is treated as a _compile target_, an _implementation detail_... something you'd abstract to have a good DX! "Tooling" still remains Frontend's primary means to its "HTML" end!

This sheer idea of getting "HTML" behind a compile wall in favor of an abstraction is also where it begins to tell that we might yet be suffering the decade-long blindspot for the platform in a new way - **embracing HTML for just the performance incentives and not for the whole idea of **_**using the platform**_! Let's just throw it out there: Frontend still isn't aligned with _using the platform_ to [give legacy tooling their well-earned rest](https://developer.chrome.com/blog/modulepreload/#:~:text=giving%20bundlers%20their%20well-earned%20rest) and [take a correction back to simplicity](https://world.hey.com/dhh/modern-web-apps-without-javascript-bundling-or-transpiling-a20f2755#:~:text=we%27re%20way%20overdue%20a%20correction%20back%20to%20simplicity%20for%20the%20frontend.)!

[Chad Fowler notes](https://twitter.com/chadfowler/status/646624348028190720) how deep-seated this is:

> 

_"The older I get, the more I realize the biggest problem to solve in tech is to get people to stop making things harder than they have to be."_

So, how about just taking the plunge... and embracing the whole web platform thing? Now, we would not only be unlocking performance on both the _user_ front and the _developer_ front, we would be doing more: **unleashing the web's full potential**! (And what could be a more ambitious goal?)

Actually, **this is why we're here**! Can we skip to the good parts?

## Introducing: Web-Native Development!

Take this as not the name of a new framework or some anti-framework movement, but as a playbook to unleashing the web's full potential.

Web-native development[5](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fn5) is an approach to web development that sees the web platform as an enabler in the whole application story, and in fact, a fundamental key to succeeding at each phase of that story - from authoring to runtime! This comes as a hard reset to the decade-long cultural shift and its pessimistic take on the web platform!

This is really about embracing and leveraging native web technologies, APIs, languages and conventions, etc, and minimizing tooling! For a fact, there comes a time in life when this is all you really want... get things done with less of the drama! "I've been in enough teams in my 20 years of programming to value this part almost more than anything else", [writes Andrea Giammarchi](https://dev.to/richharris/why-i-don-t-use-web-components-2cia#:~:text=i%27ve%20been%20in%20enough%20teams%20in%20my%2020%20years%20of%20programming%20to%20value%20this%20part%20almost%20more%20than%20anything%20else.)!

As the platform and its technologies and languages advance, we often can find multiple opportunities "to shed a lot of this tooling" and defer to native approaches. ([JS Party - Episode #89 | Changelog](https://changelog.com/jsparty/89#:~:text=and%20as%20the%20platform%20improves,%20we%20need%20to%20be%20able%20to%20shed%20a%20lot%20of%20this%20tooling.)) Good to know is that tools are only as good as being solutions to _unsolved_ problems, not _solved_ problems! Anything in exception soon begins to change the narrative to something counterproductive! Thus, for all we've bet on custom tooling, we must now beat that Sunk Cost fallacy[6](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fn6) to explore what's natively available!

At the end of the day, web-native development **gets you banking more on the web platform and less on abstractions**! You're now on the _winning_ side - rather than on the _contending_ side - of the web's "moving" story! You're now winning as the platform unfolds!

[George Katsanos explains](https://dev.to/richharris/why-i-don-t-use-web-components-2cia#:~:text=It%20should%20be,user-friendly%20interfaces.) how this allows us take back more brain cycles to finally go make something:

> 

_"It should be obvious to all of us that if we would focus all our efforts in the Platform and stop reinventing the wheel [...] we would have a lot more free time to focus on the real reason we are in this job which is to deliver fast, stable, secure, user-friendly interfaces."_

So, where's a good place to get well-rounded with the web platform? Chrome's [web.dev](https://web.dev/) developer centre is a treasure trove of resources on key web design and development subjects, maintained by the Chrome team and other industry experts! ([Here's the learning centre](https://web.dev/learn/).) And for when you need to take the web platform by its individual technologies, here is [MDN](https://developer.mozilla.org/)! (And [here's the learning centre](https://developer.mozilla.org/en-US/docs/Learn).) For showcases, here is one: [Using the platform](https://elisehe.in/2021/08/22/using-the-platform) - a delightful piece of a story on going buildless with ES6 modules and going framework-free with Web components, written by Elise Hein!

Up next is: **how far does this go in real life**? Build a twitter clone with zero tooling? Uhh, that has never been the idea, and we may never get there! The web platform is anything but a framework of its own! _At some point_, we are going to have to need higher-level abstractions over native lower-level features! Where there seems to be a bit of a bad news is that **the current state of HTML and the DOM makes that happen **_**sooner**_; it isn't long into the journey before gaps and dips in the platform gets you into forced labour! But we can easily turn this around by investing along two lines:

1. **Low-level tooling**: platform-focused initiative to _standardise_ and _factor into the platform_ common web development architectures and paradigms. (We're overdue for native-level reactivity and a more empowering component model - just to mention a few!)
2. **Higher-level tooling**: community-focused initiative to _extend_ the platform's low-level capabilities with "web-native" libraries and frameworks. (We need a new wave of modest abstractions that draw on the web platform and let us do the same!)

**This is where I **_**put my money where my mouth is**_**!** I'd like to take you on a few ideas I've been working on along these lines!

Open to some tooling ideas?

### Explore with Me...

Our journey spans a series of posts! We begin with the underlying equations and move on to a showcase of the proposals and polyfils, the userland libraries and framework!

> 

For the curious, [here's a sneak peak](https://github.com/webqit/webqit) into this project on github.

## Notes

- No frameworks were harmed in the creation of this article.

## Acknowledgements

- Special thanks to [Alex Russell](https://infrequently.org/) for the time spent reviewing this article!

## References

- [Can You Afford It?: Real-world Web Performance Budgets](https://infrequently.org/2017/10/can-you-afford-it-real-world-web-performance-budgets/) (Alex Russell)
- [2022 Page Weight Report: JavaScript Bytes](https://httparchive.org/reports/page-weight#bytesJs) (HTTP Archive)
- [The Cost of JavaScript Frameworks](https://timkadlec.com/remembers/2020-04-21-the-cost-of-javascript-frameworks/) (Tim Kadlec)
- [What drives Optimal Overhead?](https://dev.to/swyx/what-drives-optimal-overhead-2p3a) (Shawn Wang)
- [Glossary of Modern JavaScript Concepts: Part 1](https://auth0.com/blog/glossary-of-modern-javascript-concepts/) (Kim Maida)
- [Rethinking Best Practices](https://www.youtube.com/watch?v=x7cQ3mrcKaY) (Pete Hunt)
- [Why JavaScript is Eating HTML](https://css-tricks.com/why-javascript-is-eating-html/) (Mike Turley)
- [A Unified Styling Language](https://medium.com/seek-blog/a-unified-styling-language-d0c208de2660) (Mark Dalgleish)
- [Why I don't use web components](https://dev.to/richharris/why-i-don-t-use-web-components-2cia) (Rich Harris)
- [Maybe Web Components are not the Future?](https://dev.to/ryansolid/maybe-web-components-are-not-the-future-hfh) (Ryan Carniato)
- [JavaScript Frameworks & The Lost Art of HTML](https://www.deque.com/blog/javascript-frameworks-the-lost-art-of-html/) (Mark Steadman)
- [why we’re breaking up with css-in-js](https://bradfrost.com/blog/link/why-were-breaking-up-with-css-in-js/) (Brad Frost)
- [The Great Divide](https://css-tricks.com/the-great-divide/) (Chris Coyer)
- [Web Components: The Long Game](https://infrequently.org/2017/10/web-components-the-long-game/) (Alex Russell)
- [Is modern JS tooling too complicated?](https://changelog.com/jsparty/89) (JS Party – Episode #89 | Changelog)
- [The web didn't change; you did](https://remysharp.com/2021/02/11/the-web-didnt-change-you-did) (Remy Sharp)
- [Web Standards: The What, The Why, And The How](https://www.smashingmagazine.com/2019/01/web-standards-guide/) (Amy Dickens)
- [From Classicism to Metamodernism — A Short History of Web Development Series' Articles](https://dev.to/fredericbonnet/series/10459) (Frédéric Bonnet)
- [Hype Driven Development](https://blog.daftcode.pl/hype-driven-development-3469fc2e9b22) (Marek Kirejczyk)
- [What Got Me Writing Vanilla JavaScript again](https://javascript.plainenglish.io/what-got-me-writing-vanilla-js-again-2c53756c8a4c) (Hajime Yamasaki Vukelic)
- [Why We're Breaking Up with CSS-in-JS](https://dev.to/srmagura/why-were-breaking-up-wiht-css-in-js-4g9b) (Sam Magura)
- [Intro to HTML-first Frontend Frameworks](https://www.sitepen.com/blog/intro-to-html-first-frontend-frameworks) (SitePen Engineering)
- [Modern web apps without JavaScript bundling or transpiling](https://world.hey.com/dhh/modern-web-apps-without-javascript-bundling-or-transpiling-a20f2755) (David Heinemeier Hansson)

1. 

Borrowing Frank Chimero's title: [Everything Easy is Hard Again](https://frankchimero.com/blog/2018/everything-easy/) [↩](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fnref1)

1. 

Defining [The Pit of Success](https://english.stackexchange.com/questions/77535/what-does-falling-into-the-pit-of-success-mean) [↩](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fnref2)

1. 

See [Fred Brooks: No Silver Bullet (PDF)](http://worrydream.com/refs/Brooks-NoSilverBullet.pdf) [↩](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fnref3)

1. 

See [Cultural Erosion](https://geographyrevisionalevel.weebly.com/6b-cultural-erosion.html) [↩](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fnref4)

1.    

"Web-native" is a recurring theme across multiple initiatives and paradigms: [↩](https://dev.to/oxharris/rethinking-the-modern-web-5cn1#fnref5)

- [Web Native](https://webnative.tech/) - Ionic Team's presentation of the idea, but scoped to its [Capacitor](https://capacitorjs.com/) runtime.
- [Web-Native](https://dev.to/fredericbonnet/the-third-age-of-web-development-kgj#:~:text=web-native) - Frédéric Bonnet's presentation of the idea in [The Postmodernist Period](https://dev.to/fredericbonnet/the-third-age-of-web-development-kgj#:~:text=market%20(2012-today).-,the%20postmodernist%20period%20,-and%20the%20Second) of [The Third Age of Web Development](https://dev.to/fredericbonnet/the-third-age-of-web-development-kgj).
- [Modern Web](https://modern-web.dev/) - Guides, tools and libraries for modern web development built on web standards.
- [The Stackless Way](https://tutorials.yax.com/articles/the-yax-way/index.html) - Daniel Kehoe's optimistic take on web development that proposes we “use the platform” instead of frameworks and build tools.
- [Buildless](https://buildless.site/) - Pascal Schilp's paradigm for creating production websites without a build process.

Some related tags:

- #buildless - [DEV](https://dev.to/t/buildless)
- #vanilla - [DEV](https://dev.to/t/vanilla)
