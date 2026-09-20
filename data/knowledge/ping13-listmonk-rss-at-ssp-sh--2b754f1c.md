---
title: "ping13/listmonk-rss at ssp.sh"
notion_id: 2b754f1c-7d23-8169-85ea-f946240511ad
notion_url: https://app.notion.com/p/ping13-listmonk-rss-at-ssp-sh-2b754f1c7d23816985eaf946240511ad
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://github.com/ping13/listmonk-rss
tags: ["English", "DevOps", "Automation", "Productivity", "Email", "Open Source", "Tool", "Article", "Github Blog"]
---
# 



## 

- 
- 
- 
- 
- 
- 

## 

- 
- 
- 

## 

### 

![image](https://github.com/ping13/listmonk-rss/raw/main/assets/C4/architecture.png)

### 

![image](https://github.com/ping13/listmonk-rss/raw/main/assets/C4/sequence_diagram.png)

## 

### 

1. 
2. 
3. 
4. 
5. 
6. 

### 



```

```



- 
- 
- 

### 



1. 
2. 
3. 
4. 



```

```

### 



1. 
2. 
3. 

****

****

## 

### 

| Variable | Description | Required |
| --- | --- | --- |
| LISTMONK_API_USER | Listmonk API username | Yes |
| LISTMONK_API_TOKEN | Listmonk API token | Yes |
| LISTMONK_HOST | Listmonk instance URL | Yes |
| LIST_NAME | Name of the mailing list in Listmonk | Yes |
| RSS_FEED | URL of the RSS feed to monitor | Yes |
| DELAY_SEND_MINS | Minutes to delay sending after creation (default: 30). In dry run mode, this is set to 10 years. | No |
| PUSHOVER_USER_KEY | Pushover user key for notifications (optional) | No |
| PUSHOVER_API_TOKEN | Pushover API token for notifications (optional) | No |
| GH_REPOSITORY | GitHub repository in "owner/repo" format | Yes |
| GH_TOKEN | GitHub token with repo scope for state storage | Yes |

### 



- 

## 





## 


