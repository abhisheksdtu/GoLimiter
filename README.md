# GoLimiter

A simple Go web app for distributed rate limiting.

## Running

```
cd GoLimiter/cmd/server

go run main.go
```

The server listens on port 8080 by default. You can override with the `PORT` environment variable.

Health check: [http://localhost:8080/healthz](http://localhost:8080/healthz) 
