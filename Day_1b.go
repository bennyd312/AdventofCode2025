//go:build day1b

package b1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getMoves(path string) []string {
	var moves = []string{}
	file, err := os.Open("inputs/" + path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		moves = append(moves, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return moves
}

func runTests(paths []string) {
	for _, v := range paths {
		solve(v)
	}
}

func quotient(number int, divisor int) int {
	return number / divisor
}

func solve(path string) {
	var zeros int
	var current int = 50
	var moves []string = getMoves(path)

	for i := 0; i < len(moves); i++ {
		var direction string = string(moves[i][0])
		length, _ := strconv.Atoi(moves[i][1:])
		for length > 100 {
			zeros++
			length -= 100
		}
		if direction == "L" {
			if current-length < 0 {
				if current != 0 {
					zeros++
				}
				current = current - length + 100
			} else if current-length == 0 {
				zeros++
				current = 0
			} else {
				current -= length
			}
		} else {
			if current+length > 100 {
				current = current + length - 100
				zeros++
			} else if current+length == 100 {
				zeros++
				current = 0
			} else {
				current += length
			}
		}
	}
	fmt.Printf("Solution for path %d is %d\n", path, zeros)
}

func main() {
	var all_paths = []string{"1_base_1.txt", "1_base_2.txt", "1_input.txt"}
	runTests(all_paths)
}
