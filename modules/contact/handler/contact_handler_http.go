package handler

import (
	"errors"
	"golang-contact-management-restful-api/internal/transport/http"
	"golang-contact-management-restful-api/modules/contact/domain"
	"golang-contact-management-restful-api/modules/contact/models"
	"golang-contact-management-restful-api/modules/contact/usecase"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type contactHandlerHttp struct {
	app      *fiber.App
	usecase  usecase.ContactUsecase
	validate *validator.Validate
}

func NewContactHttpHandler(app *fiber.App, usecase usecase.ContactUsecase) ContactHandler {
	return &contactHandlerHttp{
		app:      app,
		usecase:  usecase,
		validate: validator.New(),
	}
}

func (handler *contactHandlerHttp) Create(ctx *fiber.Ctx) error {
	var request models.ContactCreateRequest
	if !http.ParseBody(ctx, &request) {
		return nil
	}

	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	response, err := handler.usecase.Create(ctx.Context(), username, request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(http.DataEnvelope[models.ContactResponse]{
		Data: response,
	})

}

func (handler *contactHandlerHttp) UpdateByID(ctx *fiber.Ctx) error {
	var request models.ContactUpdateRequest
	if !http.ParseBody(ctx, &request) {
		return nil
	}

	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid ID"})
	}

	response, err := handler.usecase.Update(ctx.Context(), username, id, request)
	if err != nil {
		if errors.Is(err, domain.ErrContactNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(http.ErrorResponse{Errors: err.Error()})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[models.ContactResponse]{
		Data: response,
	})
}

func (handler *contactHandlerHttp) Search(ctx *fiber.Ctx) error {
	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	size, _ := strconv.Atoi(ctx.Query("size", "10"))

	query := models.ContactSearchQuery{
		Name:  ctx.Query("name"),
		Email: ctx.Query("email"),
		Phone: ctx.Query("phone"),
		Page:  page,
		Size:  size,
	}

	results, paging, err := handler.usecase.Search(ctx.Context(), username, query)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(struct {
		Data   []models.ContactResponse `json:"data"`
		Paging models.Paging            `json:"paging"`
	}{
		Data:   results,
		Paging: paging,
	})
}

func (handler *contactHandlerHttp) GetByID(ctx *fiber.Ctx) error {
	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid ID"})
	}

	response, err := handler.usecase.FindByID(ctx.Context(), username, id)
	if err != nil {
		if errors.Is(err, domain.ErrContactNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(http.ErrorResponse{Errors: err.Error()})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[models.ContactResponse]{
		Data: response,
	})

}

func (handler *contactHandlerHttp) DeleteByID(ctx *fiber.Ctx) error {
	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid ID"})
	}

	err = handler.usecase.DeleteByID(ctx.Context(), username, id)
	if err != nil {
		if errors.Is(err, domain.ErrContactNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(http.ErrorResponse{Errors: err.Error()})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[string]{
		Data: "OK",
	})
}
