package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValidParentheses(s string) bool {
	sp := []rune{}

	for _, v := range s {
		if len(sp) == 0 {
			sp = append(sp, v)
			continue
		}

		peek := sp[len(sp)-1]

		if (v == '}' && peek == '{') ||
			(v == ')' && peek == '(') ||
			(v == ']' && peek == '[') {
			sp = sp[:len(sp)-1]
		} else {
			sp = append(sp, v)
		}

	}

	return len(sp) == 0
}

func Test_isValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			args: args{
				s: "([[{]])",
			},
			want: false,
		},
		{
			args: args{
				s: "({}[{}])",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValidParentheses(tt.args.s), tt.args.s)
		})
	}
}
