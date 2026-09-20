---
title: "Explaining The Postgres Meme"
notion_id: ec3d9b92-7c6c-4d9e-a011-9f37c6537366
notion_url: https://app.notion.com/p/Explaining-The-Postgres-Meme-ec3d9b927c6c4d9ea0119f37c6537366
last_edited: 2023-09-11T18:28:00.000Z
source_url: https://www.avestura.dev/blog/explaining-the-postgres-meme
tags: ["Article", "English", "Databases"]
---


<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





### 



### 





## 



- 
- 
- 
- 
- 
- 
- 
- 

## 



### 







```

```



### 









```

```



### 







```

```



```

```





```

```





```

```

### 



```

```

### 





```

```

### 



### 





```

```



### 





```

```

### 







```

```

### 









```

```



- 
- 
- 



### 





```

```



### 



```

```

## 



### 







```

```



### 





- 
- 
- 
- 

### 





```

```





```

```





```

```

| # | Node | Timings |  | Rows |  |  | Loops |  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
|  |  | Exclusive | Inclusive | Rows X | Actual | Plan |  |  |
|  | 1. | Hash Inner Join (cost=18.1..55.28 rows=648 width=202) (actual=0.019..0.022 rows=3 loops=1) Hash Cond: ((w.city)::text = (c.name)::text) | 0.007 ms | 0.022 ms | ↑ 216 | 3 | 648 | 1 |
|  | 2. | Seq Scan on weather as w (cost=0..13.6 rows=360 width=186) (actual=0.008..0.008 rows=3 loops=1) | 0.008 ms | 0.008 ms | ↑ 120 | 3 | 360 | 1 |
|  | 3. | Hash (cost=13.6..13.6 rows=360 width=194) (actual=0.007..0.007 rows=3 loops=1) Buckets: 1024 Batches: 1 Memory Usage: 9 kB | 0.003 ms | 0.007 ms | ↑ 120 | 3 | 360 | 1 |
|  | 4. | Seq Scan on city as c (cost=0..13.6 rows=360 width=194) (actual=0.003..0.004 rows=3 loops=1) | 0.004 ms | 0.004 ms | ↑ 120 | 3 | 360 | 1 |



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 





```

```



- 

### 





```

```

### 



- 
- 







```

```

### 





```

```



```

```

### 







```

```





- 
- 

```

```

### 





```

```





### 







```

```



### 







- 
- 
- 



## 



### 



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 





```

```





```

```

### 





```

```



```

```

### 





```

```

### 







```

```

### 





```

```





```

```







```

```







```

```





```

```



- 
- 
- 



### 





```

```

### 







### 







### 







```

```



### 







- 

## 



### 





```

```





```

```



> 



### 





- 
- 
- 
- 



- 
- 
- 
- 



| Isolation Level | Dirty Read | Nonrepeatable Read | Phantom Read | Serialization Anomaly |
| --- | --- | --- | --- | --- |
| Read uncommitted | ⚠️ Possible (✅ not in PG) | ⚠️ Possible | ⚠️ Possible | ⚠️ Possible |
| Read committed | ✅ Not possible | ⚠️ Possible | ⚠️ Possible | ⚠️ Possible |
| Repeatable read | ✅ Not possible | ✅ Not possible | ⚠️ Possible (✅ not in PG) | ⚠️ Possible |
| Serializable | ✅ Not possible | ✅ Not possible | ✅ Not possible | ✅ Not possible |



### 











- 
- 

### 





```

```

### 







```

```

### 





```

```





```

```





```

```





```

```



### 





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



### 







```

```





```

```



- 
- 



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



1. 
2. 
3. 
4. 
5. 
6. 

### 





```

```



### 





```

```



### 





```

```





```

```





```

```





```

```





```

```





```

```



```

```

## 



### 







### 





```

```





```

```



```

```





```

```



### 





### 



- 
- 
- 
- 





```

```





```

```



### 





```

```





```

```

### 





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



```

```





```

```





```

```



```

```



### 





```

```











```

```

### 





### 









## 



### 





### 





```

```

### 









### 





- 
- 
- 



```

```



### 





```

```



```

```

### 



- 
- 
- 

### 





1. 
2. 
3. 
4. 



- 

## 



### 



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->





### 





```

```





```

```





```

```







```

```



### 











```

```

- 

### 





| Problem Class | Verify Solution | Find Solution | Example |
| --- | --- | --- | --- |
| P | 😁 Easy | 😁 Easy | Multiply numbers |
| NP | 😁 Easy | 😥 Hard | 8 Queens |
| NP-hard | 😥 Hard | 😭 Hard | Best next move in Chess |



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

> 

### 







<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

- 

### 





- 

### 



- 

### 







- 

## 



### 





```

```





### 



### 

- 
- 



### 



```

```







```

```

### 





```

```

- 

### 







> 

- 

### 





```

```





```

```

- 

### 



## 





## 



- 
- 
- 
- 
- 
- 
- 
