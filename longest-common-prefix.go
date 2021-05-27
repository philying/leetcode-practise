package leetcode_practise

import "fmt"

func longestCommonPrefix(strs []string) string {
	result := ""
	newStr := ""
	if len(strs) >= 2 {
		// compare 0 and 1
		minS := len(strs[0])
		if len(strs[1]) < minS {
			minS = len(strs[1])
		}
		s0 := strs[0]
		s1 := strs[1]
		for i:=0; i<minS; i++{
			if s0[i:i+1] == s1[i:i+1] {
				newStr = newStr + s0[i:i+1]
			} else {
				if len(newStr) == 0 {
					return ""
				}
				break
			}
		}
		fmt.Println("newStr is :" + newStr)
		if len(strs) == 2 {
			result = newStr
		}
		// compare big or small
		strLen := len(newStr)
		for m:=2; m<len(strs); m++{
			s := strs[m]
			if s == "" {
				return ""
			}
			if strLen > len(s) {
				strLen = len(s)
			}
			result = ""
			for j:=0; j<strLen; j++ {
				if s[j:j+1] == newStr[j:j+1] {
					result = result + s[j:j+1]
				} else {
					return result
				}
			}
		}
	} else {
		if len(strs) == 1 {
			return strs[0]
		} else {
			return ""
		}
	}
	return result
}