package leetcode

import "testing"


func TestTrie(t *testing.T) {
	obj := Constructor()
	obj.Insert("apple")
	if !obj.Search("apple") {
		t.Fail()
	}
}