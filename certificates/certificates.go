package certificates

import (
	"fmt"
	"path/filepath"

	"tls-bridge-api/config"
)

type Service struct {
	config config.Config
}

func NewService(cfg config.Config) *Service {
	return &Service{
		config: cfg,
	}
}

func (s *Service) GetMetadata(domain string) (Metadata, error) {
	metadataPath := filepath.Join(
		s.config.CertsDir,
		domain,
		domain+".json",
	)

	return LoadMetadata(metadataPath)
}

func (s *Service) GetPFXPath(domain string) (string, error) {
	pfxPath := filepath.Join(
		s.config.CertsDir,
		domain,
		domain+".pfx",
	)

	return pfxPath, nil
}

func (s *Service) ValidateDomain(domain string) error {
	if domain == "" {
		return fmt.Errorf("domain cannot be empty")
	}

	return nil
}
