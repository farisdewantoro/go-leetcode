package leetcode

import "testing"

func romanToInt(s string) int {
	var last int
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for _, v := range s {
		val := romanNumerals[v]
		result += val
		if last != 0 && val > last {
			result -= last * 2
		}
		last = val
	}
	return result
}

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name         string
		args         args
		want         int
		explaination string
	}{
		{
			name: "test 1",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "test 2",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "test 3",
			args: args{
				s: "MCMXCIV",
			},
			want:         1994,
			explaination: "M = 1000, CM = 900, XC = 90 and IV = 4.",
		},
		{
			name: "test 4",
			args: args{
				s: "IXL",
			},
			want: 39,
		},
		// + M 1000
		// + C 100
		// + M 1000
		// - C 100
		// - C 100
		// + X 10
		// + C 100
		// - X 10
		// - X 10
		// + I 1
		// + V 5
		// - I 1
		// - I 1

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
