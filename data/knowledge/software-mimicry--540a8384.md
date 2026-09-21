---
title: "Software Mimicry"
notion_id: 540a8384-a4c8-46de-a83a-f5cf915af4ac
notion_url: https://app.notion.com/p/Software-Mimicry-540a8384a4c846dea83af5cf915af4ac
last_edited: 2023-04-22T01:02:00.000Z
source_url: https://www.hillelwayne.com/post/software-mimicry/
tags: ["System Design / Software Architecture", "Article", "Hillel Wayne", "English"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

**Mimicry** is when software X reimplements at a higher level a core feature of software Y. The produced _facsimile_ has some, but not all, of the same properties, enough to “look like” it’s the same thing but missing many of the nuances. This exists in every kind of software. One language can mimic another, a library can mimic a language, a database engine can mimic a product, etc.

### Example

_Design Patterns_ introduces the strategy pattern to apply multiple algorithms to the same call:

> INTENT: Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it. […] a Strategy object encapsulates an algorithm.

In [Design Patterns in Dynamic Programming](https://norvig.com/design-patterns/), Peter Norvig notes that strategies are just trivially replacable with higher-order functions (HOFs). Compare the algorithm “select for all of the even numbers of a list” for a “GoF language”, as compared to how you’d do with HOFs:

```plain text
// GoF pseudocode

abstract class Strategy[T, U] {
  public execute(t: T): U;
}

class IsEven inherits Strategy[Int, Bool] {
  public execute(i: Int): Bool {return i % 2;}
}

def filter(pred: Strategy[T, Bool], l: List[T]): List[T] {
  // ...
}

filter(IsEven(), [1,2,3]) -> 2

```

```plain text
// JavaScript
[1,2, 3].filter((x) => x % 2 == 0)

```

For this reason, design patterns are often (erroneously) called “missing language features”. I think it’s more interesting to think of the GoF-language as _mimicking_ HOFs with the strategy pattern _facsimile_. The strategy pattern reimplements HOFs as a code-level idiom.

### Benefits and Drawbacks

The obvious drawback is that the facsimile is considerably more verbose; in our case, a hundred characters to fifteen. It’s also fragile, as it only works for unary functions. If we wanted to pass in a function with two inputs, like `Add`, we’d need to define a new strategy.

The biggest drawback is that the facsimile doesn’t “fit” with the rest of the language. Program features are supposed to complement each other. In the JavaScript example, I used an HOF along with an anonymous function. The two features work together and with the broader language ecosystem. In contrast, our strategy facsimile is a custom class. It won’t synergize with anything else in the language, and it won’t be understood by third party tooling.

Surprisingly, there are also a couple of advantages to mimicry. There are many possible ways to design a feature. JavaScript, Haskell, and F# all do HOFs in slightly different ways, and using those languages fixes the way you use HOFs. Mimicking a feature at a user level lets us use several different designs in the same program. For example, in addition to `Strategy`, I might define `StrategyWithSetup`:

```plain text
class StrategyWithSetup inherits Strategy[T, U] {
  public do(t: T): U {
    setup();
    execute(t);
    teardown();
  }

  private setup, teardown; // blah blah blah
}

```

The biggest advantage is that you still have the _rest_ of the mimicking language. If the GoF language does 99% of what you need, and you just want HOFs for one specific case, it’s better to create a facsimile than to switch everything over to Haskell. So mimicry can be a pragmatic choice to make a tool more “right for the job”.

## Mimicry in the Wild

The most famous and notorious case of mimicry is Lisp macros. Lisp macros can extend the language with new constructs, so any language feature can be reimplemented as a lisp macro. This is why Paul Graham [calls Lisp the most powerful language](http://www.paulgraham.com/avg.html):

> By induction, the only programmers in a position to see all the differences in power between the various languages are those who understand the most powerful one. (This is probably what Eric Raymond meant about Lisp making you a better programmer.) You can’t trust the opinions of the others, because of the Blub paradox: they’re satisfied with whatever language they happen to use, because it dictates the way they think about programs.

The five languages that Eric Raymond recommends to hackers fall at various points on the power continuum. […] I think Lisp is at the top. And to support this claim I’ll tell you about one of the things I find missing when I look at the other four languages. How can you get anything done in them, I think, without macros?

This also leads to the [Lisp Curse](http://www.winestockwebdesign.com/Essays/Lisp_Curse.html): since everybody can make their own facsimile of a given feature, there’s no agreement in the broader community of which facsimile to get behind, so no way to target tooling or build an ecosystem.

Mimicry appears in other languages, too: static-typed languages can mimic dynamic languages with a `Dyn` type, leading to the (misleading) meme [“dynamic languages have just one type”](https://existentialtype.wordpress.com/2011/03/19/dynamic-languages-are-static-languages/). Before Go added generics, people would mimic generics with `interface{}`, code generators, and [unicode trickery](https://www.reddit.com/r/rust/comments/5penft/comment/dcsq64p/?context=3). JML creates a facsimile of language-level contracts with structured comments. [This Clojure article](http://www.appliedscience.studio/articles/array-programming-for-clojurists.html) implements facsimiles of several APL features without understanding [how they actually work](https://twitter.com/hillelogram/status/1343984109241167873). And, of course, there’s [Greenspun’s Tenth Rule](https://en.wikipedia.org/wiki/Greenspun%27s_tenth_rule).

Mimicry also appears in higher-level software. Most IDEs have a “vim mode” which mimics most of vim’s normal commands and some configuration, but doesn’t support things like first-class macros. People use Postgres json fields to mimic NoSQL paradigms. Whenever I change email software I have to make facsimiles of all the features the new software has that the old one doesn’t.

## Similar Ideas

“Representation” is similar to mimicry, except it’s about encoding a data structure in primitives that aren’t well-suited to it. Examples would be representing a tree in SQL tables, representing a matrix with nested arrays, or representing a graph in json.

The opposite of mimicry would be “reification”, where you turn an idiom into a software feature. My favorite example is units of measure. Lots of languages have libraries or comment annotations to represent units of measure, but only two languages, F# and [Frink](https://www.hillelwayne.com/post/frink/), turn it into an actual language feature. Another example: Rust reifies the borrow checker, which [was inspired](https://news.ycombinator.com/item?id=15715524) by [Cyclone](https://www.cs.umd.edu/~mwh/papers/ismm.pdf), a dialect of C.

## Miscellaneous Thoughts

So some other things this made me think about, but haven’t fully developed:

- Often people will use mimicry to call something uninteresting: “what’s the big deal about loops? That’s just `if` and `goto`.“. This is rightly called out when it’s someone else doing it to us. Mimicry often involves facsimiling the most obvious features of a language but not the nuances. Which ties back to the idea that languages are holistic, as are mimics.
- On the other hand, lots of people try Haskell because they got a taste of pure typed FP from language facsimiles. Mimicry can be a form of idea diffusion.
- Facsimiles can be _exported_: Someone who knows X reimplements a feature in language Y. They can also be _imported_: Someone who knows Y implements a feature they saw in language X. The difference is that in the former case, they’re most familiar with the thing they’re mimicking, while in the latter case they’re most familiar with the mimicking context. Do these lead to observably different facsimiles?
- I regularly mimic relational databases in Excel spreadsheets and feel no shame whatsoever.

Admittedly half the point of this post was to try out this format for sharing terms of I have. I often come up with terms for things I see a lot that I think might be useful for other people to have, too. Overall I think this particular format worked out, so I’ll probably do more terms in the future.

_Thanks to _[_Predrag Gruevski_](https://predr.ag/)_, _[_Quinn Wilton_](https://twitter.com/wilton_quinn)_, and _[_Lars Hupel_](https://lars.hupel.info/)_ for feedback._

_I shared an early version of this post on my _[_weekly newsletter_](https://buttondown.email/hillelwayne/)_ where I announce new blog posts and write additional essays. If you enjoyed this, _[_why not subscribe_](https://buttondown.email/hillelwayne/)_?_
