package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countStudents(t *testing.T) {
	type args struct {
		students   []int
		sandwiches []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{
				[]int{0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1},
				[]int{1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countStudents(tt.args.students, tt.args.sandwiches), "countStudents(%v, %v)", tt.args.students, tt.args.sandwiches)
		})
	}
}
