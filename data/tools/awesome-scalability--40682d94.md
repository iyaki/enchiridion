---
title: "Awesome Scalability"
notion_id: 40682d94-fde3-46a8-af1b-e941ffd3f3ff
notion_url: https://app.notion.com/p/Awesome-Scalability-40682d94fde346a8af1be941ffd3f3ff
last_edited: 2026-09-21T17:11:00.000Z
source_url: https://github.com/binhnguyennus/awesome-scalability/#
tags: ["System Design / Software Architecture", "DevOps", "SysAdmin", "Website", "English"]
---
An updated and organized reading list for illustrating the patterns of scalable, reliable, and performant large-scale systems. Concepts are explained in the articles of prominent engineers and credible references. Case studies are taken from battle-tested systems that serve millions to billions of users.

### If your system goes slow

> Understand your problems: scalability problem (fast for a single user but slow under heavy load) or performance problem (slow for a single user) by reviewing some design principles and checking how scalability and performance problems are solved at tech companies. The section of intelligence are created for those who work with data and machine learning at big (data) and deep (learning) scale.

### If your system goes down

> "Even if you lose all one day, you can build all over again if you retain your calm!" - Thuan Pham, former CTO of Uber. So, keep calm and mind the availability and stability matters!

### If you are having a system design interview

> Look at some interview notes and real-world architectures with completed diagrams to get a comprehensive view before designing your system on whiteboard. You can check some talks of engineers from tech giants to know how they build, scale, and optimize their systems. Good luck!

### If you are building your dream team

> The goal of scaling team is not growing team size but increasing team output and value. You can find out how tech companies reach that goal in various aspects: hiring, management, organization, culture, and communication in the organization section.

### Community power

> Contributions are greatly welcome! You may want to take a look at the contribution guidelines. If you see a link here that is no longer maintained or is not a good fit, please submit a pull request!

> Many long hours of hard work have gone into this project. If you find it helpful, please share on Facebook, on Twitter, on Weibo, or on your chat groups! Knowledge is power, knowledge shared is power multiplied. Thank you!

## Content

- Principle
- Scalability
- Availability
- Stability
- Performance
- Intelligence
- Architecture
- Interview
- Organization
- Talk
- Book

## Principle

