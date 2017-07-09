package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func BuildHTTPHandler(r *chi.Mux) http.Handler {
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	return r
}

func SysRoute(r *chi.Mux, sh *SysHandler) {
	r.Get("/sys/info/ping", sh.Ping)
	r.Get("/sys/info/health", sh.Health)
}
