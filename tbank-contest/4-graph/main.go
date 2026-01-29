package main

import (
	"bufio"
	"fmt"
	"os"
)

func findShortestCycle(n int, adj [][]int, startNode int) int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = -1
	}
	parent := make([]int, n+1)

	queue := []int{startNode}
	dist[startNode] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, to := range adj[curr] {
			if dist[to] == -1 {
				dist[to] = dist[curr] + 1
				parent[to] = curr
				queue = append(queue, to)
			} else if to != parent[curr] {
				return dist[curr] + dist[to] + 1
			}
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	minGirth := -1

	for i := 1; i <= n; i++ {
		res := findShortestCycle(n, adj, i)
		
		if res != -1 {
			if minGirth == -1 || res < minGirth {
				minGirth = res
			}
		}

		if minGirth == 3 {
			break
		}
	}

	fmt.Println(minGirth)
}