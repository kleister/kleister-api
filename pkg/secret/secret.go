package secret

import (
	"crypto/rand"
	"math/big"
)

const (
	letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// Generate simply generates random secrets or passwords and returns a string.
func Generate(length int) string {
	return string(Bytes(length))
}

// Bytes simply generates random secrets or passwords and returns plain bytes.
func Bytes(length int) []byte {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(
			rand.Reader,
			big.NewInt(int64(len(letters))),
		)

		if err != nil {
			return []byte{}
		}

		result[i] = letters[num.Int64()]
	}

	return result
}