- Lessons from Giant-Scale Services - Eric Brewer, UC Berkeley & Google
- Designs, Lessons and Advice from Building Large Distributed Systems - Jeff Dean, Google
- How to Design a Good API & Why it Matters - Joshua Bloch, CMU & Google
- On Efficiency, Reliability, Scaling - James Hamilton, VP at AWS
- Principles of Chaos Engineering
- Finding the Order in Chaos
- The Twelve-Factor App
- Clean Architecture
- High Cohesion and Low Coupling
- Monoliths and Microservices
- CAP Theorem and Trade-offs
- CP Databases and AP Databases
- Stateless vs Stateful Scalability
- Scale Up vs Scale Out: Hidden Costs
- ACID and BASE
- Blocking/Non-Blocking and Sync/Async
- Performance and Scalability of Databases
- Database Isolation Levels and Effects on Performance and Scalability
- The Probability of Data Loss in Large Clusters
- Data Access for Highly-Scalable Solutions: Using SQL, NoSQL, and Polyglot Persistence
- SQL vs NoSQL
- SQL vs NoSQL - Lesson Learned at Salesforce
- NoSQL Databases: Survey and Decision Guidance
- How Sharding Works
- Consistent Hashing
- Consistent Hashing: Algorithmic Tradeoffs
- Don’t be tricked by the Hashing Trick
- Uniform Consistent Hashing at Netflix
- Eventually Consistent - Werner Vogels, CTO at Amazon
- Cache is King
- Anti-Caching
- Understand Latency
- Latency Numbers Every Programmer Should Know
- The Calculus of Service Availability
- Architecture Issues When Scaling Web Applications: Bottlenecks, Database, CPU, IO
- Common Bottlenecks
- Life Beyond Distributed Transactions
- Relying on Software to Redirect Traffic Reliably at Various Layers
- Breaking Things on Purpose
- Avoid Over Engineering
- Scalability Worst Practices
- Use Solid Technologies - Don’t Re-invent the Wheel - Keep It Simple!
- Simplicity by Distributing Complexity
- Why Over-Reusing is Bad
- Performance is a Feature
- Make Performance Part of Your Workflow
- The Benefits of Server Side Rendering over Client Side Rendering
- Automate and Abstract: Lessons at Facebook
- AWS Do's and Don'ts
- (UI) Design Doesn’t Scale - Stanley Wood, Design Director at Spotify
- Linux Performance
- Building Fast and Resilient Web Applications - Ilya Grigorik
- Accept Partial Failures, Minimize Service Loss
- Design for Resiliency
- Design for Self-healing
- Design for Scaling Out
- Design for Evolution
- Learn from Mistakes
- Microservices and Orchestration Domain-Oriented Microservice Architecture at Uber Service Architecture (3 parts: Domain Gateways, Value-Added Services, BFF) at SoundCloud Container (8 parts) at Riot Games Containerization at Pinterest Evolution of Container Usage at Netflix Dockerizing MySQL at Uber Testing of Microservices at Spotify Docker in Production at Treehouse Microservice at SoundCloud Operate Kubernetes Reliably at Stripe Cross-Cluster Traffic Mirroring with Istio at Trivago Agrarian-Scale Kubernetes (3 parts) at New York Times Nanoservices at BBC PowerfulSeal: Testing Tool for Kubernetes Clusters at Bloomberg Conductor: Microservices Orchestrator at Netflix Docker Containers that Power Over 100.000 Online Shops at Shopify Microservice Architecture at Medium From bare-metal to Kubernetes at Betabrand Kubernetes at Tinder Kubernetes at Quora Kubernetes Platform at Pinterest Microservices at Nubank Payment Transaction Management in Microservices at Mercari Service Mesh at Snap GRIT: Protocol for Distributed Transactions across Microservices at eBay Rubix: Kubernetes at Palantir CRISP: Critical Path Analysis for Microservice Architectures at Uber
- Distributed Caching EVCache: Distributed In-memory Caching at Netflix EVCache Cache Warmer Infrastructure at Netflix Memsniff: Robust Memcache Traffic Analyzer at Box Caching with Consistent Hashing and Cache Smearing at Etsy Analysis of Photo Caching at Facebook Cache Efficiency Exercise at Facebook tCache: Scalable Data-aware Java Caching at Trivago Pycache: In-process Caching at Quora Reduce Memcached Memory Usage by 50% at Trivago Caching Internal Service Calls at Yelp Estimating the Cache Efficiency using Big Data at Allegro Distributed Cache at Zalando Distributed Cache for S3 at ClickHouse Application Data Caching from RAM to SSD at NetFlix Tradeoffs of Replicated Cache at Skyscanner Location Caching with Quadtrees at Yext Video Metadata Caching at Vimeo Scaling Redis at Twitter Scaling Job Queue with Redis at Slack Moving persistent data out of Redis at Github Storing Hundreds of Millions of Simple Key-Value Pairs in Redis at Instagram Redis at Trivago Optimizing Redis Storage at Deliveroo Memory Optimization in Redis at Wattpad Redis Fleet at Heroku Solving Remote Build Cache Misses (2 parts) at SoundCloud Ratings & Reviews (2 parts) at Flipkart Prefetch Caching of Items at eBay Cross-Region Caching Library at Wix Improving Distributed Caching Performance and Efficiency at Pinterest Standardize and Improve Microservices Caching at DoorDash HTTP Caching and CDN Zynga Geo Proxy: Reducing Mobile Game Latency at Zynga Google AMP at Condé Nast A/B Tests on Hosting Infrastructure (CDNs) at Deliveroo HAProxy with Kubernetes for User-facing Traffic at SoundCloud Bandaid: Service Proxy at Dropbox Service Workers at Slack CDN Services at Spotify
- Distributed Locking Chubby: Lock Service for Loosely Coupled Distributed Systems at Google Distributed Locking at Uber Distributed Locks using Redis at GoSquared ZooKeeper at Twitter Eliminating Duplicate Queries using Distributed Locking at Chartio
- Distributed Tracking, Tracing, and Measuring Zipkin: Distributed Systems Tracing at Twitter Improve Zipkin Traces using Kubernetes Pod Metadata at SoundCloud Canopy: Scalable Distributed Tracing & Analysis at Facebook Pintrace: Distributed Tracing at Pinterest XCMetrics: All-in-One Tool for Tracking Xcode Build Metrics at Spotify Real-time Distributed Tracing at LinkedIn Tracking Service Infrastructure at Scale at Shopify Distributed Tracing at HelloFresh Analyzing Distributed Trace Data at Pinterest Distributed Tracing at Uber JVM Profiler: Tracing Distributed JVM Applications at Uber Data Checking at Dropbox Tracing Distributed Systems at Showmax osquery Across the Enterprise at Palantir StatsD at Etsy
- Distributed Scheduling Distributed Task Scheduling (3 parts) at PagerDuty Building Cron at Google Distributed Cron Architecture at Quora Chronos: A Replacement for Cron at Airbnb Scheduler at Nextdoor Peloton: Unified Resource Scheduler for Diverse Cluster Workloads at Uber Fenzo: OSS Scheduler for Apache Mesos Frameworks at Netflix Airflow - Workflow Orchestration Airflow at Airbnb Airflow at Adyen Airflow at Pandora Airflow at Robinhood Airflow at Lyft Airflow at Drivy Airflow at Grab Airflow at Adobe Auditing Airflow Job Runs at Walmart MaaT: DAG-based Distributed Task Scheduler at Alibaba boundary-layer: Declarative Airflow Workflows at Etsy
- Distributed Monitoring and Alerting Unicorn: Remediation System at eBay M3: Metrics and Monitoring Platform at Uber Athena: Automated Build Health Management System at Dropbox Vortex: Monitoring Server Applications at Dropbox Nuage: Cloud Management Service at LinkedIn Telltale: Application Monitoring at Netflix ThirdEye: Monitoring Platform at LinkedIn Periskop: Exception Monitoring Service at SoundCloud Securitybot: Distributed Alerting Bot at Dropbox Monitoring System at Alibaba Real User Monitoring at Dailymotion Alerting Ecosystem at Uber Alerting Framework at Airbnb Alerting on Service-Level Objectives (SLOs) at SoundCloud Job-based Forecasting Workflow for Observability Anomaly Detection at Uber Monitoring and Alert System using Graphite and Cabot at HackerEarth Observability (2 parts) at Twitter Distributed Security Alerting at Slack Real-Time News Alerting at Bloomberg Data Pipeline Monitoring System at LinkedIn Monitoring and Observability at Picnic
- Distributed Security Aardvark and Repokid: AWS Least Privilege for Distributed, High-Velocity Development at Netflix LISA: Distributed Firewall at LinkedIn Secure Infrastructure To Store Bitcoin In The Cloud at Coinbase BinaryAlert: Real-time Serverless Malware Detection at Airbnb Scalable IAM Architecture to Secure Access to 100 AWS Accounts at Segment OAuth Audit Toolbox at Indeed Active Directory Password Blacklisting at Yelp Syscall Auditing at Scale at Slack Athenz: Fine-Grained, Role-Based Access Control at Yahoo Security Development Lifecycle at Slack Unprivileged Container Builds at Kinvolk Diffy: Differencing Engine for Digital Forensics in the Cloud at Netflix Detecting Credential Compromise in AWS at Netflix Scalable User Privacy at Spotify AVA: Audit Web Applications at Indeed TTL as a Service: Automatic Revocation of Stale Privileges at Yelp Enterprise Key Management at Slack Scalability and Authentication at Twitch Edge Authentication and Token-Agnostic Identity Propagation at Netflix Hardening Kubernetes Infrastructure with Cilium at Palantir Improving Web Vulnerability Management through Automation at Lyft Clock Skew when Syncing Password Payloads at Drobbox
- Distributed Messaging, Queuing, and Event Streaming Cape: Event Stream Processing Framework at Dropbox Brooklin: Distributed Service for Near Real-Time Data Streaming at LinkedIn Samza: Stream Processing System for Latency Insighs at LinkedIn Bullet: Forward-Looking Query Engine for Streaming Data at Yahoo EventHorizon: Tool for Watching Events Streaming at Etsy Qmessage: Distributed, Asynchronous Task Queue at Quora Cherami: Message Queue System for Transporting Async Tasks at Uber Dynein: Distributed Delayed Job Queueing System at Airbnb Timestone: Queueing System for Non-Parallelizable Workloads at Netflix Messaging Service at Riot Games Messaging System Model at Dropbox Debugging Production with Event Logging at Zillow Cross-platform In-app Messaging Orchestration Service at Netflix Video Gatekeeper at Netflix Scaling Push Messaging for Millions of Devices at Netflix Delaying Asynchronous Message Processing with RabbitMQ at Indeed Benchmarking Streaming Computation Engines at Yahoo Improving Stream Data Quality With Protobuf Schema Validation at Deliveroo Scaling Email Infrastructure at Medium Real-time Messaging at Slack Event Stream Database at Nike Event Tracking System at Udemy Event-Driven Messaging Domain-Driven Design at Alibaba Domain-Driven Design at Weebly Domain-Driven Design at Moonpig Scaling Event Sourcing for Netflix Downloads Scaling Event-Sourcing at Jet.com Event Sourcing (2 parts) at eBay Event Sourcing at FREE NOW Scalable content feed using Event Sourcing and CQRS patterns at Brainly Pub-Sub Messaging Pulsar: Pub-Sub Messaging at Scale at Yahoo Wormhole: Pub-Sub System at Facebook MemQ: Cloud Native Pub-Sub System at Pinterest Pub-Sub in Microservices at Netflix Kafka - Message Broker Kafka at LinkedIn Kafka at Pinterest Kafka at Trello Kafka at Salesforce Kafka at The New York Times Kafka at Yelp Kafka at Criteo Kafka on Kubernetes at Shopify Kafka on PaaSTA: Running Kafka on Kubernetes at Yelp (2 parts) Migrating Kafka's Zookeeper with No Downtime at Yelp Reprocessing and Dead Letter Queues with Kafka at Uber Chaperone: Audit Kafka End-to-End at Uber Finding Kafka throughput limit in infrastructure at Dropbox Cost Orchestration at Walmart InfluxDB and Kafka to Scale to Over 1 Million Metrics a Second at Hulu Scaling Kafka to Support Data Growth at PayPal Stream Data Deduplication Exactly-once Semantics with Kafka Real-time Deduping at Tapjoy Deduplication at Segment Deduplication at Mail.Ru Petabyte Scale Data Deduplication at Mixpanel
- Distributed Logging Logging at LinkedIn Scalable and Reliable Log Ingestion at Pinterest High-performance Replicated Log Service at Twitter Logging Service with Spark at CERN Accelerator Logging and Aggregation at Quora Collection and Analysis of Daemon Logs at Badoo Log Parsing with Static Code Analysis at Palantir Centralized Application Logging at eBay Enrich VPC Flow Logs at Hyper Scale to provide Network Insight at Netflix BookKeeper: Distributed Log Storage at Yahoo LogDevice: Distributed Data Store for Logs at Facebook LogFeeder: Log Collection System at Yelp DBLog: Generic Change-Data-Capture Framework at Netflix
- Distributed Searching Search Architecture at Instagram Search Architecture at eBay Search Architecture at Box Search Discovery Indexing Platform at Coupang Universal Search System at Pinterest Improving Search Engine Efficiency by over 25% at eBay Indexing and Querying Telemetry Logs with Lucene at Palantir Query Understanding at TripAdvisor Search Federation Architecture at LinkedIn (2018) Search at Slack Search Engine at DoorDash Stability and Scalability for Search at Twitter Search Service at Twitter (2014) Autocomplete Search (2 parts) at Traveloka Data-Driven Autocorrection System at Canva Adapting Search to Indian Phonetics at Flipkart Nautilus: Search Engine at Dropbox Galene: Search Architecture of LinkedIn Manas: High Performing Customized Search System at Pinterest Sherlock: Near Real Time Search Indexing at Flipkart Nebula: Storage Platform to Build Search Backends at Airbnb ELK (Elasticsearch, Logstash, Kibana) Stack Predictions in Real Time with ELK at Uber Building a scalable ELK stack at Envato ELK at Robinhood Scaling Elasticsearch Clusters at Uber Elasticsearch Performance Tuning Practice at eBay Improve Performance using Elasticsearch Plugins (2 parts) at Tinder Elasticsearch at Kickstarter Log Parsing with Logstash and Google Protocol Buffers at Trivago Fast Order Search using Data Pipeline and Elasticsearch at Yelp Moving Core Business Search to Elasticsearch at Yelp Sharding out Elasticsearch at Vinted Self-Ranking Search with Elasticsearch at Wattpad Vulcanizer: a library for operating Elasticsearch at Github
- Distributed Storage In-memory Storage MemSQL Architecture - The Fast (MVCC, InMem, LockFree, CodeGen) And Familiar (SQL) Optimizing Memcached Efficiency at Quora Real-Time Data Warehouse with MemSQL on Cisco UCS Moving to MemSQL at Tapjoy MemSQL and Kinesis for Real-time Insights at Disney MemSQL to Query Hundreds of Billions of Rows in a Dashboard at Pandora Object Storage Scaling HDFS at Uber Reasons for Choosing S3 over HDFS at Databricks File System on Amazon S3 at Quantcast Image Recovery at Scale Using S3 Versioning at Trivago Cloud Object Store at Yahoo Ambry: Distributed Immutable Object Store at LinkedIn Dynamometer: Scale Testing HDFS on Minimal Hardware with Maximum Fidelity at LinkedIn Hammerspace: Persistent, Concurrent, Off-heap Storage at Airbnb MezzFS: Mounting Object Storage in Media Processing Platform at Netflix Magic Pocket: In-house Multi-exabyte Storage System at Dropbox
- Relational Databases MySQL at Uber MySQL at Pinterest PostgreSQL at Twitch Scaling MySQL-based Financial Reporting System at Airbnb Scaling MySQL at Wix Building and Deploying MySQL Raft at Meta MaxScale (MySQL) Database Proxy at Airbnb Switching from Postgres to MySQL at Uber Handling Growth with Postgres at Instagram Scaling the Analytics Database (Postgres) at TransferWise Updating a 50 Terabyte PostgreSQL Database at Adyen Scaling Database Access for 100s of Billions of Queries per Day at PayPal Minimizing Read-Write MySQL Downtime at Yelp Migrating MySQL from 5.6 to 8.0 at Facebook Migration from HBase to MyRocks at Quora Replication MySQL Parallel Replication (4 parts) at Booking.com Mitigating MySQL Replication Lag and Reducing Read Load at Github Read Consistency with Database Replicas at Shopify Black-Box Auditing: Verifying End-to-End Replication Integrity between MySQL and Redshift at Yelp Partitioning Main MySQL Database at Airbnb Herb: Multi-DC Replication Engine for Schemaless Datastore at Uber Sharding Sharding MySQL at Pinterest Sharding MySQL at Twilio Sharding MySQL at Square Sharding MySQL at Quora Sharding Layer of Schemaless Datastore at Uber Sharding & IDs at Instagram Sharding Postgres at Notion Sharding Postgres at Figma Solr: Improving Performance for Batch Indexing at Box Geosharded Recommendations (3 parts) at Tinder Scaling Services with Shard Manager at Facebook Presto the Distributed SQL Query Engine Presto at Pinterest Presto Infrastructure at Lyft Presto at Grab Engineering Data Analytics with Presto and Apache Parquet at Uber Data Wrangling at Slack Presto in Big Data Platform on AWS at Netflix Presto Auto Scaling at Eventbrite Speed Up Presto with Alluxio Local Cache at Uber
- NoSQL Databases Key-Value Databases DynamoDB at Nike DynamoDB at Segment DynamoDB at Mapbox Manhattan: Distributed Key-Value Database at Twitter Sherpa: Distributed NoSQL Key-Value Store at Yahoo HaloDB: Embedded Key-Value Storage Engine at Yahoo MPH: Fast and Compact Immutable Key-Value Stores at Indeed Venice: Distributed Key-Value Database at Linkedin Columnar Databases Cassandra Cassandra at Instagram Storing Images in Cassandra at Walmart Storing Messages with Cassandra at Discord Scaling Cassandra Cluster at Walmart Scaling Ad Analytics with Cassandra at Yelp Scaling to 100+ Million Reads/Writes using Spark and Cassandra at Dream11 Moving Food Feed from Redis to Cassandra at Zomato Benchmarking Cassandra Scalability on AWS at Netflix Service Decomposition at Scale with Cassandra at Intuit QuickBooks Cassandra for Keeping Counts In Sync at SoundCloud Cassandra Driver Configuration for Improved Performance and Load Balancing at Glassdoor cstar: Cassandra Orchestration Tool at Spotify HBase HBase at Salesforce HBase in Facebook Messages HBase in Imgur Notification Improving HBase Backup Efficiency at Pinterest HBase at Xiaomi Redshift Redshift at GIPHY Redshift at Hudl Redshift at Drivy Document Databases eBay: Building Mission-Critical Multi-Data Center Applications with MongoDB MongoDB at Baidu: Multi-Tenant Cluster Storing 200+ Billion Documents across 160 Shards Migrating Mongo Data at Addepar The AWS and MongoDB Infrastructure of Parse (acquired by Facebook) Migrating Mountains of Mongo Data at Addepar Couchbase Ecosystem at LinkedIn SimpleDB at Zendesk Espresso: Distributed Document Store at LinkedIn Graph Databases FlockDB: Distributed Graph Database at Twitter TAO: Distributed Data Store for the Social Graph at Facebook Akutan: Distributed Knowledge Graph Store at eBay
- Time Series Databases Beringei: High-performance Time Series Storage Engine at Facebook MetricsDB: TimeSeries Database for storing metrics at Twitter Atlas: In-memory Dimensional Time Series Database at Netflix Heroic: Time Series Database at Spotify Roshi: Distributed Storage System for Time-Series Event at SoundCloud Goku: Time Series Database at Pinterest Scaling Time Series Data Storage (2 parts) at Netflix Time Series Data Abstraction Layer at Netflix Druid - Real-time Analytics Database Druid at Airbnb Druid at Walmart Druid at eBay Druid at Netflix
- Distributed Repositories, Dependencies, and Configurations Management DGit: Distributed Git at Github Stemma: Distributed Git Server at Palantir Configuration Management for Distributed Systems at Flickr Git Repository at Microsoft Solve Git Problem with Large Repositories at Microsoft Single Repository at Google Scaling Infrastructure and (Git) Workflow at Adyen Dotfiles Distribution at Booking.com Secret Detector: Preventing Secrets in Source Code at Yelp Managing Software Dependency at Scale at LinkedIn Merging Code in High-velocity Repositories at LinkedIn Dynamic Configuration at Twitter Dynamic Configuration at Mixpanel Dynamic Configuration at GoDaddy Fleet Management (3 parts) at Spotify
- Scaling Continuous Integration and Continuous Delivery Continuous Integration Stack at Facebook Continuous Integration with Distributed Repositories and Dependencies at Netflix Continuous Integration and Deployment with Bazel at Dropbox Adopting Bazel for Web at Airbnb Continuous Deployments at BuzzFeed Screwdriver: Continuous Delivery Build System for Dynamic Infrastructure at Yahoo CI/CD at Betterment CI/CD at Brainly Scaling iOS CI with Anka at Shopify Scaling Jira Server at Yelp Auto-scaling CI/CD cluster at Flexport

