/*
appconfig/appconfig_test.go
Tests the known configuration development data
*/
package appconfig

import (
	"testing"
)

func TestLoadAppConfig(t *testing.T) {
	acfg, err := Load("../testdata/config.json")

	if err != nil {
		t.Fatalf("failed to load app config: %v", err)
	}

	if acfg.Addr == "" {
		t.Errorf("expected address to be present")
	}

	if acfg.ConfigDir == "" {
		t.Errorf("expected config directory to be present")
	}

	if acfg.CertsDir == "" {
		t.Errorf("expected certs directory to be present")
	}
}
