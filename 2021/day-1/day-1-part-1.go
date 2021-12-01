package main

import (
  "strconv"
  "strings"
  "io/ioutil"
  // "fmt"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  // Reads entire file in 1 go - no thought for memory
  measurements, err := ioutil.ReadFile("day-1-input-sample.txt")
  check(err)
  measurementArray := strings.Split(string(measurements), "\n")
  var depthIncreaseCount int
  for currentIndex, measurementString := range measurementArray {
    if (currentIndex > 0){
      previousIndex := currentIndex - 1

      previousMeasurement, err := strconv.Atoi(measurementArray[previousIndex])
      currentMeasurement, err := strconv.Atoi(measurementString)
      check(err)
      if (previousMeasurement < currentMeasurement) {
        depthIncreaseCount = depthIncreaseCount + 1
      }
    }
  }
  print(depthIncreaseCount)
}