package leetcode

import (
	"fmt"
	"testing"
)


func Test443(*testing.T) {
	bytes := []byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}
	compress(bytes)

	fmt.Println(bytes)
}