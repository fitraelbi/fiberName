package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/welcome/:name?", func(c *fiber.Ctx) error {
        if c.Params("name") != "" {
            return c.SendString("Selamat datang " + c.Params("name"))
        }
        return c.SendString("Selamat datang Anonymous")
    })

    log.Fatal(app.Listen(":5000"))
}