package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func dfs(node *Node, nodeMap map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	newNode, ok := nodeMap[node.Val]
	if ok {
		return newNode
	}
	newNode = &Node{Val: node.Val}
	nodeMap[node.Val] = newNode
	for _, n := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, dfs(n, nodeMap))
	}
	return newNode
}

func cloneGraph(node *Node) *Node {
	nodeMap := make(map[int]*Node)
	return dfs(node, nodeMap)
}

func main() {

}
