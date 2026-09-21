---
title: "Migrating from Go to Rust | corrode Rust Consulting"
notion_id: 36d54f1c-7d23-811d-9e5b-ec617f8ae90c
notion_url: https://app.notion.com/p/Migrating-from-Go-to-Rust-corrode-Rust-Consulting-36d54f1c7d23811d9e5bec617f8ae90c
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://corrode.dev/learn/migration-guides/go-to-rust/
tags: ["English", "Programming", "Rust", "Go", "Software Development", "Systems Design / Software Architecture", "Article", "Tutorial", "corrode Rust Consulting"]
---
- Where Go and Rust overlap, and where they diverge.
- How Go patterns map to Rust.
- What you gain from the borrow checker.
- Where I tell people to keep Go and where Rust is worth the migration cost.
- How to migrate Go services incrementally.

![image](https://corrode.dev/learn/migration-guides/go-to-rust/go-usage.svg)

Go and Rust usage among developers, 2017–2024. Go holds steady around 17–19%; Rust has grown from 2% to 11%.

| Go tool | Rust equivalent | Notes |
| --- | --- | --- |
| `go build` | `cargo build` | Compile the project |
| `go run .` | `cargo run` | Build and run |
| `gofmt` / `goimports` | `cargo fmt` | Auto-formatter, zero config |
| `go test ./...` | `cargo test` | Testing built into the toolchain |
| `go vet ./...` | `cargo clippy` | Linter, Clippy is significantly more opinionated than `vet` |
| `go install ./cmd/foo` | `cargo install --path .` | Install a binary |
| `golangci-lint run` | `cargo clippy -- -D warnings` | Strict lint mode |
| `go doc` | `cargo doc` | Generate and view API docs |
| `pprof` | `cargo flamegraph` / `samply` | CPU profiling |
| `govulncheck` | `cargo audit` | Vulnerability scanning against an advisory database |

People often claim that a managed runtime is “good enough for most backends”, but I think they are missing the point. In my opinion, the tradeoff is more that Go optimizes for quick iteration speed whereas Rust optimizes for correctness. The fact that you have more control over memory is just a nice side effect for most production workloads. It means that you need fewer machines to do the same work, but the main reason to choose Rust is still robustness.

```plain text
func (s *Service) Handle(req *Request) error {
    // Find returns (*User, error). The error is nil for "not found";
    // the caller is expected to check user != nil, but this is very easy to forget.
    user, err := s.repo.Find(req.UserID)
    if err != nil {
        return err
    }
    return user.Account.Notify() // crashes if user is nil, or if Account is nil
}
```

```plain text
#[derive(Debug, thiserror::Error)]
pub enum UserError {
    #[error("user {0} not found")]
    NotFound(UserId),
    #[error("user already exists")]
    AlreadyExists,
    #[error(transparent)]
    Repo(#[from] RepoError),
}

pub fn rename(id: UserId, name: &str) -> Result<User, UserError> {
    let mut user = repo::get(id)?;        // ? converts RepoError -> UserError automatically
    user.name = name.to_string();
    Ok(user)
}
```

```plain text
func ReadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("reading config: %w", err)
    }
    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, fmt.Errorf("parsing config: %w", err)
    }
    return &cfg, nil
}
```

```plain text
fn read_config(path: &Path) -> Result<Config, ConfigError> {
    let data = fs::read_to_string(path)?;
    let cfg = serde_json::from_str(&data)?;
    Ok(cfg)
}
```

```plain text
func GetUser(id string) *User {
    for _, u := range users {
        if u.ID == id {
            return &u
        }
    }
    return nil
}

u := GetUser("123")
fmt.Println(u.Name)
```

```plain text
fn get_user(id: &str) -> Option<User> {
    users.iter().find(|u| u.id == id).cloned()
}

let user = get_user("123");
println!("{}", user.name); // compile error: `user` is Option<User>, not User
// You must handle both cases:
match get_user("123") {
    Some(u) => println!("{}", u.name),
    None    => println!("not found"),
}
```

```plain text
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

```plain text
pub trait Reader {
    fn read(&mut self, buf: &mut [u8]) -> std::io::Result<usize>;
}

impl Reader for MyType {
    fn read(&mut self, buf: &mut [u8]) -> std::io::Result<usize> { /* ... */ }
}
```

```plain text
tokio::spawn(async move {
    do_work(input).await;
});
```

```plain text
ch := make(chan int, 10)
go func() {
    ch <- 42
}()
v := <-ch
```

```plain text
let (tx, mut rx) = tokio::sync::mpsc::channel::<i32>(10);
tokio::spawn(async move {
    tx.send(42).await.unwrap();
});
let v = rx.recv().await.unwrap();
```

```plain text
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

```plain text
pub struct Circle {
    pub radius: f64,
}

impl Circle {
    pub fn area(&self) -> f64 {
        std::f64::consts::PI * self.radius * self.radius
    }
}
```

```plain text
let s = "héllo";
s[1] // compile error:
     // `str` cannot be indexed
     // by `{integer}`
```

Feel free to skip this section if you don’t care about generics much. 😅 In hindsight, I don’t think it’s all too important for the working engineer, but it’s part of the story of what makes Go and Rust different, and of the philosophical mindset behind both.

| Concern | Go | Rust |
| --- | --- | --- |
| HTTP server | `net/http`, `chi`, `gin`, `echo`, `fiber` | `axum` (on `hyper`) |
| HTTP client | `net/http`, `resty` | `reqwest` |
| gRPC | `google.golang.org/grpc` + `protoc-gen-go` | `tonic` + `prost` |
| OpenAPI (codegen) | `oapi-codegen` | `utoipa` (code-first) or `openapi-generator` |
| SQL | `database/sql`, `sqlc`, `sqlx`, `gorm` | `sqlx`, `sea-orm`, `diesel` |
| Migrations | `golang-migrate`, `goose` | `sqlx migrate`, `refinery` |
| JSON | `encoding/json`, `sonic`, `goccy/go-json` | `serde` + `serde_json` |
| Logging | `log/slog`, `zerolog`, `zap` | `tracing` + `tracing-subscriber` |
| Metrics | `prometheus/client_golang` | `metrics` + `metrics-exporter-prometheus` |
| Config | `viper`, `koanf` | `config` (config-rs), `figment` |
| CLI | `cobra`, `urfave/cli` | `clap` (derive) |
| Validation | `go-playground/validator` | `validator` |
| Errors | `errors`, `pkg/errors` | `thiserror` (libraries), `anyhow` (binaries) |
| Testing | `testing`, `testify`, `gomega` | built-in `#[test]`, `rstest`, `assert_matches` |
| Mocking | `mockgen`, `moq` | hand-written fakes (idiomatic), `mockall` |
| HTTP mocking | `httptest` | `httpmock`, `wiremock-rs` |
| Real deps in tests | `testcontainers-go` | `testcontainers` |
| Retry/backoff | `cenkalti/backoff` | `backon` |
| Background tasks | goroutines + `errgroup` | `tokio::spawn` + `JoinSet` |

**Start with a service that has a clear boundary.**

Don’t pick the most central, most-deployed service in your fleet. Pick the one where the contract with the rest of the system is well-defined and the blast radius is small.

**Keep the same API contract.**

If your Go service exposes a REST API, your Rust service should too: same paths, same JSON shapes, same error envelope. The migration is invisible to clients, and you can swap traffic incrementally with a gateway.

**Don’t translate idioms verbatim.**

Resist the urge to write Go-flavoured Rust. `if err != nil { return err }` becomes `?`. Goroutine-per-request becomes `tokio::spawn` only when you actually need it (axum already concurrently handles requests). Interfaces with one method usually become trait bounds on a generic, not `Box<dyn Trait>`.

**Use the compiler as a pair programmer.**

Rust’s compiler errors are usually pretty good. Read them slowly. They almost always tell you the right answer. The team members who struggle longest are the ones who fight the compiler instead of treating it as a collaborator.

**Invest in training early.**

I’ve seen teams try to do a Rust migration “on the side,” learning as they go. It rarely ends well. It’s a bit like training for a marathon by signing up for the race and then trying to run it without any prior training. You can do it, but it’s going to be painful and you might not finish.

Block off real time for learning: a workshop, [an online course](https://course.corrode.dev/), paired sessions on real code. The upfront investment pays back many times over once the team is fluent. (Hey, if you want to talk about training options, [I’m happy to chat](https://corrode.dev/services).)
