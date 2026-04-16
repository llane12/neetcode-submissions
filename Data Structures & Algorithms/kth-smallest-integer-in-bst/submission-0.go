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
	
	var inorder func(*TreeNode)
    inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		visited = append(visited, node)
		inorder(node.Right)
	}
	
	inorder(root)
	return visited[k-1].Val
}
