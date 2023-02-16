package services

import (
	"context"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
	"socialite/dto"
	"socialite/ent"
	"socialite/ent/post"
	"socialite/ent/user"
)

func CreatePost(db *ent.Client, createPostDTO *dto.CreatePostDTO, ctx context.Context) (err error, postId uuid.UUID) {
	poster, err := db.User.Get(ctx, createPostDTO.Poster)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrUserNotFound, uuid.UUID{}
		}
		return err, uuid.UUID{}
	}

	postId = uuid.New()
	src, err := createPostDTO.Image.Open()
	if err != nil {
		return err, uuid.UUID{}
	}
	defer src.Close()

	fileURL := "./posts/" + postId.String() + filepath.Ext(createPostDTO.Image.Filename)
	dst, err := os.Create(fileURL)
	// dst, err := os.Create(createPostDTO.Image.Filename)
	if err != nil {
		return err, uuid.UUID{}
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err, uuid.UUID{}
	}

	err = db.Post.Create().
		SetID(postId).
		SetPoster(poster).
		SetCaption(createPostDTO.Caption).
		SetImages([]string{postId.String() + filepath.Ext(createPostDTO.Image.Filename)}).
		SetLocation(createPostDTO.Location).
		Exec(ctx)
	if err != nil {
		return err, uuid.UUID{}
	}
	return nil, postId
}

func DeletePost(db *ent.Client, postID uuid.UUID, ctx context.Context) (err error) {
	foundPost, err := db.Post.Get(ctx, postID)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrPostNotFound
		}
		return err
	}

	for _, image := range foundPost.Images {
		err = os.Remove(image)
		if err != nil {
			return err
		}
	}

	_, err = db.Post.Delete().Where(post.IDEQ(postID)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func entPostToPostDTO(post *ent.Post) *dto.PostDTO {
	return &dto.PostDTO{
		ID:        post.ID.String(),
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Caption:   post.Caption,
		Images:    post.Images,
		Poster:    post.QueryPoster().FirstX(context.Background()),
		Location:  post.Location,
	}
}

func FindPostByID(db *ent.Client, postID uuid.UUID, ctx context.Context) (err error, postDto *dto.PostDTO) {
	post, err := db.Post.Get(ctx, postID)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrPostNotFound, nil
		}
		return err, nil
	}

	return nil, entPostToPostDTO(post)
}

func FindPosts(db *ent.Client, limit int, ctx context.Context) (err error, postsDto []*dto.PostDTO) {
	posts, err := db.Post.Query().Limit(limit).All(ctx)
	if err != nil {
		return err, nil
	}

	for _, entPost := range posts {
		postsDto = append(postsDto, entPostToPostDTO(entPost))
	}

	return nil, postsDto
}

func FindPostsByPosterID(db *ent.Client, posterID uuid.UUID, ctx context.Context) (err error, posts []*dto.PostDTO) {
	retrievedPosts, err := db.Post.Query().Where(post.HasPosterWith(user.IDEQ(posterID))).All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return err, nil
		}
	}

	for _, retrievedPost := range retrievedPosts {
		posts = append(posts, entPostToPostDTO(retrievedPost))
	}

	return nil, posts
}
