---
title: "A deeper dive into optimal PHP-FPM settings"
notion_id: c814ad93-48c2-47be-8f0a-b64c348ff12b
notion_url: https://app.notion.com/p/A-deeper-dive-into-optimal-PHP-FPM-settings-c814ad9348c247be8f0ab64c348ff12b
last_edited: 2024-04-12T17:48:00.000Z
source_url: https://geoligard.com/a-deeper-dive-into-optimal-php-fpm-settings
tags: ["Article", "Geoligard (Goran Popović)", "English", "PHP", "SysAdmin"]
---
![image](https://geoligard.com/storage/posts/March2024/PxUugm2CRgm42DGUR2yw.jpg)

single blog post image

In most cases, `PHP-FPM` settings aren't something an average developer would be in a position to examine more closely. That's ok; not everyone wants or has to spend time dealing with that kind of adjustment on the server.

Besides that, these days there are managed third-party solutions (`Laravel Forge`, `Ploi.io`, etc.) that can spin up a server for you, install all of the dependencies, including the PHP-FPM, and you just have to worry about deploying your code from their dashboard. Maybe you have a designated `DevOps` in your company or a senior developer that takes care of that kind of task. Or if you were in a position to actually set up PHP-FPM yourself, it's likely that you skimmed a few articles, made minor adjustments, or just used the defaults. And that's to be expected; there usually isn't enough time to perform a deeper dive into every server setting, especially if that's just a part of your overall assignments.

But after a while, as your app and its code evolve and as you get more and more users, you might notice that the server's become sluggish, the requests are taking more time before being processed, the memory consumption is hitting the server limits, or maybe your whole server just crashed.

Since something similar occurred on one of my servers relatively recently, I wanted to take the time and try to understand a bit better how PHP-FPM works and how different settings affect it. I went through a lot of articles on the subject, discussions, and comments, and then did some tests of my own to confirm some of the claims. Here is what I learned.

## Troubleshooting

If the issue is related to PHP-FPM, there are a few things we can do. First things first, we should check the PHP-FPM logs to see if there are any warnings. What we are interested in, in particular, are warnings about the `max_children` option. The PHP-FPM master process will spawn as many child processes as needed until the max children option is reached. Each child process is capable of processing, for example, a single request to your application at a time. So, if your `max_children` option is set to something like `5` and, for example, `10` users are simultaneously interacting with your app and sending requests, most likely you will see something like this in your logs:

```plain text
WARNING: [pool www] server reached pm.max_children setting (5), consider raising it

```

This will cause some of your requests to be delayed until enough child processes have been freed up. This is an easy command you can use to check if that kind of warning is occurring in your logs. For example, if you are using PHP-FPM 8.2, it might look something like this:

```plain text
sudo grep max_children /var/log/php8.2-fpm.log.1 /var/log/php8.2-fpm.log

```

Keep in mind that the path to the log on your system might be different, so you should double check it. Also, besides replacing your PHP version in the command, it's possible that the version isn't specified in your case and that you need to omit the version and just use `php-fpm`, like so:

```plain text
sudo grep max_children /var/log/php-fpm.log.1 /var/log/php-fpm.log

```

The same logic will apply each time I mention a command related to `php-fpm`. I'll always be using `php-fpm8.2` or `php8.2-fpm`, since that is my version, but in your case it could be something like `php-fpm7.4` (`php7.4-fpm`) or just `php-fpm`.

The easiest way to get the current values of our PHP-FPM settings without actually reading the config file is by utilizing this command:

```plain text
sudo php-fpm8.2 -tt

```

With it, we can easily find the `pm.max_children` line and confirm that indeed the `max_children` option is set to `5`.

```plain text
[19-Mar-2024 22:48:10] NOTICE:  pm.max_children = 5

```

Another thing that might interest us is to check the memory consumption on the server. By using a command like `htop` and sorting the processes by memory, we could see if server memory limits are being hit and if PHP-FPM processes are the ones that are consuming the most memory.

This could, for example, happen because of the too high `max_children` value, meaning that too many child processes have been spawned, they are all being used, and there just isn't enough memory on the server for all of them. Or, if the memory usage drops after the PHP-FPM is restarted and then gradually rises up to the limit after a while, that would usually point to a memory leak in your code. While it would be ideal to detect the memory leak and resolve the issue, locating memory leaks can sometimes be challenging, especially in big projects, and they could even be introduced by third-party libraries that might be vital for your app.

The first issue can be resolved with optimal settings, and there are some things you can do on the PHP-FPM's side regarding memory leaks as well, which we will cover a bit later.

By the way, the command to restart PHP-FPM might look like the following (but it might be different in your case, so double-check it):

```plain text
sudo service php8.2-fpm restart

```

As mentioned above, restarting PHP-FPM might provide you with a quick (but not permanent) win in case of a memory leak and buy you some time until you fix the leak or make the necessary adjustments.

## Configuring the process manager

Finally, we are ready to actually address the PHP-FPM config files and see what can be done in order to improve our particular setup. To edit the main config file, you might use a command like this:

```plain text
sudo nano /etc/php/8.2/fpm/pool.d/www.conf

```

There you can find all sorts of settings, but we will go through the most important ones that might affect the performance. The first thing we need to decide is how the process manager will control the number of child processes. There are 3 options here: `static`, `dynamic`, and `ondemand`. In most cases, the `dynamic` will be set by default. How do these options differ from one another? Let's say you determined that a maximum of 10 child processes is optional for your server.

`Static` would keep all 10 processes up at all times and is considered the fastest since all of the defined processes are already up and running, and there is no need to fork new ones as the load increases. But that also means that they will be consuming 10 processes worth of memory even if no one is visiting your website.

With `dynamic`, you could fine-tune your setup; for example, start 3 processes right away; if the load increases, fork up to 10 child processes; but when the load decreases, reduce the number to 6 processes that are waiting on connections. This option is a middle ground between memory consumption and the speed at which your app will respond to any requests, at least in theory.

And the last option, `ondemand`, means that no child processes will be spawned to start with; as the load increases, up to 10 of them will be created, and when the load decreases, you might end up again with no child processes running in the background. This option is ideal (again, in theory) for small and medium-sized apps that don't have that much traffic, for staging environments, or for servers where multiple tenants share resources. Since the child processes are recycled all the time, this option can help control any memory leaks you might be experiencing since the process will be terminated before there is time for memory to accumulate. The drawback to this option is that new processes will have to be forked all of the time, and that might affect your performance and ability to respond to requests more quickly.

Before setting one of the three options, we need to figure out the maximum load for all PHP-FPM processes. We need to determine the maximum number of child processes the server will be able to handle and set the `max_children` value appropriately. How are we going to do that? It turns out that's a bit tricky, because ideally you would have to determine how much memory a single child is using on average. The thing is that multiple processes can and usually will share some of the memory, so it is hard to pinpoint the actual memory usage of a single process.

There are a lot of scripts and articles in circulation that explain how you can calculate average memory consumption for a PHP-FPM process, but most of them didn't make much sense to me since I was getting a lot higher values than I was expecting to get.

One that seemed about right is a Python script that is mentioned in a few of those articles. You can run this command to calculate the `total` amount of memory used per program:

```plain text
cd ~ &&
wget https://raw.githubusercontent.com/pixelb/ps_mem/master/ps_mem.py &&
chmod a+x ps_mem.py &&
sudo python3 ps_mem.py

```

Note that I'm using `python3` in the last row; in your case, it might be just `python`. After running the script, you might get something like this:

```plain text
2.1 GiB + 127.5 MiB = 2.2 GiB       php-fpm8.2 (31)

```

So, this would mean that `31` PHP-FPM processes are using `2.2 GB` of memory, which would in turn mean that a single process is using about `73 MB`. Another useful command that you can always run to check the number of idle and active processes is:

```plain text
sudo service php8.2-fpm status -l

```

By checking this line here, we can see the current state of the child processes:

```plain text
Status: "Processes active: 0, idle: 30, Requests: 56116, slow: 0, Traffic: 0req/sec"

```

Here we don't have any active processes; there are `30` idle ones, and there is also `1` main process running, making it `31` total, the same as the Python script reported to us.

Keep in mind that the Python script will also report total memory usage. You can always run `htop` or `free -hl` to check current memory usage on the server and see if these numbers make sense to you and if they approximately match.

Another thing to note is that if you do have a memory leak, your memory for a single process might fill up to the `memory_limit` defined in your `php.ini` file, which is usually `128 MB` by default. So if you want to be on the safe side, you can use that value as an average value for a single PHP-FPM process.

OK, finally, we can calculate your `max_children` value. Let's suppose we have a server with `8 GB` of RAM. We know that all other programs on the server are using `2 GB`. Which leaves us with `6 GB`. And let's say we want to leave `1 GB` of buffer in case anything unexpected happens, in case some of our processes start using more memory in the future as our app grows or we decide to add some new processes. That leaves us with `5 GB` to redistribute to PHP-FPM. So, as we calculated previously, a single process is using about `73 MB` of memory. We just need to divide `5 GB` with those `73 MB` to get the `max_children` value:

```plain text
5120 (MB) / 73 (MB) = 70.14

```

So we got that the `max_children` value on our server should be `70`. That's great; no matter what option we choose for the process manager (`pm`), PHP-FPM will spawn a maximum of 70 child processes.

Now, if we decide to use the `pm = static` option, there aren't any additional options we should set. 70 child processes will be spawned right away, and they will be ready to handle requests. But remember the tradeoffs: this means `5 GB` of RAM will be used at all times by these processes.

Next, if we plan to use the `pm = ondemand` option, there is just one additional setting we should consider using, and that's `pm.process_idle_timeout`. Since child processes are spawned and terminated all the time when using the `ondemand` option, the `process_idle_timeout` settings tell PHP-FPM when to terminate a child process if it's idle (not actually being used). By default, that is set to `10 seconds`, but you can amend that value as you please if needed, although the default is pretty sensible.

And lastly, if we use `pm = dynamic`, there are a few additional settings to reconsider. First, `pm.start_servers` is the number of child processes that will be spawned right away when you start or restart PHP-FPM. Next, we have `pm.min_spare_servers`. That will set the minimum amount of idle child processes. And in the end, there is a `pm.max_spare_servers` setting that will determine the maximum amount of idle child processes.

Let's say we set these values:

```plain text
pm = dynamic
pm.max_children = 70
pm.start_servers = 20
pm.min_spare_servers = 20
pm.max_spare_servers = 40

```

Here is how that works in practice. So, in the example above, we have set `start_servers` to `20`. As soon as PHP-FPM starts, 20 child processes are spawned, and they are taking the memory that 20 processes would normally use. If there are requests coming to the app, some or all of those 20 processes will be active and handling those requests. If there is no traffic, those 20 processes will be idle, waiting on requests; they won't be terminated, and they will still be consuming memory.

Setting `min_spare_servers` to a value that is lower than `start_servers` (for example, 15) doesn't make much sense to me since 20 child processes will be spawned right away, and even if they are idle, the master process won't terminate 5 child processes to reach that minimum of 15. And you can't set `min_spare_servers` to be greater than `start_servers`, so the best course of action to me seems to be setting the `min_spare_servers` to be equal to the value of `start_servers`.

Now, if there are a lot of requests coming in and 20 child processes aren't enough to handle those requests, the master process will spawn additional child processes up to the `max_children` children value, in this case 70. And let's say 70 child processes have been spawned to handle a surge of requests. After a while, things normalize, and 70 processes are no longer needed; most or all of them become idle. In that case, the master process will terminate idle child processes up to the `max_spare_servers` value, in this case 40. And then you are left with 40 idle processes that won't be further terminated and are consuming memory that 40 child processes normally would.

So that's one of the things to keep in mind when setting these values: if there is a need to spawn more processes than `start_servers`, you will be left with that many child processes (up to the `max_spare_servers` value) running after the surge passes. If there is a need to spawn 30 child processes, you will be left with 30 of them running; if there is a need to spawn 50 child processes, you will be left with 40 (because of the `max_spare_servers`) of them running after a while. So, if you don't want to potentially end up with 40 child processes running in the background, you might want to consider lowering that value, or even keeping it the same as the `start_servers`). All of the things mentioned above will be true until you restart PHP-FPM. After the restart, 20 child processes will be spawned again based on `start_servers`.

