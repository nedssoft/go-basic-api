package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nedssoft/go-basic-api/models"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	router, _ := SetupTestRouter()

	post := models.Post{Title: "Test Post", Body: "This is a test post", UserID: 1}
	jsonValue, _ := json.Marshal(post)
	req, _ := http.NewRequest("POST", "/api/v1/posts", bytes.NewBufferString(string(jsonValue)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]models.Post
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Test Post", response["post"].Title)
	assert.Equal(t, "This is a test post", response["post"].Body)
	assert.Equal(t, uint(1), response["post"].UserID)
}

func TestGetPost(t *testing.T) {
	router, db := SetupTestRouter()

	// Create a test post
	testPost := models.Post{Title: "Test Post", Body: "This is a test post", UserID: 1}
	db.Create(&testPost)

	req, _ := http.NewRequest("GET", "/api/v1/posts/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]models.Post
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Test Post", response["post"].Title)
	assert.Equal(t, "This is a test post", response["post"].Body)
	assert.Equal(t, uint(1), response["post"].UserID)
}

func TestGetPosts(t *testing.T) {
	router, db := SetupTestRouter()

	// Create test posts
	testPosts := []models.Post{
		{Title: "Test Post 1", Body: "This is test post 1", UserID: 1},
		{Title: "Test Post 2", Body: "This is test post 2", UserID: 1},
	}
	db.Create(&testPosts)

	req, _ := http.NewRequest("GET", "/api/v1/posts", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]models.Post
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response["posts"], 2)
	assert.Equal(t, "Test Post 1", response["posts"][0].Title)
	assert.Equal(t, "Test Post 2", response["posts"][1].Title)
}