## Availability

- Resilience Engineering: Learning to Embrace Failure Resilience Engineering with Project Waterbear at LinkedIn Resiliency against Traffic Oversaturation at iHeartRadio Resiliency in Distributed Systems at GO-JEK Practical NoSQL Resilience Design Pattern for the Enterprise at eBay Ensuring Resilience to Disaster at Quora Site Resiliency at Expedia Resiliency and Disaster Recovery with Kafka at eBay Disaster Recovery for Multi-Region Kafka at Uber
- Failover The Evolution of Global Traffic Routing and Failover Testing for Disaster Recovery Failover Testing Designing a Microservices Architecture for Failure ELB for Automatic Failover at GoSquared Eliminate the Database for Higher Availability at American Express Failover with Redis Sentinel at Vinted High-availability SaaS Infrastructure at FreeAgent MySQL High Availability at GitHub MySQL High Availability at Eventbrite Business Continuity & Disaster Recovery at Walmart
- Load Balancing Introduction to Modern Network Load Balancing and Proxying Top Five (Load Balancing) Scalability Patterns Load Balancing infrastructure to support more than 1.3 billion users at Facebook DHCPLB: DHCP Load Balancer at Facebook Katran: Scalable Network Load Balancer at Facebook Deterministic Aperture: A Distributed, Load Balancing Algorithm at Twitter Load Balancing with Eureka at Netflix Edge Load Balancing at Netflix Zuul 2: Cloud Gateway at Netflix Load Balancing at Yelp Load Balancing at Github Consistent Hashing to Improve Load Balancing at Vimeo UDP Load Balancing at 500 pixel QALM: QoS Load Management Framework at Uber Traffic Steering using Rum DNS at LinkedIn Traffic Infrastructure (Edge Network) at Dropbox Intelligent DNS based load balancing at Dropbox Monitor DNS systems at Stripe Multi-DNS Architecture (3 parts) at Monday Dynamic Anycast DNS Infrastructure at Hulu
- Rate Limiting Rate Limiting for Scaling to Millions of Domains at Cloudflare Cloud Bouncer: Distributed Rate Limiting at Yahoo Scaling API with Rate Limiters at Stripe Distributed Rate Limiting at Allegro Ratequeue: Core Queueing-And-Rate-Limiting System at Twilio Quotas Service at Grab Rate Limiting at Figma
- Autoscaling Autoscaling Pinterest Autoscaling Based on Request Queuing at Square Autoscaling Jenkins at Trivago Autoscaling Pub-Sub Consumers at Spotify Autoscaling Bigtable Clusters based on CPU Load at Spotify Autoscaling AWS Step Functions Activities at Yelp Scryer: Predictive Auto Scaling Engine at Netflix Bouncer: Simple AWS Auto Scaling Rollovers at Palantir Clusterman: Autoscaling Mesos Clusters at Yelp
- Availability in Globally Distributed Storage Systems at Google
- NodeJS High Availability at Yahoo
- Operations (11 parts) at LinkedIn
- Monitoring Powers High Availability for LinkedIn Feed
- Supporting Global Events at Facebook
- High Availability at BlaBlaCar
- High Availability at Netflix
- High Availability Cloud Infrastructure at Twilio
- High Availability with Distributed DB on Kubernetes at Airbnb
- Automating Datacenter Operations at Dropbox
- Globalizing Player Accounts at Riot Games

