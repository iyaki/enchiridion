---
title: "Seperate your slow backend from the important frontend traffic with PHP-FPM"
notion_id: 979321ad-70db-4a14-ba06-df912da20857
notion_url: https://app.notion.com/p/Seperate-your-slow-backend-from-the-important-frontend-traffic-with-PHP-FPM-979321ad70db4a14ba06df912da20857
last_edited: 2026-09-21T17:21:00.000Z
source_url: https://tideways.com/profiler/blog/seperate-your-slow-backend-from-the-important-frontend-traffic-with-php-fpm
tags: ["English", "PHP", "Article", "tideways Blog"]
---
[https://tideways.com/profiler/blog/seperate-your-slow-backend-from-the-important-frontend-traffic-with-php-fpm](https://tideways.com/profiler/blog/seperate-your-slow-backend-from-the-important-frontend-traffic-with-php-fpm)

PHP Performance20.12.2016

If you are using Magento, Shopware, Oxid, a CMS or any other off the shelf software, then usually they ship both the frontend and the backend in a single application. For self-developed applications with Symfony or other frameworks this is often the case as well.

The backend/admin may then include several slow reporting, administrative operations and data exports that can take a long time. This could congest the webservers processing queue by lowering the throughput for your customers that might just click on the checkout button and see a 502 Gateway error. We described this behaviour in the last blog post on What is the best value of PHPs max_execution_time?.

Putting frontend and backend on different servers is one solution, but this can be to costly for most use-cases.

A simple solution for this problem is using different PHP-FPM pools for the frontend and backend, each with their own configuration for the maximum number of allowed requests. This is also a solution for some of the problems discussed in the post on “What is the best value for max execution time?”.

## Magento Admin and Frontend Example

How does this look like? Lets use Magento as an example, you can configure two pools in php-fpm.conf:

```plain text
; php-fpm.conf [frontend] listen = /var/run/php-fpm-frontend.sock pm = static pm.max_children = 50 [backend] listen = /var/run/php-fpm-backend.sock pm = ondemand pm.max_children = 5 pm.process_idle_timeout = 5
```

The frontend is configured for maximum of 50 simultaneous requests and the backend for a maximum of 5 simultaneous requests. The backend workers are created on demand and the frontend workers are static to avoid forking overhead. I will discuss the differences between PHP-FPM pool configurations in a future blog post.

You can then modify the Nginx vhost configuration for the Magento installation with the following switch:

```plain text
server { // .... set $fpm_socket "unix:/var/run/php-fpm-frontend.sock"; if ($uri ~* "^/admin/") { set $fpm_socket "unix:/var/run/php-fpm-backend.sock"; } location ~ .php$ { // ... fastcgi_pass $fpm_socket; } }
```

Based on the ^/admin path in the request uri it will select the different PHP FPM pool now and frontend and backend will not compete and steal each other resources anymore.

### About the author

Benjamin Founder & CEO

I am the founder and CEO of Tideways. I started the company over 10 years ago with the mission to move the PHP ecosystem forward, starting with performance. As managing director, I work across product, strategy, and the day-to-day of building a developer-focused SaaS business.

I’m a core contributor to the Doctrine open-source project and a founding board member of the PHP Foundation, which reflects my long-standing commitment to the PHP ecosystem. I particularly enjoy working at the intersection of application performance, developer experience, and the open-source community that makes PHP what it is. Outside of work, I enjoy board games, hiking, and coffee.
