package crypto_test

import (
	"bit-exporter/pkg/crypto"
	mathrend "math/rand"
	"testing"
)

const CiphersCount = 1000

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mathrend.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateTestSet(key []byte, count int) ([]crypto.CipherString, []string, error) {
	ciphers := make([]crypto.CipherString, count)
	messages := make([]string, count)
	var message string
	for i := 0; i < count; i++ {
		message = RandStringRunes(32)
		messages[i] = message
		k, err := crypto.Encrypt([]byte(message), crypto.AesCbc256_B64, key, nil)
		if err != nil {
			return ciphers, messages, err
		}
		ciphers[i] = k
	}
	return ciphers, messages, nil
}

func BenchmarkDecrypt(b *testing.B) {
	const testSetSize = 1000
	key, err := crypto.CalculateUserKey("password", "user@mail.com", 0, 50000, 0, 0)
	ciphers, messages, err := generateTestSet(key, testSetSize)
	if err != nil {
		b.Fatal(b)
	}
	var message []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testSetSize; j++ {
			message, err = crypto.Decrypt(ciphers[j], key, nil)
			if err != nil {
				b.Fatal(b)
			}
			if string(message) != messages[j] {
				b.Fatalf("Wrong message: expected '%s' got '%s'", messages[j], message)
			}
		}
	}
}
