package _11019

type Node struct{
	value int
	next *Node
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