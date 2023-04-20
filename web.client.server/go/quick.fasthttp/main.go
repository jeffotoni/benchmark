package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/jeffotoni/quick"
	"github.com/valyala/fasthttp"
)

var Domain string = os.Getenv("DOMAIN")

var fclient = &fasthttp.Client{
	MaxIdleConnDuration: 3 * time.Second,
	//MaxConnDuration:     90 * time.Second,
}

func init() {
	if len(Domain) == 0 {
		Domain = "http://127.0.0.1:3000/v1/customer/get"
	}
}

func main() {
	q := quick.New(quick.Config{
		BodyLimit:      2 * 1024 * 1024,
		MaxBodySize:    2 * 1024 * 1024,
		MaxHeaderBytes: 1 * 1024 * 1024,
		RouteCapacity:  1000,
		MoreRequests:   0, // valor de equilibrio
	})

	q.Get("/v1/client/get", Get)
	log.Println("Run Server port 0.0.0.0:8080")
	log.Println("[GET]  /v1/client/get")
	log.Println("[POST] /v1/client/post")
	q.Listen("0.0.0.0:8080")
}

func Get(c *quick.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	c.Set("Engine", "Go")
	c.Set("Location", "/v1/client/get")
	c.Set("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

	body, code, err := AdapterConnectFast(Domain, "GET", nil)
	if err != nil {
		log.Println("Error Server connect:", err, " code:", code)
		return c.Status(quick.StatusInternalServerError).SendString("")
	}
	length := strconv.Itoa(len(body))
	c.Set("Content-Length", length)
	return c.Status(code).Byte(body)
}

func AdapterConnectFast(url, method string, bodyPost []byte) (body []byte, code int, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	// if strings.ToUpper(method) == "GET" {
	// 	url = Concat(url, "/get")
	// }

	req.SetRequestURI(url)
	req.Header.SetMethod(method)

	err = fclient.Do(req, resp)
	if err != nil {
		fmt.Printf("Erro ao fazer a requisição: %s\n", err)
		return
	}

	statusCode := resp.StatusCode()
	if statusCode != fasthttp.StatusOK {
		fmt.Printf("Resposta inesperada. Status: %d\n", statusCode)
		return
	}

	body = resp.Body()
	code = statusCode
	return
}

func Concat(strs ...string) string {
	var sb strings.Builder
	for i := 0; i < len(strs); i++ {
		sb.WriteString(strs[i])
	}
	return sb.String()
}
