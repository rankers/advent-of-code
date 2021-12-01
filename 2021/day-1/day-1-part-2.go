package main

import (
  "strconv"
  "strings"
  "io/ioutil"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  // Reads entire file in 1 go - no thought for memory
  measurements, err := ioutil.ReadFile("day-1-input.txt")
  check(err)
  measurementArray := strings.Split(string(measurements), "\n")
  var depthIncreaseCount int
  var windowSize = 3

  for currentIndex, _ := range measurementArray {
    var currentWindowSum = getWindowSum(measurementArray, currentIndex, windowSize)
    var nextWindowSum = getWindowSum(measurementArray, currentIndex + 1, windowSize)

    if (currentWindowSum < nextWindowSum) {
      depthIncreaseCount = depthIncreaseCount + 1
    }
  }
  print("Increase Count: ", depthIncreaseCount)
}

func getWindowSum(measurementArray []string, currentIndex int, windowSize int) int {
  var windowSum int
  var windowEndIndex = currentIndex + windowSize

  if windowEndIndex > len(measurementArray) {
    return 0
  } else {
    for i := currentIndex; i < (windowEndIndex); i++ {
      currentMeasurement, err := strconv.Atoi(measurementArray[i])
      check(err)
      windowSum += currentMeasurement
    }
    return windowSum
  }
}