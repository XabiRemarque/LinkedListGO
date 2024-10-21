package nodeandlistops

import (
	companyops "LinkedList/elements/companyOps"
	errorops "LinkedList/elements/errorOps"
)

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func LinkedListConstructor(companies map[int]string, list *LinkedList, nodesToAdd int) *LinkedList {
	errorops.IsNodesQuantityExceedMaxLenList(nodesToAdd, len(companies))

	for i := 0; i < nodesToAdd; i++ {
		companyName, companyID := companyops.CompanyNameAndIdParser(&companies)
		newNode := newNode(companyName, companyID)
		nodeInsertion(list, newNode)
	}
	return list
}
