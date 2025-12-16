//go:build template

package none

import (
	"bufio"
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

func solve(path string) {
	//data := loadInput(path)
}

func main() {
	var paths = []string{}

	for _, path := range paths {
		solve(path)
	}

}
