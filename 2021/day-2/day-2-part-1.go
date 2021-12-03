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
  instructions, err := ioutil.ReadFile("day-2-input.txt")
  check(err)
  instructionsArray := strings.Split(string(instructions), "\n")
  var horizontalPosition, depth int = 0, 0

  for _, instruction := range instructionsArray {
    direction := strings.Split(instruction, " ")[0]
    amountStr := strings.Split(instruction, " ")[1]
    amount, err1 := strconv.Atoi(amountStr)
    check(err1)

    switch direction {
    case "forward":
        horizontalPosition += amount
    case "down":
        depth += amount
    case "up":
        depth -= amount
    }
  }
  print(depth)
  print(horizontalPosition)
  distance := depth * horizontalPosition
  print("Distance: ", distance)
}