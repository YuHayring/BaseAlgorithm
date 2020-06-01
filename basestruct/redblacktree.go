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
 */

type RedBlackTree struct {
	Val *Comparable
	parent *RedBlackTree
	left *RedBlackTree
	right *RedBlackTree
	color Color
}

func (this *RedBlackTree) root() bool {
	return this.parent.parent == nil
}

func leftRotate(node *RedBlackTree) {
	if node.parent == nil {
		panic("This is the pointer of root!")
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
		panic("This is the pointer of root!")
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
		return errors.New("this is not the pointer of root")
	}
}

func (this *RedBlackTree) insertFixUp(cur *RedBlackTree) {
	var parent, grandpa *RedBlackTree
	//当前节点非根，并且非红
	for !cur.root() && cur.parent.color == RED {
		parent = cur.parent
		grandpa = parent.parent

		//若 父节点 是 祖父节点 的左孩子
		if parent == grandpa.left {
			//Case1 叔叔是红色，交换父辈和祖辈颜色
			if grandpa.right != nil && grandpa.right.color == RED {
				grandpa.right.color = BLACK
				       parent.color = BLACK
				grandpa.color = RED
				cur = grandpa
				continue
			}

			//Case2 叔叔是黑色，当前节点是右孩子，左旋，重新处理父亲
			if parent.right == cur {
				leftRotate(parent)
				parent, cur = cur, parent
			}

			//Case3 叔叔是黑色，当前节点是左孩子,交换父辈祖辈颜色，右旋
			parent.color = BLACK
			grandpa.color = RED
			rightRotate(grandpa)

		} else {
			//父节点是祖父节点的右孩子，与上述操作对称
			//Case1 叔叔是红色，交换父辈和祖辈颜色
			if grandpa.left != nil && grandpa.left.color == RED {
				grandpa.left.color = BLACK
				parent.color = BLACK
				grandpa.color = RED
				cur = grandpa
				continue
			}

			//Case2 叔叔是黑色，当前节点是左孩子，右旋，重新处理父亲
			if parent.left == cur {
				rightRotate(parent)
				parent, cur = cur, parent
			}

			//Case3 叔叔是黑色，当前节点是右孩子,交换父辈祖辈颜色，右旋
			parent.color = BLACK
			grandpa.color = RED
			leftRotate(grandpa)
		}

		//根节点设为黑色
		this.left.color = BLACK
	}
}

func (this *RedBlackTree) DeleteVal(val *Comparable) error {
	if this.parent == nil {
		//寻找目标节点
		target := this.left
		comparing := (*target.Val).CompareTo(val)
		for comparing != 0 && target != nil {
			if comparing < 0 {
				target = target.left
			} else {
				target = target.right
			}
		}
		if target == nil {
			return nil
		}

		this.DeleteNode(target)
		return nil
	} else {
		return errors.New("this is not the pointer of root")
	}
}

func (this *RedBlackTree) DeleteNode(target *RedBlackTree) {

	var replacesChild *RedBlackTree
	var replacesParent *RedBlackTree
	var replacesColor Color

	//被删除节点有两个孩子
	if target.left != nil && target.right != nil {
		//寻找后续节点
		replace := target.right
		for replace.left != nil {
			replace = replace.left
		}

		//节点替换
		if target.parent.left == target {
			target.parent.left = replace
		} else {
			target.parent.right = replace
		}


		//需要调整的节点
		replacesChild = replace.right
		replacesParent = replace.parent

		replacesColor = replace.color

		//将替代节点删除后的缺口连接起来
		if replacesParent == target {
			//被删除节点是后续节点的父节点
			replacesParent = replace
		} else {
			//连接右子与父亲
			if replacesChild != nil {
				replacesChild.parent = replacesParent
			}
			replacesParent.left = replacesChild;

			replace.right = target.right
			target.right.parent = replace
		}

		//连接替代节点和其父亲
		replace.parent = target.parent
		replace.color = target.color
		replace.left = target.left

		target.left.parent = replace

		if replacesColor == BLACK {
			this.deleteFixUp(replacesChild)
		}
		return
	}

	//被删除的只有一个孩子

	if !(target.left == nil && target.right == nil ) {
		if target.left != nil {
			replacesChild = target.left
		} else {
			replacesChild = target.right
		}

		replacesParent = target.parent
		replacesColor = target.color

		if replacesChild != nil {
			replacesChild.parent = replacesParent
		}

		if replacesParent.left == target {
			replacesParent.left = replacesChild
		} else {
			replacesParent.right = replacesChild
		}

		if replacesColor == BLACK {
			this.deleteFixUp(replacesChild)
		}

		return
	}

	//被删除节点没有孩子
	if target.root() {
		this.left = nil
		target.parent = nil
	} else {
		//直接调整
		if target.color == BLACK {
			this.deleteFixUp(target)
		}
		//删除节点
		if target.parent != nil {
			if target == target.parent.left {
				target.parent.left = nil
			} else if target == target.parent.right {
				target.parent.right = nil
			}
			target.parent = nil
		}
	}



}




