package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type Link struct {
	URL string `json:"url"`
}

func (h *HandlerConfig) PostLink(c fiber.Ctx) error {
	body := c.Body()

	urlink := Link{}
	err := json.Unmarshal(body, &urlink)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "can't parse submitted URL")
	}

	c.SendString(urlink.URL)
	return nil
}
