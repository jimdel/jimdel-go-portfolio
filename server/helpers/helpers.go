package helpers

import "regexp"

const (
	EmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func IsValidRegex(str string, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(str)
}
