module github.com/ennyjfrick/ruleguard-logfatal/_example

go 1.23

replace github.com/ennyjfrick/ruleguard-logfatal => ../

require (
	github.com/ennyjfrick/ruleguard-logfatal v0.0.0-00010101000000-000000000000
	github.com/quasilyte/go-ruleguard/dsl v0.3.22
	github.com/rs/zerolog v1.33.0
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/zap v1.27.0
)

require (
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
)
