package middleware

import (
	"strings"

	"golang-contact-management-restful-api/modules/user/repository"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(userRepo repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		raw := strings.TrimSpace(c.Get("Authorization"))
		if raw == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
				"errors": "Unauthorized",
			})
		}

		token := raw
		if strings.HasPrefix(strings.ToLower(raw), "bearer ") {
			token = strings.TrimSpace(raw[7:])
		}
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
				"errors": "Unauthorized",
			})
		}

		user, err := userRepo.FindByToken(c.Context(), token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
				"errors": "Unauthorized",
			})
		}

		c.Locals("username", user.Username)

		return c.Next()
	}
}
