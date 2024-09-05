package utils

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed ssl/client.crt
var cert []byte

//go:embed ssl/client.key
var key []byte

//go:embed ssl/server.crt
var cacert []byte

var RootURL string = "http://localhost:18080/"

func BuildURL(segments ...string) string {
	// Join the segments with slashes
	path := strings.Join(segments, "/")

	return RootURL + path
}

func CreateClient() *http.Client {
	// clientCert, err := tls.X509KeyPair(cert, key)
	// if err != nil {
	// 	log.Fatalf("Client crt and key: %v", err)
	// }

	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(cacert)

	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{
	// 			RootCAs:      caCertPool,
	// 			Certificates: []tls.Certificate{clientCert},
	// 			ServerName:   "localhost",
	// 		},
	// 	},
	// }

	return &http.Client{} //client
}

func GetURL(url string) *http.Response {
	client := CreateClient()

	r, err := client.Get(url)
	if err != nil {
		log.Fatalf("Client Get: %v", err)
	}

	// defer r.Body.Close()

	return r
}

func BuildPostRequest(url string, body []byte, headers map[string]string, contentType string) *http.Request {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		log.Fatalf("NewRequest: %v", err)
	}

	req.Header.Set("Content-Type", contentType)

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req
}

func PostRequest(req *http.Request) []byte {
	client := CreateClient()

	r, err := client.Do(req)
	if err != nil {
		log.Fatalf("Client Do: %v", err)
	}

	defer r.Body.Close()
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Read body: %v", err)
	}

	return responseBody
}

func extractFilename(url string) string {
	segments := strings.Split(url, "/")

	return segments[len(segments)-1]
}

func DownloadFrom(url string) {
	response := GetURL(BuildURL(url))

	// Unhandled Exception
	if response.StatusCode != 200 {
		log.Fatalf("Download Code not 200")
	}

	filename := extractFilename(url)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
	defer response.Body.Close()

	// TODO: send acknowledgment?

}
