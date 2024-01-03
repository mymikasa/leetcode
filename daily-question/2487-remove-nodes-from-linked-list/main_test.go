package main

import (
	"reflect"
	"testing"
)

func Test_removeNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "case1",
			args: args{
				head: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 13,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 8,
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 8,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
