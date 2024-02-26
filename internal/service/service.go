// Package service contains game services.
package service

type ServiceKind int

const (
	Http ServiceKind = iota
	Grpc
)

var ServiceNameToEnum = map[string]ServiceKind{
	"http": Http,
	"grpc": Grpc,
}
