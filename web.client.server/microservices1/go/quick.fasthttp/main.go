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

var (
	M1_DOMAIN string = os.Getenv("M1_DOMAIN")
	M1_PATH   string = os.Getenv("M1_PATH")
	M1_PORT   string = os.Getenv("M1_PORT")

	M2_DOMAIN string = os.Getenv("M2_DOMAIN")
	M2_PATH   string = os.Getenv("M2_PATH")
	M2_PORT   string = os.Getenv("M2_PORT")

	ADDRM1, URLM2 string
)

var fclient = &fasthttp.Client{
	MaxIdleConnDuration: 2 * time.Second,
	//MaxConnDuration:     90 * time.Second,
}

func init() {
	if len(M1_DOMAIN) == 0 {
		M1_DOMAIN = "0.0.0.0"
	}
	if len(M1_PATH) == 0 {
		M1_PATH = "/v1/user"
	}
	if len(M1_PORT) == 0 {
		M1_PORT = "8080"
	}
	ADDRM1 = Concat(M1_DOMAIN, ":", M1_PORT)

	if len(M2_DOMAIN) == 0 {
		M2_DOMAIN = "http://127.0.0.1"
	}
	if len(M2_PATH) == 0 {
		M2_PATH = "/v1/avatar"
	}
	if len(M2_PORT) == 0 {
		M2_PORT = "3000"
	}
	URLM2 = Concat(M2_DOMAIN, ":", M2_PORT, M2_PATH)
}

func main() {
	q := quick.New(quick.Config{
		BodyLimit:      2 * 1024 * 1024,
		MaxBodySize:    2 * 1024 * 1024,
		MaxHeaderBytes: 1 * 1024 * 1024,
		RouteCapacity:  1000,
		MoreRequests:   269, // valor de equilibrio
	})
	q.Get(M1_PATH, Get)
	log.Println("Run Server M1_PORT", ADDRM1)
	log.Println("[GET] ", M1_PATH)
	q.Listen(ADDRM1)
}

func Get(c *quick.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	c.Set("Engine", "Go/Quick-fasthttp")
	c.Set("Location", "/v1/user")
	c.Set("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

	//println(URLM2)
	body, code, err := AdapterConnectFast(URLM2, "GET", nil)
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
