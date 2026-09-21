---
title: "Chef"
notion_id: a71d2e8a-9a54-45e2-ba0b-cc4d16913739
notion_url: https://app.notion.com/p/Chef-a71d2e8a9a5445e2ba0bcc4d16913739
last_edited: 2026-09-21T17:14:00.000Z
source_url: https://www.chef.io/products/chef-infra
tags: ["Infrastructure", "Untried", "Tool", "English"]
---
Chef Infra

## Powerful Policy-Based Configuration Management System Software

Chef® Infra® configuration management software eliminates manual efforts and ensures infrastructure remains consistent and compliant over its lifetime — even in the most complex, heterogenous, and large-scale environments.

### Confidence Through Code

Define configurations and policies as code that are testable, enforceable and can be delivered at scale as part of automated pipelines.

### Drift Elimination

Ensure configurations only change if a system diverges from the desired defined state and automatically correct configuration drift, if needed.

### Manage Disparate Systems as One

Manage Windows and Linux systems running on prem, ARM systems running in the cloud or Macs laptops running at the edge all the same way.

### Extreme Flexibility

Use simple declarative definitions for common tasks or easily extend them to support the most unique environmental requirements.

## Chef Infra Policy-Based Configuration Management Automation Architecture

Today, there isn’t a company out there that doesn't worry about security, but traditional infrastructure as code (IaC) approaches no longer scale to meet the needs of modern security-minded organizations. Traditional approaches to IaC fail to:

- Account for regulatory or business security and compliance needs
- Account for collaboration between application, infrastructure and security/compliance teams
- Allow organizations to maintain velocity with security

Using Chef to automate configuration management allows DevOps teams to define policies that are repeatable, consistent and reusable. The result is increased business agility and security because all systems and resources are continuously and automatically evaluated, corrected, and modified.

Image of Infrastructure management with Chef Workstation, Chef Infra and Chef Automate. Test Policy Create Policy Chef® Infra® Chef® Workstation™ Report State Chef® Automate™ Cloud Servers OnPrem Servers Edge Devices Push Policy

Image of Infrastructure management with Chef Workstation Test Policy Create Policy Chef® Workstation™

### Chef Workstation

### Create and test policies

Reduce risks by iterating on policy changes before pushing them to production.

Image of Infrastructure management with Chef Infra Cloud Servers OnPrem Servers Edge Devices Chef® Infra®

### Chef Infra Client

### System State Enforcement

Enforce policy by converging the system to the state declared by the various resources.

Image of Infrastructure management with Chef Automate. Chef® Automate™

### Chef Automate

### Data Aggregation and Validation

View and validate intended and actual state across all systems.

### Chef Workstation: Everything You Need to Create and Test Policy-Based Configuration Management Automation

Chef Workstation gives users all the tools they need to get started with Chef all in one easy-to-install package. Users can execute ad-hoc remote configuration tasks, perform remote scanning, create cookbooks, test software and dependencies and much more. Chef Workstation provides a local developer experience and can be run on Linux, Mac laptops, or Windows systems.

Chef Workstation includes:

- Chef Tools: Chef Infra Client, Chef InSpec and Chef Habitat
- Chef Language: Pre-built resources for managing systems as well as helpers to make authoring and distributing cookbooks easy
- Testing and Upgrade Tools: Test Kitchen, Chef Cookstyle and Chef Upgrade Lab

Whether applying existing policies from the Chef Supermarket, the Chef community or writing customized policies, Chef Workstation ensures users have everything they need to get up and running quickly.

### Create Configuration Policies: Chef Cookbooks and Recipes

A Chef Infra Cookbook is the fundamental unit of configuration and policy automation distribution. Cookbooks are used to describe the system resources under management, such as files, templates, and software packages. Resources are defined in Chef Infra recipes that describe in human-readable code the desired state of each system. The code describes the desired state and then the Chef Infra Client automatically configures the system.

### Test Configuration Policies: Test Kitchen, Chef Cookstyle and Chef InSpec

Workstation includes a number of testing tools, including Chef InSpec, that help shorten deployment cycles and inform development decisions as you build out and refine your cookbooks.

- Test Kitchen is an open source testing framework that tests cookbooks using Vagrant, Docker, VMware vSphere, or leading cloud providers.
- Chef Cookstyle is a code analysis tool that helps users write better Chef Infra cookbooks by detecting issues and automatically correcting cookbook code.
- Chef InSpec is a compliance testing solution that defines policies as code and provides continuous visibility into compliance status across all systems and teams. Test Kitchen InSpec tests can be run against a converged cookbook for easy local infrastructure validation.

### Chef Infra Client: System State Enforcement

The Chef Infra Client is a powerful agent that applies your configurations on remote Linux, macOS, Windows and cloud-based systems. It does the hard work of configuring systems and allows you to scale Chef up or down as your needs change. Security is managed by key, which can be configured with cookbooks that allow you to rotate keys, enable or disable port access to SSH and WinRM, define authorized users and groups, and further secure your environments to keep them in line with Center for Internet Security (CIS) benchmarks and other industry standards.

