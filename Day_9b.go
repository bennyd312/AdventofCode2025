//go:build template

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	x int
	y int
}

type Block struct {
	row_start int
	row_end   int
	intervals [][2]int
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func area(source Node, target Node) int {
	dx := abs(source.x-target.x) + 1
	dy := abs(source.y-target.y) + 1
	return dx * dy
}

func rowToCoords(row string) Node {
	line := strings.Split(row, ",")
	x, _ := strconv.Atoi(line[0])
	y, _ := strconv.Atoi(line[1])
	return Node{y, x}
}

func loadInput(path string) (data []string) {

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	return data
}
func contains(a, b int, intervals [][2]int) bool {
	for i := range intervals {
		if intervals[i][0] <= a && b <= intervals[i][1] {
			return true
		}
	}
	return false
}

func equalLayers(a, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func validRectangle(a, b Node, blocks []Block) bool {
	row_start := min(a.x, b.x)
	row_end := max(a.x, b.x)
	col_start := min(a.y, b.y)
	col_end := max(a.y, b.y)
	for i, block := range blocks {
		if row_end < block.row_start {
			return true
		}
		if block.row_end < row_start {
			continue
		}
		if block.row_start <= row_start || row_end <= block.row_end || row_start <= block.row_start {
			if contains(col_start, col_end, block.intervals) {
				if i == len(blocks)-1 {
					return true
				}
				continue
			} else {
				return false
			}
			continue
		} else {
		}
	}
	return false
}

func connectPoints(source, target Node, grid [][]bool) [][]bool {
	if source.x != target.x {
		ptr := source.x
		dx := 0
		if ptr < target.x {
			dx = 1
		} else {
			dx = -1
		}
		for ptr != target.x {
			grid[ptr][target.y] = true
			ptr = ptr + dx
		}
	} else {
		dy := 0
		ptr := source.y
		if ptr < target.y {
			dy++
		} else {
			dy--
		}
		for ptr != target.y {
			grid[target.x][ptr] = true
			ptr = ptr + dy
		}
	}
	grid[source.x][source.y] = true
	grid[target.x][target.y] = true
	return grid
}

func getGrid(nodes []Node, H, W int) [][]bool {
	grid := make([][]bool, H)
	for i := range grid {
		grid[i] = make([]bool, W)
	}

	previous := nodes[0]
	for i := 0; i < len(nodes); i++ {
		curr := nodes[i]
		grid = connectPoints(previous, curr, grid)
		previous = curr
	}
	grid = connectPoints(nodes[len(nodes)-1], nodes[0], grid)

	return grid
}

func getBlocks(nodes []Node, H, W int) []Block {
	grid := getGrid(nodes, H, W)
	var blocks []Block
	block := Block{0, 0, [][2]int{}}
	var previousLayer [][2]int
	var currLayer [][2]int
	for i := 0; i < H; i++ {
		x_start := -1
		inside := false
		edge := false
		for j := 0; j < W; j++ {
			if grid[i][j] == true {
				if edge {
					continue
				}
				if j < W-1 {
					if grid[i][j-1] == true || grid[i][j+1] == true {
						edge = true
						if x_start == -1 {
							x_start = j
						}
						continue
					} else {
						if inside {
							currLayer = append(currLayer, [2]int{x_start, j})
							x_start = -1
						} else {
							x_start = j
						}
						inside = !inside
					}
				}
			} else {
				if edge {
					edge = false
					curr := i
					passes := 0
					for curr != 0 {
						if grid[curr-1][j] == true {
							passes++
						}
						curr--
					}
					if passes%2 == 1 {
						inside = true
					} else {
						inside = false
					}

					if !inside {
						currLayer = append(currLayer, [2]int{x_start, j - 1})
						x_start = -1
						continue
					} else {
						continue
					}
				}
			}
		}

		if x_start != -1 {
			currLayer = append(currLayer, [2]int{x_start, W - 1})
		}
		if equalLayers(currLayer, previousLayer) {
			block.row_end = i
			currLayer = [][2]int{}
		} else {
			blocks = append(blocks, block)
			block = Block{i, i, currLayer}
			previousLayer = currLayer
			currLayer = [][2]int{}
		}
	}
	blocks = append(blocks, block)
	return blocks
}

func solve(path string) {
	solution := 0
	data := loadInput(path)
	var nodes []Node
	maxX := 0
	maxY := 0
	for i := range data {
		node := rowToCoords(data[i])
		nodes = append(nodes, node)

		if node.x > maxX {
			maxX = node.x
		}
		if node.y > maxY {
			maxY = node.y
		}

	}

	blocks := getBlocks(nodes, maxX+1, maxY+1)
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if validRectangle(nodes[i], nodes[j], blocks) {
				currArea := area(nodes[i], nodes[j])
				if solution < currArea {
					solution = currArea
				}
			}
		}
	}
	fmt.Printf("For path %s the solution is %d.\n", path, solution)

}

func main() {
	var paths = []string{"9a_simple.txt", "9a_simple1.txt", "9a_simple2.txt", "9a_simple3.txt", "9a_simple4.txt", "9a_input.txt"}
	for _, path := range paths {
		start := time.Now()
		solve(path)
		dt := time.Since(start)
		fmt.Println(dt)
	}

}
