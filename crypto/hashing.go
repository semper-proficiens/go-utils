package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashContentToString hashes some content in []byte and returns the hash in a string
func HashContentToString(content []byte) string {
	hash := sha256.Sum256(content)
	return hex.EncodeToString(hash[:])
}