In quite a few articles, there is a formula that suggests setting the `start_servers`, `min_spare_servers` and `max_spare_servers` values based on the number of CPU cores for optimal performance. The formula goes something like this:

```plain text
pm.start_servers = number of CPU cores x 4
pm.min_spare_servers = number of CPU cores x 2
pm.max_spare_servers = number of CPU cores x 4

```

The formula was supposedly constructed on some assumption about how many processes a single CPU core can handle concurrently. I'm not sure who started this trend or how those multipliers were derived exactly, but one thing seems a bit off to me, and that is setting the `min_spare_servers` to a value lower than `start_servers`, which, as I explained above, will cause the `min_spare_servers` value to never be utilized. Also, in my tests (which I'll address later on), I didn't notice any significant performance boosts from using this approach. So I don't think that this formula is something that should be blindly adopted but rather adjusted accordingly to your circumstances.

The best advice I can give regarding `max_children` and `dynamic` related settings is to monitor and fine-tune them as you go. Every situation will be different; your resources will be different; your load will be different; and your overall strategy will be different. Using the guidelines above, start with values that make sense to you, and as your app grows and changes, adjust the settings.

In this section, there is one more option worth noting, and that is `pm.max_requests`. If you do have a memory leak, this setting might be helpful to recycle your child processes after a certain number of requests. The default is set to `0` meaning that the processes won't be terminated based on that option. A sensible value might be 500 or 1000 requests, depending on your use case. If you set it to 500, for example, after 500 requests have been processed by the child process, it will be terminated (thus freeing up any accumulated memory) and then re-spawned.

