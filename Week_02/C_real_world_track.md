# Real World Track

- [Start](README.md)
- [End of week]

---

## Week 2 - OpenTelemetry

| Topic | Link |
| ----------- | ----------- |
| OpenTelemetry | [opentelemetry.io/](https://opentelemetry.io/) |
| OpenTelemetry Go SDK | [github.com/open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go) |
| Third party Opentelemetry libs |  [opentelemetry.io/docs/instrumentation/go/libraries/](https://opentelemetry.io/docs/instrumentation/go/libraries/) |
| CLI | [github.com/spf13/cobra](https://github.com/spf13/cobra) |

OpenTelemetry is a CNCF incubating project with a good degree of maturity in tracing. It's full functionality is a collection of tools, APIs, and SDKs. Use it to instrument, generate, collect, and export telemetry data (metrics, logs, and traces) to help you analyze your software’s performance and behavior. It is vendor neutral integrating with datadog and prometheus.

### HTTP Server with OpenTelemetry

Code: [http_server.go](C/01_http_server/http_server.go)

Go playground: n/a

`go run C/01_http_server/http_server.go`

### HTTP Client (same as week 1 code)

Code: [http_client.go](C.02b_http_client/http_client.go)

Go playground: n/a

`go run C/02_http_client/http_client.go -url http://localhost:8080/hello/gophers`

## Week 2 - CLI

| Topic | Link |
| ----------- | ----------- |
| CLI | [github.com/spf13/cobra](https://github.com/spf13/cobra) |

Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.  Cobra is used in many Go projects such as Kubernetes, Hugo, and Github CLI to name a few. This list contains a more extensive list of projects using Cobra.

### Cli example

Code: [C/03_cli/cli.go](C/03_cli/cli.go)

Go playground: n/a

`go run C/03_cli/cli.go get -u http://localhost:8080/hello/gophers`

---

- [Start](README.md)
- [End of week]
