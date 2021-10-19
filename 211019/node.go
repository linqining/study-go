package _11019

import (
	"strconv"
	"strings"
)

type Node struct{
	value int
	next *Node
}

func (n *Node) String()string{
	sb :=strings.Builder{}
	for n!=nil{
		sb.WriteString(strconv.Itoa(n.value))
		n = n.next
		if n!=nil{
			sb.WriteString("->")
		}
	}
	return sb.String()
}

func NewLinkList(ints []int)*Node{
	tmpNode := &Node{}
	head := tmpNode
	for _, item :=range ints{
		node := &Node{value: item}
		head.next = node
		head = head.next
	}
	rNode := tmpNode.next
	tmpNode.next=nil
	return rNode
}

func RemoveMiddleNode(head *Node) *Node {
	if head ==nil || head.next ==nil{
		return head;
	}
	// 只有一个节点
	if head.next ==nil{
		return head.next;
	}
	// 双指针，一个移动一步，一个移动两步
	pre := head
	cur := head.next.next
	for{
		if cur.next!=nil && cur.next.next !=nil{
			pre = pre.next
			cur = cur.next.next
		}else{
			break
		}
	}
	pre.next = pre.next.next
	return head
}

// 翻转链表，递归方式
func Reverse(head *Node) *Node{
	if head ==nil || head.next==nil{
		return head
	}
	cur := head
	tmp := head.next
	node := Reverse(head.next)
	tmp.next = cur
	return node
}


// 反转链表，非递归方式
func ReverseNoRecur(head *Node) *Node{
	if head ==nil || head.next ==nil{
		return head
	}
	var pre *Node

	for head.next!=nil && head.next.next!=nil{
		nextNode := head.next
		cur := head.next.next
		head.next = pre
		nextNode.next = head
		pre = nextNode
		head = cur
	}
	if head.next ==nil{
		head.next = pre
		return head
	}else{
		tmp := head.next
		head.next.next=head
		head.next=pre
		return tmp
	}
}

// 反转链表，非递归方式
func ReverseNoRecur2(head *Node) *Node{
	var pre,next *Node
	for head!=nil{
		next = head.next
		head.next = pre
		pre=head
		head = next
	}
	return pre
}

// 判断是否回文链表
func IsPalindrome(head *Node) (isPalindrome bool){
	isPalindrome = true
	if head ==nil {
		return false
	}
	begin := head
	nCount :=0
	for begin!=nil{
		nCount++
		begin = begin.next
	}
	if nCount == 1{
		return true
	}
	move := nCount/2

	var pre, cur,next,lf,rf *Node
	cur = head
	for i:=0;i<move;i++{
		next = cur.next
		cur.next=pre
		pre = cur
		cur = next
	}
	if nCount&1 ==1{
		// 奇数
		lf = pre
		rf = cur.next
	}else{
		// 偶数
		lf = pre
		rf = cur
	}

	var revCur *Node
	revCur = cur
	for i:=0;i<move;i++{
		if lf.value != rf.value && isPalindrome{
			isPalindrome = false
		}
		// 优化，判断后不改变原来的结构，将链表反转回来
		tmp := lf
		lf = lf.next
		tmp.next=revCur
		revCur = tmp
		rf = rf.next
	}
	return
}

func Pivot(head *Node,pivot int) *Node{
	if head == nil{
		return head
	}
	begin := head

	pivNode := &Node{value: pivot}
	var prev,next,headPev *Node // prev相对于pivNode的前后节点,next 相对于head的pev


	for head!=nil{
		if head.value<=pivNode.value{
			if prev==nil{
				prev = head
				begin = prev
			}else{
				prev.next = head
				prev = head
			}
			next=head.next
			prev.next=pivNode
			if headPev!=nil{
				headPev.next=next
			}
			head = next
		}else{
			headPev = head
			if pivNode.next==nil{
				pivNode.next = head
			}
			head = head.next
		}
	}
	// 断开pivotNode 链接
	if prev!=nil{
		prev.next=pivNode.next
	}
	pivNode.next = nil
	return begin
}

// 删除无序节点单链表重复出现的节点
func RemoveDuplicate(head *Node)*Node{
	if head ==nil{
		return head
	}
	begin := head
	valMap := make(map[int]bool)
	var prev *Node
	for head!=nil{
		_,ok:=valMap[head.value]
		// 存在重复的值
		if ok{
			// 第一次进来一定是走下面，因此prev一定不是nil
			prev.next = head.next
			head.next=nil
			head = prev.next
		}else{
			valMap[head.value] = true
			prev = head
			head = head.next
		}
	}
	return begin
}

func SelectionSort(n *Node)(bNode *Node,eNode *Node){
	if n==nil{
		return nil,nil
	}
	if n.next==nil{
		return n,n
	}

	prev := &Node{value: 0,next: n} //虚拟节点
	start := prev
	sepNode :=n
	last:=n
	n = n.next
	for n!=nil{
		tmp := n.next
		if n.value<= sepNode.value{
			last.next = n.next
			prev.next=n
			prev.next.next = sepNode
			prev = prev.next
		}else{
			last = n
		}
		n = tmp
	}

	lb := start.next
	start.next =nil
	prev.next = nil
	var ls, le *Node
	if lb!=nil && lb!=sepNode{
		ls,le = SelectionSort(lb)
		le.next = sepNode
	}else{
		ls = sepNode
	}

	rb := sepNode.next
	var rs, re *Node
	if rb!=nil{
		sepNode.next = nil
		rs, re = SelectionSort(rb)
		sepNode.next=rs
	}else{
		re = last
	}
	return ls, re
}