package crypto

import (
	"crypto/sha256"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/pbkdf2"
)

type KDF uint8

const (
	KDFTypePBKDF2   KDF = 0
	KDFTypeArgon2id KDF = 1
)

type KDFParams struct {
	Type        KDF
	Iterations  int
	Memory      int
	Parallelism int
}

func CalculateUserKey(
	password, email string,
	kdfType uint8,
	kdfIter, kdfMemory, kdfParallel int,
) ([]byte, error) {
	salt := []byte(strings.ToLower(email))
	passBytes := []byte(password)
	switch KDF(kdfType) {
	case KDFTypePBKDF2:
		return pbkdf2.Key(
			passBytes,
			salt,
			kdfIter,
			32,
			sha256.New,
		), nil
	case KDFTypeArgon2id:
		var argonSalt [32]byte = sha256.Sum256(salt)
		return argon2.IDKey(
			passBytes,
			argonSalt[:],
			uint32(kdfIter),
			uint32(kdfMemory*1024),
			uint8(kdfParallel),
			32,
		), nil
	default:
		return nil, fmt.Errorf("unsupported KDF type %d", kdfType)
	}
}

func DecryptMasterKey(masterKeyContent []byte, userKey []byte) ([]byte, []byte, error) {
	masterKeyCipher := CipherFromBytes(masterKeyContent)
	var key []byte
	var mac []byte
	var err error
	switch masterKeyCipher.Type {
	case AesCbc256_B64:
		fmt.Println(1)
		key, err = Decrypt(masterKeyCipher, userKey, nil)
		if err != nil {
			return key, mac, err
		}
	case AesCbc256_HmacSha256_B64:
		keyBytes, keyMacBytes := stretchKey(userKey)
		key, err = Decrypt(masterKeyCipher, keyBytes, keyMacBytes)
		if err != nil {
			return key, mac, err
		}
	default:
		return key, mac, fmt.Errorf("unsupported key cipher type %q", masterKeyCipher.Type)
	}

	if len(key) == 64 {
		key, mac = key[:32], key[32:64]
	} else if len(key) != 32 {
		return key, mac, fmt.Errorf("invalid key length: %d", len(key))
	}
	return key, mac, nil
}

func stretchKey(orig []byte) (key, keyMac []byte) {
	key = make([]byte, 32)
	keyMac = make([]byte, 32)
	var r io.Reader
	r = hkdf.Expand(sha256.New, orig, []byte("enc"))
	r.Read(key)
	r = hkdf.Expand(sha256.New, orig, []byte("mac"))
	r.Read(keyMac)
	return key, keyMac
}
