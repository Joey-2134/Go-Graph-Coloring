package main

import (
	"fmt"
	"go-graph-coloring/utils"
	"math/rand"
)

func main() {
	const (
		NUM_NODES   = 5000
		SEED        = 123456
		MAX_STEPS   = 1000
		NUM_COLOURS = 750
	)

	rng := rand.New(rand.NewSource(SEED))
	P_VALUES := []float64{0.1, 0.3, 0.6, 0.8, 0.95}

	for _, p := range P_VALUES {
		outCSVPath := fmt.Sprintf("conflicts_p_%.2f_k_%d.csv", p, NUM_COLOURS)
		graph := utils.GenerateGraph(NUM_NODES, p, rng)
		colours := utils.GenerateColours(NUM_NODES, NUM_COLOURS, rng)
		conflictsPerIteration := make([]int, 0, MAX_STEPS)

		for step := range MAX_STEPS {
			conflicts := utils.CountConflicts(graph, colours)
			conflictsPerIteration = append(conflictsPerIteration, conflicts)
			if conflicts == 0 {
				fmt.Println("No conflicts found after ", step, " steps")
				break
			}

			snapshot := make([]int, NUM_NODES)
			copy(snapshot, colours)

			for node := range NUM_NODES {
				if utils.IsConflicted(graph, snapshot, node) {
					colours[node] = utils.PickColour(graph, snapshot, node, NUM_COLOURS, rng)
				}
			}
		}

		if err := utils.WriteConflictsCSV(outCSVPath, conflictsPerIteration); err != nil {
			panic(err)
		}

		fmt.Println("Wrote conflicts CSV to:", outCSVPath)
	}
}
