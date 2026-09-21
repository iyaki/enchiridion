---
title: "Harness engineering for coding agent users"
notion_id: 34854f1c-7d23-81dc-a9d3-ebfe421305de
notion_url: https://app.notion.com/p/Harness-engineering-for-coding-agent-users-34854f1c7d2381dca9d3ebfe421305de
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://martinfowler.com/articles/harness-engineering.html
tags: ["Tool", "Article", "Martin Fowler", "English", "Software Architecture", "Artificial Intelligence (AI)", "Programming", "DevOps"]
---




![image](https://martinfowler.com/articles/harness-engineering/harness-bounded-contexts.png)





![image](https://martinfowler.com/articles/harness-engineering/harness-overview.png)



## 



- 
- 



## 



- 
- 





|  | Direction | Computational / Inferential | Example implementations |
| --- | --- | --- | --- |
| Coding conventions | feedforward | Inferential | AGENTS.md, Skills |
| Instructions how to bootstrap a new project | feedforward | Both | Skill with instructions and a bootstrap script |
| Code mods | feedforward | Computational | A tool with access to OpenRewrite recipes |
| Structural tests | feedback | Computational | A pre-commit (or coding agent) hook running ArchUnit tests that check for violations of module boundaries |
| Instructions how to review | feedback | Inferential | Skills |

## 





## 





- 
- 

![image](https://martinfowler.com/articles/harness-engineering/harness-change-lifecycle-examples.png)





- 
- 

![image](https://martinfowler.com/articles/harness-engineering/harness-continuous-feedback-examples.png)



## 





### 











### 





- 
- 

### 



- 
- 





## 





## 







## 







## 





- 
- 
- 
- 
- 



## 





## 



****
