package auth

type ResourceIfc interface {
}

type resource struct {
	moduleId		string
	endpointId		string
	typeName		string
	authRequired		bool
}

// -------------------------------------------------------------------------------------------------
// Factory Functions
// -------------------------------------------------------------------------------------------------

func NewResource() *resource {
	return &resource{}
}

