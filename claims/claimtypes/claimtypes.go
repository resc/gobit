package claimtypes

const (
	prefix200909   = "http://schemas.xmlsoap.org/ws/2009/09/identity/claims/"
	prefix200505   = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/"
	prefix200806ms = "http://schemas.microsoft.com/ws/2008/06/identity/claims/"

	// Actor Its value is http://schemas.xmlsoap.org/ws/2009/09/identity/claims/actor
	Actor = prefix200909 + "actor"
	// Anonymous is The URI for a claim that specifies the anonymous user;. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/anonymous
	Anonymous = prefix200505 + "anonymous"
	// Authentication is The URI for a claim that specifies details about whether an identity is authenticated,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/authenticated
	Authentication = prefix200505 + "authenticated"
	// AuthenticationInstant is The URI for a claim that specifies the instant at which an entity was authenticated;. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/authenticationinstant
	AuthenticationInstant = prefix200806ms + "authenticationinstant"
	// AuthenticationMethod is The URI for a claim that specifies the method with which an entity was authenticated;. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/authenticationmethod
	AuthenticationMethod = prefix200806ms + "authenticationmethod"
	// AuthorizationDecision is The URI for a claim that specifies an authorization decision on an entity;. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/authorizationdecision
	AuthorizationDecision = prefix200505 + "authorizationdecision"
	// CookiePath is The URI for a claim that specifies the cookie path;. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/cookiepath
	CookiePath = prefix200806ms + "cookiepath"
	// Country is The URI for a claim that specifies the country/region in which an entity resides,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/country
	Country = prefix200505 + "country"
	// DateOfBirth is The URI for a claim that specifies the date of birth of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/dateofbirth
	DateOfBirth = prefix200505 + "dateofbirth"
	// DenyOnlyPrimaryGroupSid is The URI for a claim that specifies the deny-only primary group SID on an entity;. A deny-only SID denies the specified entity to a securable object". Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/denyonlyprimarygroupsid
	DenyOnlyPrimaryGroupSid = prefix200806ms + "denyonlyprimarygroupsid"
	// DenyOnlyPrimarySid is The URI for a claim that specifies the deny-only primary SID on an entity;. A deny-only SID denies the specified entity to a securable object". Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/denyonlyprimarysid
	DenyOnlyPrimarySid = prefix200806ms + "denyonlyprimarysid"
	// DenyOnlySid is The URI for a claim that specifies a deny-only security identifier (SID) for an entity,. A deny-only SID denies the specified entity to a securable object". Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/denyonlysid
	DenyOnlySid = prefix200505 + "denyonlysid"
	// DenyOnlyWindowsDeviceGroup is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/denyonlywindowsdevicegroup
	DenyOnlyWindowsDeviceGroup = prefix200806ms + "denyonlywindowsdevicegroup"

	// DNS is The URI for a claim that specifies the DNS name associated with the computer name or with the alternative name of either the subject or issuer of an X.509 certificate,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/dns
	DNS = prefix200505 + "dns"
	// Dsa is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/dsa
	Dsa = prefix200806ms + "dsa"
	// Email is The URI for a claim that specifies the email address of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/email
	Email = prefix200505 + "email"
	// Expiration is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/expiration
	Expiration = prefix200806ms + "expiration"
	// Expired is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/expired
	Expired = prefix200806ms + "expired"
	// Gender is The URI for a claim that specifies the gender of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/gender
	Gender = prefix200505 + "gender"
	// GivenName is The URI for a claim that specifies the given name of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname
	GivenName = prefix200505 + "givenname"
	// GroupSid is The URI for a claim that specifies the SID for the group of an entity,. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/groupsid
	GroupSid = prefix200806ms + "groupsid"
	// Hash is The URI for a claim that specifies a hash value,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/hash
	Hash = prefix200505 + "hash"
	// HomePhone is The URI for a claim that specifies the home phone number of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/homephone
	HomePhone = prefix200505 + "homephone"
	// IsPersistent is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/ispersistent
	IsPersistent = prefix200806ms + "ispersistent"
	// Locality is The URI for a claim that specifies the locale in which an entity resides,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/locality
	Locality = prefix200505 + "locality"
	// MobilePhone is The URI for a claim that specifies the mobile phone number of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/mobilephone
	MobilePhone = prefix200505 + "mobilephone"
	// Name is The URI for a claim that specifies the name of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name
	Name = prefix200505 + "name"
	// NameIdentifier is The URI for a claim that specifies the name of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier
	NameIdentifier = prefix200505 + "nameidentifier"
	// OtherPhone is The URI for a claim that specifies the alternative phone number of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/otherphone
	OtherPhone = prefix200505 + "otherphone"
	// PostalCode is The URI for a claim that specifies the postal code of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/postalcode
	PostalCode = prefix200505 + "postalcode"
	// PrimaryGroupSid is The URI for a claim that specifies the primary group SID of an entity,. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/primarygroupsid
	PrimaryGroupSid = prefix200806ms + "primarygroupsid"
	// PrimarySid is The URI for a claim that specifies the primary SID of an entity,. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/primarysid
	PrimarySid = prefix200806ms + "primarysid"
	// Role is The URI for a claim that specifies the role of an entity,. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/role
	Role = prefix200806ms + "role"
	// Rsa is The URI for a claim that specifies an RSA key,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/rsa
	Rsa = prefix200505 + "rsa"
	// SerialNumber is The URI for a claim that specifies a serial number,. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/serialnumber
	SerialNumber = prefix200806ms + "serialnumber"
	// Sid is The URI for a claim that specifies a security identifier (SID),. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/sid
	Sid = prefix200505 + "sid"
	// Spn is The URI for a claim that specifies a service principal name (SPN) claim,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/spn
	Spn = prefix200505 + "spn"
	// StateOrProvince is The URI for a claim that specifies the state or province in which an entity resides,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/stateorprovince
	StateOrProvince = prefix200505 + "stateorprovince"
	// StreetAddress is The URI for a claim that specifies the street address of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/streetaddress
	StreetAddress = prefix200505 + "streetaddress"
	// Surname is The URI for a claim that specifies the surname of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname
	Surname = prefix200505 + "surname"
	// System is The URI for a claim that identifies the system entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/system
	System = prefix200505 + "system"
	// Thumbprint is The URI for a claim that specifies a thumbprint,. A thumbprint is a globally unique SHA-1 hash of an X.509 certificate". Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/thumbprint
	Thumbprint = prefix200505 + "thumbprint"
	// Upn is The URI for a claim that specifies a user principal name (UPN),. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/upn
	Upn = prefix200505 + "upn"
	// URI is The URI for a claim that specifies a URI,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/uri
	URI = prefix200505 + "uri"
	// UserData is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/userdata
	UserData = prefix200806ms + "userdata"
	// Version is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/version
	Version = prefix200806ms + "version"
	// Webpage is The URI for a claim that specifies the webpage of an entity,. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/webpage
	Webpage = prefix200505 + "webpage"
	// WindowsAccountName is The URI for a claim that specifies the Windows domain account name of an entity,. Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/windowsaccountname
	WindowsAccountName = prefix200806ms + "windowsaccountname"
	// WindowsDeviceClaim is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/windowsdeviceclaim
	WindowsDeviceClaim = prefix200806ms + "windowsdeviceclaim"
	// WindowsDeviceGroup is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/windowsdevicegroup
	WindowsDeviceGroup = prefix200806ms + "windowsdevicegroup"
	// WindowsFqbnVersion is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/windowsfqbnversion
	WindowsFqbnVersion = prefix200806ms + "windowsfqbnversion"
	// WindowsSubAuthority is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/windowssubauthority
	WindowsSubAuthority = prefix200806ms + "windowssubauthority"
	// WindowsUserClaim is . Its value is http://schemas.microsoft.com/ws/2008/06/identity/claims/windowsuserclaim
	WindowsUserClaim = prefix200806ms + "windowsuserclaim"
	// X500DistinguishedName is The URI for a distinguished name claim of an X.509 certificate, . The X.500 standard defines the methodology for defining distinguished names that are used by X.509 certificates. Its value is http://schemas.xmlsoap.org/ws/2005/05/identity/claims/x500distinguishedname
	X500DistinguishedName = prefix200505 + "x500distinguishedname"
)
