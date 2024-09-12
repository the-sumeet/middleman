package main

import "os/exec"

func genCert() {

	// ToDo: generate cert in pure Go.
	certPath, keyPath := getCertKeyPath()

	exec.Command("openssl", "genrsa", "-out", keyPath, "2048").Run()
	exec.Command("openssl", "pkcs12", "-export ", "-new", "rsa:4096", "-key", keyPath, "-out", certPath, "-days", "3650", "-subj", "/O=Middleman").Run()
}
