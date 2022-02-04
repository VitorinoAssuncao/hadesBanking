package pgtest

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func GetRandomDBName() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	return fmt.Sprintf("db_%d", n)
}
