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
	apiServer := &Server{
		certificates: certService,
		config:       cfg,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", healthHandler)
	mux.Handle(
		"GET /api/certificates/{domain}/metadata",
		authMiddleware(cfg, http.HandlerFunc(apiServer.metadataHandler)),
	)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
