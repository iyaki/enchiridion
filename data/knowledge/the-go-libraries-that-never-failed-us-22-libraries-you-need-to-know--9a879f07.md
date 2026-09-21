---
title: "The Go libraries that never failed us: 22 libraries you need to know"
notion_id: 9a879f07-b487-4bd1-8e97-820c3142b176
notion_url: https://app.notion.com/p/The-Go-libraries-that-never-failed-us-22-libraries-you-need-to-know-9a879f07b4874bd18e97820c3142b176
last_edited: 2023-07-10T18:28:00.000Z
source_url: https://threedots.tech/post/list-of-recommended-libraries/
tags: ["Three Dots Labs", "English", "Go", "Article", "Framework/Library"]
---
Did you have a situation when you lost a ton of time finding a Go library for your need? In theory, you can check lists like [Awesome Go](https://github.com/avelino/awesome-go) or make a choice based on GitHub stars. But Awesome Go contains over 2600 libraries, and popularity is not always the best indicator of library quality. **I often thought that it would be great to have a place where I could find just the best and battle-tested libraries I could use in my project.** Because we didn’t find such a place with Miłosz, we decided to create it.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

Based on our experience leading multiple Go teams and working on various projects, including complex financial, health, and security, we will recommend tools that could work well for different projects.

In addition to providing a list of libraries, we also want to show you some non-obvious uses for those tools and libraries. However, it’s important to note that most of these tools can be misused. We’ve included some common anti-patterns to help you avoid making those mistakes.

## Table Of Contents

- [HTTP](https://threedots.tech/post/list-of-recommended-libraries/#http) 
- [Routers](https://threedots.tech/post/list-of-recommended-libraries/#routers)
- [Middlewares](https://threedots.tech/post/list-of-recommended-libraries/#middlewares)
- [Serving static content](https://threedots.tech/post/list-of-recommended-libraries/#serving-static-content)
- [OpenAPI](https://threedots.tech/post/list-of-recommended-libraries/#openapi)
- [Generating Go server and clients](https://threedots.tech/post/list-of-recommended-libraries/#generating-go-server-and-clients)
- [Bonus: Client for JavaScript/TypeScript](https://threedots.tech/post/list-of-recommended-libraries/#bonus-client-for-javascripttypescript)
- [Alternative types of communication](https://threedots.tech/post/list-of-recommended-libraries/#alternative-types-of-communication) 
- [gRPC](https://threedots.tech/post/list-of-recommended-libraries/#grpc)
- [Messaging](https://threedots.tech/post/list-of-recommended-libraries/#messaging)
- [Database](https://threedots.tech/post/list-of-recommended-libraries/#database) 
- [SQL](https://threedots.tech/post/list-of-recommended-libraries/#sql)
- [Migrations](https://threedots.tech/post/list-of-recommended-libraries/#migrations)
- [Observability](https://threedots.tech/post/list-of-recommended-libraries/#observability) 
- [Logging](https://threedots.tech/post/list-of-recommended-libraries/#logging)
- [Metrics and tracing](https://threedots.tech/post/list-of-recommended-libraries/#metrics-and-tracing)
- [Configuration](https://threedots.tech/post/list-of-recommended-libraries/#configuration) 
- [Env variables](https://threedots.tech/post/list-of-recommended-libraries/#env-variables)
- [Building CLI](https://threedots.tech/post/list-of-recommended-libraries/#building-cli) 
- [Building CLI libraries](https://threedots.tech/post/list-of-recommended-libraries/#building-cli-libraries)
- [Testing](https://threedots.tech/post/list-of-recommended-libraries/#testing) 
- [Assertions](https://threedots.tech/post/list-of-recommended-libraries/#assertions)
- [Mocking](https://threedots.tech/post/list-of-recommended-libraries/#mocking)
- [Misc](https://threedots.tech/post/list-of-recommended-libraries/#misc) 
- [Extra types support](https://threedots.tech/post/list-of-recommended-libraries/#extra-types-support)
- [Errors](https://threedots.tech/post/list-of-recommended-libraries/#errors)
- [Useful tools](https://threedots.tech/post/list-of-recommended-libraries/#useful-tools) 
- [Misc](https://threedots.tech/post/list-of-recommended-libraries/#misc-1)
- [Live code reloading](https://threedots.tech/post/list-of-recommended-libraries/#live-code-reloading)
- [Linter](https://threedots.tech/post/list-of-recommended-libraries/#linter)
- [Formatters](https://threedots.tech/post/list-of-recommended-libraries/#formatters)
- [Example projects](https://threedots.tech/post/list-of-recommended-libraries/#example-projects) 
- [DDD & Clean Architecture](https://threedots.tech/post/list-of-recommended-libraries/#ddd--clean-architecture)
- [General purpose](https://threedots.tech/post/list-of-recommended-libraries/#general-purpose)
- [Summary](https://threedots.tech/post/list-of-recommended-libraries/#summary)

## HTTP

### Routers

As I mentioned in my [previous article](https://threedots.tech/post/best-go-framework/), it’s generally better to use libraries instead of frameworks for long-term projects. One of the most fundamental components of any service is an HTTP router. While it’s technically possible to build an application without one by using the standard library’s [http](https://pkg.go.dev/net/http) package, its routing capabilities are limited. Using a dedicated router will make your life much easier.

By design, router functionality is limited to routing requests to a proper handler. All non-standard functionalities like CORS, CSRF, error handling, HTTP logging, and authorization (that frameworks usually provide) are provided by reusable middlewares. I recommend some in the [section on middlewares](https://threedots.tech/post/list-of-recommended-libraries/#middlewares).

I use one of two router libraries in most projects: Echo or chi. Both of them are great routers, with different characteristics. They work perfectly with [OpenAPI](https://threedots.tech/post/list-of-recommended-libraries/#openapi) code generation.

### ✅ Echo [[GitHub]](https://github.com/labstack/echo) [[Docs]](https://echo.labstack.com/guide/) [[Examples]](https://echo.labstack.com/cookbook/)

Compared to chi, Echo does offer a custom `*http.Request` handler signature. Some people may find it a downside, but I think it helps to write less error-prone HTTP handlers.

If you have been writing Go for a while, you probably made this mistake at least once:

```plain text
func someHandler(w http.ResponseWriter, r *http.Request) {
	err := foo()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// you forgot the return here, bar() will be executed
	}

	bar()
}

```

Echo makes you return an error:

```plain text
func someHandler(c echo.Context) error {
	err := foo()
	if err != nil {
		return err
	}

	bar()

	return c.NoContent(http.StatusNoContent)
}

```

The advantage of Echo is the ability to define a custom [error handler](https://echo.labstack.com/guide/error-handling/). It’s not possible to do it in the same way with chi.

For detailed usage and examples, please check Echo docs.

### ✅ chi [[GitHub]](https://github.com/go-chi/chi) [[Docs]](https://pkg.go.dev/github.com/go-chi/chi) [[Examples]](https://github.com/go-chi/chi/tree/master/_examples)

Compared to Echo, chi’s handler functions are compatible with the standard library. For some people, it may be an upside; for some, it may be a downside – you should make your own judgment.

What chi does better than Echo is the format of [defining routes and grouping](https://github.com/go-chi/chi/blob/0fe6bf1ba3ac601700b7993bc4c62f6c5f707932/_examples/rest/main.go#L83). It gives you better control over middleware per path or sub-path.

```plain text
r.Route("/articles", func(r chi.Router) {
	r.With(paginate).Get("/", ListArticles)
	r.Post("/", CreateArticle)       // POST /articles
	r.Get("/search", SearchArticles) // GET /articles/search

	r.Route("/{articleID}", func(r chi.Router) {
		r.Use(ArticleCtx)            // Load the *Article on the request context
		r.Get("/", GetArticle)       // GET /articles/123
		r.Put("/", UpdateArticle)    // PUT /articles/123
		r.Delete("/", DeleteArticle) // DELETE /articles/123
	})

	// GET /articles/whats-up
	r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
})

```

### Middlewares

HTTP middlewares can give you functionalities like CORS, CSRF, error handling, HTTP logging, authorization, etc.

Echo and chi provide their set of middlewares:

- [Echo middlewares](https://echo.labstack.com/middleware/)
- [chi middlewares](https://github.com/go-chi/chi/tree/master/middleware)

Echo middlewares have a different interface, so they can’t be used in chi. Generally speaking, all standard-library compatible middlewares are compatible with chi and echo.

To use standard library-compatible middleware with echo, you need to call `echo.WrapMiddleware`:

If none of them provides the middleware you are looking for, you can check [the Awesome Go list](https://github.com/avelino/awesome-go#middlewares). All of them will be compatible with chi, Echo, and servers built just with the standard library. You can also write your own middleware. Check example middlewares for inspiration!

### Serving static content

You don’t need any library to serve static content in Go. Since Go 1.16, you can easily [embed static files into your Go binary](https://pkg.go.dev/embed).

Here’s how to do it for `Echo` and `chi`:

After running the server, assets will be available under `http://localhost:8080/static/index.html`, `http://localhost:8080/static/main.js` etc.

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### OpenAPI

Nobody likes to maintain API contracts manually. It’s annoying and counterproductive to keep multiple boring JSON’s up-to-date. OpenAPI solves this problem with a JavaScript HTTP client and Go HTTP server generated from the provided specification.

This is how an [example specification](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/tree/a0a41253db96d46d75e7ff4c7e7f95848f47dcc3/api/openapi) looks like. If you didn’t work with OpenAPI before, you can read more details in my [previous article](https://threedots.tech/post/serverless-cloud-run-firebase-modern-go-application/#openapi-swagger-client). Here, I focus on tools that we recommend for OpenAPI spec-generated code.

### Generating Go server and clients

We do not recommend using the official OpenAPI generator for the Go code. We recommend the `oapi-codegen` tool instead because of the higher quality of the generated code. It also has more functionalities.

### ✅ deepmap/oapi-codegen [[GitHub]](https://github.com/deepmap/oapi-codegen) [[Docs]](https://github.com/deepmap/oapi-codegen#readme) [[Example]](https://threedots.tech/post/serverless-cloud-run-firebase-modern-go-application/#public-http-api)

`oapi-codegen` is a great tool that doesn’t just generate models but also the entire [router](https://threedots.tech/post/list-of-recommended-libraries/#routers) definition, header’s validation, and proper parameters parsing. It works with [chi](https://threedots.tech/post/list-of-recommended-libraries/#chi) and [Echo](https://threedots.tech/post/list-of-recommended-libraries/#echo).

To generate a server, run the following:

Where `<TYPE>` for `chi` should be `chi-server`, and for `Echo` just `server`.

To generate clients:

Don’t forget to change `<GO PACKAGE>` to the desired Go package name and `<OUTPUT DIR>` to the desired output dir. 😉

Your job on the server side is just to implement the `ServerInterface` interface, like:

You can see it in action in the [Wild Workouts project](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example).

### Bonus: Client for JavaScript/TypeScript

Even if it’s the list of **recommended Go libraries**, you may need to generate code for the browser.

### ✅ openapi-generator-cli [[GitHub]](https://github.com/OpenAPITools/openapi-generator-cli) [[Docs]](https://github.com/OpenAPITools/openapi-generator-cli#readme)

In that case, we also recommend a non-official library instead of the official one.

In contrast to `oapi-codegen`, `openapi-generator-cli` is a Java tool. To avoid any JVM-related issues, we recommend generating clients using Docker:

It assumes that the spec is available locally in `api/openapi/service.yml`.

You can use `openapi-generator-cli` for TypeScript and other languages as well.

## Alternative types of communication

### gRPC

gRPC is a technology that can help you with building robust, internal communication between your services (but not only!).

I already described in detail why it’s [worth using gRPC for internal communication](https://threedots.tech/post/robust-grpc-google-cloud-run/) and how to do it.

I’ll not repeat it here and will focus on the tooling you need.

With gRPC, you have little choice for generating server and client: you should use official tooling. The good news is that you don’t need anything more because it does its job!

### ✅ protoc [[Docs]](https://grpc.io/docs/)

To generate Go code from `.proto` files, you need to install [protoc](https://grpc.io/docs/protoc-installation/) and [protoc Go Plugin](https://grpc.io/docs/quickstart/go/).

A list of supported types can be found in [Protocol Buffers Version 3 Language Specification](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec#fields). More complex built-in types like Timestamp can be found in [Well-Known Types list](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf).

### Messaging

### ✅ Watermill [[GitHub]](https://github.com/ThreeDotsLabs/watermill) [[Docs]](https://watermill.io/) [[Examples]](https://github.com/ThreeDotsLabs/watermill/tree/master/_examples)

About four years ago, when working on one of our projects, we found that there is no library that can simplify building message-driven or event-driven applications easily. To make our lives easier, we decided to write a library that will allow us to write event-driven code as easily as writing HTTP services. This is how Watermill was born.

Today, Watermill is one of the most popular Go libraries with almost 5k GitHub stars, +35 contributors, and 10 officially supported Pub/Subs.

Usually, message-broker libraries are very low-level. With Watermill, publishing messages may be as simple as:

And subscribing like:

Compared to using just the message broker’s library, Watermill provides support for some higher level functionalities like [middlewares](https://watermill.io/docs/middlewares/), [CQRS support](https://watermill.io/docs/cqrs/), or [event-forwarder](https://watermill.io/docs/forwarder/) component (that can be used to stream your messages from an SQL database to the message broker).

Today, Watermill officially supports [Kafka](https://watermill.io/pubsubs/kafka/), [GCP Pub/Sub](https://watermill.io/pubsubs/googlecloud/), [NATS](https://watermill.io/pubsubs/nats/), [RabbitMQ](https://watermill.io/pubsubs/amqp/) message brokers (Pub/Subs). It can also listen to and emit events as [HTTP hooks](https://watermill.io/pubsubs/http/), from databases like [MySQL/Postgres](https://watermill.io/pubsubs/sql/), [BoltDB](https://watermill.io/pubsubs/bolt/) or [Firestore](https://watermill.io/pubsubs/firestore/). It can also work with in-memory [Go-channel based Pub/Sub](https://watermill.io/pubsubs/gochannel/).

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## Database

### SQL

There is no golden hammer solution for interacting with SQL databases. The reason is simple: it depends greatly on what kind of data you store.

In some projects, data models are relatively simple. In some, they are very complex. Because of that, I have two libraries to recommend. You should choose one of them based on the requirements of your project.

For projects with straightforward data models, you should check `sqlx`. For a bit more complex, you should look at `SQLBoiler`.

### ✅ sqlx [[GitHub]](https://github.com/jmoiron/sqlx) [[Docs]](http://jmoiron.github.io/sqlx/)

The standard library’s `database/sql` package is rather a low-level one. `sqlx` provides a more convenient and powerful API to work with databases. It includes helper functions for common tasks like inserting and querying data and support for more advanced features like prepared statements and transactions. `sqlx` also has more advanced support for data unmarshaling (for example to structs, lists of structs, json data, etc.). As a nice bonus, `sqlx`’s interface is compatible with interfaces from `database/sql`.

But even if `sqlx` is a great library, it works well for relatively simple database models. At some level of complexity, you should consider migration to an ORM.

### ✅ SQLBoiler [[GitHub]](https://github.com/volatiletech/sqlboiler) [[Docs]](https://github.com/volatiletech/sqlboiler#table-of-contents) [[Examples]](https://github.com/volatiletech/sqlboiler#features--examples)

So far, the only ORM that fully meets our requirements is SQLBoiler. At first, how you define SQLBoiler models may surprise you. Most ORMs generate the database schema out of your Go models. SQLBoiler does the opposite: it generates Go models from your database schema.

This approach has multiple advantages. One of the most important features is stricter typing than other libraries. Thanks to that, many checks are done during compilation. You don’t need to depend on a ton of reflection and magic struct tags. In most cases, as long as the code compiles, it will work correctly.

Generating code from the database schema helps with migration from an existing database because you don’t need to re-write DB models: SQLBoiler generates them for you. So if you start with `sqlx` and move to SQLBoiler later, the migration should be pretty easy.

SQLBoiler supports PostgreSQL, MySQL, MSSQLServer 2012+, SQLite3, and CockroachDB.

### Migrations

SQLBoiler and `sqlx` don’t provide out-of-the-box support for migrations. It’s okay because you are not forced to use any particular solution.

In my recent projects, I used both `sql-migrate` and `goose`, and I was happy about them.

### ✅ sql-migrate [[GitHub]](https://github.com/rubenv/sql-migrate) [[Docs]](https://github.com/rubenv/sql-migrate#readme)

### ✅ goose [[GitHub]](https://github.com/pressly/goose) [[Docs]](https://pkg.go.dev/github.com/pressly/goose)

We like `sql-migrate` and `goose` because of their simplicity and flexibility. `sql-migrate` and `goose` can be executed as CLI tools and as part of your service.

I like to embed it into the binary of the service. Thanks to that, the migration is executed when the service starts, and it keeps the setup simple. It’s also much less complex to run. For example, `sql-migrate` with `go:embed`:

Put your migrations in `migrations/`, for example: `migrations/1_init.sql`, `migrations/2_alter_some_table.sql`, etc. Then run `Run` in your `main`.

## Observability

### Logging

The standard library’s logger doesn’t provide essential features like log levels and output formatting.

For logging, we can recommend two libraries: `Logrus` and `zap`. In contrast to `zap`, `Logrus` provides a bit nicer user API, but `zap` is faster.

You can check detailed benchmarks in [zap’s readme](https://github.com/uber-go/zap#performance).

### ✅ Logrus [[GitHub]](https://github.com/sirupsen/logrus) [[Docs]](https://pkg.go.dev/github.com/sirupsen/logrus)

### ✅ zap [[GitHub]](https://github.com/uber-go/zap) [[Docs]](http://pkg.go.dev/github.com/uber-go/zap)

### Metrics and tracing

### ✅ opencensus-go [[GitHub]](https://github.com/census-instrumentation/opencensus-go) [[Docs]](https://opencensus.io/)

OpenCensus Go is a library that helps you add metrics and tracing to your endpoints or database queries. The integration uses middleware/decorator patterns, and it doesn’t require a lot of custom code. It supports [HTTP endpoints](https://pkg.go.dev/go.opencensus.io/plugin/ochttp), [gRPC endpoints](https://pkg.go.dev/go.opencensus.io/plugin/ocgrpc), [SQL databases](https://pkg.go.dev/github.com/opencensus-integrations/ocsql), [MongoDB](https://pkg.go.dev/github.com/orijtech/mongo-go-driver), etc.

You can export traces and metrics to Prometheus, OpenZipkin, GCP Stackdriver Monitoring, Jaeger, AWS X-Ray, Datadog, Graphite, Honeycomb, or New Relic.

## Configuration

Go’s standard library doesn’t support much more configuration options than the [flag package](https://pkg.go.dev/flag). Even if it’s enough for simple CLI tools, you may need a bit more for building services.

### Env variables

### ✅ caarlos0/env [[GitHub]](https://github.com/caarlos0/env) [[Docs]](https://pkg.go.dev/github.com/caarlos0/env)

This library should provide everything you need for configuration for most applications. Compared to the standard library, it does support loading envs to structs and setting env defaults. It helps to save a lot of boilerplate for bigger configurations. It also supports embedded structs, so you can compose bigger a configuration from independent components.

### Multi-format configuration

### ✅ koanf [[GitHub]](https://github.com/knadh/koanf) [[Docs]](https://pkg.go.dev/github.com/knadh/koanf)

Koanf is an excellent tool if your project requires multiple configuration formats. It’s often the case when you write tools that are used externally (for example, CLI tools).

This is my most recent finding. Compared to other [more popular libraries](https://github.com/knadh/koanf#alternative-to-viper), `koanf` just does loading multi-format configuration right. Bonus points for a nice abstraction that allows extending parsing and loading.

Koanf does support the most important configuration formats, like `json`, `yaml`, `dotenv`, env vars, or `hcl`. They can be loaded from the filesystem, flags, and multiple external sources like `s3`, `vault`, `etcd`, or `consul`.

## Building CLI

### Building CLI libraries

### ✅ urfave/cli [[GitHub]](https://github.com/urfave/cli/) [[Docs]](https://cli.urfave.org/) [[Examples]](https://cli.urfave.org/v2/examples/greet/)

We like `urfave/cli` because of its simple interface and extensibility. We used it in multiple projects without any issues.

Compared to other alternatives, it offers a big-enough feature set while keeping the library lightweight.

## Testing

### Assertions

### ✅ testify [[GitHub]](https://github.com/stretchr/testify) [[Docs]](https://pkg.go.dev/github.com/stretchr/testify)

`testify` became the standard assertion library, and I’ve seen it in every project I worked on. It provides asserts for the most common cases and also some more complex. One of `testify`’s key features are friendly messages for all failed asserts. It makes writing and debugging tests much faster.

The library provides two ways of asserting:

- `assert` from `github.com/stretchr/testify/assert` - the test continues after failure. You should use it when you want to see multiple errors (not just the first one). Works when called in a goroutine.
- `require` from `github.com/stretchr/testify/require` - the test is interrupted after failure. You should use it when some critical condition was not met and continuing doesn’t make any sense (for example: storing to database failed). Doesn’t work when called in a goroutine.

Some example asserts:

- [Equal](https://pkg.go.dev/github.com/stretchr/testify/assert#Equal) - good enough in most cases
- [Eventually](https://pkg.go.dev/github.com/stretchr/testify/assert#Eventually) - useful for asserting asynchronous conditions
- [ElementsMatch](https://pkg.go.dev/github.com/stretchr/testify/assert#ElementsMatch) - useful for unsorted slices
- [WithinDuration](https://pkg.go.dev/github.com/stretchr/testify/assert#WithinDuration) - useful when comparing time that is not exactly equal
- [ErrorIs](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorIs)
- [JSONEq](https://pkg.go.dev/github.com/stretchr/testify/assert#JSONEq)
- [Panics](https://pkg.go.dev/github.com/stretchr/testify/assert#Panics)

### ✅ go-cmp [[GitHub]](https://github.com/google/go-cmp) [[Docs]](https://pkg.go.dev/github.com/google/go-cmp) [[Examples 1]](https://github.com/google/go-cmp/blob/master/cmp/example_test.go) [[Examples 2]](https://github.com/google/go-cmp/blob/master/cmp/cmpopts/example_test.go)

Sometimes, you must assert a complex struct in your tests skipping some fields. Or the struct contains fields that should be compared in a specific way. Or you need to ignore the slice order or time delta. It’s where `go-cmp` can help you!

To see the list of all available options, I recommend checking the godoc of [`cmp`](https://pkg.go.dev/github.com/google/go-cmp/cmp) and [`cmpopts`](https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts) package.

go-cmp can also be used outside of tests, but be careful – it’s another tool that, used irresponsibly, may hurt your project.

### ✅ gofakeit [[GitHub]](https://github.com/brianvoe/gofakeit) [[Docs]](https://pkg.go.dev/github.com/brianvoe/gofakeit)

If you need more realistic data for your tests, `gofakeit` helps.

### Mocking

### Writing mocks by hand

_Initially, I recommended one popular mocking tool here. But after some thinking, we decided that the tool is not good enough to recommend. Instead, consider an alternative mocking strategy 👇_

## Misc

### Extra types support

### ✅ google/uuid [[GitHub]](https://github.com/google/uuid) [[Docs]](https://pkg.go.dev/github.com/google/uuid)

This library generates UUIDs.

### ✅ oklog/ulid [[GitHub]](https://github.com/oklog/ulid) [[Docs]](https://pkg.go.dev/github.com/oklog/ulid)

UUIDs [may be slow to store](https://www.percona.com/blog/2014/12/19/store-uuid-optimized-way/) at a larger scale in relational databases. A solution may be using Universally Unique Lexicographically Sortable Identifier: ULIDs. ULIDs are compatible with UUIDs, are unique enough for large scale, and have shorter string representation (Crockford’s base32). ULIDs are lexicographically sortable, thanks to what building indexes should be much faster.

### ✅ shopspring/decimal [[GitHub]](https://github.com/shopspring/decimal) [[Docs]](https://pkg.go.dev/github.com/shopspring/decimal)

Go doesn’t have built-in support for decimals. `shopspring/decimal` does the job. We have used this library for a couple of years to build a large financial system.

### Errors

### ✅ hashicorp/go-multierror [[GitHub]](https://github.com/hashicorp/go-multierror) [[Docs]](https://github.com/hashicorp/go-multierror)

Did you ever need to handle an error while you were handling another error? `hashicorp/go-multierror` is here to help you!

It’s also helpful if an operation can return multiple errors, and you don’t want to return just the first one (for example, validation).

Example use cases:

or

_Note: Go 1.20 _[_will introduce_](https://github.com/golang/go/issues/53435)_ __`errors.Join`__ function. After release of Go 1.20 you should consider using it instead._

## Useful tools

### Misc

### ✅ samber/lo [[GitHub]](https://github.com/samber/lo) [[Docs]](https://pkg.go.dev/github.com/samber/lo)

Lodash-style Go library based on Go 1.18+ Generics. It may be especially useful for you if you are coming to Go from Python and missing some basic slice/map functions.

Some functions that I’m using the most:

- [Filter](https://pkg.go.dev/github.com/samber/lo#Filter)
- [Map](https://pkg.go.dev/github.com/samber/lo#Map)
- [Keys](https://pkg.go.dev/github.com/samber/lo#Keys)
- [Values](https://pkg.go.dev/github.com/samber/lo#Values)
- [Find](https://pkg.go.dev/github.com/samber/lo#Find)
- [Max](https://pkg.go.dev/github.com/samber/lo#Max)
- [Must](https://pkg.go.dev/github.com/samber/lo#Must) 😈 please use it just for tests or if you really have a good reason

Even if some may find it “non-idiomatic”, I find it useful in some cases. It’s similar to using an [ORM](https://threedots.tech/post/list-of-recommended-libraries/#sql) – as long as such libraries are used responsibly and don’t obfuscate code, they are useful.

So if you find yourself writing code like:

…it’s just better to convert it to a simple, more readable loop. 😉

### ✅ Task [[GitHub]](https://github.com/go-task/task) [[Docs]](https://taskfile.dev/)

Task is not really a Go library, but it’s a tool written in Go that may be useful for your projects.

It’s an excellent alternative to Makefile. The most important features that it offers are:

- Parallel tasks execution (supported by [task dependencies](https://taskfile.dev/usage/#task-dependencies))
- Preventing [unnecessary work](https://taskfile.dev/usage/#prevent-unnecessary-work)
- [Loading .env](https://taskfile.dev/usage/#env-files)
- [Dynamic variables](https://taskfile.dev/usage/#dynamic-variables)
- [Forwarding CLI arguments](https://taskfile.dev/usage/#forwarding-cli-arguments-to-commands)
- [Templating](https://taskfile.dev/usage/#gos-template-engine)

It’s a must-have for each of my new projects.

### Live code reloading

### ✅ reflex [[GitHub]](https://github.com/cespare/reflex) [[Docs]](https://pkg.go.dev/github.com/cespare/reflex) [[Example](https://threedots.tech/post/go-docker-dev-environment-with-go-modules-and-live-code-reloading/)]

Go doesn’t provide code live-reloading out of the box. But you can achieve it quickly with the `reflex` library.

Some time ago, Miłosz wrote an article that shows how to create a [local environment with Docker and reflex](https://threedots.tech/post/go-docker-dev-environment-with-go-modules-and-live-code-reloading/).

### Linter

### ✅ golangci-lint [[GitHub]](https://github.com/golangci/golangci-lint) [[Docs]](https://golangci-lint.run/)

golangci-lint is a linter that aggregates multiple linters and runs them in parallel and does it very fast.

Here’s [an example configuration](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/blob/b519c611e9d1248a149c89db9bcf879fd78b1e35/internal/trainer/.golangci.yml) that we use in our projects.

### ✅ go-cleanarch [[GitHub]](https://github.com/roblaszczak/go-cleanarch) [[Docs]](https://pkg.go.dev/github.com/roblaszczak/go-cleanarch#section-readme)

If you use [Clean/Hexagonal Architecture](https://threedots.tech/post/introducing-clean-architecture/) in your project, you can use this linter to ensure that The Dependency Inversion Rule and interaction between packages are kept.

### Formatters

### ✅ go fmt

The standard formatter provided by Go toolchain.

### ✅ goimports [[Docs]](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)

Goimports does all that `go fmt` does, but it also sorts imports of your Go files. It’s one of the tools that you will see widely adopted in most Go projects.

Not everybody knows, but you can also separately group your local imports with the `-local` flag.

### ✅ gofumpt [[GitHub]](https://github.com/mvdan/gofumpt) [[Docs]](https://pkg.go.dev/mvdan.cc/gofumpt#section-readme)

Just for the biggest formatting freaks! Does all that `go fmt` and `goimports` do and more!

Personally, I like gofumpt’s formatting decisions.

## Example projects

### DDD & Clean Architecture

### ✅ Wild Workouts Go DDD Example application [[GitHub]](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example)

**Wild Workouts is an example Go DDD project that we created to show how to build Go applications that are easy to develop, maintain, and fun to work with, It shows a project developed over time and with complex problems to solve.** In contrast to other example projects, it was not blindly copied from other languages.

This is the way how we build our services daily. Highly recommended if you are looking for patterns that will allow you to build more complex projects!

### General purpose

### ✅ Modern Go Application by Márk Sági-Kazár [[GitHub]](https://github.com/sagikazarmark/modern-go-application)

Another example repository that we can recommend. It doesn’t cover patterns like DDD or Clean Architecture but emphasizes infrastructure beats like observability.

## Summary

Should we check some library that is not listed here? Please let us know in the [comments](https://threedots.tech/post/list-of-recommended-libraries/#disqus_thread)!
