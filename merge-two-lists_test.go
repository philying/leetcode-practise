package leetcode

import (
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	one []int
	two []int
}

type ans struct {
	one []int
}

func TestMergeTwoLists(t *testing.T) {
}

func l2s(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}
	return res
}
