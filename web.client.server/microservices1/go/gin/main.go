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

	"github.com/gin-gonic/gin"
)

var client = &http.Client{Transport: &http.Transport{
	DisableKeepAlives: false,
	MaxIdleConns:      5,
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

func init() {
	if len(Domain) == 0 {
		Domain = "http://127.0.0.1:3000"
	}
}

func main() {
	//r := gin.Default()
	r := gin.New()

	r.Use(func(c *gin.Context) {})
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/v1/user", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Engine", "Go/Gin")
		c.Header("Location", "/v1/user")
		c.Header("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

		body, code, err := AdapterConnect("get", nil)
		if err != nil {
			log.Println("Error Server connect:", err, " code:", code)
			c.String(http.StatusBadRequest, "Error reading request body: "+err.Error())
			return
		}
		length := strconv.Itoa(len(body))
		c.Header("Content-Length", length)
		c.String(http.StatusOK, string(body))
	})

	log.Fatal(r.Run())
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
