package main

import (
	"fmt"
)

var n = 6

// [0,1],[0,2],[2,3],[2,4],[2,5]
var edges = [][]int{
	{0, 1},
	{0, 2},
	{2, 3},
	{2, 4},
	{2, 5},
}

func main() {
	fmt.Println(sumOfDistancesInTree(n, edges))
}

func sumOfDistancesInTree(n int, edges [][]int) (result []int) {
	/*
		      0 1 2 3 4 5
			0 0 1 1 0 0 0
			1 1 0 0 0 0 0
			2 1 0 0 1 1 1
			3 0 0 1 0 0 0
			4 0 0 1 0 0 0
			5 0 0 1 0 0 0
	*/
	tree := make([][]int, n+1)
	for _, e := range edges {
		i, j := e[0], e[1]

		if tree[i] == nil {
			tree[i] = make([]int, n+1)
		}

		if tree[j] == nil {
			tree[j] = make([]int, n+1)
		}

		tree[i][j] = 1
		tree[j][i] = 1
	}

	fmt.Println(countDistance(1, 5, tree))
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if i != j {
				dis, _ := countDistance(i, j, tree)
				count += dis
			}
		}
		result = append(result, count)
	}

	return result
}

func countDistance(start, end int, tree [][]int) (int, bool) {
	if start == end {
		return 0, true
	}

	for k, v := range tree[start] {
		if v == 1 {
			tree[start][k] = 0
			tree[k][start] = 0

			dis, isGetEnd := countDistance(k, end, tree)

			tree[start][k] = 1
			tree[k][start] = 1

			if isGetEnd {
				return 1 + dis, true
			}

		}
	}

	return 0, false
}
