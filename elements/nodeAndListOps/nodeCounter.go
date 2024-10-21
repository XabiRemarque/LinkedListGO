package nodeandlistops

func NodeCounter(list *LinkedList) (count int) {
	if list.head == nil {
		return
	}
	current := list.head
	for current != nil {
		current = current.next
		count++
	}
	return count
}
