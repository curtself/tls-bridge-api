/*
appconfig/apconfig.go
Configuration data for the application, such as listen address and the application directories
*/
package appconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Addr      string `json:"addr"`
	ConfigDir string `json:"config_dir"`
	CertsDir  string `json:"certs_dir"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read application config: %w", err)
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse application config: %w", err)
	}

	return cfg, nil
}
