---
title: "HASURA - Build modern apps & APIs 10x faster - Instant GraphQL APIs"
notion_id: b0c52c86-7992-4779-94c0-c4c70981dddf
notion_url: https://app.notion.com/p/HASURA-Build-modern-apps-APIs-10x-faster-Instant-GraphQL-APIs-b0c52c867992477994c0c4c70981dddf
last_edited: 2023-02-01T17:02:00.000Z
source_url: https://hasura.io/
tags: ["English", "System Design / Software Architecture", "Programming", "Productivity", "Product Management", "Project Management", "Untried", "Tool", "Service"]
---
- 
- 
- 

## 





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## 

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



![image](data:image/png;base64,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)



## 











## 

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





# 


