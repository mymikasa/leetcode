package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_advantageCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "01",
			args: args{
				nums1: []int{12, 24, 8, 32},
				nums2: []int{13, 25, 32, 11},
			},
			want: []int{24, 32, 8, 12},
		},
		{
			name: "02",
			args: args{
				nums1: []int{2, 0, 4, 1, 2},
				nums2: []int{1, 3, 0, 0, 2},
			},
			want: []int{2, 0, 2, 1, 4},
		},
		{
			name: "03",
			args: args{
				nums1: []int{15448, 14234, 13574, 19893, 6475},
				nums2: []int{14234, 6475, 19893, 15448, 13574},
			},
			want: []int{15448, 13574, 6475, 19893, 14234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, advantageCount(tt.args.nums1, tt.args.nums2), "advantageCount(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
