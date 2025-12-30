//go:build a12

package a12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func parseLinePresent(line string) (area int) {
	for i := range line {
		if line[i] == '#' {
			area++
		}
	}
	return area
}
func parseLineGrid(line string) (gridArea int, presentCount [6]int) {
	lines := strings.Split(strings.Replace(line, ":", "", 1), " ")

	gridSize := strings.Split(lines[0], "x")
	m, m_err := strconv.Atoi(gridSize[0])
	n, n_err := strconv.Atoi(gridSize[1])

	if m_err != nil {
		log.Fatal(m_err)
	}
	if n_err != nil {
		log.Fatal(n_err)
	}
	gridArea = m * n

	for i := 1; i < len(lines); i++ {
		x, err := strconv.Atoi(lines[i])
		if err != nil {
			log.Fatal(err)
		}
		presentCount[i-1] = x
	}
	return gridArea, presentCount
}

func solve(path string) {
	solution := 0
	data := loadInput(path)
	grids := false
	var presentsArea []int
	currArea := 0

	var solveGrid func(int, []int) int
	solveGrid = func(gridArea int, presentCount []int) int {
		usedArea := 0
		for i := 0; i < len(presentCount); i++ {
			usedArea += presentCount[i] * presentsArea[i]
		}

		if gridArea < usedArea {
			return 0
		} else {
			return 1
		}
	}

	for i := range data {
		line := data[i]
		if grids {
			gridArea, presentCount := parseLineGrid(line)
			solution += solveGrid(gridArea, presentCount[:])
		} else {
			if line == "" {
				presentsArea = append(presentsArea, currArea)
				currArea = 0
			} else if line[0] == '#' || line[0] == '.' {
				currArea += parseLinePresent(line)
			} else if line[len(line)-1] == ':' {
				continue
			} else {
				grids = true
				gridArea, presentCount := parseLineGrid(line)
				solution += solveGrid(gridArea, presentCount[:])
			}
		}
	}

	fmt.Printf("For path %s the solution is %v", path, solution)
}

func main() {
	var paths = []string{"12a_input.txt"}

	for _, path := range paths {
		solve(path)
	}

}
