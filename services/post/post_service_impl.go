package post

import (
	"go-agreenery/entities"
	"go-agreenery/helpers"
	"go-agreenery/repositories/post"
	"strings"
)

type postService struct {
	repository post.PostRepository
}

func NewPostService(r post.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s postService) GetPosts(filter entities.Filter) ([]entities.Post, entities.Pagination, error) {
	return s.repository.GetPosts(filter)
}

func (s postService) GetPost(id, userID string) (entities.Post, error) {
	return s.repository.GetPost(id, userID)
}

func (s postService) CreatePost(post entities.Post) (entities.Post, error) {
	var url string
	if post.MediaFile != nil {
		params := helpers.UploaderParams{
			File: post.MediaFile,
		}

		result, err := helpers.UploadFile(params)
		if err != nil {
			return entities.Post{}, err
		}

		url = result
	}

	post.Media = url

	result, err := s.repository.CreatePost(post)
	if err != nil {
		var object string
		splittedStr := strings.Split(url, "/")
		object = splittedStr[len(splittedStr)-1]

		if err := helpers.DeleteFile(object); err != nil {
			return entities.Post{}, err
		}

		return entities.Post{}, err
	}

	return result, nil
}

func (s postService) UpdatePost(post entities.Post, currUserID string) (entities.Post, error) {
	var url string
	if post.MediaFile != nil {
		postDb, err := s.repository.GetPost(post.ID, post.UserID)
		if err != nil {
			return entities.Post{}, err
		}

		var oldObj string
		if postDb.Media != "" {
			splittedStr := strings.Split(postDb.Media, "/")
			oldObj = splittedStr[len(splittedStr)-1]
		}

		params := helpers.UploaderParams{
			File:         post.MediaFile,
			OldObjectURL: oldObj,
		}

		result, err := helpers.UploadFile(params)
		if err != nil {
			return entities.Post{}, err
		}

		url = result
	}

	post.Media = url

	result, err := s.repository.UpdatePost(post, currUserID)
	if err != nil {
		var object string
		splittedStr := strings.Split(url, "/")
		object = splittedStr[len(splittedStr)-1]

		if err := helpers.DeleteFile(object); err != nil {
			return entities.Post{}, err
		}

		return entities.Post{}, err
	}

	return result, nil
}

func (s postService) DeletePost(id, currUserID string, isAdmin bool) error {
	media, err := s.repository.DeletePost(id, currUserID, isAdmin)
	if err != nil {
		return err
	}

	if media != "" {
		var oldObj string
		splittedStr := strings.Split(media, "/")
		oldObj = splittedStr[len(splittedStr)-1]

		if err := helpers.DeleteFile(oldObj); err != nil {
			return err
		}
	}

	return nil
}

func (s postService) LikePost(id, currUserID string) (entities.Post, error) {
	return s.repository.LikePost(id, currUserID)
}

func (s postService) GetPostsCountByCategory() ([]entities.Category, error) {
	return s.repository.GetPostsCountByCategory()
}
