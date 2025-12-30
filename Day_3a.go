//go:build a3

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFile(path string) []string {
	var data = []string{}

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func getJoltage(bank string) int {
	var first int = 0
	var split int = 0
	var first_str string = "0"
	for i := 0; i < len(bank)-1; i++ {
		if current, _ := strconv.Atoi(string(bank[i])); first < current {
			first = current
			first_str = string(bank[i])
			split = i
		}
		if first_str == "9" {
			break
		}
	}
	var second int = 0
	var second_str string = string(bank[second])
	for i := split + 1; i < len(bank); i++ {
		if current, _ := strconv.Atoi(string(bank[i])); second < current {
			second = current
			second_str = string(bank[i])
		}
		if second_str == "9" {
			break
		}
	}
	output, _ := strconv.Atoi(first_str + second_str)
	return output
}

func solver(path string) {
	var maxJoltage int
	data := getFile(path)
	for _, v := range data {
		maxJoltage += getJoltage(v)
	}

	fmt.Printf("For path %s the solution is %d.\n", path, maxJoltage)
}

func main() {
	var paths = []string{"3a_simple.txt", "3a_input.txt"}

	for _, path := range paths {
		solver(path)
	}
}
