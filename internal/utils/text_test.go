package utils

import "testing"

func TestPluralize_WithSingular(t *testing.T) {
	singular := "apple"
	plural := "apples"
	value := 1

	got := Pluralize(singular, plural, value)
	want := singular

	if got != want {
		t.Errorf("Pluralize() = %q; want %q", got, want)
	}
}

func TestPluralize_WithPlural(t *testing.T) {
	singular := "apple"
	plural := "apples"
	value := 2

	got := Pluralize(singular, plural, value)
	want := plural

	if got != want {
		t.Errorf("Pluralize() = %q; want %q", got, want)
	}
}

func TestPluralize_WithZero(t *testing.T) {
	singular := "apple"
	plural := "apples"
	value := 0

	got := Pluralize(singular, plural, value)
	want := plural

	if got != want {
		t.Errorf("Pluralize() = %q; want %q", got, want)
	}
}

func TestPluralize_WithNegativeValue(t *testing.T) {
	singular := "apple"
	plural := "apples"
	value := -5

	got := Pluralize(singular, plural, value)
	want := plural

	if got != want {
		t.Errorf("Pluralize() = %q; want %q", got, want)
	}
}
