package main

import (
	"fmt"
	"n-rook-problem/rook"
)

func main() {
	r := rook.NewNRook()
	r.Board.SetUsingTemplate(rook.Template8x8())
	r.Calculate(0, 0)

	fmt.Println("Max Rook : ", r.MaxRook)
	fmt.Println("Total Combination : ", len(r.Combination))
}
