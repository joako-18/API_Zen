package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type EncryptService struct {
	secretKey []byte
}

func NewEncryptService(secretKey string) (*EncryptService, error) {
	secretKeyBytes := []byte(secretKey)
	switch len(secretKeyBytes) {
	case 16, 24, 32:
		return &EncryptService{secretKey: secretKeyBytes}, nil
	default:
		return nil, errors.New("invalid key size: must be 16, 24, or 32 bytes")
	}
}

func (e *EncryptService) Encrypt(data string) (string, error) {
	block, err := aes.NewCipher(e.secretKey)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(data), nil)

	encryptedData := base64.StdEncoding.EncodeToString(append(nonce, ciphertext...))

	return encryptedData, nil
}
