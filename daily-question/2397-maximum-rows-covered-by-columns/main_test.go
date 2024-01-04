package main

import "testing"

func Test_maximumRows(t *testing.T) {
	type args struct {
		matrix    [][]int
		numSelect int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				matrix:    [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}},
				numSelect: 2,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumRows(tt.args.matrix, tt.args.numSelect); gotAns != tt.wantAns {
				t.Errorf("maximumRows() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
