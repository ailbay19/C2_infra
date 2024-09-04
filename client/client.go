package main

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
)

//go:embed ssl/client.crt
var cert []byte

//go:embed ssl/client.key
var key []byte

//go:embed ssl/server.crt
var cacert []byte

func main() {
	clientCert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Fatalf("Client crt and key: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cacert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{clientCert},
				ServerName:   "localhost",
			},
		},
	}

	r, err := client.Get("https://localhost:443/")
	if err != nil {
		log.Fatalf("Client Get: %v", err)
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Read body: %v", err)
	}

	fmt.Printf("%s\n", body)
}
