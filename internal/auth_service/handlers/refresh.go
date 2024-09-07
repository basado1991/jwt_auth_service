package handlers

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/basado1991/jwt_auth_service/internal/auth_service/utils"
	"golang.org/x/crypto/bcrypt"
)

type PostRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (r *PostRefreshRequest) Verify() error {
	if r.RefreshToken == "" {
		return errors.New("empty refresh token")
	}

	return nil
}

func (h Handler) postRefresh(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(h.Ctx, REQUEST_TIMEOUT)
	defer cancel()

	token, err := utils.ReadJwt(r, h.JwtDecoder)
	if err != nil {
		utils.WriteBadRequest(w, err)

		return
	}

	var data PostRefreshRequest
	if err := utils.ReadJson(&data, r); err != nil {
		utils.WriteBadRequest(w, err)

		return
	}

	user, err := h.Storage.GetUserById(ctx, token["id"].(string))
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteBadRequest(w, errors.New("user not exist"))
		} else {
			utils.WriteInternalError(w)
			log.Println(err)
		}
		return
	}

	givenRefreshToken, err := base64.RawURLEncoding.DecodeString(data.RefreshToken)
	if err != nil {
		utils.WriteBadRequest(w, errors.New("invalid refresh token"))
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.RefreshToken, givenRefreshToken); err != nil {
		utils.WriteBadRequest(w, errors.New("bad refresh token"))
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(token["refresh_hash"].(string)), givenRefreshToken); err != nil {
		utils.WriteBadRequest(w, errors.New("refresh token does not belong to given access token"))
		return
	}

	currentAddr := token["ip"].(string)

	tokenPair, refreshTokenHashed, err := h.makeTokenPair(token["id"].(string), r.RemoteAddr)
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

	if currentAddr != r.RemoteAddr {
		err := h.Mailer.Send(ctx, user.Email, NOTIFICATION_MESSAGE_SUBJECT, fmt.Sprintf(NOTIFICATION_MESSAGE_BODY_TEMPLATE, user.Name, r.RemoteAddr))
		if err != nil {
			log.Println(err)
		}
	}

	err = utils.WriteJsonOk(w, tokenPair)
	if err != nil {
		log.Println(err)
		return
	}
}
