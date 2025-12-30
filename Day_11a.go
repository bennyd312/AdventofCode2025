//go:build template

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	name string
	next *Node
}

func loadInput(path string) (data []string) {

	file, err := os.Open("inputs/" + path)

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

func parseLine(line string) (string, []string) {
	line = strings.Replace(line, ":", "", 1)
	allNodes := strings.Split(line, " ")
	return allNodes[0], allNodes[1:]
}

func solve(path string) {
	solution := 0
	data := loadInput(path)
	nodes := make(map[string]Node)
	edges := make(map[string][]string)
	for i := range data {
		startNode, endNodes := parseLine(data[i])
		edges[startNode] = endNodes
		nodes[startNode] = Node{startNode, nil}
	}
	head := Node{"you", nil}
	ptrHead := &head
	ptrTail := &head

	for ptrHead != nil {
		if ptrHead.name == "out" {
			solution++
			ptrHead = ptrHead.next
		} else {
			for _, name := range edges[ptrHead.name] {
				node := Node{name, nil}
				ptrTail.next = &node
				ptrTail = ptrTail.next
			}
			ptrHead = ptrHead.next
		}
	}

	fmt.Printf("For path %s the solution is %v.\n", path, solution)

}

func main() {
	var paths = []string{"11a_simple.txt", "11a_input.txt"}

	for _, path := range paths {
		solve(path)
	}

}
