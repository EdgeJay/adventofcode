package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile() []string {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	str := string(dat)
	arr := strings.Split(str, "\n")

	return arr
}

func calculate(arr []string) (count int) {
	var prevNum int = 99999

	for _, s := range arr {
		n, err := strconv.Atoi(s)
		if err == nil {
			if n > prevNum {
				count += 1
			}
			prevNum = n
		}
	}

	return
}

func main() {
	arr := readFile()
	count := calculate(arr)
	fmt.Printf("%d\n", count)
}
