---
title: "Save your disk, write files directly into RAM with /dev/shmAndrew Quinn's TILs"
notion_id: 2b754f1c-7d23-813c-92a0-cf7868bf518d
notion_url: https://app.notion.com/p/Save-your-disk-write-files-directly-into-RAM-with-dev-shmAndrew-Quinn-s-TILs-2b754f1c7d23813c92a0cf7868bf518d
last_edited: 2025-11-26T18:57:00.000Z
source_url: https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/
tags: ["Hackernoon", "English", "Linux", "System Design / Software Architecture", "DevOps", "Performance", "Article", "Note"]
---


| [`1`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-0-1)[`2`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-0-2) | `❯ mount \| grep '/dev/shm'<br>tmpfs on /dev/shm type tmpfs (rw,nosuid,nodev,inode64)` |
| --- | --- |







| [`1`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-1-1)[`2`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-1-2)[`3`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-1-3)[`4`](https://hiandrewquinn.github.io/til-site/posts/save-your-disk-write-files-directly-into-ram-with-dev-shm/#hl-1-4) | `curl -s "https://en.wiktionary.org/wiki/$word" \| \<br>  pandoc --from=html --to=plain \| \<br>  sed -n '/^Finnish$/,$p' >/dev/shm/$word.txt \<br>  && vim -c "set nowrap" /dev/shm/$word.txt` |
| --- | --- |





1. 
