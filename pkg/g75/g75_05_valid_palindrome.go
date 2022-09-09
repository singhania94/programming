package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	fmt.Printf("s: %v\n", s)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] < 48 || (s[i] > 57 && s[i] < 97) || s[i] > 122 {
			i++
		} else if s[j] < 48 || (s[j] > 57 && s[j] < 97) || s[j] > 122 {
			j--
		} else {
			if s[i] != s[j] {
				fmt.Printf("i: %v\n", i)
				fmt.Printf("j: %v\n", j)
				return false
			}
			i++
			j--
		}
	}
	return true
}
