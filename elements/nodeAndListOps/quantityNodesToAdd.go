package nodeandlistops

func quantityNodesToAdd(addingNodesQuantity, listSize, MAX_LIST_SIZE int) int {
	nodesExceedMaxListSize := quantityNodesExceedsMaxLenList(addingNodesQuantity, listSize, MAX_LIST_SIZE)
	nodesToAdd := listSize - nodesExceedMaxListSize

	if nodesToAdd < 0 {
		return 0
	}

	return nodesToAdd
}
