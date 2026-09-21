---
title: "Turso - SQLite for Production"
notion_id: 000d55f1-2bfe-4407-92c7-a6d6d66b0972
notion_url: https://app.notion.com/p/Turso-SQLite-for-Production-000d55f12bfe440792c7a6d6d66b0972
last_edited: 2024-06-05T18:35:00.000Z
source_url: https://turso.tech/
tags: ["English", "Databases", "Untried", "Service"]
---
“In e-commerce, proximity to users is vital. Turso lets us minimize round trip network latency for our global userbase and it makes a huge difference.“

Jokull Solberg, CTO

“Turso enables us to efficiently scale Astro Studio’s database per tenant architecture to as many users as we’ll ever need, on demand. It’s a game changer.“

Fred K. Schott, Co-creator

“We use Turso because it combines the efficiency of SQLite with reliability and scalability additions required by mission-critical applications.“

Rafael Umann, CEO

![image](https://turso.tech/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fbackground.1df0b096.png&w=3840&q=75)

## Best-in-class developer experience

Build with and integrate SQLite into your production applications. Whether you’re building with a single database or millions with a per-tenant architecture for mobile, web or desktop, Turso has all the developer tools you need to integrate in seconds.

[Start Building](https://api.turso.tech/auth/clerk?webui=true&type=signup)[Quickstart](https://docs.turso.tech/introduction)

```plain text

import { createClient } from "@libsql/client";

const client = createClient({
  url: "file:replica.db",
  syncUrl: "libsql://...",
  authToken: "...",
});

const result = await client.execute({
  sql: "SELECT * FROM users WHERE id = ?",
  args: [1],
});
```

### CLI

Take control of your entire database infrastructure via the command-line using the Turso CLI.

[Get Started ->](https://docs.turso.tech/cli/introduction)

### Platform API

Create databases and manage replication programmatically using the Platform API.

[Browse docs ->](https://docs.turso.tech/api-reference/introduction)

### Integrate easily

Integrate SQLite using any language or framework with open-source SDKs and libSQL protocols.

[Install ->](https://docs.turso.tech/sdk/introduction)

### Examples

Learn from ready to go apps and code examples.

[Browse examples ->](https://github.com/tursodatabase/examples)

### Scalable

Per-tenant database architecture made easy.

Offer a database per user/region/cluster/etc. Scale to millions.

- 500 Free - Starter Plan
- 10,000 - Scaler Plan
- 25,000 - Pro Plan
- Enterprise Plan

![image](https://turso.tech/_next/image?url=%2F_next%2Fstatic%2Fmedia%2FScalableImage.53076c27.png&w=3840&q=75)

### Simple

Start local, deploy when ready. Replication is easy as copy/paste to anywhere or inside your own infrastructure.

![image](https://turso.tech/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fsimple.c2e80587.png&w=384&q=75)

### Fast

Zero network latency with embedded replicas.

34+ global locations for low read & write latency.

![image](https://turso.tech/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Ffast.9a45e10b.png&w=828&q=75)

### API First

Turso comes with a powerful Platform API to build with, so you can create databases, manage replication and more.

### Open Source

Turso is powered by libSQL, the open contribution fork of SQLite, so your data is portable.

![image](https://turso.tech/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Fgithub.ebc218df.png&w=384&q=75)

### Secure and compliant

Turso is built with security for production deployments in mind, with encryption at rest and encryption in transit.

SOC2

HIPAA

- GDPR Data Placement Compatible
- DPA Available

## Turso is made for...

... and more

### Talk to us

Whether you're already using the product or thinking of trying us out, we'd love to hear from you.

![image](https://turso.tech/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Ftalk-to-us-background.59f7e65f.png&w=3840&q=75)

Talk to us
