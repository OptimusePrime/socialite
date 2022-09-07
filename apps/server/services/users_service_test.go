package services

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"socialite/models"
	"testing"
	"time"
)

func beforeEach(t *testing.T, port string) *gorm.DB {
	t.Parallel()
	return models.InitTestDatabase(t, port)
}

func TestCreateUser(t *testing.T) {
	db := beforeEach(t, "26980")

	user := models.GenerateUser()
	err := CreateUser(db, user)
	assert.NoError(t, err, "error should be nil")
}

func TestFindAllUsers(t *testing.T) {
	db := beforeEach(t, "26981")

	var createdUsers []models.User
	rand.Seed(time.Now().UnixNano())
	numberOfUsers := rand.Intn(10-1) + 1
	for i := 0; i < numberOfUsers; i++ {
		user := models.GenerateUser()
		createdUsers = append(createdUsers, user)
		err := CreateUser(db, user)
		assert.NoError(t, err, "error should be nil")
	}

	foundUsers, err := FindAllUsers(db)
	assert.NoError(t, err, "error should be nil")
	assert.Len(t, foundUsers, numberOfUsers, "length does not match")

	numberOfFoundPasswords := 0
	for _, createdUser := range createdUsers {
		for _, foundUser := range foundUsers {
			if foundUser.Password == createdUser.Password {
				numberOfFoundPasswords++
			}
		}
	}
	assert.Equal(t, len(createdUsers), numberOfFoundPasswords, "number of found passwords should equal number of created users")
	/*	var foundPasswords []string
		for _, user := range foundUsers {
			foundPasswords = append(foundPasswords, user.Password)
		}

		var createdPasswords []string
		for _, user := range users {
			createdPasswords = append(createdPasswords, user.Password)
		}



		/*	if !assert.Equal(createdPasswords, foundPasswords) {
			t.Error("found passwords should contain all created passwords")
		}*/
}

func TestFindUserByUUID(t *testing.T) {
	db := beforeEach(t, "26982")

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
	db := beforeEach(t, "26983")

	user := models.GenerateUser()
	err := CreateUser(db, user)
	assert.NoError(t, err, "error should be nil")

	foundUsers, err := FindAllUsers(db)
	assert.NoError(t, err, "error should be nil")
	assert.Len(t, foundUsers, 1, "slice length should equal 1")

	id := foundUsers[0].ID

	updatedUser := models.User{
		Name:     "Tom Scott",
		Password: "test12345",
		Email:    "tom.scott@yahoo.com",
	}
	err = UpdateOneUser(db, updatedUser, id)
	assert.NoError(t, err, "error should be nil")

	foundUser, err := FindUserByUUID(db, id)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, updatedUser.Name, foundUser.Name, "names should equal")
	assert.Equal(t, updatedUser.Password, foundUser.Password, "passwords should equal")
	assert.Equal(t, updatedUser.Email, foundUser.Email, "emails should equal")
}

func TestDeleteOneUser(t *testing.T) {
	db := beforeEach(t, "26984")

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
