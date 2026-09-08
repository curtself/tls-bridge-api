package main

import (
	"log"

	"tls-bridge-api/api"
	"tls-bridge-api/certificates"
	"tls-bridge-api/config"
)

func main() {
	cfg := config.New(
		"./testdata/config",
		"./testdata/certs",
	)

	certService := certificates.NewService(cfg)
	server := api.NewServer(cfg, certService)

	log.Println("tls-bridge-api listening on :8080")
	log.Fatal(server.ListenAndServe())
}
