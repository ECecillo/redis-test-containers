package helper

import (
	"crypto/sha256"
	"testing"
)

func BenchmarkEncryptKeyWithSha256HMAC(b *testing.B) {
	key := "test"
	secret := "mySuperSecret"
	algorithm := sha256.New

	for i := 0; i < b.N; i++ {
		EncryptKey(key, secret, algorithm)
	}
}
