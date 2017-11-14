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
	//fmt.Println(grid)
	fmt.Println("Initializing grid...")
	//initGrid(grid)
	//printGrid(grid)
	//printMap(grid)
	binaryTree(grid)
	printMap(grid)

}

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

	for i := 0; i < DIMENSION; i++ {
		top.WriteString("|")
		bottom.WriteString("+")
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

		top.Reset()
		bottom.Reset()
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

func binaryTree(g [][]Cell) {

	var neighbours []Cell
	for row := 0; row < DIMENSION; row++ {
		for col := 0; col < DIMENSION; col++ {

			//cell := g[row][col]

			if row > 0 {
				//g[row][col].north = &g[row-1][col]
				neighbours = append(neighbours, g[row-1][col])
			}

			if col < DIMENSION-1 {
				//g[row][col].east = &g[row][col+1]
				neighbours = append(neighbours, g[row][col+1])
			}

			index := generateRandomNumber(len(neighbours))
			neighbour := neighbours[index]

			//fmt.Printf("[%d,%d]", neighbour.coord.x, neighbour.coord.y)
			//fmt.Println()
			if *neighbour.south == cell {
				fmt.Println("Connect North")
			}

			if *neighbour.west == cell {
				fmt.Println("Connect West")
			}
			//fmt.Println(neighbour)

		}
	}
	/*for i, _ := range neighbours {
		fmt.Printf("[%d,%d]", neighbours[i].coord.x, neighbours[i].coord.y)
		fmt.Println()
	}*/
}

func generateRandomNumber(amount int) int {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return r.Intn(amount)

}