## Testing it all out

I've decided to do a couple of performance tests in order to confirm the theory that `static` should be the fastest when processing requests since no child processes are being forked on the fly, `dynamic` should hold the middle ground since part of the processes are already running and part of them are being forked when needed, and `ondemand` should be the slowest since it spawns and terminates processes non-stop. Here are the results.

I've used `ApacheBench` to perform these tests, and the requests were sent from another server as it is recommended, meaning the test requests weren't sent from the same server that was being tested. The server that was being tested had `16 GB of RAM`, and `4 CPU cores`. PHP-FPM was coupled with NGINX. The requests went through a Laravel app. The value for `max_children` was `80` in all test cases. The value I compared was the response time of `90%` of the requests. And in the tables below, I will only reference the difference in milliseconds between different `pm` options, `0ms being the fastest one`.

Here is an example of the command that was used for testing:

```plain text
ab -n 1000 -c 10 https://example.com

```

In the example above, we would be sending 1000 requests with a concurrency of 10 requests at a time.

Let's review the first test. I've wanted to see how different `pm` options will affect the response times when the `max_children` limits are obviously being reached. So, I've sent `25000` requests with a concurrency level of `1000` requests.

These were the additional values when the `dynamic` was used (which will be later referenced only in this format, 20/20/40):

```plain text
pm.start_servers = 20
pm.min_spare_servers = 20
pm.max_spare_servers = 40

```

