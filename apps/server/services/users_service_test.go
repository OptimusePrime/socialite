package services

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"socialite/models"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	db := models.InitTestDatabase(t, "26980")

	user := models.GenerateUser()
	err := CreateUser(db, user)
	assert.NoError(t, err, "error should be nil")
}

func TestFindAllUsers(t *testing.T) {
	t.Parallel()
	db := models.InitTestDatabase(t, "26981")

	var users []models.User
	rand.Seed(time.Now().UnixNano())
	numberOfUsers := rand.Intn(10-1) + 1
	for i := 0; i < numberOfUsers; i++ {
		user := models.GenerateUser()
		users = append(users, user)
		err := CreateUser(db, user)
		assert.NoError(t, err, "error should be nil")
	}

	foundUsers, err := FindAllUsers(db)
	assert.NoError(t, err, "error should be nil")
	assert.Len(t, foundUsers, numberOfUsers, "length does not match")

	var foundPasswords []string
	for _, user := range foundUsers {
		foundPasswords = append(foundPasswords, user.Password)
	}

	var createdPasswords []string
	for _, user := range users {
		createdPasswords = append(createdPasswords, user.Password)
	}

	if !assert.ObjectsAreEqualValues(createdPasswords, foundPasswords) {
		t.Error("found passwords should contain all created passwords")
	}
}

func TestFindUserByUUID(t *testing.T) {
	t.Parallel()
	db := models.InitTestDatabase(t, "26982")

	user := models.GenerateUser()
	err := CreateUser(db, user)
	assert.NoError(t, err, "error should be nil")

	foundUsers, err := FindAllUsers(db)
	assert.NoError(t, err, "error should be nil")
	assert.Len(t, foundUsers, 1, "slice length should equal 1")

	id := foundUsers[0].ID

	foundUser, err := FindUserByUUID(db, id)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, id, foundUser.ID, "IDs should equal")
}

func TestUpdateOneUser(t *testing.T) {
	t.Parallel()
	db := models.InitTestDatabase(t, "26983")

	user := models.GenerateUser()
	err := CreateUser(db, user)
	assert.NoError(t, err, "error should be nil")

	foundUsers, err := FindAllUsers(db)
	assert.NoError(t, err, "error should be nil")
	assert.Len(t, foundUsers, 1, "slice length should equal 1")

	id := foundUsers[0].ID

	updatedUser := models.User{
		Name: "Tom Scott",
	}
	err = UpdateOneUser(db, updatedUser, id)
	assert.NoError(t, err, "error should be nil")

	foundUser, err := FindUserByUUID(db, id)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, updatedUser.Name, foundUser.Name, "IDs should equal")
}

func TestDeleteOneUser(t *testing.T) {
	t.Parallel()
	db := models.InitTestDatabase(t, "26984")

	user := models.GenerateUser()
	err := CreateUser(db, user)
	assert.NoError(t, err, "error should be nil")

	foundUsers, err := FindAllUsers(db)
	assert.NoError(t, err, "error should be nil")
	assert.Len(t, foundUsers, 1, "slice length should equal 1")

	id := foundUsers[0].ID

	err = DeleteOneUser(db, id)
	assert.NoError(t, err, "error should be nil")

	foundUsers, err = FindAllUsers(db)
	assert.NoError(t, err, "error should be nil")
	assert.Len(t, foundUsers, 0, "slice length should equal 0")
}
