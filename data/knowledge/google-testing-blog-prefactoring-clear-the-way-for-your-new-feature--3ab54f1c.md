---
title: "Google Testing Blog: Prefactoring: Clear the Way for Your New Feature"
notion_id: 3ab54f1c-7d23-8124-bbef-e98de4af3edf
notion_url: https://app.notion.com/p/Google-Testing-Blog-Prefactoring-Clear-the-Way-for-Your-New-Feature-3ab54f1c7d238124bbefe98de4af3edf
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://testing.googleblog.com/2026/07/prefactoring-clear-way-for-your-new.html
tags: ["Software Development", "Refactoring", "Programming", "Productivity", "Article", "Google Testing Blog", "English"]
---












- 
- 
- 
- 



| Change 1 (Prefactoring)<br>Extract display name helper to remove duplication. | Change 2 (Feature)<br>Add middle name support. |
| --- | --- |
| + def get_display_name(user):<br>+    return f"{user.first_name}”  {user.last_name}<br># Profile page<br>-    display_name = f"{user.first_name} {user.last_name}"<br>+    display_name = get_display_name(user)<br># Email template<br>- greeting = f"Hi {user.first_name} {user.last_name},"<br>+ greeting = f"Hi {get_display_name(user)}," | def get_display_name(user):<br>-  return f"{user.first_name} {user.last_name}”<br>+  return f"{user.first_name} {user.middle_name} {user.last_name}" |


