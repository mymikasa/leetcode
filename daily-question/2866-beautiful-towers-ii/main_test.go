package main

import (
	"reflect"
	"testing"
)

func Test_maximumSumOfHeights(t *testing.T) {
	type args struct {
		maxHeights []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case1",
			args: args{
				maxHeights: []int{5, 3, 4, 1, 1},
			},
			want: 13,
		},
		{
			name: "case2",
			args: args{
				maxHeights: []int{3, 5, 3, 5, 1, 5, 4, 4, 4},
			},
			want: 22,
		},
		{
			name: "case3",
			args: args{
				maxHeights: []int{6, 5, 2, 1, 5, 4, 4, 2},
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSumOfHeights(tt.args.maxHeights); got != tt.want {
				t.Errorf("maximumSumOfHeights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prefixSum(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int64
		want1 []int64
	}{
		{
			name: "case1",
			args: args{
				heights: []int{6, 5, 2, 1, 5, 4, 4, 2},
			},
			want:  []int64{6, 10, 6, 4, 9, 9, 10, 9},
			want1: []int64{18, 12, 7, 5, 15, 10, 6, 2},
		},
		{
			name: "case2",
			args: args{
				heights: []int{1, 2, 3, 4, 5, 3},
			},
			want:  []int64{1, 3, 6, 10, 15, 15},
			want1: []int64{18, 12, 7, 5, 15, 10, 6, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := prefixSum(tt.args.heights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixSum() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("prefixSum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
