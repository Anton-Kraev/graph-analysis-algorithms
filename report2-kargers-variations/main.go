package main

import "fmt"

func main() {
	graph := Graph{
		{0, 1, 1, 0, 0, 0},
		{1, 0, 1, 0, 0, 0},
		{1, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1, 0},
	}

	fmt.Println(Karger(graph, 10))
	fmt.Println(KargerStein(graph, 10))
}
