package config

import (
	"testing"
)

func TestLoadCertificateConfig(t *testing.T) {
	cfg, err := LoadCertificateConfig("../testdata/config/prod-content-web.sdccd.edu.json")
	if err != nil {
		t.Fatalf("failed to load certificate config: %v", err)
	}

	if cfg.Domain != "prod-content-web.sdccd.edu" {
		t.Errorf("expected domain prod-content-web.sdccd.edu, got %s", cfg.Domain)
	}

	if cfg.AuthToken == "" {
		t.Error("expected auth token to be present")
	}
}
