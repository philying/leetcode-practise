package leetcode_practise

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReserve(t *testing.T) {
	in := 1534236469
	expected := 0
	actual := reverse(in)
	if !reflect.DeepEqual(actual, expected) {
		if actual != expected {
			//fmt.Printf("in: %d  actual: %d\n", in, actual)
			t.Errorf("reverse(%d) = %d; expected %d", in, actual, expected)
		} else {
			fmt.Println("success")
		}
	}
}
