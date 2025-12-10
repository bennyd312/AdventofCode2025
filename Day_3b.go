//go:build a3

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
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

func dropDigitGetMax(number string) string {
	size := len(number)
	nums := make([]string, 0, size)
	for i := 0; i < size; i++ {
		nums = append(nums, number[:i]+number[i+1:])
	}

	var idxLargest int
	valueLargest := big.NewInt(0)

	for i, n := range nums {
		curr := new(big.Int)
		curr.SetString(n, 10)
		if curr.Cmp(valueLargest) > 0 {
			valueLargest.Set(curr)
			idxLargest = i
		}
	}

	return nums[idxLargest]
}
func getJoltage(bank string) *big.Int {
	num := bank
	for len(num) > 12 {
		num = dropDigitGetMax(num)
	}

	result := new(big.Int)
	result.SetString(num, 10)
	return result
}

func solver(path string) {
	total := big.NewInt(0)
	data := getFile(path)
	for _, v := range data {
		total.Add(total, getJoltage(v))
	}

	fmt.Printf("For path %s the solution is %s.\n", path, total.String())
}

func main() {
	var paths = []string{"3a_test.txt", "3a_simple.txt", "3a_input.txt"}

	for _, path := range paths {
		solver(path)
	}
}
