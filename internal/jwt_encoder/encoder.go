package jwtencoder

import (
	"encoding/base64"
	"encoding/json"

	"github.com/basado1991/jwt_auth_service/internal/types"
)

type JwtEncoder struct {
  Signer JwtSigner
}

func NewJwtEncoder(signer JwtHS512Signer) *JwtEncoder {
  return &JwtEncoder{Signer: signer}
}

func (e JwtEncoder) Encode(payload map[string]any) (string, error) {
  header := types.JwtHeader {
    Type: types.JWT_TYPE,
    Algorithm: e.Signer.GetAlgorithm(),
  }
  headerMarshalled, err := json.Marshal(header)
  if err != nil {
    return "", err
  }
  payloadMarshalled, err := json.Marshal(payload)

  headerEncoded := base64.RawURLEncoding.EncodeToString(headerMarshalled)
  payloadEncoded := base64.RawURLEncoding.EncodeToString(payloadMarshalled)

  unsignedToken := headerEncoded + "." + payloadEncoded

  signature, err := e.Signer.Sign([]byte(unsignedToken))
  if err != nil {
    return "", err
  }

  token := unsignedToken + "." + base64.RawURLEncoding.EncodeToString(signature)

  return token, nil
}
