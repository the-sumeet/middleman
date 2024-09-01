package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AdguardTeam/gomitmproxy"
	"github.com/AdguardTeam/gomitmproxy/mitm"
	"github.com/AdguardTeam/gomitmproxy/proxyutil"
)

// func addCertificateToKeychain() error {

// 	// certPath := "/Users/sumeetmathpati/Projects/sumeet/middleman/testproxy/demo.crt" // Replace with your certificate path
// 	// keychainPath := "$HOME/Library/Keychains/login.keychain"                         // or ~/Library/Keychains/login.keychain

// 	commands := []string{
// 		"osascript", "-e", `do shell script "security add-trusted-cert -r trustRoot -k $HOME/Library/Keychains/login.keychain "/Users/sumeetmathpati/Projects/sumeet/middleman/testproxy/demo.crt" with prompt "Middleman wants to add SSL certificate to keychain."`,
// 	}

// 	cmd := exec.Command(commands[0], commands[1:]...)
// 	err := cmd.Run()
// 	fmt.Println("Error: ", err)
// 	if err != nil {
// 		return fmt.Errorf("failed to add certificate: %w", err)
// 	}
// 	return nil
// }

const (
	certOrganization = "Middleman"
)

type Proxy struct {
	certPath  string
	certKey   string
	proxy     *gomitmproxy.Proxy
	proxyHost string
	proxyPort string
}

func NewProxy(certPath, certKey string) *Proxy {

	proxy := Proxy{
		proxyHost: "127.0.0.1",
		proxyPort: "3333",
	}

	tlsCert, err := tls.LoadX509KeyPair(certPath, certKey)
	if err != nil {
		log.Fatal(err)
	}

	privateKey := tlsCert.PrivateKey.(*rsa.PrivateKey)

	x509c, err := x509.ParseCertificate(tlsCert.Certificate[0])
	if err != nil {
		log.Fatal(err)
	}

	mitmConfig, err := mitm.NewConfig(x509c, privateKey, nil)
	if err != nil {
		log.Fatal(err)
	}

	mitmConfig.SetValidity(time.Hour * 24 * 7)   // generate certs valid for 7 days
	mitmConfig.SetOrganization(certOrganization) // cert organization

	intPort, err := strconv.Atoi(proxy.proxyPort)
	if err != nil {
		log.Fatal(err)
	}

	gomitmproxy := gomitmproxy.NewProxy(gomitmproxy.Config{
		ListenAddr: &net.TCPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: intPort,
		},
		MITMConfig: mitmConfig,
		OnResponse: func(session *gomitmproxy.Session) *http.Response {
			log.Printf("onResponse: %s", session.Request().URL.String())

			if _, ok := session.GetProp("blocked"); ok {
				log.Printf("onResponse: was blocked")
			}

			res := session.Response()
			req := session.Request()

			if strings.Index(res.Header.Get("Content-Type"), "text/html") != 0 {
				// Do nothing with non-HTML responses
				return nil
			}

			b, err := proxyutil.ReadDecompressedBody(res)
			// Close the original body
			_ = res.Body.Close()
			if err != nil {
				return proxyutil.NewErrorResponse(req, err)
			}

			// Use latin1 before modifying the body
			// Using this 1-byte encoding will let us preserve all original characters
			// regardless of what exactly is the encoding
			body, err := proxyutil.DecodeLatin1(bytes.NewReader(b))
			if err != nil {
				return proxyutil.NewErrorResponse(session.Request(), err)
			}

			// Modifying the original body
			modifiedBody, err := proxyutil.EncodeLatin1(body + "<!-- EDITED -->")
			if err != nil {
				return proxyutil.NewErrorResponse(session.Request(), err)
			}

			res.Body = ioutil.NopCloser(bytes.NewReader(modifiedBody))
			res.Header.Del("Content-Encoding")
			res.ContentLength = int64(len(modifiedBody))
			return res
		},
	})

	proxy.proxy = gomitmproxy
	return &proxy
}

// func main() {

// 	// addCertificateToKeychain()

// 	tlsCert, err := tls.LoadX509KeyPair("example.crt", "example.key")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	privateKey := tlsCert.PrivateKey.(*rsa.PrivateKey)

// 	x509c, err := x509.ParseCertificate(tlsCert.Certificate[0])
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	mitmConfig, err := mitm.NewConfig(x509c, privateKey, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	mitmConfig.SetValidity(time.Hour * 24 * 7) // generate certs valid for 7 days
// 	mitmConfig.SetOrganization("gomitmproxy")  // cert organization

// 	proxy := gomitmproxy.NewProxy(gomitmproxy.Config{
// 		ListenAddr: &net.TCPAddr{
// 			IP:   net.IPv4(0, 0, 0, 0),
// 			Port: 3333,
// 		},
// 		MITMConfig: mitmConfig,
// 		OnResponse: func(session *gomitmproxy.Session) *http.Response {
// 			log.Printf("onResponse: %s", session.Request().URL.String())

// 			if _, ok := session.GetProp("blocked"); ok {
// 				log.Printf("onResponse: was blocked")
// 			}

// 			res := session.Response()
// 			req := session.Request()

// 			if strings.Index(res.Header.Get("Content-Type"), "text/html") != 0 {
// 				// Do nothing with non-HTML responses
// 				return nil
// 			}

// 			b, err := proxyutil.ReadDecompressedBody(res)
// 			// Close the original body
// 			_ = res.Body.Close()
// 			if err != nil {
// 				return proxyutil.NewErrorResponse(req, err)
// 			}

// 			// Use latin1 before modifying the body
// 			// Using this 1-byte encoding will let us preserve all original characters
// 			// regardless of what exactly is the encoding
// 			body, err := proxyutil.DecodeLatin1(bytes.NewReader(b))
// 			if err != nil {
// 				return proxyutil.NewErrorResponse(session.Request(), err)
// 			}

// 			// Modifying the original body
// 			modifiedBody, err := proxyutil.EncodeLatin1(body + "<!-- EDITED -->")
// 			if err != nil {
// 				return proxyutil.NewErrorResponse(session.Request(), err)
// 			}

// 			res.Body = ioutil.NopCloser(bytes.NewReader(modifiedBody))
// 			res.Header.Del("Content-Encoding")
// 			res.ContentLength = int64(len(modifiedBody))
// 			return res
// 		},
// 	})
// 	err = proxy.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	signalChannel := make(chan os.Signal, 1)
// 	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
// 	<-signalChannel

// 	// Clean up
// 	proxy.Close()
// }
