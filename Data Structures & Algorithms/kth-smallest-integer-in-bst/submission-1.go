/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var visited []*TreeNode
	
	var inorder func(*TreeNode, int)
    inorder = func(node *TreeNode, k int) {
		if node == nil {
			return
		}

		if len(visited) == k {
			return
		}

		inorder(node.Left, k)
		visited = append(visited, node)
		inorder(node.Right, k)
	}
	
	inorder(root, k)
	return visited[k-1].Val
}
