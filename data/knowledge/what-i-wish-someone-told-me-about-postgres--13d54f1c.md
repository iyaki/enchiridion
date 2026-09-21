---
title: "What I Wish Someone Told Me About Postgres"
notion_id: 13d54f1c-7d23-8191-a8a7-c17410bf31f1
notion_url: https://app.notion.com/p/What-I-Wish-Someone-Told-Me-About-Postgres-13d54f1c7d238191a8a7c17410bf31f1
last_edited: 2024-11-15T20:41:00.000Z
source_url: https://challahscript.com/what_i_wish_someone_told_me_about_postgres
tags: ["Databases", "Article", "ChallahScript (hibachrach - Hazel Bachrach)", "English"]
---








## 







## 



1. 
2. 
3. 

## 

### 



```

```



```

```



```

```



### 





| Operation | Description |
| --- | --- |
| `x IS NULL` | returns `true` if `x` evaluates to `NULL`, `false` otherwise |
| `x IS NOT NULL` | returns `true` if `x` does not evaluate to `NULL`, `false` otherwise |
| `x IS NOT DISTINCT FROM y` | the same as `x = y` but `NULL` is treated as a normal value |
| `x IS DISTINCT FROM y` | the same as `x != y`/`x <> y` but `NULL` is treated as a normal value |





```

```

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

### 



| Command | What it does |
| --- | --- |
| `\?` | List all of the shortcuts |
| `\d` | Shows list of relations (tables and sequences) as well as said relation’s owner |
| `\d+` | Same as `\d` but also includes the size and some other metadata |
| `\d table_name` | Shows the schema of a table (list of columns, including said column’s type, nullability, and default) as well as any indexes or foreign key constraints on said table |
| `\e` | Opens your default editor (set as the `$EDITOR` environment variable) to edit your query there |
| `\h SQL_KEYWORD` | Get syntax and link to docs for `SQL_KEYWORD` |



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



## 

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



```

```



```

```

## 

### 



### 



| Lock Mode | Example Statements |
| --- | --- |
| `ACCESS SHARE` | `SELECT` |
| `ROW SHARE` | `SELECT ... FOR UPDATE` |
| `ROW EXCLUSIVE` | `UPDATE`, `DELETE`, `INSERT` |
| `SHARE UPDATE EXCLUSIVE` | `CREATE INDEX CONCURRENTLY` |
| `SHARE` | `CREATE INDEX` (not `CONCURRENTLY`) |
| `ACCESS EXCLUSIVE` | Many forms of `ALTER TABLE` and `ALTER INDEX` |



| Requested Lock Mode | Existing Lock Mode |  |  |  |  |  |  |  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
|  | `ACCESS SHARE` | `ROW SHARE` | `ROW EXCL.` | `SHARE UPDATE EXCL.` | `SHARE` | `ACCESS EXCL.` |  |  |
| `ACCESS SHARE` |  |  |  |  |  | X |  |  |
| `ROW SHARE` |  |  |  |  |  | X |  |  |
| `ROW EXCL.` |  |  |  |  | X | X |  |  |
| `SHARE UPDATE EXCL.` |  |  |  | X | X | X |  |  |
| `SHARE` |  |  | X | X |  | X |  |  |
| `ACCESS EXCL.` | X | X | X | X | X | X |  |  |



| Client 1 is doing… | Client 2 wants to do a … | Can Client 2 start? |
| --- | --- | --- |
| `UPDATE` | `SELECT` | ✅ **Yes** |
| `UPDATE` | `CREATE INDEX CONCURRENTLY` | 🚫 **No, must wait** |
| `SELECT` | `CREATE INDEX` | ✅ **Yes** |
| `SELECT` | `ALTER TABLE` | 🚫 **No, must wait**[3](https://challahscript.com/what_i_wish_someone_told_me_about_postgres#fn:alter_table_special_cases) |
| `ALTER TABLE` | `SELECT` | 🚫 **No, must wait**[3](https://challahscript.com/what_i_wish_someone_told_me_about_postgres#fn:alter_table_special_cases) |



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





## 





### 



### 



### 



```

```



```

```



```

```





## 



1. 



1. 



1. 



1. 



1. 



1. 



1. 



1. 