## Stability

- Circuit Breaker Circuit Breaking in Distributed Systems Circuit Breaker for Scaling Containers Lessons in Resilience at SoundCloud Protector: Circuit Breaker for Time Series Databases at Trivago Improved Production Stability with Circuit Breakers at Heroku Circuit Breaker at Zendesk Circuit Breaker at Traveloka Circuit Breaker at Shopify
- Timeouts Fault Tolerance (Timeouts and Retries, Thread Separation, Semaphores, Circuit Breakers) at Netflix Enforce Timeout: A Reliability Methodology at DoorDash Troubleshooting a Connection Timeout Issue with tcp_tw_recycle Enabled at eBay
- Crash-safe Replication for MySQL at Booking.com
- Bulkheads: Partition and Tolerate Failure in One Part
- Steady State: Always Put Logs on Separate Disk
- Throttling: Maintain a Steady Pace
- Multi-Clustering: Improving Resiliency and Stability of a Large-scale Monolithic API Service at LinkedIn
- Determinism (4 parts) in League of Legends Server

## Performance

- Performance Optimization on OS, Storage, Database, Network Improving Performance with Background Data Prefetching at Instagram Fixing Linux filesystem performance regressions at LinkedIn Compression Techniques to Solve Network I/O Bottlenecks at eBay Optimizing Web Servers for High Throughput and Low Latency at Dropbox Linux Performance Analysis in 60.000 Milliseconds at Netflix Live Downsizing Google Cloud Persistent Disks (PD-SSD) at Mixpanel Decreasing RAM Usage by 40% Using jemalloc with Python & Celery at Zapier Reducing Memory Footprint at Slack Continuous Load Testing at Slack Performance Improvements at Pinterest Server Side Rendering at Wix 30x Performance Improvements on MySQLStreamer at Yelp Optimizing APIs at Netflix Performance Monitoring with Riemann and Clojure at Walmart Performance Tracking Dashboard for Live Games at Zynga Optimizing CAL Report Hadoop MapReduce Jobs at eBay Performance Tuning on Quartz Scheduler at eBay Profiling C++ (Part 1: Optimization, Part 2: Measurement and Analysis) at Riot Games Profiling React Server-Side Rendering at HomeAway Hardware-Assisted Video Transcoding at Dailymotion Cross Shard Transactions at 10 Million RPS at Dropbox API Profiling at Pinterest Pagelets Parallelize Server-side Processing at Yelp Improving key expiration in Redis at Twitter Ad Delivery Network Performance Optimization with Flame Graphs at MindGeek Predictive CPU isolation of containers at Netflix Improving HDFS I/O Utilization for Efficiency at Uber Cloud Jewels: Estimating kWh in the Cloud at Etsy Unthrottled: Fixing CPU Limits in the Cloud (2 parts) at Indeed
- Performance Optimization by Tuning Garbage Collection Garbage Collection in Java Applications at LinkedIn Garbage Collection in High-Throughput, Low-Latency Machine Learning Services at Adobe Garbage Collection in Redux Applications at SoundCloud Garbage Collection in Go Application at Twitch Analyzing V8 Garbage Collection Logs at Alibaba Python Garbage Collection for Dropping 50% Memory Growth Per Request at Instagram Performance Impact of Removing Out of Band Garbage Collector (OOBGC) at Github Debugging Java Memory Leaks at Allegro Optimizing JVM at Alibaba Tuning JVM Memory for Large-scale Services at Uber Solr Performance Tuning at Walmart Memory Tuning a High Throughput Microservice at Flipkart
- Performance Optimization on Image, Video, Page Load Optimizing 360 Photos at Scale at Facebook Reducing Image File Size in the Photos Infrastructure at Etsy Improving GIF Performance at Pinterest Optimizing Video Playback Performance at Pinterest Optimizing Video Stream for Low Bandwidth with Dynamic Optimizer at Netflix Adaptive Video Streaming at YouTube Reducing Video Loading Time at Dailymotion Improving Homepage Performance at Zillow The Process of Optimizing for Client Performance at Expedia Web Performance at BBC
- Performance Optimization by Brotli Compression Boosting Site Speed Using Brotli Compression at LinkedIn Brotli at Booking.com Brotli at Treebo Deploying Brotli for Static Content at Dropbox Progressive Enhancement with Brotli at Yelp Speeding Up Redis with Compression at DoorDash
- Performance Optimization on Languages and Frameworks Python at Netflix Python at scale (3 parts) at Instagram OCaml best practices (2 parts) at Issuu PHP at Slack Go at Trivago TypeScript at Etsy Kotlin for taming state at Etsy Kotlin at DoorDash BPF and Go at Bumble Ruby on Rails at GitLab Rust in production at Figma Choosing a Language Stack at WeWork Switching from Go to Rust at Discord ASP.NET Core Performance Optimization at Agoda Data Race Patterns in Go at Uber Java 21 Virtual Threads at Netflix

