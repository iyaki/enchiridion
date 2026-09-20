---
title: "Your Docker Image Is 1.2GB. Here Is How To Get It Under 80MB. | The Practical Developer"
notion_id: 36d54f1c-7d23-818f-8560-c331bd45675a
notion_url: https://app.notion.com/p/Your-Docker-Image-Is-1-2GB-Here-Is-How-To-Get-It-Under-80MB-The-Practical-Developer-36d54f1c7d23818f8560c331bd45675a
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://the-practical-developer.online/posts/docker-image-from-1gb-to-80mb/
tags: ["The Practical Developer", "English", "Docker", "Node.js", "DevOps", "Backend", "Containerization", "Article", "Tutorial"]
---


![image](https://images.unsplash.com/photo-1605745341112-85968b19335b?w=1200&q=80)







## 



```

```



```

```



## 



| Image | Size |
| --- | --- |
| `node:22` | 1.21GB |
| `node:22-slim` | 412MB |
| `node:22-alpine` | 178MB |





## 





```

```



## 





```

```







## 







## 



```

```



## 



```

```





## 

| Step | Image | Size | Saved |
| --- | --- | --- | --- |
| 0. Naive | `node:22` | 1.21GB | - |
| 1. `slim` base | `node:22-slim` | 412MB | -67% |
| 2. `.dockerignore` | `node:22-slim` | 388MB | -6% |
| 3. Multi-stage + prune | `node:22-slim` | 198MB | -49% |
| 4. Layer caching | `node:22-slim` | 198MB | (rebuild speed) |
| 5. Alpine runtime | `node:22-alpine` | 96MB | -52% |
| 6. Distroless | `distroless/nodejs22` | 78MB | -19% |



## 











## 









## 






