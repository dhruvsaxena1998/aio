// cmd/internal/helpers/apikey_generator.go

package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

const defaultAPIKeyLength = 24
const defaultHyphenInterval = 8

// GenerateAPIKey generates a random API key of a specified length (default is 24 characters).
func GenerateAPIKey(length int, hyphenInterval int) (string, error) {
	if length <= 0 {
		length = defaultAPIKeyLength
	}
	if hyphenInterval <= 0 {
		hyphenInterval = defaultHyphenInterval
	}

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	hexKey := hex.EncodeToString(bytes)
	var result strings.Builder
	for i, r := range hexKey {
		result.WriteRune(r)
		if (i+1)%hyphenInterval == 0 && i < len(hexKey)-1 {
			result.WriteRune('-')
		}
	}

	return result.String(), nil
}

// GenerateDefaultAPIKey generates a random API key with the default length (24 characters).
func GenerateDefaultAPIKey() (string, error) {
	return GenerateAPIKey(defaultAPIKeyLength, defaultHyphenInterval)
}
