package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func Decrypt(s CipherString, key, keyMac []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	switch s.Type {
	case AesCbc256_B64, AesCbc256_HmacSha256_B64:
		// continues below
	default:
		return nil, fmt.Errorf("decrypt: unsupported cipher type %q", s.Type)
	}

	if s.Type == AesCbc256_HmacSha256_B64 {
		if len(s.MAC) == 0 || len(keyMac) == 0 {
			return nil, fmt.Errorf("decrypt: cipher string type expects a MAC")
		}
		var msg []byte
		msg = append(msg, s.IV...)
		msg = append(msg, s.CT...)
		if !validMAC(msg, s.MAC, keyMac) {
			return nil, fmt.Errorf("decrypt: MAC mismatch")
		}
	}

	mode := cipher.NewCBCDecrypter(block, s.IV)
	dst := make([]byte, len(s.CT))
	mode.CryptBlocks(dst, s.CT)
	dst, err = unpadPKCS7(dst, aes.BlockSize)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func validMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func unpadPKCS7(src []byte, size int) ([]byte, error) {
	n := src[len(src)-1]
	if len(src)%size != 0 {
		return nil, fmt.Errorf("expected PKCS7 padding for block size %d, but have %d bytes", size, len(src))
	}
	if len(src) <= int(n) {
		return nil, fmt.Errorf("cannot unpad %d bytes out of a total of %d", n, len(src))
	}
	src = src[:len(src)-int(n)]
	return src, nil
}
