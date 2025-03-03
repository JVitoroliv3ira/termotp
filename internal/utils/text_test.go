package utils

import "testing"

func TestPluralize(t *testing.T) {
	testCases := []struct {
		name     string
		singular string
		plural   string
		value    int
		expected string
	}{
		{"SingularForOne", "apple", "apples", 1, "apple"},
		{"PluralForMoreThanOne", "apple", "apples", 2, "apples"},
		{"PluralForZero", "apple", "apples", 0, "apples"},
		{"PluralForNegativeValues", "apple", "apples", -7, "apples"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Pluralize(tc.singular, tc.plural, tc.value)
			if got != tc.expected {
				t.Errorf("Test %s failed: Expected '%v' but got '%v'", tc.name, tc.expected, got)
			}
		})
	}
}
