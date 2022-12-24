package str

import "strings"

func StartsWith(input string, search string) bool {
	if len(search) > len(input) {
		return false
	}
	return input[0:len(search)] == search
}

func TrimWhitespace(input string) string {
	return strings.Trim(strings.Replace(input, "\t", "", -1), " ")
}
