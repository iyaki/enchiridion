---
title: "OpenAPI.Tools - An Open Source list of great tools for Open API"
notion_id: a8f7c5ce-9403-404e-abef-c62340b69016
notion_url: https://app.notion.com/p/OpenAPI-Tools-An-Open-Source-list-of-great-tools-for-Open-API-a8f7c5ce9403404eabefc62340b69016
last_edited: 2023-05-03T19:04:00.000Z
source_url: https://openapi.tools/
tags: ["English", "Programming", "REST API", "Tool", "Service", "Website"]
---
## Tool Types

We've organised everything into categories so you can jump to the section you're interested in.

- [**Auto Generators:**](https://openapi.tools/#auto-generators) Tools that will take your code and turn it into an OpenAPI Specification document
- [**Converters:**](https://openapi.tools/#converters) Various tools to convert to and from OpenAPI and other API description formats.
- [**Data Validators:**](https://openapi.tools/#data-validators) Check to see if API requests and responses are lining up with the API description.
- [**Description Validators:**](https://openapi.tools/#description-validators) Check your API description to see if it is valid OpenAPI.
- [**Documentation:**](https://openapi.tools/#documentation) Render API Description as HTML (or maybe a PDF) so slightly less technical people can figure out how to work with the API.
- [**DSL:**](https://openapi.tools/#dsl) Writing YAML by hand is no fun, and maybe you don't want a GUI, so use a Domain Specific Language to write OpenAPI in your language of choice.
- [**Gateways:**](https://openapi.tools/#gateway) API Gateways and related tools that have integrated support for OpenAPI.
- [**GUI Editors:**](https://openapi.tools/#gui-editors) Visual editors help you design APIs without needing to memorize the entire OpenAPI specification.
- [**Learning:**](https://openapi.tools/#learning) Whether you're trying to get documentation for a third party API based on traffic, or are trying to switch to design-first at an organization with no OpenAPI at all, learning can help you move your API spec forward and keep it up to date.
- [**Miscellaneous:**](https://openapi.tools/#miscellaneous) Anything else that does stuff with OpenAPI but hasn't quite got enough to warrant its own category.
- [**Mock Servers:**](https://openapi.tools/#mock) Fake servers that take description document as input, then route incoming HTTP requests to example responses or dynamically generates examples.
- [**Parsers:**](https://openapi.tools/#parsers) Loads and read OpenAPI descriptions, so you can work with them programmatically.
- [**SDK Generators:**](https://openapi.tools/#sdk) Generate code to give to consumers, to help them avoid interacting at a HTTP level.
- [**Security:**](https://openapi.tools/#security) By poking around your OpenAPI description, some tools can look out for attack vectors you might not have noticed.
- [**Server Implementations:**](https://openapi.tools/#server) Easily create and implement resources and routes for your APIs.
- [**Testing:**](https://openapi.tools/#testing) Quickly execute API requests and validate responses on the fly through command line or GUI interfaces.
- [**Text Editors:**](https://openapi.tools/#text-editors) Text editors give you visual feedback whilst you write OpenAPI, so you can see what docs might look like.

## [Auto Generators](https://openapi.tools/#auto-generators)

Tools that will take your code and turn it into an OpenAPI Specification document

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ har-to-openapi ](https://github.com/jonluca/har-to-openapi) - Automatically generate OpenAPI 3.0 Spec by using network requests captured in one or more HAR files | TypeScript | ❌ | ✅ | ❌ |  |
| [ har2openapi ](https://github.com/dcarr178/har2openapi) - Automatically generate OpenAPI 3.0 Spec by using network requests captured in one or more HAR files | TypeScript | ❌ | ✅ | ❌ |  |

## [Converters](https://openapi.tools/#converters)

Various tools to convert to and from OpenAPI and other API description formats.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ Apimatic Transformer ](https://apimatic.io/transformer) - Transform API Descriptions to and from RAML, API Blueprint, OAI v2/v3, WSDL, etc. | SaaS | ✅ | ✅ | ✅ |  |
| [ avantation ](https://www.avantation.in/) - Generate OpenAPI 3.x specification from HAR. | TypeScript | ❌ | ✅ | ❌ |  |
| [ Counterfact ](https://counterfact.dev/) - A new (July 2022) project that converts an OpenAPI document to a full implementation that runs on ts-node. The idea is to replace the auto-generated code that returns random values with realistic code, one path at a time. If the spec is updated, you can regenerate the types and let the type checker show you which parts of the implementation need to be updated. The types can also be used in client-side code. | TypeScript / Node | ✅ | ✅ | ❌ |  |
| [ go-swagger ](https://goswagger.io/) - Unmaintained v2.0 only project seeking new maintainer, or probably a fork. Parser, validator, generates descriptions from code, or code from descriptions! | Go | ❌ | ❌ | ✅ |  |
| [ Google Gnostic ](https://github.com/googleapis/gnostic) - Compile OpenAPI descriptions into equivalent Protocol Buffer representations | Go | ❌ | ✅ | ✅ |  |
| [ JSON Schema to OpenAPI Schema ](https://www.npmjs.com/package/@openapi-contrib/json-schema-to-openapi-schema) - Due to the OpenAPI v3.0 and JSON Schema discrepancy, you can use this JS library to convert JSON Schema objects to OpenAPI Schema. | TypeScript | ❌ | ✅ | ❌ |  |
| [ Karate-IDE ](https://marketplace.visualstudio.com/items?itemName=KarateIDE.karate-ide) - Generates KarateDSL Tests and Mocks from OpenAPI 3.0 documents and so you can quickly test/explore your API. | VSCode Extension | ❌ | ✅ | ❌ |  |
| [ laravel-openapi ](https://github.com/vyuldashev/laravel-openapi) - Generate OpenAPI 3 specification for Laravel Applications. | PHP | ❌ | ✅ | ❌ |  |
| [ LucyBot api-spec-converter ](https://www.npmjs.com/package/api-spec-converter) - Convert between API description formats such as OpenAPI and RAML. | Node.js | ❌ | ✅ | ✅ |  |
| [ OAS RAML Converter ](https://mulesoft.github.io/oas-raml-converter/) - Converts between OpenAPI and RAML API specifications | Node.js | ❌ | ✅ | ✅ |  |
| [ OData OpenAPI ](https://github.com/oasis-tcs/odata-openapi) - OData 4.0, 3.0, and 2.0 to OpenAPI v3.1, v3.0, and v2.0 converter | Node.js / XSLT | ✅ | ✅ | ✅ |  |
| [ OData.OpenAPI ](https://github.com/xuzhg/OData.OpenAPI) - Convert an Edm (Entity Data Model) to OpenAPI 3.0 | .NET | ❌ | ✅ | ❌ |  |
| [ OpenAPI Filter ](https://github.com/Mermade/openapi-filter) - Filter internal components from OpenAPI Descriptions | Node.js | ✅ | ✅ | ✅ |  |
| [ OpenAPI Schema to JSON Schema ](https://www.npmjs.com/package/@openapi-contrib/openapi-schema-to-json-schema) - Due to the OpenAPI v3.0 and JSON Schema discrepancy, you can use this JS library to convert OpenAPI Schema objects to proper JSON Schema. | TypeScript | ❌ | ✅ | ❌ |  |
| [ OpenAPI TypeScript ](https://github.com/drwpow/openapi-typescript) - Convert static OpenAPI schemas to TypeScript types quickly using pure Node.js. Fast, lightweight, (almost) dependency-free, and no Java/node-gyp/running OpenAPI servers necessary. | TypeScript | ✅ | ✅ | ✅ |  |
| [ openapi-format ](https://www.npmjs.com/package/openapi-format) - A CLI to format an OpenAPI document by ordering fields in a hierarchical order, with the option to filter out flags, tags, methods, operationIDs; including the option to convert an OpenAPI 3.0 document to an OpenAPI version 3.1. | Node.js | ✅ | ✅ | ❌ |  |
| [ openapi-python-client ](https://github.com/openapi-generators/openapi-python-client) - Generate modern Python clients from OpenAPI 3.0 documents. | Python | ❌ | ✅ | ❌ |  |
| [ openapi-to-postman ](https://github.com/postmanlabs/openapi-to-postman) - Convert OpenAPI 3.0 specs to the Postman Collection (v2) format | JavaScript | ❌ | ✅ | ❌ |  |
| [ portman ](http://getportman.com/) - Port OpenAPI Spec to Postman Collection, with contract & variation tests included! | Node.js | ❌ | ✅ | ❌ |  |
| [ swagger2openapi ](https://mermade.org.uk/openapi-converter) - Upgrade files from OpenAPI v2.0 to v3.0, bundling into one mega file or respecting $refs. Part of oas-kit. | Node.js / CLI | ❌ | ✅ | ✅ |  |

## [Data Validators](https://openapi.tools/#data-validators)

Check to see if API requests and responses are lining up with the API description.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ api-codegen-ts ](https://www.npmjs.com/package/@nll/api-codegen-ts) - Generates TypeScript models, response validators, and operation controllers from OpenAPI descriptions | TypeScript | ❌ | ✅ | ✅ |  |
| [ APIFuzzer ](https://github.com/KissPeter/APIFuzzer) - Fuzz test your application using your OpenAPI definition without coding. Integrate into CI/CD, get Junit XML test result and JSON report of failures | Python | ❌ | ✅ | ✅ |  |
| [ committee ](https://github.com/interagent/committee) - Validation middleware for Rack server. This gem validates request and response using an OpenAPI Description. And convert parameter string to specific Ruby object (e.g. convert datetime string to DateTime class). | Ruby | ❌ | ✅ | ✅ |  |
| [ express-openapi-validator ](https://github.com/cdimascio/express-openapi-validator) - 🦋 Auto-validate API requests and responses in ExpressJS. | JavaScript | ❌ | ✅ | ❌ |  |
| [ JSONSchema::Validator ](https://metacpan.org/pod/JSONSchema::Validator) - A Perl library which validates request/response according to an OpenAPI specification | Perl | ❌ | ✅ | ❌ |  |
| [ kin-openapi ](https://github.com/getkin/kin-openapi) - OpenAPI 3.0 (and Swagger v2) implementation for Go (parsing, converting, validation, and more) | Go | ❌ | ✅ | ❌ |  |
| [ Mayhem for API ](https://forallsecure.com/mayhem-for-api) - Probe your REST API with an infinite stream of test cases generated automatically from your OpenAPI specification. | Any | ✅ | ✅ | ✅ |  |
| [ oas-tools ](https://github.com/isa-group/oas-tools) - NodeJS module to manage RESTful APIs defined with OpenAPI 3.0 Description over Express servers, including security validations | Node.js | ❌ | ✅ | ❌ |  |
| [ OpenAPI Enforcer ](https://www.npmjs.com/package/openapi-enforcer) - Validate your OpenAPI document, serialize, deserialize, and validate incoming requests and outgoing responses, and simplify response building. You can even produce mock data. | Node.js | ❌ | ✅ | ✅ |  |
| [ OpenAPI HttpFoundation Testing ](https://github.com/osteel/openapi-httpfoundation-testing) - Strengthen your API tests by validating HttpFoundation responses against OpenAPI definitions | PHP | ✅ | ✅ | ❌ |  |
| [ openapi-core ](https://github.com/p1c2u/openapi-core) - Validate your requests and responses against an OpenAPI 3 specification and get very verbose and human-readable descriptions of errors. You will receive a deserialized object along with validation result, so you won't need to deserialize it twice. | Python | ✅ | ✅ | ❌ |  |
| [ openapi-data-validator ](https://github.com/Rhosys/openapi-data-validator.js/blob/main/README.md) - Validate API requests against an OpenAPI schema. Lightweight, focused, and integrates with any framework | Node.js/ Javascript | ❌ | ✅ | ❌ |  |
| [ openapi-examples-validator ](https://github.com/codekie/openapi-examples-validator) - Validates embedded JSON-examples in OpenAPI-specs | JavaScript | ❌ | ✅ | ✅ |  |
| [ openapi-psr7-validator ](https://github.com/thephpleague/openapi-psr7-validator) - Using a PHP framework that supports PSR-7? Get free validation without writing a bunch of code, by registering this middleware and pointing it at your API description document. | PHP | ✅ | ✅ | ❌ |  |
| [ openapi-spring-webflux-validator ](https://github.com/cdimascio/openapi-spring-webflux-validator) - 🌱 A friendly kotlin library to validate API endpoints using an OpenAPI 3.0 or OpenAPI 2.0 specification | Java/Kotlin | ❌ | ✅ | ✅ |  |
| [ openapi-validator-bundle ](https://github.com/cydrickn/openapi-validator-bundle) - Validates Request and Response using Symfony Framework | PHP | ❌ | ✅ | ❌ |  |
| [ openapi-validator-middleware ](https://www.npmjs.com/package/openapi-validator-middleware) - Provides data validation within an Express, Koa or Fastify app according to a OpenAPI definition. It uses Ajv under the hood for validation. | Node.js | ❌ | ✅ | ✅ |  |
| [ openapi4j ](https://github.com/openapi4j/openapi4j) - Parse Description Document, validate API requests and responses using OpenAPI 3.x. | Java | ❌ | ✅ | ❌ |  |
| [ openVALIDATION ](https://docs.openvalidation.io/openapi/openapi-specification) - Allows complex validation rules to be specified in openAPI spec files using natural language. | Java | ❌ | ✅ | ❌ |  |
| [ vacuum ](https://quobix.com/vacuum) - The worlds fastest OpenAPI linter and validator. Compatible with Spectral rule-sets and designed for enterprise-grade speed and scale. | go | ✅ | ✅ | ✅ |  |

## [Documentation](https://openapi.tools/#documentation)

Render API Description as HTML (or maybe a PDF) so slightly less technical people can figure out how to work with the API.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ APIGit ](https://apigit.com/) - native Git based collaboration platform for API design, document, mock and share. | SaaS | ✅ | ✅ | ✅ |  |
| [ APIMatic Developer Experience Portal ](https://apimatic.io/developer-experience-portal) - Customizable developer portals packed with language specific documentation, client libraries, code samples, an API console and much more. | SaaS | ✅ | ✅ | ✅ |  |
| [ Apitive Studio ](https://www.apitive.com/) - A platform for Digital Product Managers and API Consultants to design REST APIs with in-built mock and documentation. | Angular 7.0, Java / Saas | ❌ | ✅ | ✅ |  |
| [ APITree ](https://apitree.com/) - HUB for managing and sharing APIs. Converts OpenAPI v2 / v3 files into beautiful API documentation. | SaaS | ❌ | ✅ | ✅ |  |
| [ BlocklyAutomation ](https://ignatandrei.github.io/BlocklyAutomation/) - Input any OpenAPI document to have generated Blocks in Blockly form to test and generate documentation. | Javascript / .NET | ✅ | ✅ | ❌ |  |
| [ Bump.sh ](https://bump.sh/?utm_source=openapi_tools&utm_medium=referral&utm_campaign=openapi) - Bump.sh generates elegant documentation and changelogs from your OpenAPI specifications. Git diff, for your API. Integrates with CI and Slack. | SaaS | ✅ | ✅ | ✅ |  |
| [ DeveloperHub ](https://developerhub.io/) - Collaboration platform for product and API documentation | SaaS | ✅ | ✅ | ✅ |  |
| [ Elements ](https://stoplight.io/open-source/elements) - Build beautiful, interactive API Docs with embeddable React or Web Components, powered by OpenAPI and Markdown | Javascript / Custom Element | ✅ | ✅ | ✅ |  |
| [ jekyll-openapi ](https://github.com/robertlove/jekyll-openapi) - An OpenAPI 3 documentation website generator built with Jekyll for use on GitHub Pages. | Jekyll | ❌ | ✅ | ❌ |  |
| [ Kong Enterprise Edition ](https://konghq.com/kong-enterprise-edition/) - Highly customizable developer portal with developer onboarding, integrated with the Kong API Gateway | Lua | ✅ | ✅ | ✅ |  |
| [ Kusk Gateway ](https://docs.kusk.io/) - Kusk-Gateway is an OpenAPI-driven API Gateway for Kubernetes. It empowers you to develop, validate, mock and deploy your APIs in a matter of minutes using both manual and automated GitOps/APIOps workflows. | Kubernetes | ❌ | ✅ | ✅ |  |
| [ LucyBot DocGen ](https://lucybot.com/docgen) - Generate a customizable website, with API documentation, console, and interactive workflows, from an OpenAPI spec | JavaScript | ❌ | ✅ | ✅ |  |
| [ MkDocs Swagger UI Tag ](https://blueswen.github.io/mkdocs-swagger-ui-tag/) - A MkDocs plugin supports for add Swagger UI in page. | Python | ❌ | ✅ | ✅ |  |
| [ MrinDoc ](https://mrin9.github.io/OpenAPI-Viewer/) - OpenAPI description document viewer. | Vue.JS | ❌ | ✅ | ✅ |  |
| [ Nexmo OAS Renderer ](https://github.com/Nexmo/nexmo-oas-renderer) - Ruby OpenAPI docs rendering, use standalone or add to your Rails app | Ruby | ❌ | ✅ | ❌ |  |
| [ oas-tools ](https://github.com/isa-group/oas-tools) - NodeJS module to manage RESTful APIs defined with OpenAPI 3.0 Description over Express servers, including security validations | Node.js | ❌ | ✅ | ❌ |  |
| [ oas3-api-snippet-enricher ](https://github.com/cdwv/oas3-api-snippet-enricher/) - Enrich your existing description documents with generated code samples | JavaScript | ❌ | ✅ | ❌ |  |
| [ OpenAPI Explorer ](https://github.com/Rhosys/openapi-explorer/blob/main/README.md) - Generate and render fully customizable API documentation, then explore and execute API requests via the integrated console. | Javascript/Custom Element | ✅ | ✅ | ✅ |  |
| [ openapi-dev-tool ](https://github.com/lyra/openapi-dev-tool) - OpenAPI Dev Tool proposes to developers a unique tool to address development and industrialization needs! | JavaScript | ❌ | ✅ | ❌ |  |
| [ openapi-viewer ](https://koumoul.com/openapi-viewer/) - Browse and test a REST API described with the OpenAPI 3.0 Specification | Vue.js | ❌ | ✅ | ❌ |  |
| [ OpenDocumenter ](https://github.com/ouropencode/OpenDocumenter) - OpenDocumenter is a automatic documentation generator for OpenAPI v3 schemas. Simply provide your schema file in JSON or YAML, then sit back and enjoy the documentation. | Vue.js | [👷](https://github.com/ouropencode/OpenDocumenter/issues/2) | ✅ | ✅ |  |
| [ RapiDoc ](https://rapidocweb.com/) - Custom Element to view OpenAPI descriptions. | Web Component | ✅ | ✅ | ✅ |  |
| [ RapiPdf ](https://mrin9.github.io/RapiPdf) - Custom Element to generate PDF from OpenAPI descriptions. | Web Component | ❌ | ✅ | ✅ |  |
| [ ReadMe ](https://readme.com/) - Build beautiful, personalized, interactive developer hubs. 🦉 | SaaS | ✅ | ✅ | ✅ |  |
| [ ReDoc ](https://rebilly.github.io/ReDoc/) - OpenAPI-generated API Reference Documentation | React.js | ✅ | ✅ | ✅ |  |
| [ RestCase Docs ](https://www.restcase.com/platform/docs) - An API-first and security-first management platform. Design visually and we will create a beautiful API documentation for your APIs. | SaaS | ❌ | ✅ | ✅ |  |
| [ Restish ](https://rest.sh/) - A CLI for REST-ish APIs with HTTP/2, built-in auth, content negotiation, caching, and more that understands and can discover OpenAPI descriptions. | CLI / Go | ❌ | ✅ | ❌ |  |
| [ Stoplight Docs ](https://stoplight.io/api-documentation) - Create beautiful, customizable, interactive API documentation generated from OpenAPI, integrated with Stoplight Studio. | SaaS | ✅ | ✅ | ✅ |  |
| [ widdershins ](https://mermade.github.io/shins) - Generate Slate/Shins markdown from OpenAPI 2.0/3.0.x | Node.js | ❌ | ✅ | ✅ |  |
| [ Zuplo (OpenAPI-based gateway and documentation) ](https://www.zuplo.com/) - Zuplo is an API gateway designed for developers. Natively powered by OpenAPI (3.1 or 3.0), zuplo offers an OpenAPI design surface, API documentation and a serverless, programmable edge gateway that includes request validation, auth, rate-limiting and more. | Web / SaaS | ✅ | ✅ | ❌ |  |

## [DSL](https://openapi.tools/#dsl)

Writing YAML by hand is no fun, and maybe you don't want a GUI, so use a Domain Specific Language to write OpenAPI in your language of choice.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ BOATS ](https://www.npmjs.com/package/boats) - BOATS allows for larger teams to contribute to multi-file OpenAPI definitions by writing Nunjucks tpl syntax in yaml with a few important helpers to ensure stricter consistency, eg operationId: <$ uniqueOpId() $>. | Node.js | ❌ | ✅ | ✅ |  |
| [ CUE ](https://cuelang.org/docs/integrations/openapi/) - CUE is an open source language, with a rich set of APIs and tooling, for defining, generating, and validating all kinds of data configuration, APIs, database schemas, code, etc. CUE currently supports generating OpenAPI through its API. | CUE | ❌ | ✅ | ❌ |  |
| [ Goa ](https://goa.design/) - Goa provides a holistic approach for developing remote APIs and microservices in Go. implementers don't have to worry about the documentation getting out of sync as Goa takes care of generating OpenAPI specifications for HTTP based services and gRPC protocol buffer files for gRPC based services | Go | ❌ | ✅ | ✅ |  |
| [ kotlin-openapi3-dsl ](https://github.com/derveloper/kotlin-openapi3-dsl) - kotlin-openapi3-dsl is a DSL written in Kotlin to write OpenAPI descriptions in plain Kotlin. | Kotlin | ❌ | ✅ | ❌ |  |
| [ Spot ](https://github.com/airtasker/spot) - A concise, developer-friendly way to describe your API contract. | TypeScript | ❌ | ✅ | ✅ |  |
| [ Supermodel ](https://supermodel.io/) - Model your data using JSON Schema, refer and remix the models freely, convert to various formats including OAS v2/v3. | SaaS | ❌ | ✅ | ✅ |  |

## [Text Editors](https://openapi.tools/#text-editors)

Text editors give you visual feedback whilst you write OpenAPI, so you can see what docs might look like.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ Atom/linter-openapi ](https://atom.io/packages/linter-openapi) - This plugin for Atom Linter will lint OpenAPI YAML files using openapi-enforcer node package. | JavaScript | ❌ | ✅ | ❌ |  |
| [ Atom/linter-swagger ](https://atom.io/packages/linter-swagger) - This plugin for Atom Linter will lint OpenAPI, both JSON and YAML using swagger-parser node package. | JavaScript | ❌ | ✅ | ✅ |  |
| [ KaiZen-OpenAPI-Editor ](https://github.com/RepreZen/KaiZen-OpenAPI-Editor) - Full-featured Eclipse editor for OpenAPI, also available on Eclipse Marketplace. | Java | ❌ | ✅ | ✅ |  |
| [ Senya Editor ](https://senya.io/) - JetBrains IDE plugin to show Swagger UI as a preview, for visual feedback as you edit. | Java | [👷](https://youtrack.jetbrains.com/issue/IDEA-294782/Add-support-for-OpenAPI-31-in-OpenAPI-Specifications-plugin) | ✅ | ✅ |  |
| [ Swagger Editor ](https://github.com/swagger-api/swagger-editor) - Design, describe, and document your API on the first open source editor fully dedicated to OpenAPI-based APIs. | Node.js | ❌ | ✅ | ✅ |  |
| [ SwaggerHub ](https://swagger.io/tools/swaggerhub/) - API design and documentation platform to improve collaboration, standardize development workflow and centralize their API discovery and consumption. | SaaS/On-Premise NodeJS | ❌ | ✅ | ✅ |  |
| [ VSCode OpenAPI ](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi) - OpenAPI extension for Visual Studio Code - new file templates, navigation, intellisense, code snippets. | Any | ❌ | ✅ | ✅ |  |
| [ VSCode OpenAPI Snippets ](https://marketplace.visualstudio.com/items?itemName=proohit.openapi-snippets) - OpenAPI Snippets for Visual Studio Code editor, includes split file validation | Any | ❌ | ✅ | ❌ |  |
| [ VSCode/openapi-lint ](https://marketplace.visualstudio.com/items?itemName=mermade.openapi-lint) - OpenAPI 2.0/3.0.x intellisense, validator and linter for Visual Studio Code | Node.js | ❌ | ✅ | ✅ |  |
| [ VSCode/Redocly OpenAPI ](https://marketplace.visualstudio.com/items?itemName=Redocly.openapi-vs-code) - Redocly OpenAPI is a Visual Studio Code extension that helps you write, validate, preview, and maintain your OpenAPI documents. | Node.js | ✅ | ✅ | ✅ |  |

## [GUI Editors](https://openapi.tools/#gui-editors)

Visual editors help you design APIs without needing to memorize the entire OpenAPI specification.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ ApiBldr ](https://www.apibldr.com/) - Web-Based API Designer for OpenAPI (swagger) and AsyncAPI specifications. | Angular 9.0 / Saas | ❌ | ✅ | ✅ |  |
| [ Apicurio Studio ](https://www.apicur.io/) - Web-Based Open Source API Design via the OpenAPI specification. | Angular 7.0, Java / Saas | ❌ | ✅ | ✅ |  |
| [ APIGit ](https://apigit.com/) - native Git based collaboration platform for API design, document, mock and share. | SaaS | ✅ | ✅ | ✅ |  |
| [ Apitive Studio ](https://www.apitive.com/) - A platform for Digital Product Managers and API Consultants to design REST APIs with in-built mock and documentation. | Angular 7.0, Java / Saas | ❌ | ✅ | ✅ |  |
| [ Flotiq - headless CMS with OpenAPI support ](https://flotiq.com/) - Visually define your Content Types, Flotiq automatically generates your own OpenAPI v3 compatible endpoints, SDKs and Postman collections. |  | ❌ | ✅ | ❌ |  |
| [ Hackolade ](https://hackolade.com/) - A visual editor for OpenAPI v2/v3, from the pioneer in data modeling for NoSQL databases. | ReactJS | ❌ | ✅ | ✅ |  |
| [ JetBrains tools (IntelliJ IDEA, PyCharm etc.) ](https://plugins.jetbrains.com/plugin/14394-openapi-specifications) - JetBrains development tools like IntelliJ IDEA, PyCharm and others come with a bundled *OpenAPI Specifications* plugin. The plugin allows you to write the OpenAPI specifications and supports you with validations, formatting, code-completion etc. It supports a *text view* as well as a rendered SwaggerUI-like *graphical interface*. | Java, Python | ❌ | ✅ | ✅ |  |
| [ OAIE Sketch ](https://www.github.com/OAIE/oaie-sketch) - Browser based OpenApi Integrated Editor with side-by side view of the yaml and an interactive graph. | Vue.js | ❌ | ✅ | ❌ |  |
| [ RestCase Designer ](https://www.restcase.com/platform/design) - A design-first API managment platform with WYSIWYG API Designer for OpenAPI and AsyncAPI specifications. | Angular 9.0 / Saas | ❌ | ✅ | ✅ |  |
| [ Stoplight Studio ](https://stoplight.io/studio) - Stoplight Studio is a GUI/text editor with linting and mocking built right in. It can run on the desktop with local files, and in the browser powered by your existing GitHub, GitLab, or BitBucket repos. | Desktop / SaaS | ✅ | ✅ | ✅ |  |
| [ Zuplo (OpenAPI-based gateway and documentation) ](https://www.zuplo.com/) - Zuplo is an API gateway designed for developers. Natively powered by OpenAPI (3.1 or 3.0), zuplo offers an OpenAPI design surface, API documentation and a serverless, programmable edge gateway that includes request validation, auth, rate-limiting and more. | Web / SaaS | ✅ | ✅ | ❌ |  |

## [Learning](https://openapi.tools/#learning)

Whether you're trying to get documentation for a third party API based on traffic, or are trying to switch to design-first at an organization with no OpenAPI at all, learning can help you move your API spec forward and keep it up to date.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ APIClarity ](https://apiclarity.io/) - Reconstruct Open API Specifications from real-time workload traffic seamlessly. | Golang, JavaScript | ❌ | ✅ | ✅ |  |
| [ InducOapi ](https://pypi.org/project/inducoapi) - A simple python module to generate OpenAPI Description Documents by supplying request/response bodies. | Python | ❌ | ✅ | ❌ |  |
| [ optic ](https://useoptic.com/) - Build your first OpenAPI description from traffic. Use Optic to patch the OpenAPI every time it detects new API behavior. | cli | ✅ | ✅ | ❌ |  |
| [ Response2Schema ](https://github.com/dsuurlant/response2schema) - Takes any JSON response and generates an OpenAPI definition document with the component schema and a default endpoint. | PHP | ❌ | ✅ | ❌ |  |
| [ Swagger Inspector ](https://swagger.io/tools/swagger-inspector/) - Run mock requests in a webapp and Swagger Inspector infers your OpenAPI description. | SaaS | ❌ | ✅ | ✅ |  |

## [Mock Servers](https://openapi.tools/#mock)

Fake servers that take description document as input, then route incoming HTTP requests to example responses or dynamically generates examples.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ API Sprout ](https://github.com/danielgtaylor/apisprout) - Lightweight, blazing fast, cross-platform OpenAPI 3 mock server with validation | cli / Docker | ❌ | ✅ | ❌ |  |
| [ APIGit ](https://apigit.com/) - native Git based collaboration platform for API design, document, mock and share. | SaaS | ✅ | ✅ | ✅ |  |
| [ Apitive Studio ](https://www.apitive.com/) - A platform for Digital Product Managers and API Consultants to design REST APIs with in-built mock and documentation. | Angular 7.0, Java / Saas | ❌ | ✅ | ✅ |  |
| [ Connexion ](https://connexion.readthedocs.io/en/latest/) - OpenAPI First framework for Python on top of Flask with automatic endpoint validation & OAuth2 support | Python | ❌ | ✅ | ✅ |  |
| [ Fakeit ](https://github.com/justinfeng/fakeit) - Create mock server from OpenAPI 3 specification with random response generation and request validation. | cli / Docker | ❌ | ✅ | ❌ |  |
| [ Falcon Heavy ](https://github.com/NotJustAToy/falcon-heavy) - The framework for building app backends and microservices via the API design-first workflow. | Python | ❌ | ✅ | ❌ |  |
| [ Karate-IDE ](https://marketplace.visualstudio.com/items?itemName=KarateIDE.karate-ide) - Generates KarateDSL Tests and Mocks from OpenAPI 3.0 documents and so you can quickly test/explore your API. | VSCode Extension | ❌ | ✅ | ❌ |  |
| [ Kusk Gateway ](https://docs.kusk.io/) - Kusk-Gateway is an OpenAPI-driven API Gateway for Kubernetes. It empowers you to develop, validate, mock and deploy your APIs in a matter of minutes using both manual and automated GitOps/APIOps workflows. | Kubernetes | ❌ | ✅ | ✅ |  |
| [ Meeshkan ](https://meeshkan.com/) - Meeshkan is an automated testing and mocking tool. It offers first-class support for GraphQL APIs, but Meeshkan is also built to handle REST APIs and third-party dependencies. | SaaS | ❌ | ✅ | ❌ |  |
| [ Microcks ](https://microcks.io/) - Kubernetes native tool for API Mocking and Testing. Turn your OAI contract examples into ready to use mocks. Use examples to test and validate implementations according spec and schema elements. | Self-hosted / SaaS | ✅ | ✅ | ✅ |  |
| [ Mockintosh ](https://mockintosh.io/) - Mocks for CloudNative Environments - Converts OpenAPI files to Mocks and use them to develop in isolated environments and test edge cases, Async call to queues such as Kafka or RabbitMQ or simulate performance & chaos testing | CLI / Docker | ✅ | ✅ | ✅ |  |
| [ MockLab ](https://www.mocklab.io/docs/getting-started/) - SaaS platform to upload your spec to create a mock server | SaaS | ❌ | ✅ | ✅ |  |
| [ OpenAPI Mocker ](https://www.npmjs.com/package/open-api-mocker) - Standalone nodejs based OpenAPI 3 mock server, docker-friendly with request validation and autoreload. | nodejs | ❌ | ✅ | ❌ |  |
| [ openapi-data-mocker ](https://github.com/ybelenko/openapi-data-mocker) - Tiny library to generate basic OpenAPI Data Types. Consider it as extended Faker package. First version able to mock most of the data formats. It doesn't support polymorphism yet, but work in progress. May be useful for writing custom unit tests. | PHP | ❌ | ✅ | ❌ |  |
| [ orval ](https://orval.dev/) - orval is able to generate client with appropriate type-signatures (TypeScript) from any valid OpenAPI v3 or Swagger v2 specification, either in yaml or json formats. Generate, valid, cache and mock in your frontend applications all with your OpenAPI specification. 🍺 | Typescript, React, Vue, Svelte, Angular | ❌ | ✅ | ✅ |  |
| [ Prism ](https://stoplight.io/prism) - Turn any OAI file into an API server with mocking, transformations, validations, and more. | cli | ✅ | ✅ | ✅ |  |
| [ Sandbox ](https://getsandbox.com/) - SaaS, self-hosted, or CLI tool for turning OpenAPI (and other) descriptions into a mock server, where you can modify behaviour, simulate downtime, and any other nonsense you can think of thanks to a built-in code editor! | SaaS / Java | ❌ | ✅ | ✅ |  |
| [ Unmock ](https://unmock.io/) - API integration testing library that intercepts outgoing requests and serves back mock data based on the OpenAPI descriptions. | Node.js | ❌ | ✅ | ❌ |  |
| [ yii2-app-api ](https://github.com/cebe/yii2-app-api) - Generate Server side API code with routing, models, data validation and database schema from an OpenAPI description. Based on Yii Framework. | PHP | ❌ | ✅ | ❌ |  |

## [Description Validators](https://openapi.tools/#description-validators)

Check your API description to see if it is valid OpenAPI.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ @redocly/openapi-cli ](https://redoc.ly/openapi-cli) - OpenAPI 3 CLI toolbox with rich validation and bundling features. | CLI & TypeScript/JavaScript | ✅ | ✅ | ✅ |  |
| [ BigstickCarpet/swagger-cli ](https://github.com/BigstickCarpet/swagger-cli) - Simple validation for OpenAPI files, supporting JSON/YAML and v2/v3 description documents. | Node.js / CLI | ❌ | ✅ | ✅ |  |
| [ Cherrybomb ](https://www.blstsecurity.com/cherrybomb) - A CLI tool that helps avoid undefined user behaviour by validating your API descriptions, to make sure key parts are not missing or vague. | Rust | ✅ | ✅ | ❌ |  |
| [ express-openapi-validator ](https://github.com/cdimascio/express-openapi-validator) - 🦋 Auto-validate API requests and responses in ExpressJS. | JavaScript | ❌ | ✅ | ❌ |  |
| [ oas-tools ](https://github.com/isa-group/oas-tools) - NodeJS module to manage RESTful APIs defined with OpenAPI 3.0 Description over Express servers, including security validations | Node.js | ❌ | ✅ | ❌ |  |
| [ OpenAPI Enforcer ](https://www.npmjs.com/package/openapi-enforcer) - Validate your OpenAPI document, serialize, deserialize, and validate incoming requests and outgoing responses, and simplify response building. You can even produce mock data. | Node.js | ❌ | ✅ | ✅ |  |
| [ OpenAPI Style Validator ](https://github.com/OpenAPITools/openapi-style-validator) - A customizable style validator to make sure your OpenAPI description follows your organization's standards. | Java, CLI | ❌ | ✅ | ✅ |  |
| [ OpenAPI Validator ](https://github.com/IBM/openapi-validator) - Configurable and extensible validator/linter for OpenAPI documents | Node.js | ❌ | ✅ | ✅ |  |
| [ openapi-examples-validator ](https://github.com/codekie/openapi-examples-validator) - Validates embedded JSON-examples in OpenAPI-specs | JavaScript | ❌ | ✅ | ✅ |  |
| [ openapi-spec-validator ](https://github.com/p1c2u/openapi-spec-validator) - OpenAPI Description validator | Python | ✅ | ✅ | ✅ |  |
| [ openapi-spring-webflux-validator ](https://github.com/cdimascio/openapi-spring-webflux-validator) - 🌱 A friendly kotlin library to validate API endpoints using an OpenAPI 3.0 or OpenAPI 2.0 specification | Java/Kotlin | ❌ | ✅ | ✅ |  |
| [ openVALIDATION ](https://docs.openvalidation.io/openapi/openapi-specification) - Allows complex validation rules to be specified in openAPI spec files using natural language. | Java | ❌ | ✅ | ❌ |  |
| [ php-openapi ](https://github.com/cebe/php-openapi) - A PHP library for manipulating and validating OpenAPI 3.0 Descriptions | PHP | [👷](https://github.com/cebe/php-openapi/pull/128) | ✅ | ❌ |  |
| [ Spectral ](https://stoplight.io/spectral) - A flexible JSON/YAML object linter with portable "rulesets" and custom functions. | CLI & TypeScript/JavaScript | ✅ | ✅ | ✅ |  |
| [ super-linter ](https://github.com/github/super-linter) - GitHub Action to lint repositories as part of CI/CD. Implements the latest version of Spectral. | CLI / Docker | ❌ | ✅ | ✅ |  |
| [ vacuum ](https://quobix.com/vacuum) - The worlds fastest OpenAPI linter and validator. Compatible with Spectral rule-sets and designed for enterprise-grade speed and scale. | go | ✅ | ✅ | ✅ |  |

## [Security](https://openapi.tools/#security)

By poking around your OpenAPI description, some tools can look out for attack vectors you might not have noticed.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ 42crunch ](https://42crunch.com/) - A unique set of integrated API security tools that allow discovery, remediation of OpenAPI vulnerabilities and runtime protection against API attacks. | SaaS | ❌ | ✅ | ✅ |  |
| [ API Insights ](https://restcase.com/platform/security) - RestCase executes hundrends of security and quality checks against the API definition, the API insights report provides detailed security scoring for prioritization, and remediation advice to help developers define the best API definition possible. | SaaS | ❌ | ✅ | ✅ |  |
| [ cats ](https://github.com/Endava/cats) - CATS is a REST API Fuzzer and negative testing tool for OpenAPI endpoints. CATS automatically generates, runs and reports tests with minimum configuration and no coding effort. Tests are self-healing and do not require maintenance. | Java | ✅ | ✅ | ✅ |  |
| [ Mayhem for API ](https://forallsecure.com/mayhem-for-api) - Probe your REST API with an infinite stream of test cases generated automatically from your OpenAPI specification. | Any | ✅ | ✅ | ✅ |  |
| [ oas-tools ](https://github.com/isa-group/oas-tools) - NodeJS module to manage RESTful APIs defined with OpenAPI 3.0 Description over Express servers, including security validations | Node.js | ❌ | ✅ | ❌ |  |
| [ openapi-fuzzer ](https://github.com/matusf/openapi-fuzzer) - Based on OpenAPI specification, openapi-fuzzer provides random data as inputs to the API endpoints in order to find bugs. | Rust | ❌ | ✅ | ❌ |  |
| [ OpenAPI3 Fuzzer ](https://pypi.org/project/openapi3-fuzzer/) - Simple fuzzer for OpenAPI 3 specification based APIs. Verifies responses and sends various attack patterns. | Python | ❌ | ✅ | ❌ |  |
| [ OWASP ZAP ](https://www.zaproxy.org/) - OWASP ZAP is a free and open source web security tool that can be used manually or completely automated. It supports importing OpenAPI v2 and v3 definitions to allow an API to be thoroughly security tested. | Java | ❌ | ✅ | ✅ |  |
| [ RESTler ](https://github.com/microsoft/restler) - RESTler is the first stateful REST API fuzzing tool for automatically testing cloud services through their REST APIs and finding security and reliability bugs in these services. RESTler analyzes the OpenAPI description of a cloud service, and then generates and executes tests that exercise the service through its REST API. During testing, it checks for specific classes of bugs and dynamically learns how the service behaves from prior service responses. | Any | ❌ | ✅ | ✅ |  |
| [ StackHawk HawkScan ](https://stackhawk.com/) - StackHawk is an application vulnerability scanner purpose built for developers to use in the DevOps pipeline. It leverages a provided OpenAPI v2 or v3 spec file for route discovery and enhanced scanning. | SaaS | ❌ | ✅ | ✅ |  |

## [SDK Generators](https://openapi.tools/#sdk)

Generate code to give to consumers, to help them avoid interacting at a HTTP level.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ api-codegen-ts ](https://www.npmjs.com/package/@nll/api-codegen-ts) - Generates TypeScript models, response validators, and operation controllers from OpenAPI descriptions | TypeScript | ❌ | ✅ | ✅ |  |
| [ APIMatic CodeGen ](https://apimatic.io/code-generation-as-a-service) - Bring in your API description (OAI v2/v3, RAML, API Blueprint, WSDL, etc.) to generate fully functional SDKs in over 10 languages. | SaaS | ✅ | ✅ | ✅ |  |
| [ Azure AutoRest ](https://github.com/Azure/autorest/) - Generates client libraries for accessing RESTful web services from an OpenAPI document. Supports C#, PowerShell, Go, Java, Node.js, TypeScript, Python, and Ruby. | TypeScript | ❌ | ✅ | ✅ |  |
| [ BlocklyAutomation ](https://ignatandrei.github.io/BlocklyAutomation/) - Input any OpenAPI document to have generated Blocks in Blockly form to test and generate documentation. | Javascript / .NET | ✅ | ✅ | ❌ |  |
| [ Counterfact ](https://counterfact.dev/) - A new (July 2022) project that converts an OpenAPI document to a full implementation that runs on ts-node. The idea is to replace the auto-generated code that returns random values with realistic code, one path at a time. If the spec is updated, you can regenerate the types and let the type checker show you which parts of the implementation need to be updated. The types can also be used in client-side code. | TypeScript / Node | ✅ | ✅ | ❌ |  |
| [ FabriKt ](https://github.com/cjbooms/fabrikt) - A sophisticated Kotlin code generation library capable of generating Jackson-annotated data classes, Spring Controller interfaces, and fault-tolerant OkHttp clients. Written in Kotlin, this library programatically generates code and is capable of handling advanced OpenApi3 specification features such as polymorphism. | Kotlin | ❌ | ✅ | ❌ |  |
| [ Flotiq - headless CMS with OpenAPI support ](https://flotiq.com/) - Visually define your Content Types, Flotiq automatically generates your own OpenAPI v3 compatible endpoints, SDKs and Postman collections. |  | ❌ | ✅ | ❌ |  |
| [ go-swagger ](https://goswagger.io/) - Unmaintained v2.0 only project seeking new maintainer, or probably a fork. Parser, validator, generates descriptions from code, or code from descriptions! | Go | ❌ | ❌ | ✅ |  |
| [ guardrail ](https://github.com/twilio/guardrail) - Principled code generation from OpenAPI descriptions | Scala, Java, ... | ❌ | ✅ | ✅ |  |
| [ janephp/open-api ](https://github.com/janephp/open-api) - Generate a PHP Client API (PSR-7 compatible) given a OpenAPI specification. | PHP | ❌ | ✅ | ✅ |  |
| [ Kiota Api Client Generator ](https://microsoft.github.io/kiota/) - Kiota is a cross platform API Client code generator that is small, fast, and optimized for API consumers to find APIs and generate client code for just the parts of the API that they need. One tool, for any OpenAPI described API, that delivers a consistent client experience in multiple languages. | C# | ❌ | ✅ | ✅ |  |
| [ NSwag ](http://nswag.org/) - OpenAPI toolchain for .NET, Web API and TypeScript | .NET | ❌ | ✅ | ✅ |  |
| [ oa-client ](https://github.com/ninofiliu/oa-client) - Flexible client helper for making and validating calls to OpenAPI backends. For Node and the browser. Runtime lib - no need for code generation! | TypeScript | ❌ | ✅ | ❌ |  |
| [ oazapfts! ](https://github.com/cellular/oazapfts) - Generate TypeScript clients from a given OpenAPI description document. | TypeScript | ❌ | ✅ | ✅ |  |
| [ OpenAPI Client Generators ](https://github.com/zijianhuang/openapiclientgen) - .NET Core command line program to generate strongly typed client API codes in C# on .NET Frameworks and .NET Core, and in TypeScript for Angular 5+, Aurelia, jQuery, AXIOS and Fetch API. | C# | ❌ | ✅ | ✅ |  |
| [ OpenAPI Commander ](https://www.npmjs.com/package/openapi-commander) - Generate a Node.js command line tool from an OpenAPI definition. | Node.js / CLI | ✅ | ✅ | ❌ |  |
| [ OpenAPI Generator ](https://openapi-generator.tech/) - A template-driven engine to generate documentation, API clients and server stubs in different languages by parsing your OpenAPI Description (community-driven fork of swagger-codegen) | Java | [👷](https://github.com/OpenAPITools/openapi-generator/issues/9083) | ✅ | ✅ |  |
| [ openapi-ts-sdk-builder ](https://github.com/nfroidure/openapi-ts-sdk-builder) - Generate a TypeScript SDK from OpenAPI 3 definitions. | Javascript / TypeScript | ✅ | ✅ | ❌ |  |
| [ restful-react ](https://github.com/contiamo/restful-react) - Generate React hooks with appropriate type-signatures from OpenAPI descriptions | React (Typescript) | ❌ | ✅ | ✅ |  |
| [ spring-openapi ](https://github.com/jrcodeza/spring-openapi) - OpenAPI v3 generator for Java Spring. Includes also client generation. Supports inheritance with discriminators, Jackson annotations and custom interceptors. | Java | ❌ | ✅ | ❌ |  |
| [ Typoas ](https://github.com/Embraser01/typoas) - Fully typed OpenAPI Typescript generator | Typescript | ❌ | ✅ | ❌ |  |
| [ Unchase.OpenAPI.Connectedservice ](https://github.com/unchase/Unchase.OpenAPI.Connectedservice) - Visual Studio extension to generate C# (TypeScript) HttpClient (or C# Controllers) code for OpenAPI web service with NSwag. | .NET | ❌ | ✅ | ✅ |  |

## [Server Implementations](https://openapi.tools/#server)

Easily create and implement resources and routes for your APIs.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ @eropple/nestjs-openapi3 ](https://github.com/eropple/nestjs-openapi3) - Integrates tightly with a NestJS application to infers complex descriptions and expresses them in its generated OpenAPI document. It then presents that document via ReDoc, and validates inputs for conformance to spec. | Node.js | ❌ | ✅ | ❌ |  |
| [ @nestjs/swagger ](https://docs.nestjs.com/recipes/swagger) - Official OpenAPI (Swagger) module for NestJS. Use decorators to define OpenAPI endpoint documentation, parameters and return types. Integrates tightly with a NestJS application. Ships with Swagger UI and serves OpenAPI v3 spec. | Node.js | ❌ | ✅ | ❌ |  |
| [ @smartrecruiters/openapi-first ](https://www.npmjs.com/package/@smartrecruiters/openapi-first) - Initializes your API express application with the description in OpenAPI 3.0 format using provided middlewares (parsers, validators, controller, defaults setting) or custom ones | Node.js | ❌ | ✅ | ❌ |  |
| [ API Platform ](https://api-platform.com/) - REST and GraphQL framework to build modern API-driven projects | PHP | ❌ | ✅ | ✅ |  |
| [ BaucisJS + baucis-openapi3 ](https://www.npmjs.com/package/baucis-openapi3) - Create REST resources with persistence on MongoDB and expose OpenAPI v.3 contracts | Node.js | ❌ | ✅ | ❌ |  |
| [ express-openapi ](https://www.npmjs.com/package/express-openapi) - An unopinionated OpenAPI framework for Express, which supports Promise based middleware, response handlers and Security Filtering. | Node.js / Typescript | [👷](https://github.com/vert-x3/vertx-web/issues/1872) | ✅ | ✅ |  |
| [ Falcon Heavy ](https://github.com/NotJustAToy/falcon-heavy) - The framework for building app backends and microservices via the API design-first workflow. | Python | ❌ | ✅ | ❌ |  |
| [ Fusio ](https://www.fusio-project.org/) - Open source API management platform | PHP | ❌ | ✅ | ❌ |  |
| [ LoopBack 4 ](https://loopback.io/) - A highly extensible object-oriented Node.js and TypeScript framework for building APIs and microservices with tight OpenAPI 3 integration. Serves Swagger UI and OpenAPI 3 spec out of the box. Generate code to interact with other OpenAPI-compliant APIs, or generate new API endpoints based on existing OpenAPI specs. | Node.js + TypeScript | ❌ | ✅ | ❌ |  |
| [ Mojolicious::Plugin::OpenApi ](https://metacpan.org/pod/Mojolicious::Plugin::OpenAPI) - Mojolicious::Plugin::OpenAPI is a plugin for Mojolicious framework that add routes and input/output validation to your Mojolicious application based on OpenAPI description documents.' | Perl | ❌ | ✅ | ✅ |  |
| [ oas-tools ](https://github.com/isa-group/oas-tools) - NodeJS module to manage RESTful APIs defined with OpenAPI 3.0 Description over Express servers, including security validations | Node.js | ❌ | ✅ | ❌ |  |
| [ OpenAPI Enforcer ](https://www.npmjs.com/package/openapi-enforcer) - Validate your OpenAPI document, serialize, deserialize, and validate incoming requests and outgoing responses, and simplify response building. You can even produce mock data. | Node.js | ❌ | ✅ | ✅ |  |
| [ OpenAPI Enforcer Middleware ](https://www.npmjs.com/package/openapi-enforcer-middleware) - An express middleware that makes it easy to write web services that follow an OpenAPI specification by leveraging the tools provided in the openapi-enforcer package. | Node.js | ❌ | ✅ | ✅ |  |
| [ openapi-backend ](https://www.npmjs.com/package/openapi-backend) - Build, Validate, Route, and Mock using OpenAPI specification. Framework-agnostic | Node.js + Typescript | ✅ | ✅ | ❌ |  |
| [ openapi-processor-spring ](https://docs.openapiprocessor.io/spring) - Generates java interfaces & model classes for Spring Boot (annotation based, MVC & WebFlux) from an openapi.yaml. Provides type mapping capabilities to adjust the generated code. Gradle support. | Java | ❌ | ✅ | ❌ |  |
| [ openapi-validator-middleware ](https://www.npmjs.com/package/openapi-validator-middleware) - Provides data validation within an Express, Koa or Fastify app according to a OpenAPI definition. It uses Ajv under the hood for validation. | Node.js | ❌ | ✅ | ✅ |  |
| [ SpringFox ](https://springfox.io/) - Automated JSON API documentation for APIs built with Spring and SpringBoot | Java, Kotlin, Groovy, or Ruby | ❌ | ✅ | ✅ |  |
| [ tsoa ](https://github.com/lukeautry/tsoa) - Creates OpenAPI docs and provides free runtime validation for your Koa, Express, Hapi (and more) services | Node.js / TypeScript | ❌ | ✅ | ✅ |  |
| [ Vert.x Web Api Contract ](https://vertx.io/docs/#web) - Create API endpoints with Vert.x 3 and OpenAPI 3 with automatic requests validation | Java, Kotlin, JavaScript, Groovy, Ruby, Ceylon or Scala | ❌ | ✅ | ❌ |  |
| [ Whook ](https://github.com/nfroidure/whook) - OpenAPI 3 based NodeJS server. | Javascript / TypeScript | ✅ | ✅ | ❌ |  |
| [ yii2-app-api ](https://github.com/cebe/yii2-app-api) - Generate Server side API code with routing, models, data validation and database schema from an OpenAPI description. Based on Yii Framework. | PHP | ❌ | ✅ | ❌ |  |

## [Miscellaneous](https://openapi.tools/#miscellaneous)

Anything else that does stuff with OpenAPI but hasn't quite got enough to warrant its own category.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ $ oas (CLI) ](https://openap.is/) - Generate OAS files from code comments and easily host them ($ npm install oas -g) | JavaScript | ❌ | ✅ | ✅ |  |
| [ @redocly/openapi-cli ](https://redoc.ly/openapi-cli) - OpenAPI 3 CLI toolbox with rich validation and bundling features. | CLI & TypeScript/JavaScript | ✅ | ✅ | ✅ |  |
| [ Django REST Framework ](https://www.django-rest-framework.org/api-guide/schemas/) - Automates generation of OpenAPI 3 description documents either as a static file (via CLI command) or a dynamic view within the Django REST Framework (DRF) application. | Python | ❌ | ✅ | ❌ |  |
| [ express-openapi-validator ](https://github.com/cdimascio/express-openapi-validator) - 🦋 Auto-validate API requests and responses in ExpressJS. | JavaScript | ❌ | ✅ | ❌ |  |
| [ Flotiq - headless CMS with OpenAPI support ](https://flotiq.com/) - Visually define your Content Types, Flotiq automatically generates your own OpenAPI v3 compatible endpoints, SDKs and Postman collections. |  | ❌ | ✅ | ❌ |  |
| [ laravel-openapi ](https://github.com/vyuldashev/laravel-openapi) - Generate OpenAPI 3 specification for Laravel Applications. | PHP | ❌ | ✅ | ❌ |  |
| [ Meta-API ](https://www.meta-api.io/) - A SaaS platform to integrate APIs using OpenAPI documents, and manipulation of data with online code editor, and automating configuration, authentication, deployment and monitoring. | SaaS | ✅ | ✅ | ✅ |  |
| [ oa-client ](https://github.com/ninofiliu/oa-client) - Flexible client helper for making and validating calls to OpenAPI backends. For Node and the browser. Runtime lib - no need for code generation! | TypeScript | ❌ | ✅ | ❌ |  |
| [ oasdiff ](https://github.com/tufin/oasdiff) - Golang module for deep comparison of two OpenAPI specifications. Available also as a command-line. | Go | ❌ | ✅ | ❌ |  |
| [ OAuth2 as OpenAPI Spec 3.0 components ](https://github.com/ybelenko/oauth2_as_oas3_components) - OAuth2 token endpoint described with OAS3 schema. All grants documented. Can be installed as NPM or Composer package. | Any | ❌ | ✅ | ❌ |  |
| [ OpenAPI Server Code Generator (oapi-codegen) ](https://github.com/deepmap/oapi-codegen) - Generate a client, server, and HTTP types for various Go HTTP servers, from an OpenAPI v3 specification | Go | ❌ | ✅ | ❌ |  |
| [ openapi-cli-tool ](https://pypi.org/project/openapi-cli-tool/) - Can list up defined API paths and bundle multi-file into one. Supports multiple file extensions. | Python | ❌ | ✅ | ❌ |  |
| [ openapi-comparator ](https://github.com/criteo/openapi-comparator) - C# library for comparing two OpenAPI specifications. | C# | ❌ | ✅ | ❌ |  |
| [ openapi-dev-tool ](https://github.com/lyra/openapi-dev-tool) - OpenAPI Dev Tool proposes to developers a unique tool to address development and industrialization needs! | JavaScript | ❌ | ✅ | ❌ |  |
| [ openapi-diff ](https://github.com/quen2404/openapi-diff) - Utility for comparing two OpenAPI specifications. | Java | ❌ | ✅ | ❌ |  |
| [ openapi-examples-validator ](https://github.com/codekie/openapi-examples-validator) - Validates embedded JSON-examples in OpenAPI-specs | JavaScript | ❌ | ✅ | ✅ |  |
| [ openapi-format ](https://www.npmjs.com/package/openapi-format) - A CLI to format an OpenAPI document by ordering fields in a hierarchical order, with the option to filter out flags, tags, methods, operationIDs; including the option to convert an OpenAPI 3.0 document to an OpenAPI version 3.1. | Node.js | ✅ | ✅ | ❌ |  |
| [ openapi-spring-webflux-validator ](https://github.com/cdimascio/openapi-spring-webflux-validator) - 🌱 A friendly kotlin library to validate API endpoints using an OpenAPI 3.0 or OpenAPI 2.0 specification | Java/Kotlin | ❌ | ✅ | ✅ |  |
| [ openVALIDATION ](https://docs.openvalidation.io/openapi/openapi-specification) - Allows complex validation rules to be specified in openAPI spec files using natural language. | Java | ❌ | ✅ | ❌ |  |
| [ optic diff ](https://www.useoptic.com/docs/openapi-diff) - Diff the effective API contract between any two versions of your OpenAPI description. Exit 1 on breaking changes. | Go | ✅ | ✅ | ❌ |  |
| [ php-openapi-faker ](https://github.com/canvural/php-openapi-faker) - Library to generate fake data for OpenAPI 3.x requests, responses and schemas. | PHP | ❌ | ✅ | ❌ |  |
| [ Restish ](https://rest.sh/) - A CLI for REST-ish APIs with HTTP/2, built-in auth, content negotiation, caching, and more that understands and can discover OpenAPI descriptions. | CLI / Go | ❌ | ✅ | ❌ |  |
| [ schema2dts ](https://github.com/nfroidure/schema2dts) - Create types definitions from an OpenAPI schema. | Javascript / TypeScript | ❌ | ✅ | ❌ |  |
| [ vacuum ](https://quobix.com/vacuum) - The worlds fastest OpenAPI linter and validator. Compatible with Spectral rule-sets and designed for enterprise-grade speed and scale. | go | ✅ | ✅ | ✅ |  |

## [Parsers](https://openapi.tools/#parsers)

Loads and read OpenAPI descriptions, so you can work with them programmatically.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ APIDevTools/swagger-parser ](https://github.com/APIDevTools/swagger-parser) - OpenAPI 2.0 and 3.0 parser and validator. Can also bundle multiple documents into one via `$ref`. | Node.js | ❌ | ✅ | ✅ |  |
| [ KaiZen OpenAPI Parser ](https://github.com/RepreZen/KaiZen-OpenApi-Parser) - High-performance Parser, Validator, and Java Object Model for OpenAPI 3.x | Java | ❌ | ✅ | ❌ |  |
| [ kin-openapi ](https://github.com/getkin/kin-openapi) - OpenAPI 3.0 (and Swagger v2) implementation for Go (parsing, converting, validation, and more) | Go | ❌ | ✅ | ✅ |  |
| [ libopenapi ](https://github.com/pb33f/libopenapi) - Enterprise grade, fully featured OpenAPI 3.1, 3.0 and Swagger parser for go. A complete toolset for reading and parsing OpenAPI and Swagger specifications. Comes complete with high and low-level APIs, diff engine, index and resolver. | go | ✅ | ✅ | ✅ |  |
| [ Microsoft/OpenAPI.NET ](https://github.com/Microsoft/OpenAPI.NET) - C# based parser with OpenAPI Description validation and migration support from V2 | .NET | ❌ | ✅ | ✅ |  |
| [ oas-tools ](https://github.com/isa-group/oas-tools) - NodeJS module to manage RESTful APIs defined with OpenAPI 3.0 Description over Express servers, including security validations | Node.js | ❌ | ✅ | ❌ |  |
| [ oas_parser ](https://github.com/Nexmo/oas_parser) - A Ruby parser for OpenAPI 3.0+ descriptions. | Ruby | ❌ | ✅ | ❌ |  |
| [ Object Oriented OpenAPI Specification ](https://github.com/goldspecdigital/oooas) - An object oriented approach to generating OpenAPI Descriptions, implemented in PHP | PHP | ❌ | ✅ | ❌ |  |
| [ openapi-filter ](https://github.com/Mermade/openapi-filter) - OpenAPI 2.0 and 3.0 filter utility. A CLI/module to filter out internal/private paths, operations, parameters, schemas etc from OpenAPI v1/OpenAPI v2/AsyncAPI definitions. Simply flag any OpenAPI object within the definition with an `x-internal` specification extension or target a OpenAPI property (tags, methods, OperationId), and it will be removed from the output. | Node.js | ❌ | ✅ | ✅ |  |
| [ openapi-format ](https://www.npmjs.com/package/openapi-format) - A CLI to format an OpenAPI document by ordering fields in a hierarchical order, with the option to filter out flags, tags, methods, operationIDs; including the option to convert an OpenAPI 3.0 document to an OpenAPI version 3.1. | Node.js | ✅ | ✅ | ❌ |  |
| [ openapi-snippet ](https://github.com/ErikWittern/openapi-snippet) - Generates code snippets in various languages & tools (cURL, Node, Python, Ruby, Java, Go, C#...), from OpenAPI documents. | Node.js | ❌ | ✅ | ✅ |  |
| [ openapi-snippet-cli ](https://github.com/richardkabiling/openapi-snippet-cli) - Adds code snippets in redoc style (x-codeSamples) to OpenAPI documents. This is a CLI wrapper for the "openapi-snippet". | Node.js | ❌ | ✅ | ✅ |  |
| [ OpenAPI-TS ](https://github.com/metadevpro/openapi3-ts) - TS Model & utils for OpenAPI 3.0.x contracts | TypeScript | ❌ | ✅ | ❌ |  |
| [ openapi3 ](https://github.com/Dorthu/openapi3) - An OpenAPI 3 Specification client, and validator, covering both description validation and limited data validation for Python 3. | Python | ❌ | ✅ | ❌ |  |
| [ OpenAPI3-Rust ](https://github.com/adwhit/openapi3-rust) - Rust serialization library for OpenAPI v3 | Rust | ❌ | ✅ | ❌ |  |
| [ openapi3_parser ](https://github.com/kevindew/openapi3_parser) - A Ruby implementation of parser and validator for the OpenAPI 3 Specification. | Ruby | ❌ | ✅ | ❌ |  |
| [ openapi4j ](https://github.com/openapi4j/openapi4j) - Parse Description Document, validate API requests and responses using OpenAPI 3.x. | Java | ❌ | ✅ | ❌ |  |
| [ php-openapi ](https://github.com/cebe/php-openapi) - A PHP library for manipulating and validating OpenAPI 3.0 Descriptions | PHP | [👷](https://github.com/cebe/php-openapi/pull/128) | ✅ | ❌ |  |
| [ psx-api ](http://phpsx.org/) - Parse and generate API specification formats | PHP | ❌ | ✅ | ✅ |  |
| [ swagger-parser ](https://github.com/swagger-api/swagger-parser) - Swagger Parser reads OpenAPI definitions into current Java POJOs. | Java | [👷](https://github.com/swagger-api/swagger-parser/pull/1730) | ✅ | ✅ |  |

## [Testing](https://openapi.tools/#testing)

Quickly execute API requests and validate responses on the fly through command line or GUI interfaces.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ Assertible ](https://assertible.com/) - Import an OpenAPI specification into Assertible to generate tests that validate JSON Schema responses and status codes on every endpoint. | SaaS | ❌ | ✅ | ✅ |  |
| [ Atlassian OpenAPI Request Validators ](https://bitbucket.org/atlassian/swagger-request-validator/src/master/) - A set of Java libraries which allow you to integrate OpenAPI Description Document validation into your testing or clients with tools like WireMock/RestAssured/MockMVC/etc... | Java | ❌ | ✅ | ✅ |  |
| [ BlocklyAutomation ](https://ignatandrei.github.io/BlocklyAutomation/) - Input any OpenAPI document to have generated Blocks in Blockly form to test and generate documentation. | Javascript / .NET | ✅ | ✅ | ❌ |  |
| [ Chai OpenAPI Response Validator ](https://github.com/openapi-library/OpenAPIValidators/tree/master/packages/chai-openapi-response-validator) - Simple Chai support for asserting that HTTP responses satisfy an OpenAPI spec. | Node.js | ❌ | ✅ | ✅ |  |
| [ Counterfact ](https://counterfact.dev/) - A new (July 2022) project that converts an OpenAPI document to a full implementation that runs on ts-node. The idea is to replace the auto-generated code that returns random values with realistic code, one path at a time. If the spec is updated, you can regenerate the types and let the type checker show you which parts of the implementation need to be updated. The types can also be used in client-side code. | TypeScript / Node | ✅ | ✅ | ❌ |  |
| [ Dredd ](http://dredd.io/) - Language-agnostic command-line tool for validating API description document against backend implementation of the API | Javascript | ❌ | ✅ | ✅ |  |
| [ EvoMaster ](https://github.com/EMResearch/EvoMaster) - A tool for automatically generating system-level test cases for RESTful APIs, using Evolutionary Algorithms and Dynamic Program Analysis. | Java/Kotlin | ❌ | ✅ | ✅ |  |
| [ hikaku ](https://github.com/codecentric/hikaku) - A library that tests if the implementation of a REST-API meets its specification. | Kotlin | ❌ | ✅ | ❌ |  |
| [ jest-openapi ](https://github.com/openapi-library/OpenAPIValidators/tree/master/packages/jest-openapi) - Additional Jest matchers for asserting that HTTP responses satisfy an OpenAPI spec. | Node.js | ❌ | ✅ | ✅ |  |
| [ Karate-IDE ](https://marketplace.visualstudio.com/items?itemName=KarateIDE.karate-ide) - Generates KarateDSL Tests and Mocks from OpenAPI 3.0 documents and so you can quickly test/explore your API. | VSCode Extension | ❌ | ✅ | ❌ |  |
| [ Mayhem for API ](https://forallsecure.com/mayhem-for-api) - Probe your REST API with an infinite stream of test cases generated automatically from your OpenAPI specification. | Any | ✅ | ✅ | ✅ |  |
| [ Meeshkan ](https://meeshkan.com/) - Meeshkan is an automated testing and mocking tool. It offers first-class support for GraphQL APIs, but Meeshkan is also built to handle REST APIs and third-party dependencies. | SaaS | ❌ | ✅ | ❌ |  |
| [ Microcks ](https://microcks.io/) - Kubernetes native tool for API Mocking and Testing. Turn your OAI contract examples into ready to use mocks. Use examples to test and validate implementations according spec and schema elements. | Self-hosted / SaaS | ✅ | ✅ | ✅ |  |
| [ OpenAPI Enforcer ](https://www.npmjs.com/package/openapi-enforcer) - Validate your OpenAPI document, serialize, deserialize, and validate incoming requests and outgoing responses, and simplify response building. You can even produce mock data. | Node.js | ❌ | ✅ | ✅ |  |
| [ openapi-dev-tool ](https://github.com/lyra/openapi-dev-tool) - OpenAPI Dev Tool proposes to developers a unique tool to address development and industrialization needs! | JavaScript | ❌ | ✅ | ❌ |  |
| [ portman ](http://getportman.com/) - Port OpenAPI Spec to Postman Collection, with contract & variation tests included! | Node.js | ❌ | ✅ | ❌ |  |
| [ ReadyAPI ](https://smartbear.com/product/ready-api/overview/) - an end to end API functional, security, performance and virtualization tool where OAS description documents can be utilized to automate the creation and validation of end to end tests, running them manually or at any point in your CI/CD pipeline. pipelines. | Java | ❌ | ✅ | ✅ |  |
| [ RESTest ](https://github.com/isa-group/RESTest) - RESTest is a framework for automated black-box testing of RESTful web APIs. It follows a model-based approach, where test cases are automatically derived from the OpenAPI description document (OAS) of the API under test. | Java | ❌ | ✅ | ✅ |  |
| [ Restish ](https://rest.sh/) - A CLI for REST-ish APIs with HTTP/2, built-in auth, content negotiation, caching, and more that understands and can discover OpenAPI descriptions. | CLI / Go | ❌ | ✅ | ❌ |  |
| [ RESTler ](https://github.com/microsoft/restler) - RESTler is the first stateful REST API fuzzing tool for automatically testing cloud services through their REST APIs and finding security and reliability bugs in these services. RESTler analyzes the OpenAPI description of a cloud service, and then generates and executes tests that exercise the service through its REST API. During testing, it checks for specific classes of bugs and dynamically learns how the service behaves from prior service responses. | Any | ❌ | ✅ | ✅ |  |
| [ Schemathesis ](https://github.com/kiwicom/schemathesis) - Reads the description document and generates test cases that will ensure that your application is compliant with its description. | Python | ❌ | ✅ | ✅ |  |
| [ Spectator ](https://github.com/hotmeteor/spectator) - Spectator provides light-weight OpenAPI testing tools you can use within your existing Laravel test suite. | PHP | [👷](https://github.com/hotmeteor/spectator/issues/100) | ✅ | ❌ |  |
| [ Swagger Inspector ](https://inspector.swagger.io/) - Swagger Inspector is a free online tool to quickly execute any API request, validate its responses and generate a corresponding OpenAPI Description. | Self-hosted/SaaS | ❌ | ✅ | ✅ |  |
| [ Tcases for OpenAPI ](https://github.com/Cornutum/tcases/blob/master/tcases-openapi/README.md) - Generates test cases directly from an OpenAPI v3 description of your API. Creates tests executable using various test frameworks. Bonus: Semantic linter reports elements that are inconsistent, superfluous, or dubious. | Java | ❌ | ✅ | ❌ |  |
| [ Unmock ](https://unmock.io/) - API integration testing library that intercepts outgoing requests and serves back mock data based on the OpenAPI descriptions. | Node.js | ❌ | ✅ | ❌ |  |
| [ vREST NG ](https://ng.vrest.io/) - vREST NG is a simple and powerful application for API Automation. It Allows to use OpenAPI specification into vREST NG to drive your API testing that validates the API responses against JSON Schema and also provides powerful response validation capabilities. | JavaScript | ❌ | ✅ | ✅ |  |

## [Gateways](https://openapi.tools/#gateway)

API Gateways and related tools that have integrated support for OpenAPI.

| Name | Language | v3.1 | v3.0 | v2.0 | GitHub |
| --- | --- | --- | --- | --- | --- |
| [ Fusio ](https://www.fusio-project.org/) - Open source API management platform | PHP | ❌ | ✅ | ❌ |  |
| [ Kong Enterprise Edition ](https://konghq.com/kong-enterprise-edition/) - Highly customizable developer portal with developer onboarding, integrated with the Kong API Gateway | Lua | ✅ | ✅ | ✅ |  |
| [ Kusk Gateway ](https://docs.kusk.io/) - Kusk-Gateway is an OpenAPI-driven API Gateway for Kubernetes. It empowers you to develop, validate, mock and deploy your APIs in a matter of minutes using both manual and automated GitOps/APIOps workflows. | Kubernetes | ❌ | ✅ | ✅ |  |
| [ Zuplo (OpenAPI-based gateway and documentation) ](https://www.zuplo.com/) - Zuplo is an API gateway designed for developers. Natively powered by OpenAPI (3.1 or 3.0), zuplo offers an OpenAPI design surface, API documentation and a serverless, programmable edge gateway that includes request validation, auth, rate-limiting and more. | Web / SaaS | ✅ | ✅ | ❌ |  |
