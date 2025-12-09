//go:build a3

package main

import (
	"bufio"
	"os"
	"log"
	"strconv"
	"fmt"
)
func getFile(path string) []string {
	var data = []string{}

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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func getJoltage(bank string) int {
	current := string(bank[0])
	size := len(bank)

	for i:=0; i<3; i++{
		for j:=0; j<size-1; j++ {
			curr, _ := strconv.Atoi(bank[i])
			next, _ := strconv.Atoi(bank[i+1])
			if next == 0 {continue;}
			if curr == 0 {continue;}
			if curr == 1 { 
				bank[i] = "0"
				continue;
		}
			if curr < next {
				bank[i] = "0"
				continue;
			} 
		}
	}
	var output string
	for i:=0; i<size; i++ {
		if char:=string(bank[i]); char != "0" {
			output += bank[i]
		}
	}

	joltage, _ = strconv.Atoi(output)

	return joltage
}

func solver(path string) {
	var maxJoltage int
	data := getFile(path)
	for _, v := range data {
		maxJoltage += getJoltage(v)
	}

	fmt.Printf("For path %s the solution is %d.\n",path,maxJoltage)
}

func main() {
	var paths = []string{"3a_simple.txt"}
	
	for _, path := range paths {
		solver(path)
	}
}