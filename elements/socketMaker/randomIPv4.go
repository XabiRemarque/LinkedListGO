package socketmaker

import (
	"math/rand"
	"strconv"
	"strings"
)

func randomIPv4() string {
	slice := make([]string, 4)
	for i := 0; i < len(slice); i++ {
		slice[i] = strconv.Itoa(rand.Intn(256))
	}
	ip := strings.Join(slice, ".")
	return ip
}
