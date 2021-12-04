package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// What I don;t understand /do better
// Primitives - rune vs int vs string vs bits/byte etc. Lots of converting between things
// How to determine more 1s or 0s, should just hold a count instead of adding up all the 1s and dividing. Riduculous
// Providing criteria (oxygen, co2) to dynamically drive 'more' or 'fewer' check

func main() {
	diagnostics, err := ioutil.ReadFile("day-3-input.txt")
	check(err)
	diagnosticsArray := strings.Split(string(diagnostics), "\n")
	oxygenRating := filterDiagnostic(0, diagnosticsArray, true)
	// co2scrubberRating := filterDiagnostic(0, diagnosticsArray, false)

	fmt.Printf("Oxyggen %v \n", oxygenRating)
	// fmt.Printf("CO2 %v \n", co2scrubberRating)
}

func filterDiagnostic(
	positionToCheck int,
	diagnosticsArray []string,
	oxygenCriteria bool) []string {

	fmt.Printf("Inbound %v \n", diagnosticsArray)

	filteredArray := []string{}

	var bitSumArray = make([]int, len(diagnosticsArray[0]))
	for _, diagnostic := range diagnosticsArray {
		for index, bit := range diagnostic {
			bitSumArray[index] += int(bit) % 2
		}
	}

	// fmt.Printf("%v \n", bitSumArray)

	for bitSumArrayIndex, columnSum := range bitSumArray {
		if bitSumArrayIndex == positionToCheck {
			fmt.Printf("Checking position %d ", positionToCheck)
			print("Col sum: ", columnSum, " ")
			print("Half array len: ", len(diagnosticsArray)/2, " ")
			print("Bigger: ", float64(columnSum) >= float64(len(diagnosticsArray))/float64(2), " ")
			print("Smaller: ", float64(columnSum) <= float64(len(diagnosticsArray))/float64(2), " ")

			if float64(columnSum) < float64(len(diagnosticsArray))/float64(2) { // CO2
				// if float64(columnSum) >= float64(len(diagnosticsArray))/float64(2) { // Oxygen
				for i := range diagnosticsArray {

					if string(diagnosticsArray[i][positionToCheck]) == "1" {

						filteredArray = append(filteredArray, diagnosticsArray[i])
					}
				}
			} else {
				for i := range diagnosticsArray {
					if string(diagnosticsArray[i][positionToCheck]) == "0" {

						filteredArray = append(filteredArray, diagnosticsArray[i])
					}
				}
			}
		}
	}

	fmt.Printf("\n Outbound  %v \n", filteredArray)

	if len(diagnosticsArray) == 1 {
		return diagnosticsArray
	} else {
		return filterDiagnostic(positionToCheck+1, filteredArray, oxygenCriteria)
	}
}

//Oxyggen [111111111001] 4089
// CO2 [011110000011] 1923
