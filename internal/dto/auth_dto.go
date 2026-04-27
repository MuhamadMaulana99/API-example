package dto

type RegisterDTO struct {
	Name string `json:"name"`

	Email string `json:"email"`

	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginDTO struct {
	Email string `json:"email"`

	Password string `json:"password"`
}
