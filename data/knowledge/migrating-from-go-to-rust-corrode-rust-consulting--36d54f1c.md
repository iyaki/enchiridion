---
title: "Migrating from Go to Rust | corrode Rust Consulting"
notion_id: 36d54f1c-7d23-811d-9e5b-ec617f8ae90c
notion_url: https://app.notion.com/p/Migrating-from-Go-to-Rust-corrode-Rust-Consulting-36d54f1c7d23811d9e5bec617f8ae90c
last_edited: 2026-09-18T00:52:00.000Z
source_url: https://corrode.dev/learn/migration-guides/go-to-rust/
tags: ["Article", "Tutorial", "corrode Rust Consulting", "English", "Programming", "Rust", "Go", "Software Development", "Systems Design / Software Architecture"]
---
- 
- 
- 
- 
- 

![image](https://corrode.dev/learn/migration-guides/go-to-rust/go-usage.svg)



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



```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```



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






















