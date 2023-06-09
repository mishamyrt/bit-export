package crypto

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
)

type CipherStringType int

var b64 = base64.StdEncoding.Strict()

// Taken from https://github.com/bitwarden/jslib/blob/f30d6f8027055507abfdefd1eeb5d9aab25cc601/src/enums/encryptionType.ts
const (
	AesCbc256_B64                     CipherStringType = 0
	AesCbc128_HmacSha256_B64          CipherStringType = 1
	AesCbc256_HmacSha256_B64          CipherStringType = 2
	Rsa2048_OaepSha256_B64            CipherStringType = 3
	Rsa2048_OaepSha1_B64              CipherStringType = 4
	Rsa2048_OaepSha256_HmacSha256_B64 CipherStringType = 5
	Rsa2048_OaepSha1_HmacSha256_B64   CipherStringType = 6
)

type CipherString struct {
	Type CipherStringType

	// Initialization vector
	IV []byte
	// Cipher text
	CT []byte
	// Message authentication code
	MAC []byte
}

func (t CipherStringType) HasMAC() bool {
	return t != AesCbc256_B64
}

func (s CipherString) IsZero() bool {
	return s.Type == 0 && s.IV == nil && s.CT == nil && s.MAC == nil
}

func (s CipherString) String() string {
	if s.IsZero() {
		return ""
	}
	if !s.Type.HasMAC() {
		return fmt.Sprintf("%d.%s|%s",
			s.Type,
			b64.EncodeToString(s.IV),
			b64.EncodeToString(s.CT),
		)
	}
	return fmt.Sprintf("%d.%s|%s|%s",
		s.Type,
		b64.EncodeToString(s.IV),
		b64.EncodeToString(s.CT),
		b64.EncodeToString(s.MAC),
	)
}

func (s CipherString) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *CipherString) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	i := bytes.IndexByte(data, '.')
	if i < 0 {
		return fmt.Errorf("cipher string does not contain a type: %q", data)
	}
	typStr := string(data[:i])
	var err error
	if t, err := strconv.Atoi(typStr); err != nil {
		return fmt.Errorf("invalid cipher string type: %q", typStr)
	} else {
		s.Type = CipherStringType(t)
	}
	switch s.Type {
	case AesCbc128_HmacSha256_B64, AesCbc256_HmacSha256_B64, AesCbc256_B64:
	default:
		return fmt.Errorf("unsupported cipher string type: %d", s.Type)
	}

	data = data[i+1:]
	parts := bytes.Split(data, []byte("|"))
	wantParts := 3
	if !s.Type.HasMAC() {
		wantParts = 2
	}
	if len(parts) != wantParts {
		return fmt.Errorf("cipher string type requires %d parts: %q", wantParts, data)
	}

	// TODO: do a single []byte allocation for all fields
	if s.IV, err = b64decode(parts[0]); err != nil {
		return err
	}
	if s.CT, err = b64decode(parts[1]); err != nil {
		return err
	}
	if s.Type.HasMAC() {
		if s.MAC, err = b64decode(parts[2]); err != nil {
			return err
		}
	}
	return nil
}

func b64decode(src []byte) ([]byte, error) {
	dst := make([]byte, b64.DecodedLen(len(src)))
	n, err := b64.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	dst = dst[:n]
	return dst, nil
}

func CipherFromBytes(b []byte) CipherString {
	var cipher CipherString
	cipher.UnmarshalText(b)
	return cipher
}

func CipherFromString(s string) CipherString {
	return CipherFromBytes([]byte(s))
}
