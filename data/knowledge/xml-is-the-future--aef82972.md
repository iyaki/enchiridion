---
title: "XML is the future"
notion_id: aef82972-65c0-4c38-8565-d2dd5b67258c
notion_url: https://app.notion.com/p/XML-is-the-future-aef8297265c04c388565d2dd5b67258c
last_edited: 2023-06-27T22:04:00.000Z
source_url: https://www.bitecode.dev/p/hype-cycles
tags: ["English", "Reflection", "Programming", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Summary

_My first hype exposure was "use the _[_Extensible Markup Language_](https://en.wikipedia.org/wiki/XML)_ for everything". Learning from it allowed me to live through the front end stack explosion, the micro-service overdose and many, many more silly trends._

_It turns out Grandma was right. Eat vegetables, exercise, sleep well._

_And use the right tool for the right job._

_Well, she didn't say that last one._

_But she could have._

## Hype cycles

When I started programming, XML was going to replace everything. HTML with [XHTML](https://en.wikipedia.org/wiki/XHTML), validation with [DTD](https://en.wikipedia.org/wiki/Document_type_definition), transformation and presentation with [XSLT](https://en.wikipedia.org/wiki/XSLT), communication with [SOAP](https://en.wikipedia.org/wiki/SOAP).

I missed the train on the OOP hype, that was the generation before me, but I read so many articles warning me about it, that I applied the reasoning to XML: let's wait and see if this is actually as good as they say before investing everything in it.

Turns out, XML was not the future. It was mostly technical debt.

It was mostly useful for things like documents, and I believe the most successful use of it are still MS Office and LibreOffice file formats. They are just zips of XML.

I was lucky to learn this lesson very early in my career: there is no silver bullet, any single tool, no matter how good it is, must be evaluated from the engineering point of view of pros and cons. Everything has a cost, and implies compromises. It's a matter of ROI. Which is hard to evaluate without experience.

Bottom line, time is once again the great equalizer, there is no substitute to observe how a complex system evolves, no matter your model of the world.

But above all, I learned that **geeks think they are rational beings, while they are completely influenced by buzz, marketing, and their emotions**. Even more so than the average person, because they believe they are less susceptible to it than normies, so they have a blind spot.

## And so it begins

XML was just the beginning of many, many waves of hype.

When MongoDB came around ([it's web scale!](https://youtu.be/b2F-DItXtZs)), suddenly you had to use [NoSQL](https://en.wikipedia.org/wiki/NoSQL) for everything. Didn't matter that there was absolutely no relation between 2 NoSQL systems. It's like labeling a country as "doesn't speak English". Didn't matter [MongoDB](https://en.wikipedia.org/wiki/MongoDB) was a terrible product at the time that was destroying your data (they did fix that, it's now a good DB to have in your toolbox). Didn't matter that most people using it didn't need free replication because their data could fit in a [SQlite](https://docs.python.org/3/library/sqlite3.html) file.

So we watched beginners put their data with no schema, no consistency, and broken validation in a big bag of blobs. The projects fail in mass.

Then the [node](https://nodejs.org/en) era arrived. [Isomorphic JavaScript](https://en.wikipedia.org/wiki/Isomorphic_JavaScript) was all the rage, you _had_ to use the same language in the frontend and the backend, and make everything [async](https://docs.python.org/3/library/asyncio.html). But JS sucked, so most JS projects were created... to avoid writing [ES5](https://en.wikipedia.org/wiki/ECMAScript#5th_Edition). I mean, no import, no namespace, terrible scoping, [schizophreniac ](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/this)`this`, [prototype based inheritance](https://en.wikipedia.org/wiki/Prototype-based_programming), weak types, come on! So we got [coffeescript](https://coffeescript.org/), then [babel](https://babeljs.io/), [webpack](https://webpack.js.org/), [typescript](https://www.typescriptlang.org/), [react + JSX,](https://legacy.reactjs.org/docs/jsx-in-depth.html) etc.

We were told to stay on top of the most modern ecosystem, and by that I mean dealing with compatibility being broken every two months. That’s the price of cutting edge tree-shaking transpilation. That, and a [left-pad](https://www.theregister.com/2016/03/23/npm_left_pad_chaos/) way of life you couldn’t debug because the map files were generated wrong.

At this point, everything needed to be a [Single Page Application](https://en.wikipedia.org/wiki/Single-page_application) with client-side routing, [immutable data structures](https://immutable-js.com/) and some kind of store. That is, if you could chose between flux, redux, alt, reflux, flummox, fluxible, fluxxor, marty.js, fynx, MacFly, DeLorean.js, fluxify, fluxury, exim, fluxtore, Redx, fluxx… No, [I’m not making that up.](https://github.com/kriasoft/react-starter-kit/issues/22)

But because you still had to pass a lot of data through the wire, and since everything had to be on the client, [GraphQL](https://graphql.org/) was born. Of course, all that stuff had terrible accessibility, [SEO](https://en.wikipedia.org/wiki/Search_engine_optimization) and first-rendering time issues, leading to the rise of [Server-Side Rendering](https://vuejs.org/guide/scaling-up/ssr.html), aka CGI with extra steps. This couldn’t stop there, so the community added hydration on top.

This turned out to be an immense addition in complexity, and created tons and tons of disposable code base, leading to, you get it, failed projects and waste of money.

Because, of course, most of those tasks could have been done with [Ruby-On-Rail](https://rubyonrails.org/), [Symfony](https://symfony.com/) or [Django](https://www.djangoproject.com/) and a pinch of [jQuery](https://jquery.com/). At least, they would have been finished with those boring techs. Instead, dead projects began to accumulate, and for one Figma shinning, you had a trail of hidden bodies behind corporate walls nobody dared to talk about.

It was taboo to speak about this madness. You were the one not getting it.

## It's a revolution!

You would think people drowning while trying to produce a basic [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) app would have been a red flag.

Instead, it inspired teams everywhere in the world to make things harder on themselves.

First, the "everything should be a [micro-service](https://en.wikipedia.org/wiki/Microservices)" crowd started to take over. Every single small website had a [docker](https://www.docker.com/resources/what-container/) container for the restish API, plus one for the front end, and one for the database. Indirection layers on top of indirection layers. To communicate between all that, why not a little message queue? [ZeroMQ](https://zeromq.org/), [RabbitMQ](https://www.rabbitmq.com/)... And a good exchange format, like [grpc](https://grpc.io/) with [protobuff](https://github.com/protocolbuffers/protobuf).

Believe it or not, it became very hard to make your todo-list app work with all those, so a solution was found: adding orchestration. [Docker swarm](https://docs.docker.com/engine/swarm/), and now [kubernetes](https://kubernetes.io/).

At this stage, so much time and money were obliterated the cloud felt like a savior: they will do all that for you, for a fee. You just had to learn their entire way of doing things, debug their black box, be locked in their ecosystem, and carefully optimize and configure - using state-of-the-art templated YAML files and hostile UIs - your entire project, so that you could only spend 10 times more on hosting, and not 10000 times by mistake.

Easy.

Second, big data arrived. You had to store every single click of your users. A/B test everything as well, so that you consistently annoy 10% of your customers and make support unbearable. Now the data you had was gigantic! And if it was not, you had to believe it, and you needed some kind of [Dynamo](https://en.wikipedia.org/wiki/Amazon_DynamoDB) data lake. Or maybe a time series db. Or a graph one. You needed something, that's for sure.

Third, all of that stuff was now very slow. It was not because of the terrible technical decisions leading to use Google level industrial architectures for your 100 request/seconds website, no. It was because you used a slow language. So let's rewrite everything in Go. Or Rust.

The compilation step is not going to have any impact on the feedback loop anyway, since the CI pipeline already takes 73 minutes.

That was the last straw, so out of tiredness, devs went back to simple ways...

Just kidding, they went in flocks to [serverless lambda](https://en.wikipedia.org/wiki/AWS_Lambda) and SaaS services you call [from the edge](https://de.wikipedia.org/wiki/Edge_Computing), cause not owning your stack is the future!

## Building products you don't need

Meanwhile, while the blog posts about burn out were increasing tenfold, somewhere at the top, leaders heard the call of money.

You can't grow without making everything social.

Gamify, gamify, gamify.

Block chain will change the universe.

You need an [AMP website](https://en.wikipedia.org/wiki/Accelerated_Mobile_Pages).

Your stuff is not competitive without Machine Learning.

If you lived through all those, you know what remains about it: almost nothing.

A few "share" buttons and "login with" workflows. Some points and badges. Graphs.

Things either died, or filled the niche they were good at, as they should.

Some were replaced by the future of today.

## Coming back to reason

I like the new hype: [YAGNI](https://en.wikipedia.org/wiki/You_aren%27t_gonna_need_it) is popular again.

Projects like [Vue](https://vuejs.org/), [HTMX](https://htmx.org/) and [unpoly](https://unpoly.com/), [alpine.js](https://alpinejs.dev/) or just vanilla are getting traction.

There is talk of coming back to using Postgres for most things.

[37signals is on the spotlight once more, because they left the cloud.](https://world.hey.com/dhh/we-have-left-the-cloud-251760fb)

It will, of course, be overdone. Because minimalism being hyped is still... hype.

You do need the cloud, containers, nosql, go, rust and js build systems. Modern software requirements, customers’ expectations and incredible new features are not to be ignored.

Just not for everything.

Nothing is ever needed for everything.
