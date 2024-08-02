package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/WatcherWhale/adguard-job/pkg/adguard"
)

var (
	baseUrl      string = os.Getenv("ADGUARD_URL")
	username     string = os.Getenv("ADGUARD_USER")
	password     string = os.Getenv("ADGUARD_PASSWORD")
	certLocation string = os.Getenv("TLS_CERT_PATH")
	keyLocation  string = os.Getenv("TLS_KEY_PATH")
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	log.Print("Reading secrets...")
	cert, err := os.ReadFile(certLocation)
	if err != nil {
		log.Panicf("Failed to get certificate: %v", err)
	}

	key, err := os.ReadFile(keyLocation)
	if err != nil {
		log.Panicf("Failed to get key: %v", err)
	}

	log.Print("Getting encryption settings...")
	settings, err := adguard.GetCurrentSettings(baseUrl, username, password)
	if err != nil {
		log.Panicf("Failed to get encryption settings: %v", err)
	}

	// Set keys
	settings.CertificateChain = string(cert)
	settings.PrivateKey = string(key)
	settings.PrivateKeySaved = false

	log.Print("Configuring new settings...")
	err = adguard.SetSettings(baseUrl, username, password, settings)
	if err != nil {
		log.Panicf("Failed to configure encryption settings: %v", err)
	}

	log.Print("Done")
}
