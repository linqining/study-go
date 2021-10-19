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
