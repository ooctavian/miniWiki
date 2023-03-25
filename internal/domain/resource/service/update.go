package service

import (
	"context"
	"strings"

	"miniWiki/internal/domain/resource/model"
	"miniWiki/internal/domain/resource/query"

	"github.com/sirupsen/logrus"
)

func (s *Resource) UpdateResource(ctx context.Context,
	request model.UpdateResourceRequest) (*model.ResourceResponse, error) {
	resource, err := s.getResource(ctx, request.ResourceId, request.AccountId)
	if err != nil {
		return nil, err
	}

	// NOTE: This could be avoided with copier
	params := query.UpdateResourceParams{
		ResourceID:  resource.ResourceID,
		Title:       ptrToString(resource.Title),
		Description: ptrToString(resource.Description),
		Link:        resource.Link,
		CategoryID:  *resource.CategoryID,
		State:       resource.State,
	}

	if request.Resource.CategoryId != nil {
		params.CategoryID = *request.Resource.CategoryId
	}

	if request.Resource.Title != nil {
		params.Title = *request.Resource.Title
	}

	if request.Resource.Description != nil {
		params.Description = *request.Resource.Description
	}

	if request.Resource.Link != nil {
		params.Link = *request.Resource.Link
	}

	if request.Resource.State != nil {
		params.State = query.ResourceState(strings.ToUpper(*request.Resource.State))
	}

	_, err = s.resourceQuerier.UpdateResource(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Infof("Failed updating in database: %v", err)
		return nil, err
	}

	return &model.ResourceResponse{
		ResourceId: resource.ResourceID,
	}, err
}

func ptrToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
