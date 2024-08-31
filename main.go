package main

import (
	"crypto/tls"
	"crypto/x509"
	"embed"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/google/martian"
	"github.com/google/martian/cors"
	"github.com/google/martian/martianhttp"
	"github.com/google/martian/mitm"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type modifier struct{}

func (m modifier) ModifyResponse(res *http.Response) error {
	return nil
}

var (
	addr           = flag.String("addr", ":8080", "host:port of the proxy")
	apiAddr        = flag.String("api-addr", ":8181", "host:port of the configuration API")
	tlsAddr        = flag.String("tls-addr", ":4443", "host:port of the proxy over TLS")
	api            = flag.String("api", "martian.proxy", "hostname for the API")
	generateCA     = flag.Bool("generate-ca-cert", false, "generate CA certificate and private key for MITM")
	cert           = flag.String("cert", "", "filepath to the CA certificate used to sign MITM certificates")
	key            = flag.String("key", "", "filepath to the private key of the CA used to sign MITM certificates")
	organization   = flag.String("organization", "Martian Proxy", "organization name for MITM certificates")
	validity       = flag.Duration("validity", time.Hour, "window of time that MITM certificates are valid")
	allowCORS      = flag.Bool("cors", false, "allow CORS requests to configure the proxy")
	harLogging     = flag.Bool("har", false, "enable HAR logging API")
	marblLogging   = flag.Bool("marbl", false, "enable MARBL logging API")
	trafficShaping = flag.Bool("traffic-shaping", false, "enable traffic shaping API")
	skipTLSVerify  = flag.Bool("skip-tls-verify", false, "skip TLS server verification; insecure")
	dsProxyURL     = flag.String("downstream-proxy-url", "", "URL of downstream proxy")
)

func configure(pattern string, handler http.Handler, mux *http.ServeMux) {
	if *allowCORS {
		handler = cors.NewHandler(handler)
	}

	// register handler for martian.proxy to be forwarded to
	// local API server
	mux.Handle(path.Join(*api, pattern), handler)

	// register handler for local API server
	p := path.Join("localhost"+*apiAddr, pattern)
	mux.Handle(p, handler)
}

func main() {

	p := martian.NewProxy()
	defer p.Close()

	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: *skipTLSVerify,
		},
	}
	p.SetRoundTripper(tr)

	if *dsProxyURL != "" {
		u, err := url.Parse(*dsProxyURL)
		if err != nil {
			log.Fatal(err)
		}
		p.SetDownstreamProxy(u)
	}

	mux := http.NewServeMux()

	var x509c *x509.Certificate
	var priv interface{}

	var err error
	x509c, priv, err = mitm.NewAuthority("martian.proxy", "Martian Authority", 30*24*time.Hour)
	if err != nil {
		log.Fatal(err)
	}

	if x509c != nil && priv != nil {
		mc, err := mitm.NewConfig(x509c, priv)
		if err != nil {
			log.Fatal(err)
		}

		mc.SetValidity(*validity)
		mc.SetOrganization(*organization)
		mc.SkipTLSVerify(*skipTLSVerify)

		p.SetMITM(mc)

		// Expose certificate authority.
		ah := martianhttp.NewAuthorityHandler(x509c)
		configure("/authority.cer", ah, mux)

		// Start TLS listener for transparent MITM.
		tl, err := net.Listen("tcp", *tlsAddr)
		if err != nil {
			log.Fatal(err)
		}

		go p.Serve(tls.NewListener(tl, mc.TLS()))
		fmt.Println("Serving MITM proxy on", *tlsAddr)
	}

	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "middleman",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
