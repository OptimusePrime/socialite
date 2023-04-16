package services

import (
	"context"
	"github.com/google/uuid"
	"socialite/dto"
	"socialite/ent"
	"socialite/ent/favourite"
	"socialite/ent/post"
	"socialite/ent/user"
)

func CreateFavourite(db *ent.Client, createFavouriteDto dto.CreateFavouriteDTO, ctx context.Context) (err error) {
	_, err = db.Post.Get(ctx, createFavouriteDto.PostID)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrPostNotFound
		}
		return err
	}

	_, err = db.User.Get(ctx, createFavouriteDto.UserID)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrUserNotFound
		}
		return err
	}

	_, err = db.Favourite.Query().Where(favourite.HasPostWith(post.IDEQ(createFavouriteDto.PostID))).Where(favourite.HasUserWith(user.IDEQ(createFavouriteDto.UserID))).First(ctx)
	if err == nil {
		return ErrFavouriteAlreadyExists
	}

	err = db.Favourite.Create().
		SetID(uuid.New()).
		SetPostID(createFavouriteDto.PostID).
		SetUserID(createFavouriteDto.UserID).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func DeleteFavourite(db *ent.Client, deleteFavouriteDto dto.DeleteFavouriteDTO, ctx context.Context) (err error) {
	_, err = db.Favourite.Delete().Where(favourite.HasUserWith(user.IDEQ(deleteFavouriteDto.UserID)), favourite.HasPostWith(post.IDEQ(deleteFavouriteDto.PostID))).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrFavouriteNotFound
		}
		return err
	}

	return nil
}

func FindFavouriteByUserId(db *ent.Client, userId uuid.UUID, ctx context.Context) (err error, favourites []*dto.FavouriteDTO) {
	retrievedFavourites, err := db.Favourite.Query().Where(favourite.HasUserWith(user.IDEQ(userId))).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrUserNotFound, nil
		}

		return err, nil
	}

	for _, retrievedFavourite := range retrievedFavourites {
		favourite := new(dto.FavouriteDTO)
		favourite.ID = retrievedFavourite.ID
		favourite.Post = *retrievedFavourite.QueryPost().FirstX(context.Background())
		favourite.User = *retrievedFavourite.QueryUser().FirstX(context.Background())
		favourites = append(favourites, favourite)
	}

	return nil, favourites
}
