//go:build a4

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func loadInput(path string) []string {
	var data []string
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data

}

func getSurroundingCoordinates(x, y, m, n int) [][]int {
	var potentialCoordinates = [][]int{{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}, {x + 1, y}, {x + 1, y + 1}, {x, y + 1}, {x - 1, y + 1}, {x - 1, y}}
	var coordinates [][]int

	for _, v := range potentialCoordinates {
		x_coord, y_coord := v[0], v[1]
		if x_coord < 0 || x_coord > n-1 {
			continue
		} else if y_coord < 0 || y_coord > m-1 {
			continue
		} else {
			coordinates = append(coordinates, v)
		}
	}

	return coordinates
}

func solve(path string) {
	changes := true
	data := loadInput(path)
	m := len(data)
	n := len(data[0])
	accessibleRolls := 0
	for changes {
		changes = false
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if data[i][j] == '.' {
					continue
				}
				coords := getSurroundingCoordinates(i, j, m, n)
				counter := 0
				for _, v := range coords {
					x, y := v[0], v[1]
					if data[x][y] == '@' {
						counter++
					}
				}
				if counter < 4 {
					data[i] = data[i][:j] + "." + data[i][j+1:]
					accessibleRolls++
					changes = true
				}
			}
		}
	}
	fmt.Printf("For path %s the solution is %d.\n", path, accessibleRolls)
}

func main() {
	var paths = []string{"4a_simple.txt", "4a_input.txt"}

	for _, path := range paths {
		solve(path)
	}
}
