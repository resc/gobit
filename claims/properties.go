package claims

import (
	"sort"
	"strings"
)

type (
	// Properties contains additional claim properties.
	Properties interface {
		Len() int
		All() func() (key string, value string, ok bool)
		Get(key string) (value string, ok bool)
		Set(key, value string)
		Del(key string) bool
	}

	properties struct {
		props []property
	}

	property struct {
		key, value string
	}
)

// NewProperties returns a new instance of the Properties interface
func NewProperties() Properties {
	return &properties{}
}

func (p *properties) Len() int {
	return len(p.props)
}

func (p *properties) All() func() (key string, value string, ok bool) {
	i := -1
	return func() (string, string, bool) {
		if i >= len(p.props)-1 {
			return "", "", false
		}
		i++
		return p.props[i].key, p.props[i].value, true
	}
}

func (p *properties) Get(key string) (string, bool) {
	i := sort.Search(p.Len(), func(i int) bool {
		return p.compare(p.props[i].key, key) == 0
	})
	if i < p.Len() {
		return p.props[i].value, true
	}
	return "", false
}

func (p *properties) Set(key, value string) {
	i := sort.Search(p.Len(), func(i int) bool {
		return p.compare(p.props[i].key, key) >= 0
	})
	if i == p.Len() {
		p.props = append(p.props, property{key, value})
	} else {
		if p.compare(p.props[i].key, key) == 0 {
			p.props[i].value = value
		} else {
			p.props = append(p.props, property{})
			copy(p.props[i+1:], p.props[i:])
			p.props[i] = property{key, value}
		}
	}
}

func (p *properties) compare(key1, key2 string) int {
	return strings.Compare(key1, key2)
}

func (p *properties) Del(key string) bool {
	lastIndex := len(p.props) - 1
	i := sort.Search(p.Len(), func(i int) bool {
		return p.compare(p.props[i].key, key) == 0
	})

	if i > lastIndex {
		return false
	}
	if i < lastIndex {
		copy(p.props[i:], p.props[i+1:])
	}
	p.props[lastIndex] = property{}
	p.props = p.props[:lastIndex]
	return true
}
