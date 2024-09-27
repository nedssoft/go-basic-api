package requests

type PostPayload struct {
	Title string `json:"title"`
	Body string `json:"body"`
	UserID uint `json:"user_id"`
}