package main

import (
	"reflect"
	"testing"
)

func Test_reverseOddLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case1",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOddLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseOddLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
