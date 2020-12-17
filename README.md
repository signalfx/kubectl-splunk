# kubectl-signalfx

A kubectl plugin for interacting with o11y-collector deployments. It can query collector pods, daemonsets, etc. with a query selector already set (`--selector app=o11y-collector`). The [support](./docs/kubectl-signalfx_support.md) command can be used to gather all relevant spec files (excluding secrets) as well as the pod log files.

Status: beta

## Installation

Download the [latest release](https://github.com/signalfx/kubectl-signalfx/releases) and copy the `kubectl-signalfx` binary to `/usr/local/bin` or somewhere in your `$PATH`.

## Usage

See [docs](docs/kubectl-signalfx.md).

## License

[Apache 2.0](./LICENSE)
