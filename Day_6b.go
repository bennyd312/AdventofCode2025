//go:build template

package b6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadInput(path string) [][]string {
	var data []string
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

	return processInput(data)
}

func processInput(data []string) [][]string {
	var newData [][]string
	rowLength := len(data)
	strLength := len(data[0])

	previous := 0
	for col := range data[0] {
		isEquation := true
		for i := 0; i < rowLength; i++ {
			if data[i][col] != ' ' {
				isEquation = false
				break
			}
		}
		if isEquation || col == strLength-1 {
			var equation []string
			for j := previous; j < col+1; j++ {
				num := ""
				for i := 0; i < rowLength-1; i++ {
					if data[i][j] != ' ' {
						num += string(data[i][j])
					}
				}
				if num != "" {
					equation = append(equation, num)
				}

			}
			equation = append(equation, strings.TrimSpace(data[rowLength-1][previous:col+1]))
			newData = append(newData, equation)
			previous = col + 1
		}
	}
	return newData
}

func sum(numbers []string) (result uint64) {
	for _, str := range numbers {
		num, _ := strconv.Atoi(str)
		result += uint64(num)
	}

	return result
}

func mul(numbers []string) (result uint64) {
	result = 1
	for _, str := range numbers {
		num, _ := strconv.Atoi(str)
		result = result * uint64(num)
	}

	return result
}

func solve(path string) {
	data := loadInput(path)
	var solution uint64
	m := len(data)
	for i := 0; i < m; i++ {
		n := len(data[i])
		column := data[i][:n-1]
		if operation := data[i][n-1]; operation == "+" {
			solution += sum(column)
		} else {
			solution += mul(column)

		}
	}

	fmt.Printf("For path %s the solution is %d.\n", path, solution)
}

func main() {
	var paths = []string{"6a_simple.txt", "6a_input.txt"}

	for _, path := range paths {
		solve(path)
	}

}
