package claims

import "github.com/resc/gobit/claims/claimvaluetypes"

type (
	// Claim is an assertion about a user.
	Claim interface {
		Issuer() string
		OriginalIssuer() string
		Subject() Identity
		Type() string
		Value() string
		ValueType() string
		Properties() Properties
		Clone() ClaimBuilder
	}
	// ClaimBuilder is used to construct a new claim.
	ClaimBuilder interface {
		WithIssuer(string) ClaimBuilder
		WithOriginalIssuer(string) ClaimBuilder
		WithSubject(Identity) ClaimBuilder
		WithValueType(string) ClaimBuilder
		WithProperty(key, value string) ClaimBuilder
		Build() Claim
	}
	claim struct {
		typ            string
		value          string
		valueType      string
		issuer         string
		originalIssuer string
		subject        Identity
		properties     Properties
	}
)

const (
	// DefaultIssuer is the default value if no issuer is specified
	DefaultIssuer = "LOCAL AUTHORITY"
)

// NewClaim constructs a new claim
func NewClaim(typ, value string) ClaimBuilder {
	return &claim{
		typ:            typ,
		value:          value,
		issuer:         DefaultIssuer,
		originalIssuer: DefaultIssuer,
		valueType:      claimvaluetypes.String,
	}
}

func (c *claim) Type() string {
	return c.typ
}

func (c *claim) Value() string {
	return c.value
}

func (c *claim) Issuer() string {
	return c.issuer
}

func (c *claim) WithIssuer(issuer string) ClaimBuilder {
	c.issuer = issuer
	return c
}

func (c *claim) OriginalIssuer() string {
	return c.originalIssuer
}

func (c *claim) WithOriginalIssuer(originalIssuer string) ClaimBuilder {
	c.originalIssuer = originalIssuer
	return c
}

func (c *claim) Subject() Identity {
	return c.subject
}

func (c *claim) WithSubject(identity Identity) ClaimBuilder {
	c.subject = identity
	return c
}

func (c *claim) ValueType() string {
	return c.valueType
}

func (c *claim) WithValueType(valueType string) ClaimBuilder {
	if valueType == "" {
		valueType = claimvaluetypes.String
	}

	c.valueType = valueType
	return c
}

func (c *claim) Properties() Properties {
	if (c.properties) == nil {
		c.properties = NewProperties()
	}
	return c.properties
}

func (c *claim) WithProperty(key, value string) ClaimBuilder {
	c.Properties().Set(key, value)
	return c
}

func (c *claim) Build() Claim {
	return c
}

func (c *claim) Clone() ClaimBuilder {
	clone := *c
	return &clone
}
