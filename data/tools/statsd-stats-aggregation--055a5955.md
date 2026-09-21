---
title: "statsd - stats aggregation"
notion_id: 055a5955-d3d6-400c-83bc-e3182dc86b5c
notion_url: https://app.notion.com/p/statsd-stats-aggregation-055a5955d3d6400c83bce3182dc86b5c
last_edited: 2022-12-21T15:40:00.000Z
source_url: https://github.com/statsd/statsd#statsd---
tags: ["English", "DevOps", "Site Reliability Engineering", "Untried", "Tool"]
---
A network daemon that runs on the Node.js platform and listens for statistics, like counters and timers, sent over UDP or TCP and sends aggregates to one or more pluggable backend services (e.g., Graphite).

## Key Concepts

- buckets Each stat is in its own "bucket". They are not predefined anywhere. Buckets can be named anything that will translate to Graphite (periods make folders, etc)
- values Each stat will have a value. How it is interpreted depends on modifiers. In general values should be integers.
- flush After the flush interval timeout (defined by config.flushInterval, default 10 seconds), stats are aggregated and sent to an upstream backend service.

## Installation and Configuration

### Docker

StatsD supports docker in three ways:

- The official container image on GitHub Container Registry
- The official container image on DockerHub
- Building the image from the bundled Dockerfile

### Manual installation

- Install Node.js (All Current and LTS Node.js versions are supported.)
- Clone the project
- Create a config file from exampleConfig.js and put it somewhere
- Start the Daemon: node stats.js /path/to/config

## Usage

The basic line protocol expects metrics to be sent in the format:

```plain text
<metricname>:<value>|<type>
```

So the simplest way to send in metrics from your command line if you have StatsD running with the default UDP server on localhost would be:

```plain text
echo "foo:1|c" | nc -u -w0 127.0.0.1 8125
```

## More Specific Topics

- Metric Types
- Graphite Integration
- Supported Servers
- Supported Backends
- Admin TCP Interface
- Server Interface
- Backend Interface
- Metric Namespacing
- StatsD Cluster Proxy

## Debugging

There are additional config variables available for debugging:

- debug - log exceptions and print out more diagnostic info
- dumpMessages - print debug info on incoming messages

For more information, check the exampleConfig.js.

## Tests

A test framework has been added using node-unit and some custom code to start and manipulate StatsD. Please add tests under test/ for any new features or bug fixes encountered. Testing a live server can be tricky, attempts were made to eliminate race conditions but it may be possible to encounter a stuck state. If doing dev work, a killall statsd will kill any stray test servers in the background (don't do this on a production machine!).

Tests can be executed with ./run_tests.sh.

## History

StatsD was originally written at Etsy and released with a blog post about how it works and why we created it.

## Inspiration

StatsD was inspired (heavily) by the project of the same name at Flickr. Here's a post where Cal Henderson described it in depth: Counting and timing. Cal re-released the code recently: Perl StatsD
