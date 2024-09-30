package leetcode

func twoSum(nums []int, target int) []int {
	sm := map[int]int{}
	for i, x := range nums {
		value, isExist := sm[target-x]
		if isExist {
			return []int{value, i}
		}
		sm[x] = i
	}
	return nil
}
