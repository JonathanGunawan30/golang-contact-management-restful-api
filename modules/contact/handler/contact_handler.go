package handler

import "github.com/gofiber/fiber/v2"

type ContactHandler interface {
	Create(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
	Search(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
}
