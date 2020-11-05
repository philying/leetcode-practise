package leetcode_practise

func twoSum(nums []int, target int) [2]int {
	var i, j = 0, 0
	var pose, posy = -1, -1
	var length = len(nums)
	var sm = make(map[int]int)
	for i = 0; i < length; i++ {
		if pose != -1 && posy != -1 {
			break
		}
		sub := target - nums[i]
		value, isExist := sm[sub]
		if isExist && ( nums[i] * 2 == target ) {
			posy = value
			pose = i
		} else {
			sm[sub] = i
		}
	}
	if pose == -1 && posy == -1 {
		for j = 0; j < length; j++ {
			value, ok := sm[nums[j]]
			if ok {
				if j != value {
					pose = value
					posy = j
					break
				}
			}
		}
	}
	var result = [2]int{posy,pose}
	return result
}

