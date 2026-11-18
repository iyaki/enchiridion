---
title: "Your Guide To AWS Instance Metadata Service (IMDS)"
notion_id: ceda763d-c50e-437f-b0c2-eab37c7dde00
notion_url: https://app.notion.com/p/Your-Guide-To-AWS-Instance-Metadata-Service-IMDS-ceda763dc50e437fb0c2eab37c7dde00
last_edited: 2023-01-13T00:55:00.000Z
source_url: https://hackernoon.com/your-guide-to-aws-instance-metadata-service-imds-bt1t3717
tags: ["English", "AWS", "Information Security", "Article", "Hackernoon"]
---
## Too Long; Didn't Read

Instance Metadata Service (IMDS) provides “data about your instance that you can use to configure or manage the running instance’s. Instance metadata is divided into categories, for example, host name, events, and security groups. Every instance has access to its own MDS using any HTTP client request, such as, curl command from the instance to http://169.169.254/latest/meta-data/. IMDSv2 uses session-oriented requests to access the locally-stored EC2 instance metadata.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## [Editor note] Schedule for Wed, Aug 11th

Metadata is “data that provides information about other data” (Wikipedia). In other words, Metadata is “data about data”.

In AWS, Instance Metadata Service (IMDS) provides “data about your instance that you can use to configure or manage the running instance. Instance metadata is divided into categories, for example, host name, events, and security groups.” ([Amazon User Guide for Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html?ref=hackernoon.com))

