package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, error := ioutil.ReadFile("day-2-input-sample.txt")
	check(error)
	fmt.Print(data)
}
