package services

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"socialite/dto"
	"testing"
	"time"
)

type PasswordGeneratorParams struct {
	length  uint32
	lower   bool
	upper   bool
	numeric bool
	special bool
	space   bool
}

type ValidateUserParams struct {
	usernameLength uint32
	emailLength    uint32
	validEmail     bool
	password       PasswordGeneratorParams
	genderLength   uint32
	nameLength     uint32
}

func testValidateUser(t *testing.T, params ValidateUserParams, expectedError error) {
	f := faker.New()
	email := f.Lorem().Text(int(params.emailLength))
	if params.validEmail {
		email += "@example.com"
	}
	u := dto.CreateUserDTO{
		Username:  f.Lorem().Text(int(params.usernameLength)),
		Email:     email,
		Password:  gofakeit.Password(params.password.lower, params.password.upper, params.password.numeric, params.password.special, params.password.space, int(params.password.length)),
		Gender:    f.Lorem().Text(int(params.genderLength)),
		Name:      f.Lorem().Text(int(params.nameLength)),
		BirthDate: time.Now(),
	}
	err := validateUser(u)
	if expectedError != nil {
		assert.EqualError(t, expectedError, err.Error(), "\""+err.Error()+"\" error should be "+"\""+expectedError.Error()+"\"")
	}
	assert.Error(t, err, "validate user error should be nil")
}

func TestValidateUser(t *testing.T) {
	t.Run("validate user should fail", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   69,
			genderLength: 69,
			password: PasswordGeneratorParams{
				length:  69,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    69,
			usernameLength: 69,
		}, nil)
	})

	t.Run("validate user should fail for username", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 69,
		}, ErrInvalidUsername)
	})

	t.Run("validate user should fail for email 'cuz of invalid length", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    69,
			usernameLength: 8,
		}, ErrInvalidEmail)
	})

	t.Run("validate user should fail for email 'cuz of invalid format", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   false,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 8,
		}, ErrInvalidEmail)
	})

	t.Run("validate user should fail for password 'cuz of upper case", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   false,
				lower:   true,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 8,
		}, ErrInvalidPassword)
	})

	t.Run("validate user should fail for password 'cuz of lower case", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   true,
				lower:   false,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 8,
		}, ErrInvalidPassword)
	})

	t.Run("validate user should fail for password 'cuz of special chars", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: false,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 8,
		}, ErrInvalidPassword)
	})

	t.Run("validate user should fail for password 'cuz of numbers", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: false,
			},
			emailLength:    8,
			usernameLength: 8,
		}, ErrInvalidPassword)
	})

	t.Run("validate user should fail for password 'cuz of length", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  69,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 8,
		}, ErrInvalidPassword)
	})

	t.Run("validate user should fail for gender", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 69,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 8,
		}, ErrInvalidGender)
	})

	t.Run("validate user should fail for name", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   69,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 8,
		}, ErrInvalidName)
	})

	t.Run("validate user should succeed", func(t *testing.T) {
		testValidateUser(t, ValidateUserParams{
			validEmail:   true,
			nameLength:   8,
			genderLength: 8,
			password: PasswordGeneratorParams{
				length:  8,
				space:   false,
				special: true,
				upper:   true,
				lower:   true,
				numeric: true,
			},
			emailLength:    8,
			usernameLength: 8,
		}, nil)
	})
}

/*import (
	"github.com/stretchr/testify/assert"
	"testing"
)*/

/*import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"math/rand"
	"socialite/ent"
	"socialite/models"
	"testing"
	"time"
)

func beforeEach(t *testing.T, port string) *gorm.DB {
	t.Parallel()
	return ent.InitTestDatabase(t, port)
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
		}
}
*/

/*func TestFindUserByUUID(t *testing.T) {
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
*/
