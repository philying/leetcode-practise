package leetcode_practise

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	in := "(){}{}"
	expected := true
	actual := isValid(in)
	if actual != expected {
		fmt.Println("error : actual is excepted is ", actual, expected)
	} else {
		fmt.Println("success!")
	}
}
