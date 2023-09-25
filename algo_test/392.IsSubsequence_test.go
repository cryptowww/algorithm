package algo_test

import (
	"cryptowww.leetcode/algo"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	if rv := algo.IsSubsequence("st", "asatt"); rv != true {
		t.Errorf("error")
	}
}
