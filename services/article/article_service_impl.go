package article

import (
	"go-agreenery/entities"
	"go-agreenery/helpers"
	"go-agreenery/repositories/article"
	"strings"
)

type articleService struct {
	repository article.ArticleRepository
}

func NewArticleService(r article.ArticleRepository) *articleService {
	return &articleService{
		repository: r,
	}
}

func (s articleService) GetArticles(filter entities.Filter) ([]entities.Article, entities.Pagination, error) {
	return s.repository.GetArticles(filter)
}

func (s articleService) GetArticle(id string) (entities.Article, error) {
	return s.repository.GetArticle(id)
}

func (s articleService) CreateArticle(article entities.Article) (entities.Article, error) {
	var url string
	if article.ThumbnailFile != nil {
		params := helpers.UploaderParams{
			File: article.ThumbnailFile,
		}

		result, err := helpers.UploadFile(params)
		if err != nil {
			return entities.Article{}, err
		}

		url = result
	}

	article.Thumbnail = url

	result, err := s.repository.CreateArticle(article)
	if err != nil {
		if article.ThumbnailFile != nil {
			var object string
			splittedStr := strings.Split(url, "/")
			object = splittedStr[len(splittedStr)-1]

			if err := helpers.DeleteFile(object); err != nil {
				return entities.Article{}, err
			}
		}

		return entities.Article{}, err
	}

	return result, nil
}

func (s articleService) UpdateArticle(article entities.Article) (entities.Article, error) {
	var url string
	if article.ThumbnailFile != nil {
		articleDb, err := s.repository.GetArticle(article.ID)
		if err != nil {
			return entities.Article{}, err
		}

		var oldObj string
		if articleDb.Thumbnail != "" {
			splittedStr := strings.Split(articleDb.Thumbnail, "/")
			oldObj = splittedStr[len(splittedStr)-1]
		}

		params := helpers.UploaderParams{
			File:         article.ThumbnailFile,
			OldObjectURL: oldObj,
		}

		result, err := helpers.UploadFile(params)
		if err != nil {
			return entities.Article{}, err
		}

		url = result
	}

	article.Thumbnail = url

	result, err := s.repository.UpdateArticle(article)
	if err != nil {
		if article.ThumbnailFile != nil {
			var object string
			splittedStr := strings.Split(url, "/")
			object = splittedStr[len(splittedStr)-1]

			if err := helpers.DeleteFile(object); err != nil {
				return entities.Article{}, err
			}
		}

		return entities.Article{}, err
	}

	return result, nil
}

func (s articleService) DeleteArticle(id string) error {
	media, err := s.repository.DeleteArticle(id)
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