## Intelligence

- Big Data Data Platform at Uber Data Platform at BMW Data Platform at Netflix Data Platform at Flipkart Data Platform at Coupang Data Platform at DoorDash Data Platform at Khan Academy Data Infrastructure at Airbnb Data Infrastructure at LinkedIn Data Infrastructure at GO-JEK Data Ingestion Infrastructure at Pinterest Data Analytics Architecture at Pinterest Data Orchestration Service at Spotify Big Data Processing (2 parts) at Spotify Big Data Processing at Uber Analytics Pipeline at Lyft Analytics Pipeline at Grammarly Analytics Pipeline at Teads ML Data Pipelines for Real-Time Fraud Prevention at PayPal Big Data Analytics and ML Techniques at LinkedIn Self-Serve Reporting Platform on Hadoop at LinkedIn Privacy-Preserving Analytics and Reporting at LinkedIn Analytics Platform for Tracking Item Availability at Walmart Real-Time Analytics for Mobile App Crashes using Apache Pinot at Uber HALO: Hardware Analytics and Lifecycle Optimization at Facebook RBEA: Real-time Analytics Platform at King AresDB: GPU-Powered Real-time Analytics Engine at Uber AthenaX: Streaming Analytics Platform at Uber Jupiter: Config Driven Adtech Batch Ingestion Platform at Uber Delta: Data Synchronization and Enrichment Platform at Netflix Keystone: Real-time Stream Processing Platform at Netflix Databook: Turning Big Data into Knowledge with Metadata at Uber Amundsen: Data Discovery & Metadata Engine at Lyft Maze: Funnel Visualization Platform at Uber Metacat: Making Big Data Discoverable and Meaningful at Netflix SpinalTap: Change Data Capture System at Airbnb Accelerator: Fast Data Processing Framework at eBay Omid: Transaction Processing Platform at Yahoo TensorFlowOnSpark: Distributed Deep Learning on Big Data Clusters at Yahoo CaffeOnSpark: Distributed Deep Learning on Big Data Clusters at Yahoo Spark on Scala: Analytics Reference Architecture at Adobe Experimentation Platform (2 parts) at Spotify Experimentation Platform at Airbnb Smart Product Platform at Zalando Log Analysis Platform at LINE Data Visualisation Platform at Myntra Building and Scaling Data Lineage at Netflix Building a scalable data management system for computer vision tasks at Pinterest Structured Data at Etsy Scaling a Mature Data Pipeline - Managing Overhead at Airbnb Spark Partitioning Strategies at Airbnb Scaling the Hadoop Distributed File System at LinkedIn Scaling Hadoop YARN cluster beyond 10,000 nodes at LinkedIn Scaling Big Data Access Controls at Pinterest
- Distributed Machine Learning Machine Learning Platform at Yelp Machine Learning Platform at Etsy Machine Learning Platform at Zalando Scaling AI/ML Infrastructure at Uber Recommendation System at Lyft Reinforcement Learning Platform at Lyft Platform for Serving Recommendations at Etsy Infrastructure to Run User Forecasts at Spotify Aroma: Using ML for Code Recommendation at Facebook Flyte: Cloud Native Machine Learning and Data Processing Platform at Lyft LyftLearn: ML Model Training Infrastructure built on Kubernetes at Lyft Horovod: Open Source Distributed Deep Learning Framework for TensorFlow at Uber Genie: Gen AI On-Call Copilot at Uber COTA: Improving Customer Care with NLP & Machine Learning at Uber Manifold: Model-Agnostic Visual Debugging Tool for Machine Learning at Uber Repo-Topix: Topic Extraction Framework at Github Concourse: Generating Personalized Content Notifications in Near-Real-Time at LinkedIn Altus Care: Applying a Chatbot to Platform Engineering at eBay PyKrylov: Accelerating Machine Learning Research at eBay Box Graph: Spontaneous Social Network at Box PricingNet: Pricing Modelling with Neural Networks at Skyscanner PinText: Multitask Text Embedding System at Pinterest SearchSage: Learning Search Query Representations at Pinterest Cannes: ML saves $1.7M a year on document previews at Dropbox Scaling Gradient Boosted Trees for Click-Through-Rate Prediction at Yelp Learning with Privacy at Scale at Apple Deep Learning for Image Classification Experiment at Mercari Deep Learning for Frame Detection in Product Images at Allegro Content-based Video Relevance Prediction at Hulu Moderating Inappropriate Video Content at Yelp Improving Photo Selection With Deep Learning at TripAdvisor Personalized Recommendations for Experiences Using Deep Learning at TripAdvisor Personalised Recommender Systems at BBC Machine Learning (2 parts) at Condé Nast Natural Language Processing and Content Analysis (2 parts) at Condé Nast Mapping the World of Music Using Machine Learning (2 parts) at iHeartRadio Machine Learning to Improve Streaming Quality at Netflix Machine Learning to Match Drivers & Riders at GO-JEK Improving Video Thumbnails with Deep Neural Nets at YouTube Quantile Regression for Delivering On Time at Instacart Cross-Lingual End-to-End Product Search with Deep Learning at Zalando Machine Learning at Jane Street Machine Learning for Ranking Answers End-to-End at Quora Ranking Platform at Booking.com Clustering Similar Stories Using LDA at Flipboard Similarity Search at Flickr Large-Scale Machine Learning Pipeline for Job Recommendations at Indeed Deep Learning from Prototype to Production at Taboola Atom Smashing using Machine Learning at CERN Mapping Tags at Medium Clustering with the Dirichlet Process Mixture Model in Scala at Monsanto Map Pins with DBSCAN & Random Forests at Foursquare Forecasting at Uber Financial Forecasting at Uber Productionizing ML with Workflows at Twitter GUI Testing Powered by Deep Learning at eBay Scaling Machine Learning to Recommend Driving Routes at Pivotal Real-Time Predictions at DoorDash Machine Learning for Indexing Text from Billions of Images at Dropbox Modeling User Journeys via Semantic Embeddings at Etsy Automated Fake Account Detection at LinkedIn Building Knowledge Graph at Airbnb Core Modeling at Instagram Neural Architecture Search (NAS) for Prohibited Item Detection at Mercari Computer Vision at Airbnb 3D Home Backend Algorithms at Zillow Long-term Forecasts at Lyft Discovering Popular Dishes with Deep Learning at Yelp SplitNet Architecture for Ad Candidate Ranking at Twitter Jobs Filter at Indeed Architecting Restaurant Wait Time Predictions at Yelp Music Personalization at Spotify Deep Learning for Domain Name Valuation at GoDaddy Similarity Clustering to Catch Fraud Rings at Stripe Personalized Search at Etsy ML Feature Serving Infrastructure at Lyft Context-Specific Bidding System at Etsy Moderating Promotional Spam and Inappropriate Content in Photos at Scale at Yelp Optimizing Payments with Machine Learning at Dropbox Scaling Media Machine Learning at Netflix Similarity Engine at eBay Machine Learning in Content Moderation at Etsy

