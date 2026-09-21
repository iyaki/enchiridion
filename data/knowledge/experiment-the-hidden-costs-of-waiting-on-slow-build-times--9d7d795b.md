---
title: "Experiment: The hidden costs of waiting on slow build times"
notion_id: 9d7d795b-40c5-4109-8047-4a6ac851c966
notion_url: https://app.notion.com/p/Experiment-The-hidden-costs-of-waiting-on-slow-build-times-9d7d795b40c5410980474a6ac851c966
last_edited: 2023-04-16T22:46:00.000Z
source_url: https://github.blog/2022-12-08-experiment-the-hidden-costs-of-waiting-on-slow-build-times/
tags: ["English", "DevOps", "Productivity", "Continuous Integration/Continuous Delivery", "Article", "Github Blog"]
---










## 











> 

## 









| **Compute power** | **Fedora 35 build** | **Fedora 36 build** | **Average time** <br>**(minutes)** | **Cost/minute for compute** | **Total cost of 1 build** | **Developer cost** <br>**(1 dev)** | **Developer cost** <br>**(5 devs)** |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 2 core | 5:24:27 | 4:54:02 | 310 | $0.008 | $2.48 | $389.98 | $1,939.98 |
| 4 core | 2:46:33 | 2:57:47 | 173 | $0.016 | $2.77 | $219.02 | $1,084.02 |
| 8 core | 1:32:13 | 1:30:41 | 92 | $0.032 | $2.94 | $117.94 | $577.94 |
| 16 core | 0:54:31 | 0:54:14 | 55 | $0.064 | $3.52 | $72.27 | $347.27 |
| 32 core | 0:36:21 | 0:32:21 | 35 | $0.128 | $4.48 | $48.23 | $223.23 |
| 64 core | 0:29:25 | 0:24:24 | 27 | $0.256 | $6.91 | $40.66 | $175.66 |









## 







| **Compute power** | **Minutes** | **Cost of 1 build** | **Partial developer cost** <br>**(1 dev)** | **Partial developer cost** <br>**(5 devs)** |
| --- | --- | --- | --- | --- |
| 2 core | 310 | $2.48 | $77.48 | $377.48 |
| 4 core | 173 | $2.77 | $77.77 | $377.77 |
| 8 core | 92 | $2.94 | $77.94 | $377.94 |
| 16 core | 55 | $3.52 | $78.52 | $378.52 |
| 32 core | 35 | $4.48 | $79.48 | $379.48 |
| 64 core | 27 | $6.91 | $81.91 | $381.91 |







| **Compute power** | **Minutes** | **Cost of 1 build** | **Partial dev cost** <br>**(1 dev, 30 mins)** | **Partial dev cost** <br>**(5 devs, 30 mins)** | **Partial dev cost** <br>**(1 dev, 15 mins)** | **Partial dev cost** <br>**(5 devs, 15 mins)** |
| --- | --- | --- | --- | --- | --- | --- |
| 2 core | 310 | $2.48 | $39.98 | $189.98 | $21.23 | $96.23 |
| 4 core | 173 | $2.77 | $40.27 | $190.27 | $21.52 | $96.52 |
| 8 core | 92 | $2.94 | $40.44 | $190.44 | $21.69 | $96.69 |
| 16 core | 55 | $3.52 | $41.02 | $191.02 | $22.27 | $97.27 |
| 32 core | 35 | $4.48 | $41.98 | $191.98 | $23.23 | $98.23 |
| 64 core | 27 | $6.91 | $44.41 | $194.41 | $25.66 | $100.66 |









## 








