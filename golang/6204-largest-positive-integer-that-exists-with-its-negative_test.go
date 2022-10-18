package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxK(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				[]int{-10, 8, 6, 7, -2, -3},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMaxK(tt.args.nums), "findMaxK(%v)", tt.args.nums)
		})
	}
}