## Architecture

- Tech Stack at Medium
- Tech Stack at Shopify
- Building Services (4 parts) at Airbnb
- Architecture of Evernote
- Architecture of Chat Service (3 parts) at Riot Games
- Architecture of League of Legends Client Update
- Architecture of Ad Platform at Twitter
- Architecture of API Gateway at Uber
- Architecture of API Gateway at Tinder
- Basic Architecture of Slack
- Lightweight Distributed Architecture to Handle Thousands of Library Releases at eBay
- Back-end at LinkedIn
- Back-end at Flickr
- Infrastructure (3 parts) at Zendesk
- Cloud Infrastructure at Grubhub
- Real-time Presence Platform at LinkedIn
- Settings Platform at LinkedIn
- Nearline System for Scale and Performance (2 parts) at Glassdoor
- Real-time User Action Counting System for Ads at Pinterest
- API Platform at Riot Games
- Games Platform at The New York Times
- Kabootar: Communication Platform at Swiggy
- Simone: Distributed Simulation Service at Netflix
- Seagull: Distributed System that Helps Running > 20 Million Tests Per Day at Yelp
- PriceAggregator: Intelligent System for Hotel Price Fetching (3 parts) at Agoda
- Phoenix: Testing Platform (3 parts) at Tinder
- Hexagonal Architecture at Netflix
- Architecture of Sticker Services at LINE
- Stack Overflow Enterprise at Palantir
- Architecture of Following Feed, Interest Feed, and Picked For You at Pinterest
- API Specification Workflow at WeWork
- Media Database at Netflix
- Member Transaction History Architecture at Walmart
- Sync Engine (2 parts) at Dropbox
- Ads Pacing Service at Twitter
- Rapid Event Notification System at Netflix
- Architectures of Finance, Banking, and Payment Systems Bank Backend at Monzo Trading Platform for Scale at Wealthsimple Core Banking System at Margo Bank Architecture of Nubank Tech Stack at TransferWise Tech Stack at Addepar Avoiding Double Payments in a Distributed Payments System at Airbnb Scaling Payments (3 parts) at Etsy Handles Millions of Digital Transactions Safely Everyday at Paytm Billing and Payment Platform at Grammarly

