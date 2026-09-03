# vulner

`vulner` reports reachable Go vulnerabilities using `golang.org/x/vuln/scan` and the Go vulnerability database. Both supported interfaces scan every package in the Go module containing the working directory; neither accepts package arguments or settings.

## Standalone command

Build the command, then run it with no arguments:

```bash
go build -o vulner ./cmd/vulner
./vulner
```

Diagnostics use `file:line:column: message (vulner)` formatting. The command exits `0` when no vulnerabilities are found, `1` for findings or scan failures, and `2` when arguments are supplied.

## golangci-lint module plugin

`.custom-gcl.yml` is the source of truth for the custom golangci-lint binary. Build it, then run it with no package arguments:

```bash
golangci-lint custom
./custom-golangci-lint run
```

Enable the module plugin in `.golangci.yml` without a `settings` section:

```yaml
version: "2"

linters:
  enable:
    - vulner
  settings:
    custom:
      vulner:
        type: module
        description: Report reachable Go vulnerabilities
```

The generated `custom-golangci-lint` executable is not versioned. Rebuild it from `.custom-gcl.yml` in CI or before release.
