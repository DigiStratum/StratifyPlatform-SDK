Our Auth model bears a number of small components that must all integrate with each other for a
given implementation in order to provide a functioning and secure system. We approach this with
an abstraction that we hope will accommodate any number of single-factor authentication schemes,
with an expectation to expand into multi-factor auth as a fast follow-on. At a minimum, we know
that this model will be adequate to suport common HTTP authentication schemes such as:<br/>
<br/>
(From https://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml)<br/>
 * Basic   [RFC7617] https://www.iana.org/go/rfc7617
 * Bearer  [RFC6750] https://www.iana.org/go/rfc6750
 * Digest  [RFC7616] https://www.iana.org/go/rfc7616
<br/>
The following interfaces are defined with their described interactions:<br/>
</br>
 * Grants - Named authorizations/permissions granted to a given Identity
 * Identity - Authenticator-specific user identity produced from provided Credential
 * Credential - Combination of information provided to authenticate an Identity (user, pass, etc)
 * Authorizor - Compares the Grants for a given Identity to some action performed against a
                Protected Resource; an app-specific Authorizor will want to factor in other
		considerations such as whether the user should be able to access some resources
		but not others with varied permissions for each. The method & manner of organizing
		this is implementation-specific

 * Session - An auth session for an authenticated Identity to retain stateful metadata for the life
             of the session.
 * SessionStore - A storage model for CRUD operations on Session resources
 * Authenticator - Identity Provider interface to authenticate Identity Credentials 
  ^^^ Maybe Authenticate() is actually a function of the Identity Provider and we don't need a
      separate layer for Authenticator...?
