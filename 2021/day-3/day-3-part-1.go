package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	diagnostics, err := ioutil.ReadFile("day-3-input.txt")
	check(err)
	diagnosticsArray := strings.Split(string(diagnostics), "\n")
	var bitSumArray = make([]int, len(diagnosticsArray[0]))
	for _, diagnostic := range diagnosticsArray {
		// fmt.Printf("%s", diagnostic)
		for index, bit := range diagnostic {
			// fmt.Println(int(bit) % 2)
			bitSumArray[index] += int(bit) % 2
		}
	}
	var gammaStrArray = make([]string, len(diagnosticsArray[0]))
	var epStrArray = make([]string, len(diagnosticsArray[0]))
	for index, columnSum := range bitSumArray {
		if columnSum > len(diagnosticsArray)/2 {
			gammaStrArray[index] = "1"
			epStrArray[index] = "0"
		} else {
			gammaStrArray[index] = "0"
			epStrArray[index] = "1"
		}
	}

	gammaString := strings.Join(gammaStrArray[:], "")
	epString := strings.Join(epStrArray[:], "")
	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	ep, _ := strconv.ParseInt(epString, 2, 64)
	print(gamma * ep)
}
