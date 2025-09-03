package handler

import (
	"errors"
	"golang-contact-management-restful-api/internal/transport/http"
	"golang-contact-management-restful-api/modules/user/domain"
	"golang-contact-management-restful-api/modules/user/models"
	"golang-contact-management-restful-api/modules/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userHandlerHttp struct {
	app      *fiber.App
	usecase  usecase.UserUsecase
	validate *validator.Validate
}

func NewUserHttpHandler(app *fiber.App, usecase usecase.UserUsecase) UserHandler {
	return &userHandlerHttp{
		app:      app,
		usecase:  usecase,
		validate: validator.New(),
	}
}

func (handler *userHandlerHttp) Register(ctx *fiber.Ctx) error {
	var request models.UserRegisterRequest
	if !http.ParseBody(ctx, &request) {
		return nil
	}

	response, err := handler.usecase.Register(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(http.DataEnvelope[models.UserResponse]{
		Data: response,
	})
}

func (handler *userHandlerHttp) Login(ctx *fiber.Ctx) error {
	var request models.UserLoginRequest
	if !http.ParseBody(ctx, &request) {
		return nil
	}

	response, err := handler.usecase.Login(ctx.Context(), request)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[models.LoginResponse]{
		Data: response,
	})
}

func (handler *userHandlerHttp) UpdateCurrent(ctx *fiber.Ctx) error {
	var request models.UserUpdateRequest
	if !http.ParseBody(ctx, &request) {
		return nil
	}

	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	response, err := handler.usecase.UpdateCurrent(ctx.Context(), username, request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[models.UserResponse]{
		Data: response,
	})

}

func (handler *userHandlerHttp) GetCurrent(ctx *fiber.Ctx) error {
	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	response, err := handler.usecase.GetCurrent(ctx.Context(), username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[models.UserResponse]{
		Data: response,
	})
}

func (handler *userHandlerHttp) Logout(ctx *fiber.Ctx) error {
	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	err = handler.usecase.Logout(ctx.Context(), username)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(http.ErrorResponse{
				Errors: err.Error(),
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[string]{
		Data: "OK",
	})
}
