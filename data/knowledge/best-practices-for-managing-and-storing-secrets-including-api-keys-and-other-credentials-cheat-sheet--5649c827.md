---
title: "Best practices for managing and storing secrets including API keys and other credentials [cheat sheet included]"
notion_id: 5649c827-cc83-474d-9254-7a8ce85c6743
notion_url: https://app.notion.com/p/Best-practices-for-managing-and-storing-secrets-including-API-keys-and-other-credentials-cheat-shee-5649c827cc83474d92547a8ce85c6743
last_edited: 2023-04-25T14:07:00.000Z
source_url: https://blog.gitguardian.com/secrets-api-management/
tags: ["Article", "CheatSheet", "GitGuardian Blog", "English", "Programming", "Information Security", "DevOps"]
---
[https://blog.gitguardian.com/secrets-api-management/](https://blog.gitguardian.com/secrets-api-management/)

👉

TL;DR: Master API key management best practices by never storing unencrypted secrets in git, enforcing automated secrets scanning, and avoiding plaintext sharing in messaging systems. Use dedicated secrets management tools, apply least privilege and short-lived keys, and implement robust rotation and monitoring strategies. Proactive governance and continuous audit are critical to minimizing risk, preventing secret sprawl, and maintaining compliance in modern development environments.

## Best Practices for Managing and Storing Secrets Including API Keys and Other Credentials [cheat sheet included]

We have compiled a list of some of the best practices to prevent API key leakage and keep secrets and credentials safe. Secrets management doesn’t have a one-size-fits-all approach, so this list considers multiple perspectives so you can be informed in deciding to or not to implement strategies.

### Table of contents

Storing and managing secrets like API keys and other credentials can be challenging. Even the most careful policies can sometimes be circumvented in exchange for convenience. We have compiled a list of some of the best practices to prevent API key leakage and keep secrets and credentials safe.

Secrets management doesn’t have a one-size-fits-all approach, so this list considers multiple perspectives so you can be informed in deciding to or not to implement strategies. Download the cheat sheet.

### Directory

- Never store unencrypted secrets in .git repositories Avoid git add * commands on git Add sensitive files in .gitignore Don’t rely on code reviews to discover secrets Use automated secrets scanning on repositories
- Don’t share your secrets unencrypted in messaging systems like Slack
- Store secrets safely
- Restrict API keys access and permissions

## Never store unencrypted secrets in .git repositories

It is common to wrongly assume that private repositories are secure vaults that are safe places to store secrets.

Private repositories are not appropriate places to store secrets.

Private repositories are high-value targets for bad actors because it is common practice to store secrets within them.

In addition, .git is designed to sprawl. Repositories get cloned onto new machines, forked into new projects and new developers regularly enter and exit a project with access to complete history. Any hard-coded secrets that exist within a private repository's history will exist in all new repositories born from that source.

If a secret enters a repository, private or public, then it should be considered compromised.

> A secret in a private repo is like a password written on a $20 bill, you might trust the person you gave it to, but that bill can end up in hundreds of peoples hands as a part of multiple transactions and within multiple cash registers.

If already committed and want to remove an API key from the git history, follow the guidelines below:

Git Clean, Git Remove file from commit - Cheatsheet - GitGuardian Blog

Exposing secrets in a git repository is bad but mistakes happen. This is a complete guide and cheatsheet to rewrite and remove files from git history

GitGuardian Blog - Automated Secrets DetectionGuest Expert

## Avoid git add * commands on git

Using wildcard commands like git add * or git add . can easily capture files that should not enter a git repository. This includes generated files, config files, and temporary source code.

To avoid that, when making a commit, you should preferably add each file by name and use git status to list tracked and untracked files. According to git-scm:

> “Remember that each file in your working directory can be in one of two states: tracked or untracked.

> Tracked files are files that were in the last snapshot; they can be unmodified, modified, or staged. In short, tracked files are files that Git knows about.Untracked files are everything else.”

Advantages

- Complete control and visibility over what files are committed
- Reduces the risk of unwanted files entering source control and API key leak
- Requires thought and consideration when adding files

Disadvantages

- Takes additional time when making a commit
- Can mistakenly miss files when committing

> TIP: Committing early and committing often will not only help navigate file history and break up otherwise large tasks, in addition it will reduce the temptation to use wildcard commands.

## Add sensitive files in .gitignore

To prevent sensitive files from ending up within git repositories a comprehensive .gitignore file should be included with all repositories and include:

- Files with environment variables like .env or configuration files like .zshrc, application.properties or .config
- Files generated by another process (such as application logs or checkpoints, unit tests/coverage reports)
- Files containing “real” data (other than test data) like database extracts.

GitHub published a collection of useful .gitignore templates. See these examples for Python or Go.

## Don’t rely on code reviews to discover secrets

It is extremely important to understand that code reviews will not always detect hard-coded secrets, especially if they are hidden in previous versions of code. The reason code reviews are not adequate protection is because reviewers are only concerned with the difference between the current and proposed states of the code; they do not consider the entire history of the project.

Why a code review won't detect a deleted API key

If secrets are committed to a development branch and later removed, these secrets won’t be visible or of importance to the reviewer. The nature of git means that if a secret gets overlooked in history it is exposed forever as anyone with access to the repository can find this secret in previous revisions of the codebase.

> TIP: As a rule, automation should be implemented wherever predefined rules can be established, like secrets detection. Human reviews should be left to check code for errors that cannot be easily predefined, such as logic.

## Use automated secrets scanning on repositories

Even when all best practices are followed, mistakes are common. When dealing with highly sensitive data, no chances should be taken. GitGuardian offers a free secrets scanning solution for developers to detect both generic API keys and specific secrets (more than 450 providers are supported!).

But the best is that it can be installed both on private and public repositories for free!

Visibility is the key to great secret management. If you don’t know you have a problem, you cannot take action to fix it. Secrets scanning provides essential visibility over your internal systems.

It is important to also consider that even the best secrets management systems and policies do not prevent newly generated secrets from entering the code base or old secrets from being extracted and included again.

Advantages

- Difficult to circumvent and ignore compared to tools that need to be manually run
- Much faster and more accurate than relying on human checking
- Can detect secrets buried within logs and history that manual reviews and searches will not uncover
- Live scanning ensures all active data leaks are captured

A common secret sprawl enabler is sending secrets in plain text over messaging services. While these systems are intended to keep messages secure, they are not intended to hold sensitive information such as secrets.

These systems are high-value targets for attackers. It only takes one compromised email or Slack account to uncover a trove of sensitive information, as secrets exposure extends beyond source code into various communication and collaboration tools. If secrets are being sent over internal systems, it also makes it possible for bad actors to move laterally between services by ‘using secrets to find secrets’.

## Store secrets safely

There is no silver bullet solution for secrets management. Different factors such as project size, team geography, and project scope, must be considered. Multiple solutions may need to coexist. Carefully consider each option, not just to meet your current needs but also to consider how each solution will scale with the growth of your project.

## Use encryption to store secrets within .git repositories

Encrypting your secrets using common tools such as git secret or SOPS and storing them within a git repository can be beneficial when working in teams as it keeps secrets synced.

However, it's important to note that this approach introduces a new challenge—managing additional security keys for encryption and decryption, as well as securely sharing them. This aspect may seem like an ongoing concern, but it's a vital aspect to address for robust security practices.

Advantages

- Your secrets are synced

Disadvantages

- You have to deal with your encryption keys securely
- No audit logs (who accessed which secret and when)
- No role-based access control (RBAC)
- Hard to rotate access. Rotating access implies revoking the key and redistributing it. The distribution part is not easy to handle with git repositories when dealing with multiple developers

## Use local environment variables when feasible

An environment variable is a dynamic object whose value is set outside of the application. This makes them easier to rotate without having to make changes within the application itself. It also removes the need to have these written within source code, making them more appropriate for handling sensitive data.

Advantages

- They are easy to change between deployed versions without changing any code
- They are less likely to be checked into the repository
- Simple and clean

Disadvantages

- This approach may not be feasible at scale when working in teams because there is no way to easily keep developers, applications, and/or infrastructure in sync.

## Use Secrets Management Tools to Enhance API Security

Secrets management systems such as Hashicorp Vault or AWS Key Management Service are encrypted systems that can safely store your secrets and tightly control access. Vaults and other managed secrets solutions are not appropriate in all cases because they are complicated to set up and need to be well maintained. Both take a considerable investment of resources.

Advantages

- It prevents secrets from sprawling
- It provides audit logs

Disadvantages

- As they introduce a single point of failure, they must be hosted on a highly-available and secure infrastructure
- All the codebase must be changed to integrate with them
- Keys giving access to the system must be carefully protected

## Restrict API access and permissions

It can be difficult to detect when an attacker is using secrets like API keys maliciously because they are often using them within their scope. By restricting the access and permissions of the API key you not only limit damage and restrict lateral movement but also provide greater visibility over when the API key is being used outside of its scope.

💡

In fact, this idea can be pushed even further by using API keys as a decoy to intercept hackers! This concept is called a honeytoken. Learn more about the Honeytoken module in GitGuardian.

## Default to minimal permission scope for APIs

When using external services, make sure the permissions of that API match the task it is fulfilling. This includes making sure you have separate APIs for read-only and read/write permissions as needed to avoid overprivileged secrets.

Many APIs also allow you to have increased control over what data can be accessed. For example, the Slack API has a large range of scopes, and restricting these scopes to meet the task's minimal requirements of the task is important to prevent an attacker from accessing sensitive data. It is common for inexperienced developers to use API keys with excessive permissions allowing them to use one key throughout their project. But this increases the potential damage of a data breach. Learn about Identity and Access Management best practices here:

IAM Best Practices [cheat sheet included]

Download our cheat sheet on IAM, Identity and Access Management, best practices. It will help you make your cloud environments more secure.

GitGuardian Blog - Automated Secrets DetectionDwayne McDaniel

## Whitelist IP addresses where appropriate

IP whitelisting provides an additional layer of security against bad actors attempting to use APIs nefariously. By providing a whitelist of IP addresses from your private network, your external services will only accept requests from those trusted sources. It is common to include a range of acceptable IP addresses or a network IP address.

In an example from GitHub, you can use IP whitelisting to prevent any untrusted sources from accessing your GitHub repositories. “The allow list for IP addresses will block access via the web, API, and Git from any IP addresses that are not on the allow list.”

Advantages

- Limited requests to select trusted sources and prevents attacks from external sources even with secret keys

Disadvantages

- Not always feasible depending on the traffic the source is expecting
- Can prevent legitimate information requests
- Needs to be maintained constantly

## Use short-lived secrets

It is common for APIs to typically provide long-lasting (or long-lived) access tokens. These tokens could last indefinitely. While this is convenient for developers, it means that a secret poses the same security risks for its entire life and increases the chances of them being used in an attack. By using short-lived secrets, the risk of undetected leaked API keys is mitigated, ensuring that even if an attacker gains access to a secret, it would be harmless, unlike most exposed secrets that remain valid for extended periods.

It is also good practice to make sure you revoke and rotate all API keys often to prevent unrevoked secrets from lurking in your systems, particularly if it is not possible to introduce a validity period on APIs.

> Imagine you own a company with hundreds of employees that all have keys to your office, keys will inevitably get lost, employees will leave the company, new keys will get cut and you will soon lose visibility over where each key is. It would be widely considered good practice to change the locks from time to time.

Advantages

- Enforces good secret hygiene
- Reduces the risk of long-term threats

Disadvantages

- Requires an active secrets management strategy

## Advanced API Key Storage and Cryptographic Protection

While environment variables provide basic separation from source code, enterprise api key management best practices demand more sophisticated storage mechanisms that leverage cryptographic protection layers. Single-purpose keys represent a fundamental principle—use dedicated keys for encryption versus authentication functions to prevent complex cryptographic side effects.

For applications requiring local key storage, implement Hardware Security Modules (HSMs) or secure enclaves that protect keys from extraction even during physical compromise. When HSMs aren't feasible, use secure key generation through cryptographically secure random number generators like Linux's /dev/urandom, wrapped within framework utility functions.

Consider implementing key derivation strategies where multiple API keys derive from a single master key using strong secret diversification methods. This approach reduces the number of root secrets while maintaining functional separation. For distributed systems, evaluate pre-shared keys (PSKs) only as a last resort for disconnected infrastructure or computationally constrained devices, ensuring future-proof algorithms like AES-256 with state-of-the-art operation modes.

## API Key Monitoring and Anomaly Detection

Comprehensive API key security requires continuous monitoring that extends beyond basic usage logging to include behavioral analysis and anomaly detection. Implement monitoring systems that track API key usage patterns, including request frequency, geographic distribution, and accessed resources to establish baseline behaviors.

Deploy alerting mechanisms for suspicious activities such as unusual request volumes, access from unexpected IP ranges, or attempts to access resources outside normal scope. Modern security platforms can correlate API key usage with user behavior analytics to identify potential account takeovers or lateral movement attempts.

Establish audit trails that capture not only successful API calls but also failed authentication attempts, permission escalations, and administrative changes to key configurations. For best practices for managing api keys in enterprise environments, integrate these monitoring capabilities with Security Information and Event Management (SIEM) systems to enable real-time threat detection and automated response workflows. Regular analysis of these audit logs helps identify patterns that may indicate compromised credentials or insider threats before they escalate into significant security incidents.

## Conclusion

Managing secrets and storing secrets is a challenge that requires vigilance from even the most experienced developer, who needs to carefully consider how they are using, storing, sharing, and distributing secrets. Unfortunately, there is no perfect checklist that a developer can follow.

Policies, tools, and strategies will differ from projects, but it is crucial for developers to understand the consequences of policies so that secrets management can be an informed, active strategy throughout the entire development process.

## Summary: Best Practices for API Key and Other Credentials Security

1. Never store unencrypted secrets in .git repositories Avoid git add * commands on git Add sensitive files in .gitignore Don’t rely on code reviews to discover secrets Use automated secrets scanning on repositories
2. Don’t share your secrets unencrypted in messaging systems like Slack
3. Store secrets safely Use encryption to store secrets within .git repositories Use environment variables Use "Secrets as a service" solutions
4. Restrict API keys access and permissions Default to minimal permission scope for APIs Whitelist IP addresses where appropriate Use short-lived secrets

If you want comprehensive guidelines on how to handle secrets with popular DevOps tools such as Docker, Jenkins, Python, AWS, or Kubernetes, have a look at this GitGuardian collection of articles:

How to handle secrets in X

In this series of articles, GitGuardian specialists answer all the questions you never dared to ask on secrets management. Secrets at the Command Line [cheat sheet included]Developers need to prevent credentials from being exposed while working on the command line. Learn how you might be at risk an…

GitGuardian Blog - Automated Secrets DetectionZiad Ghalleb

Finally, if you want to accelerate your learning journey in the world of code security, dive into our collection of cheat sheets: from how to correctly set up multiple GitHub accounts, using GitHub Actions, or manipulate secrets on the command line to understanding IAM best practices, securing infrastructure as code and Docker.

And much more!

## FAQs

What are the most critical API key management best practices for large development teams?

Core best practices include avoiding unencrypted secrets in Git repositories, enforcing automated secrets scanning, and applying least-privilege policies to all keys. Regular key rotation, centralized secrets management, and encryption of all credentials both at rest and in transit are essential to reduce risk and ensure compliance at scale.

How should organizations approach API key rotation and lifecycle management?

Automate rotation schedules based on privilege and risk: rotate high-privilege keys weekly or daily, and rotate low-privilege keys monthly. Use fallback keys only during short transition windows. If any key in a rotation chain is compromised, revoke all related keys immediately. Integrate rotation workflows with secrets management platforms for consistent auditing and governance.

Why is relying solely on code reviews insufficient for detecting hard-coded secrets?

Code reviews typically cover only recent changes and may overlook secrets committed previously and later removed. These secrets persist in repository history and remain exploitable. Automated secrets scanning provides comprehensive coverage across current code, historical commits, and developer workflows, ensuring full visibility and reducing undetected exposures.

What advanced storage options are recommended for API keys beyond environment variables?

Robust alternatives include dedicated secrets managers (e.g., HashiCorp Vault, AWS Secrets Manager), Hardware Security Modules (HSMs), and secure enclaves. These solutions offer cryptographic protection, granular permissioning, automated rotation, and auditability—key requirements for enterprise-grade API key management and regulatory compliance.

How can organizations monitor API key usage and detect anomalies?

Monitor usage metrics such as request volume, originating IP addresses, time-of-day patterns, and resource access behavior. Configure alerts for unusual activity—such as spikes in traffic, unfamiliar geographies, or unexpected permissions usage. Forward logs to SIEM platforms to enable real-time analytics, correlation, and automated incident response.

What are the risks of sharing secrets in messaging platforms like Slack?

Sharing secrets in plaintext on messaging platforms exposes them if user accounts or channels are compromised. Attackers can exploit these credentials for lateral movement or privilege escalation. Always use encrypted channels or dedicated secrets management tools to share sensitive information securely and maintain auditability.
