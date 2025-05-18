package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"strings"
)

func Proxy(targetHost string, stripPrefix string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		targetPath := strings.TrimPrefix(c.OriginalURL(), stripPrefix)
		targetURL := fmt.Sprintf("%s%s", targetHost, targetPath)

		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(resp)

		req.SetRequestURI(targetURL)
		req.Header.SetMethod(string(c.Method()))
		c.Request().Header.VisitAll(func(k, v []byte) {
			req.Header.SetBytesKV(k, v)
		})

		req.SetBody(c.Body())
		if err := fasthttp.Do(req, resp); err != nil {
			return c.Status(fiber.StatusBadGateway).SendString(err.Error())
		}

		resp.Header.VisitAll(func(k, v []byte) {
			c.Set(string(k), string(v))
		})
		return c.Status(resp.StatusCode()).Send(resp.Body())
	}
}
