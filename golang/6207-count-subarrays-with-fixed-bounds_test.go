package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		minK int
		maxK int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		{
			name: "test1",
			args: args{
				[]int{1, 3, 5, 2, 7, 5},
				1,
				5,
			},
			wantRes: 2,
		},
		{
			name: "test2",
			args: args{
				[]int{1, 1, 1, 1},
				1,
				1,
			},
			wantRes: 10,
		},
		{
			name: "test3",
			args: args{
				[]int{35054, 398719, 945315, 945315, 820417, 945315, 35054, 945315, 171832, 945315, 35054, 109750, 790964, 441974, 552913},
				35054,
				945315,
			},
			wantRes: 81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantRes, countSubarrays(tt.args.nums, tt.args.minK, tt.args.maxK), "countSubarrays(%v, %v, %v)", tt.args.nums, tt.args.minK, tt.args.maxK)
		})
	}
}
