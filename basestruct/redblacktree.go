package basestruct

import (
	"errors"
	"fmt"
)

const RED = true
const BLACK = false

type Color bool

/*
红黑树满足一下规则
1. 每个节点不是红色就是黑色
2.根节点为黑色
3.如果节点为红，其子节点必须为黑
4.任一节点至nil的任何路径，所包含的黑节点数必须相同。
5.叶子节点nil为黑色


注：这款实现代码在root上使用了一个Node作为指针指向root，即root的parent为pointer节点，pointer节点的左孩子为root
化简了对root的操作。
 */

type RedBlackTree struct {
	Val *Comparable
	parent *RedBlackTree
	left *RedBlackTree
	right *RedBlackTree
	color Color
}

func (this *RedBlackTree) isRoot() bool {
	return this.parent != nil && this.parent.parent == nil
}

func leftRotate(node *RedBlackTree) {
	if node.parent == nil {
		panic("This is the pointer of isRoot!")
	}
	//新指针存储当前节点的右孩子
	newRoot := node.right

	//将当前节点的右孩子指向新根的左孩子（当前节点的右-左孙子）
	node.right = newRoot.left
	//若父亲被变更的节点非空，设置其父亲节点
	if newRoot.left != nil {
		newRoot.left.parent = node
	}

	//新根父亲指向源根父亲
	newRoot.parent = node.parent

	//源根父亲孩子指向新根
	if node.parent.left == node {
		node.parent.left = newRoot
	} else {
		node.parent.right = newRoot
	}

	//设旧根为新根的左孩子
	newRoot.left = node

	//设新根为旧根的父亲
	node.parent = newRoot


}

func rightRotate(node *RedBlackTree) {

	if node.parent == nil {
		panic("This is the pointer of isRoot!")
	}

	//新指针存储当前节点的左孩子
	newRoot := node.left

	//将当前节点的左孩子指向新根的右孩子（当前节点的左-右孙子）
	node.left = newRoot.right
	//若父亲被变更的节点非空，设置其父亲节点
	if newRoot.right != nil {
		newRoot.right.parent = node
	}

	//新根父亲指向源根父亲
	newRoot.parent = node.parent

	//源根父亲孩子指向新根
	if node.parent.right == node {
		node.parent.right = newRoot
	} else {
		node.parent.left = newRoot
	}

	//设旧根为新根的右孩子
	newRoot.right = node

	//设新根为旧根的父亲
	node.parent = newRoot


}

func (this *RedBlackTree) Insert(val *Comparable) error {
	if this.parent == nil {

		pre := this
		cur := this.left

		//二叉查找树的正常插入
		for cur != nil {
			pre = cur
			if (*val).CompareTo(cur.Val) < 0 {
				cur = cur.left
			} else {
				cur = cur.right
			}
		}

		//新建节点，指向父节点
		node := &RedBlackTree{Val: val, color: RED, parent: pre}

		if pre.Val == nil {
			this.left = node
			this.left.color = BLACK
			return nil
		}

		//父节点的孩子指向新节点
		if (*val).CompareTo(pre.Val) < 0 {
			pre.left = node
		} else {
			pre.right = node
		}

		//修复
		this.insertFixUp(node)

		return nil
	} else {
		return errors.New("this is not the pointer of isRoot")
	}
}

