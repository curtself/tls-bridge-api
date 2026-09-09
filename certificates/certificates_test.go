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
