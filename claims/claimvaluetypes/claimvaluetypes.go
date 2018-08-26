package claimvaluetypes

const (
	// xmlSchemaNamespace contains XML schema claim value types
	xmlSchemaNamespace = "http://www.w3.org/2001/XMLSchema"
	// Base64Binary is the claim value type for the xml base64Binary datatype
	Base64Binary = xmlSchemaNamespace + "#base64Binary"
	// Base64Octet is the claim value type for the xml base64Octet datatype
	Base64Octet = xmlSchemaNamespace + "#base64Octet"
	// Boolean is the claim value type for the xml boolean datatype
	Boolean = xmlSchemaNamespace + "#boolean"
	// Date is the claim value type for the xml date datatype
	Date = xmlSchemaNamespace + "#date"
	// DateTime is the claim value type for the xml dateTime datatype
	DateTime = xmlSchemaNamespace + "#dateTime"
	// Double is the claim value type for the xml double datatype
	Double = xmlSchemaNamespace + "#double"
	// Fqbn is the claim value type for the xml fqbn datatype
	Fqbn = xmlSchemaNamespace + "#fqbn"
	// HexBinary is the claim value type for the xml hexBinary datatype
	HexBinary = xmlSchemaNamespace + "#hexBinary"
	// Integer is the claim value type for the xml integer datatype
	Integer = xmlSchemaNamespace + "#integer"
	// Integer32 is the claim value type for the xml integer32 datatype
	Integer32 = xmlSchemaNamespace + "#integer32"
	// Integer64 is the claim value type for the xml integer64 datatype
	Integer64 = xmlSchemaNamespace + "#integer64"
	// Sid is the claim value type for the xml sid datatype
	Sid = xmlSchemaNamespace + "#sid"
	// String is the claim value type for the xml string datatype
	String = xmlSchemaNamespace + "#string"
	// Time is the claim value type for the xml time datatype
	Time = xmlSchemaNamespace + "#time"
	// UInteger32 is the claim value type for the xml uinteger32 datatype
	UInteger32 = xmlSchemaNamespace + "#uinteger32"
	// UInteger64 is the claim value type for the xml uinteger64 datatype
	UInteger64 = xmlSchemaNamespace + "#uinteger64"

	// soapSchemaNamespace contains SOAP schema claim value types
	soapSchemaNamespace = "http://schemas.xmlsoap.org/"
	// DNSName is the type of a dns Name
	DNSName = soapSchemaNamespace + "claims/dns"
	// Email is the type of an email address
	Email = soapSchemaNamespace + "ws/2005/05/identity/claims/emailaddress"
	// Rsa is the type of an rsa key
	Rsa = soapSchemaNamespace + "ws/2005/05/identity/claims/rsa"
	// UpnName is the type of an UPN name
	UpnName = soapSchemaNamespace + "claims/UPN"

	// xmlSignatureConstantsNamespace contains XML Signature claims
	xmlSignatureConstantsNamespace = "http://www.w3.org/2000/09/xmldsig#"
	// DsaKeyValue is the type of a DSA key Value claim
	DsaKeyValue = xmlSignatureConstantsNamespace + "DSAKeyValue"
	// KeyInfo is the type of a DSA key info claim
	KeyInfo = xmlSignatureConstantsNamespace + "KeyInfo"
	// RsaKeyValue is the type of a RSA key value claim
	RsaKeyValue = xmlSignatureConstantsNamespace + "RSAKeyValue"

	//xqueryOperatorsNameSpace contains XQuery value types
	xqueryOperatorsNameSpace = "http://www.w3.org/TR/2002/WD-xquery-operators-20020816"
	// DaytimeDuration is the claim value type for the xml dayTimeDuration datatype
	DaytimeDuration = xqueryOperatorsNameSpace + "#dayTimeDuration"
	// YearMonthDuration is the claim value type for the xml yearMonthDuration datatype
	YearMonthDuration = xqueryOperatorsNameSpace + "#yearMonthDuration"

	// xacml10Namespace contains some oasis types
	xacml10Namespace = "urn:oasis:names:tc:xacml:1.0"
	// Rfc822Name is the value type of an RFC822 name
	Rfc822Name = xacml10Namespace + ":data-type:rfc822Name"
	// X500Name is the value type of an X500 name
	X500Name = xacml10Namespace + ":data-type:x500Name"
)
