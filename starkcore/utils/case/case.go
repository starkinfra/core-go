package _case

import (
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile("!^?=[A-Z0-9]")

func CamelToPascal(str string) string {
	snake := pattern.ReplaceAllString(str, "${1}_${2}")
	snake = pattern.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func PascalToCamel(str string) string {
	camel := pattern.ReplaceAllString(str, "${1}${2}")
	camel = pattern.ReplaceAllString(camel, "${1}${2}")
	return strings.ToLower(camel)
}

func CamelToKebab(str string) string {
	kebab := pattern.ReplaceAllString(str, "${1}-${2}")
	kebab = pattern.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(kebab)
}
