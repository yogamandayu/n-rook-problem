package rook

import (
	"fmt"
)

type NRook struct {
	MaxRook     int
	Combination []Board
	Board       *Board
	UniqueMap   map[string]int
}

//Board is the main struct of rook
type Board struct {
	Cell map[Coordinate]CellState
	Properties
}

type Properties struct {
	Size Size
}

//Size is the rook size struct
type Size struct {
	X int
	Y int
}

type Coordinate [2]int

type CellState string

const (
	CellStateEmpty CellState = " "
	CellStateWall  CellState = "#"
	CellStateRook  CellState = "*"
)

//NewNRook is initializing the n rook problem solution
func NewNRook() *NRook {
	return &NRook{
		MaxRook:     0,
		Combination: nil,
		Board:       &Board{},
		UniqueMap:   make(map[string]int),
	}
}

//SetUsingTemplate is for set the rook with custom or defined rook template
func (b *Board) SetUsingTemplate(board [][]CellState) {
	b.setSize(len(board), len(board[0]))
	b.Cell = make(map[Coordinate]CellState, len(board)*len(board[0]))
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			switch board[y][x] {
			case CellStateEmpty:
				b.SetEmpty(x, y)
			case CellStateRook:
				b.SetRook(x, y)
			case CellStateWall:
				b.SetWall(x, y)
			}
		}
	}
}

//PrintBoard is for printing the rook
func (b *Board) PrintBoard() {
	for y := 0; y < b.Size.Y; y++ {
		for x := 0; x < b.Size.X; x++ {
			fmt.Print(b.State(x, y))
		}
		fmt.Println()
	}
	fmt.Println("========")
}

//ToString is for turn board cell to string
func (b *Board) ToString() (s string) {
	for y := 0; y < b.Size.Y; y++ {
		for x := 0; x < b.Size.X; x++ {
			s += string(b.State(x, y))
		}
	}
	return
}

//SetSize is set the rook's size
func (b *Board) setSize(y, x int) {
	b.Size.Y = y
	b.Size.X = x
}

//SetRook is for filling the rook
func (b *Board) SetRook(x, y int) {
	coordinate := [2]int{x, y}
	b.Cell[coordinate] = CellStateRook
}

//IsRook is for checking if the coordinate is filled with rook
func (b *Board) IsRook(x, y int) bool {
	coordinate := [2]int{x, y}
	return b.Cell[coordinate] == CellStateRook
}

//SetEmpty is for placing the empty cell to the rook
func (b *Board) SetEmpty(x, y int) {
	coordinate := [2]int{x, y}
	b.Cell[coordinate] = CellStateEmpty
}

//IsEmpty is for checking if the coordinate is empty
func (b *Board) IsEmpty(x, y int) bool {
	coordinate := [2]int{x, y}
	return b.Cell[coordinate] == CellStateEmpty
}

//SetWall is for placing the wall to the rook
func (b *Board) SetWall(x, y int) {
	coordinate := [2]int{x, y}
	b.Cell[coordinate] = CellStateWall
}

//IsWall is for checking if the coordinate is walled
func (b *Board) IsWall(x, y int) bool {
	coordinate := [2]int{x, y}
	return b.Cell[coordinate] == CellStateWall
}

//State is used for getting the cell state
func (b *Board) State(x, y int) CellState {
	coordinate := [2]int{x, y}
	return b.Cell[coordinate]
}

func (b *Board) IsAboveOK(x, y int) bool {
	for i := y - 1; i >= 0; i-- {
		switch b.State(x, i) {
		case CellStateEmpty:
			continue
		case CellStateWall:
			return true
		case CellStateRook:
			return false
		}
	}
	return true
}

func (b *Board) IsLeftOK(x, y int) bool {
	for i := x - 1; i >= 0; i-- {
		switch b.State(i, y) {
		case CellStateEmpty:
			continue
		case CellStateWall:
			return true
		case CellStateRook:
			return false
		}
	}
	return true
}

//Calculate is for searching the combination and total filled
func (r *NRook) Calculate(startX, startY int) {
	for y := startY; y < r.Board.Size.Y; y++ {
		for x := startX; x < r.Board.Size.X; x++ {
			if !r.Board.IsEmpty(x, y) || !r.Board.IsLeftOK(x, y) || !r.Board.IsAboveOK(x, y) {
				continue
			}
			r.Board.SetRook(x, y)
			r.Calculate(x, y)
			r.Board.SetEmpty(x, y)
		}
		startX = 0
	}

	n := r.Board.CountFilled()
	if n > r.MaxRook {
		r.MaxRook = n
		r.Combination = nil
	}
	if n == r.MaxRook {
		if _, ok := r.UniqueMap[r.Board.ToString()]; !ok {
			r.UniqueMap[r.Board.ToString()]++
			r.Combination = append(r.Combination, *r.Board.DuplicateBoard())
		}
	}
}

//CountFilled is for counting Filled cell in board
func (b *Board) CountFilled() (n int) {
	for y := 0; y < b.Size.Y; y++ {
		for x := 0; x < b.Size.X; x++ {
			if b.IsRook(x, y) {
				n++
			}
		}
	}
	return
}

//DuplicateBoard is for duplicating board
func (b *Board) DuplicateBoard() *Board {
	var newBoard Board
	newBoard.setSize(b.Size.Y, b.Size.X)
	newBoard.Cell = make(map[Coordinate]CellState, b.Size.Y*b.Size.X)
	for y := 0; y < b.Size.Y; y++ {
		for x := 0; x < b.Size.X; x++ {
			switch b.State(x, y) {
			case CellStateEmpty:
				newBoard.SetEmpty(x, y)
			case CellStateRook:
				newBoard.SetRook(x, y)
			case CellStateWall:
				newBoard.SetWall(x, y)
			}
		}
	}
	return &newBoard
}