Chef InSpec provides an additional level of system state enforcement. It provides a language that describes system state expectations in a way that can be directly mapped to policies defined for Chef Infra, providing insight into how to remediate any misconfigurations uncovered in audits within the same toolkit.

Key Capabilities of the Chef Infra Client include:

- Planned Updates: Updates are published to a pipeline and nodes pull and apply those updates during scheduled check-ins.
- Unstructured Updates: Chef Infra Policy can be updated via API requests, allowing external systems (chatbots or ticketing systems) to make web requests and alter policies based on user interactions.
- Push-Based Updates: The chef-run utility further provides ad-hoc, push-based updates to groups of servers in parallel. This is useful for one-time or on-demand updates and can be configured to send data to Chef Automate to ensure changes are captured in each node’s event history.
- Dynamic Behavior Support: Chef Infra can be configured to apply policies dynamically, by collecting memory, CPU, and other profile data from each managed system and using that information for your configurations, not hard-coded values. The same dynamic information collected by Chef can be displayed or filtered from the Chef Automate dashboard for easy visibility of configuration state and trends.
- Ephemeral Resource Management: Chef’s approach to policy-based configuration automation is especially well-suited to the configuration and maintenance of ephemeral servers managed as part of CI/CD pipelines. Cloud teams can couple on-demand cloud infrastructure with last-mile configuration using Chef to ensure new systems are set-up properly the first time, regardless of whether they number in the dozens or thousands.

### Chef Automate: Data Aggregation and Validation

Chef Automate provides enterprise management and observability capabilities, and it’s included with every Chef Subscription. Automate offers visual UIs, real-time interactive dashboards, role-based access controls, third-party integrations, data APIs, and much more. Automate brings together many pieces of the Chef ecosystem and enables Infrastructure, DevOps, Security, Cloud and Release teams to easily collaborate and get work done, all while maintaining an auditable history of changes to system environments.

Every time Chef Infra Client is run, it collects a variety of system profiling information, including memory, CPU, installed packages, cloud-provider metadata and other items, that is auto-fed into Chef Automate. There, the data is organized and exposed so you and your team can execute queries to help answer questions about your environment like “What versions of which operating systems are running?”, “How many CPUs does a group of servers have?” and much more. Data can also be used to populate CMDB tools like ServiceNow and readily exported to other systems.

### Validate Changes and Ensure Compliance Across all Environments with DevOps Dashboards

The Chef Infra Client Run Status chart displays a summary of node statuses: failed, successful, or missing, as well as the total node count. The node list table shows all nodes connected to Chef Automate. The Node Details table displays the most recent policy enforcement data reported by Chef Infra Client regarding the state of a system.

The Chef Automate Infrastructure State Management Dashboard provides users with the ability to view and manage Chef Infra Server details in Automate. Using these views users can:

- Add organizations to each server.
- Review cookbooks, roles, environments, data bags, and clients for each organization.
- Search and find roles, environments, data bag items, and clients from Chef Automate.

## Why Customers Choose Chef for Configuration Management Software

### Easy to Get Started with Chef Resources, Helpers and Community

With the Chef Language users define configurations once and then can apply them across mixed fleets of Linux, Mac and Windows systems, regardless of OS version and architecture. The Chef Infra language also includes a comprehensive set of pre-built resources, helpers and cookbooks created by both Chef and the Chef Community.

- Resources: Resources are used for configuring components such as packages, files, directories, or firewalls. Today, Chef Infra Client ships with more than 150 resources for common automation tasks such as user, file, kernel_module and windows_task.
- Helpers: Helpers enable users to make configuration decisions based on operating systems, clouds, virtualization hypervisors, and more.
- Chef Developer Community: The Chef community has created thousands of freely available configuration templates (cookbooks), for Chef Infra, that can be used as-is or as the base for an organization’s own configurations.

### Agile Test-Driven Development Practices

When it comes to test-driven infrastructure, Chef wrote the book on it. Chef helped pioneer the DevOps principles and test-driven infrastructure practices that today’s agile-based delivery teams use to test and deploy systems across their entire IT estates, from smart devices running on the edge to on-premises servers to sophisticated workloads running in the cloud.

### Immutable Deployments

An important part of DevOps deployments is ensuring everything runs as it should in development, test and production environments. Policyfiles ensure that the Chef Cookbooks running in production are the same versions that were tested in development. Because policies are immutable and cannot be changed once bundled, the Chef Infra Client no longer recalculates dependency sets at the start of each run, and cookbook authors no longer have to worry about their configurations changing out from under them.

## Recommended Content

Webinar Series

### Chef Infra Best Practice Quickfire Webinar Series

Watch Webinar

Blog

### Test Your Chef Knowledge with the Chef Principles Certification

Read Blog
