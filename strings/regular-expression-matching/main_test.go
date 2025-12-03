package main

import (
	"testing"
)

// Test regular expression matching (V1 - Recursive)
// Pattern: '.' matches any single character
// Pattern: '*' matches zero or more of the preceding element
func TestIsMatchV1(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		// Empty string tests
		{
			name:     "Both empty",
			s:        "",
			p:        "",
			expected: true,
		},
		{
			name:     "Empty string, pattern with a*",
			s:        "",
			p:        "a*",
			expected: true,
		},
		{
			name:     "Empty string, pattern with .*",
			s:        "",
			p:        ".*",
			expected: true,
		},
		{
			name:     "Empty string, non-empty pattern",
			s:        "",
			p:        "a",
			expected: false,
		},

		// Exact match tests
		{
			name:     "Exact match single char",
			s:        "a",
			p:        "a",
			expected: true,
		},
		{
			name:     "Exact match multiple chars",
			s:        "abc",
			p:        "abc",
			expected: true,
		},
		{
			name:     "No match different chars",
			s:        "a",
			p:        "b",
			expected: false,
		},

		// Dot (.) pattern tests
		{
			name:     "Dot matches single char",
			s:        "a",
			p:        ".",
			expected: true,
		},
		{
			name:     "Dot matches any char",
			s:        "b",
			p:        ".",
			expected: true,
		},
		{
			name:     "Multiple dots",
			s:        "abc",
			p:        "...",
			expected: true,
		},
		{
			name:     "Dot with exact chars",
			s:        "abc",
			p:        "a.c",
			expected: true,
		},

		// Star (*) pattern tests
		{
			name:     "a* matches empty",
			s:        "",
			p:        "a*",
			expected: true,
		},
		{
			name:     "a* matches single a",
			s:        "a",
			p:        "a*",
			expected: true,
		},
		{
			name:     "a* matches multiple a",
			s:        "aaaa",
			p:        "a*",
			expected: true,
		},
		{
			name:     "a* matches nothing (b following)",
			s:        "b",
			p:        "a*b",
			expected: true,
		},
		{
			name:     "a* with b* matches aab",
			s:        "aab",
			p:        "a*b*",
			expected: true,
		},

		// .* pattern tests
		{
			name:     ".* matches everything",
			s:        "anything",
			p:        ".*",
			expected: true,
		},
		{
			name:     ".* matches empty",
			s:        "",
			p:        ".*",
			expected: true,
		},
		{
			name:     ".* in middle",
			s:        "abcdef",
			p:        "a.*f",
			expected: true,
		},

		// Complex patterns
		{
			name:     "Complex: mississippi vs mis*is*p*.",
			s:        "mississippi",
			p:        "mis*is*p*.",
			expected: false,
		},
		{
			name:     "Complex: aab vs c*a*b",
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
		{
			name:     "Complex: aa vs a",
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			name:     "Complex: aa vs a*",
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			name:     "Complex: ab vs .*",
			s:        "ab",
			p:        ".*",
			expected: true,
		},

		// Edge cases
		{
			name:     "Pattern longer than string",
			s:        "a",
			p:        "ab*",
			expected: true,
		},
		{
			name:     "Multiple stars",
			s:        "aaa",
			p:        "a*a*a*",
			expected: true,
		},
		{
			name:     "Star at beginning",
			s:        "ab",
			p:        "*ab",
			expected: false, // Invalid pattern, * must have preceding element
		},
		{
			name:     "Consecutive stars (invalid)",
			s:        "a",
			p:        "a**",
			expected: false, // Invalid pattern
		},

		// Failure cases
		{
			name:     "Pattern misMatchV1",
			s:        "abc",
			p:        "def",
			expected: false,
		},
		{
			name:     "String longer no match",
			s:        "abcd",
			p:        "abc",
			expected: false,
		},
		{
			name:     "a*b doesn't match aac",
			s:        "aac",
			p:        "a*b",
			expected: false,
		},

		// LeetCode examples
		{
			name:     "LeetCode example 1",
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			name:     "LeetCode example 2",
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			name:     "LeetCode example 3",
			s:        "ab",
			p:        ".*",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isMatchV1(test.s, test.p)
			if result != test.expected {
				t.Errorf("isMatchV1(%q, %q) = %v; want %v", test.s, test.p, result, test.expected)
			}
		})
	}
}

// Test main function
func TestMain_Coverage(t *testing.T) {
	// This test just calls main to ensure 100% coverage
	// In real projects, you typically wouldn't test main
	main()
}
