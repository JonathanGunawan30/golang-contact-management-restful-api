package handler

import "github.com/gofiber/fiber/v2"

type AddressHandler interface {
	Create(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	DeleteByID(ctx *fiber.Ctx) error
}
