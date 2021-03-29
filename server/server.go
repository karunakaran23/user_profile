package server

import (
	"net/http"
	"user_profile/handler"
	"user_profile/middleware"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) InitRoute(h *handler.Handler) {
	s.Router.HandleFunc("/user/profile", middleware.JSONandCORS(h.Userprofile))
}
