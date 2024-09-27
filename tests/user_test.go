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

func TestCreateUser(t *testing.T) {
	router, _ := SetupTestRouter()

	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "User created successfully", response["message"])
}

func TestGetUser(t *testing.T) {
	router, db := SetupTestRouter()

	user := models.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	db.Create(&user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, user.Name, response["user"].Name)
	assert.Equal(t, user.Email, response["user"].Email)
}

func TestGetNonExistentUser(t *testing.T) {
	router, _ := SetupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Failed to get user", response["error"])
}
