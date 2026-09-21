---
title: "Five Challenges for Running Reliable PHP Background Processes"
notion_id: 65988301-309e-44c7-8f89-f9e978ea7a6a
notion_url: https://app.notion.com/p/Five-Challenges-for-Running-Reliable-PHP-Background-Processes-65988301309e44c78f89f9e978ea7a6a
last_edited: 2026-09-21T17:21:00.000Z
source_url: https://tideways.com/profiler/blog/five-challenges-for-running-reliable-php-background-processes
tags: ["Article", "tideways Blog", "English", "PHP"]
---
[https://tideways.com/profiler/blog/five-challenges-for-running-reliable-php-background-processes](https://tideways.com/profiler/blog/five-challenges-for-running-reliable-php-background-processes)

PHP Performance04.03.2020

BONUS: We have discussed this topic with an expert in the PHP community in our podcast:

PHP isn’t typically thought of as a solution when creating worker or background processes, jobs that typically can last for an extended period. These can be tasks such as image processing, file repair, and mass email batch jobs.

Typically, PHP is linked with HTTP requests, requests which are short in duration and stateless in nature. However, just because of this enduring association, it doesn’t mean that PHP can’t be used for background processes. On the contrary.

If you’re considering developing background processes in PHP, then this post is for you. I’m going to step you through five points that you need to keep in mind. Hopefully, by the end, you’ll have less stress and effort as a result.

## Use Angel Processes

Background jobs can die.

They can die for many reasons, including out of memory exceptions, connection errors, as well as a number of other conditions. Given that this could happen and to avoid manual intervention to restart the job, it’s helpful to make use of an angel process.

An angel process offers three key advantages:

- Can start a job, if it’s not already started
- Can restart a failed job
- Provides debugging and logging capabilities. For example, Supervisor and Systemctl can centralize job logging and show a job’s status and history.

The angel process isn’t responsible for anything else. Consequently, the logic should be quite rudimentary, and its load should be quite light at any given time. However, there are a couple of conditions that you need to be aware of, when writing them.

## Handle Failed Jobs

No matter what language you write your jobs in, you have to know when they fail and be able to restart them when they do. PHP is no exception.

So, how do you do this? You could roll your own solution, such as a shell script, but there are many existing solutions already available. As a result, you should never need to write your own.

Here is a selection of them, delimited by operating system:

Package Description Operating System Supervisor Available for UNIX/Linux systems, Supervisor provides a simple, centralized, efficient, and extensible approach to controlling many processes. It’s been tested on Linux, Mac OS X, Solaris, and FreeBSD. You may be able to use it on Windows in combination with Cygwin. Systemd Available for UNIX/Linux systems, systemd was designed as a replacement for UNIX System V and the BSD init system. Initially developed at Red Hat, it includes features such as on-demand starting of daemons, snapshot support, and process tracking. Linux Gearman To quote the official documentation, it: ” provides a generic application framework to farm out work to other machines or processes that are better suited to do the work.” Written in C, it’s available for Linux and Windows. The Windows Service Control Manager It provides functionality similar to Supervisor and Systemd for Windows servers. Windows

For more information about supervisord, check out this post on Servers for Hackers. And for more information on Gearman and PHP, check out this post by Matthew Weier O’Phinney. If you’re considering using Gearman, be aware that it requires an angel process for its own workers. It’s too much to cover in this post, however I just wanted you to be aware of this fact.

## Monitor Background Jobs

As background jobs can last for quite some time, it’s helpful to be able to inspect their state when required. There’s often nothing worse than knowing that a job is still active, yet to have no information about how much or how little progress its made, how long it takes to complete, and any problems that it has encountered during processing.

Given that, it’s handy, though arguably not essential, to have some form of monitoring available. This could be something quite rudimentary, such as printing out a list of the currently active jobs, their start time, expected completion time, and the number of failures that they have encountered to the terminal.

The output might look like the following from Supervisor:

```plain text
$ sudo supervisorctl status alertActionWorker:alertActionWorker00 RUNNING pid 2349, uptime 0:01:35 createErrorWorker:createErrorWorker00 RUNNING pid 2345, uptime 0:01:35 createProfileWorker:createProfileWorker00 RUNNING pid 2346, uptime 0:01:35 notificationWorker:notificationWorker00 RUNNING pid 2350, uptime 0:01:35 recordMeasurementsWorker:recordMeasurementsWorker00 RUNNING pid 2339, uptime 0:01:35
```

It could also go as far as a comprehensive dashboard, such as the monitoring in Tideways:

Tideways Background Jobs Monitoring

Whichever way you go, ensure that your job is storing some form of statistical information that can be periodically polled and collated.

Many vendors provide built-in functionality to store this information, such as Microsoft’s Azure platform. Alternatively, you can store your own in one of the open-source databases, such as PostgreSQL, MySQL, or SQLite.

## Tune Memory Usage and Restart Jobs Periodically

PHP wasn’t designed with long-running processes and tasks in mind. It was designed for shorter tasks based around the nature of an HTTP request. As such, long-running tasks can be problematic, mostly because of memory consumption.

There are a couple of things that you can do to help reduce memory consumption. These include:

- Garbage Collection Analysis: Performing garbage collection analysis to help identify how your worker process uses variables and how PHP’s garbage collector interacts with your script.
- Avoid Global and Class Static Variables.
- Use cursor-like behavior for iterating over information, such as arrays, lists, and collections. By doing this, only a limited amount of memory is used at any one time.

By doing these three things, the likelihood of writing code that leaks memory is reduced, which in turn reduces the likelihood of needing to kill a process arbitrarily.

## Implement Exponential Backoff When Auto Restarting

Long-running jobs will inevitably encounter some form of network collision or outages that temporarily prevent them from continuing to execute. These can be caused by too many users being active at one point in time, one or more servers or network devices between the client and the end server being down, or even something as drastic as a DDOS attack.

The question is, what to do when these occur? One common approach is to implement Exponential Backoff in your application. Exponential Backoff is a way of spacing out repeated worker restarts for longer intervals until either:

1. A maximum number of retries is reached.
2. A maximum backoff time is reached.

It’s important to note that the time between the restarts is a random interval, up to a maximum defined value. Because value is random, it helps avoid the possibility that another background process may choose the same interval and begin executing at the same time as the current client, resulting in the same outage occurring.

The net effect of implementing an Exponential Backoff algorithm is that processes:

1. Can eventually complete successfully while not colliding with other processes.
2. Avoid creating excessive server load.

## Conclusion

That’s been a broad introduction to five challenges which PHP developers face when developing worker (background) processes. While not extensive, it’s provided an excellent introduction to them with the hope that these challenges can be avoided in the applications that you and your team are developing.

### About the author

Benjamin Founder & CEO

I am the founder and CEO of Tideways. I started the company over 10 years ago with the mission to move the PHP ecosystem forward, starting with performance. As managing director, I work across product, strategy, and the day-to-day of building a developer-focused SaaS business.

I’m a core contributor to the Doctrine open-source project and a founding board member of the PHP Foundation, which reflects my long-standing commitment to the PHP ecosystem. I particularly enjoy working at the intersection of application performance, developer experience, and the open-source community that makes PHP what it is. Outside of work, I enjoy board games, hiking, and coffee.
