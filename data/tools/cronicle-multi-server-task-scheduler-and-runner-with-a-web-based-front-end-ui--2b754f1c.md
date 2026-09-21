---
title: "Cronicle - multi-server task scheduler and runner, with a web based front-end UI"
notion_id: 2b754f1c-7d23-816b-9950-c74ece8b16ce
notion_url: https://app.notion.com/p/Cronicle-multi-server-task-scheduler-and-runner-with-a-web-based-front-end-UI-2b754f1c7d23816b9950c74ece8b16ce
last_edited: 2025-11-26T19:05:00.000Z
source_url: https://cronicle.net/
tags: ["English", "DevOps", "System Design / Software Architecture", "Node.js", "Performance", "Automation", "Tool", "cronicle.net"]
---
**Cronicle** is a multi-server task scheduler and runner, with a web based front-end UI.

It handles both scheduled, repeating and on-demand jobs, targeting any number of slave servers, with real-time stats and live log viewer.

Written in Node.js, proudly open source and MIT licensed.

Your jobs can emit progress events which are shown in real-time, along with estimated time remaining (calculated automatically). In addition to this, when your job is complete, you can emit categorized performance metrics which are displayed in a pie chart. This way, your users can see how the time was spent during your job run. [Read more](https://github.com/jhuckaby/Cronicle/blob/master/docs/Plugins.md#reporting-progress).

CPU and memory usage are tracked automatically for each of your jobs. This includes the main process that was spawned, as well as any child processes. So if you launch other processes or shell out to command-line utilities, all this will be taken into account. CPU is displayed as the percentage of one CPU core, and memory is displayed as a percentage of the configured maximum. [Read more](https://github.com/jhuckaby/Cronicle/blob/master/docs/WebUI.md#job-details-tab).

The average CPU and memory usage of your jobs is tracked over time, and you can pull up historical graphs, so you can detect trending patterns before they become a problem. In addition, your customized performance metrics are also tracked and graphed over time, if provided by your jobs (Plugins or shell scripts). [Read more](https://github.com/jhuckaby/Cronicle/blob/master/docs/WebUI.md#event-stats-tab).

Events are scheduled to run at various dates and times using a visual multi-selector widget. You can multi-select any combination of years, months, days, weekdays, hours and/or minutes. It will also repeat on a recurring basis, each time the server clock matches your selections. This is very similar to the Cron format. [Read more](https://github.com/jhuckaby/Cronicle/blob/master/docs/WebUI.md#event-timing).

You can set CPU and/or memory usage limits for each of your jobs, at the category or event level. These cause the job to be aborted if the limits are exceeded. You can also set "sustain" thresholds, so no action is taken until the limits are exceeded for a certain amount of time. [Read more](https://github.com/jhuckaby/Cronicle/blob/master/docs/WebUI.md#event-resource-limits).

Cronicle can send out an e-mails to a custom list of recipients for each job's completion. You can also specify a "web hook", which will send an HTTP POST to a custom URL that you specify, both at the start and the end of each job, and include full details in JSON format. [Read more](https://github.com/jhuckaby/Cronicle/blob/master/docs/WebUI.md#event-notification).

The Parameter system allows you to define a set of UI controls for your Plugin (text fields, text boxes, checkboxes, drop-down menus, etc.) which are then presented to the user when editing events. This can be useful if your Plugin has configurable behavior which you want to expose to your users. [Read more](https://github.com/jhuckaby/Cronicle/blob/master/docs/WebUI.md#plugin-parameters).

curl -s https://raw.githubusercontent.com/jhuckaby/Cronicle/master/bin/install.js | node
