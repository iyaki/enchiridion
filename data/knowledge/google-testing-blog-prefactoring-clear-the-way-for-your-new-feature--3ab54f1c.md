---
title: "Google Testing Blog: Prefactoring: Clear the Way for Your New Feature"
notion_id: 3ab54f1c-7d23-8124-bbef-e98de4af3edf
notion_url: https://app.notion.com/p/Google-Testing-Blog-Prefactoring-Clear-the-Way-for-Your-New-Feature-3ab54f1c7d238124bbefe98de4af3edf
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://testing.googleblog.com/2026/07/prefactoring-clear-way-for-your-new.html
tags: ["Article", "Google Testing Blog", "English", "Software Development", "Refactoring", "Programming", "Productivity"]
---
_This article was adapted from a Google _[_Tech on the Toilet_](https://testing.googleblog.com/2024/12/tech-on-toilet-driving-software.html)_ (TotT) episode. You can download a _[_printer-friendly version_](https://docs.google.com/document/d/1VVEWCA8YSct4raMepjQNuLlHyZpmKj5FYj-O7-Ufe0E/preview)_ of this TotT episode and post it in your office._

By Rahul Singal

“First make the change easy, then make the easy change.” - paraphrased from Kent Beck

You're working on a new feature, but the existing code wasn't written with future changes in mind. Trying to force the feature in directly gets complicated fast. One change leads to another, and before you know it you're already a few files deep fixing things you never planned to touch.

Prefactoring (short for "preparatory refactoring") is the practice of reworking existing code to make it more suitable for an upcoming change before you actually implement the new functionality. Instead of cleaning up code as an afterthought or trying to force a new feature into an incompatible structure, you restructure the codebase first.

Prefactoring helps you:

- Easily implement new features: Restructuring the codebase first ensures your new feature fits naturally into the code.
- Speed up reviews: It's easier to review the refactoring and the feature in separate changes.
- Avoid bugs: Isolating cleanups from functional logic can help prevent bugs.
- Roll back safely: If you need to roll back, it is much easier to revert small, focused changes.

Here is a simplified example of a prefactoring change:

| Change 1 (Prefactoring)<br>Extract display name helper to remove duplication. | Change 2 (Feature)<br>Add middle name support. |
| --- | --- |
| + def get_display_name(user):<br>+    return f"{user.first_name}”  {user.last_name}<br># Profile page<br>-    display_name = f"{user.first_name} {user.last_name}"<br>+    display_name = get_display_name(user)<br># Email template<br>- greeting = f"Hi {user.first_name} {user.last_name},"<br>+ greeting = f"Hi {get_display_name(user)}," | def get_display_name(user):<br>-  return f"{user.first_name} {user.last_name}”<br>+  return f"{user.first_name} {user.middle_name} {user.last_name}" |

You can prefactor a change that is already in review too! If your reviewer suggests a related cleanup during review, you can also extract it into a new base change to keep your current change focused on the feature. Note that not every cleanup needs to be prefactoring: you can do the cleanup in a follow up change if the cleanup doesn’t block your feature, or even in the same change if the cleanup is small enough.
