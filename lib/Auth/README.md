Our Auth model bears a number of small components that must all integrate with each other for a
given implementation in order to provide a functioning and secure system. We approach this with
an abstraction that we hope will accommodate any number of single-factor authentication schemes,
with an expectation to expand into multi-factor auth as a fast follow-on. At a minimum, we know
that this model will be adequate to suport common HTTP authentication schemes.<br/>
<br/>
Please reference our Glossary of Terms as needed:<br/>
 * https://github.com/DigiStratum/GoStratifyPlatform-SDK/wiki/Stratify-SaaS-Platform---Glossary-of-Terms 
<br/>
Reference Sources:<br/>
 * https://docs.oracle.com/cd/E13222_01/wls/docs90/secintro/realm_chap.html
 * https://www.techtarget.com/searchsecurity/definition/identity-access-management-IAM-system
 * (From https://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml)<br/>
   * Basic   [RFC7617] https://www.iana.org/go/rfc7617
   * Bearer  [RFC6750] https://www.iana.org/go/rfc6750
   * Digest  [RFC7616] https://www.iana.org/go/rfc7616

