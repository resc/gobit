package claims

type (
	// Identities is a collection of identities
	Identities interface {
		Len() int
		All() func() (Identity, bool)
		FindFirst(func(Identity) bool) (Identity, bool)
		FindAll(func(Identity) bool) []Identity
	}

	identities struct {
		items []Identity
	}
)

// NewIdentities returns a new identities collection
func NewIdentities() Identities {
	return newIdentities()
}

func newIdentities() *identities {
	return &identities{}
}

func (c *identities) Len() int {
	return len(c.items)
}

func (c *identities) FirstOrDefault() Identity {
	if len(c.items) < 1 {
		return nil
	}
	return c.items[0]
}

func (c *identities) All() func() (Identity, bool) {
	i := -1
	return func() (Identity, bool) {
		if i >= len(c.items)-1 {
			return nil, false
		}
		i++
		return c.items[i], true
	}
}

func (c *identities) FindFirst(f func(Identity) bool) (Identity, bool) {
	for i := range c.items {
		if f(c.items[i]) {
			return c.items[i], true
		}
	}
	return nil, false
}

func (c *identities) FindAll(f func(Identity) bool) []Identity {
	result := make([]Identity, 0, 4)
	for i := range c.items {
		if f(c.items[i]) {
			result = append(result, c.items[i])
		}
	}
	return result
}

func (c *identities) Add(identity Identity) {
	if identity == nil {
		return
	}
	c.items = append(c.items, identity)
}

func (c *identities) Remove(identity Identity) {
	if identity == nil {
		return
	}
	for i := range c.items {
		if c.items[i] == identity {
			a := c.items
			copy(a[i:], a[i+1:])
			a[len(a)-1] = nil
			c.items = a[:len(a)-1]
			return
		}
	}
}
