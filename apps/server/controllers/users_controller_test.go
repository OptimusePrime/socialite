package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io"
	"net/http"
	"socialite/models"
	"testing"
)

func beforeEach(t *testing.T, serverPort string, databasePort string) (func(string) string, *gorm.DB) {
	// t.Parallel()
	db := models.InitTestDatabase(t, databasePort)
	getEndpoint := InitTestServer(t, serverPort, db)
	return getEndpoint, db
}

func createOneUserAndRetrieveID(t *testing.T, getEndpoint func(string) string) uuid.UUID {
	// Create user
	user := models.GenerateUser()
	post, err := req.SetBodyJsonMarshal(user).Post(getEndpoint("users"))
	assert.NoError(t, err, "post error should be nil")
	assert.Equal(t, http.StatusCreated, post.GetStatusCode(), "post status code should be 201 created")
	var foundUsers []models.User

	// Find all users
	getAllRes, getAllErr := req.Get(getEndpoint("users"))
	assert.NoError(t, getAllErr, "get error should be nil")
	// Read body
	getAllBytes, readErr := io.ReadAll(getAllRes.Body)
	assert.NoError(t, readErr, "read error should be nil")
	// Parse body as JSON
	unmarshalErr := json.Unmarshal(getAllBytes, &foundUsers)
	assert.NoError(t, unmarshalErr, "unmarshal error should be nil")
	// Test there's only one user and set the ID to that user's ID
	assert.Len(t, foundUsers, 1, "length should equal 1")

	return foundUsers[0].ID
}

func findUserByID(t *testing.T, getEndpoint func(string) string, id uuid.UUID) models.User {
	getOneRes, getOneErr := req.Get(getEndpoint("users/") + id.String())
	assert.NoError(t, getOneErr, "get one error should be nil")
	// Read body
	getOneBytes, readAllErr := io.ReadAll(getOneRes.Body)
	assert.NoError(t, readAllErr, "read all error should be nil")
	// Parse body as JSON and test the IDs match
	var foundUser models.User
	unmarshalErr := json.Unmarshal(getOneBytes, &foundUser)
	assert.NoError(t, unmarshalErr, "unmarshal error should be nil")

	return foundUser
}

func TestCreateUserHandler(t *testing.T) {
	getEndpoint, _ := beforeEach(t, "5001", "6001")

	// Insert a user into a database and test if the status is 201
	user := models.GenerateUser()
	// apt.Post(getEndpoint("users")).JSON(user).Expect(t).Status(http.StatusCreated)
	post, err := req.SetBodyJsonMarshal(user).Post(getEndpoint("users"))
	assert.NoError(t, err, "post error should be nil")
	assert.Equal(t, http.StatusCreated, post.GetStatusCode(), "post status code should be 201 created")
}

func TestFindAllUsersHandler(t *testing.T) {
	getEndpoint, _ := beforeEach(t, "5003", "6003")

	// Create 10 users
	var createdUsers []models.User
	for i := 0; i < 10; i++ {
		user := models.GenerateUser()
		createdUsers = append(createdUsers, user)

		// Insert 10 users into a database and test if the status is 201
		res, err := req.SetBodyJsonMarshal(user).Post(getEndpoint("users"))
		assert.NoError(t, err, "response error should be nil")
		assert.Equal(t, http.StatusCreated, res.GetStatusCode())
	}

	// Get all users
	res, getErr := req.Get(getEndpoint("users"))
	assert.NoError(t, getErr, "response error should be nil")
	// Read the body
	var foundUsers []models.User
	bytes, readErr := io.ReadAll(res.Body)
	assert.NoError(t, readErr, "read error should be nil")
	// Parse body as JSON
	unmarshalErr := json.Unmarshal(bytes, &foundUsers)
	assert.NoError(t, unmarshalErr, "unmarshal error should be nil")
	// Test the length of foundUsers
	assert.Len(t, foundUsers, 10, "foundUsers length should be 10")

	// Test all created passwords are present in foundUsers
	numberOfFoundPasswords := 0
	for _, createdUser := range createdUsers {
		for _, foundUser := range foundUsers {
			if foundUser.Password == createdUser.Password {
				numberOfFoundPasswords++
			}
		}
	}
	assert.Equal(t, len(createdUsers), numberOfFoundPasswords, "number of found passwords should equal number of created users")
}

func TestFindUserByUUIDHandler(t *testing.T) {
	getEndpoint, _ := beforeEach(t, "5002", "6002")

	// Create and insert user, and retrieve the user's ID
	id := createOneUserAndRetrieveID(t, getEndpoint)

	//  Find one user based on ID and test IDs equal
	foundUser := findUserByID(t, getEndpoint, id)
	assert.Equal(t, id, foundUser.ID, "IDs should equal")
}

func TestUpdateOneUserHandler(t *testing.T) {
	getEndpoint, _ := beforeEach(t, "5004", "6004")

	// Create and insert user, and retrieve the user's ID
	id := createOneUserAndRetrieveID(t, getEndpoint)

	updatedUser := models.User{
		Name:     "Tom Scott",
		Email:    "tom.scott@gmail.com",
		Username: "Scotty",
	}

	// Update the user with id
	res, err := req.SetBodyJsonMarshal(updatedUser).Put(getEndpoint("users/" + id.String()))
	assert.NoError(t, err, "put response error should be nil")
	assert.Equal(t, http.StatusNoContent, res.GetStatusCode(), "put response status should be 204")

	// Find one user based on ID and test the user was updated
	foundUser := findUserByID(t, getEndpoint, id)
	assert.Equal(t, updatedUser.Name, foundUser.Name)
	assert.Equal(t, updatedUser.Email, foundUser.Email)
	assert.Equal(t, updatedUser.Username, foundUser.Username)
}

func TestDeleteOneUserHandler(t *testing.T) {
	getEndpoint, _ := beforeEach(t, "5005", "6005")

	// Create and insert user, and retrieve the user's ID
	id := createOneUserAndRetrieveID(t, getEndpoint)

	// Delete user
	deleteRes, deleteErr := req.Delete(getEndpoint("users/" + id.String()))
	assert.NoError(t, deleteErr, "delete response error should be nil")
	assert.Equal(t, http.StatusNoContent, deleteRes.GetStatusCode())
	// Test the user is deleted
	getRes, getErr := req.Get(getEndpoint("users/" + id.String()))
	assert.NoError(t, getErr, "get response error should be nil")
	assert.Equal(t, http.StatusNotFound, getRes.GetStatusCode())
}
