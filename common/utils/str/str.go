package str

func StartsWith(input string, search string) bool {
	if len(search) > len(input) {
		return false
	}
	return input[0:len(search)] == search
}
