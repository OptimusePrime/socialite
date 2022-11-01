package services

import (
	"context"
	"github.com/google/uuid"
	"socialite/dto"
	"socialite/ent"
	"socialite/ent/follow"
	"socialite/ent/user"
)

func CreateFollow(db *ent.Client, createFollowDto dto.CreateFollowDTO, ctx context.Context) (err error) {
	follower, err := db.User.Get(ctx, createFollowDto.Follower)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrFollowerNotFound
		}
		return err
	}

	followee, err := db.User.Get(ctx, createFollowDto.Followee)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrFolloweeNotFound
		}
		return err
	}

	_, err = db.Follow.Query().Where(follow.HasFollowerWith(user.IDEQ(follower.ID))).Where(follow.HasFolloweeWith(user.IDEQ(followee.ID))).First(ctx)
	if !ent.IsNotFound(err) {
		return ErrCannotFollowTwiceSameUser
	}

	err = db.Follow.Create().
		SetFollower(follower).
		SetFollowee(followee).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFollow(db *ent.Client, deleteFollowDTO dto.DeleteFollowDTO, ctx context.Context) (err error) {
	_, err = db.Follow.Delete().Where(follow.HasFollowerWith(user.IDEQ(deleteFollowDTO.Follower))).Where(follow.HasFolloweeWith(user.IDEQ(deleteFollowDTO.Followee))).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrFollowNotFound
		}
		return err
	}

	return nil
}

func FindFolloweesOfFollower(db *ent.Client, followerId uuid.UUID, ctx context.Context) (followees []*ent.User, err error) {
	_, err = db.User.Query().Where(user.IDEQ(followerId)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrFollowNotFound
		}
		return nil, err
	}

	follows, err := db.Follow.Query().Where(follow.HasFollowerWith(user.IDEQ(followerId))).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrFollowNotFound
		}
		return nil, err
	}
	for _, fol := range follows {
		followee, err := fol.QueryFollowee().First(ctx)
		if err != nil {
			return nil, ErrFolloweeNotFound
		}
		followees = append(followees, followee)
	}

	return followees, nil
}

func FindFollowersOfFollowee(db *ent.Client, followeeId uuid.UUID, ctx context.Context) (followers []*ent.User, err error) {
	_, err = db.User.Query().Where(user.IDEQ(followeeId)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	follows, err := db.Follow.Query().Where(follow.HasFolloweeWith(user.IDEQ(followeeId))).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrFollowNotFound
		}
		return nil, err
	}
	for _, fol := range follows {
		follower, err := fol.QueryFollower().First(ctx)
		if err != nil {
			return nil, ErrFollowerNotFound
		}
		followers = append(followers, follower)
	}

	return followers, nil
}
