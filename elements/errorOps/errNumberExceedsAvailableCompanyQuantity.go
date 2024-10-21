package errorops

import (
	"fmt"
	"log"
)

type ErrNumberExceedsAvailableCompanyQuantity struct {
	Requested int
	Rest      int
}

func (e *ErrNumberExceedsAvailableCompanyQuantity) Error() string {
	return fmt.Sprintf("Rest %d, Req %d - nodes quantity exceeds max company quantity by %d nodes", e.Rest, e.Requested, e.Requested-e.Rest)
}

func IsNodesQuantityExceedMaxLenList(requestedNum, companiesLength int) {
	if requestedNum > companiesLength {
		err := &ErrNumberExceedsAvailableCompanyQuantity{Requested: requestedNum, Rest: companiesLength}
		log.Fatalf("Error: %v", err)
	}
}
