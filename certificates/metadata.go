package certificates

import (
	"encoding/json"
	"os"
	"time"
)

type Metadata struct {
	CommonName         string          `json:"commonName"`
	SubjectDN          string          `json:"subjectDN"`
	NotBefore          time.Time       `json:"notBefore"`
	NotAfter           time.Time       `json:"notAfter"`
	Issuer             string          `json:"issuer"`
	IssuerDN           string          `json:"issuerDN"`
	SerialNumber       string          `json:"serialNumber"`
	Thumbprint         string          `json:"thumbprint"`
	Thumbprint256      string          `json:"thumbprint256"`
	IsCA               bool            `json:"isCA"`
	SelfSigned         bool            `json:"selfSigned"`
	SubjectKeyID       string          `json:"subjectKeyID"`
	AuthorityKeyID     string          `json:"authorityKeyID"`
	DNSNames           []string        `json:"dnsNames"`
	IPAddresses        []string        `json:"ipAddresses"`
	EmailAddresses     []string        `json:"emailAddresses"`
	URIs               []string        `json:"uris"`
	KeyUsage           []string        `json:"keyUsage"`
	ExtendedKeyUsage   []string        `json:"extendedKeyUsage"`
	AuthorityInfo      []AuthorityInfo `json:"authorityInfo"`
	SignatureAlgorithm string          `json:"signatureAlgorithm"`
	PublicKeyAlgorithm string         `json:"publicKeyAlgorithm"`
	PublicKeySize      int             `json:"publicKeySize"`
}

type AuthorityInfo struct {
	Method string `json:"method"`
	URI    string `json:"uri"`
}

func LoadMetadata(path string) (Metadata, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Metadata{}, err
	}

	var metadata Metadata

	if err := json.Unmarshal(data, &metadata); err != nil {
		return Metadata{}, err
	}

	return metadata, nil
}
