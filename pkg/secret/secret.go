package secret

import (
	"crypto/rand"
	"math/big"
)

const (
	letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// Generate simply generates random secrets or passwords.
func Generate(length int) string {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(
			rand.Reader,
			big.NewInt(int64(len(letters))),
		)

		if err != nil {
			return ""
		}

		result[i] = letters[num.Int64()]
	}

	return string(result)
}
