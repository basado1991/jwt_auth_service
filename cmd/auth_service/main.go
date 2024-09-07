package main

import (
	"context"
	"log"
	"os"

	authservice "github.com/basado1991/jwt_auth_service/internal/auth_service"
	"github.com/basado1991/jwt_auth_service/internal/auth_service/handlers"
	jwtdecoder "github.com/basado1991/jwt_auth_service/internal/jwt_decoder"
	jwtencoder "github.com/basado1991/jwt_auth_service/internal/jwt_encoder"
	"github.com/basado1991/jwt_auth_service/internal/mailer"
	"github.com/basado1991/jwt_auth_service/internal/storage"
)

func getEnv(name string) string {
	data, exists := os.LookupEnv(name)
	if !exists {
		log.Fatalln(name, "is not set")
	}

	return data
}

func main() {
	keyPath := getEnv("AS_KEY_PATH")
	key, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatalln(err)
	}

	addr := getEnv("AS_ADDR")
	postgresAddr := getEnv("AS_POSTGRES")

	sha512Signer := jwtencoder.NewJwtHS256Signer(key)
	jwtEncoder := jwtencoder.NewJwtEncoder(*sha512Signer)

	sha512Verifier := jwtdecoder.NewJwtHS256Verifier(key)
	jwtDecoder := jwtdecoder.NewJwtDecoder(*sha512Verifier)

	storage, err := storage.NewPostgresStorage(postgresAddr)
	if err != nil {
		log.Fatalln(err)
	}

	mailHost := getEnv("AS_MAIL_HOST")
	mailAddr := getEnv("AS_MAIL_ADDR")
	mailUsername := getEnv("AS_MAIL_USERNAME")
	mailPassword := getEnv("AS_MAIL_PASSWORD")
	mailFrom := getEnv("AS_MAIL_FROM")

	mailer := mailer.NewMailer(mailer.MailerOpts{
		Host: mailHost,
		Addr: mailAddr,

		User:     mailUsername,
		Password: mailPassword,

		From: mailFrom,
	})

	hs := handlers.Handler{
		JwtEncoder: *jwtEncoder,
		JwtDecoder: *jwtDecoder,
		Mailer:     *mailer,
		Storage:    storage,
		Ctx:        context.Background(),
	}

	authservice.Init(hs)

	log.Println("server started")
	if err := authservice.Serve(addr); err != nil {
		log.Fatalln(err)
	}
}
