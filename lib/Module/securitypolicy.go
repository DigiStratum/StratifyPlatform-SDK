package module

/*

SecurityPolicy - enforce HTTP request/response secuirty based on configured constraints

TODO: Expand to actually check HttpRequest auth with configured authenticator(s)

*/

import (
	cfg "github.com/DigiStratum/GoLib/Config"
)

// Security Policy public interface
type SecurityPolicyIfc interface {
	SetRequireAuthentication(isRequired bool)
	RequiresAuthentication() bool
}

type securityPolicy struct {
	requireAuthentication	bool
}

// Make a new one of these!
func NewSecurityPolicy(config cfg.ConfigIfc) SecurityPolicyIfc {

	// By default we do nothing
	sp := securityPolicy{
		requireAuthentication:	false,
	}

	// By configuration, we start enabling things...
	if ! config.IsEmpty() {
		if "true" == config.Get("isrequired") {
			sp.SetRequireAuthentication(true)
		}
	}

	return &sp
}

// Set whether authentication is required
func (sp *securityPolicy) SetRequireAuthentication(isRequired bool) {
	sp.requireAuthentication = isRequired
}

// Get whether authentication is required
func (sp *securityPolicy) RequiresAuthentication() bool {
	return sp.requireAuthentication
}

