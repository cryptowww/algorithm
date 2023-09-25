package algo_test

import (
	"cryptowww.leetcode/algo"
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {

	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{"roman", "XXVII", 27},
		{"roman", "LVIII", 58},
		{"roman", "MCMXCIV", 1994},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			if rv := algo.RomanToInt(c.input); rv != c.expected {
				t.Errorf("%s should be transformed %d, but got %d", c.input, c.expected, rv)
			} else {
				fmt.Printf("%s is transformed to %d\n", c.input, c.expected)
			}
		})
	}
}
