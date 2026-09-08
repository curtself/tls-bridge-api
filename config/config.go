package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ConfigDir string
	CertsDir  string
}

type CertificateConfig struct {
	Domain       string `json:"domain"`
	CertDir      string `json:"cert_dir"`
	CertFile     string `json:"certfile"`
	KeyFile      string `json:"keyfile"`
	DeployMethod string `json:"deploy_method"`
	PFXPassword  string `json:"pfx_password"`
	AuthToken    string `json:"auth_token"`
	ServerName   string `json:"server_name"`
}

func New(configDir string, certsDir string) Config {
	return Config{
		ConfigDir: configDir,
		CertsDir:  certsDir,
	}
}

func LoadCertificateConfig(path string) (CertificateConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return CertificateConfig{}, err
	}

	var cfg CertificateConfig

	if err := json.Unmarshal(data, &cfg); err != nil {
		return CertificateConfig{}, err
	}

	return cfg, nil
}
