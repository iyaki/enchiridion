---
title: "Ginkgo - Testing framework for expressive tests - Behavior Driven Development (BDD)"
notion_id: 0ae7548c-5750-4d7c-b1de-10af405244d1
notion_url: https://app.notion.com/p/Ginkgo-Testing-framework-for-expressive-tests-Behavior-Driven-Development-BDD-0ae7548c57504d7cb1de10af405244d1
last_edited: 2023-02-01T16:41:00.000Z
source_url: https://onsi.github.io/ginkgo/
tags: ["English", "Go", "Testing", "Untried", "Framework/Library"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->























## 



### 



```

```







### 





### 







### 



```

```



```

```





### 

















```

```





### 









> 

























## 



### 











### 













### 

















### 











```

```





```

```













### 







### 















### 



















### 







> 











### 







### 























> 



### 



































### 









- 
- 
- 







### 









### 







### 









### 





















### 



































### 











| Package | Contents |
| --- | --- |
| `github.com/onsi/ginkgo/v2/dsl/core` | The core DSL including all container, setup, and subject nodes (`Describe`, `Context`, `BeforeEach`, `BeforeSuite`, `It`, etc...) as well as the most commonly used functions (`RunSpecs`, `Skip`, `Fail`, `By`, `GinkgoT`) |
| `github.com/onsi/ginkgo/v2/decorators` | The decorator DSL includes all Ginkgo's decorators (e.g. `Label`, `Ordered`, `Serial`, etc...) |
| `github.com/onsi/ginkgo/v2/reporting` | The reporting DSL includes all reporting-related nodes and types (e.g. `Report`, `CurrentSpecReport`, `ReportAfterEach`, `AddReportEntry`) |
| `github.com/onsi/ginkgo/v2/table` | The table DSL includes all table-related types and functions (e.g. `DescribeTable`, `Entry`, `EntryDescription`) |



## 







### 













### 

























### 













### 













### 









### 



























### 































### 











### 















> 



### 













### 



















### 













### 





### 









### 









### 



















### 











- 
- 
- 
- 
- 
- 



| Query | Behavior |
| --- | --- |
| `ginkgo --label-filter="integration"` | Match any specs with the `integration` label |
| `ginkgo --label-filter="!slow"` | Avoid any specs labelled `slow` |
| `ginkgo --label-filter="network && !slow"` | Run specs labelled `network` that aren't `slow` |
| `ginkgo --label-filter=/library/` | Run specs with labels matching the regular expression `library` - this will match the three library-related specs in our example. |













- 
- 
- 

















### 



- 
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















### 





















### 











### 











- 
- 
- 

### 





















```

```

### 















































### 













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



















- 







## 



### 





### 



















### 



### 





























### 





### 

















### 















### 













### 









### 









### 





- 
- 
- 







### 









### 















## 





### 













- 
- 
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











### 





1. 
2. 
3. 













### 













### 













### 







### 











### 

















### 





### 





















> 





### 





### 











### 























### 









### 







### 





### 









### 







- 
- 
- 



### 







### 

















### 



















## 



### 



























### 













































### 















### 





### 







## 







### 

















### 













### 









### 











### 









- 
- 
- 
- 
- 
- 
- 
- 





### 









## 

### 











### 






