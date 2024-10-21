package nodeandlistops

import errorops "LinkedList/elements/errorOps"

const MAX_LIST_SIZE = 10

func AddNodesToList(companies map[int]string, list *LinkedList, addingNodesQuantity int) *LinkedList {
	errorops.IsNodesQuantityExceedMaxLenList(addingNodesQuantity, len(companies))

	if isAddingNodesQuantityExceedOrEqualMaxLenList(MAX_LIST_SIZE, addingNodesQuantity) {
		return tenNodeList(companies, MAX_LIST_SIZE)
	}

	listSize := NodeCounter(list)
	nodesExceedingMax := quantityNodesExceedsMaxLenList(addingNodesQuantity, listSize, MAX_LIST_SIZE)
	nodesToAdd := quantityNodesToAdd(addingNodesQuantity, listSize, MAX_LIST_SIZE)

	if isOverallNodesQuantityExceedsMaxLenList(nodesExceedingMax) {
		list = cutLinkedList(nodesToAdd, list)
	}

	LinkedListConstructor(companies, list, addingNodesQuantity)

	return list
}
