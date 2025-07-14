# Processor Service

The processor service binds to RabbitMQ, to consume and process
messages sent by the source. It performs transformations on each
event, and invokes the target with the updated event.

## Local Development

Setup RabbitMQ using the [`compose.yml`](../infra/local/compose.yml) file.

```sh
# Configure sources/infra/local/.env as required
docker compose up -d
```

Install dependencies:

```sh
npm install
```

Create a `.env` file using the [`env.template`](./env.template) as
reference.

Run the application:

```sh
npm run dev
```
