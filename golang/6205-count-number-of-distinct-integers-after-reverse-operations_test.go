package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countDistinctIntegers(t *testing.T) {
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
				[]int{2, 2, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countDistinctIntegers(tt.args.nums), "countDistinctIntegers(%v)", tt.args.nums)
		})
	}
}
