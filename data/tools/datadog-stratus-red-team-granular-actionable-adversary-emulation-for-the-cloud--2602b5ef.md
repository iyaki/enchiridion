---
title: "DataDog/stratus-red-team - Granular, Actionable Adversary Emulation for the Cloud"
notion_id: 2602b5ef-995d-4aab-b3bd-f2371643e61c
notion_url: https://app.notion.com/p/DataDog-stratus-red-team-Granular-Actionable-Adversary-Emulation-for-the-Cloud-2602b5ef995d4aabb3bdf2371643e61c
last_edited: 2023-09-13T13:38:00.000Z
source_url: https://github.com/DataDog/stratus-red-team
tags: ["English", "Information Security", "Untried", "Tool"]
---
# Stratus Red Team

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Stratus Red Team is "[Atomic Red Team](https://github.com/redcanaryco/atomic-red-team)™" for the cloud, allowing to emulate offensive attack techniques in a granular and self-contained manner.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Read the announcement blog posts:

- [https://www.datadoghq.com/blog/cyber-attack-simulation-with-stratus-red-team/](https://www.datadoghq.com/blog/cyber-attack-simulation-with-stratus-red-team/)
- [https://blog.christophetd.fr/introducing-stratus-red-team-an-adversary-emulation-tool-for-the-cloud/](https://blog.christophetd.fr/introducing-stratus-red-team-an-adversary-emulation-tool-for-the-cloud/)

## Getting Started

Stratus Red Team is a self-contained Go binary.

See the documentation at [**stratus-red-team.cloud**](https://stratus-red-team.cloud/):

- 

[Stratus Red Team Concepts](https://stratus-red-team.cloud/user-guide/getting-started/#concepts)

- 

[Installing Stratus Red Team](https://stratus-red-team.cloud/user-guide/getting-started/#installation) - Homebrew formula, Docker image and pre-built binaries available

- 

[Available Attack Techniques](https://stratus-red-team.cloud/attack-techniques/list/), mapped to MITRE ATT&CK

## Installation

- Mac OS:

```shell
brew tap datadog/stratus-red-team https://github.com/DataDog/stratus-red-team
brew install datadog/stratus-red-team/stratus-red-team

```

- 

Linux / Windows / Mac OS: Download one of the [pre-built binaries](https://github.com/datadog/stratus-red-team/releases).

- 

Docker:

```shell
IMAGE="ghcr.io/datadog/stratus-red-team"
alias stratus="docker run --rm -v $HOME/.stratus-red-team/:/root/.stratus-red-team/ -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY -e AWS_SESSION_TOKEN -e AWS_DEFAULT_REGION $IMAGE"
```

## Community

The following section lists posts and projects from the community leveraging Stratus Red Team.

Open-source projects:

- [Threatest](https://github.com/DataDog/threatest)
- [AWS Threat Detection with Stratus Red Team](https://github.com/sbasu7241/AWS-Threat-Simulation-and-Detection)

Blog posts:

- [AWS threat emulation and detection validation with Stratus Red Team and Datadog Cloud SIEM](https://www.datadoghq.com/blog/aws-threat-emulation-detection-validation-datadog/)
- [Adversary emulation on AWS with Stratus Red Team and Wazuh](https://wazuh.com/blog/adversary-emulation-on-aws-with-stratus-red-team-and-wazuh/)
- [Sky’s the Limit: Stratus Red Team for Azure](https://blog.detect.dev/posts/azure_for_stratus.html)
- [Detecting realistic AWS cloud-attacks using Azure Sentinel](https://medium.com/falconforce/falconfriday-detecting-realistic-aws-cloud-attacks-using-azure-sentinel-0xff1c-b62fd45c87dc)
- [A Data Driven Comparison of Open Source Adversary Emulation Tools](https://www.picussecurity.com/resource/blog/data-driven-comparison-between-open-source-adversary-emulation-tools)
- [Making Security Relevant in the Cloud](https://www.cloudreach.com/en/technical-blog/making-security-relevant-in-the-cloud/)
- [Detonating attacks with Datadog Stratus Red Team](https://chrisdunne.com/post/detonating-attacks-with-datadog-stratus-red-team)
- [AWS CloudTrail cheatsheet](https://invictus-ir.medium.com/aws-cloudtrail-cheat-sheet-dcf2b92e37e2)
- [Adversary emulation on GCP with Stratus Red Team and Wazuh](https://wazuh.com/blog/adversary-emulation-on-gcp-with-stratus-red-team-and-wazuh/)
- [Automated First-Response in AWS using Sigma and Athena](https://invictus-ir.medium.com/automated-first-response-in-aws-using-sigma-and-athena-615940bedc56)

Talks:

- [Purple Teaming & Adversary Emulation in the Cloud with Stratus Red Team, DEF CON Cloud Village 2022](https://www.youtube.com/watch?v=rXFFuYbkntU) (recorded after the event as the talks were not recorded)
- [Threat-Driven Development with Stratus Red Team](https://www.youtube.com/watch?v=AbWwcqLwcYI) by Ryan Marcotte Cobb
- [Cloudy With a Chance of Purple Rain: Leveraging Stratus Red Team - BSides Portland 2022](https://www.youtube.com/watch?v=Oq9ObzATZDI)

Videos:

- [Stratus Red Team: AWS EC2 Instance Credential Theft | Threat SnapShot](https://www.youtube.com/watch?v=TVS-M6DrSPw)

## Using Stratus Red Team as a Go Library

See [Examples](https://github.com/DataDog/stratus-red-team/blob/main/examples) and [Programmatic Usage](https://stratus-red-team.cloud/user-guide/programmatic-usage/).

## Development

### Building Locally

```shell
make
./bin/stratus --help
```

### Running Locally

```shell
go run cmd/stratus/*.go list
```

### Running the Tests

```shell
make test
```

### Building the Documentation

For local usage:

```shell
pip install mkdocs-material mkdocs-awesome-pages-plugin

make docs
mkdocs serve

```

### Acknowledgments

Maintainer: [@christophetd](https://twitter.com/christophetd)

Similar projects (see [how Stratus Red Team compares](https://stratus-red-team.cloud/comparison/)):

- [Atomic Red Team](https://github.com/redcanaryco/atomic-red-team) by Red Canary
- [Leonidas](https://github.com/FSecureLABS/leonidas) by F-Secure
- [pacu](https://github.com/RhinoSecurityLabs/pacu) by Rhino Security Labs
- [Amazon GuardDuty Tester](https://github.com/awslabs/amazon-guardduty-tester)
- [CloudGoat](https://github.com/RhinoSecurityLabs/cloudgoat) by Rhino Security Labs

Inspiration and relevant resources:

- [https://expel.io/blog/mind-map-for-aws-investigations/](https://expel.io/blog/mind-map-for-aws-investigations/)
- [https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/](https://rhinosecuritylabs.com/aws/aws-privilege-escalation-methods-mitigation/)
- [https://github.com/elastic/detection-rules/tree/main/rules/integrations/aws](https://github.com/elastic/detection-rules/tree/main/rules/integrations/aws)
