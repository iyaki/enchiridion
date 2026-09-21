---
title: "semgrep - Lightweight static analysis for many languages. Find bug variants with patterns that look like source code"
notion_id: 10d2ac6b-e240-484c-88c9-92905d5fccfc
notion_url: https://app.notion.com/p/semgrep-Lightweight-static-analysis-for-many-languages-Find-bug-variants-with-patterns-that-look--10d2ac6be240484c88c992905d5fccfc
last_edited: 2023-10-05T13:17:00.000Z
source_url: https://github.com/returntocorp/semgrep
tags: ["Tool", "English", "Information Security", "Programming"]
---
<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

## 



<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



| Category | Languages |
| --- | --- |
| GA | C# · Go · Java · JavaScript · JSX · JSON · PHP · Python · Ruby · Scala · Terraform · TypeScript · TSX |
| Beta | Kotlin · Rust |
| Experimental | Bash · C · C++ · Clojure · Dart · Dockerfile · Elixir · HTML · Julia · Jsonnet · Lisp · Lua · OCaml · R · Scheme · Solidity · Swift · YAML · XML · Generic (ERB, Jinja, etc.) |

### 

1. 
2. 



### 

1. 

```

```

1. 
2. 

### 

<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

1. 
2. 
3. 
4. 
5. 

### 

1. 
2. 
3. 

### 



- 
- 
- 
- 



- 
- 





### 





<!-- image hosted by Notion: its URL expires and is not preserved (ADR-05) -->

### 



| Use case | Semgrep rule |
| --- | --- |
| Ban dangerous APIs | [Prevent use of exec](https://semgrep.dev/playground/s/lglB) |
| Search routes and authentication | [Extract Spring routes](https://semgrep.dev/playground/s/Y6wD) |
| Enforce the use secure defaults | [Securely set Flask cookies](https://semgrep.dev/playground/s/6KwW) |
| Tainted data flowing into sinks | [ExpressJS dataflow into sandbox.run](https://semgrep.dev/playground/s/qEpR) |
| Enforce project best-practices | [Use assertEqual for == checks](https://semgrep.dev/playground/s/oEox), [Always check subprocess calls](https://semgrep.dev/playground/s/zENk) |
| Codify project-specific knowledge | [Verify transactions before making them](https://semgrep.dev/playground/s/p8zk) |
| Audit security hotspots | [Finding XSS in Apache Airflow](https://semgrep.dev/playground/s/KPwj), [Hardcoded credentials](https://semgrep.dev/playground/s/2Br8) |
| Audit configuration files | [Find S3 ARN uses](https://semgrep.dev/playground/s/jEKD) |
| Migrate from deprecated APIs | [DES is deprecated](https://semgrep.dev/playground/r/java.lang.security.audit.crypto.des-is-deprecated.des-is-deprecated), [Deprecated Flask APIs](https://semgrep.dev/playground/r/python.flask.maintainability.deprecated.deprecated-apis.flask-deprecated-apis), [Deprecated Bokeh APIs](https://semgrep.dev/playground/r/python.bokeh.maintainability.deprecated.deprecated_apis.bokeh-deprecated-apis) |
| Apply automatic fixes | [Use listenAndServeTLS](https://semgrep.dev/playground/s/1Ayk) |

### 



### 



### 









### 

- 
- 
- 
- 
- 
- 
- 

### 



```

```
