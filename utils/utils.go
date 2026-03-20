package utils

import "math/rand"

func CountConflicts(graph *Graph, colours []int) int {
	conflicts := 0
	for node := 0; node < graph.nodes; node++ {
		for _, neighbour := range graph.edges[node] {
			if node < neighbour && colours[node] == colours[neighbour] {
				conflicts++
			}
		}
	}
	return conflicts
}

func IsConflicted(graph *Graph, colours []int, node int) bool {
	for _, neighbour := range graph.edges[node] {
		if colours[node] == colours[neighbour] {
			return true
		}
	}
	return false
}

func PickColour(graph *Graph, colours []int, node int, numColours int, rng *rand.Rand) int {
	// build set of neighbour colours
	neighbourColours := make(map[int]bool)
	for _, neighbour := range graph.edges[node] {
		neighbourColours[colours[neighbour]] = true
	}

	// collect colours not used by any neighbour
	available := []int{}
	for colour := range numColours {
		if !neighbourColours[colour] { // if the colour is not used by any neighbour
			available = append(available, colour) // add the colour to the available colours
		}
	}

	// pick from available
	if len(available) > 0 {
		return available[rng.Intn(len(available))]
	}
	// if no available colours, pick a random colour
	return rng.Intn(numColours)
}
