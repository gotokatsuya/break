package crypto

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"strings"
)

// GenerateRandom ...
func GenerateRandom(keyLength int) (string, error) {
	k := make([]byte, keyLength)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return "", err
	}
	return strings.TrimRight(base32.StdEncoding.EncodeToString(k), "="), nil
}
