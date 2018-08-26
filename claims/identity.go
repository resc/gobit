package claims

import "github.com/resc/gobit/claims/claimtypes"

type (
	// Identity contains user identification info and authentication status
	Identity interface {
		Name() string
		AuthenticationType() string
		Actor() Identity
		Claims() Claims
		IsAuthenticated() bool
		Label() string
		SetLabel(string)
		NameClaimType() string
		RoleClaimType() string
		Clone() IdentityBuilder
	}

	// IdentityBuilder is used to create and Identity
	IdentityBuilder interface {
		WithAuthenticationType(string) IdentityBuilder
		WithActor(Identity) IdentityBuilder
		WithClaim(Claim) IdentityBuilder
		WithLabel(string) IdentityBuilder
		WithNameClaimType(string) IdentityBuilder
		WithRoleClaimType(string) IdentityBuilder
		Build() Identity
	}

	identity struct {
		actor              Identity
		authenticationType string
		label              string
		name               string
		nameClaimType      string
		roleClaimType      string
		claims             *claims
	}
)

// BuildIdentity creates a new IdentityBuilder instance
func BuildIdentity(name string) IdentityBuilder {
	return &identity{
		name:          name,
		nameClaimType: claimtypes.Name,
		roleClaimType: claimtypes.Role,
		claims:        newClaims(),
	}
}

// NewIdentity creates a new Identity instance
func NewIdentity(name, authenticationType string, actor Identity, claims ...Claim) Identity {
	b := BuildIdentity(name).
		WithAuthenticationType(authenticationType).
		WithActor(actor)
	for i := range claims {
		b.WithClaim(claims[i])
	}
	return b.Build()
}

func (id *identity) Build() Identity {
	return id
}

func (id *identity) Name() string {
	return id.name
}

func (id *identity) AuthenticationType() string {
	return id.authenticationType
}

func (id *identity) WithAuthenticationType(authenticationType string) IdentityBuilder {
	id.authenticationType = authenticationType
	return id
}

func (id *identity) Actor() Identity {
	return id.actor
}

func (id *identity) WithActor(actor Identity) IdentityBuilder {
	id.actor = actor
	return id
}

func (id *identity) Claims() Claims {
	return id.claims
}

func (id *identity) WithClaim(c Claim) IdentityBuilder {
	id.claims.Add(c)
	return id
}

func (id *identity) IsAuthenticated() bool {
	return id.authenticationType != ""
}

func (id *identity) Label() string {
	return id.label
}

func (id *identity) SetLabel(label string) {
	id.label = label
}

func (id *identity) WithLabel(label string) IdentityBuilder {
	id.SetLabel(label)
	return id
}

func (id *identity) NameClaimType() string {
	if id.nameClaimType == "" {
		return claimtypes.Name
	}
	return id.nameClaimType
}

func (id *identity) WithNameClaimType(nameClaimType string) IdentityBuilder {
	if nameClaimType == "" {
		id.nameClaimType = claimtypes.Name
	} else {
		id.nameClaimType = nameClaimType
	}
	return id
}

func (id *identity) RoleClaimType() string {
	if id.roleClaimType == "" {
		return claimtypes.Role
	}
	return id.roleClaimType
}

func (id *identity) WithRoleClaimType(roleClaimType string) IdentityBuilder {
	if roleClaimType == "" {
		id.roleClaimType = claimtypes.Role
	} else {
		id.roleClaimType = roleClaimType
	}
	return id
}

func (id *identity) Clone() IdentityBuilder {
	clone := *id
	claims := newClaims()
	clone.claims = claims
	for i := range id.claims.items {
		claims.Add(id.claims.items[i].Clone().Build())
	}
	return &clone
}