Every instance has access to its own MDS using any HTTP client request, such as, curl command from the instance to [http://169.254.169.254/latest/meta-data](http://169.254.169.254/latest/meta-data?ref=hackernoon.com)

## Enumerating IMDS

All instance Metadata categories can be found [here](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instancedata-data-categories.html?ref=hackernoon.com).

A quick review over the categories presents two fields in which we can discover some AWS credentials: iam and identity-credentials.

**IAM **– One of the most sensitive categories is the [IAM](https://aws.amazon.com/iam/?ref=hackernoon.com) category, through which the IAM roles associated with an instance, including the token that was assigned to the instance by STS, can be discovered.

**Identity credentials **– according to the documentation, these credentials that AWS uses to identify an instance to the rest of the Amazon EC2 infrastructure and are for internal use only.

## User data

User data is the set of commands you can provide to an instance at launch time. Once the instance is launched, it executes the user data script with **root privileges**. Every instance has access to its own User data using curl command from the instance to [http://169.254.169.254/latest/user-data](http://169.254.169.254/latest/meta-data?ref=hackernoon.com)

Since user data can be viewed using IMDS, information security best practices must be incorporated into writing considerations. The most common bad practice is exposing credentials (user name and passwords) as part of the user data script. A better alternative would be to transfer the credentials over a secure channel.

## IMDSv2

In mid-November of 2019, AWS announced a new IMDS name called IMDSv2. The new release can overcome Server Side Request Forgery (SSRF) vulnerabilities in web applications running on EC2, open Website Application Firewalls, open reverse proxies, and open layer 3 firewalls and NATs.

IMDSv2 uses session-oriented requests. According to IMDSv2 introduction [blog post](https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service/?ref=hackernoon.com), every request is now protected by session authentication. A session begins and ends a series of requests that software running on an EC2 instance uses to access the locally-stored EC2 instance metadata and credentials.

The software starts a session with a simple HTTP PUT request to IMDSv2. IMDSv2 returns a secret token to the software running on the EC2 instance, which will use the token as a password to make requests to IMDSv2 for metadata and credentials. There’s no limit on the number of requests within a single session, and there’s no limit on the number of IMDSv2 sessions. Sessions can last up to six hours and, for added security, a session token can only be used directly from the EC2 instance where that session began:

```plain text
TOKEN=`curl -X PUT “http://169.254.169.254/latest/api/token” -H “X-aws-ec2-metadata-token-ttl-seconds: <seconds>”` \
```

```plain text
&& curl -H “X-aws-ec2-metadata-token: $TOKEN” -v
```

```plain text
http://169.254.169.254/latest/meta-data/
```

The token is being used as a password to make requests to IMDSv2 for metadata and credentials. The token is never stored by IMDSv2 and can never be retrieved by subsequent calls, so a session and its token are effectively destroyed when the process using the token terminates.

## Known Vulnerabilities (Taken from AWS Security Blog, link below)

Several known attack types can access the Metadata server from a remote location and enumerate the AWS account utilizing an EC2 instance of the account.

**Open Website Application Firewalls**

There are Web Application Firewall (WAF) services which can be misconfigured to allow attackers unauthorized access to the network behind the WAF, including the EC2 IMDS. To be transparent, WAFs usually pass on all the headers that come with a request, and do not add their own headers, such as the standard “X-Forwarded-For” header that other kinds of proxies add. In other words, applications behind a WAF get requests just as the requester sent them.

**Open Reverse Proxies**

Reverse proxies, such as Apache httpd or Squid, can be misconfigured to allow external requests to reach internal resources, but it’s still normal for these proxies to send an X-Forwarded-For HTTP header. That header itself is used to pass on the IP address of the original caller.

**Server-side request forgery (SSRF) vulnerabilities**

SSRF vulnerabilities allow attackers to make unauthorized requests from web applications. Since these requests come from the application itself, they can be used to access internal resources that the application has access to but that were not intended to be accessible to outsiders. SSRF vulnerabilities vary in their severity, and some are immune to other types of mitigations.

For instance, blocking SSRFs through static headers in instance metadata requests is effective only when the vulnerability merely allows the attacker to control the requested URL. However, AWS analysis found many SSRF vulnerabilities that allow attackers to set arbitrary headers. These SSRF vulnerabilities impact the application’s own header processing.

## Open layer 3 firewalls and NATs

An administrator may misconfigure the following to be accessible to outsiders: layer 3 [firewalls](https://www.checkpoint.com/cyber-hub/network-security/what-is-firewall/?ref=hackernoon.com), VPNs, tunnels, or NAT devices.

**Best Practices**

Below are a set of Best Practices recommended by Check Point CloudGuard Research team in order to avoid misconfigurations which can lead to security incidents.

**1. Limit access to IMDS**

If you don’t need IMDS – turn it off.

- Minimize active surface for attackers
- Limit access to software and processes listening on accessible ports so that they won’t have the ability to present EC2 Metadata service

Let’s take a simple python server which shows instance-id. First, it will take its data from the EC2 IMDS:

Once the server is running, browsing the EC2 will return the data:

This means that every time the server will run, it will access EC2 metadata service. In essence, exposing the service with no real need to do so.

The same result is easily achieved by emitting IMDS data to a local text file at bootstrap:

The python script is now slightly different, but the result remains the same:

Let’s now disable IMDS as part of instance launch:

```plain text
aws ec2 modify-instance-metadata-options –instance-id <instance-id> –http-endpoint disabled
```

While the first script needs IMDS available at all times, the secure script will work without it.

A good practice is to disable the IMDS as part of Instance’s User data. IMDS should be disabled by default. Only those authorized will open the service, by demand.

**2. Use IMDSv2**

By default, both IMDSv1 and IMDSv2 are available to the instance.

Using aws-cli, we can force a user to use only IMDSv2:

```plain text
aws ec2 modify-instance-metadata-options –instance-id <INSTANCE-ID> –profile <AWS_PROFILE> –http-endpoint enabled –http-token required
```

Now, IMDSv1 is down:

The user must therefore use IMDSv2

**3. Lock Down IMDS to specific users**

It is recommended to restrict IMDS availability by locking down the metadata endpoint so it is only accessible to specific OS users. On Linux machines, run the following:

```plain text
ip-lockdown 169.254.169.254 root
```

This command allows the endpoint to be accessed by only the root user. This way, an attacker can only use the metadata service if they obtain root privileges.

**4. IAM – EC2 Role**

If no application needs an EC2 role – don’t assign a role to this instance.

If an EC2 role is needed:

Analyze software to create least privilege role

```plain text
permissionsLimit
```

access as in (3)

**5. Always keep your applications updated and patched**

Check Point CloudGuard Cloud Intelligence and Threat Hunting (CITH)

Check Point CloudGuard CITH identifies misconfigurations and attacks and triggers the following alerts:

- Instance Metadata Service Modification
- Abuse of Access Token, generated by STS dedicated for EC2
- Abuse of Access Token, generated by Security Service Token (STS)Anomaly Detection – Anomalous network traffic

## References

- [https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service/](https://aws.amazon.com/blogs/security/defense-in-depth-open-firewalls-reverse-proxies-ssrf-vulnerabilities-ec2-instance-metadata-service/?ref=hackernoon.com)
- [https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-instance-metadata-service.html?ref=hackernoon.com)
- [https://www.portal.reinvent.awsevents.com/connect/search.ww?trk=direct,direct#loadSearch-searchPhrase=SEC310&searchType=session&tc=0&sortBy=abbreviationSort&p](https://www.portal.reinvent.awsevents.com/connect/search.ww?trk=direct%2Cdirect&ref=hackernoon.com#loadSearch-searchPhrase=SEC310&searchType=session&tc=0&sortBy=abbreviationSort&p=)

L O A D I N G
. . . comments &
