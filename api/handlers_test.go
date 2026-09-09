package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"tls-bridge-api/certificates"
	"tls-bridge-api/config"
)

func testHandler(t *testing.T) http.Handler {
	t.Helper()

	cfg := config.New(
		"../testdata/config",
		"../testdata/certs",
	)

	certService := certificates.NewService(cfg)

	server := &Server{
		certificates: certService,
		config:       cfg,
	}

	return server.routes()
}

func TestMetadataAnonymousAccess(t *testing.T) {
	handler := testHandler(t)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/certificates/prod-content-web.sdccd.edu/metadata",
		nil,
	)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusUnauthorized,
			response.Code,
		)
	}
}

func TestMetadataAuthorizedAccess(t *testing.T) {
	handler := testHandler(t)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/certificates/prod-content-web.sdccd.edu/metadata",
		nil,
	)

	request.Header.Set(
		"Authorization",
		"Bearer ff49ca3ea21edf591ff796f60dccba666a50992a0479cee78adc528beb75a228",
	)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusOK,
			response.Code,
		)
	}
}

func TestMetadataBadTokenAccess(t *testing.T) {
	handler := testHandler(t)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/certificates/prod-content-web.sdccd.edu/metadata",
		nil,
	)

	request.Header.Set(
		"Authorization",
		"Bearer abcdefghijklmnopqrstuvwzyz1234567890",
	)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, response.Code)
	}
}

func TestMetadataWrongDomainAccess(t *testing.T) {
	handler := testHandler(t)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/certificates/admin.sdccd.edu/metadata",
		nil,
	)

	request.Header.Set(
		"Authorization",
		"Bearer ff49ca3ea21edf591ff796f60dccba666a50992a0479cee78adc528beb75a228",
	)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, response.Code)
	}
}

func TestDownloadAuthorizedAccess(t *testing.T) {
	handler := testHandler(t)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/certificates/prod-content-web.sdccd.edu/download",
		nil,
	)

	request.Header.Set(
		"Authorization",
		"Bearer ff49ca3ea21edf591ff796f60dccba666a50992a0479cee78adc528beb75a228",
	)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if response.Header().Get("Content-Type") != "application/x-pkcs12" {
		t.Errorf("expected PKCS#12 content type, got %s", response.Header().Get("Content-Type"))
	}

	expectedDisposition := `attachment; filename="prod-content-web.sdccd.edu.pfx"`

	if response.Header().Get("Content-Disposition") != expectedDisposition {
		t.Errorf(
			"expected Content-Disposition %q, got %q",
			expectedDisposition,
			response.Header().Get("Content-Disposition"),
		)
	}
	expectedPFX, err := os.ReadFile(
		"../testdata/certs/prod-content-web.sdccd.edu/prod-content-web.sdccd.edu.pfx",
	)
	if err != nil {
		t.Fatalf("failed to read expected PFX: %v", err)
	}

	if !bytes.Equal(response.Body.Bytes(), expectedPFX) {
		t.Error("response body does not match expected PFX")
	}

}

func TestDownloadAnonymousAccess(t *testing.T) {
	handler := testHandler(t)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/certificates/prod-content-web.sdccd.edu/download",
		nil,
	)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusUnauthorized,
			response.Code,
		)
	}
}

func TestDownloadBadTokenAccess(t *testing.T) {
	handler := testHandler(t)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/certificates/prod-content-web.sdccd.edu/download",
		nil,
	)

	request.Header.Set(
		"Authorization",
		"Bearer abcdefghijklmnopqrstuvwzyz1234567890",
	)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusUnauthorized,
			response.Code,
		)
	}
}

func TestDownloadMissingCertificate(t *testing.T) {
	handler := testHandler(t)

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/certificates/admin.sdccd.edu/download",
		nil,
	)

	request.Header.Set(
		"Authorization",
		"Bearer 8dd5eb5226b0b45e28933e3c7d6bb6d6da317349bbcaea0fa3e4f539a5d42554",
	)

	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusNotFound,
			response.Code,
		)
	}
}
