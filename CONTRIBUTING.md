<!--
SPDX-FileCopyrightText: 2025 Matthew Taylor <taylormw163@gmail.com>
SPDX-License-Identifier: CC-BY-NC-SA-4.0
-->

# Contributing

## Prerequisites

- Golang 1.25 or higher ([obtainable from their website](https://go.dev/dl/))
- The respository [forked and cloned on your device](https://docs.github.com/en/repositories/creating-and-managing-repositories/cloning-a-repository))

### Installing Mage

Next, you'll also probably want to install [Mage](https://github.com/magefile/mage), a make/rake-like build tool & command executer written in Go.
Run the following command in your terminal of choice:

```bash
go install github.com/magefile/mage@latest
```

Once it finishes installing, check by running `mage` - if all went well, you should get a list of available targets defined in the repo's [magefiles](../magefiles) directory.

> [!NOTE]
> Magefile targets must always be run from inside the _repository root_.

### Installing golangci-lint
The _other_ major required dependency is [golangci-lint](https://github.com/golangci/golangci-lint), used for rapidly running various linters on the entire codebase.

You can download the appropriate binary from their [releases page](https://github.com/golangci/golangci-lint/releases) and put it in your `PATH`, or run the following Bash command to chuck it into `GOPATH/bin`:
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.5.0
```

## Commands
`mage test` - Run automated tests
`mage lint` - Run `golangci-lint` code quality checks
