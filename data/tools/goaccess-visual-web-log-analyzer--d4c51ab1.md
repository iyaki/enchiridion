---
title: "GoAccess - Visual Web Log Analyzer"
notion_id: d4c51ab1-ae52-41cc-a463-9eba0bf62b76
notion_url: https://app.notion.com/p/GoAccess-Visual-Web-Log-Analyzer-d4c51ab1ae5241cca4639eba0bf62b76
last_edited: 2023-09-20T19:34:00.000Z
source_url: https://goaccess.io/
tags: ["English", "Network", "Web Development", "SysAdmin", "Site Reliability Engineering", "DevOps", "Untried", "Tool"]
---
[https://github.com/allinurl/goaccess](https://github.com/allinurl/goaccess)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->



**GoAccess** is an open source **real-time**** web log analyzer** and interactive viewer that runs in a **terminal** in *nix systems or through your **browser**.

It provides **fast** and valuable HTTP statistics for system administrators that require a visual server report on the fly.

[Live Demo](http://rt.goaccess.io/?20230807125215=)

[Download](https://goaccess.io/download)

See the [**JSON**](https://goaccess.io/json) or [**CSV**](https://goaccess.io/goaccess_csv_report.csv?20230807125215=) outputs.

GoAccess was designed to be a fast, terminal-based log analyzer. Its core idea is to quickly analyze and view web server statistics in **real time** without needing to use your browser (_great if you want to do a quick analysis of your access log via SSH, or if you simply love working in the terminal_).

While the terminal output is the default output, it has the capability to generate a complete, self-contained **real-time** [`HTML`](http://rt.goaccess.io/?20230807125215=) report (_great for analytics, monitoring and data visualization_), as well as a [`JSON`](https://goaccess.io/json), and [`CSV`](https://goaccess.io/goaccess_csv_report.csv?20230807125215=) report.

### Key Features — [See Full List](https://goaccess.io/features)

- **Fast**, **real-time**, millisecond/second updates, written in C
- **Only** ncurses as a **dependency**
- **Nearly all** web log **formats** (Apache, Nginx, Amazon S3, Elastic Load Balancing, CloudFront, Caddy, etc)
- Simply set the log format and run it against your log
- Beautiful terminal and bootstrap dashboards (Tailor GoAccess to suit your own color taste/schemes)
- and of course, [Valgrind](http://valgrind.org/) tested.
