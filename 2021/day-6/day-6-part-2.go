package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lanternfishAgeListData, _ := ioutil.ReadFile("day-6-input.txt")

	// Initialise
	// Change structure to make a count of each of the ages of fish per go.
	var lanternFishSchool = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, lanternfishAge := range strings.Split(string(lanternfishAgeListData), ",") {
		age, _ := strconv.Atoi(lanternfishAge)
		lanternFishSchool[age]++
	}

	fmt.Printf("%v\n", lanternFishSchool)

	// Now start growing
	for i := 0; i < 256; i++ {
		lanternFishSchool = ageSchool1day(lanternFishSchool)

		fmt.Printf("After day %d: %v \n", i+1, sum(lanternFishSchool))
	}

}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func ageSchool1day(lanternFishSchool []int) []int {
	zeroCount := lanternFishSchool[0]
	lanternFishSchool = lanternFishSchool[1:]
	lanternFishSchool[6] = lanternFishSchool[6] + zeroCount
	lanternFishSchool = append(lanternFishSchool, zeroCount)

	return lanternFishSchool
}
