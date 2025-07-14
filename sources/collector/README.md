# Collector Service

The collector service captures incoming events from the source, and
pushes them to RabbitMQ.

## Local Development

Setup RabbitMQ using the [`compose.yml`](../infra/local/compose.yml) file.

```sh
# Configure sources/infra/local/.env as required
docker compose up -d
```

Install dependencies:

```sh
go mod download
```

Create a `.env` file using the [`env.template`](./env.template) as
reference.

Run the application:

```sh
go run main.go
```
