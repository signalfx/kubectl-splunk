# kubectl-splunk

A kubectl plugin for interacting with splunk-otel-collector deployments. It can query collector pods, daemonsets, etc.
with a query selector already set (`--selector app=splunk-otel-collector`).
The [support](./docs/kubectl-splunk_support.md) command can be used to gather all relevant spec files (excluding
secrets) as well as the pod log files.

Status: beta

## Installation

Download the [latest release](https://github.com/signalfx/kubectl-splunk/releases) and copy the `kubectl-splunk` binary
to `/usr/local/bin` or somewhere in your `$PATH`.

## Usage

See [docs](docs/kubectl-splunk.md).

## Contributing

When updating dependencies run `make license` to regenerate `NOTIFICATIONS` file.

## License

[Apache 2.0](./LICENSE)
