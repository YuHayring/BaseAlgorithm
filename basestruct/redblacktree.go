package basestruct

/*
红黑树满足一下规则
1. 每个节点不是红色就是黑色
2.根节点为黑色
3.如果节点为红，其子节点必须为黑
4.任一节点至nil的任何路径，所包含的黑节点数必须相同。
5.叶子节点nil为黑色
 */

type RedBlackTree struct {
	parent *RedBlackTree
	left *RedBlackTree
	right *RedBlackTree
	prev *RedBlackTree
	red bool
}