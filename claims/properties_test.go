package claims

import (
	"testing"
)

func TestProperties(t *testing.T) {
	props := NewProperties()
	expected := []property{
		// keys sould be sorted here!
		property{"age", "my-age"},
		property{"name", "my-name"},
	}

	for i := range expected {
		props.Set(expected[i].key, expected[i].value)
	}

	next := props.All()
	i := -1
	for k, v, ok := next(); ok; k, v, ok = next() {
		i++
		p := expected[i]
		if p.key != k {
			t.Errorf("Expected key %s, got %s", p.key, k)
		}
		if p.value != v {
			t.Errorf("Expected value %s, got %s", p.value, v)
		}
	}
}
