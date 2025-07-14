# Publisher / Event Source

In order to emulate an application that can produce events, we use
[`vegeta`](https://github.com/tsenart/vegeta), an awesome load-testing
tool.

`vegeta` can generate sustained loads for a prolonged period, limited only
by the available resources on the host system.

To generate a sample load, create a sample payload file (or use
[sample-payload.json](./sample-payload.json)), and a target file named
`api` (filename is arbitrary) with contents as below:

```txt
POST http://<server-host:server-port>/events
Content-Type: application/json
X-Api-Key: <some-api-key>
@./sample-payload.json
```

Then trigger the load test:

```sh
vegeta attack \
    -duration=15s \
    -rate=50 \
    -workers=5 \
    -max-workers=20 \
    -targets=api | vegeta report
```

The above command will build up to a request rate of ~20 reqeusts/second, generated
by 5-20 concurrent workers. This load will be sustained for a period of 15 seconds.
Effectively, about 750 requests would be sent, in all.

At the end of the load test, `vegeta` will also generate a detailed report,
including metrics such as latency, request duration, and response status codes.
