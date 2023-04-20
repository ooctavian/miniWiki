package service

import (
	"context"
	"errors"

	model2 "miniWiki/internal/domain/category/model"
	"miniWiki/internal/domain/resource/model"
	"miniWiki/pkg/transport"

	"github.com/jackc/pgconn"
	"github.com/sirupsen/logrus"
)

func (s *Resource) CreateResource(ctx context.Context, request model.CreateResourceRequest) (*uint, error) {
	request.Resource.AuthorId = request.AccountId
	if request.Resource.CategoryName != nil {
		id, err := s.categoryService.CreateCategory(ctx, model2.CreateCategoryRequest{
			Category: model2.CreateCategory{
				Title: *request.Resource.CategoryName,
			},
		})

		if err != nil {
			logrus.WithContext(ctx).
				Errorf("Failed creating category: %v", err)
			return nil, err
		}
		request.Resource.CategoryId = *id
	}

	id, err := s.resourceRepository.InsertResource(ctx, request.Resource)
	if err != nil {
		logrus.WithContext(ctx).
			Errorf("Failed inserting in database: %v", err)
		var pgErr *pgconn.PgError
		// We can assume that duplicated key is thrown only when there is a duplicated link
		// because there is no other unique constraints on table "resource"
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, transport.NewDuplicatedKeyErr("link")
		}

		return nil, err
	}

	return &id, err
}
