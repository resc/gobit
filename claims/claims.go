package claims

type (
	// Claims is a collection of claims
	Claims interface {
		Len() int
		All() func() (Claim, bool)
		FindFirst(func(Claim) bool) (Claim, bool)
		FindAll(func(Claim) bool) []Claim
	}
	claims struct {
		items []Claim
	}
)

// NewClaims returns a new claims collection
func NewClaims() Claims {
	return newClaims()
}

// newClaims returns a new claims collection
func newClaims() *claims {
	return &claims{}
}

func (c *claims) Len() int {
	return len(c.items)
}
func (c *claims) All() func() (Claim, bool) {
	i := -1
	return func() (Claim, bool) {
		if i >= len(c.items)-1 {
			return nil, false
		}
		i++
		return c.items[i], true
	}
}
func (c *claims) FindFirst(f func(Claim) bool) (Claim, bool) {
	for i := range c.items {
		if f(c.items[i]) {
			return c.items[i], true
		}
	}
	return nil, false
}
func (c *claims) FindAll(f func(Claim) bool) []Claim {
	result := make([]Claim, 0, 4)
	for i := range c.items {
		if f(c.items[i]) {
			result = append(result, c.items[i])
		}
	}
	return result
}

func (c *claims) Add(claim Claim) {
	if claim == nil {
		return
	}
	c.items = append(c.items, claim)
}
func (c *claims) Remove(claim Claim) {
	if claim == nil {
		return
	}
	for i := range c.items {
		if c.items[i] == claim {
			a := c.items
			copy(a[i:], a[i+1:])
			a[len(a)-1] = nil
			c.items = a[:len(a)-1]
			return
		}
	}
}
