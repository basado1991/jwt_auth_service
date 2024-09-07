package authservice

import (
	"net/http"

	"github.com/basado1991/jwt_auth_service/internal/auth_service/handlers"
)

func Init(h handlers.Handler) {
  handlers.SetupRoutes(h)
}

func Serve(addr string) error {
  return http.ListenAndServe(addr, nil)
}
