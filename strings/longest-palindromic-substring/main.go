package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// Brute force approach
// O(n^3) time complexity
func longestPalindromeV1(s string) string {
	longestPalindrome := ""
	if isPalindrome(s) {
		return s
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) {
				if j-i > len(longestPalindrome) {
					longestPalindrome = s[i:j]
				}
			}
		}
	}

	return longestPalindrome
}

// Expand around center approach
// O(n^2) time complexity
// O(1) space complexity
// Checks both odd-length (single center) and even-length (two centers) palindromes
func longestPalindromeV2(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLen := 0, 1

	// Helper function to expand around center and return palindrome length
	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		// Return length of palindrome found
		return right - left - 1
	}

	// Check every possible center
	for i := 0; i < len(s); i++ {
		// Check odd-length palindromes (single center at i)
		len1 := expandAroundCenter(i, i)
		// Check even-length palindromes (center between i and i+1)
		len2 := expandAroundCenter(i, i+1)

		// Get the maximum length from both checks
		currentLen := len1
		if len2 > len1 {
			currentLen = len2
		}

		// Update if we found a longer palindrome
		if currentLen > maxLen {
			maxLen = currentLen
			// Calculate the starting position
			start = i - (currentLen-1)/2
		}
	}

	return s[start : start+maxLen]
}

func main() {
	fmt.Println(longestPalindromeV2("aabaad"))
}
