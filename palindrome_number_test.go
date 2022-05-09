package leetcode

import (
	"strconv"
	"testing"
)

func isPalindrome(x int) bool {
	num := strconv.Itoa(x)

	j := 0
	for i := len(num) - 1; i > 0; i-- {
		if num[i] != num[j] {
			return false
		}
		j++
	}
	return true
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "example 3",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
