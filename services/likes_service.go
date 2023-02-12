package services

import (
	"context"
	"github.com/google/uuid"
	"socialite/dto"
	"socialite/ent"
	"socialite/ent/like"
	"socialite/ent/post"
	"socialite/ent/user"
)

func CreateLike(db *ent.Client, createLikeDto dto.LikeDTO, ctx context.Context) (err error) {
	foundUser, err := db.User.Get(ctx, createLikeDto.User)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrUserNotFound
		}
	}

	foundPost, err := db.Post.Get(ctx, createLikeDto.Post)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrPostNotFound
		}
	}

	_, err = db.Like.Query().Where(like.HasPostWith(post.IDEQ(foundPost.ID)), like.HasUserWith(user.IDEQ(foundUser.ID))).First(ctx)
	if !ent.IsNotFound(err) {
		return ErrCannotLikeTwiceSamePost
	}

	err = db.Like.Create().
		SetPost(foundPost).
		SetUser(foundUser).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func DeleteLike(db *ent.Client, deleteLikeDto dto.LikeDTO, ctx context.Context) (err error) {
	_, err = db.Like.Delete().Where(like.HasUserWith(user.IDEQ(deleteLikeDto.User)), like.HasPostWith(post.IDEQ(deleteLikeDto.Post))).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrLikeNotFound
		}
		return err
	}

	return nil
}

func FindLikesByPostId(db *ent.Client, postId uuid.UUID, ctx context.Context) (err error, likes []dto.LikeDTO) {
	retrievedLikes, err := db.Like.Query().Where(like.HasPostWith(post.IDEQ(postId))).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrLikeNotFound, nil
		}
		return err, nil
	}

	for _, retrievedLike := range retrievedLikes {
		likes = append(likes, dto.LikeDTO{
			Post: retrievedLike.QueryPost().FirstIDX(context.Background()),
			User: retrievedLike.QueryUser().FirstIDX(context.Background()),
		})
	}

	return nil, likes
}
