package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"math"
)

func Encrypt(data []byte, typ CipherStringType, key, macKey []byte) (CipherString, error) {
	s := CipherString{}
	switch typ {
	case AesCbc256_B64, AesCbc256_HmacSha256_B64:
	default:
		return s, fmt.Errorf("encrypt: unsupported cipher type %q", s.Type)
	}
	s.Type = typ
	data = padPKCS7(data, aes.BlockSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		return s, err
	}
	s.IV = make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, s.IV); err != nil {
		return s, err
	}
	s.CT = make([]byte, len(data))
	mode := cipher.NewCBCEncrypter(block, s.IV)
	mode.CryptBlocks(s.CT, data)

	if typ == AesCbc256_HmacSha256_B64 {
		if len(macKey) == 0 {
			return s, fmt.Errorf("encrypt: cipher string type expects a MAC")
		}
		var macMessage []byte
		macMessage = append(macMessage, s.IV...)
		macMessage = append(macMessage, s.CT...)
		mac := hmac.New(sha256.New, macKey)
		mac.Write(macMessage)
		s.MAC = mac.Sum(nil)
	}

	return s, nil
}

func padPKCS7(src []byte, size int) []byte {
	// Note that we always pad, even if rem==0. This is because unpad must
	// always remove at least one byte to be unambiguous.
	rem := len(src) % size
	n := size - rem
	if n > math.MaxUint8 {
		panic(fmt.Sprintf("cannot pad over %d bytes, but got %d", math.MaxUint8, n))
	}
	padded := make([]byte, len(src)+n)
	copy(padded, src)
	for i := len(src); i < len(padded); i++ {
		padded[i] = byte(n)
	}
	return padded
}
