---
title: "TIL: Docker log rotation | nicole@web"
notion_id: 30e54f1c-7d23-817c-ac5c-f5eda98847dc
notion_url: https://app.notion.com/p/TIL-Docker-log-rotation-nicole-web-30e54f1c7d23817cac5cf5eda98847dc
last_edited: 2026-09-18T00:53:00.000Z
source_url: https://ntietz.com/blog/til-docker-log-rotation/
tags: ["Article", "Note", "ntietz's Blog", "English", "DevOps", "Docker", "System Design / Software Architecture", "Technical Debt", "Error Handling"]
---
Last week[[1]](https://ntietz.com/blog/til-docker-log-rotation/?utm_source=tldrdevops%2F#fn-lwil), when I went to publish my blog post, I ran into a surprising error: I was out of disk space. My server is used only for hosting a couple of small static sites, so I was surprised. None of the content is very large, why is the disk full?

A little investigation found the culprit. Starting from `/` and then drilling in using `du -h -d 1 .` showed me that a Docker folder was using most of the server's 25 GB disk. None of my container images are very large, so I checked in side and found a few log files that were larger than 10 GB each.

It turns out, Docker doesn't automatically rotate log files! As long as a container exists, the logs will keep growing for it. This means even if you stop a container and start it again, the logs are still there and getting bigger. I'd not thought about this before, but it turns out that when my blog is seeing heavy traffic, the logs can grow in order of megabytes per hour. And that really adds up over time.

First I did a quick check of how the logs are configured to start with. You can see the log configuration by using `docker inspect`.

```plain text
$ docker inspect --format='{{.HostConfig.LogConfig}}' my-poor-container
{json-file map[]}
```

And my container's logging was totally unconfigured! That explained a lot.

Now the fix is pretty quick. The [docs](https://docs.docker.com/engine/logging/drivers/json-file/) show us an example that works well enough here. `/etc/docker/daemon.json` didn't exist yet, so I created it and added this log configuration in.

```plain text
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m",
    "max-file": "3"
  }
}
```

The original example had `10m` for log files, but I want a little more than that. I have the disk space, and I'd like longer to investigate logs before they are truncated away.

After setting that up, I restarted the docker daemon by calling `systemctl restart docker`. But logs don't rotate yet, no! The docs told us that this applies for _new_ containers after Docker is restarted, but not for existing containers. So the final step was to stop and remove any containers I wanted rotation to work on, then recreate them.

After that, a quick check, and we've got log rotation.

```plain text
$ docker inspect --format='{{.HostConfig.LogConfig}}' my-happy-container
{json-file map[max-file:3 max-size:100m]}
```

Don't be like me, don't forget to rotate your logs!

Or, do forget. Focus on the things you enjoy, and do just enough of the other things to make it work. You can always hire someone else to solve some of the annoying, tedious, or difficult problems for you. (Hi. [Hire me!](https://cuteandfuzzylogic.com/))

1. 
2. 

_If you're looking for help on a software project, please consider _[_working with me_](https://cuteandfuzzylogic.com/)_!_

Please share this post, and subscribe to the [newsletter](https://ntietz.com/newsletter/) or [RSS feed](https://ntietz.com/atom.xml). You can email my personal email with any comments or questions.