func (this *RedBlackTree) insertFixUp(cur *RedBlackTree) {
	var parent, grandpa *RedBlackTree
	//当前节点非根，并且父亲非红
	for cur != nil && !cur.isRoot() && cur.parent.Color() == RED {
		parent = cur.Parent()
		grandpa = parent.Parent()

		//若 父节点 是 祖父节点 的左孩子
		if parent == grandpa.Left() {
			//Case1 叔叔是红色，交换父辈和祖辈颜色
			if grandpa.Right() != nil && grandpa.Right().Color() == RED {
				grandpa.Right().setColor(BLACK)
				parent.setColor(BLACK)
				grandpa.setColor(RED)
				cur = grandpa
				continue
			}

			//Case2 叔叔是黑色，当前节点是右孩子，左旋，重新处理父亲
			if parent.Right() == cur {
				leftRotate(parent)
				parent, cur = cur, parent
			}

			//Case3 叔叔是黑色，当前节点是左孩子,交换父辈祖辈颜色，右旋
			parent.setColor(BLACK)
			grandpa.setColor(RED)
			rightRotate(grandpa)

		} else {
			//父节点是祖父节点的右孩子，与上述操作对称
			//Case1 叔叔是红色，交换父辈和祖辈颜色
			if grandpa.Left() != nil && grandpa.Left().Color() == RED {
				grandpa.Left().setColor(BLACK)
				parent.setColor(BLACK)
				grandpa.setColor(RED)
				cur = grandpa
				continue
			}

			//Case2 叔叔是黑色，当前节点是左孩子，右旋，重新处理父亲
			if parent.Left() == cur {
				rightRotate(parent)
				parent, cur = cur, parent
			}

			//Case3 叔叔是黑色，当前节点是右孩子,交换父辈祖辈颜色，右旋
			parent.setColor(BLACK)
			grandpa.setColor(RED)
			leftRotate(grandpa)
		}

		//根节点设为黑色
		this.Left().setColor(BLACK)
	}
}

func (this *RedBlackTree) DeleteVal(val *Comparable) error {
	if this.parent == nil {
		//寻找目标节点
		target := this.left
		comparing := (*target.Val).CompareTo(val)
		for comparing != 0 && target != nil {
			if comparing < 0 {
				target = target.right
			} else {
				target = target.left
			}
			comparing = (*target.Val).CompareTo(val)
		}
		if target == nil {
			return nil
		}

		this.DeleteNode(target)
		return nil
	} else {
		return errors.New("this is not the pointer of isRoot")
	}
}

func (this *RedBlackTree) DeleteNode(target *RedBlackTree) {

	//被删除节点有两个孩子
	if target.left != nil && target.right != nil {
		//寻找后续节点
		replace := target.right
		for replace.left != nil {
			replace = replace.left
		}

		target.Val = replace.Val
		target = replace
	}

	//被删除的只有一个孩子

	var replace *RedBlackTree
	if target.left != nil {
		replace = target.left
	} else {
		replace = target.right
	}

	//如果只有一个孩子节点
	if replace != nil {
		//删除target
		replace.parent = target.parent
		if target == target.parent.left {
			target.parent.left = replace
		} else {
			target.parent.right = replace
		}

		target.left, target.right, target.parent = nil, nil, nil

		if target.color == BLACK {
			this.deleteFixUp(replace)
		}
	} else if target.isRoot() {
		//整棵树只有根节点
		target.left, target.right, target.parent = nil, nil, nil
		this.left = nil
	} else {
		if target.color == BLACK {
			this.deleteFixUp(target)
		}
		if target.parent != nil {
			if target == target.parent.left {
				target.parent.left = nil
			} else if target == target.parent.right{
				target.parent.right = nil
			}
			target.parent = nil
		}
	}

}




