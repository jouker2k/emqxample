package mqtt

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// encryptPass take a plaintext password and return bcrypt hash
func encryptPass(p string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	if err != nil {
		log.Error(err)
	}
	return string(h)
}
