---
title: "Don’t use booleans"
notion_id: 73f6436d-1864-4008-8c0d-e2388675b107
notion_url: https://app.notion.com/p/Don-t-use-booleans-73f6436d186440088c0de2388675b107
last_edited: 2026-09-21T16:58:00.000Z
source_url: https://www.notion.so/73f6436d186440088c0de2388675b107
tags: ["Programming", "System Design / Software Architecture", "Article", "LUU.IO", "English"]
---
Use enums instead.



_With any blanket statements like this, there are always exceptions. Though in general, I believe the use of enums is often a better choice compared to boolean, unless you really need to squeeze your data into one single physical bit._

### Readability

Recently, I came across a function call in our code-base where I challenged my team to name the arguments for some cash, and sadly, I kept my money. It looked something similar to this:

```plain text
fetch(accountId, false, true, true);

```

Obviously, the readability could improve quite a bit by having inline annotation:

```plain text
fetch(
  accountId,
  true, // include disabled,
  true, // fetch history,
  true, // fetch details
);

```

Though, as we see here, you either need a rigorous code review process or linter, or just rely on developers doing due diligence. Something like this would be much more readable:

```plain text
fetch(
  accountId,
  {ItemStatus::Active, ItemStatus::Disabled},
  FetchOptions::History | FetchOptions::Details,
);

```

Notice the difference between the `ItemStatus` and `FetchOptions` enums. One is an exlusive enum (i.e. each item can only have one of the boolean values) and one is a bit-mask flags.

### Explicit typing

Why would you want explicit typing? Imaging staying up late fixing a bug and calling this `fetch` function. If you’re like me, there’s a high probability that you’ll cross your eyes and pass true/false in the wrong places, calling `fetch(false, true, false)` instead of `fetch(true, false, false)`. Your tests would have needed to have all the permutations of the booleans to catch that issue. And good luck spotting the issue even if your tests failed.

With the enum version of the API, you can’t accidentally call `fetch(FetchOptions::History, {ItemStatus::Active})`. The compiler or interpreter should bark pretty loudly in this case.

If your programming language allows, use explicit bit-mask types instead of `int`. This will make sure that the call site cannot accidentally pass in the wrong flags.

### Behavior Driven & Extendability

Whenever I was about to use booleans and stepped back to consider using enums as an alternative, I often found myself going over the actual use-cases, or the behaviors and states of the application. It’s easy to just slap on another boolean at that point in time without considering the actual use-case or the behavior of the API. Let’s look at our `fetch` API again; boolean signature is:

```plain text
fetch(
  int accountId,
  bool includeDisabled,
  bool includeHistory,
  bool includeDetails,
);

```

Suppose we have a new piece of data called `flags`, and we want to return those `flags`. However, we need to be backward compatible to the existing callers, and only return the flags if they explicitly indicate so. So let’s update our API signature:

```plain text
fetch(
  int accountId,
  bool includeDisabled,
  bool includeHistory,
  bool includeDetails,
  bool includeFlags, // <----- This is our shiny new boolean
);

```

One issue though: the flags actually come from the details. In other words, we need to validate that if `includeFlags` is `true`, the caller must also set `includeDetails` to true, and need to document it somewhere (assuming your callers read documentation):

```plain text
if (includeFlags) {
  assert includeHistory;
}

```

Imagine having three or more enums that are coindependent, we’ll suddenly need a big compatibility matrix. Let’s see how we can extend the enum version of this API:

```plain text
enum FetchOptions {
  None = 0,
  History = 1,
  Details = 2,
}

```

```plain text
fetch(int accountId, set<ItemStatus> statuses, FetchOptions options);

```

The only thing that we need to do is to add a new option in our `FetchOptions` enum:

```plain text
enum FetchOptions {
  None = 0,
  History = 1,
  Details = 2,
  DetailsWithFlags = 4, // <----- This is our new fetch option
}

```

The caller just **can’t** pass in the wrong thing!

## Summary

With every codebase, there is always a fine balance with how much one should “over engineer”. Though, I feel that the overhead for using enums of booleans are quite low, and the benefits enums provide are worth the effort.

Copyright 2024, Steven Luu. The opinions expressed here are mine.
