package module

/*

EndpointInfo - Essential Endpoint metadata to support Discovery and Configuration stages of launch.

*/

import "fmt"

// EndpointInfo public interface
type EndpointInfoIfc interface {
	GetName() string
	GetVersion() string
	GetId() string
}

type EndpointInfo struct {
	name	string
	version	string
	id	string
}
// Make a new one of these
func NewEndpointInfo(name, version string) EndpointInfoIfc {
	return &EndpointInfo{
		name:		name,
		version:	version,
		id:		fmt.Sprintf("%s-%s", name, version),
	}
}

// Get the Name for this Endpoint
func (epi *EndpointInfo) GetName() string { return epi.name }

// Get the Version for this Endpoint
func (epi *EndpointInfo) GetVersion() string { return epi.version }

// Get the ID for this Module (must be unique within Module's scope)
func (epi *EndpointInfo) GetId() string { return epi.id }

