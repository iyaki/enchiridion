---
title: "Common Mistakes in RESTful API Design"
notion_id: 1d954f1c-7d23-81ff-bc4d-e67f4496605b
notion_url: https://app.notion.com/p/Common-Mistakes-in-RESTful-API-Design-1d954f1c7d2381ffbc4de67f4496605b
last_edited: 2026-09-21T16:57:00.000Z
source_url: https://zuplo.com/learning-center/common-pitfalls-in-restful-api-design
tags: ["English", "REST API", "Programming", "Producer (Individual Contributor)", "Article", "Zuplo Blog"]
---
Ever noticed how some APIs are a joy to work with while others make you want to throw your laptop out the window? It's rarely about the technology stack—it's almost always about design choices. Let's be honest—we've all consumed terrible APIs that made us question our career choices. And if we're being really honest, many of us have created these monsters, too.

According to [Postman's 2023 State of the API Report](https://www.postman.com/state-of-api/), nearly 70% of developers say poorly designed APIs directly impact their productivity. The difference between a good and a bad API often comes down to avoiding common mistakes in RESTful API design that create friction for the people who matter most—the developers using your API.

Ready to build APIs that developers actually want to use? Let's dive into the most common pitfalls in RESTful API design and how to avoid them.

- [Common Mistakes in RESTful API Design](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#common-mistakes-in-restful-api-design)
- [Thinking Inside-Out Vs. Outside-In](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#thinking-inside-out-vs-outside-in)
- [Improper URI Definition](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#improper-uri-definition)
- [Improper Use of HTTP Methods](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#improper-use-of-http-methods)
- [Ignoring Error Handling](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#ignoring-error-handling)
- [Lack of Versioning](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#lack-of-versioning)
- [Overcomplicating Responses](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#overcomplicating-responses)
- [Security Oversights](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#security-oversights)
- [Documentation Deficiencies](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#documentation-deficiencies)
- [RESTful API Design Solutions and Best Practices](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#restful-api-design-solutions-and-best-practices)
- [Real-World Examples and Case Studies](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#real-world-examples-and-case-studies)
- [Build Better APIs Today (Not Tomorrow)](https://zuplo.com/blog/2025/03/12/common-pitfalls-in-restful-api-design#build-better-apis-today-not-tomorrow)

## **Common Mistakes in RESTful API Design**

![image](https://cdn.zuplo.com/cdn-cgi/image/fit=contain,width=2560,format=auto/www/media/posts/2025-03-12-common-mistakes-in-restful-api/RESTful%20API%20image%201.png)

graphic of figure at desktop computer

REST has dominated API design for good reason—when done right, it creates intuitive, scalable interfaces that leverage the web's architecture.

### **Understanding RESTful API Design and Its Core Principles**

REST (Representational State Transfer) isn't a standard but an architectural style introduced by Roy Fielding in his [2000 doctoral dissertation](https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm). It's built around key constraints: client-server separation, statelessness, cacheability, uniform interface, layered system, and code-on-demand.

These principles weren't created to torture developers. They evolved from proven patterns that enable scalable, flexible systems that can evolve independently on either side of the API boundary. Mastering these concepts is key to [understanding API basics](https://zuplo.com/blog/2024/09/25/mastering-api-definitions) and helps you avoid common mistakes in RESTful API design and make smart decisions about when to follow them strictly and when flexibility makes sense.

### **The Importance of REST in API Design**

Despite newer alternatives like GraphQL and gRPC, industry surveys consistently show REST remains the dominant approach for web APIs. This staying power stems from REST's alignment with HTTP, its simplicity, and widespread developer familiarity.

REST's core strength is how it maps cleanly to HTTP's semantics, making APIs intuitive for developers who understand web fundamentals. This matters because [enhancing developer experience](https://zuplo.com/blog/2024/02/13/rickdiculous-dev-experience-for-apis) often determines whether your API thrives or dies in the marketplace. Understanding the differences between [REST vs GraphQL](https://zuplo.com/blog/2023/05/10/graphql-vs-rest-the-right-api-design-for-your-audience) can help you make informed decisions about your API design.

## **Thinking Inside-Out Vs. Outside-In**

The most fundamental mistake in RESTful API design happens before writing a single line of code: designing from the wrong perspective.

### **The Leaky Abstraction Problem**

Inside-out API design starts with your internal systems and data models and then exposes them directly through your API. It feels natural because it requires less transformation work, but it creates a leaky abstraction that forces consumers to understand your internal structure.

Consider these two endpoints for a library API:

**Inside-out approach (problematic):**

```plain text
GET /api/database/tables/book_inventory/records?status=1
```

**Outside-in approach (better):**

```plain text
GET /api/books?available=true
```

The first example exposes database implementation details that should be hidden from API consumers. The second focuses on what developers actually care about—finding available books.

### **Learning from the Best**

[Stripe's API documentation](https://stripe.com/docs/api) exemplifies outside-in design. They obsessively focus on developer experience, evidenced by their investment in extensive language-specific libraries and interactive documentation. This approach has paid off—developers everywhere cite Stripe as having one of the best API experiences, directly contributing to their market dominance despite fierce competition.

An outside-in approach means thinking like an API consumer. Ask yourself: "If I knew nothing about our internal systems, what would make the most sense to me?" This shift prevents you from creating APIs that force consumers to learn your internal domain language and structure, helping you avoid common mistakes in RESTful API design.

## **Improper URI Definition**

Your API's URI structure provides the first impression of your API's quality and usability.

### **Avoiding the Verbs-in-URLs Trap**

Many developers create URIs that reflect their storage mechanisms rather than meaningful resources. Consider these examples:

**Poor practice:**

```plain text
GET /api/getUsers
POST /api/createOrder
PUT /api/updateProduct/123
DELETE /api/deleteCustomer/456

```

**Better practice:**

```plain text
GET /api/users
POST /api/orders
PUT /api/products/123
DELETE /api/customers/456

```

The first approach uses verbs in endpoints and mixes naming conventions. The second correctly uses nouns to represent resources and lets HTTP methods convey the action, avoiding common mistakes in RESTful API design.

### **Simplifying Resource Hierarchy Complexity**

Another common mistake is creating overly complex resource hierarchies:

**Overly complex (problematic):**

```plain text
GET /api/companies/456/departments/2/employees/123/projects

```

**More flexible (better):**

```plain text
GET /api/employees/123/projects
GET /api/projects?employeeId=123

```

According to [Microsoft's API design guidelines](https://github.com/microsoft/api-guidelines), excessive nesting makes APIs brittle and difficult to evolve. The second approach provides multiple ways to access the same data depending on the consumer's needs, improving flexibility without sacrificing clarity.

### **Avoiding Query Parameter Pitfalls**

Query parameters also deserve careful thought. They should be used for filtering, sorting, and pagination—not for identifying resources:

**Incorrect:**

```plain text
GET /api/products?id=123

```

**Correct:**

```plain text
GET /api/products/123

```

The distinction might seem minor, but clarity in resource identification forms the foundation of intuitive API design.

## **Improper Use of HTTP Methods**

HTTP methods provide semantic meaning to your API operations. Using them incorrectly creates confusion and violates predictability principles.

### **The POST-for-Everything Problem**

The core HTTP methods in REST map to CRUD operations:

- GET: Retrieve resources (safe, idempotent)
- POST: Create resources
- PUT: Update resources (idempotent)
- DELETE: Remove resources (idempotent)
- PATCH: Partially update resources

A common mistake is using POST for everything. This anti-pattern forces developers to guess how your API works rather than following standard conventions.

Consider these examples:

**Problematic approach:**

```plain text
POST /api/users/search
POST /api/products/delete/123
POST /api/orders/123/cancel

```

**Better approach:**

```plain text
GET /api/users?name=Smith
DELETE /api/products/123
POST /api/orders/123/cancellations

```

The first approach overloads POST for search, deletion, and business operations. The second uses GET with query parameters for searching, DELETE for resource removal, and models the cancellation as a new resource creation.

### **Modeling Business Operations Properly**

For non-CRUD operations, consider modeling them as subresources. Instead of:

```plain text
POST /api/invoices/123/sendEmail

```

A more RESTful approach would be:

```plain text
POST /api/invoices/123/emailDeliveries

```

This approach maintains REST principles while supporting business operations beyond simple CRUD.

![image](https://zuplo.com/_next/image?url=%2F_next%2Fstatic%2Fmedia%2Ftweet.6fb8aa4c.png&w=992&q=75&dpl=dpl_Bpm7nr3v6yMKgtdQ2jhE9vZxZmJJ)

Over 10,000 developers trust Zuplo to secure, document, and monetize their APIs

[Learn More](https://zuplo.com/?utm_source=blog&utm_medium=inline-cta)

## **Ignoring Error Handling**

Nothing frustrates API consumers more than unhelpful error messages. Yet comprehensive error handling is often treated as an afterthought.

### **Bridging the Error Experience Gap**

Unclear error responses rank among developers' top API frustrations, and small wonder.

A poor error response might look like:

```plain text
{
  "error": "An error occurred"
}
```

This provides no actionable information. A better approach:

```plain text
{
  "status": 400,
  "code": "INVALID_PARAMETER",
  "message": "The email address is improperly formatted",
  "details": [
    {
      "field": "email",
      "message": "Must be a valid email address",
      "value": "not-an-email"
    }
  ],
  "help_url": "https://api.example.com/docs/errors/INVALID_PARAMETER"
}
```

The improved response includes:

- HTTP status code
- Machine-readable error code
- Human-readable message
- Detailed validation errors
- Link to documentation

Providing [consistent error responses](https://zuplo.com/blog/2023/04/11/the-power-of-problem-details) helps developers understand and handle errors appropriately.

### **Avoiding Status Code Confusion**

Status codes should follow HTTP conventions. According to [HTTP specifications](https://www.rfc-editor.org/rfc/rfc7231):

- 2xx for success
- 3xx for redirection
- 4xx for client errors
- 5xx for server errors

A common mistake is returning 200 OK with error information in the response body. This breaks client expectations and automated tooling that relies on status codes.

### **Differentiating Error Types**

Another oversight is not distinguishing between different error types. For example, differentiating between validation errors (400), authentication failures (401), authorization issues (403), and resource not found (404) helps clients respond appropriately.

## **Lack of Versioning**

APIs evolve. Without a versioning strategy, you'll eventually break client applications.

### **Understanding the Breaking Change Problem**

Even a seemingly minor change can break thousands of integrations—unless, that is, you have proper versioning in place.

### **Implementing Versioning Approaches**

Common versioning approaches include:

**URL versioning:**

```plain text
GET /api/v1/users

```

**Header versioning:**

```plain text
GET /api/users
Accept: application/vnd.company.api+json;version=1

```

**Query parameter versioning:**

```plain text
GET /api/users?version=1

```

Each approach has trade-offs. Understanding different [API versioning strategies](https://zuplo.com/blog/tags/API-Versioning) can help you choose the one that best fits your needs. Whatever approach you choose, understanding [API versioning best practices](https://zuplo.com/blog/2022/05/17/how-to-version-an-api) is essential to having a strategy before you need it.

### **Planning for Evolution**

Whatever approach you choose, the key is having a strategy before you need it. Part of managing the API lifecycle effectively includes understanding tools like the [HTTP deprecation management](https://zuplo.com/blog/2024/10/25/http-deprecation-header) for deprecating outdated endpoints.

Avoid these versioning mistakes:

- Creating new versions for every small change
- Not documenting what changed between versions
- Abandoning old versions without migration paths, making it difficult for clients to [upgrade API versions](https://zuplo.com/blog/2022/05/25/how-to-get-clients-to-move-off-old-version-of-api)
- Using semantic versioning incorrectly (minor updates that break compatibility)

Remember that versioning is about managing change without breaking integrations that depend on your API.

## **Overcomplicating Responses**

The temptation to return everything about a resource in every response leads to bloated payloads that slow down applications and complicate integration.

### **Avoiding the Kitchen Sink Response**

Consider this overly complex response:

```plain text
{
  "user": {
    "id": 123,
    "username": "jsmith",
    "email": "jsmith@example.com",
    "first_name": "John",
    "last_name": "Smith",
    "address": {
      "street": "123 Main St",
      "city": "Boston",
      "state": "MA",
      "zip": "02101",
      "country": "USA"
    },
    "phone_numbers": [
      {
        "type": "home",
        "number": "555-1234"
      },
      {
        "type": "mobile",
        "number": "555-5678"
      }
    ],
    "orders": [
      {
        "id": 1001,
        "date": "2023-01-15",
        "total": 59.99,
        "items": [
          {
            "product_id": 501,
            "name": "Widget A",
            "price": 29.99,
            "quantity": 1
          },
          {
            "product_id": 502,
            "name": "Widget B",
            "price": 15.0,
            "quantity": 2
          }
        ]
      }
      // Many more orders...
    ],
    "created_at": "2021-03-10T15:00:00Z",
    "updated_at": "2023-06-22T09:30:00Z",
    "last_login": "2023-06-22T08:45:00Z",
    "preferences": {
      "theme": "dark",
      "notifications": true,
      "language": "en-US"
    }
  }
}
```

This response includes everything about the user, whether the client needs it or not.

### **Providing a Focused Alternative**

A simplified version:

```plain text
{
  "id": 123,
  "username": "jsmith",
  "email": "jsmith@example.com",
  "name": "John Smith"
}
```

### **Employing Response Design Strategies**

To solve this problem:

- Include only essential fields by default
- Use query parameters for field selection (e.g., `fields=id,name,email`)
- Provide separate endpoints for related collections
- Consider pagination for large collections

Remember that every field you include has a maintenance cost—you'll need to support it in the future.

## **Security Oversights**

Security mistakes in RESTful API design can lead to data breaches and compliance violations. Here’s how to avoid the most common ones.

### **Addressing Authentication Weaknesses**

Frequently occurring security mistakes include:

**Inadequate Authentication:**

Using basic authentication over HTTP instead of HTTPS, or not implementing token expiration. Ideally, all APIs should enforce HTTPS and use OAuth 2.0 or similar [API authentication methods](https://zuplo.com/blog/2024/07/19/api-authentication) rather than relying solely on [API key authentication](https://zuplo.com/blog/2024/03/07/rebuttal-api-keys-can-do-everything) or [HTTP Basic Authentication](https://zuplo.com/blog/2024/07/31/simple-api-authentication).

### **Closing Authorization Gaps**

**Missing Authorization:**

Authenticating users but not checking if they're authorized to access specific resources.

### **Mitigating Data Exposure Risks**

**Information Leakage:**

Returning sensitive data in error messages or exposing internal system details. The [Facebook Cambridge Analytica scandal](https://www.nytimes.com/2018/04/04/us/politics/cambridge-analytica-scandal-fallout.html) demonstrated how excessive data exposure through APIs can have catastrophic consequences.

### **Implementing Rate Limiting Properly**

**Rate Limiting Absence:**

Not implementing rate limiting makes your API vulnerable to denial-of-service attacks and abuse. Implementing rate limiting also means properly handling and communicating rate limit exceedances, such as using [HTTP 429 Too Many Requests](https://zuplo.com/blog/2024/10/08/http-429-too-many-requests-guide) responses. [GitHub's rate limiting approach](https://docs.github.com/en/rest/overview/resources-in-the-rest-api#rate-limiting) provides an excellent example of transparent limits with clear headers showing usage.

### **Adopting Security Best Practices**

To improve security:

- Implement proper authentication with time-limited tokens
- Apply authorization at both endpoint and object levels
- Return only necessary data to minimize exposure
- Use rate limiting with clear feedback to clients
- Validate all input, no matter how trusted the source seems
- Log security events for audit trails
- Consider using API gateways that handle common security patterns

Security isn't a feature—it's a requirement. Treating it as an add-on rather than a core design principle leaves your API vulnerable. Adhering to [API security best practices](https://zuplo.com/blog/2025/01/31/api-security-best-practices) can help mitigate these risks.

## **Documentation Deficiencies**

Even a perfectly designed API will fail if developers can't figure out how to use it. Document everything, and you’re more than halfway to success.

### **Recognizing Common Documentation Failures**

Common documentation mistakes include:

- Outdated examples that no longer work
- Missing authentication instructions
- No explanation of error responses
- Incomplete endpoint listings
- Lack of code samples in multiple languages

### **Striving for Documentation Excellence**

[Stripe's documentation](https://stripe.com/docs/api) sets the gold standard with clear explanations, interactive examples, and language-specific code samples. Their approach demonstrates how documentation is a product feature, not a technical afterthought.

### **Including Documentation Essentials**

Effective API documentation includes:

- Clear getting started guide
- Authentication explanation with examples
- Complete reference for all endpoints
- Sample requests and responses
- Error code explanations
- Rate limiting details
- SDKs and client libraries
- Changelog to track updates

### **Utilizing Documentation Tools**

The [OpenAPI Specification](https://www.openapis.org/) (formerly Swagger) has become the industry standard for documenting RESTful APIs. Tools like [Swagger UI](https://swagger.io/tools/swagger-ui/) and [ReDoc](https://github.com/Redocly/redoc) can generate interactive documentation from OpenAPI definitions, reducing the effort to maintain quality documentation.

Remember that documentation is often the first interaction developers have with your API. Its quality directly impacts adoption rates.

## **RESTful API Design Solutions and Best Practices**

With an understanding of common mistakes in RESTful API design, let's look at practical solutions that can elevate your API.

### **Embracing Customer-Centric API Design**

Start with user stories that capture what API consumers need to accomplish. Create a sandbox environment where developers can test the API before production, and conduct usability testing with actual developers.

### **Providing Flexibility Through Filtering**

Rather than creating countless endpoints for different data needs, implement robust filtering. This approach provides flexibility without cluttering your API surface.

For example, instead of:

```plain text
GET /api/new-users
GET /api/active-users
GET /api/premium-users

```

Offer a single endpoint with filtering:

```plain text

GET /api/users?status=new
GET /api/users?status=active
GET /api/users?plan=premium

```

This pattern scales better as requirements grow and gives consumers more control.

### **Leveraging Modern Tools for Data Retrieval**

Consider implementing GraphQL alongside REST for complex data retrieval patterns. [GitHub's API v4](https://docs.github.com/en/graphql) demonstrates how GraphQL can complement a REST API for use cases requiring custom field selection and complex relational queries.

For bandwidth-sensitive contexts, support compression and partial responses:

```plain text
GET /api/users/123?fields=id,name,email

```

These approaches maintain RESTful simplicity while addressing performance concerns.

### **Aligning with Industry Standards**

Following established patterns reduces the learning curve for your API. Adopt conventions from [JSON](https://jsonapi.org/)

or [Microsoft API Guidelines](https://github.com/microsoft/api-guidelines) to provide consistent experiences.

These standards cover everything from pagination formats to error handling, saving you from reinventing solutions to common problems.

## **Real-World Examples and Case Studies**

Learning from successful APIs can provide valuable insights for your own designs.

### **GitHub API**

GitHub's API evolution offers valuable lessons in scaling and versioning. Their [v3 REST API](https://docs.github.com/en/rest) supports everything from simple repository operations to complex workflows.

Key lessons from GitHub:

- Consistent resource naming (repositories, issues, pull requests)
- Clear pagination with Link headers
- Comprehensive webhook integration
- Transparent rate limiting with status headers

### **Stripe API**

The [Stripe API](https://stripe.com/docs/api)’s attention to detail extends from documentation to error messages

Elements worth emulating:

- Idempotency keys for safe retries
- Expandable fields for related resources
- Consistent timestamp formats
- Backward compatibility guarantees

### **Shopify API**

[Shopify's REST API](https://shopify.dev/api/admin-rest) supports a massive ecosystem of apps and integrations. Their approach to resource design and bulk operations solves real-world scaling challenges.

Notable patterns:

- Bulk operations for performance
- Webhooks for event-driven architectures
- Scoped authentication tokens
- Metafields for extensibility

These successful APIs share common characteristics: consistent patterns, excellent documentation, robust error handling, and a focus on developer experience above all else.

## **Build Better APIs Today (Not Tomorrow)**

Creating great APIs isn't about knowing the most advanced technologies—it's about avoiding friction. Utilizing tools like a [hosted API gateway](https://zuplo.com/blog/2024/12/16/hosted-api-gateway-advantages), designing your API from the outside in, focusing on the developer experience, using HTTP methods correctly, creating intuitive URLs, building comprehensive error handling, implementing versioning early, keeping responses focused, and making security foundational.

The most successful APIs feel intuitive, behave predictably, and solve real problems for developers. By avoiding common mistakes in RESTful API design, you'll create APIs that developers actually want to use—and that's the ultimate measure of API success. Utilizing [API integration platforms](https://zuplo.com/blog/2024/11/08/building-an-api-integration-platform) can further enhance your API's capabilities. Want to build enterprise-grade API infrastructure without the complexity? [Check out Zuplo](https://portal.zuplo.com/signup?utm_source=blog))—we deploy your policies across 300 data centers worldwide in less than 5 seconds, ensuring your API is both secure and lightning-fast.

[#API Best Practices](https://zuplo.com/blog/tags/api-best-practices)
