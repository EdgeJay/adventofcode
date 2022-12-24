package files

import (
	"log"
	"os"
	"strings"
)

func ReadInputsFile(path string) []string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	str := string(dat)
	arr := strings.Split(str, "\n")

	return arr
}

func ReadInputsFileRaw(path string) string {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	str := string(dat)

	return str
}
