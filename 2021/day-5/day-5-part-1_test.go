package main

import (
	"fmt"
	"testing"
)

func TestParseLineSegment(t *testing.T) {
	lineSegmentData := "0,9 -> 5,9"
	result1, result2 := parseLineSegmentData(lineSegmentData)
	expectedCoord1 := Coordinate{0, 9}
	expectedCoord2 := Coordinate{5, 9}

	if result1 != expectedCoord1 {
		t.Errorf("Coord 1 parsed wrong, %v, %v", result1, expectedCoord1)
	}
	if result2 != expectedCoord2 {
		t.Errorf("Coord 2 parsed wrong, %v, %v", result2, expectedCoord2)
	}
}

func TestVerticalLineSegment(t *testing.T) {
	lineSegment := LineSegment{
		start: Coordinate{0, 0},
		end:   Coordinate{0, 2},
	}
	lineSegmentCoords := lineSegment.allCoords()
	result := []Coordinate{Coordinate{0, 0}, Coordinate{0, 1}, Coordinate{0, 2}}

	fmt.Printf("%v", lineSegmentCoords)

	if len(lineSegmentCoords) != len(result) {
		t.Errorf("Arr length differ")
	}
	for i, v := range result {
		if v != lineSegmentCoords[i] {
			t.Errorf("%v, want: %v.", lineSegmentCoords[i], v)
		}
	}
}
func TestIs45Degrees(t *testing.T) {
	lineSegment := LineSegment{
		start: Coordinate{1, 1},
		end:   Coordinate{3, 3},
	}
	if !lineSegment.is45Diagonal() {
		t.Errorf("Expecting %v to be diagonal", lineSegment)
	}

	lineSegment2 := LineSegment{
		start: Coordinate{3, 3},
		end:   Coordinate{1, 1},
	}
	if !lineSegment2.is45Diagonal() {
		t.Errorf("Expecting %v to be diagonal", lineSegment2)
	}

	lineSegment3 := LineSegment{
		start: Coordinate{9, 7},
		end:   Coordinate{7, 9},
	}
	if !lineSegment3.is45Diagonal() {
		t.Errorf("Expecting %v to be diagonal", lineSegment3)
	}

	// lineSegment4 := LineSegment{
	// 	start: Coordinate{1, 1},
	// 	end:   Coordinate{, },
	// }
	// if !lineSegment3.is45Diagonal() {
	// 	t.Errorf("Expecting %v to be diagonal", lineSegment3)
	// }
}

func TestDiagonalLineFallingRightSegment(t *testing.T) {
	lineSegment := LineSegment{
		start: Coordinate{1, 1},
		end:   Coordinate{3, 3},
	}
	result := lineSegment.allCoords()
	expected := []Coordinate{Coordinate{1, 1}, Coordinate{2, 2}, Coordinate{3, 3}}

	if len(result) != len(expected) {
		t.Errorf("Arr length differ, expected: %v, result: %v", expected, result)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("%v, want: %v.", result[i], v)
		}
	}
}
func TestDiagonalLineFallingLeftSegment(t *testing.T) {
	lineSegment := LineSegment{
		start: Coordinate{2, 0},
		end:   Coordinate{0, 2},
	}
	result := lineSegment.allCoords()
	expected := []Coordinate{Coordinate{1, 1}, Coordinate{2, 2}, Coordinate{3, 3}}

	if len(result) != len(expected) {
		t.Errorf("Arr length differ, expected: %v, result: %v", expected, result)
	}
	for i, v := range result {
		print(v.x == expected[i].x)
		if v.x != expected[i].x {
			t.Errorf("%v, want: %v.", result[i].x, v.x)
		}
	}
}

func TestHorizontalLineSegment(t *testing.T) {
	lineSegment := LineSegment{
		start: Coordinate{0, 0},
		end:   Coordinate{2, 0},
	}
	lineSegmentCoords := lineSegment.allCoords()
	result := []Coordinate{Coordinate{0, 0}, Coordinate{1, 0}, Coordinate{2, 0}}

	if len(lineSegmentCoords) != len(result) {
		t.Errorf("Arr length differ")
	}
	for i, v := range result {
		if v != lineSegmentCoords[i] {
			t.Errorf("%v, want: %v.", lineSegmentCoords[i], v)
		}
	}
}

func TestBackwardsLineSegment(t *testing.T) {
	lineSegment := LineSegment{
		start: Coordinate{9, 4},
		end:   Coordinate{7, 4},
	}
	lineSegmentCoords := lineSegment.allCoords()
	result := []Coordinate{Coordinate{9, 4}, Coordinate{8, 4}, Coordinate{7, 4}}

	fmt.Printf("lineSegmentCoords: %v \n", lineSegmentCoords)
	fmt.Printf("result: %v \n", result)

	if len(lineSegmentCoords) != len(result) {
		t.Errorf("Arr length differ")
	}
	for i, v := range result {
		if v != lineSegmentCoords[i] {
			t.Errorf("%v, want: %v.", lineSegmentCoords[i], v)
		}
	}
}

func TestDrawLinesOnGrid(t *testing.T) {
	lineSegment1 := LineSegment{
		start: Coordinate{0, 0},
		end:   Coordinate{2, 0},
	}
	lineSegment2 := LineSegment{
		start: Coordinate{0, 0},
		end:   Coordinate{0, 2},
	}
	grid := Grid{}
	grid.drawLine(lineSegment1)
	grid.drawLine(lineSegment2)
	result := [1000][1000]int{}
	result[0][0] = 2
	result[0][1] = 1
	result[0][2] = 1
	result[1][0] = 1
	result[2][0] = 1
	grid.print()
	if len(grid.gridPositions) != len(result) {
		t.Errorf("Arr length differ")
	}
	for i, v := range result {
		if v != grid.gridPositions[i] {
			t.Errorf("%v, want: %v.", grid.gridPositions[i], v)
		}
	}
}

func TestDrawLinesOnGridOdd(t *testing.T) {
	lineSegment1 := LineSegment{
		start: Coordinate{0, 9},
		end:   Coordinate{2, 9},
	}
	grid := Grid{}
	grid.drawLine(lineSegment1)
	result := [1000][1000]int{}
	result[9][0] = 1
	result[9][1] = 1
	result[9][2] = 1
	grid.print()
	if len(grid.gridPositions) != len(result) {
		t.Errorf("Arr length differ")
	}
	for i, v := range result {
		if v != grid.gridPositions[i] {
			t.Errorf("%v, want: %v.", grid.gridPositions[i], v)
		}
	}
}
