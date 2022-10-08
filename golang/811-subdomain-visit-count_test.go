package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	type args struct {
		cpdomains []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{
			name: "test01",
			args: args{
				cpdomains: []string{"9001 discuss.leetcode.com"},
			},
			wantRes: []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantRes, subdomainVisits(tt.args.cpdomains), "subdomainVisits(%v)", tt.args.cpdomains)
		})
	}
}
