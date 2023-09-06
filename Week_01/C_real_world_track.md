# C. Real World Track

- [Start](README.md)
- [End of week 1]

---

Week 1:

| Topic | Link |
| ----------- | ----------- |
| HTTP Client | [gobyexample.com/http-clients](https://gobyexample.com/http-clients) |
| HTTP Server | [gobyexample.com/http-servers](https://gobyexample.com/http-servers) |

## Real world examples

### HTTP Server

Code: [http_server.go](C/01_http_server.go)

go run C/01_http_server.go

Go playground: n/a

### HTTP Client

Code: [http_client.go](C/02_http_client.go)

```go
// defaults to querying the openid well_known endpoint in active directory
go run C/02_http_client.go

// alternatively query the end point
go run C/02_http_client.go -url http://localhost:8080/hello/team
```

Go playground: n/a

---

- [Start](README.md)
- [End of week 1]

`:)`
