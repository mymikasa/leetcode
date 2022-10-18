package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numComponents(t *testing.T) {
	type args struct {
		head *ListNode
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
				head: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
				nums: []int{0, 1, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numComponents(tt.args.head, tt.args.nums), "numComponents(%v, %v)", tt.args.head, tt.args.nums)
		})
	}
}
