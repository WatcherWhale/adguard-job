package adguard

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type EncryptionSettings struct {
	Enabled          bool   `json:"enabled"`
	ServerName       string `json:"server_name"`
	ForceHTTPS       bool   `json:"force_https"`
	PortHTTPS        int    `json:"port_https"`
	PortDNSOverTLS   int    `json:"port_dns_over_tls"`
	PortDNSOverQuic  int    `json:"port_dns_over_quic"`
	CertificateChain string `json:"certificate_chain"`
	PrivateKey       string `json:"private_key"`
	CertificatePath  string `json:"certificate_path"`
	PrivateKeyPath   string `json:"private_key_path"`
	ServePlainDNS    bool   `json:"serve_plain_dns"`
	PrivateKeySaved  bool   `json:"private_key_saved"`
}

func GetCurrentSettings(baseUrl string, username string, password string) (EncryptionSettings, error) {
	statusUrl, err := url.JoinPath(baseUrl, "/control/tls/status")
	if err != nil {
		return EncryptionSettings{}, err
	}

	req, err := http.NewRequest(http.MethodGet, statusUrl, nil)
	if err != nil {
		return EncryptionSettings{}, err
	}

	req.SetBasicAuth(username, password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return EncryptionSettings{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return EncryptionSettings{}, err
	}

	var settings EncryptionSettings
	err = json.Unmarshal(body, &settings)
	if err != nil {
		return EncryptionSettings{}, err
	}

	return settings, nil
}

func SetSettings(baseUrl, username, password string, settings EncryptionSettings) error {
	statusUrl, err := url.JoinPath(baseUrl, "/control/tls/configure")
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(&settings)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(jsonBytes)
	req, err := http.NewRequest(http.MethodGet, statusUrl, buf)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, password)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
