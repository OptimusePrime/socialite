package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/meilisearch/meilisearch-go"
	"socialite/dto"
	"socialite/ent"
	"socialite/ent/user"
)

func CreateUser(db *ent.Client, meili *meilisearch.Client, createUserDto dto.CreateUserDTO, ctx context.Context) (err error) {
	_, err = db.User.Query().Where(user.Username(createUserDto.Username)).First(ctx)
	if !ent.IsNotFound(err) {
		return ErrUsernameNotUnique
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
		ID:        userID,
		Name:      createUserDto.Name,
		Username:  createUserDto.Username,
		Biography: createUserDto.Biography,
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
			SetGender(createUserDto.Gender).
			SetBirthDate(createUserDto.BirthDate).
			SetAvatar(createUserDto.Avatar).
			SetBiography(createUserDto.Biography).
			Exec(ctx)
		c <- err
	}(dbChan)

	meiliErr := <-meiliChan
	dbErr := <-dbChan
	if meiliErr != nil || dbErr != nil {
		return ErrFailedCreatingUser
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

/*func FindUserByEmail(db *ent.Client, email string, ctx context.Context) (user *ent.User, err error) {

}*/

/*func UpdateOneUser(db *ent.Client, update ent.User, uuid uuid.UUID, ctx context.Context) error {
	return db.User.UpdateOneID().Set.Save(ctx)
}*/
