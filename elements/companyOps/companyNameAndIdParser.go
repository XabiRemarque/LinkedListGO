package companyops

import (
	"math/rand"
)

func CompanyNameAndIdParser(companies *map[int]string) (string, int) {
	length := len(*companies)

	randomIndex := rand.Intn(length)
	index := 0

	for id, name := range *companies {
		if index == randomIndex {
			delete(*companies, id)
			return name, id
		}
		index++
	}
	return "", 0
}
