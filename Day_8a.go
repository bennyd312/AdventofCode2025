//go:build template

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EdgeHeap []Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].length < h[j].length }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

type Edge struct {
	target Node
	source Node
	length float64
}

type Node struct {
	x int
	y int
	z int
}

func abs(number int) float64 {
	if number < 0 {
		return float64(-1 * number)
	}
	return float64(number)
}

func distance(source, target Node) float64 {
	dx := math.Pow(abs(source.x-target.x), 2)
	dy := math.Pow(abs(source.y-target.y), 2)
	dz := math.Pow(abs(source.z-target.z), 2)

	return math.Sqrt(dx + dy + dz)
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

func strListtoIntList(numbers []string) (nums []int) {
	for i := range numbers {
		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
func notCircuit(edge Edge, components map[Node]int) bool {
	if components[edge.source] == components[edge.target] && components[edge.source] != 0 {
		return false
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func updateComponents(edge Edge, id int, components map[Node]int) {
	source_id := components[edge.source]
	target_id := components[edge.target]
	if source_id == 0 && target_id != 0 {
		components[edge.source] = target_id
		return
	}
	if target_id == 0 && source_id != 0 {
		components[edge.target] = source_id
		return
	}
	if target_id == source_id { //both are zero == one element components
		components[edge.target] = id
		components[edge.source] = id
		return
	}
	lowest_id := min(source_id, target_id) // else both their id's are nonzero, we take the minimum and assign it to all associated nodes in the component == merge components
	for node := range components {
		if components[node] == target_id || components[node] == source_id {
			components[node] = lowest_id
		}
	}
	return
}

func countComponentSizes(components map[Node]int) int {
	var compSizes []int
	compCounter := make(map[int][]Node)
	for node, id := range components {
		if id != 0 {
			compCounter[id] = append(compCounter[id], node)
		}
	}
	for _, value := range compCounter {
		compSizes = append(compSizes, len(value))
	}
	sort.Slice(compSizes, func(i, j int) bool {
		return compSizes[i] > compSizes[j]
	})
	return compSizes[0] * compSizes[1] * compSizes[2]
}

func solve(path string, maxEdge int) {
	data := loadInput(path)
	var nodes []Node
	var edges []Edge
	for _, value := range data {
		coords := strListtoIntList(strings.Split(value, ","))
		node := Node{coords[0], coords[1], coords[2]}
		nodes = append(nodes, node)
	}

	n := len(nodes)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			source := nodes[i]
			target := nodes[j]
			dist := distance(source, target)
			edges = append(edges, Edge{source, target, dist})
		}
	}

	h := &EdgeHeap{}
	heap.Init(h)

	for _, edge := range edges {
		heap.Push(h, edge)
	}
	components := make(map[Node]int)
	currentEdges := 0
	id := 1
	for h.Len() > 0 && currentEdges < maxEdge {
		edge := heap.Pop(h).(Edge)
		//fmt.Println(edge)
		if notCircuit(edge, components) {
			updateComponents(edge, id, components)
			id++
		}
		currentEdges++

	}
	fmt.Printf("For path %s the solution is %d\n", path, countComponentSizes(components))
}

func main() {
	var paths = []string{"8a_simple.txt", "8a_input.txt"}
	var maxEdges = []int{10, 1000}

	for i := range paths {
		solve(paths[i], maxEdges[i])
	}

}
