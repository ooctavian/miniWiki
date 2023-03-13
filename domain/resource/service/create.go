package service

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"miniWiki/domain/resource/model"
	"miniWiki/domain/resource/query"
	"miniWiki/utils"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (s *Resource) CreateResource(ctx context.Context, request model.CreateResourceRequest) error {
	err := s.validateCategoryOwner(ctx, request.Resource.CategoryId, request.AccountId)
	if err != nil {
		return err
	}

	params := query.InsertResourceParams{
		Title:       request.Resource.Title,
		Description: request.Resource.Description,
		Link:        request.Resource.Link,
		CategoryID:  request.Resource.CategoryId,
		AuthorID:    request.AccountId,
		State:       query.ResourceState(strings.ToUpper(request.Resource.State)),
	}

	_, err = s.resourceQuerier.InsertResource(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Failed inserting in database: %v", err)
	}

	return err
}

func (s *Resource) validateCategoryOwner(ctx context.Context, categoryId int, accountId int) error {
	category, err := s.categoryQuerier.GetCategoryByID(ctx, categoryId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithContext(ctx).Errorf("Category not found: %v", err)
			return utils.NotFoundError{
				Item: "category",
				Id:   strconv.Itoa(categoryId),
			}
			return err
		}
		logrus.WithContext(ctx).Errorf("Invalid category: %v", err)
		return err
	}

	if *category.AuthorID != accountId {
		logrus.WithContext(ctx).Infof("User %d is not the owner of category %d", accountId, category.CategoryID)
		return utils.ForbiddenError{}
	}
	return nil
}