## Interview

- Designing Large-Scale Systems My Scaling Hero - Jeff Atwood (a dose of Endorphins before your interview, JK) Software Engineering Advice from Building Large-Scale Distributed Systems - Jeff Dean Introduction to Architecting Systems for Scale Anatomy of a System Design Interview 8 Things You Need to Know Before a System Design Interview Top 10 System Design Interview Questions Top 10 Common Large-Scale Software Architectural Patterns in a Nutshell Cloud Big Data Design Patterns - Lynn Langit How NOT to design Netflix in your 45-minute System Design Interview? API Best Practices: Webhooks, Deprecation, and Design
- Explaining Low-Level Systems (OS, Network/Protocol, Database, Storage) The Precise Meaning of I/O Wait Time in Linux Paxos Made Live – An Engineering Perspective How to do Distributed Locking SQL Transaction Isolation Levels Explained
- "What Happens When... and How" Questions Netflix: What Happens When You Press Play? Monzo: How Peer-To-Peer Payments Work Transit and Peering: How Your Requests Reach GitHub How Spotify Streams Music

## Organization

- Engineering Levels at SoundCloud
- Engineering Roles at Palantir
- Engineering Career Framework at Dropbox
- Scaling Engineering Teams at Twitter
- Scaling Decision-Making Across Teams at LinkedIn
- Scaling Data Science Team at GOJEK
- Scaling Agile at Zalando
- Scaling Agile at bol.com
- Lessons Learned from Scaling a Product Team at Intercom
- Hiring, Managing, and Scaling Engineering Teams at Typeform
- Scaling the Datagram Team at Instagram
- Scaling the Design Team at Flexport
- Team Model for Scaling a Design System at Salesforce
- Building Analytics Team (4 parts) at Wish
- From 2 Founders to 1000 Employees at Transferwise
- Lessons Learned Growing a UX Team from 10 to 170 at Adobe
- Five Lessons from Scaling at Pinterest
- Approach Engineering at Vinted
- Using Metrics to Improve the Development Process (and Coach People) at Indeed
- Mistakes to Avoid while Creating an Internal Product at Skyscanner
- RACI (Responsible, Accountable, Consulted, Informed) at Etsy
- Four Pillars of Leading People (Empathy, Inspiration, Trust, Honesty) at Zalando
- Pair Programming at Shopify
- Distributed Responsibility at Asana
- Rotating Engineers at Zalando
- Experiment Idea Review at Pinterest
- Tech Migrations at Spotify
- Improving Code Ownership at Yelp
- Agile Code Base at eBay
- Agile Data Engineering at Miro
- Automated Incident Management through Slack at Airbnb
- Refactor Organization at BBC
- Code Review Code Review at Palantir Code Review at LINE Code Reviews at Medium Code Review at LinkedIn Code Review at Disney Code Review at Netlify

