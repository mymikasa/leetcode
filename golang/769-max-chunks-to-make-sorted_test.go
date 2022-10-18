package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxChunksToSorted(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{
				arr: []int{1, 0, 2, 3, 4},
			},
			want: 4,
		},
		{
			name: "test02",
			args: args{
				arr: []int{1, 4, 0, 2, 3, 5},
			},
			want: 2,
		},
		{
			name: "test03",
			args: args{
				arr: []int{0, 2, 1, 4, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxChunksToSorted(tt.args.arr), "maxChunksToSorted(%v)", tt.args.arr)
		})
	}
}
