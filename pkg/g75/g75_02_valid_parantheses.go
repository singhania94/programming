package main

func IsValid(s string) bool {
	var stack string
	top := -1
	for _, c := range s {
		char := string(c)
		if char == "(" || char == "{" || char == "[" {
			top++
			stack = stack + char
		} else if top >= 0 {
			if char == ")" && string(stack[top]) == "(" ||
				char == "}" && string(stack[top]) == "{" ||
				char == "]" && string(stack[top]) == "[" {
				stack = stack[:top]
				top--
			}
		} else {
			return false
		}
	}
	if top == -1 {
		return true
	}
	return false
}
