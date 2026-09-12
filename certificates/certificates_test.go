/*
certificates/certificates_test.go
Tests the known development configuration
*/
package certificates

import (
	"testing"

	"tls-bridge-api/config"
)

func TestGetMetadata(t *testing.T) {
	cfg := config.New(
		"../testdata/config",
		"../testdata/certs",
	)

	service := NewService(cfg)

	md, err := service.GetMetadata("prod-content-web.sdccd.edu")
	if err != nil {
		t.Fatalf("failed to get metadata: %v", err)
	}

	if md.CommonName != "prod-content-web.sdccd.edu" {
		t.Errorf("expected common name prod-content-web.sdccd.edu, got %s", md.CommonName)
	}
}

func TestGetMetadataMissingDomain(t *testing.T) {
	cfg := config.New(
		"../testdata/config",
		"../testdata/certs",
	)

	service := NewService(cfg)

	_, err := service.GetMetadata("does-not-exist.example.com")
	if err == nil {
		t.Error("expected error for missing domain")
	}
}

func TestGetPFXPath(t *testing.T) {
	cfg := config.New(
		"../testdata/config",
		"../testdata/certs",
	)

	service := NewService(cfg)

	path, err := service.GetPFXPath("prod-content-web.sdccd.edu")
	if err != nil {
		t.Fatalf("failed to get PFX path: %v", err)
	}

	expected := "../testdata/certs/prod-content-web.sdccd.edu/prod-content-web.sdccd.edu.pfx"

	if path != expected {
		t.Errorf("expected path %s, got %s", expected, path)
	}
}

func TestGetPFXPathMissingDomain(t *testing.T) {
	cfg := config.New(
		"../testdata/config",
		"../testdata/certs",
	)

	service := NewService(cfg)

	_, err := service.GetPFXPath("does-not-exist.example.com")
	if err == nil {
		t.Error("expected error for missing PFX")
	}
}
