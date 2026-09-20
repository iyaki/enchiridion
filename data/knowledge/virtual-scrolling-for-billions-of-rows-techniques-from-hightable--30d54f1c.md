---
title: "Virtual Scrolling for Billions of Rows — Techniques from HighTable"
notion_id: 30d54f1c-7d23-8150-88c0-c944f9ddc942
notion_url: https://app.notion.com/p/Virtual-Scrolling-for-Billions-of-Rows-Techniques-from-HighTable-30d54f1c7d23815088c0c944f9ddc942
last_edited: 2026-02-20T01:55:00.000Z
source_url: https://rednegra.net/blog/20260212-virtual-scroll/
tags: ["Article", "Tutorial", "redbee - Medium", "English", "Web Development", "Frontend", "React", "Performance", "User Experience"]
---




![image](https://rednegra.net/blog/20260212-virtual-scroll/aW48cjwKPd-1440.jpeg)





- 
- 
- 
- 
- 
- 
- 
- 
- 

## 



```

```











## 







![image](https://rednegra.net/blog/20260212-virtual-scroll/9ZFXIAtLm3-1911.jpeg)

## 





```

```





> 



---



---



---



---



---



---



---



---



---



---

| viewport.clientHeight | 163px |
| --- | --- |
| viewport.scrollHeight | 301px |
| table.clientHeight | 301px |
| viewport.scrollTop | 0px |
| first pixel | 0px |
| last pixel | 163px |







1. 
2. 
3. 
4. 



## 



> 





---



---



---



---



---



---

---

---

---

---





```

```



```

```





```

```







> 



### 





## 







```

```



> 



```

```







| canvas.clientHeight | 301px |
| --- | --- |
| viewport.scrollTop | 0px |
| table.clientHeight | 181px |
| table.style.top | 0px |
| pixels above | 0px |





```

```



> 



```

```



> 



### 





## 











```

```



```

```



```

```





| canvas.clientHeight | 300px |
| --- | --- |
| table.clientHeight | 181px |
| viewport.scrollTop | 0px |
| table.style.top | 0px |
| downscale factor | 2,189,781,021 |
| unreachable rows / px | 72,992,695 |



> 



- 
- 
- 
- 



### 







## 







- 
- 
- 



```

```



```

```



- 
- 



```

```





| canvas.clientHeight | 300px |
| --- | --- |
| table.clientHeight | 181px |
| viewport.scrollTop | 0px |
| global anchor | 0px |
| local offset | 0px |
| table.style.top | 0px |



### 









## 



> 





> 



1. 
2. 
3. 
4. 
5. 





```

```

```

```

```

```



### 



## 




