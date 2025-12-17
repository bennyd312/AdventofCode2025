//go:build template

package b7

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func splitBeam(beams []int, splitters string) (newBeams []int) {
	size := len(beams)
	for i := 0; i < size; i++ {
		newBeams = append(newBeams, 0)
	}
	for i, beam := range beams {
		if splitters[i] == '^' && beam != 0 {
			if i == 0 {
				newBeams[1] += beam
			} else if i == size-1 {
				newBeams[i-1] += beam
			} else {
				newBeams[i-1] += beam
				newBeams[i+1] += beam
			}
			newBeams[i] = 0
		} else {
			newBeams[i] += beam
		}
	}

	return newBeams
}
func initializeBeams(beams string) (newBeams []int) {
	for _, state := range beams {
		if state == 'S' {
			newBeams = append(newBeams, 1)
		} else {
			newBeams = append(newBeams, 0)
		}
	}
	return newBeams
}

func sum(arr []int) (number int) {
	for _, v := range arr {
		number += v
	}
	return number
}

func solve(path string) {
	data := loadInput(path)
	var beams []int

	for i, _ := range data {
		if i == 0 {
			beams = initializeBeams(data[i])
		} else {
			newBeams := splitBeam(beams, data[i])
			beams = newBeams
		}
	}

	fmt.Printf("For path %s the solution is %d.\n", path, sum(beams))
}

func main() {
	var paths = []string{"7a_simple.txt", "7a_input.txt"}

	for _, path := range paths {
		solve(path)
	}

}
