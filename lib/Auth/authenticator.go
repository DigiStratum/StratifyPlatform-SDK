package auth

/*

Convert some form of authentication credential into an authenticated Identity. If authentication
fails, then no Identity will be produced.

Common auth Schemes; ref: https://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml
Basic   [RFC7617] https://www.iana.org/go/rfc7617
Bearer  [RFC6750] https://www.iana.org/go/rfc6750
Digest  [RFC7616] https://www.iana.org/go/rfc7616

*/

type AuthenticatorIfc interface {
	// Authenticate the credential and return nil+error or Identity/no error on success
	Authenticate(credential SerializedCredentialIfc) (IdentitiyIfc, error)
}

