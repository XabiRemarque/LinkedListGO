package nodeandlistops

func nodeInsertion(list *LinkedList, newNode *Node) *LinkedList {
	if list.head == nil {
		list.head = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
	return list
}
