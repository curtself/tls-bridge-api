package api

import (
	"net/http"
	"net/http/httptest"
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
	handler.ServeHTTP(response,request)

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
	handler.ServeHTTP(response,request)

	if response.Code != http.StatusUnauthorized {
		t.Errorf( "expected status %d, got %d", http.StatusUnauthorized, response.Code, )
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
	handler.ServeHTTP(response,request)

	if response.Code != http.StatusUnauthorized {
		t.Errorf( "expected status %d, got %d", http.StatusUnauthorized, response.Code, )
	}
}
