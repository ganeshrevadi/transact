# transact

A small payment system API written in Go.

## Overview

`transact` exposes an HTTP API for creating and retrieving payments. Payments
are stored in-memory and identified by a generated `pay_`-prefixed ID. Each
payment tracks an amount (in minor units), a currency, and a lifecycle status.

Payment statuses:

| Status       | Description                     |
| ------------ | ------------------------------- |
| `CREATED`    | Payment has been created        |
| `PROCESSING` | Payment is being processed      |
| `SUCCEEDED`  | Payment completed successfully  |
| `FAILED`     | Payment failed                  |

## Requirements

- Go 1.26+

## Running

```sh
go run ./cmd/api
```

The server listens on `:8080`. Check it is up:

```sh
curl http://localhost:8080/healthz
# {"status":"ok"}
```

## API

| Method | Path               | Description            |
| ------ | ------------------ | ---------------------- |
| GET    | `/healthz`         | Health check           |
| POST   | `/payments`        | Create a payment       |
| GET    | `/payments/{id}`   | Retrieve a payment     |

### Create a payment

```sh
curl -X POST http://localhost:8080/payments \
  -H 'Content-Type: application/json' \
  -d '{"amount":1000,"currency":"USD"}'
```

### Retrieve a payment

```sh
curl http://localhost:8080/payments/pay_<id>
```

## Project structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # HTTP server entrypoint
└── internal/
    └── payment/
        ├── handler.go           # HTTP handlers
        ├── service.go           # Business logic and storage
        └── model.go             # Domain types
```
