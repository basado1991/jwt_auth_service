package handlers

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/basado1991/jwt_auth_service/internal/auth_service/utils"
	"golang.org/x/crypto/bcrypt"
)

type GetAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h Handler) getAuth(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(h.Ctx, REQUEST_TIMEOUT)
	defer cancel()

	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteBadRequest(w, errors.New("id not provided"))
		return
	}

	user, err := h.Storage.GetUserById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteBadRequest(w, errors.New("user not exist"))
		} else {
			utils.WriteInternalError(w)
			log.Println(err)
		}
		return
	}

	refreshToken := make([]byte, 64)
	rand.Read(refreshToken)

	refreshTokenMarshalled := base64.RawURLEncoding.EncodeToString(refreshToken)

	refreshTokenHashed, err := bcrypt.GenerateFromPassword(refreshToken, bcrypt.DefaultCost)
	if err != nil {
		utils.WriteInternalError(w)
		log.Println(err)
		return
	}

	payload := map[string]any{
		"id":   id,
		"refresh_hash": string(refreshTokenHashed),
		"ip":   r.RemoteAddr,
		"exp":  time.Now().Add(TOKEN_EXPIRATION_TIME).Unix(),
	}
	accessToken, err := h.JwtEncoder.Encode(payload)
	if err != nil {
		utils.WriteInternalError(w)
		log.Println(err)
		return
	}

	user.RefreshToken = refreshTokenHashed

	if err := h.Storage.UpdateUser(ctx, user); err != nil {
		utils.WriteInternalError(w)
		log.Println(err)
		return
	}

	err = utils.WriteJsonOk(w, GetAuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenMarshalled,
	})
	if err != nil {
		log.Println(err)
		return
	}
}
