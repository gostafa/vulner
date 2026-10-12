# Vulner


[![`Workflow for vulner Action`](https://github.com/gostafa/vulner/actions/workflows/main.yml/badge.svg)](https://github.com/gostafa/vulner/actions/workflows/main.yml) [![codecov](https://codecov.io/gh/gostafa/vulner/graph/badge.svg)](https://codecov.io/gh/gostafa/vulner) [![CodSpeed](https://img.shields.io/endpoint?url=https://codspeed.io/badge.json)](https://app.codspeed.io/gostafa/vulner?utm_source=badge) [![LICENSE](https://img.shields.io/github/license/gostafa/vulner)](/LICENSE) 

`vulner` reports reachable Go vulnerabilities using `golang.org/x/vuln/scan`
and the Go vulnerability database. It scans every package in the Go module
containing the working directory. It does not accept flags, package patterns,
or plugin settings.

It can run as:

* a standalone CLI;
* a plugin inside a custom `golangci-lint` binary.

## Use as a CLI

### Install

```bash
go get github.com/gostafa/vulner@latest
go install github.com/gostafa/vulner/cmd/vulner@latest
```

The public API lives in `github.com/gostafa/vulner`. Call `vulner.Run(args, stdout, stderr)`
to run the standalone command. Import `github.com/gostafa/vulner/plugin` for
golangci-lint registration.

### Run

```bash
vulner

# Package patterns and flags are not accepted.
# vulner ./...
# vulner --version
```

The command takes no arguments. Any flag or package pattern is a usage error.

### Build from Source

```bash
git clone https://github.com/gostafa/vulner.git
cd vulner

go build -o ./bin/vulner ./cmd/vulner
./bin/vulner
```

## Use as a Golangci-Lint Plugin

The plugin must be included in a custom `golangci-lint` binary.

Create `.custom-gcl.yml`:

```yaml
version: v2.14.0
name: custom-golangci-lint
destination: ./bin
plugins:
  - module: github.com/gostafa/vulner
    import: github.com/gostafa/vulner/plugin
    path: .
```

Enable it in `.golangci.yml`:

```yaml
version: "2"

linters:
  default: all
  enable:
    - vulner

  settings:
    custom:
      vulner:
        type: module
```

Do not add a `settings` map under `vulner`. Non-empty plugin settings are
rejected.

Build and run the custom linter:

```bash
golangci-lint custom -v
./bin/custom-golangci-lint run ./...
```

Always run the generated `custom-golangci-lint` binary. The standard
`golangci-lint` binary does not include the plugin.

## Exit Codes

* `0`: success
* `1`: check or write error
* `2`: command usage error
