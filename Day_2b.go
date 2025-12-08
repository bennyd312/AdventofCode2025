//go:build b2

package b2

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
	file, err := os.Open(path)
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

func validateString(x string) bool {
	size := len(x)
	for i := 1; i < size; i++ {
		if m := make(map[string]int); size%i == 0 {
			parts := size / i
			for j := 0; j < parts; j++ {
				m[x[j*i:(j+1)*i]] += 1
			}
			if m[x[0:i]] == parts {
				return true
			}
		}
	}
	return false
}
func invalidIdFinder(x int, y int) int {
	var total int
	for i := x; i <= y; i++ {
		str := strconv.Itoa(i)
		if validateString(str) {
			total += i
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
