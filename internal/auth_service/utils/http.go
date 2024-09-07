package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/basado1991/jwt_auth_service/internal/auth_service/dto"
	jwtdecoder "github.com/basado1991/jwt_auth_service/internal/jwt_decoder"
)

var (
	ErrEmptyAuthorization = errors.New("empty authorization header")
	ErrInvalidBearer      = errors.New("invalid bearer token")
	ErrInternalServer     = errors.New("internal server error")
)

func WriteJsonOk(w http.ResponseWriter, data any) error {
	return WriteJson(w, http.StatusOK, data)
}

func WriteError(w http.ResponseWriter, statusCode int, payload error) error {
	return WriteJson(w, statusCode, dto.HttpError{Code: payload.Error()})
}

func WriteBadRequest(w http.ResponseWriter, payload error) error {
	return WriteError(w, http.StatusBadRequest, payload)
}

func WriteInternalError(w http.ResponseWriter) error {
	return WriteError(w, http.StatusInternalServerError, ErrInternalServer)
}

func WriteJson(w http.ResponseWriter, statusCode int, data any) error {
	marshalled, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(marshalled)
	if err != nil {
		return err
	}

	return nil
}

func ReadJson(dst dto.Verifiable, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dst)
	if err != nil {
		return err
	}

	if err := dst.Verify(); err != nil {
		return err
	}

	return nil
}

func ReadJwt(r *http.Request, jwtDecoder jwtdecoder.JwtDecoder) (map[string]any, error) {
	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		return nil, ErrEmptyAuthorization
	}
	chunks := strings.Split(bearer, " ")
	if len(chunks) < 2 {
		return nil, ErrInvalidBearer
	}
	if chunks[0] != "Bearer" {
		return nil, ErrInvalidBearer
	}

	payload, err := jwtDecoder.Decode(chunks[1])
	if err != nil {
		return nil, err
	}

	return payload, nil
}
