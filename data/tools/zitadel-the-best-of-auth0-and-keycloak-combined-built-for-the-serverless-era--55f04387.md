---
title: "ZITADEL - The best of Auth0 and Keycloak combined. Built for the serverless era"
notion_id: 55f04387-bc1b-4bcc-bd45-015d517744ec
notion_url: https://app.notion.com/p/ZITADEL-The-best-of-Auth0-and-Keycloak-combined-Built-for-the-serverless-era-55f04387bc1b4bccbd45015d517744ec
last_edited: 2023-02-01T16:41:00.000Z
source_url: https://zitadel.com/
tags: ["Tool", "Framework/Library", "English", "Information Security", "System Design / Software Architecture", "Untried"]
---
ZITADEL gives developers all they need to integrate identity management. Easy as pie. Ready when you are — because serverless. At yours or ours — because open source.

What we do

## Identity management that works for you

You want auth that's quickly set up like Auth0 but open source like Keycloak? Look no further — ZITADEL combines the ease of Auth0 and the versatility of Keycloak.

It is built for devs. And for the serverless age. Featuring multi-tenancy, open source and the cloud. You can get started in minutes on our public cloud service, run it locally or self-host. But mostly: Take care of your business, because your login is taken care of.

What we offer

## OpenSaaS: The best of two worlds

Open source and SaaS got together to bring you simplicity while enabling you to dive as deep as you like. The serverless architecture enables you to jump right in and scale up as needed. If you'd rather run your own instance on your computer, that's just one command away.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Open Source

Run your own instance locally. Or wherever you like. And for the ones who think a repo says more than words: No need to ask for keys.

[Learn more](https://zitadel.com/opensource)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Use our public cloud service (our free plan includes all features) or self-host for full control — in any case, we're happy to lend a hand.

[Learn more](https://zitadel.com/saas)

Why the best of both worlds?

## B2B: Bring your clients

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### Multi-Tenancy

B2B is baked into ZITADEL from the very beginning. Multi-tenancy enables you to provide auth for all of your clients, be it companies or end users.

[Learn more](https://zitadel.com/features#delegation)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Improve your clients' and partners' experience by allowing them to manage their resources without your interaction.

[Learn more](https://zitadel.com/usecases#self-service)

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

[Learn more](https://zitadel.com/features#brokering)

## Integrate with all your systems

Use our APIs, SDKs and our quickstarts to enable modern authentication with any of your systems.

### API-first

[Learn more](https://zitadel.com/docs/apis/introduction)

Execute custom code triggered by events to customize behavior and workflows.

[Learn more](https://zitadel.com/docs/concepts/features/actions)

### SDKs

Use our client libraries for interacting with ZITADEL or any standard OpenID Connect library.

[Learn more](https://github.com/zitadel/zitadel#client-libraries)

We provide quickstart guides in multiple languages and tools for you to start developing.

[Learn more](https://zitadel.com/docs/examples/introduction)

```plain text
provider, err := oidc.NewProvider(ctx,
"https:/[your-domain].zitadel.cloud")
if err != nil {
    // handle error
}

oauth2Config := oauth2.Config{
    ClientID:     clientID,
    ClientSecret: clientSecret,
    RedirectURL:  redirectURL,
    Endpoint: provider.Endpoint(),
    Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
}

```
