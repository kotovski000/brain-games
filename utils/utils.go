package utils

import (
	"crypto/rand"
	"math/big"
)

func RandomInt(max int64) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return int(n.Int64())
}
