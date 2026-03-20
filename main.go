package main

import (
	"fmt"
	"go-graph-coloring/utils"
	"math/rand"
)

func main() {
	const (
		NUM_NODES          = 50
		PROB_OF_CONNECTION = 0.8
		SEED               = 123456
		MAX_STEPS          = 100
	)

	rng := rand.New(rand.NewSource(SEED))

	graph := utils.GenerateGraph(NUM_NODES, PROB_OF_CONNECTION, rng)
	numColours := utils.MaxDegree(graph) + 1
	colours := utils.GenerateColours(NUM_NODES, numColours, rng)

	for step := range MAX_STEPS {
		conflicts := utils.CountConflicts(graph, colours)
		if conflicts == 0 {
			fmt.Println("No conflicts found after ", step, " steps")
			break
		}

		snapshot := make([]int, NUM_NODES)
		copy(snapshot, colours)

		for node := range NUM_NODES {
			if utils.IsConflicted(graph, snapshot, node) {
				colours[node] = utils.PickColour(graph, snapshot, node, numColours, rng)
			}
		}
	}
	fmt.Println("Colours: ", colours)
}
