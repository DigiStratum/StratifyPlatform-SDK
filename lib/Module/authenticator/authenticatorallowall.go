package authenticator

// Authentication public interface
type AuthenticationIfc interface {
	Authorize(action string) bool
}

type authenticator struct { }

// Make a new one of these!
func NewAuthenticatorAllowAll() AuthenticatorIfc {
	a := authenticator{}
	return &a
}

type authentication struct { }

// Make a new one of these!
func newAutheticationAllowAll() AuthenticationIfc {
	a := authentication{}
	return &a
}

// -------------------------------------------------------------------------------------------------
// AuthenticatorIfc

func (a *authenticator) Autheticate(credential string) AuthenticationIfc {
	// No need to check the credential - we allow all!
	return newAuthenticationAllowAll()
}

// -------------------------------------------------------------------------------------------------
// AuthenticationIfc

func (a *authentication) Authorize(action string) bool {
	// No need to check the action - we allow all!
	return true
}

