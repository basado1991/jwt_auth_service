package handlers

import (
	"context"

	jwtdecoder "github.com/basado1991/jwt_auth_service/internal/jwt_decoder"
	jwtencoder "github.com/basado1991/jwt_auth_service/internal/jwt_encoder"
	"github.com/basado1991/jwt_auth_service/internal/mailer"
	"github.com/basado1991/jwt_auth_service/internal/storage"
)

type Handler struct {
	Ctx        context.Context
	JwtEncoder jwtencoder.JwtEncoder
	JwtDecoder jwtdecoder.JwtDecoder

	Mailer mailer.Mailer

	Storage storage.Storage
}
