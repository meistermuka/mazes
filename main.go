package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// CellCoord bla bla bla
type CellCoord struct {
	x int
	y int
}

// Cell bla bla
type Cell struct {
	coord CellCoord
	data  string
	north *Cell
	south *Cell
	east  *Cell
	west  *Cell
}

// DIMENSION bla bla
const DIMENSION = 20

func main() {

	// Creating 2D grid of Cell
	grid := make([][]Cell, DIMENSION)
	for i := range grid {
		grid[i] = make([]Cell, DIMENSION)
	}

	// Setting the individual coordinates of each Cell
	for row := 0; row < DIMENSION; row++ {
		for col := 0; col < DIMENSION; col++ {
			grid[row][col].coord.x = row
			grid[row][col].coord.y = col
		}
	}

	fmt.Println("Initializing grid...")
	// Using Binary Tree style of maze init
	binaryTree(grid)
	printMap(grid)

}

// printMap basically prints the map on screen
func printMap(g [][]Cell) {
	var top, bottom, output bytes.Buffer
	topBorder := "---+"
	body := "   "
	corner := "+"
	output.WriteString("+")

	// Draw the top part
	for k := 0; k < DIMENSION; k++ {
		output.WriteString(topBorder)
	}

	output.WriteString("\n")

	for row := 0; row < DIMENSION; row++ {
		top.WriteString("|")
		bottom.WriteString("+")
		for col := 0; col < DIMENSION; col++ {
			cell := g[row][col]

			if cell.east != nil {
				top.WriteString(body)
				top.WriteString(" ")
			} else {
				top.WriteString(body)
				top.WriteString("|")
			}

			if cell.south != nil {
				bottom.WriteString(body)
				bottom.WriteString(corner)
			} else {
				bottom.WriteString("---")
				bottom.WriteString(corner)
			}

		}

		output.WriteString(top.String() + "\n")
		output.WriteString(bottom.String() + "\n")

		top.Reset()
		bottom.Reset()
	}

	fmt.Println(output.String())
}

// Basic grid init connecting all the cells together
func initGrid(g [][]Cell) {
	for row := 0; row < DIMENSION; row++ {
		for col := 0; col < DIMENSION; col++ {

			if row > 0 {
				g[row][col].north = &g[row-1][col]
			} else {
				g[row][col].north = nil
			}

			if row < DIMENSION-1 {
				g[row][col].south = &g[row+1][col]
			} else {
				g[row][col].south = nil
			}

			if col < DIMENSION-1 {
				g[row][col].east = &g[row][col+1]
			} else {
				g[row][col].east = nil
			}

			if col > 0 {
				g[row][col].west = &g[row][col-1]
			} else {
				g[row][col].west = nil
			}

		}
	}
}

// binary tree style of grid generation using random numbers based on neighbours
func binaryTree(g [][]Cell) {

	var neighbours []Cell
	for row := 0; row < DIMENSION; row++ {
		for col := 0; col < DIMENSION; col++ {
			if row > 0 {
				neighbours = append(neighbours, g[row-1][col])
			}

			if col < DIMENSION-1 {
				neighbours = append(neighbours, g[row][col+1])
			}

			index := generateRandomNumber(len(neighbours))
			neighbour := neighbours[index]

			nx := neighbour.coord.x
			ny := neighbour.coord.y

			if nx+1 < DIMENSION {
				g[nx][ny].south = &g[nx+1][ny]
				g[nx+1][ny].north = &g[nx][ny]
			}

			if ny-1 > -1 {
				g[nx][ny].west = &g[nx][ny-1]
				g[nx][ny-1].east = &g[nx][ny]
			}
		}
	}
}

func generateRandomNumber(amount int) int {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return r.Intn(amount)

}
