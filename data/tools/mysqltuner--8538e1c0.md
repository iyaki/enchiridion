---
title: "MySQLTuner"
notion_id: 8538e1c0-997c-4c12-9e1a-a922e17264be
notion_url: https://app.notion.com/p/MySQLTuner-8538e1c0997c4c129e1aa922e17264be
last_edited: 2023-04-25T14:06:00.000Z
source_url: https://github.com/major/MySQLTuner-perl/#
tags: ["English", "Databases", "Tool"]
---
[https://github.com/major/MySQLTuner-perl/#](https://github.com/major/MySQLTuner-perl/#)



MySQLTuner is a script written in Perl that allows you to review a MySQL installation quickly and make adjustments to increase performance and stability. The current configuration variables and status data is retrieved and presented in a brief format along with some basic performance suggestions.

MySQLTuner supports ~900+ indicators, KPIs, and recommendations (including Weighted Health Score, Predictive Capacity Planning, and SSL/TLS Audit) for MySQL/MariaDB/Percona Server in this latest version.

MySQLTuner is actively maintained supporting many configurations such as Galera Cluster, TokuDB, Performance schema, Linux OS metrics, InnoDB, MyISAM, Aria, ...

You can find more details on these indicators here: Indicators description.

## Useful Links

- Active Development: https://github.com/major/MySQLTuner-perl
- Releases/Tags: https://github.com/major/MySQLTuner-perl/tags
- Changelog: https://github.com/major/MySQLTuner-perl/blob/master/Changelog
- Docker Images: https://hub.docker.com/repository/docker/jmrenouard/mysqltuner/tags
- Useful References: Documentation/References
- AI Agent & MCP Server Integration Guide: Documentation/AI MCP Server Guide | Guide Serveur MCP (FR) | AGENT.md
- Interactive HTML Reports (v2.9.0+): MariaDB 11.4 E2E HTML Report Example MySQL 8.4 E2E HTML Report Example Percona 8.0 E2E HTML Report Example

## MySQLTuner needs you

MySQLTuner needs contributors for documentation, code and feedback:

- Please join us on our issue tracker at GitHub tracker.
- Contribution guide is available following MySQLTuner contributing guide
- Star MySQLTuner project at MySQLTuner Git Hub Project
- Paid support for LightPath here: jmrenouard@lightpath.fr
- Paid support for Releem available here: Releem App

### Sponsors

Active development is sponsored by:

Thanks to LightPath for providing resources (dev servers, AI subscriptions, staging & features).

## Star History

## Compatibility

Test result are available here for LTS only:

- MySQL (full support)
- Percona Server (full support)
- MariaDB (full support)
- Galera replication (full support)
- Percona XtraDB cluster (full support)
- MySQL Replication (partial support, no test environment)

Thanks to endoflife.date

- Refer to MariaDB Supported versions.
- Refer to MySQL Supported versions.

Windows Support is partial

- Windows is now supported at this time
- Successfully run MySQLtuner across WSL2 (Windows Subsystem Linux)
- https://docs.microsoft.com/en-us/windows/wsl/

UNSUPPORTED ENVIRONMENTS - NEED HELP WITH THAT

Advanced Intelligence & Ecosystem

- Weighted Health Score KPI: Overall database health assessment (0-100) based on Performance (40pts), Security (30pts), and Resilience (30pts).
- Smart Migration LTS Advisor: Identification of risks when migrating to modern LTS versions (MySQL 8.4/9.0+, MariaDB 11.x), including removed variables and deprecated authentication methods.
- Predictive Capacity Planning: Memory headroom analysis (peak vs available RAM+Swap), disk growth forecasting, and AUTO_INCREMENT capacity near max value detection.
- Cloud Autodiscovery: Native support for AWS RDS/Aurora, GCP Cloud SQL, Azure (Flexible/Managed), and DigitalOcean. Automatic detection via @@version_comment and provider-specific variables.
- Infrastructure-Aware Tuning: Detection of SSD/NVMe vs HDD storage types and ARM64/Graviton vs x86_64 architectures. Adjusts recommendations for innodb_flush_neighbors and innodb_io_capacity.
- SSL/TLS Security Audit: Session encryption check, TLS version audit (warn on TLSv1.0/1.1), certificate expiration, require_secure_transport enforcement, and remote user SSL checks.
- Authentication Plugin Auditing: Detection of insecure plugins (mysql_native_password, sha256_password), MySQL 9.x readiness diagnostics, and MariaDB ed25519/unix_socket recommendations.
- Schema Modeling & Naming Conventions: Comprehensive table structure analysis (missing PKs, surrogate key types, UTF-8 compliance, non-InnoDB tables), naming convention audit (snake_case/camelCase consistency, plural detection, boolean/date column prefixes), and foreign key analysis (unconstrained _id columns, type mismatches, CASCADE audit).
- MySQL 8.0+ / MariaDB Modeling: JSON column indexability (virtual generated columns), invisible indexes, CHECK constraints.
- Guided Auto-Fix Engine: Generation of ready-to-apply SET GLOBAL SQL statements and [mysqld] configuration blocks from the variable adjustment recommendations.
- Historical Trend Analysis: Ingest JSON output from previous runs via --compare-file to track QPS and data growth trends.
- Sysbench Integration: Parse sysbench output for QPS, TPS, and latency metrics (Avg/95th/Max) via --sysbench-file.
- Container & Systemd Log Integration: Automatic log detection from Docker, Podman, Kubectl/Kubernetes, and Systemd journal.
- AI & MCP Protocol Support: Native JSON-RPC stdio Model Context Protocol (MCP) server daemon and --agent-json structured output for AI client tools (Claude Desktop, Cursor, VS Code / Cline, Antigravity).

## 🤖 AI Agent & Model Context Protocol (MCP) Integration

MySQLTuner natively supports modern Artificial Intelligence (AI) workflows, autonomous DBA agents, and developer assistants (e.g., Claude Desktop, Cursor IDE, VS Code / Cline, Antigravity, and LangChain/LlamaIndex frameworks).

For complete technical documentation, refer to the AI & MCP Integration Guide, Guide Serveur MCP IA (FR), and AGENT.md.

### Operating Modes

1. Direct CLI Telemetry (--agent-json): Outputs zero-dependency structured JSON containing findings, impact scores (1-10), risk levels (Low, Medium, High, Critical), executable SET GLOBAL SQL statements, and pre-calculated rollback_statement baselines. perl mysqltuner.pl --agent-json --host 127.0.0.1 --user root --pass secret
2. Model Context Protocol (MCP) Server: A lightweight microservice (build/mcp_server.py and Dockerfile.mcp) communicating over standard I/O (stdio) via JSON-RPC 2.0. Resources: mysqltuner://reports/latest.json, mysqltuner://indicators/summary.json Tools: get_latest_audit, run_audit, apply_recommendation, rollback_recommendation docker run -d \ --name mysqltuner-mcp \ -e DB_HOST=127.0.0.1 -e DB_USER=root -e DB_PASSWORD=secret \ mysqltuner-mcp

## Unsupported storage engines: PRs welcome

- NDB is not supported feel free to create a Pull Request
- Archive
- Spider
- ColummStore
- Connect

## Unmaintenained stuff from MySQL or MariaDB

- MyISAM is too old and no longer active
- RockDB is not maintained anymore
- TokuDB is not maintained anymore
- XtraDB is not maintained anymore
- CVE vulnerabilities detection support from https://cve.mitre.org

MINIMAL REQUIREMENTS

- Perl 5.6 or later (with perl-doc package)
- Unix/Linux based operating system (tested on Linux, BSD variants, and Solaris variants)
- Unrestricted read access to the MySQL server (see Privileges below)

## PRIVILEGES

To run MySQLTuner with all features, the following privileges are required:

MySQL 8.0+:

```plain text
GRANT SELECT, PROCESS, SHOW DATABASES, EXECUTE, REPLICATION REPLICA, REPLICATION CLIENT, SHOW VIEW ON *.* TO 'mysqltuner'@'localhost';
```

MariaDB 10.5+:

```plain text
GRANT SELECT, PROCESS, SHOW DATABASES, EXECUTE, BINLOG MONITOR, SHOW VIEW, REPLICATION SOURCE ADMIN, REPLICA MONITOR ON *.* TO 'mysqltuner'@'localhost';
```

Legacy versions:

```plain text
GRANT SELECT, PROCESS, EXECUTE, REPLICATION CLIENT, SHOW DATABASES, SHOW VIEW ON *.* TO 'mysqltuner'@'localhost';
```

OS root access recommended for MySQL < 5.1

## WARNING

It is important for you to fully understand each change you make to a MySQL database server. If you don't understand portions of the script's output, or if you don't understand the recommendations, you should consult a knowledgeable DBA or system administrator that you trust. Always test your changes on staging environments, and always keep in mind that improvements in one area can adversely affect MySQL in other areas.

It's also important to wait at least 24 hours of uptime to get accurate results. In fact, running mysqltuner on a fresh restarted server is completely useless.

Also review the FAQ section below.

## Security recommendations

Hi directadmin user! We detected that you run mysqltuner with da_admin's credentials taken from /usr/local/directadmin/conf/my.cnf, which might bring to a password discovery! Read link for more details Issue #289.

## What is MySQLTuner checking exactly ?

All checks done by MySQLTuner are documented in MySQLTuner Internals documentation.

MySQLTuner analyzes the following areas:

- System & OS: RAM, swap, open ports, kernel parameters, load average, mount points, network cards
- Server Version: EOL detection, architecture, 64-bit recommendations
- Error Logs: Local files, Docker/Podman containers, Kubernetes pods, Systemd journal
- Cloud & Infrastructure: AWS RDS/Aurora, GCP, Azure, DigitalOcean; SSD/NVMe vs HDD; ARM64/x86_64
- Storage Engines: InnoDB (buffer pool, redo log, chunk size), MyISAM, Aria, Galera, TokuDB, RocksDB
- Security: Anonymous users, weak passwords, SSL/TLS audit, authentication plugins, CVE vulnerabilities
- Connections: Usage percentages, aborted connections, thread cache
- Performance: Sort/join/temp tables, global buffers, query cache, slow queries, memory usage
- Replication: Source/Replica status, lag, GTID, semi-sync, multi-source
- Performance Schema: Top users/hosts/statements, IO latency, lock waits, unused indexes, redundant indexes
- Schema Modeling: Primary key analysis, naming conventions, foreign keys, data types, UTF-8 compliance, JSON indexability
- Predictive: Memory headroom, disk growth forecasting, AUTO_INCREMENT capacity
- Health Score: Weighted KPI (0-100) aggregating Performance, Security, and Resilience findings

## Download/Installation

> Note: Linux distribution packages (e.g. apt install mysqltuner on Ubuntu/Debian, yum/dnf on RHEL/CentOS/Fedora) often ship a significantly older version of MySQLTuner. There is no official distribution-maintained repository that tracks the latest release. To always get the latest version, use one of the direct download methods below.

Choose one of these methods:

1. Script direct download (the simplest and shortest method):

```plain text
wget https://mysqltuner.pl/ -O mysqltuner.pl wget https://raw.githubusercontent.com/major/MySQLTuner-perl/master/basic_passwords.txt -O basic_passwords.txt wget https://raw.githubusercontent.com/major/MySQLTuner-perl/master/vulnerabilities.csv -O vulnerabilities.csv
```

1. You can download the entire repository by using git clone or git clone --depth 1 -b master followed by the cloning URL above.

```plain text
git clone --depth 1 -b master https://github.com/major/MySQLTuner-perl.git
```

1. On Apple macOS, install via Homebrew:

```plain text
brew install mysqltuner
```

1. If you are in an air-gapped environment without direct internet access, download the files on a machine that has internet access (or via a proxy host), then copy mysqltuner.pl, basic_passwords.txt, and vulnerabilities.csv to your target server.
2. Docker: Pull and run the official Docker container (Docker hub tags are available at jmrenouard/mysqltuner tags):

```plain text
docker pull jmrenouard/mysqltuner:latest # Basic execution docker run --rm -it jmrenouard/mysqltuner --host <database_host> --user <username> --pass <password> # Save generated reports (HTML/TXT) to host filesystem: docker run --rm -it -v $(pwd)/results:/results jmrenouard/mysqltuner --host <database_host> --user <username> --pass <password> # Mount custom configuration / defaults file: docker run --rm -it -v $(pwd)/my.cnf:/defaults.cnf -v $(pwd)/results:/results jmrenouard/mysqltuner --host <database_host> --user <username> --pass <password>
```

### Releases Location

- Official release notes and history are documented in the releases/ directory of this repository (e.g., releases/v2.9.2.md).
- Git release tags and downloadable source tarballs are available on GitHub Releases.

## Optional Sysschema installation for MySQL 5.6

Sysschema is installed by default under MySQL 5.7 and MySQL 8 from Oracle. By default, on MySQL 5.6/5.7/8, performance schema is enabled. For previous MySQL 5.6 version, you can follow this command to create a new database sys containing very useful view on Performance schema:

## Sysschema for MySQL old version

```plain text
curl "https://codeload.github.com/mysql/mysql-sys/zip/master" > sysschema.zip # check zip file unzip -l sysschema.zip unzip sysschema.zip cd mysql-sys-master mysql -uroot -p < sys_56.sql
```

## Sysschema for MariaDB old version

```plain text
curl "https://github.com/FromDual/mariadb-sys/archive/refs/heads/master.zip" > sysschema.zip # check zip file unzip -l sysschema.zip unzip sysschema.zip cd mariadb-sys-master mysql -u root -p < ./sys_10.sql
```

## Performance schema setup

By default, performance_schema is enabled and sysschema is installed on latest version.

By default, on MariaDB, performance schema is disabled (MariaDB<10.6).

Consider activating performance schema across your my.cnf configuration file:

```plain text
[mysqld] performance_schema = on performance-schema-consumer-events-statements-history-long = ON performance-schema-consumer-events-statements-history = ON performance-schema-consumer-events-statements-current = ON performance-schema-consumer-events-stages-current=ON performance-schema-consumer-events-stages-history=ON performance-schema-consumer-events-stages-history-long=ON performance-schema-consumer-events-transactions-current=ON performance-schema-consumer-events-transactions-history=ON performance-schema-consumer-events-transactions-history-long=ON performance-schema-consumer-events-waits-current=ON performance-schema-consumer-events-waits-history=ON performance-schema-consumer-events-waits-history-long=ON performance-schema-instrument='%=ON' max-digest-length=2048 performance-schema-max-digest-length=2018
```

## Sysschema installation for MariaDB < 10.6

Sysschema is not installed by default under MariaDB prior to 10.6 MariaDB sys

You can follow this command to create a new database sys containing a useful view on Performance schema:

```plain text
curl "https://codeload.github.com/FromDual/mariadb-sys/zip/master" > mariadb-sys.zip # check zip file unzip -l mariadb-sys.zip unzip mariadb-sys.zip cd mariadb-sys-master/ mysql -u root -p < ./sys_10.sql
```

## Errors & solutions for performance schema installation

## ERROR 1054 (42S22) at line 78 in file: './views/p_s/metrics_56.sql': Unknown column 'STATUS' in 'field list'

This error can be safely ignored Consider using a recent MySQL/MariaDB version to avoid this kind of issue during sysschema installation

In recent versions, sysschema is installed and integrated by default as sys schema (SHOW DATABASES)

## ERROR at line 21: Failed to open file './tables/sys_config_data_10.sql -- ported', error: 2 Have a look at #452 solution given by @ericx

Fixing sysctl configuration (/etc/sysctl.conf)

-- It is a system wide setting and not a database setting: Linux FS Kernel settings

You can check its values via:

```plain text
$ cat /proc/sys/fs/aio-* 65536 2305
```

For example, to set the aio-max-nr value, add the following line to the /etc/sysctl.conf file:

```plain text
fs.aio-max-nr = 1048576
```

To activate the new setting:

```plain text
sysctl -p /etc/sysctl.conf
```

## Specific usage

Usage: Minimal usage locally

```plain text
perl mysqltuner.pl --host 127.0.0.1
```

Of course, you can add the execute bit (chmod +x mysqltuner.pl) so you can execute it without calling Perl directly.

Usage: Minimal usage remotely

In previous version, --forcemem shoud be set manually, in order to be able to run an MySQLTuner analysis

Since 2.1.10, memory and swap are defined to 1Gb by default.

If you want a more accurate value according to your remote server, feel free to setup --forcemem and --forceswap to real RAM value

```plain text
perl mysqltuner.pl --host targetDNS_IP --user admin_user --pass admin_password
```

Usage: Enable maximum output information around MySQL/MariaDb without debugging

```plain text
perl mysqltuner.pl --verbose perl mysqltuner.pl --buffers --dbstat --idxstat --sysstat --pfstat --tbstat
```

Usage: Enable CVE vulnerabilities check for your MariaDB or MySQL version

```plain text
perl mysqltuner.pl --cvefile=vulnerabilities.csv
```

Usage: Write your result in a file with information displayed

```plain text
perl mysqltuner.pl --outputfile /tmp/result_mysqltuner.txt
```

Usage: Write your result in a file without outputting information

```plain text
perl mysqltuner.pl --silent --outputfile /tmp/result_mysqltuner.txt
```

Usage: Generate a standalone HTML report file (built-in, requires no CPAN or external modules)

```plain text
perl mysqltuner.pl --reportfile=mysqltuner.html
```

Usage: Dumping all information_schema and sysschema views as csv file into results subdirectory

```plain text
perl mysqltuner.pl --verbose --dumpdir=./result
```

Usage: Enable debugging information

```plain text
perl mysqltuner.pl --debug
```

Usage: Update MySQLTuner and data files (password and cve) if needed

```plain text
perl mysqltuner.pl --checkversion --updateversion
```

Usage: Integrating Sysbench performance results

```plain text
perl mysqltuner.pl --sysbench-file=/path/to/sysbench_output.txt
```

Usage: Historical Trend Analysis (compare with previous run)

```plain text
perl mysqltuner.pl --json --outputfile=run1.json # ... some time later ... perl mysqltuner.pl --compare-file=run1.json
```

Usage: Export one Markdown file per schema (schema documentation)

```plain text
perl mysqltuner.pl --verbose --schemadir=./schemas
```

Usage: Dump data with row limits and gzip compression

```plain text
perl mysqltuner.pl --verbose --dumpdir=./result --dump-limit=10000 --compress-dump
```

Usage: Container mode (analyze a database running in Docker)

```plain text
perl mysqltuner.pl --verbose --container docker:mysql_container_name
```

Usage: Table structure and naming convention analysis

```plain text
perl mysqltuner.pl --structstat
```

Usage: Filter output (show only problems)

```plain text
perl mysqltuner.pl --nogood --noinfo
```

Usage: JSON output (for automation and reporting pipelines)

```plain text
perl mysqltuner.pl --json --outputfile=report.json perl mysqltuner.pl --prettyjson
```

Usage: Non-dedicated server mode (shared hosting)

```plain text
perl mysqltuner.pl --nondedicated
```

Usage: Fast Analysis by Bypassing Table Workload Scans (Issue #986)

```plain text
perl mysqltuner.pl --skipworkload
```

Usage: AI Agent Remediation Schema (Actionable JSON)

```plain text
perl mysqltuner.pl --agent-json --outputfile=remediation.json
```

Usage: Use credentials from environment variables

```plain text
export MYSQL_USER=mysqltuner export MYSQL_PASS=secret perl mysqltuner.pl --userenv=MYSQL_USER --passenv=MYSQL_PASS
```

Usage: Use a custom defaults file

```plain text
perl mysqltuner.pl --defaults-file=/path/to/my.cnf
```

For a complete list of all available options, run perl mysqltuner.pl --help or refer to the USAGE.md documentation.

## Cloud Support

MySQLTuner now has experimental support for cloud-based MySQL services.

- --cloud: Enable cloud mode. This is a generic flag for any cloud provider.
- --azure: Enable Azure-specific support.
- --ssh-host <hostname>: The SSH host for cloud connections.
- --ssh-user <username>: The SSH user for cloud connections.
- --ssh-password <password>: The SSH password for cloud connections.
- --ssh-identity-file <path>: The path to the SSH identity file for cloud connections.

## HTML Report & Weighted Health Score

MySQLTuner calculates a Weighted Health Score KPI (overall database health assessment on a scale of 0 to 100) dynamically based on three categories:

1. Performance (40 points max): Evaluation of buffer pool read efficiency, disk temporary tables ratio, thread cache hit rate, and connection utilization.
2. Security (30 points max): Evaluation of user account configurations, weak passwords (checked offline), SSL/TLS session encryption, and authentication plugin usage.
3. Resilience (30 points max): Evaluation of replication status and lag, log setups, and database modeling findings.

Generating the HTML Report

You can generate a standalone HTML report directly with:

```plain text
perl mysqltuner.pl --reportfile=mysqltuner.html
```

This feature is built natively in pure Perl and has zero external dependencies (no CPAN modules or Python packages are required). The generated report provides an interactive dark-themed dashboard displaying:

- Overall Health Score gauge
- Detailed KPI metrics breakdown (Performance, Security, Resilience)
- Categorized recommendations lists (General, Variables to Adjust, Database Modeling, Security, System)
- Collapsible full console output log

## FAQ

Question: What are the prerequisites for running MySQL tuner ?

Before running MySQL tuner, you should have the following:

- A MySQL server installation
- Perl installed on your system
- Administrative access to your MySQL server

Question: Can MySQL tuner make changes to my configuration automatically ?

No., MySQL tuner only provides recommendations. It does not make any changes to your configuration files automatically. It is up to the user to review the suggestions and implement them as needed.

Question: How often should I run MySQL tuner ?

It is recommended to run MySQL tuner periodically, especially after significant changes to your MySQL server or its workload.

For optimal results, run the script after your server has been running for at least 24 hours to gather sufficient performance data.

Question: How do I interpret the results from MySQL tuner ?

MySQL tuner provides output in the form of suggestions and warnings.

Review each recommendation and consider implementing the changes in your MySQL configuration file (usually my.cnf or my.ini).

Be cautious when making changes and always backup your configuration file before making any modifications.

Question: Can MySQL tuner cause harm to my database or server ?

While MySQL tuner itself will not make any changes to your server, blindly implementing its recommendations without understanding the impact can cause issues.

Always ensure you understand the implications of each suggestion before applying it to your server.

Question: Does MySQL tuner support MariaDB and Percona Server ?

Yes, MySQL tuner supports MariaDB and Percona Server since they are derivatives of MySQL and share a similar architecture. The script can analyze and provide recommendations for these systems as well.

Question: What should I do if I need help with MySQL tuner or have questions about the recommendations ?

If you need help with MySQL tuner or have questions about the recommendations provided by the script, you can consult the MySQL tuner documentation, seek advice from online forums, or consult a MySQL expert.

Be cautious when implementing changes to ensure the stability and performance of your server.

Question: Will MySQLTuner fix my slow MySQL server ?

No. MySQLTuner is a read only script. It won't write to any configuration files, change the status of any daemons. It will give you an overview of your server's performance and make some basic recommendations for improvements that you can make after it completes.

Question: Can I fire my DBA now?

MySQLTuner will not replace your DBA in any form or fashion.

If your DBA constantly takes your parking spot and steals your lunch from the fridge, then you may want to consider it - but that's your call.

Once you create it, make sure it's owned by your user and the mode on the file is 0600. This should prevent the prying eyes from getting your database login credentials under normal conditions.

Question: I get "ERROR 1524 (HY000): Plugin 'unix_socket' is not loaded" even with unix_socket=OFF. How to fix?

This occurs because the MariaDB client attempts to use the unix_socket plugin by default when no user/password is provided.

- Solution 1 (Recommended): Use a ~/.my.cnf file as described above to provide explicit credentials.
- Solution 2: Pass credentials directly: perl mysqltuner.pl --user root --pass your_password.

Question: How to securely re-enable unix_socket authentication?

If you decide to use unix_socket (which allows the OS root user to log in to MariaDB root without a password), follow these steps:

1. Ensure the plugin is enabled in /etc/my.cnf: unix_socket=ON (or remove OFF).
2. In MariaDB, set the authentication plugin for the root user: ALTER USER 'root'@'localhost' IDENTIFIED VIA unix_socket;
3. Verify that the auth_socket or unix_socket plugin is ACTIVE in SHOW PLUGINS.

Question: Is there another way to secure credentials on latest MySQL and MariaDB distributions ?

You could use mysql_config_editor utilities.

```plain text
$ mysql_config_editor set --login-path=client --user=someusername --password --host=localhost Enter password: ********
```

After which, ~/.mylogin.cnf will be created with the appropriate access.

To get information about stored credentials, use the following command:

```plain text
$mysql_config_editor print [client] user = someusername password = ***** host = localhost
```

Question: It's not working on my OS! What gives?!

These kinds of things are bound to happen. Here are the details I need from you to investigate the issue:

- OS and OS version
- Architecture (x86, x86_64, IA64, Commodore 64)
- Exact MySQL version
- Where you obtained your MySQL version (OS package, source, etc)
- The full text of the error
- A copy of SHOW VARIABLES and SHOW GLOBAL STATUS output (if possible)

Question: How to perform CVE vulnerability checks?

- Download vulnerabilities.csv from this repository.
- use option --cvefile to perform CVE checks

Question: How to use mysqltuner from a remote host? Thanks to @rolandomysqldba

- You will still have to connect like a mysql client:

Connection and Authentication

```plain text
--host <hostname> Connect to a remote host to perform tests (default: localhost) --socket <socket> Use a different socket for a local connection --port <port> Port to use for connection (default: 3306) --user <username> Username to use for authentication --pass <password> Password to use for authentication --defaults-file <path> defaults file for credentials
```

Since you are using a remote host, use parameters to supply values from the OS

```plain text
--forcemem <size> Amount of RAM installed (in megabytes or with units, e.g. 15G, 1024M) --forceswap <size> Amount of swap memory configured (in megabytes or with units)
```

- You may have to contact your remote SysAdmin to ask how much RAM and swap you have

If the database has too many tables, or very large table, use this:

```plain text
--skipsize Don't enumerate tables and their types/sizes (default: on) (Recommended for servers with many tables)
```

Question: Can I install this project using homebrew on Apple Macintosh?

Yes! brew install mysqltuner can be used to install this application using homebrew on Apple Macintosh.

Question: I installed MySQLTuner via my Linux distribution's package manager (apt/yum/dnf). How do I get the latest version?

Linux distributions such as Ubuntu, Debian, RHEL, and CentOS often ship an older version of MySQLTuner in their official repositories. For example, Ubuntu 22.04 ships version 1.7.17 while the latest release may be significantly newer.

There is currently no official APT/YUM/DNF repository that tracks the latest MySQLTuner release. To get the latest version, use one of these methods:

- Direct download (recommended):

```plain text
wget https://mysqltuner.pl/ -O mysqltuner.pl wget https://raw.githubusercontent.com/major/MySQLTuner-perl/master/basic_passwords.txt -O basic_passwords.txt wget https://raw.githubusercontent.com/major/MySQLTuner-perl/master/vulnerabilities.csv -O vulnerabilities.csv chmod +x mysqltuner.pl
```

- Git clone:

```plain text
git clone --depth 1 -b master https://github.com/major/MySQLTuner-perl.git cd MySQLTuner-perl perl mysqltuner.pl
```

- Air-gapped environments: If your server has no direct internet access, download the files above on a host that has internet access (or via a proxy), then transfer mysqltuner.pl, basic_passwords.txt, and vulnerabilities.csv to the target server using scp, rsync, or another file transfer method.

## MySQLTuner and Vagrant (Legacy)

> Note: The Vagrant-based test environment is considered legacy. For modern testing, use the Docker-based test suite via make test-it or build/test_envs.sh.

Vagrant File is stored in the Vagrant subdirectory.

## Setup Docker test environments

MySQLTuner includes a Docker-based test infrastructure for multi-version validation:

```plain text
# Create and start all test containers sh build/createTestEnvs.sh # Source environment helpers source build/bashrc # Connect to a specific database mysql_percona80 sakila
```

Supported test targets (refer to MariaDB support and MySQL support for the current compatibility matrix):

- MySQL 8.0, 8.4, 9.x
- MariaDB 10.6, 10.11, 11.4, 11.8
- Percona Server 8.0

## Contributions welcome

How to contribute using Pull Request ? Follow this guide : Pull request creation

## Simple steps to create a pull request

- Fork this Github project
- Clone it to your local system
- Make a new branch
- Make your changes
- Push it back to your repo
- Click the Compare & pull request button
- Click Create pull request to open a new pull request
