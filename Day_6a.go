//go:build template

package a6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	re := regexp.MustCompile(` +`)
	for _, line := range data {
		new_line := strings.TrimSpace(re.ReplaceAllString(line, " "))
		newData = append(newData, strings.Split(new_line, " "))
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
	length := len(data)
	for col, _ := range data[0] {
		var column []string
		for row := 0; row < length-1; row++ {
			column = append(column, data[row][col])
		}
		if operation := data[length-1][col]; operation == "+" {
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
