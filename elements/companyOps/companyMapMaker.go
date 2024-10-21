package companyops

import "strings"

func MapMaker() map[int]string {
	companies := "Apple, Google, Microsoft, Amazon, Facebook, Tesla, Netflix, IBM, Samsung, Intel, Oracle, Adobe, Salesforce, Twitter, Uber, Airbnb, Spotify, Zoom, NVIDIA, PayPal, LinkedIn, Snap, Lyft, Dropbox, Slack"

	companiesSlice := strings.Split(companies, ", ")

	companiesMap := make(map[int]string)

	for i, el := range companiesSlice {
		i++
		companiesMap[i] = el
	}

	return companiesMap
}
