package dto

type CreateUserDTO struct {
	Name string `json:"name" validate:"required,min=3"`

	Email string `json:"email" validate:"required,email"`

	Password string `json:"password" validate:"required,min=8"`

	Role string `json:"role"`
}

type UpdateUserDTO struct {
	Name string `json:"name" validate:"required,min=3"`

	Email string `json:"email" validate:"required,email"`
}
