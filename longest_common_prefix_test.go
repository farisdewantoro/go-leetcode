package leetcode

import (
	"testing"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	firstLength := len(strs[0])
	result := ""

	for i := 0; i < firstLength; i++ {
		var p byte
		for _, v := range strs {
			if len(v)-1 < i {
				return result
			}
			if p == 0 {
				p = v[i]
			} else if p == v[i] {
				continue
			} else {
				return result
			}
		}
		if p != 0 {
			result += string(p)

		}
	}

	return result
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "no common prefix",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
