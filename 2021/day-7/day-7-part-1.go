package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	crabHorizontalPositionData, _ := ioutil.ReadFile("day-7-input.txt")
	crabHorizontalPositions := []int{}
	for _, crabHorizontalPositionStr := range strings.Split(string(crabHorizontalPositionData), ",") {
		pos, _ := strconv.Atoi(crabHorizontalPositionStr)
		crabHorizontalPositions = append(crabHorizontalPositions, pos)
	}

	min, max := findMinMax(crabHorizontalPositions)

	fuelDistMin := 0

	for trial := min; trial <= max; trial++ {
		fuelDist := 0
		for _, crabHorizontalPosition := range crabHorizontalPositions {
			dist := crabHorizontalPosition - trial
			if dist < 0 {
				dist = -dist
			}
			fuelDist = fuelDist + dist
		}
		if trial == min || fuelDist < fuelDistMin {
			fuelDistMin = fuelDist
		}
		fmt.Printf("Trial: %d, Fuel Used: %v, min so far %d\n", trial, fuelDist, fuelDistMin)
	}
}

func findMinMax(v []int) (int, int) {
	max := 0
	min := 0
	for i, e := range v {
		if i == 0 || e > max {
			max = e
		}
		if i == 0 || e < min {
			min = e
		}
	}
	return min, max
}
