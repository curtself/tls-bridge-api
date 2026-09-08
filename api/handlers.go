package api

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func (s *Server) metadataHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.PathValue("domain")

	metadata, err := s.certificates.GetMetadata(domain)
	if err != nil {
		http.Error(w, "certificate metadata not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(metadata); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