The results were as follows:

| Static | Dynamic | On demand |
| --- | --- | --- |
| +1223ms | +845ms | 0ms |

Turns out `ondemand`, which in theory should be the slowest option, was in fact the fastest in this test, and the `static` was unexpectedly the slowest, with over a second behind the `ondemand` setting.

In the next test, I've sent `10000` requests with a concurrency level of a `100`. I've tested with two different `dynamic` settings, one being the one from the first test (`20/20/40`) and the other one utilizing the formula based on the CPU cores (`16/8/16`). Here are the results:

| Static | Dynamic (20/20/40) | Dynamic (16/8/16) | On demand |
| --- | --- | --- | --- |
| 0ms | +14ms | +2ms | +24ms |

This time, `static` was the fastest, closely followed by the `formula-based dynamic` option, and `ondemand` was the slowest, as you would expect (based on theory). But in general, for most websites, `+24ms` isn't that much of a performance boost, and the differences between different options aren't that great.

And in the last test, I've sent only `2000` requests with a concurrency level of `16`. Again, I've used two `dynamic` settings (`20/20/40` and `16/8/16`). Let's see the results:

| Static | Dynamic (20/20/40) | Dynamic (16/8/16) | On demand |
| --- | --- | --- | --- |
| +2ms | +7ms | +10ms | 0ms |

