package requests

type UserPayload struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=100,containsany=!@#$%^&*()_+{}[]:<>?~"`
}

type UserUpdatePayload struct {
	Name string `json:"name" binding:"min=3,max=100"`
}

