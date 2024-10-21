package nodeandlistops

import (
	"fmt"
)

func EnterNodesToAdd(ordinalNumber string) (nodesQuantity int) {
	fmt.Print("Enter nodes quantity to add: ")

	for {
		rightInput, err := fmt.Scan(&nodesQuantity)
		if err != nil || rightInput != 1 || nodesQuantity == 0 {
			fmt.Print("Try again! You should enter a number bigger than 0: ")
			continue
		}
		break
	}

	fmt.Printf("Your %s number: %d\n", ordinalNumber, nodesQuantity)
	return nodesQuantity
}
