package AStar

func IDAStar[T Number](s, t int, maxThreshold T, adj [][]Edge[T], h func(v, t int) T) T {
	sNode := &Node[T]{s, 0, h(s, t), nil}
	lastNode, threshold := search(sNode, t, h(s, t), adj, h)
	for lastNode == nil && threshold <= maxThreshold {
		lastNode, threshold = search(sNode, t, threshold, adj, h)
	}
	if lastNode != nil && lastNode.v == t {
		return lastNode.cost
	}
	return MaxValue(T(0))
}

func search[T Number](node *Node[T], t int, threshold T, adj [][]Edge[T], h func(v, t int) T) (*Node[T], T) {
	f := node.cost + node.heuristic
	if f > threshold {
		return nil, f
	}
	if node.v == t {
		return node, f
	}

	minCost := MaxValue(T(0))
	for _, succ := range adj[node.v] {
		succNode := &Node[T]{succ.to, node.cost + succ.c, h(succ.to, t), node}
		path, newThreshold := search(succNode, t, threshold, adj, h)
		if path != nil {
			return path, newThreshold
		}
		minCost = min(minCost, newThreshold)
	}
	return nil, minCost
}
