package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bingoData, _ := ioutil.ReadFile("day-4-input.txt")
	bingoDataArray := strings.Split(string(bingoData), "\n")
	numberDraw, bingoDataArray := stringArrayToIntArray(strings.Split(bingoDataArray[0], ",")), bingoDataArray[1:]
	boardArray := createBoards(bingoDataArray)
	winnerIndex := -1
	for _, number := range numberDraw {
		winnerIndex = findWinner(number, boardArray)
		if winnerIndex > 0 {
			fmt.Printf("Winner Board Index: %d, Result: %d \n", winnerIndex, boardArray[winnerIndex].sumOfUnMarked()*number)
			break
		}
	}
}

// Array of boards
// Boards of 2d arrays
// All values are ints
func createBoards(bingoDataArray []string) []Board {
	var boards = []Board{}
	boardSize := 5
	batchSize := boardSize + 1 // We have 1 blank new line for every board
	batches := make([][]string, 0, (len(bingoDataArray)+batchSize-1)/batchSize)

	for batchSize < len(bingoDataArray) {
		// cutting down bingoDataArray by batch size from index 0 - batchSize
		// appending to batches from
		bingoDataArray, batches = bingoDataArray[batchSize:], append(batches, bingoDataArray[0:batchSize:batchSize])
	}
	batches = append(batches, bingoDataArray)
	for _, boardData := range batches {
		// fmt.Printf("Row: %v, length: %d \n", boardData, len(boardData))
		boards = append(boards, CreateBoard(boardData[1:])) // Assumption here is the first item is an empty line so we can drop it
	}
	return boards
}

func CreateBoard(boardData []string) Board {
	var board = Board{
		positions: make([][]Position, 5),
	}
	for i, rowStr := range boardData {
		row := stringArrayToIntArray(strings.Split(string(rowStr), " "))
		for _, value := range row {
			board.positions[i] = append(board.positions[i], Position{
				value:  value,
				marked: false,
			})
		}
	}
	// fmt.Printf("Created Board: %v \n", board)
	return board
}

type Position struct {
	value  int
	marked bool
}

type Board struct {
	positions [][]Position
}

//Value vs receiver type?
func (b *Board) winner() bool {
	return b.checkCols() || b.checkRows()
}

func (b *Board) checkRows() bool {
	anyRowAllTrue := false
	for _, row := range b.positions {
		rowAllTrue := true
		for _, position := range row {
			rowAllTrue = rowAllTrue && position.marked
		}
		if rowAllTrue {
			anyRowAllTrue = true
			break
		}
	}
	return anyRowAllTrue
}

func (b *Board) checkCols() bool {
	anyColumnAllTrue := false

	for columnIndex, _ := range b.positions {
		colAllTrue := true
		for _, row := range b.positions {
			colAllTrue = colAllTrue && row[columnIndex].marked
		}
		if colAllTrue {
			anyColumnAllTrue = true
			break
		}
	}
	return anyColumnAllTrue
}

func (b *Board) sumOfUnMarked() int {
	var sum int
	for _, row := range b.positions {
		for _, position := range row {
			if position.marked == false {
				sum += position.value
			}
		}
	}
	return sum
}

func (b *Board) mark(number int) {
	for i, row := range b.positions {
		for j, position := range row {
			if position.value == number {
				b.positions[i][j].marked = true
			}
		}
	}
}

func findWinner(number int, boardArray []Board) int {
	// fmt.Printf("Playing round with %d and boards %v \n", number, boardArray)
	for index, board := range boardArray {
		board.mark(number)
		if board.winner() {
			fmt.Printf("Board: %v wins with number %d \n", board, number)
			return index
		}
	}
	return -1
}

func stringArrayToIntArray(strArray []string) []int {
	// fmt.Printf("Converting string array: %v \n", strArray)
	var intArray = []int{}
	for _, i := range strArray {
		if i != "" {
			j, err := strconv.Atoi(strings.TrimSpace(i))
			if err != nil {
				print(i)
				panic(err)
			}
			intArray = append(intArray, j)
		}
	}
	return intArray
}
