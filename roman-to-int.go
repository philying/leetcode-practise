package leetcode_practise

import "fmt"

func romanToInt(s string) int {
	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	m["IV"] = 4
	m["IX"] = 9
	m["XL"] = 40
	m["XC"] = 90
	m["CD"] = 400
	m["CM"] = 900
	sArray := []byte(s)
	result := 0
	length := len(sArray)
	for i := 0; i < length; i++ {
		var w = string(sArray[i])
		fmt.Println(w, "====>", w)
		if w == "I" {
			if (i + 1) < length && (string(sArray[i+1]) == "V" || string(sArray[i+1]) == "X") {
				result -= 1
			} else {
				result += 1
			}
		} else if w == "X" {
			if (i + 1) < length && (string(sArray[i+1]) == "L" || string(sArray[i+1]) == "C") {
				result -= 10
			} else {
				result += 10
			}
		} else if w == "C" {
			if (i + 1) < length && (string(sArray[i+1]) == "D" || string(sArray[i+1]) == "M") {
				result -= 100
			} else {
				result += 100
			}
		} else {
			result += m[w]
		}
	}
	return result
}