package main

import "testing"

func Test_minimumMountainRemovals(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{2, 1, 1, 5, 6, 2, 3, 1},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				nums: []int{1, 3, 1},
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{nums: []int{9, 8, 1, 7, 6, 5, 4, 3, 2, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMountainRemovals(tt.args.nums); got != tt.want {
				t.Errorf("minimumMountainRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
