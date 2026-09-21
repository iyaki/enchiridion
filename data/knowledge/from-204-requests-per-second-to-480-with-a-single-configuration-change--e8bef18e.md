---
title: "From 204 requests per second to 480 with a single configuration change"
notion_id: e8bef18e-4651-42e8-9994-19f5cd97d5e1
notion_url: https://app.notion.com/p/From-204-requests-per-second-to-480-with-a-single-configuration-change-e8bef18e465142e8999419f5cd97d5e1
last_edited: 2023-01-13T17:02:00.000Z
source_url: https://getparthenon.com/blog/php-performance-tunning-from-204-to-480-with-a-single-config-change/
tags: ["English", "PHP", "Article", "Tutorial"]
---
I'm currently developing a [cookieless Free Live Chat SaaS - Blether.chat](https://blether.chat/?utm_source=getparthenon&utm_campaign=content_marketing) and as part of the development process, I decided to stress test the application and really see the difference PHP OP Cache preload made to the request per second in a real-world application. This led me down a path that resulted in me making a configuration change unrelated to the OP Cache resulting in an almost 100% increase in throughput and adding 200+ requests per second to my server's capabilities.

# What is PHP OP Cache?

First, I'll quickly explain what the OP Cache in PHP is and what preload is. PHP is an interpreted language which means at every run, the PHP interpreter will read the PHP code and convert it into operations. The OP allows the interpreter to cache the operations, so it doesn't need to reinterpret the same code repeatedly. This enables PHP to run faster as it removes a lot of work for the interpreter.

In PHP 7.4, they added the ability to include a preload for the opcache. This means you can add code that is to be cached for every PHP process without needing to be parsed, etc. Again another performance enhancer.

## Server info

For this test, I was testing on a 64GB RAM server with an AMD Ryzen 7 1700X Eight-Core Processor that had 16 CPU threads rented from [Hetzner](https://hetzner.com/).

With the testing server being a 32GB 16 CPU thread digital ocean droplet in Holland.

## Test - Start with OP Cache preloading disabled

First, I ran a load test using [k6 load testing tool](https://k6.io/) on the application without any opcache enhancements. This resulted in 187 requests per second (rps). Personally, I was quite happy since I hadn't done any optimisation so 187 rps to start with is a good place to start. The response time was 287ms, which I wasn't so happy with but I know there are lots of optimisations I can do so I'm not overly worried.

**Result RPS** 187 **Average Response** 265ms

## Test - Op Cache preloading Enabled

I then enabled PHP OPcache with the following configuration.

```plain text
; php.ini
opcache.preload=/var/www/web.blether.chat/current/config/preload.php
; required for opcache.preload:
opcache.preload_user=www-data
opcache.memory_consumption=256
opcache.max_accelerated_files=100000
opcache.validate_timestamps=0

```

A quick overview of these configurations

| name | description |
| --- | --- |
| opcache.preload | This is the php file that is to be preloaded into opcache |
| opcache.preload_user | This is the user that the opcache is going to be ran under. This should be the same as php-fpm. |
| opcache.memory_consumption | How many megabytes the opcache should consume |
| opcache.max_accelerated_files | How many files can be opened and cached |
| opcache.validate_timestamps | If the opcache should check the timestamps of the files to see if there have been any changes. In production, there should be no changes; therefore, this check can be disabled |

Once this change was made, I reran the test from before.

**Result RPS** 204 **Average Response** 244ms

There was ~10% improvement in requests per second and response times.

## Something was off

While these tests were happening I was monitoring the server activity via `htop` and New Relic. The tests used 20% of the CPU and only 10% of the memory. New Relic was saying the PHP execution time during these was **20ms**. Yet k6 was saying the time to the first byte was, on average, **244ms**. I played around both with increasing and decreasing the number of virtual users and nothing changed. Something was clearly not optimised correctly. I assumed it was something in the operation of handling the HTTP requests.

First, I tuned Nginx. This made no difference. I then looked into PHP-FPM. I found a blog on [PHP-FPM tuning](https://haydenjames.io/php-fpm-tuning-using-pm-static-max-performance/), which was super valuable.

After reading this, I changed the PHP-FPM pm setting to use static children. PHP-FPM is a process manager to run PHP it has processed, and Nginx forwards requests to PHP-FPM, and it handles them. Each process is ran as a child. By default, the children are created as they're needed. This is good for memory management purposes and allows you to use less memory overall. By switching it to static, then it has a static pool of processes to call. These processes are always running, and therefore, the memory usage is higher and constant. By using static that means children don't need to be scaled up or down as needed and PHP-FPM can just call them when one is available.

I made the following config change to PHP-FPM.

The difference was clearly noticeable on PHP-FPM restart as instead of being near instant, it took about 15 seconds to restart. The memory usage on the server jumped to 20GB out of 64GB that was available.

## Test - Opache and PHP-FPM pm.static

I reran the test with opcache still enabled. The performance on the server was instantly different. All the CPUs were at max usage. Request per second jumped up to *_480_.

**Result RPS** 480 **Average Response** 109ms

Almost a 100% increase in both the throughput and run time.

## Test - Disable OP Cache preloading and PHP-FPM pm static

Now I've seen the difference PHP-FPM pm being in static makes, I wanted to see what the difference would be if I ran the test with OP Cache preloading disabled.

**Result RPS** 442 **Average Response** 112ms

There was a clear drop in requests per second, almost 10%, with the average response time not making much change.

## Overall Results

| Test | Request Per Second | Average Response time |
| --- | --- | --- |
| PHP-FPM pm = Dynamic & no preload | 187 | 265ms |
| PHP-FPM pm = Dynamic & with preload | 204 | 244 |
| PHP-FPM pm = Static & no preload | 442 | 112ms |
| PHP-FPM pm = Static & with preload | 480 | 103ms |

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Bigger is better.

## Conclusion

This little test showed me the true power of correctly tuning PHP-FPM. By switching the pm mode from dynamic to static, I got an almost 100% increase in throughput and cut the response time in half. This will provide me a solid foundation to build my [cookieless Free Live Chat SaaS - Blether.chat](https://blether.chat/?utm_source=getparthenon&utm_campaign=content_marketing).

A shameless plug, If you're in the market for a live chat on your website and hate cookie banners like the rest of us, think about checking out [Blether.chat](https://blether.chat/?utm_source=getparthenon&utm_campaign=content_marketing). It's currently in Beta, so everything is free, and it will be using a freemium model.
