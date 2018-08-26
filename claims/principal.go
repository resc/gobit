package claims

type (
	// Principal contains user identifcation info and authorization status
	Principal interface {
		Identity() Identity
		Identities() Identities
		Claims() Claims
		IsInRole(role string) bool
	}

	principal struct {
		identities *identities
	}
)

//NewPrincipal creates a new principal from one or more identities
func NewPrincipal(identities ...Identity) Principal {
	p := &principal{
		identities: newIdentities(),
	}
	for i := range identities {
		p.identities.Add(identities[i])
	}
	return p
}

func (p *principal) Identity() Identity {
	return p.identities.FirstOrDefault()
}

func (p *principal) Identities() Identities {
	return p.identities
}

func (p *principal) Claims() Claims {
	return p
}

func (p *principal) IsInRole(role string) bool {
	for ii := range p.identities.items {
		id := p.identities.items[ii]
		next := id.Claims().All()
		for c, ok := next(); ok; c, ok = next() {
			if c.Type() == id.RoleClaimType() && c.Value() == role {
				return true
			}
		}
	}

	return false
}

func (p *principal) Len() int {
	count := 0
	next := p.Identities().All()
	for id, ok := next(); ok; id, ok = next() {
		count += id.Claims().Len()
	}
	return count
}

func (p *principal) All() func() (Claim, bool) {
	nextID := p.Identities().All()
	id, idOk := nextID()

	nextClaim := (func() (Claim, bool))(nil)
	claim, claimOk := Claim(nil), false

	if idOk {
		nextClaim = id.Claims().All()
	}

	return func() (Claim, bool) {
		if !idOk {
			return Claim(nil), false
		}

		for {
			claim, claimOk = nextClaim()
			if claimOk {
				return claim, claimOk
			}

			id, idOk = nextID()
			if idOk {
				nextClaim = id.Claims().All()
				continue
			} else {
				return Claim(nil), false
			}

		}
	}
}

func (p *principal) FindFirst(f func(Claim) bool) (Claim, bool) {
	next := p.All()
	for c, ok := next(); ok; c, ok = next() {
		if f(c) {
			return c, true
		}
	}
	return nil, false
}

func (p *principal) FindAll(f func(Claim) bool) []Claim {
	result := make([]Claim, 0, 4)
	next := p.All()
	for c, ok := next(); ok; c, ok = next() {
		if f(c) {
			result = append(result, c)
		}
	}
	return result
}
