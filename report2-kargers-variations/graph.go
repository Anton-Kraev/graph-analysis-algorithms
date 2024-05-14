package main

import (
	"math"
	"math/rand/v2"
)

type Graph [][]int

func (graph Graph) minCut(recursive bool) int {
	if len(graph) <= 6 || !recursive {
		return graph.contract(2).cutSize()
	}
	t := int(math.Ceil(1 + float64(len(graph))/math.Sqrt2))
	return min(
		graph.contract(t).minCut(recursive),
		graph.contract(t).minCut(recursive),
	)
}

func (graph Graph) cutSize() int {
	return graph[0][1]
}

func (graph Graph) contract(t int) Graph {
	modified := graph.copy()
	for len(modified) > t {
		i := rand.IntN(len(modified) - 1)
		j := rand.IntN(max(1, len(modified)-2-i)) + 1 + i
		for k := 0; k < len(modified); k++ {
			if i != k && j != k {
				modified[i][k] += modified[j][k]
				modified[j][k] = modified[len(modified)-1][k]
				modified[k][i] += modified[k][j]
				modified[k][j] = modified[k][len(modified)-1]
			}
		}
		for k := 0; k < len(modified); k++ {
			modified[k] = modified[k][:len(modified)-1]
		}
		modified = modified[:len(modified)-1]
	}
	return modified
}

func (graph Graph) copy() Graph {
	copied := make(Graph, len(graph))
	for i := 0; i < len(graph); i++ {
		copied[i] = append(copied[i], graph[i]...)
	}
	return copied
}
