package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
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
	DisableKeepAlives: true,
	//MaxIdleConns:      10,
}}

func init() {
	if len(Domain) == 0 {
		Domain = "http://127.0.0.1:3000/v1/customer"
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

	e.GET("/v1/client/get", Get)
	e.POST("/v1/client/post", Post)

	log.Println("Run Server port 0.0.0.0:8080")
	log.Println("[GET]  /v1/client/get")
	log.Println("[POST] /v1/client/post")
	e.Start("0.0.0.0:8080")
}

func Get(c echo.Context) (err error) {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Set("Engine", "Go")
	c.Response().Header().Set("Location", "/v1/client/get")
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

func Post(c echo.Context) (err error) {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Set("Engine", "Go")
	c.Response().Header().Set("Location", "/v1/client/post")
	c.Response().Header().Set("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

	payload, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Não foi possível ler o payload",
		})
	}

	start := time.Now()
	body, code, err := AdapterConnect("post", payload)
	if err != nil {
		log.Println("Error Server connect:", err, " code:", code)
		return c.String(http.StatusInternalServerError, "")
	}
	end := time.Now().Sub(start)
	log.Println("Service Adapter [POST] timeTotal:", end.String())
	length := strconv.Itoa(len(body))
	c.Response().Header().Set("Content-Length", length)
	return c.String(http.StatusOK, string(body))
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