## Talk

- Distributed Systems in One Lesson - Tim Berglund, Senior Director of Developer Experience at Confluent
- Building Real Time Infrastructure at Facebook - Jeff Barber and Shie Erlich, Software Engineer at Facebook
- Building Reliable Social Infrastructure for Google - Marc Alvidrez, Senior Manager at Google
- Building a Distributed Build System at Google Scale - Aysylu Greenberg, SDE at Google
- Site Reliability Engineering at Dropbox - Tammy Butow, Site Reliability Engineering Manager at Dropbox
- How Google Does Planet-Scale for Planet-Scale Infra - Melissa Binde, SRE Director for Google Cloud Platform
- Netflix Guide to Microservices - Josh Evans, Director of Operations Engineering at Netflix
- Achieving Rapid Response Times in Large Online Services - Jeff Dean, Google Senior Fellow
- Architecture to Handle 80K RPS Celebrity Sales at Shopify - Simon Eskildsen, Engineering Lead at Shopify
- Lessons of Scale at Facebook - Bobby Johnson, Director of Engineering at Facebook
- Performance Optimization for the Greater China Region at Salesforce - Jeff Cheng, Enterprise Architect at Salesforce
- How GIPHY Delivers a GIF to 300 Millions Users - Alex Hoang and Nima Khoshini, Services Engineers at GIPHY
- High Performance Packet Processing Platform at Alibaba - Haiyong Wang, Senior Director at Alibaba
- Solving Large-scale Data Center and Cloud Interconnection Problems - Ihab Tarazi, CTO at Equinix
- Scaling Dropbox - Kevin Modzelewski, Back-end Engineer at Dropbox
- Scaling with Performance at Facebook - Bill Jia, VP of Infrastructure at Facebook
- Scaling Live Videos to a Billion Users at Facebook - Sachin Kulkarni, Director of Engineering at Facebook
- Scaling Infrastructure at Instagram - Lisa Guo, Instagram Engineering
- Scaling Infrastructure at Twitter - Yao Yue, Staff Software Engineer at Twitter
- Scaling Infrastructure at Etsy - Bethany Macri, Engineering Manager at Etsy
- Scaling Real-time Infrastructure at Alibaba for Global Shopping Holiday - Xiaowei Jiang, Senior Director at Alibaba
- Scaling Data Infrastructure at Spotify - Matti (Lepistö) Pehrs, Spotify
- Scaling Pinterest - Marty Weiner, Pinterest’s founding engineer
- Scaling Slack - Bing Wei, Software Engineer (Infrastructure) at Slack
- Scaling Backend at Youtube - Sugu Sougoumarane, SDE at Youtube
- Scaling Backend at Uber - Matt Ranney, Chief Systems Architect at Uber
- Scaling Global CDN at Netflix - Dave Temkin, Director of Global Networks at Netflix
- Scaling Load Balancing Infra to Support 1.3 Billion Users at Facebook - Patrick Shuff, Production Engineer at Facebook
- Scaling (a NSFW site) to 200 Million Views A Day And Beyond - Eric Pickup, Lead Platform Developer at MindGeek
- Scaling Counting Infrastructure at Quora - Chun-Ho Hung and Nikhil Gar, SEs at Quora
- Scaling Git at Microsoft - Saeed Noursalehi, Principal Program Manager at Microsoft
- Scaling Multitenant Architecture Across Multiple Data Centres at Shopify - Weingarten, Engineering Lead at Shopify

## A Piece of Cake

Roses are red. Violets are blue. Binh likes sweet. Treat Binh a tiramisu? 🍰
