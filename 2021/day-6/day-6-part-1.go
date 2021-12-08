package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lanternfishAgeListData, _ := ioutil.ReadFile("day-6-sample-input.txt")

	// Initialise
	lanternFishSchool := []LanternFish{}
	for _, lanternfishAge := range strings.Split(string(lanternfishAgeListData), ",") {
		age, _ := strconv.Atoi(lanternfishAge)
		lanternFishSchool = append(lanternFishSchool, LanternFish{age: age, firstWeek: false})
	}

	// Now start growing
	for i := 0; i < 80; i++ {
		lanternFishSchool = ageSchool1day(lanternFishSchool)
		fmt.Printf("After day %d: %v \n", i+1, len(lanternFishSchool))
	}

}

func ageSchool1day(lanternFishSchool []LanternFish) []LanternFish {
	for index := range lanternFishSchool {
		agedFish, spawnNew := ageFish(lanternFishSchool[index])
		lanternFishSchool[index] = agedFish
		if spawnNew {
			lanternFishSchool = append(lanternFishSchool, LanternFish{8, false})
		}
	}

	return lanternFishSchool
}

func ageFish(fish LanternFish) (LanternFish, bool) {
	spawnFish := false
	fish.age--
	if fish.age == -1 {
		fish.age = 6
		spawnFish = true
	}
	return fish, spawnFish
}

type LanternFish struct {
	age       int
	firstWeek bool
}
