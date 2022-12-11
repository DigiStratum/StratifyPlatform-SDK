package auth

/*
Hold a list of Authorization Grants

This is meant to reflect a set of Authorizations within a given context (Scope/Realm/App), etc for a
givn user Identity. That said, the model is generalized enough that could be used in other ways as
well such as setting up Authorization packages that can be granted/revoked in bulk for administrative
purposes, etc.

Each Authorization Grant is identified by a unique string (key) and bears an optional value for a
Metadata string; both the Authorization Grant name and the metadata are context-specific. That is to
say that there could be a valid authorization grant named "fulladmin" for two entirely different
contexts. It is the consumer's responsibility to organize the AuthorizationGrants collections by
context/identity as this could take any number of forms and is outside the scope of this interface.

Our underlying HashMap data structure is organized as: [GrantName] -> [Context-Specific-Metadata]

*/

import (
	hash "github.com/DigiStratum/GoLib/Data/hashmap"
)

type GrantsIfc interface {
	// Embed HashMapIfc to inherit all of that...
	hash.HashMapIfc
}

type grants struct {
	// Embed HashMap struct to inherit all of that...
	*hash.HashMap
	scopeName	string
}

// -------------------------------------------------------------------------------------------------
// Factory Functions
// -------------------------------------------------------------------------------------------------

func NewGrants(scopeName string) *grants {
	return &grants{
		HashMap:	hash.NewHashMap(),
		scopeName:	scopeName,
	}
}

// -------------------------------------------------------------------------------------------------
// GrantsIfc Implementation
// -------------------------------------------------------------------------------------------------

func (r *grants) GetScopeName() string {
	return r.scopeName
}

