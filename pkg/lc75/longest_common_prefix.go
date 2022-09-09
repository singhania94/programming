package main

func LongestCommonPrefix(strs []string) string {
	lcp := ""
	if len(strs) == 0 {
		return lcp
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for _, str := range strs {
			if i == len(str) || c != str[i] {
				return lcp
			}
		}
		lcp = lcp + string(c)
	}
	return lcp
}
