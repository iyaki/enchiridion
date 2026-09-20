---
title: "Structured-Prompt-Driven Development (SPDD)"
notion_id: 35154f1c-7d23-81bb-87c9-e5d46b20d50a
notion_url: https://app.notion.com/p/Structured-Prompt-Driven-Development-SPDD-35154f1c7d2381bb87c9e5d46b20d50a
last_edited: 2026-04-29T02:42:00.000Z
source_url: https://martinfowler.com/articles/structured-prompt-driven/
tags: ["English", "Programming", "Artificial Intelligence (AI)", "Software Development", "DevOps", "Systems Design / Software Architecture", "Article", "Tutorial", "martinfowler"]
---


- 
- 
- 
- 





![image](https://martinfowler.com/articles/structured-prompt-driven/spdd-overview.svg)



## 







### 





![image](https://martinfowler.com/articles/structured-prompt-driven/spdd-reason-canvas.svg)



- 
- 
- 
- 



- 



- 
- 



### 











- 
- 
- 
- 







| Command | Type | Purpose |
| --- | --- | --- |
| [/spdd-story](https://github.com/gszhangwei/open-spdd/blob/v0.4.9/internal/templates/data/optional/spdd-story.md) | Optional | Breaks a large requirement into independent, deliverable user stories following the INVEST principle. |
| [/spdd-analysis](https://github.com/gszhangwei/open-spdd/blob/v0.4.9/internal/templates/data/core/spdd-analysis.md) | Core | Extracts domain keywords from requirements, scans relevant code, and produces a strategic analysis covering domain concepts, risks, and design direction. |
| [/spdd-reasons-canvas](https://github.com/gszhangwei/open-spdd/blob/v0.4.9/internal/templates/data/core/spdd-reasons-canvas.md) | Core | Generates the full REASONS Canvas — an executable blueprint from high-level rationale down to method-level operations. |
| [/spdd-generate](https://github.com/gszhangwei/open-spdd/blob/v0.4.9/internal/templates/data/core/spdd-generate.md) | Core | Reads the Canvas and generates code task by task, strictly following the operations, norms, and safeguards defined in the prompt. |
| [/spdd-api-test](https://github.com/gszhangwei/open-spdd/blob/v0.4.9/internal/templates/data/optional/spdd-api-test.md) | Optional | Generates a cURL-based API test script with structured test cases covering normal, boundary, and error scenarios. |
| [/spdd-prompt-update](https://github.com/gszhangwei/open-spdd/blob/v0.4.9/internal/templates/data/core/spdd-prompt-update.md) | Core | Incrementally updates the Canvas when requirements change (requirements → prompt → code). |
| [/spdd-sync](https://github.com/gszhangwei/open-spdd/blob/v0.4.9/internal/templates/data/core/spdd-sync.md) | Core | Synchronizes code-side changes (refactoring, fixes) back into the Canvas so the prompt stays an accurate record of the current code (code → prompt). |

## 





### 





### 



- 
- 
- 
- 
- 
- 



### 



































1. 
2. 
3. 

- 
- 
- 
- 

1. 
2. 

- 
- 
- 

1. 
2. 
3. 

### 







- 
- 
- 









- 
- 
- 
- 



### 











- 
- 
- 

### 





- 
- 
- 
- 
- 



![image](https://martinfowler.com/articles/structured-prompt-driven/example-analysis-review.png)









### 

















### 



### 









1. 
2. 
3. 





### 







![image](https://martinfowler.com/articles/structured-prompt-driven/example-script-generation.png)









![image](https://martinfowler.com/articles/structured-prompt-driven/example-test-results.png)



### 





![image](https://martinfowler.com/articles/structured-prompt-driven/code-review.svg)



### 





![image](https://martinfowler.com/articles/structured-prompt-driven/example-prompt-update-a.png)

























1. 
2. 
3. 
4. 

### 

> 





```

```















### 





![image](https://martinfowler.com/articles/structured-prompt-driven/example-regression-results.png)



### 



### 







### 









### 









### 



1. 
2. 
3. 
4. 





## 











### 



| Rating | Scenario | Notes |
| --- | --- | --- |
| ★★★★★ | Scaled, standardized delivery | High-repeat business logic that needs long-term maintainability (e.g., building many similar APIs, automating core business workflows). |
| ★★★★★ | High compliance and hard constraints | Environments where you must follow regulations, security standards, or strict architectural rules (e.g., financial core systems, multi-channel / multi-client deployments). |
| ★★★★☆ | Team collaboration and auditability | Multi-person delivery where changes must be fully traceable and reviewable end-to-end. |
| ★★★★☆ | Cross-cutting consistency work | Complex refactors where logic must stay tightly synchronized across multiple microservices or different languages. |
| ★★☆☆☆ | Firefighting hotfixes | “Stop the bleeding” production fixes where speed matters more than architectural discipline. |
| ★★☆☆☆ | Exploratory spikes | When the goal is to validate an idea quickly rather than ship production-quality software, SPDD's governance overhead won't pay back. |
| ★★☆☆☆ | One-off scripts | Disposable data cleanup or temporary scripts where SPDD's upfront cost is too high relative to the value. |
| ★☆☆☆☆ | Context black holes | When the domain is poorly defined and business rules are unclear, you can't set meaningful boundaries for the model. |
| ★☆☆☆☆ | Pure creative / visual work | Tasks driven by taste and aesthetics rather than logic (e.g., UI visual exploration, marketing copy). |

### 



| Benefit | Impact | Speed | What you get |
| --- | --- | --- | --- |
| Determinism | High | Immediate | Encode logic in a precise spec, which significantly reduces hallucination and “creative” interpretation. |
| Traceability | High | Immediate | Every meaningful change can be traced back to the structured prompt, closing the audit loop. |
| Faster reviews | High | Short-term | Code “arrives” closer to team standards, so reviews focus on logic and design, not formatting and cleanup. |
| Explainability | Medium-High | Gradual | Intent and behavior are visible at the natural-language level, lowering the cognitive load for understanding and maintenance. |
| Safer evolution | High | Long-term | Well-defined boundaries and stepwise implementation make targeted changes lower-risk and easier to iterate. |



| Area | Barrier | Nature | What it takes |
| --- | --- | --- | --- |
| Mindset shift | High | Ongoing training | Teams have to adapt to “design first” rather than “code first.” |
| Senior expertise up front | Medium-High | Per-feature | Engineers who can translate business rules into clean abstractions and design constraints. |
| Automation tooling | Medium | Infrastructure setup | Without automation, SPDD hits a throughput ceiling and struggles to keep prompts consistent. [openspdd](https://github.com/gszhangwei/open-spdd) runs the workflow in this article—from analysis and structured REASONS prompts through code and optional test support—as repeatable CLI steps, so artifacts stay versioned and reviewable instead of trapped in chat. Larger organizations may still layer a knowledge platform on top to manage and reuse assets at scale. |

## 











> 

## 









## 







****
