package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minAddToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{
				"())",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minAddToMakeValid(tt.args.s), "minAddToMakeValid(%v)", tt.args.s)
		})
	}
}
