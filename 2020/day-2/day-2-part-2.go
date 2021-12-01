package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, error := os.Open("day-2-input.txt")
	check(error)

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ": ")
		rule, password := splitLine[0], splitLine[1]
		isValid := checkValidity(rule, password)
		print(rule, password, isValid, "\n")

		if isValid {
			count++
		}
	}

	print(count)
}

func checkValidity(rule string, password string) bool {
	mixMax := strings.Split(rule, " ")[0]
	letter := strings.Split(rule, " ")[1]
	minString := strings.Split(mixMax, "-")[0]
	maxString := strings.Split(mixMax, "-")[1]

	min, err1 := strconv.Atoi(minString)
	check(err1)
	max, err2 := strconv.Atoi(maxString)
	check(err2)
	return (checkLetterAtPosition(min, letter, password) != checkLetterAtPosition(max, letter, password))
}

func checkLetterAtPosition(
	position int,
	letter string,
	password string) bool {
	charAtPosition := password[position-1]
	return string(charAtPosition) == letter
}
