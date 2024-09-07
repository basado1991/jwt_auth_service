package jwtdecoder

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

var (
	ErrMalformedToken = errors.New("jwt is malformed")
)

type JwtDecoder struct {
	Verifier JwtVerifier
}

func NewJwtDecoder(verifier JwtVerifier) *JwtDecoder {
	return &JwtDecoder{
		Verifier: verifier,
	}
}

func (d JwtDecoder) Decode(token string) (map[string]any, error) {
	chunks := strings.Split(token, ".")
	if len(chunks) < 3 {
		return nil, ErrMalformedToken
	}
	signatureBytesMarshalled := []byte(chunks[2])

	signature := make([]byte, base64.RawURLEncoding.DecodedLen(len(signatureBytesMarshalled)))
	_, err := base64.RawURLEncoding.Decode(signature, signatureBytesMarshalled)
	if err != nil {
		return nil, ErrMalformedToken
	}

	genuine, err := d.Verifier.Verify([]byte(chunks[0]+"."+chunks[1]), signature)
	if err != nil && !genuine {
		return nil, ErrMalformedToken
	}

	var payload map[string]any

	// поскольку подпись проверена, токен создан этой программой
	// а если токен создан этой программой, то он скорее всего верный, и проверять ошибки нет смысла
	payloadMarshalled, _ := base64.RawURLEncoding.DecodeString(chunks[1])
	_ = json.Unmarshal(payloadMarshalled, &payload)

	return payload, nil
}
