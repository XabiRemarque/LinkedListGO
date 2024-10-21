package nodeandlistops

func isAddingNodesQuantityExceedOrEqualMaxLenList(MAX_LIST_SIZE, addingNodesQuantity int) bool {
	return addingNodesQuantity >= MAX_LIST_SIZE
}

func tenNodeList(companies map[int]string, MAX_LIST_SIZE int) *LinkedList {
	return LinkedListConstructor(companies, NewLinkedList(), MAX_LIST_SIZE)
}
