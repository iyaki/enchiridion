---
title: "GitLab runbooks"
notion_id: 2150ae40-954a-4e8b-9a41-cb1c5c3683b8
notion_url: https://app.notion.com/p/GitLab-runbooks-2150ae40954a4e8b9a41cb1c5c3683b8
last_edited: 2023-04-21T19:18:00.000Z
source_url: https://gitlab.com/gitlab-com/runbooks#gitlab-on-call-run-books
tags: ["Book", "Guide", "English", "DevOps", "Project Management", "Product Management", "Productivity", "Help Desk", "On Call"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

This project provides a guidance for Infrastructure Reliability Engineers and Managers who are starting an on-call shift or responding to an incident. If you haven't yet, review the [Incident Management](https://about.gitlab.com/handbook/engineering/infrastructure/incident-management/index.html) page in the handbook before reading on.

## On-Call

GitLab Reliability Engineers and Managers provide 24x7 on-call coverage to ensure incidents are responded to promptly and resolved as quickly as possible.

### Shifts

We use [PagerDuty](https://gitlab.pagerduty.com/) to manage our on-call schedule and incident alerting. We currently have two escalation policies for , one for [Production Incidents](https://gitlab.pagerduty.com/escalation_policies#P7IG7DS) and the other for [Production Database Assistance](https://gitlab.pagerduty.com/escalation_policies#P1SMG60). They are staffed by SREs and DBREs, respectively, and Reliability Engineering Managers.

Currently, rotations are weekly and the day's schedule is split 12/12 hours with engineers on call as close to daytime hours as their geographical region allows. We hope to hire so that shifts are an 8/8/8 hours split, but we're not staffed sufficiently yet across timezones.

### Joining the On-Call Rotation

When a new engineer joins the team and is ready to start shadowing for an on-call rotation, [overrides](https://support.pagerduty.com/docs/editing-schedules#section-create-and-delete-overrides) should be enabled for the relevant on-call hours during that rotation. Once they have completed shadowing and are comfortable/ready to be inserted into the primary rotations, update the membership list for the appropriate schedule to [add the new team member](https://support.pagerduty.com/docs/editing-schedules#section-adding-users).

This [pagerduty forum post](https://community.pagerduty.com/t/creating-a-shadow-schedule-to-onboard-new-employees/214) was referenced when setting up the [blank shadow schedule](https://community.pagerduty.com/t/creating-a-blank-schedule/212) and initial [overrides](https://support.pagerduty.com/docs/editing-schedules#section-create-and-delete-overrides) for on-boarding new team member

## Checklists

- [Engineer on Call (EOC)](https://gitlab.com/gitlab-com/runbooks/-/blob/master/on-call/checklists/eoc.md)
- [Incident Manager on Call (IMOC)](https://about.gitlab.com/handbook/engineering/infrastructure/incident-management/#incident-manager-on-call-imoc-responsibilities)
- [Communications Manager on Call (CMOC)](https://about.gitlab.com/handbook/engineering/infrastructure/incident-management/#communications-manager-on-call-cmoc-responsibilities)

To start with the right foot let's define a set of tasks that are nice things to do before you go any further in your week

By performing these tasks we will keep the [broken window effect](https://en.wikipedia.org/wiki/Broken_windows_theory) under control, preventing future pain and mess.

## Things to keep an eye on

### Issues

First check [the on-call issues](https://gitlab.com/gitlab-com/gl-infra/reliability/-/issues?scope=all&utf8=%E2%9C%93&state=all&label_name%5B%5D=oncall) to familiarize yourself with what has been happening lately. Also, keep an eye on the [#production](https://gitlab.slack.com/channels/production) and [#incident-management](https://gitlab.slack.com/channels/incident-management) channels for discussion around any on-going issues.

### Alerts

Start by checking how many alerts are in flight right now

- go to the [fleet overview dashboard](https://dashboards.gitlab.net/d/RZmbBr7mk/gitlab-triage) and check the number of Active Alerts, it should be 0. If it is not 0 
- go to the alerts dashboard and check what is being triggered 
- [azure](https://prometheus.gitlab.com/alerts)
- [gprd prometheus](https://prometheus.gprd.gitlab.net/alerts)
- [gprd prometheus-app](https://prometheus-app.gprd.gitlab.net/alerts)
- watch the [#production](https://gitlab.slack.com/channels/production) channel or [Pagerduty](https://gitlab.pagerduty.com/) for alert notifications; each alert here should point you to the right [runbook](https://gitlab.com/gitlab-com/runbooks) to fix it.
- if they don't, you have more work to do.
- be sure to create an issue, particularly to declare toil so we can work on it and suppress it.

### Prometheus targets down

Check how many targets are not scraped at the moment. alerts are in flight right now, to do this:

- go to the [fleet overview dashboard](https://dashboards.gitlab.net/d/RZmbBr7mk/gitlab-triage) and check the number of Targets down. It should be 0. If it is not 0 
- go to the [targets down list] and check what is. 
- [azure](https://prometheus.gitlab.com/consoles/up.html)
- [gprd prometheus](https://prometheus.gprd.gitlab.net/consoles/up.html)
- [gprd prometheus-app](https://prometheus-app.gprd.gitlab.net/consoles/up.html)
- try to figure out why there is scraping problems and try to fix it. Note that sometimes there can be temporary scraping problems because of exporter errors.
- be sure to create an issue, particularly to declare toil so we can work on it and suppress it.

## Incidents

First: don't panic.

If you are feeling overwhelmed, escalate to the [IMOC](https://about.gitlab.com/handbook/engineering/infrastructure/incident-management/#incident-manager-on-call-imoc-responsibilities). Whoever is in that role can help you get other people to help with whatever is needed. Our goal is to resolve the incident in a timely manner, but sometimes that means slowing down and making sure we get the right people involved. Accuracy is as important or more than speed.

Roles for an incident can be found in the [incident management section of the handbook](https://about.gitlab.com/handbook/engineering/infrastructure/incident-management/)

If you need to declare an incident, [follow these instructions located in the handbook](https://about.gitlab.com/handbook/engineering/infrastructure/incident-management/#reporting-an-incident).

## Communication Tools

If you do end up needing to post and update about an incident, we use [Status.io](https://status.io/)

On status.io, you can [Make an incident](https://app.status.io/dashboard/5b36dc6502d06804c08349f7/incident/create) and Tweet, post to Slack, IRC, Webhooks, and email via checkboxes on creating or updating the incident.

The incident will also have an affected infrastructure section where you can pick components of the GitLab.com application and the underlying services/containers should we have an incident due to a provider.

You can update incidents with the Update Status button on an existing incident, again you can tweet, etc from that update point.

Remember to close out the incident when the issue is resolved. Also, when possible, put the issue and/or google doc in the post mortem link.

## [Reporting and incident](https://about.gitlab.com/handbook/engineering/infrastructure/incident-management/#reporting-an-incident)

## Roles

During an incident, we have [roles defined in the handbook](https://about.gitlab.com/handbook/engineering/infrastructure/incident-management/#roles-and-responsibilities)

## General guidelines for production incidents

- Is this an emergency incident? 
- Are we losing data?
- Is GitLab.com not working or offline?
- Has the incident affected users for greater than 1 hour?
- Join the `#incident management` channel
- If the _point person_ needs someone to do something, give a direct command: [_@SOMEONE_](https://gitlab.com/SOMEONE)_: please run __`this`__ command_
- Be sure to be in sync - if you are going to reboot a service, say so: _I'm bouncing server X_
- If you have conflicting information, **stop and think**, bounce ideas, escalate
- Gather information when the incident is done - logs, samples of graphs, whatever could help figuring out what happened
- use `/security` if you have any security concerns and need to pull in the Security Incident Response team

### PostgreSQL

- [PostgreSQL](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/postgres.md)
- [more postgresql](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/postgresql.md)
- [PgBouncer](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/pgbouncer/pgbouncer-1.md)
- [PostgreSQL High Availability & Failovers](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/pg-ha.md)
- [PostgreSQL switchover](https://gitlab.com/gitlab-com/runbooks/-/blob/master/howto/postgresql-switchover.md)
- [Read-only Load Balancing](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/load-balancing.md)
- [Add a new secondary replica](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/postgresql-replica.md)
- [Database backups](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/postgresql-backups-wale-walg.md)
- [Database backups restore testing](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/postgresql-backups-wale-walg.md#database-backups-restore-testing)
- [Rebuild a corrupt index](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/postgresql.md#rebuild-a-corrupt-index)
- [Checking PostgreSQL health with postgres-checkup](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/postgres-checkup.md)
- [Reducing table and index bloat using pg_repack](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/pg_repack.md)

### Frontend Services

- [GitLab Pages returns 404](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/pages/gitlab-pages.md)
- [HAProxy is missing workers](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/config_management/chef-troubleshooting.md)
- [Worker's root filesystem is running out of space](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/monitoring/filesystem_alerts.md)
- [Azure Load Balancers Misbehave](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/frontend/load-balancer-outage.md)
- [GitLab registry is down](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/registry/gitlab-registry.md)
- [Sidekiq stats no longer showing](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/sidekiq/sidekiq_stats_no_longer_showing.md)
- [Gemnasium is down](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/gemnasium_is_down.md)
- [Blocking a project causing high load](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/block-high-load-project.md)

### Supporting Services

- [Redis](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/redis/redis.md)
- [Sentry is down](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/monitoring/sentry-is-down.md)

### Gitaly

- [Gitaly error rate is too high](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/gitaly/gitaly-error-rate.md)
- [Gitaly latency is too high](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/gitaly/gitaly-latency.md)
- [Sidekiq Queues are out of control](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/sidekiq/large-sidekiq-queue.md)
- [Workers have huge load because of cat-files](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/workers-high-load.md)
- [Test pushing through all the git nodes](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/git/git.md)
- [How to gracefully restart gitaly-ruby](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/gitaly/gracefully-restart-gitaly-ruby.md)
- [Debugging gitaly with gitaly-debug](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/gitaly/gitaly-debugging-tool.md)
- [Gitaly token rotation](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/gitaly/gitaly-token-rotation.md)
- [Praefect is down](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/praefect/praefect-startup.md)
- [Praefect error rate is too high](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/praefect/praefect-error-rate.md)

### CI

- [Large number of CI pending builds](https://gitlab.com/gitlab-com/runbooks/-/blob/master/troubleshooting/ci_pending_builds.md)
- [The CI runner manager report a high number of errors](https://gitlab.com/gitlab-com/runbooks/-/blob/master/troubleshooting/ci_runner_manager_errors.md)

### Geo

- [Geo database replication](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/geo-patroni-cluster.md)

### ELK

- [`mapper_parsing_exception`](https://gitlab.com/gitlab-com/runbooks/-/blob/master/troubleshooting/elk_mapper_parsing_exception.md)[ errors](https://gitlab.com/gitlab-com/runbooks/-/blob/master/troubleshooting/elk_mapper_parsing_exception.md)

## Non-Critical

- [SSL certificate expires](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/frontend/ssl_cert.md)
- [Troubleshoot git stuck processes](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/git/git-stuck-processes.md)

## Non-Core Applications

- [version.gitlab.com](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/version/version-gitlab-com.md)

### Chef/Knife

- [General Troubleshooting](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/config_management/chef-troubleshooting.md)
- [Error executing action ](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/stale-file-handles.md)[`create`](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/stale-file-handles.md)[ on resource 'directory[/some/path]'](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/stale-file-handles.md)

### Certificates

- [Certificate runbooks](https://gitlab.com/gitlab-com/runbooks/-/blob/master/certificates/README.md)

## Learning

### Alerting and monitoring

- [GitLab monitoring overview](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/monitoring/monitoring-overview.md)
- [How to add alerts: Alerts manual](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/monitoring/alerts_manual.md)
- [How to add/update deadman switches](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/deadman-switches.md)
- [How to silence alerts](https://gitlab.com/gitlab-com/runbooks/-/blob/master/howto/silence-alerts.md)
- [Alert for SSL certificate expiration](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/alert-for-ssl-certificate-expiration.md)
- [Working with Grafana](https://gitlab.com/gitlab-com/runbooks/-/blob/master/monitoring/grafana.md)
- [Working with Prometheus](https://gitlab.com/gitlab-com/runbooks/-/blob/master/monitoring/prometheus.md)
- [Upgrade Prometheus and exporters](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/monitoring/upgrades.md)
- [Use mtail to capture metrics from logs](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/mtail.md)

### CI

- [Introduction to Shared Runners](https://gitlab.com/gitlab-com/runbooks/-/blob/master/troubleshooting/ci_introduction.md)
- [Understand CI graphs](https://gitlab.com/gitlab-com/runbooks/-/blob/master/troubleshooting/ci_graphs.md)

### Access Requests

- [Deal with various kinds of access requests](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/access-requests.md)

### Deploy

- [Get the diff between dev versions](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/dev-environment.md#figure-out-the-diff-of-deployed-versions)
- [Deploy GitLab.com](https://ops.gitlab.net/gitlab-cookbooks/chef-repo/blob/master/doc/deploying.md)
- [Rollback GitLab.com](https://gitlab.com/gitlab-org/release/docs/-/blob/master/runbooks/rollback-a-deployment.md)
- [Deploy staging.GitLab.com](https://ops.gitlab.net/gitlab-cookbooks/chef-repo/blob/master/doc/staging.md)
- [Refresh data on staging.gitlab.com](https://ops.gitlab.net/gitlab-cookbooks/chef-repo/blob/master/doc/staging.md)
- [Background Migrations](https://gitlab.com/gitlab-org/release/docs/-/blob/master/runbooks/background-migrations.md)

### Work with the fleet and the rails app

- [Reload Puma with zero downtime](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-workers.md#reload-puma-with-zero-downtime)
- [How to perform zero downtime frontend host reboot](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-workers.md#how-to-perform-zero-downtime-frontend-host-reboot)
- [Gracefully restart sidekiq jobs](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-workers.md#gracefully-restart-sidekiq-jobs)
- [Start a read-only rails console](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/Teleport/Connect_to_Rails_Console_via_Teleport.md)
- [Start a rails console in the staging environment](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/staging-environment.md#run-a-rails-console-in-staging-environment)
- [Start a redis console in the staging environment](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/staging-environment.md#run-a-redis-console-in-staging-environment)
- [Start a psql console in the staging environment](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/staging-environment.md#run-a-psql-console-in-staging-environment)
- [Force a failover with postgres](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/patroni-management.md#failoverswitchover)
- [Force a failover with redis](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-pacemaker.md#force-a-failover)
- [Use aptly](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/aptly.md)
- [Disable PackageCloud](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/stop-or-start-packagecloud.md)
- [Re-index a package in PackageCloud](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/reindex-package-in-packagecloud.md)
- [Access hosts in GCP](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/access-gcp-hosts.md)

### Restore Backups

- [Deleted Project Restoration](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/deleted-project-restore.md)
- [PostgreSQL Backups: WAL-E, WAL-G](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/patroni/postgresql-backups-wale-walg.md)
- [Work with Azure Snapshots](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/azure-snapshots.md)
- [Work with GCP Snapshots](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/gcp-snapshots.md)
- [PackageCloud Infrastructure And Recovery](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/packagecloud-infrastructure.md)

### Work with storage

- [Understanding GitLab Storage Shards](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/gitaly/storage-sharding.md)
- [How to re-balance GitLab Storage Shards](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/gitaly/storage-rebalancing.md)
- [Build and Deploy New Storage Servers](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/gitaly/storage-servers.md)
- [Manage uploads](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/uploads.md)

### Mangle front end load balancers

- [Isolate a worker by disabling the service in the LBs](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/frontend/block-things-in-haproxy.md#disable-a-whole-service-in-a-load-balancer)
- [Deny a path in the load balancers](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/frontend/block-things-in-haproxy.md#deny-a-path-with-the-delete-http-method)
- [Purchasing/Renewing SSL Certificates](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/frontend/ssl_cert-1.md)

### Work with Chef

- [Create users, rotate or remove keys from chef](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-chef.md)
- [Update packages manually for a given role](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-workers.md#update-packages-fleet-wide)
- [Rename a node already in Chef](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/rename-nodes.md)
- [Reprovisioning nodes](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/reprovisioning-nodes.md)
- [Speed up chefspec tests](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/chefspec.md#tests-are-taking-too-long-to-run)
- [Manage Chef Cookbooks](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/chef-documentation.md)
- [Chef Guidelines](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/chef-guidelines.md)
- [Chef Vault](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/chef-vault.md)
- [Debug failed provisioning](https://gitlab.com/gitlab-com/runbooks/-/blob/master/howto/debug-failed-chef-provisioning.md)

### Work with CI Infrastructure

- [Runners fleet configuration management](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/ci-runners/fleet-configuration-management/README.md)
- [Investigate Abuse Reports](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/ci-runners/ci-investigate-abuse.md)
- [Create runners manager for GitLab.com](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/ci-runners/create-runners-manager-node.md)
- [Update docker-machine](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/upgrade-docker-machine.md)
- [CI project namespace check](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/ci-runners/ci-project-namespace-check.md)

### Work with Infrastructure Providers (VMs)

- [Getting Support from GCP](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/externalvendors/GCP-rackspace-support.md)
- [Create a DO VM for a Service Engineer](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/create-do-vm-for-service-engineer.md)
- [Create VMs in Azure, add disks, etc](https://ops.gitlab.net/gitlab-cookbooks/chef-repo/blob/master/doc/azure.md#managing-vms-in-azure)
- [Bootstrap a new VM](https://ops.gitlab.net/gitlab-cookbooks/chef-repo/blob/master/doc/new-vps.md)
- [Remove existing node checklist](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/remove-node.md)

### Manually ban an IP or netblock

- [Ban a single IP using Redis and Rack Attack](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/redis/ban-an-IP-with-redis.md)
- [Ban a netblock on HAProxy](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/frontend/ban-netblocks-on-haproxy.md)

### Dealing with Spam

- [General procedures for fighting spam in snippets, issues, projects, and comments](https://docs.google.com/document/d/1V0X2aYiNTZE1npzeqDvq-dhNFZsEPsL__FqKOMXOjE8)

### Manage Marvin, our infra bot

- [Manage cog](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-cog.md)

### ElasticStack (previously Elasticsearch)

Selected elastic documents and resources:

- docs/ 
- elastic/ 
- [elastic-cloud.md](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/elastic/elastic-cloud.md) (hosted ES provider docs)
- [exercises](https://gitlab.com/gitlab-com/runbooks/-/tree/master/docs/elastic/exercises) (e.g. cluster performance tuning)
- [kibana.md](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/elastic/kibana.md)
- [README.md](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/elastic/README.md) (ES overview)
- troubleshooting/ 
- [README.md](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/elastic/troubleshooting/README.md) (troubleshooting overview)
- [scripts/](https://gitlab.com/gitlab-com/runbooks/-/blob/master/elastic/scripts) (api calls used for admin tasks documented as bash scripts)
- watchers/

### ElasticStack integration in Gitlab (indexing Gitlab data)

[elasticsearch-integration-in-gitlab.md](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/elastic/elasticsearch-integration-in-gitlab.md)

### Zoekt integration in Gitlab (indexing code, BETA)

[zoekt-integration-in-gitlab.md](https://gitlab.com/gitlab-com/runbooks/-/tree/master/docs/zoekt)

### Logging

Selected logging documents and resources:

- docs/ 
- logging/ 
- [exercises](https://gitlab.com/gitlab-com/runbooks/-/tree/master/docs/logging/exercises) (e.g. searching logs in Kibana)
- [README.md](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/logging/README.md) (logging overview) 
- [quick-start](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/logging/README.md#quick-start)
- [what-are-we-logging](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/logging/README.md#what-are-we-logging)
- [searching-logs](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/logging/README.md#searching-logs)
- [logging-infrastructure-overview](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/logging/README.md#logging-infrastructure-overview)
- troubleshooting/ 
- [README.md](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/logging/troubleshooting/README.md)

### Internal DNS

- [Managing internal DNS](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/internal_dns.md)

### Debug and monitor

- [Tracing the source of an expensive query](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/tracing-app-db-queries.md)
- [Work with Kibana (logs view)](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/logging/README.md#searching-logs)

### Secrets

- [Working with Google Cloud secrets](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/working-with-gcloud-secrets.md)

### Security

- [Working with the CloudFlare WAF/CDN](https://gitlab.com/gitlab-com/runbooks/-/blob/master/howto/externalvendors/cloudflare.md)
- [OSQuery](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/osquery.md)

### Other

- [Setup oauth2-proxy protection for web based application](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/setup-oauth2-proxy-protected-application.md)
- [Register new domain(s)](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/domain-registration.md)
- [Manage DNS entries](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-dns-entries.md)
- [Setup and Use my Yubikey](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/yubikey.md)
- [Purge Git data](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/git/purge-git-data.md)
- [Getting Started with Kubernetes and GitLab.com](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/kube/k8s-gitlab.md)
- [Using Chatops bot to run commands across the fleet](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/deploycmd.md)

### Gitter

- [MongoDB operations](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/git/gitter/mongodb-operations.md)
- [Renew the Gitter TLS certificate](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/git/gitter/renew-certificates.md)

### Manage Package Signing Keys

- [Manage Package Signing Keys](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/uncategorized/manage-package-signing-keys.md)

### Other Servers and Services

- [GitHost / GitLab Hosted](https://gitlab.com/gitlab-com/runbooks/-/blob/master/docs/git/githost.md)

### Adding runbooks rules

- Make it quick - add links for checks
- Don't make me think - write clear guidelines, write expectations
- Recommended structure 
- Symptoms - how can I quickly tell that this is what is going on
- Pre-checks - how can I be 100% sure
- Resolution - what do I have to do to fix it
- Post-checks - how can I be 100% sure that it is solved
- Rollback - optional, how can I undo my fix

Inside of the [bin](https://gitlab.com/gitlab-com/runbooks/-/tree/master/bin) directory you can find a list of scripts that can help running repetitive commands or setting up your machine to debug the infrastructure. These scripts can be bash, ruby, python or any other executable.

`glsh` in the single entrypoint to interact with the [`bin`](https://gitlab.com/gitlab-com/runbooks/-/tree/master/bin) directory. For example if you can `glsh hello` it will check if `hello` file exists inside of [`bin`](https://gitlab.com/gitlab-com/runbooks/-/tree/master/bin) directory and execute it. You can also pass multiple arguments, that the script will have access to.

Demo: [https://youtu.be/RsGgxm55YBg](https://youtu.be/RsGgxm55YBg)

```plain text
glsh hello arg1 arg2
```

## Install

```plain text
git clone git@gitlab.com:gitlab-com/runbooks.git
cd runbooks
sudo make glsh-install
```

## Update

```plain text
glsh update
```

## Create a new command

1. 

Create a new file inside of [`bin`](https://gitlab.com/gitlab-com/runbooks/-/tree/master/bin) directory: `touch bin/hello`

1.  

Populate the file with the contents that you want. The command below updates the file with a simple `echo` command.

```plain text
cat > bin/hello <<EOF
#!/usr/bin/env bash

echo "Hello from glsh"
EOF
```

1. 

Make it executable: `chmod +x bin/hello`

1. 

Run it: `glsh hello`

## Summary

Usually, following a change to the rules, you can test your new additions using:

```plain text
make verify
```

Then, regenerate the rules using:

If you get errors while doing any of these steps, then read on for more details on how to set up your local environment.

## Generating a new runbooks image

To generate a new image you must follow the git commit guidelines below, this will trigger a semantic version bump which will then cause a new pipeline that will build and tag the new image

### Git Commit Guidelines

This project uses [Semantic Versioning](https://semver.org/). We use commit messages to automatically determine the version bumps, so they should adhere to the conventions of [Conventional Commits (v1.0.0-beta.2)](https://www.conventionalcommits.org/en/v1.0.0-beta.2/).

### TL;DR

- Commit messages starting with `fix:` trigger a patch version bump
- Commit messages starting with `feat:` trigger a minor version bump
- Commit messages starting with `BREAKING CHANGE:` trigger a major version bump.
- If you don't want to publish a new image, do not use the above starting strings.

### Automatic versioning

Each push to `master` triggers a [`semantic-release`](https://semantic-release.gitbook.io/semantic-release/) CI job that determines and pushes a new version tag (if any) based on the last version tagged and the new commits pushed. Notice that this means that if a Merge Request contains, for example, several `feat:` commits, only one minor version bump will occur on merge. If your Merge Request includes several commits you may prefer to ignore the prefix on each individual commit and instead add an empty commit summarizing your changes like so:

## Tool Versioning

This project has adopted [`asdf version-manager`](https://github.com/asdf-vm/asdf) for tool versioning. Using `asdf` is recommended, although not mandatory. Please note that if you chose not to use `asdf`, you'll need to ensure that all the required binaries, an the correct versions, are installed and on your path.

## Contributor Onboarding

If you would like to contribute to this project, follow these steps to get your local development environment ready-to-go:

1. Follow the common environment setup steps described in [https://gitlab.com/gitlab-com/gl-infra/common-ci-tasks/-/blob/main/docs/developer-setup.md](https://gitlab.com/gitlab-com/gl-infra/common-ci-tasks/-/blob/main/docs/developer-setup.md).
2. Run the `./scripts/prepare-dev-env.sh` to download and install development dependencies, configure [`pre-commit`](https://gitlab.com/gitlab-com/runbooks#pre-commit-hooks)[ hooks](https://gitlab.com/gitlab-com/runbooks#pre-commit-hooks) etc.
3. That's it. You should be ready!

### Dependencies and required tooling

Following tools and libraries are required to develop dashboards locally:

- Go programming language
- Ruby programming language
- `go-jsonnet` - Jsonnet implementation written in Go
- `jsonnet-bundler` - package manager for Jsonnet
- `jq` - command line JSON processor

You can install most of them using `asdf` tool.

### Manage your dependencies using `asdf`

Before using `asdf` for the first time, install all the plugins by running:

Running this command will automatically install the versions of each tool, as specified in the `.tool-versions` file.

You don't need to use `asdf`, but in such case you will need install all dependencies manually and track their versions.

### Keeping Versions in Sync between GitLab-CI and `asdf`

`asdf` (and `.tool-versions` generally) is the SSOT for tool versions used in this repository. To keep `.tool-versions` in sync with `.gitlab-ci.yml`, there is a helper script, `./scripts/update-asdf-version-variables.sh`.

### Process for updating a tool version

1. Update the version in `.tool-versions`
2. Run `asdf install` to install latest version
3. Run `./scripts/update-asdf-version-variables.sh` to update a refresh of the `.gitlab-ci-asdf-versions.yml` file
4. Commit the changes

### Go, Jsonnet

We use `.tool-versions` to record the version of go-jsonnet that should be used for local development. The `asdf` version manager is used by some team members to automatically switch versions based on the contents of this file. It should be kept up to date. The top-level `Dockerfile` contains the version of go-jsonnet we use in CI. This should be kept in sync with `.tool-versions`, and a (non-gating) CI job enforces this.

To install [go-jsonnet](https://github.com/google/go-jsonnet), you have a few options. We recommend using `asdf` and installing via `./scripts/install-asdf-plugins.sh`.

Alternatively, you could follow that project's README to install manually. Please ensure that you install the same version as specific in `.tool-versions`.

Or via homebrew:

### `jsonnet-tool`

[`jsonnet-tool`](https://gitlab.com/gitlab-com/gl-infra/jsonnet-tool) is a small home-grown tool for generating configuration from Jsonnet files. The primary reason we use it is because it is much faster than the bash scripts we used to use for the task. Some tasks have gone from 20+ minutes to 2.5 minutes.

We recommend using asdf to manage `jsonnet-tool`. The plugin will be installed when

### Ruby

Ruby is managed through `asdf`. The version of Ruby is configured via the `.tool-versions` file. Note that previously, contributors on this project needed to configure [`legacy_version_file = yes`](https://asdf-vm.com/manage/configuration.html#legacy-version-file) but this setting is no longer required.

## Test jsonnet files

There are 2 approaches to write a test for a jsonnet file:

- Use [`jsonnetunit`](https://github.com/yugui/jsonnetunit). This method is simple and straight-forward. This approach is perfect for writing unit tests that asserts the output of a particular method. The downside is that it doesn't support jsonnet assertion and inspecting complicated result is not trivial.
- When a jsonnet file becomes more complicated, consists of multiple conditional branches and chains of methods, we should think of writing integration tests for it instead. Jsonnet Unit doesn't serve this purpose very well. Instead, let's use Rspec. Note that we probably don't want to use RSpec for testing small jsonnet functions, the idea would more be for testing error cases or complicated scenarios where we need to be more expressive about the output we expect

We have two custom matchers for writing integration tests:

### Location of test files

-  

JsonnetUnit tests must stay in the same directory and have the same name as the jsonnet file being tested but ending in `_test.jsonnet`. Some examples:

- `services/stages.libsonnet` -> `services/stages_test.jsonnet`
- `libsonnet/toolinglinks/sentry.libsonnet` -> `libsonnet/toolinglinks/sentry_test.jsonnet`
-  

RSpec tests replicates the directory structure of the Jsonnet files inside `spec` directory and must end in `_spec.rb` suffixes. Some example:

- `libsonnet/toolinglinks/grafana.libsonnet` -> `spec/libsonnet/toolinglinks/grafana_spec.rb`
- `dashboards/stage-groups/stage-group-dashboards.libsonnet` -> `spec/dashboards/stage-groups/stage-group-dashboards_spec.rb`

### How to run tests?

- Run the full Jsonnet test suite in your local environment with `make test-jsonnet && bundle exec rspec`
- Run a particular Jsonnet unit test file with `scripts/jsonnet_test.sh periodic-thanos-queries/periodic-query_test.jsonnet`
- Run a particular Jsonnet integration test file with `bundle exec rspec spec/libsonnet/toolinglinks/grafana_spec.rb`

_Note_: Verify that you have all the jsonnet dependencies downloaded before attempting to run the tests, you can automatically download the necessary dependencies by running `make jsonnet-bundle`.

## Pre-commit hooks

This project supports a set of [`pre-commit`](https://pre-commit.com/) hooks which can assist catching CI validation errors before early. While they are not required, they are recommended.

After running the `./scripts/prepare-dev-env.sh` script as described in the [Contributor Onboarding](https://gitlab.com/gitlab-com/runbooks#contributor-onboarding) section, the `pre-commit` hooks will be automatically installed and ready to go.

When running `git commit`, the hooks will check all staged changes, ensuring that they are valid. The `pre-commit` checks may in some cases automatically fix any problems. If they do this, you'll need to stage the changes and try again.


