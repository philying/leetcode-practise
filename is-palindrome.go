package leetcode

func isPalindrome(x int) bool {
	num := 0
	reverse := x
	if x < 0 {
		return false
	}
	for {
		num = num*10 + reverse%10
		reverse /= 10
		if reverse == 0 {
			break
		}
	}
	if num == x {
		return true
	}
	return false
}
