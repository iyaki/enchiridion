---
title: "Thinking beyond SQL injection: OWASP tips for secure database access"
notion_id: 0716ddcb-f2ef-4c35-b060-84f734458824
notion_url: https://app.notion.com/p/Thinking-beyond-SQL-injection-OWASP-tips-for-secure-database-access-0716ddcbf2ef4c35b06084f734458824
last_edited: 2026-09-21T17:35:00.000Z
source_url: https://github.blog/security/supply-chain-security/beyond-sql-injection-owasp-tips-secure-database-access/
tags: ["Programming", "Databases", "Information Security", "Article", "Github Blog", "English"]
---
[https://github.blog/2022-01-27-beyond-sql-injection-owasp-tips-secure-database-access/](https://github.blog/2022-01-27-beyond-sql-injection-owasp-tips-secure-database-access/)

This is part three of GitHub Security Lab’s series on the OWASP Top 10 Proactive Controls, where I provide practical guidance for OSS developers and maintainers on improving your security posture.

When you think about database security, the first thing that might come to mind is SQL injection. In 2022, SQL injection is a very well-known security vulnerability, as seen through projects such as the OWASP Top 10 risks or even XKCD’s now-famous “little Bobby Tables” cartoon. Yet as you’ll see in this post, there’s more to consider when it comes to secure database access. OWASP Top 10 Proactive Control C3 (secure database access) is especially complete and verbose. The first thing it makes clear is that database security issues are not just a problem of relational databases. NoSQL databases have their issues too and should not be considered more secure.

OWASP breaks database access security down into the following areas:

1. Secure queries
2. Secure configuration
3. Secure authentication
4. Secure communication

I’m adding a fifth security category. Even though it may not be the source of most database-related vulnerabilities, it’s the source of enough problems to deserve its own mention:

5. Secure connection

As expected, secure queries, which relates to SQL injection, is the top item. So, let’s start with that.

## Secure queries

Database query injection is one of the oldest and best known vulnerability classes. It somehow refuses to disappear and brings us new instances every year. The gist of how to query a database in a secure way is straightforward:

never let user-controlled data change the meaning of any database query.

That includes not only SQL statements but also NoSQL, OQL, GraphQL, stored procedures, and any other such user-queryable database. Note that there is a big difference between allowing user-controlled data to be part of a query and letting that data actually change the meaning of the query.

For example, let’s say I directly concatenate or interpolate user-controlled data into a query template:

```plain text
String query = "Select * from USERS where name = '" + request.getParam("username") + "' and password = '" + request.getParam("password") + "'" ;
```

Attackers will be able to change the meaning of the query by breaking out of the query context where their data is used (single-quoted string literal). In this case, the username and password are user-provided parameters that may change the meaning of the query itself.

For example, providing a value such as:

```plain text
Username: "admin" Password: "FOO' or '1'='1"
```

Will turn our templated query into:

```plain text
String query = "Select * from USERS where name = 'admin' and password = 'FOO' or '1' = '1'";
```

The resulting query will return all users and, if it is used for authentication purposes, will likely let users log in as an administrator. Alternatively, an attacker could have dropped the whole database, read arbitrary records or, in some cases, executed system commands or written files to the underlying file system.

Depending on your background, you might already be familiar with the workings of SQL injection, but what is the best way to prevent it? The best mitigation is actually not just validation of user-controlled data (for example, making sure data does not contain single quotes) or sanitization of user-controlled data (for example, removing single quotes). While these approaches are core to the input validation security in depth layer, which I’ll address in Control 5, they are also error-prone and may be insufficient. For example, some SQL injections may not require escaping out of a single-quoted context, and some sanitizers may not address all dangerous characters or might replace just their first occurrence.

The best mitigation is to use what is known as query parameterization.

Query parametrization is a feature offered by most database libraries that clearly separates the query template from the query parameters in such a way that parameters will be securely used within the template by applying the right encodings. (This topic will come up again later, in OWASP Proactive Control 4). For Java, this looks like:

```plain text
String custname = request.getParameter("customerName"); String query = "SELECT account_balance FROM user_data WHERE user_name = ? "; PreparedStatement pstmt = connection.prepareStatement( query ); pstmt.setString( 1, custname); ResultSet results = pstmt.executeQuery( );
```

If you want to learn more about query parameterization or how specific libraries implement it, OWASP has a wonderful cheat sheet for it. I recommend bookmarking this page. It’s full of resources to help developers write more secure code.

However, as stated in the OWASP C3 description, there are some parts of SQL queries that cannot be parameterized. In those cases, you may find yourself having to manually craft parts of the query using string concatenation. If you find yourself here, you may prefer more restrictive validation (for example, typing in addition to mapping of predefined values, abstracted, or indirect values). The point is to be extremely careful and make sure to properly validate the user-controlled data so that they cannot break out of the intended query context and escape into other parts of the query.

## Secure configuration

Even though SQL injections are still the top vulnerability when it comes to database security, other misconfigurations can enable additional types of attacks or escalate SQL injections to their worst form by enabling remote command execution.

This is possible because, unfortunately, database management system (DBMS) configurations are not always hardened by default. Dangerous configuration properties and their insecure defaults vary from DBMS to DBMS but, again, OWASP helps us out by providing another excellent cheat sheet that lists these insecure defaults and how to harden them. In this list, you can find recommendations, such as:

- Disabling command execution features when not needed
- Disabling stored procedures when not needed
- Disabling insecure authentication modes
- Remove sample databases installed by default
- Disabling browser services exposed by default
- Disabling file system access features if not needed
- Changing default ports
- Disabling default insecure (unencrypted) transport protocols
- Disabling unauthenticated access

## Secure authentication

I already mentioned that default database configurations might not require any form of authentication, but I think authentication is important enough to warrant its own section in a list of DBMS-related vulnerabilities. As a rule of thumb, any access to a database should be properly authenticated. This will not just prevent unauthenticated users with local access to your database ports from messing with your data, it will also provide an account of who accessed your database and which actions were performed by them. Authentication to the DBMS should be accomplished in a secure manner. This includes making sure authentication takes place only over secure channels and that credentials are properly secured.

Guess what? OWASP has another cheat sheet to help you make sure the authentication to your databases is hardened and secure. A few highlights:

- Set strong and secure passwords.
- Use per application/service accounts.
- Apply the principle of least privilege, and grant as few privileges to the account as needed.
- Perform regular account audits to verify that the accounts are still needed, have the minimum privileges, and that their passwords, keys, or tokens are rotated.

## Secure communication

A DBMS is as secure as the services, APIs, and transport methods used to reach it. When several communications options exist, you should only use those that involve encrypted communications against authenticated endpoints. OWASP has yet another cheat sheet to help you secure the communications to your databases:

- Isolate the backend database as much as possible. Place it on a separate DMZ isolated from the application servers accessing the database.
- Disable network access when possible.
- Bind services to local ports when possible.
- Limit access to service ports to specific hosts that need to access the database.
- Configure the database to only allow encrypted connections.

## Secure connection

In addition to the four items mentioned above and listed in the OWASP C3 control, I would like to bring your attention to a fifth element. In some cases, either through configuration files or administrative panels, it is possible for high-privileged users to control the DBMS connection string, either entirely or partly. Depending on the DBMS and language in use, controlling the connection string can cause a variety of problems, such as connection string parameter pollution, deserialization attacks, or JNDI datasource injection attacks. To prevent them, apply the same concepts we discussed above. For example:

- Validate that only expected characters are used.
- Encode the input for the right connection string context (host, authority, parameters, etc.).
- When possible, use a level of indirection so the final values are always retrieved from a predefined set of good known values.

## It’s a wrap

Paying attention to the five areas I reviewed in this post will help you reduce the risk of database security issues and write more secure code.

Next time, we will review the importance of encoding and escaping the data for the contexts in which the data will be used. Until then, stay secure! Oh, and don’t forget to review and bookmark the OWASP cheat sheets!

Follow GitHub Security Lab on Twitter for the latest in security research.

## Written by

## Explore more from GitHub

### Docs

Everything you need to master GitHub, all in one place.

Go to Docs

### GitHub

Build what’s next on GitHub, the place for anyone from anywhere to build anything.

Start building

### Customer stories

Meet the companies and engineering teams that build with GitHub.

Learn more

### GitHub Universe 2026

Join us October 28-29 in San Francisco or online for GitHub Universe, our flagship developer event uniting people, agents, and the world’s code.

Register now
