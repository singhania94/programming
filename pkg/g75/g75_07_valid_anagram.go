package main

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]uint)
	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	for _, c := range t {
		if count, ok := m[c]; !ok {
			return false
		} else if count == 1 {
			delete(m, c)
		} else {
			m[c]--
		}
	}

	if len(m) == 0 {
		return true
	}
	return false
}
