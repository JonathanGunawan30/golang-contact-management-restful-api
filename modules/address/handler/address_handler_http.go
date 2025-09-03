package handler

import (
	"errors"
	"golang-contact-management-restful-api/internal/transport/http"
	"golang-contact-management-restful-api/modules/address/domain"
	"golang-contact-management-restful-api/modules/address/models"
	"golang-contact-management-restful-api/modules/address/usecase"
	domain2 "golang-contact-management-restful-api/modules/contact/domain"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type addressHandlerHttp struct {
	app      *fiber.App
	usecase  usecase.AddressUsecase
	validate *validator.Validate
}

func NewAddressHttpHandler(app *fiber.App, usecase usecase.AddressUsecase) AddressHandler {
	return &addressHandlerHttp{
		app:      app,
		usecase:  usecase,
		validate: validator.New(),
	}
}

func (handler *addressHandlerHttp) Create(ctx *fiber.Ctx) error {
	var request models.AddressCreateRequest
	if !http.ParseBody(ctx, &request) {
		return nil
	}

	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	contactID, err := strconv.Atoi(ctx.Params("contactId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid ID"})
	}

	response, err := handler.usecase.Create(ctx.Context(), username, contactID, request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(http.DataEnvelope[models.AddressResponse]{
		Data: response,
	})
}

func (handler *addressHandlerHttp) UpdateByID(ctx *fiber.Ctx) error {
	var request models.AddressUpdateRequest
	if !http.ParseBody(ctx, &request) {
		return nil
	}

	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	contactID, err := strconv.Atoi(ctx.Params("contactId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid Contact ID"})
	}

	addressID, err := strconv.Atoi(ctx.Params("addressId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid Address ID"})
	}

	response, err := handler.usecase.Update(ctx.Context(), username, contactID, addressID, request)
	if err != nil {
		if errors.Is(err, domain.ErrAddressNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(http.ErrorResponse{
				Errors: err.Error(),
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[models.AddressResponse]{
		Data: response,
	})
}

func (handler *addressHandlerHttp) FindByID(ctx *fiber.Ctx) error {
	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	contactID, err := strconv.Atoi(ctx.Params("contactId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid Contact ID"})
	}

	addressID, err := strconv.Atoi(ctx.Params("addressId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid Address ID"})
	}

	response, err := handler.usecase.FindByID(ctx.Context(), username, contactID, addressID)
	if err != nil {
		if errors.Is(err, domain.ErrAddressNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(http.ErrorResponse{
				Errors: err.Error(),
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[models.AddressResponse]{
		Data: response,
	})
}

func (handler *addressHandlerHttp) FindAll(ctx *fiber.Ctx) error {
	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	contactID, err := strconv.Atoi(ctx.Params("contactId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid Contact ID"})
	}

	results, err := handler.usecase.FindAll(ctx.Context(), username, contactID)
	if err != nil {
		if errors.Is(err, domain2.ErrContactNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(http.ErrorResponse{
				Errors: err.Error(),
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{
			Errors: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(http.DataEnvelope[[]models.AddressResponse]{
		Data: results,
	})

}

func (handler *addressHandlerHttp) DeleteByID(ctx *fiber.Ctx) error {
	username, err := http.MustGetUsername(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(http.ErrorResponse{Errors: err.Error()})
	}

	contactID, err := strconv.Atoi(ctx.Params("contactId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid Contact ID"})
	}

	addressID, err := strconv.Atoi(ctx.Params("addressId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(http.ErrorResponse{Errors: "Invalid Address ID"})
	}

	err = handler.usecase.DeleteByID(ctx.Context(), username, contactID, addressID)
	if err != nil {
		if errors.Is(err, domain.ErrAddressNotFound) {
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
