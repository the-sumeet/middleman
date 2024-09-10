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

	fmt.Println(certPath)

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Create a certificate template
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{Organization: []string{"Middleman"}},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // Valid for 1 year
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		fmt.Println("Error creating certificate:", err)
		return
	}

	// Save the private key to a .key file
	privFile, err := os.Create(keyPath)
	if err != nil {
		fmt.Println("Error creating private key file:", err)
		return
	}
	defer privFile.Close()

	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	privPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	}
	if err := pem.Encode(privFile, privPem); err != nil {
		fmt.Println("Error encoding private key to PEM:", err)
		return
	}

	// Save the certificate to a .crt file
	certFile, err := os.Create(certPath)
	if err != nil {
		fmt.Println("Error creating certificate file:", err)
		return
	}
	defer certFile.Close()

	certPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}
	if err := pem.Encode(certFile, certPem); err != nil {
		fmt.Println("Error encoding certificate to PEM:", err)
		return
	}

	fmt.Println("Certificate and private key have been saved as 'certificate.crt' and 'private.key'")
}
