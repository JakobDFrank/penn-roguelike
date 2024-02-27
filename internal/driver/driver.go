// Package driver contains API implementations.
package driver

// Driver is an implementation of a server for our application.
type Driver interface {
	Serve() error
}

// DriverKind encapsulates the application's different communication protocols.
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
