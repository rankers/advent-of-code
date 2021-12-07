package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lineSegmentDataArray, _ := ioutil.ReadFile("day-5-input.txt")
	grid := Grid{}
	for _, lineSegmentData := range strings.Split(string(lineSegmentDataArray), "\n") {
		start, end := parseLineSegmentData(lineSegmentData)
		segment := LineSegment{
			start: start,
			end:   end,
		}
		grid.drawLine(segment)
	}
	grid.print()
}

type Grid struct {
	gridPositions [1000][1000]int
}

func (grid *Grid) drawLine(lineSegment LineSegment) {
	for _, coord := range lineSegment.allCoords() {
		grid.gridPositions[coord.y][coord.x] += 1
	}
}

func (grid *Grid) overlappingLines() int {
	var count int
	for _, line := range grid.gridPositions {
		for _, gridPosition := range line {
			if gridPosition > 1 {
				count++
			}
		}
	}
	return count
}

func (grid *Grid) print() {
	fmt.Printf("Count overlapping lines: %d \n", grid.overlappingLines())
}

type LineSegment struct {
	start Coordinate
	end   Coordinate
}

type Coordinate struct {
	x int
	y int
}

func (lineSegment *LineSegment) allCoords() []Coordinate {
	coordSlice := []Coordinate{}
	fmt.Printf("Working on line seg %v \n", lineSegment)

	if lineSegment.is45Diagonal() {
		return lineSegment.diagonalCoords()
	}
	// Only allowed if horizontal or vertical
	if lineSegment.start.x == lineSegment.end.x || lineSegment.start.y == lineSegment.end.y {

		coordSlice = append(coordSlice, lineSegment.start)

		if lineSegment.start.x == lineSegment.end.x { // vertical
			// Coords are ordered correctly so start is less than end
			// fmt.Printf("%v \n", lineSegment.start.y < lineSegment.start.x)
			if lineSegment.start.y < lineSegment.end.y {
				for i := lineSegment.start.y + 1; i < lineSegment.end.y; i++ {
					coordSlice = append(coordSlice, Coordinate{lineSegment.start.x, i})
				}
			} else {

				for i := lineSegment.start.y - 1; i > lineSegment.end.y; i-- {
					coordSlice = append(coordSlice, Coordinate{lineSegment.start.x, i})
				}
			}
		}
		// Coords are ordered correctly so start is less than end
		if lineSegment.start.y == lineSegment.end.y { // horizontal
			if lineSegment.start.x < lineSegment.end.x {
				for i := lineSegment.start.x + 1; i < lineSegment.end.x; i++ {
					coordSlice = append(coordSlice, Coordinate{i, lineSegment.start.y})
				}
			} else {
				for i := lineSegment.start.x - 1; i > lineSegment.end.x; i-- {
					coordSlice = append(coordSlice, Coordinate{i, lineSegment.start.y})
				}
			}
		}

		coordSlice = append(coordSlice, lineSegment.end)
	}
	return coordSlice
}

func (lineSegment *LineSegment) diagonalCoords() []Coordinate {
	coordSlice := []Coordinate{}
	switch {
	case lineSegment.start.x < lineSegment.end.x && lineSegment.start.y < lineSegment.end.y:
		fmt.Println("Dropping right")
		for i := 0; i <= lineSegment.end.x-lineSegment.start.x; i++ {
			coordSlice = append(coordSlice, Coordinate{lineSegment.start.x + i, lineSegment.start.y + i})
		}
	case lineSegment.start.x > lineSegment.end.x && lineSegment.start.y < lineSegment.end.y:
		fmt.Println("Dropping left")
		for i := 0; i <= lineSegment.start.x-lineSegment.end.x; i++ {
			coordSlice = append(coordSlice, Coordinate{lineSegment.start.x - i, lineSegment.start.y + i})
		}
	case lineSegment.start.x < lineSegment.end.x && lineSegment.start.y > lineSegment.end.y:
		fmt.Println("Rising right")
		for i := 0; i <= lineSegment.end.x-lineSegment.start.x; i++ {
			coordSlice = append(coordSlice, Coordinate{lineSegment.start.x + i, lineSegment.start.y - i})
		}
	case lineSegment.start.x > lineSegment.end.x && lineSegment.start.y > lineSegment.end.y:
		fmt.Println("Rising left")
		for i := 0; i <= lineSegment.start.x-lineSegment.end.x; i++ {
			coordSlice = append(coordSlice, Coordinate{lineSegment.start.x - i, lineSegment.start.y - i})
		}
	default:
		fmt.Println("Miss")
	}
	return coordSlice
}

func (lineSegment *LineSegment) is45Diagonal() bool {
	// minus x and minus y off each other if they're equal we're good
	xdiff := lineSegment.start.x - lineSegment.end.x
	if xdiff < 0 {
		xdiff = -xdiff
	}
	ydiff := lineSegment.start.y - lineSegment.end.y
	if ydiff < 0 {
		ydiff = -ydiff
	}
	return xdiff == ydiff
}

func parseLineSegmentData(lineSegmentData string) (Coordinate, Coordinate) {
	startCoordData, endCoordData := strings.Split(lineSegmentData, "->")[0], strings.Split(lineSegmentData, "->")[1]
	startX, _ := strconv.Atoi(strings.TrimSpace(strings.Split(startCoordData, ",")[0]))
	startY, _ := strconv.Atoi(strings.TrimSpace(strings.Split(startCoordData, ",")[1]))
	endX, _ := strconv.Atoi(strings.TrimSpace(strings.Split(endCoordData, ",")[0]))
	endY, _ := strconv.Atoi(strings.TrimSpace(strings.Split(endCoordData, ",")[1]))

	startCoord := Coordinate{
		x: startX,
		y: startY,
	}

	endCoord := Coordinate{
		x: endX,
		y: endY,
	}

	return startCoord, endCoord
}
