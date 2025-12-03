package main

import "fmt"

// isMatchV1 implements regular expression matching with '.' and '*'
// '.' matches any single character
// '*' matches zero or more of the preceding element
// Time complexity: O(2^n)
// Space complexity: O(n)
func isMatchV1(s string, p string) bool {
	// Base case: if pattern is empty, string must also be empty
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatchV1(s, p[2:]) || (firstMatch && isMatchV1(s[1:], p))
	}

	// No star: must match first character and continue with rest
	return firstMatch && isMatchV1(s[1:], p[1:])
}

// v2

func main() {
	fmt.Println(isMatchV1("aaaa", "a*"))
}
