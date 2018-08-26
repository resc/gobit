package claims

import (
	"reflect"
	"testing"
)

func Test_claim_Clone(t *testing.T) {
	tests := []struct {
		name string
		c    *claim
		want ClaimBuilder
	}{
		{"clone", NewClaim("name", "value").(*claim), NewClaim("name", "kees")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("claim.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}
