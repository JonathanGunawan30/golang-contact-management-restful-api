package handler

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	UpdateCurrent(ctx *fiber.Ctx) error
	GetCurrent(ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}
