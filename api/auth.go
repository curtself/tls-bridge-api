/*
api/auth.go
Authorization middleware for protected endpoints
*/
package api

import (
	"crypto/subtle"
	"net/http"
	"path/filepath"
	"strings"

	"tls-bridge-api/config"
)

func authMiddleware(cfg config.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		parts := strings.Fields(authorization)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		domain := r.PathValue("domain")

		certificateConfig, err := config.LoadCertificateConfig(
			filepath.Join(cfg.ConfigDir, domain+".json"),
		)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		if subtle.ConstantTimeCompare(
			[]byte(token),
			[]byte(certificateConfig.AuthToken),
		) != 1 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
