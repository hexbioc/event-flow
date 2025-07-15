# Event Flow

A sample application demonstrating efficient and reliable processing and flow
of events between unreliable systems. Head over to the
[design doc](./design/README.md) for detailed documentation.

## Project Structure

The [`design`](./design) directory hosts all documentation related to
the design of this system.

The [`sources`](./sources/) directory is sub-divided as follows:

* `emulated-source`: Event source emulation as described in the problem
  statement
* `collector`: Source code for the event collector interfacing with the event
  source
* `processor`: Source code for the event processor interfacing with the event
  target
* `infra`: Infra related configuration in the form of `docker compose` for `local`
  and `terraform` config for `cloud` (TODO)
