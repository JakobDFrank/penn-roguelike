// Package driver contains API implementations
package driver

type Driver interface {
	Serve() error
}

type DriverKind int

const (
	Http DriverKind = iota
	Grpc
	GraphQL
)

var DriverNameToEnum = map[string]DriverKind{
	"http":    Http,
	"grpc":    Grpc,
	"graphql": GraphQL,
}
