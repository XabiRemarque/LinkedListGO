package nodeandlistops

func quantityNodesExceedsMaxLenList(addingNodesQuantity, listSize, MAX_LIST_SIZE int) int {
	if listSize+addingNodesQuantity > MAX_LIST_SIZE {
		return (listSize + addingNodesQuantity) - MAX_LIST_SIZE
	}
	return 0
}

func isOverallNodesQuantityExceedsMaxLenList(nodesExceedMaxListSize int) bool {
	return nodesExceedMaxListSize > 0
}
