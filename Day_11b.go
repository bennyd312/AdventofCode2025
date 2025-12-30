package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func parseLine(line string) (string, []string) {
	line = strings.Replace(line, ":", "", 1)
	allNodes := strings.Split(line, " ")
	return allNodes[0], allNodes[1:]
}

func mul(nums ...int) int {
	solution := 1
	for i := range nums {
		solution = solution * nums[i]
	}
	return solution
}

func solve(path string) {
	data := loadInput(path)
	edges := make(map[string][]string)
	for i := range data {
		startNode, endNodes := parseLine(data[i])
		edges[startNode] = endNodes
	}

	var solvePart func(string, string) int
	solvePart = func(start string, end string) (solution int) {
		memo := make(map[string]int)

		var dfs func(string) int
		dfs = func(vertex string) int {
			if vertex == end {
				return 1
			}
			if val, ok := memo[vertex]; ok {
				return val
			}

			total := 0

			for _, v := range edges[vertex] {
				total += dfs(v)
			}
			memo[vertex] = total
			return total
		}

		return dfs(start)
	}

	solution1 := mul(
		solvePart("svr", "dac"),
		solvePart("dac", "fft"),
		solvePart("fft", "out"),
	)
	solution2 := mul(
		solvePart("svr", "fft"),
		solvePart("fft", "dac"),
		solvePart("dac", "out"),
	)
	solution := solution1 + solution2
	fmt.Printf("For path %s the solution is %v.\n", path, solution)

}

func main() {
	var paths = []string{"11b_simple.txt", "11a_input.txt"}

	for _, path := range paths {
		solve(path)
	}

}
