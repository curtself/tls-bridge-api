package api

import (
	"net/http"

	"tls-bridge-api/certificates"
	"tls-bridge-api/config"
)

type Server struct {
	certificates *certificates.Service
	config       config.Config
}

func NewServer(cfg config.Config, certService *certificates.Service) *http.Server {
	server := &Server{
		certificates: certService,
		config:       cfg,
	}

	return &http.Server{
		Addr:    ":8080",
		Handler: server.routes(),
	}
}

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", healthHandler)

	mux.Handle(
		"GET /api/certificates/{domain}/metadata",
		authMiddleware(
			s.config,
			http.HandlerFunc(s.metadataHandler),
		),
	)
	mux.Handle(
		"GET /api/certificates/{domain}/download",
		authMiddleware(
			s.config,
			http.HandlerFunc(s.downloadHandler),
		),
	)

	return mux
}
