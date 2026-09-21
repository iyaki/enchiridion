---
title: "Using systemd units for Laravel cronjobs and background processes · Command-G"
notion_id: 31854f1c-7d23-8113-a5c8-df4c178de747
notion_url: https://app.notion.com/p/Using-systemd-units-for-Laravel-cronjobs-and-background-processes-Command-G-31854f1c7d238113a5c8df4c178de747
last_edited: 2026-03-03T01:58:00.000Z
source_url: https://command-g.nl/en/articles/using-systemd-units-for-cronjobs-and-background-processes
tags: ["English", "DevOps", "System Design / Software Architecture", "Laravel", "Productivity", "Article", "Tutorial", "Command-G"]
---
Most Laravel applications will be deployed to some sort of Linux server. When deploying an app there is almost always the need to run some sort of processes in the background, for example cronjobs or a queue worker. The documentation of Laravel describes how to use Supervisor to keep the queue worker running. With this article I would like to explain how to use systemd to run background jobs.

For this guide I have deployed an example repository on a clean Ubuntu 24.04 installation. The app is deployed by the laravel user in /home/laravel/app.

## What is systemd?

Systemd is the default init system and service manager for most Linux distributions. For example Ubuntu, Fedora and Debian. Systemd initializes the system and offers dependency management for services, which we will explore later in this article. Systemd also includes a component for collecting and managing logs.

A more complete description of what Systemd offers can be found on the website of DigitalOcean.

## Why use systemd over a regular cronjob and Supervisor?

There are several reasons to prefer systemd:

- There is no need to install additional tools to run background processes.
- It offers a centralized logging system
- It offers dependency management to make sure your queue worker starts after MySQL or Redis have started
- It is possible to limit resources (CPU and memory) per unit. This could come in handy when multiple apps are running on one server.

## Using a timer for cronjobs

To run the Laravel cron, two units have to be created: one service unit to run the artisan command and one timer unit to trigger the service.

The unit for the service must be created in /etc/systemd/system/laravel-cron.service and should contain this:

```plain text
[Unit] Description=Laravel Cron [Service] Type=oneshot User=laravel WorkingDirectory=/home/laravel/app ExecStart=/usr/bin/php /home/laravel/app/artisan schedule:run
```

Each unit file has the [Unit] header which describes the unit. For this service unit four other variables are set:

- Type=oneshot describes to systemd that the command will run once and will exit after it is completed
- User=laravel tells systemd that the service should run as the laravel user
- WorkingDirectory=/home/laravel/app sets the working directory when running the service
- ExecStart=/usr/bin/php /home/laravel/app/artisan schedule:run tells systemd which command to run when executing the service.

When using systemd to run a service based on a timer we should create another unit with the same name, but its extension should end with .timer instead of .service. We want to run the laravel-cron service, so we should create the file /etc/systemd/system/laravel-cron.timer:

```plain text
[Unit] Description=Run the laravel-cron service every minute [Timer] OnCalendar=*-*-* *:*:00 [Install] WantedBy=timers.target
```

With OnCalendar=*-*-* ::00 we tell systemd to run the laravel-cron service every minute when the server is running.

The [Install] section defines which target the system must have reached when this unit is started. For timer units the convention is that this is always timers.target.

When the two units are installed we must reload the daemons with sudo systemctl daemon-reload. We can now enable the timer with sudo systemctl enable laravel-cron.timer

To verify the timer is installed and active run systemctl list-timers. This shows the list of active timers and for each timer the last time it has run and when it will run again.

## Running background jobs via units

A background process, like a queue worker or Laravel Octane, requires a single unit file. We create a unit file in /etc/systemd/system/laravel-queue.service with the following contents:

```plain text
[Unit] Description=Laravel queue worker After=network.target [Service] Type=simple User=laravel WorkingDirectory=/home/laravel/app ExecStart=/usr/bin/php /home/laravel/app/artisan queue:work --tries=3 --sleep=3 --max-time=3600 Restart=always RestartSec=5 [Install] WantedBy=multi-user.target
```

As you can see, it is a lot like the service file we created for the cronjob. The things that differ are:

- After=network.target tells systemd to start this service after the network is available. If you're using the Redis queue driver, you could add another line below After=network.target with Requires=redis.service. This is the dependency management I wrote about earlier and instructs systemd to start the service after Redis has started. Another possibility is to use Wants=redis.service. The Requires parameter is more strict and will prevent the service from starting when Redis is not running. The Wants parameter will start the service after Redis is started, but when the Redis service is not installed it will still start the service. Note that this parameter only influences the starting of the service. This will not stop the service when Redis is started.
- Type=simple tells systemd it is a simple, long-running process
- Restart=always makes sure the service is restarted after it is stopped. This allows the user laravel to restart the queue worker with php artisan queue:restart
- RestartSec=5 is the delay before restarting the service

Reload the daemons with systemctl daemon-reloadand enable the service withsystemctl enable laravel-queue and systemctl start laravel-queue.

## Viewing status and outputs of the units

The status of a unit can be checked with systemctl status {unitname}. For a timer this command shows the last time it was triggered. When viewing the status for a service, it will show if the service is running and the last 10 lines of output.

Systemd provides unified logging for services and timers. These can be viewed with journalctl. This is very useful when debugging a specific job or task. To tail the logs for the queue worker service, you should run the command journalctl -u laravel-queue -f. This will open a tail to the logs and will show each line as it appears. It's also possible to view the output from a certain timestamp. You can achieve this with the -S parameter, for example: journalctl -u laravel-queue -S '2026-02-26 12:00'.

It's also possible to filter for specific messages in the output. This can be done with the -g parameter, for example: journalctl -u laravel-queue -g 'failed'. This will show only lines which contain the word 'failed'.
