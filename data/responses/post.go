package responses

import "time"

type PostResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	UserID uint `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PostsResponse struct {
	Posts []PostResponse `json:"posts"`
}

type PostUserResponse struct {
	PostResponse
	User UserResponse `json:"user"`
}

