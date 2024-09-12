package main

// func setCA(caCert, caKey []byte) error {

// 	certPath, certKey := "example.cert", "example.key"
// 	genCert()

// 	certytes, err := os.ReadFile(certPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	keybytes, err := os.ReadFile(certKey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	goproxyCa, err := tls.X509KeyPair(certytes, keybytes)
// 	if err != nil {
// 		return err
// 	}
// 	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
// 		return err
// 	}
// 	goproxy.GoproxyCa = goproxyCa
// 	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
// 	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
// 	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
// 	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
// 	return nil
// }

// func main() {
// 	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
// 	addr := flag.String("addr", ":7777", "proxy listen address")
// 	flag.Parse()
// 	setCA(caCert, caKey)
// 	proxy := goproxy.NewProxyHttpServer()
// 	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
// 	proxy.Verbose = *verbose
// 	log.Fatal(http.ListenAndServe(*addr, proxy))
// }