On this scale, `ondemand` triumphed once again, `static` being a close second, and the `formula-based dynamic` being the last. This time, the differences between options were even less significant. The largest difference between the first place and the last one was just `10ms`.

We got a bit of an unexpected result in the end. The tests showed that the `ondemand` option was the best choice when the number of concurrent requests greatly `exceeded` our `max_children` value and when the number of concurrent requests was significantly `lower` than the `max_children` value. `Static` was the best choice when the number of concurrent requests was relatively close to the `max_children` value. But we should take into consideration that the differences in response times in the second and even more in the third test were quite low. Also, if these same tests were run again, it's quite possible that the places on the scoreboard would shift.

So, these results, of course, shouldn't be observed as something set in stone, and the same applies to theory that describes how different settings affect performance. I think that on modern servers, forking a new child process isn't an expensive operation anymore, which could significantly affect response times. At least not on the scale on which the tests were performed. That is why we can see that the `ondemand` setting isn't something to be easily dismissed, even when it comes to the speed at which the requests are processed.

The best way to move forward would be to do your own tests, since your load, settings, and operations being performed per request could be quite different, and then, based on those tests, apply the settings that seem the most performant.

## Additional settings

There are also a couple of additional settings we should go through that might prove useful in case something goes wrong with PHP-FPM or in case you need to track down slow requests.

In order to enable slowlog, which will of course log slow requests, we need to edit the same config file as before:

```plain text
sudo nano /etc/php/8.2/fpm/pool.d/www.conf

```

Then locate the slowlog section:

```plain text
slowlog = /var/log/php8.2-fpm.log.slow

```

And uncomment it. There are a couple of more related options there that you should consider uncommenting. The first one is the `request_slowlog_timeout` which is by default set to `5 seconds`. If you want to log only requests that take, for example, `3 seconds` and more, you should uncomment and amend that value. The second one is `request_slowlog_trace_depth` which is set to `20` by default. In Laravel apps, this value might be too low to go through all of the vendor functions and get to the piece of code that is actually being called in your controller, for example. So I think in most cases, `50` should be fine, but double-check if that works for you.

Here is what the whole slowlog setup might look like in the end:

```plain text
slowlog = /var/log/php8.2-fpm.log.slow
request_slowlog_timeout = 3s
request_slowlog_trace_depth = 50

```

Lastly, there is another config file we can edit and control what will happen in case our child processes start failing for whatever reason. Here is an example of how you can edit this file:

```plain text
sudo nano /etc/php/8.2/fpm/php-fpm.conf

```

In that file, we are interested in 3 options that are interconnected, and all of them are set to `0 by default` and commented out. If you plan on using them, make sure to `uncomment` them first. After that, you can set them to these values:

```plain text
emergency_restart_threshold = 10
emergency_restart_interval = 1m
process_control_timeout = 10s

```

The values used are something that you'll probably come across in some other articles as well. The first two settings tell PHP-FPM that if `10` child processes fail within a `minute`, the PHP-FPM should `restart` itself automatically. The third setting means that a child process will wait for `10 seconds` before acting on a signal sent from a master process. So in case a master process sends a KILL signal to the child process, it will have 10 seconds to finish its tasks before exiting. Of course, you can fine-tune these values to fit your needs.

Restarting PHP-FPM in a case of failure might resolve some problems, but if the problem is related to something that will reoccur even when the PHP-FPM is restarted, it will keep restarting until you can figure out what exactly happened. So you should decide for yourself if you want PHP-FPM to fail completely when something unexpected happens or if you want it to restart itself.

## Conclusion

Whew! We covered quite a few things there. We discussed the best values for the `max_children` option. We explored the advantages and flaws of different `process manager` settings and how we can test them. We also went through some additional settings that might prove useful when debugging slow requests or handling failures. I hope this was useful and that you'll be able to use the information from this article as a starting point to better monitor and fine-tune PHP-FPM on your servers.

Enjoyed the article? Consider subscribing to my [newsletter](https://geoligard.com/newsletter) for future updates.
