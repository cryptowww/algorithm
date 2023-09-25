package algo_test

import (
	"cryptowww.leetcode/algo"
	"fmt"
	"testing"
)

type Cases struct {
	name         string
	originStr    string
	goalStr      string
	canTransfrom bool
}

var cases [2]Cases

func setup() {
	cases = [...]Cases{
		{"rs", "abcde", "cdeab", true},
		{"rs", "wmkwmk", "wmkwmk", true},
	}
}

func TestRotateString(t *testing.T) {
	setup()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if rv := algo.RotateString(c.originStr, c.goalStr); rv != c.canTransfrom {
				t.Errorf("%s -> %s,  %t", c.originStr, c.goalStr, rv)
			} else {
				fmt.Println("good")
			}
		})
	}
}

func TestRotateString2(t *testing.T) {

}
