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

func maxInt(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func loadInput(path string) (data1 []string, data2 []string) {
	file, err := os.Open("inputs/" + path)

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

func sortIntervals(intervals [][]int) [][]int {
	changes := true
	n := len(intervals)
	for changes {
		changes = false
		for i := 0; i < n-1; i++ {
			if intervals[i][0] > intervals[i+1][0] {
				temp := intervals[i]
				intervals[i] = intervals[i+1]
				intervals[i+1] = temp
				changes = true
			}
		}
	}
	return intervals
}

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			last[1] = maxInt(last[1], curr[1])
		} else {
			result = append(result, curr)
		}
	}
	return result
}

func intervalSize(interval []int) int {
	return interval[1] - interval[0] + 1
}

func solve(path string) {
	solution := 0
	idRanges, _ := loadInput(path)
	var intervals [][]int

	for _, v := range idRanges {
		nums := strings.Split(v, "-")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		var interval = []int{x, y}
		intervals = append(intervals, interval)
	}
	intervals = sortIntervals(intervals)
	intervals = mergeIntervals(intervals)

	for _, interval := range intervals {
		solution += intervalSize(interval)
	}

	fmt.Printf("For path %s the solution is %d.\n", path, solution)
}

func main() {
	var paths = []string{"5a_simple.txt", "5a_input.txt"}

	for _, path := range paths {
		solve(path)
	}
}
