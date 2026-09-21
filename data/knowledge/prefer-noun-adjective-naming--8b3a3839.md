---
title: "Prefer Noun-Adjective Naming"
notion_id: 8b3a3839-8467-444a-b45f-f9cf1ec3f73e
notion_url: https://app.notion.com/p/Prefer-Noun-Adjective-Naming-8b3a38398467444ab45ff9cf1ec3f73e
last_edited: 2024-06-05T18:05:00.000Z
source_url: https://kyleshevlin.com/prefer-noun-adjective-naming/
tags: ["Article", "Kyle Shevlin", "English", "Programming"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This one will be super short.

In English grammar, the adjective comes before the noun. The green house. The large dog. The old man. Etc, etc.

Because of this, we tend to also name our functions, components, objects, etc, the same way. `PrimaryButton`, `DefaultHeader`, `StandardLayout`, etc.

I prefer to use noun-adjective naming for one simple reason: alphabetical order.

Let’s say I’m creating several components related to the same noun. I’ve been working on a streaming app for a client recently, so I’ll use a `Stream` as my noun. A `Stream` can be in three states: `created`, `ended`, or `live`.

If I follow English grammar and make components for each state, I end up with `CreatedStream`, `EndedStream`, and `LiveStream`.

Now, imagine that these are right next to all your other components in the same directory. They could be separated by dozens of components inbetween `C` and `E`, or `E` and `L`.

```plain text
- CreatedStream.tsx
- CurrentUser.tsx
- EndedStream.tsx
- FullScreen.tsx
- GoLiveButton.tsx
- LiveStream.tsx
// etc...
```

Also, try searching for them. What if you forget that it’s `Ended` and not `Completed` for that stage. You have to get the right adjective to get your search on the right path.

But if you flip them and go noun-adjective, `StreamCreated`, `StreamEnded`, and `StreamLive`, all your components are right next to each other in your directory.

```plain text
- CurrentUser.tsx
- FullScreen.tsx
- GoLiveButton.tsx
- StreamCreated.tsx
- StreamEnded.tsx
- StreamLive.tsx
// etc...
```

In addition, it’s much easier to search this way. Just type the noun `Stream` and choose which one you need.

It’s a small change, but I think it makes for far better organization in your codebase. Try it, see if you like it, too.
