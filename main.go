package main

import (
	"flag"
	"log"

	"tls-bridge-api/api"
	"tls-bridge-api/appconfig"
	"tls-bridge-api/certificates"
	"tls-bridge-api/config"
	"tls-bridge-api/version"
)

const defaultConfigPath = "/etc/tls-bridge-api/config.json"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to application configuration file")
	flag.Parse()

	appCfg, err := appconfig.Load(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New(
		appCfg.ConfigDir,
		appCfg.CertsDir,
	)

	certService := certificates.NewService(cfg)
	server := api.NewServer(cfg, certService, appCfg)

	log.Printf("tls-bridge-api listening on %s\n", server.Addr)
	log.Printf("version: %s %s %s", version.Version, version.Commit, version.BuildDate)
	log.Fatal(server.ListenAndServe())
}
