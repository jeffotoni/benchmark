package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

var Domain string = os.Getenv("DOMAIN")

var client = &http.Client{Transport: &http.Transport{
	DisableKeepAlives: false,
	MaxIdleConns:      5,
}}

func init() {
	if len(Domain) == 0 {
		Domain = "http://127.0.0.1:3000/v1/avatar"
	}
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

func (g gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/v1/user", Get)
	log.Println("Run Server port 0.0.0.0:8080")
	log.Println("[GET]  /v1/user")
	e.Start("0.0.0.0:8080")
}

func Get(c echo.Context) (err error) {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Set("Engine", "Go")
	c.Response().Header().Set("Location", "/v1/user")
	c.Response().Header().Set("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

	body, code, err := AdapterConnect("get", nil)
	if err != nil {
		log.Println("Error Server connect:", err, " code:", code)
		return c.String(http.StatusInternalServerError, "")
	}
	length := strconv.Itoa(len(body))
	c.Response().Header().Set("Content-Length", length)
	return c.String(code, string(body))
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
