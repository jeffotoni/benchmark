package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var Domain string = os.Getenv("DOMAIN")

var client = &http.Client{Transport: &http.Transport{
	DisableKeepAlives: false,
	MaxIdleConns:      5,
}}

func init() {
	if len(Domain) == 0 {
		Domain = "http://127.0.0.1:3000/v1/customer"
	}
}

func main() {

	http.HandleFunc("/v1/client/get", Get)
	http.HandleFunc("/v1/client/post", Post)

	log.Println("Run Server port 0.0.0.0:8080")
	log.Println("[GET]  /v1/client/get")
	log.Println("[POST] /v1/client/post")

	server := &http.Server{
		Addr: "0.0.0.0:8080",
	}
	log.Fatal(server.ListenAndServe())
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Engine", "net/http Go")
	w.Header().Set("Location", "/v1/client/get")
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

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Engine", "Go")
	w.Header().Set("Location", "/v1/client/post")
	w.Header().Set("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error Server ReadAll:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(``))
		return
	}

	body, code, err := AdapterConnect("post", body)
	if err != nil {
		log.Println("Error Server connect:", err, " code:", code)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(``))
		return
	}
	length := strconv.Itoa(len(body))
	w.Header().Set("Content-Length", length)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return
}

func Concat(strs ...string) string {
	var sb strings.Builder
	for i := 0; i < len(strs); i++ {
		sb.WriteString(strs[i])
	}
	return sb.String()
}

func AdapterConnect(method string, bodyPost []byte) (body []byte, code int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()

	var Url string = Domain
	var req = &http.Request{}

	if strings.ToUpper(method) == "GET" {
		Url = Concat(Url, "/get")
		req, err = http.NewRequestWithContext(ctx, "GET", Url, nil)
		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}
	} else if strings.ToUpper(method) == "POST" {
		bodysend := bytes.NewBuffer(bodyPost)
		Url = Concat(Url, "/post")
		req, err = http.NewRequestWithContext(ctx, "POST", Url, bodysend)
		if err != nil {
			fmt.Printf("Error %s", err)
			return
		}
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
