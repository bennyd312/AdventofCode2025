//go:build a10

package a10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func loadInput(path string) (data []problem) {

	file, err := os.Open("inputs/" + path)

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
		numbers = append(numbers, num)
	}

	return numbers
}

func getButtons(texts [][]string) [][]int {
	var buttons [][]int
	for i := range texts {
		if len(texts[i]) < 2 {
			continue // skip if no capture group
		}
		button := bracketToDigits(texts[i][1]) // use capture group
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
	//represents the problem - doesnt change after reading input
	goalState           []bool
	buttons             [][]int
	joltageRequirements []int
}

type configuration struct {
	//represents the current configuration for solving the problem
	chosenButtons []bool
}

type node struct {
	//queue for different configurations
	info configuration
	next *node
}

func evaluateConfig(prob problem, config configuration) bool {
	var state []bool
	for i := 0; i < len(prob.goalState); i++ {
		state = append(state, false)
	}
	for i := range config.chosenButtons {
		if config.chosenButtons[i] {
			button := prob.buttons[i]
			for _, v := range button {
				state[v] = !state[v]
			}
		}
	}

	for i := 0; i < len(state); i++ {
		if state[i] != prob.goalState[i] {
			return false
		}
	}
	return true

}

func getStartIdx(config configuration) int {
	for i := 0; i < len(config.chosenButtons); i++ {
		if config.chosenButtons[i] == true {
			return i
		}
	}
	return -1
}

func getBestConfig(prob problem) int {
	solution := len(prob.buttons)
	var initialButtons []bool
	for i := 0; i < len(prob.buttons); i++ {
		initialButtons = append(initialButtons, false)
	}

	head := node{configuration{initialButtons}, nil}
	ptrHead := &head
	ptrTail := &head

	for ptrHead != nil {
		pushes := 0
		for i := 0; i < len(initialButtons); i++ {
			if ptrHead.info.chosenButtons[i] == true {
				pushes++
			}
		}
		if solution <= pushes {
			ptrHead = ptrHead.next
			continue
		}
		if evaluateConfig(prob, ptrHead.info) {
			solution = pushes
		} else {
			start := getStartIdx(ptrHead.info)
			for i := 0; i < len(initialButtons); i++ {
				if start < i && ptrHead.info.chosenButtons[i] == false {
					newConfig := make([]bool, len(ptrHead.info.chosenButtons))
					copy(newConfig, ptrHead.info.chosenButtons)
					newConfig[i] = true
					newNode := node{configuration{newConfig}, nil}
					ptrTail.next = &newNode
					ptrTail = ptrTail.next
				}
			}
			ptrHead = ptrHead.next
		}
	}
	return solution
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
