package mqtt

import (
	"crypto/sha256"
	"encoding/hex"
)

// encryptPass take a plaintext password and return sha256 password
func encryptPass(p string) string {
	h := sha256.New()
	h.Write([]byte(p))
	return hex.EncodeToString(h.Sum(nil))
}
