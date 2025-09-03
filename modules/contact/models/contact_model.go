package models

type ContactCreateRequest struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=100"`
	LastName  string `json:"last_name" validate:"omitempty,min=3,max=100"`
	Email     string `json:"email" validate:"omitempty,email,max=100"`
	Phone     string `json:"phone" validate:"omitempty,min=5,max=20"`
}

type ContactUpdateRequest struct {
	FirstName string `json:"first_name" validate:"omitempty,min=3,max=100"`
	LastName  string `json:"last_name" validate:"omitempty,min=3,max=100"`
	Email     string `json:"email" validate:"omitempty,email,max=100"`
	Phone     string `json:"phone" validate:"omitempty,min=5,max=20"`
}

type ContactResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type ContactSearchQuery struct {
	Name  string
	Email string
	Phone string
	Page  int
	Size  int
}

type Paging struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalItem int `json:"total_item"`
}

type ContactListResponse struct {
	Data   []ContactResponse `json:"data"`
	Paging Paging            `json:"paging"`
}
