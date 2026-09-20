---
title: "From 204 requests per second to 480 with a single configuration change"
notion_id: e8bef18e-4651-42e8-9994-19f5cd97d5e1
notion_url: https://app.notion.com/p/From-204-requests-per-second-to-480-with-a-single-configuration-change-e8bef18e465142e8999419f5cd97d5e1
last_edited: 2023-01-13T17:02:00.000Z
source_url: https://getparthenon.com/blog/php-performance-tunning-from-204-to-480-with-a-single-config-change/
tags: ["PHP", "Article", "Tutorial", "English"]
---


# 





## 





## 





## 



```

```



| name | description |
| --- | --- |
| opcache.preload | This is the php file that is to be preloaded into opcache |
| opcache.preload_user | This is the user that the opcache is going to be ran under. This should be the same as php-fpm. |
| opcache.memory_consumption | How many megabytes the opcache should consume |
| opcache.max_accelerated_files | How many files can be opened and cached |
| opcache.validate_timestamps | If the opcache should check the timestamps of the files to see if there have been any changes. In production, there should be no changes; therefore, this check can be disabled |







## 











## 







## 







## 

| Test | Request Per Second | Average Response time |
| --- | --- | --- |
| PHP-FPM pm = Dynamic & no preload | 187 | 265ms |
| PHP-FPM pm = Dynamic & with preload | 204 | 244 |
| PHP-FPM pm = Static & no preload | 442 | 112ms |
| PHP-FPM pm = Static & with preload | 480 | 103ms |

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



## 




