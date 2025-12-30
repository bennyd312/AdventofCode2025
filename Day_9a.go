//go:build template

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x int
	y int
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
	return Node{x, y}
}

func loadInput(path string) (data []string) {

	file, err := os.Open("inputs/" + path)

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

func solve(path string) {
	solution := 0
	data := loadInput(path)
	var nodes []Node

	for i := range data {
		nodes = append(nodes, rowToCoords(data[i]))
	}

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			currArea := area(nodes[i], nodes[j])
			if solution < currArea {
				solution = currArea
			}
		}
	}
	fmt.Printf("For path %s the solution is %d", path, solution)

}

func main() {
	var paths = []string{"9a_simple.txt", "9a_input.txt"}

	for _, path := range paths {
		solve(path)
	}

}
