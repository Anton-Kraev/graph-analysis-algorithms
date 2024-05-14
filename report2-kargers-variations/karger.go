package main

import "math"

func Karger(graph Graph, c int) int {
	minCut := math.MaxInt
	for i := 0; i < c*len(graph)*len(graph); i++ {
		minCut = min(graph.minCut(true), minCut)
	}
	return minCut
}

func KargerStein(graph Graph, c int) int {
	minCut := math.MaxInt
	for i := 0; i < c*int(math.Log2(float64(len(graph)))); i++ {
		minCut = min(graph.minCut(true), minCut)
	}
	return minCut
}
