// Package driver contains API implementations
package driver

type Driver interface {
	Serve() error
}
