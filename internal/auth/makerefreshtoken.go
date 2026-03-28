package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() string {
	// generate 32 bytes of random data
	b := make([]byte, 32)
	_, _ = rand.Read(b)

	// convert to a hex string
	return hex.EncodeToString(b)
}
