//go:build template

package a7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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

func splitBeam(beams []bool, splitters string) (newBeams []bool, splitCounter int) {
	size := len(beams)
	for i := 0; i < size; i++ {
		newBeams = append(newBeams, false)
	}
	for i, beam := range beams {
		if splitters[i] == '^' && beam {
			splitCounter++
			if i == 0 {
				newBeams[1] = true
			} else if i == size-1 {
				newBeams[i-1] = true
			} else {
				newBeams[i-1] = true
				newBeams[i+1] = true
			}
		} else if beam {
			newBeams[i] = true
		}
	}

	return newBeams, splitCounter
}
func initializeBeams(beams string) (newBeams []bool) {
	for _, state := range beams {
		if state == 'S' {
			newBeams = append(newBeams, true)
		} else {
			newBeams = append(newBeams, false)
		}
	}
	return newBeams
}

func solve(path string) {
	var solution int
	data := loadInput(path)
	var beams []bool

	for i, _ := range data {
		if i == 0 {
			beams = initializeBeams(data[i])
		} else {
			newBeams, splitCounter := splitBeam(beams, data[i])
			solution += splitCounter
			beams = newBeams
		}
	}
	fmt.Printf("For path %s the solution is %d.\n", path, solution)
}

func main() {
	var paths = []string{"7a_simple.txt", "7a_input.txt"}
	p := fmt.Println
	for _, path := range paths {
		start := time.Now()
		solve(path)
		dt := time.Since(start)
		p(dt)

	}

}
