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
		// fmt.Printf("Data: %v, Start: %v, End: %v, All Coords: %v \n", lineSegmentData, segment.start, segment.end, segment.allCoords())
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
	// for _, line := range grid.gridPositions {
	// 	// fmt.Printf("%v\n", line)
	// }
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
