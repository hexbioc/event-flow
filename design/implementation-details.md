# Implementation Details

## Source (emulated)

The source is emulated using an open source load-testing tool,
[vegeta](https://github.com/tsenart/vegeta). The
[sources/emulated-source](../sources/emulated-source/README.md) directory
describes tool usage in detail.

## Collector Service

The collector service is built using Golang with the [Gin](https://gin-gonic.com/)
web framework. Golang is used considering it's high throughput despite a small
resource footprint.

Additionally, since I was a bit rusty with Golang, I decided to use this
opportunity to get back up to speed with its eco-system.

More details in the [sources/collector](../sources/collector/) directory.

## Message Broker

I've chosen [RabbitMQ](https://www.rabbitmq.com/) as the message broker,
considering it's reliable architecture and high throughput.

Alternatives considered:

* AWS SQS: While easy to setup, in practice the system provides very
  limited flexibiility and is unsuitable as the architecture evolves
* AWS EventBridge: While a great option, the proprietary nature would
  mean a vendor lock-in and potentially escalating costs
* Apache Kafka: Did not use as I have limited experience with the same

The current implementation uses the default exchange with a pre-defined queue
for simplicity. In practice, I would like create a dedicated-exchange and
further tune the setup as required.

## Processor Service

The processor is a custom worker implementation written in TypeScript.
It would be quicker to implement this as a part of the
[collector service](#collector-service) itself, deploying it with a different
entrypoint, but I built it separately because I wanted to:

1. Demonstrate clear de-coupling between the two systems
2. Demonstrate my TypeScript skillzz hehehee :D

The collector has several abstractions, such as:

* [worker.ts](../sources/processor/src/worker.ts): Handles mutations on
  the event payload as required, for example to add a new key-value pair
* [target.ts](../sources/processor/src/target.ts): Handles target invocation
  and rate limiting, the implementation can be modified as required

More details in the [sources/processor](../sources/processor/) directory.

## Target (mocked)

Since the target is a GraphQL service, the communication with it is
likely to be using a straight-forward REST endpoint, such as `POST /graphql`,
with the query / mutation supplied in the request body. In the interest of
time, I've mocked this request using `setTimeout` and a random response time,
bounded by 1 second.

While the above means there is no response handling, it would be fairly trivial
to do the same, for example we could reset the rate-limiting bucket to 0 on
receiving a `HTTP 429 Too Many Requests` in the response.
