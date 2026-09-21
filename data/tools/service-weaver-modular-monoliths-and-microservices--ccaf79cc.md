---
title: "Service Weaver - Modular monoliths and microservices"
notion_id: ccaf79cc-31b0-47c1-9b8d-ba5da0587366
notion_url: https://app.notion.com/p/Service-Weaver-Modular-monoliths-and-microservices-ccaf79cc31b047c19b8dba5da0587366
last_edited: 2023-03-02T17:11:00.000Z
source_url: https://serviceweaver.dev/
tags: ["Go", "Framework/Library", "English"]
---
# Write your application as a **modular binary**. Deploy it as a set of _microservices_.

Service Weaver is a programming framework for writing and deploying cloud applications.

[Read the Docs](https://serviceweaver.dev/docs.html)

Split your application into **components** written as regular Go interfaces. Don't fuss with any networking or serialization code. Focus on your business logic.

```plain text
type Adder interface {
    Add(context.Context, int, int) (int, error)
}
type adder struct{
    weaver.Implements[Adder]
}
func (adder) Add(_ context.Context, x, y int) (int, error) {
    return x + y, nil
}

```

# Step 2: Call Your Components

Call your components using regular Go method calls. No need for RPCs or HTTP requests. Forget about versioning issues. The type system guarantees **components are compatible**.

```plain text
adder, err := weaver.Get[Adder](root)
if err != nil {
    panic(err)
}
sum, err := adder.Add(ctx, 1, 2)

```

# Step 3: Deploy Your Components

Test your app locally and deploy it to the cloud. Service Weaver lets you think about **what** your code does without worrying about **where** it's running.

```plain text
$ go test .                       # Test locally.
$ go run .                        # Run in a single process.
$ weaver multi deploy weaver.toml # Run in multiple processes.
$ weaver gke deploy weaver.toml   # Run in the cloud.

```

# Step 4: Place Your Components

Run your components **wherever you want**: in the same process or on different machines. Run **as many replicas as you want**; scale up or down to match load.

Machine 1Machine 2Machine 3

# Features

## [Highly Performant ⚡](https://serviceweaver.dev/docs.html#serializable-types)

Co-located components communicate via **direct method call**. Remote components communicate using highly-efficient custom serialization and RPC protocols.

```plain text
// Automatically encoded and decoded.
type pair struct {
    weaver.AutoMarshal
    x, y int32
}

```

## [Tiny Config 🎛️](https://serviceweaver.dev/docs.html#config-files)

Deploy to the cloud **without tons of boilerplate** configuration. Here's a working config file to deploy a Service Weaver application across two regions in Google Cloud. It's less than ten lines long.

```plain text
[serviceweaver]
binary = "./example"
[gke]
regions = ["us-west1", "us-east1"]
public_listener = [
  {name = "example", hostname = "example.com"},
]

```

## [Logging, Metrics, Tracing 🔎](https://serviceweaver.dev/docs.html#logging)

Service Weaver has libraries for logging, metrics, and tracing. This telemetry is **automatically integrated** into the cloud where you deploy.

```plain text
var count = weaver.NewCounter(
    "example_count",
    "An example of a Service Weaver counter",
)
func main() {
    count.Add(1)
    // ...
}

```

## [Sharding 🗃️](https://serviceweaver.dev/docs.html#routing)

Shard requests across different component replicas.

```plain text
type Cache interface {
    Get(ctx context.Context, key string) (string, error)
    Put(ctx context.Context, key, val string) error
}
// router defines how Cache methods are routed.
type router struct{}
func (router) Get(_ context.Context, k string) string {
    return k
}
func (router) Put(_ context.Context, k, _ string) string {
    return k
}

```
