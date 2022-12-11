package auth

/*

A contextualized Authorizer for a given realm/app/scope. Combined with ScopedAuthorizationGrants,
the Authorizor will indicate whether a given Identity is authorized (allow or deny) to perform
some action relative to one or more Protected Resources.

*/

type AuthorizorIfc interface {
	// Compare Identity's ScopedAuthorizationGrants, named action against ProtectedResource
	IsAuthorized(identity Identity, action string, resource ProtectedResource) (bool, error)
}

