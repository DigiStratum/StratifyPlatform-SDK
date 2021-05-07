package authenticator

// Authentication public interface
type AuthenticationIfc interface {
	Authorize(action string) bool
}

// Authenticator public interface
type AuthenticatorIfc interface {
	Autheticate(credential string) AuthenticationIfc
}

