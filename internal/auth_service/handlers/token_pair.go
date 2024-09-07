package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h Handler) makeTokenPair(id, remoteAddr string) (*TokenPair, []byte, error) {
	refreshToken := make([]byte, 72)
	rand.Read(refreshToken)

	refreshTokenMarshalled := base64.RawURLEncoding.EncodeToString(refreshToken)

	refreshTokenHashed, err := bcrypt.GenerateFromPassword(refreshToken, bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	payload := map[string]any{
		"id":           id,
		"refresh_hash": string(refreshTokenHashed),
		"ip":           remoteAddr,
		"exp":          time.Now().Add(TOKEN_EXPIRATION_TIME).Unix(),
	}
	accessToken, err := h.JwtEncoder.Encode(payload)
	if err != nil {
		return nil, nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenMarshalled,
	}, refreshTokenHashed, nil
}
