---
title: "Expose Your Design System to LLMs ・ Hardik Pandya"
notion_id: 31a54f1c-7d23-81c7-b147-d752dd1b9e6f
notion_url: https://app.notion.com/p/Expose-Your-Design-System-to-LLMs-Hardik-Pandya-31a54f1c7d2381c7b147d752dd1b9e6f
last_edited: 2026-03-05T01:56:00.000Z
source_url: https://hvpandya.com/llm-design-systems
tags: ["hvpandya.com", "English", "UI/UX", "Design", "Artificial Intelligence (AI)", "Productivity", "Article"]
---


## 







## 







- 
- 
- 
- 
- 
- 



- 
- 
- 





### 







### 







- 
- 
- 



## 

















## 





```

```

```

```



- 
- 
- 
- 
- 

### 



### 



> 

### 





```

```



```

```



```

```



### 



1. 
2. 
3. 
4. 
5. 
6. 





> 

### 









```

```



### 



### 





- 
- 
- 
- 

## 



### 

| Tier | Files | What it covers |
| --- | --- | --- |
| Foundations + tokens | 19 | Color, typography, spacing, radius, elevation, motion, z-index, iconography, accessibility, and the master map of every CSS variable |
| Components (atoms, molecules, organisms) | 38 | Button, input, avatar, tabs, dropdown-menu, modal-dialog, form, table, navigation, content-panel |
| Patterns | 7 | Canvas-content-flow, three-column-layout, panel-expand-collapse, responsive-grid, form-layout |

### 



```

```

```

```



1. 
2. 
3. 



### 



- 
- 
- 
- 
- 



### 

| Scenario | Without legible DS | With legible DS |
| --- | --- | --- |
| Link color | AI writes `#2563EB` in one component, `#1D4ED8` in another. Both look blue. Neither is “wrong.” | `var(--color-link)`. One blue. Every component. Every session. |
| Card padding | AI writes `12px` here, `16px` there, `14px` somewhere else. All “look fine.” | `var(--space-200)` or the audit script fails with a specific suggestion. |
| Dark mode | Hardcoded `#FFFFFF` breaks. Each component needs individual fixes. | Token chain resolves per theme automatically. Zero component changes. |
| New session | AI starts fresh. Different guesses. Two sessions’ worth of inconsistency. | AI reads the same specs. Same tokens. Same output quality. |
| Design review | Manual visual comparison. “Does this look right?” “I think so.” | Automated audit: 0 errors = ship. Non-zero = specific line numbers and fix suggestions. |

## 







### 

| Metric | Before | After |
| --- | --- | --- |
| Hardcoded CSS values | 418 across 28 files | 0 |
| Spec files | 0 | 64 (3 tiers) |
| Design tokens mapped | Scattered, inconsistent | 230+ with three layer indirection |
| Upstream packages tracked | Not tracked | 39 with drift detection |
| AI output consistency | Variable, depends on session | Constrained: same spec, same tokens, same audit |



## 

### 



### 



### 



### 



### 


