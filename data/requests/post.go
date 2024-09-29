package requests

type PostPayload struct {
	Title string `json:"title" binding:"required,min=3,max=100"`
	Body string `json:"body" binding:"required,min=10,max=1000"`
	UserID uint `json:"user_id" binding:"required,gt=0"`
}

type PostUpdatePayload struct {
	Title string `json:"title" binding:"min=3,max=100"`
	Body string `json:"body" binding:"min=10,max=1000"`
}