func (this *RedBlackTree) deleteFixUp(target *RedBlackTree) {
	
	var brother,parent *RedBlackTree
	// 循环处理，条件为x不是root节点且是黑色的（因为红色不会对红黑树的性质造成破坏，所以不需要调整
	for !target.root() && ( target == nil || target.color == BLACK ) {
		parent = target.parent
		if parent.left == target {
			//目标是父亲的左孩子
			brother = parent.right
			if brother.color == RED {
				// Case 1 兄弟是红色的
				brother.color = BLACK
				parent.color = RED
				leftRotate(parent)
				brother = parent.right
			}

			if (brother.left != nil || brother.left.color == BLACK) &&
				(brother.right != nil || brother.right.color == BLACK) {
				//Case 2: x的兄弟是黑色，且兄弟的俩个孩子也都是黑色的
				brother.color = RED
				target = parent
				parent = target.parent
			} else {
				if brother.right != nil || brother.right.color == BLACK {
					// Case 3: x的兄弟是黑色的，并且兄弟的左孩子是红色，右孩子为黑色。
					brother.left.color = BLACK
					brother.color = RED
					rightRotate(brother)
					brother = parent.right
				}
				// Case 4: x的兄弟是黑色的；并且兄弟的右孩子是红色的，左孩子任意颜色。
				brother.color = parent.color
				parent.color = BLACK
				brother.right.color = BLACK
				leftRotate(parent)
				target = this.left
				break
			}

		} else {
			//目标是父亲的右孩子
			brother = parent.left
			if brother.color == RED {
				// Case 1 兄弟是红色的
				brother.color = BLACK
				parent.color = RED
				rightRotate(parent)
				brother = parent.left
			}

			if (brother.left != nil || brother.left.color == BLACK) &&
				(brother.right != nil || brother.right.color == BLACK) {
				//Case 2: x的兄弟是黑色，且兄弟的俩个孩子也都是黑色的
				brother.color = RED
				target = parent
				parent = target.parent
			} else {
				if brother.left != nil || brother.left.color == BLACK {
					// Case 3: x的兄弟是黑色的，并且兄弟的左孩子是红色，右孩子为黑色。
					brother.right.color = BLACK
					brother.color = RED
					leftRotate(brother)
					brother = parent.left
				}
				// Case 4: x的兄弟是黑色的；并且兄弟的右孩子是红色的，左孩子任意颜色。
				brother.color = parent.color
				parent.color = BLACK
				brother.left.color = BLACK
				rightRotate(parent)
				target = this.right
				break
			}
		}
	}
	if target != nil {
		target.color = BLACK
	}
}


func RBTreePrint(root *RedBlackTree, way int) {
	if way < 0 {
		fmt.Print(*root.Val)
		fmt.Print(" ")
		if root.left != nil {
			RBTreePrint(root.left, way)
		}
		if root.right != nil {
			RBTreePrint(root.right, way)
		}
	} else if way > 0 {
		if root.left != nil {
			RBTreePrint(root.left, way)
		}
		if root.right != nil {
			RBTreePrint(root.right, way)
		}
		fmt.Print(*root.Val)
		fmt.Print(" ")
	} else {
		if root.left != nil {
			RBTreePrint(root.left, way)
		}
		fmt.Print(*root.Val)
		fmt.Print(" ")
		if root.right != nil {
			RBTreePrint(root.right, way)
		}
	}
}


func (this *RedBlackTree) GetRoot() (*RedBlackTree,error) {
	if this.parent != nil {
		return nil,errors.New("This is not the pointer of root")
	}
	return this.left,nil
}