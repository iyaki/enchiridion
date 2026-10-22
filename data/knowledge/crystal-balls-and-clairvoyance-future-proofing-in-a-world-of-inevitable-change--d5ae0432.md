---
title: "Crystal balls and clairvoyance: Future proofing in a world of inevitable change"
notion_id: d5ae0432-5a1a-475d-a306-3d85f5bebe93
notion_url: https://app.notion.com/p/Crystal-balls-and-clairvoyance-Future-proofing-in-a-world-of-inevitable-change-d5ae04325a1a475da3063d85f5bebe93
last_edited: 2023-01-25T18:27:00.000Z
source_url: https://stackoverflow.blog/2022/05/19/crystal-balls-and-clairvoyance-future-proofing-in-a-world-of-inevitable-change/
tags: ["Stack Overflow Blog", "English", "Project Management", "Product Management", "System Design / Software Architecture", "Programming", "Article"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

While the future may be a mystery, you can design software to accommodate future changes. But how much future-proofing gets in the way of good design?

The concept of **future-proofing** your code is one that perennially pops up in conversations about software. It sounds magical—who _wouldn’t_ want to make your code impervious to the future?

Of course, the reality is much less rosy and much messier.

In this article, I’m going to discuss what I think people mean by “future-proofing,” how you might be able to accomplish it, and when and why it might be a bad choice.

## How does one “proof” one’s “future?”

You can think of future-proofing more accurately as **change-proofing**. We can define it as:

_“Making a design, architecture, or programming decision that allows for future changes to be easier to manage, take less time or result in fewer changes to the overall code.”_

These changes can fall into a number of different categories:

- Changes to **scale** – the project has to do more of the things it’s doing, whether that means handling more traffic or work items, processing them faster, etc.
- Changes to **requirements** – new information has entered the business (or has entered the engineering team from the business) and now the system needs to change to accommodate them.
- Changes to **technology stack** – switching to a different data store or programming language, for example.
- Changes to **integrations** – the project needs to talk to a new third-party application, either on top of or instead of an existing one.
- Changes to **schemas** – we want to change the fields that define our data objects.

## Should I stay or should I grow?

The main issue with future-proofing is that it runs slap-bang into one of the central tenets of software engineering: YAGNI, or [**You Ain’t Gonna Need It**](https://en.wikipedia.org/wiki/You_aren%27t_gonna_need_it).

This principle states that until you actually know you’re going to have a change, you shouldn’t code your software in a way that anticipates that change. Violating this principle results in bloat, confusing and unnecessary abstractions that make the code harder to understand, and often mounds of tech debt.

However, if you never cast your eyes to the future, you really can run into issues where just a little extra bit of work upfront could have saved you months of it down the road.

So… should you future-proof your code? The answer, as with practically everything in software, is, “[it depends](https://medium.com/swlh/the-golden-rule-of-software-engineering-9faaaab85e78).”

I see this as a spectrum of sorts:

Every situation is unique, of course. But some kinds of changes are much less likely to deserve future-proofing, while in others, the extra work is more likely to pay off and (just as importantly) not result in significant downsides.

## Future-proofing strategies

In general, any time we want to protect ourselves from future changes, we would engage in one of two strategies.

### Modularization

**Modularization** is the act of splitting your code up into smaller chunks. Every piece of software has some level of modularization—otherwise, your code would consist of a single enormous file that had nothing but unrolled functions and primitive types. (These programs do exist, but no one in their right mind wants anything to do with them.)

Increasing modularization allows you to **isolate** your changes to a single module, and/or to more easily **swap** modules out for each other.

You can also have modularization at different levels. Extracting code into a function is a simple act of modularization that rarely has a downside until you’re at the point where your functions are so small and numerous that they’re hard to keep track of. However, you can extract things into a class, a sub-application, a microservice, or even a separate cluster; and at each level, while the isolation and flexibility goes up, so does the cognitive and administrative overhead of managing your different components.

### Abstraction

**Abstraction** is the mental model you have of the parts of your code. A module, function, class, etc. with **low abstraction** (or higher **specificity**) more closely models the action it takes or object it represents. A thing with **higher abstraction** brings your reasoning up a level. For example, instead of your code working with `Cars`, it could work with `Vehicles`. That way, if you ever need to start handling `Motorcycles`, the change necessary to do so is mitigated by the fact that you were never really talking about `Cars` to begin with.

The higher you go with abstraction, the more flexible you make your code, but also the harder it is to understand. Everyone knows what a `Car` is, but when you start working with `Vehicles`, it becomes more difficult to work with `Car`-specific things like booster seats and steering wheels. And if you keep going higher up the abstraction tree, you might find yourself coding around `Transportation`, `ManufacturedGoods`, even plain old `Objects`, and have a very hard time figuring out how to fill up your tank.

## Types of changes

Lets go through each of the types of changes I listed above. For each, I’m recommending where on the spectrum your future-proofing efforts might lie. In general, I will be applying these suggestions for when you have **no current indication that the change will happen**. As the change becomes more likely, you’d move further towards the left side of the line.

### Changes to scale

Making your code and architecture robust and able to handle whatever you throw at it is a core facet of design. However, part of that architecture and design has to be what kind of traffic you expect it to have. The higher scale you have to handle, in general, the more complex your architecture needs to be.

Making a simple app that’s designed to be used by a handful of internal users could be done via a Rails or Django application with a minimum of JavaScript, for example. But if you need to be able to handle thousands or millions of concurrent users, that’s just not going to cut it. You’re going to need lightning-fast services or micro-services, horizontal scaling and auto-scaling, and probably a more responsive front-end and more tailored and complex caching mechanisms.

Over-architecting all of this upfront when you have no anticipation of your app actually scaling to that size is flagrant YAGNI. In fact, building your system deliberately for your current scale _even if you know you may have to rewrite it when the scale changes_ is a perfectly valid strategy, sometimes called [sacrificial architecture](https://stackoverflow.blog/2021/03/01/sacrificial-architecture-learning-from-abandoned-systems/).

Having said that, there are _some_ scaling considerations you should take into account. Leaving in lots of [N+1 queries](https://stackoverflow.com/questions/97197/what-is-the-n1-selects-problem-in-orm-object-relational-mapping) or having unnecessarily large or frequent requests is bad engineering practice all around and will leave a bad taste in your users’ mouths—regardless of how many or few they are.

### Changes to requirements

This case (where your code has to change or add to its behavior) is the hardest one to guess at, as well as the one most likely to happen.

This is a case where general modularization can help you, but I’d recommend against trying to prematurely add abstraction to your code.

In particular, if we’re talking about _business rules_, isolating the rules themselves in a way that makes them easy to test and change is an overall good practice.

As a simple example, here’s some code working with a product that has some rules around what a valid product can be considered:

```plain text
product.save if product.price >= 0 && !product.name.nil?
```

Products might be saved from many different places in your application. Rather than the calling code checking these business rules, we can extract them to their own method:

```plain text
class Product
  def valid?
    self.price >= 0 && !self.name.nil?
  end
end
product.save if product.valid?
```

We can even go further and have a separate class or module that handles this check:

```plain text
module ProductValidator
  def self.valid?(product)
    product.price >= 0 && !product.price.nil?
  end
end
product.save if ProductValidator.valid?(product)
```

These are relatively minor changes and keep your code clean and easily testable, and can isolate further changes of this sort to the one place.

The key words are _“of this sort”_, however. If the requirements changes involve completely changing how the behavior of a feature acts, or adding new features, that’s where you want to start putting your foot down more firmly on the YAGNI side of things and only building what you actually know.

Let’s imagine that your product can be purchased:

```plain text
class Product
  def purchase
    # contact the purchase service and complete transaction
  end
end
```

Wait a minute, you reason. We might eventually expand into not only buying products but also selling them! Maybe renting them? We should abstract this out:

```plain text
class Product
  def perform_action
  end
end

class PurchasableProduct
  def perform_action
    # contact the purchase service and complete transaction
  end
end
```

This is classic YAGNI over-abstraction. People reading your code won’t see you purchasing a product, you’re “performing an action” on something that _resolves_ to purchasing.

You have no reason to believe that your business actually will expand into further operations. You’ve introduced a level of abstraction that removes your code one extra step from what’s actually happening and what’s being modeled.

### Changes to technology stack

I’ll come straight out and say it—there is just no good reason to build your app in a way that you plan to eventually change your basic technology stack.

This is obviously almost impossible to do from a programming language perspective. The only way you could provide more isolation is to go further into microservices—which might make sense from an architecture perspective, but not from a future-proofing one.

As for data stores, no one—_no one_—actually changes their database from MySQL to Postgres or vice versa. If you’re using an open-source or open-standards way of interacting with your data, go all into it and don’t look back. At this point, the similarities far outweigh the differences, and if you do need to switch for whatever arcane reason, you’ll need a full regression suite anyway to make sure nothing else breaks.

(Note: This point is often used to discourage the use of [ORMs](https://en.wikipedia.org/wiki/Object%E2%80%93relational_mapping). I think ORMs have other very good advantages, but the ability to switch data technologies is not a practical argument for their use.)

### Changes to integrations

In this case, we are worried about having to change how we talk to some third-party application. These services can provide things such as metrics, tracing, alerting, logs, feature flags, object storage, deployments, etc. In this case, it’s nice to have the _freedom_ to change who you’re doing business with based on cost, feature set, ease of use, etc.

I am most inclined to ensure that we have abstraction and modularization around integrations. In many cases, these third-party applications can be fairly easily switched out for each other. Once you’re providing metrics or tracing, for example, the abstractions and ideas are very close to each other regardless of which provider you’ve signed up with.

What I prefer doing is building a _facade_ around any code that needs to talk to a third-party system. Often to keep things simple, this is a wrapper around the API of whatever provider I’ve chosen, which exposes all the functionality necessary to operate.

This facade often takes just a couple of days to build (if there isn’t [already](https://github.com/flipp-oss/deimos/) [one](https://github.com/moneta-rb/moneta) [out](https://github.com/fog/fog)[ there](https://opentelemetry.io/docs/migration/opentracing/)) and can be used in whatever project needs it. I’ll often add a set of team or company defaults that most closely match our most common use case so that new projects don’t need to figure it out for themselves.

This facade allows you to avoid _vendor lock-in_ by not tying your systems to any specific provider—but the overhead is small enough that it shouldn’t confuse you or your teammates.

One thing that is important is to allow you to _break out_ of the facade by accessing the internal client or API if necessary. If you’ve chosen your provider because of specific features that need to be used, then you shouldn’t need to tie your arm around your back to prevent you from using them. However, it means that if you do end up switching your provider, you don’t have to make any changes to the “normal” case—you only need to look at code that accesses the internal client or API and focus your work there.

Here’s a rough approximation of a feature flagging library we built for internal use:

```plain text
module FlippFlags
  def backend=(backend)
    self.backend = backend
  end

  def client(config)
    self.backend.new(config)
  end
end

class FlippFlags::Backend
  def initialize(config) # this is a Ruby version of a constructor, i.e. the "new" keyword
    @client = ... # initialize the actual client
  end

def enabled?(flag, user)
    @client.enabled?(flag, user)
end

  def internal_client
    @client
  end
end

# application code
FlippFlags.backend = FlippFlags::MyFlagProvider
client = FlippFlags.client(my_config)
client.enabled?("flag_name", user_object) # true or false
```

Here you can see that the only change you need to make is to switch the `backend` that is passed into the library, and it seamlessly switches to a different provider—or maybe even an internal class that acts the same way! Otherwise, you use the client exactly the way you’d use the internal client, but without tying your code directly to that implementation.

### Changes to schemas

The shape of our data reflects our understanding of it. As we gain more understanding, we often want to change its shape. This could mean adding or removing fields, changing field types, default values, etc.

Making our data schemas backwards and forwards compatible is possible by following some simple rules. Technologies such as [Avro](https://docs.confluent.io/platform/current/schema-registry/avro.html#summary) and [Protobuf](https://docs.buf.build/breaking/rules) specify those rules and have built them into their tooling. If you’re not using those tools, though, you can “soft-enforce” similar rules whenever you change your own schema to ensure it’s unlikely that your changes will break whatever depends on your data.

Having said that, there are cases where you simply can’t follow those rules—and that’s okay! This is when you specify a migration path to go from the old data to the new data. But the rules make it so you shouldn’t need to follow that onerous process for every single data change you go through.

## Conclusion

Future-proofing is not necessarily a _goal_ of software development so much as one of the many “[ilities](https://en.wikipedia.org/wiki/List_of_system_quality_attributes)” that affect your design. Overcorrecting on this spectrum can lead to unnecessary work, tech debt, and confusing abstractions. Finding the sweet spot, though, can save you time and effort down the road.

Tags: [future-proof](https://stackoverflow.blog/tag/future-proof/), [software architecture](https://stackoverflow.blog/tag/software-architecture/), [software engineering](https://stackoverflow.blog/tag/software-engineering/)
