package leetcode

func isValid(s string) bool {

	type Stack struct {
		items []string
	}

	stack := &Stack{}

	m := make(map[string]int)
	m["["] = 0
	m["]"] = 0
	m["{"] = 1
	m["}"] = 1
	m["("] = 2
	m[")"] = 2
	sort := make(map[string]string)
	sort["["] = "left"
	sort["]"] = "right"
	sort["{"] = "left"
	sort["}"] = "right"
	sort["("] = "left"
	sort[")"] = "right"

	for i := 0; i < len(s); i++ {
		sub := s[i : i+1]
		if len(stack.items) > 0 {
			item := stack.items[len(stack.items)-1]
			if m[item] == m[sub] && sort[item] == "left" && sort[sub] == "right" {
				//pop the top of the stack
				stack.items = stack.items[0 : len(stack.items)-1]
			} else {
				//push sub string to stack
				stack.items = append(stack.items, sub)
			}
		} else {
			//push sub string to stack
			stack.items = append(stack.items, sub)
		}
	}

	if len(stack.items) == 0 {
		return true
	}

	return false
}
