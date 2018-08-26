package claims

import (
	"reflect"
	"testing"
)

func TestNewPrincipal(t *testing.T) {
	type args struct {
		identities []Identity
	}
	tests := []struct {
		name string
		args args
		want Principal
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrincipal(tt.args.identities...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrincipal() = %v, want %v", got, tt.want)
			}
		})
	}
}
