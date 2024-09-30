package leetcode

import "fmt"

// remove duplicate numbers from nums
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 初始化慢指针 i
	i := 0
	fmt.Println("origin:", nums)
	// 快指针 j 从第二个元素开始
	//[1 1 2]
	for j := 1; j < len(nums); j++ {
		// 如果发现新元素，将其放到 i 的下一个位置
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println("format:", nums)
	// 返回唯一元素的数量
	return i + 1
}
