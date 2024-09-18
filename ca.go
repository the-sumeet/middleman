package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func genCert() {
	certPath, keyPath := getCertKeyPath()
	// Generate a private key
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Create a template for the certificate
	certTemplate := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "Middleman", OrganizationalUnit: []string{"Middleman"}},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // valid for 1 year
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		IsCA:                  true,
		BasicConstraintsValid: true,
	}

	// Create the self-signed certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, &privKey.PublicKey, privKey)
	if err != nil {
		fmt.Println("Error creating certificate:", err)
		return
	}

	// Save the private key
	keyFile, err := os.Create(keyPath)
	if err != nil {
		fmt.Println("Error creating key file:", err)
		return
	}
	defer keyFile.Close()

	err = pem.Encode(keyFile, &pem.Block{Type: "PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privKey)})
	if err != nil {
		fmt.Println("Error writing private key to file:", err)
		return
	}

	// Save the certificate
	certFile, err := os.Create(certPath)
	if err != nil {
		fmt.Println("Error creating cert file:", err)
		return
	}
	defer certFile.Close()

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	if err != nil {
		fmt.Println("Error writing certificate to file:", err)
		return
	}

	fmt.Println("Certificate and key generated successfully.")
}
