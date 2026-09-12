/*
api/handlers.go
Handles server requests
*/
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func (s *Server) downloadHandler(w http.ResponseWriter, r *http.Request) {
	domain := r.PathValue("domain")

	pfxPath, err := s.certificates.GetPFXPath(domain)
	if err != nil {
		http.Error(w, "certificate not found", http.StatusNotFound)
		return
	}

	data, err := os.ReadFile(pfxPath)
	if err != nil {
		http.Error(w, "failed to read certificate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-pkcs12")
	w.Header().Set(
		"Content-Disposition",
		`attachment; filename="`+domain+`.pfx"`,
	)

	w.Write(data)
}
