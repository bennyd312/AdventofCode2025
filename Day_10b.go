//go:build b10

package b10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

func loadInput(path string) (data []problem) {

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, lineToProblem(line))
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	return data
}

func bracketToBools(text string) []bool {
	var numbers []bool
	for i := range text {
		if text[i] == '.' {
			numbers = append(numbers, false)
		} else {
			numbers = append(numbers, true)
		}
	}

	return numbers
}

func bracketToDigits(text string) []int {
	var numbers []int
	split := strings.Split(text, ",")
	for i := range split {
		if split[i] == "" {
			continue
		}
		num, err := strconv.Atoi(split[i])
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, int(num))
	}

	return numbers
}

func getButtons(texts [][]string) [][]int {
	var buttons [][]int
	for i := range texts {
		if len(texts[i]) < 2 {
			continue
		}
		button := bracketToDigits(texts[i][1])
		buttons = append(buttons, button)
	}

	return buttons
}

func lineToProblem(text string) problem {
	reGoalState := regexp.MustCompile(`\[(.*?)]`)
	reButtons := regexp.MustCompile(`\(([^()]*)\)`)
	reJoltageReq := regexp.MustCompile(`\{(.*?)\}`)

	goalState := bracketToBools(reGoalState.FindStringSubmatch(text)[1])
	buttons := getButtons(reButtons.FindAllStringSubmatch(text, -1))
	joltageRequirements := bracketToDigits(reJoltageReq.FindStringSubmatch(text)[1])

	prob := problem{goalState, buttons, joltageRequirements}

	return prob
}

type problem struct {
	goalState           []bool
	buttons             [][]int
	joltageRequirements []int
}

func getBestConfig(prob problem) int {
	n := len(prob.joltageRequirements)
	m := len(prob.buttons)

	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, m)
	}

	// Fill matrix
	for j, btn := range prob.buttons {
		for _, idx := range btn {
			matrix[idx][j] = 1
		}
	}

	b := make([]float64, n)
	for i, val := range prob.joltageRequirements {
		b[i] = float64(val)
	}

	// Create LP
	lp := golp.NewLP(0, m)

	for i := 0; i < n; i++ {
		lp.AddConstraint(matrix[i], golp.EQ, b[i])
	}

	for j := 0; j < m; j++ {
		lp.SetInt(j, true)
	}

	obj := make([]float64, m)
	for i := range obj {
		obj[i] = 1.0
	}
	lp.SetObjFn(obj)

	status := lp.Solve()
	if status != golp.OPTIMAL {
		fmt.Println("No feasible solution for problem")
		return -1
	}

	x := lp.Variables()
	total := 0
	for _, v := range x {
		if v < 0 {
			v = 0
		}
		total += int(v + 0.5)
	}

	return total
}

func solve(path string) {
	solution := 0
	problems := loadInput(path)
	for _, problem := range problems {
		solution += getBestConfig(problem)
	}

	fmt.Printf("For path %s the solution is %v.\n", path, solution)
}

func main() {
	var paths = []string{"10a_simple.txt", "10a_input.txt"}

	for _, path := range paths {
		solve(path)
	}

}
