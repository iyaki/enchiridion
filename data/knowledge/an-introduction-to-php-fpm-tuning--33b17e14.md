---
title: "An Introduction to PHP-FPM Tuning"
notion_id: 33b17e14-47ca-4705-a41c-c83b8f187c02
notion_url: https://app.notion.com/p/An-Introduction-to-PHP-FPM-Tuning-33b17e1447ca4705a41cc83b8f187c02
last_edited: 2023-04-25T14:16:00.000Z
source_url: https://tideways.com/profiler/blog/an-introduction-to-php-fpm-tuning
tags: ["Article", "tideways Blog", "English", "PHP"]
---
[https://tideways.com/profiler/blog/an-introduction-to-php-fpm-tuning](https://tideways.com/profiler/blog/an-introduction-to-php-fpm-tuning)

PHP Performance18.02.2020

BONUS: We have discussed this topic with an expert in the PHP community in our podcast:

PHP-FPM (or Fast Process Manager) offers several advantages over mod_php, with two of the most notable being that it is more flexible to configure and currently the preferred mode of running PHP by many in the community. However, if you’re using your package manager’s default configuration settings, then you’re likely not getting the most out of it.

In this post, I’m going to give a brief overview on how to improve PHP-FPM’s performance, by discussing PHP-FPM’s three process manager types, and which one is best to use under which circumstance.

PHP-FPM can use one of three process management types:

- static
- dynamic; and
- ondemand

Let’s look at what each one is in a little bit of detail.

## Static

Static ensures a fixed number of child processes are always available to handle user requests. This is set with pm.max_children. In this mode requests don’t need to wait for new processes to startup, which makes it the fastest approach.

Assuming that you wanted to use static configuration with 10 child processes always available, you would configure it in /etc/php/8.3/fpm/pool.d/www.conf (assuming you’re using Debian/Ubunut’s PHP-FPM’s default configuration file) as follows:

```plain text
pm = static pm.max_children = 10
```

To see if the configuration change has been effective, after restarting PHP-FPM, run pstree -c -H <PHP-FPM process id> -S <PHP-FPM process id>. This will show that there are ten processes available, as in the example below.

```plain text
php-fpm8.3-+-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3 |-php-fpm8.3
```

## Dynamic

In this mode, PHP-FPM dynamically manages the number of available child processes and ensures that at least one child process is always available.

This configuration uses five configuration options; these are:

- pm.max_children: The maximum number of child processes allowed to be spawned.
- pm.start_servers: The number of child processes to start when PHP-FPM starts.
- pm.min_spare_servers: The minimum number of idle child processes PHP-FPM will create. More are created if fewer than this number are available.
- pm.max_spare_servers: The maximum number of idle child processes PHP-FPM will create. If there are more child processes available than this value, then some will be killed off.
- pm.process_idle_timeout: The idle time, in seconds, after which a child process will be killed.

Now the fun part comes; how do you calculate the values for each setting? Sebastian Buckpesch offers the following formula:

Setting Value max_children (Total RAM – Memory used for Linux, DB, etc.) / process size start_servers Number of CPU cores x 4 min_spare_servers Number of CPU cores x 2 max_spare_servers Same as start_servers

We also need to set pm.process_idle_timeout, which is the number of seconds after which an idle process will be killed. If you don’t want to calculate this from scratch, we wrapped all this into a simple to use calculator tool to get the optimal FPM configuration settings.

Let’s say that our server has two CPUs, each with four cores, and 8GB RAM. If we assume that Linux and related daemons are using around 2GB (use free -hl to get a more specific value), that leaves us with around 6192MB.

Now, how much memory is each process using? As a rule of thumb, you need to look at the memory_limit configured in PHP. To calculate a more accurate number, there’s a Python script called ps_mem.py. After running it, using sudo python ps_mem.py | grep php-fpm, you’ll get output similar to the following:

```plain text
28.4 MiB + 33.8 MiB = 62.2 MiB php-fpm8.3 (11)
```

The first column is private memory. The second column is shared memory. The third column is the total RAM used. The fourth column is the process name.

From the above, you can see that the process size is 62.2MiB. So, feeding all of that information into our formula, we arrive at the following:

```plain text
# Round the result up. (8192 - 2000) / 62.2
```

Based on that, we come to the following settings values:

Setting Value max_children 100 start_servers 32 min_spare_servers 16 max_spare_servers 32

We’ll leave the pm.process_idle_timeout to the default of 10s. Assuming that we’re happy with these settings, then we’d configure it as follows:

```plain text
pm = dynamic pm.max_children = 100 pm.start_servers = 32 pm.min_spare_servers = 16 pm.max_spare_servers = 32 pm.max_requests = 200
```

You can also use memory monitoring tools such as Tideways on a regular basis to monitor how much memory your application is using.

## ondemand

ondemand has PHP-FPM fork processes when requests are received. To configure PHP-FPM to use it, we need to set pm to dynamic, and provide values for:

- max_children
- process_idle_timeout
- max_requests

max_requests sets the number of requests each child process should execute before respawning. The documentation suggests that this setting is helpful for working around memory leaks.

Assuming that we take the same settings as for dynamic, we’d configure it as follows:

```plain text
pm = ondemand pm.max_children = 100 pm.process_idle_timeout = 10s pm.max_requests = 200
```

## Which Configuration Is Right For You?

Honestly? The answer is: “it depends“, as it always depends on the type of applications that you are running. However, here are some suggestions as to which configuration to choose. To nail the right secondary configuration values for your FPM configuration, we recommend you to look into our simple calculator.

### Low Traffic Site

If you have a low traffic site, such as one hosting a backend control panel, such as cPanel, then use ondemand. Memory will be saved as child processes will only be spawned when they’re needed and killed off when they’re no longer required. As it’s a backend, users can wait a moment or two longer while a thread spawns to handle their request.

### High Traffic Site

If you have a high traffic website, then use static and adjust the settings based on your needs over time and your available hardware resources. It might seem like overkill to have a large number of child processes always ready to receive requests.

However, high-traffic sites need to respond as quickly as possible. Therefore, it’s essential to use static so that a sufficient number of child processes are ready to do so.

By using ondemand, child processes will likely consume too much memory being spawned and killed, and the startup delay will have a performance hit.

Using dynamic likely won’t be as bad, depending on the configuration. However, you may end up with a configuration that effectively mirrors static.

## Using Multiple Pools for Frontend/Backend

Now for one final recommendation: serve the frontend and backend of your website using different pools. Say you have an e-commerce site, perhaps powered by Magento. You can look at the application as being composed of two parts:

- A frontend where customers can browse and make purchases
- A backend, where admin staff manage the shop (such as adding/removing products, categories, and tags, and reviewing ratings)

When viewed this way, it makes sense to have one pool that serves the frontend and another that serves the backend and to configure each appropriately.

For what it’s worth, you could split up any application into multiple parts using this strategy, if it makes sense to do so. Here’s how to do so.

In /etc/php/8.3/fpm/pool.d/www.conf, add the following configuration:

```plain text
; frontend [frontend] listen = /var/run/php-fpm-frontend.sock user = www-data group = www-data listen.owner = www-data listen.group = www-data pm = static pm.max_children = 5; ; backend [backend] listen = /var/run/php-fpm-backend.sock user = www-data group = www-data listen.owner = www-data listen.group = www-data pm = ondemand pm.max_children = 5 pm.process_idle_timeout = 10s
```

This creates two pools, one for the frontend, and one for the backend. Both have the same user and group, but have different process manager configurations and are connected to via different sockets.

The frontend pool uses a static configuration with a small maximum number of child processes. The backend pool uses the ondemand configuration, also with a small number of configurations. These numbers are arbitrary, as they’re for the purposes of an example.

With that saved, for your NGINX vhost file, use the following configuration:

```plain text
server { listen 80; server_name test-site.localdomain; root /var/www/test-site/public; access_log /var/log/nginx/test-site.access.log; error_log /var/log/nginx/test-site.error.log error; index index.php; set $fpm_socket "unix:/var/run/php-fpm-frontend.sock"; if ($uri ~* "^/api/") { set $fpm_socket "unix:/var/run/php-fpm-backend.sock"; } location / { try_files $uri $uri/ /index.php; location ~ .php$ { fastcgi_split_path_info ^(.+.php)(/.+)$; fastcgi_pass $fpm_socket; fastcgi_index index.php; include fastcgi.conf; fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name; } } }
```

This creates a virtual host configuration that sends requests either to the frontend or backend pool, based on the location requested. Any requests to /api are sent to the backend pool, and all other requests are routed to the frontend.

## In Conclusion

That’s been a rapid introduction to tuning PHP-FPM for better performance. We’ve looked at the three different process manager configurations, their related settings, and discussed when each configuration makes sense. We then finished up by looking at worker pools.

### About the author

Benjamin Founder & CEO

I am the founder and CEO of Tideways. I started the company over 10 years ago with the mission to move the PHP ecosystem forward, starting with performance. As managing director, I work across product, strategy, and the day-to-day of building a developer-focused SaaS business.

I’m a core contributor to the Doctrine open-source project and a founding board member of the PHP Foundation, which reflects my long-standing commitment to the PHP ecosystem. I particularly enjoy working at the intersection of application performance, developer experience, and the open-source community that makes PHP what it is. Outside of work, I enjoy board games, hiking, and coffee.
