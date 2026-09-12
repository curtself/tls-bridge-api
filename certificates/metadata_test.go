/*
certificates/metadata_test.go
Tests metadata loading against a known development configuration
*/
package certificates

import (
	"testing"
)

func TestLoadMetadata(t *testing.T) {
	md, err := LoadMetadata("../testdata/certs/prod-content-web.sdccd.edu/prod-content-web.sdccd.edu.json")

	if err != nil {
		t.Fatalf("failed to load metadata: %v", err)
	}

	if md.CommonName != "prod-content-web.sdccd.edu" {
		t.Errorf("expected common name prod-content-web.sdccd.edu, got %s", md.CommonName)
	}
	if md.NotAfter.IsZero() {
		t.Error("expected NotAfter to be populated")
	}
}
