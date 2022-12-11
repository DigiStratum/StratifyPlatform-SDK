package auth

/*
This is the identity of an authenticated "user" (maybe it's not even a human?). We will provide some convenience mechanisms for association between this Identity and ScopedAuthorizationGrants.
*/

type IdentityIfc interface {
}

type identity struct {
}

// Factory Functions

func NewIdentity() *identity {
	return &identity{ }
}

