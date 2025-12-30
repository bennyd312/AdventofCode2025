//go:build a2

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getRanges(path string) [][]string {
	var data = [][]string{}
	file, err := os.Open("inputs/" + path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, v := range line {
			tuple := strings.Split(v, "-")
			data = append(data, tuple)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func invalidIdFinder(x int, y int) int {
	var total int
	for i := x; i <= y; i++ {
		if str := strconv.Itoa(i); len(str)%2 != 1 {
			size := len(str)
			if strings.Compare(str[:size/2], str[size/2:]) == 0 {
				total += i
			}
		}
	}
	return total
}

func solve(path string) {
	var total int
	data := getRanges(path)

	for _, v := range data {
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])
		total += invalidIdFinder(x, y)
	}
	fmt.Printf("Solution for path %s is %d\n", path, total)
}

func main() {
	var paths = []string{"2a_simple.txt", "2a_input.txt", "2a_.txt"}

	for _, v := range paths {
		solve(v)
	}
}
