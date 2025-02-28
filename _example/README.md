# Examples for `github.com/ennyjfrick/ruleguard-logfatal`

This module contains examples for applicable cases of `noFatal` and `noPanic` rule groups found in `github.com/ennyjfrick/ruleguard-logfatal`, including examples using [`zap`](https://github.com/uber-go/zap), [`logrus`](https://github.com/sirupsen/logrus), [`zerolog`](https://github.com/rs/zerolog), and the standard library's [`log`](https://pkg.go.dev/log) package. 

## How to Use
1. Clone this repo and `cd` to `_example`
2. Install either [`golangci-lint`](https://golangci-lint.run) or [`ruleguard`](https://github.com/quasilyte/go-ruleguard)
3. Run the following:
```shell
# using golangci-lint
$ golangci-lint run .

# using ruleguard
$ ruleguard -rules rules.go .
```