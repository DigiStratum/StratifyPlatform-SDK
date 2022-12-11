package auth

/*

Regardless of Serialized Credential type, the implementation must somehow serialize all the parts
necessary for the Authenticator in play to be able to deserialize it back into the pieces
necessary to perform the actual credential verification.
*/

type SerializedCredentialIfc interface {
	// Serialize credential tostring, whether username, password, key, secret or otherwise
        Serialize() (*string, error)
}


