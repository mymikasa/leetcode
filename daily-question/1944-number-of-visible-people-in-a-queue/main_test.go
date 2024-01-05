package main

import (
	"reflect"
	"testing"
)

func Test_canSeePersonsCount(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test case 1",
			args: args{
				heights: []int{10, 6, 8, 5, 11, 9},
			},
			want: []int{3, 1, 2, 1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canSeePersonsCount(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canSeePersonsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
