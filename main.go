package main

import (
	"bytes"
	"fmt"
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
const DIMENSION = 10

func main() {
	//grid := [DIMENSION][DIMENSION]Cell{}
	grid := make([][]Cell, DIMENSION)
	for i := range grid {
		grid[i] = make([]Cell, DIMENSION)
	}
	//fmt.Println(grid)
	for i := 0; i < DIMENSION; i++ {
		for j := 0; j < DIMENSION; j++ {
			grid[i][j].coord.x = i
			grid[i][j].coord.y = j
		}
	}
	fmt.Println(grid)
	fmt.Println("Initializing grid...")
	initGrid(grid)
	//printGrid(grid)
	printMap(grid)

}

func printMap(g [][]Cell) {
	var top, bottom, output bytes.Buffer
	topBorder := "---+"
	body := "   "
	corner := "+"
	output.WriteString("+")
	top.WriteString("|")
	bottom.WriteString("+")

	// Draw the top part
	for k := 0; k < DIMENSION; k++ {
		output.WriteString(topBorder)
	}

	for i := 0; i < DIMENSION; i++ {
		for j := 0; j < DIMENSION; j++ {
			cell := g[i][j]

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
	}

	fmt.Println(output.String())
}

func printGrid(g [][]Cell) {
	fmt.Println("Printing Grid Info...")
	for i := 0; i < DIMENSION; i++ {
		for j := 0; j < DIMENSION; j++ {

			fmt.Println("Coords:", g[i][j].coord)

			north := g[i][j].north
			east := g[i][j].east
			south := g[i][j].south
			west := g[i][j].west

			if north != nil {
				fmt.Println("North", north.coord)
			} else {
				fmt.Println(nil)
			}

			if east != nil {
				fmt.Println("East", east.coord)
			} else {
				fmt.Println(nil)
			}

			if south != nil {
				fmt.Println("South", south.coord)
			} else {
				fmt.Println(nil)
			}

			if west != nil {
				fmt.Println("West", west.coord)
			} else {
				fmt.Println(nil)
			}

			fmt.Println()
		}
	}
}
func initGrid(g [][]Cell) {
	for i := 0; i < DIMENSION; i++ {
		for j := 0; j < DIMENSION; j++ {

			if j > 0 {
				g[i][j].north = &g[i][j-1]
			} else {
				g[i][j].north = nil
			}

			if j < DIMENSION-1 {
				g[i][j].south = &g[i][j+1]
			} else {
				g[i][j].south = nil
			}

			if i < DIMENSION-1 {
				g[i][j].east = &g[i+1][j]
			} else {
				g[i][j].east = nil
			}

			if i > 0 {
				g[i][j].west = &g[i-1][j]
			} else {
				g[i][j].west = nil
			}

		}
	}
}
