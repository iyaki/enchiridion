---
title: "Open Policy Agent"
notion_id: 89f537de-2331-4108-a720-a56b3d409ee9
notion_url: https://app.notion.com/p/Open-Policy-Agent-89f537de23314108a720a56b3d409ee9
last_edited: 2026-09-21T17:12:00.000Z
source_url: https://www.openpolicyagent.org/
tags: ["English", "Information Security", "DevOps", "Untried", "Tool"]
---
## OPA is a policy engine that streamlines policy management across your stack for improved development, security and audit capability.

data.json

```plain text
{}
```

input.json

```plain text
{ "account": { "state": "open" }, "user": { "risk_score": "low" }, "transaction": { "amount": 950 }}
```

```plain text
# Run your first Rego policy!package paymentsdefault allow := falseallow if { input.account.state == "open" input.user.risk_score in ["low", "medium"] input.transaction.amount <= 1000}# Open in the Rego Playground to see the full example.
```

Loading...

Developer Productivity: OPA helps teams focus on delivering business value by decoupling policy from application logic. Security & platform teams centrally manage shared policies, while developer teams extend them as needed within the policy system.

Performance: Rego, our domain-specific policy language, is built for speed. By operating on pre-loaded, in-memory data, OPA acts as a fast policy decision point for your applications.

Audit & Compliance: OPA generates comprehensive audit trails for every policy decision. This detailed history supports auditing and compliance efforts and enables decisions to be replayed for analysis or debugging.

Interested to see more? Checkout the Maintainer Track Session from KubeCon.

## Context-aware, Expressive, Fast, Portable

OPA is a general-purpose policy engine that unifies policy enforcement across the stack. OPA provides a high-level declarative language that lets you specify policy for a wide range of use cases. You can use OPA to enforce policies in applications, proxies, Kubernetes, CI/CD pipelines, API gateways, and more.

The examples below are interactive! Use the Evaluate button to see output for the given rego and input. Then edit the rego or the input (or both) to see how the output changes.

- API
- Envoy
- Kubernetes
- AI Tool Calling

Applications can directly integrate with OPA using the SDKs or REST API. This is great when your application needs to make domain specific runtime decisions.

policy.rego

```plain text
package application.authz# Only owner can update the pet's information. Ownership# information is provided as part of the request data from# the application.default allow := falseallow if { input.method == "PUT" some petid input.path = ["pets", petid] input.user == input.owner}
```

Output

```plain text
{ "allow": false }
```

Loading...

input.json

```plain text
{ "role": "staff", "owner": "bob@example.com", "path": [ "pets", "pet113-987" ], "user": "alice@example.com"}
```

data.json

```plain text
{}
```

Open in OPA Playground
