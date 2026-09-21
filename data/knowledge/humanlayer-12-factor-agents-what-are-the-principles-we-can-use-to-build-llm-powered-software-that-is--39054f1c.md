---
title: "humanlayer/12-factor-agents: What are the principles we can use to build LLM-powered software that is actually good enough to put in the hands of production customers?"
notion_id: 39054f1c-7d23-8157-8964-e35d08727132
notion_url: https://app.notion.com/p/humanlayer-12-factor-agents-What-are-the-principles-we-can-use-to-build-LLM-powered-software-that-i-39054f1c7d2381578964e35d08727132
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://github.com/humanlayer/12-factor-agents
tags: ["English", "Artificial Intelligence (AI)", "Software Development", "Programming", "Systems Design / Software Architecture", "Product Management", "Tool", "Article", "GitHub"]
---
# 

![image](https://camo.githubusercontent.com/ca34f9d4db9ab9f34416e899ff805504820e1b798ea511881859e82eabc2488d/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f436f64652d417061636865253230322e302d626c75652e737667)

![image](https://camo.githubusercontent.com/e96ba20d4ad3391c17fd32a72efa44a1c9b55ea9daf65f3991e4edc1a194d504/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f436f6e74656e742d434325323042592d2d5341253230342e302d6c69676874677265792e737667)

![image](https://camo.githubusercontent.com/e9cd644682259047523514c2968835aca35c9529c76a3f36eb85e7cb57d8d624/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f636861742d646973636f72642d353836354632)

![image](https://camo.githubusercontent.com/6c112d2088c809e72887f779f27623a10f1b8647cfaa848c9184d9889278c07b/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6169646f74656e67696e6565722d636f6e665f74616c6b5f2831376d292d7768697465)

![image](https://camo.githubusercontent.com/1774ff77c6a8df66b41e1bed843fb53a3d8d1e23fb495895d44945c8d88434ed/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f796f75747562652d646565705f646976652d6372696d736f6e)











![image](https://camo.githubusercontent.com/118510b0c1b6853f572787d983136d3858b1afd41bae236a59d5ba1c56b4c5f3/68747470733a2f2f7374617469632e73636172662e73682f612e706e673f782d707869643d32616361643939612d633264392d343864662d383666352d396361383036316237626639)

![image](https://private-user-images.githubusercontent.com/3730605/430151074-23286ad8-7bef-4902-b371-88ff6a22e998.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3ODI3NTIxMjIsIm5iZiI6MTc4Mjc1MTgyMiwicGF0aCI6Ii8zNzMwNjA1LzQzMDE1MTA3NC0yMzI4NmFkOC03YmVmLTQ5MDItYjM3MS04OGZmNmEyMmU5OTgucG5nP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9QUtJQVZDT0RZTFNBNTNQUUs0WkElMkYyMDI2MDYyOSUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA2MjlUMTY1MDIyWiZYLUFtei1FeHBpcmVzPTMwMCZYLUFtei1TaWduYXR1cmU9MzdiZTQ0MTc1NTBjZDkxNGZlNmI0ZTYzNWFkMmJkOTRiMWQ2Yzc5MTk2NzM0YmMyM2YzYWIwMTI0YmJkYjc2ZiZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmcmVzcG9uc2UtY29udGVudC10eXBlPWltYWdlJTJGcG5nIn0.yRlPK7nyOPJCzAcX8LreqgWmZ-s6_UlPpufFbprvjO8)













> 





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
- 
- 
- 
- 

### 

## 



### 



![image](https://github.com/humanlayer/12-factor-agents/raw/main/img/010-software-dag.png)

### 



![image](https://github.com/humanlayer/12-factor-agents/raw/main/img/015-dag-orchestrators.png)

### 



![image](https://github.com/humanlayer/12-factor-agents/raw/main/img/025-agent-dag.png)



![image](https://github.com/humanlayer/12-factor-agents/raw/main/img/026-agent-dag-lines.png)



### 





1. 
2. 
3. 
4. 

```

```





****

****

## 





1. 
2. 
3. 
4. 
5. 
6. 
7. 

****

### 



1. 
2. 
3. 
4. 
5. 

> 

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
- 
- 
- 
- 

## 

- 

## 

- 
- 
- 
- 
- 
- 
- 

## 



![image](https://avatars.githubusercontent.com/u/3730605?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/50557586?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/66259401?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/18105223?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/4084885?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/39267118?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/1882972?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/380402?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/16674643?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/85041180?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/36044389?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/7169731?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/15862501?v=4&s=80)

![image](https://avatars.githubusercontent.com/u/160066852?v=4&s=80)

## 




