//go:build a4

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadInput(path string) (data1 []string, data2 []string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	secondLine := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			secondLine = true
		} else if !secondLine {
			data1 = append(data1, line)
		} else {
			data2 = append(data2, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return

}

func solve(path string) {
	solution := 0
	idRanges, ids := loadInput(path)
	var intervals [][]int

	for _, v := range idRanges {
		nums := strings.Split(v, "-")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		var interval = []int{x, y}
		intervals = append(intervals, interval)
	}

	for _, v := range ids {
		num, _ := strconv.Atoi(v)
		found := false
		for _, interval := range intervals {
			if interval[0] <= num && num <= interval[1] {
				solution++
				found = true
			}
			if found {
				break
			}
		}
	}
	fmt.Printf("For path %s the solution is %d.\n", path, solution)
}

func main() {
	var paths = []string{"5a_simple.txt", "5a_input.txt"}

	for _, path := range paths {
		solve(path)
	}
}
