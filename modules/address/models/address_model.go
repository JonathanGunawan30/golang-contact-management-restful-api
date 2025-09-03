package models

type AddressCreateRequest struct {
	Street     string `json:"street" validate:"omitempty,min=3,max=255"`
	City       string `json:"city" validate:"omitempty,min=3,max=100"`
	Province   string `json:"province" validate:"omitempty,min=3,max=100"`
	Country    string `json:"country" validate:"required,min=3,max=100"`
	PostalCode string `json:"postal_code" validate:"required,min=3,max=10"`
}

type AddressUpdateRequest struct {
	Street     string `json:"street" validate:"omitempty,min=3,max=255"`
	City       string `json:"city" validate:"omitempty,min=3,max=100"`
	Province   string `json:"province" validate:"omitempty,min=3,max=100"`
	Country    string `json:"country" validate:"omitempty,min=3,max=100"`
	PostalCode string `json:"postal_code" validate:"omitempty,min=3,max=10"`
}

type AddressResponse struct {
	ID         int    `json:"id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}
