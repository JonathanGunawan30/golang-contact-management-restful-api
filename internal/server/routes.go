package server

import (
	addressHandlerPkg "golang-contact-management-restful-api/modules/address/handler"
	contactHandlerPkg "golang-contact-management-restful-api/modules/contact/handler"
	userHandlerPkg "golang-contact-management-restful-api/modules/user/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, userHandler userHandlerPkg.UserHandler, auth fiber.Handler) {
	api := app.Group("/api")

	api.Post("/users", userHandler.Register)
	api.Post("/users/login", userHandler.Login)

	u := api.Group("/users", auth)
	u.Patch("/current", userHandler.UpdateCurrent)
	u.Get("/current", userHandler.GetCurrent)
	u.Delete("/logout", userHandler.Logout)

}

func RegisterContactRoutes(app *fiber.App, contactHandler contactHandlerPkg.ContactHandler, auth fiber.Handler) {
	api := app.Group("/api", auth)

	api.Post("/contacts", contactHandler.Create)
	api.Get("/contacts", contactHandler.Search)
	api.Get("/contacts/:id", contactHandler.GetByID)
	api.Put("/contacts/:id", contactHandler.UpdateByID)
	api.Delete("/contacts/:id", contactHandler.DeleteByID)
}

func RegisterAddressRoutes(app *fiber.App, addressHandler addressHandlerPkg.AddressHandler, auth fiber.Handler) {
	api := app.Group("/api", auth)

	api.Post("/contacts/:contactId/addresses", addressHandler.Create)
	api.Get("/contacts/:contactId/addresses", addressHandler.FindAll)
	api.Get("/contacts/:contactId/addresses/:addressId", addressHandler.FindByID)
	api.Put("/contacts/:contactId/addresses/:addressId", addressHandler.UpdateByID)
	api.Delete("/contacts/:contactId/addresses/:addressId", addressHandler.DeleteByID)
}
