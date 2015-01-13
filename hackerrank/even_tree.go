// https://www.hackerrank.com/challenges/even-tree
package main

import "fmt"

type node struct {
	id       int
	count    int
	children []*node
}

func count(root *node) int {
	//fmt.Println(root.id)
	if len(root.children) == 0 {
		root.count = 1
	} else {
		sum := 1
		for _, child := range root.children {
			sum += count(child)
		}

		root.count = sum
	}

	return root.count
}

func main() {

	// N is the number of vertices and M is the number of edges.
	var n, m int
	_, err := fmt.Scanf("%d %d", &n, &m)
	if err != nil {
		panic(err)
	}

	// use matrix to save tree
	var nodes []node

	nodes = make([]node, n)

	for i := 0; i < n; i++ {
		nodes[i] = node{id: i + 1}
		nodes[i].children = make([]*node, 0)
	}

	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Scanf("%d %d", &v1, &v2)
		nodes[v2-1].children = append(nodes[v2-1].children, &nodes[v1-1])
	}

	count(&nodes[0])

	answer := 0
	for i := 1; i < n; i++ {
		if nodes[i].count%2 == 0 {
			answer += 1
		}
	}

	fmt.Println(answer)
}
