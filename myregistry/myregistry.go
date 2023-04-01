package myregistry

import (
	"google.golang.org/grpc/metadata"
)

type ServiceNodeInfo struct {
	host     string
	port     string
	Address  string
	Metadata metadata.MD
}

type Registrar interface {
	Register(service *ServiceNodeInfo) error
	Unregister(service *ServiceNodeInfo) error
	Close()
}
