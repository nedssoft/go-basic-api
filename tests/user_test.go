package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nedssoft/go-basic-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/nedssoft/go-basic-api/data/responses"
)

const (
	name  = "Test User"
	email = "test@example.com"
	password = "password$123"
)

func TestCreateUser(t *testing.T) {
	router, _ := SetupTestRouter()
	user := models.User{
		Name:  name,
		Email: email,
		Password: password,
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]responses.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, name, response["user"].Name)
	assert.Equal(t, email, response["user"].Email)
}

func TestGetUser(t *testing.T) {
	router, db := SetupTestRouter()
	user := models.User{
		Name:  name,
		Email: email,
		Password: password,
	}
	db.Create(&user)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, name, response["user"].Name)
	assert.Equal(t, email, response["user"].Email)
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

func TestDeleteUser(t *testing.T) {
	router, db := SetupTestRouter()

	// Create a test user
	testUser := models.User{Name: name, Email: email, Password: password}
	db.Create(&testUser)

	req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "User deleted", response["message"])

	// Check if the user was deleted
	var deletedUser models.User
	db.First(&deletedUser, 1)
	assert.Equal(t, uint(0), deletedUser.ID)
}

func TestUpdateUser(t *testing.T) {
	router, db := SetupTestRouter()

	// Create a test user
	testUser := models.User{Name: name, Email: email, Password: password}
	db.Create(&testUser)


	updatedUser := models.User{Name: "Updated User"}
	jsonValue, _ := json.Marshal(updatedUser)
	req, _ := http.NewRequest("PUT", "/api/v1/users/1", bytes.NewBufferString(string(jsonValue)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Updated User", response["user"].Name)
}

