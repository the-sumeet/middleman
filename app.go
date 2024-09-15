package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	// "database/sql"

	"github.com/elazarl/goproxy"
	// _ "modernc.org/sqlite"
)

type App struct {
	ctx             context.Context
	proxy           *goproxy.ProxyHttpServer
	proxyStartStoop chan bool
	database        Database
	requests        []http.Request
}

type ReturnValue struct {
	Redirects []Redirect `json:"redirects"`
	// Requests  []http.Request `json:"requests"`
	Error string `json:"error"`
}

func NewApp() *App {

	database := FileDatabase{
		filePath: "database.json",
	}
	database.load()
	genCert()

	proxy := goproxy.NewProxyHttpServer()

	app := &App{
		proxy:           proxy,
		database:        &database,
		proxyStartStoop: make(chan bool),
	}

	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartProxy() {
	//
	certPath, certKey := getCertKeyPath()
	fmt.Println("CertPath: ", certPath)
	fmt.Println("CertKey: ", certKey)
	certytes, err := os.ReadFile(certPath)
	if err != nil {
		panic(err)
	}
	certytes = []byte(`-----BEGIN CERTIFICATE-----
MIICpDCCAYwCCQCk3AX/3JQsGTANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQKDAlt
aWRkbGVtYW4wHhcNMjQwOTE1MDg0NzA0WhcNMjQxMDE1MDg0NzA0WjAUMRIwEAYD
VQQKDAltaWRkbGVtYW4wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCW
dOtbdzacm9puzw0z0pCW5EpEdC7yAUHvZczsCP1Bgs9kqJNgtuGzHfY8lFZwyw/d
UgBCBdasEobnco7HkIqRPIiEZ8HJDKj7KM/E2N+ZX3fKnkOyRflT/AMGBM4rKE+/
2JOctAhJJCcJfYdEHQ+wDJ69WWUjZaXUjzT0swn39h1MgNFBADcj/tb+LwKGNpv4
xqZetrJJFBHz3KE33yGbFXzmiz4keYixQ1StZaxJZVghuN0DN4azIOY7THFR2mFC
ZY1EMA0JXov/RmBlptMDjYlAPulRDxAIMg8W3sXKD1gbtZrCzFLUHROP8Z7KA0nz
ZKv380YToqYI81hdPtD9AgMBAAEwDQYJKoZIhvcNAQELBQADggEBACMcF9oqfCed
3e12ajL7HKj5Egf9SziVT0qLEypQDNFwjtciqF9E0hlvxtstp6q0+5JetnvsUkoB
7erqOEk07GfIvnJi2we4AeG00tZHPYAUViUz4xLqUpRORr4WHUY0tmaZJWSHWRc9
kfyUpiCQLFRfrOP1XeCkJlnB4IphfuLA5lqhzhDuB2tX7zsY1YJ6ewnH2ShWa9W3
z/x4Mk+IwoYW/Sppr6d6mvMShfEzTJi4B+BoZg2pc3bJHFNn3ay4ECQGd0pwaAWq
Qo1Oz4EEllEy4Ei1UTlULc+oH6zjNWHaxi+gosvjtAlILhnf8eM574pZGakjUcqg
HmLNo/G+K10=
-----END CERTIFICATE-----`)

	keybytes, err := os.ReadFile(certKey)
	if err != nil {
		panic(err)
	}
	keybytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAlnTrW3c2nJvabs8NM9KQluRKRHQu8gFB72XM7Aj9QYLPZKiT
YLbhsx32PJRWcMsP3VIAQgXWrBKG53KOx5CKkTyIhGfByQyo+yjPxNjfmV93yp5D
skX5U/wDBgTOKyhPv9iTnLQISSQnCX2HRB0PsAyevVllI2Wl1I809LMJ9/YdTIDR
QQA3I/7W/i8Chjab+MamXraySRQR89yhN98hmxV85os+JHmIsUNUrWWsSWVYIbjd
AzeGsyDmO0xxUdphQmWNRDANCV6L/0ZgZabTA42JQD7pUQ8QCDIPFt7Fyg9YG7Wa
wsxS1B0Tj/GeygNJ82Sr9/NGE6KmCPNYXT7Q/QIDAQABAoIBAAQ7vfOQ4yR87iR4
BvMSVacpPT231ypQBr7pql5p0lZpi9RVOfBatPcJPMhvlA8QZ53elMtGxseIresN
Oh87XerAPUccdENn68lFWLM9Nu1l1kUJNvZpKZ7HVH5y3CoytCu9uEmMdz83mTGX
KpssBij58e15+nKIhaX5cZABCaLEavs6C0h83UL19ogXQeJVKfRTtjUdCpvOlu0I
84Cb1ykXXqAGsIO67NXc5oxKv6z1QGaKCTKt4O5SHrzdVEf+dN0x4UBFFlF0eajv
m+VW36976oor0yu1pKpXSKs6wxLc1LYlrnTEoi8B9D5Sv/oJKlpveCykr5erwnWM
cwM+QIECgYEAxZ1cVllJEzFEiDZoZbNMmoIDsPuXLwfBt3ayCrpv+Vg98yvfhcCt
S08Nprh8KTsp8UwDg16IMUy+Q1hwsix5LVVFxbesenwfDJ4XlZZdyl/eAP/m+uuI
UK0Kv3SXsxn06necVR3xxpcojSh4wY5QccGhqMDcbliNrlhH9xd5lokCgYEAwujA
uvhatyF7WZn93q+BkRVZpvkpjAW8mC5UKW6ILa9VFBYdfBR3MrknLy3t9X24ttRC
b8Gwf5N/XZzhJAeRZWXNQN1D6EXHrCoHzgKLh3qlz5jdk2O/VmdgpVvU6eylUOgR
+B17vAuDWEZLAlSHediojOrkRr6A9UmIUTYLydUCgYEAwbdEcRENZVcCi1RqemeN
TFjvLWs0BNJhv2sHlSS154PtFpeHgiIivpQ8GZb3f6OTtgqB1yGv+EP2ryXfM4oN
L7dLskofNeK+vSusiuLgBiZ8BhbVYlvJQOyggJXWr6deQwoFohq8i6RaCCYIWhUc
HqdlxXtpmnIMUUARK9NDSaECgYEAtTIoKilEPEd2IuBT8Ld8XmJYzC+Kfk+++kLn
nvTQyJfdIiVFF9r2zULvuJ6cP8K2+9DsSnToHlIC8AYuD46xjnBLlmec/8wPSnBw
fQZErJhKmWnlY2YxtKO6Zz+t+iIztblpKx5Nr42Md34xsWLf51iRqR+dRF9KB75q
agUyhFkCgYEAjgONcfLLv7aEj+NhvwvAXosBMY7Njj1zHlXv1exr8PDLvhq1GY0l
oEXyHy4zKCDR81GKjkzMznJb2/mxSFKvanLyDB7WMU82lKJS+goZU73IG5reEBq0
siBcOWcnlKJ9IChFJHP8CsgurRd8yYNQiqhCXkGGKyXDwwj+/PYlitM=
-----END RSA PRIVATE KEY-----`)

	goproxyCa, err := tls.X509KeyPair(certytes, keybytes)
	if err != nil {
		panic(err)
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		panic(err)
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	//

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	go func() {

		log.Println("Proxy Starting")
		a.proxy.Verbose = true
		// a.proxy.OnRequest().DoFunc(
		// 	func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		// 		fmt.Println(r.URL)
		// 		r.Header.Set("X-GoProxy", "yxorPoG-X")
		// 		return r, nil
		// 	})

		a.proxy.OnRequest().DoFunc(a.getOnRequest())

		go func() {
			err := http.Serve(l, a.proxy)
			if err != nil {
				fmt.Println("Proxy Stopped: ", err)
			}
			log.Println("Proxy Started")
		}()

		fmt.Println("Waiting for stop")
		<-a.proxyStartStoop
		fmt.Println("Stopping")
		l.Close()

		log.Println("Proxy Stopped")
	}()
	log.Println("Exit")

}

func (a *App) StopProxy() {
	log.Println("Proxy Stopping")
	a.proxyStartStoop <- true
	log.Println("Proxy Stopped")
}

func (a *App) getOnRequest() func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		requestCopy := *r
		a.requests = append(a.requests, requestCopy)
		r.Header.Set("X-GoProxy", "yxorPoG-X")
		return r, nil
	}
}

func (a *App) getOnResponse() func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	return func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {

		redirects := a.database.GetRedirects()
		for _, redirect := range redirects {
			if redirect.matches(resp) {
				resp.Header.Set("Location", redirect.ToValue)
				resp.StatusCode = 307
			}
		}

		return resp
	}
}

func (a *App) GetRedirects() ReturnValue {
	return ReturnValue{
		Redirects: a.database.GetRedirects(),
	}
}

func (a *App) SaveRedirect(redirectId int, redirect Redirect) {
	a.database.SaveRedirect(redirectId, redirect)
}

func (a *App) AddRedirect(redirect Redirect) {
	a.database.AddRedirect(redirect)
}

func (a *App) RemoveRedirect(redirectId int) {
	a.database.RemoveRedirect(redirectId)
}
