package utils

import "math/rand"

type Graph struct {
	nodes int
	edges map[int][]int
}

func MaxDegree(graph *Graph) int {
	maxDegree := 0
	for _, edges := range graph.edges {
		if len(edges) > maxDegree {
			maxDegree = len(edges)
		}
	}
	return maxDegree
}

func GenerateColours(numNodes int, numColours int, rng *rand.Rand) []int {
	colours := make([]int, numNodes)
	for i := range colours {
		colours[i] = rng.Intn(numColours)
	}
	return colours
}

func GenerateGraph(n int, p float64, rng *rand.Rand) *Graph {
	graph := &Graph{
		nodes: n,
		edges: make(map[int][]int),
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if rng.Float64() < p {
				graph.edges[i] = append(graph.edges[i], j)
				graph.edges[j] = append(graph.edges[j], i)
			}
		}
	}

	return graph
}
