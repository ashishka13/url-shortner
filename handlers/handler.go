package handlers

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type HandlerConfig struct {
}

func Controller() {
	app := fiber.New()
	h := HandlerConfig{}
	app.Post("/links/*", h.PostLink(fiber.NewDefaultCtx(app)))

	log.Fatal(app.Listen(":3000"))
}
