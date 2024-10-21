package socketmaker

import (
	"math/rand"
)

func randomPort() int {
	return rand.Intn(1024)
}
