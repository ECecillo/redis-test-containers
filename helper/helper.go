package helper

import (
	"crypto/hmac"
	"encoding/hex"
	"hash"
)

func EncryptKey(key string, secret string, algorithm func() hash.Hash) string {
	hmac := hmac.New(algorithm, []byte(secret))
	hmac.Write([]byte(key))

	encryptedKey := hmac.Sum(nil)
	hmacHex := hex.EncodeToString(encryptedKey)

	return hmacHex
}
