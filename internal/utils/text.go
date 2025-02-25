package utils

func Pluralize(singular, plural string, value int) string {
	if value == 1 {
		return singular
	}
	return plural
}
