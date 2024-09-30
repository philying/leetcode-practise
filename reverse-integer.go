package leetcode

func reverse(x int) int {
	count := 0
	y := x
	for y != 0 {
		y /= 10
		count += 1
	}
	i := 0
	value := 0
	for i < count {
		num := x / pow(10, i) % 10
		value += num * pow(10, count-i-1)
		i += 1
	}
	max := pow(2, 31)
	min := max * -1
	if value > max || value < min {
		value = 0
	}
	return value
}

func pow(x, n int) int {
	ret := 1
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
