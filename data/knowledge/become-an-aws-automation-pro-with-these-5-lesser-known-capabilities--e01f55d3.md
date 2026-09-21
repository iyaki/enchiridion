---
title: "Become An AWS Automation Pro With These 5 Lesser-Known Capabilities"
notion_id: e01f55d3-420b-4b03-b549-41b19aeed9da
notion_url: https://app.notion.com/p/Become-An-AWS-Automation-Pro-With-These-5-Lesser-Known-Capabilities-e01f55d3420b4b03b54941b19aeed9da
last_edited: 2026-09-21T17:39:00.000Z
source_url: https://hackernoon.com/become-an-aws-automation-pro-with-these-5-lesser-known-capabilities-xn8135s0
tags: ["Article", "Hackernoon", "English", "AWS", "DevOps", "SysAdmin"]
---
[https://hackernoon.com/become-an-aws-automation-pro-with-these-5-lesser-known-capabilities-xn8135s0](https://hackernoon.com/become-an-aws-automation-pro-with-these-5-lesser-known-capabilities-xn8135s0)

Chances are you have used Amazon Web Services (AWS) at one point or another, whether to spin up a machine for a quick job, to run a full-fledged production application, or something in between.

Whatever your use of AWS—for development, testing, or production apps, automation can save time and money, and improve the resilience of your workloads.

Here are five lesser-known AWS capabilities you can leverage to become an AWS automation pro.

## 1. Automation on AWS Systems Manager

AWS Systems Manager can help manage and maintain AWS resources. It lets you build low-code automation workflows, monitor how tasks execute, and get notifications via CloudWatch Events. This is extremely useful for updating and patching machine images, as well as taking care of chores on EC2 or other managed services.

To set up automation on Systems Manager, you create a runbook called an “automation document”, written in JSON or YAML. There is a library of common tasks like restarting EC2 instances or updating an Amazon Machine Image (AMI). You can schedule up to 1,000 tasks per account—25 simultaneous tasks with up to 75 child tasks each.

See the Automations Action Reference for a full list of things you can do with Systems Manager automation.

## 2. Automating EBS Snapshots

Most AWS services include built-in automation features such as notifications, scheduling, and diagnostics. Automation is usually managed by labeling volumes or resources, and applying policies that define when, under what conditions, and by whom automated actions should be performed. Creating backups is a particularly useful example.

On AWS EC2, you can store incremental backups of compute instances as EBS snapshots, and then restore instances from there.

The challenge begins when a large number of EC2 instances and EBS volumes need to be backed up continuously. You can solve this using Amazon’s Lifecycle Manager for EBS Snapshots feature, in combination with smart tagging of your volumes. This lets you easily automate large-scale backups.

## 3. Using Multiple AWS Accounts for Development Teams

When automating infrastructure on AWS on a large scale, it can be very convenient to split up your deployment into multiple AWS accounts. Amazon provides a tool called AWS Organizations which lets you manage billing and security for multiple Amazon accounts in one place.

Why use multiple Amazon accounts? For example, you can give each developer or team their own AWS account, and use separate accounts for development, testing, and production. In a development account, you can give the developer full admin-level permissions of the entire environment, with complete isolation from production environments.

Developers can deploy entire stacks using CloudFormation without overwriting or damaging production environments. Isolation into separate accounts also lets you control costs—you can give each environment its own budget and quotas, so dev/test environments cannot get out of control in terms of costs.

Use AWS Secrets Manager to share credentials and other sensitive information between different AWS accounts, without having to embed secrets in your code.

## 4. Define a Duration for your Spot Instances

You can create Spot instances with a defined duration. These are spot instances that run continuously for the duration you select (but keep in mind the discount will be a bit lower). You can use them for anything that takes a known period of time to complete—for example, analytics, batch processing, or build jobs.

When you define a duration in a spot request, Amazon delivers the spot instance as soon as capacity is available, and it runs continuously until the duration ends. For example, the following code requests a spot instance that will run three instances, uninterrupted, for two hours:

```plain text
aws ec2 request-spot-instances --instance-count 3 --block-duration-minutes 120 --type "one-time" --launch-specification file://my-specs.json
```

## 5. Using SAM to Build and Test Serverless

AWS provides the Serverless Application Model (SAM), which lets you simplify serverless development using infrastructure as code (IaC). SAM is an open specification built on CloudFormation, Amazon’s template engine, and primary IaC platform.

Practically speaking, SAM gives you a set of command-line tools that let you build an entire serverless stack with just a few lines of a configuration. You can add Lambda function code into your configuration, and deploy everything together.

Another nice feature is SAM Local, which lets you test serverless applications on a local machine before deploying them. SAM Local runs a Docker-based environment on your local machine letting you test functions with simulations of event sources like S3 and DynamoDB. This lets you validate SAM IaC templates before actually running them and incurring costs on AWS.

## Conclusion

In this article, I covered five ways you can automate your way to AWS nirvana:

- Use AWS Systems Manager can help you create runbooks for automating cloud systems based on CloudWatch events
- Automate EBS snapshots to improve backups and reduce storage costs
- Split up your AWS deployment into several accounts to provide more isolation and manageability
- Define a duration for spot instances to make them more predictable
- Use the Serverless Application Model (SAM) to automate serverless tasks using infrastructure as code

Happy automation!
