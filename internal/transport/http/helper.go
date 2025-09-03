package http

import (
	"github.com/gofiber/fiber/v2"
)

func ParseBody[T any](c *fiber.Ctx, dst *T) bool {
	if err := c.BodyParser(dst); err != nil {
		_ = c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Errors: "invalid request body",
		})
		return false
	}
	return true
}

func PointerToString(p *string) string {
	if p != nil {
		return *p
	}
	return ""
}

func StringToPointerIfNotEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func MustGetUsername(c *fiber.Ctx) (string, error) {
	username, ok := c.Locals("username").(string)
	if !ok || username == "" {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return username, nil
}
