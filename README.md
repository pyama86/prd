# prd

Prd computes the highest, lowest, average of the specified file

## Usage

```
$ prd example.jp-cpu-user-d.rrd | jq
{
　"name": "example.jp-cpu-user-d.rrd",
  "max": 2010,
  "min": 1566,
  "avg": 1790
}
```

## Install

for OSX
```
$ brew install rrdtool
```

To install, use `go get`:

```bash
$ go get -d github.com/pyama86/prd
```

## Contribution

1. Fork ([https://github.com/pyama86/prd/fork](https://github.com/pyama86/prd/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[pyama86](https://github.com/pyama86)
