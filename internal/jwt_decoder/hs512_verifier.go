package jwtdecoder

import (
	"crypto/hmac"
	"crypto/sha512"
)

type JwtHS512Verifier struct {
	key []byte
}

func NewJwtHS256Verifier(key []byte) *JwtHS512Verifier {
	return &JwtHS512Verifier{key: key}
}

func (v JwtHS512Verifier) Verify(data []byte, signature []byte) (bool, error) {
	mac := hmac.New(sha512.New, v.key)
	if _, err := mac.Write(data); err != nil {
		return false, err
	}
	expectedMac := mac.Sum(nil)
	return hmac.Equal(expectedMac, signature), nil
}
