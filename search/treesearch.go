package search

import obj "../basestruct"

func DFS(root *obj.TreeNode, visit func(node *obj.TreeNode)) {
	visit(root)
	if root.Left != nil {
		DFS(root.Left, visit)
	}
	if root.Right != nil {
		DFS(root.Right, visit)
	}
}

func BFS(root *obj.TreeNode, visit func(node *obj.TreeNode)) {
	target, next := 0, 1
	queues := make([]*obj.Queue, 2)
	queues[0] = obj.ConstructorOfQueue()
	queues[1] = obj.ConstructorOfQueue()
	queues[0].Push(root)
	for !queues[target].Empty() {
		for !queues[target].Empty() {
			node := queues[target].Pull().(*obj.TreeNode)
			if node.Left != nil {
				queues[next].Push(node.Left)
			}
			if node.Right != nil {
				queues[next].Push(node.Right)
			}
			visit(node)
		}
		target, next = next, target
	}

}
