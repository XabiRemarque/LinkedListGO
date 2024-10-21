package nodeandlistops

func cutLinkedList(nodesToAdd int, list *LinkedList) *LinkedList {
	if isOverallNodesQuantityExceedsMaxLenList(nodesToAdd) {
		current := list.head
		for i := 1; i < nodesToAdd; i++ {
			if current.next != nil {
				current = current.next
			}
		}
		current.next = nil
	}
	return list
}
