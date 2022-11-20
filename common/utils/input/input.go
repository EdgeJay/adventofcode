package input

func ProcessEachLine[V any](lines []string, cb func(line string) V) []V {
	var result []V

	for _, ln := range lines {
		result = append(result, cb(ln))
	}

	return result
}
