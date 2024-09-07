package jwtencoder

import (
	"crypto/hmac"
	"crypto/sha512"
)

type JwtHS512Signer struct {
	key []byte
}

func NewJwtHS256Signer(key []byte) *JwtHS512Signer {
	return &JwtHS512Signer{key: key}
}

func (s JwtHS512Signer) Sign(data []byte) ([]byte, error) {
	mac := hmac.New(sha512.New, s.key)
	_, err := mac.Write(data)
	if err != nil {
		return nil, err
	}

	return mac.Sum(nil), nil
}

func (s JwtHS512Signer) GetAlgorithm() string {
	return "HS512"
}
