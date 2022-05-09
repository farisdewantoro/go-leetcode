package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPowerOfTwo(n int) bool {
	return (n&(n-1) == 0) && n != 0
}

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			args: args{
				n: 14,
			},
			want: false,
		},
		{
			args: args{
				n: 12,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPowerOfTwo(tt.args.n)
			assert.Equal(t, tt.want, got, tt.args)
		})
	}
}
