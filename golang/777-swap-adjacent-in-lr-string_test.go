package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canTransform(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				start: "RXXLRXRXL",
				end:   "XRLXXRRLX",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				start: "RXXLRXRXL",
				end:   "XRRXXRRLX",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				start: "R",
				end:   "X",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				start: "RL",
				end:   "LR",
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				start: "LXXLXRLXXL",
				end:   "XLLXRXLXLX",
			},
			want: false,
		},
		{
			name: "test6",
			args: args{
				start: "XXXXXLXXXX",
				end:   "LXXXXXXXXX",
			},
			want: true,
		},
		{
			name: "test7",
			args: args{
				start: "XXXLXXRXXXX",
				end:   "LXXXXXXXXXR",
			},
			want: true,
		},
		{
			name: "test8",
			args: args{
				start: "XLXRRXXRXX",
				end:   "LXXXXXXRRR",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canTransform(tt.args.start, tt.args.end), "canTransform(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}
