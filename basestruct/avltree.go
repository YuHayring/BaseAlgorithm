package basestruct

import "../util"

type AVLTree struct {
	Val    int
	Height int
	Left   *AVLTree
	Right  *AVLTree
}

//搜索
func (this *AVLTree) Search(value int) *AVLTree {
	if this == nil {
		return nil
	}
	if value < this.Val {
		return this.Left.Search(value)
	} else if value > this.Val {
		return this.Right.Search(value)
	} else {
		return this
	}
}

func (this *AVLTree) leftRotate() *AVLTree {
	newRoot := this.Right
	this.Right, newRoot.Left = newRoot.Left, this
	this.Height = util.MaxInt(this.Left.Height, this.Right.Height) + 1
	newRoot.Height = util.MaxInt(newRoot.Left.Height, newRoot.Right.Height) + 1
	return newRoot
}

func (this *AVLTree) rightRotate() *AVLTree {
	newRoot := this.Left
	this.Left, newRoot.Right = newRoot.Right, this
	this.Height = util.MaxInt(this.Left.Height, this.Right.Height) +1
	newRoot.Height = util.MaxInt(newRoot.Left.Height, newRoot.Right.Height) + 1
	return newRoot
}

func (this *AVLTree) rightLeftRotate() *AVLTree {
	this.Right = this.Right.rightRotate()
	return this.leftRotate()
}

func (this *AVLTree) leftThenRightRotate() *AVLTree {
	this.Left = this.Left.leftRotate()
	return this.rightRotate()
}

func (this *AVLTree) makeBalance() *AVLTree {
	var newRoot *AVLTree
	if this.Right.Height - this.Left.Height == 2 {
		if this.Right.Right.Height > this.Right.Left.Height {
			newRoot = this.leftRotate()
		}else {
			newRoot = this.rightLeftRotate()
		}
	}else if this.Left.Height - this.Right.Height == 2 {
		if this.Left.Left.Height > this.Left.Right.Height {
			newRoot = this.rightRotate()
		} else {
			newRoot = this.leftThenRightRotate()
		}
	}
	return newRoot
}

func (this *AVLTree) Insert(value int) *AVLTree {
	if this == nil {
		return &AVLTree{value,1,nil,nil}
	}
	var newRoot *AVLTree
	if value < this.Val {
		this.Left = this.Left.Insert(value)
		newRoot = this.makeBalance()
	}else if value > this.Val {
		this.Right = this.Right.Insert(value)
		newRoot = this.makeBalance()
	}else {
		newRoot = this
		return newRoot
	}
	newRoot.Height = util.MaxInt(this.Left.Height, this.Right.Height) + 1
	return newRoot
}

func (this *AVLTree) Delete(value int) *AVLTree {
	if this == nil {
		return this
	}
	if value < this.Val {
		this.Left = this.Left.Delete(value)
		this.Height = util.MaxInt(this.Left.Height, this.Right.Height) + 1
		return this.makeBalance()
	}else if value > this.Val{
		this.Right = this.Right.Delete(value)
		this.Height = util.MaxInt(this.Left.Height, this.Right.Height) + 1
		return this.makeBalance()
	}else {
		newRoot := this
		if this.Left != nil && this.Right != nil {
			rightMin := this.Right
			for rightMin.Left.Left != nil {
				rightMin = rightMin.Left
			}
			this.Val = rightMin.Val
			this.Right = this.Right.Delete(this.Val)
			newRoot = this
		} else if this.Left !=nil {
			newRoot = this.Left
		}else {//只有一个右孩子或没孩子
			newRoot = this.Right
		}
		if newRoot != nil {
			newRoot.Height = util.MaxInt(this.Left.Height, this.Right.Height) + 1
			return newRoot.makeBalance()
		} else {
			return nil
		}
	}
}



