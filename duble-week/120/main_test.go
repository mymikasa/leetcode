package main

import "testing"

func Test_incremovableSubarrayCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				[]int{1, 2, 3, 4},
			},
			want: 10,
		},
		{
			name: "case2",
			args: args{
				[]int{6, 5, 7, 8},
			},
			want: 7,
		},
		{
			name: "case3",
			args: args{
				[]int{8, 7, 6, 6},
			},
			want: 3,
		},
		{
			name: "case3",
			args: args{
				[]int{6, 7, 8, 5, 7, 2, 9, 8},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incremovableSubarrayCount(tt.args.nums); got != tt.want {
				t.Errorf("incremovableSubarrayCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
