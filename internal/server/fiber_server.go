package server

import (
	"fmt"
	"golang-contact-management-restful-api/config"

	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app *fiber.App
	cfg *config.Config
}

func NewFiberServer(cfg *config.Config) Server {
	return &fiberServer{
		app: fiber.New(),
		cfg: cfg,
	}
}

func (server *fiberServer) Start() error {
	return server.app.Listen(fmt.Sprintf(":%s", server.cfg.Server.Port))
}

func (server *fiberServer) GetEngine() *fiber.App {
	return server.app
}
