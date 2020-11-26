package leetcode

import "testing"

type test139Input struct {
	s string
	word []string
}

type test139Case struct {
	input test139Input
	output bool
}

func Test139(t *testing.T) {
	cases := []test139Case{
		{
			input: test139Input{"leetcode", []string{"leet", "code"}},
			output: true,
		},
		{
			input: test139Input{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}},
			output: false,
		},
	}

	for _, c := range cases {
		actual := wordBreak3(c.input.s, c.input.word)
		if  actual != c.output {
			t.Errorf("Input: %v, Output: %v, Expect: %v", c.input, actual, c.output)
		}
	}
}
