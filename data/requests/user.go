package requests

type UserPayload struct {
	Name string `json:"name"`
	Email string `json:"email"`
}