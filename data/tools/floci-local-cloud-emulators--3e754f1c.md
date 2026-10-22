---
title: "Floci — Local Cloud Emulators"
notion_id: 3e754f1c-7d23-8174-9160-deb595b4b471
notion_url: https://app.notion.com/p/Floci-Local-Cloud-Emulators-3e754f1c7d2381749160deb595b4b471
last_edited: 2026-09-26T03:28:00.000Z
source_url: https://floci.io/
tags: ["Official Website", "English", "Cloud", "DevOps", "Automation", "Serverless", "AWS", "Azure", "GCP", "OCI", "Tool", "Service"]
---
[Skip to main content](https://floci.io/#main-content)

## Any cloud.Locally.

Run AWS, Azure, GCP, and OCI locally in milliseconds: the fast, credential-free loop your team and its AI agents need to ship faster.

Cloud Emulators

## Pick your cloud. Start instantly.

Each emulator is a standalone MIT-licensed binary. No auth tokens, no feature gates, no cloud account needed.

floci

Drop-in replacement for LocalStack. Same port 4566, 119 services, native binary. Switch with zero code changes.

S3SQSLambdaDynamoDBRDSEKS+113 more

floci-az

Covers Blob, Queue, Table, Functions, App Config, Key Vault, Event Hubs and Service Bus. Native speed, MIT license, no Azure account needed.

BlobQueueTableFunctionsKey VaultEvent HubsService Bus+21 more

floci-gcp

Covers GCS, Pub/Sub, Firestore, Cloud Run, Cloud SQL, GKE and more. All 25 services on one port, native binary. MIT license, no GCP account needed.

GCSPub/SubFirestoreDatastoreSecret ManagerIAM+19 more

floci-oci

Covers Object Storage, Identity, Queue, Streaming, KMS, Vault and Functions. All 8 services on one port. MIT license, no Oracle Cloud account needed.

Object StorageQueueStreamingKMSVaultFunctions+2 more

Built for AI-assisted development

## Give your AI agents a cloud they can't break.

Coding agents write cloud code faster than ever, but they can't safely run it against a real account. Point them at Floci instead: a local cloud to build, run, and verify against. Instant, free, and credential-free.

0

### No credentials to leak.

Agents connect with throwaway keys. No real cloud secrets in the agent's context or sandbox. Nothing to exfiltrate, nothing to bill.

24ms

### Inner-loop fast.

24 ms cold start, 13 MiB idle. Agents spin up, test, and tear down inside the edit loop instead of waiting on round-trips to a remote account.

real

### Real signal, not mock theater.

Lambda, RDS, Redis, and Kafka run for real, so code an agent verifies locally behaves the same in production. No mock-shaped false positives.

✗

### Zero blast radius.

Let agents iterate aggressively. Worst case, they reset a local container, not your staging environment or your cloud bill.

export AWS_ENDPOINT_URL=http://localhost:4566

pytest tests/

Works with every SDK, CLI, Terraform/OpenTofu module, and test runner your agents already use, with no special integration.

Developer Tools

## One toolchain. All clouds.

Manage and explore every Floci emulator from a unified CLI or visual dashboard.

Why Floci

## Built for developers who ship.

No gatekeeping. No pricing tiers. No surprises. A local cloud that starts in milliseconds and works everywhere.

$0

### MIT Licensed. Forever free.

Fork it, embed it, extend it. No "community edition" sunset, no enterprise feature flags. Every service is available to every developer, always.

✗

### No auth token. Ever.

Pull the Docker image and go. No sign-ups, no API keys, no telemetry. LocalStack started requiring an auth token in March 2026. Floci never will.

24ms

### Native binary speed.

Compiled with GraalVM Mandrel. Starts in 24ms, idles at 13 MiB, fast enough for inner-loop iteration. Your CI, your laptop, and your AI agents will all thank you.

real

### Real engines, not mocks.

Lambda runs in real Docker containers. RDS uses real PostgreSQL/MySQL. ElastiCache runs real Redis. 100% protocol fidelity, no surprises in production.

Use Cases

## Where teams put Floci to work.

From the inner loop to CI to the classroom — anywhere a real cloud account is too slow, too risky, or too expensive.

CI

### Ephemeral test environments.

Spin up the full stack inside every pipeline job. 24 ms startup adds nothing to the build, each job gets an isolated cloud, and teardown is free.

λ

### Local-first development.

Build against S3, Pub/Sub, or Service Bus on your laptop — offline if you want. No shared dev account to trip over, no credentials to rotate.

tf

### Infrastructure-as-code dry runs.

Apply Terraform, OpenTofu, or CloudFormation against localhost first. Catch typos, drift, and bad refactors before they ever touch a real account.

101

### Workshops & onboarding.

Teach cloud services with zero billing risk. Every student runs the whole stack on their own machine — no accounts to provision, nothing to clean up.

Quick Start

## One install. Every service, locally.

No account, no token. Pick your platform and go.

curl -fsSL https://floci.io/install.sh | sh

floci start && eval $(floci env)

aws s3 mb s3://my-bucket

echo "Why pay for S3 when floci is free? 🎉" > hello-floci.txt

aws s3 cp hello-floci.txt s3://my-bucket/

aws s3 cp s3://my-bucket/hello-floci.txt hello-back.txt

cat hello-back.txt

119 AWS services ready on :4566. S3 · SQS · DynamoDB · Lambda · RDS · +114 more
