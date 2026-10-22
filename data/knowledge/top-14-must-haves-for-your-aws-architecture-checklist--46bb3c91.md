---
title: "Top 14 Must-Haves for Your AWS Architecture Checklist"
notion_id: 46bb3c91-2c9f-42ab-84d3-59701cea4279
notion_url: https://app.notion.com/p/Top-14-Must-Haves-for-Your-AWS-Architecture-Checklist-46bb3c912c9f42ab84d359701cea4279
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://hackernoon.com/top-14-must-haves-for-your-aws-architecture-checklist-0c2q35sv
tags: ["Article", "Hackernoon", "English", "AWS", "DevOps"]
---
[https://hackernoon.com/top-14-must-haves-for-your-aws-architecture-checklist-0c2q35sv](https://hackernoon.com/top-14-must-haves-for-your-aws-architecture-checklist-0c2q35sv)

After analyzing the requirements of your application, you came up with an AWS architecture. Good job! This blog post contains a checklist helping you to validate your architecture. Make sure that you haven’t forgotten any important aspects.

I’ve created an AWS diagram using Cloudcraft and will use the following architecture as an example when going through my AWS Architecture Checklist. The architecture shows a typical web application consisting of a layer for synchronous request processing and a layer for asynchronous job processing. My architecture is made up of the following components:

- Route 53 a globally distributed Domain Name System (DNS)
- CloudFront a Content Delivery Network (CDN)
- Simple Storage Service (S3) an object store
- Application Load Balancer (ALB) distributing incoming HTTPS requests among a fleet of EC2 Instances
- EC2 Instance a virtual machine
- Auto Scaling Group manages a fleet of EC2 instances
- Simple Queue Service (SQS) a message queue
- Relational Database Service (RDS) Aurora a PostgreSQL-compatible database
- Elastic File System (EFS) a network file system (NFS)

✅ Multi-AZ

Is every component distributed among at least two availability zones? AWS partitions their regions in so-called availability zones (AZ). An availability zone consists of one or multiple isolated data centers.

There are three types of AWS services:

1. Services that are operating among multiple AZs by default
2. Services that come with a built-in option for Multi-AZ deployments
3. Services that you need to deploy among multiple AZs yourself

There is no action needed for services that are operating among multiple AZs by default. However, make sure you have verified that information with the help of the AWS documentation.

Make sure you have specified Multi-AZ as a requirement for services that come with a built-in option for Multi-AZ deployments.

Double-check whether your application is running on multiple EC2 instances. Make sure to distribute the EC2 instances among at least two AZs in parallel. Be aware that it might not be possible to operate a legacy application on more than one machine at the same time. In that case, you need to know that AWS does not even offer an SLA for EC2 instances running in a single AZ only. You have to find a way to recover operation in another AZ in case of a disaster on your own.

Back to my example, architecture.

- Multi-AZ by default: Route 53, CloudFront, S3, SQS
- Multi-AZ optional: ALB, RDS Aurora, EFS, Auto Scaling
- Action required: EC2 Instances

In my example, I’m using Auto Scaling Groups to make sure EC2 instances are distributed among two Availability Zones.

✅ Multi-Region

Is it necessary to deploy your application to multiple regions? As discussed in the previous section, AWS partitions its regions into isolated availability zones. That’s one of the approaches to avoid failures to cause outages affecting the whole region.

Depending on your availability requirements, you might need to deploy your application among multiple regions. For example, use US East (Ohio) and US West (Oregon) to operate your application on the east and west coast in parallel.

However, availability comes with a cost. It is complicated and expensive to go to the Multi-Region path. Make sure that there is a real business case for the availability requirements. Until today, Multi-Region deployments are exceptions, not the norm.

In my example, I’m not making use of a 2nd region to improve availability.

✅ Stateless Server

Is your application persisting data on the virtual machines – either in memory or on disk? Check all EC2 instances in your architecture. Aim to implement the concept of a stateless server. Instead of persisting data on a single EC2 instance, outsource storing data to the database (for example, RDS Aurora), an object store (e.g., S3), or an NFS filesystem (EFS for Linux or FSx for Windows File Server).

Doing so enables you to add or remove EC2 instances on demand, roll out new versions of your application without any service interruption, replace failed EC2 instances automatically without losing any data.

Back to the example. No data is stored on EC2 instances. Instead, the following services are used to persist data:

- S3 the object store
- RDS Aurora a PostgreSQL database
- EFS a network file system (NFS) or FSx for Windows File Server

✅ Auto Recovery

Do all components recover from failure automatically? Many AWS services do recover from failure without human interaction:

- The Application Load Balancer uses a DNS name to be able to shift away from an AZ affected by an outage.
- RDS Aurora will failover to a replica instance automatically within minutes after the primary instance becomes unavailable.
- S3 is fault-tolerant by design and will route incoming requests to healthy nodes automatically.

It is important to note that EC2 does not recover from failure automatically. By default, a failed virtual machine is broken until you fix the issue manually. However, there are two options to recover failed instances automatically:

An Auto Scaling Group uses built-in EC2 health checks or the health check status from a linked load balancer to detect failed instances, terminates failed instances, and replaces them by launching a new instance.

EC2 Auto Recovery allows you to recover a single EC2 instance automatically. The built-in EC2 health checks are used to detect failures.

In my architecture, the Auto Scaling Groups (ASG) will recover failed instances automatically. In one case, the health check of the Application Load Balancer (ALB) is used to detect failed instances. In the other case, the ASG uses the built-in EC2 health check to detect failures.

✅ Decouple Client-Server-Communication

Are clients communicating with EC2 instances directly? Try to decouple the client from the server by introducing a load balancer or a queue. Doing is a requirement to scale on-demand and replace EC2 instances on the fly.

AWS provides three types of load balancers:

1. Classic Load Balancer (CLB) a legacy load balancer that you should not use for new architectures anymore.
2. Application Load Balancer (ALB) a layer-7 load balancer that you should use for HTTP/HTTPS-based communication.
3. Network Load Balancer (NLB) a load balancer that you should use for all communication that does not use HTTP/HTTPS.

A load balancer is an excellent option to decouple your client-server-communication. Especially if the client expects an immediate response to each request.

However, use asynchronous decoupling whenever possible. AWS offers a bunch of messaging services that you can use to do so:

- Simple Queue Service (SQS) a highly distributed message queue Kinesis an event stream for ordered messaging
- Amazon MQ provides managed Apache ActiveMQ message brokers
- Amazon Managed Streaming for Apache Kafka (Amazon MSK)

Back to my example. I’m using an Application Load Balancer to decouple the HTTPS communication between client and server. The browser does not need to know which server is available to answer requests. Instead, the browser sends requests to the ALB, which is responsible for forwarding the request to a healthy server. For asynchronous requests, I’m using SQS in my example. The web servers add jobs to a queue that workers are picking up and processing asynchronously.

✅ Auto Scaling

The promise of a public cloud is that we have to spend less on idle resources. As capacity is available on-demand, all you need to do is to scale your infrastructure automatically. Doing so is especially important when making use of EC2 instances.

In that case, make sure your architecture leverages Auto Scaling Groups to adjust the number of running EC2 instances. There are two options to do so:

- Based on a schedule.
- Based on a utilization metric.

In general, scaling based on a schedule is the more straightforward option. For example, you could set the number of EC2 instances to 4 during working hours and to 2 during the night and weekends. Scaling a fleet of EC2 instances is very simple when using SQS. In that case, use the number of messages in the queue to decide how many instances you need.

A side note: in some cases (e.g., development or test environments), it might be possible to scale down to 0, which means to shut down your application completely.

In my example, I’m using Auto Scaling Groups to adjust the number of EC2 instances automatically.

✅ Backup and Restore

Have you thought about a way to backup all persistent data? I recommend using AWS Backup to create snapshots of your data periodically. Currently, AWS Backup supports the following services:

- EC2 Instances
- EBS Volumes
- RDS
- DynamoDB
- EFS
- AWS Storage Gateway

Make sure the RPO (Recovery Point Objective) is defined, and the AWS Backup plans are configured accordingly.

In my example architecture, I’m using AWS Backup to create snapshots of the RDS Aurora database and the EFS file system.

If so, how long will it take to restore a backup? This question is critical to decide whether your architecture will be able to fulfill your RTO (Recovery Time Objective). Unfortunately, AWS typically does not offer any guarantees for the duration of a restore.

✅ Encrypt Data-In-Transit

Werner Vogels (CTO of Amazon) is quoted with: “Dance like nobody’s watching. Encrypt like everyone is.” You should take this advice to heart: encrypt all data in transit. No matter if the data is traversing the public Internet or just your private network (VPC).

Use the Amazon Certificate Manager (ACM) to manage TLS/SSL certificates and make sure you have specified HTTPS/TLS connections for incoming traffic to your load balancers only.

Further, you should make use of the TLS/SSL encrypted endpoints that database services like RDS are offering. The communication with AWS API’s (e.g., S3 or SQS) uses HTTPS and is encrypted by default.

Back to my example, all external and internal communication is either HTTPS or using an encrypted channel (RDS).

✅ Encrypt Data-At-Rest

Furthermore, you should verify that all components also encrypt data-at-rest. Most AWS services allow you to specify a key managed by the AWS Key Management Service (KMS) to encrypt and decrypt your data.

- Your architecture should specify the requirements for the KMS encryption keys.
- AWS managed CMKs (customer master key) are easy to use but do not allow fine granular access control.
- Customer-managed CMKs are more complicated to use but allow you to control access and use keys among multiple AWS accounts.
- Use key material generated by AWS or bring your key material.

Depending on your requirements, it might be necessary to use AWS CloudHSM to protect your master keys.

In my example architecture, I’m using KSM with AWS-managed CMKs to encrypt data-at-rest stored in RDS and EFS. S3 contains publicly available files only. Therefore I’m not encrypting that data at all.

✅ Minimize Network I/O

AWS charges for network traffic. First of all, you have to pay for traffic from AWS to the Internet. On top of that, you are also paying for traffic between availability zones.

Therefore, one goal for your architecture is to minimize network I/O. Do so by compressing data, reduce communication overhead, disable cross-region load balancing, for example.

✅ Maximum I/O Throughput

Verify the maximum I/O throughput for every network connection in your architecture. Handle this checklist item with care. I’ve seen too many architectures failing in production because of insufficient I/O throughput.

- The EC2 instance type defines the maximum network throughput between the ALB and your EC2 instance.
- The EC2 instance type defines the maximum network throughput between the EC2 instance and SQS/RDS/EFS.
- The RDS instance type defines the maximum network throughput between the RDS instance and the EC2 instances.

On top of that, there is also a hidden network connection between your EC2 instance and the EBS volume. The maximum I/O throughput to disk is also limited by the configuration of the EBS volume as well as the maximum network throughput from the EC2 instance to the EBS volume. Some instance types come with a dedicated network connection for EBS volumes as well. Head over to the AWS documentation to learn more: Amazon EBS Volume Types and Amazon EBS–Optimized Instances.

✅ Caching

Have you thought about adding caching layers to reduce the number of read requests? Introducing caching decreases the load on all underlying parts of your architecture. But caching comes with a downside, you have to find a way to invalidate the cache.

AWS offers a few services that allow you to introduce caching to your architecture:

- CloudFront, the content delivery network, caches responses to HTTP requests.
- ElastiCache provides in-memory databases (Redis or memcached) that you can integrate into your applications, for example, to cache responses from the database or to store pre-calculated results.
- DynamoDB DAX is a caching layer for Amazon’s NoSQL database.

In my example, I’m using CloudFront to cache the static contents (e.g., JavaScript and CSS files) of the web application.

✅ Protect Network Boundaries

From a network security perspective, it is crucial to reduce the attack surface from the Internet to a minimum. Go through your architecture and make sure you have specified network access only for a minimum of endpoints.

Your network architecture should consider that as well. It is typical to divide your VPC into public and private subnets. EC2 Instances and database services should be placed into a private subnet to avoid incoming traffic from the Internet at all.

In my example, the only public endpoints are CloudFront and the Application Load Balancer. The Application Load Balancer is deployed to the public subnets.

✅ Least Privilege Principle for IAM

Identity and Access Management (IAM) controls who can access or administer your cloud resources. Make sure your architecture follows the Least Privilege Principle when it comes to accessing AWS APIs.

That is easier said than done. Your architecture does not need to go into every detail here. But I suggest that you specify that EC2 instances, ECS tasks, and Lambda functions should make use of IAM roles to authenticate for AWS API calls.

Summary

My AWS Architecture Checklist summarizes best practices that I have learned the hard way when drafting and implementing architectures for AWS. The checklist covers the basics and helps you not to lose sight of a single critical point.

I’m closing with a template of my AWS Architecture Checklist for you to fill out.

◻️ Multi-AZ

◻️ Multi-Region

◻️ Stateless Server

◻️ Auto Recovery

◻️ Decouple Client-Server-Communication

◻️ Auto Scaling

◻️ Backup and Restore

◻️ Encrypt Data-In-Transit

◻️ Encrypt Data-At-Rest

◻️ Minimize Network I/O

◻️ Maximum I/O Throughput

◻️ Caching

◻️ Protect Network Boundaries

◻️ Least Privilege Principle for IAM

Previously published at https://blog.cloudcraft.co/aws-architecture-checklist/
