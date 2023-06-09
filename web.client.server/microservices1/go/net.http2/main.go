package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	// "golang.org/x/net/http2"
)

var client = &http.Client{Transport: &http.Transport{
	DisableKeepAlives: true,
	//MaxIdleConns:      10,
}}

var (
	Domain = os.Getenv("DOMAIN")
	key    = "key.pem"
	cert   = "cert.pem"

	cfsslkey  = "./cfssl/server-key.pem"
	cfsslcert = "./cfssl/server.pem"

	serverCrt = "./certs/server.crt"
	serverKey = "./certs/server.key"
)

func main() {
	http.HandleFunc("/v1/user", Get)
	log.Println("Run Server port 0.0.0.0:8080")
	log.Println("[GET]  /v1/user")

	server := &http.Server{
		//Addr: "0.0.0.0:443",
		Addr: "0.0.0.0:8080",
	}

	//http2.ConfigureServer(server, &http2.Server{})
	//log.Fatal(server.ListenAndServeTLS(cert, key))
	log.Fatal(server.ListenAndServe())
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Engine", "Go/nethttp")
	w.Header().Set("Location", "/v1/user")
	w.Header().Set("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

	body, code, err := AdapterConnect("get", nil)
	if err != nil {
		log.Println("Error Server connect:", err, " code:", code)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(``))
		return
	}
	length := strconv.Itoa(len(body))
	w.Header().Set("Content-Length", length)
	w.WriteHeader(code)
	w.Write(body)
	return
}

func AdapterConnect(method string, bodyPost []byte) (body []byte, code int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()

	// send POST
	var Url string = Domain
	var req = &http.Request{}

	req, err = http.NewRequestWithContext(ctx, "GET", Url, nil)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	code = resp.StatusCode

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	return
}
