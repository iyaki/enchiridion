---
title: "Cerbos - Scalable, open-source authorization layer for implementing roles and permissions"
notion_id: 538dfd1a-df26-491d-a113-501441493c05
notion_url: https://app.notion.com/p/Cerbos-Scalable-open-source-authorization-layer-for-implementing-roles-and-permissions-538dfd1adf26491da113501441493c05
last_edited: 2023-04-22T01:16:00.000Z
source_url: https://cerbos.dev/
tags: ["English", "Programming", "Information Security", "Untried", "Tool", "Service"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

# The scalable, open-source authorization layer for implementing roles and permissions

## Do not reinvent

## user permissions

## access management

## authorization

## ABAC

## RBAC

## user permissions

[Get started](https://docs.cerbos.dev/cerbos/latest/quickstart.html)

[How it works](https://cerbos.dev/how-it-works)

Revolutionize the way you authorize with Cerbos

Join leading companies using Cerbos

## Decoupled, low-latency

## and scalable

### Build features, not plumbing

Decouple and centrally manage the authorization logic across all applications and services to reduce repetition, gain visibility and push access changes instantly across the fleet.

### Fine grained access management

Test and deploy fine-grained access control policies with confidence using CI/CD/GitOps workflow.

### Bring your own identity

Cerbos works with any identity provider from popular services like [Auth0](https://cerbos.dev/ecosystem/cerbos-auth0),   [Okta](https://cerbos.dev/ecosystem/cerbos-okta),   [FusionAuth](https://cerbos.dev/ecosystem/cerbos-fusionauth),   [Magic](https://cerbos.dev/ecosystem/cerbos-magic),   [WorkOS](https://cerbos.dev/ecosystem/cerbos-workos)    to your own, bespoke directory system.

```plain text
if (await cerbos.isAllowed({ principal: user, resource, action: "edit" })) {
  // allowed
}
```

```plain text
if (user.email.includes("@mycompany.com") ||
  (user.company.package === "premium" && user.groups.includes("managers"))
) {
  if(user.region === resource.region) {
    // access allowed
    AuditLog.record("ALLOWED", "edit", user, resource);
  } else {
    // access denied
    AuditLog.record("DENIED", "edit", user, resource);
  }
} else {
  // access denied
  AuditLog.record("DENIED", "edit", user, resource);
}
```

Before

After

## Human-readable access

## configuration

### Configure rather than code

Low-code, human-readable configuration that provides wider organizational visibility and enables collaboration for enforcing security policies and auditing compliance requirements.

### Extensible roles and conditions

Go beyond basic role-based-access-control with context-aware role definitions and attributes.

### Multiple use cases

Cerbos policies are flexible enough to model a wide range of domains including multi-tenant SaaS systems, feature flags, product packaging and more.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

```plain text
---
apiVersion: api.cerbos.dev/v1
resourcePolicy:
  resource: contact
  version: default
  rules:

    # admins can do all actions
    - actions: ["*"]
      effect: EFFECT_ALLOW
      roles:
        - admin

    # user role can read and list contacts
    - actions: ["read", "list"]
      effect: EFFECT_ALLOW
      roles:
        - user

    # user role can create and update contacts
    # if they are a member of the sales group
    - actions: ["create", "update"]
      effect: EFFECT_ALLOW
      roles:
        - user
      condition:
        match:
          expr: ("sales" in request.principal.attr.groups)

    # user role can delete a contact if they are one of:
    # - the owner of the contact
    # - member of the sales-manager group
    # - member of the product group
    - actions: ["delete"]
      effect: EFFECT_ALLOW
      roles:
        - user
      condition:
        match:
          any:
            of:
              - expr: request.resource.attr.owner == request.principal.id
              - expr: ("sales-manager" in request.principal.attr.groups)
              - expr: ("product" in request.principal.attr.groups)
```

## Open-core, language-

## agnostic and auditable

### No vendor or cloud lock-in

Cerbos is stateless and self-hosted. Run on any public/private cloud, serverless platform or even your own datacenter. Everything stays within your perimeter and 100% within your control.

### Polyglot

Cerbos exposes a simple, language-agnostic API that can be used from any part of your stack from legacy apps and monoliths to microservices.

### Full auditing and accountability

Audit access controls with real-time change logs and meet ISO27001 and SOC2 requirements. Integrate with security information and event management providers to help avoid advanced persistent threats to security.

### What our users say about Cerbos

"Just discovered your embedded testing framework. This is probably the best balance between hyperfocused functionality and embedded tooling I've ever seen in an open source project. Damn, good work!"

### Rob, Principal Engineer@ [Utility Warehouse](https://uw.co.uk/)

"It's weird to say an outside company has our back, but Cerbos does. It's the people. It's their open-source code: it's high quality, you can read it, it does what it says on the tin"

### Joe, Software Engineer@ [9fin](https://9fin.com/)

"It's a good feeling being able to say yes to almost any permissioning requirement."
      
        "Cerbos is small, contained and easy to implement. It 100% delivers on the promise of abstracting away the complexity of decision making."

### David, Senior Software Engineer@ [Salesroom](https://www.salesroom.com/)

"We're not worried about scaling because we can easily increase our load on Cerbos. It will also be easy for us to change how we're distributing policies as we reach different points of scale."

### Joe, CEO & Co-Founder@ [Nook](https://www.nook.io/)

"We went from one user - every role, to a world where there are many users - many roles. And the product, it relies on Cerbos to actually bring the value that we want to bring to customers. All of our customers are relying on Cerbos, by relying on the product, which is of course relying on Cerbos."

### Chuck, Head of Engineering@ [Salesroom](https://www.salesroom.com/)

"Instead of thinking of how much time Cerbos has saved us, I think about how much time it didn't cost us. It didn't cost us any time. Cerbos just works. I don't have to think about it. It's as simple as that."

### Steve, Staff Engineer@ [NTWRK](https://thentwrk.com/)

"Cerbos just works for us, which is a very welcome predicament."

### Engin, Head of Product and Growth & Co-Founder@ [Debite](https://www.debite.io/)

"If it wasn't for Cerbos, one thing is for sure - we would've launched later than we did. As a result, we would have less customers. And the maintenance part is also very important. Our technical team would be dealing with daily stuff regarding access controls, access logs. Now, we don't have to spend any time on that."

### Rounak, Founding Engineer@ [CommandK](https://www.commandk.dev/)

"Cerbos policy writing is quite flexible, and deploying as a unit microservice as well. Cerbos "doesn't get in the way" once integrated, that's the best part."

### Romina, Tech Lead@ [Wizeline](https://www.wizeline.com/)

"It is easy to implement and provides a solution for a problem that is often not properly addressed."

### Henry, CTO & Co-Founder@ [Nook](https://www.nook.io/)

"Having the separation of the permissions from the code base just makes the code base more elegant. It makes the permissioning more elegant. It means they're centralized, so they're not tied to specific endpoints. And ultimately it means that different business owners have the ability to actually make updates."

### Rasmus, CTO@ [Firtal](https://www.firtal.com/)

"Just discovered your embedded testing framework. This is probably the best balance between hyperfocused functionality and embedded tooling I've ever seen in an open source project. Damn, good work!"

### Rob, Principal Engineer@ [Utility Warehouse](https://uw.co.uk/)

"It's weird to say an outside company has our back, but Cerbos does. It's the people. It's their open-source code: it's high quality, you can read it, it does what it says on the tin"

### Joe, Software Engineer@ [9fin](https://9fin.com/)

"It's a good feeling being able to say yes to almost any permissioning requirement."
      
        "Cerbos is small, contained and easy to implement. It 100% delivers on the promise of abstracting away the complexity of decision making."

### David, Senior Software Engineer@ [Salesroom](https://www.salesroom.com/)

"We're not worried about scaling because we can easily increase our load on Cerbos. It will also be easy for us to change how we're distributing policies as we reach different points of scale."

### Joe, CEO & Co-Founder@ [Nook](https://www.nook.io/)

"We went from one user - every role, to a world where there are many users - many roles. And the product, it relies on Cerbos to actually bring the value that we want to bring to customers. All of our customers are relying on Cerbos, by relying on the product, which is of course relying on Cerbos."

## Bring your own identity

Use any identity provider to authenticate your users. Use Cerbos to enforce access controls.

## SDKs & Integrations

Get productive quickly using our SDKs, quickstart guides and integrations with popular frameworks.

## Deployment Models

Cerbos is stateless and can run anywhere - from bare-metal, to Kubernetes, to serverless, to edge functions.

### These are some questions that may be running through your mind:

### As a Developer

I am tired of the toil and risk of managing access controls.  I need off-the-shelf controls that I don't have to build  myself.

How do I take advantage of modern stacks, serverless architectures and a stateless approach?

How do I manage authorization at scale for SaaS multi-tenant environments?

### As a Product Manager

How do I support complex requirements for enterprise clients, each one having different organizational structures and access control needs?

It also takes far too long to add, retire and tier new features especially as we keep changing our product packaging.

### As a CTO / CISO

I want our team to focus on core application development.

Who tried to or performed what action? How do I check the access logs?

I need to use attribute-based access controls, ABAC, for a zero-trust approach while taking the approval burden off of the IT department.

### And here is how Cerbos can help you:

Cerbos is a self-hosted, open source authorization layer that separates your authorization logic from your core application code.

### Quickly change your authorization logic.

With Cerbos, authorization logic is changed through configuration, not code, making the process faster and accessible to a wider range of stakeholders.

### Reduce the risks from changes to your authorization logic.

Cerbos is very opinionated, reducing the risks that an engineer's error will cause a security incident. And there's no downtime risk, because you're not changing anything in the application's core code. Cerbos also allows you to manage your policies according to Gitops principles.

### Expand the number of stakeholders who can change authorization logic.

With Cerbos, business leaders can see how permissions are structured and make changes to the authorization logic without talking to a developer. There's no risk they'll crash the application, because the authorization layer is decoupled — and they're configuring, not coding.

## Talk to us to find out more!

## Why we built Cerbos

We built Cerbos based on our experience at Google, Microsoft, Elastic, Qubit and CGI, because we have experienced first hand the problems with creating, maintaining and scaling permissions management.

## Recent Articles

# [How to use Cerbos effectively](https://cerbos.dev/blog/how-to-use-cerbos-effectively)

[Charith Ellawala on Apr Tue, 04DOCUMENTATIONGUIDE](https://cerbos.dev/blog/how-to-use-cerbos-effectively)

# [How Loop achieved reliable and scalable authorization with Cerbos](https://cerbos.dev/blog/how-loop-achieved-reliable-and-scalable-authorization-with-cerbos)

# [Contributor podcast: Decoupling authorization](https://cerbos.dev/blog/contributor-podcast-decoupling-authorization)

## Let's keep in touch!

Please subscribe below to get notified about
all the new features and updates from Cerbos.
