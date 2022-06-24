package rand

import (
	"math/rand"
	"time"
)

func R() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
