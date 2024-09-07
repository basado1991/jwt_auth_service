package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/basado1991/jwt_auth_service/internal/auth_service/utils"
)

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

	tokenPair, refreshTokenHashed, err := h.makeTokenPair(id, r.RemoteAddr)
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

	err = utils.WriteJsonOk(w, tokenPair)
	if err != nil {
		log.Println(err)
		return
	}
}
