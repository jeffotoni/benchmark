package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var Domain string = os.Getenv("DOMAIN")

var fclient = &fasthttp.Client{
	//MaxConnsPerHost: 10, // Ajuste de acordo com suas necessidades
	MaxIdleConnDuration: 3 * time.Second,
}

func init() {
	if len(Domain) == 0 {
		Domain = "http://127.0.0.1:3000/v1/avatar"
	}
}
func main() {
	app := fiber.New()
	app.Get("/v1/user", Get)
	log.Println("Run Server port 0.0.0.0:8080")
	log.Println("[GET]  /v1/user")
	app.Listen("0.0.0.0:8080")
}

func Get(c *fiber.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	c.Set("Engine", "Go/Fiber-fasthttp")
	c.Set("Location", "/v1/user")
	c.Set("Date", time.Now().Format("2006-01-02T15:04:05.000Z"))

	body, code, err := AdapterConnectFast(Domain, "GET", nil)
	if err != nil {
		log.Println("Error Server connect:", err, " code:", code)
		return c.Status(fiber.StatusInternalServerError).SendString("")
	}
	length := strconv.Itoa(len(body))
	c.Set("Content-Length", length)
	return c.Status(code).Send(body)
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
