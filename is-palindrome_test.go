package leetcode_practise

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	in := 12212
	expected := false
	actual := isPalindrome(in)
	if !reflect.DeepEqual(actual, expected) {
		if actual != expected {
			//fmt.Printf("in: %d  actual: %d\n", in, actual)
			t.Errorf("reverse(%d) = %t; expected %t", in, actual, expected)
		} else {
			fmt.Println("success")
		}
	}
}
