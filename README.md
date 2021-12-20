>ℹ️&nbsp;&nbsp;SignalFx was acquired by Splunk in October 2019. See [Splunk SignalFx](https://www.splunk.com/en_us/investor-relations/acquisitions/signalfx.html) for more information.

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

Submission of an executed [Splunk Contributor License Agreement](https://www.splunk.com/en_us/form/contributions.html) is a prerequisite of contributing to kubectl-splunk. 

## License

[Apache 2.0](./LICENSE)
