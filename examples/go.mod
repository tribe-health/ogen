module github.com/ogen-go/ogen/examples

go 1.17

require (
	github.com/go-faster/errors v0.5.0
	github.com/go-faster/jx v0.23.3
	github.com/google/uuid v1.3.0
	github.com/ogen-go/ogen v0.0.0
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/metric v0.24.0
	go.opentelemetry.io/otel/trace v1.2.0
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/goccy/go-yaml v1.9.4 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/sys v0.0.0-20211030160813-b3129d9d1021 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

replace github.com/ogen-go/ogen v0.0.0 => ./..
