package leetcode_practise

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	in := "MCMXCIV"
	expected := 1994
	actual := romanToInt(in)
	if !reflect.DeepEqual(actual, expected) {
		if actual != expected {
			//fmt.Printf("in: %d  actual: %d\n", in, actual)
			t.Errorf("romanToInt(%s) = %d; expected %d", in, actual, expected)
		} else {
			fmt.Println("success")
		}
	}
}