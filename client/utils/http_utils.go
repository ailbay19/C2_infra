package utils

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
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
var resultsPath string = "results"

var client_id string

func SetClientId(id string) {
	client_id = id
	fmt.Printf("Client ID: %s", client_id)
}

func BuildURL(segments ...string) string {
	// Join the segments with slashes
	path := strings.Join(segments, "/")

	return RootURL + path
}

func CreateClient() *http.Client {
	// clientCert, err := tls.X509KeyPair(cert, key)
	// if err != nil {
	// 		return nil
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
	// fmt.Print("HERE")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}

	if client_id != "" {
		req.Header.Set("X-Client-ID", client_id)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil
	}

	return res
}

func BuildPostRequest(url string, body []byte, headers map[string]string, contentType string) *http.Request {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", contentType)

	if client_id != "" {
		req.Header.Set("X-Client-ID", client_id)
	}

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
		return nil
	}

	defer r.Body.Close()
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
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
		return
	}

	filename := extractFilename(url)

	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// TODO: send acknowledgment?

	// Always make executable?
	err = os.Chmod(filename, 0744)
	if err != nil {
		return
	}
}

func SendResults(result []byte) {
	url := BuildURL(resultsPath)

	req := BuildPostRequest(url, result, nil, "application/octet-stream")

	_ = PostRequest(req)
}
