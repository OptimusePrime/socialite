package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
	"net/mail"
	"socialite/dto"
	"socialite/ent"
	"socialite/ent/user"
	"time"
)

/*func validateEmail(email string, length uint8) (isValid bool) {
	if uint8(len(email)) > length {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validatePassword(password string, length uint8) (isValid bool) {
	if uint8(len(password)) > length {
		return false
	}
}*/

func validateUser(u dto.CreateUserDTO) (err error) {
	fmt.Println(u)
	if len(u.Name) < 3 || len(u.Name) > 16 {
		return ErrInvalidName
	}
	if len(u.Username) < 3 || len(u.Username) > 16 {
		return ErrInvalidUsername
	}
	if _, err = mail.ParseAddress(u.Email); len(u.Email) > 48 || err != nil {
		return ErrInvalidEmail
	}
	if len(u.Gender) > 16 {
		return ErrInvalidGender
	}
	if len(u.Password) < 8 && len(u.Password) > 32 {
		return ErrInvalidPassword
	}

	/*	validate = validator.New()
		err := validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
			password := fl.Field().String()
			_, err := regexp.Match("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]$", []byte(password))
			if err != nil {
				return false
			}
			return true
		})
		if err != nil {
			return false
		}

		err = validate.Struct(u)
		if err != nil {
			return false
		}*/
	return nil
}

func CreateUser(db *ent.Client, meili *meilisearch.Client, createUserDto dto.CreateUserDTO, ctx context.Context) (err error) {
	_, err = db.User.Query().Where(user.Username(createUserDto.Username)).First(ctx)
	if !ent.IsNotFound(err) {
		return ErrUsernameNotUnique
	}

	if err := validateUser(createUserDto); err != nil {
		return err
	}

	_, err = db.User.Query().Where(user.Email(createUserDto.Email)).First(ctx)
	if !ent.IsNotFound(err) {
		return ErrEmailNotUnique
	}

	params := Argon2Parameters{
		memory:      6 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	hashedPassword, err := HashPassword(createUserDto.Password, params)
	if err != nil {
		return ErrFailedHashingPassword
	}

	userID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	userDocument := dto.CreateUserDocumentDTO{
		ID:       userID,
		Name:     createUserDto.Name,
		Username: createUserDto.Username,
	}

	meiliChan := make(chan error)
	dbChan := make(chan error)

	go func(c chan error) {
		_, err := meili.Index("users").AddDocuments(userDocument)
		c <- err
	}(meiliChan)

	go func(c chan error) {
		err := db.User.
			Create().
			SetID(userID).
			SetEmail(createUserDto.Email).
			SetUsername(createUserDto.Username).
			SetName(createUserDto.Name).
			SetPassword(hashedPassword).
			// SetGender(createUserDto.Gender).
			SetBirthDate(time.Now()).
			// SetAvatar(" ").
			// SetBiography("").
			Exec(ctx)
		c <- err
	}(dbChan)

	meiliErr := <-meiliChan
	dbErr := <-dbChan
	if meiliErr != nil {
		return meiliErr
	}
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func FindUserByUUID(db *ent.Client, uuid uuid.UUID, ctx context.Context) (user *ent.User, err error) {
	return db.User.Get(ctx, uuid)
}

func FindAllUsers(db *ent.Client, ctx context.Context) (users []*ent.User, err error) {
	return db.User.Query().All(ctx)
}

func DeleteOneUser(db *ent.Client /*, uuid uuid.UUID*/, meili *meilisearch.Client, ctx context.Context, accessToken string) (err error) {
	err, isValid, userID := ValidateJWTAccessToken(accessToken)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrInvalidAccessToken
	}

	_, err = meili.Index("users").DeleteDocument(userID.String())
	if err != nil {
		return err
	}

	return db.User.DeleteOneID(*userID).Exec(ctx)
}

func UpdateUser(db *ent.Client, meili *meilisearch.Client, ctx context.Context, updateUserDto dto.UpdateUserDTO) (err error) {
	err, isValid, userId := ValidateJWTAccessToken(updateUserDto.AccessToken)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrInvalidAccessToken
	}

	err = db.User.UpdateOneID(*userId).
		SetName(updateUserDto.Name).
		SetUsername(updateUserDto.Username).
		SetBiography(updateUserDto.Biography).
		SetGender(updateUserDto.Gender).
		SetPronouns(updateUserDto.Pronouns).
		Exec(ctx)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

/*func FindUserByEmail(db *ent.Client, email string, ctx context.Context) (user *ent.User, err error) {

}*/

/*func UpdateOneUser(db *ent.Client, update ent.User, uuid uuid.UUID, ctx context.Context) error {
	return db.User.UpdateOneID().Set.Save(ctx)
}*/
