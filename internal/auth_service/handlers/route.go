package handlers

import "net/http"

func SetupRoutes(h Handler) {
	http.HandleFunc("GET /auth", h.getAuth)
	http.HandleFunc("POST /refresh", h.postRefresh)
}
