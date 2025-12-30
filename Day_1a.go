//go:build day1a

package a1

import (
	"bufio"
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

func main() {
	var zeros int
	var current int = 50
	var moves []string = getMoves("1_input.txt")

	for i := 0; i < len(moves); i++ {
		var direction string = string(moves[i][0])
		length, _ := strconv.Atoi(moves[i][1:])
		if direction == "L" {
			var difference int = current - length
			if difference < 0 {
				current = difference % 100
			} else {
				current = difference
			}
		} else {
			var difference int = current + length
			if difference > 99 {
				current = difference % 100
			} else {
				current = difference
			}
		}
		if current == 0 {
			zeros++
		}
	}
	println(zeros)
}