func (this *RedBlackTree) deleteFixUp(target *RedBlackTree) {

	var brother,parent *RedBlackTree
	// 循环处理，条件为x不是root节点且是黑色的（因为红色不会对红黑树的性质造成破坏，所以不需要调整
	for target == nil || !target.isRoot() && target.Color() == BLACK {
		if target == nil {
			parent = nil
		} else {
			parent = target.Parent()
		}


		if parent.Left() == target {
			//目标是父亲的左孩子
			brother = parent.Right()
			if brother.Color() == RED {
				// Case 1 兄弟是红色的
				brother.setColor(BLACK)
				parent.setColor(RED)
				leftRotate(parent)
				brother = parent.Right()
			}

			if (brother.Left().Color() == BLACK) &&
				(brother.Right().Color() == BLACK) {
				//Case 2: x的兄弟是黑色，且兄弟的俩个孩子也都是黑色的
				brother.setColor(RED)
				target = parent
				parent = target.Parent()
			} else {
				if brother.Right().Color() == BLACK {
					// Case 3: x的兄弟是黑色的，并且兄弟的左孩子是红色，右孩子为黑色。
					brother.Left().setColor(BLACK)
					brother.setColor(RED)
					rightRotate(brother)
					brother = parent.Right()
				}
				// Case 4: x的兄弟是黑色的；并且兄弟的右孩子是红色的，左孩子任意颜色。
				brother.setColor(parent.Color())
				parent.setColor(BLACK)
				brother.Right().setColor(BLACK)
				leftRotate(parent)
				target = this.Left()
				break
			}

		} else {
			//目标是父亲的右孩子
			brother = parent.Left()
			if brother.Color() == RED {
				// Case 1 兄弟是红色的
				brother.setColor(BLACK)
				parent.setColor(RED)
				rightRotate(parent)
				brother = parent.Left()
			}

			if (brother.Left() != nil || brother.Left().Color() == BLACK) &&
				(brother.Right() != nil || brother.Right().Color() == BLACK) {
				//Case 2: x的兄弟是黑色，且兄弟的俩个孩子也都是黑色的
				brother.setColor(RED)
				target = parent
				parent = target.Parent()
			} else {
				if brother.Left() != nil || brother.Left().Color() == BLACK {
					// Case 3: x的兄弟是黑色的，并且兄弟的左孩子是红色，右孩子为黑色。
					brother.Right().setColor(BLACK)
					brother.setColor(RED)
					leftRotate(brother)
					brother = parent.Left()
				}
				// Case 4: x的兄弟是黑色的；并且兄弟的右孩子是红色的，左孩子任意颜色。
				brother.setColor(parent.Color())
				parent.setColor(BLACK)
				brother.Left().setColor(BLACK)
				rightRotate(parent)
				target = this.Right()
				break
			}
		}
	}
	if target != nil {
		target.setColor(BLACK)
	}
}


func RBTreePrint(root *RedBlackTree, way int) {
	fmt.Print("{")
	if way < 0 {
		if root.color {
			fmt.Print("R")
		} else {
			fmt.Print("B")
		}
		fmt.Print(*root.Val)
		fmt.Print(" ")
		if root.left != nil {
			RBTreePrint(root.left, way)
		} else {
			fmt.Print("nil")
		}
		if root.right != nil {
			RBTreePrint(root.right, way)
		} else {
			fmt.Print("nil")
		}
	} else if way > 0 {
		if root.left != nil {
			RBTreePrint(root.left, way)
		} else {
			fmt.Print("nil")
		}
		if root.right != nil {
			RBTreePrint(root.right, way)
		} else {
			fmt.Print("nil")
		}
		if root.color {
			fmt.Print("R")
		} else {
			fmt.Print("B")
		}
		fmt.Print(*root.Val)
		fmt.Print(" ")
	} else {
		if root.left != nil {
			RBTreePrint(root.left, way)
		} else {
			fmt.Print("nil")
		}
		if root.color {
			fmt.Print("R")
		} else {
			fmt.Print("B")
		}
		fmt.Print(*root.Val)
		fmt.Print(" ")
		if root.right != nil {
			RBTreePrint(root.right, way)
		} else {
			fmt.Print("nil")
		}
	}

	fmt.Print("}")
}


func (this *RedBlackTree) GetRoot() (*RedBlackTree,error) {
	if this.parent != nil {
		return nil,errors.New("This is not the pointer of isRoot")
	}
	return this.left,nil
}

func (this *RedBlackTree) Left() *RedBlackTree {
	if this == nil {
		return nil
	}
	return this.left
}

func (this *RedBlackTree) Parent() *RedBlackTree {
	if this == nil {
		return nil
	}
	return this.parent
}

func (this *RedBlackTree) Right() *RedBlackTree {
	if this == nil {
		return nil
	}
	return this.right
}

func (this *RedBlackTree) setColor(color Color) {
	if this != nil {
		this.color = color
	}
}

func (this *RedBlackTree) Color() Color {
	if this == nil {
		return BLACK
	}
	return this.color
}