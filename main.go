package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func init() {

}

func main() {
	app := fiber.New()

	app.Get("/keys/:key?", func(c *fiber.Ctx) error {
		key := c.Params("key")
		if key == "" || len(key) == 0 {
			return c.SendString("0")
		}
		k := GetKey(key)
		if k == nil {
			return c.SendString("0")
		}
		times := k.Times - k.Used
		if times < 0 {
			times = 0
		}
		return c.SendString(k.Secret + "," + fmt.Sprint(times))
	})

	app.Get("/use", func(c *fiber.Ctx) error {
		key := c.Query("key")
		use := c.Query("use")
		if key == "" {
			return c.SendString("0")
		}
		if times, err := strconv.Atoi(use); err == nil && times > 0 {
			Update(key, times)
			return c.SendString("1")
		}
		return c.SendString("0")
	})

	app.Get("/help", func(c *fiber.Ctx) error {
		return c.SendString("https://zx.wanmei.com/")
	})

	log.Fatal(app.Listen(Conf.Addr))
}
