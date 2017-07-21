# yaml-interactor

## dependence manage

- golang 1.8+
- glide 0.12.3+

## format

```bash
# see which code needs format
find . -path ./vendor -prune -o -name '*.go' -print | xargs gofmt -d | grep '^'
# format
glide nv | xargs go fmt
```

## before run
```bash
# install dependence
glide install
# update dependence
glide up
```

## usage
```
complete a yaml file in command line

Usage:
  yaml-file-generator [flags]

Flags:
  -h, --help            help for yaml-file-generator
  -s, --source string   configure file to load
  -t, --target string   configure file to write

```