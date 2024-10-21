package main

import (
	companyops "LinkedList/elements/companyOps"
	nodeandlistops "LinkedList/elements/nodeAndListOps"
	"fmt"
)

func main() {
	companies := companyops.MapMaker()

	LinkedList := nodeandlistops.NewLinkedList()

	firstQuantity := nodeandlistops.EnterNodesToAdd("1st")

	addedLinkedList := nodeandlistops.LinkedListConstructor(companies, LinkedList, firstQuantity)

	secondQuantity := nodeandlistops.EnterNodesToAdd("2nd")

	addedLinkedList = nodeandlistops.AddNodesToList(companies, addedLinkedList, secondQuantity)

	fmt.Printf("\nNodes quantity: %d\n\n", nodeandlistops.NodeCounter(addedLinkedList))

	addedLinkedList.NodeOutput()
}
