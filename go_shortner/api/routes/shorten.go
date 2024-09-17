package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/master-of-none/go_shortner/helpers"
	"time"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse the JSON"})
	}

	//! Rate Limiting
	// Check the IP and check whether it's stored in Database and then decrement the rate by 1

	//! Check the input sent by User is URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}
	//! Check for the Domain Errors
	// Create the function in helper packages.
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Domain Not Found"})
	}
	//! Enforce the https or SSL
	// Create function in helpers
	body.URL = helpers.EnforceHTTP(body.URL)
}
