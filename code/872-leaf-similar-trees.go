type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, seq []int) []int {
	if root == nil {
		return seq
	}
	if root.Left== nil && root.Right == nil {
		seq = append(seq, root.Val)
	} else {
		if root.Left != nil {
			seq = dfs(root.Left, seq)
		}
		if root.Right != nil {
			seq = dfs(root.Right, seq)
		}
	}
	return seq
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	seq1 := dfs(root1, []int{})
	seq2 := dfs(root2, []int{})
	fmt.Println(seq1)
	fmt.Println(seq2)
	if len(seq1) != len(seq2) {
		return false
	}
	for i:=0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			return false
		}
	}
	return true
}
