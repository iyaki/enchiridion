---
title: "Software Security Primer"
notion_id: ce6fb890-fbd1-4137-b0b3-5d77d5f2fa2d
notion_url: https://app.notion.com/p/Software-Security-Primer-ce6fb890fbd14137b0b35d77d5f2fa2d
last_edited: 2023-04-25T14:25:00.000Z
source_url: https://hackernoon.com/software-security-primer-sku3wl9
tags: ["Article", "Hackernoon", "English", "Programming", "Information Security"]
---
[https://hackernoon.com/software-security-primer-sku3wl9](https://hackernoon.com/software-security-primer-sku3wl9)

As a developer, when you think of security, what comes to mind? Is it clear what are different aspects that need to be handled to make a software secure? And why you need to do that?

The objective of this article is to provide that view and articulate what controls should be in place and how. Details of how to apply those controls are not covered as they would require separate articles - and lot of content is available on the web anyway.

## What is Software Security?

Security implementation of a software application can be classified in two parts:

1. Pre-deployment - building a secure software
2. Post-deployment - security of the environment where the software is running

Software Security is pre-deployment. It is the process of identifying risks and building controls (or Countermeasures as it is called in security terminology) in the software itself while it is being built.

Software Security is the focus of this post. We will see controls commonly used and what risks they mitigate.

For more details, see What is Software Security by Gary McGraw

## Nature of security

A software application would generally have two aspects:

a. Services that provide some functionality

b. Data generated and consumed by the services

Security can be defined as defending the services and the data from unauthorized and malicious usage at all times.

Defending is the key word here. Defense literally means an act of resisting an attack. That means attacks can happen anytime and any number of times, and we need to keep protecting the system from these attacks.

That's what makes security of a software application very difficult because it is not easy to get it right all the time.

Also, there is no room for error. One incident may be enough to destroy the reputation and business built over the years.

## How to approach security?

The following is the minimum that should be done:

Follow industry best practices: Rather than inventing our own mechanisms, follow the best practices. These best practices have come out of years of learning and it is better to rely on that tried and tested knowledge.

Secure by design: Make security a criteria in all aspects like Requirements, Architecture, DevOps, Engineering Practices like code reviews, unit testing, application testing, etc. This would ensure that everything is built secure by default - security is not thought of after developing the functionality.

## What would make a software secure?

CIA triad (or AIC triad) is commonly used for Information Security, but it also helps to better understand security in general. A software should have the following properties to be secure:

Confidentiality: Authorized access to services and data e.g., to view your GMail mailbox and send emails, you must successfully login to GMail.

Integrity: Integrity and authenticity of the data presented (security in transit) and stored (security at rest) e.g., the emails shown by GMail are exactly how they were sent by the sender and the content is same as that stored on the GMail servers.

Availability: Software application is always available to all authorized users for all legitimate use-cases e.g., GMail services are always available to its users.

## Design Principles

See Secure Design Principles by OWASP for details.

Defense in depth is a very important principle to follow while designing the software. In simple terms, it means to have controls at different layers. It will ensure breach at one layer will not result in total failure. We'll see some examples later in the article.

## Controls

The following are the controls that are commonly used. A combination of these controls can be implemented depending on the risk.

NOTE: These are commonly used controls. This is not an exhaustive list e.g., though it is not included above, applying signatures on every request and response can be used if the risk is high and require that level of security.

We will look at each of these. However, they build on the security building blocks like authentication, cryptography, certificates, etc. If you are not familiar with these concepts, it may be helpful to go through Software Security Building Blocks.

## Authentication

Why?

Authentication is required to perform authorization. Authentication will establish the identity of the entity trying to access the services, and based on that authorization can be applied.

How?

There can be two types of authentication:

Authenticating an end user:

OAuth2 and OpenID Connect are widely used protocols for this. There are number of identity management services that support these protocols.

Using a standard protocol rather than custom authentication mechanism would make integrations with other systems easier.

Authenticating a client:

Two commonly used methods to authenticate a Client (a program calling an API):

OAuth2 - Client Credentials Grant: This is an OAuth2 workflow that allows a client to authenticate itself using a client-id and a client-secret and acquire a token which it can then use to call APIs.

TLS Mutual Authentication: Both the Client and the Server trust each other's X.509 certificates used to establish TLS connection. Connection is established only if trusted parties try to connect. This is suitable in server-to-server communications e.g., when a backend service calls an external API.

For Web Applications that rely on cookie based session tracking, authenticating legitimate requests from client (Browser) is required to avoid Cross Site Request Forgery (CSRF). Watch this video to see a demo of CSRF.

## Authorization - for Confidentiality

Why?

Authorization in this context is required for the following objective:

Authenticated entity - a user or a client - can access only the data that they are supposed to access e.g., when you login to GMail, GMail does not show someone else's emails to you.

How?

All read operations on data associated with a user should have the following checks.

1. The logged-in user is same as the user associated with data e.g., User Id on a payment transaction record must match the User Id of logged in user. The same check can be applied to clients if required.
2. If #1 is not true, the logged-in user has enough permissions to access other user's data.

## Authorization - for Integrity

Why?

Authorization in this context is required for the following objective:

Authenticated entity - a user or a client - can perform only those data manipulation actions that they are authorized to e.g., when you login to Outlook, you cannot manage someone else's calendar unless the other person has given you permissions to manage his or her calendar.

How?

All operations should have the following checks:

1. Is the logged-in user authorized to perform the operation?
2. If Role Based Access Control (RBAC) is used, is a role required to perform the operation assigned to the logged-in user? Spring Security and .Net Framework easily allow such access control using declarative methods - annotations \ attributes.
3. The logged-in user is same as the user associated with data that is being modified e.g., the User Id on a payment transaction request must match the User Id of the logged in user.

Client Authorization:

In the case of APIs that can be consumed by multiple clients, authorizing clients is also important e.g., in a Microservice architecture, an API gateway can block requests if a client is not authorized to use an API.

In the above design, an admin can disable a user so it is allowed from the Admin Portal client application only. If the Web Portal application tries to call the API, it is blocked. OAuth2 JWT tokens carry Client information and this can be done easily by inspecting tokens.

NOTE: The same check should also be applied on the Microservice - remember Defense in depth? In case someone mistakenly removes the check on API Gateway, microservice will still block the client.

## Data Privacy

Why?

Personally Identifiable Information (PII) is any data that can identify an individual e.g., Mobile Number, Email address, etc. Unauthorized access to such information can be a breach of privacy.

How?

Identify PIIs captured by the software.

Do not display PIIs without masking e.g., mobile number as 98XXX XXX76.

All communications like email, SMS, etc. should mask PIIs.

Do not write PIIs to the log files in plain text.

## Input Validation

Why?

Accepting invalid input can compromise the integrity of the data e.g., accepting a past date on a payment transaction request would result in incorrect state of the system.

Lack of input validation can pose a vulnerability for Cross Site Scripting (XSS) attacks. See CROSS-SITE SCRIPTING (XSS) TUTORIAL for a quick introduction to XSS.

How?

Inputs should be validated at all layers: Defense in depth!

Sanitize all text input to remove <script> tags. Otherwise the <script> tags would be stored with text and it might end-up in a malicious script being executed when the text is displayed in the browser later on - that's XSS in a nutshell.

If a file is being uploaded as part of some functionality, it should also be treated as an input and be validated - scanned and its format validated if required.

## TLS \ HTTPS

Why?

For security of data when it is in transit. It protects against Man-in-the-middle attack which can compromise both integrity and confidentiality of the system.

How?

Always use TLS 1.2 (currently latest) and HTTPS for all internet traffic.

In mobile applications, implement SSL Pinning to make sure the mobile app always talks to "your" server only. See Android Security: SSL Pinning for details.

## Encryption and Hashing

Why?

Security of data at rest (when stored) is as important as security in transit. Storing encrypted or hashed data will prevent unauthorized access to data at rest.

How?

Identify data that needs to be stored encrypted or as a hash value. Hash of a password is stored so it is irreversible even with the keys.

Encryption key and the data should not be stored together e.g., in same database.

Challenges:

1. Key Protection: Encrypted data is as secure as the key it is encrypted with. Therefore, management of encryption key is very important when encrypting data at rest.
2. Performance trade-off: Having to decrypt data on all read operations will impact the overall responsiveness of the system if frequently used data fields are stored encrypted.
3. Key rotation: Once data is stored with an encryption key, it has to be decrypted with the same key. When data starts accumulating over the years, rotating the encryption key becomes a challenge because existing data needs to either (a) be re-encrypted with new key, which may be challenging in a live system or (b) handle multiple keys and key-data mapping in program itself, which can become very complicated.

Cloud services like AWS Key Management Service (KMS) provide key protection and key rotation out of the box. Access to KMS keys and its encryption APIs can be controlled using IAM permissions and the client can encrypt / decrypt data without accessing the key itself.

## Auditing and Logging

Why?

Auditing helps to track the change in state and who caused that change e.g., if a user's name was updated, audit records will capture when it was updated and who updated it. This can be useful in event of an incident.

Logging helps to trace errors and troubleshoot issues.

How?

Auditing is straightforward. Hibernate makes it very easy in Java.

Logging too is straightforward, but in distributed architecture using ELK Stack or something similar is important to see all logs in one place.

## Secrets Management

Why?

Secrets are database password, API access credentials, etc. Most of these are configurations but if they are left in plain text, it poses a risk as they can be misused.

Let's say an organization uses an external SMS service provider to send SMS. If the API access credentials for this external service are not protected, there is a risk of someone being able to send SMS with the identity of the organization.

How?

You can use AWS Secrets Manager, Vault, etc. to store such secrets. These services can also automatically rotate secrets.

Alternative is to keep secrets encrypted in normal configuration - properties file or web.config, and the encryption key is stored in AWS Secrets Manager, Vault, etc.

In Single Page Applications (SPA) and Mobile Applications, generally a JWT token would be used to call APIs: token is passed in Authorization header with every request. Since it is a bearer token, anyone with the token can use the APIs. So the token should not be stored on clients more than required .

NOTE: There is trade-off between security and usability. Most applications (web and mobile) keep users logged in forever, which means they have to keep a token with long validity on the client: in browser local storage or mobile. This is a risk. So a decision has to be made based on nature of the application.

## Sender Policy Framework

Why?

Let's say an application sends emails from connect@myapp.co address. If someone else is able to send emails using the same domain myapp.co, it can impact reputation and business of the organization. This is called Spoofing. Sender Policy Framework prevents such misuse of domain and assures authenticity of emails.

How?

SPF is configured with a DNS entry. See What is SPF for details.

## Resiliency

Why?

Resiliency here means ability of the application to be responsive (operational) in the event of a failure. See Reactive Manifesto for more details.

In the context of security, we can loosely define resiliency to the following:

1. In the event of a sudden increase in load (with malicious intent i.e. Denial of Service (DoS) attack), can the system remain operational? All cloud providers have services that provide protection against DoS \ DDoS attacks, but these are reactive in nature and may take a little time to react. The application should be able to withstand that otherwise the protection is of no use.
2. How quickly the application can recover in the event of a downtime?

How?

There are different modern architecture patterns to achieve resiliency, but in any case - even for monolith architecture, the system can be resilient with following implementation:

Elasticity: Elasticity means ability to scale up or down depending on the load. It is easy to implement with Load Balancing and Auto Scaling in any cloud environment.

Automation: Automated builds and deployments is critical to quickly recover in the case of a failure. Unit testing and Test Automation would also help to do quick releases for urgent patches.

Configuration Simplicity: Configurations keep getting added to a software as it evolves and often it become very difficult to correctly configure the software. This is important to have quick recovery in case of a failure. See Application Configuration from Operations Viewpoint for impact of configuration on operations and ways to mitigate the challenges.

## Conclusion

Security is difficult if it is ignored, but Identifying risks and putting in place appropriate controls by following industry best practices and secure by design principles can make it a less daunting task.

Also published at https://dev.to/pathiknd/software-security-overview-3ldi
